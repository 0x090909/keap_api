package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use EmailsBatchAddPostResponseable instead.
type EmailsBatchAddResponse struct {
	EmailsBatchAddPostResponse
}

// NewEmailsBatchAddResponse instantiates a new EmailsBatchAddResponse and sets the default values.
func NewEmailsBatchAddResponse() *EmailsBatchAddResponse {
	m := &EmailsBatchAddResponse{
		EmailsBatchAddPostResponse: *NewEmailsBatchAddPostResponse(),
	}
	return m
}

// CreateEmailsBatchAddResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEmailsBatchAddResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewEmailsBatchAddResponse(), nil
}

// Deprecated: This class is obsolete. Use EmailsBatchAddPostResponseable instead.
type EmailsBatchAddResponseable interface {
	EmailsBatchAddPostResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
