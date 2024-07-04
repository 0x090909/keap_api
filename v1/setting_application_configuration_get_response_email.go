package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SettingApplicationConfigurationGetResponse_email struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The append_contact_key_to_links property
	append_contact_key_to_links *bool
	// The default_opt_in_link property
	default_opt_in_link *string
	// The default_opt_out_link property
	default_opt_out_link *string
	// The hide_emails_to_and_from_domains property
	hide_emails_to_and_from_domains *string
	// The whitelisted_domains property
	whitelisted_domains *string
}

// NewSettingApplicationConfigurationGetResponse_email instantiates a new SettingApplicationConfigurationGetResponse_email and sets the default values.
func NewSettingApplicationConfigurationGetResponse_email() *SettingApplicationConfigurationGetResponse_email {
	m := &SettingApplicationConfigurationGetResponse_email{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateSettingApplicationConfigurationGetResponse_emailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSettingApplicationConfigurationGetResponse_emailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSettingApplicationConfigurationGetResponse_email(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SettingApplicationConfigurationGetResponse_email) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAppendContactKeyToLinks gets the append_contact_key_to_links property value. The append_contact_key_to_links property
// returns a *bool when successful
func (m *SettingApplicationConfigurationGetResponse_email) GetAppendContactKeyToLinks() *bool {
	return m.append_contact_key_to_links
}

// GetDefaultOptInLink gets the default_opt_in_link property value. The default_opt_in_link property
// returns a *string when successful
func (m *SettingApplicationConfigurationGetResponse_email) GetDefaultOptInLink() *string {
	return m.default_opt_in_link
}

// GetDefaultOptOutLink gets the default_opt_out_link property value. The default_opt_out_link property
// returns a *string when successful
func (m *SettingApplicationConfigurationGetResponse_email) GetDefaultOptOutLink() *string {
	return m.default_opt_out_link
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SettingApplicationConfigurationGetResponse_email) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["append_contact_key_to_links"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAppendContactKeyToLinks(val)
		}
		return nil
	}
	res["default_opt_in_link"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDefaultOptInLink(val)
		}
		return nil
	}
	res["default_opt_out_link"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDefaultOptOutLink(val)
		}
		return nil
	}
	res["hide_emails_to_and_from_domains"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetHideEmailsToAndFromDomains(val)
		}
		return nil
	}
	res["whitelisted_domains"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetWhitelistedDomains(val)
		}
		return nil
	}
	return res
}

// GetHideEmailsToAndFromDomains gets the hide_emails_to_and_from_domains property value. The hide_emails_to_and_from_domains property
// returns a *string when successful
func (m *SettingApplicationConfigurationGetResponse_email) GetHideEmailsToAndFromDomains() *string {
	return m.hide_emails_to_and_from_domains
}

// GetWhitelistedDomains gets the whitelisted_domains property value. The whitelisted_domains property
// returns a *string when successful
func (m *SettingApplicationConfigurationGetResponse_email) GetWhitelistedDomains() *string {
	return m.whitelisted_domains
}

// Serialize serializes information the current object
func (m *SettingApplicationConfigurationGetResponse_email) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteBoolValue("append_contact_key_to_links", m.GetAppendContactKeyToLinks())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("default_opt_in_link", m.GetDefaultOptInLink())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("default_opt_out_link", m.GetDefaultOptOutLink())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("hide_emails_to_and_from_domains", m.GetHideEmailsToAndFromDomains())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("whitelisted_domains", m.GetWhitelistedDomains())
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
func (m *SettingApplicationConfigurationGetResponse_email) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAppendContactKeyToLinks sets the append_contact_key_to_links property value. The append_contact_key_to_links property
func (m *SettingApplicationConfigurationGetResponse_email) SetAppendContactKeyToLinks(value *bool) {
	m.append_contact_key_to_links = value
}

// SetDefaultOptInLink sets the default_opt_in_link property value. The default_opt_in_link property
func (m *SettingApplicationConfigurationGetResponse_email) SetDefaultOptInLink(value *string) {
	m.default_opt_in_link = value
}

// SetDefaultOptOutLink sets the default_opt_out_link property value. The default_opt_out_link property
func (m *SettingApplicationConfigurationGetResponse_email) SetDefaultOptOutLink(value *string) {
	m.default_opt_out_link = value
}

// SetHideEmailsToAndFromDomains sets the hide_emails_to_and_from_domains property value. The hide_emails_to_and_from_domains property
func (m *SettingApplicationConfigurationGetResponse_email) SetHideEmailsToAndFromDomains(value *string) {
	m.hide_emails_to_and_from_domains = value
}

// SetWhitelistedDomains sets the whitelisted_domains property value. The whitelisted_domains property
func (m *SettingApplicationConfigurationGetResponse_email) SetWhitelistedDomains(value *string) {
	m.whitelisted_domains = value
}

type SettingApplicationConfigurationGetResponse_emailable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAppendContactKeyToLinks() *bool
	GetDefaultOptInLink() *string
	GetDefaultOptOutLink() *string
	GetHideEmailsToAndFromDomains() *string
	GetWhitelistedDomains() *string
	SetAppendContactKeyToLinks(value *bool)
	SetDefaultOptInLink(value *string)
	SetDefaultOptOutLink(value *string)
	SetHideEmailsToAndFromDomains(value *string)
	SetWhitelistedDomains(value *string)
}
