package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use OpportunitiesItemWithOpportunityGetResponseable instead.
type OpportunitiesItemWithOpportunityResponse struct {
	OpportunitiesItemWithOpportunityGetResponse
}

// NewOpportunitiesItemWithOpportunityResponse instantiates a new OpportunitiesItemWithOpportunityResponse and sets the default values.
func NewOpportunitiesItemWithOpportunityResponse() *OpportunitiesItemWithOpportunityResponse {
	m := &OpportunitiesItemWithOpportunityResponse{
		OpportunitiesItemWithOpportunityGetResponse: *NewOpportunitiesItemWithOpportunityGetResponse(),
	}
	return m
}

// CreateOpportunitiesItemWithOpportunityResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOpportunitiesItemWithOpportunityResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewOpportunitiesItemWithOpportunityResponse(), nil
}

// Deprecated: This class is obsolete. Use OpportunitiesItemWithOpportunityGetResponseable instead.
type OpportunitiesItemWithOpportunityResponseable interface {
	OpportunitiesItemWithOpportunityGetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
