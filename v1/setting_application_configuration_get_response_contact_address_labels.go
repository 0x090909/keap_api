package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SettingApplicationConfigurationGetResponse_contact_address_labels struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The line_1 property
	line_1 *string
	// The line_2 property
	line_2 *string
	// The line_3 property
	line_3 *string
}

// NewSettingApplicationConfigurationGetResponse_contact_address_labels instantiates a new SettingApplicationConfigurationGetResponse_contact_address_labels and sets the default values.
func NewSettingApplicationConfigurationGetResponse_contact_address_labels() *SettingApplicationConfigurationGetResponse_contact_address_labels {
	m := &SettingApplicationConfigurationGetResponse_contact_address_labels{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateSettingApplicationConfigurationGetResponse_contact_address_labelsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSettingApplicationConfigurationGetResponse_contact_address_labelsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSettingApplicationConfigurationGetResponse_contact_address_labels(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SettingApplicationConfigurationGetResponse_contact_address_labels) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SettingApplicationConfigurationGetResponse_contact_address_labels) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["line_1"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLine1(val)
		}
		return nil
	}
	res["line_2"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLine2(val)
		}
		return nil
	}
	res["line_3"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLine3(val)
		}
		return nil
	}
	return res
}

// GetLine1 gets the line_1 property value. The line_1 property
// returns a *string when successful
func (m *SettingApplicationConfigurationGetResponse_contact_address_labels) GetLine1() *string {
	return m.line_1
}

// GetLine2 gets the line_2 property value. The line_2 property
// returns a *string when successful
func (m *SettingApplicationConfigurationGetResponse_contact_address_labels) GetLine2() *string {
	return m.line_2
}

// GetLine3 gets the line_3 property value. The line_3 property
// returns a *string when successful
func (m *SettingApplicationConfigurationGetResponse_contact_address_labels) GetLine3() *string {
	return m.line_3
}

// Serialize serializes information the current object
func (m *SettingApplicationConfigurationGetResponse_contact_address_labels) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("line_1", m.GetLine1())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("line_2", m.GetLine2())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("line_3", m.GetLine3())
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
func (m *SettingApplicationConfigurationGetResponse_contact_address_labels) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetLine1 sets the line_1 property value. The line_1 property
func (m *SettingApplicationConfigurationGetResponse_contact_address_labels) SetLine1(value *string) {
	m.line_1 = value
}

// SetLine2 sets the line_2 property value. The line_2 property
func (m *SettingApplicationConfigurationGetResponse_contact_address_labels) SetLine2(value *string) {
	m.line_2 = value
}

// SetLine3 sets the line_3 property value. The line_3 property
func (m *SettingApplicationConfigurationGetResponse_contact_address_labels) SetLine3(value *string) {
	m.line_3 = value
}

type SettingApplicationConfigurationGetResponse_contact_address_labelsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetLine1() *string
	GetLine2() *string
	GetLine3() *string
	SetLine1(value *string)
	SetLine2(value *string)
	SetLine3(value *string)
}
