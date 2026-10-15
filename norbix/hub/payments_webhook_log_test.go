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

func newPaymentsTestModule(baseURL string) *PaymentsModule {
	return &PaymentsModule{t: transport.New(&transport.Config{
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

// GET /{version}/payments/webhooks/log (gateway campaign audit/payments, item 8):
// the generated request's integrationId and limit land on the query, no body,
// and a received-webhook row decodes into the generated response.
func TestGetPaymentsWebhookLog(t *testing.T) {
	var gotMethod, gotPath, gotRawQuery string
	var gotBody []byte
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod, gotPath, gotRawQuery = r.Method, r.URL.Path, r.URL.RawQuery
		gotBody, _ = io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"list":[{"integrationId":"int_stripe","source":"Stripe","env":"PROD","eventName":"payment_intent.succeeded","providerEventId":"evt_1","statusCode":200,"description":"Processed.","receivedOn":"2026-10-15T08:00:00Z"}]}`)
	}))
	defer srv.Close()

	reqDto := dtos.GetPaymentsWebhookLog{IntegrationId: "int_stripe", Limit: 20}
	raw, _ := json.Marshal(reqDto)
	var req map[string]any
	if err := json.Unmarshal(raw, &req); err != nil {
		t.Fatal(err)
	}

	var out dtos.GetPaymentsWebhookLogResponse
	if err := newPaymentsTestModule(srv.URL).GetPaymentsWebhookLog(context.Background(), req, &out); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if gotMethod != http.MethodGet {
		t.Errorf("verb: got %q want GET", gotMethod)
	}
	if gotPath != "/v2/payments/webhooks/log" {
		t.Errorf("path: got %q", gotPath)
	}
	gotQuery, _ := url.ParseQuery(gotRawQuery)
	wantQuery := url.Values{"integrationId": {"int_stripe"}, "limit": {"20"}}
	if !reflect.DeepEqual(gotQuery, wantQuery) {
		t.Errorf("query: got %v want %v", gotQuery, wantQuery)
	}
	if len(gotBody) != 0 {
		t.Errorf("body: got %s want none", gotBody)
	}
	if len(out.List) != 1 {
		t.Fatalf("rows: got %d want 1", len(out.List))
	}
	row := out.List[0]
	if row.EventName != "payment_intent.succeeded" || row.StatusCode != 200 || row.ProviderEventId != "evt_1" || row.Description != "Processed." {
		t.Errorf("row: got %+v", *row)
	}
}
