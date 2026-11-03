package api

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/norbix-code/sdk-go/v2/norbix/api/dtos"
	"github.com/norbix-code/sdk-go/v2/norbix/internal/transport"
)

// Every notification route on the API surface (Email, SMS and Push campaigns
// plus the signed-in user's own push devices and messages) and the module
// method that calls it. Each case asserts the verb and the fully-resolved path
// — version substituted, route tokens filled — so a wrong path, a swapped verb
// or a parameter in the wrong place fails here instead of at runtime.
//
// Nothing leaves the process: every call goes to a local test server, so no
// provider is ever contacted.

const (
	testAPICampaignID     = "camp_1"
	testAPIBatchID        = "batch_1"
	testAPINotificationID = "notif_1"
	testAPIDeviceID       = "pnd_1"
	testAPIMessageID      = "ntf_1"
)

type apiNotificationsCase struct {
	name string
	verb string
	path string
	call func(ctx context.Context, m *NotificationsModule) error
}

func apiNotificationsCases() []apiNotificationsCase {
	body := map[string]any{"probe": "value"}

	return []apiNotificationsCase{
		// --- own push devices and messages (end-user)
		{"RegisterOwnPushDevice", http.MethodPost, "/v2/notifications/push/devices",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.RegisterOwnPushDevice(ctx, body, nil)
			}},
		{"GetOwnPushDevices", http.MethodGet, "/v2/notifications/push/devices",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetOwnPushDevices(ctx, body, nil)
			}},
		{"DeleteOwnPushDevice", http.MethodDelete, "/v2/notifications/push/devices/pnd_1",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.DeleteOwnPushDevice(ctx, testAPIDeviceID, body, nil)
			}},
		{"DeleteOwnPushDeviceByToken", http.MethodDelete, "/v2/notifications/push/devices",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.DeleteOwnPushDeviceByToken(ctx, body, nil)
			}},
		{"GetOwnPushMessages", http.MethodGet, "/v2/notifications/push/messages",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetOwnPushMessages(ctx, body, nil)
			}},
		{"MarkOwnPushMessageRead", http.MethodPost, "/v2/notifications/push/messages/ntf_1/read",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.MarkOwnPushMessageRead(ctx, testAPIMessageID, body, nil)
			}},
		{"MarkAllOwnPushMessagesRead", http.MethodPost, "/v2/notifications/push/messages/read",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.MarkAllOwnPushMessagesRead(ctx, body, nil)
			}},
		// --- email campaigns
		{"CreateEmailCampaign", http.MethodPost, "/v2/notifications/email/campaigns",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.CreateEmailCampaign(ctx, body, nil)
			}},
		{"DeleteEmailCampaign", http.MethodDelete, "/v2/notifications/email/campaigns/camp_1",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.DeleteEmailCampaign(ctx, testAPICampaignID, body, nil)
			}},
		{"GetEmailCampaign", http.MethodGet, "/v2/notifications/email/campaigns/camp_1",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetEmailCampaign(ctx, testAPICampaignID, body, nil)
			}},
		{"GetEmailCampaigns", http.MethodGet, "/v2/notifications/email/campaigns",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetEmailCampaigns(ctx, body, nil)
			}},
		{"GetEmailCampaignBatches", http.MethodGet, "/v2/notifications/email/campaigns/camp_1/batches",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetEmailCampaignBatches(ctx, testAPICampaignID, body, nil)
			}},
		{"GetEmailCampaignBatchNotification", http.MethodGet, "/v2/notifications/email/campaigns/camp_1/batches/batch_1/notif_1",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetEmailCampaignBatchNotification(ctx, testAPICampaignID, testAPIBatchID, testAPINotificationID, body, nil)
			}},
		{"GetEmailCampaignBatchNotifications", http.MethodGet, "/v2/notifications/email/campaigns/camp_1/batches/batch_1",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetEmailCampaignBatchNotifications(ctx, testAPICampaignID, testAPIBatchID, body, nil)
			}},
		{"GetEmailCampaignStatistics", http.MethodGet, "/v2/notifications/email/campaigns/camp_1/stats",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetEmailCampaignStatistics(ctx, testAPICampaignID, body, nil)
			}},
		{"PreviewEmailNotification", http.MethodGet, "/v2/notifications/email/preview",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.PreviewEmailNotification(ctx, body, nil)
			}},
		{"RetryFailedEmailNotifications", http.MethodPost, "/v2/notifications/email/campaigns/camp_1/retry-failed",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.RetryFailedEmailNotifications(ctx, testAPICampaignID, body, nil)
			}},
		{"StopEmailCampaign", http.MethodPost, "/v2/notifications/email/campaigns/camp_1/stop",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.StopEmailCampaign(ctx, testAPICampaignID, body, nil)
			}},
		{"GetEmailCampaignMessages", http.MethodGet, "/v2/notifications/emails/campaigns/camp_1/messages",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetEmailCampaignMessages(ctx, testAPICampaignID, body, nil)
			}},
		// --- sms campaigns
		{"CreateSmsCampaign", http.MethodPost, "/v2/notifications/sms/campaigns",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.CreateSmsCampaign(ctx, body, nil)
			}},
		{"DeleteSmsCampaign", http.MethodDelete, "/v2/notifications/sms/campaigns/camp_1",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.DeleteSmsCampaign(ctx, testAPICampaignID, body, nil)
			}},
		{"GetSmsCampaign", http.MethodGet, "/v2/notifications/sms/campaigns/camp_1",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetSmsCampaign(ctx, testAPICampaignID, body, nil)
			}},
		{"GetSmsCampaigns", http.MethodGet, "/v2/notifications/sms/campaigns",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetSmsCampaigns(ctx, body, nil)
			}},
		{"GetSmsCampaignBatches", http.MethodGet, "/v2/notifications/sms/campaigns/camp_1/batches",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetSmsCampaignBatches(ctx, testAPICampaignID, body, nil)
			}},
		{"GetSmsCampaignBatchNotification", http.MethodGet, "/v2/notifications/sms/campaigns/camp_1/batches/batch_1/notif_1",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetSmsCampaignBatchNotification(ctx, testAPICampaignID, testAPIBatchID, testAPINotificationID, body, nil)
			}},
		{"GetSmsCampaignBatchNotifications", http.MethodGet, "/v2/notifications/sms/campaigns/camp_1/batches/batch_1",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetSmsCampaignBatchNotifications(ctx, testAPICampaignID, testAPIBatchID, body, nil)
			}},
		{"GetSmsCampaignStatistics", http.MethodGet, "/v2/notifications/sms/campaigns/camp_1/stats",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetSmsCampaignStatistics(ctx, testAPICampaignID, body, nil)
			}},
		{"PreviewSmsNotification", http.MethodGet, "/v2/notifications/sms/preview",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.PreviewSmsNotification(ctx, body, nil)
			}},
		{"RetryFailedSmsNotifications", http.MethodPost, "/v2/notifications/sms/campaigns/camp_1/retry-failed",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.RetryFailedSmsNotifications(ctx, testAPICampaignID, body, nil)
			}},
		{"StopSmsCampaign", http.MethodPost, "/v2/notifications/sms/campaigns/camp_1/stop",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.StopSmsCampaign(ctx, testAPICampaignID, body, nil)
			}},
		{"GetSmsCampaignMessages", http.MethodGet, "/v2/notifications/sms/campaigns/camp_1/messages",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetSmsCampaignMessages(ctx, testAPICampaignID, body, nil)
			}},
		// --- push campaigns
		{"GetPushCampaignAudienceCount", http.MethodPost, "/v2/notifications/push/campaigns/audience-count",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetPushCampaignAudienceCount(ctx, body, nil)
			}},
		{"CreatePushCampaign", http.MethodPost, "/v2/notifications/push/campaigns",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.CreatePushCampaign(ctx, body, nil)
			}},
		{"DeletePushCampaign", http.MethodDelete, "/v2/notifications/push/campaigns/camp_1",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.DeletePushCampaign(ctx, testAPICampaignID, body, nil)
			}},
		{"GetPushCampaign", http.MethodGet, "/v2/notifications/push/campaigns/camp_1",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetPushCampaign(ctx, testAPICampaignID, body, nil)
			}},
		{"GetPushCampaigns", http.MethodGet, "/v2/notifications/push/campaigns",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetPushCampaigns(ctx, body, nil)
			}},
		{"GetPushCampaignBatches", http.MethodGet, "/v2/notifications/push/campaigns/camp_1/batches",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetPushCampaignBatches(ctx, testAPICampaignID, body, nil)
			}},
		{"GetPushCampaignBatchNotification", http.MethodGet, "/v2/notifications/push/campaigns/camp_1/batches/batch_1/notif_1",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetPushCampaignBatchNotification(ctx, testAPICampaignID, testAPIBatchID, testAPINotificationID, body, nil)
			}},
		{"GetPushCampaignBatchNotifications", http.MethodGet, "/v2/notifications/push/campaigns/camp_1/batches/batch_1",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetPushCampaignBatchNotifications(ctx, testAPICampaignID, testAPIBatchID, body, nil)
			}},
		{"GetPushCampaignStatistics", http.MethodGet, "/v2/notifications/push/campaigns/camp_1/stats",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetPushCampaignStatistics(ctx, testAPICampaignID, body, nil)
			}},
		{"PreviewPushNotification", http.MethodGet, "/v2/notifications/push/preview",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.PreviewPushNotification(ctx, body, nil)
			}},
		{"RetryFailedPushNotifications", http.MethodPost, "/v2/notifications/push/campaigns/camp_1/retry-failed",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.RetryFailedPushNotifications(ctx, testAPICampaignID, body, nil)
			}},
		{"StopPushCampaign", http.MethodPost, "/v2/notifications/push/campaigns/camp_1/stop",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.StopPushCampaign(ctx, testAPICampaignID, body, nil)
			}},
		{"GetPushCampaignMessages", http.MethodGet, "/v2/notifications/push/campaigns/camp_1/messages",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GetPushCampaignMessages(ctx, testAPICampaignID, body, nil)
			}},
	}
}

