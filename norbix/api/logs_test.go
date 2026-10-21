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

// answerLogs starts a fake gateway that counts the calls and answers every one
// with the given status, Retry-After (none when empty) and body.
func answerLogs(t *testing.T, status int, retryAfter, body string) (*LogsModule, *int32) {
	t.Helper()
	var calls int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&calls, 1)
		w.Header().Set("Content-Type", "application/json")
		if retryAfter != "" {
			w.Header().Set("Retry-After", retryAfter)
		}
		w.WriteHeader(status)
		_, _ = io.WriteString(w, body)
	}))
	t.Cleanup(srv.Close)
	return newAPILogsTestModule(srv.URL), &calls
}

const logsQuotaBody = `{"responseStatus":{"errorCode":"CM-ERRORS-LOGGING-003","message":"Log quota used up.","errors":[{"errorCode":"CM-ERRORS-LOGGING-003","message":"Log quota used up."}]}}`

func writeOneLog(m *LogsModule) error {
	return m.WriteLogs(context.Background(),
		map[string]any{"entries": []any{map[string]any{"title": "x"}}}, nil)
}

// A 429 with Retry-After: 37 hands the 37 seconds to the caller, on the
// RateLimitError and on its base error, and the POST is still sent once.
func TestAPILogsWriteLogsQuotaRefusalCarriesRetryAfter(t *testing.T) {
	m, calls := answerLogs(t, http.StatusTooManyRequests, "37", logsQuotaBody)

	err := writeOneLog(m)

	var rl *norbixerr.RateLimitError
	if !stderrors.As(err, &rl) {
		t.Fatalf("error: got %T %v, want *errors.RateLimitError", err, err)
	}
	if rl.Base.Code != "CM-ERRORS-LOGGING-003" {
		t.Errorf("code: got %q want CM-ERRORS-LOGGING-003", rl.Base.Code)
	}
	if got := rl.RetryAfterSeconds(); got == nil || *got != 37 {
		t.Errorf("RetryAfterSeconds(): got %v want 37", got)
	}
	var base *norbixerr.Error
	if !stderrors.As(err, &base) || base.RetryAfterSeconds == nil || *base.RetryAfterSeconds != 37 {
		t.Errorf("base RetryAfterSeconds: got %v want 37", base.RetryAfterSeconds)
	}
	if n := atomic.LoadInt32(calls); n != 1 {
		t.Errorf("calls: got %d want 1 (a POST is not retried)", n)
	}
}

// A 429 that names an HTTP date gives the seconds until that date.
func TestAPILogsWriteLogsQuotaRefusalWithAnHTTPDate(t *testing.T) {
	at := time.Now().Add(120 * time.Second).UTC().Format(http.TimeFormat)
	m, _ := answerLogs(t, http.StatusTooManyRequests, at, logsQuotaBody)

	var rl *norbixerr.RateLimitError
	if err := writeOneLog(m); !stderrors.As(err, &rl) {
		t.Fatalf("error: got %T %v, want *errors.RateLimitError", err, err)
	}
	if got := rl.RetryAfterSeconds(); got == nil || *got < 115 || *got > 121 {
		t.Errorf("RetryAfterSeconds(): got %v want about 120", got)
	}
}

// A 429 without Retry-After, or with one that cannot be read, leaves it nil.
func TestAPILogsWriteLogsQuotaRefusalWithoutRetryAfter(t *testing.T) {
	for _, header := range []string{"", "soon", "-5"} {
		m, calls := answerLogs(t, http.StatusTooManyRequests, header, logsQuotaBody)

		var rl *norbixerr.RateLimitError
		if err := writeOneLog(m); !stderrors.As(err, &rl) {
			t.Fatalf("Retry-After %q: got %T %v, want *errors.RateLimitError", header, err, err)
		}
		if got := rl.RetryAfterSeconds(); got != nil {
			t.Errorf("Retry-After %q: RetryAfterSeconds() got %d want nil", header, *got)
		}
		if n := atomic.LoadInt32(calls); n != 1 {
			t.Errorf("Retry-After %q: calls got %d want 1", header, n)
		}
	}
}

// A 500 has no Retry-After: the base error leaves it nil.
func TestAPILogsWriteLogs500HasNoRetryAfter(t *testing.T) {
	m, calls := answerLogs(t, http.StatusInternalServerError, "",
		`{"responseStatus":{"errorCode":"CM-ERRORS-500","message":"Boom"}}`)

	err := writeOneLog(m)

	var base *norbixerr.Error
	if !stderrors.As(err, &base) {
		t.Fatalf("error: got %T %v, want *errors.Error", err, err)
	}
	if base.Status != http.StatusInternalServerError {
		t.Errorf("status: got %d want 500", base.Status)
	}
	if base.RetryAfterSeconds != nil {
		t.Errorf("RetryAfterSeconds: got %d want nil", *base.RetryAfterSeconds)
	}
	if n := atomic.LoadInt32(calls); n != 1 {
		t.Errorf("calls: got %d want 1 (a POST is not retried)", n)
	}
}
