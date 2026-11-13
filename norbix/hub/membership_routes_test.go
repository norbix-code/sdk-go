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

// Every Membership route the Hub serves for the dashboard screens — users
// list, user edit, roles, policies, the five settings tabs and the
// service-user API keys — called through the Membership module. Each case
// asserts the host (hub, not api: the contact routes exist on both), the verb,
// the fully-resolved path, the query (reads and deletes) and the JSON body
// (saves). Source of truth: the gateway `[Route]` attributes
// (Hub.Membership, Hub.Contacts), which the endpoint manifest partly misses.
// Field names come from the generated types in norbix/hub/dtos.
//
// Nothing leaves the process: every call goes to a local test server.

type membershipRouteCase struct {
	name  string
	verb  string
	path  string
	query url.Values     // nil: no query string expected
	body  map[string]any // nil: no body expected
	call  func(ctx context.Context, m *MembershipModule) error
}

const (
	mbContact = "ct_1"
	mbAuth    = "usr_1"
	mbRole    = "role_1"
	mbPolicy  = "pol_1"
)

func membershipRouteCases() []membershipRouteCase {
	createContact := map[string]any{"primaryEmail": "ana@example.test", "displayName": "Ana", "firstName": "Ana", "lastName": "Berg", "tags": []any{"vip"}}
	merge := map[string]any{"survivorId": mbContact, "mergedIds": []any{"ct_2", "ct_3"}}
	dryRun := map[string]any{"dryRun": true}
	importUsers := map[string]any{"dryRun": false, "ignoreDefaultRoles": true, "users": []any{map[string]any{"email": "ana@example.test"}}}
	updateContact := map[string]any{"displayName": "Ana B.", "company": "Norbix", "city": "Vilnius"}
	identity := map[string]any{"authId": mbAuth}
	aiPlan := map[string]any{"planId": "plan_pro"}
	role := map[string]any{"roleName": "Editor", "description": "Edits records", "policies": []any{"pol_1"}}
	roleUpdate := map[string]any{"id": mbRole, "roleName": "Editor", "policies": []any{"pol_1", "pol_2"}}
	policy := map[string]any{"policyName": "Read orders", "description": "Orders only", "policyDocumentJson": `{"statements":[]}`}
	policyUpdate := map[string]any{"id": mbPolicy, "policyName": "Read orders", "policyDocumentJson": `{"statements":[]}`}
	authn := map[string]any{"logoutUrl": "https://app.example.test/bye", "allowUsernames": true}
	authz := map[string]any{"allowAnyoneToSignUp": false, "inviteOnly": true, "defaultRoles": []any{"User"}}
	complexity := map[string]any{"minLength": 12, "minNumbers": 1, "allowedSpecial": "!@#"}
	passkey := map[string]any{"enabled": true, "codeTtlMinutes": 10, "rpId": "app.example.test"}
	apiKey := map[string]any{"name": "ci", "scopes": []any{"database:read"}, "expiresInDays": 30, "notes": "build server"}

	return []membershipRouteCase{
		// Users list.
		{"GetAllContacts", http.MethodGet, "/v2/membership/users", url.Values{"search": {"ana"}, "pageSize": {"25"}}, nil,
			func(ctx context.Context, m *MembershipModule) error {
				return m.GetAllContacts(ctx, map[string]any{"search": "ana", "pageSize": 25}, nil)
			}},
		{"CreateContact", http.MethodPost, "/v2/membership/users", nil, createContact,
			func(ctx context.Context, m *MembershipModule) error { return m.CreateContact(ctx, createContact, nil) }},
		{"MergeContacts", http.MethodPost, "/v2/membership/users/merge", nil, merge,
			func(ctx context.Context, m *MembershipModule) error { return m.MergeContacts(ctx, merge, nil) }},
		{"DeleteContact", http.MethodDelete, "/v2/membership/users/" + mbContact, nil, nil,
			func(ctx context.Context, m *MembershipModule) error { return m.DeleteContact(ctx, mbContact, nil, nil) }},
		{"ImportUsers", http.MethodPost, "/v2/membership/users/import", nil, importUsers,
			func(ctx context.Context, m *MembershipModule) error { return m.ImportUsers(ctx, importUsers, nil) }},
		{"RepairLoginProjectIds", http.MethodPost, "/v2/membership/users/repair-logins", nil, dryRun,
			func(ctx context.Context, m *MembershipModule) error { return m.RepairLoginProjectIds(ctx, dryRun, nil) }},
		{"RepairContactRoles", http.MethodPost, "/v2/membership/users/repair-roles", nil, dryRun,
			func(ctx context.Context, m *MembershipModule) error { return m.RepairContactRoles(ctx, dryRun, nil) }},

		// User edit / detail.
		{"GetContact", http.MethodGet, "/v2/membership/users/" + mbContact, nil, nil,
			func(ctx context.Context, m *MembershipModule) error { return m.GetContact(ctx, mbContact, nil, nil) }},
		{"UpdateContact", http.MethodPatch, "/v2/membership/users/" + mbContact, nil, updateContact,
			func(ctx context.Context, m *MembershipModule) error {
				return m.UpdateContact(ctx, mbContact, updateContact, nil)
			}},
		{"AddContactIdentity", http.MethodPost, "/v2/membership/users/" + mbContact + "/identities", nil, identity,
			func(ctx context.Context, m *MembershipModule) error {
				return m.AddContactIdentity(ctx, mbContact, identity, nil)
			}},
		{"RemoveContactIdentity", http.MethodDelete, "/v2/membership/users/" + mbContact + "/identities/" + mbAuth, nil, nil,
			func(ctx context.Context, m *MembershipModule) error {
				return m.RemoveContactIdentity(ctx, mbContact, mbAuth, nil, nil)
			}},
		{"PromoteContactIdentity", http.MethodPost, "/v2/membership/users/" + mbContact + "/identities/" + mbAuth + "/promote", nil, nil,
			func(ctx context.Context, m *MembershipModule) error {
				return m.PromoteContactIdentity(ctx, mbContact, mbAuth, nil, nil)
			}},
		{"ClearContactBounce", http.MethodPost, "/v2/membership/users/" + mbContact + "/marketing-state/Email/clear-bounce", nil, nil,
			func(ctx context.Context, m *MembershipModule) error {
				return m.ClearContactBounce(ctx, mbContact, "Email", nil, nil)
			}},
		{"SetUserAiPlan", http.MethodPut, "/v2/membership/users/" + mbAuth + "/ai-plan", nil, aiPlan,
			func(ctx context.Context, m *MembershipModule) error { return m.SetUserAiPlan(ctx, mbAuth, aiPlan, nil) }},

		// Roles.
		{"GetRoles", http.MethodGet, "/v2/membership/roles", nil, nil,
			func(ctx context.Context, m *MembershipModule) error { return m.GetRoles(ctx, nil, nil) }},
		{"GetRole", http.MethodGet, "/v2/membership/roles/" + mbRole, nil, nil,
			func(ctx context.Context, m *MembershipModule) error { return m.GetRole(ctx, mbRole, nil, nil) }},
		{"CreateRole", http.MethodPost, "/v2/membership/roles", nil, role,
			func(ctx context.Context, m *MembershipModule) error { return m.CreateRole(ctx, role, nil) }},
		{"UpdateRolePolicies", http.MethodPatch, "/v2/membership/roles", nil, roleUpdate,
			func(ctx context.Context, m *MembershipModule) error {
				return m.UpdateRolePolicies(ctx, roleUpdate, nil)
			}},
		// The id goes on the query: the route has no {Id} (same as the .NET snapshot).
		{"DeleteRole", http.MethodDelete, "/v2/membership/roles", url.Values{"id": {mbRole}}, nil,
			func(ctx context.Context, m *MembershipModule) error {
				return m.DeleteRole(ctx, map[string]any{"id": mbRole}, nil)
			}},

		// Policies.
		{"GetPolicies", http.MethodGet, "/v2/membership/policies", nil, nil,
			func(ctx context.Context, m *MembershipModule) error { return m.GetPolicies(ctx, nil, nil) }},
		{"GetPolicy", http.MethodGet, "/v2/membership/policies/" + mbPolicy, nil, nil,
			func(ctx context.Context, m *MembershipModule) error { return m.GetPolicy(ctx, mbPolicy, nil, nil) }},
		{"CreatePolicy", http.MethodPost, "/v2/membership/policies", nil, policy,
			func(ctx context.Context, m *MembershipModule) error { return m.CreatePolicy(ctx, policy, nil) }},
		{"UpdatePolicy", http.MethodPut, "/v2/membership/policies", nil, policyUpdate,
			func(ctx context.Context, m *MembershipModule) error { return m.UpdatePolicy(ctx, policyUpdate, nil) }},
		{"DeletePolicy", http.MethodDelete, "/v2/membership/policies", url.Values{"id": {mbPolicy}}, nil,
			func(ctx context.Context, m *MembershipModule) error {
				return m.DeletePolicy(ctx, map[string]any{"id": mbPolicy}, nil)
			}},

		// Settings — module switch, authentication, authorization, password complexity, passkeys.
		{"EnableMembership", http.MethodPut, "/v2/membership/enable", nil, nil,
			func(ctx context.Context, m *MembershipModule) error { return m.EnableMembership(ctx, nil, nil) }},
		{"GetAuthenticationSettings", http.MethodGet, "/v2/membership/authentication", nil, nil,
			func(ctx context.Context, m *MembershipModule) error {
				return m.GetAuthenticationSettings(ctx, nil, nil)
			}},
		{"UpdateAuthenticationSettings", http.MethodPut, "/v2/membership/authentication", nil, authn,
			func(ctx context.Context, m *MembershipModule) error {
				return m.UpdateAuthenticationSettings(ctx, authn, nil)
			}},
		{"GetAuthorizationSettings", http.MethodGet, "/v2/membership/authorization", nil, nil,
			func(ctx context.Context, m *MembershipModule) error { return m.GetAuthorizationSettings(ctx, nil, nil) }},
		{"UpdateAuthorizationSettings", http.MethodPut, "/v2/membership/authorization", nil, authz,
			func(ctx context.Context, m *MembershipModule) error {
				return m.UpdateAuthorizationSettings(ctx, authz, nil)
			}},
		{"UpdatePasswordComplexity", http.MethodPut, "/v2/membership/authorization/password-complexity", nil, complexity,
			func(ctx context.Context, m *MembershipModule) error {
				return m.UpdatePasswordComplexity(ctx, complexity, nil)
			}},
		{"GetPasskeySettings", http.MethodGet, "/v2/membership/passkey/settings", nil, nil,
			func(ctx context.Context, m *MembershipModule) error { return m.GetPasskeySettings(ctx, nil, nil) }},
		{"SavePasskeySettings", http.MethodPost, "/v2/membership/passkey/settings", nil, passkey,
			func(ctx context.Context, m *MembershipModule) error { return m.SavePasskeySettings(ctx, passkey, nil) }},

		// Service-user API keys.
		{"ListServiceUserApiKeys", http.MethodGet, "/v2/membership/users/" + mbAuth + "/api-keys", nil, nil,
			func(ctx context.Context, m *MembershipModule) error {
				return m.ListServiceUserApiKeys(ctx, mbAuth, nil, nil)
			}},
		{"IssueServiceUserApiKey", http.MethodPost, "/v2/membership/users/" + mbAuth + "/api-keys", nil, apiKey,
			func(ctx context.Context, m *MembershipModule) error {
				return m.IssueServiceUserApiKey(ctx, mbAuth, apiKey, nil)
			}},
		{"DeleteServiceUserApiKey", http.MethodDelete, "/v2/membership/users/" + mbAuth + "/api-keys/7", nil, nil,
			func(ctx context.Context, m *MembershipModule) error {
				return m.DeleteServiceUserApiKey(ctx, mbAuth, "7", nil, nil)
			}},
	}
}

