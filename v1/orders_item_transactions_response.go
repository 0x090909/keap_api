package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use OrdersItemTransactionsGetResponseable instead.
type OrdersItemTransactionsResponse struct {
	OrdersItemTransactionsGetResponse
}

// NewOrdersItemTransactionsResponse instantiates a new OrdersItemTransactionsResponse and sets the default values.
func NewOrdersItemTransactionsResponse() *OrdersItemTransactionsResponse {
	m := &OrdersItemTransactionsResponse{
		OrdersItemTransactionsGetResponse: *NewOrdersItemTransactionsGetResponse(),
	}
	return m
}

// CreateOrdersItemTransactionsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrdersItemTransactionsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewOrdersItemTransactionsResponse(), nil
}

// Deprecated: This class is obsolete. Use OrdersItemTransactionsGetResponseable instead.
type OrdersItemTransactionsResponseable interface {
	OrdersItemTransactionsGetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
