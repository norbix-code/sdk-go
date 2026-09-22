// Package errors defines the error taxonomy returned by the Norbix SDK.
//
// Every failure surfaces as a *Error (or a typed variant such as
// *NotFoundError) so callers can branch on either the HTTP status, the stable
// machine code, or the concrete Go type via errors.As.
package errors

import (
	"encoding/json"
	"fmt"
)

// Code values used by the SDK itself (network / client-side failures). HTTP
// failures carry the backend errorCode instead.
const (
	CodeError            = "NORBIX_ERROR"
	CodeNetwork          = "NORBIX_NETWORK_ERROR"
	CodeNotAuthenticated = "NORBIX_NOT_AUTHENTICATED"
	CodeAccountScope     = "NORBIX_ACCOUNT_SCOPE_REQUIRED"
	CodeMissingPathParam = "NORBIX_MISSING_PATH_PARAM"
	CodeAuthentication   = "NORBIX_AUTHENTICATION_ERROR"
	CodeNotFound         = "NORBIX_NOT_FOUND"
	CodeRateLimit        = "NORBIX_RATE_LIMIT"
	CodeValidation       = "NORBIX_VALIDATION_ERROR"
)

// Item is one error inside the gateway's responseStatus.errors list.
type Item struct {
	// ErrorCode is the gateway's own code, e.g. "CM-ERRORS-FILES-016".
	ErrorCode string `json:"errorCode,omitempty"`
	// Message is the gateway's own text.
	Message string `json:"message,omitempty"`
	// FieldName names the request field the error is about, when it is about one.
	FieldName string `json:"fieldName,omitempty"`
	// Context carries any extra values the gateway attached to this error.
	Context map[string]any `json:"context,omitempty"`
}

// Error is the base error for all Norbix SDK failures.
type Error struct {
	// Message is the human-readable description. For a gateway failure it is
	// the gateway's own text, taken from the first entry of Errors.
	Message string
	// Status is the HTTP status code (0 for client-side / network errors).
	// It is 200 when the gateway refused the call inside a 200 answer.
	Status int
	// Code is the stable machine-readable code (SDK constant or the gateway's
	// errorCode, taken from the first entry of Errors).
	Code string
	// Errors holds every error the gateway sent, in the order it sent them.
	Errors []Item
	// Details is the decoded error body from the backend, when present.
	Details map[string]any
	// Body is the answer exactly as it arrived, for whoever needs the rest of it.
	Body []byte
	// wrapped is an optional underlying error (e.g. the net/http error).
	wrapped error
}

// HTTPStatus is the HTTP status of the answer. Same value as Status; this is
// the name the other Norbix SDKs use for it.
func (e *Error) HTTPStatus() int { return e.Status }

// New builds a base *Error.
func New(message, code string) *Error {
	return &Error{Message: message, Code: code, Details: map[string]any{}}
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s (%d): %s", e.Code, e.Status, e.Message)
}

// Unwrap exposes the wrapped error for errors.Is / errors.As.
func (e *Error) Unwrap() error { return e.wrapped }

// WithWrapped returns a copy carrying the underlying error.
func (e *Error) WithWrapped(err error) *Error {
	e.wrapped = err
	return e
}

// AuthenticationError is returned for HTTP 401 / 403.
type AuthenticationError struct{ Base *Error }

func (e *AuthenticationError) Error() string { return e.Base.Error() }
func (e *AuthenticationError) Unwrap() error { return e.Base }

// NotFoundError is returned for HTTP 404.
type NotFoundError struct{ Base *Error }

func (e *NotFoundError) Error() string { return e.Base.Error() }
func (e *NotFoundError) Unwrap() error { return e.Base }

// RateLimitError is returned for HTTP 429.
type RateLimitError struct{ Base *Error }

func (e *RateLimitError) Error() string { return e.Base.Error() }
func (e *RateLimitError) Unwrap() error { return e.Base }

// ValidationError is returned for HTTP 400 / 422.
type ValidationError struct{ Base *Error }

func (e *ValidationError) Error() string { return e.Base.Error() }
func (e *ValidationError) Unwrap() error { return e.Base }

