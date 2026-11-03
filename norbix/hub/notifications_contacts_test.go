package hub

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/norbix-code/sdk-go/v2/norbix/internal/transport"
)

// Contacts moved from /notifications/contacts to /membership/users on the
// gateway; the old path answers 404. The deprecated Notifications contact
// methods keep their names and call the membership routes. Consent and
// unsubscribe are served by the API host only.
//
// Nothing leaves the process: every call goes to a local test server.

func TestNoModuleCallsTheOldContactsRoute(t *testing.T) {
	files, err := filepath.Glob("*.go")
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range files {
		if strings.HasSuffix(f, "_test.go") {
			continue
		}
		src, err := os.ReadFile(f)
		if err != nil {
			t.Fatal(err)
		}
		if strings.Contains(string(src), "/notifications/contacts") {
			t.Errorf("%s calls /notifications/contacts; contacts live on /membership/users", f)
		}
	}
}

func TestNotificationsContactMethodsCallMembershipRoutes(t *testing.T) {
	cases := []struct {
		name   string
		host   string
		method string
		path   string
		call   func(ctx context.Context, m *NotificationsModule) error
	}{
		{"GetAllContacts", "hub", "GET", "/v2/membership/users",
			func(ctx context.Context, m *NotificationsModule) error { return m.GetAllContacts(ctx, nil, nil) }},
		{"CreateContact", "hub", "POST", "/v2/membership/users",
			func(ctx context.Context, m *NotificationsModule) error { return m.CreateContact(ctx, nil, nil) }},
		{"MergeContacts", "hub", "POST", "/v2/membership/users/merge",
			func(ctx context.Context, m *NotificationsModule) error { return m.MergeContacts(ctx, nil, nil) }},
		{"DeleteContact", "hub", "DELETE", "/v2/membership/users/ct_1",
			func(ctx context.Context, m *NotificationsModule) error { return m.DeleteContact(ctx, "ct_1", nil, nil) }},
		{"GetContact", "hub", "GET", "/v2/membership/users/ct_1",
			func(ctx context.Context, m *NotificationsModule) error { return m.GetContact(ctx, "ct_1", nil, nil) }},
		{"AddContactIdentity", "hub", "POST", "/v2/membership/users/ct_1/identities",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.AddContactIdentity(ctx, "ct_1", nil, nil)
			}},
		{"RemoveContactIdentity", "hub", "DELETE", "/v2/membership/users/ct_1/identities/usr_1",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.RemoveContactIdentity(ctx, "ct_1", "usr_1", nil, nil)
			}},
		{"PromoteContactIdentity", "hub", "POST", "/v2/membership/users/ct_1/identities/usr_1/promote",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.PromoteContactIdentity(ctx, "ct_1", "usr_1", nil, nil)
			}},
		{"GrantContactConsent", "api", "POST", "/v2/membership/users/ct_1/marketing-state/Email/consent",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.GrantContactConsent(ctx, "ct_1", "Email", nil, nil)
			}},
		{"UnsubscribeContact", "api", "POST", "/v2/membership/users/ct_1/marketing-state/Email/unsubscribe",
			func(ctx context.Context, m *NotificationsModule) error {
				return m.UnsubscribeContact(ctx, "ct_1", "Email", nil, nil)
			}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var gotHost, gotMethod, gotPath string
			handler := func(host string) http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
					gotHost, gotMethod, gotPath = host, r.Method, r.URL.Path
					w.WriteHeader(http.StatusNoContent)
				}
			}
			api := httptest.NewServer(handler("api"))
			defer api.Close()
			hub := httptest.NewServer(handler("hub"))
			defer hub.Close()

			m := &NotificationsModule{t: transport.New(&transport.Config{
				ProjectID:  "proj_1",
				APIKey:     "key_1",
				BaseURLAPI: api.URL,
				BaseURLHub: hub.URL,
				APIVersion: "v2",
				HubVersion: "v2",
				Timeout:    5 * time.Second,
				Env:        "PROD",
			}, nil)}
			if err := c.call(context.Background(), m); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if gotHost != c.host || gotMethod != c.method || gotPath != c.path {
				t.Errorf("got %s %s on %s, want %s %s on %s", gotMethod, gotPath, gotHost, c.method, c.path, c.host)
			}
		})
	}
}
