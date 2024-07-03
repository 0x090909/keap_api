package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use PaymentMethodConfigsPostResponseable instead.
type PaymentMethodConfigsResponse struct {
	PaymentMethodConfigsPostResponse
}

// NewPaymentMethodConfigsResponse instantiates a new PaymentMethodConfigsResponse and sets the default values.
func NewPaymentMethodConfigsResponse() *PaymentMethodConfigsResponse {
	m := &PaymentMethodConfigsResponse{
		PaymentMethodConfigsPostResponse: *NewPaymentMethodConfigsPostResponse(),
	}
	return m
}

// CreatePaymentMethodConfigsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreatePaymentMethodConfigsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewPaymentMethodConfigsResponse(), nil
}

// Deprecated: This class is obsolete. Use PaymentMethodConfigsPostResponseable instead.
type PaymentMethodConfigsResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	PaymentMethodConfigsPostResponseable
}
