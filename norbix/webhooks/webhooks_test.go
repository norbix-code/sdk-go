package webhooks

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/norbix-code/sdk-go/v2/norbix"
)

func sign(secret, ts string, body []byte) string { return ComputeSignature(secret, ts, body) }

func TestVerifySignatureOK(t *testing.T) {
	secret := "shh"
	body := []byte(`{"id":"d1","event":"membership.user.registered"}`)
	ts := fmt.Sprintf("%d", time.Now().Unix())
	if err := VerifySignature(secret, body, sign(secret, ts, body), ts, 300); err != nil {
		t.Fatalf("expected ok, got %v", err)
	}
}

func TestVerifySignatureMismatch(t *testing.T) {
	body := []byte(`{}`)
	ts := fmt.Sprintf("%d", time.Now().Unix())
	err := VerifySignature("shh", body, "sha256=deadbeef", ts, 300)
	var se *SignatureError
	if !errors.As(err, &se) {
		t.Fatalf("expected *SignatureError, got %T", err)
	}
}

func TestVerifyTimestampTolerance(t *testing.T) {
	secret := "shh"
	body := []byte(`{}`)
	old := fmt.Sprintf("%d", time.Now().Add(-2*time.Hour).Unix())
	err := VerifySignature(secret, body, sign(secret, old, body), old, 300)
	if err == nil {
		t.Fatal("expected stale timestamp to fail")
	}
}

func TestReceiverDispatchTypedMembership(t *testing.T) {
	secret := "shh"
	user := `{"id":"u1","email":"a@b.com","userName":"alice"}`
	data := `{"id":"u1","to":` + user + `}`
	body := []byte(`{"id":"d1","event":"membership.user.registered","createdOn":"2026-01-01T00:00:00Z","accountId":"acc","projectId":"proj","data":` + data + `}`)
	ts := fmt.Sprintf("%d", time.Now().Unix())

	r := New(Options{Secret: secret})
	type u struct {
		ID    string `json:"id"`
		Email string `json:"email"`
	}
	var got u
	var gotMetaUser string
	r.On(EventMembershipUserRegistered, func(ctx context.Context, payload json.RawMessage, e Event) error {
		v, err := DecodePayload[u](payload)
		if err != nil {
			return err
		}
		got = v
		gotMetaUser = e.Metadata.UserID
		return nil
	})

	res, err := r.Handle(context.Background(), HandleInput{
		RawBody: body,
		Headers: map[string][]string{
			HeaderSignature: {sign(secret, ts, body)},
			HeaderTimestamp: {ts},
			HeaderEvent:     {"membership.user.registered"},
		},
	})
	if err != nil {
		t.Fatalf("handle error: %v", err)
	}
	if !res.Handled || got.Email != "a@b.com" || gotMetaUser != "u1" {
		t.Errorf("dispatch failed: handled=%v got=%+v meta=%q", res.Handled, got, gotMetaUser)
	}
	if res.Verified == nil || !*res.Verified {
		t.Errorf("expected verified=true")
	}
}

func TestReceiverMutationDatabase(t *testing.T) {
	data := `{"schemaName":"orders","id":"r1","from":{"n":1},"to":{"n":2}}`
	body := []byte(`{"id":"d2","event":"database.record.updated","accountId":"a","projectId":"p","data":` + data + `}`)
	r := New(Options{}) // no secret -> skip verify
	type doc struct {
		N int `json:"n"`
	}
	var mut Mutation[doc]
	var schema string
	r.On(EventDatabaseRecordUpdated, func(ctx context.Context, payload json.RawMessage, e Event) error {
		m, err := DecodeMutation[doc](payload)
		if err != nil {
			return err
		}
		mut = m
		schema = e.Metadata.SchemaName
		return nil
	})
	if _, err := r.Handle(context.Background(), HandleInput{RawBody: body}); err != nil {
		t.Fatalf("handle: %v", err)
	}
	if mut.From.N != 1 || mut.To.N != 2 || schema != "orders" {
		t.Errorf("mutation/meta wrong: %+v schema=%q", mut, schema)
	}
}

func TestReceiverGuardProjectMismatch(t *testing.T) {
	body := []byte(`{"id":"d3","event":"files.file.uploaded","accountId":"a","projectId":"WRONG","data":{}}`)
	r := New(Options{ProjectID: "RIGHT"})
	_, err := r.Handle(context.Background(), HandleInput{RawBody: body})
	var se *SignatureError
	if !errors.As(err, &se) {
		t.Fatalf("expected guard rejection, got %T %v", err, err)
	}
}

