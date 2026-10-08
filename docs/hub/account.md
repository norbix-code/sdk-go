# Account and projects — Go

[← Back to project README](../../README.md)

The project endpoints on the Hub, and the `hub.AccountModule` method that
calls each one: create and read projects, project settings (name, brand,
CORS, languages, regions, legal documents), the Admin Portal, project AI
settings, AI service users and the developer MCP endpoint. The two public
project routes live on the API as `api.PublicModule`.

Every route here is covered by a test that asserts the verb and the
fully-resolved path: `norbix/hub/account_project_test.go`,
`norbix/hub/account_ai_test.go` and `norbix/api/public_test.go`.

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
