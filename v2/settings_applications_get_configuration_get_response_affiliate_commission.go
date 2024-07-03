package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SettingsApplicationsGetConfigurationGetResponse_affiliate_commission struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The calculation_type property
	calculation_type *string
	// The levels property
	levels *int32
	// The participant_types property
	participant_types *string
}

// NewSettingsApplicationsGetConfigurationGetResponse_affiliate_commission instantiates a new SettingsApplicationsGetConfigurationGetResponse_affiliate_commission and sets the default values.
func NewSettingsApplicationsGetConfigurationGetResponse_affiliate_commission() *SettingsApplicationsGetConfigurationGetResponse_affiliate_commission {
	m := &SettingsApplicationsGetConfigurationGetResponse_affiliate_commission{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateSettingsApplicationsGetConfigurationGetResponse_affiliate_commissionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSettingsApplicationsGetConfigurationGetResponse_affiliate_commissionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSettingsApplicationsGetConfigurationGetResponse_affiliate_commission(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_affiliate_commission) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCalculationType gets the calculation_type property value. The calculation_type property
// returns a *string when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_affiliate_commission) GetCalculationType() *string {
	return m.calculation_type
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_affiliate_commission) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["calculation_type"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCalculationType(val)
		}
		return nil
	}
	res["levels"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLevels(val)
		}
		return nil
	}
	res["participant_types"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetParticipantTypes(val)
		}
		return nil
	}
	return res
}

// GetLevels gets the levels property value. The levels property
// returns a *int32 when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_affiliate_commission) GetLevels() *int32 {
	return m.levels
}

// GetParticipantTypes gets the participant_types property value. The participant_types property
// returns a *string when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_affiliate_commission) GetParticipantTypes() *string {
	return m.participant_types
}

// Serialize serializes information the current object
func (m *SettingsApplicationsGetConfigurationGetResponse_affiliate_commission) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("calculation_type", m.GetCalculationType())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("levels", m.GetLevels())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("participant_types", m.GetParticipantTypes())
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
func (m *SettingsApplicationsGetConfigurationGetResponse_affiliate_commission) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCalculationType sets the calculation_type property value. The calculation_type property
func (m *SettingsApplicationsGetConfigurationGetResponse_affiliate_commission) SetCalculationType(value *string) {
	m.calculation_type = value
}

// SetLevels sets the levels property value. The levels property
func (m *SettingsApplicationsGetConfigurationGetResponse_affiliate_commission) SetLevels(value *int32) {
	m.levels = value
}

// SetParticipantTypes sets the participant_types property value. The participant_types property
func (m *SettingsApplicationsGetConfigurationGetResponse_affiliate_commission) SetParticipantTypes(value *string) {
	m.participant_types = value
}

type SettingsApplicationsGetConfigurationGetResponse_affiliate_commissionable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCalculationType() *string
	GetLevels() *int32
	GetParticipantTypes() *string
	SetCalculationType(value *string)
	SetLevels(value *int32)
	SetParticipantTypes(value *string)
}
