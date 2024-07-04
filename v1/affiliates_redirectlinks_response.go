package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use AffiliatesRedirectlinksGetResponseable instead.
type AffiliatesRedirectlinksResponse struct {
	AffiliatesRedirectlinksGetResponse
}

// NewAffiliatesRedirectlinksResponse instantiates a new AffiliatesRedirectlinksResponse and sets the default values.
func NewAffiliatesRedirectlinksResponse() *AffiliatesRedirectlinksResponse {
	m := &AffiliatesRedirectlinksResponse{
		AffiliatesRedirectlinksGetResponse: *NewAffiliatesRedirectlinksGetResponse(),
	}
	return m
}

// CreateAffiliatesRedirectlinksResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAffiliatesRedirectlinksResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAffiliatesRedirectlinksResponse(), nil
}

// Deprecated: This class is obsolete. Use AffiliatesRedirectlinksGetResponseable instead.
type AffiliatesRedirectlinksResponseable interface {
	AffiliatesRedirectlinksGetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
