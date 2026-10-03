package hub

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/norbix-code/sdk-go/v2/norbix/internal/transport"
)

// Every LLM and MCP integration route (project audit, item E1) and the
// AiModule method that calls it. Each case asserts the verb and the
// fully-resolved path. Nothing leaves the process: every call goes to a local
// test server, no provider is ever contacted.
//
// SetLlmIntegrationAsDefault and the embedding routes are in
// ai_embeddings_test.go.

type aiIntegrationCase struct {
	name string
	verb string
	path string
	call func(ctx context.Context, m *AiModule) error
}

func aiIntegrationCases() []aiIntegrationCase {
	body := map[string]any{"probe": "value"}

	return []aiIntegrationCase{
		{"SaveLlmIntegration", http.MethodPost, "/v2/ai/integrations/llms/",
			func(ctx context.Context, m *AiModule) error {
				return m.SaveLlmIntegration(ctx, body, nil)
			}},
		{"TestLlmIntegration", http.MethodPost, "/v2/ai/integrations/llms/test",
			func(ctx context.Context, m *AiModule) error {
				return m.TestLlmIntegration(ctx, body, nil)
			}},
		{"GetLlmIntegrations", http.MethodGet, "/v2/ai/integrations/llms/integrations",
			func(ctx context.Context, m *AiModule) error {
				return m.GetLlmIntegrations(ctx, body, nil)
			}},
		{"GetLlmIntegration", http.MethodGet, "/v2/ai/integrations/llms/id_1",
			func(ctx context.Context, m *AiModule) error {
				return m.GetLlmIntegration(ctx, "id_1", body, nil)
			}},
		{"DeleteLlmIntegration", http.MethodDelete, "/v2/ai/integrations/llms/id_1",
			func(ctx context.Context, m *AiModule) error {
				return m.DeleteLlmIntegration(ctx, "id_1", body, nil)
			}},
		{"EnableLlmIntegration", http.MethodPut, "/v2/ai/integrations/llms/id_1/enable",
			func(ctx context.Context, m *AiModule) error {
				return m.EnableLlmIntegration(ctx, "id_1", body, nil)
			}},
		{"DisableLlmIntegration", http.MethodPut, "/v2/ai/integrations/llms/id_1/disable",
			func(ctx context.Context, m *AiModule) error {
				return m.DisableLlmIntegration(ctx, "id_1", body, nil)
			}},
		{"SaveMcpIntegration", http.MethodPost, "/v2/ai/integrations/mcp/",
			func(ctx context.Context, m *AiModule) error {
				return m.SaveMcpIntegration(ctx, body, nil)
			}},
		{"TestMcpIntegration", http.MethodPost, "/v2/ai/integrations/mcp/test",
			func(ctx context.Context, m *AiModule) error {
				return m.TestMcpIntegration(ctx, body, nil)
			}},
		{"GetMcpIntegrations", http.MethodGet, "/v2/ai/integrations/mcp/integrations",
			func(ctx context.Context, m *AiModule) error {
				return m.GetMcpIntegrations(ctx, body, nil)
			}},
		{"GetMcpIntegration", http.MethodGet, "/v2/ai/integrations/mcp/id_1",
			func(ctx context.Context, m *AiModule) error {
				return m.GetMcpIntegration(ctx, "id_1", body, nil)
			}},
		{"DeleteMcpIntegration", http.MethodDelete, "/v2/ai/integrations/mcp/id_1",
			func(ctx context.Context, m *AiModule) error {
				return m.DeleteMcpIntegration(ctx, "id_1", body, nil)
			}},
		{"EnableMcpIntegration", http.MethodPut, "/v2/ai/integrations/mcp/id_1/enable",
			func(ctx context.Context, m *AiModule) error {
				return m.EnableMcpIntegration(ctx, "id_1", body, nil)
			}},
		{"DisableMcpIntegration", http.MethodPut, "/v2/ai/integrations/mcp/id_1/disable",
			func(ctx context.Context, m *AiModule) error {
				return m.DisableMcpIntegration(ctx, "id_1", body, nil)
			}},
	}
}

func newAiIntegrationTestModule(baseURL string) *AiModule {
	return &AiModule{t: transport.New(&transport.Config{
		ProjectID:  "proj_1",
		APIKey:     "key_1",
		BaseURLAPI: baseURL,
		BaseURLHub: baseURL,
		APIVersion: "v2",
		HubVersion: "v2",
		Timeout:    5 * time.Second,
		Env:        "PROD",
	}, nil)}
}

func TestAiIntegrationEndpointsHitTheExpectedRoute(t *testing.T) {
	for _, c := range aiIntegrationCases() {
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

			if err := c.call(context.Background(), newAiIntegrationTestModule(srv.URL)); err != nil {
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

// The LLM + MCP integration surface is 14 routes; a changed count means a
// route arrived or left untested.
func TestAiIntegrationSurfaceSize(t *testing.T) {
	const want = 14
	if got := len(aiIntegrationCases()); got != want {
		t.Errorf("ai integration endpoint count: got %d want %d", got, want)
	}
}
