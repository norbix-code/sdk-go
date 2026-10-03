package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/norbix-code/sdk-go/v2/norbix/internal/transport"
)

// Every wave-3 endUserChat route and the module method that calls it. Each case
// asserts the verb and the fully-resolved path, so a wrong path or a swapped
// verb fails here instead of at runtime. Nothing leaves the process: every
// call goes to a local test server, no provider is ever contacted.

type endUserChatCase struct {
	name string
	verb string
	path string
	call func(ctx context.Context, m *AiModule) error
}

func endUserChatCases() []endUserChatCase {
	body := map[string]any{"probe": "value"}

	return []endUserChatCase{
		{"GetEndUserChatAvailability", http.MethodGet, "/v2/ai/chat/availability",
			func(ctx context.Context, m *AiModule) error {
				return m.GetEndUserChatAvailability(ctx, body, nil)
			}},
		{"ListEndUserChatSessions", http.MethodGet, "/v2/ai/chat/sessions",
			func(ctx context.Context, m *AiModule) error {
				return m.ListEndUserChatSessions(ctx, body, nil)
			}},
		{"CreateEndUserChatSession", http.MethodPost, "/v2/ai/chat/sessions",
			func(ctx context.Context, m *AiModule) error {
				return m.CreateEndUserChatSession(ctx, body, nil)
			}},
		{"GetEndUserChatSession", http.MethodGet, "/v2/ai/chat/sessions/sessionId_1",
			func(ctx context.Context, m *AiModule) error {
				return m.GetEndUserChatSession(ctx, "sessionId_1", body, nil)
			}},
		{"RenameEndUserChatSession", http.MethodPatch, "/v2/ai/chat/sessions/sessionId_1",
			func(ctx context.Context, m *AiModule) error {
				return m.RenameEndUserChatSession(ctx, "sessionId_1", body, nil)
			}},
		{"DeleteEndUserChatSession", http.MethodDelete, "/v2/ai/chat/sessions/sessionId_1",
			func(ctx context.Context, m *AiModule) error {
				return m.DeleteEndUserChatSession(ctx, "sessionId_1", body, nil)
			}},
		{"PinEndUserChatSession", http.MethodPut, "/v2/ai/chat/sessions/sessionId_1/pin",
			func(ctx context.Context, m *AiModule) error {
				return m.PinEndUserChatSession(ctx, "sessionId_1", body, nil)
			}},
		{"ArchiveEndUserChatSession", http.MethodPut, "/v2/ai/chat/sessions/sessionId_1/archive",
			func(ctx context.Context, m *AiModule) error {
				return m.ArchiveEndUserChatSession(ctx, "sessionId_1", body, nil)
			}},
		{"GetEndUserChatEntries", http.MethodGet, "/v2/ai/chat/sessions/sessionId_1/entries",
			func(ctx context.Context, m *AiModule) error {
				return m.GetEndUserChatEntries(ctx, "sessionId_1", body, nil)
			}},
		{"SetEndUserChatEntryFeedback", http.MethodPut, "/v2/ai/chat/sessions/sessionId_1/entries/entryId_1/feedback",
			func(ctx context.Context, m *AiModule) error {
				return m.SetEndUserChatEntryFeedback(ctx, "sessionId_1", "entryId_1", body, nil)
			}},
		{"ListEndUserChatAttachments", http.MethodGet, "/v2/ai/chat/sessions/sessionId_1/attachments",
			func(ctx context.Context, m *AiModule) error {
				return m.ListEndUserChatAttachments(ctx, "sessionId_1", body, nil)
			}},
		{"UploadEndUserChatAttachment", http.MethodPost, "/v2/ai/chat/sessions/sessionId_1/attachments",
			func(ctx context.Context, m *AiModule) error {
				return m.UploadEndUserChatAttachment(ctx, "sessionId_1", body, nil)
			}},
		{"DeleteEndUserChatAttachment", http.MethodDelete, "/v2/ai/chat/attachments/attachmentId_1",
			func(ctx context.Context, m *AiModule) error {
				return m.DeleteEndUserChatAttachment(ctx, "attachmentId_1", body, nil)
			}},
		{"ListEndUserChatMemory", http.MethodGet, "/v2/ai/chat/memory",
			func(ctx context.Context, m *AiModule) error {
				return m.ListEndUserChatMemory(ctx, body, nil)
			}},
		{"ForgetEndUserChatMemory", http.MethodDelete, "/v2/ai/chat/memory/noteId_1",
			func(ctx context.Context, m *AiModule) error {
				return m.ForgetEndUserChatMemory(ctx, "noteId_1", body, nil)
			}},
		{"StartEndUserChatTurn", http.MethodPost, "/v2/ai/chat/turn",
			func(ctx context.Context, m *AiModule) error {
				return m.StartEndUserChatTurn(ctx, body, nil)
			}},
	}
}

func newEnduserchatTestModule(baseURL string) *AiModule {
	return &AiModule{t: transport.New(&transport.Config{
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

func TestEnduserchatEndpointsHitTheExpectedRoute(t *testing.T) {
	for _, c := range endUserChatCases() {
		t.Run(c.name, func(t *testing.T) {
			var gotMethod, gotPath, gotAuth, gotProject string
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotMethod = r.Method
				gotPath = r.URL.Path
				gotAuth = r.Header.Get("Authorization")
				gotProject = r.Header.Get("X-CM-ProjectId")
				w.WriteHeader(http.StatusNoContent)
			}))
			defer srv.Close()

			if err := c.call(context.Background(), newEnduserchatTestModule(srv.URL)); err != nil {
				t.Fatalf("%s: unexpected error: %v", c.name, err)
			}
			if gotMethod != c.verb {
				t.Errorf("%s verb: got %q want %q", c.name, gotMethod, c.verb)
			}
			if gotPath != c.path {
				t.Errorf("%s path: got %q want %q", c.name, gotPath, c.path)
			}
			if gotAuth != "Bearer key_1" {
				t.Errorf("%s auth header: got %q", c.name, gotAuth)
			}
			if gotProject != "proj_1" {
				t.Errorf("%s project header: got %q", c.name, gotProject)
			}
		})
	}
}

// The wave-3 endUserChat surface is 16 routes; a changed count means a route arrived
// or left untested.
func TestEnduserchatSurfaceSize(t *testing.T) {
	const want = 16
	if got := len(endUserChatCases()); got != want {
		t.Errorf("endUserChat endpoint count: got %d want %d", got, want)
	}
}
