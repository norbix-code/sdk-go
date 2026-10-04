package hub

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/norbix-code/sdk-go/v2/norbix/hub/dtos"
	"github.com/norbix-code/sdk-go/v2/norbix/internal/transport"
)

// Every Sms endpoint the gateway exposes, and the module method that calls
// it. Each case asserts the verb and the fully-resolved path — version
// substituted, route tokens filled — so a wrong path or a swapped verb fails
// here instead of at runtime. The 34 routes are the Sms rows of the gateway's
// endpoint manifest (typegen coverage/endpoints.hub.json).
//
// Nothing leaves the process: every call goes to a local test server, so no
// SMS provider is ever contacted.

type smsCase struct {
	name string
	verb string
	path string
	call func(ctx context.Context, m *NotificationsModule) error
}

func smsCases() []smsCase {
	body := map[string]any{"probe": "value"}

	return []smsCase{
		// --- module
		{"EnableSms", http.MethodGet, "/v2/notifications/sms/enable",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.EnableSms(ctx, nil, nil)
			}},
		{"DisableSms", http.MethodGet, "/v2/notifications/sms/disable",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.DisableSms(ctx, nil, nil)
			}},
		{"GetSmsDisableDependencies", http.MethodGet, "/v2/notifications/sms/disable-dependencies",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetSmsDisableDependencies(ctx, nil, nil)
			}},
		{"GetSmsSettings", http.MethodGet, "/v2/notifications/sms/settings",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetSmsSettings(ctx, nil, nil)
			}},
		{"PreviewSmsNotification", http.MethodGet, "/v2/notifications/sms/preview",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.PreviewSmsNotification(ctx, map[string]any{"hash": "signed-link"}, nil)
			}},

		// --- integrations
		{"GetSmsIntegrations", http.MethodGet, "/v2/notifications/sms/integrations",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetSmsIntegrations(ctx, nil, nil)
			}},
		{"SaveSmsIntegration", http.MethodPost, "/v2/notifications/sms/integrations",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.SaveSmsIntegration(ctx, body, nil)
			}},
		{"GetSmsIntegration", http.MethodGet, "/v2/notifications/sms/integrations/" + testIntegrationID,
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetSmsIntegration(ctx, testIntegrationID, nil, nil)
			}},
		{"EnableSmsIntegration", http.MethodPut, "/v2/notifications/sms/integrations/" + testIntegrationID + "/enable",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.EnableSmsIntegration(ctx, testIntegrationID, nil, nil)
			}},
		{"DisableSmsIntegration", http.MethodPut, "/v2/notifications/sms/integrations/" + testIntegrationID + "/disable",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.DisableSmsIntegration(ctx, testIntegrationID, nil, nil)
			}},
		{"SetSmsIntegrationAsDefault", http.MethodPut, "/v2/notifications/sms/integrations/" + testIntegrationID + "/default",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.SetSmsIntegrationAsDefault(ctx, testIntegrationID, nil, nil)
			}},
		{"DeleteSmsIntegration", http.MethodDelete, "/v2/notifications/sms/integrations/" + testIntegrationID,
			func(ctx context.Context, m *NotificationsModule) error {
				return m.DeleteSmsIntegration(ctx, testIntegrationID, nil, nil)
			}},
		{"TestSmsIntegration", http.MethodPost, "/v2/notifications/sms/integrations/test",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.TestSmsIntegration(ctx, body, nil)
			}},
		{"ConfirmSmsIntegrationHumanDelivery", http.MethodPost, "/v2/notifications/sms/integrations/confirm-human-delivery",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.ConfirmSmsIntegrationHumanDelivery(ctx, body, nil)
			}},

		// --- templates
		{"GetSmsTemplates", http.MethodGet, "/v2/notifications/sms/templates",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetSmsTemplates(ctx, nil, nil)
			}},
		{"CreateSmsTemplate", http.MethodPost, "/v2/notifications/sms/templates",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.CreateSmsTemplate(ctx, body, nil)
			}},
		{"UpdateSmsTemplate", http.MethodPut, "/v2/notifications/sms/templates",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.UpdateSmsTemplate(ctx, body, nil)
			}},
		{"GetSmsTemplate", http.MethodGet, "/v2/notifications/sms/templates/" + testTemplateID,
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetSmsTemplate(ctx, testTemplateID, nil, nil)
			}},
		{"DeleteSmsTemplate", http.MethodDelete, "/v2/notifications/sms/templates/" + testTemplateID,
			func(ctx context.Context, m *NotificationsModule) error {
				return m.DeleteSmsTemplate(ctx, testTemplateID, nil, nil)
			}},
		{"ArchiveSmsTemplate", http.MethodPut, "/v2/notifications/sms/templates/" + testTemplateID + "/archive",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.ArchiveSmsTemplate(ctx, testTemplateID, nil, nil)
			}},
		{"UnArchiveSmsTemplate", http.MethodPut, "/v2/notifications/sms/templates/" + testTemplateID + "/unarchive",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.UnArchiveSmsTemplate(ctx, testTemplateID, nil, nil)
			}},
		{"CloneSmsTemplate", http.MethodPost, "/v2/notifications/sms/templates/" + testTemplateID + "/clone",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.CloneSmsTemplate(ctx, testTemplateID, nil, nil)
			}},
		{"GetSmsMessageContentTokens", http.MethodGet, "/v2/notifications/sms/templates/" + testTemplateID + "/tokens",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetSmsMessageContentTokens(ctx, testTemplateID, nil, nil)
			}},
		{"RenderSms", http.MethodPost, "/v2/notifications/sms/templates/render",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.RenderSms(ctx, body, nil)
			}},

		// --- campaigns
		{"GetSmsCampaigns", http.MethodGet, "/v2/notifications/sms/campaigns",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetSmsCampaigns(ctx, nil, nil)
			}},
		{"CreateSmsCampaign", http.MethodPost, "/v2/notifications/sms/campaigns",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.CreateSmsCampaign(ctx, body, nil)
			}},
		{"GetSmsCampaign", http.MethodGet, "/v2/notifications/sms/campaigns/" + testCampaignID,
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetSmsCampaign(ctx, testCampaignID, nil, nil)
			}},
		{"DeleteSmsCampaign", http.MethodDelete, "/v2/notifications/sms/campaigns/" + testCampaignID,
			func(ctx context.Context, m *NotificationsModule) error {
				return m.DeleteSmsCampaign(ctx, testCampaignID, nil, nil)
			}},
		{"StopSmsCampaign", http.MethodPost, "/v2/notifications/sms/campaigns/" + testCampaignID + "/stop",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.StopSmsCampaign(ctx, testCampaignID, nil, nil)
			}},
		{"GetSmsCampaignBatches", http.MethodGet, "/v2/notifications/sms/campaigns/" + testCampaignID + "/batches",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetSmsCampaignBatches(ctx, testCampaignID, nil, nil)
			}},
		{"GetSmsCampaignBatchNotifications", http.MethodGet, "/v2/notifications/sms/campaigns/" + testCampaignID + "/batches/" + testBatchID,
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetSmsCampaignBatchNotifications(ctx, testCampaignID, testBatchID, nil, nil)
			}},
		{"GetSmsCampaignBatchNotification", http.MethodGet, "/v2/notifications/sms/campaigns/" + testCampaignID + "/batches/" + testBatchID + "/" + testNotificationID,
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetSmsCampaignBatchNotification(ctx, testCampaignID, testBatchID, testNotificationID, nil, nil)
			}},
		{"GetSmsCampaignStatistics", http.MethodGet, "/v2/notifications/sms/campaigns/" + testCampaignID + "/stats",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetSmsCampaignStatistics(ctx, testCampaignID, nil, nil)
			}},
		{"GetSmsCampaignMessages", http.MethodGet, "/v2/notifications/sms/campaigns/" + testCampaignID + "/messages",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetSmsCampaignMessages(ctx, testCampaignID, nil, nil)
			}},
	}
}

