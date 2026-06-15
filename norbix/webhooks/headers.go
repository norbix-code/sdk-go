package webhooks

// Outbound Norbix webhook delivery header names.
//
// Source of truth: gateway WebhookDeliveryClient (NOT the inbound API headers
// like norbix-account-id / norbix-project-id used for studio -> gateway calls).
const (
	HeaderEvent       = "X-Norbix-Event"
	HeaderDelivery    = "X-Norbix-Delivery"
	HeaderIdempotency = "Idempotency-Key"
	HeaderAccount     = "X-Norbix-Account"
	HeaderProject     = "X-Norbix-Project"
	HeaderIntegration = "X-Norbix-Integration"
	HeaderDestination = "X-Norbix-Destination"
	HeaderSignature   = "X-Norbix-Signature"
	HeaderTimestamp   = "X-Norbix-Timestamp"
)
