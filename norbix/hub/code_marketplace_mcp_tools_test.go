package hub

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/norbix-code/sdk-go/v2/norbix/hub/dtos"
)

// GET /{version}/code/marketplace/integrations/{IntegrationViewId}/tools returns
// the stored tool catalog of an MCP connection (gateway item mcp-step1,
// Code.Marketplace.md §10.2). The route table in code_marketplace_test.go pins
// verb and path; this test pins the connection id in the path, no query, no
// body, and that the catalog decodes into the generated response.
func TestGetMarketplaceIntegrationToolsDecodesTheCatalog(t *testing.T) {
	var gotMethod, gotPath, gotRawQuery string
	var gotBody []byte
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod, gotPath, gotRawQuery = r.Method, r.URL.Path, r.URL.RawQuery
		gotBody, _ = io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"tools":[{"name":"create_issue","title":"Create issue","description":"Open a new issue in a repository","inputSchema":"{\"type\":\"object\",\"required\":[\"title\"]}","readOnlyHint":false,"destructiveHint":false},{"name":"get_me","readOnlyHint":true}],"listedAt":"2026-10-13T08:00:00Z","catalogHash":"4ce4259bec02f283","truncated":false}`)
	}))
	defer srv.Close()

	var out dtos.GetMarketplaceIntegrationToolsResponse
	if err := newCodeTestModule(srv.URL).GetMarketplaceIntegrationTools(context.Background(), "int_7Hq2mcpServer", map[string]any{}, &out); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if gotMethod != http.MethodGet {
		t.Errorf("verb: got %q want GET", gotMethod)
	}
	if gotPath != "/v2/code/marketplace/integrations/int_7Hq2mcpServer/tools" {
		t.Errorf("path: got %q", gotPath)
	}
	if gotRawQuery != "" {
		t.Errorf("query: got %q want none", gotRawQuery)
	}
	if len(gotBody) != 0 {
		t.Errorf("body: got %s want none", gotBody)
	}

	if len(out.Tools) != 2 || out.CatalogHash != "4ce4259bec02f283" || out.Truncated {
		t.Fatalf("catalog: got %+v", out)
	}
	if out.Tools[0].Name != "create_issue" || out.Tools[0].InputSchema != `{"type":"object","required":["title"]}` || out.Tools[0].ReadOnlyHint {
		t.Errorf("tool 0: got %+v", *out.Tools[0])
	}
	if out.Tools[1].Name != "get_me" || !out.Tools[1].ReadOnlyHint {
		t.Errorf("tool 1: got %+v", *out.Tools[1])
	}
}
