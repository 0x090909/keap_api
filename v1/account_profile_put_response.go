package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccountProfilePutResponse information about the company that owns/uses this account
type AccountProfilePutResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The address property
	address AccountProfilePutResponse_addressable
	// The goals of this business, ie. Grow Business, Convert more leads
	business_goals []string
	// The business_primary_color property
	business_primary_color *string
	// The business_secondary_color property
	business_secondary_color *string
	// The type of business
	business_type *string
	// The currency_code property
	currency_code *string
	// The email property
	email *string
	// The language_tag property
	language_tag *string
	// The logo_url property
	logo_url *string
	// The name property
	name *string
	// The phone property
	phone *string
	// The phone_ext property
	phone_ext *string
	// The time_zone property
	time_zone *string
	// The website property
	website *string
}

// NewAccountProfilePutResponse instantiates a new AccountProfilePutResponse and sets the default values.
func NewAccountProfilePutResponse() *AccountProfilePutResponse {
	m := &AccountProfilePutResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateAccountProfilePutResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAccountProfilePutResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAccountProfilePutResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AccountProfilePutResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAddress gets the address property value. The address property
// returns a AccountProfilePutResponse_addressable when successful
func (m *AccountProfilePutResponse) GetAddress() AccountProfilePutResponse_addressable {
	return m.address
}

// GetBusinessGoals gets the business_goals property value. The goals of this business, ie. Grow Business, Convert more leads
// returns a []string when successful
func (m *AccountProfilePutResponse) GetBusinessGoals() []string {
	return m.business_goals
}

// GetBusinessPrimaryColor gets the business_primary_color property value. The business_primary_color property
// returns a *string when successful
func (m *AccountProfilePutResponse) GetBusinessPrimaryColor() *string {
	return m.business_primary_color
}

// GetBusinessSecondaryColor gets the business_secondary_color property value. The business_secondary_color property
// returns a *string when successful
func (m *AccountProfilePutResponse) GetBusinessSecondaryColor() *string {
	return m.business_secondary_color
}

// GetBusinessType gets the business_type property value. The type of business
// returns a *string when successful
func (m *AccountProfilePutResponse) GetBusinessType() *string {
	return m.business_type
}

// GetCurrencyCode gets the currency_code property value. The currency_code property
// returns a *string when successful
func (m *AccountProfilePutResponse) GetCurrencyCode() *string {
	return m.currency_code
}

// GetEmail gets the email property value. The email property
// returns a *string when successful
func (m *AccountProfilePutResponse) GetEmail() *string {
	return m.email
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AccountProfilePutResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["address"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateAccountProfilePutResponse_addressFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAddress(val.(AccountProfilePutResponse_addressable))
		}
		return nil
	}
	res["business_goals"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfPrimitiveValues("string")
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]string, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = *(v.(*string))
				}
			}
			m.SetBusinessGoals(res)
		}
		return nil
	}
	res["business_primary_color"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetBusinessPrimaryColor(val)
		}
		return nil
	}
	res["business_secondary_color"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetBusinessSecondaryColor(val)
		}
		return nil
	}
	res["business_type"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetBusinessType(val)
		}
		return nil
	}
	res["currency_code"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCurrencyCode(val)
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
	res["language_tag"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLanguageTag(val)
		}
		return nil
	}
	res["logo_url"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLogoUrl(val)
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
	return res
}

// GetLanguageTag gets the language_tag property value. The language_tag property
// returns a *string when successful
func (m *AccountProfilePutResponse) GetLanguageTag() *string {
	return m.language_tag
}

// GetLogoUrl gets the logo_url property value. The logo_url property
// returns a *string when successful
func (m *AccountProfilePutResponse) GetLogoUrl() *string {
	return m.logo_url
}

// GetName gets the name property value. The name property
// returns a *string when successful
func (m *AccountProfilePutResponse) GetName() *string {
	return m.name
}

// GetPhone gets the phone property value. The phone property
// returns a *string when successful
func (m *AccountProfilePutResponse) GetPhone() *string {
	return m.phone
}

// GetPhoneExt gets the phone_ext property value. The phone_ext property
// returns a *string when successful
func (m *AccountProfilePutResponse) GetPhoneExt() *string {
	return m.phone_ext
}

// GetTimeZone gets the time_zone property value. The time_zone property
// returns a *string when successful
func (m *AccountProfilePutResponse) GetTimeZone() *string {
	return m.time_zone
}

