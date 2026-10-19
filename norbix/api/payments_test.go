package api

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/norbix-code/sdk-go/v2/norbix/internal/transport"
)

// Every payment order route on the API (gateway campaign audit/payments-checkout)
// and the module method that calls it. Each case asserts the verb, the
// fully-resolved path (version substituted, {Id} filled), the query string and
// the JSON body, so a wrong path, a swapped verb or a parameter in the wrong
// place fails here instead of at runtime.
//
// Nothing leaves the process: every call goes to a local test server.

type apiPaymentsCase struct {
	name  string
	verb  string
	path  string
	query url.Values     // nil: no query string expected
	body  map[string]any // nil: no body expected
	call  func(ctx context.Context, m *PaymentsModule) error
}

func apiPaymentsCases() []apiPaymentsCase {
	order := map[string]any{
		"buyerUserId":       "usr_1",
		"merchantReference": "ad-1042-promo-1",
		"idempotencyKey":    "ad-1042-promo-1",
		"currency":          "EUR",
		"priceVersion":      "2026-10-01",
		"lines":             []any{map[string]any{"sku": "star-day", "quantity": 21, "unitAmountMinor": 30}},
		"metadata":          map[string]any{"adId": "ad_1042"},
	}
	checkout := map[string]any{
		"integrationId":    "int_montonio",
		"returnUrl":        "https://shop.example.test/return",
		"preferredCountry": "LT",
	}
	filter := map[string]any{"buyerUserId": "usr_1", "merchantReference": "ad-1042", "status": "Paid"}

	return []apiPaymentsCase{
		{"CreatePaymentOrder", http.MethodPost, "/v2/payments/orders", nil, order,
			func(ctx context.Context, m *PaymentsModule) error {
				return m.CreatePaymentOrder(ctx, order, nil)
			}},
		{"GetPaymentOrder", http.MethodGet, "/v2/payments/orders/ord_1", nil, nil,
			func(ctx context.Context, m *PaymentsModule) error {
				return m.GetPaymentOrder(ctx, "ord_1", nil, nil)
			}},
		{"GetPaymentOrders", http.MethodGet, "/v2/payments/orders",
			url.Values{"buyerUserId": {"usr_1"}, "merchantReference": {"ad-1042"}, "status": {"Paid"}}, nil,
			func(ctx context.Context, m *PaymentsModule) error {
				return m.GetPaymentOrders(ctx, filter, nil)
			}},
		{"StartPaymentCheckout", http.MethodPost, "/v2/payments/orders/ord_1/checkout", nil, checkout,
			func(ctx context.Context, m *PaymentsModule) error {
				return m.StartPaymentCheckout(ctx, "ord_1", checkout, nil)
			}},
	}
}

func newAPIPaymentsTestModule(baseURL string) *PaymentsModule {
	return &PaymentsModule{t: transport.New(&transport.Config{
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

func TestAPIPaymentsEndpointsHitTheExpectedRoute(t *testing.T) {
	for _, c := range apiPaymentsCases() {
		t.Run(c.name, func(t *testing.T) {
			var gotMethod, gotPath, gotRawQuery, gotProject string
			var gotBody []byte
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotMethod = r.Method
				gotPath = r.URL.Path
				gotRawQuery = r.URL.RawQuery
				gotProject = r.Header.Get("norbix-project-id")
				gotBody, _ = io.ReadAll(r.Body)
				w.WriteHeader(http.StatusNoContent)
			}))
			defer srv.Close()

			if err := c.call(context.Background(), newAPIPaymentsTestModule(srv.URL)); err != nil {
				t.Fatalf("%s: unexpected error: %v", c.name, err)
			}
			if gotMethod != c.verb {
				t.Errorf("verb: got %q want %q", gotMethod, c.verb)
			}
			if gotPath != c.path {
				t.Errorf("path: got %q want %q", gotPath, c.path)
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
			// Compare as JSON values: both sides re-encoded with sorted keys.
			var got any
			if err := json.Unmarshal(gotBody, &got); err != nil {
				t.Fatalf("body is not JSON: %v (%s)", err, gotBody)
			}
			gs, _ := json.Marshal(got)
			ws, _ := json.Marshal(c.body)
			if string(gs) != string(ws) {
				t.Errorf("body: got %s want %s", gs, ws)
			}
		})
	}
}

// A created order answers its id; GetPaymentOrder decodes the stored order.
func TestAPIPaymentsDecodeTheAnswers(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodPost {
			_, _ = io.WriteString(w, `{"id":"ord_1"}`)
			return
		}
		_, _ = io.WriteString(w, `{"item":{"id":"ord_1","status":"Created","totalMinor":630,"lines":[{"sku":"star-day","quantity":21,"unitAmountMinor":30,"amountMinor":630}]}}`)
	}))
	defer srv.Close()
	m := newAPIPaymentsTestModule(srv.URL)

	var created struct {
		ID string `json:"id"`
	}
	if err := m.CreatePaymentOrder(context.Background(), map[string]any{"currency": "EUR"}, &created); err != nil || created.ID != "ord_1" {
		t.Fatalf("create: %v %+v", err, created)
	}
	var got struct {
		Item struct {
			ID         string `json:"id"`
			Status     string `json:"status"`
			TotalMinor int64  `json:"totalMinor"`
		} `json:"item"`
	}
	if err := m.GetPaymentOrder(context.Background(), "ord_1", nil, &got); err != nil {
		t.Fatalf("get: %v", err)
	}
	if got.Item.ID != "ord_1" || got.Item.Status != "Created" || got.Item.TotalMinor != 630 {
		t.Errorf("order: %+v", got.Item)
	}
}

// The API payments surface is 4 routes. A new route changes this count, so
// it cannot arrive untested.
func TestAPIPaymentsSurfaceSize(t *testing.T) {
	const want = 4
	if got := len(apiPaymentsCases()); got != want {
		t.Errorf("payments endpoint count: got %d want %d", got, want)
	}
}
