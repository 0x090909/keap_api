package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AccountProfilePutRequestBody_address struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The country_code property
	country_code *string
	// The line1 property
	line1 *string
	// The line2 property
	line2 *string
	// The locality property
	locality *string
	// Field used to store postal codes containing a combination of letters and numbers ex. 'EC1A', 'S1 2HE', '75000'
	postal_code *string
	// The region property
	region *string
	// Mainly used in the United States, this is typically numeric. ex. '85001', '90002' Note: this is to be used instead of 'postal_code', not in addition to.
	zip_code *string
	// Last four of a full zip code ex. '8244', '4320'. This field is supplemental to the zip_code field, otherwise will be ignored.
	zip_four *string
}

// NewAccountProfilePutRequestBody_address instantiates a new AccountProfilePutRequestBody_address and sets the default values.
func NewAccountProfilePutRequestBody_address() *AccountProfilePutRequestBody_address {
	m := &AccountProfilePutRequestBody_address{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateAccountProfilePutRequestBody_addressFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAccountProfilePutRequestBody_addressFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAccountProfilePutRequestBody_address(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AccountProfilePutRequestBody_address) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCountryCode gets the country_code property value. The country_code property
// returns a *string when successful
func (m *AccountProfilePutRequestBody_address) GetCountryCode() *string {
	return m.country_code
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AccountProfilePutRequestBody_address) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
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
func (m *AccountProfilePutRequestBody_address) GetLine1() *string {
	return m.line1
}

// GetLine2 gets the line2 property value. The line2 property
// returns a *string when successful
func (m *AccountProfilePutRequestBody_address) GetLine2() *string {
	return m.line2
}

// GetLocality gets the locality property value. The locality property
// returns a *string when successful
func (m *AccountProfilePutRequestBody_address) GetLocality() *string {
	return m.locality
}

// GetPostalCode gets the postal_code property value. Field used to store postal codes containing a combination of letters and numbers ex. 'EC1A', 'S1 2HE', '75000'
// returns a *string when successful
func (m *AccountProfilePutRequestBody_address) GetPostalCode() *string {
	return m.postal_code
}

// GetRegion gets the region property value. The region property
// returns a *string when successful
func (m *AccountProfilePutRequestBody_address) GetRegion() *string {
	return m.region
}

// GetZipCode gets the zip_code property value. Mainly used in the United States, this is typically numeric. ex. '85001', '90002' Note: this is to be used instead of 'postal_code', not in addition to.
// returns a *string when successful
func (m *AccountProfilePutRequestBody_address) GetZipCode() *string {
	return m.zip_code
}

// GetZipFour gets the zip_four property value. Last four of a full zip code ex. '8244', '4320'. This field is supplemental to the zip_code field, otherwise will be ignored.
// returns a *string when successful
func (m *AccountProfilePutRequestBody_address) GetZipFour() *string {
	return m.zip_four
}

// Serialize serializes information the current object
func (m *AccountProfilePutRequestBody_address) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *AccountProfilePutRequestBody_address) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCountryCode sets the country_code property value. The country_code property
func (m *AccountProfilePutRequestBody_address) SetCountryCode(value *string) {
	m.country_code = value
}

// SetLine1 sets the line1 property value. The line1 property
func (m *AccountProfilePutRequestBody_address) SetLine1(value *string) {
	m.line1 = value
}

// SetLine2 sets the line2 property value. The line2 property
func (m *AccountProfilePutRequestBody_address) SetLine2(value *string) {
	m.line2 = value
}

// SetLocality sets the locality property value. The locality property
func (m *AccountProfilePutRequestBody_address) SetLocality(value *string) {
	m.locality = value
}

// SetPostalCode sets the postal_code property value. Field used to store postal codes containing a combination of letters and numbers ex. 'EC1A', 'S1 2HE', '75000'
func (m *AccountProfilePutRequestBody_address) SetPostalCode(value *string) {
	m.postal_code = value
}

// SetRegion sets the region property value. The region property
func (m *AccountProfilePutRequestBody_address) SetRegion(value *string) {
	m.region = value
}

// SetZipCode sets the zip_code property value. Mainly used in the United States, this is typically numeric. ex. '85001', '90002' Note: this is to be used instead of 'postal_code', not in addition to.
func (m *AccountProfilePutRequestBody_address) SetZipCode(value *string) {
	m.zip_code = value
}

// SetZipFour sets the zip_four property value. Last four of a full zip code ex. '8244', '4320'. This field is supplemental to the zip_code field, otherwise will be ignored.
func (m *AccountProfilePutRequestBody_address) SetZipFour(value *string) {
	m.zip_four = value
}

type AccountProfilePutRequestBody_addressable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCountryCode() *string
	GetLine1() *string
	GetLine2() *string
	GetLocality() *string
	GetPostalCode() *string
	GetRegion() *string
	GetZipCode() *string
	GetZipFour() *string
	SetCountryCode(value *string)
	SetLine1(value *string)
	SetLine2(value *string)
	SetLocality(value *string)
	SetPostalCode(value *string)
	SetRegion(value *string)
	SetZipCode(value *string)
	SetZipFour(value *string)
}
