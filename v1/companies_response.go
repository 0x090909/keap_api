package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use CompaniesGetResponseable instead.
type CompaniesResponse struct {
	CompaniesGetResponse
}

// NewCompaniesResponse instantiates a new CompaniesResponse and sets the default values.
func NewCompaniesResponse() *CompaniesResponse {
	m := &CompaniesResponse{
		CompaniesGetResponse: *NewCompaniesGetResponse(),
	}
	return m
}

// CreateCompaniesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCompaniesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCompaniesResponse(), nil
}

// Deprecated: This class is obsolete. Use CompaniesGetResponseable instead.
type CompaniesResponseable interface {
	CompaniesGetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
