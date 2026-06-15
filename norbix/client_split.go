package norbix

import (
	"context"
	"net/http"

	"github.com/norbix-code/sdk-go/norbix/api"
	"github.com/norbix-code/sdk-go/norbix/hub"
	"github.com/norbix-code/sdk-go/norbix/internal/transport"
)

// APIClient is an API-only client with flat module access (c.Database, c.Chat, ...).
type APIClient struct {
	auth
	*api.Namespace
	t *transport.Transport
}

// NewAPI builds an API-only client.
func NewAPI(opts Options, httpClient ...*http.Client) (*APIClient, error) {
	cfg, err := buildConfig(opts, "NorbixApi")
	if err != nil {
		return nil, err
	}
	var hc *http.Client
	if len(httpClient) > 0 {
		hc = httpClient[0]
	}
	t := transport.New(cfg, hc)
	return &APIClient{auth: auth{t: t}, Namespace: api.NewNamespace(t), t: t}, nil
}

// Login authenticates a user and stores the returned bearer token.
func (c *APIClient) Login(ctx context.Context, creds LoginCredentials) (map[string]any, error) {
	return c.login(ctx, creds)
}

// Close is a no-op.
func (c *APIClient) Close() error { return nil }

// HubClient is a Hub-only client with flat module access (c.Account, c.Email, ...).
type HubClient struct {
	auth
	*hub.Namespace
	t *transport.Transport
}

// NewHub builds a Hub-only client.
func NewHub(opts Options, httpClient ...*http.Client) (*HubClient, error) {
	cfg, err := buildConfig(opts, "NorbixHub")
	if err != nil {
		return nil, err
	}
	var hc *http.Client
	if len(httpClient) > 0 {
		hc = httpClient[0]
	}
	t := transport.New(cfg, hc)
	return &HubClient{auth: auth{t: t}, Namespace: hub.NewNamespace(t), t: t}, nil
}

// Login authenticates a user and stores the returned bearer token.
func (c *HubClient) Login(ctx context.Context, creds LoginCredentials) (map[string]any, error) {
	return c.login(ctx, creds)
}

// Close is a no-op.
func (c *HubClient) Close() error { return nil }
