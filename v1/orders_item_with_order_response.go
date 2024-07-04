package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use OrdersItemWithOrderGetResponseable instead.
type OrdersItemWithOrderResponse struct {
	OrdersItemWithOrderGetResponse
}

// NewOrdersItemWithOrderResponse instantiates a new OrdersItemWithOrderResponse and sets the default values.
func NewOrdersItemWithOrderResponse() *OrdersItemWithOrderResponse {
	m := &OrdersItemWithOrderResponse{
		OrdersItemWithOrderGetResponse: *NewOrdersItemWithOrderGetResponse(),
	}
	return m
}

// CreateOrdersItemWithOrderResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrdersItemWithOrderResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewOrdersItemWithOrderResponse(), nil
}

// Deprecated: This class is obsolete. Use OrdersItemWithOrderGetResponseable instead.
type OrdersItemWithOrderResponseable interface {
	OrdersItemWithOrderGetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
