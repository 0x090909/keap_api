package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use EmailsUnsyncPostResponseable instead.
type EmailsUnsyncResponse struct {
	EmailsUnsyncPostResponse
}

// NewEmailsUnsyncResponse instantiates a new EmailsUnsyncResponse and sets the default values.
func NewEmailsUnsyncResponse() *EmailsUnsyncResponse {
	m := &EmailsUnsyncResponse{
		EmailsUnsyncPostResponse: *NewEmailsUnsyncPostResponse(),
	}
	return m
}

// CreateEmailsUnsyncResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEmailsUnsyncResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewEmailsUnsyncResponse(), nil
}

// Deprecated: This class is obsolete. Use EmailsUnsyncPostResponseable instead.
type EmailsUnsyncResponseable interface {
	EmailsUnsyncPostResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
