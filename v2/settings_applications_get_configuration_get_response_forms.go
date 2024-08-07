package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SettingsApplicationsGetConfigurationGetResponse_forms struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The spam_filters_enabled property
	spam_filters_enabled *bool
}

// NewSettingsApplicationsGetConfigurationGetResponse_forms instantiates a new SettingsApplicationsGetConfigurationGetResponse_forms and sets the default values.
func NewSettingsApplicationsGetConfigurationGetResponse_forms() *SettingsApplicationsGetConfigurationGetResponse_forms {
	m := &SettingsApplicationsGetConfigurationGetResponse_forms{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateSettingsApplicationsGetConfigurationGetResponse_formsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSettingsApplicationsGetConfigurationGetResponse_formsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSettingsApplicationsGetConfigurationGetResponse_forms(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_forms) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_forms) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["spam_filters_enabled"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSpamFiltersEnabled(val)
		}
		return nil
	}
	return res
}

// GetSpamFiltersEnabled gets the spam_filters_enabled property value. The spam_filters_enabled property
// returns a *bool when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_forms) GetSpamFiltersEnabled() *bool {
	return m.spam_filters_enabled
}

// Serialize serializes information the current object
func (m *SettingsApplicationsGetConfigurationGetResponse_forms) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteBoolValue("spam_filters_enabled", m.GetSpamFiltersEnabled())
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
func (m *SettingsApplicationsGetConfigurationGetResponse_forms) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetSpamFiltersEnabled sets the spam_filters_enabled property value. The spam_filters_enabled property
func (m *SettingsApplicationsGetConfigurationGetResponse_forms) SetSpamFiltersEnabled(value *bool) {
	m.spam_filters_enabled = value
}

type SettingsApplicationsGetConfigurationGetResponse_formsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetSpamFiltersEnabled() *bool
	SetSpamFiltersEnabled(value *bool)
}
