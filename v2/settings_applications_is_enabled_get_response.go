package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SettingsApplicationsIsEnabledGetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The enabled property
	enabled *bool
}

// NewSettingsApplicationsIsEnabledGetResponse instantiates a new SettingsApplicationsIsEnabledGetResponse and sets the default values.
func NewSettingsApplicationsIsEnabledGetResponse() *SettingsApplicationsIsEnabledGetResponse {
	m := &SettingsApplicationsIsEnabledGetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateSettingsApplicationsIsEnabledGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSettingsApplicationsIsEnabledGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSettingsApplicationsIsEnabledGetResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SettingsApplicationsIsEnabledGetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetEnabled gets the enabled property value. The enabled property
// returns a *bool when successful
func (m *SettingsApplicationsIsEnabledGetResponse) GetEnabled() *bool {
	return m.enabled
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SettingsApplicationsIsEnabledGetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["enabled"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetEnabled(val)
		}
		return nil
	}
	return res
}

// Serialize serializes information the current object
func (m *SettingsApplicationsIsEnabledGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteBoolValue("enabled", m.GetEnabled())
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
func (m *SettingsApplicationsIsEnabledGetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetEnabled sets the enabled property value. The enabled property
func (m *SettingsApplicationsIsEnabledGetResponse) SetEnabled(value *bool) {
	m.enabled = value
}

type SettingsApplicationsIsEnabledGetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetEnabled() *bool
	SetEnabled(value *bool)
}
