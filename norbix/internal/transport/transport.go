// Package transport is the shared HTTP layer for the Norbix SDK.
//
// It mirrors the Python/JS transports: it injects auth + scope headers
// (X-CM-ProjectId, X-CM-AccountId, norbix-env, nb-region), builds URLs from a
// templated path, serialises GET/DELETE params as the query string and
// POST/PUT/PATCH params as a JSON body, and retries idempotent requests on
// 429 / 5xx with exponential backoff.
package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	norbixerr "github.com/norbix-code/sdk-go/norbix/errors"
)

// Target selects which base URL a request goes to.
type Target string

const (
	TargetAPI Target = "api"
	TargetHub Target = "hub"
)

// Scope describes the authentication requirement of an endpoint.
type Scope string

const (
	ScopeProject         Scope = "project"
	ScopeAccount         Scope = "account"
	ScopeUnauthenticated Scope = "unauthenticated"
)

const (
	defaultMaxRetries = 3
	baseRetryDelay    = 250 * time.Millisecond
)

var idempotentVerbs = map[string]bool{http.MethodGet: true, http.MethodDelete: true}

// Config holds resolved client configuration. Fields the auth helpers mutate
// (bearer token, region, env, scope) are updated in place by the client.
type Config struct {
	APIKey      string
	BearerToken string
	ProjectID   string
	AccountID   string

	BaseURLAPI string
	BaseURLHub string
	APIVersion string
	HubVersion string

	Timeout time.Duration
	Env     string
	Region  string

	// Whether each base URL is the SDK default (only defaults get a region
	// subdomain rewrite; user-supplied URLs are never touched).
	BaseURLAPIIsDefault bool
	BaseURLHubIsDefault bool

	DefaultHeaders map[string]string
}

// Transport performs requests against the Norbix API and Hub.
type Transport struct {
	Cfg    *Config
	client *http.Client
}

// New builds a Transport. Pass a nil http.Client to use a default one.
func New(cfg *Config, client *http.Client) *Transport {
	if client == nil {
		client = &http.Client{}
	}
	return &Transport{Cfg: cfg, client: client}
}

// Request is a single SDK call. Module methods build one of these.
type Request struct {
	Target     Target
	Path       string
	Method     string
	PathParams map[string]string
	// Body is sent as JSON for write verbs and as the query string for GET/DELETE.
	Body  map[string]any
	Scope Scope

	// Per-call overrides (empty means "use client default").
	Timeout     time.Duration
	BearerToken string
	Env         string
	Region      string
}

// rawResponse is the fully-buffered result of a request attempt.
type rawResponse struct {
	status  int
	body    []byte
	headers http.Header
}

// Send executes req and decodes the JSON response into out (which may be nil
// to discard the body, or a *map[string]any / pointer to a typed struct).
func (t *Transport) Send(ctx context.Context, req Request, out any) error {
	if ctx == nil {
		ctx = context.Background()
	}
	if req.Scope == ScopeAccount && t.Cfg.AccountID == "" {
		return norbixerr.New(
			"This endpoint is account-scoped. Configure AccountID on the client.",
			norbixerr.CodeAccountScope,
		)
	}

	base := t.Cfg.BaseURLAPI
	version := t.Cfg.APIVersion
	if req.Target == TargetHub {
		base = t.Cfg.BaseURLHub
		version = t.Cfg.HubVersion
	}

	urlStr, body, err := buildURLAndBody(base, req.Path, version, req.Method, req.PathParams, req.Body)
	if err != nil {
		return err
	}

	headers := map[string]string{"Accept": "application/json"}
	for k, v := range t.Cfg.DefaultHeaders {
		headers[k] = v
	}

	if req.Scope != ScopeUnauthenticated {
		token := req.BearerToken
		if token == "" {
			token = t.Cfg.BearerToken
		}
		if token == "" {
			token = t.Cfg.APIKey
		}
		if token == "" {
			return norbixerr.New(
				"Not authenticated. Provide APIKey / BearerToken or login first.",
				norbixerr.CodeNotAuthenticated,
			)
		}
		headers["Authorization"] = "Bearer " + token
	}

	headers["X-CM-ProjectId"] = t.Cfg.ProjectID
	if t.Cfg.AccountID != "" {
		headers["X-CM-AccountId"] = t.Cfg.AccountID
	}

	// Environment selector: per-call override wins; "PROD" is the backend
	// default so the header is omitted for it.
	env := req.Env
	if env == "" {
		env = t.Cfg.Env
	}
	if env != "" && env != "PROD" {
		headers["norbix-env"] = env
	}

	// Region selector: per-call override wins; no default region exists, so the
	// header is omitted when unset.
	region := req.Region
	if region == "" {
		region = t.Cfg.Region
	}
	if region != "" {
		headers["nb-region"] = region
	}

	if body != nil {
		headers["Content-Type"] = "application/json"
	}

	timeout := req.Timeout
	if timeout == 0 {
		timeout = t.Cfg.Timeout
	}

	resp, err := t.doWithRetries(ctx, req.Method, urlStr, headers, body, timeout)
	if err != nil {
		return norbixerr.New(err.Error(), norbixerr.CodeNetwork).WithWrapped(err)
	}

	if resp.status >= 400 {
		data := map[string]any{}
		_ = json.Unmarshal(resp.body, &data)
		message := stringField(data, "message")
		if message == "" {
			if len(resp.body) > 0 {
				message = string(resp.body)
			} else {
				message = "Request failed"
			}
		}
		code := stringField(data, "errorCode")
		if code == "" {
			code = fmt.Sprintf("HTTP_%d", resp.status)
		}
		return norbixerr.FromHTTP(message, resp.status, code, data)
	}

	if out == nil || resp.status == http.StatusNoContent || len(resp.body) == 0 {
		return nil
	}
	if err := json.Unmarshal(resp.body, out); err != nil {
		return norbixerr.New("failed to decode response: "+err.Error(), norbixerr.CodeError).WithWrapped(err)
	}
	return nil
}

