package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use OrdersItemPaymentsPostResponseable instead.
type OrdersItemPaymentsResponse struct {
	OrdersItemPaymentsPostResponse
}

// NewOrdersItemPaymentsResponse instantiates a new OrdersItemPaymentsResponse and sets the default values.
func NewOrdersItemPaymentsResponse() *OrdersItemPaymentsResponse {
	m := &OrdersItemPaymentsResponse{
		OrdersItemPaymentsPostResponse: *NewOrdersItemPaymentsPostResponse(),
	}
	return m
}

// CreateOrdersItemPaymentsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrdersItemPaymentsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewOrdersItemPaymentsResponse(), nil
}

// Deprecated: This class is obsolete. Use OrdersItemPaymentsPostResponseable instead.
type OrdersItemPaymentsResponseable interface {
	OrdersItemPaymentsPostResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
