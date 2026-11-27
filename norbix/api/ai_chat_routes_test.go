package api

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	norbixerr "github.com/norbix-code/sdk-go/v2/norbix/errors"
	"github.com/norbix-code/sdk-go/v2/norbix/internal/transport"
)

// Matrix section 08 AI chat, API side: every end-user AI chat and AI tool
// route of AiModule, and the chat stream. Each route case asserts the verb,
// the fully-resolved path on the API (the Hub base is an unreachable
// *.invalid host), the exact query string, the JSON body and the key +
// project headers. Field names are the gateway's (Api.AI ChatTurn.cs,
// ChatSessions.cs, ChatItems.cs, EndUserAiTools.cs). The older table in
// ai_test.go points both hosts at one server and checks neither query nor
// body.
//
// The stream tests run AiChat against a fake SSE server: the request it
// opens (host, path, channel, auth) and how it parses the frames.
//
// Nothing leaves the process and no LLM is contacted.

type aiChatRouteCase struct {
	name  string
	verb  string
	path  string
	query url.Values     // nil: no query string expected
	body  map[string]any // nil: no body expected
	call  func(ctx context.Context, m *AiModule, req map[string]any) error
	req   map[string]any
}

func aiChatRouteCases() []aiChatRouteCase {
	return []aiChatRouteCase{
		{"StartEndUserChatTurn", http.MethodPost, "/v2/ai/chat/turn", nil, map[string]any{"sessionId": "aics_1", "message": "Where is my order?"},
			func(ctx context.Context, m *AiModule, req map[string]any) error {
				return m.StartEndUserChatTurn(ctx, req, nil)
			}, map[string]any{"sessionId": "aics_1", "message": "Where is my order?"}},
		{"GetEndUserChatAvailability", http.MethodGet, "/v2/ai/chat/availability", nil, nil,
			func(ctx context.Context, m *AiModule, req map[string]any) error {
				return m.GetEndUserChatAvailability(ctx, req, nil)
			}, nil},
		{"ListEndUserChatSessions", http.MethodGet, "/v2/ai/chat/sessions", url.Values{"take": {"10"}, "includeArchived": {"true"}}, nil,
			func(ctx context.Context, m *AiModule, req map[string]any) error {
				return m.ListEndUserChatSessions(ctx, req, nil)
			}, map[string]any{"take": 10, "includeArchived": true}},
		{"CreateEndUserChatSession", http.MethodPost, "/v2/ai/chat/sessions", nil, map[string]any{"assistantId": "ast_1", "title": "Orders"},
			func(ctx context.Context, m *AiModule, req map[string]any) error {
				return m.CreateEndUserChatSession(ctx, req, nil)
			}, map[string]any{"assistantId": "ast_1", "title": "Orders"}},
		{"GetEndUserChatSession", http.MethodGet, "/v2/ai/chat/sessions/aics_1", nil, nil,
			func(ctx context.Context, m *AiModule, req map[string]any) error {
				return m.GetEndUserChatSession(ctx, "aics_1", req, nil)
			}, nil},
		{"RenameEndUserChatSession", http.MethodPatch, "/v2/ai/chat/sessions/aics_1", nil, map[string]any{"title": "Returns"},
			func(ctx context.Context, m *AiModule, req map[string]any) error {
				return m.RenameEndUserChatSession(ctx, "aics_1", req, nil)
			}, map[string]any{"title": "Returns"}},
		{"DeleteEndUserChatSession", http.MethodDelete, "/v2/ai/chat/sessions/aics_1", nil, nil,
			func(ctx context.Context, m *AiModule, req map[string]any) error {
				return m.DeleteEndUserChatSession(ctx, "aics_1", req, nil)
			}, nil},
		{"PinEndUserChatSession", http.MethodPut, "/v2/ai/chat/sessions/aics_1/pin", nil, map[string]any{"pinned": true},
			func(ctx context.Context, m *AiModule, req map[string]any) error {
				return m.PinEndUserChatSession(ctx, "aics_1", req, nil)
			}, map[string]any{"pinned": true}},
		{"ArchiveEndUserChatSession", http.MethodPut, "/v2/ai/chat/sessions/aics_1/archive", nil, map[string]any{"archived": false},
			func(ctx context.Context, m *AiModule, req map[string]any) error {
				return m.ArchiveEndUserChatSession(ctx, "aics_1", req, nil)
			}, map[string]any{"archived": false}},
		{"GetEndUserChatEntries", http.MethodGet, "/v2/ai/chat/sessions/aics_1/entries", url.Values{"afterSeq": {"3"}, "take": {"50"}}, nil,
			func(ctx context.Context, m *AiModule, req map[string]any) error {
				return m.GetEndUserChatEntries(ctx, "aics_1", req, nil)
			}, map[string]any{"afterSeq": 3, "take": 50}},
		{"SetEndUserChatEntryFeedback", http.MethodPut, "/v2/ai/chat/sessions/aics_1/entries/ent_1/feedback", nil, map[string]any{"feedback": "down"},
			func(ctx context.Context, m *AiModule, req map[string]any) error {
				return m.SetEndUserChatEntryFeedback(ctx, "aics_1", "ent_1", req, nil)
			}, map[string]any{"feedback": "down"}},
		{"DecideEndUserChatAction", http.MethodPost, "/v2/ai/chat/sessions/aics_1/actions/ent_7", nil, map[string]any{"decision": "apply"},
			func(ctx context.Context, m *AiModule, req map[string]any) error {
				return m.DecideEndUserChatAction(ctx, "aics_1", "ent_7", req, nil)
			}, map[string]any{"decision": "apply"}},
		{"ListEndUserChatAttachments", http.MethodGet, "/v2/ai/chat/sessions/aics_1/attachments", nil, nil,
			func(ctx context.Context, m *AiModule, req map[string]any) error {
				return m.ListEndUserChatAttachments(ctx, "aics_1", req, nil)
			}, nil},
		{"UploadEndUserChatAttachment", http.MethodPost, "/v2/ai/chat/sessions/aics_1/attachments", nil, map[string]any{"fileName": "receipt.pdf", "contentType": "application/pdf", "base64Content": "JVBERg=="},
			func(ctx context.Context, m *AiModule, req map[string]any) error {
				return m.UploadEndUserChatAttachment(ctx, "aics_1", req, nil)
			}, map[string]any{"fileName": "receipt.pdf", "contentType": "application/pdf", "base64Content": "JVBERg=="}},
		{"DeleteEndUserChatAttachment", http.MethodDelete, "/v2/ai/chat/attachments/att_1", nil, nil,
			func(ctx context.Context, m *AiModule, req map[string]any) error {
				return m.DeleteEndUserChatAttachment(ctx, "att_1", req, nil)
			}, nil},
		{"ListEndUserChatMemory", http.MethodGet, "/v2/ai/chat/memory", url.Values{"take": {"25"}}, nil,
			func(ctx context.Context, m *AiModule, req map[string]any) error {
				return m.ListEndUserChatMemory(ctx, req, nil)
			}, map[string]any{"take": 25}},
		{"ForgetEndUserChatMemory", http.MethodDelete, "/v2/ai/chat/memory/mem_1", nil, nil,
			func(ctx context.Context, m *AiModule, req map[string]any) error {
				return m.ForgetEndUserChatMemory(ctx, "mem_1", req, nil)
			}, nil},
		{"GetEndUserAiTools", http.MethodGet, "/v2/ai/tools", nil, nil,
			func(ctx context.Context, m *AiModule, req map[string]any) error {
				return m.GetEndUserAiTools(ctx, req, nil)
			}, nil},
		{"InvokeEndUserAiTool", http.MethodPost, "/v2/ai/tools/get_my_orders", nil, map[string]any{"argumentsJson": "{\"status\":\"open\"}"},
			func(ctx context.Context, m *AiModule, req map[string]any) error {
				return m.InvokeEndUserAiTool(ctx, "get_my_orders", req, nil)
			}, map[string]any{"argumentsJson": "{\"status\":\"open\"}"}},
	}
}

