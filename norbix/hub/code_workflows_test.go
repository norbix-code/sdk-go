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

// The 26 workflow (beta) routes of the Code module (gateway campaign
// audit/workflows, docs/architecture/Code.Workflows.md §7) and the module
// method that calls each one. Every case asserts the verb, the fully-resolved
// path (version and ids substituted) and where the parameters land: the query
// string for GET / DELETE, the JSON body for POST / PUT. Nothing leaves the
// process.

const (
	testWorkflowID    = "wf_6Qp1dW2kRm8NsT4vYc0Hb"
	testWorkflowRunID = "wfr_4hT2b9QkLm3XcT1yRfGh7a"
	testResumeToken   = "signed.resume.token"
)

func workflowCases() []codeCase {
	definition := map[string]any{
		"name":  "Order approval",
		"start": "approve",
		"steps": []any{map[string]any{"id": "approve", "kind": "wait.approval"}},
	}
	return []codeCase{
		{"GetCodeWorkflows", http.MethodGet, "/v2/code/workflows",
			url.Values{"isEnabled": {"true"}, "pageSize": {"20"}, "startingAfter": {"wf_1"}}, nil,
			func(ctx context.Context, m *CodeModule) error {
				return m.GetCodeWorkflows(ctx, map[string]any{"isEnabled": true, "pageSize": 20, "startingAfter": "wf_1"}, nil)
			}},
		{"SaveCodeWorkflow", http.MethodPut, "/v2/code/workflows", nil, map[string]any{"definition": definition},
			func(ctx context.Context, m *CodeModule) error {
				return m.SaveCodeWorkflow(ctx, map[string]any{"definition": definition}, nil)
			}},
		{"ValidateCodeWorkflow", http.MethodPost, "/v2/code/workflows/validate", nil, map[string]any{"definition": definition},
			func(ctx context.Context, m *CodeModule) error {
				return m.ValidateCodeWorkflow(ctx, map[string]any{"definition": definition}, nil)
			}},
		{"GetCodeWorkflowSchema", http.MethodGet, "/v2/code/workflows/schema", nil, nil,
			func(ctx context.Context, m *CodeModule) error {
				return m.GetCodeWorkflowSchema(ctx, nil, nil)
			}},
		{"EnableCodeWorkflows", http.MethodPut, "/v2/code/workflows/enable", nil, map[string]any{},
			func(ctx context.Context, m *CodeModule) error {
				return m.EnableCodeWorkflows(ctx, map[string]any{}, nil)
			}},
		{"DisableCodeWorkflows", http.MethodPut, "/v2/code/workflows/disable", nil, map[string]any{},
			func(ctx context.Context, m *CodeModule) error {
				return m.DisableCodeWorkflows(ctx, map[string]any{}, nil)
			}},
		{"GetCodeWorkflow", http.MethodGet, "/v2/code/workflows/" + testWorkflowID,
			url.Values{"pinnedVersion": {"3"}}, nil,
			func(ctx context.Context, m *CodeModule) error {
				return m.GetCodeWorkflow(ctx, testWorkflowID, map[string]any{"pinnedVersion": 3}, nil)
			}},
		{"DeleteCodeWorkflow", http.MethodDelete, "/v2/code/workflows/" + testWorkflowID, nil, nil,
			func(ctx context.Context, m *CodeModule) error {
				return m.DeleteCodeWorkflow(ctx, testWorkflowID, nil, nil)
			}},
		{"PublishCodeWorkflow", http.MethodPost, "/v2/code/workflows/" + testWorkflowID + "/publish", nil, map[string]any{"confirmed": true},
			func(ctx context.Context, m *CodeModule) error {
				return m.PublishCodeWorkflow(ctx, testWorkflowID, map[string]any{"confirmed": true}, nil)
			}},
		{"GetCodeWorkflowDraft", http.MethodGet, "/v2/code/workflows/" + testWorkflowID + "/draft", nil, nil,
			func(ctx context.Context, m *CodeModule) error {
				return m.GetCodeWorkflowDraft(ctx, testWorkflowID, nil, nil)
			}},
		{"DiscardCodeWorkflowDraft", http.MethodDelete, "/v2/code/workflows/" + testWorkflowID + "/draft", nil, nil,
			func(ctx context.Context, m *CodeModule) error {
				return m.DiscardCodeWorkflowDraft(ctx, testWorkflowID, nil, nil)
			}},
		{"GetCodeWorkflowVersions", http.MethodGet, "/v2/code/workflows/" + testWorkflowID + "/versions", nil, nil,
			func(ctx context.Context, m *CodeModule) error {
				return m.GetCodeWorkflowVersions(ctx, testWorkflowID, nil, nil)
			}},
		{"GetCodeWorkflowVersionDiff", http.MethodGet, "/v2/code/workflows/" + testWorkflowID + "/versions/diff",
			url.Values{"fromVersion": {"1"}, "toVersion": {"2"}}, nil,
			func(ctx context.Context, m *CodeModule) error {
				return m.GetCodeWorkflowVersionDiff(ctx, testWorkflowID, map[string]any{"fromVersion": 1, "toVersion": 2}, nil)
			}},
		{"EnableCodeWorkflow", http.MethodPut, "/v2/code/workflows/" + testWorkflowID + "/enable", nil, map[string]any{},
			func(ctx context.Context, m *CodeModule) error {
				return m.EnableCodeWorkflow(ctx, testWorkflowID, map[string]any{}, nil)
			}},
		{"DisableCodeWorkflow", http.MethodPut, "/v2/code/workflows/" + testWorkflowID + "/disable", nil, map[string]any{},
			func(ctx context.Context, m *CodeModule) error {
				return m.DisableCodeWorkflow(ctx, testWorkflowID, map[string]any{}, nil)
			}},
		{"InvokeCodeWorkflow", http.MethodPost, "/v2/code/workflows/" + testWorkflowID + "/invoke", nil,
			map[string]any{"payload": map[string]any{"orderId": "ord_1"}, "runVersion": "draft", "waitSeconds": 10},
			func(ctx context.Context, m *CodeModule) error {
				return m.InvokeCodeWorkflow(ctx, testWorkflowID, map[string]any{"payload": map[string]any{"orderId": "ord_1"}, "waitSeconds": 10, "runVersion": "draft"}, nil)
			}},
		{"GetCodeWorkflowRuns", http.MethodGet, "/v2/code/workflows/" + testWorkflowID + "/runs",
			url.Values{"from": {"2026-10-01T00:00:00Z"}, "pageSize": {"50"}, "status": {"failed"}}, nil,
			func(ctx context.Context, m *CodeModule) error {
				return m.GetCodeWorkflowRuns(ctx, testWorkflowID, map[string]any{"status": "failed", "from": "2026-10-01T00:00:00Z", "pageSize": 50}, nil)
			}},
		{"ResumeCodeWorkflowRunByKey", http.MethodPost, "/v2/code/workflows/" + testWorkflowID + "/resume", nil,
			map[string]any{"decision": "approved", "key": "order-1042", "payload": map[string]any{"note": "ok"}, "stepId": "approve"},
			func(ctx context.Context, m *CodeModule) error {
				return m.ResumeCodeWorkflowRunByKey(ctx, testWorkflowID, map[string]any{"key": "order-1042", "stepId": "approve", "decision": "approved", "payload": map[string]any{"note": "ok"}}, nil)
			}},
		{"GetCodeWorkflowWaitingRuns", http.MethodGet, "/v2/code/workflows/runs",
			url.Values{"assignee": {"me"}, "key": {"order-1042"}, "status": {"waiting"}, "waitingKind": {"approval"}, "workflowId": {testWorkflowID}}, nil,
			func(ctx context.Context, m *CodeModule) error {
				return m.GetCodeWorkflowWaitingRuns(ctx, map[string]any{"status": "waiting", "waitingKind": "approval", "assignee": "me", "workflowId": testWorkflowID, "key": "order-1042"}, nil)
			}},
		{"GetCodeWorkflowRun", http.MethodGet, "/v2/code/workflows/runs/" + testWorkflowRunID, nil, nil,
			func(ctx context.Context, m *CodeModule) error {
				return m.GetCodeWorkflowRun(ctx, testWorkflowRunID, nil, nil)
			}},
		{"GetCodeWorkflowRunLogs", http.MethodGet, "/v2/code/workflows/runs/" + testWorkflowRunID + "/logs",
			url.Values{"level": {"error"}, "pageSize": {"100"}, "stepId": {"approve"}}, nil,
			func(ctx context.Context, m *CodeModule) error {
				return m.GetCodeWorkflowRunLogs(ctx, testWorkflowRunID, map[string]any{"stepId": "approve", "level": "error", "pageSize": 100}, nil)
			}},
		{"ResumeCodeWorkflowRun", http.MethodPost, "/v2/code/workflows/runs/" + testWorkflowRunID + "/resume", nil,
			map[string]any{"decision": "rejected", "payload": map[string]any{"reason": "over budget"}, "stepId": "approve"},
			func(ctx context.Context, m *CodeModule) error {
				return m.ResumeCodeWorkflowRun(ctx, testWorkflowRunID, map[string]any{"stepId": "approve", "decision": "rejected", "payload": map[string]any{"reason": "over budget"}}, nil)
			}},
		{"RetryCodeWorkflowRunStep", http.MethodPost, "/v2/code/workflows/runs/" + testWorkflowRunID + "/retry", nil,
			map[string]any{"stepId": "charge"},
			func(ctx context.Context, m *CodeModule) error {
				return m.RetryCodeWorkflowRunStep(ctx, testWorkflowRunID, map[string]any{"stepId": "charge"}, nil)
			}},
		{"CancelCodeWorkflowRun", http.MethodPost, "/v2/code/workflows/runs/" + testWorkflowRunID + "/cancel", nil,
			map[string]any{"reason": "duplicate order"},
			func(ctx context.Context, m *CodeModule) error {
				return m.CancelCodeWorkflowRun(ctx, testWorkflowRunID, map[string]any{"reason": "duplicate order"}, nil)
			}},
		{"OpenCodeWorkflowResumeLink", http.MethodGet, "/v2/code/workflows/resume/" + testResumeToken, nil, nil,
			func(ctx context.Context, m *CodeModule) error {
				return m.OpenCodeWorkflowResumeLink(ctx, testResumeToken, nil, nil)
			}},
		{"UseCodeWorkflowResumeLink", http.MethodPost, "/v2/code/workflows/resume/" + testResumeToken, nil, nil,
			func(ctx context.Context, m *CodeModule) error {
				return m.UseCodeWorkflowResumeLink(ctx, testResumeToken, nil, nil)
			}},
	}
}

