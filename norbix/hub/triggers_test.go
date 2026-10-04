package hub

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/norbix-code/sdk-go/v2/norbix/hub/dtos"
	"github.com/norbix-code/sdk-go/v2/norbix/internal/transport"
)

// The cross-module trigger endpoints, and the notification fields a trigger
// action carries on save. Nothing leaves the process: every call goes to a
// local test server.

func newTriggersTestTransport(baseURL string) *transport.Transport {
	return transport.New(&transport.Config{
		ProjectID:  "proj_1",
		AccountID:  "acct_1",
		APIKey:     "key_1",
		BaseURLAPI: baseURL,
		BaseURLHub: baseURL,
		APIVersion: "v2",
		HubVersion: "v2",
		Timeout:    5 * time.Second,
		Env:        "PROD",
	}, nil)
}

func TestGetTriggersNeedingAttentionHitsTheExpectedRoute(t *testing.T) {
	var gotMethod, gotPath, gotType, gotProject string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		gotType = r.URL.Query().Get("triggerType")
		gotProject = r.Header.Get("X-CM-ProjectId")
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"items":[{"triggerId":"trg_1","triggerType":"Schema","reason":"template misses language de","atUtc":"2026-10-04T10:00:00Z"}]}`)
	}))
	defer srv.Close()

	m := &TriggersModule{t: newTriggersTestTransport(srv.URL)}
	var out dtos.GetTriggersNeedingAttentionResponse
	if err := m.GetTriggersNeedingAttention(context.Background(), map[string]any{"triggerType": "Schema"}, &out); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet {
		t.Errorf("verb: got %q want GET", gotMethod)
	}
	if gotPath != "/v2/triggers/attention" {
		t.Errorf("path: got %q", gotPath)
	}
	if gotType != "Schema" {
		t.Errorf("triggerType query: got %q want Schema", gotType)
	}
	if gotProject != "proj_1" {
		t.Errorf("project header: got %q", gotProject)
	}
	if len(out.Items) != 1 || out.Items[0].TriggerId != "trg_1" || out.Items[0].TriggerType != dtos.TriggerTypeSchema || out.Items[0].Reason == "" {
		t.Errorf("decoded items: got %+v", out.Items)
	}
}

func TestHubNamespaceWiresTriggers(t *testing.T) {
	ns := NewNamespace(newTriggersTestTransport("http://127.0.0.1:1"))
	if ns.Triggers == nil || ns.Triggers.t == nil {
		t.Fatal("hub namespace has no Triggers module")
	}
}

// A notification trigger action names its provider (`integrationId`, required
// by the gateway) and may pin a template `language` and an `initiatorId`. The
// typed action DTOs flatten those into the `action` object of the save body.
func TestSaveTriggerSendsNotificationActionFields(t *testing.T) {
	base := dtos.TriggerActionDto{IntegrationId: testIntegrationID}
	actions := map[string]any{
		"Email": dtos.TriggerActionEmailDto{TriggerActionDto: withType(base, dtos.TriggerActionTypeEmail), TemplateId: testTemplateID, Language: "de", InitiatorId: "user_1"},
		"Push":  dtos.TriggerActionPushDto{TriggerActionDto: withType(base, dtos.TriggerActionTypePush), TemplateId: testTemplateID, Language: "de", InitiatorId: "user_1"},
		"Sms":   dtos.TriggerActionSmsDto{TriggerActionDto: withType(base, dtos.TriggerActionTypeSms), TemplateId: testTemplateID, Language: "de", InitiatorId: "user_1"},
	}
	for kind, action := range actions {
		t.Run(kind, func(t *testing.T) {
			var got map[string]any
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				b, _ := io.ReadAll(r.Body)
				_ = json.Unmarshal(b, &got)
				w.WriteHeader(http.StatusNoContent)
			}))
			defer srv.Close()

			m := &DatabaseModule{t: newTriggersTestTransport(srv.URL)}
			err := m.SaveSchemaTrigger(context.Background(), map[string]any{
				"type":   "Schema",
				"name":   "notify on insert",
				"action": toMap(t, action),
			}, nil)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			sent, ok := got["action"].(map[string]any)
			if !ok {
				t.Fatalf("action missing from body: %v", got)
			}
			assertFields(t, sent, map[string]any{
				"type":          kind,
				"integrationId": testIntegrationID,
				"templateId":    testTemplateID,
				"language":      "de",
				"initiatorId":   "user_1",
			})
		})
	}
}

func withType(a dtos.TriggerActionDto, kind dtos.TriggerActionType) dtos.TriggerActionDto {
	a.Type = kind
	return a
}

func toMap(t *testing.T, v any) map[string]any {
	t.Helper()
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	var m map[string]any
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatal(err)
	}
	return m
}
