package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use OrdersGetResponseable instead.
type OrdersResponse struct {
	OrdersGetResponse
}

// NewOrdersResponse instantiates a new OrdersResponse and sets the default values.
func NewOrdersResponse() *OrdersResponse {
	m := &OrdersResponse{
		OrdersGetResponse: *NewOrdersGetResponse(),
	}
	return m
}

// CreateOrdersResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrdersResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewOrdersResponse(), nil
}

// Deprecated: This class is obsolete. Use OrdersGetResponseable instead.
type OrdersResponseable interface {
	OrdersGetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
