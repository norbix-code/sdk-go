package norbix

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
)

// The first three labels are the server's own values (the project DTO
// hostLabel, a project's DNS record, a project's API host), not this
// decoder's — devops docs/tasks/clients-project-host.md.
var projectHostVectors = [][2]string{
	{"pr_5XvmrTA6JVggZdmEE5bGdm", "pr-022c4aca9f9e4df787ab0a20c83c44b6"},
	{"pr_3vlTD52e9VK8lXzdv0SmWG", "pr-184336d00e2141c3be5138e8dcf52e81"},
	{"pr_89W9MJvt0StwxsEvyohlJ", "pr-60d101b1938747659947d507192a5204"},
	{"pr_xCeyU0VjHvjVi5MpOYSGY", "pr-6b8b6182960c416098d153d3eb0e601f"},
}

var notAProjectID = []string{
	"p1",
	"proj_123",
	"pr_",
	"pr_!!!",
	"pr_zzzzzzzzzzzzzzzzzzzzzzz", // 23 characters
	"pr_zzzzzzzzzzzzzzzzzzzzzz",  // 22 characters, 2^128 or more
	"pr_0",                       // not a version-4 GUID
	"acc_oOAAIfAzbauZ6f0NHEHjP",
	"",
}

const (
	virginia     = "pr_5XvmrTA6JVggZdmEE5bGdm"
	virginiaHost = "https://pr-022c4aca9f9e4df787ab0a20c83c44b6.api.norbix.ai/"
	germany      = "pr_3vlTD52e9VK8lXzdv0SmWG"
	germanyHost  = "https://pr-184336d00e2141c3be5138e8dcf52e81.api.norbix.ai/"
)

// recorder is an http.RoundTripper that answers every request with `{}` and
// keeps the last one, so a test sees the URL and headers with no network.
type recorder struct {
	url     string
	headers http.Header
	body    string
}

func (r *recorder) RoundTrip(req *http.Request) (*http.Response, error) {
	r.url = req.URL.String()
	r.headers = req.Header.Clone()
	body := r.body
	if body == "" {
		body = "{}"
	}
	return &http.Response{
		StatusCode: http.StatusOK,
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       io.NopCloser(strings.NewReader(body)),
		Request:    req,
	}, nil
}

// hostClient builds a full client on a recorder, with the NORBIX_* variables
// that pick a host cleared.
func hostClient(t *testing.T, o hostTestOptions) (*Client, *recorder) {
	t.Helper()
	for _, k := range []string{"NORBIX_PROJECT_ID", "NORBIX_API_URL", "NORBIX_HUB_URL", "NORBIX_REGION", "NORBIX_ENV", "NORBIX_API_VERSION", "NORBIX_HUB_VERSION", "NORBIX_ACCOUNT_ID"} {
		t.Setenv(k, "")
	}
	for k, v := range o.envOverride {
		t.Setenv(k, v)
	}
	rec := &recorder{}
	c, err := New(o.Options, &http.Client{Transport: rec})
	if err != nil {
		t.Fatal(err)
	}
	return c, rec
}

func apiCall(t *testing.T, c *Client) {
	t.Helper()
	if err := c.API.Database.Find(context.Background(), "orders", nil, nil); err != nil {
		t.Fatal(err)
	}
}

func hubCall(t *testing.T, c *Client) {
	t.Helper()
	if err := c.Hub.Regions.List(context.Background(), nil, nil); err != nil {
		t.Fatal(err)
	}
}

func wantPrefix(t *testing.T, got, prefix string) {
	t.Helper()
	if !strings.HasPrefix(got, prefix) {
		t.Errorf("got %q, want it to start with %q", got, prefix)
	}
}

func wantHeader(t *testing.T, rec *recorder, name, want string) {
	t.Helper()
	if got := rec.headers.Get(name); got != want {
		t.Errorf("%s = %q, want %q", name, got, want)
	}
}

func TestProjectHostLabelVectors(t *testing.T) {
	for _, v := range projectHostVectors {
		label, ok := ProjectHostLabel(v[0])
		if !ok || label != v[1] {
			t.Errorf("ProjectHostLabel(%q) = %q, %v; want %q", v[0], label, ok, v[1])
		}
	}
}

