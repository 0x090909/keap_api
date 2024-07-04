package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SettingApplicationConfigurationGetResponse_appointment struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The appointment_types property
	appointment_types *string
	// The share_opportunity_with_related_user property
	share_opportunity_with_related_user *bool
}

// NewSettingApplicationConfigurationGetResponse_appointment instantiates a new SettingApplicationConfigurationGetResponse_appointment and sets the default values.
func NewSettingApplicationConfigurationGetResponse_appointment() *SettingApplicationConfigurationGetResponse_appointment {
	m := &SettingApplicationConfigurationGetResponse_appointment{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateSettingApplicationConfigurationGetResponse_appointmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSettingApplicationConfigurationGetResponse_appointmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSettingApplicationConfigurationGetResponse_appointment(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SettingApplicationConfigurationGetResponse_appointment) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAppointmentTypes gets the appointment_types property value. The appointment_types property
// returns a *string when successful
func (m *SettingApplicationConfigurationGetResponse_appointment) GetAppointmentTypes() *string {
	return m.appointment_types
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SettingApplicationConfigurationGetResponse_appointment) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["appointment_types"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAppointmentTypes(val)
		}
		return nil
	}
	res["share_opportunity_with_related_user"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetShareOpportunityWithRelatedUser(val)
		}
		return nil
	}
	return res
}

// GetShareOpportunityWithRelatedUser gets the share_opportunity_with_related_user property value. The share_opportunity_with_related_user property
// returns a *bool when successful
func (m *SettingApplicationConfigurationGetResponse_appointment) GetShareOpportunityWithRelatedUser() *bool {
	return m.share_opportunity_with_related_user
}

// Serialize serializes information the current object
func (m *SettingApplicationConfigurationGetResponse_appointment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("appointment_types", m.GetAppointmentTypes())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("share_opportunity_with_related_user", m.GetShareOpportunityWithRelatedUser())
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
func (m *SettingApplicationConfigurationGetResponse_appointment) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAppointmentTypes sets the appointment_types property value. The appointment_types property
func (m *SettingApplicationConfigurationGetResponse_appointment) SetAppointmentTypes(value *string) {
	m.appointment_types = value
}

// SetShareOpportunityWithRelatedUser sets the share_opportunity_with_related_user property value. The share_opportunity_with_related_user property
func (m *SettingApplicationConfigurationGetResponse_appointment) SetShareOpportunityWithRelatedUser(value *bool) {
	m.share_opportunity_with_related_user = value
}

type SettingApplicationConfigurationGetResponse_appointmentable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAppointmentTypes() *string
	GetShareOpportunityWithRelatedUser() *bool
	SetAppointmentTypes(value *string)
	SetShareOpportunityWithRelatedUser(value *bool)
}
