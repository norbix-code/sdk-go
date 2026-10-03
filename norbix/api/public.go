package api

import (
	"context"

	"github.com/norbix-code/sdk-go/v2/norbix/internal/transport"
)

// PublicModule groups the public project endpoints on the API (project audit,
// item E1). They drive the Admin Portal before anyone signs in.
//
// These calls carry NO sign-in: no Authorization header is sent, even when the
// client is signed in. The project id travels in the path. The answer holds
// only what the project chose to make public; an unknown project answers an
// empty document, never an error that tells a stranger it exists.
type PublicModule struct{ t *transport.Transport }

// GetPublicProjectConfig performs GET /{version}/public/projects/{ProjectId}/config (scope: unauthenticated).
// It returns the project's public brand and sign-in config.
func (m *PublicModule) GetPublicProjectConfig(ctx context.Context, projectId string, out any) error {
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/public/projects/{ProjectId}/config",
		Method:     "GET",
		PathParams: map[string]string{"ProjectId": projectId},
		Scope:      transport.ScopeUnauthenticated,
	}, out)
}

// GetPublicProjectLegal performs GET /{version}/public/projects/{ProjectId}/legal/{Kind} (scope: unauthenticated).
// kind is "terms" or "privacy". The document is there only when the project
// exposes it (hub.AccountModule.UpdateProjectExposeLegal).
func (m *PublicModule) GetPublicProjectLegal(ctx context.Context, projectId string, kind string, out any) error {
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/public/projects/{ProjectId}/legal/{Kind}",
		Method:     "GET",
		PathParams: map[string]string{"ProjectId": projectId, "Kind": kind},
		Scope:      transport.ScopeUnauthenticated,
	}, out)
}
