package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OrdersItemPaymentPlanPutResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The auto_charge property
	auto_charge *bool
	// The credit_card_id property
	credit_card_id *int64
	// The days_between_payments property
	days_between_payments *int32
	// The initial_payment_amount property
	initial_payment_amount *float64
	// The initial_payment_date property
	initial_payment_date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
	// The number_of_payments property
	number_of_payments *int32
	// The payment_gateway property
	payment_gateway OrdersItemPaymentPlanPutResponse_payment_gatewayable
	// The plan_start_date property
	plan_start_date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
}

// NewOrdersItemPaymentPlanPutResponse instantiates a new OrdersItemPaymentPlanPutResponse and sets the default values.
func NewOrdersItemPaymentPlanPutResponse() *OrdersItemPaymentPlanPutResponse {
	m := &OrdersItemPaymentPlanPutResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateOrdersItemPaymentPlanPutResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrdersItemPaymentPlanPutResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewOrdersItemPaymentPlanPutResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OrdersItemPaymentPlanPutResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAutoCharge gets the auto_charge property value. The auto_charge property
// returns a *bool when successful
func (m *OrdersItemPaymentPlanPutResponse) GetAutoCharge() *bool {
	return m.auto_charge
}

// GetCreditCardId gets the credit_card_id property value. The credit_card_id property
// returns a *int64 when successful
func (m *OrdersItemPaymentPlanPutResponse) GetCreditCardId() *int64 {
	return m.credit_card_id
}

// GetDaysBetweenPayments gets the days_between_payments property value. The days_between_payments property
// returns a *int32 when successful
func (m *OrdersItemPaymentPlanPutResponse) GetDaysBetweenPayments() *int32 {
	return m.days_between_payments
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrdersItemPaymentPlanPutResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["auto_charge"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAutoCharge(val)
		}
		return nil
	}
	res["credit_card_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCreditCardId(val)
		}
		return nil
	}
	res["days_between_payments"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDaysBetweenPayments(val)
		}
		return nil
	}
	res["initial_payment_amount"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetFloat64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetInitialPaymentAmount(val)
		}
		return nil
	}
	res["initial_payment_date"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetDateOnlyValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetInitialPaymentDate(val)
		}
		return nil
	}
	res["number_of_payments"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNumberOfPayments(val)
		}
		return nil
	}
	res["payment_gateway"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateOrdersItemPaymentPlanPutResponse_payment_gatewayFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPaymentGateway(val.(OrdersItemPaymentPlanPutResponse_payment_gatewayable))
		}
		return nil
	}
	res["plan_start_date"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetDateOnlyValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPlanStartDate(val)
		}
		return nil
	}
	return res
}

// GetInitialPaymentAmount gets the initial_payment_amount property value. The initial_payment_amount property
// returns a *float64 when successful
func (m *OrdersItemPaymentPlanPutResponse) GetInitialPaymentAmount() *float64 {
	return m.initial_payment_amount
}

// GetInitialPaymentDate gets the initial_payment_date property value. The initial_payment_date property
// returns a *DateOnly when successful
func (m *OrdersItemPaymentPlanPutResponse) GetInitialPaymentDate() *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly {
	return m.initial_payment_date
}

// GetNumberOfPayments gets the number_of_payments property value. The number_of_payments property
// returns a *int32 when successful
func (m *OrdersItemPaymentPlanPutResponse) GetNumberOfPayments() *int32 {
	return m.number_of_payments
}

// GetPaymentGateway gets the payment_gateway property value. The payment_gateway property
// returns a OrdersItemPaymentPlanPutResponse_payment_gatewayable when successful
func (m *OrdersItemPaymentPlanPutResponse) GetPaymentGateway() OrdersItemPaymentPlanPutResponse_payment_gatewayable {
	return m.payment_gateway
}

// GetPlanStartDate gets the plan_start_date property value. The plan_start_date property
// returns a *DateOnly when successful
func (m *OrdersItemPaymentPlanPutResponse) GetPlanStartDate() *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly {
	return m.plan_start_date
}

// Serialize serializes information the current object
func (m *OrdersItemPaymentPlanPutResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteBoolValue("auto_charge", m.GetAutoCharge())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("credit_card_id", m.GetCreditCardId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("days_between_payments", m.GetDaysBetweenPayments())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteFloat64Value("initial_payment_amount", m.GetInitialPaymentAmount())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteDateOnlyValue("initial_payment_date", m.GetInitialPaymentDate())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("number_of_payments", m.GetNumberOfPayments())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("payment_gateway", m.GetPaymentGateway())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteDateOnlyValue("plan_start_date", m.GetPlanStartDate())
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
func (m *OrdersItemPaymentPlanPutResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAutoCharge sets the auto_charge property value. The auto_charge property
func (m *OrdersItemPaymentPlanPutResponse) SetAutoCharge(value *bool) {
	m.auto_charge = value
}

// SetCreditCardId sets the credit_card_id property value. The credit_card_id property
func (m *OrdersItemPaymentPlanPutResponse) SetCreditCardId(value *int64) {
	m.credit_card_id = value
}

// SetDaysBetweenPayments sets the days_between_payments property value. The days_between_payments property
func (m *OrdersItemPaymentPlanPutResponse) SetDaysBetweenPayments(value *int32) {
	m.days_between_payments = value
}

// SetInitialPaymentAmount sets the initial_payment_amount property value. The initial_payment_amount property
func (m *OrdersItemPaymentPlanPutResponse) SetInitialPaymentAmount(value *float64) {
	m.initial_payment_amount = value
}

// SetInitialPaymentDate sets the initial_payment_date property value. The initial_payment_date property
func (m *OrdersItemPaymentPlanPutResponse) SetInitialPaymentDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
	m.initial_payment_date = value
}

// SetNumberOfPayments sets the number_of_payments property value. The number_of_payments property
func (m *OrdersItemPaymentPlanPutResponse) SetNumberOfPayments(value *int32) {
	m.number_of_payments = value
}

// SetPaymentGateway sets the payment_gateway property value. The payment_gateway property
func (m *OrdersItemPaymentPlanPutResponse) SetPaymentGateway(value OrdersItemPaymentPlanPutResponse_payment_gatewayable) {
	m.payment_gateway = value
}

// SetPlanStartDate sets the plan_start_date property value. The plan_start_date property
func (m *OrdersItemPaymentPlanPutResponse) SetPlanStartDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
	m.plan_start_date = value
}

type OrdersItemPaymentPlanPutResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAutoCharge() *bool
	GetCreditCardId() *int64
	GetDaysBetweenPayments() *int32
	GetInitialPaymentAmount() *float64
	GetInitialPaymentDate() *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
	GetNumberOfPayments() *int32
	GetPaymentGateway() OrdersItemPaymentPlanPutResponse_payment_gatewayable
	GetPlanStartDate() *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
	SetAutoCharge(value *bool)
	SetCreditCardId(value *int64)
	SetDaysBetweenPayments(value *int32)
	SetInitialPaymentAmount(value *float64)
	SetInitialPaymentDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
	SetNumberOfPayments(value *int32)
	SetPaymentGateway(value OrdersItemPaymentPlanPutResponse_payment_gatewayable)
	SetPlanStartDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
}
