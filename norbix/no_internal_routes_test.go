package norbix

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// owner 2026-10-22: /internal/* is not SDK surface. The gateway's
// /internal/_typegen route is type-generator tooling; no module may call it.
func TestNoModuleCallsAnInternalRoute(t *testing.T) {
	for _, dir := range []string{"api", "hub"} {
		files, err := filepath.Glob(filepath.Join(dir, "*.go"))
		if err != nil {
			t.Fatal(err)
		}
		for _, f := range files {
			if strings.HasSuffix(f, "_test.go") {
				continue
			}
			src, err := os.ReadFile(f)
			if err != nil {
				t.Fatal(err)
			}
			if strings.Contains(string(src), `"/internal/`) {
				t.Errorf("%s calls an /internal/* route", f)
			}
		}
	}
}
