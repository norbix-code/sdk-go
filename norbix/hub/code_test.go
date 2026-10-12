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

	"github.com/norbix-code/sdk-go/v2/norbix/internal/transport"
)

// The five cloud-function routes of the Code module (gateway campaign
// audit/code, docs/architecture/Code.Functions.md §3.3–3.4, §5.1) and the
// module method that calls each one. Every case asserts the verb, the
// fully-resolved path (version substituted) and where the parameters land —
// the query string for the GET routes, the JSON body for invoke — so a wrong
// path, a swapped verb or a parameter in the wrong place fails here instead
// of at runtime. The module file is hand-maintained since the generator was
// lost, so these cases are also its specification.
//
// Nothing leaves the process: every call goes to a local test server.

const (
	testCodeIntegrationID = "int_1"
	testCodeFunctionID    = "orders-on-update"
	testCodeTriggerID     = "trg_1"
)

type codeCase struct {
	name  string
	verb  string
	path  string
	query url.Values     // nil: no query string expected
	body  map[string]any // nil: no body expected
	call  func(ctx context.Context, m *CodeModule) error
}

func codeFunctionCases() []codeCase {
	return []codeCase{
		{"GetCodeFunctions", http.MethodGet, "/v2/code/functions",
			url.Values{"integrationId": {testCodeIntegrationID}, "prefix": {"orders-"}, "tag": {"norbix=true"}, "pageSize": {"10"}}, nil,
			func(ctx context.Context, m *CodeModule) error {
				return m.GetCodeFunctions(ctx, map[string]any{
					"integrationId": testCodeIntegrationID,
					"prefix":        "orders-",
					"tag":           "norbix=true",
					"pageSize":      10,
				}, nil)
			}},
		{"GetCodeFunctionQualifiers", http.MethodGet, "/v2/code/functions/qualifiers",
			url.Values{"integrationId": {testCodeIntegrationID}, "functionId": {testCodeFunctionID}}, nil,
			func(ctx context.Context, m *CodeModule) error {
				return m.GetCodeFunctionQualifiers(ctx, map[string]any{
					"integrationId": testCodeIntegrationID,
					"functionId":    testCodeFunctionID,
				}, nil)
			}},
		{"InvokeCodeFunction", http.MethodPost, "/v2/code/functions/invoke", nil,
			map[string]any{
				"functionRef":    map[string]any{"integrationId": testCodeIntegrationID, "functionId": testCodeFunctionID, "qualifier": "prod"},
				"payload":        map[string]any{"orderId": "6ac54a59212c29912eaddd81", "dryRun": true},
				"invocationType": "Sync",
				"runAsUserId":    "usr_service_1",
			},
			func(ctx context.Context, m *CodeModule) error {
				return m.InvokeCodeFunction(ctx, map[string]any{
					"functionRef":    map[string]any{"integrationId": testCodeIntegrationID, "functionId": testCodeFunctionID, "qualifier": "prod"},
					"payload":        map[string]any{"orderId": "6ac54a59212c29912eaddd81", "dryRun": true},
					"invocationType": "Sync",
					"runAsUserId":    "usr_service_1",
				}, nil)
			}},
		{"GetCodeFunctionEventSample", http.MethodGet, "/v2/code/functions/event-sample",
			url.Values{"source": {"database"}, "event": {"database.record.updated"}, "schemaId": {"sch_1"}}, nil,
			func(ctx context.Context, m *CodeModule) error {
				return m.GetCodeFunctionEventSample(ctx, map[string]any{
					"source":   "database",
					"event":    "database.record.updated",
					"schemaId": "sch_1",
				}, nil)
			}},
		{"GetCodeFunctionRuns", http.MethodGet, "/v2/code/functions/runs",
			url.Values{"triggerId": {testCodeTriggerID}, "limit": {"20"}}, nil,
			func(ctx context.Context, m *CodeModule) error {
				return m.GetCodeFunctionRuns(ctx, map[string]any{"triggerId": testCodeTriggerID, "limit": 20}, nil)
			}},
	}
}

func newCodeTestModule(baseURL string) *CodeModule {
	return &CodeModule{t: transport.New(&transport.Config{
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

func TestCodeFunctionEndpointsHitTheExpectedRoute(t *testing.T) {
	for _, c := range codeFunctionCases() {
		t.Run(c.name, func(t *testing.T) {
			var gotMethod, gotPath, gotRawQuery string
			var gotBody []byte
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotMethod = r.Method
				gotPath = r.URL.Path
				gotRawQuery = r.URL.RawQuery
				gotBody, _ = io.ReadAll(r.Body)
				w.WriteHeader(http.StatusNoContent)
			}))
			defer srv.Close()

			if err := c.call(context.Background(), newCodeTestModule(srv.URL)); err != nil {
				t.Fatalf("%s: unexpected error: %v", c.name, err)
			}
			if gotMethod != c.verb {
				t.Errorf("verb: got %q want %q", gotMethod, c.verb)
			}
			if gotPath != c.path {
				t.Errorf("path: got %q want %q", gotPath, c.path)
			}

			// Query: exactly the expected parameters, nothing else.
			gotQuery, _ := url.ParseQuery(gotRawQuery)
			if c.query == nil && gotRawQuery != "" {
				t.Errorf("query: got %q want none", gotRawQuery)
			}
			if c.query != nil {
				gq, _ := json.Marshal(gotQuery)
				wq, _ := json.Marshal(c.query)
				if string(gq) != string(wq) {
					t.Errorf("query: got %s want %s", gq, wq)
				}
			}

			// Body: none for the reads, the request for invoke.
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
			assertFields(t, got, c.body)
		})
	}
}

// The Code module is 38 routes: the 10 integration ones the generator wrote,
// the 23 marketplace ones (code_marketplace_test.go) and the five
// cloud-function ones above. A route added to the gateway without a method
// here, or a method added without a case, moves one of the numbers.
func TestCodeSurfaceSize(t *testing.T) {
	const wantMethods = 38
	if got := reflect.TypeOf(&CodeModule{}).NumMethod(); got != wantMethods {
		t.Errorf("CodeModule method count: got %d want %d", got, wantMethods)
	}
	const wantFunctionCases = 5
	if got := len(codeFunctionCases()); got != wantFunctionCases {
		t.Errorf("code/functions case count: got %d want %d", got, wantFunctionCases)
	}
}
