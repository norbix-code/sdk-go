package hub

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/norbix-code/sdk-go/v2/norbix/hub/dtos"
)

// Every trigger route of every host — Database (schema), Files, Membership,
// Payments, AI — and the six shared trigger routes, each with the verb, the
// fully-resolved path, the query and where the request goes (query for a
// read, body for a save). Source of truth: the gateway's `[Route]` attributes
// (the shared routes are missing from the endpoint manifest — emitter bug).
// Nothing leaves the process: every call goes to a local test server.

type triggerRouteCase struct {
	name  string
	verb  string
	path  string
	query url.Values     // nil: no query string expected
	body  map[string]any // nil: no body expected
	call  func(ctx context.Context, ns *Namespace) error
}

const trgID = "trg_1"

func saveBody(kind string) map[string]any {
	return map[string]any{"trigger": map[string]any{"type": kind, "name": "On event", "action": map[string]any{"type": "Email", "integrationId": "int_1", "templateId": "tpl_1"}}}
}

func triggerRouteCases() []triggerRouteCase {
	return []triggerRouteCase{
		// Database (schema) triggers — the reference host.
		{"GetSchemaTriggers", http.MethodGet, "/v2/database/schemas/triggers", url.Values{"schemaId": {"sch_1"}}, nil,
			func(ctx context.Context, ns *Namespace) error {
				return ns.Database.GetSchemaTriggers(ctx, map[string]any{"schemaId": "sch_1"}, nil)
			}},
		{"GetSchemaTrigger", http.MethodGet, "/v2/database/schemas/triggers/" + trgID, nil, nil,
			func(ctx context.Context, ns *Namespace) error {
				return ns.Database.GetSchemaTrigger(ctx, trgID, nil, nil)
			}},
		{"SaveSchemaTrigger", http.MethodPost, "/v2/database/schemas/triggers", nil, saveBody("Schema"),
			func(ctx context.Context, ns *Namespace) error {
				return ns.Database.SaveSchemaTrigger(ctx, saveBody("Schema"), nil)
			}},
		{"EnableSchemaTrigger", http.MethodPatch, "/v2/database/schemas/triggers/" + trgID + "/enable", nil, nil,
			func(ctx context.Context, ns *Namespace) error {
				return ns.Database.EnableSchemaTrigger(ctx, trgID, nil, nil)
			}},
		{"DisableSchemaTrigger", http.MethodPatch, "/v2/database/schemas/triggers/" + trgID + "/disable", nil, nil,
			func(ctx context.Context, ns *Namespace) error {
				return ns.Database.DisableSchemaTrigger(ctx, trgID, nil, nil)
			}},
		{"DeleteSchemaTrigger", http.MethodDelete, "/v2/database/schemas/triggers/" + trgID, nil, nil,
			func(ctx context.Context, ns *Namespace) error {
				return ns.Database.DeleteSchemaTrigger(ctx, trgID, nil, nil)
			}},

		// Files triggers.
		{"GetFilesTriggers", http.MethodGet, "/v2/files/triggers", url.Values{"pageSize": {"10"}}, nil,
			func(ctx context.Context, ns *Namespace) error {
				return ns.Files.GetFilesTriggers(ctx, map[string]any{"pageSize": 10}, nil)
			}},
		{"GetFilesTrigger", http.MethodGet, "/v2/files/triggers/" + trgID, nil, nil,
			func(ctx context.Context, ns *Namespace) error { return ns.Files.GetFilesTrigger(ctx, trgID, nil, nil) }},
		{"SaveFilesTrigger", http.MethodPost, "/v2/files/triggers", nil, saveBody("Files"),
			func(ctx context.Context, ns *Namespace) error {
				return ns.Files.SaveFilesTrigger(ctx, saveBody("Files"), nil)
			}},
		{"EnableFilesTrigger", http.MethodPatch, "/v2/files/triggers/" + trgID + "/enable", nil, nil,
			func(ctx context.Context, ns *Namespace) error {
				return ns.Files.EnableFilesTrigger(ctx, trgID, nil, nil)
			}},
		{"DisableFilesTrigger", http.MethodPatch, "/v2/files/triggers/" + trgID + "/disable", nil, nil,
			func(ctx context.Context, ns *Namespace) error {
				return ns.Files.DisableFilesTrigger(ctx, trgID, nil, nil)
			}},
		{"DeleteFilesTrigger", http.MethodDelete, "/v2/files/triggers/" + trgID, nil, nil,
			func(ctx context.Context, ns *Namespace) error {
				return ns.Files.DeleteFilesTrigger(ctx, trgID, nil, nil)
			}},

		// Membership triggers.
		{"GetMembershipTriggers", http.MethodGet, "/v2/membership/triggers", nil, nil,
			func(ctx context.Context, ns *Namespace) error {
				return ns.Membership.GetMembershipTriggers(ctx, nil, nil)
			}},
		{"GetMembershipTrigger", http.MethodGet, "/v2/membership/triggers/" + trgID, nil, nil,
			func(ctx context.Context, ns *Namespace) error {
				return ns.Membership.GetMembershipTrigger(ctx, trgID, nil, nil)
			}},
		{"SaveMembershipTrigger", http.MethodPost, "/v2/membership/triggers", nil, saveBody("Membership"),
			func(ctx context.Context, ns *Namespace) error {
				return ns.Membership.SaveMembershipTrigger(ctx, saveBody("Membership"), nil)
			}},
		{"EnableMembershipTrigger", http.MethodPatch, "/v2/membership/triggers/" + trgID + "/enable", nil, nil,
			func(ctx context.Context, ns *Namespace) error {
				return ns.Membership.EnableMembershipTrigger(ctx, trgID, nil, nil)
			}},
		{"DisableMembershipTrigger", http.MethodPatch, "/v2/membership/triggers/" + trgID + "/disable", nil, nil,
			func(ctx context.Context, ns *Namespace) error {
				return ns.Membership.DisableMembershipTrigger(ctx, trgID, nil, nil)
			}},
		{"DeleteMembershipTrigger", http.MethodDelete, "/v2/membership/triggers/" + trgID, nil, nil,
			func(ctx context.Context, ns *Namespace) error {
				return ns.Membership.DeleteMembershipTrigger(ctx, trgID, nil, nil)
			}},

		// Payments triggers.
		{"GetPaymentsTriggers", http.MethodGet, "/v2/payments/triggers", nil, nil,
			func(ctx context.Context, ns *Namespace) error { return ns.Payments.GetPaymentsTriggers(ctx, nil, nil) }},
		{"GetPaymentsTrigger", http.MethodGet, "/v2/payments/triggers/" + trgID, nil, nil,
			func(ctx context.Context, ns *Namespace) error {
				return ns.Payments.GetPaymentsTrigger(ctx, trgID, nil, nil)
			}},
		{"SavePaymentsTrigger", http.MethodPost, "/v2/payments/triggers", nil, saveBody("Payments"),
			func(ctx context.Context, ns *Namespace) error {
				return ns.Payments.SavePaymentsTrigger(ctx, saveBody("Payments"), nil)
			}},
		{"EnablePaymentsTrigger", http.MethodPatch, "/v2/payments/triggers/" + trgID + "/enable", nil, nil,
			func(ctx context.Context, ns *Namespace) error {
				return ns.Payments.EnablePaymentsTrigger(ctx, trgID, nil, nil)
			}},
		{"DisablePaymentsTrigger", http.MethodPatch, "/v2/payments/triggers/" + trgID + "/disable", nil, nil,
			func(ctx context.Context, ns *Namespace) error {
				return ns.Payments.DisablePaymentsTrigger(ctx, trgID, nil, nil)
			}},
		{"DeletePaymentsTrigger", http.MethodDelete, "/v2/payments/triggers/" + trgID, nil, nil,
			func(ctx context.Context, ns *Namespace) error {
				return ns.Payments.DeletePaymentsTrigger(ctx, trgID, nil, nil)
			}},

		// AI triggers (credits and quota events).
		{"GetAiProjectTriggers", http.MethodGet, "/v2/ai/triggers", nil, nil,
			func(ctx context.Context, ns *Namespace) error { return ns.Ai.GetAiProjectTriggers(ctx, nil, nil) }},
		{"GetAiProjectTrigger", http.MethodGet, "/v2/ai/triggers/" + trgID, nil, nil,
			func(ctx context.Context, ns *Namespace) error { return ns.Ai.GetAiProjectTrigger(ctx, trgID, nil, nil) }},
		{"SaveAiProjectTrigger", http.MethodPost, "/v2/ai/triggers", nil, saveBody("Ai"),
			func(ctx context.Context, ns *Namespace) error {
				return ns.Ai.SaveAiProjectTrigger(ctx, saveBody("Ai"), nil)
			}},
		{"EnableAiProjectTrigger", http.MethodPatch, "/v2/ai/triggers/" + trgID + "/enable", nil, nil,
			func(ctx context.Context, ns *Namespace) error {
				return ns.Ai.EnableAiProjectTrigger(ctx, trgID, nil, nil)
			}},
		{"DisableAiProjectTrigger", http.MethodPatch, "/v2/ai/triggers/" + trgID + "/disable", nil, nil,
			func(ctx context.Context, ns *Namespace) error {
				return ns.Ai.DisableAiProjectTrigger(ctx, trgID, nil, nil)
			}},
		{"DeleteAiProjectTrigger", http.MethodDelete, "/v2/ai/triggers/" + trgID, nil, nil,
			func(ctx context.Context, ns *Namespace) error {
				return ns.Ai.DeleteAiProjectTrigger(ctx, trgID, nil, nil)
			}},

		// Shared: whatever the host.
		{"GetTriggerRuns", http.MethodGet, "/v2/triggers/runs", url.Values{"triggerType": {"Files"}, "triggerId": {trgID}, "status": {"failed"}}, nil,
			func(ctx context.Context, ns *Namespace) error {
				return ns.Triggers.GetTriggerRuns(ctx, map[string]any{"triggerType": "Files", "triggerId": trgID, "status": "failed"}, nil)
			}},
		{"GetTriggerRun", http.MethodGet, "/v2/triggers/runs/trr_1", url.Values{"triggerType": {"Ai"}}, nil,
			func(ctx context.Context, ns *Namespace) error {
				return ns.Triggers.GetTriggerRun(ctx, "trr_1", map[string]any{"triggerType": "Ai"}, nil)
			}},
		{"GetTriggersNeedingAttention", http.MethodGet, "/v2/triggers/attention", url.Values{"triggerType": {"Membership"}}, nil,
			func(ctx context.Context, ns *Namespace) error {
				return ns.Triggers.GetTriggersNeedingAttention(ctx, map[string]any{"triggerType": "Membership"}, nil)
			}},
		{"GetTriggerDefinition", http.MethodGet, "/v2/triggers/" + trgID + "/definition", url.Values{"triggerType": {"Schema"}, "schemaId": {"sch_1"}}, nil,
			func(ctx context.Context, ns *Namespace) error {
				return ns.Triggers.GetTriggerDefinition(ctx, trgID, map[string]any{"triggerType": "Schema", "schemaId": "sch_1"}, nil)
			}},
		{"TestTriggerPreExecuteCode", http.MethodPost, "/v2/triggers/pre-execute-code/test", nil,
			map[string]any{"code": "return New.status == \"published\";", "triggerType": "Schema", "newRecord": "{\"status\":\"published\"}", "oldRecord": "{\"status\":\"draft\"}", "eventKind": "Updated"},
			func(ctx context.Context, ns *Namespace) error {
				return ns.Triggers.TestTriggerPreExecuteCode(ctx, map[string]any{"code": "return New.status == \"published\";", "triggerType": "Schema", "newRecord": "{\"status\":\"published\"}", "oldRecord": "{\"status\":\"draft\"}", "eventKind": "Updated"}, nil)
			}},
		{"CheckTriggerExecutionIdentity", http.MethodPost, "/v2/triggers/execution-identity/check", nil,
			map[string]any{"triggerType": "Payments", "executionIdentityId": "usr_1", "action": map[string]any{"type": "Email", "integrationId": "int_1"}},
			func(ctx context.Context, ns *Namespace) error {
				return ns.Triggers.CheckTriggerExecutionIdentity(ctx, map[string]any{"triggerType": "Payments", "executionIdentityId": "usr_1", "action": map[string]any{"type": "Email", "integrationId": "int_1"}}, nil)
			}},
	}
}

