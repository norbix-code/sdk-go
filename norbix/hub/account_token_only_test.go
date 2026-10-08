package hub

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/norbix-code/sdk-go/v2/norbix/internal/transport"
)

// The gateway takes the account from the signed-in session on every
// /account/... route below (UserSession.UserAuth.AccountId), or from the
// project id in the path, or needs no account at all. So each method works
// with a token only — a client with no AccountID — like the TypeScript,
// .NET, Kotlin and Swift SDKs. One sub-test per method, against a throw-away
// local server: verb, resolved path, auth header, and no norbix-account-id
// header. Never a real gateway.
//
// Four routes are public on the gateway (no [Authenticate]): sign-up
// (CreateAccount), accepting an invitation (CreateTeamMemberFromInvitation),
// the region list (GetAccountRegions, Regions.List) and VerifyAccount. They
// go out with NO token at all; VerifyAccount sends the account id once, in
// the request (the gateway reads VerifyAccount.AccountId, Account/Verify.cs).

type tokenOnlyHub struct {
	Account *AccountModule
	Regions *RegionsModule
}

type tokenOnlyCase struct {
	name string
	verb string
	path string
	call func(ctx context.Context, h tokenOnlyHub) error
}

func newTokenOnlyHub(baseURL string) tokenOnlyHub {
	t := transport.New(&transport.Config{
		ProjectID:  "proj_1",
		AccountID:  "", // no account id: the point of these tests
		APIKey:     "key_1",
		BaseURLAPI: baseURL,
		BaseURLHub: baseURL,
		APIVersion: "v2",
		HubVersion: "v2",
		Timeout:    5 * time.Second,
		Env:        "PROD",
	}, nil)
	return tokenOnlyHub{Account: &AccountModule{t: t}, Regions: &RegionsModule{t: t}}
}

