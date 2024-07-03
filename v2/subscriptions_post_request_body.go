package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SubscriptionsPostRequestBody struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// If true, it will disable the check to see if there is already an identical subscription for the contact. Default is false.
	allow_duplicate *bool
	// Only works if the product the product subscription is for is taxable. Default is false.
	allow_tax *bool
	// Defaults to true.
	auto_charge *bool
	// Must be 0 or greater. Default is the price in the product subscription.
	billing_amount *float64
	// The contact_id property
	contact_id *int64
	// The first day the subscription will bill, in EST. Must not be in the past. Default is today.
	first_bill_date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
	// Default is the contact's most recently used card, if auto charge is true. Default is 0 otherwise.
	payment_method_id *string
	// Must be greater than 0. Default is 1.
	quantity *int32
	// The sale_affiliate_id property
	sale_affiliate_id *int64
	// Id of the product subscription.
	subscription_plan_id *int64
}

// NewSubscriptionsPostRequestBody instantiates a new SubscriptionsPostRequestBody and sets the default values.
func NewSubscriptionsPostRequestBody() *SubscriptionsPostRequestBody {
	m := &SubscriptionsPostRequestBody{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateSubscriptionsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSubscriptionsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSubscriptionsPostRequestBody(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SubscriptionsPostRequestBody) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAllowDuplicate gets the allow_duplicate property value. If true, it will disable the check to see if there is already an identical subscription for the contact. Default is false.
// returns a *bool when successful
func (m *SubscriptionsPostRequestBody) GetAllowDuplicate() *bool {
	return m.allow_duplicate
}

// GetAllowTax gets the allow_tax property value. Only works if the product the product subscription is for is taxable. Default is false.
// returns a *bool when successful
func (m *SubscriptionsPostRequestBody) GetAllowTax() *bool {
	return m.allow_tax
}

// GetAutoCharge gets the auto_charge property value. Defaults to true.
// returns a *bool when successful
func (m *SubscriptionsPostRequestBody) GetAutoCharge() *bool {
	return m.auto_charge
}

// GetBillingAmount gets the billing_amount property value. Must be 0 or greater. Default is the price in the product subscription.
// returns a *float64 when successful
func (m *SubscriptionsPostRequestBody) GetBillingAmount() *float64 {
	return m.billing_amount
}

// GetContactId gets the contact_id property value. The contact_id property
// returns a *int64 when successful
func (m *SubscriptionsPostRequestBody) GetContactId() *int64 {
	return m.contact_id
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SubscriptionsPostRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["allow_duplicate"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAllowDuplicate(val)
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
	res["first_bill_date"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetDateOnlyValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFirstBillDate(val)
		}
		return nil
	}
	res["payment_method_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPaymentMethodId(val)
		}
		return nil
	}
	res["quantity"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
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
	return res
}

// GetFirstBillDate gets the first_bill_date property value. The first day the subscription will bill, in EST. Must not be in the past. Default is today.
// returns a *DateOnly when successful
func (m *SubscriptionsPostRequestBody) GetFirstBillDate() *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly {
	return m.first_bill_date
}

// GetPaymentMethodId gets the payment_method_id property value. Default is the contact's most recently used card, if auto charge is true. Default is 0 otherwise.
// returns a *string when successful
func (m *SubscriptionsPostRequestBody) GetPaymentMethodId() *string {
	return m.payment_method_id
}

// GetQuantity gets the quantity property value. Must be greater than 0. Default is 1.
// returns a *int32 when successful
func (m *SubscriptionsPostRequestBody) GetQuantity() *int32 {
	return m.quantity
}

// GetSaleAffiliateId gets the sale_affiliate_id property value. The sale_affiliate_id property
// returns a *int64 when successful
func (m *SubscriptionsPostRequestBody) GetSaleAffiliateId() *int64 {
	return m.sale_affiliate_id
}

// GetSubscriptionPlanId gets the subscription_plan_id property value. Id of the product subscription.
// returns a *int64 when successful
func (m *SubscriptionsPostRequestBody) GetSubscriptionPlanId() *int64 {
	return m.subscription_plan_id
}

// Serialize serializes information the current object
func (m *SubscriptionsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteBoolValue("allow_duplicate", m.GetAllowDuplicate())
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
		err := writer.WriteInt64Value("contact_id", m.GetContactId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteDateOnlyValue("first_bill_date", m.GetFirstBillDate())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("payment_method_id", m.GetPaymentMethodId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("quantity", m.GetQuantity())
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
		err := writer.WriteInt64Value("subscription_plan_id", m.GetSubscriptionPlanId())
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
func (m *SubscriptionsPostRequestBody) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAllowDuplicate sets the allow_duplicate property value. If true, it will disable the check to see if there is already an identical subscription for the contact. Default is false.
func (m *SubscriptionsPostRequestBody) SetAllowDuplicate(value *bool) {
	m.allow_duplicate = value
}

// SetAllowTax sets the allow_tax property value. Only works if the product the product subscription is for is taxable. Default is false.
func (m *SubscriptionsPostRequestBody) SetAllowTax(value *bool) {
	m.allow_tax = value
}

// SetAutoCharge sets the auto_charge property value. Defaults to true.
func (m *SubscriptionsPostRequestBody) SetAutoCharge(value *bool) {
	m.auto_charge = value
}

// SetBillingAmount sets the billing_amount property value. Must be 0 or greater. Default is the price in the product subscription.
func (m *SubscriptionsPostRequestBody) SetBillingAmount(value *float64) {
	m.billing_amount = value
}

// SetContactId sets the contact_id property value. The contact_id property
func (m *SubscriptionsPostRequestBody) SetContactId(value *int64) {
	m.contact_id = value
}

// SetFirstBillDate sets the first_bill_date property value. The first day the subscription will bill, in EST. Must not be in the past. Default is today.
func (m *SubscriptionsPostRequestBody) SetFirstBillDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
	m.first_bill_date = value
}

// SetPaymentMethodId sets the payment_method_id property value. Default is the contact's most recently used card, if auto charge is true. Default is 0 otherwise.
func (m *SubscriptionsPostRequestBody) SetPaymentMethodId(value *string) {
	m.payment_method_id = value
}

// SetQuantity sets the quantity property value. Must be greater than 0. Default is 1.
func (m *SubscriptionsPostRequestBody) SetQuantity(value *int32) {
	m.quantity = value
}

// SetSaleAffiliateId sets the sale_affiliate_id property value. The sale_affiliate_id property
func (m *SubscriptionsPostRequestBody) SetSaleAffiliateId(value *int64) {
	m.sale_affiliate_id = value
}

// SetSubscriptionPlanId sets the subscription_plan_id property value. Id of the product subscription.
func (m *SubscriptionsPostRequestBody) SetSubscriptionPlanId(value *int64) {
	m.subscription_plan_id = value
}

type SubscriptionsPostRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAllowDuplicate() *bool
	GetAllowTax() *bool
	GetAutoCharge() *bool
	GetBillingAmount() *float64
	GetContactId() *int64
	GetFirstBillDate() *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
	GetPaymentMethodId() *string
	GetQuantity() *int32
	GetSaleAffiliateId() *int64
	GetSubscriptionPlanId() *int64
	SetAllowDuplicate(value *bool)
	SetAllowTax(value *bool)
	SetAutoCharge(value *bool)
	SetBillingAmount(value *float64)
	SetContactId(value *int64)
	SetFirstBillDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
	SetPaymentMethodId(value *string)
	SetQuantity(value *int32)
	SetSaleAffiliateId(value *int64)
	SetSubscriptionPlanId(value *int64)
}
