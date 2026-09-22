# norbix-go

Official Go SDK for [Norbix](https://norbix.ai). It mirrors the TypeScript,
Python, .NET, Dart, Kotlin and Swift SDKs: a shared transport (auth, scope,
environment and region headers, retries), full API and Hub endpoint coverage,
generated DTO structs, and an inbound webhook receiver with HMAC verification.

> Status: generated foundation. Endpoint methods accept/return JSON
> (`map[string]any` in, any `out` pointer) so you can decode into the generated
> DTO structs in `norbix/api/dtos` and `norbix/hub/dtos`.

## Install

```bash
go get github.com/norbix-code/sdk-go/norbix
```

Requires Go 1.22+.

## Quick start

```go
package main

import (
	"context"
	"fmt"

	"github.com/norbix-code/sdk-go/norbix"
	"github.com/norbix-code/sdk-go/norbix/hub/dtos"
)

func main() {
	// Service-to-service: long-lived API key. Reads NORBIX_* env vars when fields are empty.
	c, err := norbix.New(norbix.Options{
		APIKey:    "sk_...",
		ProjectID: "proj_...",
	})
	if err != nil {
		panic(err)
	}
	defer c.Close()

	ctx := context.Background()

	// API call — decode straight into a generated DTO.
	var schemas map[string]any
	if err := c.API.Database.GetDatabaseSchemas(ctx, nil, &schemas); err != nil {
		panic(err)
	}
	fmt.Println(schemas)

	// API call with a path parameter and a typed result.
	var user dtos.UserDto
	_ = c.API.Membership.GetUser(ctx, "user_123", nil, &user)
}
```

### User-on-behalf-of (login)

```go
c, _ := norbix.New(norbix.Options{ProjectID: "proj_..."})
_, err := c.Login(ctx, norbix.LoginCredentials{UserName: "alice", Password: "..."})
// subsequent calls use the bearer token returned by Login
```

### Client shapes

| Constructor | Access |
|-------------|--------|
| `norbix.New(opts)` | `c.API.<Module>.<Method>` / `c.Hub.<Module>.<Method>` |
| `norbix.NewAPI(opts)` | flat: `c.Database`, `c.Membership`, `c.Chat`, ... |
| `norbix.NewHub(opts)` | flat: `c.Account`, `c.Email`, `c.Webhooks`, ... |

### Configuration

Every `Options` field falls back to an environment variable, then a default:

| Option | Env var | Default |
|--------|---------|---------|
| `ProjectID` | `NORBIX_PROJECT_ID` | required |
| `APIKey` | `NORBIX_API_KEY` | — |
| `BearerToken` | `NORBIX_BEARER_TOKEN` | — |
| `AccountID` | `NORBIX_ACCOUNT_ID` | — |
| `Env` | `NORBIX_ENV` | `PROD` (no header sent) |
| `Region` | `NORBIX_REGION` | none |
| `BaseURLAPI` | `NORBIX_API_URL` | `https://api.norbix.ai` |
| `BaseURLHub` | `NORBIX_HUB_URL` | `https://hub.norbix.ai` |

Setting a `Region` on a default base URL composes it as a subdomain, e.g.
`https://nb-eu-germany.api.norbix.ai`. A custom base URL is never rewritten.
Per-request env/region/token overrides and switching at runtime are available via
`c.SetEnv`, `c.SetRegion`, `c.SetBearerToken`, `c.SetScope`.

### Errors

Failures return `*errors.Error` or a typed variant. Use `errors.As`:

```go
import norbixerr "github.com/norbix-code/sdk-go/norbix/errors"

var e *norbixerr.Error
if errors.As(err, &e) {
	fmt.Println(e.HTTPStatus(), e.Code, e.Message)
	for _, item := range e.Errors {
		fmt.Println(item.ErrorCode, item.FieldName, item.Message)
	}
	_ = e.Body // the answer exactly as it arrived
}
```

Types: `AuthenticationError` (401/403), `NotFoundError` (404),
`RateLimitError` (429), `ValidationError` (400/422), base `Error` otherwise.
Idempotent verbs (GET/DELETE) retry on 429/5xx with exponential backoff.

`Message` and `Code` are the gateway's own. The gateway puts them inside
`responseStatus.errors[]`, so the SDK reads that list first, takes the first
entry for `Message` / `Code`, and keeps every entry in `Errors`. Only when the
body has no `responseStatus` are the top-level `message` and `errorCode` read.
`Request failed (HTTP <status>)` with the code `HTTP_<status>` is the last
fallback, used when the body says nothing — a 500 page that is not JSON, say.

#### Breaking change — a refused call now returns an error

The gateway answers a business refusal (an unknown id, a rule that says no)
with **HTTP 200** and `responseStatus.isSuccess = false`. The SDK used to fill
`out` and return `nil`, so code carried on as if the call had worked. It now
returns a `*errors.Error` with `Status` 200 and the gateway's message and code.

If your code checked `out.ResponseStatus.IsSuccess` itself, check `err`
instead. Endpoints that answer with raw bytes rather than a document (file
download, the public file link — the ones you pass a `*[]byte` to) are not JSON
and are unchanged.

## Working with terms

A **taxonomy** is a named tree of **terms** (labels). A term can have one parent (a clean hierarchy) or several parents (the same item under many categories). Pick the call that matches what you want:

| I want to… | Call | Returns |
| --- | --- | --- |
| Get a taxonomy's terms as a flat list | `FindTerms` | a paginated `list` of terms |
| Get only the children of one term | `FindTermsChildren` | a `list` of child terms (direct + multi-parent) |
| Get a taxonomy's terms as a ready-made tree | `FindTermTree` | a `tree` of nested term nodes |
| Get the taxonomy structure (e.g. Countries → Cities) | `FindTaxonomyTree` | a `tree` of taxonomy nodes |

The examples below all use one example `services` taxonomy shaped like this:

```text
Indoors
  └─ Air conditioning
       └─ Wall-mounted
Outdoors
  └─ Solar panels
```

### List a taxonomy's terms (flat)

**Goal:** show every term of `services` in a simple list, in display order.

```go
var out map[string]any
if err := c.API.Database.FindTerms(ctx, "services", nil, &out); err != nil {
	panic(err)
}
```

```json
{
  "list": {
    "items": [
      { "id": "term_indoors",   "taxonomyName": "services", "parentId": null,           "order": 1, "name": "Indoors" },
      { "id": "term_air_con",   "taxonomyName": "services", "parentId": "term_indoors", "order": 1, "name": "Air conditioning" },
      { "id": "term_wall",      "taxonomyName": "services", "parentId": "term_air_con", "order": 1, "name": "Wall-mounted" },
      { "id": "term_outdoors",  "taxonomyName": "services", "parentId": null,           "order": 2, "name": "Outdoors" },
      { "id": "term_solar",     "taxonomyName": "services", "parentId": "term_outdoors","order": 1, "name": "Solar panels" }
    ],
    "hasMore": false, "hasPrevious": false, "startingAfter": null, "endingBefore": null
  },
  "responseStatus": { "isSuccess": true }
}
```

The list is flat — every term is one row, with its `parentId` telling you where it sits. The nesting is not built for you here (use `FindTermTree` for that).

### List only top-level terms (filtered)

**Goal:** show just the roots (no parent) — for the first level of a menu.

```go
var out map[string]any
if err := c.API.Database.FindTerms(ctx, "services", map[string]any{"filter": "{ \"parentId\": null }"}, &out); err != nil {
	panic(err)
}
```

```json
{
  "list": {
    "items": [
      { "id": "term_indoors",  "taxonomyName": "services", "parentId": null, "order": 1, "name": "Indoors" },
      { "id": "term_outdoors", "taxonomyName": "services", "parentId": null, "order": 2, "name": "Outdoors" }
    ],
    "hasMore": false, "hasPrevious": false, "startingAfter": null, "endingBefore": null
  },
  "responseStatus": { "isSuccess": true }
}
```

`filter` is an optional MongoDB filter, ANDed with the taxonomy. Use it to fetch one level at a time (lazy tree loading) or to find terms by any field.

### Get a term's children

**Goal:** the user expanded *Indoors* — load what is directly under it.

```go
var out map[string]any
if err := c.API.Database.FindTermsChildren(ctx, "services", "term_indoors", nil, &out); err != nil {
	panic(err)
}
```

```json
{
  "list": {
    "items": [
      {
        "id": "term_air_con",
        "taxonomyName": "services",
        "parentId": "term_indoors",
        "order": 1,
        "name": "Air conditioning",
        "multiParents": [
          { "taxonomyId": "tax_service_types", "parentId": "term_indoors",          "name": "Indoors" },
          { "taxonomyId": "tax_service_types", "parentId": "term_energy_efficient", "name": "Energy efficient" }
        ]
      }
    ],
    "hasMore": false, "hasPrevious": false
  },
  "responseStatus": { "isSuccess": true }
}
```

This returns **both** direct children (their `parentId` is `term_indoors`) **and** multi-parent children (terms that list `term_indoors` in `multiParents`). Parent names are already resolved, so no second lookup.

### Multi-parent: one product in several categories

**Goal:** in a `products` taxonomy, a *Relaxing massage oil* belongs to *For couples*, *Gift ideas*, **and** *Body care*. Listing the children of **any** of those categories returns it.

```go
var out map[string]any
if err := c.API.Database.FindTermsChildren(ctx, "products", "term_gift_ideas", nil, &out); err != nil {
	panic(err)
}
```

```json
{
  "list": {
    "items": [
      {
        "id": "term_relaxing_oil",
        "taxonomyName": "products",
        "name": "Relaxing massage oil",
        "multiParents": [
          { "taxonomyId": "tax_categories", "parentId": "term_for_couples", "name": "For couples" },
          { "taxonomyId": "tax_categories", "parentId": "term_gift_ideas",  "name": "Gift ideas" },
          { "taxonomyId": "tax_categories", "parentId": "term_body_care",   "name": "Body care" }
        ]
      }
    ],
    "hasMore": false, "hasPrevious": false
  },
  "responseStatus": { "isSuccess": true }
}
```

One product, three category links — no duplicate listings. The same product would also come back from the children of `term_for_couples` and `term_body_care`.

### Get the whole term tree in one call

**Goal:** render the full `services` tree at once, already nested.

```go
var out map[string]any
if err := c.API.Database.FindTermTree(ctx, "services", nil, &out); err != nil {
	panic(err)
}
```

```json
{
  "tree": [
    {
      "id": "term_indoors",
      "name": "Indoors",
      "order": 1,
      "children": [
        {
          "id": "term_air_con",
          "name": "Air conditioning",
          "order": 1,
          "children": [
            { "id": "term_wall", "name": "Wall-mounted", "order": 1, "children": null }
          ]
        }
      ]
    },
    {
      "id": "term_outdoors",
      "name": "Outdoors",
      "order": 2,
      "children": [
        { "id": "term_solar", "name": "Solar panels", "order": 1, "children": null }
      ]
    }
  ],
  "responseStatus": { "isSuccess": true }
}
```

Roots are in `tree`; each node carries its own `children`; a leaf has `children: null`. The tree arrives ready to render — no client-side tree building.

### Get only a sub-tree, capped by depth

**Goal:** start from *Indoors* and go at most 2 levels deep.

```go
var out map[string]any
if err := c.API.Database.FindTermTree(ctx, "services", map[string]any{"rootTermId": "term_indoors", "depth": 2}, &out); err != nil {
	panic(err)
}
```

```json
{
  "tree": [
    {
      "id": "term_indoors",
      "name": "Indoors",
      "order": 1,
      "children": [
        { "id": "term_air_con", "name": "Air conditioning", "order": 1, "children": null }
      ]
    }
  ],
  "responseStatus": { "isSuccess": true }
}
```

With `depth` 2 you get *Indoors* (level 1) and *Air conditioning* (level 2); *Wall-mounted* (level 3) is cut off, so *Air conditioning* shows `children: null`.

### Get the taxonomy structure tree — without terms

**Goal:** see how taxonomies relate to each other (e.g. a `Cities` taxonomy whose parent is `Countries`), structure only.

```go
var out map[string]any
if err := c.API.Database.FindTaxonomyTree(ctx, nil, &out); err != nil {
	panic(err)
}
```

```json
{
  "tree": [
    {
      "viewId": "txn_countries",
      "taxonomyName": "Countries",
      "taxonomySlug": "countries",
      "parentId": null,
      "children": [
        { "viewId": "txn_cities", "taxonomyName": "Cities", "taxonomySlug": "cities", "parentId": "txn_countries", "children": null, "terms": null }
      ],
      "terms": null
    }
  ],
  "responseStatus": { "isSuccess": true }
}
```

This is the **taxonomy** tree, not the term tree: nodes are taxonomies. Every `terms` is `null` because we did not ask for terms.

### Get the taxonomy structure tree — with terms

**Goal:** same structure, but also pull each taxonomy's terms in the same call.

```go
var out map[string]any
if err := c.API.Database.FindTaxonomyTree(ctx, map[string]any{"includeTerms": true}, &out); err != nil {
	panic(err)
}
```

```json
{
  "tree": [
    {
      "viewId": "txn_countries",
      "taxonomyName": "Countries",
      "taxonomySlug": "countries",
      "parentId": null,
      "terms": [
        { "id": "term_lt", "name": "Lithuania", "order": 1, "children": null },
        { "id": "term_lv", "name": "Latvia",    "order": 2, "children": null }
      ],
      "children": [
        {
          "viewId": "txn_cities",
          "taxonomyName": "Cities",
          "taxonomySlug": "cities",
          "parentId": "txn_countries",
          "terms": [
            { "id": "term_vilnius", "name": "Vilnius", "order": 1, "children": null },
            { "id": "term_kaunas",  "name": "Kaunas",  "order": 2, "children": null }
          ],
          "children": null
        }
      ]
    }
  ],
  "responseStatus": { "isSuccess": true }
}
```

Now each taxonomy node's `terms` holds that taxonomy's full term tree (same shape as `FindTermTree`) — *Countries* carries its countries, *Cities* carries its cities.

> Every term-reading call also accepts an optional `databaseIntegrationId` key in the request map to target a non-default database.

## Webhooks

The `norbix/webhooks` package verifies, parses, normalises, and dispatches
inbound deliveries. Signature algorithm matches the gateway:
`X-Norbix-Signature: sha256=<hex>` where the HMAC-SHA256 input is
`"<timestamp>.<rawBody>"`.

```go
r := webhooks.New(webhooks.Options{}) // reads NORBIX_WEBHOOK_SIGNING_SECRET etc.

r.On(webhooks.EventMembershipUserRegistered,
	func(ctx context.Context, payload json.RawMessage, e webhooks.Event) error {
		user, _ := webhooks.DecodePayload[dtos.UserDto](payload)
		log.Printf("registered %s (%s)", user.UserName, e.Metadata.UserID)
		return nil
	})

// Mutation events ({from, to}) decode with DecodeMutation:
r.On(webhooks.EventDatabaseRecordUpdated,
	func(ctx context.Context, payload json.RawMessage, e webhooks.Event) error {
		m, _ := webhooks.DecodeMutation[MyDoc](payload)
		_ = m.From; _ = m.To
		return nil
	})

// Catch-all logger for every catalog event:
r.OnAll(webhooks.EventNames, func(ctx context.Context, env webhooks.Envelope, c webhooks.Context) error {
	log.Printf("%s %s", env.Event, env.ID)
	return nil
})

res, err := r.Handle(ctx, webhooks.HandleInput{RawBody: raw, Headers: req.Header})
```

In your HTTP handler, respond `401` on `*webhooks.SignatureError` and `400` on
`*webhooks.ParseError`. A full runnable server is in
[`examples/webhook-receiver`](examples/webhook-receiver/main.go).

Normalisation matches the JS SDK: entity events hand you the entity, mutation
events `{from, to}`, batch events the array; wrapper ids (record id, schema,
user id) are lifted onto `event.Metadata`.

## Package layout

```
norbix/
  client.go, client_split.go, config.go   full / API-only / Hub-only clients
  errors/                                 error taxonomy
  internal/transport/                     shared HTTP layer + tests
  api/        <module>.go, namespace.go   API endpoint modules
  api/dtos/                               generated API DTO structs
  hub/        <module>.go, namespace.go   Hub endpoint modules
  hub/dtos/                               generated Hub DTO structs
  webhooks/                               receiver, signature, events + tests
examples/webhook-receiver/                runnable webhook sink
```

## Codegen notes

`norbix/api/dtos/dtos.go` and `norbix/hub/dtos/dtos.go` are generated from the
Norbix gateway's own contract, straight from a running gateway. They carry a
`DO NOT EDIT` header: a change made here is lost on the next run. If a type is
wrong, the gateway or the generator is wrong.

The endpoint modules (`norbix/api/*.go`, `norbix/hub/*.go`) are **not**
generated today — the script that made them is gone. They are maintained by
hand, with a test per method, and they keep their old header until a module
generator exists.

### Regenerate the types

Start the gateway (API on `:5002`, Hub on `:5001`), then run the generator from
the private `typegen` toolchain — one command, both files:

```bash
python3 <typegen>/languages/go/generate.py --out .
go build ./norbix/... && go vet ./norbix/... && gofmt -l norbix/ && go test ./norbix/...
```

`--api-url` and `--hub-url` point it somewhere else; `--api-file` / `--hub-file`
read a saved contract export instead of a live gateway. Two runs on the same
gateway give byte-identical files, so a non-empty `git diff` after a run is a
real contract change — read it before committing.

## License

MIT
