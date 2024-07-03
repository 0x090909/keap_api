package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ContactsItemNotesItemWithNote_GetResponseable instead.
type ContactsItemNotesItemWithNote_Response struct {
	ContactsItemNotesItemWithNote_GetResponse
}

// NewContactsItemNotesItemWithNote_Response instantiates a new ContactsItemNotesItemWithNote_Response and sets the default values.
func NewContactsItemNotesItemWithNote_Response() *ContactsItemNotesItemWithNote_Response {
	m := &ContactsItemNotesItemWithNote_Response{
		ContactsItemNotesItemWithNote_GetResponse: *NewContactsItemNotesItemWithNote_GetResponse(),
	}
	return m
}

// CreateContactsItemNotesItemWithNote_ResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsItemNotesItemWithNote_ResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsItemNotesItemWithNote_Response(), nil
}

// Deprecated: This class is obsolete. Use ContactsItemNotesItemWithNote_GetResponseable instead.
type ContactsItemNotesItemWithNote_Responseable interface {
	ContactsItemNotesItemWithNote_GetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
