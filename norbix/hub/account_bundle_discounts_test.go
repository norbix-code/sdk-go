package hub

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	nberrors "github.com/norbix-code/sdk-go/v2/norbix/errors"
	"github.com/norbix-code/sdk-go/v2/norbix/hub/dtos"
)

// Discount codes on the Managed Service bundles. The gateway takes the
// account from the session for both routes, so every test runs WITHOUT an
// AccountID (captureAccount). Each one asserts the verb, the fully-resolved
// path and the body, against a throw-away local server. Nothing leaves the
// process.

func TestCheckBundleDiscountSendsCodeBundleAndPeriodAndReadsTheDiscount(t *testing.T) {
	var out dtos.CheckBundleDiscountResponse
	got := captureAccount(t,
		`{"discount":{"code":"LAUNCH20","percentOff":20,"currency":"usd","months":3,"endsAt":"2027-02-18T00:00:00Z"}}`,
		func(m *AccountModule) error {
			return m.CheckBundleDiscount(context.Background(), map[string]any{
				"code":          "LAUNCH20",
				"bundle":        dtos.ManagedServiceBundleGrowth,
				"billingPeriod": dtos.SubscriptionBillingPeriodMonthly,
			}, &out)
		})

	if got.method != http.MethodPost || got.path != "/v2/account/stripe/discount/check" {
		t.Errorf("route: got %s %s want POST /v2/account/stripe/discount/check", got.method, got.path)
	}
	if len(got.query) != 0 {
		t.Errorf("no query expected, got %v", got.query)
	}
	assertFields(t, got.body, map[string]any{"code": "LAUNCH20", "bundle": "Growth", "billingPeriod": "Monthly"})

	d := out.Discount
	if d == nil || d.Code != "LAUNCH20" {
		t.Fatalf("discount not decoded: %+v", d)
	}
	if d.PercentOff == nil || *d.PercentOff != 20 {
		t.Errorf("percentOff: got %v want 20", d.PercentOff)
	}
	if d.AmountOff != nil {
		t.Errorf("amountOff was not sent, want nil, got %v", *d.AmountOff)
	}
	if d.Months == nil || *d.Months != 3 || d.EndsAt != "2027-02-18T00:00:00Z" {
		t.Errorf("duration not decoded: months %v endsAt %q", d.Months, d.EndsAt)
	}
}

func TestApplyBundleDiscountSendsTheCodeInAPostBody(t *testing.T) {
	got := captureAccount(t, "", func(m *AccountModule) error {
		return m.ApplyBundleDiscount(context.Background(), map[string]any{"code": "LAUNCH20"}, nil)
	})

	if got.method != http.MethodPost || got.path != "/v2/account/stripe/discount" {
		t.Errorf("route: got %s %s want POST /v2/account/stripe/discount", got.method, got.path)
	}
	if len(got.query) != 0 {
		t.Errorf("no query expected, got %v", got.query)
	}
	assertFields(t, got.body, map[string]any{"code": "LAUNCH20"})
}

func TestApplyBundleDiscountRefusalReachesTheCallerWithCodeAndReason(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"responseStatus":{"errorCode":"CM-ERRORS-ACCOUNT-027",` +
			`"message":"This discount code was already used.",` +
			`"errors":[{"errorCode":"CM-ERRORS-ACCOUNT-027","message":"This discount code was already used.",` +
			`"context":{"reason":"already_used"}}]}}`))
	}))
	defer srv.Close()

	err := newProjectTestModule(srv.URL, "").ApplyBundleDiscount(context.Background(),
		map[string]any{"code": "LAUNCH20"}, nil)
	if err == nil {
		t.Fatal("want an error for a refused code, got nil")
	}

	var validation *nberrors.ValidationError
	if !errors.As(err, &validation) {
		t.Fatalf("error: got %T %v want *errors.ValidationError", err, err)
	}
	var base *nberrors.Error
	if !errors.As(err, &base) {
		t.Fatalf("base error not reachable: %T %v", err, err)
	}
	if base.Status != http.StatusBadRequest || base.Code != "CM-ERRORS-ACCOUNT-027" {
		t.Errorf("error: got %d %q want 400 CM-ERRORS-ACCOUNT-027", base.Status, base.Code)
	}
	if len(base.Errors) != 1 || base.Errors[0].Context["reason"] != "already_used" {
		t.Errorf("reason not kept: %+v", base.Errors)
	}
}

func TestCreateStripeCheckoutSessionSendsTheDiscountCode(t *testing.T) {
	got := captureAccount(t, `{}`, func(m *AccountModule) error {
		return m.CreateStripeCheckoutSession(context.Background(), map[string]any{
			"subscriptionType": "ManagedService",
			"bundle":           dtos.ManagedServiceBundleStarter,
			"billingPeriod":    dtos.SubscriptionBillingPeriodMonthly,
			"discountCode":     "LAUNCH20",
		}, nil)
	})

	if got.method != http.MethodPost || got.path != "/v2/account/stripe/create-checkout-session" {
		t.Errorf("route: got %s %s want POST /v2/account/stripe/create-checkout-session", got.method, got.path)
	}
	assertFields(t, got.body, map[string]any{"bundle": "Starter", "discountCode": "LAUNCH20"})
}

func TestGetAccountStatusReadsTheDiscount(t *testing.T) {
	var out dtos.GetAccountStatusResponse
	got := captureAccount(t,
		`{"item":{"accountId":"acc_1","bundle":"Growth","discount":{"code":"FLAT10","amountOff":1000,"currency":"usd"}}}`,
		func(m *AccountModule) error {
			return m.GetAccountStatus(context.Background(), nil, &out)
		})

	if got.method != http.MethodGet || got.path != "/v2/account/status" {
		t.Errorf("route: got %s %s want GET /v2/account/status", got.method, got.path)
	}
	if out.Item == nil || out.Item.Discount == nil {
		t.Fatalf("discount not decoded: %+v", out.Item)
	}
	d := out.Item.Discount
	if d.Code != "FLAT10" || d.AmountOff == nil || *d.AmountOff != 1000 || d.Currency != "usd" {
		t.Errorf("discount: got %+v", d)
	}
	if d.Months != nil {
		t.Errorf("months not sent (forever), want nil, got %v", *d.Months)
	}
}
