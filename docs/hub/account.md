# Account and projects — Go

[← Back to project README](../../README.md)

The project endpoints on the Hub, and the `hub.AccountModule` method that
calls each one: create and read projects, project settings (name, brand,
CORS, languages, regions, legal documents), the Admin Portal, project AI
settings, AI service users, the subscription bundles, discount codes, trial
limits and Flex clusters, and
the developer MCP endpoint. The two public
project routes live on the API as `api.PublicModule`.

Every route here is covered by a test that asserts the verb and the
fully-resolved path: `norbix/hub/account_project_test.go`,
`norbix/hub/account_ai_test.go` and `norbix/api/public_test.go`. Every
project settings route (matrix section 21) is also in
`norbix/hub/account_settings_routes_test.go`, which checks the host, the query
string and the body too.

Each method takes a context, any route ids as plain arguments, an untyped
request body, and a pointer to decode the response into (pass `nil` to discard
it).

**Scope.** Every method on this page is marked *project*: it needs only a key
or a bearer token, no `AccountID`. The gateway takes the account from the
signed-in session, or from the project id in the path, and never reads the
`norbix-account-id` header. Test: `norbix/hub/account_token_only_test.go`.

```go
client, err := norbix.New(norbix.Options{
    ProjectID: "proj_123",
    APIKey:    "sk_live_...",
})
if err != nil {
    log.Fatal(err)
}

var project map[string]any
err = client.Hub.Account.GetProject(ctx, "proj_123", nil, &project)
```

## Projects

| method | verb | path | scope |
|---|---|---|---|
| `GetProjects(ctx, req, out)` | `GET` | `/account/projects` | project |
| `CreateProject(ctx, req, out)` | `POST` | `/account/projects` | project |
| `GetProject(ctx, projectId, req, out)` | `GET` | `/account/projects/{projectId}` | project |
| `DeleteProject(ctx, projectId, req, out)` | `DELETE` | `/account/projects/{projectId}` | project |
| `EnableProject(ctx, projectId, req, out)` | `PATCH` | `/account/projects/{projectId}/enable` | project |
| `DisableProject(ctx, projectId, req, out)` | `PATCH` | `/account/projects/{projectId}/disable` | project |
| `GetProjectTokens(ctx, projectId, req, out)` | `GET` | `/account/projects/{projectId}/tokens` | project |
| `GetProjectHealth(ctx, projectId, req, out)` | `GET` | `/account/projects/{projectId}/health` | project |
| `WaitForProjectActive(ctx, projectId, req, out)` | `GET` | `/account/projects/{projectId}/wait-active` | project |

A new project provisions its database in the background. Call
`WaitForProjectActive` once after `CreateProject`, before you build inside the
project; a timeout is not an error, read `isActive` in the answer.
`GetProjectHealth` lists the project's problems (each with a Cloud fix path)
and its user / e-mail / API call numbers; pass `"refresh": true` to compute it
now and `"periodDays"` (7, 30 or 90) for another window.

```go
var ready map[string]any
err := client.Hub.Account.WaitForProjectActive(ctx, "proj_123", map[string]any{"timeoutSeconds": 60}, &ready)
```

## Environments

