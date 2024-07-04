package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
)

type AffiliatesCommissionsGetResponse_commissions struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The amount_earned property
	amount_earned *float32
	// The contact_first_name property
	contact_first_name *string
	// The contact_id property
	contact_id *int64
	// The contact_last_name property
	contact_last_name *string
	// The date_earned property
	date_earned *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	// The description property
	description *string
	// The invoice_id property
	invoice_id *int64
	// The product_name property
	product_name *string
	// The sales_affiliate_id property
	sales_affiliate_id *int64
	// The sold_by_first_name property
	sold_by_first_name *string
	// The sold_by_last_name property
	sold_by_last_name *string
}

// NewAffiliatesCommissionsGetResponse_commissions instantiates a new AffiliatesCommissionsGetResponse_commissions and sets the default values.
func NewAffiliatesCommissionsGetResponse_commissions() *AffiliatesCommissionsGetResponse_commissions {
	m := &AffiliatesCommissionsGetResponse_commissions{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateAffiliatesCommissionsGetResponse_commissionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAffiliatesCommissionsGetResponse_commissionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAffiliatesCommissionsGetResponse_commissions(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AffiliatesCommissionsGetResponse_commissions) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAmountEarned gets the amount_earned property value. The amount_earned property
// returns a *float32 when successful
func (m *AffiliatesCommissionsGetResponse_commissions) GetAmountEarned() *float32 {
	return m.amount_earned
}

// GetContactFirstName gets the contact_first_name property value. The contact_first_name property
// returns a *string when successful
func (m *AffiliatesCommissionsGetResponse_commissions) GetContactFirstName() *string {
	return m.contact_first_name
}

// GetContactId gets the contact_id property value. The contact_id property
// returns a *int64 when successful
func (m *AffiliatesCommissionsGetResponse_commissions) GetContactId() *int64 {
	return m.contact_id
}

// GetContactLastName gets the contact_last_name property value. The contact_last_name property
// returns a *string when successful
func (m *AffiliatesCommissionsGetResponse_commissions) GetContactLastName() *string {
	return m.contact_last_name
}

// GetDateEarned gets the date_earned property value. The date_earned property
// returns a *Time when successful
func (m *AffiliatesCommissionsGetResponse_commissions) GetDateEarned() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.date_earned
}

// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *AffiliatesCommissionsGetResponse_commissions) GetDescription() *string {
	return m.description
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AffiliatesCommissionsGetResponse_commissions) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["amount_earned"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetFloat32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAmountEarned(val)
		}
		return nil
	}
	res["contact_first_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetContactFirstName(val)
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
	res["contact_last_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetContactLastName(val)
		}
		return nil
	}
	res["date_earned"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetTimeValue()
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
	res["sales_affiliate_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSalesAffiliateId(val)
		}
		return nil
	}
	res["sold_by_first_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSoldByFirstName(val)
		}
		return nil
	}
	res["sold_by_last_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSoldByLastName(val)
		}
		return nil
	}
	return res
}

// GetInvoiceId gets the invoice_id property value. The invoice_id property
// returns a *int64 when successful
func (m *AffiliatesCommissionsGetResponse_commissions) GetInvoiceId() *int64 {
	return m.invoice_id
}

// GetProductName gets the product_name property value. The product_name property
// returns a *string when successful
func (m *AffiliatesCommissionsGetResponse_commissions) GetProductName() *string {
	return m.product_name
}

// GetSalesAffiliateId gets the sales_affiliate_id property value. The sales_affiliate_id property
// returns a *int64 when successful
func (m *AffiliatesCommissionsGetResponse_commissions) GetSalesAffiliateId() *int64 {
	return m.sales_affiliate_id
}

// GetSoldByFirstName gets the sold_by_first_name property value. The sold_by_first_name property
// returns a *string when successful
func (m *AffiliatesCommissionsGetResponse_commissions) GetSoldByFirstName() *string {
	return m.sold_by_first_name
}

