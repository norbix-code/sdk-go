package hub

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/norbix-code/sdk-go/v2/norbix/hub/dtos"
)

// Cloning an email, SMS or push template answers the copy's id
// (dtos.IdResponse, gateway 2026-10-23): the source template stays as it is
// and the copy gets a new id. Each case answers the clone with the gateway's
// {"id": "…"} body and checks that the caller reads the copy's id back.
// Nothing leaves the process.
func TestTemplateCloneReturnsTheCopyID(t *testing.T) {
	const sourceID, copyID = "tpl_source_1", "tpl_copy_1"

	cases := []struct {
		name string
		path string
		call func(ctx context.Context, m *NotificationsModule, out *dtos.IdResponse) error
	}{
		{"CloneEmailTemplate", "/v2/notifications/email/templates/" + sourceID + "/clone",
			func(ctx context.Context, m *NotificationsModule, out *dtos.IdResponse) error {
				return m.CloneEmailTemplate(ctx, sourceID, nil, out)
			}},
		{"CloneSmsTemplate", "/v2/notifications/sms/templates/" + sourceID + "/clone",
			func(ctx context.Context, m *NotificationsModule, out *dtos.IdResponse) error {
				return m.CloneSmsTemplate(ctx, sourceID, nil, out)
			}},
		{"ClonePushTemplate", "/v2/notifications/push/templates/" + sourceID + "/clone",
			func(ctx context.Context, m *NotificationsModule, out *dtos.IdResponse) error {
				return m.ClonePushTemplate(ctx, sourceID, nil, out)
			}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var gotMethod, gotPath string
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotMethod, gotPath = r.Method, r.URL.Path
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{"id":"` + copyID + `"}`))
			}))
			defer srv.Close()

			var out dtos.IdResponse
			if err := c.call(context.Background(), newPushTestModule(srv.URL), &out); err != nil {
				t.Fatalf("%s: unexpected error: %v", c.name, err)
			}
			if gotMethod != http.MethodPost || gotPath != c.path {
				t.Errorf("%s: sent %s %s, want POST %s", c.name, gotMethod, gotPath, c.path)
			}
			if out.Id != copyID {
				t.Errorf("%s: copy id = %q, want %q", c.name, out.Id, copyID)
			}
		})
	}
}
