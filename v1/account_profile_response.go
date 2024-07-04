package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccountProfileResponse information about the company that owns/uses this account
// Deprecated: This class is obsolete. Use AccountProfileGetResponseable instead.
type AccountProfileResponse struct {
	AccountProfileGetResponse
}

// NewAccountProfileResponse instantiates a new AccountProfileResponse and sets the default values.
func NewAccountProfileResponse() *AccountProfileResponse {
	m := &AccountProfileResponse{
		AccountProfileGetResponse: *NewAccountProfileGetResponse(),
	}
	return m
}

// CreateAccountProfileResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAccountProfileResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAccountProfileResponse(), nil
}

// Deprecated: This class is obsolete. Use AccountProfileGetResponseable instead.
type AccountProfileResponseable interface {
	AccountProfileGetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
