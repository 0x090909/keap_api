package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ContactsGetResponse_contacts_links struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The id property
	id *string
	// The link_type_id property
	link_type_id *string
	// The link_type_name property
	link_type_name *string
	// The linked_contact_id property
	linked_contact_id *string
}

// NewContactsGetResponse_contacts_links instantiates a new ContactsGetResponse_contacts_links and sets the default values.
func NewContactsGetResponse_contacts_links() *ContactsGetResponse_contacts_links {
	m := &ContactsGetResponse_contacts_links{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateContactsGetResponse_contacts_linksFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsGetResponse_contacts_linksFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsGetResponse_contacts_links(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContactsGetResponse_contacts_links) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContactsGetResponse_contacts_links) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
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
	res["link_type_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLinkTypeId(val)
		}
		return nil
	}
	res["link_type_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLinkTypeName(val)
		}
		return nil
	}
	res["linked_contact_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLinkedContactId(val)
		}
		return nil
	}
	return res
}

// GetId gets the id property value. The id property
// returns a *string when successful
func (m *ContactsGetResponse_contacts_links) GetId() *string {
	return m.id
}

// GetLinkedContactId gets the linked_contact_id property value. The linked_contact_id property
// returns a *string when successful
func (m *ContactsGetResponse_contacts_links) GetLinkedContactId() *string {
	return m.linked_contact_id
}

// GetLinkTypeId gets the link_type_id property value. The link_type_id property
// returns a *string when successful
func (m *ContactsGetResponse_contacts_links) GetLinkTypeId() *string {
	return m.link_type_id
}

// GetLinkTypeName gets the link_type_name property value. The link_type_name property
// returns a *string when successful
func (m *ContactsGetResponse_contacts_links) GetLinkTypeName() *string {
	return m.link_type_name
}

// Serialize serializes information the current object
func (m *ContactsGetResponse_contacts_links) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("id", m.GetId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("linked_contact_id", m.GetLinkedContactId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("link_type_id", m.GetLinkTypeId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("link_type_name", m.GetLinkTypeName())
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
func (m *ContactsGetResponse_contacts_links) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetId sets the id property value. The id property
func (m *ContactsGetResponse_contacts_links) SetId(value *string) {
	m.id = value
}

// SetLinkedContactId sets the linked_contact_id property value. The linked_contact_id property
func (m *ContactsGetResponse_contacts_links) SetLinkedContactId(value *string) {
	m.linked_contact_id = value
}

// SetLinkTypeId sets the link_type_id property value. The link_type_id property
func (m *ContactsGetResponse_contacts_links) SetLinkTypeId(value *string) {
	m.link_type_id = value
}

// SetLinkTypeName sets the link_type_name property value. The link_type_name property
func (m *ContactsGetResponse_contacts_links) SetLinkTypeName(value *string) {
	m.link_type_name = value
}

type ContactsGetResponse_contacts_linksable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetId() *string
	GetLinkedContactId() *string
	GetLinkTypeId() *string
	GetLinkTypeName() *string
	SetId(value *string)
	SetLinkedContactId(value *string)
	SetLinkTypeId(value *string)
	SetLinkTypeName(value *string)
}
