// Command webhook-receiver is a runnable example of a Norbix webhook sink.
//
// Point your destination's Endpoint URL at http://<host>:8080/webhooks and set
// NORBIX_WEBHOOK_SIGNING_SECRET (reveal it via hub.Webhooks). It mirrors the
// reference Node handler: one typed handler for membership.user.registered plus
// an OnAll logger for every catalog event.
package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/norbix-code/sdk-go/norbix/hub/dtos"
	"github.com/norbix-code/sdk-go/norbix/webhooks"
)

func buildReceiver() *webhooks.Receiver {
	// Options default from env: NORBIX_WEBHOOK_SIGNING_SECRET,
	// NORBIX_WEBHOOK_TOLERANCE_SECONDS, NORBIX_PROJECT_ID, NORBIX_ACCOUNT_ID.
	r := webhooks.New(webhooks.Options{})

	// Typed handler — decode the payload into the generated UserDto.
	r.On(webhooks.EventMembershipUserRegistered,
		func(ctx context.Context, payload json.RawMessage, e webhooks.Event) error {
			user, err := webhooks.DecodePayload[dtos.UserDto](payload)
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

	// Catch-all logger — runs for every listed event, even alongside On.
	r.OnAll(webhooks.EventNames, func(ctx context.Context, env webhooks.Envelope, c webhooks.Context) error {
		log.Printf("──────────── Norbix webhook ────────────")
		log.Printf("handler: %s delivery=%s trigger=%s verified=%v",
			env.Event, env.ID, env.TriggerID, c.Verified)
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
			"handled":    res.Handled,
		})
	})

	log.Println("listening on :8080/webhooks")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
