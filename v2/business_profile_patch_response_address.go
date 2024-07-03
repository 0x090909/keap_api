package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type BusinessProfilePatchResponse_address struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// An ISO 3166-1 Country Code (https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3)
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
}

// NewBusinessProfilePatchResponse_address instantiates a new BusinessProfilePatchResponse_address and sets the default values.
func NewBusinessProfilePatchResponse_address() *BusinessProfilePatchResponse_address {
	m := &BusinessProfilePatchResponse_address{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateBusinessProfilePatchResponse_addressFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBusinessProfilePatchResponse_addressFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewBusinessProfilePatchResponse_address(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *BusinessProfilePatchResponse_address) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCountryCode gets the country_code property value. An ISO 3166-1 Country Code (https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3)
// returns a *string when successful
func (m *BusinessProfilePatchResponse_address) GetCountryCode() *string {
	return m.country_code
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *BusinessProfilePatchResponse_address) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
	return res
}

// GetLine1 gets the line1 property value. The line1 property
// returns a *string when successful
func (m *BusinessProfilePatchResponse_address) GetLine1() *string {
	return m.line1
}

// GetLine2 gets the line2 property value. The line2 property
// returns a *string when successful
func (m *BusinessProfilePatchResponse_address) GetLine2() *string {
	return m.line2
}

// GetLocality gets the locality property value. The municipality to which the address belongs
// returns a *string when successful
func (m *BusinessProfilePatchResponse_address) GetLocality() *string {
	return m.locality
}

// GetPostalCode gets the postal_code property value. The postal_code property
// returns a *string when successful
func (m *BusinessProfilePatchResponse_address) GetPostalCode() *string {
	return m.postal_code
}

// GetRegion gets the region property value. The long-name descriptive version of the Region Code
// returns a *string when successful
func (m *BusinessProfilePatchResponse_address) GetRegion() *string {
	return m.region
}

// Serialize serializes information the current object
func (m *BusinessProfilePatchResponse_address) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
		err := writer.WriteAdditionalData(m.GetAdditionalData())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BusinessProfilePatchResponse_address) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCountryCode sets the country_code property value. An ISO 3166-1 Country Code (https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3)
func (m *BusinessProfilePatchResponse_address) SetCountryCode(value *string) {
	m.country_code = value
}

// SetLine1 sets the line1 property value. The line1 property
func (m *BusinessProfilePatchResponse_address) SetLine1(value *string) {
	m.line1 = value
}

// SetLine2 sets the line2 property value. The line2 property
func (m *BusinessProfilePatchResponse_address) SetLine2(value *string) {
	m.line2 = value
}

// SetLocality sets the locality property value. The municipality to which the address belongs
func (m *BusinessProfilePatchResponse_address) SetLocality(value *string) {
	m.locality = value
}

// SetPostalCode sets the postal_code property value. The postal_code property
func (m *BusinessProfilePatchResponse_address) SetPostalCode(value *string) {
	m.postal_code = value
}

// SetRegion sets the region property value. The long-name descriptive version of the Region Code
func (m *BusinessProfilePatchResponse_address) SetRegion(value *string) {
	m.region = value
}

type BusinessProfilePatchResponse_addressable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCountryCode() *string
	GetLine1() *string
	GetLine2() *string
	GetLocality() *string
	GetPostalCode() *string
	GetRegion() *string
	SetCountryCode(value *string)
	SetLine1(value *string)
	SetLine2(value *string)
	SetLocality(value *string)
	SetPostalCode(value *string)
	SetRegion(value *string)
}