func TestParseEnvelopeInvalid(t *testing.T) {
	_, err := ParseEnvelope([]byte(`{"event":"x"}`))
	var pe *ParseError
	if !errors.As(err, &pe) {
		t.Fatalf("expected *ParseError, got %T", err)
	}
}

func TestReceiverEventIDSharedAcrossDeliveries(t *testing.T) {
	// One record change → two deliveries to the same destination: the plain
	// webhook (triggerId null) and a schema Webhook trigger. Different ids,
	// one eventId.
	plain := []byte(`{"id":"d-plain","eventId":"evt-1","event":"database.record.inserted","accountId":"a","projectId":"p","triggerId":null,"data":{"schemaName":"orders","id":"r1"}}`)
	trig := []byte(`{"id":"d-trig","eventId":"evt-1","event":"database.record.inserted","accountId":"a","projectId":"p","triggerId":"t1","data":{"schemaName":"orders","id":"r1"}}`)

	r := New(Options{})
	seen := map[string]bool{}
	processed := 0
	var gotEventIDs, gotDeliveryIDs []string
	r.On(EventDatabaseRecordInserted, func(ctx context.Context, payload json.RawMessage, e Event) error {
		gotEventIDs = append(gotEventIDs, e.EventID)
		gotDeliveryIDs = append(gotDeliveryIDs, e.DeliveryID)
		if seen[e.EventID] {
			return nil
		}
		seen[e.EventID] = true
		processed++
		return nil
	})

	for _, body := range [][]byte{plain, trig} {
		res, err := r.Handle(context.Background(), HandleInput{RawBody: body})
		if err != nil {
			t.Fatalf("handle: %v", err)
		}
		if res.EventID != "evt-1" {
			t.Errorf("result EventID = %q, want evt-1", res.EventID)
		}
	}
	if processed != 1 {
		t.Errorf("processed = %d, want 1 (dedupe on EventID)", processed)
	}
	if fmt.Sprint(gotEventIDs) != "[evt-1 evt-1]" || fmt.Sprint(gotDeliveryIDs) != "[d-plain d-trig]" {
		t.Errorf("ids wrong: eventIds=%v deliveryIds=%v", gotEventIDs, gotDeliveryIDs)
	}
}

func TestReceiverEventIDFallsBackToID(t *testing.T) {
	// Older gateway: no eventId on the envelope → EventID falls back to id.
	body := []byte(`{"id":"d-old","event":"files.file.uploaded","accountId":"a","projectId":"p","data":{}}`)
	env, err := ParseEnvelope(body)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if env.EventID != "" || env.EffectiveEventID() != "d-old" {
		t.Errorf("EventID=%q Effective=%q, want empty / d-old", env.EventID, env.EffectiveEventID())
	}

	r := New(Options{})
	var got string
	r.On(EventFilesFileUploaded, func(ctx context.Context, payload json.RawMessage, e Event) error {
		got = e.EventID
		return nil
	})
	res, err := r.Handle(context.Background(), HandleInput{RawBody: body})
	if err != nil {
		t.Fatalf("handle: %v", err)
	}
	if got != "d-old" || res.EventID != "d-old" {
		t.Errorf("fallback failed: event=%q result=%q", got, res.EventID)
	}
}

// payments.webhook.received (gateway campaign audit/payments): the payload is
// the normalized payment fact, the provider ids land on Metadata.
func TestReceiverPaymentsWebhookReceived(t *testing.T) {
	fact := `{"integrationId":"int_stripe","provider":"Stripe","providerEventId":"evt_1","eventName":"payment_intent.succeeded","status":"Succeeded","amountMinor":1099,"currency":"USD","externalIds":{"customer":"cus_1"},"receivedOn":"2026-10-15T08:00:00Z"}`
	body := []byte(`{"id":"d3","eventId":"e3","event":"payments.webhook.received","accountId":"a","projectId":"p","triggerId":"trg_p","data":` + fact + `}`)
	r := New(Options{})
	type payment struct {
		EventName   string `json:"eventName"`
		Status      string `json:"status"`
		AmountMinor int64  `json:"amountMinor"`
		Currency    string `json:"currency"`
	}
	var got payment
	var meta EventMetadata
	r.On(EventPaymentsWebhookReceived, func(ctx context.Context, payload json.RawMessage, e Event) error {
		v, err := DecodePayload[payment](payload)
		if err != nil {
			return err
		}
		got = v
		meta = e.Metadata
		return nil
	})
	res, err := r.Handle(context.Background(), HandleInput{RawBody: body})
	if err != nil {
		t.Fatalf("handle: %v", err)
	}
	if !res.Handled || got.EventName != "payment_intent.succeeded" || got.AmountMinor != 1099 || got.Currency != "USD" || got.Status != "Succeeded" {
		t.Errorf("payload wrong: handled=%v got=%+v", res.Handled, got)
	}
	if meta.IntegrationID != "int_stripe" || meta.PaymentProvider != "Stripe" || meta.PaymentProviderEventID != "evt_1" || meta.PaymentEventName != "payment_intent.succeeded" {
		t.Errorf("metadata wrong: %+v", meta)
	}
}

