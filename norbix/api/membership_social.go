package api

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"regexp"

	"github.com/norbix-code/sdk-go/v2/norbix/internal/transport"
)

// SocialSignInURL is the URL that starts a project social sign-in:
// GET {api}/{version}/membership/auth/social/{provider}?projectId=…[&env=…][&mode=…][&returnUrl=…].
//
// It is not a call: send the browser there (on a phone, the system browser)
// and the gateway redirects to the provider. A navigation carries no headers,
// so the project id and a non-PROD env travel in the query. provider is
// google, apple, microsoft, github, facebook, x or okta. mode picks the
// integration's OAuth mode ("" = the only mode, else the first by name).
// returnURL is the app page to come back to: absolute https (http only on
// localhost), origin in the project's allowed origins or the project's own
// hosted Admin Portal; "" uses the mode's pages. The page gets
// ?code=…&provider=… (or ?error=…); trade the code with
// ExchangeSocialSignInCode within 2 minutes. For PKCE use
// SocialSignInURLWithOptions.
//
// Hand-written: the start route answers a 302, so it has no call method.
func (m *MembershipModule) SocialSignInURL(provider, mode, returnURL string) (string, error) {
	return m.SocialSignInURLWithOptions(provider, SocialSignInOptions{Mode: mode, ReturnURL: returnURL})
}

// SocialSignInOptions are the optional parts of the social sign-in start URL.
type SocialSignInOptions struct {
	// Mode is the integration's OAuth mode ("" = the only mode, else the first by name).
	Mode string
	// ReturnURL is the app page to come back to ("" = the mode's pages).
	ReturnURL string
	// CodeChallenge is the PKCE S256 challenge from NewPKCE; the exchange then
	// needs its CodeVerifier as "codeVerifier".
	CodeChallenge string
}

// SocialSignInURLWithOptions is SocialSignInURL with PKCE:
// …[&mode=…][&returnUrl=…][&codeChallenge=…].
func (m *MembershipModule) SocialSignInURLWithOptions(provider string, opts SocialSignInOptions) (string, error) {
	return m.t.NavigationURL(
		transport.TargetAPI,
		"/{version}/membership/auth/social/{provider}",
		map[string]string{"provider": provider},
		[][2]string{{"mode", opts.Mode}, {"returnUrl", opts.ReturnURL}, {"codeChallenge", opts.CodeChallenge}},
	)
}

// PKCEPair is a PKCE (RFC 7636, S256) verifier and its challenge for social
// sign-in and link identity. Keep CodeVerifier (memory on a device, the
// session on a server) and never put it in a URL; send CodeChallenge on the
// start ("codeChallenge") and CodeVerifier on the exchange or the link
// finish ("codeVerifier"). A code started with a challenge works only with
// the matching verifier: CM-ERRORS-MEMBERSHIP-058 otherwise, and the code is
// used up. A challenge that is not S256 is CM-ERRORS-MEMBERSHIP-059.
type PKCEPair struct {
	CodeVerifier  string
	CodeChallenge string
}

// NewPKCE makes a verifier from 32 crypto/rand bytes (base64url, 43
// characters) and its S256 challenge.
func NewPKCE() (PKCEPair, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return PKCEPair{}, err
	}
	verifier := base64.RawURLEncoding.EncodeToString(b)
	challenge, err := PKCEChallenge(verifier)
	return PKCEPair{CodeVerifier: verifier, CodeChallenge: challenge}, err
}

var pkceVerifier = regexp.MustCompile(`^[A-Za-z0-9\-._~]{43,128}$`)

// PKCEChallenge is the S256 challenge of a verifier: base64url (no padding)
// of SHA-256 over its ASCII bytes.
func PKCEChallenge(verifier string) (string, error) {
	if !pkceVerifier.MatchString(verifier) {
		return "", errors.New("norbix: a PKCE code verifier is 43 to 128 characters of A-Z a-z 0-9 - . _ ~")
	}
	sum := sha256.Sum256([]byte(verifier))
	return base64.RawURLEncoding.EncodeToString(sum[:]), nil
}
