package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ContactsModelGetResponseable instead.
type ContactsModelResponse struct {
	ContactsModelGetResponse
}

// NewContactsModelResponse instantiates a new ContactsModelResponse and sets the default values.
func NewContactsModelResponse() *ContactsModelResponse {
	m := &ContactsModelResponse{
		ContactsModelGetResponse: *NewContactsModelGetResponse(),
	}
	return m
}

// CreateContactsModelResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsModelResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsModelResponse(), nil
}

// Deprecated: This class is obsolete. Use ContactsModelGetResponseable instead.
type ContactsModelResponseable interface {
	ContactsModelGetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
