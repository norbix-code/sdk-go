package norbix

import "testing"

func TestComposeRegionalURL(t *testing.T) {
	got := composeRegionalURL(DefaultBaseURLAPI, "nb-eu-germany")
	want := "https://nb-eu-germany.api.norbix.ai"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
	if composeRegionalURL(DefaultBaseURLAPI, "") != DefaultBaseURLAPI {
		t.Error("empty region must not rewrite")
	}
}

func TestBuildConfigRegionRewritesDefaultsOnly(t *testing.T) {
	cfg, err := buildConfig(Options{
		ProjectID:  "p",
		APIKey:     "k",
		Region:     "nb-eu-germany",
		BaseURLHub: "https://my.hub.example", // custom -> must NOT be rewritten
	}, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if cfg.BaseURLAPI != "https://nb-eu-germany.api.norbix.ai" {
		t.Errorf("default api url should be rewritten, got %q", cfg.BaseURLAPI)
	}
	if cfg.BaseURLHub != "https://my.hub.example" {
		t.Errorf("custom hub url must be untouched, got %q", cfg.BaseURLHub)
	}
}

func TestBuildConfigRequiresProjectID(t *testing.T) {
	t.Setenv("NORBIX_PROJECT_ID", "")
	if _, err := buildConfig(Options{APIKey: "k"}, "Test"); err == nil {
		t.Fatal("expected projectID required error")
	}
}

func TestBuildConfigDefaultsToV3(t *testing.T) {
	t.Setenv("NORBIX_API_VERSION", "")
	t.Setenv("NORBIX_HUB_VERSION", "")
	cfg, err := buildConfig(Options{ProjectID: "p", APIKey: "k"}, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if cfg.APIVersion != "v3" || cfg.HubVersion != "v3" {
		t.Errorf("default versions: got api %q hub %q, want v3", cfg.APIVersion, cfg.HubVersion)
	}

	t.Setenv("NORBIX_API_VERSION", "v2")
	cfg, err = buildConfig(Options{ProjectID: "p", APIKey: "k", HubVersion: "v4"}, "Test")
	if err != nil {
		t.Fatal(err)
	}
	if cfg.APIVersion != "v2" || cfg.HubVersion != "v4" {
		t.Errorf("overrides: got api %q hub %q, want v2 / v4", cfg.APIVersion, cfg.HubVersion)
	}
}