func newAPINotificationsTestModule(baseURL, apiKey string) *NotificationsModule {
	return &NotificationsModule{t: transport.New(&transport.Config{
		ProjectID:  "proj_1",
		APIKey:     apiKey,
		BaseURLAPI: baseURL,
		BaseURLHub: "http://hub.invalid",
		APIVersion: "v2",
		HubVersion: "v2",
		Timeout:    5 * time.Second,
		Env:        "PROD",
	}, nil)}
}

func TestAPINotificationsEndpointsHitTheExpectedRoute(t *testing.T) {
	for _, c := range apiNotificationsCases() {
		t.Run(c.name, func(t *testing.T) {
			var gotMethod, gotPath, gotQuery, gotProject string
			var gotBody []byte
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotMethod = r.Method
				gotPath = r.URL.Path
				gotQuery = r.URL.RawQuery
				gotProject = r.Header.Get("norbix-project-id")
				gotBody, _ = io.ReadAll(r.Body)
				w.WriteHeader(http.StatusNoContent)
			}))
			defer srv.Close()

			if err := c.call(context.Background(), newAPINotificationsTestModule(srv.URL, "key_1")); err != nil {
				t.Fatalf("%s: unexpected error: %v", c.name, err)
			}
			if gotMethod != c.verb {
				t.Errorf("verb: got %q want %q", gotMethod, c.verb)
			}
			if gotPath != c.path {
				t.Errorf("path: got %q want %q", gotPath, c.path)
			}
			if gotProject != "proj_1" {
				t.Errorf("project header: got %q want %q", gotProject, "proj_1")
			}
			// The probe body travels as the query string on GET / DELETE and as
			// JSON otherwise — never both, never neither.
			switch c.verb {
			case http.MethodGet, http.MethodDelete:
				if gotQuery != "probe=value" || len(gotBody) != 0 {
					t.Errorf("read verb: query %q body %q", gotQuery, gotBody)
				}
			default:
				if gotQuery != "" || !strings.Contains(string(gotBody), `"probe":"value"`) {
					t.Errorf("write verb: query %q body %q", gotQuery, gotBody)
				}
			}
		})
	}
}

