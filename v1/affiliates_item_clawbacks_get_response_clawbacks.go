package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AffiliatesItemClawbacksGetResponse_clawbacks struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The amount property
	amount *float64
	// The contact_id property
	contact_id *int64
	// The date_earned property
	date_earned *string
	// The description property
	description *string
	// The family_name property
	family_name *string
	// The given_name property
	given_name *string
	// The invoice_id property
	invoice_id *int64
	// The product_name property
	product_name *string
	// The sale_affiliate_id property
	sale_affiliate_id *int64
	// The sold_by_family_name property
	sold_by_family_name *string
	// The sold_by_given_name property
	sold_by_given_name *string
	// The subscription_plan_name property
	subscription_plan_name *string
}

// NewAffiliatesItemClawbacksGetResponse_clawbacks instantiates a new AffiliatesItemClawbacksGetResponse_clawbacks and sets the default values.
func NewAffiliatesItemClawbacksGetResponse_clawbacks() *AffiliatesItemClawbacksGetResponse_clawbacks {
	m := &AffiliatesItemClawbacksGetResponse_clawbacks{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateAffiliatesItemClawbacksGetResponse_clawbacksFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAffiliatesItemClawbacksGetResponse_clawbacksFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAffiliatesItemClawbacksGetResponse_clawbacks(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AffiliatesItemClawbacksGetResponse_clawbacks) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAmount gets the amount property value. The amount property
// returns a *float64 when successful
func (m *AffiliatesItemClawbacksGetResponse_clawbacks) GetAmount() *float64 {
	return m.amount
}

// GetContactId gets the contact_id property value. The contact_id property
// returns a *int64 when successful
func (m *AffiliatesItemClawbacksGetResponse_clawbacks) GetContactId() *int64 {
	return m.contact_id
}

// GetDateEarned gets the date_earned property value. The date_earned property
// returns a *string when successful
func (m *AffiliatesItemClawbacksGetResponse_clawbacks) GetDateEarned() *string {
	return m.date_earned
}

// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *AffiliatesItemClawbacksGetResponse_clawbacks) GetDescription() *string {
	return m.description
}

// GetFamilyName gets the family_name property value. The family_name property
// returns a *string when successful
func (m *AffiliatesItemClawbacksGetResponse_clawbacks) GetFamilyName() *string {
	return m.family_name
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AffiliatesItemClawbacksGetResponse_clawbacks) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
	res["date_earned"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDateEarned(val)
		}
		return nil
	}
	res["description"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDescription(val)
		}
		return nil
	}
	res["family_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFamilyName(val)
		}
		return nil
	}
	res["given_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetGivenName(val)
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
	res["product_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetProductName(val)
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
	res["sold_by_family_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSoldByFamilyName(val)
		}
		return nil
	}
	res["sold_by_given_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSoldByGivenName(val)
		}
		return nil
	}
	res["subscription_plan_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSubscriptionPlanName(val)
		}
		return nil
	}
	return res
}

// GetGivenName gets the given_name property value. The given_name property
// returns a *string when successful
func (m *AffiliatesItemClawbacksGetResponse_clawbacks) GetGivenName() *string {
	return m.given_name
}

// GetInvoiceId gets the invoice_id property value. The invoice_id property
// returns a *int64 when successful
func (m *AffiliatesItemClawbacksGetResponse_clawbacks) GetInvoiceId() *int64 {
	return m.invoice_id
}

// GetProductName gets the product_name property value. The product_name property
// returns a *string when successful
func (m *AffiliatesItemClawbacksGetResponse_clawbacks) GetProductName() *string {
	return m.product_name
}

// GetSaleAffiliateId gets the sale_affiliate_id property value. The sale_affiliate_id property
// returns a *int64 when successful
func (m *AffiliatesItemClawbacksGetResponse_clawbacks) GetSaleAffiliateId() *int64 {
	return m.sale_affiliate_id
}

// GetSoldByFamilyName gets the sold_by_family_name property value. The sold_by_family_name property
// returns a *string when successful
func (m *AffiliatesItemClawbacksGetResponse_clawbacks) GetSoldByFamilyName() *string {
	return m.sold_by_family_name
}

// GetSoldByGivenName gets the sold_by_given_name property value. The sold_by_given_name property
// returns a *string when successful
func (m *AffiliatesItemClawbacksGetResponse_clawbacks) GetSoldByGivenName() *string {
	return m.sold_by_given_name
}

