package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SettingApplicationConfigurationGetResponse_opportunity_states_loss struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The reasons property
	reasons *string
	// The stage property
	stage *string
}

// NewSettingApplicationConfigurationGetResponse_opportunity_states_loss instantiates a new SettingApplicationConfigurationGetResponse_opportunity_states_loss and sets the default values.
func NewSettingApplicationConfigurationGetResponse_opportunity_states_loss() *SettingApplicationConfigurationGetResponse_opportunity_states_loss {
	m := &SettingApplicationConfigurationGetResponse_opportunity_states_loss{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateSettingApplicationConfigurationGetResponse_opportunity_states_lossFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSettingApplicationConfigurationGetResponse_opportunity_states_lossFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSettingApplicationConfigurationGetResponse_opportunity_states_loss(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SettingApplicationConfigurationGetResponse_opportunity_states_loss) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SettingApplicationConfigurationGetResponse_opportunity_states_loss) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
func (m *SettingApplicationConfigurationGetResponse_opportunity_states_loss) GetReasons() *string {
	return m.reasons
}

// GetStage gets the stage property value. The stage property
// returns a *string when successful
func (m *SettingApplicationConfigurationGetResponse_opportunity_states_loss) GetStage() *string {
	return m.stage
}

// Serialize serializes information the current object
func (m *SettingApplicationConfigurationGetResponse_opportunity_states_loss) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *SettingApplicationConfigurationGetResponse_opportunity_states_loss) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetReasons sets the reasons property value. The reasons property
func (m *SettingApplicationConfigurationGetResponse_opportunity_states_loss) SetReasons(value *string) {
	m.reasons = value
}

// SetStage sets the stage property value. The stage property
func (m *SettingApplicationConfigurationGetResponse_opportunity_states_loss) SetStage(value *string) {
	m.stage = value
}

type SettingApplicationConfigurationGetResponse_opportunity_states_lossable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetReasons() *string
	GetStage() *string
	SetReasons(value *string)
	SetStage(value *string)
}