`Hub.Environments` lists, creates and deletes the environments:
`List(ctx, req, out)` (`GET /account/projects/environments`),
`Create(ctx, req, out)` (`POST`, send `environmentName` and the new
environment's database `integration`) and
`Delete(ctx, environmentName, req, out)` (`DELETE …/{environmentName}`).
The ladder, promotion and admin addresses are on `Hub.Account`:

| method | verb | path | scope |
|---|---|---|---|
| `SetEnvironmentRank(ctx, environmentName, req, out)` | `PATCH` | `/account/projects/environments/{environmentName}/rank` | project |
| `PromoteEnvironment(ctx, req, out)` | `POST` | `/account/projects/environments/promote` | project |
| `RollbackPromotion(ctx, req, out)` | `POST` | `/account/projects/environments/promote/rollback` | project |
| `RepairEnvironmentBuiltIns(ctx, req, out)` | `POST` | `/account/projects/environments/built-ins/repair` | project |
| `GetEnvironmentAdminUrl(ctx, environmentName, req, out)` | `GET` | `/account/projects/environments/{environmentName}/admin-url` | project |
| `UpdateEnvironmentAdminUrl(ctx, environmentName, req, out)` | `PATCH` | `/account/projects/environments/{environmentName}/admin-url` | project |

Preview a promotion first (`"dryRun": true`), then run it. The answer's
`fromVersion` is the anchor `RollbackPromotion` takes:

```go
var plan map[string]any
err := client.Hub.Account.PromoteEnvironment(ctx, map[string]any{
    "sourceEnv": "TEST", "targetEnv": "PROD", "dryRun": true,
}, &plan)
```

## Project settings

All are `PATCH /account/projects/{projectId}/settings/<segment>`, scope project.

| method | segment |
|---|---|
| `UpdateProjectName` | `name` |
| `UpdateProjectDescription` | `description` |
| `UpdateProjectLogo` | `logo` |
| `UpdateProjectIcon` | `icon` |
| `UpdateProjectMainColor` | `main-color` |
| `UpdateProjectAccentColor` | `accent-color` |
| `UpdateProjectUrl` | `url` |
| `UpdateProjectLanguages` | `languages` |
| `UpdateProjectDefaultLanguage` | `default-language` |
| `UpdateProjectRegions` | `regions` |
| `UpdateProjectAllowedOrigins` | `origins` (CORS) |
| `UpdateProjectAdminUrl` | `admin-url` |
| `UpdateProjectLegalDocuments` | `legal` |
| `UpdateProjectExposeLegal` | `legal/expose` |
| `UpdateProjectExposeBrand` | `brand/expose` |
| `UpdateProjectExposeAuth` | `auth/expose` |

### Check template languages

`CheckProjectLanguages(ctx, projectId, req, out)` — `POST
/account/projects/{projectId}/settings/languages/check`, scope project. It
lists the Email, Push and SMS templates that miss a project language. Send the
proposed `defaultLanguage` and/or `languages` before you save them with
`UpdateProjectLanguages`; send neither to check the current settings.

```go
var gaps dtos.CheckProjectLanguagesResponse
err := client.Hub.Account.CheckProjectLanguages(ctx, "proj_123", map[string]any{
    "languages": []string{"en", "de"},
}, &gaps)
for _, g := range gaps.Templates {
    fmt.Println(g.Module, g.TemplateName, g.MissingLanguages)
}
```

### CORS

The browser origins allowed to call the project:

```go
err := client.Hub.Account.UpdateProjectAllowedOrigins(ctx, "proj_123", map[string]any{
    "origins": []string{"https://app.example.com"},
}, nil)
```

### Legal documents

Save the Terms and Privacy Policy as Markdown, then make them public. Once
exposed, anyone can read them without signing in:

```go
_ = client.Hub.Account.UpdateProjectLegalDocuments(ctx, "proj_123", map[string]any{
    "termsMarkdown":   "# Terms\n...",
    "privacyMarkdown": "# Privacy\n...",
}, nil)
_ = client.Hub.Account.UpdateProjectExposeLegal(ctx, "proj_123", map[string]any{"exposed": true}, nil)

var terms map[string]any
_ = client.API.Public.GetPublicProjectLegal(ctx, "proj_123", "terms", &terms)
```

### What the public Admin Portal config shows

`api.PublicModule.GetPublicProjectConfig` (no sign-in) returns the brand by
default and hides the sign-in methods and password policy. Two switches change
that:

```go
// Hide the brand (name, colors, logo, icon) from the public config.
_ = client.Hub.Account.UpdateProjectExposeBrand(ctx, "proj_123", map[string]any{"exposed": false}, nil)
// Show the sign-in methods (email / phone / username) and the password policy.
_ = client.Hub.Account.UpdateProjectExposeAuth(ctx, "proj_123", map[string]any{"exposed": true}, nil)
```

## Admin Portal

| method | verb | path | scope |
|---|---|---|---|
| `SetAdminPortalEnabled(ctx, projectId, req, out)` | `PUT` | `/account/projects/{projectId}/admin-portal/enabled` | project |
| `GetAdminPortalStructure(ctx, projectId, req, out)` | `GET` | `/account/projects/{projectId}/admin-portal/structure` | project |
| `AssignAdminPortalServiceUser(ctx, projectId, req, out)` | `PUT` | `/account/projects/{projectId}/settings/admin-portal/service-user` | project |
| `UpdateProjectAdminUrl(ctx, projectId, req, out)` | `PATCH` | `/account/projects/{projectId}/settings/admin-url` | project |

### Public project config (API, no sign-in)

`api.PublicModule` sends no `Authorization` header, even when the client has
a key — these links must work for anyone. An unknown project answers an empty
document, not an error.

| method | verb | path |
|---|---|---|
| `GetPublicProjectConfig(ctx, projectId, out)` | `GET` | `/public/projects/{ProjectId}/config` |
| `GetPublicProjectLegal(ctx, projectId, kind, out)` | `GET` | `/public/projects/{ProjectId}/legal/{Kind}` (`terms` or `privacy`) |
| `GetPublicProjectBrandAsset(ctx, projectId, kind, out)` | `GET` | `/public/projects/{ProjectId}/brand/{Kind}` (`logo` or `icon`) |

The brand asset answers the image itself: pass a `*[]byte` as `out`. It is
there only while the project exposes its brand; every miss is a plain 404.

```go
var cfg map[string]any
err := client.API.Public.GetPublicProjectConfig(ctx, "proj_123", &cfg)
```

## Project AI settings

| method | verb | path | scope |
|---|---|---|---|
| `GetProjectAiSettings(ctx, projectId, req, out)` | `GET` | `/account/projects/{projectId}/ai/settings` | project |
| `UpdateProjectAiSettings(ctx, projectId, req, out)` | `PUT` | `/account/projects/{projectId}/ai/settings` | project |
| `CreateProjectAiAssistant(ctx, projectId, req, out)` | `POST` | `/account/projects/{projectId}/ai/assistants` | project |
| `UpdateProjectAiAssistant(ctx, projectId, assistantId, req, out)` | `PUT` | `/account/projects/{projectId}/ai/assistants/{assistantId}` | project |
| `DeleteProjectAiAssistant(ctx, projectId, assistantId, req, out)` | `DELETE` | `/account/projects/{projectId}/ai/assistants/{assistantId}` | project |
| `GetProjectAiUsage(ctx, projectId, req, out)` | `GET` | `/account/projects/{projectId}/ai/usage` | project |
| `UpdateProjectAiKnowledge(ctx, projectId, req, out)` | `PUT` | `/account/projects/{projectId}/ai/knowledge` | project |
| `GetProjectAiPlans(ctx, projectId, req, out)` | `GET` | `/account/projects/{projectId}/ai/plans` | project |
| `UpdateProjectAiPlans(ctx, projectId, req, out)` | `PUT` | `/account/projects/{projectId}/ai/plans` | project |
| `UpdateProjectAiPlanAssignments(ctx, projectId, req, out)` | `PUT` | `/account/projects/{projectId}/ai/plans/assignments` | project |
| `GetProjectAiUserPlans(ctx, projectId, req, out)` | `GET` | `/account/projects/{projectId}/ai/plans/users` | project |
| `CreateAiCreditPackCheckout(ctx, projectId, req, out)` | `POST` | `/account/projects/{projectId}/ai/credits/checkout` | project |

`UpdateProjectAiPlans` replaces the whole list (`"plans"`): a plan left out is
removed. `UpdateProjectAiPlanAssignments` replaces `"roles"` and
`"defaultPlanId"` together. Read `GetProjectAiPlans` first and send the list
back with your change. `CreateAiCreditPackCheckout` returns a Stripe checkout
address for a credit pack (`"pack"`: `pack-10`, `pack-50` or `pack-200`);
nothing is charged until the payer completes the checkout.

```go
var settings map[string]any
err := client.Hub.Account.GetProjectAiSettings(ctx, "proj_123", nil, &settings)
```

## AI service users

A service user is a scoped identity for an AI agent (Claude Code, Cursor, …).
Its API key is shown **once** — in the answer of `CreateAiServiceUser` and
`RotateAiServiceUserKey` (pass `"revokeKeyId"` to retire the old key in the
same call). Later reads show only a key id and a hint.

| method | verb | path | scope |
|---|---|---|---|
| `CreateAiServiceUser(ctx, req, out)` | `POST` | `/account/ai/service-users` | project |
| `ListAiServiceUsers(ctx, req, out)` | `GET` | `/account/ai/service-users` | project |
| `DeleteAiServiceUser(ctx, id, req, out)` | `DELETE` | `/account/ai/service-users/{Id}` | project |
| `RotateAiServiceUserKey(ctx, id, req, out)` | `POST` | `/account/ai/service-users/{Id}/keys` | project |
| `RevokeAiServiceUserKey(ctx, id, keyId, req, out)` | `DELETE` | `/account/ai/service-users/{Id}/keys/{KeyId}` | project |

```go
var created map[string]any
err := client.Hub.Account.CreateAiServiceUser(ctx, map[string]any{
    "name": "Claude Code on my laptop",
    "scope": map[string]any{
        "reach": "project", "projectId": "proj_123",
        "rights": "read", "envs": []string{"TEST"},
    },
}, &created)
// Store created's key now; it is not shown again.
```

## Sign-up, invitations, regions and account verification (no token)

These four routes are public on the gateway, so the methods send **no**
`Authorization` header and need no key, no bearer token and no `AccountID`
(scope *unauthenticated*). `VerifyAccount` takes the account id once, in the
request. Test: `TestPublicAccountRoutesWorkWithNoTokenAndNoAccountID` in
`norbix/hub/account_token_only_test.go`.

| method | verb | path | scope |
|---|---|---|---|
| `CreateAccount(ctx, req, out)` | `POST` | `/account` | unauthenticated |
| `CreateTeamMemberFromInvitation(ctx, req, out)` | `POST` | `/account/team/member` | unauthenticated |
| `GetAccountRegions(ctx, req, out)` (also `Hub.Regions.List`) | `GET` | `/account/regions` | unauthenticated |
| `VerifyAccount(ctx, req, out)` | `GET` | `/account/verify` | unauthenticated |

```go
hub, err := norbix.NewHub(norbix.Options{ProjectID: "proj_123"}) // no key, no token
if err != nil {
    log.Fatal(err)
}
err = hub.Account.VerifyAccount(ctx, map[string]any{
    "accountId": "acct_123", // from the verification link
    "token":     "tok_...",
}, nil)
```

## Your account user and the team

The signed-in account user's own profile and phone, and the account team. The
gateway takes the account from the signed-in session, so these need only a key
or a bearer token — no `AccountID`. Tests: `norbix/hub/account_me_test.go`.

| method | verb | path | scope |
|---|---|---|---|
| `GetMyAccountUserProfile(ctx, req, out)` | `GET` | `/account/me` | project |
| `UpdateMyAccountUserPhone(ctx, req, out)` | `PUT` | `/account/me/phone` | project |
| `GetAccountCollaborators(ctx, req, out)` | `GET` | `/account/collaborators` | project |

The phone (E.164 format, `+` and the country code; an empty value clears it) is
the number "Account users" SMS campaigns send to. Members without a phone are
skipped.

```go
err := client.Hub.Account.UpdateMyAccountUserPhone(ctx, map[string]any{
    "phone": "+37060000000",
}, nil)

var me dtos.GetMyAccountUserProfileResponse
err = client.Hub.Account.GetMyAccountUserProfile(ctx, nil, &me)
fmt.Println(me.Item.Email, me.Item.GeneralInfo.Phone)
```

The team list pages with flat fields — `pageSize` (default 20),
`startingAfter`, `endingBefore` — and `projectId` (without
`includeAccountOwner`) narrows it to one project's collaborators:

```go
var team dtos.GetAccountCollaboratorsResponse
err := client.Hub.Account.GetAccountCollaborators(ctx, map[string]any{
    "projectId": "proj_123",
    "pageSize":  50,
}, &team)
```

## Subscription, bundles and Flex clusters

The Managed Service is sold as three bundles: `Starter`, `Growth` and `Scale`
(`dtos.ManagedServiceBundle`), paid `Monthly` or `Annual`
(`dtos.SubscriptionBillingPeriod`). Each bundle has a limit of enabled
projects. Tests: `norbix/hub/account_bundles_flex_test.go`.

| method | verb | path | scope |
|---|---|---|---|
| `GetAccountStatus(ctx, req, out)` | `GET` | `/account/status` | project |
| `GetManagedServiceBundles(ctx, req, out)` | `GET` | `/account/stripe/bundles` | project |
| `CreateStripeCheckoutSession(ctx, req, out)` | `POST` | `/account/stripe/create-checkout-session` | project |
| `ChangeManagedServiceBundle(ctx, req, out)` | `POST` | `/account/stripe/change-bundle` | project |
| `CheckBundleDiscount(ctx, req, out)` | `POST` | `/account/stripe/discount/check` | project |
| `ApplyBundleDiscount(ctx, req, out)` | `POST` | `/account/stripe/discount` | project |
| `GetTrialAvailability(ctx, req, out)` | `GET` | `/account/trial-availability` | project |
| `JoinTrialWaitlist(ctx, req, out)` | `POST` | `/account/trial-waitlist` | project |
| `GetStripeBillingPortalUrl(ctx, req, out)` | `POST` | `/account/stripe/get-portal-url` | project |
| `SetExtraRegions(ctx, req, out)` | `POST` | `/account/stripe/extra-regions` | project |
| `GetAccountUsageBilling(ctx, req, out)` | `GET` | `/account/usage-billing` | project |
| `GetFlexClusters(ctx, req, out)` | `GET` | `/account/flex-clusters` | project |
| `UpgradeFlexCluster(ctx, flexClusterId, req, out)` | `POST` | `/account/flex-clusters/{flexClusterId}/upgrade` | project |

- `GetManagedServiceBundles` lists the three bundles with the project limit
  (`ProjectCap`) and the `Monthly` and `Annual` price. `Amount` is in minor
  units (cents), `Currency` is lower case (`"usd"`).
- `CreateStripeCheckoutSession` for the Managed Service needs `bundle` and
  `billingPeriod`. It no longer takes `projectCap`.
- `ChangeManagedServiceBundle` moves the account to another bundle or billing
  period. The gateway refuses with `CM-ERRORS-ACCOUNT-023` when the account
  has more enabled projects than the new bundle allows.
- `SetExtraRegions` sets how many extra regions (`extraRegions`, 0 removes
  them) the Starter or Growth plan carries; Stripe changes the subscription
  pro rata. The change reaches the account a moment later, through Stripe's
  webhook: read `GetAccountStatus` again after the answer. Discount codes:
  see below.
- `GetAccountUsageBilling` returns the usage-based billing
  (`dtos.UsageBillingDto`): the Norbix-managed Atlas charges per period and
  cluster, and the workflow run counts.
- Every billing route of this table is also in
  `norbix/hub/account_settings_routes_test.go` (host, verb, path, query, body).
- `GetAccountStatus` now also returns `Bundle`, `BillingPeriod`, `IsTrial`,
  `RenewsOn` and `EnabledProjects` (`dtos.AccountStatusDto`).

```go
var bundles dtos.GetManagedServiceBundlesResponse
err := client.Hub.Account.GetManagedServiceBundles(ctx, nil, &bundles)

err = client.Hub.Account.ChangeManagedServiceBundle(ctx, map[string]any{
    "bundle":        dtos.ManagedServiceBundleGrowth,
    "billingPeriod": dtos.SubscriptionBillingPeriodAnnual,
}, nil)
```

### Discount codes

A discount code takes a share (`PercentOff`) or a fixed sum (`AmountOff`, minor
units, with `Currency`) off the bundle price, for `Months` months (nil means
for as long as the subscription runs) or until `EndsAt`
(`dtos.BundleDiscountDto`). Tests: `norbix/hub/account_bundle_discounts_test.go`.

- `CheckBundleDiscount` checks a code and returns what it gives
  (`dtos.CheckBundleDiscountResponse`); it changes nothing. Send `code` and,
  when you know them, `bundle` and `billingPeriod`.
- `ApplyBundleDiscount` puts a code on the running bundle subscription, so it
  changes what the account pays. Send `code`. With no subscription the gateway
  answers `CM-ERRORS-PROJECTS-039`.
- `CreateStripeCheckoutSession` takes the code as `discountCode`.
- `GetAccountStatus` returns the discount in force as `Discount`.

A refused code fails with HTTP 400 and `CM-ERRORS-ACCOUNT-027`
(`*errors.ValidationError`). The reason is in the error's context:
`not_found`, `not_started`, `expired`, `not_for_bundle`, `monthly_only`,
`wrong_billing_period`, `already_has_discount` or `already_used`.

```go
var check dtos.CheckBundleDiscountResponse
err := client.Hub.Account.CheckBundleDiscount(ctx, map[string]any{
    "code":          "LAUNCH20",
    "bundle":        dtos.ManagedServiceBundleGrowth,
    "billingPeriod": dtos.SubscriptionBillingPeriodMonthly,
}, &check)

err = client.Hub.Account.ApplyBundleDiscount(ctx, map[string]any{"code": "LAUNCH20"}, nil)
var nerr *nberrors.Error // github.com/norbix-code/sdk-go/v2/norbix/errors
if errors.As(err, &nerr) && nerr.Code == "CM-ERRORS-ACCOUNT-027" && len(nerr.Errors) > 0 {
    reason, _ := nerr.Errors[0].Context["reason"].(string) // e.g. "already_used"
    _ = reason
}
```

### Trial limits

A free Managed Service trial can be limited per region. Tests:
`norbix/hub/account_trial_limits_test.go`.

- `GetTrialAvailability` asks whether a free trial can start in a region now,
  before opening a checkout (`dtos.TrialAvailabilityResponse`). Send `region`
  (it goes in the query string).
  - `LimitsOn` false: this installation has no trial limits; start the trial
    as before.
  - `TrialsOpen` false: the region's trials are full. Offer buying now without
    a trial (checkout with `withoutTrial: true`) or joining the waiting list
    when `WaitlistOpen` is true.
  - `OnWaitlist`: the owner email is already on the list.
  - `Refusal`: why this account's trial was refused after checkout
    (`CardAlreadyHadTrial`), or empty.
- `JoinTrialWaitlist` puts the account owner's email on the trial waiting list
  for a region (`dtos.TrialWaitlistResponse`). Send `region`, `consent` (must
  be true) and `consentVersion` (names the consent text shown). When a slot
  frees, Norbix emails an invitation link valid a few days. `Status` is
  `added`, `alreadyListed` or `alreadyCustomer` (an email that already had a
  trial or pays is never added); `Message` can be shown as is. Errors:
  `CM-ERRORS-ACCOUNT-044` the list is switched off, `CM-ERRORS-ACCOUNT-043`
  unknown region, a validation error when `consent` is not true.
- `CreateStripeCheckoutSession` takes `region`, `withoutTrial` and
  `trialInvite`. A trial checkout in a full region fails with HTTP 400 and
  `CM-ERRORS-ACCOUNT-040`. The invitation token from the waiting list goes in
  `trialInvite` and starts the trial once even when the region is full; an
  expired invitation or someone else's fails with `CM-ERRORS-ACCOUNT-042`.

The website's own waiting-list form posts to `/account/trial-waitlist/intake`
with a shared key header. That route is for the Norbix website server only, so
the SDK has no method for it.

```go
var avail dtos.TrialAvailabilityResponse
err := client.Hub.Account.GetTrialAvailability(ctx, map[string]any{
    "region": "nb-eu-germany",
}, &avail)
if avail.LimitsOn && !avail.TrialsOpen && avail.WaitlistOpen && !avail.OnWaitlist {
    var joined dtos.TrialWaitlistResponse
    err = client.Hub.Account.JoinTrialWaitlist(ctx, map[string]any{
        "region":         "nb-eu-germany",
        "consent":        true,
        "consentVersion": "2026-11",
    }, &joined)
    // joined.Status: "added", "alreadyListed" or "alreadyCustomer"
}
```

### Flex clusters

`GetFlexClusters` lists the account's Norbix-managed Atlas clusters
(`dtos.FlexClusterDto`): tier (`FLEX`, `M10` … `M40`), status (`Empty`,
`Provisioning`, `Active`, `Failed`), the databases on the cluster, its usage
and the limits it is close to (`NearLimits`). `CanCreateNew` says whether the
account can add a cluster; when it cannot, `CannotCreateReason` is
`TrialOneCluster` or `ClusterLimitReached`. `UpgradeFlexCluster` queues the
upgrade of one cluster by one step on the ladder FLEX → M10 → M20 → M30 → M40,
within the plan (trial: FLEX only · Starter: up to M20 · Growth: up to M30 ·
Scale: up to M40). Each cluster carries `NextTierCode` (the next step, empty when
there is none) and `TopTierCode` (the plan's highest tier); send `NextTierCode`
as `tierCode`.

```go
var clusters dtos.GetFlexClustersResponse
err := client.Hub.Account.GetFlexClusters(ctx, nil, &clusters)

c := clusters.Clusters[0]
if c.NextTierCode != "" {
    err = client.Hub.Account.UpgradeFlexCluster(ctx, c.Id, map[string]any{
        "tierCode": c.NextTierCode,
    }, nil)
}
```

To put a new managed database on a cluster that already exists, send its id
as `flexClusterId` when you save the integration; see
[Database](database.md).

## Developer MCP endpoint

`/account/mcp` is the Hub's Model Context Protocol server. MCP clients
(Claude, Cursor) talk to it on their own; these methods are for tests and
scripts that speak JSON-RPC directly. The session travels in the
`Mcp-Session-Id` header (`hub.McpSessionHeader`).

| method | verb | path | scope |
|---|---|---|---|
| `Mcp(ctx, sessionId, msg, out)` | `POST` | `/account/mcp` | project |
| `McpOpenStream(ctx, sessionId, out)` | `GET` | `/account/mcp` | project |
| `McpEndSession(ctx, sessionId, out)` | `DELETE` | `/account/mcp` | project |

```go
var tools map[string]any
err := client.Hub.Account.Mcp(ctx, sessionID, map[string]any{
    "jsonrpc": "2.0", "id": 1, "method": "tools/list",
}, &tools)
```

Limits: the SDK reads the whole answer before it returns, and it does not
return response headers. So the `Mcp-Session-Id` that `initialize` answers
with cannot be read through the SDK, and `McpOpenStream` returns only when the
server closes the stream or the client timeout is reached. Pass a `*[]byte` as
`out` when the answer may be an SSE stream rather than JSON.