func tokenOnlyCases() []tokenOnlyCase {
	body := map[string]any{"probe": "value"}

	return []tokenOnlyCase{
		{"GetAccountProfile", http.MethodGet, "/v2/account/profile",
			func(ctx context.Context, h tokenOnlyHub) error { return h.Account.GetAccountProfile(ctx, body, nil) }},
		{"UpdateAccountProfile", http.MethodPut, "/v2/account/profile",
			func(ctx context.Context, h tokenOnlyHub) error { return h.Account.UpdateAccountProfile(ctx, body, nil) }},
		{"ResendAccountVerificationToken", http.MethodGet, "/v2/account/verify/resend",
			func(ctx context.Context, h tokenOnlyHub) error {
				return h.Account.ResendAccountVerificationToken(ctx, body, nil)
			}},
		{"GetAccountStatus", http.MethodGet, "/v2/account/status",
			func(ctx context.Context, h tokenOnlyHub) error { return h.Account.GetAccountStatus(ctx, body, nil) }},
		{"CreateStripeCheckoutSession", http.MethodPost, "/v2/account/stripe/create-checkout-session",
			func(ctx context.Context, h tokenOnlyHub) error {
				return h.Account.CreateStripeCheckoutSession(ctx, body, nil)
			}},
		{"GetStripeBillingPortalUrl", http.MethodPost, "/v2/account/stripe/get-portal-url",
			func(ctx context.Context, h tokenOnlyHub) error {
				return h.Account.GetStripeBillingPortalUrl(ctx, body, nil)
			}},
		{"DeleteNotificationsGroup", http.MethodDelete, "/v2/account/projects/projectId_1/notifications/settings/group",
			func(ctx context.Context, h tokenOnlyHub) error {
				return h.Account.DeleteNotificationsGroup(ctx, "projectId_1", body, nil)
			}},
		{"DeleteNotificationsTag", http.MethodDelete, "/v2/account/projects/projectId_1/notifications/settings/tag",
			func(ctx context.Context, h tokenOnlyHub) error {
				return h.Account.DeleteNotificationsTag(ctx, "projectId_1", body, nil)
			}},
		{"RemoveTagFromNotificationsGroup", http.MethodDelete, "/v2/account/projects/projectId_1/notifications/settings/group/tag",
			func(ctx context.Context, h tokenOnlyHub) error {
				return h.Account.RemoveTagFromNotificationsGroup(ctx, "projectId_1", body, nil)
			}},
		{"SaveNotificationsGroup", http.MethodPost, "/v2/account/projects/projectId_1/notifications/settings/group",
			func(ctx context.Context, h tokenOnlyHub) error {
				return h.Account.SaveNotificationsGroup(ctx, "projectId_1", body, nil)
			}},
		{"SaveNotificationsTag", http.MethodPost, "/v2/account/projects/projectId_1/notifications/settings/tag",
			func(ctx context.Context, h tokenOnlyHub) error {
				return h.Account.SaveNotificationsTag(ctx, "projectId_1", body, nil)
			}},
		{"CreateProject", http.MethodPost, "/v2/account/projects",
			func(ctx context.Context, h tokenOnlyHub) error { return h.Account.CreateProject(ctx, body, nil) }},
		{"DeleteProject", http.MethodDelete, "/v2/account/projects/projectId_1",
			func(ctx context.Context, h tokenOnlyHub) error {
				return h.Account.DeleteProject(ctx, "projectId_1", body, nil)
			}},
		{"GetProject", http.MethodGet, "/v2/account/projects/projectId_1",
			func(ctx context.Context, h tokenOnlyHub) error {
				return h.Account.GetProject(ctx, "projectId_1", body, nil)
			}},
		{"GetProjects", http.MethodGet, "/v2/account/projects",
			func(ctx context.Context, h tokenOnlyHub) error { return h.Account.GetProjects(ctx, body, nil) }},
		{"GetProjectTokens", http.MethodGet, "/v2/account/projects/projectId_1/tokens",
			func(ctx context.Context, h tokenOnlyHub) error {
				return h.Account.GetProjectTokens(ctx, "projectId_1", body, nil)
			}},
		{"UpdateProjectAccentColor", http.MethodPatch, "/v2/account/projects/projectId_1/settings/accent-color",
			func(ctx context.Context, h tokenOnlyHub) error {
				return h.Account.UpdateProjectAccentColor(ctx, "projectId_1", body, nil)
			}},
		{"UpdateProjectIcon", http.MethodPatch, "/v2/account/projects/projectId_1/settings/icon",
			func(ctx context.Context, h tokenOnlyHub) error {
				return h.Account.UpdateProjectIcon(ctx, "projectId_1", body, nil)
			}},
		{"UpdateProjectLogo", http.MethodPatch, "/v2/account/projects/projectId_1/settings/logo",
			func(ctx context.Context, h tokenOnlyHub) error {
				return h.Account.UpdateProjectLogo(ctx, "projectId_1", body, nil)
			}},
		{"UpdateProjectMainColor", http.MethodPatch, "/v2/account/projects/projectId_1/settings/main-color",
			func(ctx context.Context, h tokenOnlyHub) error {
				return h.Account.UpdateProjectMainColor(ctx, "projectId_1", body, nil)
			}},
		{"UpdateProjectAllowedOrigins", http.MethodPatch, "/v2/account/projects/projectId_1/settings/origins",
			func(ctx context.Context, h tokenOnlyHub) error {
				return h.Account.UpdateProjectAllowedOrigins(ctx, "projectId_1", body, nil)
			}},
		{"UpdateProjectDefaultLanguage", http.MethodPatch, "/v2/account/projects/projectId_1/settings/default-language",
			func(ctx context.Context, h tokenOnlyHub) error {
				return h.Account.UpdateProjectDefaultLanguage(ctx, "projectId_1", body, nil)
			}},
		{"UpdateProjectDescription", http.MethodPatch, "/v2/account/projects/projectId_1/settings/description",
			func(ctx context.Context, h tokenOnlyHub) error {
				return h.Account.UpdateProjectDescription(ctx, "projectId_1", body, nil)
			}},
		{"DisableProject", http.MethodPatch, "/v2/account/projects/projectId_1/disable",
			func(ctx context.Context, h tokenOnlyHub) error {
				return h.Account.DisableProject(ctx, "projectId_1", body, nil)
			}},
		{"EnableProject", http.MethodPatch, "/v2/account/projects/projectId_1/enable",
			func(ctx context.Context, h tokenOnlyHub) error {
				return h.Account.EnableProject(ctx, "projectId_1", body, nil)
			}},
		{"UpdateProjectLanguages", http.MethodPatch, "/v2/account/projects/projectId_1/settings/languages",
			func(ctx context.Context, h tokenOnlyHub) error {
				return h.Account.UpdateProjectLanguages(ctx, "projectId_1", body, nil)
			}},
		{"CheckProjectLanguages", http.MethodPost, "/v2/account/projects/projectId_1/settings/languages/check",
			func(ctx context.Context, h tokenOnlyHub) error {
				return h.Account.CheckProjectLanguages(ctx, "projectId_1", body, nil)
			}},
		{"UpdateProjectUrl", http.MethodPatch, "/v2/account/projects/projectId_1/settings/url",
			func(ctx context.Context, h tokenOnlyHub) error {
				return h.Account.UpdateProjectUrl(ctx, "projectId_1", body, nil)
			}},
		{"UpdateProjectName", http.MethodPatch, "/v2/account/projects/projectId_1/settings/name",
			func(ctx context.Context, h tokenOnlyHub) error {
				return h.Account.UpdateProjectName(ctx, "projectId_1", body, nil)
			}},
		{"UpdateProjectRegions", http.MethodPatch, "/v2/account/projects/projectId_1/settings/regions",
			func(ctx context.Context, h tokenOnlyHub) error {
				return h.Account.UpdateProjectRegions(ctx, "projectId_1", body, nil)
			}},
		{"SendInviteToTeamMember", http.MethodPost, "/v2/account/team/member/invite",
			func(ctx context.Context, h tokenOnlyHub) error {
				return h.Account.SendInviteToTeamMember(ctx, body, nil)
			}},
		{"GetLicenses", http.MethodGet, "/v2/account/licenses",
			func(ctx context.Context, h tokenOnlyHub) error { return h.Account.GetLicenses(ctx, body, nil) }},
		{"UpdateProjectAdminUrl", http.MethodPatch, "/v2/account/projects/projectId_1/settings/admin-url",
			func(ctx context.Context, h tokenOnlyHub) error {
				return h.Account.UpdateProjectAdminUrl(ctx, "projectId_1", body, nil)
			}},
		{"UpdateProjectLegalDocuments", http.MethodPatch, "/v2/account/projects/projectId_1/settings/legal",
			func(ctx context.Context, h tokenOnlyHub) error {
				return h.Account.UpdateProjectLegalDocuments(ctx, "projectId_1", body, nil)
			}},
		{"UpdateProjectExposeLegal", http.MethodPatch, "/v2/account/projects/projectId_1/settings/legal/expose",
			func(ctx context.Context, h tokenOnlyHub) error {
				return h.Account.UpdateProjectExposeLegal(ctx, "projectId_1", body, nil)
			}},
		{"UpdateProjectExposeBrand", http.MethodPatch, "/v2/account/projects/projectId_1/settings/brand/expose",
			func(ctx context.Context, h tokenOnlyHub) error {
				return h.Account.UpdateProjectExposeBrand(ctx, "projectId_1", body, nil)
			}},
		{"UpdateProjectExposeAuth", http.MethodPatch, "/v2/account/projects/projectId_1/settings/auth/expose",
			func(ctx context.Context, h tokenOnlyHub) error {
				return h.Account.UpdateProjectExposeAuth(ctx, "projectId_1", body, nil)
			}},
		{"AssignAdminPortalServiceUser", http.MethodPut, "/v2/account/projects/projectId_1/settings/admin-portal/service-user",
			func(ctx context.Context, h tokenOnlyHub) error {
				return h.Account.AssignAdminPortalServiceUser(ctx, "projectId_1", body, nil)
			}},
		{"CreateAiServiceUser", http.MethodPost, "/v2/account/ai/service-users",
			func(ctx context.Context, h tokenOnlyHub) error { return h.Account.CreateAiServiceUser(ctx, body, nil) }},
		{"ListAiServiceUsers", http.MethodGet, "/v2/account/ai/service-users",
			func(ctx context.Context, h tokenOnlyHub) error { return h.Account.ListAiServiceUsers(ctx, body, nil) }},
		{"DeleteAiServiceUser", http.MethodDelete, "/v2/account/ai/service-users/id_1",
			func(ctx context.Context, h tokenOnlyHub) error {
				return h.Account.DeleteAiServiceUser(ctx, "id_1", body, nil)
			}},
		{"RotateAiServiceUserKey", http.MethodPost, "/v2/account/ai/service-users/id_1/keys",
			func(ctx context.Context, h tokenOnlyHub) error {
				return h.Account.RotateAiServiceUserKey(ctx, "id_1", body, nil)
			}},
		{"RevokeAiServiceUserKey", http.MethodDelete, "/v2/account/ai/service-users/id_1/keys/keyId_1",
			func(ctx context.Context, h tokenOnlyHub) error {
				return h.Account.RevokeAiServiceUserKey(ctx, "id_1", "keyId_1", body, nil)
			}},
		{"Regions.UpdateProjectRegions", http.MethodPatch, "/v2/account/projects/projectId_1/settings/regions",
			func(ctx context.Context, h tokenOnlyHub) error {
				return h.Regions.UpdateProjectRegions(ctx, "projectId_1", body, nil)
			}},
	}
}

