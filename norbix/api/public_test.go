package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"reflect"
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

// The brand logo / icon answers the image itself: the bytes come back
// untouched in a *[]byte, from the API host, with no Authorization header.
func TestGetPublicProjectBrandAssetReturnsTheImageBytes(t *testing.T) {
	png := []byte{0x89, 'P', 'N', 'G', 0x0d, 0x0a, 0x1a, 0x0a}
	var gotMethod, gotPath, gotQuery string
	gotAuth := "unset"
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod, gotPath, gotQuery = r.Method, r.URL.Path, r.URL.RawQuery
		gotAuth = r.Header.Get("Authorization")
		w.Header().Set("Content-Type", "image/png")
		_, _ = w.Write(png)
	}))
	defer srv.Close()

	want := publicCase{"GetPublicProjectBrandAsset", http.MethodGet, "/v2/public/projects/proj_9/brand/logo", nil}
	var out []byte
	if err := newPublicTestModule(srv.URL).GetPublicProjectBrandAsset(context.Background(), "proj_9", "logo", &out); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != want.verb {
		t.Errorf("verb: got %q want %q", gotMethod, want.verb)
	}
	if gotPath != want.path {
		t.Errorf("path: got %q want %q", gotPath, want.path)
	}
	if gotQuery != "" {
		t.Errorf("query: got %q want empty", gotQuery)
	}
	if gotAuth != "" {
		t.Errorf("auth header: got %q want none", gotAuth)
	}
	if string(out) != string(png) {
		t.Errorf("bytes: got %v want %v", out, png)
	}
}

// The public surface is 3 routes: the two JSON cases above and the brand
// asset in its own test. A changed method count means a route arrived or left
// untested.
func TestPublicSurfaceSize(t *testing.T) {
	if got := len(publicCases()); got != 2 {
		t.Errorf("public JSON endpoint count: got %d want 2", got)
	}
	if got := reflect.TypeOf(&PublicModule{}).NumMethod(); got != 3 {
		t.Errorf("PublicModule method count: got %d want 3", got)
	}
}
