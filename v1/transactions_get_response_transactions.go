package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TransactionsGetResponse_transactions struct {
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
	orders []TransactionsGetResponse_transactions_ordersable
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

// NewTransactionsGetResponse_transactions instantiates a new TransactionsGetResponse_transactions and sets the default values.
func NewTransactionsGetResponse_transactions() *TransactionsGetResponse_transactions {
	m := &TransactionsGetResponse_transactions{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateTransactionsGetResponse_transactionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionsGetResponse_transactionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTransactionsGetResponse_transactions(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransactionsGetResponse_transactions) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAmount gets the amount property value. The amount property
// returns a *float64 when successful
func (m *TransactionsGetResponse_transactions) GetAmount() *float64 {
	return m.amount
}

// GetCollectionMethod gets the collection_method property value. The collection_method property
// returns a *string when successful
func (m *TransactionsGetResponse_transactions) GetCollectionMethod() *string {
	return m.collection_method
}

// GetContactId gets the contact_id property value. The contact_id property
// returns a *int64 when successful
func (m *TransactionsGetResponse_transactions) GetContactId() *int64 {
	return m.contact_id
}

// GetCurrency gets the currency property value. The currency property
// returns a *string when successful
func (m *TransactionsGetResponse_transactions) GetCurrency() *string {
	return m.currency
}

// GetErrors gets the errors property value. The errors property
// returns a *string when successful
func (m *TransactionsGetResponse_transactions) GetErrors() *string {
	return m.errors
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransactionsGetResponse_transactions) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
		val, err := n.GetCollectionOfObjectValues(CreateTransactionsGetResponse_transactions_ordersFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]TransactionsGetResponse_transactions_ordersable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(TransactionsGetResponse_transactions_ordersable)
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
func (m *TransactionsGetResponse_transactions) GetGateway() *string {
	return m.gateway
}

// GetGatewayAccountName gets the gateway_account_name property value. The gateway_account_name property
// returns a *string when successful
func (m *TransactionsGetResponse_transactions) GetGatewayAccountName() *string {
	return m.gateway_account_name
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *TransactionsGetResponse_transactions) GetId() *int64 {
	return m.id
}

// GetOrderIds gets the order_ids property value. The order_ids property
// returns a *string when successful
func (m *TransactionsGetResponse_transactions) GetOrderIds() *string {
	return m.order_ids
}

// GetOrders gets the orders property value. The orders property
// returns a []TransactionsGetResponse_transactions_ordersable when successful
func (m *TransactionsGetResponse_transactions) GetOrders() []TransactionsGetResponse_transactions_ordersable {
	return m.orders
}

// GetPaymentDate gets the paymentDate property value. The paymentDate property
// returns a *string when successful
func (m *TransactionsGetResponse_transactions) GetPaymentDate() *string {
	return m.paymentDate
}

// GetStatus gets the status property value. The status property
// returns a *string when successful
func (m *TransactionsGetResponse_transactions) GetStatus() *string {
	return m.status
}

// GetTest gets the test property value. The test property
// returns a *bool when successful
func (m *TransactionsGetResponse_transactions) GetTest() *bool {
	return m.test
}

// GetTransactionDate gets the transaction_date property value. The transaction_date property
// returns a *string when successful
func (m *TransactionsGetResponse_transactions) GetTransactionDate() *string {
	return m.transaction_date
}

// GetTypeEscaped gets the type property value. The type property
// returns a *string when successful
func (m *TransactionsGetResponse_transactions) GetTypeEscaped() *string {
	return m.typeEscaped
}

// Serialize serializes information the current object
func (m *TransactionsGetResponse_transactions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *TransactionsGetResponse_transactions) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAmount sets the amount property value. The amount property
func (m *TransactionsGetResponse_transactions) SetAmount(value *float64) {
	m.amount = value
}

// SetCollectionMethod sets the collection_method property value. The collection_method property
func (m *TransactionsGetResponse_transactions) SetCollectionMethod(value *string) {
	m.collection_method = value
}

// SetContactId sets the contact_id property value. The contact_id property
func (m *TransactionsGetResponse_transactions) SetContactId(value *int64) {
	m.contact_id = value
}

// SetCurrency sets the currency property value. The currency property
func (m *TransactionsGetResponse_transactions) SetCurrency(value *string) {
	m.currency = value
}

// SetErrors sets the errors property value. The errors property
func (m *TransactionsGetResponse_transactions) SetErrors(value *string) {
	m.errors = value
}

// SetGateway sets the gateway property value. The gateway property
func (m *TransactionsGetResponse_transactions) SetGateway(value *string) {
	m.gateway = value
}

// SetGatewayAccountName sets the gateway_account_name property value. The gateway_account_name property
func (m *TransactionsGetResponse_transactions) SetGatewayAccountName(value *string) {
	m.gateway_account_name = value
}

// SetId sets the id property value. The id property
func (m *TransactionsGetResponse_transactions) SetId(value *int64) {
	m.id = value
}

// SetOrderIds sets the order_ids property value. The order_ids property
func (m *TransactionsGetResponse_transactions) SetOrderIds(value *string) {
	m.order_ids = value
}

// SetOrders sets the orders property value. The orders property
func (m *TransactionsGetResponse_transactions) SetOrders(value []TransactionsGetResponse_transactions_ordersable) {
	m.orders = value
}

// SetPaymentDate sets the paymentDate property value. The paymentDate property
func (m *TransactionsGetResponse_transactions) SetPaymentDate(value *string) {
	m.paymentDate = value
}

// SetStatus sets the status property value. The status property
func (m *TransactionsGetResponse_transactions) SetStatus(value *string) {
	m.status = value
}

// SetTest sets the test property value. The test property
func (m *TransactionsGetResponse_transactions) SetTest(value *bool) {
	m.test = value
}

// SetTransactionDate sets the transaction_date property value. The transaction_date property
func (m *TransactionsGetResponse_transactions) SetTransactionDate(value *string) {
	m.transaction_date = value
}

// SetTypeEscaped sets the type property value. The type property
func (m *TransactionsGetResponse_transactions) SetTypeEscaped(value *string) {
	m.typeEscaped = value
}

type TransactionsGetResponse_transactionsable interface {
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
	GetOrders() []TransactionsGetResponse_transactions_ordersable
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
	SetOrders(value []TransactionsGetResponse_transactions_ordersable)
	SetPaymentDate(value *string)
	SetStatus(value *string)
	SetTest(value *bool)
	SetTransactionDate(value *string)
	SetTypeEscaped(value *string)
}
