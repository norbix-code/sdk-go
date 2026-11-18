package transport

import (
	"encoding/hex"
	"math/big"
	"net"
	"net/url"
	"regexp"
	"strings"
)

// DefaultBaseURLAPI is the SDK's built-in API address.
const DefaultBaseURLAPI = "https://api.norbix.ai"

// ProjectHostLabelPlaceholder marks where a project's label goes in an API URL
// template, e.g. "https://{projectHostLabel}.api.customer.com" (an install's
// GET /v3/echo apiUrlTemplate).
const ProjectHostLabelPlaceholder = "{projectHostLabel}"

const norbixAPIDomain = "api.norbix.ai"

var projectLabelPattern = regexp.MustCompile(`(?i)^pr-[0-9a-f]{32}$`)

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

// ProjectAPIBaseURL is where a project's API calls go, from the API base URL
// the client has (devops docs/tasks/clients-project-host.md, round 2):
//
//   - a template with {projectHostLabel} is filled with the project's label;
//   - an address that is already a project's host ("https://pr-<32 hex>.…")
//     gets the current project's label, so host and norbix-project-id always
//     name the same project;
//   - a norbix.ai API address over https, "https://api.norbix.ai" or
//     "https://<region>.api.norbix.ai", becomes "https://<label>.api.norbix.ai"
//     (the project's DNS record follows its region);
//   - any other address is used as given.
//
// shared is the same address without a project label (browser links such as
// the social sign-in link are built on it). With a project id that is not a
// Norbix id, project is shared. Scheme, port and path are kept; no trailing
// slash is added. It never panics.
func ProjectAPIBaseURL(baseURL, projectID string) (project, shared string) {
	label, ok := ProjectHostLabel(projectID)

	if strings.Contains(baseURL, ProjectHostLabelPlaceholder) {
		shared = strings.ReplaceAll(baseURL, ProjectHostLabelPlaceholder+".", "")
		if !ok {
			return shared, shared
		}
		return strings.ReplaceAll(baseURL, ProjectHostLabelPlaceholder, label), shared
	}

	u, err := url.Parse(baseURL)
	if err != nil || u.Host == "" {
		return baseURL, baseURL
	}
	labels := strings.Split(strings.ToLower(u.Hostname()), ".")

	if len(labels) > 1 && projectLabelPattern.MatchString(labels[0]) {
		shared = withHost(u, strings.Join(labels[1:], "."))
		if !ok {
			return shared, shared
		}
		return withHost(u, label+"."+strings.Join(labels[1:], ".")), shared
	}

	host := strings.Join(labels, ".")
	isNorbixAPI := host == norbixAPIDomain ||
		(len(labels) == 4 && labels[0] != "" && strings.HasSuffix(host, "."+norbixAPIDomain))
	if isNorbixAPI && strings.EqualFold(u.Scheme, "https") && ok {
		return withHost(u, label+"."+norbixAPIDomain), baseURL
	}
	return baseURL, baseURL
}

// withHost is u with another host name; scheme, port, path and query are kept.
func withHost(u *url.URL, hostname string) string {
	c := *u
	if port := u.Port(); port != "" {
		c.Host = net.JoinHostPort(hostname, port)
	} else {
		c.Host = hostname
	}
	return c.String()
}

// apiBase is the API base URL of a call: ProjectAPIBaseURL of Cfg.BaseURLAPI
// and the current project id, computed on every call, so the host and the
// norbix-project-id header always name the same project.
func (t *Transport) apiBase() string {
	project, _ := ProjectAPIBaseURL(t.Cfg.BaseURLAPI, t.Cfg.ProjectID)
	return project
}

// sharedAPIBase is Cfg.BaseURLAPI without a project label, the host of
// browser links (NavigationURL).
func (t *Transport) sharedAPIBase() string {
	_, shared := ProjectAPIBaseURL(t.Cfg.BaseURLAPI, t.Cfg.ProjectID)
	return shared
}
