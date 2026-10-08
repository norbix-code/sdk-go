package hub

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/norbix-code/sdk-go/v2/norbix/hub/dtos"
)

// Trigger "order" and "break on failure": the gateway added `order` (whole
// number, lower runs first) and `breakOnError` (stop the later triggers of the
// same event when this one fails) to every trigger. The save sends the
// caller's map, so both reach the body as given — order 0 included — and the
// typed read DTO decodes both.

func TestSaveSchemaTriggerSendsOrderAndBreakOnError(t *testing.T) {
	var body map[string]map[string]any
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		raw, _ := io.ReadAll(r.Body)
		_ = json.Unmarshal(raw, &body)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	err := newHubDatabaseTestModule(srv.URL).SaveSchemaTrigger(context.Background(), map[string]any{
		"trigger": map[string]any{"type": "Schema", "schemaId": "sch_1", "name": "first", "order": 0, "breakOnError": true},
	}, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := body["trigger"]
	if got["order"] != float64(0) || got["breakOnError"] != true {
		t.Errorf("trigger body: got order %v breakOnError %v, want 0 / true", got["order"], got["breakOnError"])
	}
}

func TestSchemaTriggerReadDecodesOrderAndBreakOnError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"trigger":{"viewId":"trg_1","name":"first","schemaId":"sch_1","order":2,"breakOnError":true}}`))
	}))
	defer srv.Close()

	var one dtos.GetSchemaTriggerResponse
	if err := newHubDatabaseTestModule(srv.URL).GetSchemaTrigger(context.Background(), "trg_1", nil, &one); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if one.Trigger == nil || one.Trigger.Order != 2 || !one.Trigger.BreakOnError {
		t.Errorf("trigger: got %+v, want order 2 and breakOnError true", one.Trigger)
	}
}
