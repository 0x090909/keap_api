package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ContactsLinksTypesGetResponse_contact_link_types struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The id property
	id *string
	// The max_links property
	max_links *int64
	// The name property
	name *string
}

// NewContactsLinksTypesGetResponse_contact_link_types instantiates a new ContactsLinksTypesGetResponse_contact_link_types and sets the default values.
func NewContactsLinksTypesGetResponse_contact_link_types() *ContactsLinksTypesGetResponse_contact_link_types {
	m := &ContactsLinksTypesGetResponse_contact_link_types{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateContactsLinksTypesGetResponse_contact_link_typesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsLinksTypesGetResponse_contact_link_typesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsLinksTypesGetResponse_contact_link_types(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContactsLinksTypesGetResponse_contact_link_types) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContactsLinksTypesGetResponse_contact_link_types) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
	res["max_links"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetMaxLinks(val)
		}
		return nil
	}
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

// GetId gets the id property value. The id property
// returns a *string when successful
func (m *ContactsLinksTypesGetResponse_contact_link_types) GetId() *string {
	return m.id
}

// GetMaxLinks gets the max_links property value. The max_links property
// returns a *int64 when successful
func (m *ContactsLinksTypesGetResponse_contact_link_types) GetMaxLinks() *int64 {
	return m.max_links
}

// GetName gets the name property value. The name property
// returns a *string when successful
func (m *ContactsLinksTypesGetResponse_contact_link_types) GetName() *string {
	return m.name
}

// Serialize serializes information the current object
func (m *ContactsLinksTypesGetResponse_contact_link_types) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("id", m.GetId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("max_links", m.GetMaxLinks())
		if err != nil {
			return err
		}
	}
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
func (m *ContactsLinksTypesGetResponse_contact_link_types) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetId sets the id property value. The id property
func (m *ContactsLinksTypesGetResponse_contact_link_types) SetId(value *string) {
	m.id = value
}

// SetMaxLinks sets the max_links property value. The max_links property
func (m *ContactsLinksTypesGetResponse_contact_link_types) SetMaxLinks(value *int64) {
	m.max_links = value
}

// SetName sets the name property value. The name property
func (m *ContactsLinksTypesGetResponse_contact_link_types) SetName(value *string) {
	m.name = value
}

type ContactsLinksTypesGetResponse_contact_link_typesable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetId() *string
	GetMaxLinks() *int64
	GetName() *string
	SetId(value *string)
	SetMaxLinks(value *int64)
	SetName(value *string)
}
