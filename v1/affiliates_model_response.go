package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use AffiliatesModelGetResponseable instead.
type AffiliatesModelResponse struct {
	AffiliatesModelGetResponse
}

// NewAffiliatesModelResponse instantiates a new AffiliatesModelResponse and sets the default values.
func NewAffiliatesModelResponse() *AffiliatesModelResponse {
	m := &AffiliatesModelResponse{
		AffiliatesModelGetResponse: *NewAffiliatesModelGetResponse(),
	}
	return m
}

// CreateAffiliatesModelResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAffiliatesModelResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAffiliatesModelResponse(), nil
}

// Deprecated: This class is obsolete. Use AffiliatesModelGetResponseable instead.
type AffiliatesModelResponseable interface {
	AffiliatesModelGetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
