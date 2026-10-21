# Logs — Go

[← Back to project README](../../README.md)

Logs has two sides:

- **Write** — your app sends its own entries with `client.API.Logs.WriteLogs`
  (`POST /logs` on the **Api** host).
- **Read and manage** — the log list, one entry, clean-up, settings and the
  log integrations (where copies of the entries go) are on the **Hub**:
  `client.Hub.Logs`.

Every method takes a context, any route id as a plain argument, an untyped
request (`map[string]any` — sent as the query string for `GET` / `DELETE`, as
the JSON body otherwise), and a pointer to decode the response into (pass `nil`
to discard it). The request and response types are in `norbix/api/dtos` and
`norbix/hub/dtos`. Tests: `norbix/api/logs_test.go`, `norbix/hub/logs_test.go`,
`norbix/hub/module_toggle_test.go`.

## Methods

| method | verb | path | params go in |
|---|---|---|---|
| `API.Logs.WriteLogs(ctx, req, out)` | `POST` | `/logs` (Api) | body: `entries` |
| `Hub.Logs.EnableLogging(ctx, req, out)` | `PUT` | `/logs/enable` | — |
| `Hub.Logs.DisableLogging(ctx, req, out)` | `PUT` | `/logs/disable` | — |
| `Hub.Logs.GetLogs(ctx, req, out)` | `GET` | `/logs` | query (see below) |
| `Hub.Logs.GetLog(ctx, id, req, out)` | `GET` | `/logs/{Id}` | path: `id` |
| `Hub.Logs.GetLogsByCorrelationId(ctx, req, out)` | `GET` | `/logs/audit` | query: `targetCorrelationId` |
| `Hub.Logs.CleanLogs(ctx, req, out)` | `POST` | `/logs/clean` | — |
| `Hub.Logs.GetLogSettings(ctx, req, out)` | `GET` | `/logs/settings` | — |
| `Hub.Logs.SaveLogSettings(ctx, req, out)` | `POST` | `/logs/settings` | body: `skipCloudDashboardLogs`, `skipHttpBodyMeta`, `aiChatLoggingEnabled` |
| `Hub.Logs.GetLoggingIntegrations(ctx, req, out)` | `GET` | `/logs/integrations` | — |
| `Hub.Logs.GetLoggingIntegration(ctx, id, req, out)` | `GET` | `/logs/integrations/{id}` | path: `id` |
| `Hub.Logs.SaveLoggingIntegration(ctx, req, out)` | `POST` | `/logs/integrations` | body: `integration` |
| `Hub.Logs.TestLoggingIntegration(ctx, req, out)` | `POST` | `/logs/integrations/test` | body: `integrationId` |
| `Hub.Logs.EnableLoggingIntegration(ctx, id, req, out)` | `PUT` | `/logs/integrations/{Id}/enable` | path: `id` |
| `Hub.Logs.DisableLoggingIntegration(ctx, id, req, out)` | `PUT` | `/logs/integrations/{Id}/disable` | path: `id` |
| `Hub.Logs.DeleteLoggingIntegration(ctx, id, req, out)` | `DELETE` | `/logs/integrations/{Id}` | path: `id` |

Every path is prefixed with the version (`/v2/...`).

## Write your app's entries

```go
import apidtos "github.com/norbix-code/sdk-go/v2/norbix/api/dtos"

var out apidtos.WriteLogsResponse
err := client.API.Logs.WriteLogs(ctx, map[string]any{
	"entries": []*apidtos.WriteLogEntry{{
		Level:         "error",
		Title:         "Checkout failed",
		Message:       "card declined",
		EventCode:     "checkout.failed", // stored as app:checkout.failed
		CorrelationId: "req_42",
		Meta:          map[string]string{"orderId": "ord_1"},
	}},
}, &out)
// out.Accepted == 1
```

One call takes 1 to 100 entries. Each entry:

| field | required | rule |
|---|---|---|
| `title` | yes | at most 200 characters |
| `level` | no | `information`, `warning` or `error` |
| `message` | no | at most 8000 characters |
| `eventCode` | no | `a-z 0-9 _ . : -`, at most 64; stored as `app:<code>` |
| `correlationId` | no | letters, digits, `_` and `-`, at most 64 |
| `timestamp` | no | RFC 3339; a time with no zone is read as UTC; at most 7 days old and 5 minutes ahead; the gateway stamps it when left out |
| `meta` | no | text → text, at most 20 keys |

