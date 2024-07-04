package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use AffiliatesItemPaymentsGetResponseable instead.
type AffiliatesItemPaymentsResponse struct {
	AffiliatesItemPaymentsGetResponse
}

// NewAffiliatesItemPaymentsResponse instantiates a new AffiliatesItemPaymentsResponse and sets the default values.
func NewAffiliatesItemPaymentsResponse() *AffiliatesItemPaymentsResponse {
	m := &AffiliatesItemPaymentsResponse{
		AffiliatesItemPaymentsGetResponse: *NewAffiliatesItemPaymentsGetResponse(),
	}
	return m
}

// CreateAffiliatesItemPaymentsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAffiliatesItemPaymentsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAffiliatesItemPaymentsResponse(), nil
}

// Deprecated: This class is obsolete. Use AffiliatesItemPaymentsGetResponseable instead.
type AffiliatesItemPaymentsResponseable interface {
	AffiliatesItemPaymentsGetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
