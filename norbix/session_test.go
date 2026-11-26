package norbix

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

// Sessions (owner decision D18): a project sign-in token past half its window
// gets a fresh one in the norbix-session-token answer header. The client swaps
// it in only when it sent its own stored bearer token on that call, and tells
// the app through Options.OnSessionToken. Every case starts its own fake
// gateway; nothing leaves the process.

type sessionGateway struct {
	mu       sync.Mutex
	auth     []string       // Authorization header of each call, in order
	body     map[string]any // last JSON body
	fresh    string         // norbix-session-token value to answer ("" = no header)
	setEmpty bool           // answer the header with an empty value
}

func (g *sessionGateway) seen() []string {
	g.mu.Lock()
	defer g.mu.Unlock()
	return append([]string(nil), g.auth...)
}

func newSessionClient(t *testing.T, opts Options) (*Client, *sessionGateway) {
	t.Helper()
	t.Setenv("NORBIX_API_KEY", "")
	t.Setenv("NORBIX_BEARER_TOKEN", "")
	g := &sessionGateway{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		raw, _ := io.ReadAll(r.Body)
		g.mu.Lock()
		g.auth = append(g.auth, r.Header.Get("Authorization"))
		g.body = nil
		if len(raw) > 0 {
			_ = json.Unmarshal(raw, &g.body)
		}
		if g.fresh != "" {
			w.Header().Set("norbix-session-token", g.fresh)
		} else if g.setEmpty {
			w.Header()["Norbix-Session-Token"] = []string{""}
		}
		g.mu.Unlock()
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path == "/auth" {
			_, _ = w.Write([]byte(`{"bearerToken":"jwt_signed_in","userId":"u_1"}`))
			return
		}
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	t.Cleanup(srv.Close)
	opts.ProjectID = "proj_1"
	opts.BaseURLAPI = srv.URL
	opts.BaseURLHub = srv.URL
	c, err := New(opts)
	if err != nil {
		t.Fatal(err)
	}
	return c, g
}

func echo(t *testing.T, c *Client) {
	t.Helper()
	if err := c.API.Echo.Echo(context.Background(), nil, nil); err != nil {
		t.Fatalf("echo: %v", err)
	}
}

func TestSessionTokenHeaderSwapsTheBearerTokenAndCallsTheCallback(t *testing.T) {
	var saved []string
	c, g := newSessionClient(t, Options{
		BearerToken:    "jwt_old",
		OnSessionToken: func(token string) { saved = append(saved, token) },
	})
	g.fresh = "jwt_new"
	echo(t, c)

	if c.BearerToken() != "jwt_new" {
		t.Fatalf("token not swapped: %q", c.BearerToken())
	}
	if len(saved) != 1 || saved[0] != "jwt_new" {
		t.Fatalf("callback: got %v want [jwt_new]", saved)
	}

	g.fresh = ""
	echo(t, c)
	if got := g.seen(); got[1] != "Bearer jwt_new" {
		t.Errorf("next call sent %q, want the fresh token", got[1])
	}
	if len(saved) != 1 {
		t.Errorf("callback called again without a header: %v", saved)
	}
}

func TestSessionTokenHeaderAbsentOrEmptyLeavesTheTokenAlone(t *testing.T) {
	called := false
	c, g := newSessionClient(t, Options{
		BearerToken:    "jwt_old",
		OnSessionToken: func(string) { called = true },
	})
	echo(t, c)
	g.setEmpty = true
	echo(t, c)

	if c.BearerToken() != "jwt_old" || called {
		t.Fatalf("token changed without a header value: %q, callback %v", c.BearerToken(), called)
	}
}

func TestSessionTokenHeaderIsIgnoredForAnAPIKeyClient(t *testing.T) {
	called := false
	c, g := newSessionClient(t, Options{
		APIKey:         "key_1",
		OnSessionToken: func(string) { called = true },
	})
	g.fresh = "jwt_new"
	echo(t, c)
	echo(t, c)

	if c.BearerToken() != "" || called {
		t.Fatalf("an API-key client took a token: %q, callback %v", c.BearerToken(), called)
	}
	for _, a := range g.seen() {
		if a != "Bearer key_1" {
			t.Errorf("sent %q, want the API key", a)
		}
	}
}

func TestSessionTokenHeaderAfterLogoutChangesNothing(t *testing.T) {
	c, g := newSessionClient(t, Options{BearerToken: "jwt_old"})
	g.fresh = "jwt_new"
	c.Logout()
	c.SetAPIKey("key_1")
	echo(t, c)
	if c.BearerToken() != "" {
		t.Fatalf("a signed-out client took a token: %q", c.BearerToken())
	}
}

func TestLoginSendsRememberMe(t *testing.T) {
	for _, remember := range []bool{false, true} {
		c, g := newSessionClient(t, Options{})
		_, err := c.Login(context.Background(), LoginCredentials{
			UserName: "ada", Password: "secret", RememberMe: remember,
		})
		if err != nil {
			t.Fatal(err)
		}
		if g.body["rememberMe"] != remember {
			t.Errorf("rememberMe: got %v want %v (body %v)", g.body["rememberMe"], remember, g.body)
		}
		if c.BearerToken() != "jwt_signed_in" {
			t.Errorf("login token not stored: %q", c.BearerToken())
		}
	}
}
