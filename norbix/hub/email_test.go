package hub

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/norbix-code/sdk-go/v2/norbix/internal/transport"
)

// Every Email endpoint the gateway exposes, and the module method that calls
// it: the project routes under /{version}/notifications/email (and the older
// /{version}/notifications/emails/campaigns/{campaignId}/messages), and the
// public link routes under /{version}/email/. Each case asserts the verb and
// the fully-resolved path — version substituted, route tokens filled — so a
// wrong path or a swapped verb fails here instead of at runtime. The routes
// and verbs are the gateway's (its endpoint manifest), not read from this SDK.
//
// Nothing leaves the process: every call goes to a local test server, so no
// email provider is ever contacted.

const testEmailID = "id_1"

type emailCase struct {
	name string
	verb string
	path string
	call func(ctx context.Context, n *NotificationsModule, e *EmailModule) error
}

func emailCases() []emailCase {
	body := map[string]any{"probe": "value"}

	return []emailCase{
		{"OneClickUnsubscribe", http.MethodPost, "/v2/email/one-click-unsubscribe",
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return e.OneClickUnsubscribe(ctx, body, nil)
			}},
		{"GetEmailPreferencesByLink", http.MethodGet, "/v2/email/preferences",
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return e.GetEmailPreferencesByLink(ctx, nil, nil)
			}},
		{"GetEmailCampaigns", http.MethodGet, "/v2/notifications/email/campaigns",
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.GetEmailCampaigns(ctx, nil, nil)
			}},
		{"CreateEmailCampaign", http.MethodPost, "/v2/notifications/email/campaigns",
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.CreateEmailCampaign(ctx, body, nil)
			}},
		{"DeleteEmailCampaign", http.MethodDelete, "/v2/notifications/email/campaigns/" + testEmailID,
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.DeleteEmailCampaign(ctx, testEmailID, nil, nil)
			}},
		{"StopEmailCampaign", http.MethodPost, "/v2/notifications/email/campaigns/" + testEmailID + "/stop",
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.StopEmailCampaign(ctx, testEmailID, body, nil)
			}},
		{"GetEmailCampaign", http.MethodGet, "/v2/notifications/email/campaigns/" + testEmailID,
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.GetEmailCampaign(ctx, testEmailID, nil, nil)
			}},
		{"GetEmailCampaignBatches", http.MethodGet, "/v2/notifications/email/campaigns/" + testEmailID + "/batches",
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.GetEmailCampaignBatches(ctx, testEmailID, nil, nil)
			}},
		{"GetEmailCampaignBatchNotifications", http.MethodGet, "/v2/notifications/email/campaigns/" + testEmailID + "/batches/" + testBatchID,
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.GetEmailCampaignBatchNotifications(ctx, testEmailID, testBatchID, nil, nil)
			}},
		{"GetEmailCampaignBatchNotification", http.MethodGet, "/v2/notifications/email/campaigns/" + testEmailID + "/batches/" + testBatchID + "/" + testNotificationID,
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.GetEmailCampaignBatchNotification(ctx, testEmailID, testBatchID, testNotificationID, nil, nil)
			}},
		{"GetEmailCampaignStatistics", http.MethodGet, "/v2/notifications/email/campaigns/" + testEmailID + "/stats",
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.GetEmailCampaignStatistics(ctx, testEmailID, nil, nil)
			}},
		{"DisableEmail", http.MethodGet, "/v2/notifications/email/disable",
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.DisableEmail(ctx, nil, nil)
			}},
		{"GetEmailDisableDependencies", http.MethodGet, "/v2/notifications/email/disable-dependencies",
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.GetEmailDisableDependencies(ctx, nil, nil)
			}},
		{"EnableEmail", http.MethodGet, "/v2/notifications/email/enable",
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.EnableEmail(ctx, nil, nil)
			}},
		{"GetEmailFooters", http.MethodGet, "/v2/notifications/email/footers",
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.GetEmailFooters(ctx, nil, nil)
			}},
		{"SaveEmailFooter", http.MethodPost, "/v2/notifications/email/footers",
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.SaveEmailFooter(ctx, body, nil)
			}},
		{"DeleteEmailFooter", http.MethodDelete, "/v2/notifications/email/footers/" + testEmailID,
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.DeleteEmailFooter(ctx, testEmailID, nil, nil)
			}},
		{"GetEmailFooter", http.MethodGet, "/v2/notifications/email/footers/" + testEmailID,
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.GetEmailFooter(ctx, testEmailID, nil, nil)
			}},
		{"GetEmailIntegrations", http.MethodGet, "/v2/notifications/email/integrations",
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.GetEmailIntegrations(ctx, nil, nil)
			}},
		{"SaveEmailIntegration", http.MethodPost, "/v2/notifications/email/integrations",
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.SaveEmailIntegration(ctx, body, nil)
			}},
		{"ConfirmEmailIntegrationHumanDelivery", http.MethodPost, "/v2/notifications/email/integrations/confirm-human-delivery",
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.ConfirmEmailIntegrationHumanDelivery(ctx, body, nil)
			}},
		{"CheckEmailIntegrationDomainHealth", http.MethodPost, "/v2/notifications/email/integrations/domain-health",
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.CheckEmailIntegrationDomainHealth(ctx, body, nil)
			}},
		{"TestEmailIntegration", http.MethodPost, "/v2/notifications/email/integrations/test",
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.TestEmailIntegration(ctx, body, nil)
			}},
		{"DeleteEmailIntegration", http.MethodDelete, "/v2/notifications/email/integrations/" + testEmailID,
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.DeleteEmailIntegration(ctx, testEmailID, nil, nil)
			}},
		{"SetEmailsIntegrationAsDefault", http.MethodPut, "/v2/notifications/email/integrations/" + testEmailID + "/default",
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.SetEmailsIntegrationAsDefault(ctx, testEmailID, body, nil)
			}},
		{"DisableEmailIntegration", http.MethodPut, "/v2/notifications/email/integrations/" + testEmailID + "/disable",
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.DisableEmailIntegration(ctx, testEmailID, body, nil)
			}},
		{"EnableEmailIntegration", http.MethodPut, "/v2/notifications/email/integrations/" + testEmailID + "/enable",
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.EnableEmailIntegration(ctx, testEmailID, body, nil)
			}},
		{"GetEmailIntegration", http.MethodGet, "/v2/notifications/email/integrations/" + testEmailID,
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.GetEmailIntegration(ctx, testEmailID, nil, nil)
			}},
		{"PreviewEmailNotification", http.MethodGet, "/v2/notifications/email/preview",
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.PreviewEmailNotification(ctx, nil, nil)
			}},
		{"GetEmailSettings", http.MethodGet, "/v2/notifications/email/settings",
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.GetEmailSettings(ctx, nil, nil)
			}},
		{"GetEmailSignatures", http.MethodGet, "/v2/notifications/email/signatures",
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.GetEmailSignatures(ctx, nil, nil)
			}},
		{"SaveEmailSignature", http.MethodPost, "/v2/notifications/email/signatures",
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.SaveEmailSignature(ctx, body, nil)
			}},
		{"DeleteEmailSignature", http.MethodDelete, "/v2/notifications/email/signatures/" + testEmailID,
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.DeleteEmailSignature(ctx, testEmailID, nil, nil)
			}},
		{"GetEmailSignature", http.MethodGet, "/v2/notifications/email/signatures/" + testEmailID,
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.GetEmailSignature(ctx, testEmailID, nil, nil)
			}},
		{"GetSystemEmailTemplates", http.MethodGet, "/v2/notifications/email/system-templates",
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.GetSystemEmailTemplates(ctx, nil, nil)
			}},
		{"GetSystemEmailTemplate", http.MethodGet, "/v2/notifications/email/system-templates/" + testEmailID,
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.GetSystemEmailTemplate(ctx, testEmailID, nil, nil)
			}},
		{"GetEmailTemplates", http.MethodGet, "/v2/notifications/email/templates",
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.GetEmailTemplates(ctx, nil, nil)
			}},
		{"CreateEmailTemplate", http.MethodPost, "/v2/notifications/email/templates",
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.CreateEmailTemplate(ctx, body, nil)
			}},
		{"UpdateEmailTemplate", http.MethodPut, "/v2/notifications/email/templates",
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.UpdateEmailTemplate(ctx, body, nil)
			}},
		{"AttachFileToTemplate", http.MethodPost, "/v2/notifications/email/templates/attachments",
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.AttachFileToTemplate(ctx, body, nil)
			}},
		{"GetMjml", http.MethodPost, "/v2/notifications/email/templates/mjml",
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.GetMjml(ctx, body, nil)
			}},
		{"DeleteEmailTemplate", http.MethodDelete, "/v2/notifications/email/templates/" + testEmailID,
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.DeleteEmailTemplate(ctx, testEmailID, nil, nil)
			}},
		{"ArchiveEmailTemplate", http.MethodPut, "/v2/notifications/email/templates/" + testEmailID + "/archive",
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.ArchiveEmailTemplate(ctx, testEmailID, body, nil)
			}},
		{"CloneEmailTemplate", http.MethodPost, "/v2/notifications/email/templates/" + testEmailID + "/clone",
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.CloneEmailTemplate(ctx, testEmailID, body, nil)
			}},
		{"UnArchiveEmailTemplate", http.MethodPut, "/v2/notifications/email/templates/" + testEmailID + "/unarchive",
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.UnArchiveEmailTemplate(ctx, testEmailID, body, nil)
			}},
		{"GetEmailTemplate", http.MethodGet, "/v2/notifications/email/templates/" + testEmailID,
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.GetEmailTemplate(ctx, testEmailID, nil, nil)
			}},
		{"GetEmailTemplateAvailableTokens", http.MethodGet, "/v2/notifications/email/templates/" + testEmailID + "/tokens",
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.GetEmailTemplateAvailableTokens(ctx, testEmailID, nil, nil)
			}},
		{"SaveEmailValidationIntegration", http.MethodPost, "/v2/notifications/email/validation/integrations",
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.SaveEmailValidationIntegration(ctx, body, nil)
			}},
		{"TestEmailValidationIntegration", http.MethodPost, "/v2/notifications/email/validation/integrations/test",
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.TestEmailValidationIntegration(ctx, body, nil)
			}},
		{"GetEmailCampaignMessages", http.MethodGet, "/v2/notifications/emails/campaigns/" + testCampaignID + "/messages",
			func(ctx context.Context, n *NotificationsModule, e *EmailModule) error {
				return n.GetEmailCampaignMessages(ctx, testCampaignID, nil, nil)
			}},
	}
}