func TestProjectHostLabelRefusesWhatIsNotAProjectID(t *testing.T) {
	for _, id := range notAProjectID {
		if label, ok := ProjectHostLabel(id); ok || label != "" {
			t.Errorf("ProjectHostLabel(%q) = %q, %v; want no label", id, label, ok)
		}
	}
}

func TestProjectHostAPICallsGoToTheProjectHostAndStillSendTheProjectHeader(t *testing.T) {
	c, rec := hostClient(t, hostOpts(Options{ProjectID: virginia, APIKey: "k"}))

	apiCall(t, c)

	wantPrefix(t, rec.url, virginiaHost)
	wantHeader(t, rec, "norbix-project-id", virginia)
	wantHeader(t, rec, "nb-region", "")
}

func TestProjectHostHubCallsStayOnTheSharedHost(t *testing.T) {
	c, rec := hostClient(t, hostOpts(Options{ProjectID: virginia, APIKey: "k"}))

	hubCall(t, c)

	wantPrefix(t, rec.url, "https://hub.norbix.ai/")
	wantHeader(t, rec, "norbix-project-id", virginia)
}

func TestProjectHostFromTheProjectIDEnvironmentVariable(t *testing.T) {
	opts := hostOpts(Options{APIKey: "k"})
	opts.envOverride = map[string]string{"NORBIX_PROJECT_ID": germany}
	c, rec := hostClient(t, opts)

	apiCall(t, c)

	wantPrefix(t, rec.url, germanyHost)
	wantHeader(t, rec, "norbix-project-id", germany)
}

func TestProjectHostLoginGoesToTheProjectHost(t *testing.T) {
	c, rec := hostClient(t, hostOpts(Options{ProjectID: virginia}))
	rec.body = `{"bearerToken":"t"}`

	if _, err := c.Login(context.Background(), LoginCredentials{UserName: "alice", Password: "x"}); err != nil {
		t.Fatal(err)
	}

	if rec.url != virginiaHost+"auth" {
		t.Errorf("login url = %q, want %q", rec.url, virginiaHost+"auth")
	}
}

func TestProjectHostAnIDThatIsNotANorbixIDKeepsTheSharedHost(t *testing.T) {
	c, rec := hostClient(t, hostOpts(Options{ProjectID: "p1", APIKey: "k"}))

	apiCall(t, c)

	wantPrefix(t, rec.url, "https://api.norbix.ai/")
	wantHeader(t, rec, "norbix-project-id", "p1")
}

func TestProjectHostSetScopeMovesTheCallsToTheNewProject(t *testing.T) {
	c, rec := hostClient(t, hostOpts(Options{ProjectID: virginia, APIKey: "k"}))

	c.SetScope(germany, "")
	apiCall(t, c)
	wantPrefix(t, rec.url, germanyHost)
	wantHeader(t, rec, "norbix-project-id", germany)

	c.SetScope("p1", "")
	apiCall(t, c)
	wantPrefix(t, rec.url, "https://api.norbix.ai/")
	wantHeader(t, rec, "norbix-project-id", "p1")

	c.SetScope(virginia, "")
	apiCall(t, c)
	wantPrefix(t, rec.url, virginiaHost)
	wantHeader(t, rec, "norbix-project-id", virginia)
}

func TestProjectHostTheAPIOnlyClientUsesTheProjectHost(t *testing.T) {
	hostClient(t, hostOpts(Options{ProjectID: "p1"})) // clears the environment
	rec := &recorder{}
	c, err := NewAPI(Options{ProjectID: virginia, APIKey: "k"}, &http.Client{Transport: rec})
	if err != nil {
		t.Fatal(err)
	}

	if err := c.Database.Find(context.Background(), "orders", nil, nil); err != nil {
		t.Fatal(err)
	}

	wantPrefix(t, rec.url, virginiaHost)
}

