package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ContactsUnlinkPostRequestBody struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The contact1_id property
	contact1_id *string
	// The contact2_id property
	contact2_id *string
	// The link_type_id property
	link_type_id *string
}

// NewContactsUnlinkPostRequestBody instantiates a new ContactsUnlinkPostRequestBody and sets the default values.
func NewContactsUnlinkPostRequestBody() *ContactsUnlinkPostRequestBody {
	m := &ContactsUnlinkPostRequestBody{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateContactsUnlinkPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsUnlinkPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsUnlinkPostRequestBody(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContactsUnlinkPostRequestBody) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetContact1Id gets the contact1_id property value. The contact1_id property
// returns a *string when successful
func (m *ContactsUnlinkPostRequestBody) GetContact1Id() *string {
	return m.contact1_id
}

// GetContact2Id gets the contact2_id property value. The contact2_id property
// returns a *string when successful
func (m *ContactsUnlinkPostRequestBody) GetContact2Id() *string {
	return m.contact2_id
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContactsUnlinkPostRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
	return res
}

// GetLinkTypeId gets the link_type_id property value. The link_type_id property
// returns a *string when successful
func (m *ContactsUnlinkPostRequestBody) GetLinkTypeId() *string {
	return m.link_type_id
}

// Serialize serializes information the current object
func (m *ContactsUnlinkPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
		err := writer.WriteAdditionalData(m.GetAdditionalData())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ContactsUnlinkPostRequestBody) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetContact1Id sets the contact1_id property value. The contact1_id property
func (m *ContactsUnlinkPostRequestBody) SetContact1Id(value *string) {
	m.contact1_id = value
}

// SetContact2Id sets the contact2_id property value. The contact2_id property
func (m *ContactsUnlinkPostRequestBody) SetContact2Id(value *string) {
	m.contact2_id = value
}

// SetLinkTypeId sets the link_type_id property value. The link_type_id property
func (m *ContactsUnlinkPostRequestBody) SetLinkTypeId(value *string) {
	m.link_type_id = value
}

type ContactsUnlinkPostRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetContact1Id() *string
	GetContact2Id() *string
	GetLinkTypeId() *string
	SetContact1Id(value *string)
	SetContact2Id(value *string)
	SetLinkTypeId(value *string)
}
