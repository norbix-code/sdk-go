package hub

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

// The Project settings routes of matrix section 21 that had no Go method or
// no exact-route test before the settings coverage campaign (2026-11-18):
// project health and wait-active, the environment ladder, the AI plans and
// knowledge, the account usage billing and the bundle discount / extra
// regions. Each case asserts the verb, the fully-resolved path on the HUB
// (the API base is an unreachable *.invalid host), the exact query string,
// the JSON body and the key + project headers. Field names are the gateway's
// (Hub.Account Project/*.cs, Project/Settings/Ai/*.cs, Account/Stripe/*.cs).
//
// Nothing leaves the process: every call goes to a local test server.

type accountSettingsCase struct {
	name  string
	verb  string
	path  string
	query url.Values     // nil: no query string expected
	body  map[string]any // nil: no body expected
	call  func(ctx context.Context, m *AccountModule, req map[string]any) error
	req   map[string]any
}

func accountSettingsCases() []accountSettingsCase {
	const p = "/v2/account/projects/pr_1"
	const e = "/v2/account/projects/environments"
	return []accountSettingsCase{
		// General info
		{"GetProjectHealth", http.MethodGet, p + "/health", url.Values{"periodDays": {"7"}, "refresh": {"true"}}, nil,
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.GetProjectHealth(ctx, "pr_1", req, nil)
			}, map[string]any{"refresh": true, "periodDays": 7}},
		{"WaitForProjectActive", http.MethodGet, p + "/wait-active", url.Values{"timeoutSeconds": {"60"}}, nil,
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.WaitForProjectActive(ctx, "pr_1", req, nil)
			}, map[string]any{"timeoutSeconds": 60}},

		// Environments (list / create / delete live on EnvironmentsModule, below)
		{"SetEnvironmentRank", http.MethodPatch, e + "/STAGING/rank", nil, map[string]any{"rank": 2},
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.SetEnvironmentRank(ctx, "STAGING", req, nil)
			}, map[string]any{"rank": 2}},
		{"PromoteEnvironment", http.MethodPost, e + "/promote", nil,
			map[string]any{"sourceEnv": "TEST", "targetEnv": "PROD", "dryRun": true,
				"integrationChoices": []any{map[string]any{"id": "int_db", "action": "map", "mapToId": "int_prod_db"}}},
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.PromoteEnvironment(ctx, req, nil)
			}, map[string]any{"sourceEnv": "TEST", "targetEnv": "PROD", "dryRun": true,
				"integrationChoices": []any{map[string]any{"id": "int_db", "action": "map", "mapToId": "int_prod_db"}}}},
		{"RollbackPromotion", http.MethodPost, e + "/promote/rollback", nil,
			map[string]any{"targetEnv": "PROD", "fromVersion": 279, "dryRun": false, "confirmRecordLoss": true},
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.RollbackPromotion(ctx, req, nil)
			}, map[string]any{"targetEnv": "PROD", "fromVersion": 279, "dryRun": false, "confirmRecordLoss": true}},
		{"RepairEnvironmentBuiltIns", http.MethodPost, e + "/built-ins/repair", nil,
			map[string]any{"environmentName": "TEST", "dryRun": true},
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.RepairEnvironmentBuiltIns(ctx, req, nil)
			}, map[string]any{"environmentName": "TEST", "dryRun": true}},
		{"GetEnvironmentAdminUrl", http.MethodGet, e + "/TEST/admin-url", nil, nil,
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.GetEnvironmentAdminUrl(ctx, "TEST", req, nil)
			}, nil},
		{"UpdateEnvironmentAdminUrl", http.MethodPatch, e + "/TEST/admin-url", nil,
			map[string]any{"url": "https://test-admin.example.com"},
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.UpdateEnvironmentAdminUrl(ctx, "TEST", req, nil)
			}, map[string]any{"url": "https://test-admin.example.com"}},

		// AI settings tab
		{"UpdateProjectAiKnowledge", http.MethodPut, p + "/ai/knowledge", nil, map[string]any{"embedFiles": true},
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.UpdateProjectAiKnowledge(ctx, "pr_1", req, nil)
			}, map[string]any{"embedFiles": true}},
		{"GetProjectAiPlans", http.MethodGet, p + "/ai/plans", nil, nil,
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.GetProjectAiPlans(ctx, "pr_1", req, nil)
			}, nil},
		{"UpdateProjectAiPlans", http.MethodPut, p + "/ai/plans", nil,
			map[string]any{"plans": []any{map[string]any{"id": "aip_1", "name": "Free", "quotaUnit": "Credits", "monthlyQuota": 100}}},
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.UpdateProjectAiPlans(ctx, "pr_1", req, nil)
			}, map[string]any{"plans": []any{map[string]any{"id": "aip_1", "name": "Free", "quotaUnit": "Credits", "monthlyQuota": 100}}}},
		{"UpdateProjectAiPlanAssignments", http.MethodPut, p + "/ai/plans/assignments", nil,
			map[string]any{"roles": []any{map[string]any{"roleId": "rl_pro", "planId": "aip_2"}}, "defaultPlanId": "aip_1"},
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.UpdateProjectAiPlanAssignments(ctx, "pr_1", req, nil)
			}, map[string]any{"roles": []any{map[string]any{"roleId": "rl_pro", "planId": "aip_2"}}, "defaultPlanId": "aip_1"}},
		{"GetProjectAiUserPlans", http.MethodGet, p + "/ai/plans/users", nil, nil,
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.GetProjectAiUserPlans(ctx, "pr_1", req, nil)
			}, nil},
		{"CreateAiCreditPackCheckout", http.MethodPost, p + "/ai/credits/checkout", nil,
			map[string]any{"pack": "pack-10", "returnUrl": "/projects/shop/settings?tab=AI"},
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.CreateAiCreditPackCheckout(ctx, "pr_1", req, nil)
			}, map[string]any{"pack": "pack-10", "returnUrl": "/projects/shop/settings?tab=AI"}},

		// Billing
		{"GetAccountUsageBilling", http.MethodGet, "/v2/account/usage-billing", nil, nil,
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.GetAccountUsageBilling(ctx, req, nil)
			}, nil},
		{"CheckBundleDiscount", http.MethodPost, "/v2/account/stripe/discount/check", nil,
			map[string]any{"code": "LAUNCH20", "bundle": "Growth", "billingPeriod": "Annual"},
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.CheckBundleDiscount(ctx, req, nil)
			}, map[string]any{"code": "LAUNCH20", "bundle": "Growth", "billingPeriod": "Annual"}},
		{"ApplyBundleDiscount", http.MethodPost, "/v2/account/stripe/discount", nil, map[string]any{"code": "LAUNCH20"},
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.ApplyBundleDiscount(ctx, req, nil)
			}, map[string]any{"code": "LAUNCH20"}},
		{"SetExtraRegions", http.MethodPost, "/v2/account/stripe/extra-regions", nil, map[string]any{"extraRegions": 1},
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.SetExtraRegions(ctx, req, nil)
			}, map[string]any{"extraRegions": 1}},
	}
}

