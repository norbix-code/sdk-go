package norbix

import (
	"os"
	"strings"
	"time"

	"github.com/norbix-code/sdk-go/v2/norbix/internal/transport"
)

// Defaults applied when neither an explicit option nor an env var is set.
const (
	DefaultBaseURLAPI = transport.DefaultBaseURLAPI
	DefaultBaseURLHub = "https://hub.norbix.ai"
	DefaultVersion    = "v3"
	DefaultTimeout    = 30 * time.Second

	// ProjectHostLabelPlaceholder marks where a project's label goes in an API
	// URL template, e.g. "https://{projectHostLabel}.api.customer.com" (an
	// install on its own domain gives it in GET /v3/echo -> apiUrlTemplate).
	ProjectHostLabelPlaceholder = transport.ProjectHostLabelPlaceholder
)

// Options configures a Norbix client. The zero value is valid; unset fields
// fall back to NORBIX_* environment variables, then to the package defaults.
type Options struct {
	ProjectID   string
	APIKey      string
	BearerToken string
	AccountID   string

	// Env selects the project environment ("PROD" is the default and sends no
	// header). Set e.g. "STAGING" to target a non-prod environment.
	Env string
	// Region targets a Norbix region (e.g. "nb-eu-germany") and is sent as the
	// nb-region header. On the default Hub URL it is composed as a subdomain.
	// API calls on a norbix.ai API URL go to the project's own host instead
	// (see ProjectAPIBaseURL); the region subdomain is used for them only when
	// the project id is not a Norbix project id.
	Region string

	BaseURLAPI string
	BaseURLHub string
	APIVersion string
	HubVersion string

	Timeout        time.Duration
	DefaultHeaders map[string]string
}

func env(key string) string { return os.Getenv(key) }

func firstNonEmpty(vals ...string) string {
	for _, v := range vals {
		if v != "" {
			return v
		}
	}
	return ""
}

// ProjectHostLabel is the DNS label of a project's own API host:
// "pr_5XvmrTA6JVggZdmEE5bGdm" -> "pr-022c4aca9f9e4df787ab0a20c83c44b6".
//
// See ProjectAPIBaseURL for where it is used. ok is false for anything that is
// not a Norbix project id (a test id, another prefix, a value that is not a
// version-4 GUID) and the client then keeps the shared host.
func ProjectHostLabel(projectID string) (label string, ok bool) {
	return transport.ProjectHostLabel(projectID)
}

// ProjectAPIBaseURL is where a project's API calls go, from the API base URL
// the client has (BaseURLAPI / NORBIX_API_URL, or the built-in default with
// the region composed). The client applies it on every call with the current
// project:
//
//	https://api.norbix.ai, https://nb-eu-germany.api.norbix.ai -> https://pr-<hex>.api.norbix.ai
//	https://{projectHostLabel}.api.customer.com               -> https://pr-<hex>.api.customer.com
//	https://pr-<hex of another project>.api.customer.com      -> https://pr-<hex>.api.customer.com
//	anything else (https://api.customer.com, localhost, ...)  -> as given
//
// shared is the address without a project label; the social sign-in link is
// built on it. With a project id that is not a Norbix id, project is shared.
// Scheme, port and path are kept. Hub calls are not affected. It never panics.
func ProjectAPIBaseURL(baseURL, projectID string) (project, shared string) {
	return transport.ProjectAPIBaseURL(baseURL, projectID)
}

// composeRegionalURL prefixes the region code as a subdomain of an SDK-default
// base URL, e.g. ("https://api.norbix.ai", "nb-eu-germany") ->
// "https://nb-eu-germany.api.norbix.ai". Only applied to SDK defaults.
func composeRegionalURL(defaultURL, region string) string {
	if region == "" {
		return defaultURL
	}
	scheme := "https://"
	host := defaultURL
	if i := strings.Index(defaultURL, "://"); i >= 0 {
		scheme = defaultURL[:i+3]
		host = defaultURL[i+3:]
	}
	return scheme + region + "." + host
}

// buildConfig resolves Options + env vars + defaults into a transport.Config.
func buildConfig(o Options, clientName string) (*transport.Config, error) {
	projectID := firstNonEmpty(o.ProjectID, env("NORBIX_PROJECT_ID"))
	if projectID == "" {
		return nil, errProjectIDRequired(clientName)
	}

	baseAPI := firstNonEmpty(o.BaseURLAPI, env("NORBIX_API_URL"), DefaultBaseURLAPI)
	baseHub := firstNonEmpty(o.BaseURLHub, env("NORBIX_HUB_URL"), DefaultBaseURLHub)
	apiVersion := firstNonEmpty(o.APIVersion, env("NORBIX_API_VERSION"), DefaultVersion)
	hubVersion := firstNonEmpty(o.HubVersion, env("NORBIX_HUB_VERSION"), DefaultVersion)
	region := firstNonEmpty(o.Region, env("NORBIX_REGION"))

	apiIsDefault := baseAPI == DefaultBaseURLAPI
	hubIsDefault := baseHub == DefaultBaseURLHub

	if region != "" {
		if apiIsDefault {
			baseAPI = composeRegionalURL(DefaultBaseURLAPI, region)
		}
		if hubIsDefault {
			baseHub = composeRegionalURL(DefaultBaseURLHub, region)
		}
	}

	envName := firstNonEmpty(o.Env, env("NORBIX_ENV"), "PROD")

	timeout := o.Timeout
	if timeout == 0 {
		timeout = DefaultTimeout
	}

	headers := o.DefaultHeaders
	if headers == nil {
		headers = map[string]string{}
	}

	return &transport.Config{
		APIKey:              firstNonEmpty(o.APIKey, env("NORBIX_API_KEY")),
		BearerToken:         firstNonEmpty(o.BearerToken, env("NORBIX_BEARER_TOKEN")),
		ProjectID:           projectID,
		AccountID:           firstNonEmpty(o.AccountID, env("NORBIX_ACCOUNT_ID")),
		BaseURLAPI:          baseAPI,
		BaseURLHub:          baseHub,
		APIVersion:          apiVersion,
		HubVersion:          hubVersion,
		Timeout:             timeout,
		Env:                 envName,
		Region:              region,
		BaseURLAPIIsDefault: apiIsDefault,
		BaseURLHubIsDefault: hubIsDefault,
		DefaultHeaders:      headers,
	}, nil
}
