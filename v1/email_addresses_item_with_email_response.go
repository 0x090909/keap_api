package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use EmailAddressesItemWithEmailPutResponseable instead.
type EmailAddressesItemWithEmailResponse struct {
	EmailAddressesItemWithEmailPutResponse
}

// NewEmailAddressesItemWithEmailResponse instantiates a new EmailAddressesItemWithEmailResponse and sets the default values.
func NewEmailAddressesItemWithEmailResponse() *EmailAddressesItemWithEmailResponse {
	m := &EmailAddressesItemWithEmailResponse{
		EmailAddressesItemWithEmailPutResponse: *NewEmailAddressesItemWithEmailPutResponse(),
	}
	return m
}

// CreateEmailAddressesItemWithEmailResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEmailAddressesItemWithEmailResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewEmailAddressesItemWithEmailResponse(), nil
}

// Deprecated: This class is obsolete. Use EmailAddressesItemWithEmailPutResponseable instead.
type EmailAddressesItemWithEmailResponseable interface {
	EmailAddressesItemWithEmailPutResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
