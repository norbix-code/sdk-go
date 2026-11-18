package hub

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/norbix-code/sdk-go/v2/norbix/internal/transport"
)

// Matrix sections 06 AI · MCP, 07 AI · LLM integration and 08 AI chat, Hub
// side: every LLM and MCP integration route of AiModule (the MCP ones are
// deprecated — save is refused with CM-ERRORS-…-030 — but still served) and
// the AI tool routes of AccountModule. The developer AI chat routes
// (/account/chat/*) were retired on 2026-11-18. Each case asserts the verb,
// the fully-resolved path on the HUB (the API base is an unreachable
// *.invalid host), the exact query string, the JSON body and the key +
// project headers. Field names are the gateway's (Hub.AI
// Integrations/{Llms,Mcps}/*.cs, AiToolsBridge.cs).
//
// Nothing leaves the process: every call goes to a local test server, and no
// LLM or MCP server is contacted.

type aiClusterCase[M any] struct {
	name  string
	verb  string
	path  string
	query url.Values     // nil: no query string expected
	body  map[string]any // nil: no body expected
	call  func(ctx context.Context, m M, req map[string]any) error
	req   map[string]any
}

func aiClusterIntegrationCases() []aiClusterCase[*AiModule] {
	return []aiClusterCase[*AiModule]{
		{"GetLlmIntegrations", http.MethodGet, "/v2/ai/integrations/llms/integrations", url.Values{"pageSize": {"50"}, "startingAfter": {"int_llm0"}}, nil,
			func(ctx context.Context, m *AiModule, req map[string]any) error {
				return m.GetLlmIntegrations(ctx, req, nil)
			}, map[string]any{"pageSize": 50, "startingAfter": "int_llm0"}},
		{"GetLlmIntegration", http.MethodGet, "/v2/ai/integrations/llms/int_llm1", nil, nil,
			func(ctx context.Context, m *AiModule, req map[string]any) error {
				return m.GetLlmIntegration(ctx, "int_llm1", req, nil)
			}, nil},
		{"SaveLlmIntegration", http.MethodPost, "/v2/ai/integrations/llms/", nil, map[string]any{"integration": map[string]any{"provider": "OpenAI", "integrationName": "OpenAI", "apiKey": "sk-test", "defaultModel": "gpt-4o-mini", "isDefault": true, "isEnabled": true}},
			func(ctx context.Context, m *AiModule, req map[string]any) error {
				return m.SaveLlmIntegration(ctx, req, nil)
			}, map[string]any{"integration": map[string]any{"provider": "OpenAI", "integrationName": "OpenAI", "apiKey": "sk-test", "defaultModel": "gpt-4o-mini", "isDefault": true, "isEnabled": true}}},
		{"TestLlmIntegration", http.MethodPost, "/v2/ai/integrations/llms/test", nil, map[string]any{"integrationId": "int_llm1"},
			func(ctx context.Context, m *AiModule, req map[string]any) error {
				return m.TestLlmIntegration(ctx, req, nil)
			}, map[string]any{"integrationId": "int_llm1"}},
		{"EnableLlmIntegration", http.MethodPut, "/v2/ai/integrations/llms/int_llm1/enable", nil, nil,
			func(ctx context.Context, m *AiModule, req map[string]any) error {
				return m.EnableLlmIntegration(ctx, "int_llm1", req, nil)
			}, nil},
		{"DisableLlmIntegration", http.MethodPut, "/v2/ai/integrations/llms/int_llm1/disable", nil, nil,
			func(ctx context.Context, m *AiModule, req map[string]any) error {
				return m.DisableLlmIntegration(ctx, "int_llm1", req, nil)
			}, nil},
		{"SetLlmIntegrationAsDefault", http.MethodPut, "/v2/ai/integrations/llms/int_llm1/default", nil, nil,
			func(ctx context.Context, m *AiModule, req map[string]any) error {
				return m.SetLlmIntegrationAsDefault(ctx, "int_llm1", req, nil)
			}, nil},
		{"DeleteLlmIntegration", http.MethodDelete, "/v2/ai/integrations/llms/int_llm1", nil, nil,
			func(ctx context.Context, m *AiModule, req map[string]any) error {
				return m.DeleteLlmIntegration(ctx, "int_llm1", req, nil)
			}, nil},
		{"GetMcpIntegrations", http.MethodGet, "/v2/ai/integrations/mcp/integrations", url.Values{"pageSize": {"20"}}, nil,
			func(ctx context.Context, m *AiModule, req map[string]any) error {
				return m.GetMcpIntegrations(ctx, req, nil)
			}, map[string]any{"pageSize": 20}},
		{"GetMcpIntegration", http.MethodGet, "/v2/ai/integrations/mcp/int_mcp1", nil, nil,
			func(ctx context.Context, m *AiModule, req map[string]any) error {
				return m.GetMcpIntegration(ctx, "int_mcp1", req, nil)
			}, nil},
		{"SaveMcpIntegration", http.MethodPost, "/v2/ai/integrations/mcp/", nil, map[string]any{"integration": map[string]any{"provider": "GitHub", "integrationName": "GitHub", "name": "GitHub", "category": "Code", "serverUrl": "https://api.githubcopilot.com/mcp/", "isEnabled": true}},
			func(ctx context.Context, m *AiModule, req map[string]any) error {
				return m.SaveMcpIntegration(ctx, req, nil)
			}, map[string]any{"integration": map[string]any{"provider": "GitHub", "integrationName": "GitHub", "name": "GitHub", "category": "Code", "serverUrl": "https://api.githubcopilot.com/mcp/", "isEnabled": true}}},
		{"TestMcpIntegration", http.MethodPost, "/v2/ai/integrations/mcp/test", nil, map[string]any{"integrationId": "int_mcp1"},
			func(ctx context.Context, m *AiModule, req map[string]any) error {
				return m.TestMcpIntegration(ctx, req, nil)
			}, map[string]any{"integrationId": "int_mcp1"}},
		{"EnableMcpIntegration", http.MethodPut, "/v2/ai/integrations/mcp/int_mcp1/enable", nil, nil,
			func(ctx context.Context, m *AiModule, req map[string]any) error {
				return m.EnableMcpIntegration(ctx, "int_mcp1", req, nil)
			}, nil},
		{"DisableMcpIntegration", http.MethodPut, "/v2/ai/integrations/mcp/int_mcp1/disable", nil, nil,
			func(ctx context.Context, m *AiModule, req map[string]any) error {
				return m.DisableMcpIntegration(ctx, "int_mcp1", req, nil)
			}, nil},
		{"DeleteMcpIntegration", http.MethodDelete, "/v2/ai/integrations/mcp/int_mcp1", nil, nil,
			func(ctx context.Context, m *AiModule, req map[string]any) error {
				return m.DeleteMcpIntegration(ctx, "int_mcp1", req, nil)
			}, nil},
	}
}

