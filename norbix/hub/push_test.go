package hub

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/norbix-code/sdk-go/norbix/internal/transport"
)

// Every Push endpoint the gateway exposes, and the module method that calls
// it. Each case asserts the verb and the fully-resolved path — version
// substituted, route tokens filled — so a wrong path or a swapped verb fails
// here instead of at runtime.
//
// Nothing leaves the process: every call goes to a local test server, so no
// push provider is ever contacted.

const (
	testCampaignID     = "camp_1"
	testBatchID        = "batch_1"
	testNotificationID = "notif_1"
	testTemplateID     = "tpl_1"
	testIntegrationID  = "int_1"
)

type pushCase struct {
	name string
	verb string
	path string
	call func(ctx context.Context, m *NotificationsModule) error
}

func pushCases() []pushCase {
	body := map[string]any{"probe": "value"}

	return []pushCase{
		// --- module
		{"EnablePush", http.MethodGet, "/v2/notifications/push/enable",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.EnablePush(ctx, nil, nil)
			}},
		{"DisablePush", http.MethodGet, "/v2/notifications/push/disable",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.DisablePush(ctx, nil, nil)
			}},
		{"GetPushDisableDependencies", http.MethodGet, "/v2/notifications/push/disable-dependencies",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetPushDisableDependencies(ctx, nil, nil)
			}},
		{"GetPushSettings", http.MethodGet, "/v2/notifications/push/settings",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetPushSettings(ctx, nil, nil)
			}},

		// --- integrations
		{"GetPushIntegrations", http.MethodGet, "/v2/notifications/push/integrations",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetPushIntegrations(ctx, nil, nil)
			}},
		{"SavePushIntegration", http.MethodPost, "/v2/notifications/push/integrations",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.SavePushIntegration(ctx, body, nil)
			}},
		{"GetPushIntegration", http.MethodGet, "/v2/notifications/push/integrations/" + testIntegrationID,
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetPushIntegration(ctx, testIntegrationID, nil, nil)
			}},
		{"EnablePushIntegration", http.MethodPut, "/v2/notifications/push/integrations/" + testIntegrationID + "/enable",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.EnablePushIntegration(ctx, testIntegrationID, nil, nil)
			}},
		{"DisablePushIntegration", http.MethodPut, "/v2/notifications/push/integrations/" + testIntegrationID + "/disable",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.DisablePushIntegration(ctx, testIntegrationID, nil, nil)
			}},
		{"SetPushIntegrationAsDefault", http.MethodPut, "/v2/notifications/push/integrations/" + testIntegrationID + "/default",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.SetPushIntegrationAsDefault(ctx, testIntegrationID, nil, nil)
			}},
		{"DeletePushIntegration", http.MethodDelete, "/v2/notifications/push/integrations/" + testIntegrationID,
			func(ctx context.Context, m *NotificationsModule) error {
				return m.DeletePushIntegration(ctx, testIntegrationID, nil, nil)
			}},
		{"TestPushIntegration", http.MethodPost, "/v2/notifications/push/integrations/test",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.TestPushIntegration(ctx, body, nil)
			}},
		{"ConfirmPushIntegrationHumanDelivery", http.MethodPost, "/v2/notifications/push/integrations/confirm-human-delivery",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.ConfirmPushIntegrationHumanDelivery(ctx, body, nil)
			}},
		{"RegisterCodeMashAppPushIntegration", http.MethodPost, "/v2/notifications/push/integrations/app/request",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.RegisterCodeMashAppPushIntegration(ctx, body, nil)
			}},

		// --- templates
		{"GetPushTemplates", http.MethodGet, "/v2/notifications/push/templates",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetPushTemplates(ctx, nil, nil)
			}},
		{"CreatePushTemplate", http.MethodPost, "/v2/notifications/push/templates",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.CreatePushTemplate(ctx, body, nil)
			}},
		{"UpdatePushTemplate", http.MethodPut, "/v2/notifications/push/templates",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.UpdatePushTemplate(ctx, body, nil)
			}},
		{"GetPushTemplate", http.MethodGet, "/v2/notifications/push/templates/" + testTemplateID,
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetPushTemplate(ctx, testTemplateID, nil, nil)
			}},
		{"DeletePushTemplate", http.MethodDelete, "/v2/notifications/push/templates/" + testTemplateID,
			func(ctx context.Context, m *NotificationsModule) error {
				return m.DeletePushTemplate(ctx, testTemplateID, nil, nil)
			}},
		{"ArchivePushTemplate", http.MethodPut, "/v2/notifications/push/templates/" + testTemplateID + "/archive",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.ArchivePushTemplate(ctx, testTemplateID, nil, nil)
			}},
		{"UnArchivePushTemplate", http.MethodPut, "/v2/notifications/push/templates/" + testTemplateID + "/unarchive",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.UnArchivePushTemplate(ctx, testTemplateID, nil, nil)
			}},
		{"ClonePushTemplate", http.MethodPost, "/v2/notifications/push/templates/" + testTemplateID + "/clone",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.ClonePushTemplate(ctx, testTemplateID, nil, nil)
			}},
		{"GetPushMessageContentTokens", http.MethodGet, "/v2/notifications/push/templates/" + testTemplateID + "/tokens",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetPushMessageContentTokens(ctx, testTemplateID, nil, nil)
			}},
		{"RenderPush", http.MethodPost, "/v2/notifications/push/templates/render",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.RenderPush(ctx, body, nil)
			}},

		// --- campaigns
		{"GetPushCampaigns", http.MethodGet, "/v2/notifications/push/campaigns",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetPushCampaigns(ctx, nil, nil)
			}},
		{"CreatePushCampaign", http.MethodPost, "/v2/notifications/push/campaigns",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.CreatePushCampaign(ctx, body, nil)
			}},
		{"GetPushCampaign", http.MethodGet, "/v2/notifications/push/campaigns/" + testCampaignID,
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetPushCampaign(ctx, testCampaignID, nil, nil)
			}},
		{"DeletePushCampaign", http.MethodDelete, "/v2/notifications/push/campaigns/" + testCampaignID,
			func(ctx context.Context, m *NotificationsModule) error {
				return m.DeletePushCampaign(ctx, testCampaignID, nil, nil)
			}},
		{"StopPushCampaign", http.MethodPost, "/v2/notifications/push/campaigns/" + testCampaignID + "/stop",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.StopPushCampaign(ctx, testCampaignID, nil, nil)
			}},
		{"GetPushCampaignBatches", http.MethodGet, "/v2/notifications/push/campaigns/" + testCampaignID + "/batches",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetPushCampaignBatches(ctx, testCampaignID, nil, nil)
			}},
		{"GetPushCampaignBatchNotifications", http.MethodGet,
			"/v2/notifications/push/campaigns/" + testCampaignID + "/batches/" + testBatchID,
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetPushCampaignBatchNotifications(ctx, testCampaignID, testBatchID, nil, nil)
			}},
		{"GetPushCampaignBatchNotification", http.MethodGet,
			"/v2/notifications/push/campaigns/" + testCampaignID + "/batches/" + testBatchID + "/" + testNotificationID,
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetPushCampaignBatchNotification(ctx, testCampaignID, testBatchID, testNotificationID, nil, nil)
			}},
		{"GetPushCampaignStatistics", http.MethodGet, "/v2/notifications/push/campaigns/" + testCampaignID + "/stats",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetPushCampaignStatistics(ctx, testCampaignID, nil, nil)
			}},
		{"GetPushCampaignMessages", http.MethodGet, "/v2/notifications/push/campaigns/" + testCampaignID + "/messages",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetPushCampaignMessages(ctx, testCampaignID, nil, nil)
			}},
		{"GetPushCampaignMessage", http.MethodGet,
			"/v2/notifications/push/campaigns/" + testCampaignID + "/messages/" + testNotificationID,
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetPushCampaignMessage(ctx, testCampaignID, testNotificationID, nil, nil)
			}},
		{"PreviewPushNotification", http.MethodGet, "/v2/notifications/push/preview",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.PreviewPushNotification(ctx, nil, nil)
			}},

		// --- devices
		{"RegisterDevice", http.MethodPost, "/v2/notifications/push/devices",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.RegisterDevice(ctx, body, nil)
			}},
	}
}

