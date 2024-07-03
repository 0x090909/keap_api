package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ContactsItemNotesItemWithNote_GetResponse_assigned_to_user struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The email_address property
	email_address *string
	// The family_name property
	family_name *string
	// The given_name property
	given_name *string
	// The id property
	id *string
}

// NewContactsItemNotesItemWithNote_GetResponse_assigned_to_user instantiates a new ContactsItemNotesItemWithNote_GetResponse_assigned_to_user and sets the default values.
func NewContactsItemNotesItemWithNote_GetResponse_assigned_to_user() *ContactsItemNotesItemWithNote_GetResponse_assigned_to_user {
	m := &ContactsItemNotesItemWithNote_GetResponse_assigned_to_user{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateContactsItemNotesItemWithNote_GetResponse_assigned_to_userFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsItemNotesItemWithNote_GetResponse_assigned_to_userFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsItemNotesItemWithNote_GetResponse_assigned_to_user(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContactsItemNotesItemWithNote_GetResponse_assigned_to_user) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetEmailAddress gets the email_address property value. The email_address property
// returns a *string when successful
func (m *ContactsItemNotesItemWithNote_GetResponse_assigned_to_user) GetEmailAddress() *string {
	return m.email_address
}

// GetFamilyName gets the family_name property value. The family_name property
// returns a *string when successful
func (m *ContactsItemNotesItemWithNote_GetResponse_assigned_to_user) GetFamilyName() *string {
	return m.family_name
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContactsItemNotesItemWithNote_GetResponse_assigned_to_user) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["email_address"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetEmailAddress(val)
		}
		return nil
	}
	res["family_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFamilyName(val)
		}
		return nil
	}
	res["given_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetGivenName(val)
		}
		return nil
	}
	res["id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetId(val)
		}
		return nil
	}
	return res
}

// GetGivenName gets the given_name property value. The given_name property
// returns a *string when successful
func (m *ContactsItemNotesItemWithNote_GetResponse_assigned_to_user) GetGivenName() *string {
	return m.given_name
}

// GetId gets the id property value. The id property
// returns a *string when successful
func (m *ContactsItemNotesItemWithNote_GetResponse_assigned_to_user) GetId() *string {
	return m.id
}

// Serialize serializes information the current object
func (m *ContactsItemNotesItemWithNote_GetResponse_assigned_to_user) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("email_address", m.GetEmailAddress())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("family_name", m.GetFamilyName())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("given_name", m.GetGivenName())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("id", m.GetId())
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
func (m *ContactsItemNotesItemWithNote_GetResponse_assigned_to_user) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetEmailAddress sets the email_address property value. The email_address property
func (m *ContactsItemNotesItemWithNote_GetResponse_assigned_to_user) SetEmailAddress(value *string) {
	m.email_address = value
}

// SetFamilyName sets the family_name property value. The family_name property
func (m *ContactsItemNotesItemWithNote_GetResponse_assigned_to_user) SetFamilyName(value *string) {
	m.family_name = value
}

// SetGivenName sets the given_name property value. The given_name property
func (m *ContactsItemNotesItemWithNote_GetResponse_assigned_to_user) SetGivenName(value *string) {
	m.given_name = value
}

// SetId sets the id property value. The id property
func (m *ContactsItemNotesItemWithNote_GetResponse_assigned_to_user) SetId(value *string) {
	m.id = value
}

type ContactsItemNotesItemWithNote_GetResponse_assigned_to_userable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetEmailAddress() *string
	GetFamilyName() *string
	GetGivenName() *string
	GetId() *string
	SetEmailAddress(value *string)
	SetFamilyName(value *string)
	SetGivenName(value *string)
	SetId(value *string)
}
