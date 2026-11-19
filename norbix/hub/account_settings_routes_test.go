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

// Every AccountModule route of matrix section 21 (Project settings): general
// info, brand, access, regions, languages, the environment ladder, admin
// portal, legal + expose switches, AI, termination and billing (the three
// EnvironmentsModule routes are in TestEnvironmentsModuleRoutes below, the
// public API-host routes in norbix/api/public_test.go). The older tables of
// these routes (account_project_test.go, account_ai_test.go) point both hosts
// at one server and check neither query nor body. Each case asserts the verb,
// the fully-resolved path on the HUB
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
	fileRef := map[string]any{"resource": map[string]any{"id": "fl_1"}, "integrationId": "int_files", "provider": "Local", "path": "brand/logo.png", "isPublic": true}
	assistant := map[string]any{"name": "Support", "systemPrompt": "Be brief.", "toolsets": []any{"own:records"}, "memoryEnabled": false, "isDefault": true}
	return []accountSettingsCase{
		// General info
		{"GetProjects", http.MethodGet, "/v2/account/projects", nil, nil,
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.GetProjects(ctx, req, nil)
			}, nil},
		{"CreateProject", http.MethodPost, "/v2/account/projects", nil, map[string]any{"projectName": "Shop", "primaryRegion": "nb-eu-germany", "description": "Orders", "integration": map[string]any{"provider": "MongoDbConnectionString", "integrationName": "Shop database", "connectionString": "mongodb://db/shop"}},
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.CreateProject(ctx, req, nil)
			}, map[string]any{"projectName": "Shop", "primaryRegion": "nb-eu-germany", "description": "Orders", "integration": map[string]any{"provider": "MongoDbConnectionString", "integrationName": "Shop database", "connectionString": "mongodb://db/shop"}}},
		{"GetProject", http.MethodGet, "/v2/account/projects/pr_1", nil, nil,
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.GetProject(ctx, "pr_1", req, nil)
			}, nil},
		{"GetProjectTokens", http.MethodGet, p + "/tokens", url.Values{"recipientId": {"usr_1"}}, nil,
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.GetProjectTokens(ctx, "pr_1", req, nil)
			}, map[string]any{"recipientId": "usr_1"}},
		// Brand
		{"UpdateProjectName", http.MethodPatch, p + "/settings/name", nil, map[string]any{"name": "Shop"},
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.UpdateProjectName(ctx, "pr_1", req, nil)
			}, map[string]any{"name": "Shop"}},
		{"UpdateProjectDescription", http.MethodPatch, p + "/settings/description", nil, map[string]any{"description": "Orders"},
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.UpdateProjectDescription(ctx, "pr_1", req, nil)
			}, map[string]any{"description": "Orders"}},
		{"UpdateProjectUrl", http.MethodPatch, p + "/settings/url", nil, map[string]any{"url": "https://shop.example.com"},
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.UpdateProjectUrl(ctx, "pr_1", req, nil)
			}, map[string]any{"url": "https://shop.example.com"}},
		{"UpdateProjectIcon", http.MethodPatch, p + "/settings/icon", nil, map[string]any{"fileResource": fileRef},
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.UpdateProjectIcon(ctx, "pr_1", req, nil)
			}, map[string]any{"fileResource": fileRef}},
		{"UpdateProjectLogo", http.MethodPatch, p + "/settings/logo", nil, map[string]any{"fileResource": fileRef},
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.UpdateProjectLogo(ctx, "pr_1", req, nil)
			}, map[string]any{"fileResource": fileRef}},
		{"UpdateProjectMainColor", http.MethodPatch, p + "/settings/main-color", nil, map[string]any{"color": "#1F6FEB"},
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.UpdateProjectMainColor(ctx, "pr_1", req, nil)
			}, map[string]any{"color": "#1F6FEB"}},
		{"UpdateProjectAccentColor", http.MethodPatch, p + "/settings/accent-color", nil, map[string]any{"color": "#F78166"},
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.UpdateProjectAccentColor(ctx, "pr_1", req, nil)
			}, map[string]any{"color": "#F78166"}},
		// Access (CORS) and regions
		{"UpdateProjectAllowedOrigins", http.MethodPatch, p + "/settings/origins", nil, map[string]any{"origins": []any{"https://app.example.com"}, "removeAdminPortalOrigin": true},
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.UpdateProjectAllowedOrigins(ctx, "pr_1", req, nil)
			}, map[string]any{"origins": []any{"https://app.example.com"}, "removeAdminPortalOrigin": true}},
		{"UpdateProjectRegions", http.MethodPatch, p + "/settings/regions", nil, map[string]any{"additionalRegions": []any{"nb-us-east"}},
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.UpdateProjectRegions(ctx, "pr_1", req, nil)
			}, map[string]any{"additionalRegions": []any{"nb-us-east"}}},
		// Languages
		{"UpdateProjectLanguages", http.MethodPatch, p + "/settings/languages", nil, map[string]any{"languages": []any{"en", "lt"}},
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.UpdateProjectLanguages(ctx, "pr_1", req, nil)
			}, map[string]any{"languages": []any{"en", "lt"}}},
		{"UpdateProjectDefaultLanguage", http.MethodPatch, p + "/settings/default-language", nil, map[string]any{"defaultLanguage": "lt"},
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.UpdateProjectDefaultLanguage(ctx, "pr_1", req, nil)
			}, map[string]any{"defaultLanguage": "lt"}},
		{"CheckProjectLanguages", http.MethodPost, p + "/settings/languages/check", nil, map[string]any{"languages": []any{"en", "de"}, "defaultLanguage": "en"},
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.CheckProjectLanguages(ctx, "pr_1", req, nil)
			}, map[string]any{"languages": []any{"en", "de"}, "defaultLanguage": "en"}},
		// Admin portal
		{"GetAdminPortalStructure", http.MethodGet, p + "/admin-portal/structure", nil, nil,
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.GetAdminPortalStructure(ctx, "pr_1", req, nil)
			}, nil},
		{"AssignAdminPortalServiceUser", http.MethodPut, p + "/settings/admin-portal/service-user", nil, map[string]any{"serviceUserId": "su_1"},
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.AssignAdminPortalServiceUser(ctx, "pr_1", req, nil)
			}, map[string]any{"serviceUserId": "su_1"}},
		{"UpdateProjectAdminUrl", http.MethodPatch, p + "/settings/admin-url", nil, map[string]any{"url": "https://admin.example.com"},
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.UpdateProjectAdminUrl(ctx, "pr_1", req, nil)
			}, map[string]any{"url": "https://admin.example.com"}},
		// Legal + expose switches
		{"UpdateProjectLegalDocuments", http.MethodPatch, p + "/settings/legal", nil, map[string]any{"termsMarkdown": "# Terms", "privacyMarkdown": "# Privacy"},
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.UpdateProjectLegalDocuments(ctx, "pr_1", req, nil)
			}, map[string]any{"termsMarkdown": "# Terms", "privacyMarkdown": "# Privacy"}},
		// AI settings
		{"GetProjectAiSettings", http.MethodGet, p + "/ai/settings", nil, nil,
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.GetProjectAiSettings(ctx, "pr_1", req, nil)
			}, nil},
		{"UpdateProjectAiSettings", http.MethodPut, p + "/ai/settings", nil, map[string]any{"enabled": true, "defaultLlmIntegrationId": "int_llm", "defaultModel": "gpt-4o-mini"},
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.UpdateProjectAiSettings(ctx, "pr_1", req, nil)
			}, map[string]any{"enabled": true, "defaultLlmIntegrationId": "int_llm", "defaultModel": "gpt-4o-mini"}},
		{"CreateProjectAiAssistant", http.MethodPost, p + "/ai/assistants", nil, assistant,
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.CreateProjectAiAssistant(ctx, "pr_1", req, nil)
			}, assistant},
		{"UpdateProjectAiAssistant", http.MethodPut, p + "/ai/assistants/ast_1", nil, assistant,
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.UpdateProjectAiAssistant(ctx, "pr_1", "ast_1", req, nil)
			}, assistant},
		{"DeleteProjectAiAssistant", http.MethodDelete, p + "/ai/assistants/ast_1", nil, nil,
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.DeleteProjectAiAssistant(ctx, "pr_1", "ast_1", req, nil)
			}, nil},
		{"GetProjectAiUsage", http.MethodGet, p + "/ai/usage", url.Values{"top": {"5"}}, nil,
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.GetProjectAiUsage(ctx, "pr_1", req, nil)
			}, map[string]any{"top": 5}},
		// Termination
		{"DisableProject", http.MethodPatch, p + "/disable", nil, nil,
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.DisableProject(ctx, "pr_1", req, nil)
			}, nil},
		{"EnableProject", http.MethodPatch, p + "/enable", nil, nil,
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.EnableProject(ctx, "pr_1", req, nil)
			}, nil},
		{"DeleteProject", http.MethodDelete, "/v2/account/projects/pr_1", nil, nil,
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.DeleteProject(ctx, "pr_1", req, nil)
			}, nil},
		// Billing
		{"GetStripeBillingPortalUrl", http.MethodPost, "/v2/account/stripe/get-portal-url", nil, map[string]any{"subscriptionType": "ManagedService", "returnUrl": "https://cloud.example.com/billing"},
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.GetStripeBillingPortalUrl(ctx, req, nil)
			}, map[string]any{"subscriptionType": "ManagedService", "returnUrl": "https://cloud.example.com/billing"}},
		{"GetManagedServiceBundles", http.MethodGet, "/v2/account/stripe/bundles", nil, nil,
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.GetManagedServiceBundles(ctx, req, nil)
			}, nil},
		{"ChangeManagedServiceBundle", http.MethodPost, "/v2/account/stripe/change-bundle", nil, map[string]any{"bundle": "Growth", "billingPeriod": "Annual"},
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.ChangeManagedServiceBundle(ctx, req, nil)
			}, map[string]any{"bundle": "Growth", "billingPeriod": "Annual"}},
		{"CreateStripeCheckoutSession", http.MethodPost, "/v2/account/stripe/create-checkout-session", nil, map[string]any{"subscriptionType": "ManagedService", "bundle": "Starter", "billingPeriod": "Monthly"},
			func(ctx context.Context, m *AccountModule, req map[string]any) error {
				return m.CreateStripeCheckoutSession(ctx, req, nil)
			}, map[string]any{"subscriptionType": "ManagedService", "bundle": "Starter", "billingPeriod": "Monthly"}},

		// The routes that had no Go method before 2026-11-18: general info
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

		// AI settings (new methods)
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

		// Billing (new methods)
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

// 51 AccountModule routes (54 Hub routes of the section minus the three
// EnvironmentsModule ones); a changed count means a route arrived or left
// untested.
func TestAccountSettingsSurfaceSize(t *testing.T) {
	const want = 51
	if got := len(accountSettingsCases()); got != want {
		t.Errorf("account settings endpoint count: got %d want %d", got, want)
	}
}
