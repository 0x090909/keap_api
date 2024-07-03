package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SettingsApplicationsGetConfigurationGetResponse_application_company struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The caller_id_number property
	caller_id_number *string
	// The city property
	city *string
	// The country property
	country *string
	// The email property
	email *string
	// The name property
	name *string
	// The phone property
	phone *string
	// The phone_ext property
	phone_ext *string
	// The state property
	state *string
	// The street_address_1 property
	street_address_1 *string
	// The street_address_2 property
	street_address_2 *string
	// The web_logo_url property
	web_logo_url *string
	// The website property
	website *string
	// The zip property
	zip *string
}

// NewSettingsApplicationsGetConfigurationGetResponse_application_company instantiates a new SettingsApplicationsGetConfigurationGetResponse_application_company and sets the default values.
func NewSettingsApplicationsGetConfigurationGetResponse_application_company() *SettingsApplicationsGetConfigurationGetResponse_application_company {
	m := &SettingsApplicationsGetConfigurationGetResponse_application_company{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateSettingsApplicationsGetConfigurationGetResponse_application_companyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSettingsApplicationsGetConfigurationGetResponse_application_companyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSettingsApplicationsGetConfigurationGetResponse_application_company(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_application_company) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCallerIdNumber gets the caller_id_number property value. The caller_id_number property
// returns a *string when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_application_company) GetCallerIdNumber() *string {
	return m.caller_id_number
}

// GetCity gets the city property value. The city property
// returns a *string when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_application_company) GetCity() *string {
	return m.city
}

// GetCountry gets the country property value. The country property
// returns a *string when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_application_company) GetCountry() *string {
	return m.country
}

// GetEmail gets the email property value. The email property
// returns a *string when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_application_company) GetEmail() *string {
	return m.email
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_application_company) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["caller_id_number"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCallerIdNumber(val)
		}
		return nil
	}
	res["city"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCity(val)
		}
		return nil
	}
	res["country"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCountry(val)
		}
		return nil
	}
	res["email"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetEmail(val)
		}
		return nil
	}
	res["name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetName(val)
		}
		return nil
	}
	res["phone"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPhone(val)
		}
		return nil
	}
	res["phone_ext"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPhoneExt(val)
		}
		return nil
	}
	res["state"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetState(val)
		}
		return nil
	}
	res["street_address_1"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetStreetAddress1(val)
		}
		return nil
	}
	res["street_address_2"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetStreetAddress2(val)
		}
		return nil
	}
	res["web_logo_url"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetWebLogoUrl(val)
		}
		return nil
	}
	res["website"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetWebsite(val)
		}
		return nil
	}
	res["zip"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetZip(val)
		}
		return nil
	}
	return res
}

// GetName gets the name property value. The name property
// returns a *string when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_application_company) GetName() *string {
	return m.name
}

// GetPhone gets the phone property value. The phone property
// returns a *string when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_application_company) GetPhone() *string {
	return m.phone
}

// GetPhoneExt gets the phone_ext property value. The phone_ext property
// returns a *string when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_application_company) GetPhoneExt() *string {
	return m.phone_ext
}

// GetState gets the state property value. The state property
// returns a *string when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_application_company) GetState() *string {
	return m.state
}

// GetStreetAddress1 gets the street_address_1 property value. The street_address_1 property
// returns a *string when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_application_company) GetStreetAddress1() *string {
	return m.street_address_1
}

// GetStreetAddress2 gets the street_address_2 property value. The street_address_2 property
// returns a *string when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_application_company) GetStreetAddress2() *string {
	return m.street_address_2
}

// GetWebLogoUrl gets the web_logo_url property value. The web_logo_url property
// returns a *string when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_application_company) GetWebLogoUrl() *string {
	return m.web_logo_url
}

// GetWebsite gets the website property value. The website property
// returns a *string when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_application_company) GetWebsite() *string {
	return m.website
}

// GetZip gets the zip property value. The zip property
// returns a *string when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_application_company) GetZip() *string {
	return m.zip
}

