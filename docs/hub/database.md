# Database — Go

[← Back to project README](../../README.md)

The Database module covers schemas (drafts, versions, settings, list
settings, embed, schema bundles), records in collections, taxonomies and
their terms (including the trees and the merged tree), schema triggers,
database integrations, saved aggregates and turning the module on or off.
Imports are not part of this guide.

There are two clients:

- **Hub** (`client.Hub.Database`, `hub.DatabaseModule`) — the project
  owner's side: everything, including record work done from the dashboard.
- **Api** (`client.API.Database`, `api.DatabaseModule`) — the end user's
  side: read schemas and taxonomies, work with records, and `FindOwn` for the
  records the signed-in user is responsible for.

Every method is covered by `norbix/hub/database_test.go` (67 routes) and
`norbix/api/database_test.go` (22 routes). Each test asserts the verb, the
fully-resolved path, the query string, the JSON body and the key + project
headers.

Each method takes a context, any route id as a plain argument, an untyped
request (`map[string]any` — sent as the query string for `GET` / `DELETE`, as
the JSON body otherwise), and a pointer to decode the response into (pass
`nil` to discard it). Every path is prefixed with the version (`/v2/...`).

**Records are JSON strings.** The gateway takes a record, a filter, an update
and a pipeline as JSON text, not as nested objects: `"document"`,
`"documents"`, `"filter"`, `"update"`, `"replacement"` and `"pipeline"` are
strings. Pass `"databaseIntegrationId"` to work on a database other than the
project default.

## Examples

### Add a record and read it back (Hub)

```go
var inserted map[string]any
err := client.Hub.Database.InsertRecord(ctx, "products",
	map[string]any{"document": `{"title":"Shoe","price":12}`}, &inserted)

var page map[string]any
err = client.Hub.Database.FindRecords(ctx, "products", map[string]any{
	"filter":   `{"price":{"$gt":10}}`,
	"sortBy":   "price",
	"pageSize": 20,
}, &page)
```

### Records the signed-in user is responsible for (Api)

```go
var mine map[string]any
err := client.API.Database.FindOwn(ctx, "products",
	map[string]any{"pageSize": 20}, &mine)
```

### Taxonomy trees

```go
// Every taxonomy, with its terms.
var tree map[string]any
err := client.Hub.Database.GetDatabaseTaxonomyTree(ctx,
	map[string]any{"includeTerms": true}, &tree)

// The terms of "categories" merged with the terms of the taxonomies it
// depends on (Api: FindMergedTermTree).
var merged map[string]any
err = client.Hub.Database.GetDatabaseMergedTermTree(ctx, "categories", nil, &merged)
```

The taxonomy list (`GetDatabaseTaxonomies`, `dtos.TaxonomyListProjection`)
carries `Description`, `Dependencies`, `ParentName` and `DependencyRefs`: one
`dtos.TaxonomyRef{Id, Name}` per dependency, in the order of `Dependencies`.
A dependency that no longer exists keeps its place with an empty `Name`.
(`DependencyNames` is gone.)

```go
var list struct {
	Result []dtos.TaxonomyListProjection `json:"result"`
}
err := client.Hub.Database.GetDatabaseTaxonomies(ctx, nil, &list)
for _, ref := range list.Result[0].DependencyRefs {
	fmt.Println(ref.Id, ref.Name) // Name == "" → the dependency was deleted
}
```

### Update or delete many records

`"update"` is the **bare field document**: the gateway applies it with
`$set` itself. A body with `$` operators (`$set`, `$inc`, …) is refused with
`CM-ERRORS-DATABASE-035`.

An empty filter (`"{}"`) touches every record of the collection, so it is
refused with `CM-ERRORS-DATABASE-037` unless the request also sets
`"allRecords": true`. For update many, a missing filter counts as `"{}"`.

```go
// Only the records that match the filter.
err := client.API.Database.UpdateMany(ctx, "products", map[string]any{
	"filter": `{"brand":"acme"}`,
	"update": `{"price":12}`,
}, nil)

// Every record, on purpose.
err = client.API.Database.DeleteMany(ctx, "products", map[string]any{
	"filter":     "{}",
	"allRecords": true,
}, nil)
```

Callers that only have own-record rights (`createAsUser`, `updateOwn`,
`deleteOwn`) may call insert many, update many and delete many; those calls
touch only the caller's own records.

