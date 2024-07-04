package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use OrdersItemPaymentPlanPutResponseable instead.
type OrdersItemPaymentPlanResponse struct {
	OrdersItemPaymentPlanPutResponse
}

// NewOrdersItemPaymentPlanResponse instantiates a new OrdersItemPaymentPlanResponse and sets the default values.
func NewOrdersItemPaymentPlanResponse() *OrdersItemPaymentPlanResponse {
	m := &OrdersItemPaymentPlanResponse{
		OrdersItemPaymentPlanPutResponse: *NewOrdersItemPaymentPlanPutResponse(),
	}
	return m
}

// CreateOrdersItemPaymentPlanResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrdersItemPaymentPlanResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewOrdersItemPaymentPlanResponse(), nil
}

// Deprecated: This class is obsolete. Use OrdersItemPaymentPlanPutResponseable instead.
type OrdersItemPaymentPlanResponseable interface {
	OrdersItemPaymentPlanPutResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
