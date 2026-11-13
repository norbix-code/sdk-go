package api

import (
	"net/url"
	"testing"

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
