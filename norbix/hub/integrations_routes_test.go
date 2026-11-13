package hub

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/norbix-code/sdk-go/v2/norbix/hub/dtos"
	"github.com/norbix-code/sdk-go/v2/norbix/internal/transport"
)

// Every integration route of the Code, Logs, Membership and Payments modules
// (the gateway's `[Route]` list under Hub.<Module>/Integrations/*.cs, branch
// refactoringV2) and the two shared quick-connect reads, with the module
// method that calls each one. Each case asserts the verb, the fully-resolved
// path (version and route id substituted, no `{…}` token left) and where the
// request lands — the query string for GET, the JSON body otherwise. The
// Files and AI integration routes have their own tables (files_test.go,
// ai_integrations_test.go).
//
// Nothing leaves the process: every call goes to a local test server, no
// provider is contacted.

const integrationRouteID = "int_1"

type integrationRouteCase struct {
	name string
	verb string
	path string
	call func(ctx context.Context, h *Namespace) error
}

func integrationRouteCases() []integrationRouteCase {
	body := map[string]any{"probe": "value"}
	id := integrationRouteID
	return []integrationRouteCase{
		// Code — Hub.Code/Integrations
		{"Code.GetCodeIntegrations", http.MethodGet, "/v2/code/integrations",
			func(ctx context.Context, h *Namespace) error { return h.Code.GetCodeIntegrations(ctx, body, nil) }},
		{"Code.GetCodeIntegration", http.MethodGet, "/v2/code/integrations/" + id,
			func(ctx context.Context, h *Namespace) error { return h.Code.GetCodeIntegration(ctx, id, body, nil) }},
		{"Code.SaveCodeIntegration", http.MethodPost, "/v2/code/integrations",
			func(ctx context.Context, h *Namespace) error { return h.Code.SaveCodeIntegration(ctx, body, nil) }},
		{"Code.TestCodeIntegration", http.MethodPost, "/v2/code/integrations/test",
			func(ctx context.Context, h *Namespace) error { return h.Code.TestCodeIntegration(ctx, body, nil) }},
		{"Code.ConfirmCodeIntegrationHumanDelivery", http.MethodPost, "/v2/code/integrations/confirm-human-delivery",
			func(ctx context.Context, h *Namespace) error {
				return h.Code.ConfirmCodeIntegrationHumanDelivery(ctx, body, nil)
			}},
		{"Code.EnableCodeIntegration", http.MethodPut, "/v2/code/integrations/" + id + "/enable",
			func(ctx context.Context, h *Namespace) error { return h.Code.EnableCodeIntegration(ctx, id, body, nil) }},
		{"Code.DisableCodeIntegration", http.MethodPut, "/v2/code/integrations/" + id + "/disable",
			func(ctx context.Context, h *Namespace) error {
				return h.Code.DisableCodeIntegration(ctx, id, body, nil)
			}},
		{"Code.DeleteCodeIntegration", http.MethodDelete, "/v2/code/integrations/" + id,
			func(ctx context.Context, h *Namespace) error { return h.Code.DeleteCodeIntegration(ctx, id, body, nil) }},

		// Logs — Hub.Logs/Integrations
		{"Logs.GetLoggingIntegrations", http.MethodGet, "/v2/logs/integrations",
			func(ctx context.Context, h *Namespace) error { return h.Logs.GetLoggingIntegrations(ctx, body, nil) }},
		{"Logs.GetLoggingIntegration", http.MethodGet, "/v2/logs/integrations/" + id,
			func(ctx context.Context, h *Namespace) error { return h.Logs.GetLoggingIntegration(ctx, id, body, nil) }},
		{"Logs.SaveLoggingIntegration", http.MethodPost, "/v2/logs/integrations",
			func(ctx context.Context, h *Namespace) error { return h.Logs.SaveLoggingIntegration(ctx, body, nil) }},
		{"Logs.TestLoggingIntegration", http.MethodPost, "/v2/logs/integrations/test",
			func(ctx context.Context, h *Namespace) error { return h.Logs.TestLoggingIntegration(ctx, body, nil) }},
		{"Logs.EnableLoggingIntegration", http.MethodPut, "/v2/logs/integrations/" + id + "/enable",
			func(ctx context.Context, h *Namespace) error {
				return h.Logs.EnableLoggingIntegration(ctx, id, body, nil)
			}},
		{"Logs.DisableLoggingIntegration", http.MethodPut, "/v2/logs/integrations/" + id + "/disable",
			func(ctx context.Context, h *Namespace) error {
				return h.Logs.DisableLoggingIntegration(ctx, id, body, nil)
			}},
		{"Logs.DeleteLoggingIntegration", http.MethodDelete, "/v2/logs/integrations/" + id,
			func(ctx context.Context, h *Namespace) error {
				return h.Logs.DeleteLoggingIntegration(ctx, id, body, nil)
			}},

		// Membership — Hub.Membership/Integrations (sign-in providers; no test route)
		{"Membership.GetMembershipIntegrations", http.MethodGet, "/v2/membership/integrations",
			func(ctx context.Context, h *Namespace) error {
				return h.Membership.GetMembershipIntegrations(ctx, body, nil)
			}},
		{"Membership.GetMembershipIntegration", http.MethodGet, "/v2/membership/integrations/" + id,
			func(ctx context.Context, h *Namespace) error {
				return h.Membership.GetMembershipIntegration(ctx, id, body, nil)
			}},
		{"Membership.SaveMembershipIntegration", http.MethodPost, "/v2/membership/integrations",
			func(ctx context.Context, h *Namespace) error {
				return h.Membership.SaveMembershipIntegration(ctx, body, nil)
			}},
		{"Membership.EnableMembershipIntegration", http.MethodPut, "/v2/membership/integrations/" + id + "/enable",
			func(ctx context.Context, h *Namespace) error {
				return h.Membership.EnableMembershipIntegration(ctx, id, body, nil)
			}},
		{"Membership.DisableMembershipIntegration", http.MethodPut, "/v2/membership/integrations/" + id + "/disable",
			func(ctx context.Context, h *Namespace) error {
				return h.Membership.DisableMembershipIntegration(ctx, id, body, nil)
			}},
		{"Membership.DeleteMembershipIntegration", http.MethodDelete, "/v2/membership/integrations/" + id,
			func(ctx context.Context, h *Namespace) error {
				return h.Membership.DeleteMembershipIntegration(ctx, id, body, nil)
			}},

		// Payments — Hub.Payments/Integrations, quick connect included
		{"Payments.GetPaymentsIntegrations", http.MethodGet, "/v2/payments/integrations",
			func(ctx context.Context, h *Namespace) error {
				return h.Payments.GetPaymentsIntegrations(ctx, body, nil)
			}},
		{"Payments.GetPaymentsIntegration", http.MethodGet, "/v2/payments/integrations/" + id,
			func(ctx context.Context, h *Namespace) error {
				return h.Payments.GetPaymentsIntegration(ctx, id, body, nil)
			}},
		{"Payments.SavePaymentsIntegration", http.MethodPost, "/v2/payments/integrations",
			func(ctx context.Context, h *Namespace) error {
				return h.Payments.SavePaymentsIntegration(ctx, body, nil)
			}},
		{"Payments.TestPaymentsIntegration", http.MethodPost, "/v2/payments/integrations/test",
			func(ctx context.Context, h *Namespace) error {
				return h.Payments.TestPaymentsIntegration(ctx, body, nil)
			}},
		{"Payments.ConfirmPaymentsIntegrationHumanDelivery", http.MethodPost, "/v2/payments/integrations/confirm-human-delivery",
			func(ctx context.Context, h *Namespace) error {
				return h.Payments.ConfirmPaymentsIntegrationHumanDelivery(ctx, body, nil)
			}},
		{"Payments.EnablePaymentsIntegration", http.MethodPut, "/v2/payments/integrations/" + id + "/enable",
			func(ctx context.Context, h *Namespace) error {
				return h.Payments.EnablePaymentsIntegration(ctx, id, body, nil)
			}},
		{"Payments.DisablePaymentsIntegration", http.MethodPut, "/v2/payments/integrations/" + id + "/disable",
			func(ctx context.Context, h *Namespace) error {
				return h.Payments.DisablePaymentsIntegration(ctx, id, body, nil)
			}},
		{"Payments.DeletePaymentsIntegration", http.MethodDelete, "/v2/payments/integrations/" + id,
			func(ctx context.Context, h *Namespace) error {
				return h.Payments.DeletePaymentsIntegration(ctx, id, body, nil)
			}},
		{"Payments.StartPaymentsIntegrationConnect", http.MethodPost, "/v2/payments/integrations/connect/start",
			func(ctx context.Context, h *Namespace) error {
				return h.Payments.StartPaymentsIntegrationConnect(ctx, body, nil)
			}},
		{"Payments.CompletePaymentsIntegrationConnect", http.MethodPost, "/v2/payments/integrations/connect/qc_1/complete",
			func(ctx context.Context, h *Namespace) error {
				return h.Payments.CompletePaymentsIntegrationConnect(ctx, "qc_1", body, nil)
			}},
		{"Payments.DisconnectPaymentsIntegrationConnect", http.MethodPost, "/v2/payments/integrations/" + id + "/connect/disconnect",
			func(ctx context.Context, h *Namespace) error {
				return h.Payments.DisconnectPaymentsIntegrationConnect(ctx, id, body, nil)
			}},

		// Shared quick-connect reads — Hub.Payments/Integrations/Connect/QuickConnect.cs
		{"Integrations.GetQuickConnectProviders", http.MethodGet, "/v2/integrations/connect/providers",
			func(ctx context.Context, h *Namespace) error {
				return h.Integrations.GetQuickConnectProviders(ctx, body, nil)
			}},
		{"Integrations.GetIntegrationConnectTargets", http.MethodGet, "/v2/integrations/connect/qc_1/targets",
			func(ctx context.Context, h *Namespace) error {
				return h.Integrations.GetIntegrationConnectTargets(ctx, "qc_1", body, nil)
			}},
	}
}

