package hub

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/norbix-code/sdk-go/v2/norbix/internal/transport"
)

// Turning a whole module on or off changes project state, so the gateway
// serves these 18 routes as PUT (gateway refactoringV2; before, they were
// GET). A GET still works on the gateway for a while, but it is deprecated and
// logs a warning on every call, so the SDK must send PUT. Each case asserts
// the verb, the resolved path and that no body or query string is sent.
//
// Nothing leaves the process: every call goes to a local test server.

type moduleToggleCase struct {
	name string
	path string
	call func(ctx context.Context, t *transport.Transport) error
}

func moduleToggleCases() []moduleToggleCase {
	return []moduleToggleCase{
		{"EnableDatabase", "/v2/database/enable", func(ctx context.Context, t *transport.Transport) error {
			return (&DatabaseModule{t: t}).EnableDatabase(ctx, nil, nil)
		}},
		{"DisableDatabase", "/v2/database/disable", func(ctx context.Context, t *transport.Transport) error {
			return (&DatabaseModule{t: t}).DisableDatabase(ctx, nil, nil)
		}},
		{"EnableFiles", "/v2/files/enable", func(ctx context.Context, t *transport.Transport) error {
			return (&FilesModule{t: t}).EnableFiles(ctx, nil, nil)
		}},
		{"DisableFiles", "/v2/files/disable", func(ctx context.Context, t *transport.Transport) error {
			return (&FilesModule{t: t}).DisableFiles(ctx, nil, nil)
		}},
		{"EnablePush", "/v2/notifications/push/enable", func(ctx context.Context, t *transport.Transport) error {
			return (&NotificationsModule{t: t}).EnablePush(ctx, nil, nil)
		}},
		{"DisablePush", "/v2/notifications/push/disable", func(ctx context.Context, t *transport.Transport) error {
			return (&NotificationsModule{t: t}).DisablePush(ctx, nil, nil)
		}},
		{"EnableSms", "/v2/notifications/sms/enable", func(ctx context.Context, t *transport.Transport) error {
			return (&NotificationsModule{t: t}).EnableSms(ctx, nil, nil)
		}},
		{"DisableSms", "/v2/notifications/sms/disable", func(ctx context.Context, t *transport.Transport) error {
			return (&NotificationsModule{t: t}).DisableSms(ctx, nil, nil)
		}},
		{"EnableEmail", "/v2/notifications/email/enable", func(ctx context.Context, t *transport.Transport) error {
			return (&NotificationsModule{t: t}).EnableEmail(ctx, nil, nil)
		}},
		{"DisableEmail", "/v2/notifications/email/disable", func(ctx context.Context, t *transport.Transport) error {
			return (&NotificationsModule{t: t}).DisableEmail(ctx, nil, nil)
		}},
		{"EnablePayments", "/v2/payments/enable", func(ctx context.Context, t *transport.Transport) error {
			return (&PaymentsModule{t: t}).EnablePayments(ctx, nil, nil)
		}},
		{"DisablePayments", "/v2/payments/disable", func(ctx context.Context, t *transport.Transport) error {
			return (&PaymentsModule{t: t}).DisablePayments(ctx, nil, nil)
		}},
		{"EnableLogging", "/v2/logs/enable", func(ctx context.Context, t *transport.Transport) error {
			return (&LogsModule{t: t}).EnableLogging(ctx, nil, nil)
		}},
		{"DisableLogging", "/v2/logs/disable", func(ctx context.Context, t *transport.Transport) error {
			return (&LogsModule{t: t}).DisableLogging(ctx, nil, nil)
		}},
		{"EnableMembership", "/v2/membership/enable", func(ctx context.Context, t *transport.Transport) error {
			return (&MembershipModule{t: t}).EnableMembership(ctx, nil, nil)
		}},
		{"DisableMembership", "/v2/membership/disable", func(ctx context.Context, t *transport.Transport) error {
			return (&MembershipModule{t: t}).DisableMembership(ctx, nil, nil)
		}},
		{"EnableCode", "/v2/code/enable", func(ctx context.Context, t *transport.Transport) error {
			return (&CodeModule{t: t}).EnableCode(ctx, nil, nil)
		}},
		{"DisableCode", "/v2/code/disable", func(ctx context.Context, t *transport.Transport) error {
			return (&CodeModule{t: t}).DisableCode(ctx, nil, nil)
		}},
	}
}

func newModuleToggleTransport(baseURL string) *transport.Transport {
	return transport.New(&transport.Config{
		ProjectID:  "proj_1",
		AccountID:  "acct_1",
		APIKey:     "key_1",
		BaseURLAPI: "http://api.invalid",
		BaseURLHub: baseURL,
		APIVersion: "v2",
		HubVersion: "v2",
		Timeout:    5 * time.Second,
		Env:        "PROD",
	}, nil)
}

func TestModuleEnableDisableUsePut(t *testing.T) {
	for _, c := range moduleToggleCases() {
		t.Run(c.name, func(t *testing.T) {
			var gotMethod, gotPath, gotRawQuery, gotProject string
			var gotBody []byte
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotMethod = r.Method
				gotPath = r.URL.Path
				gotRawQuery = r.URL.RawQuery
				gotProject = r.Header.Get("X-CM-ProjectId")
				gotBody, _ = io.ReadAll(r.Body)
				w.WriteHeader(http.StatusNoContent)
			}))
			defer srv.Close()

			if err := c.call(context.Background(), newModuleToggleTransport(srv.URL)); err != nil {
				t.Fatalf("%s: unexpected error: %v", c.name, err)
			}
			if gotMethod != http.MethodPut {
				t.Errorf("verb: got %q want %q", gotMethod, http.MethodPut)
			}
			if gotPath != c.path {
				t.Errorf("path: got %q want %q", gotPath, c.path)
			}
			if gotRawQuery != "" {
				t.Errorf("query: got %q want none", gotRawQuery)
			}
			if len(gotBody) != 0 {
				t.Errorf("body: got %s want none", gotBody)
			}
			if gotProject != "proj_1" {
				t.Errorf("project header: got %q want %q", gotProject, "proj_1")
			}
		})
	}
}

// The gateway has exactly 18 module on/off routes (9 modules). A new module
// toggle must be added here too, so its verb is checked.
func TestModuleToggleSurfaceSize(t *testing.T) {
	if got := len(moduleToggleCases()); got != 18 {
		t.Fatalf("module toggle cases: got %d want 18", got)
	}
}
