package norbix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	norbixerr "github.com/norbix-code/sdk-go/v2/norbix/errors"
)

// FunctionEventSpecVersion is the only envelope version this SDK reads.
const FunctionEventSpecVersion = "1"

// CodeFunctionEvent is the envelope Norbix sends to a cloud function (AWS
// Lambda, Azure Functions, Google Cloud Functions): the event name, the data
// of the source event, the context — including a short-lived bearer token
// for the run-as service user — and the developer's resolved config. One
// shape for every source: database, membership, files, payments, ai,
// scheduler, manual.
//
// Parse it with ParseFunctionEvent, pick the typed view that matches Event
// (AsDatabaseRecordChange, AsUserChange, …) and build a client for the
// run-as user with ClientOptions:
//
//	event, err := norbix.ParseFunctionEvent(raw)
//	var old, next *Order
//	_, err = event.AsDatabaseRecordChange(&old, &next)
//	c, err := norbix.New(event.ClientOptions())
//
// The raw slots stay on Data (json.RawMessage, nil when absent); the whole
// body is Raw().
type CodeFunctionEvent struct {
	SpecVersion string          `json:"specVersion"`
	ID          string          `json:"id"`        // this delivery
	EventID     string          `json:"eventId"`   // the change; a webhook delivery of the same change carries the same id
	Event       string          `json:"event"`     // the handler key: database.record.updated, membership.user.registered, scheduler.task.fired, manual.invoke, …
	Source      string          `json:"source"`    // database | membership | files | payments | ai | scheduler | manual
	CreatedOn   time.Time       `json:"createdOn"` // when the envelope was built
	Context     FunctionContext `json:"context"`   // ids, deployment urls, the caller (as data) and the run-as user with its token
	Data        FunctionData    `json:"data"`      // the raw slots — one is set per source
	Config      json.RawMessage `json:"config"`    // the developer's config, bound at fire time; nil when the action has none

	raw json.RawMessage
}

// FunctionContext is the context of a CodeFunctionEvent: where it happened
// and who the function may act as.
type FunctionContext struct {
	AccountID     string          `json:"accountId"`
	ProjectID     string          `json:"projectId"`
	Environment   string          `json:"environment"` // PROD / TEST / … — the environment the event happened in
	Region        string          `json:"region"`      // the project's region code, when it has one
	HubURL        string          `json:"hubUrl"`
	APIURL        string          `json:"apiUrl"`
	EchoURL       string          `json:"echoUrl"`
	TriggerID     string          `json:"triggerId"` // which trigger fired it (or TaskID; neither for a manual call)
	TaskID        string          `json:"taskId"`
	CorrelationID string          `json:"correlationId"`
	Caller        *FunctionCaller `json:"caller"` // who caused the source event, as data, never authority; nil when the source has no caller
	RunAs         *FunctionRunAs  `json:"runAs"`  // the service user the function acts as
}

// FunctionCaller is who caused the source event — data for the function, not
// a credential.
type FunctionCaller struct {
	UserID string `json:"userId"`
	Email  string `json:"email"`
	Kind   string `json:"kind"` // user | apiKey | system
}

// FunctionRunAs is the service user the function acts as, with a short-lived
// bearer token for it (15 minutes).
type FunctionRunAs struct {
	UserID      string     `json:"userId"`
	AccessToken string     `json:"accessToken"`
	ExpiresAt   *time.Time `json:"expiresAt"`
}

// FunctionData holds the raw data slots of the envelope; one is set per
// source, the others are nil. The As* views decode them.
type FunctionData struct {
	Database json.RawMessage `json:"database"`
	User     json.RawMessage `json:"user"`
	File     json.RawMessage `json:"file"`
	Payment  json.RawMessage `json:"payment"`
	Order    json.RawMessage `json:"order"` // payments.order.* only
	AI       json.RawMessage `json:"ai"`
	Task     json.RawMessage `json:"task"`
	Payload  json.RawMessage `json:"payload"` // manual invoke, scheduler static payload
}

// DatabaseRecordChange is the Database slot: which record changed and the
// documents before and after. Above 200 KB for both together the documents
// are left out, Truncated is true and OldRef / NewRef say where to fetch them.
type DatabaseRecordChange struct {
	Collection    string          `json:"collection"`
	SchemaID      string          `json:"schemaId"`
	SchemaVersion *int            `json:"schemaVersion"`
	RecordID      string          `json:"recordId"`
	Old           json.RawMessage `json:"old"` // nil on insert or when truncated
	New           json.RawMessage `json:"new"` // nil on delete or when truncated
	Truncated     bool            `json:"truncated"`
	OldRef        *DocumentRef    `json:"oldRef"`
	NewRef        *DocumentRef    `json:"newRef"`
}

