package norbix

import (
	"context"
	stderrors "errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/norbix-code/sdk-go/v2/norbix/api/dtos"
	norbixerr "github.com/norbix-code/sdk-go/v2/norbix/errors"
)

// E-mail verification for password sign-up: the confirm answer decodes into
// dtos.EmailAddressVerifiedResponse, and the sign-in refusal the gateway sends
// while the project blocks unverified logins (HTTP 401,
// CM-ERRORS-MEMBERSHIP-USERS-052) is recognised by IsEmailNotVerified.
//
// Every case starts its own fake gateway; nothing leaves the process.

func TestEmailNotVerifiedCodeIsTheGatewayCode(t *testing.T) {
	if norbixerr.CodeEmailNotVerified != "CM-ERRORS-MEMBERSHIP-USERS-052" {
		t.Errorf("code: got %q", norbixerr.CodeEmailNotVerified)
	}
}

func TestVerifyEmailAddressDecodesTheAnswer(t *testing.T) {
	t.Setenv("NORBIX_API_KEY", "")
	t.Setenv("NORBIX_BEARER_TOKEN", "")
	var gotBody string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf := make([]byte, 256)
		n, _ := r.Body.Read(buf)
		gotBody = string(buf[:n])
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"authId":"auth_1","verifiedAtUtc":"2026-11-22T10:00:00Z","responseStatus":{"isSuccess":true}}`))
	}))
	t.Cleanup(srv.Close)
	c, err := New(Options{ProjectID: "proj_1", BaseURLAPI: srv.URL, BaseURLHub: srv.URL, APIVersion: "v3", HubVersion: "v2"})
	if err != nil {
		t.Fatalf("client: %v", err)
	}

	var out dtos.EmailAddressVerifiedResponse
	if err := c.API.Membership.VerifyEmailAddress(context.Background(), map[string]any{"token": "evt_1"}, &out); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotBody != `{"token":"evt_1"}` {
		t.Errorf("body: got %s", gotBody)
	}
	if out.AuthId != "auth_1" || out.VerifiedAtUtc != "2026-11-22T10:00:00Z" {
		t.Errorf("answer: got %+v", out)
	}
}

func TestA401WithCode052IsEmailNotVerified(t *testing.T) {
	c := newAnsweringServer(t, http.StatusUnauthorized, `{"responseStatus":{"isSuccess":false,"errors":[
		{"errorCode":"CM-ERRORS-MEMBERSHIP-USERS-052","message":"Verify your e-mail address before you sign in."}]}}`,
		"application/json")

	err := aCall(c)
	if !norbixerr.IsEmailNotVerified(err) {
		t.Fatalf("IsEmailNotVerified: got false for %v", err)
	}
	var auth *norbixerr.AuthenticationError
	if !stderrors.As(err, &auth) {
		t.Errorf("type: got %T want *errors.AuthenticationError", err)
	}
	if !norbixerr.IsEmailNotVerified(fmt.Errorf("sign in: %w", err)) {
		t.Error("IsEmailNotVerified: a wrapped refusal must still be recognised")
	}
}

func TestCode052LaterInTheErrorListIsEmailNotVerified(t *testing.T) {
	c := newAnsweringServer(t, http.StatusUnauthorized, `{"responseStatus":{"isSuccess":false,"errors":[
		{"errorCode":"CM-ERRORS-MEMBERSHIP-USERS-001","message":"Sign-in failed."},
		{"errorCode":"CM-ERRORS-MEMBERSHIP-USERS-052","message":"Verify your e-mail address."}]}}`,
		"application/json")

	if err := aCall(c); !norbixerr.IsEmailNotVerified(err) {
		t.Fatalf("IsEmailNotVerified: got false for %v", err)
	}
}

func TestOtherErrorsAreNotEmailNotVerified(t *testing.T) {
	c := newAnsweringServer(t, http.StatusUnauthorized, `{"responseStatus":{"isSuccess":false,"errors":[
		{"errorCode":"CM-ERRORS-MEMBERSHIP-USERS-050","message":"Credentials changed."}]}}`,
		"application/json")

	if err := aCall(c); norbixerr.IsEmailNotVerified(err) {
		t.Errorf("IsEmailNotVerified: got true for %v", err)
	}
	if norbixerr.IsEmailNotVerified(nil) {
		t.Error("IsEmailNotVerified(nil): got true")
	}
	if norbixerr.IsEmailNotVerified(stderrors.New("CM-ERRORS-MEMBERSHIP-USERS-052")) {
		t.Error("IsEmailNotVerified: a plain error carrying the text is not the refusal")
	}
}
