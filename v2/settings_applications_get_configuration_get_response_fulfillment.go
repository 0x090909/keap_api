package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SettingsApplicationsGetConfigurationGetResponse_fulfillment struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The default_message_fields property
	default_message_fields *string
	// The default_message_to property
	default_message_to *string
}

// NewSettingsApplicationsGetConfigurationGetResponse_fulfillment instantiates a new SettingsApplicationsGetConfigurationGetResponse_fulfillment and sets the default values.
func NewSettingsApplicationsGetConfigurationGetResponse_fulfillment() *SettingsApplicationsGetConfigurationGetResponse_fulfillment {
	m := &SettingsApplicationsGetConfigurationGetResponse_fulfillment{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateSettingsApplicationsGetConfigurationGetResponse_fulfillmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSettingsApplicationsGetConfigurationGetResponse_fulfillmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSettingsApplicationsGetConfigurationGetResponse_fulfillment(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_fulfillment) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetDefaultMessageFields gets the default_message_fields property value. The default_message_fields property
// returns a *string when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_fulfillment) GetDefaultMessageFields() *string {
	return m.default_message_fields
}

// GetDefaultMessageTo gets the default_message_to property value. The default_message_to property
// returns a *string when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_fulfillment) GetDefaultMessageTo() *string {
	return m.default_message_to
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_fulfillment) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["default_message_fields"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDefaultMessageFields(val)
		}
		return nil
	}
	res["default_message_to"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDefaultMessageTo(val)
		}
		return nil
	}
	return res
}

// Serialize serializes information the current object
func (m *SettingsApplicationsGetConfigurationGetResponse_fulfillment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("default_message_fields", m.GetDefaultMessageFields())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("default_message_to", m.GetDefaultMessageTo())
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
func (m *SettingsApplicationsGetConfigurationGetResponse_fulfillment) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetDefaultMessageFields sets the default_message_fields property value. The default_message_fields property
func (m *SettingsApplicationsGetConfigurationGetResponse_fulfillment) SetDefaultMessageFields(value *string) {
	m.default_message_fields = value
}

// SetDefaultMessageTo sets the default_message_to property value. The default_message_to property
func (m *SettingsApplicationsGetConfigurationGetResponse_fulfillment) SetDefaultMessageTo(value *string) {
	m.default_message_to = value
}

type SettingsApplicationsGetConfigurationGetResponse_fulfillmentable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetDefaultMessageFields() *string
	GetDefaultMessageTo() *string
	SetDefaultMessageFields(value *string)
	SetDefaultMessageTo(value *string)
}