func TestProjectHostTheSocialSignInLinkStaysOnTheSharedHost(t *testing.T) {
	c, _ := hostClient(t, hostOpts(Options{ProjectID: virginia, APIKey: "k"}))

	link, err := c.API.Membership.SocialSignInURL("google", "", "")
	if err != nil {
		t.Fatal(err)
	}

	wantPrefix(t, link, "https://api.norbix.ai/v3/membership/auth/social/google?")
	if !strings.Contains(link, "projectId="+virginia) {
		t.Errorf("link %q does not carry the project id", link)
	}
}

func TestProjectHostWithARegion(t *testing.T) {
	c, rec := hostClient(t, hostOpts(Options{ProjectID: germany, APIKey: "k", Region: "nb-eu-germany"}))

	// API: the project host with no region label; nb-region is still sent.
	apiCall(t, c)
	wantPrefix(t, rec.url, germanyHost)
	wantHeader(t, rec, "nb-region", "nb-eu-germany")
	wantHeader(t, rec, "norbix-project-id", germany)

	// Hub: the regional host.
	hubCall(t, c)
	wantPrefix(t, rec.url, "https://nb-eu-germany.hub.norbix.ai/")
	wantHeader(t, rec, "nb-region", "nb-eu-germany")
}

func TestProjectHostSetRegionChangesTheHubHostOnly(t *testing.T) {
	c, rec := hostClient(t, hostOpts(Options{ProjectID: germany, APIKey: "k"}))

	c.SetRegion("nb-eu-germany")
	apiCall(t, c)
	wantPrefix(t, rec.url, germanyHost)
	wantHeader(t, rec, "nb-region", "nb-eu-germany")
	hubCall(t, c)
	wantPrefix(t, rec.url, "https://nb-eu-germany.hub.norbix.ai/")

	c.SetRegion("")
	apiCall(t, c)
	wantPrefix(t, rec.url, germanyHost)
	wantHeader(t, rec, "nb-region", "")
	hubCall(t, c)
	wantPrefix(t, rec.url, "https://hub.norbix.ai/")
}

func TestProjectHostARegionWithAnIDThatDoesNotDecodeKeepsTheRegionalHost(t *testing.T) {
	c, rec := hostClient(t, hostOpts(Options{ProjectID: "p1", APIKey: "k", Region: "nb-eu-germany"}))

	apiCall(t, c)

	wantPrefix(t, rec.url, "https://nb-eu-germany.api.norbix.ai/")
	wantHeader(t, rec, "nb-region", "nb-eu-germany")
}

func TestProjectHostTheSocialSignInLinkUsesTheRegionalSharedHost(t *testing.T) {
	c, _ := hostClient(t, hostOpts(Options{ProjectID: germany, APIKey: "k", Region: "nb-eu-germany"}))

	link, err := c.API.Membership.SocialSignInURL("google", "", "")
	if err != nil {
		t.Fatal(err)
	}

	wantPrefix(t, link, "https://nb-eu-germany.api.norbix.ai/")
}

func TestProjectHostAnotherAddressIsUsedAsGiven(t *testing.T) {
	c, rec := hostClient(t, hostOpts(Options{
		ProjectID:  virginia,
		APIKey:     "k",
		Region:     "nb-eu-germany",
		BaseURLAPI: "https://api.norbix.io",
		BaseURLHub: "https://hub.norbix.io",
	}))

	apiCall(t, c)
	wantPrefix(t, rec.url, "https://api.norbix.io/")
	wantHeader(t, rec, "norbix-project-id", virginia)
	wantHeader(t, rec, "nb-region", "nb-eu-germany")

	hubCall(t, c)
	wantPrefix(t, rec.url, "https://hub.norbix.io/")

	c.SetScope(germany, "")
	apiCall(t, c)
	wantPrefix(t, rec.url, "https://api.norbix.io/")
	wantHeader(t, rec, "norbix-project-id", germany)

	link, err := c.API.Membership.SocialSignInURL("google", "", "")
	if err != nil {
		t.Fatal(err)
	}
	wantPrefix(t, link, "https://api.norbix.io/")
}

func TestProjectHostAnotherAPIURLFromTheEnvironmentIsUsedAsGiven(t *testing.T) {
	opts := hostOpts(Options{ProjectID: virginia, APIKey: "k"})
	opts.envOverride = map[string]string{"NORBIX_API_URL": "https://api.norbix.io"}
	c, rec := hostClient(t, opts)

	apiCall(t, c)

	wantPrefix(t, rec.url, "https://api.norbix.io/")
}

