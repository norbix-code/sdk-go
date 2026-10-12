package norbix

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	norbixerr "github.com/norbix-code/sdk-go/v2/norbix/errors"
	"github.com/norbix-code/sdk-go/v2/norbix/hub/dtos"
)

// ParseFunctionEvent reads the envelope a cloud function receives (gateway
// CodeFunctionEventDto, written camelCase with camelCase enums and nulls
// left out) and the typed views decode its slots into the caller's types.
// One sample per source below, shaped like GET /code/functions/event-sample
// answers.
//
// Nothing leaves the process: the client test sends the run-as token to a
// local test server.

const functionEventContext = `"context": {
    "accountId": "acc_1",
    "projectId": "proj_1",
    "environment": "TEST",
    "region": "nb-eu-germany",
    "hubUrl": "https://hub.example.test",
    "apiUrl": "https://api.example.test",
    "echoUrl": "https://api.example.test/v3/echo",
    "triggerId": "trg_1",
    "correlationId": "corr-1",
    "caller": { "userId": "usr_caller_1", "email": "ada@example.test", "kind": "user" },
    "runAs": { "userId": "usr_service_1", "accessToken": "eyJ.run-as.token", "expiresAt": "2026-10-12T08:15:00Z" }
  }`

func functionEventSample(event, source, data, config string) []byte {
	body := `{
  "specVersion": "1",
  "id": "9f1c0000-0000-4000-8000-000000000001",
  "eventId": "7a2e0000-0000-4000-8000-000000000002",
  "event": "` + event + `",
  "source": "` + source + `",
  "createdOn": "2026-10-12T08:00:00Z",
  ` + functionEventContext + `,
  "data": ` + data
	if config != "" {
		body += `,
  "config": ` + config
	}
	return []byte(body + `
}`)
}

var (
	sampleDatabaseUpdated = functionEventSample("database.record.updated", "database", `{
    "database": {
      "collection": "orders",
      "schemaId": "sch_1",
      "schemaVersion": 3,
      "recordId": "6ac54a59212c29912eaddd81",
      "old": { "_id": "6ac54a59212c29912eaddd81", "status": "new", "total": 42.5 },
      "new": { "_id": "6ac54a59212c29912eaddd81", "status": "processing", "total": 42.5 },
      "truncated": false
    }
  }`, `{ "acceptStatus": "processing", "imageWidth": 568 }`)

	sampleDatabaseInserted = functionEventSample("database.record.inserted", "database", `{
    "database": {
      "collection": "orders",
      "schemaId": "sch_1",
      "schemaVersion": 3,
      "recordId": "6ac54a59212c29912eaddd82",
      "new": { "_id": "6ac54a59212c29912eaddd82", "status": "new", "total": 10 },
      "truncated": false
    }
  }`, "")

	sampleDatabaseTruncated = functionEventSample("database.record.replaced", "database", `{
    "database": {
      "collection": "orders",
      "schemaId": "sch_1",
      "recordId": "6ac54a59212c29912eaddd83",
      "truncated": true,
      "oldRef": { "kind": "document", "projectId": "proj_1", "integrationId": "int_db_1", "databaseIntegrationId": "int_db_1", "schemaId": "sch_1", "externalId": "6ac54a59212c29912eaddd83" },
      "newRef": { "kind": "document", "projectId": "proj_1", "integrationId": "int_db_1", "databaseIntegrationId": "int_db_1", "schemaId": "sch_1", "externalId": "6ac54a59212c29912eaddd83" }
    }
  }`, "")

	sampleMembershipUpdated = functionEventSample("membership.user.updated", "membership", `{
    "user": {
      "userId": "usr_member_1",
      "old": { "id": "usr_member_1", "email": "ada@example.test", "displayName": "Ada" },
      "new": { "id": "usr_member_1", "email": "ada@example.test", "displayName": "Ada Lovelace" }
    }
  }`, "")

	sampleFileUploaded = functionEventSample("files.file.uploaded", "files", `{
    "file": {
      "new": { "id": "fil_1", "fileName": "invoice.pdf", "contentType": "application/pdf" },
      "path": "invoices/2026/invoice.pdf"
    }
  }`, "")

	samplePaymentReceived = functionEventSample("payments.webhook.received", "payments", `{
    "payment": {
      "integrationId": "int_pay_1",
      "provider": "Stripe",
      "providerEventId": "evt_stripe_1",
      "eventName": "payment_intent.succeeded",
      "status": "succeeded",
      "amountMinor": 4250,
      "currency": "EUR",
      "externalIds": { "paymentIntent": "pi_1", "customer": "cus_1" },
      "occurredOn": "2026-10-12T07:59:58Z",
      "receivedOn": "2026-10-12T08:00:00Z"
    }
  }`, "")

	sampleAICreditsWarning = functionEventSample("ai.credits.warning", "ai", `{
    "ai": {
      "payload": { "creditsLeft": 120, "threshold": 0.1, "plan": "starter" },
      "targetUserId": "usr_member_1"
    }
  }`, "")

	sampleSchedulerFired = functionEventSample("scheduler.task.fired", "scheduler", `{
    "task": { "taskId": "tsk_1", "name": "Nightly payouts", "firedAt": "2026-10-12T03:00:00Z" },
    "payload": { "orderId": "6ac54a59212c29912eaddd81", "dryRun": false }
  }`, `{ "acceptStatus": "paid", "imageWidth": 0 }`)

	sampleManualInvoke = functionEventSample("manual.invoke", "manual", `{
    "payload": { "orderId": "6ac54a59212c29912eaddd81", "dryRun": true }
  }`, "")
)

