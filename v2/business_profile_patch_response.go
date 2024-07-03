package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BusinessProfilePatchResponse profile information about the business that owns/uses this account
type BusinessProfilePatchResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The address property
	address BusinessProfilePatchResponse_addressable
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

// NewBusinessProfilePatchResponse instantiates a new BusinessProfilePatchResponse and sets the default values.
func NewBusinessProfilePatchResponse() *BusinessProfilePatchResponse {
	m := &BusinessProfilePatchResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateBusinessProfilePatchResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBusinessProfilePatchResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewBusinessProfilePatchResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BusinessProfilePatchResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAddress gets the address property value. The address property
// returns a BusinessProfilePatchResponse_addressable when successful
func (m *BusinessProfilePatchResponse) GetAddress() BusinessProfilePatchResponse_addressable {
	return m.address
}

// GetBusinessGoals gets the business_goals property value. The goals of this business, ie. Grow Business, Convert more leads
// returns a []string when successful
func (m *BusinessProfilePatchResponse) GetBusinessGoals() []string {
	return m.business_goals
}

// GetBusinessPrimaryColor gets the business_primary_color property value. The business_primary_color property
// returns a *string when successful
func (m *BusinessProfilePatchResponse) GetBusinessPrimaryColor() *string {
	return m.business_primary_color
}

// GetBusinessSecondaryColor gets the business_secondary_color property value. The business_secondary_color property
// returns a *string when successful
func (m *BusinessProfilePatchResponse) GetBusinessSecondaryColor() *string {
	return m.business_secondary_color
}

// GetBusinessType gets the business_type property value. The type of business
// returns a *string when successful
func (m *BusinessProfilePatchResponse) GetBusinessType() *string {
	return m.business_type
}

// GetCurrencyCode gets the currency_code property value. ISO 4217 Currency Code
// returns a *string when successful
func (m *BusinessProfilePatchResponse) GetCurrencyCode() *string {
	return m.currency_code
}

// GetEmail gets the email property value. The email property
// returns a *string when successful
func (m *BusinessProfilePatchResponse) GetEmail() *string {
	return m.email
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BusinessProfilePatchResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["address"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateBusinessProfilePatchResponse_addressFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAddress(val.(BusinessProfilePatchResponse_addressable))
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
func (m *BusinessProfilePatchResponse) GetLanguageTag() *string {
	return m.language_tag
}

// GetLogoUrl gets the logo_url property value. The logo_url property
// returns a *string when successful
func (m *BusinessProfilePatchResponse) GetLogoUrl() *string {
	return m.logo_url
}

// GetName gets the name property value. The name property
// returns a *string when successful
func (m *BusinessProfilePatchResponse) GetName() *string {
	return m.name
}

// GetPhone gets the phone property value. The phone property
// returns a *string when successful
func (m *BusinessProfilePatchResponse) GetPhone() *string {
	return m.phone
}

// GetTimeZone gets the time_zone property value. The time_zone property
// returns a *string when successful
func (m *BusinessProfilePatchResponse) GetTimeZone() *string {
	return m.time_zone
}

// GetWebsite gets the website property value. The website property
// returns a *string when successful
func (m *BusinessProfilePatchResponse) GetWebsite() *string {
	return m.website
}

// Serialize serializes information the current object
func (m *BusinessProfilePatchResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *BusinessProfilePatchResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAddress sets the address property value. The address property
func (m *BusinessProfilePatchResponse) SetAddress(value BusinessProfilePatchResponse_addressable) {
	m.address = value
}

// SetBusinessGoals sets the business_goals property value. The goals of this business, ie. Grow Business, Convert more leads
func (m *BusinessProfilePatchResponse) SetBusinessGoals(value []string) {
	m.business_goals = value
}

// SetBusinessPrimaryColor sets the business_primary_color property value. The business_primary_color property
func (m *BusinessProfilePatchResponse) SetBusinessPrimaryColor(value *string) {
	m.business_primary_color = value
}

// SetBusinessSecondaryColor sets the business_secondary_color property value. The business_secondary_color property
func (m *BusinessProfilePatchResponse) SetBusinessSecondaryColor(value *string) {
	m.business_secondary_color = value
}

// SetBusinessType sets the business_type property value. The type of business
func (m *BusinessProfilePatchResponse) SetBusinessType(value *string) {
	m.business_type = value
}

// SetCurrencyCode sets the currency_code property value. ISO 4217 Currency Code
func (m *BusinessProfilePatchResponse) SetCurrencyCode(value *string) {
	m.currency_code = value
}

// SetEmail sets the email property value. The email property
func (m *BusinessProfilePatchResponse) SetEmail(value *string) {
	m.email = value
}

// SetLanguageTag sets the language_tag property value. The language_tag property
func (m *BusinessProfilePatchResponse) SetLanguageTag(value *string) {
	m.language_tag = value
}

// SetLogoUrl sets the logo_url property value. The logo_url property
func (m *BusinessProfilePatchResponse) SetLogoUrl(value *string) {
	m.logo_url = value
}

// SetName sets the name property value. The name property
func (m *BusinessProfilePatchResponse) SetName(value *string) {
	m.name = value
}

// SetPhone sets the phone property value. The phone property
func (m *BusinessProfilePatchResponse) SetPhone(value *string) {
	m.phone = value
}

// SetTimeZone sets the time_zone property value. The time_zone property
func (m *BusinessProfilePatchResponse) SetTimeZone(value *string) {
	m.time_zone = value
}

// SetWebsite sets the website property value. The website property
func (m *BusinessProfilePatchResponse) SetWebsite(value *string) {
	m.website = value
}

type BusinessProfilePatchResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAddress() BusinessProfilePatchResponse_addressable
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
	SetAddress(value BusinessProfilePatchResponse_addressable)
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
