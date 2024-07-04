package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SubscriptionsGetResponse_subscriptions struct {
	// The active property
	active *bool
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The allow_tax property
	allow_tax *bool
	// The auto_charge property
	auto_charge *bool
	// The billing_amount property
	billing_amount *float64
	// The billing_frequency property
	billing_frequency *int32
	// The contact_id property
	contact_id *int64
	// The credit_card_id property
	credit_card_id *int64
	// The end_date property
	end_date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
	// The id property
	id *int64
	// The next_bill_date property
	next_bill_date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
	// The payment_gateway_id property
	payment_gateway_id *int64
	// The product_id property
	product_id *int64
	// The quantity property
	quantity *int64
	// The sale_affiliate_id property
	sale_affiliate_id *int64
	// The start_date property
	start_date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
	// The subscription_plan_id property
	subscription_plan_id *int64
	// The use_default_payment_gateway property
	use_default_payment_gateway *bool
}

// NewSubscriptionsGetResponse_subscriptions instantiates a new SubscriptionsGetResponse_subscriptions and sets the default values.
func NewSubscriptionsGetResponse_subscriptions() *SubscriptionsGetResponse_subscriptions {
	m := &SubscriptionsGetResponse_subscriptions{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateSubscriptionsGetResponse_subscriptionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSubscriptionsGetResponse_subscriptionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSubscriptionsGetResponse_subscriptions(), nil
}

// GetActive gets the active property value. The active property
// returns a *bool when successful
func (m *SubscriptionsGetResponse_subscriptions) GetActive() *bool {
	return m.active
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SubscriptionsGetResponse_subscriptions) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAllowTax gets the allow_tax property value. The allow_tax property
// returns a *bool when successful
func (m *SubscriptionsGetResponse_subscriptions) GetAllowTax() *bool {
	return m.allow_tax
}

// GetAutoCharge gets the auto_charge property value. The auto_charge property
// returns a *bool when successful
func (m *SubscriptionsGetResponse_subscriptions) GetAutoCharge() *bool {
	return m.auto_charge
}

// GetBillingAmount gets the billing_amount property value. The billing_amount property
// returns a *float64 when successful
func (m *SubscriptionsGetResponse_subscriptions) GetBillingAmount() *float64 {
	return m.billing_amount
}

// GetBillingFrequency gets the billing_frequency property value. The billing_frequency property
// returns a *int32 when successful
func (m *SubscriptionsGetResponse_subscriptions) GetBillingFrequency() *int32 {
	return m.billing_frequency
}

// GetContactId gets the contact_id property value. The contact_id property
// returns a *int64 when successful
func (m *SubscriptionsGetResponse_subscriptions) GetContactId() *int64 {
	return m.contact_id
}

// GetCreditCardId gets the credit_card_id property value. The credit_card_id property
// returns a *int64 when successful
func (m *SubscriptionsGetResponse_subscriptions) GetCreditCardId() *int64 {
	return m.credit_card_id
}

// GetEndDate gets the end_date property value. The end_date property
// returns a *DateOnly when successful
func (m *SubscriptionsGetResponse_subscriptions) GetEndDate() *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly {
	return m.end_date
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SubscriptionsGetResponse_subscriptions) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["active"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetActive(val)
		}
		return nil
	}
	res["allow_tax"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAllowTax(val)
		}
		return nil
	}
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
	res["billing_amount"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetFloat64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetBillingAmount(val)
		}
		return nil
	}
	res["billing_frequency"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetBillingFrequency(val)
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
	res["end_date"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetDateOnlyValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetEndDate(val)
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
	res["next_bill_date"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetDateOnlyValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNextBillDate(val)
		}
		return nil
	}
	res["payment_gateway_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPaymentGatewayId(val)
		}
		return nil
	}
	res["product_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetProductId(val)
		}
		return nil
	}
	res["quantity"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetQuantity(val)
		}
		return nil
	}
	res["sale_affiliate_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSaleAffiliateId(val)
		}
		return nil
	}
	res["start_date"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetDateOnlyValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetStartDate(val)
		}
		return nil
	}
	res["subscription_plan_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSubscriptionPlanId(val)
		}
		return nil
	}
	res["use_default_payment_gateway"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetUseDefaultPaymentGateway(val)
		}
		return nil
	}
	return res
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *SubscriptionsGetResponse_subscriptions) GetId() *int64 {
	return m.id
}