type testOrder struct {
	ID     string  `json:"_id"`
	Status string  `json:"status"`
	Total  float64 `json:"total"`
}

type testPerson struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	DisplayName string `json:"displayName"`
}

type testPayout struct {
	OrderID string `json:"orderId"`
	DryRun  bool   `json:"dryRun"`
}

type testSettings struct {
	AcceptStatus string `json:"acceptStatus"`
	ImageWidth   int    `json:"imageWidth"`
}

func mustParseFunctionEvent(t *testing.T, raw []byte) *CodeFunctionEvent {
	t.Helper()
	ev, err := ParseFunctionEvent(raw)
	if err != nil {
		t.Fatalf("ParseFunctionEvent: %v", err)
	}
	return ev
}

func errorCode(err error) string {
	var e *norbixerr.Error
	if errors.As(err, &e) {
		return e.Code
	}
	return ""
}

func TestParseFunctionEventReadsTheEnvelope(t *testing.T) {
	ev := mustParseFunctionEvent(t, sampleDatabaseUpdated)

	if ev.SpecVersion != "1" || ev.ID != "9f1c0000-0000-4000-8000-000000000001" || ev.EventID != "7a2e0000-0000-4000-8000-000000000002" {
		t.Errorf("ids: %+v", ev)
	}
	if ev.Event != "database.record.updated" || ev.Source != "database" {
		t.Errorf("event / source: %q %q", ev.Event, ev.Source)
	}
	if !ev.CreatedOn.Equal(time.Date(2026, 10, 12, 8, 0, 0, 0, time.UTC)) {
		t.Errorf("createdOn: %v", ev.CreatedOn)
	}
	c := ev.Context
	if c.AccountID != "acc_1" || c.ProjectID != "proj_1" || c.Environment != "TEST" || c.Region != "nb-eu-germany" {
		t.Errorf("context ids: %+v", c)
	}
	if c.HubURL != "https://hub.example.test" || c.APIURL != "https://api.example.test" || c.EchoURL != "https://api.example.test/v3/echo" {
		t.Errorf("context urls: %+v", c)
	}
	if c.TriggerID != "trg_1" || c.TaskID != "" || c.CorrelationID != "corr-1" {
		t.Errorf("context trigger / task: %+v", c)
	}
	if c.Caller == nil || c.Caller.UserID != "usr_caller_1" || c.Caller.Email != "ada@example.test" || c.Caller.Kind != "user" {
		t.Errorf("caller: %+v", c.Caller)
	}
	if c.RunAs == nil || c.RunAs.UserID != "usr_service_1" || c.RunAs.AccessToken != "eyJ.run-as.token" || c.RunAs.ExpiresAt == nil {
		t.Errorf("runAs: %+v", c.RunAs)
	}
	if ev.Data.Database == nil || ev.Data.User != nil || ev.Data.File != nil || ev.Data.Payment != nil || ev.Data.AI != nil || ev.Data.Task != nil || ev.Data.Payload != nil {
		t.Errorf("slots: %+v", ev.Data)
	}
	var cfg testSettings
	if err := ev.DecodeConfig(&cfg); err != nil {
		t.Fatalf("DecodeConfig: %v", err)
	}
	if cfg.AcceptStatus != "processing" || cfg.ImageWidth != 568 {
		t.Errorf("config: %+v", cfg)
	}
	if string(ev.Raw()) != string(sampleDatabaseUpdated) {
		t.Error("Raw() must be the body as it arrived")
	}
}

