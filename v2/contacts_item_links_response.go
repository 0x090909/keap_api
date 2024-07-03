package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ContactsItemLinksGetResponseable instead.
type ContactsItemLinksResponse struct {
	ContactsItemLinksGetResponse
}

// NewContactsItemLinksResponse instantiates a new ContactsItemLinksResponse and sets the default values.
func NewContactsItemLinksResponse() *ContactsItemLinksResponse {
	m := &ContactsItemLinksResponse{
		ContactsItemLinksGetResponse: *NewContactsItemLinksGetResponse(),
	}
	return m
}

// CreateContactsItemLinksResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsItemLinksResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsItemLinksResponse(), nil
}

// Deprecated: This class is obsolete. Use ContactsItemLinksGetResponseable instead.
type ContactsItemLinksResponseable interface {
	ContactsItemLinksGetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
