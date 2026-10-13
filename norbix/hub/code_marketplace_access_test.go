package hub

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/norbix-code/sdk-go/v2/norbix/hub/dtos"
)

// GET /{version}/code/marketplace/integrations/{IntegrationViewId}/access returns
// the access matrix of a connection (gateway item access, Code.Marketplace.md
// §12, route 25). The route table in code_marketplace_test.go pins verb, path
// and body for routes 25–27; this test pins that the matrix decodes into the
// generated response, with the generated MarketplaceAccessLevel values.
func TestGetMarketplaceIntegrationAccessDecodesTheMatrix(t *testing.T) {
	var gotMethod, gotPath string
	var gotBody []byte
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod, gotPath = r.Method, r.URL.Path
		gotBody, _ = io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"roles":[{"roleId":"acc_6Qp1dW2kRm8NsT4vYc0Hb_nr_3kLm9QwErTyU","name":"Projects Developer","scope":"account","isSystem":true,"package":"None","custom":[],"defaultNotApplied":"The person who enabled the connection may not manage policies."},{"roleId":"pr_1_nr_2","name":"Support","scope":"project","package":"ReadOnly","isCustom":true,"custom":[{"policyId":"pol_9","name":"GitHub — Support","usedByRoles":2}]}],"packages":[{"kind":"ReadOnly","policyId":"pol_1","name":"GitHub — Read-only","exists":true}],"tools":[{"name":"get_me","readOnly":true,"grantable":true}],"functions":[{"functionId":"func_1","name":"Create issue","functionKey":"create_issue"}]}`)
	}))
	defer srv.Close()

	var out dtos.GetMarketplaceIntegrationAccessResponse
	if err := newCodeTestModule(srv.URL).GetMarketplaceIntegrationAccess(context.Background(), "int_6Qp1dW2kRm8NsT4vYc0Hb", map[string]any{}, &out); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if gotMethod != http.MethodGet || gotPath != "/v2/code/marketplace/integrations/int_6Qp1dW2kRm8NsT4vYc0Hb/access" {
		t.Errorf("request: got %s %s", gotMethod, gotPath)
	}
	if len(gotBody) != 0 {
		t.Errorf("body: got %s want none", gotBody)
	}
	if len(out.Roles) != 2 || len(out.Packages) != 1 || len(out.Tools) != 1 || len(out.Functions) != 1 {
		t.Fatalf("matrix: got %+v", out)
	}
	if out.Roles[0].Package != dtos.MarketplaceAccessLevelNone || out.Roles[0].DefaultNotApplied == "" {
		t.Errorf("role 0: got %+v", *out.Roles[0])
	}
	if out.Roles[1].Package != dtos.MarketplaceAccessLevelReadOnly || !out.Roles[1].IsCustom || out.Roles[1].Custom[0].UsedByRoles != 2 {
		t.Errorf("role 1: got %+v", *out.Roles[1])
	}
	if out.Packages[0].Kind != dtos.MarketplaceAccessLevelReadOnly || !out.Tools[0].Grantable || out.Functions[0].FunctionKey != "create_issue" {
		t.Errorf("packages / tools / functions: got %+v %+v %+v", *out.Packages[0], *out.Tools[0], *out.Functions[0])
	}
}