func TestEveryTriggerRouteOfEveryHostHitsTheExpectedRoute(t *testing.T) {
	for _, c := range triggerRouteCases() {
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

			if err := c.call(context.Background(), NewNamespace(newTriggersTestTransport(srv.URL))); err != nil {
				t.Fatalf("%s: unexpected error: %v", c.name, err)
			}
			if gotMethod != c.verb {
				t.Errorf("verb: got %q want %q", gotMethod, c.verb)
			}
			if gotPath != c.path {
				t.Errorf("path: got %q want %q", gotPath, c.path)
			}
			if gotAuth != "Bearer key_1" {
				t.Errorf("authorization: got %q", gotAuth)
			}
			if gotProject != "proj_1" {
				t.Errorf("project header: got %q", gotProject)
			}

			// Query: exactly the expected parameters, nothing else.
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

			// Body: none for reads, deletes and toggles; the request otherwise.
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
			assertFields(t, got, c.body)
		})
	}
}

// Five hosts × six routes + six shared routes. A host or a shared route added
// to the gateway changes this count, so it cannot arrive untested.
func TestTriggerRouteSurfaceSize(t *testing.T) {
	const want = 5*6 + 6
	if got := len(triggerRouteCases()); got != want {
		t.Errorf("trigger route count: got %d want %d", got, want)
	}
}

