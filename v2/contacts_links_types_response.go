package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ContactsLinksTypesGetResponseable instead.
type ContactsLinksTypesResponse struct {
	ContactsLinksTypesGetResponse
}

// NewContactsLinksTypesResponse instantiates a new ContactsLinksTypesResponse and sets the default values.
func NewContactsLinksTypesResponse() *ContactsLinksTypesResponse {
	m := &ContactsLinksTypesResponse{
		ContactsLinksTypesGetResponse: *NewContactsLinksTypesGetResponse(),
	}
	return m
}

// CreateContactsLinksTypesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsLinksTypesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsLinksTypesResponse(), nil
}

// Deprecated: This class is obsolete. Use ContactsLinksTypesGetResponseable instead.
type ContactsLinksTypesResponseable interface {
	ContactsLinksTypesGetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
