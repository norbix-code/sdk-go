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

	"github.com/norbix-code/sdk-go/v2/norbix/hub/dtos"
)

// GET /{version}/code/marketplace/functions/runs takes exactly one of
// functionId / triggerId / taskId (gateway item gw-runs, Code.Marketplace.md §9).
// The route table in code_marketplace_test.go covers triggerId; this test pins a
// function's own runs: the generated request's functionId and limit land on the
// query, nothing else, no body — and the new row fields decode into the
// generated response.
func TestGetMarketplaceFunctionRunsByFunctionId(t *testing.T) {
	var gotMethod, gotPath, gotRawQuery string
	var gotBody []byte
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod, gotPath, gotRawQuery = r.Method, r.URL.Path, r.URL.RawQuery
		gotBody, _ = io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"runs":[{"functionId":"func_1p4R3gTobGZpD6fXvHT4rd","source":"api-key","event":"invoke","succeeded":true,"status":"ok","durationMs":42,"correlationId":"0f8e2c9a4b7d4e1f9a3c5b6d7e8f9012","hasDetails":true}]}`)
	}))
	defer srv.Close()

	// The request built from the generated type, sent as the module's map.
	reqDto := dtos.GetMarketplaceFunctionRuns{FunctionId: "func_1p4R3gTobGZpD6fXvHT4rd", Limit: 50}
	raw, _ := json.Marshal(reqDto)
	var req map[string]any
	if err := json.Unmarshal(raw, &req); err != nil {
		t.Fatal(err)
	}

	var out dtos.GetMarketplaceFunctionRunsResponse
	if err := newCodeTestModule(srv.URL).GetMarketplaceFunctionRuns(context.Background(), req, &out); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if gotMethod != http.MethodGet {
		t.Errorf("verb: got %q want GET", gotMethod)
	}
	if gotPath != "/v2/code/marketplace/functions/runs" {
		t.Errorf("path: got %q", gotPath)
	}
	gotQuery, _ := url.ParseQuery(gotRawQuery)
	wantQuery := url.Values{"functionId": {"func_1p4R3gTobGZpD6fXvHT4rd"}, "limit": {"50"}}
	if !reflect.DeepEqual(gotQuery, wantQuery) {
		t.Errorf("query: got %v want %v", gotQuery, wantQuery)
	}
	if len(gotBody) != 0 {
		t.Errorf("body: got %s want none", gotBody)
	}

	if len(out.Runs) != 1 {
		t.Fatalf("runs: got %d want 1", len(out.Runs))
	}
	run := out.Runs[0]
	if run.Status != "ok" || run.CorrelationId != "0f8e2c9a4b7d4e1f9a3c5b6d7e8f9012" || !run.HasDetails || run.Source != "api-key" || run.Event != "invoke" {
		t.Errorf("row: got %+v", *run)
	}
}
