package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TagsItemContactsGetResponse_contacts_contact struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The email property
	email *string
	// The first_name property
	first_name *string
	// The id property
	id *int64
	// The last_name property
	last_name *string
}

// NewTagsItemContactsGetResponse_contacts_contact instantiates a new TagsItemContactsGetResponse_contacts_contact and sets the default values.
func NewTagsItemContactsGetResponse_contacts_contact() *TagsItemContactsGetResponse_contacts_contact {
	m := &TagsItemContactsGetResponse_contacts_contact{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateTagsItemContactsGetResponse_contacts_contactFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTagsItemContactsGetResponse_contacts_contactFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTagsItemContactsGetResponse_contacts_contact(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TagsItemContactsGetResponse_contacts_contact) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetEmail gets the email property value. The email property
// returns a *string when successful
func (m *TagsItemContactsGetResponse_contacts_contact) GetEmail() *string {
	return m.email
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TagsItemContactsGetResponse_contacts_contact) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
	res["first_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFirstName(val)
		}
		return nil
	}
	res["id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetId(val)
		}
		return nil
	}
	res["last_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLastName(val)
		}
		return nil
	}
	return res
}

// GetFirstName gets the first_name property value. The first_name property
// returns a *string when successful
func (m *TagsItemContactsGetResponse_contacts_contact) GetFirstName() *string {
	return m.first_name
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *TagsItemContactsGetResponse_contacts_contact) GetId() *int64 {
	return m.id
}

// GetLastName gets the last_name property value. The last_name property
// returns a *string when successful
func (m *TagsItemContactsGetResponse_contacts_contact) GetLastName() *string {
	return m.last_name
}

// Serialize serializes information the current object
func (m *TagsItemContactsGetResponse_contacts_contact) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("email", m.GetEmail())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("first_name", m.GetFirstName())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("id", m.GetId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("last_name", m.GetLastName())
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
func (m *TagsItemContactsGetResponse_contacts_contact) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetEmail sets the email property value. The email property
func (m *TagsItemContactsGetResponse_contacts_contact) SetEmail(value *string) {
	m.email = value
}

// SetFirstName sets the first_name property value. The first_name property
func (m *TagsItemContactsGetResponse_contacts_contact) SetFirstName(value *string) {
	m.first_name = value
}

// SetId sets the id property value. The id property
func (m *TagsItemContactsGetResponse_contacts_contact) SetId(value *int64) {
	m.id = value
}

// SetLastName sets the last_name property value. The last_name property
func (m *TagsItemContactsGetResponse_contacts_contact) SetLastName(value *string) {
	m.last_name = value
}

type TagsItemContactsGetResponse_contacts_contactable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetEmail() *string
	GetFirstName() *string
	GetId() *int64
	GetLastName() *string
	SetEmail(value *string)
	SetFirstName(value *string)
	SetId(value *int64)
	SetLastName(value *string)
}
