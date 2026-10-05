package hub

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/norbix-code/sdk-go/v2/norbix/hub/dtos"
)

// The signed-in account user's own profile and phone, and the team list.
// The gateway takes the account from the session for all three
// (Hub.Account/Account/Team/GetAll.cs and the account/me services), so the
// client needs only a token: every test here runs WITHOUT an AccountID and
// the call must still go out. Nothing leaves the process.

type capturedRequest struct {
	method string
	path   string
	query  map[string][]string
	body   map[string]any
}

func captureAccount(t *testing.T, answer string, call func(m *AccountModule) error) capturedRequest {
	t.Helper()
	var got capturedRequest
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got.method = r.Method
		got.path = r.URL.Path
		got.query = r.URL.Query()
		raw, _ := io.ReadAll(r.Body)
		if len(raw) > 0 {
			_ = json.Unmarshal(raw, &got.body)
		}
		if answer == "" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(answer))
	}))
	defer srv.Close()

	if err := call(newProjectTestModule(srv.URL, "")); err != nil {
		t.Fatalf("unexpected error (no AccountID set): %v", err)
	}
	return got
}

func TestGetMyAccountUserProfileReadsTheProfileAndPhone(t *testing.T) {
	var out dtos.GetMyAccountUserProfileResponse
	got := captureAccount(t,
		`{"item":{"id":"user_1","email":"ada@example.test","generalInfo":{"phone":"+37060000000"}}}`,
		func(m *AccountModule) error {
			return m.GetMyAccountUserProfile(context.Background(), nil, &out)
		})

	if got.method != http.MethodGet || got.path != "/v2/account/me" {
		t.Errorf("route: got %s %s want GET /v2/account/me", got.method, got.path)
	}
	if out.Item == nil || out.Item.Email != "ada@example.test" {
		t.Fatalf("profile not decoded: %+v", out.Item)
	}
	if out.Item.GeneralInfo == nil || out.Item.GeneralInfo.Phone != "+37060000000" {
		t.Errorf("phone not decoded: %+v", out.Item.GeneralInfo)
	}
}

func TestUpdateMyAccountUserPhoneSendsThePhoneInAPutBody(t *testing.T) {
	got := captureAccount(t, "", func(m *AccountModule) error {
		return m.UpdateMyAccountUserPhone(context.Background(), map[string]any{"phone": "+37060000000"}, nil)
	})

	if got.method != http.MethodPut || got.path != "/v2/account/me/phone" {
		t.Errorf("route: got %s %s want PUT /v2/account/me/phone", got.method, got.path)
	}
	assertFields(t, got.body, map[string]any{"phone": "+37060000000"})
}

// The team list pages with flat query fields (no nested "pagingArgs") and can
// be narrowed to one project's collaborators.
func TestGetAccountCollaboratorsSendsFlatPagingAndTheProjectFilter(t *testing.T) {
	got := captureAccount(t, "", func(m *AccountModule) error {
		return m.GetAccountCollaborators(context.Background(), map[string]any{
			"projectId":     "proj_2",
			"pageSize":      50,
			"startingAfter": "member_20",
		}, nil)
	})

	if got.method != http.MethodGet || got.path != "/v2/account/collaborators" {
		t.Errorf("route: got %s %s want GET /v2/account/collaborators", got.method, got.path)
	}
	want := map[string]string{"projectId": "proj_2", "pageSize": "50", "startingAfter": "member_20"}
	for k, v := range want {
		if vals := got.query[k]; len(vals) != 1 || vals[0] != v {
			t.Errorf("query %s: got %v want %q", k, vals, v)
		}
	}
	if _, ok := got.query["pagingArgs"]; ok {
		t.Errorf("query must not carry pagingArgs: %v", got.query)
	}
}
