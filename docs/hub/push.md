# Push — Go

[← Back to project README](../../README.md)

Every Push endpoint the gateway exposes, and the `hub.NotificationsModule`
method that calls it. All 37 are covered by `norbix/hub/push_test.go`, which
asserts the verb and the fully-resolved path of each one.

Each method takes a context, any route ids as plain arguments, an untyped
request body, and a pointer to decode the response into (pass `nil` to discard
it).

```go
client, err := norbix.New(norbix.Options{
    ProjectID: "proj_123",
    APIKey:    "sk_live_...",
})
if err != nil {
    log.Fatal(err)
}

if err := client.Hub.Notifications.EnablePush(ctx, nil, nil); err != nil {
    log.Fatal(err)
}

var templates map[string]any
err = client.Hub.Notifications.GetPushTemplates(ctx, nil, &templates)
```


## Module

| method | verb | path |
|---|---|---|
| `DisablePush(ctx, req, out)` | `GET` | `/notifications/push/disable` |
| `GetPushDisableDependencies(ctx, req, out)` | `GET` | `/notifications/push/disable-dependencies` |
| `EnablePush(ctx, req, out)` | `GET` | `/notifications/push/enable` |
| `GetPushSettings(ctx, req, out)` | `GET` | `/notifications/push/settings` |

## Integrations

| method | verb | path |
|---|---|---|
| `GetPushIntegrations(ctx, req, out)` | `GET` | `/notifications/push/integrations` |
| `SavePushIntegration(ctx, req, out)` | `POST` | `/notifications/push/integrations` |
| `RegisterCodeMashAppPushIntegration(ctx, req, out)` | `POST` | `/notifications/push/integrations/app/request` |
| `ConfirmPushIntegrationHumanDelivery(ctx, req, out)` | `POST` | `/notifications/push/integrations/confirm-human-delivery` |
| `TestPushIntegration(ctx, req, out)` | `POST` | `/notifications/push/integrations/test` |
| `DeletePushIntegration(ctx, id, req, out)` | `DELETE` | `/notifications/push/integrations/{Id}` |
| `SetPushIntegrationAsDefault(ctx, id, req, out)` | `PUT` | `/notifications/push/integrations/{Id}/default` |
| `DisablePushIntegration(ctx, id, req, out)` | `PUT` | `/notifications/push/integrations/{Id}/disable` |
| `EnablePushIntegration(ctx, id, req, out)` | `PUT` | `/notifications/push/integrations/{Id}/enable` |
| `GetPushIntegration(ctx, id, req, out)` | `GET` | `/notifications/push/integrations/{id}` |

## Templates

| method | verb | path |
|---|---|---|
| `GetPushTemplates(ctx, req, out)` | `GET` | `/notifications/push/templates` |
| `CreatePushTemplate(ctx, req, out)` | `POST` | `/notifications/push/templates` |
| `UpdatePushTemplate(ctx, req, out)` | `PUT` | `/notifications/push/templates` |
| `RenderPush(ctx, req, out)` | `POST` | `/notifications/push/templates/render` |
| `DeletePushTemplate(ctx, id, req, out)` | `DELETE` | `/notifications/push/templates/{Id}` |
| `ArchivePushTemplate(ctx, id, req, out)` | `PUT` | `/notifications/push/templates/{Id}/archive` |
| `ClonePushTemplate(ctx, id, req, out)` | `POST` | `/notifications/push/templates/{Id}/clone` |
| `UnArchivePushTemplate(ctx, id, req, out)` | `PUT` | `/notifications/push/templates/{Id}/unarchive` |
| `GetPushTemplate(ctx, id, req, out)` | `GET` | `/notifications/push/templates/{id}` |
| `GetPushMessageContentTokens(ctx, id, req, out)` | `GET` | `/notifications/push/templates/{id}/tokens` |

## Campaigns

| method | verb | path |
|---|---|---|
| `GetPushCampaigns(ctx, req, out)` | `GET` | `/notifications/push/campaigns` |
| `CreatePushCampaign(ctx, req, out)` | `POST` | `/notifications/push/campaigns` |
| `DeletePushCampaign(ctx, id, req, out)` | `DELETE` | `/notifications/push/campaigns/{Id}` |
| `StopPushCampaign(ctx, id, req, out)` | `POST` | `/notifications/push/campaigns/{Id}/stop` |
| `GetPushCampaignMessages(ctx, campaignId, req, out)` | `GET` | `/notifications/push/campaigns/{campaignId}/messages` |
| `GetPushCampaignMessage(ctx, campaignId, id, req, out)` | `GET` | `/notifications/push/campaigns/{campaignId}/messages/{id}` |
| `GetPushCampaign(ctx, id, req, out)` | `GET` | `/notifications/push/campaigns/{id}` |
| `GetPushCampaignBatches(ctx, id, req, out)` | `GET` | `/notifications/push/campaigns/{id}/batches` |
| `GetPushCampaignBatchNotifications(ctx, id, batchId, req, out)` | `GET` | `/notifications/push/campaigns/{id}/batches/{batchId}` |
| `GetPushCampaignBatchNotification(ctx, id, batchId, notificationId, req, out)` | `GET` | `/notifications/push/campaigns/{id}/batches/{batchId}/{notificationId}` |
| `GetPushCampaignStatistics(ctx, id, req, out)` | `GET` | `/notifications/push/campaigns/{id}/stats` |
| `PreviewPushNotification(ctx, req, out)` | `GET` | `/notifications/push/preview` |