// The API notifications surface is 44 routes: 12 Email, 12 SMS, 20 Push (13
// campaign + 7 own device / message). A new route changes this count, so it
// cannot arrive untested.
func TestAPINotificationsSurfaceSize(t *testing.T) {
	const want = 44
	if got := len(apiNotificationsCases()); got != want {
		t.Errorf("notifications endpoint count: got %d want %d", got, want)
	}
	seen := map[string]bool{}
	for _, c := range apiNotificationsCases() {
		key := c.verb + " " + c.path
		if seen[key] {
			t.Errorf("route listed twice: %s", key)
		}
		seen[key] = true
	}
}

// The three preview routes open with a signed link alone: the "hash" is the
// key, so a client with no API key and no bearer token must still send the
// request. A client that has a key still sends it.
func TestAPIPreviewWithSignedLinkNeedsNoSignIn(t *testing.T) {
	previews := []struct {
		name string
		path string
		call func(ctx context.Context, m *NotificationsModule, req map[string]any) error
	}{
		{"email", "/v2/notifications/email/preview",
			func(ctx context.Context, m *NotificationsModule, req map[string]any) error {
				return m.PreviewEmailNotification(ctx, req, nil)
			}},
		{"sms", "/v2/notifications/sms/preview",
			func(ctx context.Context, m *NotificationsModule, req map[string]any) error {
				return m.PreviewSmsNotification(ctx, req, nil)
			}},
		{"push", "/v2/notifications/push/preview",
			func(ctx context.Context, m *NotificationsModule, req map[string]any) error {
				return m.PreviewPushNotification(ctx, req, nil)
			}},
	}
	for _, p := range previews {
		t.Run(p.name, func(t *testing.T) {
			var gotPath, gotQuery, gotAuth, gotKey string
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotPath = r.URL.Path
				gotQuery = r.URL.RawQuery
				gotAuth = r.Header.Get("Authorization")
				gotKey = r.Header.Get("x-api-key")
				w.WriteHeader(http.StatusNoContent)
			}))
			defer srv.Close()

			// No key, no token: the request still goes out, with the hash.
			err := p.call(context.Background(), newAPINotificationsTestModule(srv.URL, ""), map[string]any{"hash": "signed.link"})
			if err != nil {
				t.Fatalf("no sign-in: unexpected error: %v", err)
			}
			if gotPath != p.path || gotQuery != "hash=signed.link" {
				t.Errorf("no sign-in: path %q query %q", gotPath, gotQuery)
			}
			if gotAuth != "" || gotKey != "" {
				t.Errorf("no sign-in: auth headers sent: Authorization=%q x-api-key=%q", gotAuth, gotKey)
			}

			// With a key: the key is still sent.
			err = p.call(context.Background(), newAPINotificationsTestModule(srv.URL, "key_1"), map[string]any{"hash": "signed.link"})
			if err != nil {
				t.Fatalf("with key: unexpected error: %v", err)
			}
			if gotAuth == "" && gotKey == "" {
				t.Errorf("with key: no auth header sent")
			}
		})
	}
}

