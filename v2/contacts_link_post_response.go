package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ContactsLinkPostResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The contact1_id property
	contact1_id *string
	// The contact2_id property
	contact2_id *string
	// The link_type_id property
	link_type_id *string
	// The link_type_name property
	link_type_name *string
}

// NewContactsLinkPostResponse instantiates a new ContactsLinkPostResponse and sets the default values.
func NewContactsLinkPostResponse() *ContactsLinkPostResponse {
	m := &ContactsLinkPostResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateContactsLinkPostResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsLinkPostResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsLinkPostResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContactsLinkPostResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetContact1Id gets the contact1_id property value. The contact1_id property
// returns a *string when successful
func (m *ContactsLinkPostResponse) GetContact1Id() *string {
	return m.contact1_id
}

// GetContact2Id gets the contact2_id property value. The contact2_id property
// returns a *string when successful
func (m *ContactsLinkPostResponse) GetContact2Id() *string {
	return m.contact2_id
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContactsLinkPostResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["contact1_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetContact1Id(val)
		}
		return nil
	}
	res["contact2_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetContact2Id(val)
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
	return res
}

// GetLinkTypeId gets the link_type_id property value. The link_type_id property
// returns a *string when successful
func (m *ContactsLinkPostResponse) GetLinkTypeId() *string {
	return m.link_type_id
}

// GetLinkTypeName gets the link_type_name property value. The link_type_name property
// returns a *string when successful
func (m *ContactsLinkPostResponse) GetLinkTypeName() *string {
	return m.link_type_name
}

// Serialize serializes information the current object
func (m *ContactsLinkPostResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("contact1_id", m.GetContact1Id())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("contact2_id", m.GetContact2Id())
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
func (m *ContactsLinkPostResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetContact1Id sets the contact1_id property value. The contact1_id property
func (m *ContactsLinkPostResponse) SetContact1Id(value *string) {
	m.contact1_id = value
}

// SetContact2Id sets the contact2_id property value. The contact2_id property
func (m *ContactsLinkPostResponse) SetContact2Id(value *string) {
	m.contact2_id = value
}

// SetLinkTypeId sets the link_type_id property value. The link_type_id property
func (m *ContactsLinkPostResponse) SetLinkTypeId(value *string) {
	m.link_type_id = value
}

// SetLinkTypeName sets the link_type_name property value. The link_type_name property
func (m *ContactsLinkPostResponse) SetLinkTypeName(value *string) {
	m.link_type_name = value
}

type ContactsLinkPostResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetContact1Id() *string
	GetContact2Id() *string
	GetLinkTypeId() *string
	GetLinkTypeName() *string
	SetContact1Id(value *string)
	SetContact2Id(value *string)
	SetLinkTypeId(value *string)
	SetLinkTypeName(value *string)
}
