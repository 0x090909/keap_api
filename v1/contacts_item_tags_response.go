package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ContactsItemTagsGetResponseable instead.
type ContactsItemTagsResponse struct {
	ContactsItemTagsGetResponse
}

// NewContactsItemTagsResponse instantiates a new ContactsItemTagsResponse and sets the default values.
func NewContactsItemTagsResponse() *ContactsItemTagsResponse {
	m := &ContactsItemTagsResponse{
		ContactsItemTagsGetResponse: *NewContactsItemTagsGetResponse(),
	}
	return m
}

// CreateContactsItemTagsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsItemTagsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsItemTagsResponse(), nil
}

// Deprecated: This class is obsolete. Use ContactsItemTagsGetResponseable instead.
type ContactsItemTagsResponseable interface {
	ContactsItemTagsGetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
