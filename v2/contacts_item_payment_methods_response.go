package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ContactsItemPaymentMethodsGetResponseable instead.
type ContactsItemPaymentMethodsResponse struct {
	ContactsItemPaymentMethodsGetResponse
}

// NewContactsItemPaymentMethodsResponse instantiates a new ContactsItemPaymentMethodsResponse and sets the default values.
func NewContactsItemPaymentMethodsResponse() *ContactsItemPaymentMethodsResponse {
	m := &ContactsItemPaymentMethodsResponse{
		ContactsItemPaymentMethodsGetResponse: *NewContactsItemPaymentMethodsGetResponse(),
	}
	return m
}

// CreateContactsItemPaymentMethodsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsItemPaymentMethodsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsItemPaymentMethodsResponse(), nil
}

// Deprecated: This class is obsolete. Use ContactsItemPaymentMethodsGetResponseable instead.
type ContactsItemPaymentMethodsResponseable interface {
	ContactsItemPaymentMethodsGetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
