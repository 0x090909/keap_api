package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SettingApplicationConfigurationGetResponse_template struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The default_country_code property
	default_country_code *int32
	// The default_user_id property
	default_user_id *int64
}

// NewSettingApplicationConfigurationGetResponse_template instantiates a new SettingApplicationConfigurationGetResponse_template and sets the default values.
func NewSettingApplicationConfigurationGetResponse_template() *SettingApplicationConfigurationGetResponse_template {
	m := &SettingApplicationConfigurationGetResponse_template{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateSettingApplicationConfigurationGetResponse_templateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSettingApplicationConfigurationGetResponse_templateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSettingApplicationConfigurationGetResponse_template(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SettingApplicationConfigurationGetResponse_template) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetDefaultCountryCode gets the default_country_code property value. The default_country_code property
// returns a *int32 when successful
func (m *SettingApplicationConfigurationGetResponse_template) GetDefaultCountryCode() *int32 {
	return m.default_country_code
}

// GetDefaultUserId gets the default_user_id property value. The default_user_id property
// returns a *int64 when successful
func (m *SettingApplicationConfigurationGetResponse_template) GetDefaultUserId() *int64 {
	return m.default_user_id
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SettingApplicationConfigurationGetResponse_template) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["default_country_code"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDefaultCountryCode(val)
		}
		return nil
	}
	res["default_user_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDefaultUserId(val)
		}
		return nil
	}
	return res
}

// Serialize serializes information the current object
func (m *SettingApplicationConfigurationGetResponse_template) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt32Value("default_country_code", m.GetDefaultCountryCode())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("default_user_id", m.GetDefaultUserId())
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
func (m *SettingApplicationConfigurationGetResponse_template) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetDefaultCountryCode sets the default_country_code property value. The default_country_code property
func (m *SettingApplicationConfigurationGetResponse_template) SetDefaultCountryCode(value *int32) {
	m.default_country_code = value
}

// SetDefaultUserId sets the default_user_id property value. The default_user_id property
func (m *SettingApplicationConfigurationGetResponse_template) SetDefaultUserId(value *int64) {
	m.default_user_id = value
}

type SettingApplicationConfigurationGetResponse_templateable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetDefaultCountryCode() *int32
	GetDefaultUserId() *int64
	SetDefaultCountryCode(value *int32)
	SetDefaultUserId(value *int64)
}