func TestAccountRoutesWorkWithATokenAndNoAccountID(t *testing.T) {
	for _, c := range tokenOnlyCases() {
		t.Run(c.name, func(t *testing.T) {
			var gotMethod, gotPath, gotAuth, gotAccount string
			hit := false
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				hit = true
				gotMethod = r.Method
				gotPath = r.URL.Path
				gotAuth = r.Header.Get("Authorization")
				gotAccount = r.Header.Get("norbix-account-id")
				w.WriteHeader(http.StatusNoContent)
			}))
			defer srv.Close()

			if err := c.call(context.Background(), newTokenOnlyHub(srv.URL)); err != nil {
				t.Fatalf("%s: want the call to go out with a token only, got %v", c.name, err)
			}
			if !hit {
				t.Fatalf("%s: the request never reached the server", c.name)
			}
			if gotMethod != c.verb || gotPath != c.path {
				t.Errorf("%s route: got %s %s want %s %s", c.name, gotMethod, gotPath, c.verb, c.path)
			}
			if gotAuth != "Bearer key_1" {
				t.Errorf("%s auth header: got %q want %q", c.name, gotAuth, "Bearer key_1")
			}
			if gotAccount != "" {
				t.Errorf("%s: no norbix-account-id header expected, got %q", c.name, gotAccount)
			}
		})
	}
}

