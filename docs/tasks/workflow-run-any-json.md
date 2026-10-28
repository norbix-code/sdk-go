# Workflow runs with non-object step values (Go SDK)
This file: /Users/djovaisas/Projects/norbix/worktrees/norbix-go/fix/workflow-run-any-json/docs/tasks/workflow-run-any-json.md (branch fix/workflow-run-any-json)

## Goal
A Workflows (Beta) run whose steps hold a string, number or list (a `transform` or `condition` step) can be read with sdk-go, and a regeneration keeps it.
Not in scope: other SDKs (TS, .NET, Python already read these runs), gateway changes, the endpoint modules.

## Plan
1. [done] fix(typegen:go:types): the generator maps C# `object` (export `Object`) to Go `any` instead of `map[string]any`; generator tests updated
2. [done] gen(sdk-go:dtos): regenerated hub and api dtos.go from a private Docker stack (gateway refactoringV2 5f386f6bf); the old generator gives byte-identical files, so the whole diff is the `Object` change
3. [done] test(sdk-go:workflows): a run whose step input is a string and output a number decodes, and StepOutput reads the number
4. [done] release(typegen): generator fix fast-forwarded to main as 53fb4b5 with nbx-ship
5. [doing] release(sdk-go): pull request merged with nbx-ship --wait-release, minor version (v2.42.0 expected)
6. [todo] docs(docs:sdk:go): remove the "Known issue in v2.41.0" hint in the Workflows (Beta) section of go-lang.md, ship with nbx-ship

## Changes
| file (absolute, branch) | what changed | step |
|------|--------------|------|
| https://github.com/codemash-io/sdks-typegen/blob/53fb4b5/languages/go/generate.py (main 53fb4b5) | `Object` → `any` in PRIMITIVES; `Mapper.value` docstring rewritten | 1 |
| https://github.com/codemash-io/sdks-typegen/blob/53fb4b5/tests/go_generator_test.py (main 53fb4b5) | bare `Object` now expects `any`; new test for `Object`, `Object[]` and an inline object literal | 1 |
| https://github.com/codemash-io/sdks-typegen/blob/53fb4b5/languages/go/README.md (main 53fb4b5) | mapping table row for `Object` | 1 |
| /Users/djovaisas/Projects/norbix/worktrees/norbix-go/fix/workflow-run-any-json/norbix/hub/dtos/dtos.go (fix/workflow-run-any-json) | 18 fields `map[string]any` → `any` (and `[]map[string]any` → `[]any`) | 2 |
| /Users/djovaisas/Projects/norbix/worktrees/norbix-go/fix/workflow-run-any-json/norbix/api/dtos/dtos.go (fix/workflow-run-any-json) | 10 fields, same change | 2 |
| /Users/djovaisas/Projects/norbix/worktrees/norbix-go/fix/workflow-run-any-json/norbix/hub/code_workflows_helpers_test.go (fix/workflow-run-any-json) | `TestWaitForWorkflowRunReadsNonObjectStepValues`: string inputs; number, bool and list outputs through WaitForWorkflowRun + StepOutput | 3 |
| /Users/djovaisas/Projects/norbix/worktrees/norbix-go/fix/workflow-run-any-json/norbix/api/dtos/dtos_test.go (fix/workflow-run-any-json) | `TestWorkflowRunReadsNonObjectStepValues`: the api run decodes a string input and number outputs | 3 |

## Findings
fix(sdk-go:workflows): step input/output and run output were typed as JSON objects, so reading a run with a transform or condition step failed — doing
    where: /Users/djovaisas/Projects/norbix/worktrees/norbix-go/fix/workflow-run-any-json/norbix/hub/dtos/dtos.go:6583 (branch fix/workflow-run-any-json)