## Devices

| method | verb | path |
|---|---|---|
| `RegisterDevice(ctx, req, out)` | `POST` | `/notifications/push/devices` |
| `GetPushDevices(ctx, req, out)` | `GET` | `/notifications/push/devices` |
| `GetPushDevice(ctx, id, req, out)` | `GET` | `/notifications/push/devices/{id}` |

## Choosing who a campaign goes to

`CreatePushCampaign` takes the audience in the request body. The server picks
the shape from the `source` field, so send `source` plus that audience's own
fields:

| audience | `source` | own fields |
|---|---|---|
| everyone in the project | `allUsers` | `rolesNames`, `userTags` (both optional filters) |
| a named list of project users | `specifiedUsers` | `userRecipients` |
| a named list of account users | `accountUsers` | `userRecipients` |
| rows of a database collection | `collection` | `schemaName`, `fields` (the record fields that hold the recipient), `fieldType` (`User` or `Email`), optional `roleNames`, `languages` |
| raw device tokens | `devices` | `devices`: a list of `{ token, deliveryFamily }`, `deliveryFamily` one of `Ios`, `Android`, `Chrome`, `Safari`, `Expo` |

Every target also takes `templateId` (required) and the optional `integrationId`,
`language`, `notes`, `campaignTime` (Unix seconds) and `mappedTokens`. Note the
spelling: `rolesNames` on `allUsers`, but `roleNames` on `collection` — the
gateway names them differently.

```go
req := map[string]any{
    "campaign": map[string]any{
        "source":     "allUsers",
        "templateId": "tpl_123",
        "userTags":   []string{"beta"},
    },
}
err := client.Hub.Notifications.CreatePushCampaign(ctx, req, nil)
```

Send `source` as the name, not a number — the server reads it as a string.

## Choosing a push provider

`SavePushIntegration` works the same way, with a `provider` field:

| provider | `provider` value | own fields |
|---|---|---|
| Fake (sandbox, never sends) | `Fake` | none |
| Android / Firebase | `AndroidFirebase` | `projectId`, `clientEmail`, `serviceAccountJson` |
| Apple APNs | `AppleApns` | `teamId`, `appBundleId`, `keyId`, `privateKey`, `isProduction` |
| Chrome extension | `ChromePush` | `extensionId` (a GUID), `vapidPublicKey`, `vapidPrivateKey`, optional `subject` |
| Chrome web | `ChromeWeb` | `vapidPublicKey`, `vapidPrivateKey`, optional `subject` |
| Edge web | `EdgeWeb` | `vapidPublicKey`, `vapidPrivateKey`, optional `subject` |
| Firefox web | `FirefoxWeb` | `vapidPublicKey`, `vapidPrivateKey`, optional `subject` |
| Safari | `SafariPush` | `websitePushId`, `certificateP12Base64`, `certificatePassword` |

Every provider also takes `integrationName` and `isEnabled`; send `integrationId`
to update an existing one. The Chrome extension value is `ChromePush`. The
generated types also list `CodeMashChromePlugin` and other `CodeMash*` values —
the server rejects those here with "Unsupported provider".

Use `Fake` in tests and local development. It accepts a send and contacts no
push service, so nothing reaches a real device.

## Registering a device

A device is registered for one user. Send the device under `pushDeviceDto`
(`deviceOs` and `token` are required; `deviceId`, `brand`, `manufacturer`,
`modelName`, `deviceName`, `deviceType` are optional) plus `userId`:

```go
req := map[string]any{
    "userId": "user_123",
    "pushDeviceDto": map[string]any{"deviceOs": "iOS", "token": "<device token>"},
}
err := client.Hub.Notifications.RegisterDevice(ctx, req, nil)
```

## Listing registered devices

`GetPushDevices` returns the devices registered in the project, each with the
user it belongs to. Narrow it with `userId`, `deviceKey` (the provider token)
or `platform` (`ios`, `android`, `chrome`, `safari`, `expo`); a word outside
that list is refused rather than answered with an empty page.

```go
var devices map[string]any
err := client.Hub.Notifications.GetPushDevices(ctx, map[string]any{
    "platform": "ios",
}, &devices)
```

Devices are stored inside their user, so a page is a page of **users** and
carries every matching device those users hold. Follow `hasMore` rather than
stopping at the first short page.

`GetPushDevice` takes one device id and answers with the device and its owner:

```go
var device map[string]any
err := client.Hub.Notifications.GetPushDevice(ctx, "pnd_123", nil, &device)
```

## Known gaps

| what | why |
|---|---|
| `CheckIntegrationAvailability` and `TestCodeMashIosAppIntegration` | these two methods point at routes the gateway has commented out, so they are not callable. They are left in place until the gateway decides whether to finish or drop them. |
| `GetPushCampaignMessage` | the gateway route declares an `{id}` token that no request field matches, so the endpoint is unreliable until that is fixed. |

