package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TagsItemContactsGetResponse_contacts struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The email property
	email *string
	// The family_name property
	family_name *string
	// The given_name property
	given_name *string
	// The id property
	id *string
}

// NewTagsItemContactsGetResponse_contacts instantiates a new TagsItemContactsGetResponse_contacts and sets the default values.
func NewTagsItemContactsGetResponse_contacts() *TagsItemContactsGetResponse_contacts {
	m := &TagsItemContactsGetResponse_contacts{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateTagsItemContactsGetResponse_contactsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTagsItemContactsGetResponse_contactsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTagsItemContactsGetResponse_contacts(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TagsItemContactsGetResponse_contacts) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetEmail gets the email property value. The email property
// returns a *string when successful
func (m *TagsItemContactsGetResponse_contacts) GetEmail() *string {
	return m.email
}

// GetFamilyName gets the family_name property value. The family_name property
// returns a *string when successful
func (m *TagsItemContactsGetResponse_contacts) GetFamilyName() *string {
	return m.family_name
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TagsItemContactsGetResponse_contacts) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
func (m *TagsItemContactsGetResponse_contacts) GetGivenName() *string {
	return m.given_name
}

// GetId gets the id property value. The id property
// returns a *string when successful
func (m *TagsItemContactsGetResponse_contacts) GetId() *string {
	return m.id
}

// Serialize serializes information the current object
func (m *TagsItemContactsGetResponse_contacts) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("email", m.GetEmail())
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
func (m *TagsItemContactsGetResponse_contacts) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetEmail sets the email property value. The email property
func (m *TagsItemContactsGetResponse_contacts) SetEmail(value *string) {
	m.email = value
}

// SetFamilyName sets the family_name property value. The family_name property
func (m *TagsItemContactsGetResponse_contacts) SetFamilyName(value *string) {
	m.family_name = value
}

// SetGivenName sets the given_name property value. The given_name property
func (m *TagsItemContactsGetResponse_contacts) SetGivenName(value *string) {
	m.given_name = value
}

// SetId sets the id property value. The id property
func (m *TagsItemContactsGetResponse_contacts) SetId(value *string) {
	m.id = value
}

type TagsItemContactsGetResponse_contactsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetEmail() *string
	GetFamilyName() *string
	GetGivenName() *string
	GetId() *string
	SetEmail(value *string)
	SetFamilyName(value *string)
	SetGivenName(value *string)
	SetId(value *string)
}
