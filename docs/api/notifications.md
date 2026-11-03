# Notifications — Go (API)

[← Back to project README](../../README.md)

Every notification endpoint on the **API** surface, and the `api.NotificationsModule`
method that calls it: Email, SMS and Push campaigns plus the signed-in user's own
push devices and messages. All 44 are covered by `norbix/api/notifications_test.go`,
which asserts the verb and the fully-resolved path of each one.

The Hub has the same campaign paths for the admin surface — those live on
`client.Hub.Notifications` and are documented in [docs/hub/push.md](../hub/push.md).
This module is the data-API twin: the one a project's own backend calls with an
API key, or an app calls with the user's bearer token.

Each method takes a context, any route ids as plain arguments, an untyped
request body (sent as the query string on `GET` / `DELETE`, as JSON otherwise),
and a pointer to decode the response into (pass `nil` to discard it). The
generated request and response structs live in `norbix/api/dtos`.

```go
client, err := norbix.New(norbix.Options{
    ProjectID: "proj_123",
    APIKey:    "sk_live_...",
})
if err != nil {
    log.Fatal(err)
}

var campaigns dtos.GetEmailCampaignsResponse
err = client.API.Notifications.GetEmailCampaigns(ctx, map[string]any{"pageSize": 20}, &campaigns)
```

## Own push devices and messages

The end-user side of Push: an app signed in as a user (bearer token from
`client.Login`) registers its device, reads the push messages it received and
marks them read. Each route sees only the caller's own rows — another user's
device or message is "not found".

| method | verb | path |
|---|---|---|
| `RegisterOwnPushDevice(ctx, req, out)` | `POST` | `/notifications/push/devices` |
| `GetOwnPushDevices(ctx, req, out)` | `GET` | `/notifications/push/devices` |
| `DeleteOwnPushDevice(ctx, id, req, out)` | `DELETE` | `/notifications/push/devices/{Id}` |
| `DeleteOwnPushDeviceByToken(ctx, req, out)` | `DELETE` | `/notifications/push/devices` |
| `GetOwnPushMessages(ctx, req, out)` | `GET` | `/notifications/push/messages` |
| `MarkOwnPushMessageRead(ctx, id, req, out)` | `POST` | `/notifications/push/messages/{Id}/read` |
| `MarkAllOwnPushMessagesRead(ctx, req, out)` | `POST` | `/notifications/push/messages/read` |

`RegisterOwnPushDevice` takes `provider` (`fcm` or `apns` — the **native** token;
an Expo push token is refused) and `token`, plus optional `platform`, `osName`,
`osVersion`, `appVersion`, `deviceName`, `modelName`, `brand`, `manufacturer`,
`installationId`, `locale`, `timeZone`. Sending the same token again refreshes
the device and answers the same id; a token another user held moves to the
caller.

```go
var device dtos.IdResponse
err := client.API.Notifications.RegisterOwnPushDevice(ctx, map[string]any{
    "provider": "fcm",
    "token":    "<native device token>",
    "platform": "android",
    "locale":   "lt-LT",
}, &device)

// Sign-out from this device: no id needed, safe to repeat.
err = client.API.Notifications.DeleteOwnPushDeviceByToken(ctx, map[string]any{"token": "<native device token>"}, nil)
```

`GetOwnPushMessages` is paged (`pageNumber`, `pageSize`) and newest first;
`unreadOnly: true` keeps the unread ones. The answer decodes into
`dtos.GetOwnPushMessagesResponse` (`List.Items` of `dtos.OwnPushMessageDto`).

```go
var inbox dtos.GetOwnPushMessagesResponse
err := client.API.Notifications.GetOwnPushMessages(ctx, map[string]any{"unreadOnly": true, "pageSize": 20}, &inbox)
for _, m := range inbox.List.Items {
    _ = client.API.Notifications.MarkOwnPushMessageRead(ctx, m.Id, nil, nil)
}
```

## Email campaigns

