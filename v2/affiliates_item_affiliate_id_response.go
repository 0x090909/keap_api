package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use AffiliatesItemAffiliate_idGetResponseable instead.
type AffiliatesItemAffiliate_idResponse struct {
	AffiliatesItemAffiliate_idGetResponse
}

// NewAffiliatesItemAffiliate_idResponse instantiates a new AffiliatesItemAffiliate_idResponse and sets the default values.
func NewAffiliatesItemAffiliate_idResponse() *AffiliatesItemAffiliate_idResponse {
	m := &AffiliatesItemAffiliate_idResponse{
		AffiliatesItemAffiliate_idGetResponse: *NewAffiliatesItemAffiliate_idGetResponse(),
	}
	return m
}

// CreateAffiliatesItemAffiliate_idResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAffiliatesItemAffiliate_idResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAffiliatesItemAffiliate_idResponse(), nil
}

// Deprecated: This class is obsolete. Use AffiliatesItemAffiliate_idGetResponseable instead.
type AffiliatesItemAffiliate_idResponseable interface {
	AffiliatesItemAffiliate_idGetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
