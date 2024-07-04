package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use AffiliatesGetResponseable instead.
type AffiliatesResponse struct {
	AffiliatesGetResponse
}

// NewAffiliatesResponse instantiates a new AffiliatesResponse and sets the default values.
func NewAffiliatesResponse() *AffiliatesResponse {
	m := &AffiliatesResponse{
		AffiliatesGetResponse: *NewAffiliatesGetResponse(),
	}
	return m
}

// CreateAffiliatesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAffiliatesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAffiliatesResponse(), nil
}

// Deprecated: This class is obsolete. Use AffiliatesGetResponseable instead.
type AffiliatesResponseable interface {
	AffiliatesGetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
