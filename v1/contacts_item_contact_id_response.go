package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ContactsItemContactIdPatchResponseable instead.
type ContactsItemContactIdResponse struct {
	ContactsItemContactIdPatchResponse
}

// NewContactsItemContactIdResponse instantiates a new ContactsItemContactIdResponse and sets the default values.
func NewContactsItemContactIdResponse() *ContactsItemContactIdResponse {
	m := &ContactsItemContactIdResponse{
		ContactsItemContactIdPatchResponse: *NewContactsItemContactIdPatchResponse(),
	}
	return m
}

// CreateContactsItemContactIdResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsItemContactIdResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsItemContactIdResponse(), nil
}

// Deprecated: This class is obsolete. Use ContactsItemContactIdPatchResponseable instead.
type ContactsItemContactIdResponseable interface {
	ContactsItemContactIdPatchResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