func newAiChatRoutesTestModule(apiURL string) *AiModule {
	return &AiModule{t: transport.New(&transport.Config{
		ProjectID:  "proj_1",
		AccountID:  "acct_1",
		APIKey:     "key_1",
		BaseURLAPI: apiURL,
		BaseURLHub: "http://hub.invalid",
		APIVersion: "v2",
		HubVersion: "v2",
		Timeout:    5 * time.Second,
		Env:        "PROD",
	}, nil)}
}

func TestAiChatRoutesHitTheExpectedRoute(t *testing.T) {
	for _, c := range aiChatRouteCases() {
		t.Run(c.name, func(t *testing.T) {
			var gotMethod, gotPath, gotRawQuery, gotAuth, gotProject string
			var gotBody []byte
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotMethod = r.Method
				gotPath = r.URL.Path
				gotRawQuery = r.URL.RawQuery
				gotAuth = r.Header.Get("Authorization")
				gotProject = r.Header.Get("norbix-project-id")
				gotBody, _ = io.ReadAll(r.Body)
				w.WriteHeader(http.StatusNoContent)
			}))
			defer srv.Close()

			if err := c.call(context.Background(), newAiChatRoutesTestModule(srv.URL), c.req); err != nil {
				t.Fatalf("%s: unexpected error: %v", c.name, err)
			}
			if gotMethod != c.verb {
				t.Errorf("verb: got %q want %q", gotMethod, c.verb)
			}
			if gotPath != c.path {
				t.Errorf("path: got %q want %q", gotPath, c.path)
			}
			if gotAuth != "Bearer key_1" {
				t.Errorf("authorization: got %q want %q", gotAuth, "Bearer key_1")
			}
			if gotProject != "proj_1" {
				t.Errorf("project header: got %q want %q", gotProject, "proj_1")
			}

			if c.query == nil && gotRawQuery != "" {
				t.Errorf("query: got %q want none", gotRawQuery)
			}
			if c.query != nil {
				gotQuery, _ := url.ParseQuery(gotRawQuery)
				gq, _ := json.Marshal(gotQuery)
				wq, _ := json.Marshal(c.query)
				if string(gq) != string(wq) {
					t.Errorf("query: got %s want %s", gq, wq)
				}
			}

			if c.body == nil {
				if len(gotBody) != 0 {
					t.Errorf("body: got %s want none", gotBody)
				}
				return
			}
			var got map[string]any
			if err := json.Unmarshal(gotBody, &got); err != nil {
				t.Fatalf("body is not JSON: %v (%s)", err, gotBody)
			}
			gb, _ := json.Marshal(got)
			wb, _ := json.Marshal(c.body)
			if string(gb) != string(wb) {
				t.Errorf("body: got %s want %s", gb, wb)
			}
		})
	}
}

