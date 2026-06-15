package webhooks

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// Handler is a typed handler — first arg is the normalised payload (still
// JSON-encoded; decode with DecodePayload / DecodeMutation), second is metadata.
type Handler func(ctx context.Context, payload json.RawMessage, event Event) error

// RawHandler is invoked for OnAll registrations with the full envelope + context.
type RawHandler func(ctx context.Context, env Envelope, c Context) error

// Options configure a Receiver. Empty fields fall back to environment variables.
type Options struct {
	// Secret is the project signing secret. When set, verification runs on every
	// delivery. Defaults to NORBIX_WEBHOOK_SIGNING_SECRET. Omit to skip verify.
	Secret string
	// ToleranceSeconds rejects timestamps outside this window.
	// Defaults to NORBIX_WEBHOOK_TOLERANCE_SECONDS, else 300.
	ToleranceSeconds int
	// ProjectID, when set, rejects deliveries whose envelope projectId differs.
	// Defaults to NORBIX_PROJECT_ID.
	ProjectID string
	// AccountID, when set, rejects deliveries whose envelope accountId differs.
	// Defaults to NORBIX_ACCOUNT_ID.
	AccountID string
}

// Receiver verifies, parses, normalises, and dispatches inbound deliveries.
type Receiver struct {
	handlers  map[string]Handler
	onAll     map[string][]RawHandler
	secret    string
	tolerance int
	projectID string
	accountID string
}

// New builds a Receiver, resolving unset options from the environment.
func New(opts Options) *Receiver {
	tolerance := opts.ToleranceSeconds
	if tolerance == 0 {
		if v := os.Getenv("NORBIX_WEBHOOK_TOLERANCE_SECONDS"); v != "" {
			if n, err := strconv.Atoi(v); err == nil {
				tolerance = n
			}
		}
	}
	if tolerance == 0 {
		tolerance = 300
	}
	return &Receiver{
		handlers:  map[string]Handler{},
		onAll:     map[string][]RawHandler{},
		secret:    orEnv(opts.Secret, "NORBIX_WEBHOOK_SIGNING_SECRET"),
		tolerance: tolerance,
		projectID: orEnv(opts.ProjectID, "NORBIX_PROJECT_ID"),
		accountID: orEnv(opts.AccountID, "NORBIX_ACCOUNT_ID"),
	}
}

func orEnv(v, key string) string {
	if v != "" {
		return v
	}
	return os.Getenv(key)
}

// On registers a typed handler for one event. The latest registration wins.
func (r *Receiver) On(event string, h Handler) *Receiver {
	r.handlers[event] = h
	return r
}

// OnAll registers a raw handler for many events, always invoked after On.
func (r *Receiver) OnAll(events []string, h RawHandler) *Receiver {
	for _, e := range events {
		r.onAll[e] = append(r.onAll[e], h)
	}
	return r
}

// Handle verifies (when a secret is configured), parses, normalises, and
// dispatches the delivery. It returns a *SignatureError on bad signature or
// guard mismatch (respond 401) and a *ParseError on a malformed body (400).
func (r *Receiver) Handle(ctx context.Context, in HandleInput) (HandleResult, error) {
	headers := ParseHeaders(in.Headers)
	var verified *bool

	if !in.SkipVerify && r.secret != "" {
		if err := VerifySignature(r.secret, in.RawBody, headers.Signature, headers.Timestamp, r.tolerance); err != nil {
			return HandleResult{}, err
		}
		t := true
		verified = &t
	}

	env, err := ParseEnvelope(in.RawBody)
	if err != nil {
		return HandleResult{}, err
	}

	if r.projectID != "" && env.ProjectID != r.projectID {
		return HandleResult{}, newSignatureError(
			fmt.Sprintf("delivery projectId %s does not match configured %s", env.ProjectID, r.projectID),
		)
	}
	if r.accountID != "" && env.AccountID != r.accountID {
		return HandleResult{}, newSignatureError(
			fmt.Sprintf("delivery accountId %s does not match configured %s", env.AccountID, r.accountID),
		)
	}

	c := Context{Path: in.Path, Headers: headers, Verified: verified}
	handled := false

	if h, ok := r.handlers[env.Event]; ok {
		norm := Normalize(env)
		event := Event{
			Name:          env.Event,
			DeliveryID:    env.ID,
			CreatedOn:     env.CreatedOn,
			TriggerID:     env.TriggerID,
			AccountID:     firstNonEmptyStr(headers.AccountID, env.AccountID),
			ProjectID:     firstNonEmptyStr(headers.ProjectID, env.ProjectID),
			IntegrationID: headers.IntegrationID,
			DestinationID: headers.DestinationID,
			Verified:      verified,
			Metadata:      norm.Metadata,
			Raw:           env,
		}
		if err := h(ctx, norm.Payload, event); err != nil {
			return HandleResult{}, err
		}
		handled = true
	}

	for _, h := range r.onAll[env.Event] {
		if err := h(ctx, env, c); err != nil {
			return HandleResult{}, err
		}
	}

	return HandleResult{
		Received:   true,
		Event:      env.Event,
		DeliveryID: env.ID,
		Verified:   verified,
		Handled:    handled,
		TriggerID:  env.TriggerID,
	}, nil
}

func firstNonEmptyStr(a, b string) string {
	if a != "" {
		return a
	}
	return b
}
