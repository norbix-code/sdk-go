package hub

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/norbix-code/sdk-go/v2/norbix/hub/dtos"
	"github.com/norbix-code/sdk-go/v2/norbix/internal/transport"
)

func newIntegrationsTestModule(baseURL string) *IntegrationsModule {
	return &IntegrationsModule{t: transport.New(&transport.Config{
		ProjectID:  "proj_1",
		AccountID:  "acct_1",
		APIKey:     "key_1",
		BaseURLAPI: baseURL,
		BaseURLHub: baseURL,
		APIVersion: "v2",
		HubVersion: "v2",
		Timeout:    5 * time.Second,
		Env:        "PROD",
	}, nil)}
}

// GET /{version}/integrations/keyless/setup (gateway feat/keyless-cloud): module,
// cloud and the integration id go on the query, no body, and the answer decodes
// into the generated response — the setup templates included.
func TestGetKeylessIntegrationSetup(t *testing.T) {
	var gotMethod, gotPath string
	var gotQuery map[string][]string
	var gotBody []byte
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod, gotPath, gotQuery = r.Method, r.URL.Path, r.URL.Query()
		gotBody, _ = io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"integrationId":"int_1","env":"PROD","module":"files","cloud":"aws",`+
			`"issuer":"https://issuer.example","available":true,"subject":"p/proj_1/e/PROD/files/i/int_1",`+
			`"audience":"sts.amazonaws.com","script":"aws iam create-role ...",`+
			`"templates":[{"kind":"terraform","fileName":"norbix-files-aws.tf","content":"resource \"aws_iam_role\" ..."}]}`)
	}))
	defer srv.Close()

	req := toRequestMap(t, dtos.GetKeylessIntegrationSetup{Module: "files", Cloud: "aws", IntegrationId: "int_1"})

	var out dtos.GetKeylessIntegrationSetupResponse
	if err := newIntegrationsTestModule(srv.URL).GetKeylessIntegrationSetup(context.Background(), req, &out); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if gotMethod != http.MethodGet {
		t.Errorf("verb: got %q want GET", gotMethod)
	}
	if gotPath != "/v2/integrations/keyless/setup" {
		t.Errorf("path: got %q", gotPath)
	}
	for k, want := range map[string]string{"module": "files", "cloud": "aws", "integrationId": "int_1"} {
		if got := gotQuery[k]; len(got) != 1 || got[0] != want {
			t.Errorf("query %s: got %v want %q", k, got, want)
		}
	}
	if len(gotBody) != 0 {
		t.Errorf("a GET must not carry a body, got %q", gotBody)
	}
	want := dtos.GetKeylessIntegrationSetupResponse{
		IntegrationId: "int_1", Env: "PROD", Module: "files", Cloud: "aws",
		Issuer: "https://issuer.example", Available: true, Subject: "p/proj_1/e/PROD/files/i/int_1",
		Audience: "sts.amazonaws.com", Script: "aws iam create-role ...",
		Templates: []*dtos.KeylessIntegrationSetupTemplate{
			{Kind: "terraform", FileName: "norbix-files-aws.tf", Content: `resource "aws_iam_role" ...`},
		},
	}
	if !reflect.DeepEqual(out, want) {
		t.Errorf("response: got %+v want %+v", out, want)
	}
}

// When the gateway cannot sign keyless tokens, Available is false and the
// reason comes through.
func TestGetKeylessIntegrationSetupUnavailable(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, `{"integrationId":"int_2","module":"email","cloud":"azure","available":false,"unavailableReason":"no signing key"}`)
	}))
	defer srv.Close()

	var out dtos.GetKeylessIntegrationSetupResponse
	req := toRequestMap(t, dtos.GetKeylessIntegrationSetup{Module: "email", Cloud: "azure"})
	if err := newIntegrationsTestModule(srv.URL).GetKeylessIntegrationSetup(context.Background(), req, &out); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.Available || out.UnavailableReason != "no signing key" {
		t.Errorf("got available=%v reason=%q", out.Available, out.UnavailableReason)
	}
}
