package transport

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	norbixerr "github.com/norbix-code/sdk-go/norbix/errors"
)

func newTestTransport(base string) *Transport {
	return New(&Config{
		ProjectID:  "proj_1",
		APIKey:     "key_1",
		BaseURLAPI: base,
		BaseURLHub: base,
		APIVersion: "v2",
		HubVersion: "v2",
		Timeout:    5 * time.Second,
		Env:        "PROD",
	}, nil)
}

func TestSendInjectsHeadersAndVersion(t *testing.T) {
	var gotPath, gotAuth, gotProject string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		gotAuth = r.Header.Get("Authorization")
		gotProject = r.Header.Get("X-CM-ProjectId")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	defer srv.Close()

	tr := newTestTransport(srv.URL)
	var out map[string]any
	err := tr.Send(context.Background(), Request{
		Target: TargetAPI,
		Path:   "/{version}/echo",
		Method: http.MethodGet,
		Scope:  ScopeProject,
	}, &out)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotPath != "/v2/echo" {
		t.Errorf("path: got %q want /v2/echo", gotPath)
	}
	if gotAuth != "Bearer key_1" {
		t.Errorf("auth: got %q", gotAuth)
	}
	if gotProject != "proj_1" {
		t.Errorf("project header: got %q", gotProject)
	}
	if out["ok"] != true {
		t.Errorf("body decode: got %v", out)
	}
}

func TestSendPathParamsAndBody(t *testing.T) {
	var gotPath string
	var gotBody map[string]any
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		b, _ := io.ReadAll(r.Body)
		_ = json.Unmarshal(b, &gotBody)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	tr := newTestTransport(srv.URL)
	err := tr.Send(context.Background(), Request{
		Target:     TargetHub,
		Path:       "/{version}/database/collections/{collectionName}",
		Method:     http.MethodPost,
		PathParams: map[string]string{"collectionName": "orders"},
		Body:       map[string]any{"name": "n1", "skip": nil},
		Scope:      ScopeProject,
	}, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotPath != "/v2/database/collections/orders" {
		t.Errorf("path: got %q", gotPath)
	}
	if gotBody["name"] != "n1" {
		t.Errorf("body name: got %v", gotBody)
	}
	if _, present := gotBody["skip"]; present {
		t.Errorf("nil body fields must be dropped, got %v", gotBody)
	}
}

func TestMissingPathParam(t *testing.T) {
	tr := newTestTransport("http://example.invalid")
	err := tr.Send(context.Background(), Request{
		Target: TargetAPI,
		Path:   "/{version}/x/{id}",
		Method: http.MethodGet,
		Scope:  ScopeProject,
	}, nil)
	if err == nil {
		t.Fatal("expected error for missing path param")
	}
	var ne *norbixerr.Error
	if !asNorbix(err, &ne) || ne.Code != norbixerr.CodeMissingPathParam {
		t.Errorf("expected CodeMissingPathParam, got %v", err)
	}
}

func TestAccountScopeRequires(t *testing.T) {
	tr := newTestTransport("http://example.invalid")
	err := tr.Send(context.Background(), Request{
		Target: TargetHub, Path: "/{version}/account/profile", Method: http.MethodGet, Scope: ScopeAccount,
	}, nil)
	if err == nil {
		t.Fatal("expected account scope error")
	}
}

func TestErrorMapping404(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"message":"nope","errorCode":"X"}`))
	}))
	defer srv.Close()
	tr := newTestTransport(srv.URL)
	err := tr.Send(context.Background(), Request{Target: TargetAPI, Path: "/{version}/x", Method: http.MethodGet, Scope: ScopeProject}, nil)
	var nf *norbixerr.NotFoundError
	if !asNorbix(err, &nf) {
		t.Fatalf("expected *NotFoundError, got %T: %v", err, err)
	}
	if nf.Base.Status != 404 {
		t.Errorf("status: got %d", nf.Base.Status)
	}
}

func TestRetriesOn503ForGET(t *testing.T) {
	var calls int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := atomic.AddInt32(&calls, 1)
		if n < 2 {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	defer srv.Close()
	tr := newTestTransport(srv.URL)
	var out map[string]any
	if err := tr.Send(context.Background(), Request{Target: TargetAPI, Path: "/{version}/x", Method: http.MethodGet, Scope: ScopeProject}, &out); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if atomic.LoadInt32(&calls) < 2 {
		t.Errorf("expected a retry, calls=%d", calls)
	}
}

func TestNoRetryOnPOST(t *testing.T) {
	var calls int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&calls, 1)
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer srv.Close()
	tr := newTestTransport(srv.URL)
	_ = tr.Send(context.Background(), Request{Target: TargetAPI, Path: "/{version}/x", Method: http.MethodPost, Scope: ScopeProject, Body: map[string]any{"a": 1}}, nil)
	if atomic.LoadInt32(&calls) != 1 {
		t.Errorf("POST must not retry, calls=%d", calls)
	}
}

// asNorbix is a tiny errors.As shim usable for both base and typed errors.
func asNorbix[T error](err error, target *T) bool {
	for err != nil {
		if t, ok := err.(T); ok {
			*target = t
			return true
		}
		type unwrapper interface{ Unwrap() error }
		if u, ok := err.(unwrapper); ok {
			err = u.Unwrap()
		} else {
			break
		}
	}
	return false
}
