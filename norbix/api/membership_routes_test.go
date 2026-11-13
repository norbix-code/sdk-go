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

// Every Membership route the API host serves — the user list and the user
// edit screens (contacts and the /membership/auth user routes), user
// registration and invitation, and the end-user passkey / password / token
// routes (/membership/userauth/*) — called through the API Membership
// module. Each case asserts the host (api, not hub: the contact routes exist
// on both), the verb, the fully-resolved path, the query (reads and deletes)
// and the JSON body (saves). Source of truth: the gateway `[Route]`
// attributes (Api.Membership, Hub.Contacts served again on the API host),
// which the endpoint manifest partly misses. Field names come from the
// generated types in norbix/api/dtos.
//
// Nothing leaves the process: every call goes to a local test server.

type apiMembershipCase struct {
	name  string
	verb  string
	path  string
	query url.Values     // nil: no query string expected
	body  map[string]any // nil: no body expected
	call  func(ctx context.Context, m *MembershipModule) error
}

const (
	amContact = "ct_1"
	amUser    = "usr_1"
	amAuth    = "auth_1"
	amCred    = "cred_1"
	amDb      = "int_db"
)

func apiMembershipCases() []apiMembershipCase {
	createContact := map[string]any{"primaryEmail": "ana@example.test", "displayName": "Ana", "tags": []any{"vip"}}
	merge := map[string]any{"survivorId": amContact, "mergedIds": []any{"ct_2"}}
	dryRun := map[string]any{"dryRun": true}
	updateContact := map[string]any{"displayName": "Ana B.", "language": "lt"}
	identity := map[string]any{"authId": amAuth}
	aiPlan := map[string]any{"planId": "plan_pro"}

	invite := map[string]any{"email": "ana@example.test", "fullName": "Ana Berg", "roles": []any{"Editor"}}
	finish := map[string]any{"token": "inv_tok", "password": "S3cret-pass!", "firstName": "Ana"}
	email := map[string]any{"email": "ana@example.test", "password": "S3cret-pass!", "roles": []any{"User"}}
	emailPerm := map[string]any{"email": "ana@example.test", "password": "S3cret-pass!"}
	guest := map[string]any{"userGeneralInfo": map[string]any{"displayName": "Guest"}}
	phone := map[string]any{"phone": "+37060000000"}
	service := map[string]any{"userGeneralInfo": map[string]any{"displayName": "ci-bot"}, "roles": []any{"Service"}}
	userName := map[string]any{"userName": "ana", "password": "S3cret-pass!", "roles": []any{"User"}}
	userNamePerm := map[string]any{"userName": "ana", "password": "S3cret-pass!"}

	updateUser := map[string]any{"id": amUser, "userGeneralInfo": map[string]any{"displayName": "Ana B."}}
	assignRoles := map[string]any{"id": amUser, "roles": []any{"Editor"}}
	byID := map[string]any{"id": amUser}
	prefs := map[string]any{"blockAllMarketingMessages": false, "blockedTags": []any{"promo"}}
	// PKCE: the RFC 7636 appendix B verifier and its S256 challenge.
	link := map[string]any{"code": "link_code", "codeVerifier": "dBjftJeZ4CVP-mB92K27uhbUJU1p1r_wW1gFWFOEjXk"}
	linkStart := map[string]any{"provider": "google", "mode": "app", "returnUrl": "https://app.example.test/settings",
		"codeChallenge": "E9Melhoa2OwvFrEMTJguCHaoeK1t8URWbuGJSstw-cM"}
	exchange := map[string]any{"code": "one_time", "codeVerifier": "dBjftJeZ4CVP-mB92K27uhbUJU1p1r_wW1gFWFOEjXk"}
	tag := map[string]any{"subscribed": true}
	mapAuth := map[string]any{"authId": amAuth, "databaseIntegrationId": amDb}
	setRoles := map[string]any{"roles": []any{"Editor", "User"}}
	consent := map[string]any{"lawfulBasis": "Consent", "source": "signup-form"}
	unsubscribe := map[string]any{"reason": "too many emails"}

	mail := map[string]any{"email": "ana@example.test"}
	confirmMail := map[string]any{"email": "ana@example.test", "code": "123456"}
	refresh := map[string]any{"refreshToken": "rt_1"}
	regOptions := map[string]any{"verificationToken": "vt_1"}
	verifyAuthn := map[string]any{"ceremonyId": "cer_1", "assertionResponse": map[string]any{"id": "cred"}}
	verifyReg := map[string]any{"verificationToken": "vt_1", "ceremonyId": "cer_1", "attestationResponse": map[string]any{"id": "cred"}, "friendlyName": "Laptop"}
	rename := map[string]any{"friendlyName": "Phone"}
	changePassword := map[string]any{"currentPassword": "old-pass", "newPassword": "new-pass-1!"}
	confirmReset := map[string]any{"token": "rst_1", "newPassword": "new-pass-1!"}
	consume := map[string]any{"token": "ml_1"}
	recovery := map[string]any{"email": "ana@example.test", "recoveryCode": "AAAA-BBBB"}

	return []apiMembershipCase{
		// Users list.
		{"GetAllContacts", http.MethodGet, "/v2/membership/users", url.Values{"search": {"ana"}, "pageSize": {"25"}}, nil,
			func(ctx context.Context, m *MembershipModule) error {
				return m.GetAllContacts(ctx, map[string]any{"search": "ana", "pageSize": 25}, nil)
			}},
		{"CreateContact", http.MethodPost, "/v2/membership/users", nil, createContact,
			func(ctx context.Context, m *MembershipModule) error { return m.CreateContact(ctx, createContact, nil) }},
		{"MergeContacts", http.MethodPost, "/v2/membership/users/merge", nil, merge,
			func(ctx context.Context, m *MembershipModule) error { return m.MergeContacts(ctx, merge, nil) }},
		{"RepairContactRoles", http.MethodPost, "/v2/membership/users/repair-roles", nil, dryRun,
			func(ctx context.Context, m *MembershipModule) error { return m.RepairContactRoles(ctx, dryRun, nil) }},
		{"DeleteContact", http.MethodDelete, "/v2/membership/users/" + amContact, nil, nil,
			func(ctx context.Context, m *MembershipModule) error { return m.DeleteContact(ctx, amContact, nil, nil) }},
		{"GetUsers", http.MethodGet, "/v2/membership/auth", url.Values{"search": {"ana"}, "roleNames": {"Editor"}}, nil,
			func(ctx context.Context, m *MembershipModule) error {
				return m.GetUsers(ctx, map[string]any{"search": "ana", "roleNames": []string{"Editor"}}, nil)
			}},
		// The id goes on the query: the route has no {id} (same as the .NET snapshot).
		{"DeleteUser", http.MethodDelete, "/v2/membership/auth", url.Values{"id": {amUser}}, nil,
			func(ctx context.Context, m *MembershipModule) error { return m.DeleteUser(ctx, byID, nil) }},

		// User — new: invitation and the eight registration routes.
		{"InviteUser", http.MethodPost, "/v2/membership/auth/invite", nil, invite,
			func(ctx context.Context, m *MembershipModule) error { return m.InviteUser(ctx, invite, nil) }},
		{"FinishInvitation", http.MethodPost, "/v2/membership/auth/invitation/finish", nil, finish,
			func(ctx context.Context, m *MembershipModule) error { return m.FinishInvitation(ctx, finish, nil) }},
		{"SaveEmailUser", http.MethodPost, "/v2/membership/auth/register/email", nil, email,
			func(ctx context.Context, m *MembershipModule) error { return m.SaveEmailUser(ctx, email, nil) }},
		{"SaveEmailUserNameWithPermissions", http.MethodPost, "/v2/membership/auth/register/email-with-permissions", nil, emailPerm,
			func(ctx context.Context, m *MembershipModule) error {
				return m.SaveEmailUserNameWithPermissions(ctx, emailPerm, nil)
			}},
		{"SaveGuestUser", http.MethodPost, "/v2/membership/auth/register/guest", nil, guest,
			func(ctx context.Context, m *MembershipModule) error { return m.SaveGuestUser(ctx, guest, nil) }},
		{"SavePhoneUser", http.MethodPost, "/v2/membership/auth/register/phone", nil, phone,
			func(ctx context.Context, m *MembershipModule) error { return m.SavePhoneUser(ctx, phone, nil) }},
		{"SavePhoneUserNameWithPermissions", http.MethodPost, "/v2/membership/auth/register/phone-with-permissions", nil, phone,
			func(ctx context.Context, m *MembershipModule) error {
				return m.SavePhoneUserNameWithPermissions(ctx, phone, nil)
			}},
		{"SaveSystemUserWithPermissions", http.MethodPost, "/v2/membership/auth/register/service", nil, service,
			func(ctx context.Context, m *MembershipModule) error {
				return m.SaveSystemUserWithPermissions(ctx, service, nil)
			}},
		{"SaveUserNameUser", http.MethodPost, "/v2/membership/auth/register/user-name", nil, userName,
			func(ctx context.Context, m *MembershipModule) error { return m.SaveUserNameUser(ctx, userName, nil) }},
		{"SaveUserNameWithPermissions", http.MethodPost, "/v2/membership/auth/register/user-name-with-permissions", nil, userNamePerm,
			func(ctx context.Context, m *MembershipModule) error {
				return m.SaveUserNameWithPermissions(ctx, userNamePerm, nil)
			}},

		// User edit / detail — contact routes on the API host.
		{"GetContact", http.MethodGet, "/v2/membership/users/" + amContact, nil, nil,
			func(ctx context.Context, m *MembershipModule) error { return m.GetContact(ctx, amContact, nil, nil) }},
		{"UpdateContact", http.MethodPatch, "/v2/membership/users/" + amContact, nil, updateContact,
			func(ctx context.Context, m *MembershipModule) error {
				return m.UpdateContact(ctx, amContact, updateContact, nil)
			}},
		{"AddContactIdentity", http.MethodPost, "/v2/membership/users/" + amContact + "/identities", nil, identity,
			func(ctx context.Context, m *MembershipModule) error {
				return m.AddContactIdentity(ctx, amContact, identity, nil)
			}},
		{"RemoveContactIdentity", http.MethodDelete, "/v2/membership/users/" + amContact + "/identities/" + amAuth, nil, nil,
			func(ctx context.Context, m *MembershipModule) error {
				return m.RemoveContactIdentity(ctx, amContact, amAuth, nil, nil)
			}},
		{"PromoteContactIdentity", http.MethodPost, "/v2/membership/users/" + amContact + "/identities/" + amAuth + "/promote", nil, nil,
			func(ctx context.Context, m *MembershipModule) error {
				return m.PromoteContactIdentity(ctx, amContact, amAuth, nil, nil)
			}},
		{"ClearContactBounce", http.MethodPost, "/v2/membership/users/" + amContact + "/marketing-state/Sms/clear-bounce", nil, nil,
			func(ctx context.Context, m *MembershipModule) error {
				return m.ClearContactBounce(ctx, amContact, "Sms", nil, nil)
			}},
		{"SetUserAiPlan", http.MethodPut, "/v2/membership/users/" + amUser + "/ai-plan", nil, aiPlan,
			func(ctx context.Context, m *MembershipModule) error { return m.SetUserAiPlan(ctx, amUser, aiPlan, nil) }},

		// User edit / detail — /membership/auth user routes.
		{"UpdateUser", http.MethodPut, "/v2/membership/auth", nil, updateUser,
			func(ctx context.Context, m *MembershipModule) error { return m.UpdateUser(ctx, updateUser, nil) }},
		{"AssignRolePermissions", http.MethodPut, "/v2/membership/auth/assign-roles", nil, assignRoles,
			func(ctx context.Context, m *MembershipModule) error {
				return m.AssignRolePermissions(ctx, assignRoles, nil)
			}},
		{"BlockUser", http.MethodPatch, "/v2/membership/auth/block", nil, byID,
			func(ctx context.Context, m *MembershipModule) error { return m.BlockUser(ctx, byID, nil) }},
		{"UnblockUser", http.MethodPatch, "/v2/membership/auth/unblock", nil, byID,
			func(ctx context.Context, m *MembershipModule) error { return m.UnblockUser(ctx, byID, nil) }},
		{"GetUser", http.MethodGet, "/v2/membership/auth/" + amUser, url.Values{"databaseIntegrationId": {amDb}}, nil,
			func(ctx context.Context, m *MembershipModule) error {
				return m.GetUser(ctx, amUser, map[string]any{"databaseIntegrationId": amDb}, nil)
			}},
		{"GetUserPreferences", http.MethodGet, "/v2/membership/auth/" + amUser + "/preferences", nil, nil,
			func(ctx context.Context, m *MembershipModule) error {
				return m.GetUserPreferences(ctx, amUser, nil, nil)
			}},
		{"UpdateUserPreferences", http.MethodPut, "/v2/membership/auth/" + amUser + "/preferences", nil, prefs,
			func(ctx context.Context, m *MembershipModule) error {
				return m.UpdateUserPreferences(ctx, amUser, prefs, nil)
			}},
		{"StartLinkIdentity", http.MethodPost, "/v2/membership/auth/" + amUser + "/link-identity/start", nil, linkStart,
			func(ctx context.Context, m *MembershipModule) error {
				return m.StartLinkIdentity(ctx, amUser, linkStart, nil)
			}},
		{"LinkIdentity", http.MethodPost, "/v2/membership/auth/" + amUser + "/link-identity", nil, link,
			func(ctx context.Context, m *MembershipModule) error { return m.LinkIdentity(ctx, amUser, link, nil) }},
		{"ExchangeSocialSignInCode", http.MethodPost, "/v2/membership/auth/social/google/exchange", nil, exchange,
			func(ctx context.Context, m *MembershipModule) error {
				return m.ExchangeSocialSignInCode(ctx, "google", exchange, nil)
			}},
		{"SetContactTagSubscription", http.MethodPut, "/v2/membership/users/" + amContact + "/marketing-state/Email/Marketing/tags/news", nil, tag,
			func(ctx context.Context, m *MembershipModule) error {
				return m.SetContactTagSubscription(ctx, amContact, "Email", "Marketing", "news", tag, nil)
			}},
		{"MapAuthToUser", http.MethodPost, "/v2/membership/users/" + amUser + "/map-auth", nil, mapAuth,
			func(ctx context.Context, m *MembershipModule) error {
				return m.MapAuthToUser(ctx, amUser, mapAuth, nil)
			}},
		{"SetContactRoles", http.MethodPut, "/v2/membership/users/" + amUser + "/roles", nil, setRoles,
			func(ctx context.Context, m *MembershipModule) error {
				return m.SetContactRoles(ctx, amUser, setRoles, nil)
			}},
		{"GrantContactConsent", http.MethodPost, "/v2/membership/users/" + amContact + "/marketing-state/Email/consent", nil, consent,
			func(ctx context.Context, m *MembershipModule) error {
				return m.GrantContactConsent(ctx, amContact, "Email", consent, nil)
			}},
		{"UnsubscribeContact", http.MethodPost, "/v2/membership/users/" + amContact + "/marketing-state/Email/unsubscribe", nil, unsubscribe,
			func(ctx context.Context, m *MembershipModule) error {
				return m.UnsubscribeContact(ctx, amContact, "Email", unsubscribe, nil)
			}},

		// Passkey & user-auth (API only, no screen).
		{"ConfirmEmailVerification", http.MethodPost, "/v2/membership/userauth/email/confirm-verification", nil, confirmMail,
			func(ctx context.Context, m *MembershipModule) error {
				return m.ConfirmEmailVerification(ctx, confirmMail, nil)
			}},
		{"StartEmailVerification", http.MethodPost, "/v2/membership/userauth/email/start-verification", nil, mail,
			func(ctx context.Context, m *MembershipModule) error { return m.StartEmailVerification(ctx, mail, nil) }},
		{"HasPasskey", http.MethodPost, "/v2/membership/userauth/has-passkey", nil, mail,
			func(ctx context.Context, m *MembershipModule) error { return m.HasPasskey(ctx, mail, nil) }},
		{"PasskeyLogout", http.MethodPost, "/v2/membership/userauth/logout", nil, refresh,
			func(ctx context.Context, m *MembershipModule) error { return m.PasskeyLogout(ctx, refresh, nil) }},
		{"PasskeyAuthenticationOptions", http.MethodPost, "/v2/membership/userauth/passkey/authentication-options", nil, mail,
			func(ctx context.Context, m *MembershipModule) error {
				return m.PasskeyAuthenticationOptions(ctx, mail, nil)
			}},
		{"PasskeyRegistrationOptions", http.MethodPost, "/v2/membership/userauth/passkey/registration-options", nil, regOptions,
			func(ctx context.Context, m *MembershipModule) error {
				return m.PasskeyRegistrationOptions(ctx, regOptions, nil)
			}},
		{"VerifyPasskeyAuthentication", http.MethodPost, "/v2/membership/userauth/passkey/verify-authentication", nil, verifyAuthn,
			func(ctx context.Context, m *MembershipModule) error {
				return m.VerifyPasskeyAuthentication(ctx, verifyAuthn, nil)
			}},
		{"VerifyPasskeyRegistration", http.MethodPost, "/v2/membership/userauth/passkey/verify-registration", nil, verifyReg,
			func(ctx context.Context, m *MembershipModule) error {
				return m.VerifyPasskeyRegistration(ctx, verifyReg, nil)
			}},
		{"ListPasskeys", http.MethodGet, "/v2/membership/userauth/passkeys", nil, nil,
			func(ctx context.Context, m *MembershipModule) error { return m.ListPasskeys(ctx, nil, nil) }},
		{"RenamePasskey", http.MethodPost, "/v2/membership/userauth/passkeys/" + amCred + "/rename", nil, rename,
			func(ctx context.Context, m *MembershipModule) error { return m.RenamePasskey(ctx, amCred, rename, nil) }},
		{"RevokePasskey", http.MethodPost, "/v2/membership/userauth/passkeys/" + amCred + "/revoke", nil, nil,
			func(ctx context.Context, m *MembershipModule) error { return m.RevokePasskey(ctx, amCred, nil, nil) }},
		{"ChangePassword", http.MethodPost, "/v2/membership/userauth/password/change", nil, changePassword,
			func(ctx context.Context, m *MembershipModule) error {
				return m.ChangePassword(ctx, changePassword, nil)
			}},
		{"ConfirmPasswordReset", http.MethodPost, "/v2/membership/userauth/password/reset/confirm", nil, confirmReset,
			func(ctx context.Context, m *MembershipModule) error {
				return m.ConfirmPasswordReset(ctx, confirmReset, nil)
			}},
		{"RequestPasswordReset", http.MethodPost, "/v2/membership/userauth/password/reset/request", nil, mail,
			func(ctx context.Context, m *MembershipModule) error { return m.RequestPasswordReset(ctx, mail, nil) }},
		{"SendTestPasswordReset", http.MethodPost, "/v2/membership/userauth/password/reset/test", nil, mail,
			func(ctx context.Context, m *MembershipModule) error { return m.SendTestPasswordReset(ctx, mail, nil) }},
		{"ConsumeMagicLink", http.MethodPost, "/v2/membership/userauth/recovery/magic-link/consume", nil, consume,
			func(ctx context.Context, m *MembershipModule) error { return m.ConsumeMagicLink(ctx, consume, nil) }},
		{"RequestMagicLink", http.MethodPost, "/v2/membership/userauth/recovery/magic-link/request", nil, mail,
			func(ctx context.Context, m *MembershipModule) error { return m.RequestMagicLink(ctx, mail, nil) }},
		{"UseRecoveryCode", http.MethodPost, "/v2/membership/userauth/recovery/use-code", nil, recovery,
			func(ctx context.Context, m *MembershipModule) error { return m.UseRecoveryCode(ctx, recovery, nil) }},
		{"RefreshPasskeyToken", http.MethodPost, "/v2/membership/userauth/token/refresh", nil, refresh,
			func(ctx context.Context, m *MembershipModule) error { return m.RefreshPasskeyToken(ctx, refresh, nil) }},
	}
}

