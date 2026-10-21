// Hand-written, not generated (gateway campaign audit/logging, Logs stage A).
//
// Written in the shape of the other modules (payments.go, files.go), whose
// generator was lost; the request and response types are the generated dtos
// (WriteLogs, WriteLogEntry, WriteLogsResponse).

package api

import (
	"context"

	"github.com/norbix-code/sdk-go/v2/norbix/internal/transport"
)

// LogsModule groups the log endpoints on the API: an app writes its own log
// entries into the project's Logs. Reading them is on the Hub
// (client.Hub.Logs.GetLogs / GetLog).
type LogsModule struct{ t *transport.Transport }

// WriteLogs performs POST /{version}/logs (scope: project).
// Body: entries, 1 to 100 of them (request type dtos.WriteLogs). Each entry
// (dtos.WriteLogEntry) has a title (required, at most 200 characters) and
// optionally level ("information", "warning" or "error"), message (at most
// 8000), eventCode (a-z 0-9 _ . : -, at most 64, stored as app:<code>),
// correlationId (letters, digits, _ and -, at most 64), timestamp (UTC, at
// most 7 days old and 5 minutes ahead) and meta (a string map, at most 20
// keys). Answers dtos.WriteLogsResponse {accepted: n}.
//
// The call returns at once: the entries are queued and reach the project's
// integrations afterwards; a failed delivery is never reported to the app.
// Refusals: HTTP 400 for a broken entry, HTTP 400 CM-ERRORS-LOGGING-002 when
// the Logs module is off, and HTTP 429 CM-ERRORS-LOGGING-003 (a
// *errors.RateLimitError) when the project's per-minute quota is used up — the
// whole call is refused; the gateway's Retry-After header says how many
// seconds to wait. The SDK does not retry a POST.
func (m *LogsModule) WriteLogs(ctx context.Context, req map[string]any, out any) error {
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/logs",
		Method:     "POST",
		PathParams: nil,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}