func TestAsDatabaseRecordChangeDecodesOldAndNew(t *testing.T) {
	ev := mustParseFunctionEvent(t, sampleDatabaseUpdated)

	var old, next *testOrder
	change, err := ev.AsDatabaseRecordChange(&old, &next)
	if err != nil {
		t.Fatalf("AsDatabaseRecordChange: %v", err)
	}
	if change.Collection != "orders" || change.SchemaID != "sch_1" || change.SchemaVersion == nil || *change.SchemaVersion != 3 || change.RecordID != "6ac54a59212c29912eaddd81" {
		t.Errorf("change: %+v", change)
	}
	if change.Truncated || change.OldRef != nil || change.NewRef != nil {
		t.Errorf("not truncated: %+v", change)
	}
	if old == nil || old.Status != "new" || old.Total != 42.5 {
		t.Errorf("old: %+v", old)
	}
	if next == nil || next.Status != "processing" || next.ID != "6ac54a59212c29912eaddd81" {
		t.Errorf("new: %+v", next)
	}
	if change.Old == nil || change.New == nil {
		t.Error("the raw documents stay on the change")
	}
}

func TestAsDatabaseRecordChangeLeavesAnAbsentDocumentNil(t *testing.T) {
	ev := mustParseFunctionEvent(t, sampleDatabaseInserted)

	var old, next *testOrder
	change, err := ev.AsDatabaseRecordChange(&old, &next)
	if err != nil {
		t.Fatalf("AsDatabaseRecordChange: %v", err)
	}
	if old != nil {
		t.Errorf("insert has no old, got %+v", old)
	}
	if next == nil || next.Status != "new" || next.Total != 10 {
		t.Errorf("new: %+v", next)
	}
	if change.Old != nil || change.New == nil {
		t.Errorf("raw slots: old=%s new=%s", change.Old, change.New)
	}

	// nil outs are allowed: the caller only wants the metadata.
	if _, err := ev.AsDatabaseRecordChange(nil, nil); err != nil {
		t.Errorf("nil outs: %v", err)
	}
}

func TestAsDatabaseRecordChangeCarriesTheRefsWhenTruncated(t *testing.T) {
	ev := mustParseFunctionEvent(t, sampleDatabaseTruncated)

	var old, next *testOrder
	change, err := ev.AsDatabaseRecordChange(&old, &next)
	if err != nil {
		t.Fatalf("AsDatabaseRecordChange: %v", err)
	}
	if !change.Truncated || old != nil || next != nil || change.SchemaVersion != nil {
		t.Errorf("truncated change: %+v old=%v new=%v", change, old, next)
	}
	for name, ref := range map[string]*DocumentRef{"oldRef": change.OldRef, "newRef": change.NewRef} {
		if ref == nil || ref.Kind != "document" || ref.DatabaseIntegrationID != "int_db_1" || ref.SchemaID != "sch_1" || ref.ExternalID != "6ac54a59212c29912eaddd83" || ref.ProjectID != "proj_1" {
			t.Errorf("%s: %+v", name, ref)
		}
	}
}

func TestAsUserChangeDecodesTheUsers(t *testing.T) {
	ev := mustParseFunctionEvent(t, sampleMembershipUpdated)

	var old, next testPerson
	change, err := ev.AsUserChange(&old, &next)
	if err != nil {
		t.Fatalf("AsUserChange: %v", err)
	}
	if change.UserID != "usr_member_1" {
		t.Errorf("userId: %q", change.UserID)
	}
	if old.DisplayName != "Ada" || next.DisplayName != "Ada Lovelace" || next.Email != "ada@example.test" {
		t.Errorf("old=%+v new=%+v", old, next)
	}
}

func TestAsFileChangeDecodesTheFile(t *testing.T) {
	ev := mustParseFunctionEvent(t, sampleFileUploaded)

	var old, next map[string]string
	change, err := ev.AsFileChange(&old, &next)
	if err != nil {
		t.Fatalf("AsFileChange: %v", err)
	}
	if change.Path != "invoices/2026/invoice.pdf" || change.Old != nil {
		t.Errorf("change: %+v", change)
	}
	if old != nil || next["fileName"] != "invoice.pdf" || next["contentType"] != "application/pdf" {
		t.Errorf("old=%v new=%v", old, next)
	}
}

