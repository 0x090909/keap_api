package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ContactsLinkPostResponseable instead.
type ContactsLinkResponse struct {
	ContactsLinkPostResponse
}

// NewContactsLinkResponse instantiates a new ContactsLinkResponse and sets the default values.
func NewContactsLinkResponse() *ContactsLinkResponse {
	m := &ContactsLinkResponse{
		ContactsLinkPostResponse: *NewContactsLinkPostResponse(),
	}
	return m
}

// CreateContactsLinkResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsLinkResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsLinkResponse(), nil
}

// Deprecated: This class is obsolete. Use ContactsLinkPostResponseable instead.
type ContactsLinkResponseable interface {
	ContactsLinkPostResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