func newAccountSettingsTestModule(baseURL string) *AccountModule {
	return &AccountModule{t: transport.New(&transport.Config{
		ProjectID:  "proj_1",
		APIKey:     "key_1",
		BaseURLAPI: "http://api.invalid",
		BaseURLHub: baseURL,
		APIVersion: "v2",
		HubVersion: "v2",
		Timeout:    5 * time.Second,
		Env:        "PROD",
	}, nil)}
}

func TestAccountSettingsRoutesHitTheExpectedRoute(t *testing.T) {
	for _, c := range accountSettingsCases() {
		t.Run(c.name, func(t *testing.T) {
			var gotMethod, gotPath, gotRawQuery, gotAuth, gotProject string
			var gotBody []byte
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotMethod = r.Method
				gotPath = r.URL.Path
				gotRawQuery = r.URL.RawQuery
				gotAuth = r.Header.Get("Authorization")
				gotProject = r.Header.Get("norbix-project-id")
				gotBody, _ = io.ReadAll(r.Body)
				w.WriteHeader(http.StatusNoContent)
			}))
			defer srv.Close()

			if err := c.call(context.Background(), newAccountSettingsTestModule(srv.URL), c.req); err != nil {
				t.Fatalf("%s: unexpected error: %v", c.name, err)
			}
			if gotMethod != c.verb {
				t.Errorf("verb: got %q want %q", gotMethod, c.verb)
			}
			if gotPath != c.path {
				t.Errorf("path: got %q want %q", gotPath, c.path)
			}
			if gotAuth != "Bearer key_1" {
				t.Errorf("authorization: got %q want %q", gotAuth, "Bearer key_1")
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
			var got map[string]any
			if err := json.Unmarshal(gotBody, &got); err != nil {
				t.Fatalf("body is not JSON: %v (%s)", err, gotBody)
			}
			gb, _ := json.Marshal(got)
			wb, _ := json.Marshal(c.body)
			if string(gb) != string(wb) {
				t.Errorf("body: got %s want %s", gb, wb)
			}
		})
	}
}