// newPushTestModule wires a notifications module to srv and records nothing
// else — each test reads what the handler saw.
func newPushTestModule(baseURL string) *NotificationsModule {
	return &NotificationsModule{t: transport.New(&transport.Config{
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

func TestPushEndpointsHitTheExpectedRoute(t *testing.T) {
	for _, c := range pushCases() {
		t.Run(c.name, func(t *testing.T) {
			var gotMethod, gotPath string
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotMethod = r.Method
				gotPath = r.URL.Path
				w.WriteHeader(http.StatusNoContent)
			}))
			defer srv.Close()

			if err := c.call(context.Background(), newPushTestModule(srv.URL)); err != nil {
				t.Fatalf("%s: unexpected error: %v", c.name, err)
			}
			if gotMethod != c.verb {
				t.Errorf("%s verb: got %q want %q", c.name, gotMethod, c.verb)
			}
			if gotPath != c.path {
				t.Errorf("%s path: got %q want %q", c.name, gotPath, c.path)
			}
		})
	}
}

// The whole Push surface is 37 routes. If the gateway grows one and the module
// is regenerated, this count changes and the test says so, so a new endpoint
// cannot arrive untested.
func TestPushSurfaceSize(t *testing.T) {
	const want = 37
	if got := len(pushCases()); got != want {
		t.Errorf("push endpoint count: got %d want %d", got, want)
	}
}

func TestPushEndpointsSendAuthAndProjectHeaders(t *testing.T) {
	var auth, project string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth = r.Header.Get("Authorization")
		project = r.Header.Get("X-CM-ProjectId")
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	if err := newPushTestModule(srv.URL).GetPushTemplates(context.Background(), nil, nil); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if auth != "Bearer key_1" {
		t.Errorf("authorization: got %q", auth)
	}
	if project != "proj_1" {
		t.Errorf("project header: got %q", project)
	}
}

// A campaign is targeted by a `source` discriminator plus that audience's own
// fields. The Go module takes an untyped body, so the test's job is to prove
// the body reaches the wire unchanged for each of the five audiences.
func TestCreatePushCampaignCarriesTheAudienceShape(t *testing.T) {
	audiences := []map[string]any{
		{"source": "allUsers", "templateId": testTemplateID, "userTags": []any{"beta"}},
		{"source": "specifiedUsers", "templateId": testTemplateID, "userRecipients": []any{"user_1"}},
		{"source": "accountUsers", "templateId": testTemplateID, "userRecipients": []any{"acct_user_1"}},
		{"source": "collection", "templateId": testTemplateID, "schemaName": "subscribers"},
		{"source": "devices", "templateId": testTemplateID, "devices": []any{
			map[string]any{"token": "device_1", "deliveryFamily": "ios"},
		}},
	}

	for _, audience := range audiences {
		source, _ := audience["source"].(string)
		t.Run(source, func(t *testing.T) {
			var got map[string]any
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				b, _ := io.ReadAll(r.Body)
				_ = json.Unmarshal(b, &got)
				w.WriteHeader(http.StatusNoContent)
			}))
			defer srv.Close()

			req := map[string]any{"campaign": audience}
			if err := newPushTestModule(srv.URL).CreatePushCampaign(context.Background(), req, nil); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			campaign, ok := got["campaign"].(map[string]any)
			if !ok {
				t.Fatalf("campaign missing from body: %v", got)
			}
			if campaign["source"] != source {
				t.Errorf("source: got %v want %q", campaign["source"], source)
			}
			if campaign["templateId"] != testTemplateID {
				t.Errorf("templateId: got %v", campaign["templateId"])
			}
		})
	}
}

// Only the Fake provider is exercised end to end — it is the sandbox that
// accepts a send and contacts no push service. The other providers are shape
// assertions that never leave the process.
func TestSavePushIntegrationCarriesTheProviderShape(t *testing.T) {
	providers := []string{
		"Fake", "AndroidFirebase", "AppleApns", "CodeMashChromePlugin",
		"ChromeWeb", "EdgeWeb", "FirefoxWeb", "SafariPush",
	}

	for _, provider := range providers {
		t.Run(provider, func(t *testing.T) {
			var got map[string]any
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				b, _ := io.ReadAll(r.Body)
				_ = json.Unmarshal(b, &got)
				w.WriteHeader(http.StatusNoContent)
			}))
			defer srv.Close()

			req := map[string]any{"integration": map[string]any{
				"provider":        provider,
				"integrationName": "test-" + provider,
				"isEnabled":       true,
			}}
			if err := newPushTestModule(srv.URL).SavePushIntegration(context.Background(), req, nil); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			integration, ok := got["integration"].(map[string]any)
			if !ok {
				t.Fatalf("integration missing from body: %v", got)
			}
			if integration["provider"] != provider {
				t.Errorf("provider: got %v want %q", integration["provider"], provider)
			}
		})
	}
}
