package webhooks

// Webhook error codes.
const (
	CodeSignatureInvalid = "WEBHOOK_SIGNATURE_INVALID"
	CodeParseInvalid     = "WEBHOOK_PARSE_INVALID"
)

// Error is the base type for webhook receiver failures.
type Error struct {
	Message string
	Code    string
}

func (e *Error) Error() string { return e.Message }

// SignatureError indicates a failed HMAC verification or a guard mismatch
// (worth responding 401). Use errors.As to detect it.
type SignatureError struct{ Base *Error }

func (e *SignatureError) Error() string { return e.Base.Error() }
func (e *SignatureError) Unwrap() error { return e.Base }

func newSignatureError(msg string) *SignatureError {
	return &SignatureError{&Error{Message: msg, Code: CodeSignatureInvalid}}
}

// ParseError indicates a malformed envelope or headers (worth responding 400).
type ParseError struct{ Base *Error }

func (e *ParseError) Error() string { return e.Base.Error() }
func (e *ParseError) Unwrap() error { return e.Base }

func newParseError(msg string) *ParseError {
	return &ParseError{&Error{Message: msg, Code: CodeParseInvalid}}
}
