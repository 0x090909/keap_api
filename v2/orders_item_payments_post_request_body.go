package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
)

type OrdersItemPaymentsPostRequestBody struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The apply_to_commissions property
	apply_to_commissions *bool
	// The date property
	date *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	// The notes property
	notes *string
	// The payment_amount property
	payment_amount *string
	// The payment_method_id property
	payment_method_id *int64
}

// NewOrdersItemPaymentsPostRequestBody instantiates a new OrdersItemPaymentsPostRequestBody and sets the default values.
func NewOrdersItemPaymentsPostRequestBody() *OrdersItemPaymentsPostRequestBody {
	m := &OrdersItemPaymentsPostRequestBody{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateOrdersItemPaymentsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrdersItemPaymentsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewOrdersItemPaymentsPostRequestBody(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OrdersItemPaymentsPostRequestBody) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetApplyToCommissions gets the apply_to_commissions property value. The apply_to_commissions property
// returns a *bool when successful
func (m *OrdersItemPaymentsPostRequestBody) GetApplyToCommissions() *bool {
	return m.apply_to_commissions
}

// GetDate gets the date property value. The date property
// returns a *Time when successful
func (m *OrdersItemPaymentsPostRequestBody) GetDate() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.date
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrdersItemPaymentsPostRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["apply_to_commissions"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetApplyToCommissions(val)
		}
		return nil
	}
	res["date"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetTimeValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDate(val)
		}
		return nil
	}
	res["notes"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNotes(val)
		}
		return nil
	}
	res["payment_amount"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPaymentAmount(val)
		}
		return nil
	}
	res["payment_method_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPaymentMethodId(val)
		}
		return nil
	}
	return res
}

// GetNotes gets the notes property value. The notes property
// returns a *string when successful
func (m *OrdersItemPaymentsPostRequestBody) GetNotes() *string {
	return m.notes
}

// GetPaymentAmount gets the payment_amount property value. The payment_amount property
// returns a *string when successful
func (m *OrdersItemPaymentsPostRequestBody) GetPaymentAmount() *string {
	return m.payment_amount
}

// GetPaymentMethodId gets the payment_method_id property value. The payment_method_id property
// returns a *int64 when successful
func (m *OrdersItemPaymentsPostRequestBody) GetPaymentMethodId() *int64 {
	return m.payment_method_id
}

// Serialize serializes information the current object
func (m *OrdersItemPaymentsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteBoolValue("apply_to_commissions", m.GetApplyToCommissions())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteTimeValue("date", m.GetDate())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("notes", m.GetNotes())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("payment_amount", m.GetPaymentAmount())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("payment_method_id", m.GetPaymentMethodId())
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
func (m *OrdersItemPaymentsPostRequestBody) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetApplyToCommissions sets the apply_to_commissions property value. The apply_to_commissions property
func (m *OrdersItemPaymentsPostRequestBody) SetApplyToCommissions(value *bool) {
	m.apply_to_commissions = value
}

// SetDate sets the date property value. The date property
func (m *OrdersItemPaymentsPostRequestBody) SetDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	m.date = value
}

// SetNotes sets the notes property value. The notes property
func (m *OrdersItemPaymentsPostRequestBody) SetNotes(value *string) {
	m.notes = value
}

// SetPaymentAmount sets the payment_amount property value. The payment_amount property
func (m *OrdersItemPaymentsPostRequestBody) SetPaymentAmount(value *string) {
	m.payment_amount = value
}

// SetPaymentMethodId sets the payment_method_id property value. The payment_method_id property
func (m *OrdersItemPaymentsPostRequestBody) SetPaymentMethodId(value *int64) {
	m.payment_method_id = value
}

type OrdersItemPaymentsPostRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetApplyToCommissions() *bool
	GetDate() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetNotes() *string
	GetPaymentAmount() *string
	GetPaymentMethodId() *int64
	SetApplyToCommissions(value *bool)
	SetDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	SetNotes(value *string)
	SetPaymentAmount(value *string)
	SetPaymentMethodId(value *int64)
}
