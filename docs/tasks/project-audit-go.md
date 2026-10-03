# Project audit — Go SDK (item E1)

## Goal

Make the Go SDK's Project module complete: add the missing project settings,
admin portal, public config, developer MCP endpoint and AI service user
methods, prove every Project-module method with a route test, and document the
module.

Not in scope: AI plans, AI knowledge, AI credits (decided internal); any
gateway or TS SDK change.

## Plan

1. docs(project): task file — done
2. feat(project): Hub methods — `UpdateProjectAdminUrl`, `UpdateProjectLegalDocuments`,
   `UpdateProjectExposeLegal`, `GetAdminPortalStructure`, `AssignAdminPortalServiceUser`,
   AI service users (create / list / delete / rotate key / revoke key), developer MCP
   endpoint (POST / GET / DELETE) — done, `norbix/hub/account.go`
3. feat(project): API `PublicModule` — `GetPublicProjectConfig`, `GetPublicProjectLegal`
   (served by the API host, no sign-in) — done, `norbix/api/public.go`
4. test(project): route tests for every Project-module method (new and existing) and
   every LLM + MCP integration method in `ai.go` — done
5. docs(project): `docs/hub/account.md`, `docs/hub/ai.md`, README links — done
6. check: `go build ./...`, `go vet ./...`, `go test ./...`, gofmt — done (build ok, vet ok, gofmt clean; test: 5 packages ok, 0 failures, 259 passing subtests in `./norbix/...`)
7. push branch and open one pull request to `main` — done, https://github.com/norbix-code/sdk-go/pull/20
8. feat(project): expose brand and expose auth switches (item B3b, gateway routes from item B1) — `UpdateProjectExposeBrand`, `UpdateProjectExposeAuth` + route tests + docs — done (build ok, vet ok, gofmt clean, `go test ./...` ok, 355 passing tests/subtests in `./norbix/...`)

Decisions (taken, not open):

- Scope per new method follows the Go neighbours, not TS (TS says `project` for
  everything): `.../settings/*` → `ScopeAccount` like `UpdateProjectName`;
  `.../admin-portal/*` → `ScopeProject` like `SetAdminPortalEnabled`;
  `/account/ai/service-users*` → `ScopeAccount` (an account-level resource);
  `/account/mcp` → `ScopeProject` (needs a key, works with an AI service user key
  and no account id); `/public/projects/*` → `ScopeUnauthenticated` (gateway says
  public; same as `GetPublicFile`).
- MCP needs the `Mcp-Session-Id` header on GET / DELETE, so `transport.Request`
  gets an optional `Headers` map (additive, no behaviour change for other calls).

## Changes

| file | what changed | plan step # |
|---|---|---|
| `docs/tasks/project-audit-go.md` | this file | 1 |
| `norbix/internal/transport/transport.go` | optional per-request `Headers` | 2 |
| `norbix/hub/account.go` | 13 new methods (5 project settings / admin portal, 5 AI service users, 3 MCP) | 2 |
| `norbix/api/public.go` | new `PublicModule` with 2 methods | 3 |
| `norbix/api/namespace.go` | wire `Public` | 3 |
| `norbix/hub/account_project_test.go` | route tests: 31 Project-module routes (verb, path, auth / project / account headers), account-scope refusal without AccountID, MCP session header + JSON-RPC body | 4 |
| `norbix/hub/ai_integrations_test.go` | route tests: 14 LLM + MCP integration routes (`SetLlmIntegrationAsDefault` was already in `ai_embeddings_test.go`) | 4 |
| `norbix/api/public_test.go` | route tests: 2 public routes, no auth header even with a key | 4 |
| `docs/hub/account.md`, `docs/hub/ai.md` | module docs | 5 |
| `README.md` | "Module guides" links | 5 |
| `norbix/hub/account.go` | `UpdateProjectExposeBrand`, `UpdateProjectExposeAuth` (PATCH `.../settings/brand/expose`, `.../settings/auth/expose`, scope account) | 8 |
| `norbix/hub/account_project_test.go` | 2 route cases; surface size 31 → 33 | 8 |
| `docs/hub/account.md` | settings table rows + "What the public Admin Portal config shows" | 8 |

## Findings

- Scope mismatch: TS SDK sends every Project-module call with scope `project`; Go
  uses `account` for `/account/projects/{projectId}/settings/*`. Only effect: Go
  refuses the call locally when `AccountID` is not set. Left open (`norbix/hub/account.go`).
- TS SDK sends the two public routes with scope `project` (an auth header is sent);
  the gateway marks them public / unauthenticated. Go sends no auth. Left open (TS
  `src/api/public.ts`).
- TS has the public routes on both `src/api/public.ts` and `src/hub/public.ts`; the
  gateway declares them only in `Isidos.CodeMash.Services.Api/Heartbeat/`. Go puts them on
  the API host only. Left open.
- MCP: the transport buffers the whole answer, so `McpOpenStream` (GET, a long SSE
  stream) returns only when the server closes the stream or the timeout hits; and the
  transport does not return response headers, so the `Mcp-Session-Id` that
  `initialize` answers with cannot be read through the SDK. Left open (needs a
  streaming transport call).
- Gateway route `GET /{version}/account/projects/{projectId}/wait-active` has no Go
  method. Left open (not in this item's list).
- `ai.go` has `SetLlmIntegrationAsDefault` (`PUT /ai/integrations/llms/{Id}/default`) — the
  MCP integration family has no default route in the gateway, so no gap.
- `account.go` / `ai.go` carry "Code generated by gen_modules.py. DO NOT EDIT." but the
  generator is not in this repo; methods are hand-added (as earlier waves did). Left open.
- README never linked `docs/hub/push.md`; fixed here (Module guides section).
- types(project): `norbix/hub/dtos/dtos.go` has no `UpdateProjectExposeBrand` / `UpdateProjectExposeAuth` request structs — the file is generated (`DO NOT EDIT`) from a running gateway, so not hand-edited; the methods take `map[string]any` and do not need them — left open, regenerate the types later (README "Regenerate the types").

## Rejected / moved out

- AI plans / knowledge / credits routes — decided internal, not added.

## Needs you

- [ ] Review and merge https://github.com/norbix-code/sdk-go/pull/20 (not merged by the agent).

## Open questions

None.