## Errors to expect

Every refusal is an `*errors.Error`; `Code` is the gateway's code and
`Errors[0].Context` holds its extra values.

| code | when | context |
|---|---|---|
| `CM-ERRORS-DATABASE-031` | Api `FindTerms` / `FindTermsChildren`: the filter has `$where`, `$function` or `$accumulator` | `Operator` |
| `CM-ERRORS-DATABASE-035` | update one / update many: the update has `$` operators | `Operator` |
| `CM-ERRORS-DATABASE-036` | insert one / insert many / replace: the record is not a JSON object ("Invalid record document"; was `-005`) | `Index` (insert many) |
| `CM-ERRORS-DATABASE-037` | update many / delete many: empty filter without `allRecords` | `Operation` |
| `CM-ERRORS-MEMBERSHIP-USERS-012` | change responsibility: the new owner is not a user of the project in the request environment | — |
| `CM-ERRORS-SCHEMA-002` | rename: another schema in the same environment already uses the name | `SchemaName` |
| `CM-ERRORS-SCHEMA-018` | delete schema: a saved aggregate uses it as its start or a joined collection | `SchemaId`, `BlockerAggregateIds`, `BlockerAggregateNames` |
| `CM-ERRORS-TAXONOMIES-005` | term read by a name over 40 characters | — |
| `CM-ERRORS-TAXONOMIES-010` | term read or merged tree: no taxonomy has that name | — |
| `CM-ERRORS-TAXONOMIES-011` | whole-taxonomy tree, merged tree or `includeTerms`: more than 5000 terms (read a sub-tree with `rootTermId` + `depth`) | `MaxTerms` |
| `CM-ERRORS-TRIGGERS-002` | schema trigger enable / disable / delete: no copy in the request environment; save: the trigger id belongs to another schema | — |

Other behaviour to know:

- Update, replace and change responsibility do not match soft-deleted
  records: such a record is "not found", and update many skips it.
- `GetDatabaseTaxonomyTree` with `includeTerms` fails when the term read
  fails (it used to return the taxonomies without terms).
- Term reads by name (term tree, list, children, merged tree; Hub and Api)
  need `database:read` on `database:term:<taxonomy id>`; the merged tree asks
  it for every nested taxonomy too.
- `SaveDatabaseTaxonomy` with the `viewId` of an existing taxonomy is an
  update and needs `database:update` on `database:taxonomy:<viewId>`; without
  one it is a create (`database:create` on all).
- `TestDatabaseAggregate` needs `database:create` or `database:update` on
  `database:aggregate:<schemaId>`, plus read. Read-only callers are refused.
- `RenameDatabaseSchema` takes only `"title"` (`renameUniqueName` is gone).

### Schema embed and list settings

```go
err := client.Hub.Database.UpdateDatabaseSchemaEmbed(ctx, schemaID, map[string]any{
	"embed": dtos.SchemaEmbedSettingsDto{
		Enabled:                true,
		Fields:                 []string{"title", "description"},
		EmbeddingIntegrationId: "emb_…",
	},
}, nil)

err = client.Hub.Database.UpdateDatabaseSchemaListSettings(ctx, schemaID, map[string]any{
	"settings": dtos.SchemaListSettingsDto{
		Columns:     []*dtos.SchemaListColumnDto{{Field: "title"}},
		DefaultSort: &dtos.SchemaListSortDto{Field: "title", Order: 1},
	},
}, nil)
```

## Methods

### Hub — `client.Hub.Database` (67 methods)

**Module**

| method | verb | path |
|---|---|---|
| `DisableDatabase(ctx, req, out)` | `PUT` | `/database/disable` |
| `EnableDatabase(ctx, req, out)` | `PUT` | `/database/enable` |

**Records and collections**

