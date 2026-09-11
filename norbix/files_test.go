package norbix

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Files endpoints, one case per endpoint: 12 on the hub (the dashboard API)
// and 8 on the public API.
//
// Every case starts its own fake HTTP server and its own client, so the order
// the tests run in does not matter and no real storage provider is contacted.
// Each case checks the three things that can silently break: the method is on
// the module at all, the request goes to the right path with the right verb,
// and the values the caller passed end up in the query string or the body.

const filesIntegrationID = "33333333-3333-3333-3333-333333333333"

// recorded is what the fake server saw.
type recorded struct {
	method string
	path   string
	query  string
	body   map[string]any
}

// newFilesClient starts a fake gateway and returns a client pointed at it,
// together with the record of the last request it received.
func newFilesClient(t *testing.T, responseBody string) (*Client, *recorded) {
	t.Helper()
	got := &recorded{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got.method = r.Method
		got.path = r.URL.Path
		got.query = r.URL.RawQuery
		if raw, _ := io.ReadAll(r.Body); len(raw) > 0 {
			_ = json.Unmarshal(raw, &got.body)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(responseBody))
	}))
	t.Cleanup(srv.Close)

	c, err := New(Options{
		ProjectID:  "proj_1",
		APIKey:     "key_1",
		BaseURLAPI: srv.URL,
		BaseURLHub: srv.URL,
		APIVersion: "v2",
		HubVersion: "v2",
	})
	if err != nil {
		t.Fatalf("client: %v", err)
	}
	return c, got
}

func expectCall(t *testing.T, got *recorded, wantMethod, wantPath string) {
	t.Helper()
	if got.method != wantMethod {
		t.Errorf("method: got %q want %q", got.method, wantMethod)
	}
	if got.path != wantPath {
		t.Errorf("path: got %q want %q", got.path, wantPath)
	}
}

// ---- hub: the module itself -------------------------------------------------

func TestHubEnableFiles(t *testing.T) {
	c, got := newFilesClient(t, `{}`)
	var out map[string]any
	if err := c.Hub.Files.EnableFiles(context.Background(), nil, &out); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expectCall(t, got, http.MethodGet, "/v2/files/enable")
}

func TestHubDisableFiles(t *testing.T) {
	c, got := newFilesClient(t, `{}`)
	var out map[string]any
	if err := c.Hub.Files.DisableFiles(context.Background(), nil, &out); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expectCall(t, got, http.MethodGet, "/v2/files/disable")
}

// ---- hub: browsing ----------------------------------------------------------

func TestHubGetFolderFiles(t *testing.T) {
	c, got := newFilesClient(t, `{"folders":["invoices/"]}`)
	var out struct {
		Folders []string `json:"folders"`
	}
	err := c.Hub.Files.GetFolderFiles(context.Background(), map[string]any{
		"filesIntegrationId": filesIntegrationID,
		"path":               "invoices/",
	}, &out)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expectCall(t, got, http.MethodGet, "/v2/files/folder")
	if len(out.Folders) != 1 || out.Folders[0] != "invoices/" {
		t.Errorf("folders: got %v", out.Folders)
	}
}

func TestHubGetFile(t *testing.T) {
	c, got := newFilesClient(t, `{"file":{"path":"invoices/invoice.pdf"},"isPublic":false}`)
	var out struct {
		File struct {
			Path string `json:"path"`
		} `json:"file"`
		IsPublic bool `json:"isPublic"`
	}
	err := c.Hub.Files.GetFile(context.Background(), map[string]any{
		"filesIntegrationId": filesIntegrationID,
		"path":               "invoices/invoice.pdf",
	}, &out)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expectCall(t, got, http.MethodGet, "/v2/files/item")
	if out.File.Path != "invoices/invoice.pdf" {
		t.Errorf("file path: got %q", out.File.Path)
	}
}

// ---- hub: integrations ------------------------------------------------------

