package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ContactsItemWithContact_GetResponse_custom_fields struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The content property
	content ContactsItemWithContact_GetResponse_custom_fields_contentable
	// The id property
	id *string
}

// NewContactsItemWithContact_GetResponse_custom_fields instantiates a new ContactsItemWithContact_GetResponse_custom_fields and sets the default values.
func NewContactsItemWithContact_GetResponse_custom_fields() *ContactsItemWithContact_GetResponse_custom_fields {
	m := &ContactsItemWithContact_GetResponse_custom_fields{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateContactsItemWithContact_GetResponse_custom_fieldsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsItemWithContact_GetResponse_custom_fieldsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsItemWithContact_GetResponse_custom_fields(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContactsItemWithContact_GetResponse_custom_fields) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetContent gets the content property value. The content property
// returns a ContactsItemWithContact_GetResponse_custom_fields_contentable when successful
func (m *ContactsItemWithContact_GetResponse_custom_fields) GetContent() ContactsItemWithContact_GetResponse_custom_fields_contentable {
	return m.content
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContactsItemWithContact_GetResponse_custom_fields) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["content"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateContactsItemWithContact_GetResponse_custom_fields_contentFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetContent(val.(ContactsItemWithContact_GetResponse_custom_fields_contentable))
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

// GetId gets the id property value. The id property
// returns a *string when successful
func (m *ContactsItemWithContact_GetResponse_custom_fields) GetId() *string {
	return m.id
}

// Serialize serializes information the current object
func (m *ContactsItemWithContact_GetResponse_custom_fields) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("content", m.GetContent())
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
func (m *ContactsItemWithContact_GetResponse_custom_fields) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetContent sets the content property value. The content property
func (m *ContactsItemWithContact_GetResponse_custom_fields) SetContent(value ContactsItemWithContact_GetResponse_custom_fields_contentable) {
	m.content = value
}

// SetId sets the id property value. The id property
func (m *ContactsItemWithContact_GetResponse_custom_fields) SetId(value *string) {
	m.id = value
}

type ContactsItemWithContact_GetResponse_custom_fieldsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetContent() ContactsItemWithContact_GetResponse_custom_fields_contentable
	GetId() *string
	SetContent(value ContactsItemWithContact_GetResponse_custom_fields_contentable)
	SetId(value *string)
}
