package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use EmailsSyncPostResponseable instead.
type EmailsSyncResponse struct {
	EmailsSyncPostResponse
}

// NewEmailsSyncResponse instantiates a new EmailsSyncResponse and sets the default values.
func NewEmailsSyncResponse() *EmailsSyncResponse {
	m := &EmailsSyncResponse{
		EmailsSyncPostResponse: *NewEmailsSyncPostResponse(),
	}
	return m
}

// CreateEmailsSyncResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEmailsSyncResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewEmailsSyncResponse(), nil
}

// Deprecated: This class is obsolete. Use EmailsSyncPostResponseable instead.
type EmailsSyncResponseable interface {
	EmailsSyncPostResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