func newIntegrationRoutesNamespace(baseURL string) *Namespace {
	return NewNamespace(transport.New(&transport.Config{
		ProjectID:  "proj_1",
		AccountID:  "acct_1",
		APIKey:     "key_1",
		BaseURLAPI: "http://api.invalid",
		BaseURLHub: baseURL,
		APIVersion: "v2",
		HubVersion: "v2",
		Timeout:    5 * time.Second,
		Env:        "PROD",
	}, nil))
}

func TestIntegrationRoutesHitTheExpectedRoute(t *testing.T) {
	cases := integrationRouteCases()
	if len(cases) != 34 {
		t.Fatalf("cases: got %d want 34 (code 8, logs 7, membership 6, payments 11, shared 2)", len(cases))
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var gotMethod, gotPath, gotRawQuery, gotProject string
			var gotBody []byte
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotMethod, gotPath, gotRawQuery = r.Method, r.URL.Path, r.URL.RawQuery
				gotProject = r.Header.Get("norbix-project-id")
				gotBody, _ = io.ReadAll(r.Body)
				w.Header().Set("Content-Type", "application/json")
				_, _ = io.WriteString(w, `{}`)
			}))
			defer srv.Close()

			if err := c.call(context.Background(), newIntegrationRoutesNamespace(srv.URL)); err != nil {
				t.Fatalf("%s: unexpected error: %v", c.name, err)
			}
			if gotMethod != c.verb {
				t.Errorf("verb: got %q want %q", gotMethod, c.verb)
			}
			if gotPath != c.path {
				t.Errorf("path: got %q want %q", gotPath, c.path)
			}
			if strings.ContainsAny(gotPath, "{}") {
				t.Errorf("path: a route token was left unfilled: %q", gotPath)
			}
			if gotProject != "proj_1" {
				t.Errorf("project header: got %q want proj_1", gotProject)
			}
			// The request lands on the query for GET and DELETE, in the JSON body otherwise.
			if c.verb == http.MethodGet || c.verb == http.MethodDelete {
				if gotRawQuery != "probe=value" {
					t.Errorf("query: got %q want probe=value", gotRawQuery)
				}
				if len(gotBody) != 0 {
					t.Errorf("body: got %s want none", gotBody)
				}
				return
			}
			if gotRawQuery != "" {
				t.Errorf("query: got %q want none", gotRawQuery)
			}
			var got map[string]any
			if err := json.Unmarshal(gotBody, &got); err != nil {
				t.Fatalf("body is not JSON: %v (%s)", err, gotBody)
			}
			if got["probe"] != "value" {
				t.Errorf("body: got %v want probe=value", got)
			}
		})
	}
}

