package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TransactionsItemWithTransactionGetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The amount property
	amount *float64
	// The collection_method property
	collection_method *string
	// The contact_id property
	contact_id *int64
	// The currency property
	currency *string
	// The errors property
	errors *string
	// The gateway property
	gateway *string
	// The gateway_account_name property
	gateway_account_name *string
	// The id property
	id *int64
	// The order_ids property
	order_ids *string
	// The orders property
	orders []TransactionsItemWithTransactionGetResponse_ordersable
	// The paymentDate property
	paymentDate *string
	// The status property
	status *string
	// The test property
	test *bool
	// The transaction_date property
	transaction_date *string
	// The type property
	typeEscaped *string
}

// NewTransactionsItemWithTransactionGetResponse instantiates a new TransactionsItemWithTransactionGetResponse and sets the default values.
func NewTransactionsItemWithTransactionGetResponse() *TransactionsItemWithTransactionGetResponse {
	m := &TransactionsItemWithTransactionGetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateTransactionsItemWithTransactionGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionsItemWithTransactionGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTransactionsItemWithTransactionGetResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransactionsItemWithTransactionGetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAmount gets the amount property value. The amount property
// returns a *float64 when successful
func (m *TransactionsItemWithTransactionGetResponse) GetAmount() *float64 {
	return m.amount
}

// GetCollectionMethod gets the collection_method property value. The collection_method property
// returns a *string when successful
func (m *TransactionsItemWithTransactionGetResponse) GetCollectionMethod() *string {
	return m.collection_method
}

// GetContactId gets the contact_id property value. The contact_id property
// returns a *int64 when successful
func (m *TransactionsItemWithTransactionGetResponse) GetContactId() *int64 {
	return m.contact_id
}

// GetCurrency gets the currency property value. The currency property
// returns a *string when successful
func (m *TransactionsItemWithTransactionGetResponse) GetCurrency() *string {
	return m.currency
}

// GetErrors gets the errors property value. The errors property
// returns a *string when successful
func (m *TransactionsItemWithTransactionGetResponse) GetErrors() *string {
	return m.errors
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransactionsItemWithTransactionGetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["amount"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetFloat64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAmount(val)
		}
		return nil
	}
	res["collection_method"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCollectionMethod(val)
		}
		return nil
	}
	res["contact_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetContactId(val)
		}
		return nil
	}
	res["currency"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCurrency(val)
		}
		return nil
	}
	res["errors"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetErrors(val)
		}
		return nil
	}
	res["gateway"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetGateway(val)
		}
		return nil
	}
	res["gateway_account_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetGatewayAccountName(val)
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
	res["order_ids"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetOrderIds(val)
		}
		return nil
	}
	res["orders"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateTransactionsItemWithTransactionGetResponse_ordersFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]TransactionsItemWithTransactionGetResponse_ordersable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(TransactionsItemWithTransactionGetResponse_ordersable)
				}
			}
			m.SetOrders(res)
		}
		return nil
	}
	res["paymentDate"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPaymentDate(val)
		}
		return nil
	}
	res["status"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetStatus(val)
		}
		return nil
	}
	res["test"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTest(val)
		}
		return nil
	}
	res["transaction_date"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTransactionDate(val)
		}
		return nil
	}
	res["type"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTypeEscaped(val)
		}
		return nil
	}
	return res
}

// GetGateway gets the gateway property value. The gateway property
// returns a *string when successful
func (m *TransactionsItemWithTransactionGetResponse) GetGateway() *string {
	return m.gateway
}

// GetGatewayAccountName gets the gateway_account_name property value. The gateway_account_name property
// returns a *string when successful
func (m *TransactionsItemWithTransactionGetResponse) GetGatewayAccountName() *string {
	return m.gateway_account_name
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *TransactionsItemWithTransactionGetResponse) GetId() *int64 {
	return m.id
}

// GetOrderIds gets the order_ids property value. The order_ids property
// returns a *string when successful
func (m *TransactionsItemWithTransactionGetResponse) GetOrderIds() *string {
	return m.order_ids
}

// GetOrders gets the orders property value. The orders property
// returns a []TransactionsItemWithTransactionGetResponse_ordersable when successful
func (m *TransactionsItemWithTransactionGetResponse) GetOrders() []TransactionsItemWithTransactionGetResponse_ordersable {
	return m.orders
}

// GetPaymentDate gets the paymentDate property value. The paymentDate property
// returns a *string when successful
func (m *TransactionsItemWithTransactionGetResponse) GetPaymentDate() *string {
	return m.paymentDate
}

// GetStatus gets the status property value. The status property
// returns a *string when successful
func (m *TransactionsItemWithTransactionGetResponse) GetStatus() *string {
	return m.status
}

// GetTest gets the test property value. The test property
// returns a *bool when successful
func (m *TransactionsItemWithTransactionGetResponse) GetTest() *bool {
	return m.test
}

