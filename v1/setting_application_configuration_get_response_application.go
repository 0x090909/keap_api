package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SettingApplicationConfigurationGetResponse_application struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The company property
	company SettingApplicationConfigurationGetResponse_application_companyable
	// The default_view_locale property
	default_view_locale *string
	// The features_enabled property
	features_enabled SettingApplicationConfigurationGetResponse_application_features_enabledable
	// The time_zone property
	time_zone *string
}

// NewSettingApplicationConfigurationGetResponse_application instantiates a new SettingApplicationConfigurationGetResponse_application and sets the default values.
func NewSettingApplicationConfigurationGetResponse_application() *SettingApplicationConfigurationGetResponse_application {
	m := &SettingApplicationConfigurationGetResponse_application{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateSettingApplicationConfigurationGetResponse_applicationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSettingApplicationConfigurationGetResponse_applicationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSettingApplicationConfigurationGetResponse_application(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SettingApplicationConfigurationGetResponse_application) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCompany gets the company property value. The company property
// returns a SettingApplicationConfigurationGetResponse_application_companyable when successful
func (m *SettingApplicationConfigurationGetResponse_application) GetCompany() SettingApplicationConfigurationGetResponse_application_companyable {
	return m.company
}

// GetDefaultViewLocale gets the default_view_locale property value. The default_view_locale property
// returns a *string when successful
func (m *SettingApplicationConfigurationGetResponse_application) GetDefaultViewLocale() *string {
	return m.default_view_locale
}

// GetFeaturesEnabled gets the features_enabled property value. The features_enabled property
// returns a SettingApplicationConfigurationGetResponse_application_features_enabledable when successful
func (m *SettingApplicationConfigurationGetResponse_application) GetFeaturesEnabled() SettingApplicationConfigurationGetResponse_application_features_enabledable {
	return m.features_enabled
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SettingApplicationConfigurationGetResponse_application) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["company"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSettingApplicationConfigurationGetResponse_application_companyFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCompany(val.(SettingApplicationConfigurationGetResponse_application_companyable))
		}
		return nil
	}
	res["default_view_locale"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDefaultViewLocale(val)
		}
		return nil
	}
	res["features_enabled"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSettingApplicationConfigurationGetResponse_application_features_enabledFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFeaturesEnabled(val.(SettingApplicationConfigurationGetResponse_application_features_enabledable))
		}
		return nil
	}
	res["time_zone"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTimeZone(val)
		}
		return nil
	}
	return res
}

// GetTimeZone gets the time_zone property value. The time_zone property
// returns a *string when successful
func (m *SettingApplicationConfigurationGetResponse_application) GetTimeZone() *string {
	return m.time_zone
}

// Serialize serializes information the current object
func (m *SettingApplicationConfigurationGetResponse_application) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("company", m.GetCompany())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("default_view_locale", m.GetDefaultViewLocale())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("features_enabled", m.GetFeaturesEnabled())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("time_zone", m.GetTimeZone())
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
func (m *SettingApplicationConfigurationGetResponse_application) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCompany sets the company property value. The company property
func (m *SettingApplicationConfigurationGetResponse_application) SetCompany(value SettingApplicationConfigurationGetResponse_application_companyable) {
	m.company = value
}

// SetDefaultViewLocale sets the default_view_locale property value. The default_view_locale property
func (m *SettingApplicationConfigurationGetResponse_application) SetDefaultViewLocale(value *string) {
	m.default_view_locale = value
}

// SetFeaturesEnabled sets the features_enabled property value. The features_enabled property
func (m *SettingApplicationConfigurationGetResponse_application) SetFeaturesEnabled(value SettingApplicationConfigurationGetResponse_application_features_enabledable) {
	m.features_enabled = value
}

// SetTimeZone sets the time_zone property value. The time_zone property
func (m *SettingApplicationConfigurationGetResponse_application) SetTimeZone(value *string) {
	m.time_zone = value
}

type SettingApplicationConfigurationGetResponse_applicationable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCompany() SettingApplicationConfigurationGetResponse_application_companyable
	GetDefaultViewLocale() *string
	GetFeaturesEnabled() SettingApplicationConfigurationGetResponse_application_features_enabledable
	GetTimeZone() *string
	SetCompany(value SettingApplicationConfigurationGetResponse_application_companyable)
	SetDefaultViewLocale(value *string)
	SetFeaturesEnabled(value SettingApplicationConfigurationGetResponse_application_features_enabledable)
	SetTimeZone(value *string)
}