// Quick connect end to end on the generated types: the start request carries
// the provider, the name, the mode, the inputs and the reconnect id under
// their wire names and decodes the authorize URL; complete carries the target
// and isEnabled; the answers decode into the generated responses.
func TestPaymentsQuickConnectBodiesAndAnswers(t *testing.T) {
	var gotPath string
	var gotBody []byte
	answer := `{}`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		gotBody, _ = io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, answer)
	}))
	defer srv.Close()
	h := newIntegrationRoutesNamespace(srv.URL)
	ctx := context.Background()

	answer = `{"authorizeUrl":"https://shop.example.com/wc-auth/v1/authorize?x=1","connectId":"qc_1",` +
		`"redirectUrl":"https://hub.example/v2/integrations/connect/callback/woocommerce","expiresAt":"2026-11-03T10:10:00Z"}`
	start := toRequestMap(t, dtos.StartPaymentsIntegrationConnect{
		Provider:        "woocommerce",
		IntegrationName: "Shop",
		Mode:            "test",
		Input:           map[string]string{"storeUrl": "https://shop.example.com"},
		IntegrationId:   integrationRouteID,
	})
	var started dtos.StartIntegrationConnectResponse
	if err := h.Payments.StartPaymentsIntegrationConnect(ctx, start, &started); err != nil {
		t.Fatalf("start: %v", err)
	}
	var sent map[string]any
	if err := json.Unmarshal(gotBody, &sent); err != nil {
		t.Fatalf("start body is not JSON: %v (%s)", err, gotBody)
	}
	assertFields(t, sent, map[string]any{
		"provider":        "woocommerce",
		"integrationName": "Shop",
		"mode":            "test",
		"input":           map[string]any{"storeUrl": "https://shop.example.com"},
		"integrationId":   integrationRouteID,
	})
	if started.ConnectId != "qc_1" || started.AuthorizeUrl != "https://shop.example.com/wc-auth/v1/authorize?x=1" {
		t.Errorf("start answer: got %+v", started)
	}

	answer = `{"connectId":"qc_1","module":"payments","provider":"woocommerce","integrationName":"Shop",` +
		`"accountName":"Shop Ltd","targets":[{"id":"shop_42","name":"Main shop","detail":"https://shop.example.com"}]}`
	var targets dtos.GetIntegrationConnectTargetsResponse
	if err := h.Integrations.GetIntegrationConnectTargets(ctx, started.ConnectId, nil, &targets); err != nil {
		t.Fatalf("targets: %v", err)
	}
	if gotPath != "/v2/integrations/connect/qc_1/targets" {
		t.Errorf("targets path: got %q", gotPath)
	}
	if len(targets.Targets) != 1 || targets.Targets[0].Id != "shop_42" || targets.Targets[0].Name != "Main shop" {
		t.Errorf("targets answer: got %+v", targets)
	}

	answer = `{"id":"int_new"}`
	complete := toRequestMap(t, dtos.CompletePaymentsIntegrationConnect{ConnectId: "qc_1", TargetId: targets.Targets[0].Id, IsEnabled: true})
	var saved dtos.IdResponse
	if err := h.Payments.CompletePaymentsIntegrationConnect(ctx, "qc_1", complete, &saved); err != nil {
		t.Fatalf("complete: %v", err)
	}
	if gotPath != "/v2/payments/integrations/connect/qc_1/complete" {
		t.Errorf("complete path: got %q", gotPath)
	}
	if err := json.Unmarshal(gotBody, &sent); err != nil {
		t.Fatalf("complete body is not JSON: %v (%s)", err, gotBody)
	}
	assertFields(t, sent, map[string]any{"targetId": "shop_42", "isEnabled": true})
	if saved.Id != "int_new" {
		t.Errorf("complete answer: got %+v", saved)
	}

	answer = `{"wasConnected":true,"providerAnswered":false}`
	var ended dtos.DisconnectPaymentsIntegrationConnectResponse
	if err := h.Payments.DisconnectPaymentsIntegrationConnect(ctx, saved.Id, nil, &ended); err != nil {
		t.Fatalf("disconnect: %v", err)
	}
	if gotPath != "/v2/payments/integrations/int_new/connect/disconnect" {
		t.Errorf("disconnect path: got %q", gotPath)
	}
	if !ended.WasConnected || ended.ProviderAnswered {
		t.Errorf("disconnect answer: got %+v", ended)
	}

	answer = `{"providers":[{"module":"payments","provider":"woocommerce","available":true,"callbackKind":"redirect",` +
		`"requiredInputs":["storeUrl"],"modes":["test","live"],"redirectUrl":"https://hub.example/cb"}]}`
	var providers dtos.GetQuickConnectProvidersResponse
	if err := h.Integrations.GetQuickConnectProviders(ctx, toRequestMap(t, dtos.GetQuickConnectProviders{Module: "payments"}), &providers); err != nil {
		t.Fatalf("providers: %v", err)
	}
	if len(providers.Providers) != 1 || providers.Providers[0].Provider != "woocommerce" || !providers.Providers[0].Available {
		t.Errorf("providers answer: got %+v", providers)
	}
}
