package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use CompaniesItemWithCompany_GetResponseable instead.
type CompaniesItemWithCompany_Response struct {
	CompaniesItemWithCompany_GetResponse
}

// NewCompaniesItemWithCompany_Response instantiates a new CompaniesItemWithCompany_Response and sets the default values.
func NewCompaniesItemWithCompany_Response() *CompaniesItemWithCompany_Response {
	m := &CompaniesItemWithCompany_Response{
		CompaniesItemWithCompany_GetResponse: *NewCompaniesItemWithCompany_GetResponse(),
	}
	return m
}

// CreateCompaniesItemWithCompany_ResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCompaniesItemWithCompany_ResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCompaniesItemWithCompany_Response(), nil
}

// Deprecated: This class is obsolete. Use CompaniesItemWithCompany_GetResponseable instead.
type CompaniesItemWithCompany_Responseable interface {
	CompaniesItemWithCompany_GetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
