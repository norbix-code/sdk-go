package hub

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/norbix-code/sdk-go/v2/norbix/internal/transport"
)

// The three notification preview routes open with a signed link alone: the
// "hash" is the key, so a client with no API key and no bearer token must
// still send the request. A client that has a key still sends it.
//
// Nothing leaves the process: every call goes to a local test server.

func newPreviewTestModule(baseURL, apiKey string) *NotificationsModule {
	return &NotificationsModule{t: transport.New(&transport.Config{
		ProjectID:  "proj_1",
		APIKey:     apiKey,
		BaseURLAPI: baseURL,
		BaseURLHub: baseURL,
		APIVersion: "v2",
		HubVersion: "v2",
		Timeout:    5 * time.Second,
		Env:        "PROD",
	}, nil)}
}

func TestPreviewWithSignedLinkNeedsNoSignIn(t *testing.T) {
	previews := []struct {
		name string
		path string
		call func(ctx context.Context, m *NotificationsModule, req map[string]any) error
	}{
		{"push", "/v2/notifications/push/preview",
			func(ctx context.Context, m *NotificationsModule, req map[string]any) error {
				return m.PreviewPushNotification(ctx, req, nil)
			}},
		{"email", "/v2/notifications/email/preview",
			func(ctx context.Context, m *NotificationsModule, req map[string]any) error {
				return m.PreviewEmailNotification(ctx, req, nil)
			}},
		{"sms", "/v2/notifications/sms/preview",
			func(ctx context.Context, m *NotificationsModule, req map[string]any) error {
				return m.PreviewSmsNotification(ctx, req, nil)
			}},
	}
	auths := []struct {
		name     string
		apiKey   string
		wantAuth string
	}{
		{"no credentials", "", ""},
		{"api key", "key_1", "Bearer key_1"},
	}

	for _, p := range previews {
		for _, a := range auths {
			t.Run(p.name+"/"+a.name, func(t *testing.T) {
				var gotPath, gotHash, gotAuth string
				var sawAuth bool
				srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					gotPath = r.URL.Path
					gotHash = r.URL.Query().Get("hash")
					gotAuth = r.Header.Get("Authorization")
					_, sawAuth = r.Header["Authorization"]
					w.WriteHeader(http.StatusNoContent)
				}))
				defer srv.Close()

				m := newPreviewTestModule(srv.URL, a.apiKey)
				if err := p.call(context.Background(), m, map[string]any{"hash": "signed-link-abc"}); err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
				if gotPath != p.path {
					t.Errorf("path: got %q want %q", gotPath, p.path)
				}
				if gotHash != "signed-link-abc" {
					t.Errorf("hash query: got %q want %q", gotHash, "signed-link-abc")
				}
				if gotAuth != a.wantAuth {
					t.Errorf("authorization: got %q want %q", gotAuth, a.wantAuth)
				}
				if a.wantAuth == "" && sawAuth {
					t.Error("authorization header was sent, want none")
				}
			})
		}
	}
}