```csharp
// gateway/src/Isidos.CodeMash.Gateway.Contracts/Code/Workflows/WorkflowRunDto.cs:125-131 (refactoringV2)
    public object? Input { get; init; }          // <-- here: any JSON value (an expression string for transform/condition)
    public bool InputTruncated { get; init; }
    public object? Output { get; init; }         // <-- here: any JSON value (a transform may return a number or a list)
```
```go
// before — norbix/hub/dtos/dtos.go:6583-6585 (main, v2.41.0)
	Input           map[string]any       `json:"input,omitempty"`   // <-- here: refuses "input": "$.amount"
	InputTruncated  bool                 `json:"inputTruncated"`
	Output          map[string]any       `json:"output,omitempty"`  // <-- here: refuses "output": 42
```
```python
# before — typegen languages/go/generate.py:292 (main caee5a5)
    "Object": "map[string]any",    # <-- here: every C# `object` became a JSON-object map
```

decision(typegen:go:types): map every C# `object` to `any`, not only the three workflow fields — the contract says "any JSON value" for all of them, and an allow-list would be a hand-kept list in a generated file again — done
    where: https://github.com/codemash-io/sdks-typegen/blob/53fb4b5/languages/go/generate.py#L292 (main 53fb4b5)
    Fields that change type (`map[string]any` → `any`, `[]map[string]any` → `[]any`): workflow run — `WorkflowRunStepDto.Input` / `Output`, `WorkflowRunDto.Output` (hub, api); taxonomy — `TermDto.Meta`, `TermTreeDto.Meta` (hub, api); database — `FindResponse` / `FindRecordsResponse.Result`, `FindOneResponse` / `FindOneRecordResponse.Result`, `AggregateResponse` / `AggregateRecordsResponse`, `ExecuteAggregateResponse` / `ExecuteRecordsAggregateResponse`, `TestDatabaseAggregateResponse`, `DistinctResponse` / `DistinctRecordValuesResponse.Values` (hub, api); code — `InvokeCodeFunctionResponse.Output`, `InvokeCodeWorkflowResponse.Output`, `InvokeMarketplaceFunctionResponse.Output`, `HttpResult.Response`, `ResolvedResourceEntry.Resolved`, `DiagnosticPackStepResultDto.Result` (hub). Breaking for a Go caller that indexes them (type-assert first); released as a minor (major frozen). No SDK code and no docs snippet reads these typed fields (build, vet, tests green; docs grep empty).

fix(sdk-go:database): distinct values were typed as a list of JSON objects, so distinct values of a text or number field could not be decoded into the typed response — done (same change)
    where: /Users/djovaisas/Projects/norbix/worktrees/norbix-go/fix/workflow-run-any-json/norbix/hub/dtos/dtos.go:8370 (branch fix/workflow-run-any-json)
```go
// before — norbix/hub/dtos/dtos.go:8368-8371 (main, v2.41.0)
type DistinctRecordValuesResponse struct {
	ResponseBase
	Values []map[string]any `json:"values,omitempty"`   // <-- here: ["paid","open"] cannot unmarshal into a map
}
```
```go
// after — same file:8368-8371 (fix/workflow-run-any-json)
type DistinctRecordValuesResponse struct {
	ResponseBase
	Values []any `json:"values,omitempty"`
}
```

decision(sdk-go:workflows): `WorkflowRunDto.Request` and `Trigger` stay `map[string]any` — in C# they are `Dictionary<string, object?>`, always a JSON object — done
    where: /Users/djovaisas/Projects/norbix/gateway/src/Isidos.CodeMash.Gateway.Contracts/Code/Workflows/WorkflowRunDto.cs:84 (branch refactoringV2)

chore(git:typegen): nbx-doctor warning — the typegen main checkout is on `fix/go-references-build-ignore` with an untracked `coverage/` folder; not mine, untouched — todo
    where: /Users/djovaisas/Projects/norbix/sdks/typegen (branch fix/go-references-build-ignore)

## Rejected / moved out
- decision(typegen:go:types): `json.RawMessage` instead of `any` — rejected — `any` matches how the generator already writes dictionary values, and StepOutput decodes from either by a JSON round trip — new ticket: none

## Needs you
(none)

## Open questions
(none)
