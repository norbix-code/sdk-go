package api

import (
	"context"

	"github.com/norbix-code/sdk-go/v2/norbix/internal/transport"
)

// AiModule groups the end-user AI chat endpoints on the API, for a signed-in
// project user. StartEndUserChatTurn answers at once with a turnId; the answer
// streams over the gateway's SSE endpoint on the user's own channel
// "ai-chat:{projectId}:{authId}" (events ai.chat.turn.* and ai.chat.session.*).
// A subscription to another user's channel is refused with HTTP 403 and
// responseStatus.errorCode "AiChatChannelRefused" before the stream starts —
// do not retry it.
type AiModule struct{ t *transport.Transport }

// GetEndUserChatAvailability performs GET /{version}/ai/chat/availability (scope: project).
func (m *AiModule) GetEndUserChatAvailability(ctx context.Context, req map[string]any, out any) error {
	return m.t.Send(ctx, transport.Request{
		Target: transport.TargetAPI,
		Path:   "/{version}/ai/chat/availability",
		Method: "GET",
		Body:   req,
		Scope:  transport.ScopeProject,
	}, out)
}

// ListEndUserChatSessions performs GET /{version}/ai/chat/sessions (scope: project).
func (m *AiModule) ListEndUserChatSessions(ctx context.Context, req map[string]any, out any) error {
	return m.t.Send(ctx, transport.Request{
		Target: transport.TargetAPI,
		Path:   "/{version}/ai/chat/sessions",
		Method: "GET",
		Body:   req,
		Scope:  transport.ScopeProject,
	}, out)
}

// CreateEndUserChatSession performs POST /{version}/ai/chat/sessions (scope: project).
func (m *AiModule) CreateEndUserChatSession(ctx context.Context, req map[string]any, out any) error {
	return m.t.Send(ctx, transport.Request{
		Target: transport.TargetAPI,
		Path:   "/{version}/ai/chat/sessions",
		Method: "POST",
		Body:   req,
		Scope:  transport.ScopeProject,
	}, out)
}

// GetEndUserChatSession performs GET /{version}/ai/chat/sessions/{SessionId} (scope: project).
func (m *AiModule) GetEndUserChatSession(ctx context.Context, sessionId string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"SessionId": sessionId,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/ai/chat/sessions/{SessionId}",
		Method:     "GET",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// RenameEndUserChatSession performs PATCH /{version}/ai/chat/sessions/{SessionId} (scope: project).
func (m *AiModule) RenameEndUserChatSession(ctx context.Context, sessionId string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"SessionId": sessionId,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/ai/chat/sessions/{SessionId}",
		Method:     "PATCH",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// DeleteEndUserChatSession performs DELETE /{version}/ai/chat/sessions/{SessionId} (scope: project).
func (m *AiModule) DeleteEndUserChatSession(ctx context.Context, sessionId string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"SessionId": sessionId,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/ai/chat/sessions/{SessionId}",
		Method:     "DELETE",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// PinEndUserChatSession performs PUT /{version}/ai/chat/sessions/{SessionId}/pin (scope: project).
func (m *AiModule) PinEndUserChatSession(ctx context.Context, sessionId string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"SessionId": sessionId,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/ai/chat/sessions/{SessionId}/pin",
		Method:     "PUT",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// ArchiveEndUserChatSession performs PUT /{version}/ai/chat/sessions/{SessionId}/archive (scope: project).
func (m *AiModule) ArchiveEndUserChatSession(ctx context.Context, sessionId string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"SessionId": sessionId,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/ai/chat/sessions/{SessionId}/archive",
		Method:     "PUT",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// GetEndUserChatEntries performs GET /{version}/ai/chat/sessions/{SessionId}/entries (scope: project).
func (m *AiModule) GetEndUserChatEntries(ctx context.Context, sessionId string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"SessionId": sessionId,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/ai/chat/sessions/{SessionId}/entries",
		Method:     "GET",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// SetEndUserChatEntryFeedback performs PUT /{version}/ai/chat/sessions/{SessionId}/entries/{EntryId}/feedback (scope: project).
func (m *AiModule) SetEndUserChatEntryFeedback(ctx context.Context, sessionId string, entryId string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"SessionId": sessionId,
		"EntryId":   entryId,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/ai/chat/sessions/{SessionId}/entries/{EntryId}/feedback",
		Method:     "PUT",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// ListEndUserChatAttachments performs GET /{version}/ai/chat/sessions/{SessionId}/attachments (scope: project).
func (m *AiModule) ListEndUserChatAttachments(ctx context.Context, sessionId string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"SessionId": sessionId,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/ai/chat/sessions/{SessionId}/attachments",
		Method:     "GET",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// UploadEndUserChatAttachment performs POST /{version}/ai/chat/sessions/{SessionId}/attachments (scope: project).
func (m *AiModule) UploadEndUserChatAttachment(ctx context.Context, sessionId string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"SessionId": sessionId,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/ai/chat/sessions/{SessionId}/attachments",
		Method:     "POST",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// DeleteEndUserChatAttachment performs DELETE /{version}/ai/chat/attachments/{AttachmentId} (scope: project).
func (m *AiModule) DeleteEndUserChatAttachment(ctx context.Context, attachmentId string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"AttachmentId": attachmentId,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/ai/chat/attachments/{AttachmentId}",
		Method:     "DELETE",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// ListEndUserChatMemory performs GET /{version}/ai/chat/memory (scope: project).
func (m *AiModule) ListEndUserChatMemory(ctx context.Context, req map[string]any, out any) error {
	return m.t.Send(ctx, transport.Request{
		Target: transport.TargetAPI,
		Path:   "/{version}/ai/chat/memory",
		Method: "GET",
		Body:   req,
		Scope:  transport.ScopeProject,
	}, out)
}

// ForgetEndUserChatMemory performs DELETE /{version}/ai/chat/memory/{NoteId} (scope: project).
func (m *AiModule) ForgetEndUserChatMemory(ctx context.Context, noteId string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"NoteId": noteId,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/ai/chat/memory/{NoteId}",
		Method:     "DELETE",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// StartEndUserChatTurn performs POST /{version}/ai/chat/turn (scope: project).
func (m *AiModule) StartEndUserChatTurn(ctx context.Context, req map[string]any, out any) error {
	return m.t.Send(ctx, transport.Request{
		Target: transport.TargetAPI,
		Path:   "/{version}/ai/chat/turn",
		Method: "POST",
		Body:   req,
		Scope:  transport.ScopeProject,
	}, out)
}
