package hub

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/norbix-code/sdk-go/v2/norbix/internal/transport"
)

// Every Project-module route (project audit, item E1) and the AccountModule
// method that calls it: project CRUD, project settings, Admin Portal, AI
// service users and the developer MCP endpoint. Each case asserts the verb,
// the fully-resolved path and the scope headers, so a wrong path, a swapped
// verb or a lost header fails here instead of at runtime. Nothing leaves the
// process: every call goes to a local test server.
//
// The AI settings / assistants / usage routes are in account_ai_test.go.

type projectCase struct {
	name  string
	verb  string
	path  string
	scope string // "account" also sends X-CM-AccountId and needs AccountID
	call  func(ctx context.Context, m *AccountModule) error
}

func projectCases() []projectCase {
	body := map[string]any{"probe": "value"}

	return []projectCase{
		{"CreateProject", http.MethodPost, "/v2/account/projects", "account",
			func(ctx context.Context, m *AccountModule) error {
				return m.CreateProject(ctx, body, nil)
			}},
		{"GetProjects", http.MethodGet, "/v2/account/projects", "account",
			func(ctx context.Context, m *AccountModule) error {
				return m.GetProjects(ctx, body, nil)
			}},
		{"GetProject", http.MethodGet, "/v2/account/projects/projectId_1", "account",
			func(ctx context.Context, m *AccountModule) error {
				return m.GetProject(ctx, "projectId_1", body, nil)
			}},
		{"DeleteProject", http.MethodDelete, "/v2/account/projects/projectId_1", "account",
			func(ctx context.Context, m *AccountModule) error {
				return m.DeleteProject(ctx, "projectId_1", body, nil)
			}},
		{"EnableProject", http.MethodPatch, "/v2/account/projects/projectId_1/enable", "account",
			func(ctx context.Context, m *AccountModule) error {
				return m.EnableProject(ctx, "projectId_1", body, nil)
			}},
		{"DisableProject", http.MethodPatch, "/v2/account/projects/projectId_1/disable", "account",
			func(ctx context.Context, m *AccountModule) error {
				return m.DisableProject(ctx, "projectId_1", body, nil)
			}},
		{"GetProjectTokens", http.MethodGet, "/v2/account/projects/projectId_1/tokens", "account",
			func(ctx context.Context, m *AccountModule) error {
				return m.GetProjectTokens(ctx, "projectId_1", body, nil)
			}},
		{"UpdateProjectName", http.MethodPatch, "/v2/account/projects/projectId_1/settings/name", "account",
			func(ctx context.Context, m *AccountModule) error {
				return m.UpdateProjectName(ctx, "projectId_1", body, nil)
			}},
		{"UpdateProjectDescription", http.MethodPatch, "/v2/account/projects/projectId_1/settings/description", "account",
			func(ctx context.Context, m *AccountModule) error {
				return m.UpdateProjectDescription(ctx, "projectId_1", body, nil)
			}},
		{"UpdateProjectLogo", http.MethodPatch, "/v2/account/projects/projectId_1/settings/logo", "account",
			func(ctx context.Context, m *AccountModule) error {
				return m.UpdateProjectLogo(ctx, "projectId_1", body, nil)
			}},
		{"UpdateProjectIcon", http.MethodPatch, "/v2/account/projects/projectId_1/settings/icon", "account",
			func(ctx context.Context, m *AccountModule) error {
				return m.UpdateProjectIcon(ctx, "projectId_1", body, nil)
			}},
		{"UpdateProjectMainColor", http.MethodPatch, "/v2/account/projects/projectId_1/settings/main-color", "account",
			func(ctx context.Context, m *AccountModule) error {
				return m.UpdateProjectMainColor(ctx, "projectId_1", body, nil)
			}},
		{"UpdateProjectAccentColor", http.MethodPatch, "/v2/account/projects/projectId_1/settings/accent-color", "account",
			func(ctx context.Context, m *AccountModule) error {
				return m.UpdateProjectAccentColor(ctx, "projectId_1", body, nil)
			}},
		{"UpdateProjectAllowedOrigins", http.MethodPatch, "/v2/account/projects/projectId_1/settings/origins", "account",
			func(ctx context.Context, m *AccountModule) error {
				return m.UpdateProjectAllowedOrigins(ctx, "projectId_1", body, nil)
			}},
		{"UpdateProjectUrl", http.MethodPatch, "/v2/account/projects/projectId_1/settings/url", "account",
			func(ctx context.Context, m *AccountModule) error {
				return m.UpdateProjectUrl(ctx, "projectId_1", body, nil)
			}},
		{"UpdateProjectLanguages", http.MethodPatch, "/v2/account/projects/projectId_1/settings/languages", "account",
			func(ctx context.Context, m *AccountModule) error {
				return m.UpdateProjectLanguages(ctx, "projectId_1", body, nil)
			}},
		{"CheckProjectLanguages", http.MethodPost, "/v2/account/projects/projectId_1/settings/languages/check", "account",
			func(ctx context.Context, m *AccountModule) error {
				return m.CheckProjectLanguages(ctx, "projectId_1", body, nil)
			}},
		{"UpdateProjectDefaultLanguage", http.MethodPatch, "/v2/account/projects/projectId_1/settings/default-language", "account",
			func(ctx context.Context, m *AccountModule) error {
				return m.UpdateProjectDefaultLanguage(ctx, "projectId_1", body, nil)
			}},
		{"UpdateProjectRegions", http.MethodPatch, "/v2/account/projects/projectId_1/settings/regions", "account",
			func(ctx context.Context, m *AccountModule) error {
				return m.UpdateProjectRegions(ctx, "projectId_1", body, nil)
			}},
		{"UpdateProjectAdminUrl", http.MethodPatch, "/v2/account/projects/projectId_1/settings/admin-url", "account",
			func(ctx context.Context, m *AccountModule) error {
				return m.UpdateProjectAdminUrl(ctx, "projectId_1", body, nil)
			}},
		{"UpdateProjectLegalDocuments", http.MethodPatch, "/v2/account/projects/projectId_1/settings/legal", "account",
			func(ctx context.Context, m *AccountModule) error {
				return m.UpdateProjectLegalDocuments(ctx, "projectId_1", body, nil)
			}},
		{"UpdateProjectExposeLegal", http.MethodPatch, "/v2/account/projects/projectId_1/settings/legal/expose", "account",
			func(ctx context.Context, m *AccountModule) error {
				return m.UpdateProjectExposeLegal(ctx, "projectId_1", body, nil)
			}},
		{"UpdateProjectExposeBrand", http.MethodPatch, "/v2/account/projects/projectId_1/settings/brand/expose", "account",
			func(ctx context.Context, m *AccountModule) error {
				return m.UpdateProjectExposeBrand(ctx, "projectId_1", body, nil)
			}},
		{"UpdateProjectExposeAuth", http.MethodPatch, "/v2/account/projects/projectId_1/settings/auth/expose", "account",
			func(ctx context.Context, m *AccountModule) error {
				return m.UpdateProjectExposeAuth(ctx, "projectId_1", body, nil)
			}},
		{"AssignAdminPortalServiceUser", http.MethodPut, "/v2/account/projects/projectId_1/settings/admin-portal/service-user", "account",
			func(ctx context.Context, m *AccountModule) error {
				return m.AssignAdminPortalServiceUser(ctx, "projectId_1", body, nil)
			}},
		{"GetAdminPortalStructure", http.MethodGet, "/v2/account/projects/projectId_1/admin-portal/structure", "project",
			func(ctx context.Context, m *AccountModule) error {
				return m.GetAdminPortalStructure(ctx, "projectId_1", body, nil)
			}},
		{"CreateAiServiceUser", http.MethodPost, "/v2/account/ai/service-users", "account",
			func(ctx context.Context, m *AccountModule) error {
				return m.CreateAiServiceUser(ctx, body, nil)
			}},
		{"ListAiServiceUsers", http.MethodGet, "/v2/account/ai/service-users", "account",
			func(ctx context.Context, m *AccountModule) error {
				return m.ListAiServiceUsers(ctx, body, nil)
			}},
		{"DeleteAiServiceUser", http.MethodDelete, "/v2/account/ai/service-users/aisu_1", "account",
			func(ctx context.Context, m *AccountModule) error {
				return m.DeleteAiServiceUser(ctx, "aisu_1", body, nil)
			}},
		{"RotateAiServiceUserKey", http.MethodPost, "/v2/account/ai/service-users/aisu_1/keys", "account",
			func(ctx context.Context, m *AccountModule) error {
				return m.RotateAiServiceUserKey(ctx, "aisu_1", body, nil)
			}},
		{"RevokeAiServiceUserKey", http.MethodDelete, "/v2/account/ai/service-users/aisu_1/keys/key_1", "account",
			func(ctx context.Context, m *AccountModule) error {
				return m.RevokeAiServiceUserKey(ctx, "aisu_1", "key_1", body, nil)
			}},
		{"Mcp", http.MethodPost, "/v2/account/mcp", "project",
			func(ctx context.Context, m *AccountModule) error {
				return m.Mcp(ctx, "sess_1", body, nil)
			}},
		{"McpOpenStream", http.MethodGet, "/v2/account/mcp", "project",
			func(ctx context.Context, m *AccountModule) error {
				return m.McpOpenStream(ctx, "sess_1", nil)
			}},
		{"McpEndSession", http.MethodDelete, "/v2/account/mcp", "project",
			func(ctx context.Context, m *AccountModule) error {
				return m.McpEndSession(ctx, "sess_1", nil)
			}},
	}
}

