package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use AffiliatesProgramsGetResponseable instead.
type AffiliatesProgramsResponse struct {
	AffiliatesProgramsGetResponse
}

// NewAffiliatesProgramsResponse instantiates a new AffiliatesProgramsResponse and sets the default values.
func NewAffiliatesProgramsResponse() *AffiliatesProgramsResponse {
	m := &AffiliatesProgramsResponse{
		AffiliatesProgramsGetResponse: *NewAffiliatesProgramsGetResponse(),
	}
	return m
}

// CreateAffiliatesProgramsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAffiliatesProgramsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAffiliatesProgramsResponse(), nil
}

// Deprecated: This class is obsolete. Use AffiliatesProgramsGetResponseable instead.
type AffiliatesProgramsResponseable interface {
	AffiliatesProgramsGetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
