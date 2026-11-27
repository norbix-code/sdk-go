package api

import (
	"context"
	"encoding/json"

	norbixerr "github.com/norbix-code/sdk-go/v2/norbix/errors"
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

// DecideEndUserChatAction performs POST /{version}/ai/chat/sessions/{SessionId}/actions/{EntryId} (scope: project).
// It is the signed-in user's own decision on an action the chat held: a tool
// that changes something (sends, pays, creates, deletes) is not run by the
// assistant but lands in the chat as an "action.pending" entry. req carries
// "decision": "apply" runs it once, as the user; "dismiss" runs nothing. The
// answer's "status" is "applied", "failed" or "skipped"; a second decision is
// refused with HTTP 409.
func (m *AiModule) DecideEndUserChatAction(ctx context.Context, sessionId string, entryId string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"SessionId": sessionId,
		"EntryId":   entryId,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/ai/chat/sessions/{SessionId}/actions/{EntryId}",
		Method:     "POST",
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

// ── AI chat coverage (2026-11-18): the end-user AI tool inventory. ──

// GetEndUserAiTools performs GET /{version}/ai/tools (scope: project).
//
// Lists the AI tools a signed-in project user's chat may call (own-scope
// tools only): name, description, toolsets, confirmation, parameters.
func (m *AiModule) GetEndUserAiTools(ctx context.Context, req map[string]any, out any) error {
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/ai/tools",
		Method:     "GET",
		PathParams: nil,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// InvokeEndUserAiTool performs POST /{version}/ai/tools/{ToolName} (scope: project).
//
// Calls one end-user AI tool by name with "argumentsJson" (the arguments as a
// JSON string) and returns its result text.
func (m *AiModule) InvokeEndUserAiTool(ctx context.Context, toolName string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"ToolName": toolName,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/ai/tools/{ToolName}",
		Method:     "POST",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// AiChatChannel is the end-user AI chat channel of one project user,
// "ai-chat:{projectId}:{authId}" (TypeScript twin: aiChatChannel).
func AiChatChannel(projectId, authId string) string {
	return "ai-chat:" + projectId + ":" + authId
}

// AiChatEvent is one realtime envelope of the end-user AI chat stream.
// EventName is ai.chat.turn.started / .token / .tool / .completed / .failed
// or ai.chat.session.* ; Payload carries the turn id and, for a token, the
// text. FrameID is the SSE frame id.
type AiChatEvent struct {
	Channel        string         `json:"channel"`
	EventName      string         `json:"eventName"`
	ProjectID      string         `json:"projectId"`
	NotificationID string         `json:"notificationId,omitempty"`
	CreatedAtUtc   string         `json:"createdAtUtc"`
	Payload        map[string]any `json:"payload,omitempty"`
	FrameID        string         `json:"-"`
}

// AiChat streams the answer of the end-user AI chat (TypeScript twin:
// norbix.aiChat). It subscribes to the signed-in user's own channel on the
// API host — GET /event-stream?channels=ai-chat:{projectId}:{authId} with the
// client's token and project — and calls onEvent for every chat envelope
// until onEvent returns false, the server closes the stream (nil) or ctx
// ends. Frames that are not an envelope (heartbeats, subscription notices)
// are skipped. Start a turn with StartEndUserChatTurn; its turnId is in each
// event's payload.
//
// Another user's channel is refused with HTTP 403 and the error code
// "AiChatChannelRefused" before the stream starts; AiChat returns that error
// and never retries it.
func (m *AiModule) AiChat(ctx context.Context, authId string, onEvent func(AiChatEvent) bool) error {
	projectId := m.t.Cfg.ProjectID
	if projectId == "" {
		return norbixerr.New("AiChat: the client has no project id.", norbixerr.CodeError)
	}
	if authId == "" {
		return norbixerr.New("AiChat: authId is required.", norbixerr.CodeError)
	}
	return m.t.Stream(ctx, transport.Request{
		Target: transport.TargetAPI,
		Path:   "/event-stream",
		Body:   map[string]any{"channels": AiChatChannel(projectId, authId)},
		Scope:  transport.ScopeProject,
	}, func(f transport.SSEEvent) bool {
		var env AiChatEvent
		if json.Unmarshal([]byte(f.Data), &env) != nil || env.EventName == "" {
			return true
		}
		env.FrameID = f.ID
		return onEvent(env)
	})
}
