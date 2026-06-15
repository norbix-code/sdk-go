// Package norbix is the official Go SDK for Norbix. It exposes a full client
// (Client) with api.* and hub.* namespaces, plus API-only and Hub-only
// variants with flat module access.
//
//	c, _ := norbix.New(norbix.Options{APIKey: "...", ProjectID: "..."})
//	defer c.Close()
//	var schemas map[string]any
//	_ = c.API.Database.GetDatabaseSchemas(ctx, nil, &schemas)
package norbix

import (
	"context"
	"net/http"

	"github.com/norbix-code/sdk-go/norbix/api"
	norbixerr "github.com/norbix-code/sdk-go/norbix/errors"
	"github.com/norbix-code/sdk-go/norbix/hub"
	"github.com/norbix-code/sdk-go/norbix/internal/transport"
)

func errProjectIDRequired(clientName string) error {
	return norbixerr.New(
		clientName+": ProjectID is required (set NORBIX_PROJECT_ID or pass Options.ProjectID).",
		norbixerr.CodeError,
	)
}

// LoginCredentials is the payload for Client.Login.
type LoginCredentials struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	Provider string `json:"provider,omitempty"`
}

// auth bundles the shared mutators every client embeds.
type auth struct {
	t *transport.Transport
}

// SetBearerToken sets the bearer token used for subsequent requests.
func (a *auth) SetBearerToken(token string) { a.t.Cfg.BearerToken = token }

// SetAPIKey sets the API key used for subsequent requests.
func (a *auth) SetAPIKey(key string) { a.t.Cfg.APIKey = key }

// Logout clears the bearer token.
func (a *auth) Logout() { a.t.Cfg.BearerToken = "" }

// SetScope switches the project (and optionally account) for subsequent requests.
func (a *auth) SetScope(projectID, accountID string) {
	a.t.Cfg.ProjectID = projectID
	a.t.Cfg.AccountID = accountID
}

// SetEnv switches the project environment (norbix-env header). Pass "" or "PROD"
// to return to production.
func (a *auth) SetEnv(envName string) {
	if envName == "" {
		envName = "PROD"
	}
	a.t.Cfg.Env = envName
}

// GetEnv returns the current project environment.
func (a *auth) GetEnv() string { return a.t.Cfg.Env }

// SetRegion switches the Norbix region (nb-region header). Pass "" to clear it.
// SDK-default base URLs are recomposed; user-supplied base URLs are untouched.
func (a *auth) SetRegion(region string) {
	a.t.Cfg.Region = region
	if a.t.Cfg.BaseURLAPIIsDefault {
		a.t.Cfg.BaseURLAPI = composeRegionalURL(DefaultBaseURLAPI, region)
	}
	if a.t.Cfg.BaseURLHubIsDefault {
		a.t.Cfg.BaseURLHub = composeRegionalURL(DefaultBaseURLHub, region)
	}
}

// GetRegion returns the current Norbix region ("" when unset).
func (a *auth) GetRegion() string { return a.t.Cfg.Region }

// IsAuthenticated reports whether a bearer token or API key is configured.
func (a *auth) IsAuthenticated() bool {
	return a.t.Cfg.BearerToken != "" || a.t.Cfg.APIKey != ""
}

// login performs the shared /auth POST and stores the returned bearer token.
func (a *auth) login(ctx context.Context, creds LoginCredentials) (map[string]any, error) {
	if creds.Provider == "" {
		creds.Provider = "credentials"
	}
	body := map[string]any{
		"userName": creds.UserName,
		"password": creds.Password,
		"provider": creds.Provider,
	}
	var result map[string]any
	err := a.t.Send(ctx, transport.Request{
		Target: transport.TargetAPI,
		Path:   "/auth",
		Method: http.MethodPost,
		Body:   body,
		Scope:  transport.ScopeUnauthenticated,
	}, &result)
	if err != nil {
		return nil, err
	}
	if token, ok := result["bearerToken"].(string); ok && token != "" {
		a.SetBearerToken(token)
	}
	return result, nil
}

// Client is the full Norbix SDK client with api.* and hub.* namespaces.
type Client struct {
	auth
	API *api.Namespace
	Hub *hub.Namespace
	t   *transport.Transport
}

// New builds a full client. Pass an *http.Client to reuse connections (nil is fine).
func New(opts Options, httpClient ...*http.Client) (*Client, error) {
	cfg, err := buildConfig(opts, "Norbix")
	if err != nil {
		return nil, err
	}
	var hc *http.Client
	if len(httpClient) > 0 {
		hc = httpClient[0]
	}
	t := transport.New(cfg, hc)
	return &Client{
		auth: auth{t: t},
		API:  api.NewNamespace(t),
		Hub:  hub.NewNamespace(t),
		t:    t,
	}, nil
}

// Login authenticates a user and stores the returned bearer token on the client.
func (c *Client) Login(ctx context.Context, creds LoginCredentials) (map[string]any, error) {
	return c.login(ctx, creds)
}

// Close is a no-op kept for symmetry / future connection pooling.
func (c *Client) Close() error { return nil }
