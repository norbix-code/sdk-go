package hub

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/norbix-code/sdk-go/v2/norbix/hub/dtos"
	"github.com/norbix-code/sdk-go/v2/norbix/internal/transport"
)

// The Hub log routes added for the Logs campaign (finding L12: the Go SDK had
// no GetLogs, GetLog, CleanLogs, GetLogSettings or SaveLogSettings) and the
// integration test every saved integration needs. Each test asserts the verb,
// the fully-resolved path, the query string or JSON body, and decodes the
// answer into the generated response type.
//
// Nothing leaves the process: every call goes to a local test server.

type logsSeen struct {
	method, path, rawQuery, project string
	body                            []byte
}

// newLogsTestServer answers every call with answer and records what it got.
func newLogsTestServer(t *testing.T, answer string) (*LogsModule, *logsSeen) {
	t.Helper()
	seen := &logsSeen{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		seen.method, seen.path, seen.rawQuery = r.Method, r.URL.Path, r.URL.RawQuery
		seen.project = r.Header.Get("norbix-project-id")
		seen.body, _ = io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, answer)
	}))
	t.Cleanup(srv.Close)
	return &LogsModule{t: transport.New(&transport.Config{
		ProjectID:  "proj_1",
		AccountID:  "acct_1",
		APIKey:     "key_1",
		BaseURLAPI: "http://api.invalid",
		BaseURLHub: srv.URL,
		APIVersion: "v2",
		HubVersion: "v2",
		Timeout:    5 * time.Second,
		Env:        "PROD",
	}, nil)}, seen
}

func (s *logsSeen) expect(t *testing.T, verb, path string) {
	t.Helper()
	if s.method != verb {
		t.Errorf("verb: got %q want %q", s.method, verb)
	}
	if s.path != path {
		t.Errorf("path: got %q want %q", s.path, path)
	}
	if s.project != "proj_1" {
		t.Errorf("project header: got %q want proj_1", s.project)
	}
}

func (s *logsSeen) expectNoBody(t *testing.T) {
	t.Helper()
	if len(s.body) != 0 {
		t.Errorf("body: got %s want none", s.body)
	}
}

func (s *logsSeen) expectBody(t *testing.T, want map[string]any) {
	t.Helper()
	var got any
	if err := json.Unmarshal(s.body, &got); err != nil {
		t.Fatalf("body is not JSON: %v (%s)", err, s.body)
	}
	gs, _ := json.Marshal(got)
	ws, _ := json.Marshal(want)
	if string(gs) != string(ws) {
		t.Errorf("body: got %s want %s", gs, ws)
	}
}

const logEntryJSON = `{"id":"6714a0c2f1e2d3c4b5a69788","timestamp":"2026-10-21T08:00:00Z","module":"App","level":"Error","eventCode":"app:checkout.failed","title":"Checkout failed","message":"card declined","correlationId":"req_1","env":"PROD","origin":"app","meta":{"orderId":"ord_1"}}`

func TestLogsGetLogs(t *testing.T) {
	m, seen := newLogsTestServer(t, `{"list":{"items":[`+logEntryJSON+`],"hasMore":true,"startingAfter":"6714a0c2f1e2d3c4b5a69788"}}`)
	req := map[string]any{
		"level":            "Error",
		"logCorrelationId": "req_1",
		"eventCode":        "app:checkout.failed",
		"metaKey":          "orderId",
		"metaValue":        "ord_1",
		"pageSize":         20,
	}

	var out dtos.GetLogsResponse
	if err := m.GetLogs(context.Background(), req, &out); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	seen.expect(t, http.MethodGet, "/v2/logs")
	seen.expectNoBody(t)
	gotQuery, _ := url.ParseQuery(seen.rawQuery)
	wantQuery := url.Values{
		"level":            {"Error"},
		"logCorrelationId": {"req_1"},
		"eventCode":        {"app:checkout.failed"},
		"metaKey":          {"orderId"},
		"metaValue":        {"ord_1"},
		"pageSize":         {"20"},
	}
	if !reflect.DeepEqual(gotQuery, wantQuery) {
		t.Errorf("query: got %v want %v", gotQuery, wantQuery)
	}
	if len(out.List.Items) != 1 || !out.List.HasMore {
		t.Fatalf("list: got %+v", out.List)
	}
	if e := out.List.Items[0]; e.Title != "Checkout failed" || e.Origin != "app" || e.Meta["orderId"] != "ord_1" {
		t.Errorf("entry: got %+v", *e)
	}
}

