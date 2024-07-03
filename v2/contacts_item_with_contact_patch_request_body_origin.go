package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ContactsItemWithContact_PatchRequestBody_origin struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The ip_address property
	ip_address *string
}

// NewContactsItemWithContact_PatchRequestBody_origin instantiates a new ContactsItemWithContact_PatchRequestBody_origin and sets the default values.
func NewContactsItemWithContact_PatchRequestBody_origin() *ContactsItemWithContact_PatchRequestBody_origin {
	m := &ContactsItemWithContact_PatchRequestBody_origin{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateContactsItemWithContact_PatchRequestBody_originFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsItemWithContact_PatchRequestBody_originFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsItemWithContact_PatchRequestBody_origin(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContactsItemWithContact_PatchRequestBody_origin) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContactsItemWithContact_PatchRequestBody_origin) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["ip_address"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetIpAddress(val)
		}
		return nil
	}
	return res
}

// GetIpAddress gets the ip_address property value. The ip_address property
// returns a *string when successful
func (m *ContactsItemWithContact_PatchRequestBody_origin) GetIpAddress() *string {
	return m.ip_address
}

// Serialize serializes information the current object
func (m *ContactsItemWithContact_PatchRequestBody_origin) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("ip_address", m.GetIpAddress())
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
func (m *ContactsItemWithContact_PatchRequestBody_origin) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetIpAddress sets the ip_address property value. The ip_address property
func (m *ContactsItemWithContact_PatchRequestBody_origin) SetIpAddress(value *string) {
	m.ip_address = value
}

type ContactsItemWithContact_PatchRequestBody_originable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetIpAddress() *string
	SetIpAddress(value *string)
}
