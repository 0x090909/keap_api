package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ContactsItemWithContact_GetResponseable instead.
type ContactsItemWithContact_Response struct {
	ContactsItemWithContact_GetResponse
}

// NewContactsItemWithContact_Response instantiates a new ContactsItemWithContact_Response and sets the default values.
func NewContactsItemWithContact_Response() *ContactsItemWithContact_Response {
	m := &ContactsItemWithContact_Response{
		ContactsItemWithContact_GetResponse: *NewContactsItemWithContact_GetResponse(),
	}
	return m
}

// CreateContactsItemWithContact_ResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsItemWithContact_ResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsItemWithContact_Response(), nil
}

// Deprecated: This class is obsolete. Use ContactsItemWithContact_GetResponseable instead.
type ContactsItemWithContact_Responseable interface {
	ContactsItemWithContact_GetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
