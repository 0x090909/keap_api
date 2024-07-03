package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BusinessProfilePatchRequestBody update profile information about the business that owns/uses this account
type BusinessProfilePatchRequestBody struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The address property
	address BusinessProfilePatchRequestBody_addressable
	// The goals of this business, ie. Grow Business, Convert more leads
	business_goals []string
	// The business_primary_color property
	business_primary_color *string
	// The business_secondary_color property
	business_secondary_color *string
	// ISO 4217 Currency Code
	currency_code *string
	// The email property
	email *string
	// The name property
	name *string
	// The phone property
	phone *string
	// The website property
	website *string
}

// NewBusinessProfilePatchRequestBody instantiates a new BusinessProfilePatchRequestBody and sets the default values.
func NewBusinessProfilePatchRequestBody() *BusinessProfilePatchRequestBody {
	m := &BusinessProfilePatchRequestBody{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateBusinessProfilePatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBusinessProfilePatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewBusinessProfilePatchRequestBody(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BusinessProfilePatchRequestBody) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAddress gets the address property value. The address property
// returns a BusinessProfilePatchRequestBody_addressable when successful
func (m *BusinessProfilePatchRequestBody) GetAddress() BusinessProfilePatchRequestBody_addressable {
	return m.address
}

// GetBusinessGoals gets the business_goals property value. The goals of this business, ie. Grow Business, Convert more leads
// returns a []string when successful
func (m *BusinessProfilePatchRequestBody) GetBusinessGoals() []string {
	return m.business_goals
}

// GetBusinessPrimaryColor gets the business_primary_color property value. The business_primary_color property
// returns a *string when successful
func (m *BusinessProfilePatchRequestBody) GetBusinessPrimaryColor() *string {
	return m.business_primary_color
}

// GetBusinessSecondaryColor gets the business_secondary_color property value. The business_secondary_color property
// returns a *string when successful
func (m *BusinessProfilePatchRequestBody) GetBusinessSecondaryColor() *string {
	return m.business_secondary_color
}

// GetCurrencyCode gets the currency_code property value. ISO 4217 Currency Code
// returns a *string when successful
func (m *BusinessProfilePatchRequestBody) GetCurrencyCode() *string {
	return m.currency_code
}

// GetEmail gets the email property value. The email property
// returns a *string when successful
func (m *BusinessProfilePatchRequestBody) GetEmail() *string {
	return m.email
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BusinessProfilePatchRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["address"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateBusinessProfilePatchRequestBody_addressFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAddress(val.(BusinessProfilePatchRequestBody_addressable))
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

// GetName gets the name property value. The name property
// returns a *string when successful
func (m *BusinessProfilePatchRequestBody) GetName() *string {
	return m.name
}

// GetPhone gets the phone property value. The phone property
// returns a *string when successful
func (m *BusinessProfilePatchRequestBody) GetPhone() *string {
	return m.phone
}

// GetWebsite gets the website property value. The website property
// returns a *string when successful
func (m *BusinessProfilePatchRequestBody) GetWebsite() *string {
	return m.website
}

// Serialize serializes information the current object
func (m *BusinessProfilePatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *BusinessProfilePatchRequestBody) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAddress sets the address property value. The address property
func (m *BusinessProfilePatchRequestBody) SetAddress(value BusinessProfilePatchRequestBody_addressable) {
	m.address = value
}

// SetBusinessGoals sets the business_goals property value. The goals of this business, ie. Grow Business, Convert more leads
func (m *BusinessProfilePatchRequestBody) SetBusinessGoals(value []string) {
	m.business_goals = value
}

// SetBusinessPrimaryColor sets the business_primary_color property value. The business_primary_color property
func (m *BusinessProfilePatchRequestBody) SetBusinessPrimaryColor(value *string) {
	m.business_primary_color = value
}

// SetBusinessSecondaryColor sets the business_secondary_color property value. The business_secondary_color property
func (m *BusinessProfilePatchRequestBody) SetBusinessSecondaryColor(value *string) {
	m.business_secondary_color = value
}

// SetCurrencyCode sets the currency_code property value. ISO 4217 Currency Code
func (m *BusinessProfilePatchRequestBody) SetCurrencyCode(value *string) {
	m.currency_code = value
}

// SetEmail sets the email property value. The email property
func (m *BusinessProfilePatchRequestBody) SetEmail(value *string) {
	m.email = value
}

// SetName sets the name property value. The name property
func (m *BusinessProfilePatchRequestBody) SetName(value *string) {
	m.name = value
}

// SetPhone sets the phone property value. The phone property
func (m *BusinessProfilePatchRequestBody) SetPhone(value *string) {
	m.phone = value
}

// SetWebsite sets the website property value. The website property
func (m *BusinessProfilePatchRequestBody) SetWebsite(value *string) {
	m.website = value
}

type BusinessProfilePatchRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAddress() BusinessProfilePatchRequestBody_addressable
	GetBusinessGoals() []string
	GetBusinessPrimaryColor() *string
	GetBusinessSecondaryColor() *string
	GetCurrencyCode() *string
	GetEmail() *string
	GetName() *string
	GetPhone() *string
	GetWebsite() *string
	SetAddress(value BusinessProfilePatchRequestBody_addressable)
	SetBusinessGoals(value []string)
	SetBusinessPrimaryColor(value *string)
	SetBusinessSecondaryColor(value *string)
	SetCurrencyCode(value *string)
	SetEmail(value *string)
	SetName(value *string)
	SetPhone(value *string)
	SetWebsite(value *string)
}
