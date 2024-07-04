package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use OrdersItemItemsPostResponseable instead.
type OrdersItemItemsResponse struct {
	OrdersItemItemsPostResponse
}

// NewOrdersItemItemsResponse instantiates a new OrdersItemItemsResponse and sets the default values.
func NewOrdersItemItemsResponse() *OrdersItemItemsResponse {
	m := &OrdersItemItemsResponse{
		OrdersItemItemsPostResponse: *NewOrdersItemItemsPostResponse(),
	}
	return m
}

// CreateOrdersItemItemsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrdersItemItemsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewOrdersItemItemsResponse(), nil
}

// Deprecated: This class is obsolete. Use OrdersItemItemsPostResponseable instead.
type OrdersItemItemsResponseable interface {
	OrdersItemItemsPostResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