func TestHubGetFilesIntegrations(t *testing.T) {
	c, got := newFilesClient(t, `{"defaultIntegrationId":"`+filesIntegrationID+`"}`)
	var out struct {
		DefaultIntegrationID string `json:"defaultIntegrationId"`
	}
	if err := c.Hub.Files.GetFilesIntegrations(context.Background(), nil, &out); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expectCall(t, got, http.MethodGet, "/v2/files/integrations")
	if out.DefaultIntegrationID != filesIntegrationID {
		t.Errorf("default integration: got %q", out.DefaultIntegrationID)
	}
}

func TestHubGetFilesIntegration(t *testing.T) {
	c, got := newFilesClient(t, `{"item":{}}`)
	var out map[string]any
	err := c.Hub.Files.GetFilesIntegration(context.Background(), filesIntegrationID, nil, &out)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expectCall(t, got, http.MethodGet, "/v2/files/integrations/"+filesIntegrationID)
}

func TestHubSaveFilesIntegration(t *testing.T) {
	c, got := newFilesClient(t, `{"id":"`+filesIntegrationID+`"}`)
	var out map[string]any
	err := c.Hub.Files.SaveFilesIntegration(context.Background(), map[string]any{
		"integration": map[string]any{
			"integrationName": "Invoices bucket",
			"bucketName":      "norbix-invoices",
			"isEnabled":       true,
		},
	}, &out)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expectCall(t, got, http.MethodPost, "/v2/files/integrations")
	integration, ok := got.body["integration"].(map[string]any)
	if !ok || integration["bucketName"] != "norbix-invoices" {
		t.Errorf("body did not carry the integration settings: %v", got.body)
	}
}

func TestHubDeleteFilesIntegration(t *testing.T) {
	c, got := newFilesClient(t, `{}`)
	var out map[string]any
	err := c.Hub.Files.DeleteFilesIntegration(context.Background(), filesIntegrationID, nil, &out)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expectCall(t, got, http.MethodDelete, "/v2/files/integrations/"+filesIntegrationID)
}

func TestHubEnableFilesIntegration(t *testing.T) {
	c, got := newFilesClient(t, `{}`)
	var out map[string]any
	err := c.Hub.Files.EnableFilesIntegration(context.Background(), filesIntegrationID, nil, &out)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expectCall(t, got, http.MethodPut, "/v2/files/integrations/"+filesIntegrationID+"/enable")
}

func TestHubDisableFilesIntegration(t *testing.T) {
	c, got := newFilesClient(t, `{}`)
	var out map[string]any
	err := c.Hub.Files.DisableFilesIntegration(context.Background(), filesIntegrationID, nil, &out)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expectCall(t, got, http.MethodPut, "/v2/files/integrations/"+filesIntegrationID+"/disable")
}

func TestHubSetFilesIntegrationAsDefault(t *testing.T) {
	c, got := newFilesClient(t, `{}`)
	var out map[string]any
	err := c.Hub.Files.SetFilesIntegrationAsDefault(context.Background(), filesIntegrationID, nil, &out)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expectCall(t, got, http.MethodPut, "/v2/files/integrations/"+filesIntegrationID+"/default")
}

func TestHubTestFilesIntegration(t *testing.T) {
	c, got := newFilesClient(t, `{"items":[{"operation":"ListFiles","result":"Ok"}]}`)
	var out struct {
		Items []struct {
			Operation string `json:"operation"`
			Result    string `json:"result"`
		} `json:"items"`
	}
	err := c.Hub.Files.TestFilesIntegration(context.Background(), map[string]any{
		"integrationId": filesIntegrationID,
	}, &out)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expectCall(t, got, http.MethodPost, "/v2/files/integrations/test")
	if got.body["integrationId"] != filesIntegrationID {
		t.Errorf("body did not carry the integration id: %v", got.body)
	}
	if len(out.Items) != 1 || out.Items[0].Result != "Ok" {
		t.Errorf("results: got %v", out.Items)
	}
}

// ---- public API -------------------------------------------------------------

