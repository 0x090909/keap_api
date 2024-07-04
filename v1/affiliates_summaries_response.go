package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use AffiliatesSummariesGetResponseable instead.
type AffiliatesSummariesResponse struct {
	AffiliatesSummariesGetResponse
}

// NewAffiliatesSummariesResponse instantiates a new AffiliatesSummariesResponse and sets the default values.
func NewAffiliatesSummariesResponse() *AffiliatesSummariesResponse {
	m := &AffiliatesSummariesResponse{
		AffiliatesSummariesGetResponse: *NewAffiliatesSummariesGetResponse(),
	}
	return m
}

// CreateAffiliatesSummariesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAffiliatesSummariesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAffiliatesSummariesResponse(), nil
}

// Deprecated: This class is obsolete. Use AffiliatesSummariesGetResponseable instead.
type AffiliatesSummariesResponseable interface {
	AffiliatesSummariesGetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