// The three environment routes of EnvironmentsModule (the Cloud
// Environments tab lists, creates and deletes through them).
func TestEnvironmentsModuleRoutes(t *testing.T) {
	cases := []struct {
		name string
		verb string
		path string
		body map[string]any
		call func(ctx context.Context, m *EnvironmentsModule) error
	}{
		{"List", http.MethodGet, "/v2/account/projects/environments", nil,
			func(ctx context.Context, m *EnvironmentsModule) error { return m.List(ctx, nil, nil) }},
		{"Create", http.MethodPost, "/v2/account/projects/environments",
			map[string]any{"environmentName": "STAGING", "integration": map[string]any{"provider": "MongoDbConnectionString", "integrationName": "Staging database", "connectionString": "mongodb://db/staging"}},
			func(ctx context.Context, m *EnvironmentsModule) error {
				return m.Create(ctx, map[string]any{"environmentName": "STAGING", "integration": map[string]any{"provider": "MongoDbConnectionString", "integrationName": "Staging database", "connectionString": "mongodb://db/staging"}}, nil)
			}},
		{"Delete", http.MethodDelete, "/v2/account/projects/environments/STAGING", nil,
			func(ctx context.Context, m *EnvironmentsModule) error { return m.Delete(ctx, "STAGING", nil, nil) }},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var gotMethod, gotPath, gotRawQuery, gotAuth string
			var gotBody []byte
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotMethod, gotPath, gotRawQuery = r.Method, r.URL.Path, r.URL.RawQuery
				gotAuth = r.Header.Get("Authorization")
				gotBody, _ = io.ReadAll(r.Body)
				w.WriteHeader(http.StatusNoContent)
			}))
			defer srv.Close()

			m := &EnvironmentsModule{t: newAccountSettingsTestModule(srv.URL).t}
			if err := c.call(context.Background(), m); err != nil {
				t.Fatalf("%s: unexpected error: %v", c.name, err)
			}
			if gotMethod != c.verb || gotPath != c.path {
				t.Errorf("route: got %s %s want %s %s", gotMethod, gotPath, c.verb, c.path)
			}
			if gotRawQuery != "" {
				t.Errorf("query: got %q want none", gotRawQuery)
			}
			if gotAuth != "Bearer key_1" {
				t.Errorf("authorization: got %q", gotAuth)
			}
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
			g2, _ := json.Marshal(got)
			w2, _ := json.Marshal(c.body)
			if string(g2) != string(w2) {
				t.Errorf("body: got %s want %s", g2, w2)
			}
		})
	}
}

// 18 AccountModule routes; a changed count means a route arrived or left
// untested.
func TestAccountSettingsSurfaceSize(t *testing.T) {
	const want = 18
	if got := len(accountSettingsCases()); got != want {
		t.Errorf("account settings endpoint count: got %d want %d", got, want)
	}
}
