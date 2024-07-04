package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OrdersItemTransactionsGetResponse_transactions_orders_contact struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The company_name property
	company_name *string
	// The email property
	email *string
	// The first_name property
	first_name *string
	// The id property
	id *int64
	// The job_title property
	job_title *string
	// The last_name property
	last_name *string
}

// NewOrdersItemTransactionsGetResponse_transactions_orders_contact instantiates a new OrdersItemTransactionsGetResponse_transactions_orders_contact and sets the default values.
func NewOrdersItemTransactionsGetResponse_transactions_orders_contact() *OrdersItemTransactionsGetResponse_transactions_orders_contact {
	m := &OrdersItemTransactionsGetResponse_transactions_orders_contact{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateOrdersItemTransactionsGetResponse_transactions_orders_contactFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrdersItemTransactionsGetResponse_transactions_orders_contactFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewOrdersItemTransactionsGetResponse_transactions_orders_contact(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OrdersItemTransactionsGetResponse_transactions_orders_contact) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCompanyName gets the company_name property value. The company_name property
// returns a *string when successful
func (m *OrdersItemTransactionsGetResponse_transactions_orders_contact) GetCompanyName() *string {
	return m.company_name
}

// GetEmail gets the email property value. The email property
// returns a *string when successful
func (m *OrdersItemTransactionsGetResponse_transactions_orders_contact) GetEmail() *string {
	return m.email
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrdersItemTransactionsGetResponse_transactions_orders_contact) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["company_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCompanyName(val)
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
	res["job_title"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetJobTitle(val)
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
	return res
}

// GetFirstName gets the first_name property value. The first_name property
// returns a *string when successful
func (m *OrdersItemTransactionsGetResponse_transactions_orders_contact) GetFirstName() *string {
	return m.first_name
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *OrdersItemTransactionsGetResponse_transactions_orders_contact) GetId() *int64 {
	return m.id
}

// GetJobTitle gets the job_title property value. The job_title property
// returns a *string when successful
func (m *OrdersItemTransactionsGetResponse_transactions_orders_contact) GetJobTitle() *string {
	return m.job_title
}

// GetLastName gets the last_name property value. The last_name property
// returns a *string when successful
func (m *OrdersItemTransactionsGetResponse_transactions_orders_contact) GetLastName() *string {
	return m.last_name
}

// Serialize serializes information the current object
func (m *OrdersItemTransactionsGetResponse_transactions_orders_contact) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("company_name", m.GetCompanyName())
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
		err := writer.WriteStringValue("job_title", m.GetJobTitle())
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
		err := writer.WriteAdditionalData(m.GetAdditionalData())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OrdersItemTransactionsGetResponse_transactions_orders_contact) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCompanyName sets the company_name property value. The company_name property
func (m *OrdersItemTransactionsGetResponse_transactions_orders_contact) SetCompanyName(value *string) {
	m.company_name = value
}

// SetEmail sets the email property value. The email property
func (m *OrdersItemTransactionsGetResponse_transactions_orders_contact) SetEmail(value *string) {
	m.email = value
}

// SetFirstName sets the first_name property value. The first_name property
func (m *OrdersItemTransactionsGetResponse_transactions_orders_contact) SetFirstName(value *string) {
	m.first_name = value
}

// SetId sets the id property value. The id property
func (m *OrdersItemTransactionsGetResponse_transactions_orders_contact) SetId(value *int64) {
	m.id = value
}

// SetJobTitle sets the job_title property value. The job_title property
func (m *OrdersItemTransactionsGetResponse_transactions_orders_contact) SetJobTitle(value *string) {
	m.job_title = value
}

// SetLastName sets the last_name property value. The last_name property
func (m *OrdersItemTransactionsGetResponse_transactions_orders_contact) SetLastName(value *string) {
	m.last_name = value
}

type OrdersItemTransactionsGetResponse_transactions_orders_contactable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCompanyName() *string
	GetEmail() *string
	GetFirstName() *string
	GetId() *int64
	GetJobTitle() *string
	GetLastName() *string
	SetCompanyName(value *string)
	SetEmail(value *string)
	SetFirstName(value *string)
	SetId(value *int64)
	SetJobTitle(value *string)
	SetLastName(value *string)
}
