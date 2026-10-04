# Scheduler — Go

[← Back to project README](../../README.md)

The scheduler runs a task on a cron schedule. Every Scheduler endpoint the
gateway exposes has a `hub.SchedulerModule` method (`client.Hub.Scheduler`).
All 8 are covered by `norbix/hub/scheduler_test.go`, which asserts the verb,
the fully-resolved path, the query string and the JSON body of each one.

Each method takes a context, any route id as a plain argument, an untyped
request (`map[string]any` — sent as the query string for `GET` / `DELETE`, as
the JSON body otherwise), and a pointer to decode the response into (pass `nil`
to discard it).

## Methods

| method | verb | path | params go in |
|---|---|---|---|
| `EnableScheduler(ctx, req, out)` | `PUT` | `/scheduler/enable` | — |
| `DisableScheduler(ctx, req, out)` | `PUT` | `/scheduler/disable` | — |
| `GetSchedulerTasks(ctx, req, out)` | `GET` | `/scheduler/tasks` | query: `pageSize`, `startingAfter`, `endingBefore`, `type`, `enabled` |
| `GetSchedulerTask(ctx, id, req, out)` | `GET` | `/scheduler/tasks/{id}` | path: `id` |
| `SaveSchedulerTask(ctx, req, out)` | `POST` | `/scheduler/tasks` | body (see below) |
| `EnableSchedulerTask(ctx, id, req, out)` | `PUT` | `/scheduler/tasks/{Id}/enable` | path: `id` |
| `DisableSchedulerTask(ctx, id, req, out)` | `PUT` | `/scheduler/tasks/{Id}/disable` | path: `id` |
| `DeleteSchedulerTask(ctx, id, req, out)` | `DELETE` | `/scheduler/tasks/{Id}` | path: `id` |

Every path is prefixed with the Hub version (`/v2/...`). Task ids look like
`tsk_…`.

> **Module enable / disable are `PUT`.** Up to v2.6 the SDK sent `GET`; the
> gateway now answers only `PUT` on `/scheduler/enable` and
> `/scheduler/disable`, so the call from an older SDK fails. Upgrade the SDK.

## Turn the module on

```go
if err := client.Hub.Scheduler.EnableScheduler(ctx, nil, nil); err != nil {
	log.Fatal(err)
}
```

## Save a task (create or update)

**Only `EmailCampaign` tasks run today.** The other `SchedulerTaskType`
values (`PushCampaign`, `SmsCampaign`, `CodeFunctionalCall`, `WebhookCall`)
exist in the type but the gateway refuses them on save.

The request:

| field | required | meaning |
|---|---|---|
| `initiatorUserId` | yes | the user the task runs as (`usr_…`): you, or a service user of this project |
| `name` | yes | display name |
| `cron` | yes | exactly **5 fields** (minute hour day-of-month month day-of-week), evaluated in **UTC** — `0 9 * * 1` is every Monday 09:00 UTC |
| `isEnabled` | yes | start enabled or not |
| `stopOnError` | yes | disable the task after a failed run |
| `task` | yes | what to run — `dtos.EmailCampaignSchedulerTaskRequest` |
| `taskId` | on update | the `tsk_…` id to update; leave it out to create |
| `description` | no | notes |

The task is `dtos.EmailCampaignSchedulerTaskRequest`: `Type` must be
`dtos.SchedulerTaskTypeEmailCampaign` (sent as `"type": "EmailCampaign"`),
`Campaign` is the email campaign to send (the same shape the email campaign
endpoints take: `source` + `templateId` + the audience fields of that source),
and `DatabaseIntegrationId` is optional.

```go
import "github.com/norbix-code/sdk-go/v2/norbix/hub/dtos"

var saved dtos.IdResponse
err := client.Hub.Scheduler.SaveSchedulerTask(ctx, map[string]any{
	"initiatorUserId": "usr_123",
	"name":            "Weekly digest",
	"cron":            "0 9 * * 1", // Mondays 09:00 UTC
	"isEnabled":       true,
	"stopOnError":     false,
	"task": dtos.EmailCampaignSchedulerTaskRequest{
		Type: dtos.SchedulerTaskTypeEmailCampaign,
		Campaign: &dtos.EmailCampaignRequest{
			Source:        dtos.EmailCampaignRecipientsSourceTypesAllUsers,
			TemplateId:    "etpl_123",
			IntegrationId: "eint_123", // the email provider, required
		},
	},
}, &saved)
if err != nil {
	log.Fatal(err)
}
fmt.Println(saved.Id) // tsk_…
```

This goes on the wire as:

```json
{
  "initiatorUserId": "usr_123",
  "name": "Weekly digest",
  "cron": "0 9 * * 1",
  "isEnabled": true,
  "stopOnError": false,
  "task": {
    "type": "EmailCampaign",
    "campaign": { "source": "AllUsers", "templateId": "etpl_123" }
  }
}
```

`dtos.EmailCampaignRequest` holds only the fields every source shares. To add
the audience fields of a source (for example `rolesNames` / `userTags` for
`AllUsers`, `userRecipients` for `SpecifiedUsers`), pass the campaign as a map
— or the matching `dtos.EmailTo…DeliverySettingsRequest` struct with `Source`
set — instead:

```go
"task": map[string]any{
	"type": "EmailCampaign",
	"campaign": dtos.EmailToAllUsersDeliverySettingsRequest{
		Source:     dtos.EmailCampaignRecipientsSourceTypesAllUsers,
		RolesNames: []string{"subscriber"},
		EmailCampaignRequest: dtos.EmailCampaignRequest{TemplateId: "etpl_123"},
	},
},
```

To update, send the same request with `"taskId": "tsk_…"`.

## Read tasks

```go
var list dtos.GetSchedulerTasksResponse
err := client.Hub.Scheduler.GetSchedulerTasks(ctx, map[string]any{
	"pageSize": 20,
	"type":     "EmailCampaign", // optional filter
	"enabled":  true,            // optional filter
}, &list)

var one dtos.GetSchedulerTaskResponse
err = client.Hub.Scheduler.GetSchedulerTask(ctx, "tsk_123", nil, &one)
// one.Item.PayloadJson holds the saved task body as JSON.
```

## Pause, resume, delete

```go
err := client.Hub.Scheduler.DisableSchedulerTask(ctx, "tsk_123", nil, nil)
err = client.Hub.Scheduler.EnableSchedulerTask(ctx, "tsk_123", nil, nil)
err = client.Hub.Scheduler.DeleteSchedulerTask(ctx, "tsk_123", nil, nil)
```
