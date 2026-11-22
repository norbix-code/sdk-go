package norbix

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	norbixerr "github.com/norbix-code/sdk-go/v2/norbix/errors"
)

// Routes the gateway does not authenticate — public sign-up, password reset,
// echo, the quick-connect provider list — called through the real module
// methods by a client that has a project id and nothing else. Before the fix
// they were marked project scope, so the SDK answered
// NORBIX_NOT_AUTHENTICATED and sent nothing.
//
// Every case starts its own fake gateway and its own client; nothing leaves
// the process.

const signupProjectID = "proj_1"

// signupSeen is what the fake gateway saw: how many requests, and the last one.
type signupSeen struct {
	hits   int
	method string
	path   string
	header http.Header
}

// newSignupClient starts a fake gateway and returns a client pointed at it.
// Pass an empty apiKey for a visitor who has never signed in. The NORBIX_*
// credentials of the machine running the tests are blanked, so "no key"
// really means no key.
func newSignupClient(t *testing.T, apiKey string) (*Client, *signupSeen) {
	t.Helper()
	t.Setenv("NORBIX_API_KEY", "")
	t.Setenv("NORBIX_BEARER_TOKEN", "")

	seen := &signupSeen{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		seen.hits++
		seen.method = r.Method
		seen.path = r.URL.Path
		seen.header = r.Header.Clone()
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{}`))
	}))
	t.Cleanup(srv.Close)

	c, err := New(Options{
		ProjectID:  signupProjectID,
		APIKey:     apiKey,
		BaseURLAPI: srv.URL,
		BaseURLHub: srv.URL,
		APIVersion: "v3",
		HubVersion: "v2",
	})
	if err != nil {
		t.Fatalf("client: %v", err)
	}
	return c, seen
}

func TestPublicRoutesAreSentWithoutAToken(t *testing.T) {
	cases := []struct {
		name   string
		method string
		path   string
		call   func(context.Context, *Client) error
	}{
		{"register with e-mail", http.MethodPost, "/v3/membership/auth/register/email",
			func(ctx context.Context, c *Client) error {
				return c.API.Membership.SaveEmailUser(ctx, map[string]any{"email": "ana@example.test", "password": "S3cret-pass!"}, nil)
			}},
		{"register with user name", http.MethodPost, "/v3/membership/auth/register/user-name",
			func(ctx context.Context, c *Client) error {
				return c.API.Membership.SaveUserNameUser(ctx, map[string]any{"userName": "ana", "password": "S3cret-pass!"}, nil)
			}},
		{"password reset request", http.MethodPost, "/v3/membership/userauth/password/reset/request",
			func(ctx context.Context, c *Client) error {
				return c.API.Membership.RequestPasswordReset(ctx, map[string]any{"email": "ana@example.test"}, nil)
			}},
		{"password reset confirm", http.MethodPost, "/v3/membership/userauth/password/reset/confirm",
			func(ctx context.Context, c *Client) error {
				return c.API.Membership.ConfirmPasswordReset(ctx, map[string]any{"token": "reset_tok", "password": "S3cret-pass!"}, nil)
			}},
		{"e-mail address verification confirm", http.MethodPost, "/v3/membership/userauth/email-verification/confirm",
			func(ctx context.Context, c *Client) error {
				return c.API.Membership.VerifyEmailAddress(ctx, map[string]any{"token": "evt_1"}, nil)
			}},
		{"e-mail address verification resend", http.MethodPost, "/v3/membership/userauth/email-verification/resend",
			func(ctx context.Context, c *Client) error {
				return c.API.Membership.ResendEmailAddressVerification(ctx, map[string]any{"email": "ana@example.test"}, nil)
			}},
		{"api echo", http.MethodGet, "/v3/echo",
			func(ctx context.Context, c *Client) error { return c.API.Echo.Echo(ctx, nil, nil) }},
		{"hub echo", http.MethodGet, "/v2/echo",
			func(ctx context.Context, c *Client) error { return c.Hub.Echo.Echo(ctx, nil, nil) }},
		{"quick-connect providers", http.MethodGet, "/v2/integrations/connect/providers",
			func(ctx context.Context, c *Client) error {
				return c.Hub.Integrations.GetQuickConnectProviders(ctx, nil, nil)
			}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			c, seen := newSignupClient(t, "")

			if err := tc.call(context.Background(), c); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if seen.hits != 1 {
				t.Fatalf("requests: got %d want 1", seen.hits)
			}
			if seen.method != tc.method {
				t.Errorf("method: got %q want %q", seen.method, tc.method)
			}
			if seen.path != tc.path {
				t.Errorf("path: got %q want %q", seen.path, tc.path)
			}
			if auth, ok := seen.header["Authorization"]; ok {
				t.Errorf("Authorization: must not be sent without a token, got %q", auth)
			}
			if got := seen.header.Get("norbix-project-id"); got != signupProjectID {
				t.Errorf("norbix-project-id: got %q want %q", got, signupProjectID)
			}
		})
	}
}

// A client that does have a key still sends it on a public route: the gateway
// then knows who registered the user.
func TestRegisterWithEmailSendsTheAPIKeyWhenThereIsOne(t *testing.T) {
	c, seen := newSignupClient(t, "key_1")

	err := c.API.Membership.SaveEmailUser(context.Background(),
		map[string]any{"email": "ana@example.test", "password": "S3cret-pass!"}, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if seen.hits != 1 {
		t.Fatalf("requests: got %d want 1", seen.hits)
	}
	if seen.method != http.MethodPost || seen.path != "/v3/membership/auth/register/email" {
		t.Errorf("call: got %s %s", seen.method, seen.path)
	}
	if got := seen.header.Get("Authorization"); got != "Bearer key_1" {
		t.Errorf("Authorization: got %q want %q", got, "Bearer key_1")
	}
	if got := seen.header.Get("norbix-project-id"); got != signupProjectID {
		t.Errorf("norbix-project-id: got %q want %q", got, signupProjectID)
	}
}

// A route that needs sign-in is unchanged: no token, no request.
func TestBlockUserStillNeedsAToken(t *testing.T) {
	c, seen := newSignupClient(t, "")

	err := c.API.Membership.BlockUser(context.Background(), map[string]any{"id": "usr_1"}, nil)

	var ne *norbixerr.Error
	if !errors.As(err, &ne) {
		t.Fatalf("expected a *errors.Error, got %v", err)
	}
	if ne.Code != "NORBIX_NOT_AUTHENTICATED" {
		t.Errorf("code: got %q want %q", ne.Code, "NORBIX_NOT_AUTHENTICATED")
	}
	if seen.hits != 0 {
		t.Errorf("requests: got %d want 0 — nothing may reach the gateway", seen.hits)
	}
}