// 17 end-user chat routes (with the held-action decision) + 2 AI tool routes.
func TestAiChatRoutesSurfaceSize(t *testing.T) {
	if got := len(aiChatRouteCases()); got != 19 {
		t.Errorf("end-user AI chat route count: got %d want 19", got)
	}
}

func TestAiChatChannelName(t *testing.T) {
	if got := AiChatChannel("pr_x", "usr_y"); got != "ai-chat:pr_x:usr_y" {
		t.Errorf("channel: got %q", got)
	}
}

// One turn as the gateway streams it: a heartbeat, a subscription notice that
// is not an envelope, then started → two tokens (one data field split over
// two lines is not used by the gateway, but is legal SSE) → completed.
const aiChatStreamBody = ": heartbeat\n\n" +
	"event: cmd.onConnect\ndata: {\"id\":\"sub_1\"}\n\n" +
	"id: 1\nevent: cmd.ai.chat.turn.started\ndata: {\"channel\":\"ai-chat\",\"eventName\":\"ai.chat.turn.started\",\"projectId\":\"proj_1\",\"createdAtUtc\":\"2026-11-18T10:00:00Z\",\"payload\":{\"turnId\":\"trn_1\"}}\n\n" +
	"id: 2\nevent: cmd.ai.chat.turn.token\ndata: {\"channel\":\"ai-chat\",\"eventName\":\"ai.chat.turn.token\",\"projectId\":\"proj_1\",\r\ndata: \"createdAtUtc\":\"2026-11-18T10:00:01Z\",\"payload\":{\"turnId\":\"trn_1\",\"text\":\"Your order \"}}\r\n\r\n" +
	"id: 3\nevent: cmd.ai.chat.turn.token\ndata: {\"channel\":\"ai-chat\",\"eventName\":\"ai.chat.turn.token\",\"projectId\":\"proj_1\",\"createdAtUtc\":\"2026-11-18T10:00:02Z\",\"payload\":{\"turnId\":\"trn_1\",\"text\":\"has shipped.\"}}\n\n" +
	"id: 4\nevent: cmd.ai.chat.turn.completed\ndata: {\"channel\":\"ai-chat\",\"eventName\":\"ai.chat.turn.completed\",\"projectId\":\"proj_1\",\"createdAtUtc\":\"2026-11-18T10:00:03Z\",\"payload\":{\"turnId\":\"trn_1\"}}\n\n"