| method | verb | path |
|---|---|---|
| `CreateEmailCampaign(ctx, req, out)` | `POST` | `/notifications/email/campaigns` |
| `DeleteEmailCampaign(ctx, id, req, out)` | `DELETE` | `/notifications/email/campaigns/{Id}` |
| `GetEmailCampaign(ctx, id, req, out)` | `GET` | `/notifications/email/campaigns/{id}` |
| `GetEmailCampaigns(ctx, req, out)` | `GET` | `/notifications/email/campaigns` |
| `GetEmailCampaignBatches(ctx, id, req, out)` | `GET` | `/notifications/email/campaigns/{id}/batches` |
| `GetEmailCampaignBatchNotification(ctx, id, batchId, notificationId, req, out)` | `GET` | `/notifications/email/campaigns/{id}/batches/{batchId}/{notificationId}` |
| `GetEmailCampaignBatchNotifications(ctx, id, batchId, req, out)` | `GET` | `/notifications/email/campaigns/{id}/batches/{batchId}` |
| `GetEmailCampaignStatistics(ctx, id, req, out)` | `GET` | `/notifications/email/campaigns/{id}/stats` |
| `PreviewEmailNotification(ctx, req, out)` | `GET` | `/notifications/email/preview` |
| `RetryFailedEmailNotifications(ctx, id, req, out)` | `POST` | `/notifications/email/campaigns/{Id}/retry-failed` |
| `StopEmailCampaign(ctx, id, req, out)` | `POST` | `/notifications/email/campaigns/{Id}/stop` |
| `GetEmailCampaignMessages(ctx, campaignId, req, out)` | `GET` | `/notifications/emails/campaigns/{campaignId}/messages` |

Note the `emails` segment of `GetEmailCampaignMessages`: it is the only route
under `/notifications/emails/`, and the path is copied as the gateway exposes it.

## SMS campaigns

| method | verb | path |
|---|---|---|
| `CreateSmsCampaign(ctx, req, out)` | `POST` | `/notifications/sms/campaigns` |
| `DeleteSmsCampaign(ctx, id, req, out)` | `DELETE` | `/notifications/sms/campaigns/{id}` |
| `GetSmsCampaign(ctx, id, req, out)` | `GET` | `/notifications/sms/campaigns/{id}` |
| `GetSmsCampaigns(ctx, req, out)` | `GET` | `/notifications/sms/campaigns` |
| `GetSmsCampaignBatches(ctx, id, req, out)` | `GET` | `/notifications/sms/campaigns/{id}/batches` |
| `GetSmsCampaignBatchNotification(ctx, id, batchId, notificationId, req, out)` | `GET` | `/notifications/sms/campaigns/{id}/batches/{batchId}/{notificationId}` |
| `GetSmsCampaignBatchNotifications(ctx, id, batchId, req, out)` | `GET` | `/notifications/sms/campaigns/{id}/batches/{batchId}` |
| `GetSmsCampaignStatistics(ctx, id, req, out)` | `GET` | `/notifications/sms/campaigns/{id}/stats` |
| `PreviewSmsNotification(ctx, req, out)` | `GET` | `/notifications/sms/preview` |
| `RetryFailedSmsNotifications(ctx, id, req, out)` | `POST` | `/notifications/sms/campaigns/{Id}/retry-failed` |
| `StopSmsCampaign(ctx, id, req, out)` | `POST` | `/notifications/sms/campaigns/{Id}/stop` |
| `GetSmsCampaignMessages(ctx, campaignId, req, out)` | `GET` | `/notifications/sms/campaigns/{campaignId}/messages` |

## Push campaigns

| method | verb | path |
|---|---|---|
| `GetPushCampaignAudienceCount(ctx, req, out)` | `POST` | `/notifications/push/campaigns/audience-count` |
| `CreatePushCampaign(ctx, req, out)` | `POST` | `/notifications/push/campaigns` |
| `DeletePushCampaign(ctx, id, req, out)` | `DELETE` | `/notifications/push/campaigns/{Id}` |
| `GetPushCampaign(ctx, id, req, out)` | `GET` | `/notifications/push/campaigns/{id}` |
| `GetPushCampaigns(ctx, req, out)` | `GET` | `/notifications/push/campaigns` |
| `GetPushCampaignBatches(ctx, id, req, out)` | `GET` | `/notifications/push/campaigns/{id}/batches` |
| `GetPushCampaignBatchNotification(ctx, id, batchId, notificationId, req, out)` | `GET` | `/notifications/push/campaigns/{id}/batches/{batchId}/{notificationId}` |
| `GetPushCampaignBatchNotifications(ctx, id, batchId, req, out)` | `GET` | `/notifications/push/campaigns/{id}/batches/{batchId}` |
| `GetPushCampaignStatistics(ctx, id, req, out)` | `GET` | `/notifications/push/campaigns/{id}/stats` |
| `PreviewPushNotification(ctx, req, out)` | `GET` | `/notifications/push/preview` |
| `RetryFailedPushNotifications(ctx, id, req, out)` | `POST` | `/notifications/push/campaigns/{Id}/retry-failed` |
| `StopPushCampaign(ctx, id, req, out)` | `POST` | `/notifications/push/campaigns/{Id}/stop` |
| `GetPushCampaignMessages(ctx, campaignId, req, out)` | `GET` | `/notifications/push/campaigns/{campaignId}/messages` |