// Every other route needs the project credentials: without a key or a token
// the transport refuses the call before anything leaves the process.
func TestAPINotificationsRefuseWithoutCredentials(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("request reached the server: %s %s", r.Method, r.URL.Path)
	}))
	defer srv.Close()
	m := newAPINotificationsTestModule(srv.URL, "")
	if err := m.GetOwnPushDevices(context.Background(), nil, nil); err == nil {
		t.Fatal("expected an error without credentials")
	}
}

// The own-push answers decode straight into the generated dtos: the device
// list is {items}, the message list is a paged {list}.
func TestAPIOwnPushAnswersDecodeIntoTheDtos(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/v2/notifications/push/devices":
			_, _ = io.WriteString(w, `{"id":"pnd_1"}`)
		case r.URL.Path == "/v2/notifications/push/devices":
			_, _ = io.WriteString(w, `{"items":[{"id":"pnd_1","provider":"fcm","platform":"android","token":"tok"}]}`)
		case r.URL.Path == "/v2/notifications/push/messages":
			_, _ = io.WriteString(w, `{"list":{"items":[{"id":"ntf_1","campaignId":"camp_1","title":"Hi","body":"There","isRead":false}],"totalCount":1,"pageNumber":0,"pageSize":20}}`)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer srv.Close()
	m := newAPINotificationsTestModule(srv.URL, "key_1")
	ctx := context.Background()

	var created dtos.IdResponse
	if err := m.RegisterOwnPushDevice(ctx, map[string]any{"provider": "fcm", "token": "tok"}, &created); err != nil || created.Id != "pnd_1" {
		t.Fatalf("register: %v %+v", err, created)
	}
	var devices dtos.GetOwnPushDevicesResponse
	if err := m.GetOwnPushDevices(ctx, nil, &devices); err != nil {
		t.Fatalf("devices: %v", err)
	}
	if len(devices.Items) != 1 || devices.Items[0].Id != "pnd_1" || devices.Items[0].Provider != "fcm" {
		t.Errorf("devices: %+v", devices.Items)
	}
	var messages dtos.GetOwnPushMessagesResponse
	if err := m.GetOwnPushMessages(ctx, map[string]any{"unreadOnly": true}, &messages); err != nil {
		t.Fatalf("messages: %v", err)
	}
	if len(messages.List.Items) != 1 || messages.List.Items[0].Id != "ntf_1" || messages.List.Items[0].IsRead {
		t.Errorf("messages: %+v", messages.List)
	}
	raw, _ := json.Marshal(messages.List.Items[0])
	if !strings.Contains(string(raw), `"campaignId":"camp_1"`) {
		t.Errorf("message round-trip: %s", raw)
	}
}
