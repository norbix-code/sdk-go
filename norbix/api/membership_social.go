package api

import "github.com/norbix-code/sdk-go/v2/norbix/internal/transport"

// SocialSignInURL is the URL that starts a project social sign-in:
// GET {api}/{version}/membership/auth/social/{provider}?projectId=…[&env=…][&mode=…][&returnUrl=…].
//
// It is not a call: send the browser there (on a phone, the system browser)
// and the gateway redirects to the provider. A navigation carries no headers,
// so the project id and a non-PROD env travel in the query. provider is
// google, apple, microsoft, github, facebook, x or okta. mode picks the
// integration's OAuth mode ("" = the only mode, else the first by name).
// returnURL is the app page to come back to: absolute https (http only on
// localhost), origin in the project's allowed origins; "" uses the mode's
// pages. The page gets ?code=…&provider=… (or ?error=…); trade the code with
// ExchangeSocialSignInCode within 2 minutes.
//
// Hand-written: the start route answers a 302, so it has no call method.
func (m *MembershipModule) SocialSignInURL(provider, mode, returnURL string) (string, error) {
	return m.t.NavigationURL(
		transport.TargetAPI,
		"/{version}/membership/auth/social/{provider}",
		map[string]string{"provider": provider},
		[][2]string{{"mode", mode}, {"returnUrl", returnURL}},
	)
}
