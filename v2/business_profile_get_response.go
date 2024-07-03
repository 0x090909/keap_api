package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BusinessProfileGetResponse profile information about the business that owns/uses this account
type BusinessProfileGetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The address property
	address BusinessProfileGetResponse_addressable
	// The goals of this business, ie. Grow Business, Convert more leads
	business_goals []string
	// The business_primary_color property
	business_primary_color *string
	// The business_secondary_color property
	business_secondary_color *string
	// The type of business
	business_type *string
	// ISO 4217 Currency Code
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
	// The time_zone property
	time_zone *string
	// The website property
	website *string
}

// NewBusinessProfileGetResponse instantiates a new BusinessProfileGetResponse and sets the default values.
func NewBusinessProfileGetResponse() *BusinessProfileGetResponse {
	m := &BusinessProfileGetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateBusinessProfileGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBusinessProfileGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewBusinessProfileGetResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BusinessProfileGetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAddress gets the address property value. The address property
// returns a BusinessProfileGetResponse_addressable when successful
func (m *BusinessProfileGetResponse) GetAddress() BusinessProfileGetResponse_addressable {
	return m.address
}

// GetBusinessGoals gets the business_goals property value. The goals of this business, ie. Grow Business, Convert more leads
// returns a []string when successful
func (m *BusinessProfileGetResponse) GetBusinessGoals() []string {
	return m.business_goals
}

// GetBusinessPrimaryColor gets the business_primary_color property value. The business_primary_color property
// returns a *string when successful
func (m *BusinessProfileGetResponse) GetBusinessPrimaryColor() *string {
	return m.business_primary_color
}

// GetBusinessSecondaryColor gets the business_secondary_color property value. The business_secondary_color property
// returns a *string when successful
func (m *BusinessProfileGetResponse) GetBusinessSecondaryColor() *string {
	return m.business_secondary_color
}

// GetBusinessType gets the business_type property value. The type of business
// returns a *string when successful
func (m *BusinessProfileGetResponse) GetBusinessType() *string {
	return m.business_type
}

// GetCurrencyCode gets the currency_code property value. ISO 4217 Currency Code
// returns a *string when successful
func (m *BusinessProfileGetResponse) GetCurrencyCode() *string {
	return m.currency_code
}

// GetEmail gets the email property value. The email property
// returns a *string when successful
func (m *BusinessProfileGetResponse) GetEmail() *string {
	return m.email
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BusinessProfileGetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["address"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateBusinessProfileGetResponse_addressFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAddress(val.(BusinessProfileGetResponse_addressable))
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
func (m *BusinessProfileGetResponse) GetLanguageTag() *string {
	return m.language_tag
}

// GetLogoUrl gets the logo_url property value. The logo_url property
// returns a *string when successful
func (m *BusinessProfileGetResponse) GetLogoUrl() *string {
	return m.logo_url
}

// GetName gets the name property value. The name property
// returns a *string when successful
func (m *BusinessProfileGetResponse) GetName() *string {
	return m.name
}

// GetPhone gets the phone property value. The phone property
// returns a *string when successful
func (m *BusinessProfileGetResponse) GetPhone() *string {
	return m.phone
}

// GetTimeZone gets the time_zone property value. The time_zone property
// returns a *string when successful
func (m *BusinessProfileGetResponse) GetTimeZone() *string {
	return m.time_zone
}

// GetWebsite gets the website property value. The website property
// returns a *string when successful
func (m *BusinessProfileGetResponse) GetWebsite() *string {
	return m.website
}

// Serialize serializes information the current object
func (m *BusinessProfileGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *BusinessProfileGetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAddress sets the address property value. The address property
func (m *BusinessProfileGetResponse) SetAddress(value BusinessProfileGetResponse_addressable) {
	m.address = value
}

// SetBusinessGoals sets the business_goals property value. The goals of this business, ie. Grow Business, Convert more leads
func (m *BusinessProfileGetResponse) SetBusinessGoals(value []string) {
	m.business_goals = value
}

// SetBusinessPrimaryColor sets the business_primary_color property value. The business_primary_color property
func (m *BusinessProfileGetResponse) SetBusinessPrimaryColor(value *string) {
	m.business_primary_color = value
}

// SetBusinessSecondaryColor sets the business_secondary_color property value. The business_secondary_color property
func (m *BusinessProfileGetResponse) SetBusinessSecondaryColor(value *string) {
	m.business_secondary_color = value
}

// SetBusinessType sets the business_type property value. The type of business
func (m *BusinessProfileGetResponse) SetBusinessType(value *string) {
	m.business_type = value
}

// SetCurrencyCode sets the currency_code property value. ISO 4217 Currency Code
func (m *BusinessProfileGetResponse) SetCurrencyCode(value *string) {
	m.currency_code = value
}

// SetEmail sets the email property value. The email property
func (m *BusinessProfileGetResponse) SetEmail(value *string) {
	m.email = value
}

// SetLanguageTag sets the language_tag property value. The language_tag property
func (m *BusinessProfileGetResponse) SetLanguageTag(value *string) {
	m.language_tag = value
}

// SetLogoUrl sets the logo_url property value. The logo_url property
func (m *BusinessProfileGetResponse) SetLogoUrl(value *string) {
	m.logo_url = value
}

// SetName sets the name property value. The name property
func (m *BusinessProfileGetResponse) SetName(value *string) {
	m.name = value
}

// SetPhone sets the phone property value. The phone property
func (m *BusinessProfileGetResponse) SetPhone(value *string) {
	m.phone = value
}

// SetTimeZone sets the time_zone property value. The time_zone property
func (m *BusinessProfileGetResponse) SetTimeZone(value *string) {
	m.time_zone = value
}

// SetWebsite sets the website property value. The website property
func (m *BusinessProfileGetResponse) SetWebsite(value *string) {
	m.website = value
}

type BusinessProfileGetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAddress() BusinessProfileGetResponse_addressable
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
	GetTimeZone() *string
	GetWebsite() *string
	SetAddress(value BusinessProfileGetResponse_addressable)
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
	SetTimeZone(value *string)
	SetWebsite(value *string)
}
