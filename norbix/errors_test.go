package norbix

import (
	"context"
	stderrors "errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/norbix-code/sdk-go/norbix/api/dtos"
	norbixerr "github.com/norbix-code/sdk-go/norbix/errors"
)

// What the caller sees when a call fails (10b-files slice ERRORS, #66 and #67).
//
// Two rules are pinned here. First, the message and the error code are the
// gateway's own: the gateway puts them inside responseStatus.errors[], so
// reading the top of that block gave every caller "Request failed" and
// "HTTP_404" — that was #66. Second, a call fails when the gateway says it
// failed, even with HTTP 200 and responseStatus.isSuccess=false — that was #67.
//
// Every case starts its own fake gateway and its own client, so the order the
// tests run in does not matter and no real server is contacted.

// newAnsweringServer answers every request with one body and one status.
func newAnsweringServer(t *testing.T, status int, body string, contentType string) *Client {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", contentType)
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
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
	return c
}

// aCall makes any call; the transport is the same for all of them.
func aCall(c *Client) error {
	var out dtos.TestFilesIntegrationResponse
	return c.API.Files.TestFilesIntegration(context.Background(), "int_7", nil, &out)
}

func baseError(t *testing.T, err error) *norbixerr.Error {
	t.Helper()
	if err == nil {
		t.Fatal("expected an error, got none")
	}
	var base *norbixerr.Error
	if !stderrors.As(err, &base) {
		t.Fatalf("error type: got %T want *errors.Error", err)
	}
	return base
}

// (a) HTTP 400 with two errors inside responseStatus.errors.
func TestA400TakesMessageAndCodeFromTheFirstErrorAndKeepsThemAll(t *testing.T) {
	c := newAnsweringServer(t, http.StatusBadRequest, `{"responseStatus":{"isSuccess":false,"errors":[
		{"message":"File name is required","errorCode":"CM-ERRORS-FILES-002","fieldName":"fileName"},
		{"message":"Folder does not exist","errorCode":"CM-ERRORS-FILES-016","context":{"Provider":"Local"}}
	]}}`, "application/json")

	base := baseError(t, aCall(c))

	if base.Message != "File name is required" {
		t.Errorf("message: got %q", base.Message)
	}
	if base.Code != "CM-ERRORS-FILES-002" {
		t.Errorf("code: got %q", base.Code)
	}
	if base.HTTPStatus() != http.StatusBadRequest {
		t.Errorf("status: got %d want 400", base.HTTPStatus())
	}
	if len(base.Errors) != 2 {
		t.Fatalf("errors: got %d want 2", len(base.Errors))
	}
	if base.Errors[0].FieldName != "fileName" {
		t.Errorf("first error fieldName: got %q", base.Errors[0].FieldName)
	}
	if base.Errors[1].ErrorCode != "CM-ERRORS-FILES-016" {
		t.Errorf("second error code: got %q", base.Errors[1].ErrorCode)
	}
	if base.Errors[1].Context["Provider"] != "Local" {
		t.Errorf("second error context: got %v", base.Errors[1].Context)
	}
	if len(base.Body) == 0 {
		t.Error("the raw body was not kept")
	}
	// A 400 stays the SDK's validation error type.
	var validation *norbixerr.ValidationError
	if !stderrors.As(aCall(c), &validation) {
		t.Error("a 400 should still be a *errors.ValidationError")
	}
}

// (b) HTTP 200 whose body says the call failed.
func TestA200ThatSaysItFailedIsAnError(t *testing.T) {
	c := newAnsweringServer(t, http.StatusOK, `{"responseStatus":{"isSuccess":false,"errors":[
		{"message":"File not found: \"a/b.txt\" does not exist in Local (int_7).","errorCode":"CM-ERRORS-FILES-016"}
	]}}`, "application/json")

	base := baseError(t, aCall(c))

	if base.HTTPStatus() != http.StatusOK {
		t.Errorf("status: got %d want 200", base.HTTPStatus())
	}
	if base.Message != `File not found: "a/b.txt" does not exist in Local (int_7).` {
		t.Errorf("message: got %q", base.Message)
	}
	if base.Code != "CM-ERRORS-FILES-016" {
		t.Errorf("code: got %q", base.Code)
	}
}

// (c) HTTP 200 that says the call worked — unchanged.
func TestA200ThatSaysItWorkedStillComesBackAsAValue(t *testing.T) {
	c := newAnsweringServer(t, http.StatusOK,
		`{"items":[{"operation":"UploadFile","result":"OK"}],"responseStatus":{"isSuccess":true}}`,
		"application/json")

	var out dtos.TestFilesIntegrationResponse
	if err := c.API.Files.TestFilesIntegration(context.Background(), "int_7", nil, &out); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(out.Items) != 1 || out.Items[0].Result != "OK" {
		t.Errorf("items: got %+v", out.Items)
	}
}

func TestA200WithNoResponseStatusStillComesBackAsAValue(t *testing.T) {
	c := newAnsweringServer(t, http.StatusOK, `{"items":[]}`, "application/json")

	if err := aCall(c); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

// (d) a 500 whose body is not JSON at all.
func TestA500WithABodyThatIsNotJSONUsesTheFallbackText(t *testing.T) {
	c := newAnsweringServer(t, http.StatusInternalServerError, "<html>Bad Gateway</html>", "text/html")

	base := baseError(t, aCall(c))

	if base.Message != "Request failed (HTTP 500)" {
		t.Errorf("message: got %q", base.Message)
	}
	if base.Code != "HTTP_500" {
		t.Errorf("code: got %q want the HTTP_500 fallback", base.Code)
	}
	if string(base.Body) != "<html>Bad Gateway</html>" {
		t.Errorf("body: got %q", string(base.Body))
	}
}

func TestAnEmptyErrorBodyUsesTheFallbackTextToo(t *testing.T) {
	c := newAnsweringServer(t, http.StatusNotFound, `{}`, "application/json")

	base := baseError(t, aCall(c))

	if base.Message != "Request failed (HTTP 404)" {
		t.Errorf("message: got %q", base.Message)
	}
}

func TestReadsTheTopOfTheBodyWhenThereIsNoResponseStatus(t *testing.T) {
	c := newAnsweringServer(t, http.StatusConflict,
		`{"message":"Already exists","errorCode":"CM-ERRORS-FILES-009"}`, "application/json")

	base := baseError(t, aCall(c))

	if base.Message != "Already exists" {
		t.Errorf("message: got %q", base.Message)
	}
	if base.Code != "CM-ERRORS-FILES-009" {
		t.Errorf("code: got %q", base.Code)
	}
}