func newSmsTestModule(baseURL string) *NotificationsModule {
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

func TestSmsEndpointsHitTheExpectedRoute(t *testing.T) {
	for _, c := range smsCases() {
		t.Run(c.name, func(t *testing.T) {
			var gotMethod, gotPath string
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotMethod = r.Method
				gotPath = r.URL.Path
				w.WriteHeader(http.StatusNoContent)
			}))
			defer srv.Close()

			if err := c.call(context.Background(), newSmsTestModule(srv.URL)); err != nil {
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

// The whole Sms surface is 34 routes — the Sms rows of the gateway's endpoint
// manifest. If the gateway grows one and the module gains a method, this
// count changes and the test says so, so a new endpoint cannot arrive untested.
// (The old SmsRazorSyntaxCheck is not among them: the gateway never served
// POST /{version}/notifications/sms/templates/razor-syntax-check.)
func TestSmsSurfaceSize(t *testing.T) {
	const want = 34
	if got := len(smsCases()); got != want {
		t.Errorf("sms endpoint count: got %d want %d", got, want)
	}
}

func TestSmsEndpointsSendAuthAndProjectHeaders(t *testing.T) {
	var auth, project string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth = r.Header.Get("Authorization")
		project = r.Header.Get("X-CM-ProjectId")
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	if err := newSmsTestModule(srv.URL).GetSmsTemplates(context.Background(), nil, nil); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if auth != "Bearer key_1" {
		t.Errorf("authorization: got %q", auth)
	}
	if project != "proj_1" {
		t.Errorf("project header: got %q", project)
	}
}

// The four audience blocks of an Sms campaign, with the fields the gateway
// reads for each one. Source of truth: gateway Hub.Sms/Campaigns/Create.cs —
// `deliveryType` picks the block (AllUsers / SpecifiedUsers / Collection /
// PhoneNumbers) and each block repeats its own `recipientsSourceType`.
var smsCampaignTargets = []struct {
	deliveryType string
	block        string
	fields       map[string]any
}{
	{"AllUsers", "allUsers", map[string]any{
		"recipientsSourceType": "AllUsers",
		"rolesNames":           []any{"authenticated"},
		"userTags":             []any{"beta"},
	}},
	{"SpecifiedUsers", "specifiedUsers", map[string]any{
		"recipientsSourceType": "SpecifiedUsers",
		"recipients":           []any{"user_1"},
	}},
	{"Collection", "collection", map[string]any{
		"recipientsSourceType": "Collection",
		"schemaName":           "subscribers",
		"fields":               []any{"owner"},
		"fieldType":            "User",
	}},
	{"PhoneNumbers", "phoneNumbers", map[string]any{
		"recipientsSourceType": "PhoneNumbers",
		"phoneNumbers":         []any{"+37060000000"},
	}},
}

func TestCreateSmsCampaignSendsEachTargetShape(t *testing.T) {
	for _, target := range smsCampaignTargets {
		t.Run(target.deliveryType, func(t *testing.T) {
			got := sendAndCapture(t, func(m *NotificationsModule) error {
				return m.CreateSmsCampaign(context.Background(), map[string]any{
					"templateId":            testTemplateID,
					"integrationId":         testIntegrationID,
					"databaseIntegrationId": "db_int_1",
					"deliveryType":          target.deliveryType,
					target.block:            target.fields,
				}, nil)
			})

			// The discriminator and the template reach the wire as names, not
			// numbers: the gateway maps `deliveryType` by its enum name. The
			// SMS provider (`integrationId`, required by the gateway) sits at
			// the top level, next to and distinct from `databaseIntegrationId`.
			assertFields(t, got, map[string]any{
				"templateId":            testTemplateID,
				"integrationId":         testIntegrationID,
				"databaseIntegrationId": "db_int_1",
				"deliveryType":          target.deliveryType,
			})
			sent, ok := got[target.block].(map[string]any)
			if !ok {
				t.Fatalf("%s block missing from body: %v", target.block, got)
			}
			assertFields(t, sent, target.fields)
		})
	}
}

func TestSmsCampaignDeliveryTypesMatchTheGateway(t *testing.T) {
	// SmsCampaignRecipientsSourceTypes in the gateway has exactly these five.
	want := map[dtos.SmsCampaignRecipientsSourceTypes]bool{
		dtos.SmsCampaignRecipientsSourceTypesAllUsers:       true,
		dtos.SmsCampaignRecipientsSourceTypesSpecifiedUsers: true,
		dtos.SmsCampaignRecipientsSourceTypesAccountUsers:   true,
		dtos.SmsCampaignRecipientsSourceTypesPhoneNumbers:   true,
		dtos.SmsCampaignRecipientsSourceTypesCollection:     true,
	}
	for _, v := range []dtos.SmsCampaignRecipientsSourceTypes{"AllUsers", "SpecifiedUsers", "AccountUsers", "PhoneNumbers", "Collection"} {
		if !want[v] {
			t.Errorf("generated enum lacks %q", v)
		}
	}
	if len(want) != 5 {
		t.Errorf("delivery types: got %d want 5", len(want))
	}
}
