package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use AffiliatesCommissionsGetResponseable instead.
type AffiliatesCommissionsResponse struct {
	AffiliatesCommissionsGetResponse
}

// NewAffiliatesCommissionsResponse instantiates a new AffiliatesCommissionsResponse and sets the default values.
func NewAffiliatesCommissionsResponse() *AffiliatesCommissionsResponse {
	m := &AffiliatesCommissionsResponse{
		AffiliatesCommissionsGetResponse: *NewAffiliatesCommissionsGetResponse(),
	}
	return m
}

// CreateAffiliatesCommissionsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAffiliatesCommissionsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAffiliatesCommissionsResponse(), nil
}

// Deprecated: This class is obsolete. Use AffiliatesCommissionsGetResponseable instead.
type AffiliatesCommissionsResponseable interface {
	AffiliatesCommissionsGetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