// GetSoldByLastName gets the sold_by_last_name property value. The sold_by_last_name property
// returns a *string when successful
func (m *AffiliatesCommissionsGetResponse_commissions) GetSoldByLastName() *string {
	return m.sold_by_last_name
}

// Serialize serializes information the current object
func (m *AffiliatesCommissionsGetResponse_commissions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteFloat32Value("amount_earned", m.GetAmountEarned())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("contact_first_name", m.GetContactFirstName())
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
		err := writer.WriteStringValue("contact_last_name", m.GetContactLastName())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteTimeValue("date_earned", m.GetDateEarned())
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
		err := writer.WriteInt64Value("sales_affiliate_id", m.GetSalesAffiliateId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("sold_by_first_name", m.GetSoldByFirstName())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("sold_by_last_name", m.GetSoldByLastName())
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
func (m *AffiliatesCommissionsGetResponse_commissions) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAmountEarned sets the amount_earned property value. The amount_earned property
func (m *AffiliatesCommissionsGetResponse_commissions) SetAmountEarned(value *float32) {
	m.amount_earned = value
}

// SetContactFirstName sets the contact_first_name property value. The contact_first_name property
func (m *AffiliatesCommissionsGetResponse_commissions) SetContactFirstName(value *string) {
	m.contact_first_name = value
}

// SetContactId sets the contact_id property value. The contact_id property
func (m *AffiliatesCommissionsGetResponse_commissions) SetContactId(value *int64) {
	m.contact_id = value
}

// SetContactLastName sets the contact_last_name property value. The contact_last_name property
func (m *AffiliatesCommissionsGetResponse_commissions) SetContactLastName(value *string) {
	m.contact_last_name = value
}

// SetDateEarned sets the date_earned property value. The date_earned property
func (m *AffiliatesCommissionsGetResponse_commissions) SetDateEarned(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	m.date_earned = value
}

// SetDescription sets the description property value. The description property
func (m *AffiliatesCommissionsGetResponse_commissions) SetDescription(value *string) {
	m.description = value
}

// SetInvoiceId sets the invoice_id property value. The invoice_id property
func (m *AffiliatesCommissionsGetResponse_commissions) SetInvoiceId(value *int64) {
	m.invoice_id = value
}

// SetProductName sets the product_name property value. The product_name property
func (m *AffiliatesCommissionsGetResponse_commissions) SetProductName(value *string) {
	m.product_name = value
}

// SetSalesAffiliateId sets the sales_affiliate_id property value. The sales_affiliate_id property
func (m *AffiliatesCommissionsGetResponse_commissions) SetSalesAffiliateId(value *int64) {
	m.sales_affiliate_id = value
}

// SetSoldByFirstName sets the sold_by_first_name property value. The sold_by_first_name property
func (m *AffiliatesCommissionsGetResponse_commissions) SetSoldByFirstName(value *string) {
	m.sold_by_first_name = value
}

// SetSoldByLastName sets the sold_by_last_name property value. The sold_by_last_name property
func (m *AffiliatesCommissionsGetResponse_commissions) SetSoldByLastName(value *string) {
	m.sold_by_last_name = value
}

type AffiliatesCommissionsGetResponse_commissionsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAmountEarned() *float32
	GetContactFirstName() *string
	GetContactId() *int64
	GetContactLastName() *string
	GetDateEarned() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetDescription() *string
	GetInvoiceId() *int64
	GetProductName() *string
	GetSalesAffiliateId() *int64
	GetSoldByFirstName() *string
	GetSoldByLastName() *string
	SetAmountEarned(value *float32)
	SetContactFirstName(value *string)
	SetContactId(value *int64)
	SetContactLastName(value *string)
	SetDateEarned(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	SetDescription(value *string)
	SetInvoiceId(value *int64)
	SetProductName(value *string)
	SetSalesAffiliateId(value *int64)
	SetSoldByFirstName(value *string)
	SetSoldByLastName(value *string)
}