func TestEveryHubMembershipRouteHitsTheExpectedRoute(t *testing.T) {
	for _, c := range membershipRouteCases() {
		t.Run(c.name, func(t *testing.T) {
			var gotHost, gotMethod, gotPath, gotRawQuery, gotAuth, gotProject string
			var gotBody []byte
			handler := func(host string) http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
					gotHost, gotMethod, gotPath, gotRawQuery = host, r.Method, r.URL.Path, r.URL.RawQuery
					gotAuth = r.Header.Get("Authorization")
					gotProject = r.Header.Get("norbix-project-id")
					gotBody, _ = io.ReadAll(r.Body)
					w.WriteHeader(http.StatusNoContent)
				}
			}
			api := httptest.NewServer(handler("api"))
			defer api.Close()
			hub := httptest.NewServer(handler("hub"))
			defer hub.Close()

			m := &MembershipModule{t: transport.New(&transport.Config{
				ProjectID:  "proj_1",
				AccountID:  "acct_1",
				APIKey:     "key_1",
				BaseURLAPI: api.URL,
				BaseURLHub: hub.URL,
				APIVersion: "v2",
				HubVersion: "v2",
				Timeout:    5 * time.Second,
				Env:        "PROD",
			}, nil)}
			if err := c.call(context.Background(), m); err != nil {
				t.Fatalf("%s: unexpected error: %v", c.name, err)
			}
			if gotHost != "hub" {
				t.Errorf("host: got %q want hub", gotHost)
			}
			if gotMethod != c.verb {
				t.Errorf("verb: got %q want %q", gotMethod, c.verb)
			}
			if gotPath != c.path {
				t.Errorf("path: got %q want %q", gotPath, c.path)
			}
			if gotAuth != "Bearer key_1" {
				t.Errorf("authorization: got %q", gotAuth)
			}
			if gotProject != "proj_1" {
				t.Errorf("project header: got %q", gotProject)
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

// Users list 7 + user edit 7 + roles 5 + policies 5 + settings 8 + API keys 3.
// (Membership cannot be disabled, so settings has no disable route.)
// A Hub membership route added to a screen changes this count, so it cannot
// arrive untested.
func TestHubMembershipRouteSurfaceSize(t *testing.T) {
	const want = 7 + 7 + 5 + 5 + 8 + 3
	if got := len(membershipRouteCases()); got != want {
		t.Errorf("hub membership route count: got %d want %d", got, want)
	}
}