// GetWebsite gets the website property value. The website property
// returns a *string when successful
func (m *AccountProfilePutResponse) GetWebsite() *string {
	return m.website
}

// Serialize serializes information the current object
func (m *AccountProfilePutResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("address", m.GetAddress())
		if err != nil {
			return err
		}
	}
	if m.GetBusinessGoals() != nil {
		err := writer.WriteCollectionOfStringValues("business_goals", m.GetBusinessGoals())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("business_primary_color", m.GetBusinessPrimaryColor())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("business_secondary_color", m.GetBusinessSecondaryColor())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("business_type", m.GetBusinessType())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("currency_code", m.GetCurrencyCode())
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
		err := writer.WriteStringValue("language_tag", m.GetLanguageTag())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("logo_url", m.GetLogoUrl())
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
		err := writer.WriteStringValue("time_zone", m.GetTimeZone())
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
		err := writer.WriteAdditionalData(m.GetAdditionalData())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccountProfilePutResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAddress sets the address property value. The address property
func (m *AccountProfilePutResponse) SetAddress(value AccountProfilePutResponse_addressable) {
	m.address = value
}

// SetBusinessGoals sets the business_goals property value. The goals of this business, ie. Grow Business, Convert more leads
func (m *AccountProfilePutResponse) SetBusinessGoals(value []string) {
	m.business_goals = value
}

// SetBusinessPrimaryColor sets the business_primary_color property value. The business_primary_color property
func (m *AccountProfilePutResponse) SetBusinessPrimaryColor(value *string) {
	m.business_primary_color = value
}

// SetBusinessSecondaryColor sets the business_secondary_color property value. The business_secondary_color property
func (m *AccountProfilePutResponse) SetBusinessSecondaryColor(value *string) {
	m.business_secondary_color = value
}

// SetBusinessType sets the business_type property value. The type of business
func (m *AccountProfilePutResponse) SetBusinessType(value *string) {
	m.business_type = value
}

// SetCurrencyCode sets the currency_code property value. The currency_code property
func (m *AccountProfilePutResponse) SetCurrencyCode(value *string) {
	m.currency_code = value
}

// SetEmail sets the email property value. The email property
func (m *AccountProfilePutResponse) SetEmail(value *string) {
	m.email = value
}

// SetLanguageTag sets the language_tag property value. The language_tag property
func (m *AccountProfilePutResponse) SetLanguageTag(value *string) {
	m.language_tag = value
}

// SetLogoUrl sets the logo_url property value. The logo_url property
func (m *AccountProfilePutResponse) SetLogoUrl(value *string) {
	m.logo_url = value
}

// SetName sets the name property value. The name property
func (m *AccountProfilePutResponse) SetName(value *string) {
	m.name = value
}

// SetPhone sets the phone property value. The phone property
func (m *AccountProfilePutResponse) SetPhone(value *string) {
	m.phone = value
}

// SetPhoneExt sets the phone_ext property value. The phone_ext property
func (m *AccountProfilePutResponse) SetPhoneExt(value *string) {
	m.phone_ext = value
}

// SetTimeZone sets the time_zone property value. The time_zone property
func (m *AccountProfilePutResponse) SetTimeZone(value *string) {
	m.time_zone = value
}

// SetWebsite sets the website property value. The website property
func (m *AccountProfilePutResponse) SetWebsite(value *string) {
	m.website = value
}

type AccountProfilePutResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAddress() AccountProfilePutResponse_addressable
	GetBusinessGoals() []string
	GetBusinessPrimaryColor() *string
	GetBusinessSecondaryColor() *string
	GetBusinessType() *string
	GetCurrencyCode() *string
	GetEmail() *string
	GetLanguageTag() *string
	GetLogoUrl() *string
	GetName() *string
	GetPhone() *string
	GetPhoneExt() *string
	GetTimeZone() *string
	GetWebsite() *string
	SetAddress(value AccountProfilePutResponse_addressable)
	SetBusinessGoals(value []string)
	SetBusinessPrimaryColor(value *string)
	SetBusinessSecondaryColor(value *string)
	SetBusinessType(value *string)
	SetCurrencyCode(value *string)
	SetEmail(value *string)
	SetLanguageTag(value *string)
	SetLogoUrl(value *string)
	SetName(value *string)
	SetPhone(value *string)
	SetPhoneExt(value *string)
	SetTimeZone(value *string)
	SetWebsite(value *string)
}
