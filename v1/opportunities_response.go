package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use OpportunitiesGetResponseable instead.
type OpportunitiesResponse struct {
	OpportunitiesGetResponse
}

// NewOpportunitiesResponse instantiates a new OpportunitiesResponse and sets the default values.
func NewOpportunitiesResponse() *OpportunitiesResponse {
	m := &OpportunitiesResponse{
		OpportunitiesGetResponse: *NewOpportunitiesGetResponse(),
	}
	return m
}

// CreateOpportunitiesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOpportunitiesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewOpportunitiesResponse(), nil
}

// Deprecated: This class is obsolete. Use OpportunitiesGetResponseable instead.
type OpportunitiesResponseable interface {
	OpportunitiesGetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
