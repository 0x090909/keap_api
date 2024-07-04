package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use SubscriptionsGetResponseable instead.
type SubscriptionsResponse struct {
	SubscriptionsGetResponse
}

// NewSubscriptionsResponse instantiates a new SubscriptionsResponse and sets the default values.
func NewSubscriptionsResponse() *SubscriptionsResponse {
	m := &SubscriptionsResponse{
		SubscriptionsGetResponse: *NewSubscriptionsGetResponse(),
	}
	return m
}

// CreateSubscriptionsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSubscriptionsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSubscriptionsResponse(), nil
}

// Deprecated: This class is obsolete. Use SubscriptionsGetResponseable instead.
type SubscriptionsResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	SubscriptionsGetResponseable
}