func TestCodeWorkflowEndpointsHitTheExpectedRoute(t *testing.T) {
	for _, c := range workflowCases() {
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

func TestCodeWorkflowCaseCount(t *testing.T) {
	const want = 26
	if got := len(workflowCases()); got != want {
		t.Errorf("code/workflows case count: got %d want %d", got, want)
	}
}

// The signed resume link needs no session: the token in the path is the key.
// A client with no API key and no bearer token still sends both link routes,
// with no Authorization header; a client that has a key still sends it.
func TestCodeWorkflowResumeLinkNeedsNoSignIn(t *testing.T) {
	auths := []struct {
		name     string
		apiKey   string
		wantAuth string
	}{
		{"no credentials", "", ""},
		{"api key", "key_1", "Bearer key_1"},
	}
	calls := []struct {
		name string
		verb string
		call func(ctx context.Context, m *CodeModule) error
	}{
		{"OpenCodeWorkflowResumeLink", http.MethodGet, func(ctx context.Context, m *CodeModule) error {
			return m.OpenCodeWorkflowResumeLink(ctx, testResumeToken, nil, nil)
		}},
		{"UseCodeWorkflowResumeLink", http.MethodPost, func(ctx context.Context, m *CodeModule) error {
			return m.UseCodeWorkflowResumeLink(ctx, testResumeToken, nil, nil)
		}},
	}
	for _, a := range auths {
		for _, c := range calls {
			t.Run(a.name+"/"+c.name, func(t *testing.T) {
				var gotMethod, gotPath, gotAuth string
				var sawAuth bool
				srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					gotMethod, gotPath = r.Method, r.URL.Path
					gotAuth = r.Header.Get("Authorization")
					_, sawAuth = r.Header["Authorization"]
					w.WriteHeader(http.StatusNoContent)
				}))
				defer srv.Close()

				m := &CodeModule{t: transport.New(&transport.Config{
					APIKey:     a.apiKey,
					BaseURLAPI: srv.URL,
					BaseURLHub: srv.URL,
					APIVersion: "v2",
					HubVersion: "v2",
					Timeout:    5 * time.Second,
				}, nil)}
				if err := c.call(context.Background(), m); err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if gotMethod != c.verb {
					t.Errorf("verb: got %q want %q", gotMethod, c.verb)
				}
				if gotPath != "/v2/code/workflows/resume/"+testResumeToken {
					t.Errorf("path: got %q", gotPath)
				}
				if a.wantAuth == "" && sawAuth {
					t.Errorf("authorization header sent without credentials: %q", gotAuth)
				}
				if a.wantAuth != "" && gotAuth != a.wantAuth {
					t.Errorf("authorization: got %q want %q", gotAuth, a.wantAuth)
				}
			})
		}
	}
}
