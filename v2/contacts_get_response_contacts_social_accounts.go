package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ContactsGetResponse_contacts_social_accounts struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The name property
	name *string
}

// NewContactsGetResponse_contacts_social_accounts instantiates a new ContactsGetResponse_contacts_social_accounts and sets the default values.
func NewContactsGetResponse_contacts_social_accounts() *ContactsGetResponse_contacts_social_accounts {
	m := &ContactsGetResponse_contacts_social_accounts{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateContactsGetResponse_contacts_social_accountsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsGetResponse_contacts_social_accountsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsGetResponse_contacts_social_accounts(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContactsGetResponse_contacts_social_accounts) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContactsGetResponse_contacts_social_accounts) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetName(val)
		}
		return nil
	}
	return res
}

// GetName gets the name property value. The name property
// returns a *string when successful
func (m *ContactsGetResponse_contacts_social_accounts) GetName() *string {
	return m.name
}

// Serialize serializes information the current object
func (m *ContactsGetResponse_contacts_social_accounts) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("name", m.GetName())
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
func (m *ContactsGetResponse_contacts_social_accounts) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetName sets the name property value. The name property
func (m *ContactsGetResponse_contacts_social_accounts) SetName(value *string) {
	m.name = value
}

type ContactsGetResponse_contacts_social_accountsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetName() *string
	SetName(value *string)
}
