// Hand-written, not generated (gateway campaign audit/payments-checkout).
//
// Written in the shape of the other modules (files.go), whose generator was
// lost; the request and response types are the generated dtos
// (CreatePaymentOrderRequest, GetPaymentOrderResponse, …) once the types are
// regenerated from a gateway with payment orders.

package api

import (
	"context"

	"github.com/norbix-code/sdk-go/v2/norbix/internal/transport"
)

// PaymentsModule groups the payment order and checkout endpoints on the API.
//
// A payment order is what a buyer pays for: lines priced in minor units, a
// merchant reference and a price version. Create it, start a provider checkout
// (Montonio, …) and send the buyer to checkout.paymentUrl. A verified provider
// callback moves the order on; a webhook destination or a Code function hears
// it as payments.order.paid / …refunded / …expired.
type PaymentsModule struct{ t *transport.Transport }

// CreatePaymentOrder performs POST /{version}/payments/orders (scope: project).
// Body: buyerUserId, merchantReference, currency, priceVersion, lines (each
// sku, title, quantity, unitAmountMinor), optional idempotencyKey and metadata
// (a string map). Answers {id: "ord_…"}; the same body with the same
// idempotencyKey answers the same id.
func (m *PaymentsModule) CreatePaymentOrder(ctx context.Context, req map[string]any, out any) error {
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/payments/orders",
		Method:     "POST",
		PathParams: nil,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// GetPaymentOrder performs GET /{version}/payments/orders/{Id} (scope: project).
// Answers {item: order} — lines, totals, status, checkouts, settledMinor,
// refundedMinor, ….
func (m *PaymentsModule) GetPaymentOrder(ctx context.Context, id string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"Id": id,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/payments/orders/{Id}",
		Method:     "GET",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// GetPaymentOrders performs GET /{version}/payments/orders (scope: project).
// Optional query: buyerUserId, merchantReference, status. Answers
// {list: {items: [...], hasMore, …}}.
func (m *PaymentsModule) GetPaymentOrders(ctx context.Context, req map[string]any, out any) error {
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/payments/orders",
		Method:     "GET",
		PathParams: nil,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// StartPaymentCheckout performs POST /{version}/payments/orders/{Id}/checkout (scope: project).
// Body: integrationId, returnUrl, optional method, preferredCountry,
// preferredProvider. Answers checkout (paymentUrl to send the buyer to,
// expiresAtUtc, …), reused (an open checkout was handed back) and the order.
func (m *PaymentsModule) StartPaymentCheckout(ctx context.Context, id string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"Id": id,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/payments/orders/{Id}/checkout",
		Method:     "POST",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}