func TestAsPaymentReadsTheWebhookFact(t *testing.T) {
	ev := mustParseFunctionEvent(t, samplePaymentReceived)

	p, err := ev.AsPayment()
	if err != nil {
		t.Fatalf("AsPayment: %v", err)
	}
	if p.IntegrationID != "int_pay_1" || p.Provider != "Stripe" || p.ProviderEventID != "evt_stripe_1" || p.EventName != "payment_intent.succeeded" || p.Status != "succeeded" {
		t.Errorf("payment: %+v", p)
	}
	if p.AmountMinor == nil || *p.AmountMinor != 4250 || p.Currency != "EUR" {
		t.Errorf("amount: %+v", p)
	}
	if p.ExternalIDs["paymentIntent"] != "pi_1" || p.ExternalIDs["customer"] != "cus_1" {
		t.Errorf("externalIds: %v", p.ExternalIDs)
	}
	if p.OccurredOn == nil || !p.OccurredOn.Equal(time.Date(2026, 10, 12, 7, 59, 58, 0, time.UTC)) || !p.ReceivedOn.Equal(time.Date(2026, 10, 12, 8, 0, 0, 0, time.UTC)) {
		t.Errorf("times: %v %v", p.OccurredOn, p.ReceivedOn)
	}
}

func TestAsAIUsageDecodesThePayload(t *testing.T) {
	ev := mustParseFunctionEvent(t, sampleAICreditsWarning)

	var payload map[string]any
	usage, err := ev.AsAIUsage(&payload)
	if err != nil {
		t.Fatalf("AsAIUsage: %v", err)
	}
	if usage.TargetUserID != "usr_member_1" || payload["plan"] != "starter" || payload["creditsLeft"] != float64(120) {
		t.Errorf("usage=%+v payload=%v", usage, payload)
	}
}

func TestAsScheduledTaskAndPayload(t *testing.T) {
	ev := mustParseFunctionEvent(t, sampleSchedulerFired)

	task, err := ev.AsScheduledTask()
	if err != nil {
		t.Fatalf("AsScheduledTask: %v", err)
	}
	if task.TaskID != "tsk_1" || task.Name != "Nightly payouts" || !task.FiredAt.Equal(time.Date(2026, 10, 12, 3, 0, 0, 0, time.UTC)) {
		t.Errorf("task: %+v", task)
	}
	var payout testPayout
	if err := ev.Payload(&payout); err != nil {
		t.Fatalf("Payload: %v", err)
	}
	if payout.OrderID != "6ac54a59212c29912eaddd81" || payout.DryRun {
		t.Errorf("payload: %+v", payout)
	}
}

func TestManualInvokeHasAPayloadAndNoConfig(t *testing.T) {
	ev := mustParseFunctionEvent(t, sampleManualInvoke)

	var payout testPayout
	if err := ev.Payload(&payout); err != nil {
		t.Fatalf("Payload: %v", err)
	}
	if !payout.DryRun || payout.OrderID != "6ac54a59212c29912eaddd81" {
		t.Errorf("payload: %+v", payout)
	}
	if ev.Config != nil {
		t.Errorf("config must be nil when the action has none, got %s", ev.Config)
	}
	cfg := testSettings{AcceptStatus: "untouched"}
	if err := ev.DecodeConfig(&cfg); err != nil || cfg.AcceptStatus != "untouched" {
		t.Errorf("DecodeConfig without config: err=%v cfg=%+v", err, cfg)
	}
}

func TestAViewOfAnotherSourceNamesTheSlotAndTheEvent(t *testing.T) {
	ev := mustParseFunctionEvent(t, sampleMembershipUpdated)

	_, err := ev.AsDatabaseRecordChange(nil, nil)
	if errorCode(err) != norbixerr.CodeFunctionEventNoSlot {
		t.Fatalf("want %s, got %v", norbixerr.CodeFunctionEventNoSlot, err)
	}
	if !strings.Contains(err.Error(), `"membership.user.updated"`) || !strings.Contains(err.Error(), "data.database") {
		t.Errorf("message must name the event and the slot: %v", err)
	}
}

func TestADocumentThatDoesNotDecodeIsAnError(t *testing.T) {
	ev := mustParseFunctionEvent(t, sampleDatabaseUpdated)

	var wrong *int
	_, err := ev.AsDatabaseRecordChange(&wrong, nil)
	if errorCode(err) != norbixerr.CodeFunctionEventInvalid || !strings.Contains(err.Error(), "data.database.old") {
		t.Errorf("want %s naming data.database.old, got %v", norbixerr.CodeFunctionEventInvalid, err)
	}
}

