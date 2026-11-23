package hub

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/norbix-code/sdk-go/v2/norbix/hub/dtos"
)

// GET /{version}/integrations/{IntegrationId}/secrets (gateway docs/tasks/secret-fields.md):
// the id goes in the path, env on the query, and the rows decode without any value.
func TestGetIntegrationSecrets(t *testing.T) {
	var gotMethod, gotPath, gotEnv string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod, gotPath, gotEnv = r.Method, r.URL.Path, r.URL.Query().Get("env")
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"result":[`+
			`{"key":"secretKey","field":"secretKey","isSet":true,"isOptional":false},`+
			`{"key":"webhookSigningSecret","field":"webhookSigningSecret","isSet":false,"isOptional":true}]}`)
	}))
	defer srv.Close()

	var out dtos.GetIntegrationSecretsResponse
	req := map[string]any{"env": "TEST"}
	if err := newIntegrationsTestModule(srv.URL).GetIntegrationSecrets(context.Background(), "int_1", req, &out); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodGet || gotPath != "/v2/integrations/int_1/secrets" {
		t.Errorf("route: got %s %s", gotMethod, gotPath)
	}
	if gotEnv != "TEST" {
		t.Errorf("env: got %q want TEST", gotEnv)
	}
	want := []*dtos.IntegrationSecretDto{
		{Key: "secretKey", Field: "secretKey", IsSet: true, IsOptional: false},
		{Key: "webhookSigningSecret", Field: "webhookSigningSecret", IsSet: false, IsOptional: true},
	}
	if !reflect.DeepEqual(out.Result, want) {
		t.Errorf("rows: got %+v want %+v", out.Result, want)
	}
}

// POST /{version}/integrations/{IntegrationId}/secrets/remove: the keys go in the body;
// a needed key is refused with CM-ERRORS-SECRETS-012.
func TestRemoveIntegrationSecrets(t *testing.T) {
	var gotPath string
	var gotBody map[string]any
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		raw, _ := io.ReadAll(r.Body)
		_ = json.Unmarshal(raw, &gotBody)
		w.Header().Set("Content-Type", "application/json")
		if keys, _ := gotBody["keys"].([]any); len(keys) == 1 && keys[0] == "secretKey" {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = io.WriteString(w, `{"responseStatus":{"isSuccess":false,"errorCode":"CM-ERRORS-SECRETS-012","message":"needed"}}`)
			return
		}
		_, _ = io.WriteString(w, `{"id":"int_1"}`)
	}))
	defer srv.Close()

	m := newIntegrationsTestModule(srv.URL)
	var out dtos.IdResponse
	if err := m.RemoveIntegrationSecrets(context.Background(), "int_1", map[string]any{"keys": []string{"webhookSigningSecret"}}, &out); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotPath != "/v2/integrations/int_1/secrets/remove" {
		t.Errorf("path: got %q", gotPath)
	}
	if keys, _ := gotBody["keys"].([]any); len(keys) != 1 || keys[0] != "webhookSigningSecret" {
		t.Errorf("body keys: got %v", gotBody["keys"])
	}
	if out.Id != "int_1" {
		t.Errorf("id: got %q", out.Id)
	}

	if err := m.RemoveIntegrationSecrets(context.Background(), "int_1", map[string]any{"keys": []string{"secretKey"}}, nil); err == nil {
		t.Fatal("a needed key must answer an error")
	}
}
