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
