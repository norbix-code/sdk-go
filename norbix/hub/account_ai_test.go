package hub

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/norbix-code/sdk-go/v2/norbix/internal/transport"
)

// Every wave-3 projectAi route and the module method that calls it. Each case
// asserts the verb and the fully-resolved path, so a wrong path or a swapped
// verb fails here instead of at runtime. Nothing leaves the process: every
// call goes to a local test server, no provider is ever contacted.

type projectAiCase struct {
	name string
	verb string
	path string
	call func(ctx context.Context, m *AccountModule) error
}

func projectAiCases() []projectAiCase {
	body := map[string]any{"probe": "value"}

	return []projectAiCase{
		{"GetProjectAiSettings", http.MethodGet, "/v2/account/projects/projectId_1/ai/settings",
			func(ctx context.Context, m *AccountModule) error {
				return m.GetProjectAiSettings(ctx, "projectId_1", body, nil)
			}},
		{"UpdateProjectAiSettings", http.MethodPut, "/v2/account/projects/projectId_1/ai/settings",
			func(ctx context.Context, m *AccountModule) error {
				return m.UpdateProjectAiSettings(ctx, "projectId_1", body, nil)
			}},
		{"CreateProjectAiAssistant", http.MethodPost, "/v2/account/projects/projectId_1/ai/assistants",
			func(ctx context.Context, m *AccountModule) error {
				return m.CreateProjectAiAssistant(ctx, "projectId_1", body, nil)
			}},
		{"UpdateProjectAiAssistant", http.MethodPut, "/v2/account/projects/projectId_1/ai/assistants/assistantId_1",
			func(ctx context.Context, m *AccountModule) error {
				return m.UpdateProjectAiAssistant(ctx, "projectId_1", "assistantId_1", body, nil)
			}},
		{"DeleteProjectAiAssistant", http.MethodDelete, "/v2/account/projects/projectId_1/ai/assistants/assistantId_1",
			func(ctx context.Context, m *AccountModule) error {
				return m.DeleteProjectAiAssistant(ctx, "projectId_1", "assistantId_1", body, nil)
			}},
		{"GetProjectAiUsage", http.MethodGet, "/v2/account/projects/projectId_1/ai/usage",
			func(ctx context.Context, m *AccountModule) error {
				return m.GetProjectAiUsage(ctx, "projectId_1", body, nil)
			}},
		{"SetAdminPortalEnabled", http.MethodPut, "/v2/account/projects/projectId_1/admin-portal/enabled",
			func(ctx context.Context, m *AccountModule) error {
				return m.SetAdminPortalEnabled(ctx, "projectId_1", body, nil)
			}},
	}
}

func newProjectaiTestModule(baseURL string) *AccountModule {
	return &AccountModule{t: transport.New(&transport.Config{
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

func TestProjectaiEndpointsHitTheExpectedRoute(t *testing.T) {
	for _, c := range projectAiCases() {
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

			if err := c.call(context.Background(), newProjectaiTestModule(srv.URL)); err != nil {
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

// The wave-3 projectAi surface is 7 routes; a changed count means a route arrived
// or left untested.
func TestProjectaiSurfaceSize(t *testing.T) {
	const want = 7
	if got := len(projectAiCases()); got != want {
		t.Errorf("projectAi endpoint count: got %d want %d", got, want)
	}
}