// payments.order.paid (gateway campaign audit/payments-checkout): the payload is
// the order event (norbix.PaymentOrderEvent); the order ids and the provider
// ids of its payment land on Metadata.
func TestReceiverPaymentsOrderPaid(t *testing.T) {
	orderEvent := `{"orderEventId":"ord_2ToIiRM03X63TZaJ2dJdv2:2","event":"payments.order.paid","from":"CheckoutStarted","to":"Paid","reason":"paid in attempt 1",` +
		`"order":{"id":"ord_2ToIiRM03X63TZaJ2dJdv2","env":"PROD","buyerUserId":"usr_1","merchantReference":"ad-1042-promo-1","currency":"EUR","priceVersion":"2026-10-01",` +
		`"lines":[{"sku":"star-day","title":"Star × day","quantity":21,"unitAmountMinor":30,"amountMinor":630}],"totalMinor":630,"metadata":{"adId":"ad_1042"},` +
		`"status":"Paid","checkouts":[],"settledMinor":630,"refundedMinor":0,"providerPaymentId":"mnt-1","mismatches":[]},` +
		`"payment":{"integrationId":"int_montonio","provider":"Montonio","providerEventId":"order:mnt-1:PAID","orderId":"ord_2ToIiRM03X63TZaJ2dJdv2","eventName":"order.paid","status":"Succeeded","amountMinor":630,"currency":"EUR","receivedOn":"2026-10-19T09:54:43Z"}}`
	body := []byte(`{"id":"dlv_o","eventId":"evt_order_1","event":"payments.order.paid","accountId":"a","projectId":"p","triggerId":"trg_o","data":` + orderEvent + `}`)
	r := New(Options{})
	var got norbix.PaymentOrderEvent
	var meta EventMetadata
	r.On(EventPaymentsOrderPaid, func(ctx context.Context, payload json.RawMessage, e Event) error {
		v, err := DecodePayload[norbix.PaymentOrderEvent](payload)
		if err != nil {
			return err
		}
		got = v
		meta = e.Metadata
		return nil
	})
	res, err := r.Handle(context.Background(), HandleInput{RawBody: body})
	if err != nil {
		t.Fatalf("handle: %v", err)
	}
	if !res.Handled || got.OrderEventID != "ord_2ToIiRM03X63TZaJ2dJdv2:2" || got.To != "Paid" || got.Order.Status != "Paid" || got.Order.TotalMinor != 630 || got.Payment == nil || got.Payment.ProviderEventID != "order:mnt-1:PAID" {
		t.Errorf("payload wrong: handled=%v got=%+v", res.Handled, got)
	}
	if meta.IntegrationID != "int_montonio" || meta.PaymentProvider != "Montonio" || meta.PaymentProviderEventID != "order:mnt-1:PAID" || meta.PaymentEventName != "order.paid" {
		t.Errorf("payment metadata wrong: %+v", meta)
	}
	if meta.OrderID != "ord_2ToIiRM03X63TZaJ2dJdv2" || meta.OrderEventID != "ord_2ToIiRM03X63TZaJ2dJdv2:2" || meta.OrderStatus != "Paid" {
		t.Errorf("order metadata wrong: %+v", meta)
	}
}

func TestPaymentsOrderEventsAreInTheCatalog(t *testing.T) {
	want := []string{EventPaymentsOrderPaid, EventPaymentsOrderPartiallyRefunded, EventPaymentsOrderRefunded, EventPaymentsOrderReversed, EventPaymentsOrderExpired}
	names := map[string]bool{}
	for _, n := range EventNames {
		names[n] = true
	}
	for i, n := range []string{"payments.order.paid", "payments.order.partially_refunded", "payments.order.refunded", "payments.order.reversed", "payments.order.expired"} {
		if want[i] != n || !names[n] {
			t.Errorf("event %q: constant %q, in catalog %v", n, want[i], names[n])
		}
	}
}
