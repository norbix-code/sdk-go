package transport

import (
	"encoding/hex"
	"math/big"
	"strings"
)

// DefaultBaseURLAPI is the SDK's built-in API address. Only this address is
// replaced by a project's own host; a user-supplied URL is used as given.
const DefaultBaseURLAPI = "https://api.norbix.ai"

const (
	base62Alphabet  = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	projectIDPrefix = "pr_"
)

// guidHexOrder is the .NET "N" format: the first three groups of a GUID are
// stored little-endian.
var guidHexOrder = [16]int{3, 2, 1, 0, 5, 4, 7, 6, 8, 9, 10, 11, 12, 13, 14, 15}

var big62 = big.NewInt(62)

// ProjectHostLabel is the DNS label of a project:
// "pr_5XvmrTA6JVggZdmEE5bGdm" -> "pr-022c4aca9f9e4df787ab0a20c83c44b6".
// The project id ("pr_" + base62) is the same GUID the gateway writes as hex,
// so the label is built with no call. ok is false for anything that is not a
// Norbix project id (a test id, another prefix, a value that is not a
// version-4 GUID); it never panics.
func ProjectHostLabel(projectID string) (label string, ok bool) {
	if !strings.HasPrefix(projectID, projectIDPrefix) {
		return "", false
	}
	digits := projectID[len(projectIDPrefix):]
	// 62^22 is just above 2^128: a GUID never needs more than 22 digits.
	if len(digits) == 0 || len(digits) > 22 {
		return "", false
	}

	value := new(big.Int)
	for i := 0; i < len(digits); i++ {
		digit := strings.IndexByte(base62Alphabet, digits[i])
		if digit < 0 {
			return "", false
		}
		value.Mul(value, big62)
		value.Add(value, big.NewInt(int64(digit)))
	}
	if value.BitLen() > 128 {
		return "", false
	}

	// The number is the GUID's bytes, least significant first (.NET
	// Guid.ToByteArray()); FillBytes writes the most significant first.
	var be [16]byte
	value.FillBytes(be[:])
	var b [16]byte
	for i := range b {
		b[i] = be[15-i]
	}

	// Only a version-4 GUID is a Norbix id (the gateway refuses the rest).
	if b[7]>>4 != 4 || b[8]>>6 != 2 {
		return "", false
	}

	var ordered [16]byte
	for i, src := range guidHexOrder {
		ordered[i] = b[src]
	}
	return "pr-" + hex.EncodeToString(ordered[:]), true
}

// apiBase is the API base URL of a call. On the SDK's built-in address a
// project is served on its own host, "https://pr-<32 hex>.api.norbix.ai", with
// no region label (the project's DNS record already points at its region). It
// is computed from the current project id on every call, so the host and the
// norbix-project-id header always name the same project. A user-supplied URL,
// and a project id that does not decode, keep Cfg.BaseURLAPI.
func (t *Transport) apiBase() string {
	if t.Cfg.BaseURLAPIIsDefault {
		if label, ok := ProjectHostLabel(t.Cfg.ProjectID); ok {
			return strings.Replace(DefaultBaseURLAPI, "://", "://"+label+".", 1)
		}
	}
	return t.Cfg.BaseURLAPI
}
