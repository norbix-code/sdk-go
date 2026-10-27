// Hand-written, not generated (gateway campaign audit/workflows).
//
// Written in the shape of the other modules (payments.go); the request and
// response types are the generated dtos (GetCodeWorkflowWaitingRuns,
// ResumeCodeWorkflowRun, WorkflowRunsResponse, WorkflowRunResponse,
// CodeWorkflowResumeLinkResponse, …).

package api

import (
	"context"

	"github.com/norbix-code/sdk-go/v2/norbix/internal/transport"
)

// CodeModule groups the Code endpoints on the API: the workflow (beta) routes
// an app user needs to answer a waiting run — the inbox of runs waiting on
// them, the resume by run id and the signed resume link from a notification.
// Workflows are in beta; bugs may occur. Path ids are positional arguments;
// leave them out of req.
type CodeModule struct{ t *transport.Transport }

// Beta: GetCodeWorkflowWaitingRuns performs GET /{version}/code/workflows/runs (scope: project).
// Query: status? (default waiting), waitingKind?, assignee? ("me" or a usr_ id), workflowId?, key?,
// pageSize?, startingAfter?, endingBefore?. The inbox of runs waiting on a person (dtos.WorkflowRunsResponse).
func (m *CodeModule) GetCodeWorkflowWaitingRuns(ctx context.Context, req map[string]any, out any) error {
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/code/workflows/runs",
		Method:     "GET",
		PathParams: nil,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// Beta: ResumeCodeWorkflowRun performs POST /{version}/code/workflows/runs/{RunId}/resume (scope: project).
// Body: stepId?, decision? ("approved" | "rejected"), payload?. Answers dtos.WorkflowRunResponse.
func (m *CodeModule) ResumeCodeWorkflowRun(ctx context.Context, runId string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"RunId": runId,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/code/workflows/runs/{RunId}/resume",
		Method:     "POST",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// Beta: OpenCodeWorkflowResumeLink performs GET /{version}/code/workflows/resume/{Token} (scope: optional — public; the signed token in the link is the key, so no API key or session is needed).
// Checks a signed resume link and shows what it would do; nothing changes.
// Answers dtos.CodeWorkflowResumeLinkResponse.
func (m *CodeModule) OpenCodeWorkflowResumeLink(ctx context.Context, token string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"Token": token,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/code/workflows/resume/{Token}",
		Method:     "GET",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeOptional,
	}, out)
}

// Beta: UseCodeWorkflowResumeLink performs POST /{version}/code/workflows/resume/{Token} (scope: optional — public; the signed token in the link is the key, so no API key or session is needed).
// Answers a signed resume link: resumes the run as the link's assignee. Answers dtos.WorkflowRunResponse.
func (m *CodeModule) UseCodeWorkflowResumeLink(ctx context.Context, token string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"Token": token,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/code/workflows/resume/{Token}",
		Method:     "POST",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeOptional,
	}, out)
}
