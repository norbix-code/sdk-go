package hub

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
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

// sendAndCapture posts through call and returns the JSON body the server saw.
func sendAndCapture(t *testing.T, call func(m *NotificationsModule) error) map[string]any {
	t.Helper()
	var got map[string]any
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		_ = json.Unmarshal(b, &got)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	if err := call(newPushTestModule(srv.URL)); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got == nil {
		t.Fatal("no JSON body reached the server")
	}
	return got
}

// assertFields fails for every key of want that is missing from got or holds
// a different value. Values are compared as JSON, so nested lists and objects
// count too.
func assertFields(t *testing.T, got, want map[string]any) {
	t.Helper()
	for key, w := range want {
		g, ok := got[key]
		if !ok {
			t.Errorf("field %q missing from body", key)
			continue
		}
		gb, _ := json.Marshal(g)
		wb, _ := json.Marshal(w)
		if string(gb) != string(wb) {
			t.Errorf("field %q: got %s want %s", key, gb, wb)
		}
	}
}

// The five campaign targets, with the fields the gateway reads for each one.
// Source of truth: gateway Hub.Push/Campaigns/Create.PushTo*.cs (the field
// names) and Create_.cs (PushCampaignRequestDtoJsonConverter picks the record
// from `source`, case-insensitive). Note `rolesNames` on allUsers but
// `roleNames` on collection — the gateway spells them differently.
var pushCampaignTargets = []struct {
	source string
	fields map[string]any
}{
	{"allUsers", map[string]any{
		"rolesNames": []any{"admin"},
		"userTags":   []any{"beta"},
	}},
	{"specifiedUsers", map[string]any{
		"userRecipients": []any{"user_1"},
	}},
	{"accountUsers", map[string]any{
		"userRecipients": []any{"acct_user_1"},
	}},
	{"collection", map[string]any{
		"schemaName": "subscribers",
		"fields":     []any{"owner"},
		"fieldType":  "User",
	}},
	{"devices", map[string]any{
		"devices": []any{map[string]any{"token": "device_1", "deliveryFamily": "Ios"}},
	}},
}

func TestCreatePushCampaignSendsEachTargetShape(t *testing.T) {
	for _, target := range pushCampaignTargets {
		t.Run(target.source, func(t *testing.T) {
			campaign := map[string]any{
				"source":     target.source,
				"templateId": testTemplateID,
			}
			for k, v := range target.fields {
				campaign[k] = v
			}

			got := sendAndCapture(t, func(m *NotificationsModule) error {
				return m.CreatePushCampaign(context.Background(),
					map[string]any{"campaign": campaign}, nil)
			})

			sent, ok := got["campaign"].(map[string]any)
			if !ok {
				t.Fatalf("campaign missing from body: %v", got)
			}
			// The discriminator has to reach the wire as the name: the gateway
			// reads it with GetString() and fails on a number.
			assertFields(t, sent, map[string]any{"source": target.source, "templateId": testTemplateID})
			assertFields(t, sent, target.fields)
		})
	}
}

func TestPushCampaignTargetsMatchTheGateway(t *testing.T) {
	// PushCampaignRecipientsSourceTypes in the gateway has exactly these five.
	want := map[string]bool{
		"allusers": true, "specifiedusers": true, "accountusers": true,
		"collection": true, "devices": true,
	}
	if len(pushCampaignTargets) != len(want) {
		t.Fatalf("targets: got %d want %d", len(pushCampaignTargets), len(want))
	}
	for _, target := range pushCampaignTargets {
		if !want[strings.ToLower(target.source)] {
			t.Errorf("source %q is not a gateway target", target.source)
		}
	}
}

