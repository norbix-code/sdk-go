package api

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/norbix-code/sdk-go/v2/norbix/internal/transport"
)

// The four workflow (beta) routes an app user calls on the API (gateway
// campaign audit/workflows, docs/architecture/Code.Workflows.md §7, item
// gw-wait-notify) and the module method that calls each one. Each case
// asserts the verb, the fully-resolved path (version substituted, ids
// filled), the query string and the JSON body. The hub base URL is invalid,
// so a route sent to the wrong host fails here.
//
// Nothing leaves the process: every call goes to a local test server.

const (
	testAPIWorkflowRunID = "wfr_4hT2b9QkLm3XcT1yRfGh7a"
	testAPIResumeToken   = "signed.resume.token"
)

type apiCodeCase struct {
	name  string
	verb  string
	path  string
	query url.Values     // nil: no query string expected
	body  map[string]any // nil: no body expected
	call  func(ctx context.Context, m *CodeModule) error
}

func apiCodeCases() []apiCodeCase {
	resume := map[string]any{"stepId": "approve", "decision": "approved", "payload": map[string]any{"note": "ok"}}
	return []apiCodeCase{
		{"GetCodeWorkflowWaitingRuns", http.MethodGet, "/v2/code/workflows/runs",
			url.Values{"assignee": {"me"}, "pageSize": {"20"}, "waitingKind": {"approval"}}, nil,
			func(ctx context.Context, m *CodeModule) error {
				return m.GetCodeWorkflowWaitingRuns(ctx, map[string]any{"assignee": "me", "waitingKind": "approval", "pageSize": 20}, nil)
			}},
		{"ResumeCodeWorkflowRun", http.MethodPost, "/v2/code/workflows/runs/" + testAPIWorkflowRunID + "/resume", nil, resume,
			func(ctx context.Context, m *CodeModule) error {
				return m.ResumeCodeWorkflowRun(ctx, testAPIWorkflowRunID, resume, nil)
			}},
		{"OpenCodeWorkflowResumeLink", http.MethodGet, "/v2/code/workflows/resume/" + testAPIResumeToken, nil, nil,
			func(ctx context.Context, m *CodeModule) error {
				return m.OpenCodeWorkflowResumeLink(ctx, testAPIResumeToken, nil, nil)
			}},
		{"UseCodeWorkflowResumeLink", http.MethodPost, "/v2/code/workflows/resume/" + testAPIResumeToken, nil, nil,
			func(ctx context.Context, m *CodeModule) error {
				return m.UseCodeWorkflowResumeLink(ctx, testAPIResumeToken, nil, nil)
			}},
	}
}

func newAPICodeTestModule(baseURL, apiKey string) *CodeModule {
	return &CodeModule{t: transport.New(&transport.Config{
		ProjectID:  "proj_1",
		APIKey:     apiKey,
		BaseURLAPI: baseURL,
		BaseURLHub: "http://hub.invalid",
		APIVersion: "v2",
		HubVersion: "v2",
		Timeout:    5 * time.Second,
		Env:        "PROD",
	}, nil)}
}

func TestAPICodeWorkflowEndpointsHitTheExpectedRoute(t *testing.T) {
	for _, c := range apiCodeCases() {
		t.Run(c.name, func(t *testing.T) {
			var gotMethod, gotPath, gotRawQuery string
			var gotBody []byte
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotMethod = r.Method
				gotPath = r.URL.Path
				gotRawQuery = r.URL.RawQuery
				gotBody, _ = io.ReadAll(r.Body)
				w.WriteHeader(http.StatusNoContent)
			}))
			defer srv.Close()

			if err := c.call(context.Background(), newAPICodeTestModule(srv.URL, "key_1")); err != nil {
				t.Fatalf("%s: unexpected error: %v", c.name, err)
			}
			if gotMethod != c.verb {
				t.Errorf("verb: got %q want %q", gotMethod, c.verb)
			}
			if gotPath != c.path {
				t.Errorf("path: got %q want %q", gotPath, c.path)
			}
			if c.query == nil && gotRawQuery != "" {
				t.Errorf("query: got %q want none", gotRawQuery)
			}
			if c.query != nil {
				gotQuery, _ := url.ParseQuery(gotRawQuery)
				gq, _ := json.Marshal(gotQuery)
				wq, _ := json.Marshal(c.query)
				if string(gq) != string(wq) {
					t.Errorf("query: got %s want %s", gq, wq)
				}
			}
			if c.body == nil {
				if len(gotBody) != 0 {
					t.Errorf("body: got %s want none", gotBody)
				}
				return
			}
			var got map[string]any
			if err := json.Unmarshal(gotBody, &got); err != nil {
				t.Fatalf("body is not JSON: %v (%s)", err, gotBody)
			}
			gb, _ := json.Marshal(got)
			wb, _ := json.Marshal(c.body)
			if string(gb) != string(wb) {
				t.Errorf("body: got %s want %s", gb, wb)
			}
		})
	}
}

// The signed resume link needs no session: an app with no API key and no
// signed-in user still sends both link routes, with no Authorization header.
func TestAPICodeWorkflowResumeLinkNeedsNoSignIn(t *testing.T) {
	calls := []apiCodeCase{apiCodeCases()[2], apiCodeCases()[3]}
	for _, c := range calls {
		t.Run(c.name, func(t *testing.T) {
			var gotPath string
			var sawAuth bool
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotPath = r.URL.Path
				_, sawAuth = r.Header["Authorization"]
				w.WriteHeader(http.StatusNoContent)
			}))
			defer srv.Close()

			if err := c.call(context.Background(), newAPICodeTestModule(srv.URL, "")); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if gotPath != c.path {
				t.Errorf("path: got %q want %q", gotPath, c.path)
			}
			if sawAuth {
				t.Error("authorization header sent without credentials")
			}
		})
	}
}

// Without a key the project routes still refuse before sending: only the
// link routes are public.
func TestAPICodeWorkflowInboxNeedsSignIn(t *testing.T) {
	m := newAPICodeTestModule("http://127.0.0.1:1", "")
	if err := m.GetCodeWorkflowWaitingRuns(context.Background(), nil, nil); err == nil {
		t.Fatal("expected a not-authenticated error")
	}
}

func TestAPICodeSurfaceSize(t *testing.T) {
	const want = 4
	if got := len(apiCodeCases()); got != want {
		t.Errorf("api code case count: got %d want %d", got, want)
	}
	if got := reflect.TypeOf(&CodeModule{}).NumMethod(); got != want {
		t.Errorf("api CodeModule method count: got %d want %d", got, want)
	}
}
