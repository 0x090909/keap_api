package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SettingsApplicationsGetConfigurationGetResponse_application_features_enabled struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The marketing property
	marketing *bool
}

// NewSettingsApplicationsGetConfigurationGetResponse_application_features_enabled instantiates a new SettingsApplicationsGetConfigurationGetResponse_application_features_enabled and sets the default values.
func NewSettingsApplicationsGetConfigurationGetResponse_application_features_enabled() *SettingsApplicationsGetConfigurationGetResponse_application_features_enabled {
	m := &SettingsApplicationsGetConfigurationGetResponse_application_features_enabled{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateSettingsApplicationsGetConfigurationGetResponse_application_features_enabledFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSettingsApplicationsGetConfigurationGetResponse_application_features_enabledFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSettingsApplicationsGetConfigurationGetResponse_application_features_enabled(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_application_features_enabled) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_application_features_enabled) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["marketing"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetMarketing(val)
		}
		return nil
	}
	return res
}

// GetMarketing gets the marketing property value. The marketing property
// returns a *bool when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_application_features_enabled) GetMarketing() *bool {
	return m.marketing
}

// Serialize serializes information the current object
func (m *SettingsApplicationsGetConfigurationGetResponse_application_features_enabled) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteBoolValue("marketing", m.GetMarketing())
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
func (m *SettingsApplicationsGetConfigurationGetResponse_application_features_enabled) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetMarketing sets the marketing property value. The marketing property
func (m *SettingsApplicationsGetConfigurationGetResponse_application_features_enabled) SetMarketing(value *bool) {
	m.marketing = value
}

type SettingsApplicationsGetConfigurationGetResponse_application_features_enabledable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetMarketing() *bool
	SetMarketing(value *bool)
}