func TestProjectHostANorbixAPIAddressWrittenOutGoesToTheProjectHost(t *testing.T) {
	opts := hostOpts(Options{APIKey: "k"})
	opts.envOverride = map[string]string{"NORBIX_PROJECT_ID": virginia, "NORBIX_API_URL": "https://api.norbix.ai"}
	c, rec := hostClient(t, opts)
	apiCall(t, c)
	wantPrefix(t, rec.url, virginiaHost)

	regional, rec := hostClient(t, hostOpts(Options{
		ProjectID:  germany,
		APIKey:     "k",
		BaseURLAPI: "https://nb-eu-germany.api.norbix.ai",
	}))
	apiCall(t, regional)
	wantPrefix(t, rec.url, germanyHost)
	wantHeader(t, rec, "norbix-project-id", germany)

	link, err := regional.API.Membership.SocialSignInURL("google", "", "")
	if err != nil {
		t.Fatal(err)
	}
	wantPrefix(t, link, "https://nb-eu-germany.api.norbix.ai/")
}

const customerTemplate = "https://{projectHostLabel}.api.customer.com"

func TestProjectHostATemplateIsFilledAndFollowsSetScope(t *testing.T) {
	c, rec := hostClient(t, hostOpts(Options{
		ProjectID:  virginia,
		APIKey:     "k",
		BaseURLAPI: customerTemplate,
		BaseURLHub: "https://hub.customer.com",
	}))

	apiCall(t, c)
	wantPrefix(t, rec.url, "https://pr-022c4aca9f9e4df787ab0a20c83c44b6.api.customer.com/")
	wantHeader(t, rec, "norbix-project-id", virginia)

	c.SetScope(germany, "")
	apiCall(t, c)
	wantPrefix(t, rec.url, "https://pr-184336d00e2141c3be5138e8dcf52e81.api.customer.com/")
	wantHeader(t, rec, "norbix-project-id", germany)

	link, err := c.API.Membership.SocialSignInURL("google", "", "")
	if err != nil {
		t.Fatal(err)
	}
	wantPrefix(t, link, "https://api.customer.com/v3/")

	hubCall(t, c)
	wantPrefix(t, rec.url, "https://hub.customer.com/")
}

func TestProjectHostATemplateFromTheEnvironmentWithAPort(t *testing.T) {
	opts := hostOpts(Options{APIKey: "k"})
	opts.envOverride = map[string]string{
		"NORBIX_PROJECT_ID": virginia,
		"NORBIX_API_URL":    "http://{projectHostLabel}.api.local.norbix.test:8443",
	}
	c, rec := hostClient(t, opts)

	apiCall(t, c)

	wantPrefix(t, rec.url, "http://pr-022c4aca9f9e4df787ab0a20c83c44b6.api.local.norbix.test:8443/v3/")
}

func TestProjectHostATemplateWithAnIDThatIsNotANorbixID(t *testing.T) {
	c, rec := hostClient(t, hostOpts(Options{ProjectID: "p1", APIKey: "k", BaseURLAPI: customerTemplate}))

	apiCall(t, c)

	wantPrefix(t, rec.url, "https://api.customer.com/")
}

func TestProjectHostAPastedProjectHostFollowsTheClientsProject(t *testing.T) {
	c, rec := hostClient(t, hostOpts(Options{
		ProjectID:  germany,
		APIKey:     "k",
		BaseURLAPI: "https://pr-022c4aca9f9e4df787ab0a20c83c44b6.api.customer.com",
	}))

	apiCall(t, c)
	wantPrefix(t, rec.url, "https://pr-184336d00e2141c3be5138e8dcf52e81.api.customer.com/")
	wantHeader(t, rec, "norbix-project-id", germany)

	link, err := c.API.Membership.SocialSignInURL("google", "", "")
	if err != nil {
		t.Fatal(err)
	}
	wantPrefix(t, link, "https://api.customer.com/")
}

