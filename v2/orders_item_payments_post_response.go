package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OrdersItemPaymentsPostResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The id property
	id *int64
	// The invoice_id property
	invoice_id *int64
	// The payment_amount property
	payment_amount *float64
	// The payment_status property
	payment_status *string
	// The transaction_id property
	transaction_id *int64
}

// NewOrdersItemPaymentsPostResponse instantiates a new OrdersItemPaymentsPostResponse and sets the default values.
func NewOrdersItemPaymentsPostResponse() *OrdersItemPaymentsPostResponse {
	m := &OrdersItemPaymentsPostResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateOrdersItemPaymentsPostResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrdersItemPaymentsPostResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewOrdersItemPaymentsPostResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OrdersItemPaymentsPostResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrdersItemPaymentsPostResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
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
	res["invoice_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetInvoiceId(val)
		}
		return nil
	}
	res["payment_amount"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetFloat64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPaymentAmount(val)
		}
		return nil
	}
	res["payment_status"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPaymentStatus(val)
		}
		return nil
	}
	res["transaction_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTransactionId(val)
		}
		return nil
	}
	return res
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *OrdersItemPaymentsPostResponse) GetId() *int64 {
	return m.id
}

// GetInvoiceId gets the invoice_id property value. The invoice_id property
// returns a *int64 when successful
func (m *OrdersItemPaymentsPostResponse) GetInvoiceId() *int64 {
	return m.invoice_id
}

// GetPaymentAmount gets the payment_amount property value. The payment_amount property
// returns a *float64 when successful
func (m *OrdersItemPaymentsPostResponse) GetPaymentAmount() *float64 {
	return m.payment_amount
}

// GetPaymentStatus gets the payment_status property value. The payment_status property
// returns a *string when successful
func (m *OrdersItemPaymentsPostResponse) GetPaymentStatus() *string {
	return m.payment_status
}

// GetTransactionId gets the transaction_id property value. The transaction_id property
// returns a *int64 when successful
func (m *OrdersItemPaymentsPostResponse) GetTransactionId() *int64 {
	return m.transaction_id
}

// Serialize serializes information the current object
func (m *OrdersItemPaymentsPostResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt64Value("id", m.GetId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("invoice_id", m.GetInvoiceId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteFloat64Value("payment_amount", m.GetPaymentAmount())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("payment_status", m.GetPaymentStatus())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("transaction_id", m.GetTransactionId())
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
func (m *OrdersItemPaymentsPostResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetId sets the id property value. The id property
func (m *OrdersItemPaymentsPostResponse) SetId(value *int64) {
	m.id = value
}

// SetInvoiceId sets the invoice_id property value. The invoice_id property
func (m *OrdersItemPaymentsPostResponse) SetInvoiceId(value *int64) {
	m.invoice_id = value
}

// SetPaymentAmount sets the payment_amount property value. The payment_amount property
func (m *OrdersItemPaymentsPostResponse) SetPaymentAmount(value *float64) {
	m.payment_amount = value
}

// SetPaymentStatus sets the payment_status property value. The payment_status property
func (m *OrdersItemPaymentsPostResponse) SetPaymentStatus(value *string) {
	m.payment_status = value
}

// SetTransactionId sets the transaction_id property value. The transaction_id property
func (m *OrdersItemPaymentsPostResponse) SetTransactionId(value *int64) {
	m.transaction_id = value
}

type OrdersItemPaymentsPostResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetId() *int64
	GetInvoiceId() *int64
	GetPaymentAmount() *float64
	GetPaymentStatus() *string
	GetTransactionId() *int64
	SetId(value *int64)
	SetInvoiceId(value *int64)
	SetPaymentAmount(value *float64)
	SetPaymentStatus(value *string)
	SetTransactionId(value *int64)
}