// Calls that must go out with no credentials at all, though the client has a key.
var noCredentialRoutes = map[string]struct{}{"ExchangeSocialSignInCode": {}}

func TestEveryAPIMembershipRouteHitsTheExpectedRoute(t *testing.T) {
	for _, c := range apiMembershipCases() {
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
			if gotHost != "api" {
				t.Errorf("host: got %q want api", gotHost)
			}
			if gotMethod != c.verb {
				t.Errorf("verb: got %q want %q", gotMethod, c.verb)
			}
			if gotPath != c.path {
				t.Errorf("path: got %q want %q", gotPath, c.path)
			}
			wantAuth := "Bearer key_1"
			if _, open := noCredentialRoutes[c.name]; open {
				// The gateway judges any token it gets, even on an open route, and
				// refuses an old one, so these calls send none.
				wantAuth = ""
			}
			if gotAuth != wantAuth {
				t.Errorf("authorization: got %q want %q", gotAuth, wantAuth)
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

// Users list 7 + user new 10 + user edit 20 + passkey & user-auth 19 + social
// sign-in and link 2 (exchange, link start; link finish is in user edit).
// An API membership route added to a screen changes this count, so it cannot
// arrive untested.
func TestAPIMembershipRouteSurfaceSize(t *testing.T) {
	const want = 7 + 10 + 20 + 19 + 2
	if got := len(apiMembershipCases()); got != want {
		t.Errorf("api membership route count: got %d want %d", got, want)
	}
}