| method | verb | path |
|---|---|---|
| `SeedCollectionRecords(ctx, req, out)` | `POST` | `/database/collections/seed` |
| `FindRecords(ctx, collectionName, req, out)` | `GET` | `/database/collections/{collectionName}` |
| `InsertRecord(ctx, collectionName, req, out)` | `POST` | `/database/collections/{collectionName}` |
| `AggregateRecords(ctx, collectionName, req, out)` | `POST` | `/database/collections/{collectionName}/aggregate` |
| `ExecuteRecordsAggregate(ctx, collectionName, aggregateId, req, out)` | `POST` | `/database/collections/{collectionName}/aggregates/{aggregateId}/execute` |
| `CountRecords(ctx, collectionName, req, out)` | `GET` | `/database/collections/{collectionName}/count` |
| `DistinctRecordValues(ctx, collectionName, req, out)` | `GET` | `/database/collections/{collectionName}/distinct` |
| `GetCollectionIndexes(ctx, collectionName, req, out)` | `GET` | `/database/collections/{collectionName}/indexes` |
| `DeleteManyRecords(ctx, collectionName, req, out)` | `DELETE` | `/database/collections/{collectionName}/many` |
| `InsertManyRecords(ctx, collectionName, req, out)` | `POST` | `/database/collections/{collectionName}/many` |
| `UpdateManyRecords(ctx, collectionName, req, out)` | `PUT` | `/database/collections/{collectionName}/many` |
| `DeleteRecord(ctx, collectionName, id, req, out)` | `DELETE` | `/database/collections/{collectionName}/{id}` |
| `FindOneRecord(ctx, collectionName, id, req, out)` | `GET` | `/database/collections/{collectionName}/{id}` |
| `UpdateOneRecord(ctx, collectionName, id, req, out)` | `PUT` | `/database/collections/{collectionName}/{id}` |
| `ReplaceRecord(ctx, collectionName, id, req, out)` | `PUT` | `/database/collections/{collectionName}/{id}/replace` |
| `ChangeRecordResponsibility(ctx, collectionName, id, req, out)` | `PUT` | `/database/collections/{collectionName}/{id}/responsibility` |

**Schemas**

| method | verb | path |
|---|---|---|
| `DeleteDatabaseSchema(ctx, id, req, out)` | `DELETE` | `/database/schemas/{Id}` |
| `DiscardDatabaseSchemaDraft(ctx, id, req, out)` | `DELETE` | `/database/schemas/{Id}/draft` |
| `GetDatabaseSchema(ctx, id, req, out)` | `GET` | `/database/schemas/{id}` |
| `GetDatabaseSchemas(ctx, req, out)` | `GET` | `/database/schemas` |
| `GetDatabaseSchemaDraft(ctx, id, req, out)` | `GET` | `/database/schemas/{Id}/draft` |
| `GetDatabaseSchemaVersionDiff(ctx, id, req, out)` | `GET` | `/database/schemas/{Id}/versions/diff` |
| `GetDatabaseSchemaVersions(ctx, id, req, out)` | `GET` | `/database/schemas/{Id}/versions` |
| `PublishDatabaseSchema(ctx, id, req, out)` | `POST` | `/database/schemas/{Id}/publish` |
| `RenameDatabaseSchema(ctx, id, req, out)` | `PUT` | `/database/schemas/{Id}/rename` |
| `SaveDatabaseSchema(ctx, req, out)` | `POST` | `/database/schemas` |
| `UpdateDatabaseSchemaDraft(ctx, id, req, out)` | `PUT` | `/database/schemas/{Id}/draft` |
| `UpdateDatabaseSchemaSettings(ctx, id, req, out)` | `PUT` | `/database/schemas/{Id}/settings` |
| `ApplyDatabaseSchemaBundle(ctx, req, out)` | `POST` | `/database/schemas/apply-bundle` |
| `UpdateDatabaseSchemaEmbed(ctx, id, req, out)` | `PUT` | `/database/schemas/{Id}/embed` |
| `GetDatabaseSchemaListSettings(ctx, id, req, out)` | `GET` | `/database/schemas/{Id}/list-settings` |
| `UpdateDatabaseSchemaListSettings(ctx, id, req, out)` | `PUT` | `/database/schemas/{Id}/list-settings` |

`DeleteDatabaseSchema` also drops the schema's records — its MongoDB collection, with its indexes — in the request environment (the client's `Env`, on every active database integration of that environment). For a schema with AI embed on, its records are also removed from the AI knowledge. The delete is still refused when a saved aggregate or a schema trigger uses the schema; nothing is dropped then. The request and response shape did not change, and a retry is safe (idempotent).