// Serialize serializes information the current object
func (m *SettingsApplicationsGetConfigurationGetResponse_application_company) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("caller_id_number", m.GetCallerIdNumber())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("city", m.GetCity())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("country", m.GetCountry())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("email", m.GetEmail())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("name", m.GetName())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("phone", m.GetPhone())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("phone_ext", m.GetPhoneExt())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("state", m.GetState())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("street_address_1", m.GetStreetAddress1())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("street_address_2", m.GetStreetAddress2())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("website", m.GetWebsite())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("web_logo_url", m.GetWebLogoUrl())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("zip", m.GetZip())
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
func (m *SettingsApplicationsGetConfigurationGetResponse_application_company) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCallerIdNumber sets the caller_id_number property value. The caller_id_number property
func (m *SettingsApplicationsGetConfigurationGetResponse_application_company) SetCallerIdNumber(value *string) {
	m.caller_id_number = value
}

// SetCity sets the city property value. The city property
func (m *SettingsApplicationsGetConfigurationGetResponse_application_company) SetCity(value *string) {
	m.city = value
}

// SetCountry sets the country property value. The country property
func (m *SettingsApplicationsGetConfigurationGetResponse_application_company) SetCountry(value *string) {
	m.country = value
}

// SetEmail sets the email property value. The email property
func (m *SettingsApplicationsGetConfigurationGetResponse_application_company) SetEmail(value *string) {
	m.email = value
}

// SetName sets the name property value. The name property
func (m *SettingsApplicationsGetConfigurationGetResponse_application_company) SetName(value *string) {
	m.name = value
}

// SetPhone sets the phone property value. The phone property
func (m *SettingsApplicationsGetConfigurationGetResponse_application_company) SetPhone(value *string) {
	m.phone = value
}

// SetPhoneExt sets the phone_ext property value. The phone_ext property
func (m *SettingsApplicationsGetConfigurationGetResponse_application_company) SetPhoneExt(value *string) {
	m.phone_ext = value
}

// SetState sets the state property value. The state property
func (m *SettingsApplicationsGetConfigurationGetResponse_application_company) SetState(value *string) {
	m.state = value
}

// SetStreetAddress1 sets the street_address_1 property value. The street_address_1 property
func (m *SettingsApplicationsGetConfigurationGetResponse_application_company) SetStreetAddress1(value *string) {
	m.street_address_1 = value
}

// SetStreetAddress2 sets the street_address_2 property value. The street_address_2 property
func (m *SettingsApplicationsGetConfigurationGetResponse_application_company) SetStreetAddress2(value *string) {
	m.street_address_2 = value
}

// SetWebLogoUrl sets the web_logo_url property value. The web_logo_url property
func (m *SettingsApplicationsGetConfigurationGetResponse_application_company) SetWebLogoUrl(value *string) {
	m.web_logo_url = value
}

// SetWebsite sets the website property value. The website property
func (m *SettingsApplicationsGetConfigurationGetResponse_application_company) SetWebsite(value *string) {
	m.website = value
}

// SetZip sets the zip property value. The zip property
func (m *SettingsApplicationsGetConfigurationGetResponse_application_company) SetZip(value *string) {
	m.zip = value
}

type SettingsApplicationsGetConfigurationGetResponse_application_companyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCallerIdNumber() *string
	GetCity() *string
	GetCountry() *string
	GetEmail() *string
	GetName() *string
	GetPhone() *string
	GetPhoneExt() *string
	GetState() *string
	GetStreetAddress1() *string
	GetStreetAddress2() *string
	GetWebLogoUrl() *string
	GetWebsite() *string
	GetZip() *string
	SetCallerIdNumber(value *string)
	SetCity(value *string)
	SetCountry(value *string)
	SetEmail(value *string)
	SetName(value *string)
	SetPhone(value *string)
	SetPhoneExt(value *string)
	SetState(value *string)
	SetStreetAddress1(value *string)
	SetStreetAddress2(value *string)
	SetWebLogoUrl(value *string)
	SetWebsite(value *string)
	SetZip(value *string)
}
