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

// Trial limits on the Managed Service: trial availability per region, the
// trial waiting list and the trial fields on the checkout. The gateway takes
// the account from the session, so every test runs WITHOUT an AccountID
// (captureAccount). Each one asserts the verb, the fully-resolved path and
// the query or body, against a throw-away local server. Nothing leaves the
// process. The website-only intake route (/account/trial-waitlist/intake)
// has no method, so it has no test here.

func TestGetTrialAvailabilitySendsTheRegionInTheQueryAndReadsTheAnswer(t *testing.T) {
	var out dtos.TrialAvailabilityResponse
	got := captureAccount(t,
		`{"limitsOn":true,"region":"nb-eu-germany","trialsOpen":false,"waitlistOpen":true,"onWaitlist":false,"refusal":""}`,
		func(m *AccountModule) error {
			return m.GetTrialAvailability(context.Background(), map[string]any{
				"region": "nb-eu-germany",
			}, &out)
		})

	if got.method != http.MethodGet || got.path != "/v2/account/trial-availability" {
		t.Errorf("route: got %s %s want GET /v2/account/trial-availability", got.method, got.path)
	}
	if r := got.query["region"]; len(r) != 1 || r[0] != "nb-eu-germany" {
		t.Errorf("region query: got %v want [nb-eu-germany]", got.query)
	}
	if got.body != nil {
		t.Errorf("no body expected on GET, got %v", got.body)
	}

	if !out.LimitsOn || out.Region != "nb-eu-germany" {
		t.Errorf("limitsOn/region: got %v %q", out.LimitsOn, out.Region)
	}
	if out.TrialsOpen {
		t.Errorf("trialsOpen: got true want false")
	}
	if !out.WaitlistOpen {
		t.Errorf("waitlistOpen: got false want true")
	}
	if out.OnWaitlist || out.Refusal != "" {
		t.Errorf("onWaitlist/refusal: got %v %q", out.OnWaitlist, out.Refusal)
	}
}

func TestJoinTrialWaitlistSendsRegionAndConsentInAPostBodyAndReadsTheStatus(t *testing.T) {
	var out dtos.TrialWaitlistResponse
	got := captureAccount(t,
		`{"status":"added","message":"You are on the waiting list."}`,
		func(m *AccountModule) error {
			return m.JoinTrialWaitlist(context.Background(), map[string]any{
				"region":         "nb-eu-germany",
				"consent":        true,
				"consentVersion": "2026-11",
			}, &out)
		})

	if got.method != http.MethodPost || got.path != "/v2/account/trial-waitlist" {
		t.Errorf("route: got %s %s want POST /v2/account/trial-waitlist", got.method, got.path)
	}
	if len(got.query) != 0 {
		t.Errorf("no query expected, got %v", got.query)
	}
	assertFields(t, got.body, map[string]any{"region": "nb-eu-germany", "consent": true, "consentVersion": "2026-11"})

	if out.Status != "added" || out.Message != "You are on the waiting list." {
		t.Errorf("answer: got %q %q", out.Status, out.Message)
	}
}

func TestCreateStripeCheckoutSessionSendsTheTrialFields(t *testing.T) {
	got := captureAccount(t, `{}`, func(m *AccountModule) error {
		return m.CreateStripeCheckoutSession(context.Background(), map[string]any{
			"subscriptionType": "ManagedService",
			"bundle":           dtos.ManagedServiceBundleStarter,
			"billingPeriod":    dtos.SubscriptionBillingPeriodMonthly,
			"region":           "nb-eu-germany",
			"withoutTrial":     true,
			"trialInvite":      "inv_123",
		}, nil)
	})

	if got.method != http.MethodPost || got.path != "/v2/account/stripe/create-checkout-session" {
		t.Errorf("route: got %s %s want POST /v2/account/stripe/create-checkout-session", got.method, got.path)
	}
	assertFields(t, got.body, map[string]any{
		"bundle":       "Starter",
		"region":       "nb-eu-germany",
		"withoutTrial": true,
		"trialInvite":  "inv_123",
	})
}

func TestCreateStripeCheckoutSessionFullRegionReachesTheCallerWithTheCode(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"responseStatus":{"errorCode":"CM-ERRORS-ACCOUNT-040",` +
			`"message":"Free trials in this region are full right now.",` +
			`"errors":[{"errorCode":"CM-ERRORS-ACCOUNT-040","message":"Free trials in this region are full right now."}]}}`))
	}))
	defer srv.Close()

	err := newProjectTestModule(srv.URL, "").CreateStripeCheckoutSession(context.Background(),
		map[string]any{
			"subscriptionType": "ManagedService",
			"bundle":           dtos.ManagedServiceBundleStarter,
			"billingPeriod":    dtos.SubscriptionBillingPeriodMonthly,
			"region":           "nb-eu-germany",
		}, nil)
	if err == nil {
		t.Fatal("want an error for a full region, got nil")
	}

	var validation *nberrors.ValidationError
	if !errors.As(err, &validation) {
		t.Fatalf("error: got %T %v want *errors.ValidationError", err, err)
	}
	var base *nberrors.Error
	if !errors.As(err, &base) {
		t.Fatalf("base error not reachable: %T %v", err, err)
	}
	if base.Status != http.StatusBadRequest || base.Code != "CM-ERRORS-ACCOUNT-040" {
		t.Errorf("error: got %d %q want 400 CM-ERRORS-ACCOUNT-040", base.Status, base.Code)
	}
}
