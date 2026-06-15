package webhooks

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

// headerValue reads a header case-insensitively from a map[string][]string.
func headerValue(headers map[string][]string, name string) string {
	if v, ok := headers[name]; ok && len(v) > 0 {
		return v[0]
	}
	lower := strings.ToLower(name)
	for k, v := range headers {
		if strings.ToLower(k) == lower && len(v) > 0 {
			return v[0]
		}
	}
	return ""
}

// ParseHeaders extracts Norbix delivery headers from an inbound request.
func ParseHeaders(headers map[string][]string) DeliveryHeaders {
	deliveryID := headerValue(headers, HeaderDelivery)
	if deliveryID == "" {
		deliveryID = headerValue(headers, HeaderIdempotency)
	}
	return DeliveryHeaders{
		Event:          headerValue(headers, HeaderEvent),
		DeliveryID:     deliveryID,
		IdempotencyKey: headerValue(headers, HeaderIdempotency),
		AccountID:      headerValue(headers, HeaderAccount),
		ProjectID:      headerValue(headers, HeaderProject),
		IntegrationID:  headerValue(headers, HeaderIntegration),
		DestinationID:  headerValue(headers, HeaderDestination),
		Signature:      headerValue(headers, HeaderSignature),
		Timestamp:      headerValue(headers, HeaderTimestamp),
	}
}

// ParseEnvelope decodes and validates the JSON envelope from the raw body.
func ParseEnvelope(rawBody []byte) (Envelope, error) {
	var env Envelope
	if err := json.Unmarshal(rawBody, &env); err != nil {
		return Envelope{}, newParseError("Webhook body is not valid JSON")
	}
	if env.ID == "" {
		return Envelope{}, newParseError("Webhook envelope missing id")
	}
	if env.Event == "" {
		return Envelope{}, newParseError("Webhook envelope missing event")
	}
	return env, nil
}

// ComputeSignature returns the expected X-Norbix-Signature for a body.
//
// Algorithm (gateway WebhookDeliveryClient.Sign):
//
//	sha256=<hex>  where hex = HMAC-SHA256(secret, "<timestamp>.<rawBody>")
func ComputeSignature(secret, timestamp string, rawBody []byte) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(timestamp))
	mac.Write([]byte("."))
	mac.Write(rawBody)
	return "sha256=" + hex.EncodeToString(mac.Sum(nil))
}

// VerifySignature checks the signature against the raw body. toleranceSeconds
// rejects timestamps older/newer than the window (0 disables the time check).
func VerifySignature(secret string, rawBody []byte, signature, timestamp string, toleranceSeconds int) error {
	if signature == "" {
		return newSignatureError("missing X-Norbix-Signature header")
	}
	if timestamp == "" {
		return newSignatureError("missing X-Norbix-Timestamp header")
	}
	if toleranceSeconds > 0 {
		sent, err := strconv.ParseFloat(timestamp, 64)
		if err != nil {
			return newSignatureError("X-Norbix-Timestamp is not a number")
		}
		age := math.Abs(float64(time.Now().Unix()) - sent)
		if age > float64(toleranceSeconds) {
			return newSignatureError(
				fmt.Sprintf("timestamp outside %ds tolerance (age %ds)", toleranceSeconds, int(age)),
			)
		}
	}
	expected := ComputeSignature(secret, timestamp, rawBody)
	if subtle.ConstantTimeCompare([]byte(expected), []byte(signature)) != 1 {
		return newSignatureError("signature mismatch")
	}
	return nil
}
