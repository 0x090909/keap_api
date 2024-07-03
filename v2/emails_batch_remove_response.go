package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use EmailsBatchRemovePostResponseable instead.
type EmailsBatchRemoveResponse struct {
	EmailsBatchRemovePostResponse
}

// NewEmailsBatchRemoveResponse instantiates a new EmailsBatchRemoveResponse and sets the default values.
func NewEmailsBatchRemoveResponse() *EmailsBatchRemoveResponse {
	m := &EmailsBatchRemoveResponse{
		EmailsBatchRemovePostResponse: *NewEmailsBatchRemovePostResponse(),
	}
	return m
}

// CreateEmailsBatchRemoveResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEmailsBatchRemoveResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewEmailsBatchRemoveResponse(), nil
}

// Deprecated: This class is obsolete. Use EmailsBatchRemovePostResponseable instead.
type EmailsBatchRemoveResponseable interface {
	EmailsBatchRemovePostResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