func (t *Transport) doWithRetries(
	ctx context.Context,
	method, urlStr string,
	headers map[string]string,
	body []byte,
	timeout time.Duration,
) (*rawResponse, error) {
	var last *rawResponse
	for attempt := 0; attempt < defaultMaxRetries; attempt++ {
		raw, err := t.doOnce(ctx, method, urlStr, headers, body, timeout)
		if err != nil {
			// Network/timeout error: retry only idempotent verbs.
			if !idempotentVerbs[method] || attempt >= defaultMaxRetries-1 {
				return nil, err
			}
			if sleepErr := backoff(ctx, attempt); sleepErr != nil {
				return nil, sleepErr
			}
			continue
		}
		last = raw
		if raw.status < 400 {
			return raw, nil
		}
		if raw.status < 500 && raw.status != 429 {
			return raw, nil
		}
		if !idempotentVerbs[method] || attempt >= defaultMaxRetries-1 {
			return raw, nil
		}
		if sleepErr := backoff(ctx, attempt); sleepErr != nil {
			return nil, sleepErr
		}
	}
	return last, nil
}

// doOnce performs a single request and buffers the whole body so the per-attempt
// timeout context can be released safely.
func (t *Transport) doOnce(
	ctx context.Context,
	method, urlStr string,
	headers map[string]string,
	body []byte,
	timeout time.Duration,
) (*rawResponse, error) {
	reqCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var reader io.Reader
	if body != nil {
		reader = bytes.NewReader(body)
	}
	httpReq, err := http.NewRequestWithContext(reqCtx, method, urlStr, reader)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		httpReq.Header.Set(k, v)
	}

	resp, err := t.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return &rawResponse{status: resp.StatusCode, body: buf, headers: resp.Header}, nil
}

func backoff(ctx context.Context, attempt int) error {
	delay := time.Duration(float64(baseRetryDelay)*pow2(attempt)) +
		time.Duration(rand.Int63n(int64(100*time.Millisecond)))
	select {
	case <-time.After(delay):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func pow2(n int) float64 {
	r := 1.0
	for i := 0; i < n; i++ {
		r *= 2
	}
	return r
}

// buildURLAndBody substitutes path params, fills {version}, then either
// query-encodes (GET/DELETE) or JSON-encodes (others) the remaining body.
func buildURLAndBody(
	base, path, version, method string,
	pathParams map[string]string,
	request map[string]any,
) (string, []byte, error) {
	normalized := strings.ReplaceAll(path, "{version}", version)

	for {
		start := strings.Index(normalized, "{")
		if start < 0 {
			break
		}
		rel := strings.Index(normalized[start:], "}")
		if rel < 0 {
			break
		}
		end := start + rel
		token := normalized[start+1 : end]
		val, ok := pathParams[token]
		if !ok {
			return "", nil, norbixerr.New(
				fmt.Sprintf("Missing path parameter '%s' for path %s", token, path),
				norbixerr.CodeMissingPathParam,
			)
		}
		normalized = normalized[:start] + val + normalized[end+1:]
	}

	full := strings.TrimRight(base, "/") + "/" + strings.TrimLeft(normalized, "/")

	remaining := map[string]any{}
	for k, v := range request {
		if v != nil {
			remaining[k] = v
		}
	}

	if method == http.MethodGet || method == http.MethodDelete {
		if len(remaining) == 0 {
			return full, nil, nil
		}
		q := url.Values{}
		for k, v := range remaining {
			for _, item := range toIterable(v) {
				q.Add(k, stringify(item))
			}
		}
		return full + "?" + q.Encode(), nil, nil
	}

	if len(remaining) == 0 {
		return full, nil, nil
	}
	b, err := json.Marshal(remaining)
	if err != nil {
		return "", nil, norbixerr.New("failed to encode request body: "+err.Error(), norbixerr.CodeError)
	}
	return full, b, nil
}

func toIterable(v any) []any {
	if list, ok := v.([]any); ok {
		return list
	}
	return []any{v}
}

func stringify(v any) string {
	switch t := v.(type) {
	case bool:
		if t {
			return "true"
		}
		return "false"
	case string:
		return t
	default:
		return fmt.Sprintf("%v", v)
	}
}

func stringField(m map[string]any, key string) string {
	if v, ok := m[key]; ok {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}
