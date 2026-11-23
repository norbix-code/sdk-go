package hub

import (
	"context"

	"github.com/norbix-code/sdk-go/v2/norbix/internal/transport"
)

// IntegrationsModule groups the integrations endpoints on the HUB — helpers
// shared by every module's integrations (files, email, code, …).
type IntegrationsModule struct{ t *transport.Transport }

// GetKeylessIntegrationSetup performs GET /{version}/integrations/keyless/setup (scope: project).
//
// It returns what a developer needs before saving a keyless cloud integration:
// the integration id to save with, the token issuer, subject and audience, whether
// keyless works on this gateway (with the reason when it does not) and the cloud
// setup script, and templates: ready CloudFormation / Bicep / Terraform files for
// this integration (Templates: Kind, FileName, Content). The request takes
// "module" (files, email or code), "cloud" (azure, aws or google) and an optional
// "integrationId" (an existing integration); see
// dtos.GetKeylessIntegrationSetup and dtos.GetKeylessIntegrationSetupResponse.
func (m *IntegrationsModule) GetKeylessIntegrationSetup(ctx context.Context, req map[string]any, out any) error {
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetHub,
		Path:       "/{version}/integrations/keyless/setup",
		Method:     "GET",
		PathParams: nil,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// GetQuickConnectProviders performs GET /{version}/integrations/connect/providers (scope: optional).
//
// It lists the providers that can quick-connect (one-click OAuth) on this
// install: available or not, the inputs to ask for before the provider's page
// opens (for example storeUrl), the modes (test / live) and the redirect URL
// Norbix registered. The request takes an optional "module" (for example
// payments; empty = every module); see dtos.GetQuickConnectProviders and
// dtos.GetQuickConnectProvidersResponse. Read it before
// Payments.StartPaymentsIntegrationConnect.
func (m *IntegrationsModule) GetQuickConnectProviders(ctx context.Context, req map[string]any, out any) error {
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetHub,
		Path:       "/{version}/integrations/connect/providers",
		Method:     "GET",
		PathParams: nil,
		Body:       req,
		Scope:      transport.ScopeOptional,
	}, out)
}

// GetIntegrationConnectTargets performs GET /{version}/integrations/connect/{ConnectId}/targets (scope: project).
//
// Once the person approved on the provider's page, it lists the accounts
// (profiles, shops) the provider offered for the connect id — pick one and
// finish with Payments.CompletePaymentsIntegrationConnect; see
// dtos.GetIntegrationConnectTargetsResponse (Targets: Id, Name, Detail).
func (m *IntegrationsModule) GetIntegrationConnectTargets(ctx context.Context, connectId string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"ConnectId": connectId,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetHub,
		Path:       "/{version}/integrations/connect/{ConnectId}/targets",
		Method:     "GET",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// GetIntegrationSecrets performs GET /{version}/integrations/{IntegrationId}/secrets (scope: project).
//
// It lists which secrets of an integration (any module) are saved and which are
// optional — never a value or any part of one. Each row has Key, Field (the save
// request property that fills it), IsSet and IsOptional; see
// dtos.GetIntegrationSecretsResponse. The list is empty for keyless and
// Norbix-owned integrations and for MCP, Marketplace and Webhook integrations.
// An optional "env" in req picks the environment.
func (m *IntegrationsModule) GetIntegrationSecrets(ctx context.Context, integrationId string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"IntegrationId": integrationId,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetHub,
		Path:       "/{version}/integrations/{IntegrationId}/secrets",
		Method:     "GET",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// RemoveIntegrationSecrets performs POST /{version}/integrations/{IntegrationId}/secrets/remove (scope: project).
//
// It removes optional secrets of an integration. The request takes "keys" (as
// GetIntegrationSecrets lists them) and an optional "env". A key the integration
// needs is refused with CM-ERRORS-SECRETS-012 (400) and nothing is written; see
// dtos.RemoveIntegrationSecretsRequest and dtos.IdResponse.
func (m *IntegrationsModule) RemoveIntegrationSecrets(ctx context.Context, integrationId string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"IntegrationId": integrationId,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetHub,
		Path:       "/{version}/integrations/{IntegrationId}/secrets/remove",
		Method:     "POST",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}