`GetPushCampaignAudienceCount` takes the same `campaign` body as
`CreatePushCampaign` and answers how many devices it would reach, before
anything is created.

## Creating a campaign

`CreateEmailCampaign` and `CreatePushCampaign` take the campaign under
`campaign` (`dtos.EmailCampaignRequest` / `dtos.PushCampaignRequest`), an
optional `databaseIntegrationId` and `sendNow`; `CreateSmsCampaign` takes the
fields flat (`templateId`, `deliveryType` and the matching audience block —
`allUsers`, `specifiedUsers`, `accountUsers`, `collection` or `phoneNumbers` —
plus optional `integrationId`, `language`, `initiatorId`, `sendNow`). All three
answer `dtos.IdResponse`. The audience shapes are the Hub's: see
[Choosing who a campaign goes to](../hub/push.md#choosing-who-a-campaign-goes-to).

```go
var created dtos.IdResponse
err := client.API.Notifications.CreatePushCampaign(ctx, map[string]any{
    "campaign": map[string]any{
        "templateId":   "tpl_123",
        "source":       "allUsers",
    },
    "sendNow": true,
}, &created)
```

## Reading a campaign

`GetEmailCampaigns`, `GetSmsCampaigns` and `GetPushCampaigns` are paged
(`pageNumber`, `pageSize`) and filter by `templateId`, `from` / `to` (a time
window) and, for Email and SMS, `campaignId`; Email also takes `emailAddress`.
`Get…Campaign` answers one campaign, `Get…CampaignStatistics` its counters
(`dtos.CampaignStatsDto`), `Get…CampaignBatches` the batches it was sent in,
`Get…CampaignBatchNotifications` one batch's rows and
`Get…CampaignBatchNotification` one row with its delivery `State`, `Attempts`
and `DeliveryEvents` (see the README section *Notifications: delivery state,
attempts, retry the failed ones*). `Get…CampaignMessages` lists the rendered
messages of a campaign, optionally narrowed to one `campaignBatchId`.

## Stop and retry

`Stop…Campaign` stops a running campaign; what was already sent stays sent.
`RetryFailed…Notifications` re-sends the campaign's retryable failures. The
body may carry `notificationIds`, `categories`, `includeAmbiguous` and an
`idempotencyKey`; the answer (`dtos.RetryFailed{Email,Sms,Push}NotificationsResponse`)
says how many were selected and skipped and whether the call was a replay.

```go
var retried dtos.RetryFailedEmailNotificationsResponse
err := client.API.Notifications.RetryFailedEmailNotifications(ctx, "camp_123", map[string]any{
    "idempotencyKey": "retry-2026-11-03-1",
}, &retried)
```

## Preview a notification with its signed link (no sign-in)

A preview link carries a signed `hash`. The hash alone is the key: the client
needs no API key and no login, so it is safe to open from a page a customer
clicks. If the client does have a key or a token, it is still sent.

```go
// No APIKey, no login — only the project id the client always needs.
client, err := norbix.New(norbix.Options{ProjectID: "proj_123"})
if err != nil {
    log.Fatal(err)
}

var preview dtos.PreviewEmailNotificationResponse
err = client.API.Notifications.PreviewEmailNotification(ctx,
    map[string]any{"hash": signedLink}, &preview)
```

`PreviewSmsNotification` and `PreviewPushNotification` work the same way. A
signed-in member can pass `projectId` + `notificationId` instead of `hash`.
The gateway answers `401` for a bad or expired link, `403` for a missing
permission and `400` for a malformed request.

Every other route in this module needs the project credentials: without an API
key or a bearer token the client returns an error before anything is sent.
