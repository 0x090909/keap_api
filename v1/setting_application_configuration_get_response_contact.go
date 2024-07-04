package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SettingApplicationConfigurationGetResponse_contact struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The address_labels property
	address_labels SettingApplicationConfigurationGetResponse_contact_address_labelsable
	// The contact_types property
	contact_types *string
	// The default_new_contact_form property
	default_new_contact_form *string
	// The disable_contact_edit_in_client_login property
	disable_contact_edit_in_client_login *bool
	// The fax_types property
	fax_types *string
	// The phone_types property
	phone_types *string
	// The suffix_types property
	suffix_types *string
	// The title_types property
	title_types *string
}

// NewSettingApplicationConfigurationGetResponse_contact instantiates a new SettingApplicationConfigurationGetResponse_contact and sets the default values.
func NewSettingApplicationConfigurationGetResponse_contact() *SettingApplicationConfigurationGetResponse_contact {
	m := &SettingApplicationConfigurationGetResponse_contact{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateSettingApplicationConfigurationGetResponse_contactFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSettingApplicationConfigurationGetResponse_contactFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSettingApplicationConfigurationGetResponse_contact(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SettingApplicationConfigurationGetResponse_contact) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAddressLabels gets the address_labels property value. The address_labels property
// returns a SettingApplicationConfigurationGetResponse_contact_address_labelsable when successful
func (m *SettingApplicationConfigurationGetResponse_contact) GetAddressLabels() SettingApplicationConfigurationGetResponse_contact_address_labelsable {
	return m.address_labels
}

// GetContactTypes gets the contact_types property value. The contact_types property
// returns a *string when successful
func (m *SettingApplicationConfigurationGetResponse_contact) GetContactTypes() *string {
	return m.contact_types
}

// GetDefaultNewContactForm gets the default_new_contact_form property value. The default_new_contact_form property
// returns a *string when successful
func (m *SettingApplicationConfigurationGetResponse_contact) GetDefaultNewContactForm() *string {
	return m.default_new_contact_form
}

// GetDisableContactEditInClientLogin gets the disable_contact_edit_in_client_login property value. The disable_contact_edit_in_client_login property
// returns a *bool when successful
func (m *SettingApplicationConfigurationGetResponse_contact) GetDisableContactEditInClientLogin() *bool {
	return m.disable_contact_edit_in_client_login
}

// GetFaxTypes gets the fax_types property value. The fax_types property
// returns a *string when successful
func (m *SettingApplicationConfigurationGetResponse_contact) GetFaxTypes() *string {
	return m.fax_types
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SettingApplicationConfigurationGetResponse_contact) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["address_labels"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSettingApplicationConfigurationGetResponse_contact_address_labelsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAddressLabels(val.(SettingApplicationConfigurationGetResponse_contact_address_labelsable))
		}
		return nil
	}
	res["contact_types"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetContactTypes(val)
		}
		return nil
	}
	res["default_new_contact_form"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDefaultNewContactForm(val)
		}
		return nil
	}
	res["disable_contact_edit_in_client_login"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDisableContactEditInClientLogin(val)
		}
		return nil
	}
	res["fax_types"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFaxTypes(val)
		}
		return nil
	}
	res["phone_types"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPhoneTypes(val)
		}
		return nil
	}
	res["suffix_types"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSuffixTypes(val)
		}
		return nil
	}
	res["title_types"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTitleTypes(val)
		}
		return nil
	}
	return res
}

// GetPhoneTypes gets the phone_types property value. The phone_types property
// returns a *string when successful
func (m *SettingApplicationConfigurationGetResponse_contact) GetPhoneTypes() *string {
	return m.phone_types
}

// GetSuffixTypes gets the suffix_types property value. The suffix_types property
// returns a *string when successful
func (m *SettingApplicationConfigurationGetResponse_contact) GetSuffixTypes() *string {
	return m.suffix_types
}

// GetTitleTypes gets the title_types property value. The title_types property
// returns a *string when successful
func (m *SettingApplicationConfigurationGetResponse_contact) GetTitleTypes() *string {
	return m.title_types
}

// Serialize serializes information the current object
func (m *SettingApplicationConfigurationGetResponse_contact) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("address_labels", m.GetAddressLabels())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("contact_types", m.GetContactTypes())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("default_new_contact_form", m.GetDefaultNewContactForm())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("disable_contact_edit_in_client_login", m.GetDisableContactEditInClientLogin())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("fax_types", m.GetFaxTypes())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("phone_types", m.GetPhoneTypes())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("suffix_types", m.GetSuffixTypes())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("title_types", m.GetTitleTypes())
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
func (m *SettingApplicationConfigurationGetResponse_contact) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAddressLabels sets the address_labels property value. The address_labels property
func (m *SettingApplicationConfigurationGetResponse_contact) SetAddressLabels(value SettingApplicationConfigurationGetResponse_contact_address_labelsable) {
	m.address_labels = value
}

// SetContactTypes sets the contact_types property value. The contact_types property
func (m *SettingApplicationConfigurationGetResponse_contact) SetContactTypes(value *string) {
	m.contact_types = value
}

// SetDefaultNewContactForm sets the default_new_contact_form property value. The default_new_contact_form property
func (m *SettingApplicationConfigurationGetResponse_contact) SetDefaultNewContactForm(value *string) {
	m.default_new_contact_form = value
}

// SetDisableContactEditInClientLogin sets the disable_contact_edit_in_client_login property value. The disable_contact_edit_in_client_login property
func (m *SettingApplicationConfigurationGetResponse_contact) SetDisableContactEditInClientLogin(value *bool) {
	m.disable_contact_edit_in_client_login = value
}

// SetFaxTypes sets the fax_types property value. The fax_types property
func (m *SettingApplicationConfigurationGetResponse_contact) SetFaxTypes(value *string) {
	m.fax_types = value
}

// SetPhoneTypes sets the phone_types property value. The phone_types property
func (m *SettingApplicationConfigurationGetResponse_contact) SetPhoneTypes(value *string) {
	m.phone_types = value
}

// SetSuffixTypes sets the suffix_types property value. The suffix_types property
func (m *SettingApplicationConfigurationGetResponse_contact) SetSuffixTypes(value *string) {
	m.suffix_types = value
}

// SetTitleTypes sets the title_types property value. The title_types property
func (m *SettingApplicationConfigurationGetResponse_contact) SetTitleTypes(value *string) {
	m.title_types = value
}

type SettingApplicationConfigurationGetResponse_contactable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAddressLabels() SettingApplicationConfigurationGetResponse_contact_address_labelsable
	GetContactTypes() *string
	GetDefaultNewContactForm() *string
	GetDisableContactEditInClientLogin() *bool
	GetFaxTypes() *string
	GetPhoneTypes() *string
	GetSuffixTypes() *string
	GetTitleTypes() *string
	SetAddressLabels(value SettingApplicationConfigurationGetResponse_contact_address_labelsable)
	SetContactTypes(value *string)
	SetDefaultNewContactForm(value *string)
	SetDisableContactEditInClientLogin(value *bool)
	SetFaxTypes(value *string)
	SetPhoneTypes(value *string)
	SetSuffixTypes(value *string)
	SetTitleTypes(value *string)
}
