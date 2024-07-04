package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use AffiliatesItemAffiliateIdGetResponseable instead.
type AffiliatesItemAffiliateIdResponse struct {
	AffiliatesItemAffiliateIdGetResponse
}

// NewAffiliatesItemAffiliateIdResponse instantiates a new AffiliatesItemAffiliateIdResponse and sets the default values.
func NewAffiliatesItemAffiliateIdResponse() *AffiliatesItemAffiliateIdResponse {
	m := &AffiliatesItemAffiliateIdResponse{
		AffiliatesItemAffiliateIdGetResponse: *NewAffiliatesItemAffiliateIdGetResponse(),
	}
	return m
}

// CreateAffiliatesItemAffiliateIdResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAffiliatesItemAffiliateIdResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAffiliatesItemAffiliateIdResponse(), nil
}

// Deprecated: This class is obsolete. Use AffiliatesItemAffiliateIdGetResponseable instead.
type AffiliatesItemAffiliateIdResponseable interface {
	AffiliatesItemAffiliateIdGetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
