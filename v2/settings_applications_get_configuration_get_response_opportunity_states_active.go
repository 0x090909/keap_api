package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SettingsApplicationsGetConfigurationGetResponse_opportunity_states_active struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The stages property
	stages *string
}

// NewSettingsApplicationsGetConfigurationGetResponse_opportunity_states_active instantiates a new SettingsApplicationsGetConfigurationGetResponse_opportunity_states_active and sets the default values.
func NewSettingsApplicationsGetConfigurationGetResponse_opportunity_states_active() *SettingsApplicationsGetConfigurationGetResponse_opportunity_states_active {
	m := &SettingsApplicationsGetConfigurationGetResponse_opportunity_states_active{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateSettingsApplicationsGetConfigurationGetResponse_opportunity_states_activeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSettingsApplicationsGetConfigurationGetResponse_opportunity_states_activeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSettingsApplicationsGetConfigurationGetResponse_opportunity_states_active(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_opportunity_states_active) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_opportunity_states_active) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["stages"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetStages(val)
		}
		return nil
	}
	return res
}

// GetStages gets the stages property value. The stages property
// returns a *string when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_opportunity_states_active) GetStages() *string {
	return m.stages
}

// Serialize serializes information the current object
func (m *SettingsApplicationsGetConfigurationGetResponse_opportunity_states_active) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("stages", m.GetStages())
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
func (m *SettingsApplicationsGetConfigurationGetResponse_opportunity_states_active) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetStages sets the stages property value. The stages property
func (m *SettingsApplicationsGetConfigurationGetResponse_opportunity_states_active) SetStages(value *string) {
	m.stages = value
}

type SettingsApplicationsGetConfigurationGetResponse_opportunity_states_activeable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetStages() *string
	SetStages(value *string)
}
