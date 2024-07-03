package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SettingsApplicationsGetConfigurationGetResponse_opportunity struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The default_stage property
	default_stage *string
	// The states property
	states SettingsApplicationsGetConfigurationGetResponse_opportunity_statesable
}

// NewSettingsApplicationsGetConfigurationGetResponse_opportunity instantiates a new SettingsApplicationsGetConfigurationGetResponse_opportunity and sets the default values.
func NewSettingsApplicationsGetConfigurationGetResponse_opportunity() *SettingsApplicationsGetConfigurationGetResponse_opportunity {
	m := &SettingsApplicationsGetConfigurationGetResponse_opportunity{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateSettingsApplicationsGetConfigurationGetResponse_opportunityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSettingsApplicationsGetConfigurationGetResponse_opportunityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSettingsApplicationsGetConfigurationGetResponse_opportunity(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_opportunity) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetDefaultStage gets the default_stage property value. The default_stage property
// returns a *string when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_opportunity) GetDefaultStage() *string {
	return m.default_stage
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_opportunity) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["default_stage"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDefaultStage(val)
		}
		return nil
	}
	res["states"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSettingsApplicationsGetConfigurationGetResponse_opportunity_statesFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetStates(val.(SettingsApplicationsGetConfigurationGetResponse_opportunity_statesable))
		}
		return nil
	}
	return res
}

// GetStates gets the states property value. The states property
// returns a SettingsApplicationsGetConfigurationGetResponse_opportunity_statesable when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_opportunity) GetStates() SettingsApplicationsGetConfigurationGetResponse_opportunity_statesable {
	return m.states
}

// Serialize serializes information the current object
func (m *SettingsApplicationsGetConfigurationGetResponse_opportunity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("default_stage", m.GetDefaultStage())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("states", m.GetStates())
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
func (m *SettingsApplicationsGetConfigurationGetResponse_opportunity) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetDefaultStage sets the default_stage property value. The default_stage property
func (m *SettingsApplicationsGetConfigurationGetResponse_opportunity) SetDefaultStage(value *string) {
	m.default_stage = value
}

// SetStates sets the states property value. The states property
func (m *SettingsApplicationsGetConfigurationGetResponse_opportunity) SetStates(value SettingsApplicationsGetConfigurationGetResponse_opportunity_statesable) {
	m.states = value
}

type SettingsApplicationsGetConfigurationGetResponse_opportunityable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetDefaultStage() *string
	GetStates() SettingsApplicationsGetConfigurationGetResponse_opportunity_statesable
	SetDefaultStage(value *string)
	SetStates(value SettingsApplicationsGetConfigurationGetResponse_opportunity_statesable)
}
