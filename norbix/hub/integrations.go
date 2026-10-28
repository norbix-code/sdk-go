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
// setup script. The request takes "module" (files, email or code), "cloud" (azure,
// aws or google) and an optional "integrationId" (an existing integration); see
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
