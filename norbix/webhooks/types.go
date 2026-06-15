package webhooks

import "encoding/json"

// Envelope is the JSON payload POSTed to every webhook destination.
type Envelope struct {
	// ID is the stable delivery id — dedupe retries on this (also X-Norbix-Delivery).
	ID string `json:"id"`
	// Event is the logical event name, e.g. database.record.inserted.
	Event string `json:"event"`
	// CreatedOn is the ISO-8601 UTC emit time.
	CreatedOn string `json:"createdOn"`
	AccountID string `json:"accountId"`
	ProjectID string `json:"projectId"`
	// TriggerID is the trigger that produced this delivery, if any.
	TriggerID string `json:"triggerId,omitempty"`
	// Data is the raw module-specific payload (record, user, file metadata, ...).
	Data json.RawMessage `json:"data"`
}

// DeliveryHeaders are the parsed Norbix delivery headers on an inbound POST.
type DeliveryHeaders struct {
	Event          string
	DeliveryID     string
	IdempotencyKey string
	AccountID      string
	ProjectID      string
	IntegrationID  string
	DestinationID  string
	Signature      string
	Timestamp      string
}

// EventMetadata holds identifiers lifted off the wire payload. Fields are
// populated only for events that carry them.
type EventMetadata struct {
	UserID        string   `json:"userId,omitempty"`
	SchemaID      string   `json:"schemaId,omitempty"`
	SchemaName    string   `json:"schemaName,omitempty"`
	RecordID      string   `json:"recordId,omitempty"`
	RecordIDs     []string `json:"recordIds,omitempty"`
	IntegrationID string   `json:"integrationId,omitempty"`
}

// Event is the metadata passed to a handler alongside the decoded payload.
type Event struct {
	Name          string
	DeliveryID    string
	CreatedOn     string
	TriggerID     string
	CorrelationID string
	AccountID     string
	ProjectID     string
	IntegrationID string
	DestinationID string
	// Verified is true when the signature was verified; nil when verification
	// was skipped (no secret configured).
	Verified *bool
	Metadata EventMetadata
	// Raw is the escape hatch: the full envelope for unmapped fields.
	Raw Envelope
}

// Context is passed to OnAll handlers alongside the envelope.
type Context struct {
	Path     string
	Headers  DeliveryHeaders
	Verified *bool
}

// HandleInput is the inbound request to Handle.
type HandleInput struct {
	// RawBody is the exact UTF-8 request body bytes (required for signature verify).
	RawBody []byte
	// Headers are the incoming request headers (http.Header or any map[string][]string).
	Headers map[string][]string
	// Path is the optional request path, for logging.
	Path string
	// SkipVerify forces verification off even when a secret is configured.
	SkipVerify bool
}

// HandleResult is the outcome of a delivery.
type HandleResult struct {
	Received   bool
	Event      string
	DeliveryID string
	Verified   *bool
	Handled    bool
	TriggerID  string
}

// Mutation is the {from, to} shape carried by update/replace events.
type Mutation[T any] struct {
	From T `json:"from"`
	To   T `json:"to"`
}
