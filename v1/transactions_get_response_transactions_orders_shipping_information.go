package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TransactionsGetResponse_transactions_orders_shipping_information struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The city property
	city *string
	// The company property
	company *string
	// The country property
	country *string
	// The first_name property
	first_name *string
	// The id property
	id *int64
	// The invoiceToCompany property
	invoiceToCompany *bool
	// The last_name property
	last_name *string
	// The middle_name property
	middle_name *string
	// The phone property
	phone *string
	// The state property
	state *string
	// The street1 property
	street1 *string
	// The street2 property
	street2 *string
	// The zip property
	zip *string
}

// NewTransactionsGetResponse_transactions_orders_shipping_information instantiates a new TransactionsGetResponse_transactions_orders_shipping_information and sets the default values.
func NewTransactionsGetResponse_transactions_orders_shipping_information() *TransactionsGetResponse_transactions_orders_shipping_information {
	m := &TransactionsGetResponse_transactions_orders_shipping_information{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateTransactionsGetResponse_transactions_orders_shipping_informationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionsGetResponse_transactions_orders_shipping_informationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTransactionsGetResponse_transactions_orders_shipping_information(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransactionsGetResponse_transactions_orders_shipping_information) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCity gets the city property value. The city property
// returns a *string when successful
func (m *TransactionsGetResponse_transactions_orders_shipping_information) GetCity() *string {
	return m.city
}

// GetCompany gets the company property value. The company property
// returns a *string when successful
func (m *TransactionsGetResponse_transactions_orders_shipping_information) GetCompany() *string {
	return m.company
}

// GetCountry gets the country property value. The country property
// returns a *string when successful
func (m *TransactionsGetResponse_transactions_orders_shipping_information) GetCountry() *string {
	return m.country
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransactionsGetResponse_transactions_orders_shipping_information) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
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
	res["id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetId(val)
		}
		return nil
	}
	res["invoiceToCompany"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetInvoiceToCompany(val)
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
	res["street1"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetStreet1(val)
		}
		return nil
	}
	res["street2"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetStreet2(val)
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

// GetFirstName gets the first_name property value. The first_name property
// returns a *string when successful
func (m *TransactionsGetResponse_transactions_orders_shipping_information) GetFirstName() *string {
	return m.first_name
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *TransactionsGetResponse_transactions_orders_shipping_information) GetId() *int64 {
	return m.id
}

// GetInvoiceToCompany gets the invoiceToCompany property value. The invoiceToCompany property
// returns a *bool when successful
func (m *TransactionsGetResponse_transactions_orders_shipping_information) GetInvoiceToCompany() *bool {
	return m.invoiceToCompany
}

// GetLastName gets the last_name property value. The last_name property
// returns a *string when successful
func (m *TransactionsGetResponse_transactions_orders_shipping_information) GetLastName() *string {
	return m.last_name
}

// GetMiddleName gets the middle_name property value. The middle_name property
// returns a *string when successful
func (m *TransactionsGetResponse_transactions_orders_shipping_information) GetMiddleName() *string {
	return m.middle_name
}

// GetPhone gets the phone property value. The phone property
// returns a *string when successful
func (m *TransactionsGetResponse_transactions_orders_shipping_information) GetPhone() *string {
	return m.phone
}

// GetState gets the state property value. The state property
// returns a *string when successful
func (m *TransactionsGetResponse_transactions_orders_shipping_information) GetState() *string {
	return m.state
}

// GetStreet1 gets the street1 property value. The street1 property
// returns a *string when successful
func (m *TransactionsGetResponse_transactions_orders_shipping_information) GetStreet1() *string {
	return m.street1
}

// GetStreet2 gets the street2 property value. The street2 property
// returns a *string when successful
func (m *TransactionsGetResponse_transactions_orders_shipping_information) GetStreet2() *string {
	return m.street2
}

// GetZip gets the zip property value. The zip property
// returns a *string when successful
func (m *TransactionsGetResponse_transactions_orders_shipping_information) GetZip() *string {
	return m.zip
}

// Serialize serializes information the current object
func (m *TransactionsGetResponse_transactions_orders_shipping_information) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("city", m.GetCity())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("company", m.GetCompany())
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
		err := writer.WriteStringValue("first_name", m.GetFirstName())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("id", m.GetId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("invoiceToCompany", m.GetInvoiceToCompany())
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
		err := writer.WriteStringValue("state", m.GetState())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("street1", m.GetStreet1())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("street2", m.GetStreet2())
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
func (m *TransactionsGetResponse_transactions_orders_shipping_information) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCity sets the city property value. The city property
func (m *TransactionsGetResponse_transactions_orders_shipping_information) SetCity(value *string) {
	m.city = value
}

// SetCompany sets the company property value. The company property
func (m *TransactionsGetResponse_transactions_orders_shipping_information) SetCompany(value *string) {
	m.company = value
}

// SetCountry sets the country property value. The country property
func (m *TransactionsGetResponse_transactions_orders_shipping_information) SetCountry(value *string) {
	m.country = value
}

// SetFirstName sets the first_name property value. The first_name property
func (m *TransactionsGetResponse_transactions_orders_shipping_information) SetFirstName(value *string) {
	m.first_name = value
}

// SetId sets the id property value. The id property
func (m *TransactionsGetResponse_transactions_orders_shipping_information) SetId(value *int64) {
	m.id = value
}

// SetInvoiceToCompany sets the invoiceToCompany property value. The invoiceToCompany property
func (m *TransactionsGetResponse_transactions_orders_shipping_information) SetInvoiceToCompany(value *bool) {
	m.invoiceToCompany = value
}

// SetLastName sets the last_name property value. The last_name property
func (m *TransactionsGetResponse_transactions_orders_shipping_information) SetLastName(value *string) {
	m.last_name = value
}

// SetMiddleName sets the middle_name property value. The middle_name property
func (m *TransactionsGetResponse_transactions_orders_shipping_information) SetMiddleName(value *string) {
	m.middle_name = value
}

// SetPhone sets the phone property value. The phone property
func (m *TransactionsGetResponse_transactions_orders_shipping_information) SetPhone(value *string) {
	m.phone = value
}

// SetState sets the state property value. The state property
func (m *TransactionsGetResponse_transactions_orders_shipping_information) SetState(value *string) {
	m.state = value
}

// SetStreet1 sets the street1 property value. The street1 property
func (m *TransactionsGetResponse_transactions_orders_shipping_information) SetStreet1(value *string) {
	m.street1 = value
}

// SetStreet2 sets the street2 property value. The street2 property
func (m *TransactionsGetResponse_transactions_orders_shipping_information) SetStreet2(value *string) {
	m.street2 = value
}

// SetZip sets the zip property value. The zip property
func (m *TransactionsGetResponse_transactions_orders_shipping_information) SetZip(value *string) {
	m.zip = value
}

type TransactionsGetResponse_transactions_orders_shipping_informationable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCity() *string
	GetCompany() *string
	GetCountry() *string
	GetFirstName() *string
	GetId() *int64
	GetInvoiceToCompany() *bool
	GetLastName() *string
	GetMiddleName() *string
	GetPhone() *string
	GetState() *string
	GetStreet1() *string
	GetStreet2() *string
	GetZip() *string
	SetCity(value *string)
	SetCompany(value *string)
	SetCountry(value *string)
	SetFirstName(value *string)
	SetId(value *int64)
	SetInvoiceToCompany(value *bool)
	SetLastName(value *string)
	SetMiddleName(value *string)
	SetPhone(value *string)
	SetState(value *string)
	SetStreet1(value *string)
	SetStreet2(value *string)
	SetZip(value *string)
}
