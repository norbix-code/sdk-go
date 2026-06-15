// Package errors defines the error taxonomy returned by the Norbix SDK.
//
// Every failure surfaces as a *Error (or a typed variant such as
// *NotFoundError) so callers can branch on either the HTTP status, the stable
// machine code, or the concrete Go type via errors.As.
package errors

import "fmt"

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

// Error is the base error for all Norbix SDK failures.
type Error struct {
	// Message is the human-readable description.
	Message string
	// Status is the HTTP status code (0 for client-side / network errors).
	Status int
	// Code is the stable machine-readable code (SDK constant or backend errorCode).
	Code string
	// Details is the decoded error body from the backend, when present.
	Details map[string]any
	// wrapped is an optional underlying error (e.g. the net/http error).
	wrapped error
}

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