// A Code action (run a cloud function) and a Workflow action (start a
// workflow), typed, reach the wire under "trigger.action" of every host's save
// with the gateway's field names: functionRef / invocationType / config, and
// workflowRef / payload. No runAsUserId: the trigger's executionIdentityId
// says who runs.
func TestTriggerSaveSendsCodeAndWorkflowActions(t *testing.T) {
	code := dtos.TriggerActionCodeDto{
		TriggerActionDto: dtos.TriggerActionDto{Type: dtos.TriggerActionTypeCode},
		FunctionRef:      &dtos.FunctionRefDto{IntegrationId: "int_lambda", FunctionId: "notify-sales", Qualifier: "live"},
		InvocationType:   dtos.CodeInvocationTypeAsync,
		Config:           `{"source":"go"}`,
	}
	workflow := dtos.TriggerActionWorkflowDto{
		TriggerActionDto: dtos.TriggerActionDto{Type: dtos.TriggerActionTypeWorkflow},
		WorkflowRef:      "wf_onboard",
		Payload:          map[string]string{"source": "go"},
	}
	wantCode := map[string]any{
		"type":           "Code",
		"functionRef":    map[string]any{"integrationId": "int_lambda", "functionId": "notify-sales", "qualifier": "live"},
		"invocationType": "Async",
		"config":         `{"source":"go"}`,
	}
	wantWorkflow := map[string]any{
		"type":        "Workflow",
		"workflowRef": "wf_onboard",
		"payload":     map[string]any{"source": "go"},
	}

	saves := []struct {
		host string
		kind string
		path string
		save func(ctx context.Context, ns *Namespace, body map[string]any) error
	}{
		{"Database", "Schema", "/v2/database/schemas/triggers", func(ctx context.Context, ns *Namespace, b map[string]any) error {
			return ns.Database.SaveSchemaTrigger(ctx, b, nil)
		}},
		{"Files", "Files", "/v2/files/triggers", func(ctx context.Context, ns *Namespace, b map[string]any) error {
			return ns.Files.SaveFilesTrigger(ctx, b, nil)
		}},
		{"Membership", "Membership", "/v2/membership/triggers", func(ctx context.Context, ns *Namespace, b map[string]any) error {
			return ns.Membership.SaveMembershipTrigger(ctx, b, nil)
		}},
		{"Payments", "Payments", "/v2/payments/triggers", func(ctx context.Context, ns *Namespace, b map[string]any) error {
			return ns.Payments.SavePaymentsTrigger(ctx, b, nil)
		}},
		{"Ai", "Ai", "/v2/ai/triggers", func(ctx context.Context, ns *Namespace, b map[string]any) error {
			return ns.Ai.SaveAiProjectTrigger(ctx, b, nil)
		}},
	}
	actions := []struct {
		name   string
		action any
		want   map[string]any
	}{
		{"Code", code, wantCode},
		{"Workflow", workflow, wantWorkflow},
	}
	for _, s := range saves {
		for _, a := range actions {
			t.Run(s.host+"/"+a.name, func(t *testing.T) {
				var gotPath string
				var got map[string]any
				srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					gotPath = r.URL.Path
					b, _ := io.ReadAll(r.Body)
					_ = json.Unmarshal(b, &got)
					w.WriteHeader(http.StatusNoContent)
				}))
				defer srv.Close()

				body := map[string]any{"trigger": map[string]any{"type": s.kind, "name": "run " + a.name, "executionIdentityId": "usr_1", "action": toMap(t, a.action)}}
				if err := s.save(context.Background(), NewNamespace(newTriggersTestTransport(srv.URL)), body); err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if gotPath != s.path {
					t.Errorf("path: got %q want %q", gotPath, s.path)
				}
				trigger, ok := got["trigger"].(map[string]any)
				if !ok {
					t.Fatalf("trigger missing from body: %v", got)
				}
				sent, ok := trigger["action"].(map[string]any)
				if !ok {
					t.Fatalf("action missing from trigger: %v", trigger)
				}
				if _, has := sent["runAsUserId"]; has {
					t.Errorf("action carries runAsUserId: %v", sent)
				}
				if trigger["executionIdentityId"] != "usr_1" {
					t.Errorf("executionIdentityId: got %v", trigger["executionIdentityId"])
				}
				assertFields(t, sent, a.want)
			})
		}
	}
}