func newNoTokenHub(baseURL string) tokenOnlyHub {
	t := transport.New(&transport.Config{
		ProjectID:  "proj_1",
		BaseURLAPI: baseURL,
		BaseURLHub: baseURL,
		APIVersion: "v2",
		HubVersion: "v2",
		Timeout:    5 * time.Second,
		Env:        "PROD",
	}, nil) // no APIKey, no BearerToken, no AccountID
	return tokenOnlyHub{Account: &AccountModule{t: t}, Regions: &RegionsModule{t: t}}
}

func TestPublicAccountRoutesWorkWithNoTokenAndNoAccountID(t *testing.T) {
	signUp := map[string]any{"email": "ada@example.test", "password": "pw", "displayName": "Ada"}
	invite := map[string]any{"token": "invite-1", "password": "pw", "displayName": "Ada"}
	verify := map[string]any{"accountId": "acct_1", "token": "verify-1"}

	cases := []struct {
		name  string
		verb  string
		path  string
		query string // GET: the query the gateway reads
		body  map[string]any
		call  func(ctx context.Context, h tokenOnlyHub) error
	}{
		{"CreateAccount", http.MethodPost, "/v2/account", "", signUp,
			func(ctx context.Context, h tokenOnlyHub) error { return h.Account.CreateAccount(ctx, signUp, nil) }},
		{"CreateTeamMemberFromInvitation", http.MethodPost, "/v2/account/team/member", "", invite,
			func(ctx context.Context, h tokenOnlyHub) error {
				return h.Account.CreateTeamMemberFromInvitation(ctx, invite, nil)
			}},
		{"GetAccountRegions", http.MethodGet, "/v2/account/regions", "", nil,
			func(ctx context.Context, h tokenOnlyHub) error { return h.Account.GetAccountRegions(ctx, nil, nil) }},
		{"Regions.List", http.MethodGet, "/v2/account/regions", "", nil,
			func(ctx context.Context, h tokenOnlyHub) error { return h.Regions.List(ctx, nil, nil) }},
		{"VerifyAccount", http.MethodGet, "/v2/account/verify", "accountId=acct_1&token=verify-1", nil,
			func(ctx context.Context, h tokenOnlyHub) error { return h.Account.VerifyAccount(ctx, verify, nil) }},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var gotMethod, gotPath, gotQuery, gotAuth, gotAccount string
			var gotBody map[string]any
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotMethod, gotPath, gotQuery = r.Method, r.URL.Path, r.URL.Query().Encode()
				gotAuth = r.Header.Get("Authorization")
				gotAccount = r.Header.Get("norbix-account-id")
				raw, _ := io.ReadAll(r.Body)
				if len(raw) > 0 {
					_ = json.Unmarshal(raw, &gotBody)
				}
				w.WriteHeader(http.StatusNoContent)
			}))
			defer srv.Close()

			if err := c.call(context.Background(), newNoTokenHub(srv.URL)); err != nil {
				t.Fatalf("%s: want the call to go out with no token, got %v", c.name, err)
			}
			if gotMethod != c.verb || gotPath != c.path {
				t.Errorf("%s route: got %s %s want %s %s", c.name, gotMethod, gotPath, c.verb, c.path)
			}
			if gotAuth != "" || gotAccount != "" {
				t.Errorf("%s: no Authorization / norbix-account-id header expected, got %q / %q", c.name, gotAuth, gotAccount)
			}
			if gotQuery != c.query {
				t.Errorf("%s query: got %q want %q", c.name, gotQuery, c.query)
			}
			if c.body != nil && !reflect.DeepEqual(gotBody, c.body) {
				t.Errorf("%s body: got %v want %v", c.name, gotBody, c.body)
			}
		})
	}
}
