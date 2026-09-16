package norbix

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Public file links — the five endpoints slice PUB added to the gateway
// (10b-files, slice SDK-2).
//
// Every case starts its own fake gateway and its own client, so the order the
// tests run in does not matter and no real storage provider is contacted.

const publicFileID = "nbpf_7hK2abc"

// newPublicFileServer answers every request with the given bytes and records
// what it saw. Unlike newFilesClient it does not pretend the body is JSON,
// because a public link answers with a file.
func newPublicFileServer(t *testing.T, payload []byte, status int) (*httptest.Server, *recorded, *http.Header) {
	t.Helper()
	got := &recorded{}
	seen := &http.Header{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got.method = r.Method
		got.path = r.URL.Path
		got.query = r.URL.RawQuery
		*seen = r.Header.Clone()
		w.Header().Set("Content-Type", "application/pdf")
		w.WriteHeader(status)
		_, _ = w.Write(payload)
	}))
	t.Cleanup(srv.Close)
	return srv, got, seen
}

// clientFor builds a client pointed at srv. Pass an empty apiKey for somebody
// who has never signed in.
func clientFor(t *testing.T, url, apiKey string) *Client {
	t.Helper()
	c, err := New(Options{
		ProjectID:  "proj_1",
		APIKey:     apiKey,
		BaseURLAPI: url,
		BaseURLHub: url,
		APIVersion: "v3",
		HubVersion: "v2",
	})
	if err != nil {
		t.Fatalf("client: %v", err)
	}
	return c
}

// ---- hub: make it public, make it private -----------------------------------

func TestHubPublicEndpointsEachHitTheirOwnRoute(t *testing.T) {
	cases := []struct {
		name string
		call func(*Client, context.Context, map[string]any, any) error
		path string
	}{
		{"MakeFilePublic", func(c *Client, ctx context.Context, req map[string]any, out any) error {
			return c.Hub.Files.MakeFilePublic(ctx, req, out)
		}, "/v2/files/item/public"},
		{"MakeFilePrivate", func(c *Client, ctx context.Context, req map[string]any, out any) error {
			return c.Hub.Files.MakeFilePrivate(ctx, req, out)
		}, "/v2/files/item/private"},
		{"MakeFolderPublic", func(c *Client, ctx context.Context, req map[string]any, out any) error {
			return c.Hub.Files.MakeFolderPublic(ctx, req, out)
		}, "/v2/files/folder/public"},
		{"MakeFolderPrivate", func(c *Client, ctx context.Context, req map[string]any, out any) error {
			return c.Hub.Files.MakeFolderPrivate(ctx, req, out)
		}, "/v2/files/folder/private"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			c, got := newFilesClient(t, `{}`)
			var out map[string]any
			err := tc.call(c, context.Background(), map[string]any{
				"filesIntegrationId": filesIntegrationID,
				"path":               "invoices/invoice.pdf",
			}, &out)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			expectCall(t, got, http.MethodPost, tc.path)
			if got.body["filesIntegrationId"] != filesIntegrationID {
				t.Errorf("integration id: got %v", got.body["filesIntegrationId"])
			}
			if got.body["path"] != "invoices/invoice.pdf" {
				t.Errorf("path: got %v", got.body["path"])
			}
		})
	}
}

func TestMakeFilePublicGivesBackTheMintedID(t *testing.T) {
	c, _ := newFilesClient(t, `{"id":"`+publicFileID+`","status":"Success"}`)
	var out struct {
		ID string `json:"id"`
	}
	err := c.Hub.Files.MakeFilePublic(context.Background(), map[string]any{
		"filesIntegrationId": filesIntegrationID,
		"path":               "invoices/invoice.pdf",
	}, &out)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if out.ID != publicFileID {
		t.Errorf("id: got %q want %q", out.ID, publicFileID)
	}
}

// ---- api: the link anyone can open ------------------------------------------