// GetTransactionDate gets the transaction_date property value. The transaction_date property
// returns a *string when successful
func (m *TransactionsItemWithTransactionGetResponse) GetTransactionDate() *string {
	return m.transaction_date
}

// GetTypeEscaped gets the type property value. The type property
// returns a *string when successful
func (m *TransactionsItemWithTransactionGetResponse) GetTypeEscaped() *string {
	return m.typeEscaped
}

// Serialize serializes information the current object
func (m *TransactionsItemWithTransactionGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteFloat64Value("amount", m.GetAmount())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("collection_method", m.GetCollectionMethod())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("contact_id", m.GetContactId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("currency", m.GetCurrency())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("errors", m.GetErrors())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("gateway", m.GetGateway())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("gateway_account_name", m.GetGatewayAccountName())
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
	if m.GetOrders() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOrders()))
		for i, v := range m.GetOrders() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("orders", cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("order_ids", m.GetOrderIds())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("paymentDate", m.GetPaymentDate())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("status", m.GetStatus())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("test", m.GetTest())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("transaction_date", m.GetTransactionDate())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("type", m.GetTypeEscaped())
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
func (m *TransactionsItemWithTransactionGetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAmount sets the amount property value. The amount property
func (m *TransactionsItemWithTransactionGetResponse) SetAmount(value *float64) {
	m.amount = value
}

// SetCollectionMethod sets the collection_method property value. The collection_method property
func (m *TransactionsItemWithTransactionGetResponse) SetCollectionMethod(value *string) {
	m.collection_method = value
}

// SetContactId sets the contact_id property value. The contact_id property
func (m *TransactionsItemWithTransactionGetResponse) SetContactId(value *int64) {
	m.contact_id = value
}

// SetCurrency sets the currency property value. The currency property
func (m *TransactionsItemWithTransactionGetResponse) SetCurrency(value *string) {
	m.currency = value
}

// SetErrors sets the errors property value. The errors property
func (m *TransactionsItemWithTransactionGetResponse) SetErrors(value *string) {
	m.errors = value
}

// SetGateway sets the gateway property value. The gateway property
func (m *TransactionsItemWithTransactionGetResponse) SetGateway(value *string) {
	m.gateway = value
}

// SetGatewayAccountName sets the gateway_account_name property value. The gateway_account_name property
func (m *TransactionsItemWithTransactionGetResponse) SetGatewayAccountName(value *string) {
	m.gateway_account_name = value
}

// SetId sets the id property value. The id property
func (m *TransactionsItemWithTransactionGetResponse) SetId(value *int64) {
	m.id = value
}

// SetOrderIds sets the order_ids property value. The order_ids property
func (m *TransactionsItemWithTransactionGetResponse) SetOrderIds(value *string) {
	m.order_ids = value
}

// SetOrders sets the orders property value. The orders property
func (m *TransactionsItemWithTransactionGetResponse) SetOrders(value []TransactionsItemWithTransactionGetResponse_ordersable) {
	m.orders = value
}

// SetPaymentDate sets the paymentDate property value. The paymentDate property
func (m *TransactionsItemWithTransactionGetResponse) SetPaymentDate(value *string) {
	m.paymentDate = value
}

// SetStatus sets the status property value. The status property
func (m *TransactionsItemWithTransactionGetResponse) SetStatus(value *string) {
	m.status = value
}

// SetTest sets the test property value. The test property
func (m *TransactionsItemWithTransactionGetResponse) SetTest(value *bool) {
	m.test = value
}

// SetTransactionDate sets the transaction_date property value. The transaction_date property
func (m *TransactionsItemWithTransactionGetResponse) SetTransactionDate(value *string) {
	m.transaction_date = value
}

// SetTypeEscaped sets the type property value. The type property
func (m *TransactionsItemWithTransactionGetResponse) SetTypeEscaped(value *string) {
	m.typeEscaped = value
}

type TransactionsItemWithTransactionGetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAmount() *float64
	GetCollectionMethod() *string
	GetContactId() *int64
	GetCurrency() *string
	GetErrors() *string
	GetGateway() *string
	GetGatewayAccountName() *string
	GetId() *int64
	GetOrderIds() *string
	GetOrders() []TransactionsItemWithTransactionGetResponse_ordersable
	GetPaymentDate() *string
	GetStatus() *string
	GetTest() *bool
	GetTransactionDate() *string
	GetTypeEscaped() *string
	SetAmount(value *float64)
	SetCollectionMethod(value *string)
	SetContactId(value *int64)
	SetCurrency(value *string)
	SetErrors(value *string)
	SetGateway(value *string)
	SetGatewayAccountName(value *string)
	SetId(value *int64)
	SetOrderIds(value *string)
	SetOrders(value []TransactionsItemWithTransactionGetResponse_ordersable)
	SetPaymentDate(value *string)
	SetStatus(value *string)
	SetTest(value *bool)
	SetTransactionDate(value *string)
	SetTypeEscaped(value *string)
}
