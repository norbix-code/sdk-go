package api

import (
	"context"
	"encoding/json"
	stderrors "errors"
	"io"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"github.com/norbix-code/sdk-go/v2/norbix/api/dtos"
	norbixerr "github.com/norbix-code/sdk-go/v2/norbix/errors"
	"github.com/norbix-code/sdk-go/v2/norbix/internal/transport"
)

// POST /{version}/logs on the API (gateway campaign audit/logging, Logs stage
// A): an app writes its own entries. Nothing leaves the process: every call
// goes to a local test server.

func newAPILogsTestModule(baseURL string) *LogsModule {
	return &LogsModule{t: transport.New(&transport.Config{
		ProjectID:  "proj_1",
		AccountID:  "acct_1",
		APIKey:     "key_1",
		BaseURLAPI: baseURL,
		BaseURLHub: "http://hub.invalid",
		APIVersion: "v2",
		HubVersion: "v2",
		Timeout:    5 * time.Second,
		Env:        "PROD",
	}, nil)}
}

// The entries are built from the generated types and go out as the JSON body;
// the answer decodes into the generated response.
func TestAPILogsWriteLogs(t *testing.T) {
	var gotMethod, gotPath, gotRawQuery, gotProject string
	var gotBody []byte
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod, gotPath, gotRawQuery = r.Method, r.URL.Path, r.URL.RawQuery
		gotProject = r.Header.Get("norbix-project-id")
		gotBody, _ = io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"accepted":2}`)
	}))
	defer srv.Close()

	entries := []*dtos.WriteLogEntry{
		{
			Level:         "error",
			Title:         "Checkout failed",
			Message:       "card declined",
			EventCode:     "checkout.failed",
			CorrelationId: "req_1",
			Timestamp:     "2026-10-21T08:00:00Z",
			Meta:          map[string]string{"orderId": "ord_1"},
		},
		{Title: "Checkout started"},
	}

	var out dtos.WriteLogsResponse
	err := newAPILogsTestModule(srv.URL).WriteLogs(context.Background(), map[string]any{"entries": entries}, &out)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if gotMethod != http.MethodPost {
		t.Errorf("verb: got %q want POST", gotMethod)
	}
	if gotPath != "/v2/logs" {
		t.Errorf("path: got %q want /v2/logs", gotPath)
	}
	if gotRawQuery != "" {
		t.Errorf("query: got %q want none", gotRawQuery)
	}
	if gotProject != "proj_1" {
		t.Errorf("project header: got %q want proj_1", gotProject)
	}
	var got any
	if err := json.Unmarshal(gotBody, &got); err != nil {
		t.Fatalf("body is not JSON: %v (%s)", err, gotBody)
	}
	want := map[string]any{"entries": []any{
		map[string]any{
			"level":         "error",
			"title":         "Checkout failed",
			"message":       "card declined",
			"eventCode":     "checkout.failed",
			"correlationId": "req_1",
			"timestamp":     "2026-10-21T08:00:00Z",
			"meta":          map[string]any{"orderId": "ord_1"},
		},
		map[string]any{"title": "Checkout started"},
	}}
	gs, _ := json.Marshal(got)
	ws, _ := json.Marshal(want)
	if string(gs) != string(ws) {
		t.Errorf("body: got %s want %s", gs, ws)
	}
	if out.Accepted != 2 {
		t.Errorf("accepted: got %v want 2", out.Accepted)
	}
}

// The project's per-minute quota is used up: the gateway refuses the whole
// call with 429 CM-ERRORS-LOGGING-003. The SDK hands back a RateLimitError
// carrying the gateway's code and does not resend the POST.
func TestAPILogsWriteLogsQuotaRefusal(t *testing.T) {
	var calls int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&calls, 1)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Retry-After", "17")
		w.WriteHeader(http.StatusTooManyRequests)
		_, _ = io.WriteString(w, `{"responseStatus":{"errorCode":"CM-ERRORS-LOGGING-003","message":"Log quota used up.","errors":[{"errorCode":"CM-ERRORS-LOGGING-003","message":"Log quota used up."}]}}`)
	}))
	defer srv.Close()

	err := newAPILogsTestModule(srv.URL).WriteLogs(context.Background(),
		map[string]any{"entries": []any{map[string]any{"title": "x"}}}, nil)

	var rl *norbixerr.RateLimitError
	if !stderrors.As(err, &rl) {
		t.Fatalf("error: got %T %v, want *errors.RateLimitError", err, err)
	}
	if rl.Base.Status != http.StatusTooManyRequests || rl.Base.Code != "CM-ERRORS-LOGGING-003" {
		t.Errorf("error: got status %d code %q", rl.Base.Status, rl.Base.Code)
	}
	if n := atomic.LoadInt32(&calls); n != 1 {
		t.Errorf("calls: got %d want 1 (a POST is not retried)", n)
	}
}
