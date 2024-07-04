package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
)

type OrdersItemPayments struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The amount property
	amount *float64
	// The id property
	id *int64
	// The invoice_id property
	invoice_id *int64
	// The last_updated property
	last_updated *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	// The note property
	note *string
	// The pay_date property
	pay_date *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	// The pay_status property
	pay_status *string
	// The payment_id property
	payment_id *int64
	// The refund_invoice_payment_id property
	refund_invoice_payment_id *int64
	// The skip_commission property
	skip_commission *bool
}

// NewOrdersItemPayments instantiates a new OrdersItemPayments and sets the default values.
func NewOrdersItemPayments() *OrdersItemPayments {
	m := &OrdersItemPayments{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateOrdersItemPaymentsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrdersItemPaymentsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewOrdersItemPayments(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OrdersItemPayments) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAmount gets the amount property value. The amount property
// returns a *float64 when successful
func (m *OrdersItemPayments) GetAmount() *float64 {
	return m.amount
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrdersItemPayments) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
	res["last_updated"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetTimeValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLastUpdated(val)
		}
		return nil
	}
	res["note"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNote(val)
		}
		return nil
	}
	res["pay_date"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetTimeValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPayDate(val)
		}
		return nil
	}
	res["pay_status"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPayStatus(val)
		}
		return nil
	}
	res["payment_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPaymentId(val)
		}
		return nil
	}
	res["refund_invoice_payment_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRefundInvoicePaymentId(val)
		}
		return nil
	}
	res["skip_commission"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSkipCommission(val)
		}
		return nil
	}
	return res
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *OrdersItemPayments) GetId() *int64 {
	return m.id
}

// GetInvoiceId gets the invoice_id property value. The invoice_id property
// returns a *int64 when successful
func (m *OrdersItemPayments) GetInvoiceId() *int64 {
	return m.invoice_id
}

// GetLastUpdated gets the last_updated property value. The last_updated property
// returns a *Time when successful
func (m *OrdersItemPayments) GetLastUpdated() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.last_updated
}

// GetNote gets the note property value. The note property
// returns a *string when successful
func (m *OrdersItemPayments) GetNote() *string {
	return m.note
}

// GetPayDate gets the pay_date property value. The pay_date property
// returns a *Time when successful
func (m *OrdersItemPayments) GetPayDate() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.pay_date
}

// GetPaymentId gets the payment_id property value. The payment_id property
// returns a *int64 when successful
func (m *OrdersItemPayments) GetPaymentId() *int64 {
	return m.payment_id
}

// GetPayStatus gets the pay_status property value. The pay_status property
// returns a *string when successful
func (m *OrdersItemPayments) GetPayStatus() *string {
	return m.pay_status
}

// GetRefundInvoicePaymentId gets the refund_invoice_payment_id property value. The refund_invoice_payment_id property
// returns a *int64 when successful
func (m *OrdersItemPayments) GetRefundInvoicePaymentId() *int64 {
	return m.refund_invoice_payment_id
}

// GetSkipCommission gets the skip_commission property value. The skip_commission property
// returns a *bool when successful
func (m *OrdersItemPayments) GetSkipCommission() *bool {
	return m.skip_commission
}

// Serialize serializes information the current object
func (m *OrdersItemPayments) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteFloat64Value("amount", m.GetAmount())
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
		err := writer.WriteInt64Value("invoice_id", m.GetInvoiceId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteTimeValue("last_updated", m.GetLastUpdated())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("note", m.GetNote())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("payment_id", m.GetPaymentId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteTimeValue("pay_date", m.GetPayDate())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("pay_status", m.GetPayStatus())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("refund_invoice_payment_id", m.GetRefundInvoicePaymentId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("skip_commission", m.GetSkipCommission())
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
func (m *OrdersItemPayments) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAmount sets the amount property value. The amount property
func (m *OrdersItemPayments) SetAmount(value *float64) {
	m.amount = value
}

// SetId sets the id property value. The id property
func (m *OrdersItemPayments) SetId(value *int64) {
	m.id = value
}

// SetInvoiceId sets the invoice_id property value. The invoice_id property
func (m *OrdersItemPayments) SetInvoiceId(value *int64) {
	m.invoice_id = value
}

// SetLastUpdated sets the last_updated property value. The last_updated property
func (m *OrdersItemPayments) SetLastUpdated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	m.last_updated = value
}

// SetNote sets the note property value. The note property
func (m *OrdersItemPayments) SetNote(value *string) {
	m.note = value
}

// SetPayDate sets the pay_date property value. The pay_date property
func (m *OrdersItemPayments) SetPayDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	m.pay_date = value
}

// SetPaymentId sets the payment_id property value. The payment_id property
func (m *OrdersItemPayments) SetPaymentId(value *int64) {
	m.payment_id = value
}

// SetPayStatus sets the pay_status property value. The pay_status property
func (m *OrdersItemPayments) SetPayStatus(value *string) {
	m.pay_status = value
}

// SetRefundInvoicePaymentId sets the refund_invoice_payment_id property value. The refund_invoice_payment_id property
func (m *OrdersItemPayments) SetRefundInvoicePaymentId(value *int64) {
	m.refund_invoice_payment_id = value
}

// SetSkipCommission sets the skip_commission property value. The skip_commission property
func (m *OrdersItemPayments) SetSkipCommission(value *bool) {
	m.skip_commission = value
}

type OrdersItemPaymentsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAmount() *float64
	GetId() *int64
	GetInvoiceId() *int64
	GetLastUpdated() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetNote() *string
	GetPayDate() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetPaymentId() *int64
	GetPayStatus() *string
	GetRefundInvoicePaymentId() *int64
	GetSkipCommission() *bool
	SetAmount(value *float64)
	SetId(value *int64)
	SetInvoiceId(value *int64)
	SetLastUpdated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	SetNote(value *string)
	SetPayDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	SetPaymentId(value *int64)
	SetPayStatus(value *string)
	SetRefundInvoicePaymentId(value *int64)
	SetSkipCommission(value *bool)
}
