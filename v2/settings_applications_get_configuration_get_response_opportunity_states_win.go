package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SettingsApplicationsGetConfigurationGetResponse_opportunity_states_win struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The reasons property
	reasons *string
	// The stage property
	stage *string
}

// NewSettingsApplicationsGetConfigurationGetResponse_opportunity_states_win instantiates a new SettingsApplicationsGetConfigurationGetResponse_opportunity_states_win and sets the default values.
func NewSettingsApplicationsGetConfigurationGetResponse_opportunity_states_win() *SettingsApplicationsGetConfigurationGetResponse_opportunity_states_win {
	m := &SettingsApplicationsGetConfigurationGetResponse_opportunity_states_win{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateSettingsApplicationsGetConfigurationGetResponse_opportunity_states_winFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSettingsApplicationsGetConfigurationGetResponse_opportunity_states_winFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSettingsApplicationsGetConfigurationGetResponse_opportunity_states_win(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_opportunity_states_win) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_opportunity_states_win) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["reasons"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetReasons(val)
		}
		return nil
	}
	res["stage"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetStage(val)
		}
		return nil
	}
	return res
}

// GetReasons gets the reasons property value. The reasons property
// returns a *string when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_opportunity_states_win) GetReasons() *string {
	return m.reasons
}

// GetStage gets the stage property value. The stage property
// returns a *string when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_opportunity_states_win) GetStage() *string {
	return m.stage
}

// Serialize serializes information the current object
func (m *SettingsApplicationsGetConfigurationGetResponse_opportunity_states_win) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("reasons", m.GetReasons())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("stage", m.GetStage())
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
func (m *SettingsApplicationsGetConfigurationGetResponse_opportunity_states_win) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetReasons sets the reasons property value. The reasons property
func (m *SettingsApplicationsGetConfigurationGetResponse_opportunity_states_win) SetReasons(value *string) {
	m.reasons = value
}

// SetStage sets the stage property value. The stage property
func (m *SettingsApplicationsGetConfigurationGetResponse_opportunity_states_win) SetStage(value *string) {
	m.stage = value
}

type SettingsApplicationsGetConfigurationGetResponse_opportunity_states_winable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetReasons() *string
	GetStage() *string
	SetReasons(value *string)
	SetStage(value *string)
}