// DocumentRef says where a truncated document can be fetched.
type DocumentRef struct {
	Kind                  string `json:"kind"`
	ProjectID             string `json:"projectId"`
	IntegrationID         string `json:"integrationId"`
	DatabaseIntegrationID string `json:"databaseIntegrationId"`
	SchemaID              string `json:"schemaId"`
	ExternalID            string `json:"externalId"`
}

// UserChange is the Membership slot: the user before and after (OnInvited
// has only new.email).
type UserChange struct {
	UserID string          `json:"userId"`
	Old    json.RawMessage `json:"old"`
	New    json.RawMessage `json:"new"`
}

// FileChange is the Files slot: upload → New, delete → Old (today only the path).
type FileChange struct {
	Old  json.RawMessage `json:"old"`
	New  json.RawMessage `json:"new"`
	Path string          `json:"path"`
}

// Payment is the Payments slot: the webhook fact, no "before".
type Payment struct {
	IntegrationID   string            `json:"integrationId"`
	Provider        string            `json:"provider"`
	ProviderEventID string            `json:"providerEventId"`
	EventName       string            `json:"eventName"`
	OrderID         string            `json:"orderId"` // the payment order (ord_…) the callback is about, when it is about one
	Status          string            `json:"status"`
	AmountMinor     *int64            `json:"amountMinor"`
	Currency        string            `json:"currency"`
	ExternalIDs     map[string]string `json:"externalIds"`
	OccurredOn      *time.Time        `json:"occurredOn"`
	ReceivedOn      time.Time         `json:"receivedOn"`
}

// PaymentOrderEvent is the Payments order slot (payments.order.*): a stored
// payment order changed. Order is the order as stored after the change,
// Payment the provider callback that changed it. OrderEventID is the
// de-duplication key: it is the same on every retry of the same change.
type PaymentOrderEvent struct {
	OrderEventID string       `json:"orderEventId"`
	Event        string       `json:"event"` // payments.order.paid | partially_refunded | refunded | reversed | expired
	From         string       `json:"from"`  // order status before the change
	To           string       `json:"to"`    // order status after the change
	Reason       string       `json:"reason"`
	Order        PaymentOrder `json:"order"`
	Payment      *Payment     `json:"payment"`
}

// PaymentOrder is a stored payment order: lines priced in minor units, the
// total, its status and the provider checkouts started for it.
type PaymentOrder struct {
	ID                string                 `json:"id"`
	Env               string                 `json:"env"`
	BuyerUserID       string                 `json:"buyerUserId"`
	MerchantReference string                 `json:"merchantReference"`
	Currency          string                 `json:"currency"`
	PriceVersion      string                 `json:"priceVersion"`
	Lines             []PaymentOrderLine     `json:"lines"`
	TotalMinor        int64                  `json:"totalMinor"`
	Metadata          map[string]string      `json:"metadata"`
	Status            string                 `json:"status"` // Created | CheckoutStarted | Paid | PartiallyRefunded | Refunded | Reversed | Expired | …
	CreatedOnUtc      *time.Time             `json:"createdOnUtc"`
	UpdatedOnUtc      *time.Time             `json:"updatedOnUtc"`
	Checkouts         []PaymentOrderCheckout `json:"checkouts"`
	SettledMinor      *int64                 `json:"settledMinor"`
	RefundedMinor     int64                  `json:"refundedMinor"`
	PaidOnUtc         *time.Time             `json:"paidOnUtc"`
	ProviderPaymentID string                 `json:"providerPaymentId"`
	Mismatches        []string               `json:"mismatches"`
}

// PaymentOrderLine is one priced line of a payment order.
type PaymentOrderLine struct {
	SKU             string `json:"sku"`
	Title           string `json:"title"`
	Quantity        int64  `json:"quantity"`
	UnitAmountMinor int64  `json:"unitAmountMinor"`
	AmountMinor     int64  `json:"amountMinor"`
}

// PaymentOrderCheckout is one provider checkout started for a payment order.
type PaymentOrderCheckout struct {
	Attempt           int        `json:"attempt"`
	IntegrationID     string     `json:"integrationId"`
	Provider          string     `json:"provider"`
	ProviderOrderID   string     `json:"providerOrderId"`
	ProviderReference string     `json:"providerReference"`
	PaymentURL        string     `json:"paymentUrl"`
	Method            string     `json:"method"`
	StartedOnUtc      *time.Time `json:"startedOnUtc"`
	ExpiresAtUtc      *time.Time `json:"expiresAtUtc"`
}

// AIUsage is the AI slot: the usage payload and the user it concerns.
type AIUsage struct {
	Payload      json.RawMessage `json:"payload"`
	TargetUserID string          `json:"targetUserId"`
}

