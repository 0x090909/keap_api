package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ContactsModelCustomFieldsPostResponseable instead.
type ContactsModelCustomFieldsResponse struct {
	ContactsModelCustomFieldsPostResponse
}

// NewContactsModelCustomFieldsResponse instantiates a new ContactsModelCustomFieldsResponse and sets the default values.
func NewContactsModelCustomFieldsResponse() *ContactsModelCustomFieldsResponse {
	m := &ContactsModelCustomFieldsResponse{
		ContactsModelCustomFieldsPostResponse: *NewContactsModelCustomFieldsPostResponse(),
	}
	return m
}

// CreateContactsModelCustomFieldsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsModelCustomFieldsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsModelCustomFieldsResponse(), nil
}

// Deprecated: This class is obsolete. Use ContactsModelCustomFieldsPostResponseable instead.
type ContactsModelCustomFieldsResponseable interface {
	ContactsModelCustomFieldsPostResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