func TestProjectHostAPlainAddressIsUsedAsGiven(t *testing.T) {
	for _, api := range []string{
		"https://api.customer.com",
		"https://api.norbix.io",
		"http://localhost:5002",
		"http://api.norbix.ai",
	} {
		c, rec := hostClient(t, hostOpts(Options{ProjectID: virginia, APIKey: "k", BaseURLAPI: api}))
		apiCall(t, c)
		wantPrefix(t, rec.url, api+"/")
	}
}

func TestProjectAPIBaseURL(t *testing.T) {
	const (
		vLabel = "pr-022c4aca9f9e4df787ab0a20c83c44b6"
		gLabel = "pr-184336d00e2141c3be5138e8dcf52e81"
	)
	cases := []struct {
		base, id, project, shared string
	}{
		{"https://api.norbix.ai", virginia, "https://" + vLabel + ".api.norbix.ai", "https://api.norbix.ai"},
		{"https://api.norbix.ai/v3", virginia, "https://" + vLabel + ".api.norbix.ai/v3", "https://api.norbix.ai/v3"},
		{"https://nb-eu-germany.api.norbix.ai", germany, "https://" + gLabel + ".api.norbix.ai", "https://nb-eu-germany.api.norbix.ai"},
		{"https://api.norbix.ai", "p1", "https://api.norbix.ai", "https://api.norbix.ai"},
		{"https://nb-eu-germany.api.norbix.ai", "p1", "https://nb-eu-germany.api.norbix.ai", "https://nb-eu-germany.api.norbix.ai"},
		{"https://a.b.api.norbix.ai", virginia, "https://a.b.api.norbix.ai", "https://a.b.api.norbix.ai"},
		{"http://api.norbix.ai", virginia, "http://api.norbix.ai", "http://api.norbix.ai"},
		{"https://{projectHostLabel}.api.customer.com/v3", virginia, "https://" + vLabel + ".api.customer.com/v3", "https://api.customer.com/v3"},
		{"https://{projectHostLabel}.api.customer.com", "p1", "https://api.customer.com", "https://api.customer.com"},
		{"https://PR-022C4ACA9F9E4DF787AB0A20C83C44B6.api.customer.com/", germany, "https://" + gLabel + ".api.customer.com/", "https://api.customer.com/"},
		{"https://" + vLabel + ".api.norbix.ai", "p1", "https://api.norbix.ai", "https://api.norbix.ai"},
		{"http://" + vLabel + ".localhost:5002/v3", germany, "http://" + gLabel + ".localhost:5002/v3", "http://localhost:5002/v3"},
		{"https://api.customer.com", virginia, "https://api.customer.com", "https://api.customer.com"},
		{"http://localhost:5002", virginia, "http://localhost:5002", "http://localhost:5002"},
		{"", virginia, "", ""},
		{"not a url", virginia, "not a url", "not a url"},
		{"://bad", virginia, "://bad", "://bad"},
	}
	for _, tc := range cases {
		project, shared := ProjectAPIBaseURL(tc.base, tc.id)
		if project != tc.project || shared != tc.shared {
			t.Errorf("ProjectAPIBaseURL(%q, %q) = %q, %q; want %q, %q", tc.base, tc.id, project, shared, tc.project, tc.shared)
		}
	}
	if ProjectHostLabelPlaceholder != "{projectHostLabel}" {
		t.Errorf("ProjectHostLabelPlaceholder = %q", ProjectHostLabelPlaceholder)
	}
}

func TestProjectHostACustomAPIURLWithTheDefaultHubURL(t *testing.T) {
	c, rec := hostClient(t, hostOpts(Options{
		ProjectID:  virginia,
		APIKey:     "k",
		Region:     "nb-eu-germany",
		BaseURLAPI: "http://localhost:5002",
	}))

	apiCall(t, c)
	wantPrefix(t, rec.url, "http://localhost:5002/")

	// Only the Hub follows the region.
	hubCall(t, c)
	wantPrefix(t, rec.url, "https://nb-eu-germany.hub.norbix.ai/")
}

// hostTestOptions is Options plus the environment variables a test sets.
type hostTestOptions struct {
	Options
	envOverride map[string]string
}

func hostOpts(o Options) hostTestOptions { return hostTestOptions{Options: o} }
