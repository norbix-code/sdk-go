package webhooks

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	"time"
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
