package hub

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// The 24 marketplace routes of the Code module (gateway campaign audit/marketplace,
// docs/architecture/Code.Marketplace.md §14; route 24 = MCP tool catalog, item mcp-step1) and the module method that calls each
// one. Every case asserts the verb, the fully-resolved path (version and ids
// substituted) and where the parameters land: the query string for GET / DELETE,
// the JSON body for POST / PUT. Nothing leaves the process.

func marketplaceCases() []codeCase {
	return []codeCase{
		{"GetMarketplaceListings", http.MethodGet, "/v2/code/marketplace/listings", url.Values{"search": {"slack"}}, nil,
			func(ctx context.Context, m *CodeModule) error {
				return m.GetMarketplaceListings(ctx, map[string]any{"search": "slack"}, nil)
			}},
		{"GetMarketplaceListingFunctionTokens", http.MethodGet, "/v2/code/marketplace/listings/ml_slack_v1/functions/slack.post-message/tokens", nil, nil,
			func(ctx context.Context, m *CodeModule) error {
				return m.GetMarketplaceListingFunctionTokens(ctx, "ml_slack_v1", "slack.post-message", map[string]any{}, nil)
			}},
		{"GetMarketplaceIntegrations", http.MethodGet, "/v2/code/marketplace/integrations", nil, nil,
			func(ctx context.Context, m *CodeModule) error {
				return m.GetMarketplaceIntegrations(ctx, map[string]any{}, nil)
			}},
		{"GetMarketplaceIntegration", http.MethodGet, "/v2/code/marketplace/integrations/int_6Qp1dW2kRm8NsT4vYc0Hb", nil, nil,
			func(ctx context.Context, m *CodeModule) error {
				return m.GetMarketplaceIntegration(ctx, "int_6Qp1dW2kRm8NsT4vYc0Hb", map[string]any{}, nil)
			}},
		{"SaveMarketplaceIntegration", http.MethodPost, "/v2/code/marketplace/integrations", nil, map[string]any{"integration": map[string]any{"listingViewId": "ml_slack_v1", "config": map[string]any{"workspace": "acme"}}, "secrets": map[string]any{"botToken": "xoxb-test"}},
			func(ctx context.Context, m *CodeModule) error {
				return m.SaveMarketplaceIntegration(ctx, map[string]any{"integration": map[string]any{"listingViewId": "ml_slack_v1", "config": map[string]any{"workspace": "acme"}}, "secrets": map[string]any{"botToken": "xoxb-test"}}, nil)
			}},
		{"DeleteMarketplaceIntegration", http.MethodDelete, "/v2/code/marketplace/integrations/int_6Qp1dW2kRm8NsT4vYc0Hb", nil, nil,
			func(ctx context.Context, m *CodeModule) error {
				return m.DeleteMarketplaceIntegration(ctx, "int_6Qp1dW2kRm8NsT4vYc0Hb", map[string]any{}, nil)
			}},
		{"EnableMarketplaceIntegration", http.MethodPost, "/v2/code/marketplace/integrations/int_6Qp1dW2kRm8NsT4vYc0Hb/enable", nil, map[string]any{},
			func(ctx context.Context, m *CodeModule) error {
				return m.EnableMarketplaceIntegration(ctx, "int_6Qp1dW2kRm8NsT4vYc0Hb", map[string]any{}, nil)
			}},
		{"DisableMarketplaceIntegration", http.MethodPost, "/v2/code/marketplace/integrations/int_6Qp1dW2kRm8NsT4vYc0Hb/disable", nil, map[string]any{},
			func(ctx context.Context, m *CodeModule) error {
				return m.DisableMarketplaceIntegration(ctx, "int_6Qp1dW2kRm8NsT4vYc0Hb", map[string]any{}, nil)
			}},
		{"GetMarketplaceFunctions", http.MethodGet, "/v2/code/marketplace/integrations/int_6Qp1dW2kRm8NsT4vYc0Hb/functions", nil, nil,
			func(ctx context.Context, m *CodeModule) error {
				return m.GetMarketplaceFunctions(ctx, "int_6Qp1dW2kRm8NsT4vYc0Hb", map[string]any{}, nil)
			}},
		{"GetMarketplaceFunction", http.MethodGet, "/v2/code/marketplace/integrations/int_6Qp1dW2kRm8NsT4vYc0Hb/functions/func_4hT2b9QkLm3XcT1yRfGh7a", nil, nil,
			func(ctx context.Context, m *CodeModule) error {
				return m.GetMarketplaceFunction(ctx, "int_6Qp1dW2kRm8NsT4vYc0Hb", "func_4hT2b9QkLm3XcT1yRfGh7a", map[string]any{}, nil)
			}},
		{"SaveMarketplaceFunction", http.MethodPost, "/v2/code/marketplace/integrations/int_6Qp1dW2kRm8NsT4vYc0Hb/functions", nil, map[string]any{"function": map[string]any{"functionKey": "slack.post-message", "displayName": "Post order", "arguments": map[string]any{"channel": "#orders"}}},
			func(ctx context.Context, m *CodeModule) error {
				return m.SaveMarketplaceFunction(ctx, "int_6Qp1dW2kRm8NsT4vYc0Hb", map[string]any{"function": map[string]any{"functionKey": "slack.post-message", "displayName": "Post order", "arguments": map[string]any{"channel": "#orders"}}}, nil)
			}},
		{"DeleteMarketplaceFunction", http.MethodDelete, "/v2/code/marketplace/integrations/int_6Qp1dW2kRm8NsT4vYc0Hb/functions/func_4hT2b9QkLm3XcT1yRfGh7a", nil, nil,
			func(ctx context.Context, m *CodeModule) error {
				return m.DeleteMarketplaceFunction(ctx, "int_6Qp1dW2kRm8NsT4vYc0Hb", "func_4hT2b9QkLm3XcT1yRfGh7a", map[string]any{}, nil)
			}},
		{"EnableMarketplaceFunction", http.MethodPost, "/v2/code/marketplace/integrations/int_6Qp1dW2kRm8NsT4vYc0Hb/functions/func_4hT2b9QkLm3XcT1yRfGh7a/enable", nil, map[string]any{},
			func(ctx context.Context, m *CodeModule) error {
				return m.EnableMarketplaceFunction(ctx, "int_6Qp1dW2kRm8NsT4vYc0Hb", "func_4hT2b9QkLm3XcT1yRfGh7a", map[string]any{}, nil)
			}},
		{"DisableMarketplaceFunction", http.MethodPost, "/v2/code/marketplace/integrations/int_6Qp1dW2kRm8NsT4vYc0Hb/functions/func_4hT2b9QkLm3XcT1yRfGh7a/disable", nil, map[string]any{},
			func(ctx context.Context, m *CodeModule) error {
				return m.DisableMarketplaceFunction(ctx, "int_6Qp1dW2kRm8NsT4vYc0Hb", "func_4hT2b9QkLm3XcT1yRfGh7a", map[string]any{}, nil)
			}},
		{"GetMarketplaceFunctionTokens", http.MethodGet, "/v2/code/marketplace/integrations/int_6Qp1dW2kRm8NsT4vYc0Hb/functions/func_4hT2b9QkLm3XcT1yRfGh7a/tokens", nil, nil,
			func(ctx context.Context, m *CodeModule) error {
				return m.GetMarketplaceFunctionTokens(ctx, "int_6Qp1dW2kRm8NsT4vYc0Hb", "func_4hT2b9QkLm3XcT1yRfGh7a", map[string]any{}, nil)
			}},
		{"InvokeMarketplaceFunction", http.MethodPost, "/v2/code/marketplace/functions/func_4hT2b9QkLm3XcT1yRfGh7a/invoke", nil, map[string]any{"payload": map[string]any{"text": "Order 6ac54a59 shipped"}},
			func(ctx context.Context, m *CodeModule) error {
				return m.InvokeMarketplaceFunction(ctx, "func_4hT2b9QkLm3XcT1yRfGh7a", map[string]any{"payload": map[string]any{"text": "Order 6ac54a59 shipped"}}, nil)
			}},
		{"GetMarketplaceListing", http.MethodGet, "/v2/code/marketplace/listings/ml_slack_v1", nil, nil,
			func(ctx context.Context, m *CodeModule) error {
				return m.GetMarketplaceListing(ctx, "ml_slack_v1", map[string]any{}, nil)
			}},
		{"TestMarketplaceIntegration", http.MethodPost, "/v2/code/marketplace/integrations/int_6Qp1dW2kRm8NsT4vYc0Hb/test", nil, map[string]any{},
			func(ctx context.Context, m *CodeModule) error {
				return m.TestMarketplaceIntegration(ctx, "int_6Qp1dW2kRm8NsT4vYc0Hb", map[string]any{}, nil)
			}},
		{"GetMarketplaceIntegrationTools", http.MethodGet, "/v2/code/marketplace/integrations/int_6Qp1dW2kRm8NsT4vYc0Hb/tools", nil, nil,
			func(ctx context.Context, m *CodeModule) error {
				return m.GetMarketplaceIntegrationTools(ctx, "int_6Qp1dW2kRm8NsT4vYc0Hb", map[string]any{}, nil)
			}},
		{"ReplaceMarketplaceIntegrationSecrets", http.MethodPut, "/v2/code/marketplace/integrations/int_6Qp1dW2kRm8NsT4vYc0Hb/secrets", nil, map[string]any{"secrets": map[string]any{"botToken": "xoxb-new"}},
			func(ctx context.Context, m *CodeModule) error {
				return m.ReplaceMarketplaceIntegrationSecrets(ctx, "int_6Qp1dW2kRm8NsT4vYc0Hb", map[string]any{"secrets": map[string]any{"botToken": "xoxb-new"}}, nil)
			}},
		{"RevealMarketplaceIntegrationSecrets", http.MethodPost, "/v2/code/marketplace/integrations/int_6Qp1dW2kRm8NsT4vYc0Hb/secrets/reveal", nil, map[string]any{},
			func(ctx context.Context, m *CodeModule) error {
				return m.RevealMarketplaceIntegrationSecrets(ctx, "int_6Qp1dW2kRm8NsT4vYc0Hb", map[string]any{}, nil)
			}},
		{"SetMarketplaceIntegrationTokenMappings", http.MethodPut, "/v2/code/marketplace/integrations/int_6Qp1dW2kRm8NsT4vYc0Hb/token-mappings", nil, map[string]any{"tokenMappings": []any{map[string]any{"key": "channel", "value": "#orders", "resolver": "Custom"}}},
			func(ctx context.Context, m *CodeModule) error {
				return m.SetMarketplaceIntegrationTokenMappings(ctx, "int_6Qp1dW2kRm8NsT4vYc0Hb", map[string]any{"tokenMappings": []any{map[string]any{"key": "channel", "value": "#orders", "resolver": "Custom"}}}, nil)
			}},
		{"GetMarketplaceFunctionCatalog", http.MethodGet, "/v2/code/marketplace/integrations/int_6Qp1dW2kRm8NsT4vYc0Hb/catalog", nil, nil,
			func(ctx context.Context, m *CodeModule) error {
				return m.GetMarketplaceFunctionCatalog(ctx, "int_6Qp1dW2kRm8NsT4vYc0Hb", map[string]any{}, nil)
			}},
		{"GetMarketplaceFunctionRuns", http.MethodGet, "/v2/code/marketplace/functions/runs", url.Values{"triggerId": {"trg_1p4R3gTobGZpD6fXvHT4rd"}, "limit": {"5"}}, nil,
			func(ctx context.Context, m *CodeModule) error {
				return m.GetMarketplaceFunctionRuns(ctx, map[string]any{"triggerId": "trg_1p4R3gTobGZpD6fXvHT4rd", "limit": 5}, nil)
			}},
	}
}

func TestMarketplaceEndpointsHitTheExpectedRoute(t *testing.T) {
	for _, c := range marketplaceCases() {
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

			if err := c.call(context.Background(), newCodeTestModule(srv.URL)); err != nil {
				t.Fatalf("%s: unexpected error: %v", c.name, err)
			}
			if gotMethod != c.verb {
				t.Errorf("verb: got %q want %q", gotMethod, c.verb)
			}
			if gotPath != c.path {
				t.Errorf("path: got %q want %q", gotPath, c.path)
			}
			gotQuery, _ := url.ParseQuery(gotRawQuery)
			if c.query == nil && gotRawQuery != "" {
				t.Errorf("query: got %q want none", gotRawQuery)
			}
			if c.query != nil {
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
			if len(gotBody) == 0 {
				got = map[string]any{}
			} else if err := json.Unmarshal(gotBody, &got); err != nil {
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

func TestMarketplaceCaseCount(t *testing.T) {
	const want = 24
	if got := len(marketplaceCases()); got != want {
		t.Errorf("code/marketplace case count: got %d want %d", got, want)
	}
}