func newEmailTestTransport(baseURL, apiKey string) *transport.Transport {
	return transport.New(&transport.Config{
		ProjectID:  "proj_1",
		AccountID:  "acct_1",
		APIKey:     apiKey,
		BaseURLAPI: baseURL,
		BaseURLHub: baseURL,
		APIVersion: "v2",
		HubVersion: "v2",
		Timeout:    5 * time.Second,
		Env:        "PROD",
	}, nil)
}

func TestEmailEndpointsHitTheExpectedRoute(t *testing.T) {
	for _, c := range emailCases() {
		t.Run(c.name, func(t *testing.T) {
			var gotMethod, gotPath string
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotMethod = r.Method
				gotPath = r.URL.Path
				w.WriteHeader(http.StatusNoContent)
			}))
			defer srv.Close()

			tr := newEmailTestTransport(srv.URL, "key_1")
			if err := c.call(context.Background(), &NotificationsModule{t: tr}, &EmailModule{t: tr}); err != nil {
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

// The whole Email surface is 50 routes (gateway endpoint manifest, Email
// campaign 2026-10: the duplicate campaigns/{campaignId}/messages/{notificationId}
// read was removed, GET /{version}/email/preferences was added). If the
// gateway grows or drops one, this count changes and the test says so.
func TestEmailSurfaceSize(t *testing.T) {
	const want = 50
	if got := len(emailCases()); got != want {
		t.Errorf("email endpoint count: got %d want %d", got, want)
	}
}

func TestEmailEndpointsSendAuthAndProjectHeaders(t *testing.T) {
	var auth, project string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth = r.Header.Get("Authorization")
		project = r.Header.Get("X-CM-ProjectId")
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	n := &NotificationsModule{t: newEmailTestTransport(srv.URL, "key_1")}
	if err := n.StopEmailCampaign(context.Background(), testCampaignID, nil, nil); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if auth != "Bearer key_1" {
		t.Errorf("authorization: got %q", auth)
	}
	if project != "proj_1" {
		t.Errorf("project header: got %q", project)
	}
}

// GET /{version}/email/preferences opens with the signed unsubscribe-link
// token alone, so a client with no API key and no bearer token must still
// send it — with no Authorization header. A client that has a key still
// sends it.
func TestEmailPreferencesByLinkNeedsNoSignIn(t *testing.T) {
	auths := []struct {
		name     string
		apiKey   string
		wantAuth string
	}{
		{"no credentials", "", ""},
		{"api key", "key_1", "Bearer key_1"},
	}
	for _, a := range auths {
		t.Run(a.name, func(t *testing.T) {
			var gotPath, gotToken, gotAuth string
			var sawAuth bool
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotPath = r.URL.Path
				gotToken = r.URL.Query().Get("token")
				gotAuth = r.Header.Get("Authorization")
				_, sawAuth = r.Header["Authorization"]
				w.WriteHeader(http.StatusNoContent)
			}))
			defer srv.Close()

			e := &EmailModule{t: newEmailTestTransport(srv.URL, a.apiKey)}
			if err := e.GetEmailPreferencesByLink(context.Background(), map[string]any{"token": "signed-link-abc"}, nil); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if gotPath != "/v2/email/preferences" {
				t.Errorf("path: got %q", gotPath)
			}
			if gotToken != "signed-link-abc" {
				t.Errorf("token query: got %q", gotToken)
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
