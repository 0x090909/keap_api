package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ContactsPostResponse_addresses struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The long-name descriptive version of the Country Code
	country *string
	// An ISO 3166-2 Country Code (https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3)
	country_code *string
	// The line1 property
	line1 *string
	// The line2 property
	line2 *string
	// The municipality to which the address belongs
	locality *string
	// The postal_code property
	postal_code *string
	// The long-name descriptive version of the Region Code
	region *string
	// An ISO 3166-2 Province Code, such as one of the US States (https://en.wikipedia.org/wiki/ISO_3166-2:US)
	region_code *string
	// The zip_code property
	zip_code *string
	// The zip_four property
	zip_four *string
}

// NewContactsPostResponse_addresses instantiates a new ContactsPostResponse_addresses and sets the default values.
func NewContactsPostResponse_addresses() *ContactsPostResponse_addresses {
	m := &ContactsPostResponse_addresses{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateContactsPostResponse_addressesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsPostResponse_addressesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsPostResponse_addresses(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContactsPostResponse_addresses) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCountry gets the country property value. The long-name descriptive version of the Country Code
// returns a *string when successful
func (m *ContactsPostResponse_addresses) GetCountry() *string {
	return m.country
}

// GetCountryCode gets the country_code property value. An ISO 3166-2 Country Code (https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3)
// returns a *string when successful
func (m *ContactsPostResponse_addresses) GetCountryCode() *string {
	return m.country_code
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContactsPostResponse_addresses) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
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
	res["country_code"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCountryCode(val)
		}
		return nil
	}
	res["line1"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLine1(val)
		}
		return nil
	}
	res["line2"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLine2(val)
		}
		return nil
	}
	res["locality"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLocality(val)
		}
		return nil
	}
	res["postal_code"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPostalCode(val)
		}
		return nil
	}
	res["region"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRegion(val)
		}
		return nil
	}
	res["region_code"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRegionCode(val)
		}
		return nil
	}
	res["zip_code"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetZipCode(val)
		}
		return nil
	}
	res["zip_four"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetZipFour(val)
		}
		return nil
	}
	return res
}

// GetLine1 gets the line1 property value. The line1 property
// returns a *string when successful
func (m *ContactsPostResponse_addresses) GetLine1() *string {
	return m.line1
}

// GetLine2 gets the line2 property value. The line2 property
// returns a *string when successful
func (m *ContactsPostResponse_addresses) GetLine2() *string {
	return m.line2
}

// GetLocality gets the locality property value. The municipality to which the address belongs
// returns a *string when successful
func (m *ContactsPostResponse_addresses) GetLocality() *string {
	return m.locality
}

// GetPostalCode gets the postal_code property value. The postal_code property
// returns a *string when successful
func (m *ContactsPostResponse_addresses) GetPostalCode() *string {
	return m.postal_code
}

// GetRegion gets the region property value. The long-name descriptive version of the Region Code
// returns a *string when successful
func (m *ContactsPostResponse_addresses) GetRegion() *string {
	return m.region
}

// GetRegionCode gets the region_code property value. An ISO 3166-2 Province Code, such as one of the US States (https://en.wikipedia.org/wiki/ISO_3166-2:US)
// returns a *string when successful
func (m *ContactsPostResponse_addresses) GetRegionCode() *string {
	return m.region_code
}

// GetZipCode gets the zip_code property value. The zip_code property
// returns a *string when successful
func (m *ContactsPostResponse_addresses) GetZipCode() *string {
	return m.zip_code
}

// GetZipFour gets the zip_four property value. The zip_four property
// returns a *string when successful
func (m *ContactsPostResponse_addresses) GetZipFour() *string {
	return m.zip_four
}

// Serialize serializes information the current object
func (m *ContactsPostResponse_addresses) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("country", m.GetCountry())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("country_code", m.GetCountryCode())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("line1", m.GetLine1())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("line2", m.GetLine2())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("locality", m.GetLocality())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("postal_code", m.GetPostalCode())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("region", m.GetRegion())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("region_code", m.GetRegionCode())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("zip_code", m.GetZipCode())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("zip_four", m.GetZipFour())
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
func (m *ContactsPostResponse_addresses) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCountry sets the country property value. The long-name descriptive version of the Country Code
func (m *ContactsPostResponse_addresses) SetCountry(value *string) {
	m.country = value
}

// SetCountryCode sets the country_code property value. An ISO 3166-2 Country Code (https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3)
func (m *ContactsPostResponse_addresses) SetCountryCode(value *string) {
	m.country_code = value
}

// SetLine1 sets the line1 property value. The line1 property
func (m *ContactsPostResponse_addresses) SetLine1(value *string) {
	m.line1 = value
}

// SetLine2 sets the line2 property value. The line2 property
func (m *ContactsPostResponse_addresses) SetLine2(value *string) {
	m.line2 = value
}

// SetLocality sets the locality property value. The municipality to which the address belongs
func (m *ContactsPostResponse_addresses) SetLocality(value *string) {
	m.locality = value
}

// SetPostalCode sets the postal_code property value. The postal_code property
func (m *ContactsPostResponse_addresses) SetPostalCode(value *string) {
	m.postal_code = value
}

// SetRegion sets the region property value. The long-name descriptive version of the Region Code
func (m *ContactsPostResponse_addresses) SetRegion(value *string) {
	m.region = value
}

// SetRegionCode sets the region_code property value. An ISO 3166-2 Province Code, such as one of the US States (https://en.wikipedia.org/wiki/ISO_3166-2:US)
func (m *ContactsPostResponse_addresses) SetRegionCode(value *string) {
	m.region_code = value
}

// SetZipCode sets the zip_code property value. The zip_code property
func (m *ContactsPostResponse_addresses) SetZipCode(value *string) {
	m.zip_code = value
}

// SetZipFour sets the zip_four property value. The zip_four property
func (m *ContactsPostResponse_addresses) SetZipFour(value *string) {
	m.zip_four = value
}

type ContactsPostResponse_addressesable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCountry() *string
	GetCountryCode() *string
	GetLine1() *string
	GetLine2() *string
	GetLocality() *string
	GetPostalCode() *string
	GetRegion() *string
	GetRegionCode() *string
	GetZipCode() *string
	GetZipFour() *string
	SetCountry(value *string)
	SetCountryCode(value *string)
	SetLine1(value *string)
	SetLine2(value *string)
	SetLocality(value *string)
	SetPostalCode(value *string)
	SetRegion(value *string)
	SetRegionCode(value *string)
	SetZipCode(value *string)
	SetZipFour(value *string)
}