`GetDatabaseSchemas` returns only the schemas of the request environment (the client's `Env`, sent as the `norbix-env` header; `PROD` when none is set). Each row carries that environment in `Env`. The paging cursors (`startingAfter` / `endingBefore`) are schema view ids (`sch_…`); a cursor saved before this gateway change no longer matches.

**Schema triggers**

| method | verb | path |
|---|---|---|
| `DeleteSchemaTrigger(ctx, triggerId, req, out)` | `DELETE` | `/database/schemas/triggers/{triggerId}` |
| `DisableSchemaTrigger(ctx, triggerId, req, out)` | `PATCH` | `/database/schemas/triggers/{triggerId}/disable` |
| `EnableSchemaTrigger(ctx, triggerId, req, out)` | `PATCH` | `/database/schemas/triggers/{triggerId}/enable` |
| `GetSchemaTrigger(ctx, id, req, out)` | `GET` | `/database/schemas/triggers/{id}` |
| `GetSchemaTriggers(ctx, req, out)` | `GET` | `/database/schemas/triggers` |
| `SaveSchemaTrigger(ctx, req, out)` | `POST` | `/database/schemas/triggers` |

Schema triggers live per environment. `GetSchemaTriggers` lists only the
request environment's triggers (the client's `Env`, sent as `norbix-env`;
`PROD` when none is set), and each row carries `Env`. Enable, disable and
delete act on the copy in that environment. `GetSchemaTrigger` returns
`dtos.SchemaTriggerDto` with `Env`; its `SchemaId` is the owning schema
(`sch_…`) — it used to hold the trigger's own id by mistake.

**Taxonomies and terms**

| method | verb | path |
|---|---|---|
| `DeleteDatabaseTaxonomy(ctx, id, req, out)` | `DELETE` | `/database/taxonomies/{Id}` |
| `GetDatabaseTaxonomy(ctx, id, req, out)` | `GET` | `/database/taxonomies/{id}` |
| `GetDatabaseTaxonomies(ctx, req, out)` | `GET` | `/database/taxonomies` |
| `SaveDatabaseTaxonomy(ctx, req, out)` | `POST` | `/database/taxonomies` |
| `DeleteDatabaseTaxonomyTerm(ctx, taxonomyId, id, req, out)` | `DELETE` | `/database/taxonomies/{TaxonomyId}/terms/{Id}` |
| `DeleteManyDatabaseTaxonomyTerms(ctx, taxonomyId, req, out)` | `DELETE` | `/database/taxonomies/{TaxonomyId}/terms/many` |
| `GetDatabaseTaxonomyTerm(ctx, taxonomyId, id, req, out)` | `GET` | `/database/taxonomies/{TaxonomyId}/terms/{Id}` |
| `SaveDatabaseTaxonomyTerm(ctx, taxonomyId, req, out)` | `POST` | `/database/taxonomies/{TaxonomyId}/terms` |
| `UpdateDatabaseTaxonomyTerm(ctx, taxonomyId, id, req, out)` | `PUT` | `/database/taxonomies/{TaxonomyId}/terms/{Id}` |
| `GetDatabaseTaxonomyTree(ctx, req, out)` | `GET` | `/database/taxonomies/tree` |
| `GetDatabaseMergedTermTree(ctx, taxonomyName, req, out)` | `GET` | `/database/taxonomies/{TaxonomyName}/merged-tree` |
| `GetDatabaseTaxonomyTermTree(ctx, taxonomyName, req, out)` | `GET` | `/database/taxonomies/{TaxonomyName}/terms/tree` |

**Database integrations**

| method | verb | path |
|---|---|---|
| `DeleteDatabaseIntegration(ctx, id, req, out)` | `DELETE` | `/database/integrations/{Id}` |
| `DisableDatabaseIntegration(ctx, id, req, out)` | `PUT` | `/database/integrations/{Id}/disable` |
| `EnableDatabaseIntegration(ctx, id, req, out)` | `PUT` | `/database/integrations/{Id}/enable` |
| `GetDatabaseIntegration(ctx, id, req, out)` | `GET` | `/database/integrations/{id}` |
| `GetDatabaseIntegrations(ctx, req, out)` | `GET` | `/database/integrations` |
| `SaveDatabaseIntegration(ctx, req, out)` | `POST` | `/database/integrations` |
| `SetDatabaseIntegrationAsDefault(ctx, id, req, out)` | `PUT` | `/database/integrations/{Id}/default` |
| `GetAllowedFlexTiers(ctx, req, out)` | `GET` | `/database/integrations/flex-tiers` |
| `TestDatabaseIntegration(ctx, req, out)` | `POST` | `/database/integrations/test` |
| `RevealManagedFlexConnectionString(ctx, id, req, out)` | `GET` | `/database/integrations/{Id}/connection-string` |

