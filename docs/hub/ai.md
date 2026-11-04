# AI integrations — Go

[← Back to project README](../../README.md)

The project's LLM and MCP integrations on the Hub, and the `hub.AiModule`
method that calls each one. An LLM integration connects a model provider; an
MCP integration connects an outside MCP server whose tools the project's AI
can use. Every route is covered by `norbix/hub/ai_integrations_test.go`
(`SetLlmIntegrationAsDefault` and the embedding routes by
`norbix/hub/ai_embeddings_test.go`), which asserts the verb and the
fully-resolved path.

Each method takes a context, any route ids as plain arguments, an untyped
request body, and a pointer to decode the response into (pass `nil` to discard
it). All are project-scoped: a key or bearer token is enough.

Project AI settings, assistants and AI service users are on
`hub.AccountModule` — see [account.md](account.md).

```go
client, err := norbix.New(norbix.Options{
    ProjectID: "proj_123",
    APIKey:    "sk_live_...",
})
if err != nil {
    log.Fatal(err)
}

var llms map[string]any
err = client.Hub.Ai.GetLlmIntegrations(ctx, nil, &llms)
```

## LLM integrations

| method | verb | path |
|---|---|---|
| `GetLlmIntegrations(ctx, req, out)` | `GET` | `/ai/integrations/llms/integrations` |
| `SaveLlmIntegration(ctx, req, out)` | `POST` | `/ai/integrations/llms/` |
| `TestLlmIntegration(ctx, req, out)` | `POST` | `/ai/integrations/llms/test` |
| `GetLlmIntegration(ctx, id, req, out)` | `GET` | `/ai/integrations/llms/{id}` |
| `DeleteLlmIntegration(ctx, id, req, out)` | `DELETE` | `/ai/integrations/llms/{Id}` |
| `EnableLlmIntegration(ctx, id, req, out)` | `PUT` | `/ai/integrations/llms/{Id}/enable` |
| `DisableLlmIntegration(ctx, id, req, out)` | `PUT` | `/ai/integrations/llms/{Id}/disable` |
| `SetLlmIntegrationAsDefault(ctx, id, req, out)` | `PUT` | `/ai/integrations/llms/{Id}/default` |

## MCP integrations

| method | verb | path |
|---|---|---|
| `GetMcpIntegrations(ctx, req, out)` | `GET` | `/ai/integrations/mcp/integrations` |
| `SaveMcpIntegration(ctx, req, out)` | `POST` | `/ai/integrations/mcp/` |
| `TestMcpIntegration(ctx, req, out)` | `POST` | `/ai/integrations/mcp/test` |
| `GetMcpIntegration(ctx, id, req, out)` | `GET` | `/ai/integrations/mcp/{id}` |
| `DeleteMcpIntegration(ctx, id, req, out)` | `DELETE` | `/ai/integrations/mcp/{Id}` |
| `EnableMcpIntegration(ctx, id, req, out)` | `PUT` | `/ai/integrations/mcp/{Id}/enable` |
| `DisableMcpIntegration(ctx, id, req, out)` | `PUT` | `/ai/integrations/mcp/{Id}/disable` |

```go
// Turn an MCP integration off without deleting it.
err := client.Hub.Ai.DisableMcpIntegration(ctx, "mcp_123", nil, nil)
```

## AI triggers

A trigger on an AI project event: `"type": "Ai"`, `"when"` one of
`OnCreditsWarning`, `OnCreditsExhausted` (the project's AI credits),
`OnQuotaWarning`, `OnQuotaExhausted` (a user's AI plan quota), and any trigger
action. Every route is asserted in `norbix/hub/triggers_routes_test.go`.

| Method | Verb | Route |
|---|---|---|
| `SaveAiProjectTrigger(ctx, req, out)` | `POST` | `/ai/triggers` |
| `GetAiProjectTriggers(ctx, req, out)` | `GET` | `/ai/triggers` |
| `GetAiProjectTrigger(ctx, id, req, out)` | `GET` | `/ai/triggers/{id}` |
| `EnableAiProjectTrigger(ctx, triggerId, req, out)` | `PATCH` | `/ai/triggers/{triggerId}/enable` |
| `DisableAiProjectTrigger(ctx, triggerId, req, out)` | `PATCH` | `/ai/triggers/{triggerId}/disable` |
| `DeleteAiProjectTrigger(ctx, triggerId, req, out)` | `DELETE` | `/ai/triggers/{triggerId}` |

```go
var saved dtos.IdResponse
err := client.Hub.Ai.SaveAiProjectTrigger(ctx, map[string]any{
	"trigger": map[string]any{
		"type": "Ai", "name": "Credits running out", "when": "OnCreditsWarning",
		"action": map[string]any{"type": "Email", "integrationId": "int_mail", "templateId": "tpl_credits"},
	},
}, &saved)
```
