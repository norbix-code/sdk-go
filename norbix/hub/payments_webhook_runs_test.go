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

const webhookRunRow = `{"id":"pwr_1","callId":"pwc_1","env":"PROD","integrationId":"int_stripe","source":"Stripe",` +
	`"receivedOn":"2026-10-26T08:00:00Z","statusCode":400,"outcome":"Refused","reason":"signature-mismatch",` +
	`"description":"Signature does not match.","eventName":"payment_intent.succeeded","providerEventId":"evt_1",` +
	`"order":{"orderId":"ord_1","kind":"NotMatched","reason":"no order"},` +
	`"actions":[{"triggerId":"trg_email","action":"Email","status":"Sent","norbixEvent":"payments.order.paid","at":"2026-10-26T08:00:01Z"},` +
	`{"triggerId":"trg_code","action":"Code","status":"Skipped","reason":"condition false","norbixEvent":"payments.webhook.received","at":"2026-10-26T08:00:01Z"}],` +
	`"part":1,"parts":1}`

func toRequestMap(t *testing.T, v any) map[string]any {
	t.Helper()
	raw, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	var m map[string]any
	if err := json.Unmarshal(raw, &m); err != nil {
		t.Fatal(err)
	}
	return m
}

// GET /{version}/payments/webhooks/runs (gateway 86f24ecd2): every filter and
// the cursor paging land on the query, no body, and a run row with its order
// and trigger actions decodes into the generated paginated response.
func TestGetPaymentsWebhookRuns(t *testing.T) {
	var gotMethod, gotPath, gotRawQuery string
	var gotBody []byte
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod, gotPath, gotRawQuery = r.Method, r.URL.Path, r.URL.RawQuery
		gotBody, _ = io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"list":{"items":[`+webhookRunRow+`],"hasMore":true,"hasPrevious":false,"startingAfter":"pwr_1"}}`)
	}))
	defer srv.Close()

	req := toRequestMap(t, dtos.GetPaymentsWebhookRuns{
		CodeMashListPaginationRequestBase: dtos.CodeMashListPaginationRequestBase{
			PageSize:      ptr(20.0),
			StartingAfter: "pwr_prev",
		},
		IntegrationId: "int_stripe",
		EventName:     "payment_intent.succeeded",
		TriggerId:     "trg_email",
		Outcome:       "refused",
		FromUtc:       "2026-10-25T00:00:00Z",
		ToUtc:         "2026-10-26T00:00:00Z",
	})

	var out dtos.GetPaymentsWebhookRunsResponse
	if err := newPaymentsTestModule(srv.URL).GetPaymentsWebhookRuns(context.Background(), req, &out); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if gotMethod != http.MethodGet {
		t.Errorf("verb: got %q want GET", gotMethod)
	}
	if gotPath != "/v2/payments/webhooks/runs" {
		t.Errorf("path: got %q", gotPath)
	}
	gotQuery, _ := url.ParseQuery(gotRawQuery)
	wantQuery := url.Values{
		"integrationId": {"int_stripe"},
		"eventName":     {"payment_intent.succeeded"},
		"triggerId":     {"trg_email"},
		"outcome":       {"refused"},
		"fromUtc":       {"2026-10-25T00:00:00Z"},
		"toUtc":         {"2026-10-26T00:00:00Z"},
		"pageSize":      {"20"},
		"startingAfter": {"pwr_prev"},
	}
	if !reflect.DeepEqual(gotQuery, wantQuery) {
		t.Errorf("query: got %v want %v", gotQuery, wantQuery)
	}
	if len(gotBody) != 0 {
		t.Errorf("body: got %s want none", gotBody)
	}
	if len(out.List.Items) != 1 || !out.List.HasMore || out.List.StartingAfter != "pwr_1" {
		t.Fatalf("page: got %+v", out.List)
	}
	run := out.List.Items[0]
	if run.Id != "pwr_1" || run.StatusCode != 400 || run.Outcome != "Refused" || run.Reason != "signature-mismatch" {
		t.Errorf("run: got %+v", *run)
	}
	if run.Order == nil || run.Order.Kind != "NotMatched" {
		t.Errorf("order: got %+v", run.Order)
	}
	if len(run.Actions) != 2 || run.Actions[0].Status != "Sent" || run.Actions[1].Status != "Skipped" {
		t.Errorf("actions: got %+v", run.Actions)
	}
}

// GET /{version}/payments/webhooks/runs/{Id}: the run id goes in the path and
// the run plus its masked raw call decode into the generated response.
func TestGetPaymentsWebhookRun(t *testing.T) {
	var gotMethod, gotPath string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod, gotPath = r.Method, r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"item":{"run":`+webhookRunRow+`,"raw":{"callId":"pwc_1","receivedOn":"2026-10-26T08:00:00Z",`+
			`"contentType":"application/json","body":"{\"id\":\"evt_1\"}","truncated":false,"maskedFields":2,`+
			`"headers":{"Stripe-Signature":"t=1"},"answerStatus":400,"parts":1}}}`)
	}))
	defer srv.Close()

	var out dtos.GetPaymentsWebhookRunResponse
	if err := newPaymentsTestModule(srv.URL).GetPaymentsWebhookRun(context.Background(), "pwr_1", nil, &out); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if gotMethod != http.MethodGet {
		t.Errorf("verb: got %q want GET", gotMethod)
	}
	if gotPath != "/v2/payments/webhooks/runs/pwr_1" {
		t.Errorf("path: got %q", gotPath)
	}
	if out.Item == nil || out.Item.Run == nil || out.Item.Run.Id != "pwr_1" {
		t.Fatalf("item: got %+v", out.Item)
	}
	raw := out.Item.Raw
	if raw == nil || raw.CallId != "pwc_1" || raw.AnswerStatus != 400 || raw.Headers["Stripe-Signature"] != "t=1" {
		t.Errorf("raw: got %+v", raw)
	}
}