func TestGetPublicFileReturnsTheBytesUnchanged(t *testing.T) {
	payload := []byte("%PDF-1.7 hello")
	srv, got, _ := newPublicFileServer(t, payload, http.StatusOK)

	var pdf []byte
	err := clientFor(t, srv.URL, "").API.Files.GetPublicFile(
		context.Background(), publicFileID, "invoice.pdf", &pdf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectCall(t, got, http.MethodGet, "/v3/files/public/"+publicFileID+"/invoice.pdf")
	if string(pdf) != string(payload) {
		t.Errorf("bytes: got %q want %q", pdf, payload)
	}
}

func TestGetPublicFileSendsNoAuthorizationHeader(t *testing.T) {
	srv, _, seen := newPublicFileServer(t, []byte("x"), http.StatusOK)

	var out []byte
	// A key IS configured — and must still not be sent. That is the whole
	// point: a public link is not a session.
	err := clientFor(t, srv.URL, "key_1").API.Files.GetPublicFile(
		context.Background(), publicFileID, "invoice.pdf", &out)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if auth := seen.Get("Authorization"); auth != "" {
		t.Errorf("Authorization was sent: %q", auth)
	}
}

func TestGetPublicFileWorksWithNoCredentialsAtAll(t *testing.T) {
	srv, _, _ := newPublicFileServer(t, []byte("x"), http.StatusOK)

	var out []byte
	if err := clientFor(t, srv.URL, "").API.Files.GetPublicFile(
		context.Background(), publicFileID, "invoice.pdf", &out); err != nil {
		t.Fatalf("a link must work with no sign-in: %v", err)
	}
}

func TestAFolderRelativePathKeepsItsSlashes(t *testing.T) {
	srv, got, _ := newPublicFileServer(t, []byte("x"), http.StatusOK)

	var out []byte
	err := clientFor(t, srv.URL, "").API.Files.GetPublicFile(
		context.Background(), "nbpf_folder1", "2026/q1/report.pdf", &out)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectCall(t, got, http.MethodGet, "/v3/files/public/nbpf_folder1/2026/q1/report.pdf")
}

func TestALinkThatPointsAtNothingIsAnError(t *testing.T) {
	srv, _, _ := newPublicFileServer(t, []byte(""), http.StatusNotFound)

	var out []byte
	err := clientFor(t, srv.URL, "").API.Files.GetPublicFile(
		context.Background(), "nbpf_gone", "invoice.pdf", &out)
	if err == nil {
		t.Fatal("expected an error for a 404")
	}
}

// A provider that signs its own links (Amazon S3, Azure Blob, Google Cloud
// Storage) makes the gateway answer 302. net/http follows it on its own, so
// the bytes arrive from the provider.
func TestItFollowsTheRedirectASigningProviderAnswersWith(t *testing.T) {
	payload := []byte("bytes-from-the-provider")
	provider := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write(payload)
	}))
	t.Cleanup(provider.Close)

	gateway := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, provider.URL+"/signed", http.StatusFound)
	}))
	t.Cleanup(gateway.Close)

	var out []byte
	err := clientFor(t, gateway.URL, "").API.Files.GetPublicFile(
		context.Background(), publicFileID, "invoice.pdf", &out)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(out) != string(payload) {
		t.Errorf("bytes: got %q want %q", out, payload)
	}
}

// The two new fields the read side now answers with.
func TestAListingCarriesTheNewPublicFields(t *testing.T) {
	body := `{
	  "list": {"items": [{
	     "path": "invoices/invoice.pdf",
	     "isPublic": true,
	     "publicUrl": "https://api.norbix.io/v3/files/public/nbpf_7hK2abc/invoice.pdf"}]},
	  "folders": ["invoices", "drafts"],
	  "publicFolders": [{"path":"invoices","publicId":"nbpf_folder1","inherited":false}]
	}`
	c, _ := newFilesClient(t, body)

	var out struct {
		List struct {
			Items []struct {
				Path      string `json:"path"`
				IsPublic  bool   `json:"isPublic"`
				PublicURL string `json:"publicUrl"`
			} `json:"items"`
		} `json:"list"`
		Folders       []string `json:"folders"`
		PublicFolders []struct {
			Path      string `json:"path"`
			PublicID  string `json:"publicId"`
			Inherited bool   `json:"inherited"`
		} `json:"publicFolders"`
	}
	if err := c.API.Files.ListFiles(context.Background(), filesIntegrationID, nil, &out); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !out.List.Items[0].IsPublic {
		t.Error("isPublic did not survive the round trip")
	}
	if out.List.Items[0].PublicURL == "" {
		t.Error("publicUrl did not survive the round trip")
	}
	if len(out.PublicFolders) != 1 || out.PublicFolders[0].PublicID != "nbpf_folder1" {
		t.Errorf("publicFolders: got %v", out.PublicFolders)
	}
	// The plain folder list is untouched, so an older client keeps working.
	if len(out.Folders) != 2 {
		t.Errorf("folders: got %v", out.Folders)
	}
}

// Guard against the bug this slice had to fix in the transport: a file read
// into a plain map would have been parsed as JSON and failed.
func TestBytesAreNotParsedAsJson(t *testing.T) {
	srv, _, _ := newPublicFileServer(t, []byte("\x89PNG\r\n\x1a\n not json at all"), http.StatusOK)

	var out []byte
	if err := clientFor(t, srv.URL, "").API.Files.GetPublicFile(
		context.Background(), publicFileID, "logo.png", &out); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(out) == 0 {
		t.Fatal("no bytes came back")
	}

	// And the proof that it would have failed the other way round:
	var asJSON map[string]any
	err := clientFor(t, srv.URL, "").API.Files.GetPublicFile(
		context.Background(), publicFileID, "logo.png", &asJSON)
	if err == nil {
		t.Error("a PNG decoded into a map without complaining")
	}
}