func newProjectTestModule(baseURL string, accountID string) *AccountModule {
	return &AccountModule{t: transport.New(&transport.Config{
		ProjectID:  "proj_1",
		AccountID:  accountID,
		APIKey:     "key_1",
		BaseURLAPI: baseURL,
		BaseURLHub: baseURL,
		APIVersion: "v2",
		HubVersion: "v2",
		Timeout:    5 * time.Second,
		Env:        "PROD",
	}, nil)}
}

func TestProjectEndpointsHitTheExpectedRoute(t *testing.T) {
	for _, c := range projectCases() {
		t.Run(c.name, func(t *testing.T) {
			var gotMethod, gotPath, gotAuth, gotProject, gotAccount string
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotMethod = r.Method
				gotPath = r.URL.Path
				gotAuth = r.Header.Get("Authorization")
				gotProject = r.Header.Get("X-CM-ProjectId")
				gotAccount = r.Header.Get("X-CM-AccountId")
				w.WriteHeader(http.StatusNoContent)
			}))
			defer srv.Close()

			if err := c.call(context.Background(), newProjectTestModule(srv.URL, "acct_1")); err != nil {
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
			if gotAccount != "acct_1" {
				t.Errorf("%s account header: got %q", c.name, gotAccount)
			}
		})
	}
}