**Saved aggregates**

| method | verb | path |
|---|---|---|
| `DeleteDatabaseAggregate(ctx, id, req, out)` | `DELETE` | `/database/aggregates/{Id}` |
| `GetDatabaseAggregate(ctx, id, req, out)` | `GET` | `/database/aggregates/{Id}` |
| `GetDatabaseAggregates(ctx, req, out)` | `GET` | `/database/aggregates` |
| `SaveDatabaseAggregate(ctx, req, out)` | `POST` | `/database/aggregates` |
| `TestDatabaseAggregate(ctx, req, out)` | `POST` | `/database/aggregates/test` |

A saved aggregate (`dtos.MongoDbAggregateDto`) lists the collections its
pipeline joins in `JoinedCollections`.

### Api — `client.API.Database` (22 methods)

**Records and collections**

| method | verb | path |
|---|---|---|
| `Aggregate(ctx, collectionName, req, out)` | `POST` | `/database/collections/{collectionName}/aggregate` |
| `ChangeResponsibility(ctx, collectionName, id, req, out)` | `PUT` | `/database/collections/{collectionName}/{id}/responsibility` |
| `Count(ctx, collectionName, req, out)` | `GET` | `/database/collections/{collectionName}/count` |
| `DeleteMany(ctx, collectionName, req, out)` | `DELETE` | `/database/collections/{collectionName}/many` |
| `DeleteOne(ctx, collectionName, id, req, out)` | `DELETE` | `/database/collections/{collectionName}/{id}` |
| `Distinct(ctx, collectionName, req, out)` | `GET` | `/database/collections/{collectionName}/distinct` |
| `ExecuteAggregate(ctx, collectionName, aggregateId, req, out)` | `POST` | `/database/collections/{collectionName}/aggregates/{aggregateId}/execute` |
| `Find(ctx, collectionName, req, out)` | `GET` | `/database/collections/{collectionName}` |
| `FindOne(ctx, collectionName, id, req, out)` | `GET` | `/database/collections/{collectionName}/{id}` |
| `InsertMany(ctx, collectionName, req, out)` | `POST` | `/database/collections/{collectionName}/many` |
| `InsertOne(ctx, collectionName, req, out)` | `POST` | `/database/collections/{collectionName}` |
| `ReplaceOne(ctx, collectionName, id, req, out)` | `PUT` | `/database/collections/{collectionName}/{id}/replace` |
| `UpdateMany(ctx, collectionName, req, out)` | `PUT` | `/database/collections/{collectionName}/many` |
| `UpdateOne(ctx, collectionName, id, req, out)` | `PUT` | `/database/collections/{collectionName}/{id}` |
| `FindOwn(ctx, collectionName, req, out)` | `GET` | `/database/collections/{collectionName}/own` |

**Schemas**

| method | verb | path |
|---|---|---|
| `GetDatabaseSchema(ctx, id, req, out)` | `GET` | `/database/schemas/{id}` |
| `GetDatabaseSchemas(ctx, req, out)` | `GET` | `/database/schemas` |

`GetDatabaseSchemas` returns only the schemas of the request environment (the client's `Env`, sent as the `norbix-env` header; `PROD` when none is set). Each row carries that environment in `Env`. The paging cursors (`startingAfter` / `endingBefore`) are schema view ids (`sch_…`); a cursor saved before this gateway change no longer matches.

**Taxonomies and terms**

| method | verb | path |
|---|---|---|
| `FindTerms(ctx, taxonomyName, req, out)` | `GET` | `/database/taxonomies/{taxonomyName}/terms` |
| `FindTermsChildren(ctx, taxonomyName, parentId, req, out)` | `GET` | `/database/taxonomies/{taxonomyName}/terms/{parentId}/children` |
| `FindTermTree(ctx, taxonomyName, req, out)` | `GET` | `/database/taxonomies/{taxonomyName}/terms/tree` |
| `FindTaxonomyTree(ctx, req, out)` | `GET` | `/database/taxonomies/tree` |
| `FindMergedTermTree(ctx, taxonomyName, req, out)` | `GET` | `/database/taxonomies/{taxonomyName}/merged-tree` |