func aiClusterToolCases() []aiClusterCase[*AccountModule] {
	return []aiClusterCase[*AccountModule]{
		{"GetAiTools", http.MethodGet, "/v2/account/ai/tools", url.Values{"toolset": {"ai:database-read"}}, nil,
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.GetAiTools(ctx, req, nil)
			}, map[string]any{"toolset": "ai:database-read"}},
		{"InvokeAiTool", http.MethodPost, "/v2/account/ai/tools/get_schemas", nil, map[string]any{"argumentsJson": "{\"take\":5}"},
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.InvokeAiTool(ctx, "get_schemas", req, nil)
			}, map[string]any{"argumentsJson": "{\"take\":5}"}},
	}
}

func newAiClusterTestTransport(hubURL string) *transport.Transport {
	return transport.New(&transport.Config{
		ProjectID:  "proj_1",
		AccountID:  "acct_1",
		APIKey:     "key_1",
		BaseURLAPI: "http://api.invalid",
		BaseURLHub: hubURL,
		APIVersion: "v2",
		HubVersion: "v2",
		Timeout:    5 * time.Second,
		Env:        "PROD",
	}, nil)
}

func runAiClusterCases[M any](t *testing.T, cases []aiClusterCase[M], module func(hubURL string) M) {
	t.Helper()
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var gotMethod, gotPath, gotRawQuery, gotAuth, gotProject string
			var gotBody []byte
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotMethod = r.Method
				gotPath = r.URL.Path
				gotRawQuery = r.URL.RawQuery
				gotAuth = r.Header.Get("Authorization")
				gotProject = r.Header.Get("norbix-project-id")
				gotBody, _ = io.ReadAll(r.Body)
				w.WriteHeader(http.StatusNoContent)
			}))
			defer srv.Close()

			if err := c.call(context.Background(), module(srv.URL), c.req); err != nil {
				t.Fatalf("%s: unexpected error: %v", c.name, err)
			}
			if gotMethod != c.verb {
				t.Errorf("verb: got %q want %q", gotMethod, c.verb)
			}
			if gotPath != c.path {
				t.Errorf("path: got %q want %q", gotPath, c.path)
			}
			if gotAuth != "Bearer key_1" {
				t.Errorf("authorization: got %q want %q", gotAuth, "Bearer key_1")
			}
			if gotProject != "proj_1" {
				t.Errorf("project header: got %q want %q", gotProject, "proj_1")
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

func TestAiIntegrationRoutesHitTheExpectedRoute(t *testing.T) {
	runAiClusterCases(t, aiClusterIntegrationCases(), func(hubURL string) *AiModule {
		return &AiModule{t: newAiClusterTestTransport(hubURL)}
	})
}

func TestAiToolRoutesHitTheExpectedRoute(t *testing.T) {
	runAiClusterCases(t, aiClusterToolCases(), func(hubURL string) *AccountModule {
		return &AccountModule{t: newAiClusterTestTransport(hubURL)}
	})
}

// 8 LLM + 7 MCP integration routes; 2 AI tool routes. A changed count means
// a route arrived or left untested.
func TestAiClusterSurfaceSize(t *testing.T) {
	if got := len(aiClusterIntegrationCases()); got != 15 {
		t.Errorf("LLM + MCP integration route count: got %d want 15", got)
	}
	if got := len(aiClusterToolCases()); got != 2 {
		t.Errorf("AI tool route count: got %d want 2", got)
	}
}