// An account-scoped call is refused before it leaves the process when the
// client has no AccountID; a project-scoped call still goes out.
func TestProjectEndpointsAccountScope(t *testing.T) {
	for _, c := range projectCases() {
		t.Run(c.name, func(t *testing.T) {
			hit := false
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				hit = true
				w.WriteHeader(http.StatusNoContent)
			}))
			defer srv.Close()

			err := c.call(context.Background(), newProjectTestModule(srv.URL, ""))
			if c.scope == "account" {
				if err == nil || hit {
					t.Errorf("%s: want a local refusal without AccountID, got err=%v hit=%v", c.name, err, hit)
				}
				return
			}
			if err != nil || !hit {
				t.Errorf("%s: want the call to go out, got err=%v hit=%v", c.name, err, hit)
			}
		})
	}
}

// The MCP calls name their session in the Mcp-Session-Id header, accept JSON
// or an SSE stream, and POST the JSON-RPC message as the body.
func TestMcpSendsSessionHeaderAndJsonRpcBody(t *testing.T) {
	var gotSession, gotAccept string
	var gotBody map[string]any
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotSession = r.Header.Get("Mcp-Session-Id")
		gotAccept = r.Header.Get("Accept")
		raw, _ := io.ReadAll(r.Body)
		_ = json.Unmarshal(raw, &gotBody)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":{"tools":[]}}`))
	}))
	defer srv.Close()

	msg := map[string]any{"jsonrpc": "2.0", "id": 1, "method": "tools/list"}
	var out map[string]any
	if err := newProjectTestModule(srv.URL, "").Mcp(context.Background(), "sess_1", msg, &out); err != nil {
		t.Fatalf("Mcp: unexpected error: %v", err)
	}
	if gotSession != "sess_1" {
		t.Errorf("Mcp-Session-Id: got %q want %q", gotSession, "sess_1")
	}
	if gotAccept != "application/json, text/event-stream" {
		t.Errorf("Accept: got %q", gotAccept)
	}
	if gotBody["method"] != "tools/list" || gotBody["jsonrpc"] != "2.0" {
		t.Errorf("JSON-RPC body: got %v", gotBody)
	}
	if _, ok := out["result"]; !ok {
		t.Errorf("answer not decoded: got %v", out)
	}

	// No session yet ("initialize"): the header is left out.
	gotSession = "unset"
	if err := newProjectTestModule(srv.URL, "").Mcp(context.Background(), "", msg, nil); err != nil {
		t.Fatalf("Mcp without session: unexpected error: %v", err)
	}
	if gotSession != "" {
		t.Errorf("Mcp-Session-Id without session: got %q want empty", gotSession)
	}
}

// The Project-module surface tested here is 34 routes; a changed count
// means a route arrived or left untested.
func TestProjectSurfaceSize(t *testing.T) {
	const want = 34
	if got := len(projectCases()); got != want {
		t.Errorf("project endpoint count: got %d want %d", got, want)
	}
}
