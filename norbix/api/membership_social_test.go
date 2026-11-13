package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/norbix-code/sdk-go/v2/norbix/internal/transport"
)

func socialModule(env string) *MembershipModule {
	return &MembershipModule{t: transport.New(&transport.Config{
		ProjectID:  "proj_1",
		APIKey:     "key_1",
		BaseURLAPI: "https://api.test",
		BaseURLHub: "https://hub.test",
		APIVersion: "v3",
		HubVersion: "v3",
		Env:        env,
	}, nil)}
}

func TestSocialSignInURLCarriesTheProjectAndNoEnvForProd(t *testing.T) {
	got, err := socialModule("PROD").SocialSignInURL("google", "", "")
	if err != nil {
		t.Fatal(err)
	}
	want := "https://api.test/v3/membership/auth/social/google?projectId=proj_1"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestSocialSignInURLEncodesModeAndReturnURL(t *testing.T) {
	got, err := socialModule("PROD").SocialSignInURL(
		"github", "Admin portal & app", "https://app.example.com/oauth/callback?next=/a b&x=1")
	if err != nil {
		t.Fatal(err)
	}
	want := "https://api.test/v3/membership/auth/social/github?projectId=proj_1" +
		"&mode=Admin+portal+%26+app" +
		"&returnUrl=https%3A%2F%2Fapp.example.com%2Foauth%2Fcallback%3Fnext%3D%2Fa+b%26x%3D1"
	if got != want {
		t.Errorf("got  %q\nwant %q", got, want)
	}
	u, _ := url.Parse(got)
	if m := u.Query().Get("mode"); m != "Admin portal & app" {
		t.Errorf("mode decodes to %q", m)
	}
	if r := u.Query().Get("returnUrl"); r != "https://app.example.com/oauth/callback?next=/a b&x=1" {
		t.Errorf("returnUrl decodes to %q", r)
	}
}

func TestSocialSignInURLPutsANonProdEnvInTheQuery(t *testing.T) {
	got, err := socialModule("TEST").SocialSignInURL("okta", "", "")
	if err != nil {
		t.Fatal(err)
	}
	want := "https://api.test/v3/membership/auth/social/okta?projectId=proj_1&env=TEST"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestSocialSignInURLEscapesTheProvider(t *testing.T) {
	got, _ := socialModule("PROD").SocialSignInURL("a/b", "", "")
	want := "https://api.test/v3/membership/auth/social/a%2Fb?projectId=proj_1"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// The gateway judges any token it gets, even on an open route: a JWT from
// before a password change is refused 401 and a token of another env 403. So
// the exchange sends no credentials, though the client has a key and a JWT.
func TestExchangeSocialSignInCodeSendsNoCredentials(t *testing.T) {
	var gotAuth, gotKey, gotProject string
	var calls int
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		gotAuth = r.Header.Get("Authorization")
		gotKey = r.Header.Get("X-Api-Key")
		gotProject = r.Header.Get("norbix-project-id")
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()
	m := &MembershipModule{t: transport.New(&transport.Config{
		ProjectID: "proj_1", APIKey: "key_1", BearerToken: "old.jwt",
		BaseURLAPI: srv.URL, BaseURLHub: srv.URL, APIVersion: "v3", HubVersion: "v3",
		Timeout: 5 * time.Second,
	}, nil)}
	if err := m.ExchangeSocialSignInCode(context.Background(), "google", map[string]any{"code": "c"}, nil); err != nil {
		t.Fatal(err)
	}
	if calls != 1 || gotAuth != "" || gotKey != "" {
		t.Errorf("calls %d, Authorization %q, X-Api-Key %q: want one call and no credentials", calls, gotAuth, gotKey)
	}
	if gotProject != "proj_1" {
		t.Errorf("norbix-project-id: got %q", gotProject)
	}
}

func TestPKCEChallengeMatchesTheRFC7636Vector(t *testing.T) {
	got, err := PKCEChallenge("dBjftJeZ4CVP-mB92K27uhbUJU1p1r_wW1gFWFOEjXk")
	if err != nil || got != "E9Melhoa2OwvFrEMTJguCHaoeK1t8URWbuGJSstw-cM" {
		t.Errorf("got %q, %v", got, err)
	}
	if _, err := PKCEChallenge("short"); err == nil {
		t.Error("a short verifier must be refused")
	}
}

func TestNewPKCEMakesAVerifierAndItsChallenge(t *testing.T) {
	a, err := NewPKCE()
	if err != nil {
		t.Fatal(err)
	}
	b, _ := NewPKCE()
	if len(a.CodeVerifier) != 43 || a.CodeVerifier == b.CodeVerifier {
		t.Errorf("verifier %q / %q", a.CodeVerifier, b.CodeVerifier)
	}
	if c, _ := PKCEChallenge(a.CodeVerifier); c != a.CodeChallenge {
		t.Errorf("challenge %q want %q", a.CodeChallenge, c)
	}
}

func TestSocialSignInURLWithOptionsCarriesTheChallengeLast(t *testing.T) {
	const challenge = "E9Melhoa2OwvFrEMTJguCHaoeK1t8URWbuGJSstw-cM"
	got, err := socialModule("PROD").SocialSignInURLWithOptions("google",
		SocialSignInOptions{Mode: "app", ReturnURL: "https://app.test/cb", CodeChallenge: challenge})
	if err != nil {
		t.Fatal(err)
	}
	want := "https://api.test/v3/membership/auth/social/google?projectId=proj_1&mode=app" +
		"&returnUrl=https%3A%2F%2Fapp.test%2Fcb&codeChallenge=" + challenge
	if got != want {
		t.Errorf("got  %q\nwant %q", got, want)
	}
}