// ScheduledTask is the Scheduler slot: which task fired and when.
type ScheduledTask struct {
	TaskID  string    `json:"taskId"`
	Name    string    `json:"name"`
	FiredAt time.Time `json:"firedAt"`
}

// ParseFunctionEvent reads the JSON body a cloud function received. It
// refuses a body that is not an envelope (*errors.Error, code
// NORBIX_FUNCTION_EVENT_INVALID) and an envelope whose specVersion this SDK
// does not read (NORBIX_FUNCTION_EVENT_UNSUPPORTED).
func ParseFunctionEvent(raw []byte) (*CodeFunctionEvent, error) {
	var head struct {
		SpecVersion *string `json:"specVersion"`
	}
	if err := json.Unmarshal(raw, &head); err != nil {
		return nil, norbixerr.New("The function event is not valid JSON: "+err.Error(), norbixerr.CodeFunctionEventInvalid)
	}
	if head.SpecVersion == nil {
		return nil, norbixerr.New("The function event has no specVersion; it is not a Norbix code function envelope.", norbixerr.CodeFunctionEventInvalid)
	}
	if *head.SpecVersion != FunctionEventSpecVersion {
		return nil, norbixerr.New(
			fmt.Sprintf("The function event has specVersion %q; this SDK reads specVersion %q. Update github.com/norbix-code/sdk-go.", *head.SpecVersion, FunctionEventSpecVersion),
			norbixerr.CodeFunctionEventUnsupported,
		)
	}

	var ev CodeFunctionEvent
	if err := json.Unmarshal(raw, &ev); err != nil {
		return nil, norbixerr.New("The function event does not match the envelope shape: "+err.Error(), norbixerr.CodeFunctionEventInvalid)
	}
	if ev.Event == "" {
		return nil, norbixerr.New("The function event has no event name.", norbixerr.CodeFunctionEventInvalid)
	}

	ev.Config = presentOrNil(ev.Config)
	ev.Data.Database = presentOrNil(ev.Data.Database)
	ev.Data.User = presentOrNil(ev.Data.User)
	ev.Data.File = presentOrNil(ev.Data.File)
	ev.Data.Payment = presentOrNil(ev.Data.Payment)
	ev.Data.Order = presentOrNil(ev.Data.Order)
	ev.Data.AI = presentOrNil(ev.Data.AI)
	ev.Data.Task = presentOrNil(ev.Data.Task)
	ev.Data.Payload = presentOrNil(ev.Data.Payload)
	ev.raw = append(json.RawMessage(nil), raw...)
	return &ev, nil
}

// Raw is the envelope as it arrived, for a field this type has no view of.
func (e *CodeFunctionEvent) Raw() json.RawMessage { return e.raw }

// AsDatabaseRecordChange reads the Database slot and decodes old / new into
// oldOut / newOut (pointers to the caller's record type, either may be nil).
// A document that is absent — insert has no old, delete no new, a truncated
// change neither — leaves its out untouched. Errors with
// NORBIX_FUNCTION_EVENT_NO_SLOT when the envelope carries no Database slot.
func (e *CodeFunctionEvent) AsDatabaseRecordChange(oldOut, newOut any) (DatabaseRecordChange, error) {
	var change DatabaseRecordChange
	if err := e.decodeSlot("database", e.Data.Database, &change); err != nil {
		return DatabaseRecordChange{}, err
	}
	change.Old = presentOrNil(change.Old)
	change.New = presentOrNil(change.New)
	if err := decodeDocument("data.database.old", change.Old, oldOut); err != nil {
		return DatabaseRecordChange{}, err
	}
	if err := decodeDocument("data.database.new", change.New, newOut); err != nil {
		return DatabaseRecordChange{}, err
	}
	return change, nil
}

// AsUserChange reads the Membership slot and decodes old / new into oldOut /
// newOut (either may be nil; an absent document leaves its out untouched).
func (e *CodeFunctionEvent) AsUserChange(oldOut, newOut any) (UserChange, error) {
	var change UserChange
	if err := e.decodeSlot("user", e.Data.User, &change); err != nil {
		return UserChange{}, err
	}
	change.Old = presentOrNil(change.Old)
	change.New = presentOrNil(change.New)
	if err := decodeDocument("data.user.old", change.Old, oldOut); err != nil {
		return UserChange{}, err
	}
	if err := decodeDocument("data.user.new", change.New, newOut); err != nil {
		return UserChange{}, err
	}
	return change, nil
}

// AsFileChange reads the Files slot and decodes old / new into oldOut /
// newOut (either may be nil; an absent document leaves its out untouched).
func (e *CodeFunctionEvent) AsFileChange(oldOut, newOut any) (FileChange, error) {
	var change FileChange
	if err := e.decodeSlot("file", e.Data.File, &change); err != nil {
		return FileChange{}, err
	}
	change.Old = presentOrNil(change.Old)
	change.New = presentOrNil(change.New)
	if err := decodeDocument("data.file.old", change.Old, oldOut); err != nil {
		return FileChange{}, err
	}
	if err := decodeDocument("data.file.new", change.New, newOut); err != nil {
		return FileChange{}, err
	}
	return change, nil
}

