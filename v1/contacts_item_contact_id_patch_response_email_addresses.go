package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ContactsItemContactIdPatchResponse_email_addresses struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The email property
	email *string
}

// NewContactsItemContactIdPatchResponse_email_addresses instantiates a new ContactsItemContactIdPatchResponse_email_addresses and sets the default values.
func NewContactsItemContactIdPatchResponse_email_addresses() *ContactsItemContactIdPatchResponse_email_addresses {
	m := &ContactsItemContactIdPatchResponse_email_addresses{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateContactsItemContactIdPatchResponse_email_addressesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsItemContactIdPatchResponse_email_addressesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsItemContactIdPatchResponse_email_addresses(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContactsItemContactIdPatchResponse_email_addresses) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetEmail gets the email property value. The email property
// returns a *string when successful
func (m *ContactsItemContactIdPatchResponse_email_addresses) GetEmail() *string {
	return m.email
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContactsItemContactIdPatchResponse_email_addresses) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["email"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetEmail(val)
		}
		return nil
	}
	return res
}

// Serialize serializes information the current object
func (m *ContactsItemContactIdPatchResponse_email_addresses) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("email", m.GetEmail())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteAdditionalData(m.GetAdditionalData())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ContactsItemContactIdPatchResponse_email_addresses) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetEmail sets the email property value. The email property
func (m *ContactsItemContactIdPatchResponse_email_addresses) SetEmail(value *string) {
	m.email = value
}

type ContactsItemContactIdPatchResponse_email_addressesable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetEmail() *string
	SetEmail(value *string)
}
