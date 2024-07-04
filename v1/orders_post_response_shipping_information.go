package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OrdersPostResponse_shipping_information struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The company property
	company *string
	// The country_code property
	country_code *string
	// The first_name property
	first_name *string
	// The is_invoice_to_company property
	is_invoice_to_company *bool
	// The last_name property
	last_name *string
	// The line1 property
	line1 *string
	// The line2 property
	line2 *string
	// The locality property
	locality *string
	// The middle_name property
	middle_name *string
	// The phone property
	phone *string
	// The region property
	region *string
	// The zip_code property
	zip_code *string
	// The zip_four property
	zip_four *string
}

// NewOrdersPostResponse_shipping_information instantiates a new OrdersPostResponse_shipping_information and sets the default values.
func NewOrdersPostResponse_shipping_information() *OrdersPostResponse_shipping_information {
	m := &OrdersPostResponse_shipping_information{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateOrdersPostResponse_shipping_informationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrdersPostResponse_shipping_informationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewOrdersPostResponse_shipping_information(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OrdersPostResponse_shipping_information) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCompany gets the company property value. The company property
// returns a *string when successful
func (m *OrdersPostResponse_shipping_information) GetCompany() *string {
	return m.company
}

// GetCountryCode gets the country_code property value. The country_code property
// returns a *string when successful
func (m *OrdersPostResponse_shipping_information) GetCountryCode() *string {
	return m.country_code
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrdersPostResponse_shipping_information) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["company"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCompany(val)
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
	res["first_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFirstName(val)
		}
		return nil
	}
	res["is_invoice_to_company"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetIsInvoiceToCompany(val)
		}
		return nil
	}
	res["last_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLastName(val)
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
	res["middle_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetMiddleName(val)
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

// GetFirstName gets the first_name property value. The first_name property
// returns a *string when successful
func (m *OrdersPostResponse_shipping_information) GetFirstName() *string {
	return m.first_name
}

// GetIsInvoiceToCompany gets the is_invoice_to_company property value. The is_invoice_to_company property
// returns a *bool when successful
func (m *OrdersPostResponse_shipping_information) GetIsInvoiceToCompany() *bool {
	return m.is_invoice_to_company
}

// GetLastName gets the last_name property value. The last_name property
// returns a *string when successful
func (m *OrdersPostResponse_shipping_information) GetLastName() *string {
	return m.last_name
}

// GetLine1 gets the line1 property value. The line1 property
// returns a *string when successful
func (m *OrdersPostResponse_shipping_information) GetLine1() *string {
	return m.line1
}

// GetLine2 gets the line2 property value. The line2 property
// returns a *string when successful
func (m *OrdersPostResponse_shipping_information) GetLine2() *string {
	return m.line2
}

// GetLocality gets the locality property value. The locality property
// returns a *string when successful
func (m *OrdersPostResponse_shipping_information) GetLocality() *string {
	return m.locality
}

// GetMiddleName gets the middle_name property value. The middle_name property
// returns a *string when successful
func (m *OrdersPostResponse_shipping_information) GetMiddleName() *string {
	return m.middle_name
}

// GetPhone gets the phone property value. The phone property
// returns a *string when successful
func (m *OrdersPostResponse_shipping_information) GetPhone() *string {
	return m.phone
}

// GetRegion gets the region property value. The region property
// returns a *string when successful
func (m *OrdersPostResponse_shipping_information) GetRegion() *string {
	return m.region
}

// GetZipCode gets the zip_code property value. The zip_code property
// returns a *string when successful
func (m *OrdersPostResponse_shipping_information) GetZipCode() *string {
	return m.zip_code
}

// GetZipFour gets the zip_four property value. The zip_four property
// returns a *string when successful
func (m *OrdersPostResponse_shipping_information) GetZipFour() *string {
	return m.zip_four
}

// Serialize serializes information the current object
func (m *OrdersPostResponse_shipping_information) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("company", m.GetCompany())
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
		err := writer.WriteStringValue("first_name", m.GetFirstName())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("is_invoice_to_company", m.GetIsInvoiceToCompany())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("last_name", m.GetLastName())
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
		err := writer.WriteStringValue("middle_name", m.GetMiddleName())
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
func (m *OrdersPostResponse_shipping_information) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCompany sets the company property value. The company property
func (m *OrdersPostResponse_shipping_information) SetCompany(value *string) {
	m.company = value
}

// SetCountryCode sets the country_code property value. The country_code property
func (m *OrdersPostResponse_shipping_information) SetCountryCode(value *string) {
	m.country_code = value
}

// SetFirstName sets the first_name property value. The first_name property
func (m *OrdersPostResponse_shipping_information) SetFirstName(value *string) {
	m.first_name = value
}

// SetIsInvoiceToCompany sets the is_invoice_to_company property value. The is_invoice_to_company property
func (m *OrdersPostResponse_shipping_information) SetIsInvoiceToCompany(value *bool) {
	m.is_invoice_to_company = value
}

// SetLastName sets the last_name property value. The last_name property
func (m *OrdersPostResponse_shipping_information) SetLastName(value *string) {
	m.last_name = value
}

// SetLine1 sets the line1 property value. The line1 property
func (m *OrdersPostResponse_shipping_information) SetLine1(value *string) {
	m.line1 = value
}

// SetLine2 sets the line2 property value. The line2 property
func (m *OrdersPostResponse_shipping_information) SetLine2(value *string) {
	m.line2 = value
}

// SetLocality sets the locality property value. The locality property
func (m *OrdersPostResponse_shipping_information) SetLocality(value *string) {
	m.locality = value
}

// SetMiddleName sets the middle_name property value. The middle_name property
func (m *OrdersPostResponse_shipping_information) SetMiddleName(value *string) {
	m.middle_name = value
}

// SetPhone sets the phone property value. The phone property
func (m *OrdersPostResponse_shipping_information) SetPhone(value *string) {
	m.phone = value
}

// SetRegion sets the region property value. The region property
func (m *OrdersPostResponse_shipping_information) SetRegion(value *string) {
	m.region = value
}

// SetZipCode sets the zip_code property value. The zip_code property
func (m *OrdersPostResponse_shipping_information) SetZipCode(value *string) {
	m.zip_code = value
}

// SetZipFour sets the zip_four property value. The zip_four property
func (m *OrdersPostResponse_shipping_information) SetZipFour(value *string) {
	m.zip_four = value
}

type OrdersPostResponse_shipping_informationable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCompany() *string
	GetCountryCode() *string
	GetFirstName() *string
	GetIsInvoiceToCompany() *bool
	GetLastName() *string
	GetLine1() *string
	GetLine2() *string
	GetLocality() *string
	GetMiddleName() *string
	GetPhone() *string
	GetRegion() *string
	GetZipCode() *string
	GetZipFour() *string
	SetCompany(value *string)
	SetCountryCode(value *string)
	SetFirstName(value *string)
	SetIsInvoiceToCompany(value *bool)
	SetLastName(value *string)
	SetLine1(value *string)
	SetLine2(value *string)
	SetLocality(value *string)
	SetMiddleName(value *string)
	SetPhone(value *string)
	SetRegion(value *string)
	SetZipCode(value *string)
	SetZipFour(value *string)
}
