package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SettingApplicationEnabledGetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The value property
	value *string
}

// NewSettingApplicationEnabledGetResponse instantiates a new SettingApplicationEnabledGetResponse and sets the default values.
func NewSettingApplicationEnabledGetResponse() *SettingApplicationEnabledGetResponse {
	m := &SettingApplicationEnabledGetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateSettingApplicationEnabledGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSettingApplicationEnabledGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSettingApplicationEnabledGetResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SettingApplicationEnabledGetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SettingApplicationEnabledGetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetValue(val)
		}
		return nil
	}
	return res
}

// GetValue gets the value property value. The value property
// returns a *string when successful
func (m *SettingApplicationEnabledGetResponse) GetValue() *string {
	return m.value
}

// Serialize serializes information the current object
func (m *SettingApplicationEnabledGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("value", m.GetValue())
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
func (m *SettingApplicationEnabledGetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetValue sets the value property value. The value property
func (m *SettingApplicationEnabledGetResponse) SetValue(value *string) {
	m.value = value
}

type SettingApplicationEnabledGetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetValue() *string
	SetValue(value *string)
}