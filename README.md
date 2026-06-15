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

var nf *norbixerr.NotFoundError
if errors.As(err, &nf) {
	// nf.Base.Status == 404, nf.Base.Code, nf.Base.Details
}
```

Types: `AuthenticationError` (401/403), `NotFoundError` (404),
`RateLimitError` (429), `ValidationError` (400/422), base `Error` otherwise.
Idempotent verbs (GET/DELETE) retry on 429/5xx with exponential backoff.

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

Endpoint modules are generated from the same route map the other SDKs use; DTO
structs are generated from the ServiceStack TypeScript DTOs (`hub2.dtos.ts`,
`api2.dtos.ts`) because ServiceStack has no native Go type exporter. Generated
files carry a `DO NOT EDIT` header.

## License

MIT
