// Command webhook-receiver is a runnable example of a Norbix webhook sink.
//
// Point your destination's Endpoint URL at http://<host>:8080/webhooks and set
// NORBIX_WEBHOOK_SIGNING_SECRET (reveal it via hub.Webhooks). It mirrors the
// reference Node handler: one typed handler for membership.user.registered, a
// database.record.inserted handler that de-duplicates on eventId, plus an OnAll
// logger for every catalog event.
package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/norbix-code/sdk-go/v2/norbix/hub/dtos"
	"github.com/norbix-code/sdk-go/v2/norbix/webhooks"
)

// seenEvents remembers the change ids already processed. One record change can
// arrive as several deliveries (the plain webhook delivery and each schema
// Webhook-trigger delivery): each has its own id, all share one eventId.
// In production keep this in a shared store with a TTL (Redis, a DB table).
type seenEvents struct {
	mu   sync.Mutex
	seen map[string]bool
}

// firstTime reports whether eventID is new, and records it.
func (s *seenEvents) firstTime(eventID string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.seen[eventID] {
		return false
	}
	s.seen[eventID] = true
	return true
}

func buildReceiver() *webhooks.Receiver {
	dedupe := &seenEvents{seen: map[string]bool{}}

	// Options default from env: NORBIX_WEBHOOK_SIGNING_SECRET,
	// NORBIX_WEBHOOK_TOLERANCE_SECONDS, NORBIX_PROJECT_ID, NORBIX_ACCOUNT_ID.
	r := webhooks.New(webhooks.Options{})

	// Typed handler — decode the payload into the generated AuthDto (the gateway renamed UserDto → AuthDto).
	r.On(webhooks.EventMembershipUserRegistered,
		func(ctx context.Context, payload json.RawMessage, e webhooks.Event) error {
			user, err := webhooks.DecodePayload[dtos.AuthDto](payload)
			if err != nil {
				return err
			}
			name := user.Email
			if name == "" {
				name = user.UserName
			}
			log.Printf("[webhook] user registered id=%s email=%s", e.Metadata.UserID, name)
			return nil
		})

	// Record changes — de-duplicate on the event id, not the delivery id.
	// EffectiveEventID falls back to id for gateways that do not send eventId.
	r.On(webhooks.EventDatabaseRecordInserted,
		func(ctx context.Context, payload json.RawMessage, e webhooks.Event) error {
			if !dedupe.firstTime(e.EventID) {
				log.Printf("[webhook] skip duplicate change eventId=%s delivery=%s trigger=%s",
					e.EventID, e.DeliveryID, e.TriggerID)
				return nil // still answer 2xx so the gateway does not retry
			}
			log.Printf("[webhook] record inserted schema=%s id=%s eventId=%s",
				e.Metadata.SchemaName, e.Metadata.RecordID, e.EventID)
			return nil
		})

	// Catch-all logger — runs for every listed event, even alongside On.
	r.OnAll(webhooks.EventNames, func(ctx context.Context, env webhooks.Envelope, c webhooks.Context) error {
		log.Printf("──────────── Norbix webhook ────────────")
		log.Printf("handler: %s delivery=%s eventId=%s trigger=%s verified=%v",
			env.Event, env.ID, env.EffectiveEventID(), env.TriggerID, c.Verified)
		log.Printf("payload: %s", string(env.Data))
		return nil
	})
	return r
}

func main() {
	receiver := buildReceiver()

	http.HandleFunc("/webhooks", func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		// The raw body is required for signature verification — read it exactly.
		raw, err := io.ReadAll(req.Body)
		if err != nil {
			http.Error(w, "read error", http.StatusBadRequest)
			return
		}

		res, err := receiver.Handle(req.Context(), webhooks.HandleInput{
			RawBody: raw,
			Headers: req.Header,
			Path:    req.URL.Path,
		})
		if err != nil {
			var sigErr *webhooks.SignatureError
			if errors.As(err, &sigErr) {
				http.Error(w, sigErr.Error(), http.StatusUnauthorized) // 401
				return
			}
			http.Error(w, err.Error(), http.StatusBadRequest) // 400 (parse, etc.)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"received":   res.Received,
			"event":      res.Event,
			"deliveryId": res.DeliveryID,
			"eventId":    res.EventID,
			"handled":    res.Handled,
		})
	})

	log.Println("listening on :8080/webhooks")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
