package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use CompaniesModelGetResponseable instead.
type CompaniesModelResponse struct {
	CompaniesModelGetResponse
}

// NewCompaniesModelResponse instantiates a new CompaniesModelResponse and sets the default values.
func NewCompaniesModelResponse() *CompaniesModelResponse {
	m := &CompaniesModelResponse{
		CompaniesModelGetResponse: *NewCompaniesModelGetResponse(),
	}
	return m
}

// CreateCompaniesModelResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCompaniesModelResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCompaniesModelResponse(), nil
}

// Deprecated: This class is obsolete. Use CompaniesModelGetResponseable instead.
type CompaniesModelResponseable interface {
	CompaniesModelGetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