func TestAiChatStreamOpensTheCallersChannelOnTheAPIHost(t *testing.T) {
	var gotMethod, gotPath, gotRawQuery, gotAuth, gotProject, gotAccept string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod, gotPath, gotRawQuery = r.Method, r.URL.Path, r.URL.RawQuery
		gotAuth = r.Header.Get("Authorization")
		gotProject = r.Header.Get("norbix-project-id")
		gotAccept = r.Header.Get("Accept")
		w.Header().Set("Content-Type", "text/event-stream")
		_, _ = io.WriteString(w, aiChatStreamBody)
	}))
	defer srv.Close()

	var got []AiChatEvent
	err := newAiChatRoutesTestModule(srv.URL).AiChat(context.Background(), "usr_1", func(e AiChatEvent) bool {
		got = append(got, e)
		return true
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet || gotPath != "/event-stream" {
		t.Errorf("route: got %s %s want GET /event-stream", gotMethod, gotPath)
	}
	if gotRawQuery != "channels=ai-chat%3Aproj_1%3Ausr_1" {
		t.Errorf("query: got %q", gotRawQuery)
	}
	if gotAuth != "Bearer key_1" || gotProject != "proj_1" || gotAccept != "text/event-stream" {
		t.Errorf("headers: auth %q project %q accept %q", gotAuth, gotProject, gotAccept)
	}

	var names, text []string
	for _, e := range got {
		names = append(names, e.EventName+"#"+e.FrameID)
		if s, ok := e.Payload["text"].(string); ok {
			text = append(text, s)
		}
		if e.Payload["turnId"] != "trn_1" {
			t.Errorf("%s: turnId %v", e.EventName, e.Payload["turnId"])
		}
	}
	want := "ai.chat.turn.started#1 ai.chat.turn.token#2 ai.chat.turn.token#3 ai.chat.turn.completed#4"
	if strings.Join(names, " ") != want {
		t.Errorf("events: got %q want %q", strings.Join(names, " "), want)
	}
	if strings.Join(text, "") != "Your order has shipped." {
		t.Errorf("tokens: got %q", strings.Join(text, ""))
	}
}

func TestAiChatStreamStopsWhenTheHandlerSaysSo(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, aiChatStreamBody)
	}))
	defer srv.Close()

	n := 0
	err := newAiChatRoutesTestModule(srv.URL).AiChat(context.Background(), "usr_1", func(e AiChatEvent) bool {
		n++
		return e.EventName != "ai.chat.turn.token"
	})
	if err != nil || n != 2 {
		t.Errorf("got %d events, err %v; want 2, nil", n, err)
	}
}

func TestAiChatStreamRefusalIsReturnedNotRetried(t *testing.T) {
	calls := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		_, _ = io.WriteString(w, `{"responseStatus":{"errorCode":"AiChatChannelRefused","message":"Not your channel."}}`)
	}))
	defer srv.Close()

	err := newAiChatRoutesTestModule(srv.URL).AiChat(context.Background(), "usr_other", func(AiChatEvent) bool { return true })
	var nerr *norbixerr.Error
	if !errors.As(err, &nerr) {
		t.Fatalf("want a *norbixerr.Error, got %T %v", err, err)
	}
	if nerr.Status != http.StatusForbidden || nerr.Code != "AiChatChannelRefused" {
		t.Errorf("error: status %d code %q", nerr.Status, nerr.Code)
	}
	if calls != 1 {
		t.Errorf("calls: got %d want 1", calls)
	}
}

func TestAiChatStreamNeedsAProjectAndAnAuthId(t *testing.T) {
	m := newAiChatRoutesTestModule("http://api.invalid")
	if err := m.AiChat(context.Background(), "", func(AiChatEvent) bool { return true }); err == nil {
		t.Error("empty authId: want an error")
	}
	m.t.Cfg.ProjectID = ""
	if err := m.AiChat(context.Background(), "usr_1", func(AiChatEvent) bool { return true }); err == nil {
		t.Error("no project: want an error")
	}
}