// FromBody builds the error a gateway answer describes.
//
// The gateway puts its message and its error code inside
// responseStatus.errors[], not at the top of the block, so that list is read
// first: the first entry gives Message and Code, and every entry is kept in
// Errors. Only when the body has no responseStatus are the top-level message
// and errorCode read. "Request failed (HTTP <status>)" is the last fallback,
// used when the body says nothing at all — a 500 page that is not JSON, say.
func FromBody(raw []byte, status int) error {
	details := map[string]any{}
	_ = json.Unmarshal(raw, &details)

	// source is responseStatus when the body has one, the body itself when it
	// has none — so the top-level fields are read only in the second case.
	source := details
	if rs := responseStatusOf(details); rs != nil {
		source = rs
	}

	items := itemsOf(source["errors"])
	message := ""
	code := ""
	for _, it := range items {
		if it.Message != "" || it.ErrorCode != "" {
			message, code = it.Message, it.ErrorCode
			break
		}
	}
	if message == "" {
		message = stringField(source, "message")
	}
	if code == "" {
		code = stringField(source, "errorCode")
	}
	if message == "" {
		message = fmt.Sprintf("Request failed (HTTP %d)", status)
	}
	if code == "" {
		// Callers switch on Code, so it is never left empty; the gateway's own
		// code wins whenever the gateway sent one.
		code = fmt.Sprintf("HTTP_%d", status)
	}

	err := FromHTTP(message, status, code, details)
	base := baseOf(err)
	base.Errors = items
	base.Body = raw
	return err
}

// responseStatusOf returns the responseStatus block, whatever its casing.
func responseStatusOf(body map[string]any) map[string]any {
	for _, key := range []string{"responseStatus", "ResponseStatus"} {
		if v, ok := body[key]; ok {
			if m, ok := v.(map[string]any); ok {
				return m
			}
		}
	}
	return nil
}

// SaysItFailed reports whether the body carries responseStatus.isSuccess=false.
//
// The gateway answers a business refusal — an unknown id, a rule that says no —
// with HTTP 200 and that flag. Without this check the SDK would hand such an
// answer back as a value and the caller would carry on as if the call had
// worked (10b-files, issue #67).
func SaysItFailed(raw []byte) bool {
	body := map[string]any{}
	if err := json.Unmarshal(raw, &body); err != nil {
		return false
	}
	rs := responseStatusOf(body)
	if rs == nil {
		return false
	}
	for _, key := range []string{"isSuccess", "IsSuccess"} {
		if v, ok := rs[key]; ok {
			if b, ok := v.(bool); ok {
				return !b
			}
		}
	}
	return false
}

func itemsOf(v any) []Item {
	list, ok := v.([]any)
	if !ok {
		return nil
	}
	items := make([]Item, 0, len(list))
	for _, entry := range list {
		m, ok := entry.(map[string]any)
		if !ok {
			continue
		}
		item := Item{
			ErrorCode: stringField(m, "errorCode"),
			Message:   stringField(m, "message"),
			FieldName: stringField(m, "fieldName"),
		}
		if ctx, ok := m["context"].(map[string]any); ok {
			item.Context = ctx
		}
		items = append(items, item)
	}
	return items
}

func stringField(m map[string]any, key string) string {
	if v, ok := m[key]; ok {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

// baseOf digs the *Error out of any of the typed wrappers.
func baseOf(err error) *Error {
	switch e := err.(type) {
	case *Error:
		return e
	case *NotFoundError:
		return e.Base
	case *RateLimitError:
		return e.Base
	case *AuthenticationError:
		return e.Base
	case *ValidationError:
		return e.Base
	default:
		return &Error{}
	}
}

// FromHTTP maps an HTTP failure to the most specific error type available.
func FromHTTP(message string, status int, code string, details map[string]any) error {
	base := &Error{Message: message, Status: status, Code: code, Details: details}
	if base.Details == nil {
		base.Details = map[string]any{}
	}
	switch {
	case status == 404:
		return &NotFoundError{base}
	case status == 429:
		return &RateLimitError{base}
	case status == 401 || status == 403:
		return &AuthenticationError{base}
	case status == 400 || status == 422:
		return &ValidationError{base}
	default:
		return base
	}
}
