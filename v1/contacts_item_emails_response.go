package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ContactsItemEmailsGetResponseable instead.
type ContactsItemEmailsResponse struct {
	ContactsItemEmailsGetResponse
}

// NewContactsItemEmailsResponse instantiates a new ContactsItemEmailsResponse and sets the default values.
func NewContactsItemEmailsResponse() *ContactsItemEmailsResponse {
	m := &ContactsItemEmailsResponse{
		ContactsItemEmailsGetResponse: *NewContactsItemEmailsGetResponse(),
	}
	return m
}

// CreateContactsItemEmailsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsItemEmailsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsItemEmailsResponse(), nil
}

// Deprecated: This class is obsolete. Use ContactsItemEmailsGetResponseable instead.
type ContactsItemEmailsResponseable interface {
	ContactsItemEmailsGetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
