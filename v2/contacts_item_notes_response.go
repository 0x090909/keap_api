package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ContactsItemNotesGetResponseable instead.
type ContactsItemNotesResponse struct {
	ContactsItemNotesGetResponse
}

// NewContactsItemNotesResponse instantiates a new ContactsItemNotesResponse and sets the default values.
func NewContactsItemNotesResponse() *ContactsItemNotesResponse {
	m := &ContactsItemNotesResponse{
		ContactsItemNotesGetResponse: *NewContactsItemNotesGetResponse(),
	}
	return m
}

// CreateContactsItemNotesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsItemNotesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsItemNotesResponse(), nil
}

// Deprecated: This class is obsolete. Use ContactsItemNotesGetResponseable instead.
type ContactsItemNotesResponseable interface {
	ContactsItemNotesGetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