// GetSubscriptionPlanName gets the subscription_plan_name property value. The subscription_plan_name property
// returns a *string when successful
func (m *AffiliatesItemClawbacksGetResponse_clawbacks) GetSubscriptionPlanName() *string {
	return m.subscription_plan_name
}

// Serialize serializes information the current object
func (m *AffiliatesItemClawbacksGetResponse_clawbacks) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteFloat64Value("amount", m.GetAmount())
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
		err := writer.WriteStringValue("date_earned", m.GetDateEarned())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("description", m.GetDescription())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("family_name", m.GetFamilyName())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("given_name", m.GetGivenName())
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
		err := writer.WriteStringValue("product_name", m.GetProductName())
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
		err := writer.WriteStringValue("sold_by_family_name", m.GetSoldByFamilyName())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("sold_by_given_name", m.GetSoldByGivenName())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("subscription_plan_name", m.GetSubscriptionPlanName())
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
func (m *AffiliatesItemClawbacksGetResponse_clawbacks) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAmount sets the amount property value. The amount property
func (m *AffiliatesItemClawbacksGetResponse_clawbacks) SetAmount(value *float64) {
	m.amount = value
}

// SetContactId sets the contact_id property value. The contact_id property
func (m *AffiliatesItemClawbacksGetResponse_clawbacks) SetContactId(value *int64) {
	m.contact_id = value
}

// SetDateEarned sets the date_earned property value. The date_earned property
func (m *AffiliatesItemClawbacksGetResponse_clawbacks) SetDateEarned(value *string) {
	m.date_earned = value
}

// SetDescription sets the description property value. The description property
func (m *AffiliatesItemClawbacksGetResponse_clawbacks) SetDescription(value *string) {
	m.description = value
}

// SetFamilyName sets the family_name property value. The family_name property
func (m *AffiliatesItemClawbacksGetResponse_clawbacks) SetFamilyName(value *string) {
	m.family_name = value
}

// SetGivenName sets the given_name property value. The given_name property
func (m *AffiliatesItemClawbacksGetResponse_clawbacks) SetGivenName(value *string) {
	m.given_name = value
}

// SetInvoiceId sets the invoice_id property value. The invoice_id property
func (m *AffiliatesItemClawbacksGetResponse_clawbacks) SetInvoiceId(value *int64) {
	m.invoice_id = value
}

// SetProductName sets the product_name property value. The product_name property
func (m *AffiliatesItemClawbacksGetResponse_clawbacks) SetProductName(value *string) {
	m.product_name = value
}

// SetSaleAffiliateId sets the sale_affiliate_id property value. The sale_affiliate_id property
func (m *AffiliatesItemClawbacksGetResponse_clawbacks) SetSaleAffiliateId(value *int64) {
	m.sale_affiliate_id = value
}

// SetSoldByFamilyName sets the sold_by_family_name property value. The sold_by_family_name property
func (m *AffiliatesItemClawbacksGetResponse_clawbacks) SetSoldByFamilyName(value *string) {
	m.sold_by_family_name = value
}

// SetSoldByGivenName sets the sold_by_given_name property value. The sold_by_given_name property
func (m *AffiliatesItemClawbacksGetResponse_clawbacks) SetSoldByGivenName(value *string) {
	m.sold_by_given_name = value
}

// SetSubscriptionPlanName sets the subscription_plan_name property value. The subscription_plan_name property
func (m *AffiliatesItemClawbacksGetResponse_clawbacks) SetSubscriptionPlanName(value *string) {
	m.subscription_plan_name = value
}

type AffiliatesItemClawbacksGetResponse_clawbacksable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAmount() *float64
	GetContactId() *int64
	GetDateEarned() *string
	GetDescription() *string
	GetFamilyName() *string
	GetGivenName() *string
	GetInvoiceId() *int64
	GetProductName() *string
	GetSaleAffiliateId() *int64
	GetSoldByFamilyName() *string
	GetSoldByGivenName() *string
	GetSubscriptionPlanName() *string
	SetAmount(value *float64)
	SetContactId(value *int64)
	SetDateEarned(value *string)
	SetDescription(value *string)
	SetFamilyName(value *string)
	SetGivenName(value *string)
	SetInvoiceId(value *int64)
	SetProductName(value *string)
	SetSaleAffiliateId(value *int64)
	SetSoldByFamilyName(value *string)
	SetSoldByGivenName(value *string)
	SetSubscriptionPlanName(value *string)
}
