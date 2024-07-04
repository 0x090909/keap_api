package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SettingApplicationConfigurationGetResponse_opportunity_states struct {
	// The active property
	active SettingApplicationConfigurationGetResponse_opportunity_states_activeable
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The loss property
	loss SettingApplicationConfigurationGetResponse_opportunity_states_lossable
	// The win property
	win SettingApplicationConfigurationGetResponse_opportunity_states_winable
}

// NewSettingApplicationConfigurationGetResponse_opportunity_states instantiates a new SettingApplicationConfigurationGetResponse_opportunity_states and sets the default values.
func NewSettingApplicationConfigurationGetResponse_opportunity_states() *SettingApplicationConfigurationGetResponse_opportunity_states {
	m := &SettingApplicationConfigurationGetResponse_opportunity_states{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateSettingApplicationConfigurationGetResponse_opportunity_statesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSettingApplicationConfigurationGetResponse_opportunity_statesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSettingApplicationConfigurationGetResponse_opportunity_states(), nil
}

// GetActive gets the active property value. The active property
// returns a SettingApplicationConfigurationGetResponse_opportunity_states_activeable when successful
func (m *SettingApplicationConfigurationGetResponse_opportunity_states) GetActive() SettingApplicationConfigurationGetResponse_opportunity_states_activeable {
	return m.active
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SettingApplicationConfigurationGetResponse_opportunity_states) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SettingApplicationConfigurationGetResponse_opportunity_states) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["active"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSettingApplicationConfigurationGetResponse_opportunity_states_activeFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetActive(val.(SettingApplicationConfigurationGetResponse_opportunity_states_activeable))
		}
		return nil
	}
	res["loss"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSettingApplicationConfigurationGetResponse_opportunity_states_lossFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLoss(val.(SettingApplicationConfigurationGetResponse_opportunity_states_lossable))
		}
		return nil
	}
	res["win"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSettingApplicationConfigurationGetResponse_opportunity_states_winFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetWin(val.(SettingApplicationConfigurationGetResponse_opportunity_states_winable))
		}
		return nil
	}
	return res
}

// GetLoss gets the loss property value. The loss property
// returns a SettingApplicationConfigurationGetResponse_opportunity_states_lossable when successful
func (m *SettingApplicationConfigurationGetResponse_opportunity_states) GetLoss() SettingApplicationConfigurationGetResponse_opportunity_states_lossable {
	return m.loss
}

// GetWin gets the win property value. The win property
// returns a SettingApplicationConfigurationGetResponse_opportunity_states_winable when successful
func (m *SettingApplicationConfigurationGetResponse_opportunity_states) GetWin() SettingApplicationConfigurationGetResponse_opportunity_states_winable {
	return m.win
}

// Serialize serializes information the current object
func (m *SettingApplicationConfigurationGetResponse_opportunity_states) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("active", m.GetActive())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("loss", m.GetLoss())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("win", m.GetWin())
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

// SetActive sets the active property value. The active property
func (m *SettingApplicationConfigurationGetResponse_opportunity_states) SetActive(value SettingApplicationConfigurationGetResponse_opportunity_states_activeable) {
	m.active = value
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SettingApplicationConfigurationGetResponse_opportunity_states) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetLoss sets the loss property value. The loss property
func (m *SettingApplicationConfigurationGetResponse_opportunity_states) SetLoss(value SettingApplicationConfigurationGetResponse_opportunity_states_lossable) {
	m.loss = value
}

// SetWin sets the win property value. The win property
func (m *SettingApplicationConfigurationGetResponse_opportunity_states) SetWin(value SettingApplicationConfigurationGetResponse_opportunity_states_winable) {
	m.win = value
}

type SettingApplicationConfigurationGetResponse_opportunity_statesable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetActive() SettingApplicationConfigurationGetResponse_opportunity_states_activeable
	GetLoss() SettingApplicationConfigurationGetResponse_opportunity_states_lossable
	GetWin() SettingApplicationConfigurationGetResponse_opportunity_states_winable
	SetActive(value SettingApplicationConfigurationGetResponse_opportunity_states_activeable)
	SetLoss(value SettingApplicationConfigurationGetResponse_opportunity_states_lossable)
	SetWin(value SettingApplicationConfigurationGetResponse_opportunity_states_winable)
}