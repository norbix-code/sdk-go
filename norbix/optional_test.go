package norbix

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	hubdtos "github.com/norbix-code/sdk-go/v2/norbix/hub/dtos"
)

// asMap turns a generated request type into the map a module method takes,
// the way an app does it: through its JSON tags.
func asMap(t *testing.T, v any) map[string]any {
	t.Helper()
	raw, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var m map[string]any
	if err := json.Unmarshal(raw, &m); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	return m
}

func TestBoolAndFloat64ReturnTheirValue(t *testing.T) {
	if b := Bool(false); b == nil || *b {
		t.Errorf("Bool(false): got %v", b)
	}
	if b := Bool(true); b == nil || !*b {
		t.Errorf("Bool(true): got %v", b)
	}
	if f := Float64(0); f == nil || *f != 0 {
		t.Errorf("Float64(0): got %v", f)
	}
}

// L32: SaveLogSettings' fields are C# bool? — absent means "unchanged". A
// false must reach the gateway, and a field left nil must not be sent.
func TestSaveLogSettingsSendsFalseAndLeavesNilOut(t *testing.T) {
	c, got := newFilesClient(t, `{}`)
	req := asMap(t, hubdtos.SaveLogSettings{
		SkipCloudDashboardLogs: Bool(false),
		SkipHttpBodyMeta:       Bool(true),
	})

	if err := c.Hub.Logs.SaveLogSettings(context.Background(), req, nil); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectCall(t, got, http.MethodPost, "/v2/logs/settings")
	if v, ok := got.body["skipCloudDashboardLogs"]; !ok || v != false {
		t.Errorf(`body: want "skipCloudDashboardLogs":false, got %v`, got.body)
	}
	if v, ok := got.body["skipHttpBodyMeta"]; !ok || v != true {
		t.Errorf(`body: want "skipHttpBodyMeta":true, got %v`, got.body)
	}
	if _, ok := got.body["aiChatLoggingEnabled"]; ok {
		t.Errorf("body: aiChatLoggingEnabled was nil and must not be sent, got %v", got.body)
	}
}

// IsEnabled on an integration request is a required C# bool: it is always
// sent, so an integration can be saved switched off.
func TestSaveLoggingIntegrationSendsIsEnabledFalse(t *testing.T) {
	c, got := newFilesClient(t, `{"id":"int_mongo"}`)
	req := map[string]any{
		"integration": hubdtos.MongoDbLoggingIntegrationRequest{
			LoggingIntegrationRequest: hubdtos.LoggingIntegrationRequest{
				IntegrationName: "App logs",
				IsEnabled:       false,
			},
			Provider:         hubdtos.LoggingProviderMongoDB,
			ConnectionString: "mongodb://example.test",
		},
	}

	if err := c.Hub.Logs.SaveLoggingIntegration(context.Background(), req, nil); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectCall(t, got, http.MethodPost, "/v2/logs/integrations")
	integration, _ := got.body["integration"].(map[string]any)
	if v, ok := integration["isEnabled"]; !ok || v != false {
		t.Errorf(`body: want "integration":{"isEnabled":false,...}, got %v`, got.body)
	}
}

// A nullable number (C# int?) is a *float64: 0 is sent, nil is not.
func TestOptionalNumberSendsZeroAndLeavesNilOut(t *testing.T) {
	c, got := newFilesClient(t, `{"list":[]}`)

	req := asMap(t, hubdtos.GetPaymentsWebhookLog{IntegrationId: "int_stripe", Limit: Float64(0)})
	if err := c.Hub.Payments.GetPaymentsWebhookLog(context.Background(), req, nil); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.query != "integrationId=int_stripe&limit=0" {
		t.Errorf("query with Limit 0: got %q", got.query)
	}

	req = asMap(t, hubdtos.GetPaymentsWebhookLog{IntegrationId: "int_stripe"})
	if err := c.Hub.Payments.GetPaymentsWebhookLog(context.Background(), req, nil); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.query != "integrationId=int_stripe" {
		t.Errorf("query with Limit nil: got %q", got.query)
	}
}
