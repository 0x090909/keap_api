package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use OpportunitiesModelGetResponseable instead.
type OpportunitiesModelResponse struct {
	OpportunitiesModelGetResponse
}

// NewOpportunitiesModelResponse instantiates a new OpportunitiesModelResponse and sets the default values.
func NewOpportunitiesModelResponse() *OpportunitiesModelResponse {
	m := &OpportunitiesModelResponse{
		OpportunitiesModelGetResponse: *NewOpportunitiesModelGetResponse(),
	}
	return m
}

// CreateOpportunitiesModelResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOpportunitiesModelResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewOpportunitiesModelResponse(), nil
}

// Deprecated: This class is obsolete. Use OpportunitiesModelGetResponseable instead.
type OpportunitiesModelResponseable interface {
	OpportunitiesModelGetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