// GetNextBillDate gets the next_bill_date property value. The next_bill_date property
// returns a *DateOnly when successful
func (m *SubscriptionsGetResponse_subscriptions) GetNextBillDate() *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly {
	return m.next_bill_date
}

// GetPaymentGatewayId gets the payment_gateway_id property value. The payment_gateway_id property
// returns a *int64 when successful
func (m *SubscriptionsGetResponse_subscriptions) GetPaymentGatewayId() *int64 {
	return m.payment_gateway_id
}

// GetProductId gets the product_id property value. The product_id property
// returns a *int64 when successful
func (m *SubscriptionsGetResponse_subscriptions) GetProductId() *int64 {
	return m.product_id
}

// GetQuantity gets the quantity property value. The quantity property
// returns a *int64 when successful
func (m *SubscriptionsGetResponse_subscriptions) GetQuantity() *int64 {
	return m.quantity
}

// GetSaleAffiliateId gets the sale_affiliate_id property value. The sale_affiliate_id property
// returns a *int64 when successful
func (m *SubscriptionsGetResponse_subscriptions) GetSaleAffiliateId() *int64 {
	return m.sale_affiliate_id
}

// GetStartDate gets the start_date property value. The start_date property
// returns a *DateOnly when successful
func (m *SubscriptionsGetResponse_subscriptions) GetStartDate() *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly {
	return m.start_date
}

// GetSubscriptionPlanId gets the subscription_plan_id property value. The subscription_plan_id property
// returns a *int64 when successful
func (m *SubscriptionsGetResponse_subscriptions) GetSubscriptionPlanId() *int64 {
	return m.subscription_plan_id
}

// GetUseDefaultPaymentGateway gets the use_default_payment_gateway property value. The use_default_payment_gateway property
// returns a *bool when successful
func (m *SubscriptionsGetResponse_subscriptions) GetUseDefaultPaymentGateway() *bool {
	return m.use_default_payment_gateway
}