func TestAPIListFiles(t *testing.T) {
	c, got := newFilesClient(t, `{"folders":["invoices/"]}`)
	var out map[string]any
	err := c.API.Files.ListFiles(context.Background(), filesIntegrationID, map[string]any{
		"path": "invoices/",
	}, &out)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expectCall(t, got, http.MethodGet, "/v2/files/"+filesIntegrationID)
}

func TestAPIGetFileInfo(t *testing.T) {
	c, got := newFilesClient(t, `{"file":{"path":"invoices/invoice.pdf"}}`)
	var out map[string]any
	err := c.API.Files.GetFileInfo(context.Background(), filesIntegrationID, map[string]any{
		"path": "invoices/invoice.pdf",
	}, &out)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expectCall(t, got, http.MethodGet, "/v2/files/"+filesIntegrationID+"/info")
}

func TestAPIGetSignedURL(t *testing.T) {
	c, got := newFilesClient(t, `{"url":"https://storage.example/signed"}`)
	var out struct {
		URL string `json:"url"`
	}
	err := c.API.Files.GetSignedUrl(context.Background(), filesIntegrationID, map[string]any{
		"path": "invoices/invoice.pdf",
	}, &out)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expectCall(t, got, http.MethodGet, "/v2/files/"+filesIntegrationID+"/sign")
	if out.URL != "https://storage.example/signed" {
		t.Errorf("url: got %q", out.URL)
	}
}

func TestAPIRequestUploadURL(t *testing.T) {
	c, got := newFilesClient(t, `{"url":"https://storage.example/put"}`)
	var out struct {
		URL string `json:"url"`
	}
	err := c.API.Files.RequestUploadUrl(context.Background(), filesIntegrationID, map[string]any{
		"path":        "invoices/invoice.pdf",
		"contentType": "application/pdf",
	}, &out)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expectCall(t, got, http.MethodPost, "/v2/files/"+filesIntegrationID+"/upload-url")
	if got.body["contentType"] != "application/pdf" {
		t.Errorf("body did not carry the content type: %v", got.body)
	}
}

func TestAPICommitUpload(t *testing.T) {
	c, got := newFilesClient(t, `{}`)
	var out map[string]any
	err := c.API.Files.CommitUpload(context.Background(), filesIntegrationID, map[string]any{
		"path":        "invoices/invoice.pdf",
		"contentType": "application/pdf",
		"sizeBytes":   1024,
	}, &out)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expectCall(t, got, http.MethodPost, "/v2/files/"+filesIntegrationID+"/commit")
	if got.body["path"] != "invoices/invoice.pdf" {
		t.Errorf("body did not carry the path: %v", got.body)
	}
}

func TestAPIDownloadFile(t *testing.T) {
	c, got := newFilesClient(t, `{}`)
	var out map[string]any
	err := c.API.Files.DownloadFileApi(context.Background(), filesIntegrationID, map[string]any{
		"path": "invoices/invoice.pdf",
	}, &out)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expectCall(t, got, http.MethodGet, "/v2/files/"+filesIntegrationID+"/download")
}

func TestAPIDeleteFile(t *testing.T) {
	c, got := newFilesClient(t, `{}`)
	var out map[string]any
	err := c.API.Files.DeleteFileApi(context.Background(), filesIntegrationID, map[string]any{
		"path": "invoices/invoice.pdf",
	}, &out)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expectCall(t, got, http.MethodDelete, "/v2/files/"+filesIntegrationID)
}

func TestAPIDeleteManyFiles(t *testing.T) {
	c, got := newFilesClient(t, `{}`)
	var out map[string]any
	err := c.API.Files.DeleteManyFilesApi(context.Background(), filesIntegrationID, map[string]any{
		"paths": []string{"invoices/a.pdf", "invoices/b.pdf"},
	}, &out)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expectCall(t, got, http.MethodDelete, "/v2/files/"+filesIntegrationID+"/bulk")
}
