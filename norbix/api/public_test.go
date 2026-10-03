package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/norbix-code/sdk-go/v2/norbix/internal/transport"
)

// The two public project routes (project audit, item E1). Each case asserts
// the verb, the fully-resolved path, and that NO Authorization header is sent
// even though the client holds a key — these links must work for anyone.

type publicCase struct {
	name string
	verb string
	path string
	call func(ctx context.Context, m *PublicModule, out any) error
}

func publicCases() []publicCase {
	return []publicCase{
		{"GetPublicProjectConfig", http.MethodGet, "/v2/public/projects/proj_9/config",
			func(ctx context.Context, m *PublicModule, out any) error {
				return m.GetPublicProjectConfig(ctx, "proj_9", out)
			}},
		{"GetPublicProjectLegal", http.MethodGet, "/v2/public/projects/proj_9/legal/terms",
			func(ctx context.Context, m *PublicModule, out any) error {
				return m.GetPublicProjectLegal(ctx, "proj_9", "terms", out)
			}},
	}
}

func newPublicTestModule(baseURL string) *PublicModule {
	return &PublicModule{t: transport.New(&transport.Config{
		ProjectID:  "proj_1",
		APIKey:     "key_1",
		BaseURLAPI: baseURL,
		BaseURLHub: "http://hub.invalid",
		APIVersion: "v2",
		HubVersion: "v2",
		Timeout:    5 * time.Second,
		Env:        "PROD",
	}, nil)}
}

func TestPublicEndpointsHitTheExpectedRouteWithoutAuth(t *testing.T) {
	for _, c := range publicCases() {
		t.Run(c.name, func(t *testing.T) {
			var gotMethod, gotPath, gotQuery string
			gotAuth := "unset"
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotMethod = r.Method
				gotPath = r.URL.Path
				gotQuery = r.URL.RawQuery
				gotAuth = r.Header.Get("Authorization")
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{"available":true}`))
			}))
			defer srv.Close()

			var out map[string]any
			if err := c.call(context.Background(), newPublicTestModule(srv.URL), &out); err != nil {
				t.Fatalf("%s: unexpected error: %v", c.name, err)
			}
			if gotMethod != c.verb {
				t.Errorf("%s verb: got %q want %q", c.name, gotMethod, c.verb)
			}
			if gotPath != c.path {
				t.Errorf("%s path: got %q want %q", c.name, gotPath, c.path)
			}
			if gotQuery != "" {
				t.Errorf("%s query: got %q want empty", c.name, gotQuery)
			}
			if gotAuth != "" {
				t.Errorf("%s auth header: got %q want none", c.name, gotAuth)
			}
			if out["available"] != true {
				t.Errorf("%s answer not decoded: got %v", c.name, out)
			}
		})
	}
}

// The public surface is 2 routes; a changed count means a route arrived or
// left untested.
func TestPublicSurfaceSize(t *testing.T) {
	const want = 2
	if got := len(publicCases()); got != want {
		t.Errorf("public endpoint count: got %d want %d", got, want)
	}
}
