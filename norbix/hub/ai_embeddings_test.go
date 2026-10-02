package hub

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/norbix-code/sdk-go/v2/norbix/internal/transport"
)

// Every wave-3 embeddings route and the module method that calls it. Each case
// asserts the verb and the fully-resolved path, so a wrong path or a swapped
// verb fails here instead of at runtime. Nothing leaves the process: every
// call goes to a local test server, no provider is ever contacted.

type embeddingsCase struct {
	name string
	verb string
	path string
	call func(ctx context.Context, m *AiModule) error
}

func embeddingsCases() []embeddingsCase {
	body := map[string]any{"probe": "value"}

	return []embeddingsCase{
		{"GetEmbeddingIntegrations", http.MethodGet, "/v2/ai/integrations/embeddings",
			func(ctx context.Context, m *AiModule) error {
				return m.GetEmbeddingIntegrations(ctx, body, nil)
			}},
		{"SaveEmbeddingIntegration", http.MethodPost, "/v2/ai/integrations/embeddings",
			func(ctx context.Context, m *AiModule) error {
				return m.SaveEmbeddingIntegration(ctx, body, nil)
			}},
		{"GetEmbeddingIntegration", http.MethodGet, "/v2/ai/integrations/embeddings/id_1",
			func(ctx context.Context, m *AiModule) error {
				return m.GetEmbeddingIntegration(ctx, "id_1", body, nil)
			}},
		{"DeleteEmbeddingIntegration", http.MethodDelete, "/v2/ai/integrations/embeddings/id_1",
			func(ctx context.Context, m *AiModule) error {
				return m.DeleteEmbeddingIntegration(ctx, "id_1", body, nil)
			}},
		{"TestEmbeddingIntegration", http.MethodPost, "/v2/ai/integrations/embeddings/id_1/test",
			func(ctx context.Context, m *AiModule) error {
				return m.TestEmbeddingIntegration(ctx, "id_1", body, nil)
			}},
		{"SetLlmIntegrationAsDefault", http.MethodPut, "/v2/ai/integrations/llms/id_1/default",
			func(ctx context.Context, m *AiModule) error {
				return m.SetLlmIntegrationAsDefault(ctx, "id_1", body, nil)
			}},
	}
}

func newEmbeddingsTestModule(baseURL string) *AiModule {
	return &AiModule{t: transport.New(&transport.Config{
		ProjectID:  "proj_1",
		AccountID:  "acct_1",
		APIKey:     "key_1",
		BaseURLAPI: baseURL,
		BaseURLHub: baseURL,
		APIVersion: "v2",
		HubVersion: "v2",
		Timeout:    5 * time.Second,
		Env:        "PROD",
	}, nil)}
}

func TestEmbeddingsEndpointsHitTheExpectedRoute(t *testing.T) {
	for _, c := range embeddingsCases() {
		t.Run(c.name, func(t *testing.T) {
			var gotMethod, gotPath, gotAuth, gotProject string
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotMethod = r.Method
				gotPath = r.URL.Path
				gotAuth = r.Header.Get("Authorization")
				gotProject = r.Header.Get("X-CM-ProjectId")
				w.WriteHeader(http.StatusNoContent)
			}))
			defer srv.Close()

			if err := c.call(context.Background(), newEmbeddingsTestModule(srv.URL)); err != nil {
				t.Fatalf("%s: unexpected error: %v", c.name, err)
			}
			if gotMethod != c.verb {
				t.Errorf("%s verb: got %q want %q", c.name, gotMethod, c.verb)
			}
			if gotPath != c.path {
				t.Errorf("%s path: got %q want %q", c.name, gotPath, c.path)
			}
			if gotAuth != "Bearer key_1" {
				t.Errorf("%s auth header: got %q", c.name, gotAuth)
			}
			if gotProject != "proj_1" {
				t.Errorf("%s project header: got %q", c.name, gotProject)
			}
		})
	}
}

// The wave-3 embeddings surface is 6 routes; a changed count means a route arrived
// or left untested.
func TestEmbeddingsSurfaceSize(t *testing.T) {
	const want = 6
	if got := len(embeddingsCases()); got != want {
		t.Errorf("embeddings endpoint count: got %d want %d", got, want)
	}
}
