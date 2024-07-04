package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ContactsPutResponse_custom_fields struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The content property
	content ContactsPutResponse_custom_fields_contentable
	// The id property
	id *int64
}

// NewContactsPutResponse_custom_fields instantiates a new ContactsPutResponse_custom_fields and sets the default values.
func NewContactsPutResponse_custom_fields() *ContactsPutResponse_custom_fields {
	m := &ContactsPutResponse_custom_fields{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateContactsPutResponse_custom_fieldsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsPutResponse_custom_fieldsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsPutResponse_custom_fields(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContactsPutResponse_custom_fields) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetContent gets the content property value. The content property
// returns a ContactsPutResponse_custom_fields_contentable when successful
func (m *ContactsPutResponse_custom_fields) GetContent() ContactsPutResponse_custom_fields_contentable {
	return m.content
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContactsPutResponse_custom_fields) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["content"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateContactsPutResponse_custom_fields_contentFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetContent(val.(ContactsPutResponse_custom_fields_contentable))
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
	return res
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *ContactsPutResponse_custom_fields) GetId() *int64 {
	return m.id
}

// Serialize serializes information the current object
func (m *ContactsPutResponse_custom_fields) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("content", m.GetContent())
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
		err := writer.WriteAdditionalData(m.GetAdditionalData())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ContactsPutResponse_custom_fields) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetContent sets the content property value. The content property
func (m *ContactsPutResponse_custom_fields) SetContent(value ContactsPutResponse_custom_fields_contentable) {
	m.content = value
}

// SetId sets the id property value. The id property
func (m *ContactsPutResponse_custom_fields) SetId(value *int64) {
	m.id = value
}

type ContactsPutResponse_custom_fieldsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetContent() ContactsPutResponse_custom_fields_contentable
	GetId() *int64
	SetContent(value ContactsPutResponse_custom_fields_contentable)
	SetId(value *int64)
}
