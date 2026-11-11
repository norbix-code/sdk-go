package hub

import (
	"context"
	"net/http"
	"testing"

	"github.com/norbix-code/sdk-go/v2/norbix/hub/dtos"
)

// Managed Service bundles and the account's Norbix-managed Atlas clusters.
// The gateway takes the account from the session for all four routes, so
// every test runs WITHOUT an AccountID (captureAccount). Each one asserts the
// verb, the fully-resolved path and where the parameters go, against a
// throw-away local server. Nothing leaves the process.

func TestGetManagedServiceBundlesReadsTheBundlesAndPrices(t *testing.T) {
	var out dtos.GetManagedServiceBundlesResponse
	got := captureAccount(t,
		`{"bundles":[{"bundle":"Starter","projectCap":3,`+
			`"monthly":{"amount":4900,"currency":"usd"},"annual":{"amount":49000,"currency":"usd"}}]}`,
		func(m *AccountModule) error {
			return m.GetManagedServiceBundles(context.Background(), nil, &out)
		})

	if got.method != http.MethodGet || got.path != "/v2/account/stripe/bundles" {
		t.Errorf("route: got %s %s want GET /v2/account/stripe/bundles", got.method, got.path)
	}
	if len(got.query) != 0 || len(got.body) != 0 {
		t.Errorf("no parameters expected: query %v body %v", got.query, got.body)
	}
	if len(out.Bundles) != 1 {
		t.Fatalf("bundles not decoded: %+v", out.Bundles)
	}
	b := out.Bundles[0]
	if b.Bundle != dtos.ManagedServiceBundleStarter || b.ProjectCap != 3 {
		t.Errorf("bundle: got %q with limit %v want Starter with limit 3", b.Bundle, b.ProjectCap)
	}
	if b.Monthly == nil || b.Monthly.Amount != 4900 || b.Monthly.Currency != "usd" {
		t.Errorf("monthly price not decoded: %+v", b.Monthly)
	}
	if b.Annual == nil || b.Annual.Amount != 49000 {
		t.Errorf("annual price not decoded: %+v", b.Annual)
	}
}

func TestChangeManagedServiceBundleSendsBundleAndPeriodInAPostBody(t *testing.T) {
	got := captureAccount(t, "", func(m *AccountModule) error {
		return m.ChangeManagedServiceBundle(context.Background(), map[string]any{
			"bundle":        dtos.ManagedServiceBundleGrowth,
			"billingPeriod": dtos.SubscriptionBillingPeriodAnnual,
		}, nil)
	})

	if got.method != http.MethodPost || got.path != "/v2/account/stripe/change-bundle" {
		t.Errorf("route: got %s %s want POST /v2/account/stripe/change-bundle", got.method, got.path)
	}
	if len(got.query) != 0 {
		t.Errorf("no query expected, got %v", got.query)
	}
	assertFields(t, got.body, map[string]any{"bundle": "Growth", "billingPeriod": "Annual"})
}

func TestGetFlexClustersReadsClustersDatabasesAndUsage(t *testing.T) {
	var out dtos.GetFlexClustersResponse
	got := captureAccount(t,
		`{"clusters":[{"id":"fc_1","displayName":"Main","atlasClusterName":"nbx-1","norbixRegionCode":"eu-central",`+
			`"tierCode":"FLEX","status":"Active","nearLimits":["storage"],`+
			`"databases":[{"projectId":"proj_1","projectName":"Shop","env":"PROD","integrationId":"int_1","integrationName":"Main db"}],`+
			`"usage":{"databaseCount":1,"storageBytes":2048}}],`+
			`"canCreateNew":false,"cannotCreateReason":"ClusterLimitReached"}`,
		func(m *AccountModule) error {
			return m.GetFlexClusters(context.Background(), nil, &out)
		})

	if got.method != http.MethodGet || got.path != "/v2/account/flex-clusters" {
		t.Errorf("route: got %s %s want GET /v2/account/flex-clusters", got.method, got.path)
	}
	if len(got.query) != 0 || len(got.body) != 0 {
		t.Errorf("no parameters expected: query %v body %v", got.query, got.body)
	}
	if out.CanCreateNew || out.CannotCreateReason != "ClusterLimitReached" {
		t.Errorf("create flag: got %v %q want false ClusterLimitReached", out.CanCreateNew, out.CannotCreateReason)
	}
	if len(out.Clusters) != 1 {
		t.Fatalf("clusters not decoded: %+v", out.Clusters)
	}
	c := out.Clusters[0]
	if c.Id != "fc_1" || c.TierCode != "FLEX" || c.Status != "Active" {
		t.Errorf("cluster: got %+v", c)
	}
	if len(c.Databases) != 1 || c.Databases[0].IntegrationId != "int_1" || c.Databases[0].Env != "PROD" {
		t.Errorf("databases not decoded: %+v", c.Databases)
	}
	if c.Usage == nil || c.Usage.DatabaseCount != 1 || c.Usage.StorageBytes == nil || *c.Usage.StorageBytes != 2048 {
		t.Errorf("usage not decoded: %+v", c.Usage)
	}
	if c.Usage != nil && c.Usage.Connections != nil {
		t.Errorf("connections were not sent, want nil, got %v", *c.Usage.Connections)
	}
}

func TestUpgradeFlexClusterPutsTheIdInThePathAndTheTierInTheBody(t *testing.T) {
	got := captureAccount(t, "", func(m *AccountModule) error {
		return m.UpgradeFlexCluster(context.Background(), "fc_1", map[string]any{"tierCode": "M10"}, nil)
	})

	if got.method != http.MethodPost || got.path != "/v2/account/flex-clusters/fc_1/upgrade" {
		t.Errorf("route: got %s %s want POST /v2/account/flex-clusters/fc_1/upgrade", got.method, got.path)
	}
	if len(got.query) != 0 {
		t.Errorf("no query expected, got %v", got.query)
	}
	assertFields(t, got.body, map[string]any{"tierCode": "M10"})
	if _, inBody := got.body["flexClusterId"]; inBody {
		t.Errorf("flexClusterId belongs in the path, not in the body: %v", got.body)
	}
}

func TestCreateStripeCheckoutSessionSendsBundleAndBillingPeriod(t *testing.T) {
	var out dtos.CreateStripeCheckoutSessionResponse
	got := captureAccount(t, `{}`, func(m *AccountModule) error {
		return m.CreateStripeCheckoutSession(context.Background(), map[string]any{
			"subscriptionType": "ManagedService",
			"bundle":           dtos.ManagedServiceBundleStarter,
			"billingPeriod":    dtos.SubscriptionBillingPeriodMonthly,
		}, &out)
	})

	if got.method != http.MethodPost || got.path != "/v2/account/stripe/create-checkout-session" {
		t.Errorf("route: got %s %s want POST /v2/account/stripe/create-checkout-session", got.method, got.path)
	}
	assertFields(t, got.body, map[string]any{"bundle": "Starter", "billingPeriod": "Monthly"})
}