func TestANewerSpecVersionIsRefused(t *testing.T) {
	raw := strings.Replace(string(sampleDatabaseUpdated), `"specVersion": "1"`, `"specVersion": "2"`, 1)

	_, err := ParseFunctionEvent([]byte(raw))
	if errorCode(err) != norbixerr.CodeFunctionEventUnsupported {
		t.Fatalf("want %s, got %v", norbixerr.CodeFunctionEventUnsupported, err)
	}
	if !strings.Contains(err.Error(), `specVersion "2"`) {
		t.Errorf("message must name the version: %v", err)
	}
}

func TestABodyThatIsNotAnEnvelopeIsRefused(t *testing.T) {
	for name, raw := range map[string]string{
		"not json":       `not json`,
		"no specVersion": `{"event":"x"}`,
		"an array":       `[1,2]`,
		"no event":       `{"specVersion":"1","data":{}}`,
	} {
		_, err := ParseFunctionEvent([]byte(raw))
		if errorCode(err) != norbixerr.CodeFunctionEventInvalid {
			t.Errorf("%s: want %s, got %v", name, norbixerr.CodeFunctionEventInvalid, err)
		}
	}
}

func TestClientOptionsSendTheRunAsTokenToTheEventHubURL(t *testing.T) {
	var seen http.Header
	var seenPath string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		seen = r.Header.Clone()
		seenPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"version":"test"}`))
	}))
	defer srv.Close()

	raw := strings.Replace(string(sampleDatabaseUpdated), `"hubUrl": "https://hub.example.test"`, `"hubUrl": "`+srv.URL+`"`, 1)
	ev := mustParseFunctionEvent(t, []byte(raw))

	opts := ev.ClientOptions()
	if opts.BearerToken != "eyJ.run-as.token" || opts.ProjectID != "proj_1" || opts.AccountID != "acc_1" || opts.Env != "TEST" || opts.Region != "nb-eu-germany" {
		t.Errorf("options: %+v", opts)
	}
	if opts.BaseURLHub != srv.URL || opts.BaseURLAPI != "https://api.example.test" || opts.APIKey != "" {
		t.Errorf("urls: %+v", opts)
	}

	c, err := New(opts)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	defer c.Close()
	var out map[string]any
	if err := c.Hub.Echo.Echo(context.Background(), nil, &out); err != nil {
		t.Fatalf("Echo: %v", err)
	}
	if seenPath != "/v3/echo" {
		t.Errorf("path: %q", seenPath)
	}
	if got := seen.Get("Authorization"); got != "Bearer eyJ.run-as.token" {
		t.Errorf("Authorization: %q", got)
	}
	if seen.Get("norbix-project-id") != "proj_1" || seen.Get("norbix-account-id") != "acc_1" || seen.Get("norbix-env") != "TEST" || seen.Get("nb-region") != "nb-eu-germany" {
		t.Errorf("headers: %v", seen)
	}
}

func TestClientOptionsWithoutARunAsTokenHaveNoBearer(t *testing.T) {
	raw := strings.Replace(string(samplePaymentReceived),
		`"runAs": { "userId": "usr_service_1", "accessToken": "eyJ.run-as.token", "expiresAt": "2026-10-12T08:15:00Z" }`,
		`"runAs": { "userId": "usr_service_1" }`, 1)
	ev := mustParseFunctionEvent(t, []byte(raw))

	if opts := ev.ClientOptions(); opts.BearerToken != "" || opts.ProjectID != "proj_1" {
		t.Errorf("options: %+v", opts)
	}
}

// The generated envelope type decodes a real envelope: the dictionary
// slots (old / new documents, payload, config) hold scalar values, so they
// are map[string]any, not map[string]map[string]any — the earlier
// generator output refused `{"status":"paid"}` with "cannot unmarshal string
// into Go struct field ... of type map[string]interface {}".
func TestGeneratedEnvelopeDecodesScalarSlotValues(t *testing.T) {
	raw := []byte(`{"data":{"database":{"old":{"status":"paid"}}}}`)

	var ev dtos.CodeFunctionEventDto
	if err := json.Unmarshal(raw, &ev); err != nil {
		t.Fatalf("Unmarshal into dtos.CodeFunctionEventDto: %v", err)
	}
	if ev.Data == nil || ev.Data.Database == nil {
		t.Fatalf("data.database missing: %+v", ev)
	}
	if got := ev.Data.Database.Old["status"]; got != "paid" {
		t.Errorf("old.status = %#v, want %q", got, "paid")
	}
}