func TestLogsGetLog(t *testing.T) {
	m, seen := newLogsTestServer(t, `{"item":`+logEntryJSON+`}`)

	var out dtos.GetLogResponse
	if err := m.GetLog(context.Background(), "6714a0c2f1e2d3c4b5a69788", nil, &out); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	seen.expect(t, http.MethodGet, "/v2/logs/6714a0c2f1e2d3c4b5a69788")
	seen.expectNoBody(t)
	if seen.rawQuery != "" {
		t.Errorf("query: got %q want none", seen.rawQuery)
	}
	if out.Item == nil || out.Item.Id != "6714a0c2f1e2d3c4b5a69788" || out.Item.EventCode != "app:checkout.failed" {
		t.Errorf("item: got %+v", out.Item)
	}
}

func TestLogsCleanLogs(t *testing.T) {
	m, seen := newLogsTestServer(t, `{}`)

	var out dtos.CleanLogsResponse
	if err := m.CleanLogs(context.Background(), nil, &out); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Nothing to send: the env the entries are removed for is the request's
	// own (client Env or the norbix-env header), so the POST has no body.
	seen.expect(t, http.MethodPost, "/v2/logs/clean")
	seen.expectNoBody(t)
}

func TestLogsGetLogSettings(t *testing.T) {
	m, seen := newLogsTestServer(t, `{"skipCloudDashboardLogs":true,"skipHttpBodyMeta":false,"aiChatLoggingEnabled":true,"hasNorbixLogging":true}`)

	var out dtos.GetLogSettingsResponse
	if err := m.GetLogSettings(context.Background(), nil, &out); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	seen.expect(t, http.MethodGet, "/v2/logs/settings")
	seen.expectNoBody(t)
	if !out.SkipCloudDashboardLogs || out.SkipHttpBodyMeta || !out.AiChatLoggingEnabled || !out.HasNorbixLogging {
		t.Errorf("settings: got %+v", out)
	}
}

// A false must reach the gateway. The typed request (*bool fields) is proven
// in norbix/optional_test.go; this one sends a plain map.
func TestLogsSaveLogSettings(t *testing.T) {
	m, seen := newLogsTestServer(t, `{}`)
	req := map[string]any{
		"skipCloudDashboardLogs": true,
		"skipHttpBodyMeta":       false,
		"aiChatLoggingEnabled":   false,
	}

	var out dtos.SaveLogSettingsResponse
	if err := m.SaveLogSettings(context.Background(), req, &out); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	seen.expect(t, http.MethodPost, "/v2/logs/settings")
	seen.expectBody(t, req)
}

func TestLogsTestLoggingIntegration(t *testing.T) {
	m, seen := newLogsTestServer(t, `{"items":[{"operation":"Connection","result":"Success"}]}`)
	req := map[string]any{"integrationId": "int_mongo"}

	var out dtos.TestLoggingIntegrationResponse
	if err := m.TestLoggingIntegration(context.Background(), req, &out); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	seen.expect(t, http.MethodPost, "/v2/logs/integrations/test")
	seen.expectBody(t, req)
	if len(out.Items) != 1 {
		t.Errorf("items: got %d want 1", len(out.Items))
	}
}

// A MongoDB integration is saved from the generated request type; the
// provider, the connection string and the optional database name reach the
// body under their wire names.
func TestLogsSaveMongoDbLoggingIntegration(t *testing.T) {
	m, seen := newLogsTestServer(t, `{"id":"int_mongo"}`)
	integration := dtos.MongoDbLoggingIntegrationRequest{
		LoggingIntegrationRequest: dtos.LoggingIntegrationRequest{
			IntegrationName: "Mongo logs",
			IsEnabled:       true,
		},
		Provider:         dtos.LoggingProviderMongoDB,
		ConnectionString: "mongodb://logs.example.test:27017",
		DatabaseName:     "app_logs",
	}

	var out dtos.IdResponse
	if err := m.SaveLoggingIntegration(context.Background(), map[string]any{"integration": integration}, &out); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	seen.expect(t, http.MethodPost, "/v2/logs/integrations")
	seen.expectBody(t, map[string]any{"integration": map[string]any{
		"provider":         "MongoDB",
		"integrationName":  "Mongo logs",
		"isEnabled":        true,
		"databaseName":     "app_logs",
		"connectionString": "mongodb://logs.example.test:27017",
	}})
	if out.Id != "int_mongo" {
		t.Errorf("id: got %q", out.Id)
	}
}
