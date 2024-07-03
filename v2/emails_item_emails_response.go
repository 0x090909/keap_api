package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use EmailsItemEmailsGetResponseable instead.
type EmailsItemEmailsResponse struct {
	EmailsItemEmailsGetResponse
}

// NewEmailsItemEmailsResponse instantiates a new EmailsItemEmailsResponse and sets the default values.
func NewEmailsItemEmailsResponse() *EmailsItemEmailsResponse {
	m := &EmailsItemEmailsResponse{
		EmailsItemEmailsGetResponse: *NewEmailsItemEmailsGetResponse(),
	}
	return m
}

// CreateEmailsItemEmailsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEmailsItemEmailsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewEmailsItemEmailsResponse(), nil
}

// Deprecated: This class is obsolete. Use EmailsItemEmailsGetResponseable instead.
type EmailsItemEmailsResponseable interface {
	EmailsItemEmailsGetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