// Serialize serializes information the current object
func (m *SubscriptionsGetResponse_subscriptions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteBoolValue("active", m.GetActive())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("allow_tax", m.GetAllowTax())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("auto_charge", m.GetAutoCharge())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteFloat64Value("billing_amount", m.GetBillingAmount())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("billing_frequency", m.GetBillingFrequency())
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
		err := writer.WriteInt64Value("credit_card_id", m.GetCreditCardId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteDateOnlyValue("end_date", m.GetEndDate())
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
		err := writer.WriteDateOnlyValue("next_bill_date", m.GetNextBillDate())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("payment_gateway_id", m.GetPaymentGatewayId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("product_id", m.GetProductId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("quantity", m.GetQuantity())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("sale_affiliate_id", m.GetSaleAffiliateId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteDateOnlyValue("start_date", m.GetStartDate())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("subscription_plan_id", m.GetSubscriptionPlanId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("use_default_payment_gateway", m.GetUseDefaultPaymentGateway())
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

// SetActive sets the active property value. The active property
func (m *SubscriptionsGetResponse_subscriptions) SetActive(value *bool) {
	m.active = value
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SubscriptionsGetResponse_subscriptions) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAllowTax sets the allow_tax property value. The allow_tax property
func (m *SubscriptionsGetResponse_subscriptions) SetAllowTax(value *bool) {
	m.allow_tax = value
}

// SetAutoCharge sets the auto_charge property value. The auto_charge property
func (m *SubscriptionsGetResponse_subscriptions) SetAutoCharge(value *bool) {
	m.auto_charge = value
}

// SetBillingAmount sets the billing_amount property value. The billing_amount property
func (m *SubscriptionsGetResponse_subscriptions) SetBillingAmount(value *float64) {
	m.billing_amount = value
}

// SetBillingFrequency sets the billing_frequency property value. The billing_frequency property
func (m *SubscriptionsGetResponse_subscriptions) SetBillingFrequency(value *int32) {
	m.billing_frequency = value
}

// SetContactId sets the contact_id property value. The contact_id property
func (m *SubscriptionsGetResponse_subscriptions) SetContactId(value *int64) {
	m.contact_id = value
}

// SetCreditCardId sets the credit_card_id property value. The credit_card_id property
func (m *SubscriptionsGetResponse_subscriptions) SetCreditCardId(value *int64) {
	m.credit_card_id = value
}

// SetEndDate sets the end_date property value. The end_date property
func (m *SubscriptionsGetResponse_subscriptions) SetEndDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
	m.end_date = value
}

// SetId sets the id property value. The id property
func (m *SubscriptionsGetResponse_subscriptions) SetId(value *int64) {
	m.id = value
}

// SetNextBillDate sets the next_bill_date property value. The next_bill_date property
func (m *SubscriptionsGetResponse_subscriptions) SetNextBillDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
	m.next_bill_date = value
}

// SetPaymentGatewayId sets the payment_gateway_id property value. The payment_gateway_id property
func (m *SubscriptionsGetResponse_subscriptions) SetPaymentGatewayId(value *int64) {
	m.payment_gateway_id = value
}

// SetProductId sets the product_id property value. The product_id property
func (m *SubscriptionsGetResponse_subscriptions) SetProductId(value *int64) {
	m.product_id = value
}

// SetQuantity sets the quantity property value. The quantity property
func (m *SubscriptionsGetResponse_subscriptions) SetQuantity(value *int64) {
	m.quantity = value
}

// SetSaleAffiliateId sets the sale_affiliate_id property value. The sale_affiliate_id property
func (m *SubscriptionsGetResponse_subscriptions) SetSaleAffiliateId(value *int64) {
	m.sale_affiliate_id = value
}

// SetStartDate sets the start_date property value. The start_date property
func (m *SubscriptionsGetResponse_subscriptions) SetStartDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
	m.start_date = value
}

// SetSubscriptionPlanId sets the subscription_plan_id property value. The subscription_plan_id property
func (m *SubscriptionsGetResponse_subscriptions) SetSubscriptionPlanId(value *int64) {
	m.subscription_plan_id = value
}

// SetUseDefaultPaymentGateway sets the use_default_payment_gateway property value. The use_default_payment_gateway property
func (m *SubscriptionsGetResponse_subscriptions) SetUseDefaultPaymentGateway(value *bool) {
	m.use_default_payment_gateway = value
}

type SubscriptionsGetResponse_subscriptionsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetActive() *bool
	GetAllowTax() *bool
	GetAutoCharge() *bool
	GetBillingAmount() *float64
	GetBillingFrequency() *int32
	GetContactId() *int64
	GetCreditCardId() *int64
	GetEndDate() *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
	GetId() *int64
	GetNextBillDate() *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
	GetPaymentGatewayId() *int64
	GetProductId() *int64
	GetQuantity() *int64
	GetSaleAffiliateId() *int64
	GetStartDate() *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
	GetSubscriptionPlanId() *int64
	GetUseDefaultPaymentGateway() *bool
	SetActive(value *bool)
	SetAllowTax(value *bool)
	SetAutoCharge(value *bool)
	SetBillingAmount(value *float64)
	SetBillingFrequency(value *int32)
	SetContactId(value *int64)
	SetCreditCardId(value *int64)
	SetEndDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
	SetId(value *int64)
	SetNextBillDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
	SetPaymentGatewayId(value *int64)
	SetProductId(value *int64)
	SetQuantity(value *int64)
	SetSaleAffiliateId(value *int64)
	SetStartDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
	SetSubscriptionPlanId(value *int64)
	SetUseDefaultPaymentGateway(value *bool)
}
