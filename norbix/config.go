package norbix

import (
	"os"
	"strings"
	"time"

	"github.com/norbix-code/sdk-go/norbix/internal/transport"
)

// Defaults applied when neither an explicit option nor an env var is set.
const (
	DefaultBaseURLAPI = "https://api.norbix.ai"
	DefaultBaseURLHub = "https://hub.norbix.ai"
	DefaultVersion    = "v2"
	DefaultTimeout    = 30 * time.Second
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
	// Region targets a Norbix region (e.g. "nb-eu-germany"). When set on a
	// default base URL the region is composed as a subdomain.
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
