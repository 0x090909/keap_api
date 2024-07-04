package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use UsersGetResponseable instead.
type UsersResponse struct {
	UsersGetResponse
}

// NewUsersResponse instantiates a new UsersResponse and sets the default values.
func NewUsersResponse() *UsersResponse {
	m := &UsersResponse{
		UsersGetResponse: *NewUsersGetResponse(),
	}
	return m
}

// CreateUsersResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUsersResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewUsersResponse(), nil
}

// Deprecated: This class is obsolete. Use UsersGetResponseable instead.
type UsersResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	UsersGetResponseable
}