// The eight providers the gateway's save converter accepts, with the fields
// each one validates. Source of truth: gateway Hub.Push/Integrations/Save_.cs
// (PushIntegrationRequestDtoJsonConverter — the switch arms) and
// Save.<Provider>.cs (the fields). The Chrome extension is `ChromePush`:
// `CodeMashChromePlugin` also parses as a PushProvider, but no switch arm
// takes it, so the gateway answers "Unsupported provider".
//
// Every value is a dummy. Only Fake is ever meant for a send; none of these
// leaves the process.
var pushIntegrationProviders = []struct {
	provider string
	fields   map[string]any
}{
	{"Fake", map[string]any{}},
	{"AndroidFirebase", map[string]any{
		"projectId":          "firebase-project",
		"clientEmail":        "push@firebase-project.iam.example.com",
		"serviceAccountJson": `{"type":"service_account"}`,
	}},
	{"AppleApns", map[string]any{
		"teamId":       "TEAM123456",
		"appBundleId":  "com.example.app",
		"keyId":        "KEY1234567",
		"privateKey":   "-----BEGIN PRIVATE KEY-----dummy",
		"isProduction": false,
	}},
	{"ChromePush", map[string]any{
		"extensionId":     "3f2504e0-4f89-11d3-9a0c-0305e82c3301",
		"vapidPublicKey":  "vapid-public",
		"vapidPrivateKey": "vapid-private",
		"subject":         "mailto:push@example.com",
	}},
	{"ChromeWeb", map[string]any{
		"vapidPublicKey":  "vapid-public",
		"vapidPrivateKey": "vapid-private",
		"subject":         "mailto:push@example.com",
	}},
	{"EdgeWeb", map[string]any{
		"vapidPublicKey":  "vapid-public",
		"vapidPrivateKey": "vapid-private",
	}},
	{"FirefoxWeb", map[string]any{
		"vapidPublicKey":  "vapid-public",
		"vapidPrivateKey": "vapid-private",
	}},
	{"SafariPush", map[string]any{
		"websitePushId":        "web.com.example",
		"certificateP12Base64": "ZHVtbXk=",
		"certificatePassword":  "dummy",
	}},
}

func TestSavePushIntegrationSendsEachProviderShape(t *testing.T) {
	for _, p := range pushIntegrationProviders {
		t.Run(p.provider, func(t *testing.T) {
			integration := map[string]any{
				"provider":        p.provider,
				"integrationName": "test-" + p.provider,
				"isEnabled":       true,
			}
			for k, v := range p.fields {
				integration[k] = v
			}

			got := sendAndCapture(t, func(m *NotificationsModule) error {
				return m.SavePushIntegration(context.Background(),
					map[string]any{"integration": integration}, nil)
			})

			sent, ok := got["integration"].(map[string]any)
			if !ok {
				t.Fatalf("integration missing from body: %v", got)
			}
			assertFields(t, sent, map[string]any{
				"provider":        p.provider,
				"integrationName": "test-" + p.provider,
				"isEnabled":       true,
			})
			assertFields(t, sent, p.fields)
		})
	}
}

func TestPushIntegrationProvidersMatchTheGateway(t *testing.T) {
	// The switch arms of PushIntegrationRequestDtoJsonConverter.Read.
	want := map[string]bool{
		"AppleApns": true, "AndroidFirebase": true, "SafariPush": true,
		"ChromeWeb": true, "FirefoxWeb": true, "EdgeWeb": true,
		"ChromePush": true, "Fake": true,
	}
	if len(pushIntegrationProviders) != len(want) {
		t.Fatalf("providers: got %d want %d", len(pushIntegrationProviders), len(want))
	}
	for _, p := range pushIntegrationProviders {
		if !want[p.provider] {
			t.Errorf("provider %q has no switch arm in the gateway", p.provider)
		}
	}
}

// RegisterDevice carries the device under `pushDeviceDto` plus the owning
// user. Source of truth: gateway Hub.Push/Devices/Create.cs and
// Contracts/Notifications/Push/Devices/PushDeviceDto.cs (`deviceOs` and
// `token` are required).
func TestRegisterDeviceSendsTheDeviceShape(t *testing.T) {
	device := map[string]any{"deviceOs": "iOS", "token": "device_1", "modelName": "iPhone"}

	got := sendAndCapture(t, func(m *NotificationsModule) error {
		return m.RegisterDevice(context.Background(),
			map[string]any{"pushDeviceDto": device, "userId": "user_1"}, nil)
	})

	assertFields(t, got, map[string]any{"pushDeviceDto": device, "userId": "user_1"})
}
