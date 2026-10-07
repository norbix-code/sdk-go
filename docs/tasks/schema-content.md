# Schema content — Go SDK (`audit/schema-content`)

Client side of the gateway campaign `audit/schema-content`
(gateway task file: `gateway/docs/tasks/schema-content.md`, items in
`gateway/docs/testing/testing-plan/23-schema-content/items/`). The gateway
branch is not on `refactoringV2` yet, so this branch ships to `main` only
after the campaign lands: the pull request is opened and left open.

## Goal

Make the Go SDK follow the schema-content contract: regenerated DTOs, files
read by id, `expandReferences` on record reads, `arrayFilters` on record
updates, nested documents and the new record error codes documented.

Not in scope: typed request / response methods (the Go modules stay untyped,
`map[string]any` in, `any` out), the AI catalog hint, the cloud builder.

## Plan

1. `done` — regenerate `norbix/{api,hub}/dtos/dtos.go` from the campaign hosts
   (Hub `:49826`, Api `:49877`) with `typegen/languages/go/generate.py`;
   refresh `references/*.dtos.go` from `/types/go` through
   `fix_csharp_generics.py`.
2. `done` — files by id: `API.Files.GetFileById(ctx, filesIntegrationId, id, …)`
   (`GET /files/{filesIntegrationId}/by-id/{id}`) and
   `Hub.Files.GetFileById(ctx, …)` (`GET /files/item/by-id`), one fake-server
   test each.
3. `done` — records: `expandReferences` on `Find` / `FindOne` / `FindOwn`
   (Api) and `FindRecords` / `FindOneRecord` (Hub), `arrayFilters` on
   `UpdateOne` / `UpdateMany` (Api) and `UpdateOneRecord` /
   `UpdateManyRecords` (Hub) — method comments, a typed
   `ExpandedReference{Id, Display}` helper, route tests with the new
   parameters, decode test for an expanded reference, refusal tests for
   `CM-ERRORS-DATABASE-056` / `-053`.
4. `done` — docs: `docs/hub/database.md` (expand references, array filters,
   nested documents, the new schema field DTOs, error codes 039–056), one
   README line for files by id, this task file.
5. `done` — full suite green (`go build ./... && go vet ./... && gofmt -l . &&
   go test ./...`), `nbx-ship --no-merge`, pull request reported.

## Changes

| file | what changed | step |
|---|---|---|
| `norbix/api/dtos/dtos.go`, `norbix/hub/dtos/dtos.go` | regenerated: `ObjectFieldDto`, `ArrayFieldDto`, `JsonFieldDto`, `CurrencyDefaultDto`; `Default` / `Unique` / `DisplayField` / `MultipleOf` / `Minimum` / `Maximum` / `MinItems` / `MaxItems` / `AllowedFileType` / `MaxSizeMb` / `Slug` on the field DTOs; `ExpandReferences` on the find requests; `ArrayFilters` on the update requests; `GetFileByIdRequest` (Api) / `GetFileById` (Hub) + `GetFileByIdResponse` | 1 |
| `references/api.dtos.go`, `references/hub.dtos.go` | refreshed raw `/types/go` export (build-ignored reference) | 1 |
| `norbix/api/files.go`, `norbix/hub/files.go` | `GetFileById` | 2 |
| `norbix/files_test.go` | `TestAPIGetFileById`, `TestHubGetFileById`; header counts 13 / 10 | 2 |
| `norbix/api/database.go`, `norbix/hub/database.go` | method comments for `expandReferences` / `arrayFilters` | 3 |
| `norbix/api/references.go`, `norbix/hub/references.go` | `ExpandedReference` (`{id, display}`) | 3 |
| `norbix/api/database_test.go`, `norbix/hub/database_test.go` | route cases carry `expandReferences` / `arrayFilters`; decode test; refusal cases 053 / 056 | 3 |
| `docs/hub/database.md`, `README.md` | guide + README line | 4 |
| `docs/tasks/schema-content.md` | this file | 4 |

## Findings

- The generator prints `named but never declared, written as any: Func,
  JsonObject` for the Hub export — present on `main` as well, not caused by
  this contract. Left open (typegen).
- `ArrayFilters` is a JSON **string** on the wire (an extended-JSON array),
  same as `filter` / `update`; the generated Go field is `string`. Correct,
  documented.
- The gateway has no DTO for an expanded reference (it is built as a BSON
  document), so the SDK ships a hand-written `ExpandedReference` helper next
  to the modules, not in the generated `dtos` package.

## Rejected / moved out

- Nothing.

## Needs you

- Merge only after the gateway campaign branch `audit/schema-content` is on
  `refactoringV2` (the SDK then targets a contract that `main` serves).

## Open questions

- None.
