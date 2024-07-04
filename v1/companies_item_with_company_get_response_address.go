package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CompaniesItemWithCompanyGetResponse_address struct {
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
	// The region property
	region *string
	// The zip_code property
	zip_code *string
	// The zip_four property
	zip_four *string
}

// NewCompaniesItemWithCompanyGetResponse_address instantiates a new CompaniesItemWithCompanyGetResponse_address and sets the default values.
func NewCompaniesItemWithCompanyGetResponse_address() *CompaniesItemWithCompanyGetResponse_address {
	m := &CompaniesItemWithCompanyGetResponse_address{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateCompaniesItemWithCompanyGetResponse_addressFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCompaniesItemWithCompanyGetResponse_addressFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCompaniesItemWithCompanyGetResponse_address(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CompaniesItemWithCompanyGetResponse_address) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCountryCode gets the country_code property value. The country_code property
// returns a *string when successful
func (m *CompaniesItemWithCompanyGetResponse_address) GetCountryCode() *string {
	return m.country_code
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CompaniesItemWithCompanyGetResponse_address) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
func (m *CompaniesItemWithCompanyGetResponse_address) GetLine1() *string {
	return m.line1
}

// GetLine2 gets the line2 property value. The line2 property
// returns a *string when successful
func (m *CompaniesItemWithCompanyGetResponse_address) GetLine2() *string {
	return m.line2
}

// GetLocality gets the locality property value. The locality property
// returns a *string when successful
func (m *CompaniesItemWithCompanyGetResponse_address) GetLocality() *string {
	return m.locality
}

// GetRegion gets the region property value. The region property
// returns a *string when successful
func (m *CompaniesItemWithCompanyGetResponse_address) GetRegion() *string {
	return m.region
}

// GetZipCode gets the zip_code property value. The zip_code property
// returns a *string when successful
func (m *CompaniesItemWithCompanyGetResponse_address) GetZipCode() *string {
	return m.zip_code
}

// GetZipFour gets the zip_four property value. The zip_four property
// returns a *string when successful
func (m *CompaniesItemWithCompanyGetResponse_address) GetZipFour() *string {
	return m.zip_four
}

// Serialize serializes information the current object
func (m *CompaniesItemWithCompanyGetResponse_address) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *CompaniesItemWithCompanyGetResponse_address) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCountryCode sets the country_code property value. The country_code property
func (m *CompaniesItemWithCompanyGetResponse_address) SetCountryCode(value *string) {
	m.country_code = value
}

// SetLine1 sets the line1 property value. The line1 property
func (m *CompaniesItemWithCompanyGetResponse_address) SetLine1(value *string) {
	m.line1 = value
}

// SetLine2 sets the line2 property value. The line2 property
func (m *CompaniesItemWithCompanyGetResponse_address) SetLine2(value *string) {
	m.line2 = value
}

// SetLocality sets the locality property value. The locality property
func (m *CompaniesItemWithCompanyGetResponse_address) SetLocality(value *string) {
	m.locality = value
}

// SetRegion sets the region property value. The region property
func (m *CompaniesItemWithCompanyGetResponse_address) SetRegion(value *string) {
	m.region = value
}

// SetZipCode sets the zip_code property value. The zip_code property
func (m *CompaniesItemWithCompanyGetResponse_address) SetZipCode(value *string) {
	m.zip_code = value
}

// SetZipFour sets the zip_four property value. The zip_four property
func (m *CompaniesItemWithCompanyGetResponse_address) SetZipFour(value *string) {
	m.zip_four = value
}

type CompaniesItemWithCompanyGetResponse_addressable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCountryCode() *string
	GetLine1() *string
	GetLine2() *string
	GetLocality() *string
	GetRegion() *string
	GetZipCode() *string
	GetZipFour() *string
	SetCountryCode(value *string)
	SetLine1(value *string)
	SetLine2(value *string)
	SetLocality(value *string)
	SetRegion(value *string)
	SetZipCode(value *string)
	SetZipFour(value *string)
}
