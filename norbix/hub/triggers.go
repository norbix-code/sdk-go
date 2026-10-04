package hub

import (
	"context"

	"github.com/norbix-code/sdk-go/v2/norbix/internal/transport"
)

// TriggersModule groups the trigger endpoints that span every module on the
// HUB. Saving, listing and deleting a trigger stays on its own module
// (Database.SaveSchemaTrigger, Files.SaveFilesTrigger, ...).
type TriggersModule struct{ t *transport.Transport }

// GetTriggersNeedingAttention performs GET /{version}/triggers/attention (scope: project).
//
// It lists the triggers of one kind ("triggerType", for example
// "Schema") whose last run was stopped before sending anything — for
// example because a template no longer has every project language. The
// response carries "items": one entry per trigger with its "triggerId",
// "triggerType", "reason" and "atUtc" (dtos.GetTriggersNeedingAttentionResponse).
func (m *TriggersModule) GetTriggersNeedingAttention(ctx context.Context, req map[string]any, out any) error {
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetHub,
		Path:       "/{version}/triggers/attention",
		Method:     "GET",
		PathParams: nil,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}
