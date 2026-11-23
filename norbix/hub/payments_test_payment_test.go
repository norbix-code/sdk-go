package hub

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/norbix-code/sdk-go/v2/norbix/hub/dtos"
)

// POST /{version}/payments/integrations/test-payment (review batch 3): the
// integration id and return URL go in the JSON body and the pay link decodes
// into the generated response.
func TestStartPaymentsIntegrationTestPayment(t *testing.T) {
	var gotMethod, gotPath string
	var gotBody map[string]any
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod, gotPath = r.Method, r.URL.Path
		raw, _ := io.ReadAll(r.Body)
		_ = json.Unmarshal(raw, &gotBody)
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"payUrl":"https://pay.example.test/1","reference":"ref-1","amountMinor":1,"currency":"EUR"}`)
	}))
	defer srv.Close()

	req := map[string]any{"integrationId": "int_neopay", "returnUrl": "https://app.example.test/back"}
	var out dtos.StartPaymentsIntegrationTestPaymentResponse
	if err := newPaymentsTestModule(srv.URL).StartPaymentsIntegrationTestPayment(context.Background(), req, &out); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("verb: got %q want POST", gotMethod)
	}
	if gotPath != "/v2/payments/integrations/test-payment" {
		t.Errorf("path: got %q", gotPath)
	}
	if gotBody["integrationId"] != "int_neopay" || gotBody["returnUrl"] != "https://app.example.test/back" {
		t.Errorf("body: got %v", gotBody)
	}
	if out.PayUrl != "https://pay.example.test/1" || out.Reference != "ref-1" || out.AmountMinor != 1 || out.Currency != "EUR" {
		t.Errorf("response: got %+v", out)
	}
}
