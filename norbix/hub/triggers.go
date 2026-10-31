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

// CheckTriggerExecutionIdentity performs POST /{version}/triggers/execution-identity/check (scope: project).
//
// It answers whether the caller may name "executionIdentityId" (empty = the
// environment's default) for this "triggerType" and "action", and what the
// identity lacks for that action. Nothing is saved
// (dtos.CheckTriggerExecutionIdentityResponse).
func (m *TriggersModule) CheckTriggerExecutionIdentity(ctx context.Context, req map[string]any, out any) error {
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetHub,
		Path:       "/{version}/triggers/execution-identity/check",
		Method:     "POST",
		PathParams: nil,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// GetTriggerRuns performs GET /{version}/triggers/runs (scope: project).
//
// It lists the runs of one trigger, newest first, kept 30 days. Send
// "triggerType" and "triggerId", and optionally "status" (queued, running,
// succeeded, failed or skipped) and the paging fields
// (dtos.GetTriggerRunsResponse).
func (m *TriggersModule) GetTriggerRuns(ctx context.Context, req map[string]any, out any) error {
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetHub,
		Path:       "/{version}/triggers/runs",
		Method:     "GET",
		PathParams: nil,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// GetTriggerRun performs GET /{version}/triggers/runs/{Id} (scope: project).
//
// One run (id trr_…). "triggerType" is required: a run of another type answers
// not found (dtos.GetTriggerRunResponse).
func (m *TriggersModule) GetTriggerRun(ctx context.Context, id string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"Id": id,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetHub,
		Path:       "/{version}/triggers/runs/{Id}",
		Method:     "GET",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}