The call returns at once: the entries are queued and reach the project's
integrations afterwards. A failed delivery to an integration is never reported
to the app — the project owner sees it in the Logs list.

Refusals:

| HTTP | code | when | Go type |
|---|---|---|---|
| 400 | — | an entry breaks a rule above | `*errors.ValidationError` |
| 400 | `CM-ERRORS-LOGGING-002` | the Logs module is off | `*errors.ValidationError` |
| 429 | `CM-ERRORS-LOGGING-003` | the project's per-minute quota (entries and bytes) is used up; the whole call is refused | `*errors.RateLimitError` |

The SDK never resends a `POST`. On a 429 the gateway sends a `Retry-After`
header with the seconds to wait; the SDK does not surface response headers on
its errors yet, so wait and try again (a few seconds up to a minute).

A narrow key for an app: a service user whose only grant is the built-in
policy "Logs — write" (`logging:create`).

## Read the logs

```go
import "github.com/norbix-code/sdk-go/v2/norbix/hub/dtos"

var page dtos.GetLogsResponse
err := client.Hub.Logs.GetLogs(ctx, map[string]any{
	"level":            "Error",
	"logCorrelationId": "req_42",
	"pageSize":         50,
}, &page)
for _, e := range page.List.Items {
	fmt.Println(e.Timestamp, e.Level, e.Title, e.Origin)
}

var one dtos.GetLogResponse
err = client.Hub.Logs.GetLog(ctx, page.List.Items[0].Id, nil, &one)
```

`GetLogs` filters: `level`, `module`, `logCorrelationId`, `eventCode`,
`search` (title and message), `metaKey` + `metaValue`, `fromUtc`, `toUtc`;
paging: `pageSize`, `startingAfter`, `endingBefore` (the list is newest first,
`page.List.HasMore` says whether there is more). Filter by
**`logCorrelationId`**, not `correlationId` — that one is the current
request's own id and matches nothing.

The list shows only the request env's entries (the client's `Env`, PROD when
none). A project without a Norbix Logging integration gets empty pages.
`CleanLogs` removes the request env's entries; other envs keep theirs.

## Settings

```go
err := client.Hub.Logs.SaveLogSettings(ctx, map[string]any{
	"skipCloudDashboardLogs": true,
	"skipHttpBodyMeta":       false,
	"aiChatLoggingEnabled":   false,
}, nil)
```

Send all three values as a map. The generated `dtos.SaveLogSettings` tags
every field `omitempty`, so a `false` in that struct is left out of the body.

## Integrations — save, then test

Save the integration with its provider request type inside `integration`, then
**call `TestLoggingIntegration`. Nothing is delivered to an integration until
its last test passed** — a saved but untested integration gets no entries, and
neither does one whose last test failed.

```go
var saved dtos.IdResponse
err := client.Hub.Logs.SaveLoggingIntegration(ctx, map[string]any{
	"integration": dtos.MongoDbLoggingIntegrationRequest{
		LoggingIntegrationRequest: dtos.LoggingIntegrationRequest{
			IntegrationName: "App logs",
			IsEnabled:       true,
		},
		Provider:         dtos.LoggingProviderMongoDB,
		ConnectionString: "mongodb://…",
		DatabaseName:     "app_logs", // optional
	},
}, &saved)
if err != nil {
	log.Fatal(err)
}

var tested dtos.TestLoggingIntegrationResponse
err = client.Hub.Logs.TestLoggingIntegration(ctx, map[string]any{"integrationId": saved.Id}, &tested)
```

Provider notes:

- **MongoDB** — `dtos.MongoDbLoggingIntegrationRequest`: `connectionString`
  (required), `databaseName` (optional).
- **Datadog** — `dtos.DataDogLoggingIntegrationRequest`: `site` must be one of
  `datadoghq.com`, `us3.datadoghq.com`, `us5.datadoghq.com`, `datadoghq.eu`,
  `ap1.datadoghq.com`, `ap2.datadoghq.com`, `ddog-gov.com`; plus `serviceName`,
  `environment`, `apiKey`.
- **AWS CloudWatch** — `dtos.AwsCloudWatchLoggingIntegrationRequest`: `region`,
  `logGroup`, `logStream` (both created on the first write when missing),
  `accessKey`, `secretKey` of an IAM key with `logs:CreateLogGroup`,
  `logs:CreateLogStream` and `logs:PutLogEvents`.

Always set `Provider` on the provider request — the gateway picks the request
type by it.