// AsPayment reads the Payments slot. On a payments.order.* event it is the
// provider callback that changed the order.
func (e *CodeFunctionEvent) AsPayment() (Payment, error) {
	var payment Payment
	if err := e.decodeSlot("payment", e.Data.Payment, &payment); err != nil {
		return Payment{}, err
	}
	return payment, nil
}

// AsPaymentOrder reads the Payments order slot (payments.order.*): the
// stored order after the change, with OrderEventID to de-duplicate on.
// AsPayment still returns the callback that made the change. Errors with
// NORBIX_FUNCTION_EVENT_NO_SLOT when the envelope carries no data.order.
func (e *CodeFunctionEvent) AsPaymentOrder() (PaymentOrderEvent, error) {
	var order PaymentOrderEvent
	if err := e.decodeSlot("order", e.Data.Order, &order); err != nil {
		return PaymentOrderEvent{}, err
	}
	return order, nil
}

// AsAIUsage reads the AI slot and decodes its payload into payloadOut (may be nil).
func (e *CodeFunctionEvent) AsAIUsage(payloadOut any) (AIUsage, error) {
	var usage AIUsage
	if err := e.decodeSlot("ai", e.Data.AI, &usage); err != nil {
		return AIUsage{}, err
	}
	usage.Payload = presentOrNil(usage.Payload)
	if err := decodeDocument("data.ai.payload", usage.Payload, payloadOut); err != nil {
		return AIUsage{}, err
	}
	return usage, nil
}

// AsScheduledTask reads the Scheduler slot.
func (e *CodeFunctionEvent) AsScheduledTask() (ScheduledTask, error) {
	var task ScheduledTask
	if err := e.decodeSlot("task", e.Data.Task, &task); err != nil {
		return ScheduledTask{}, err
	}
	return task, nil
}

// Payload decodes the caller's JSON object of a manual call (or a scheduler
// task's static payload) into out. An envelope without a payload leaves out
// untouched and returns nil; check Data.Payload when the difference matters.
func (e *CodeFunctionEvent) Payload(out any) error {
	return decodeDocument("data.payload", e.Data.Payload, out)
}

// DecodeConfig decodes the developer's config, bound at fire time, into out.
// An action without config leaves out untouched and returns nil; check Config
// when the difference matters.
func (e *CodeFunctionEvent) DecodeConfig(out any) error {
	return decodeDocument("config", e.Config, out)
}

// ClientOptions are the Options a client takes to act as the run-as user:
// the bearer token from Context.RunAs (valid 15 minutes), the project,
// account, environment and region of the event and the deployment's Hub and
// API urls. Without a run-as token BearerToken is empty and the first call
// fails with NORBIX_NOT_AUTHENTICATED.
func (e *CodeFunctionEvent) ClientOptions() Options {
	ctx := e.Context
	opts := Options{
		ProjectID:  ctx.ProjectID,
		AccountID:  ctx.AccountID,
		Env:        ctx.Environment,
		Region:     ctx.Region,
		BaseURLHub: ctx.HubURL,
		BaseURLAPI: ctx.APIURL,
	}
	if ctx.RunAs != nil {
		opts.BearerToken = ctx.RunAs.AccessToken
	}
	return opts
}

func (e *CodeFunctionEvent) decodeSlot(name string, slot json.RawMessage, out any) error {
	if slot == nil {
		return norbixerr.New(
			fmt.Sprintf("This %q event (source %q) carries no data.%s slot.", e.Event, e.Source, name),
			norbixerr.CodeFunctionEventNoSlot,
		)
	}
	if err := json.Unmarshal(slot, out); err != nil {
		return norbixerr.New(fmt.Sprintf("data.%s does not match the envelope shape: %s", name, err), norbixerr.CodeFunctionEventInvalid)
	}
	return nil
}

func decodeDocument(path string, doc json.RawMessage, out any) error {
	if out == nil || doc == nil {
		return nil
	}
	if err := json.Unmarshal(doc, out); err != nil {
		return norbixerr.New(fmt.Sprintf("%s does not decode into %T: %s", path, out, err), norbixerr.CodeFunctionEventInvalid)
	}
	return nil
}

// presentOrNil turns an absent or JSON null slot into nil so callers can
// test a single condition.
func presentOrNil(raw json.RawMessage) json.RawMessage {
	trimmed := bytes.TrimSpace(raw)
	if len(trimmed) == 0 || bytes.Equal(trimmed, []byte("null")) {
		return nil
	}
	return raw
}
