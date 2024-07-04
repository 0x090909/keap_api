package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OrdersPostResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The contact property
	contact OrdersPostResponse_contactable
	// The creation_date property
	creation_date *string
	// The id property
	id *int64
	// The lead_affiliate_id property
	lead_affiliate_id *int64
	// The modification_date property
	modification_date *string
	// The notes property
	notes *string
	// The order_date property
	order_date *string
	// The order_items property
	order_items []OrdersPostResponse_order_itemsable
	// The payment_plan property
	payment_plan OrdersPostResponse_payment_planable
	// The refund_total property
	refund_total *float64
	// The sales_affiliate_id property
	sales_affiliate_id *int64
	// The shipping_information property
	shipping_information OrdersPostResponse_shipping_informationable
	// The status property
	status *string
	// The title property
	title *string
	// The total property
	total *float64
	// The total_due property
	total_due *float64
	// The total_paid property
	total_paid *float64
}

// NewOrdersPostResponse instantiates a new OrdersPostResponse and sets the default values.
func NewOrdersPostResponse() *OrdersPostResponse {
	m := &OrdersPostResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateOrdersPostResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrdersPostResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewOrdersPostResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OrdersPostResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetContact gets the contact property value. The contact property
// returns a OrdersPostResponse_contactable when successful
func (m *OrdersPostResponse) GetContact() OrdersPostResponse_contactable {
	return m.contact
}

// GetCreationDate gets the creation_date property value. The creation_date property
// returns a *string when successful
func (m *OrdersPostResponse) GetCreationDate() *string {
	return m.creation_date
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrdersPostResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["contact"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateOrdersPostResponse_contactFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetContact(val.(OrdersPostResponse_contactable))
		}
		return nil
	}
	res["creation_date"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCreationDate(val)
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
	res["lead_affiliate_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLeadAffiliateId(val)
		}
		return nil
	}
	res["modification_date"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetModificationDate(val)
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
	res["order_date"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetOrderDate(val)
		}
		return nil
	}
	res["order_items"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateOrdersPostResponse_order_itemsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]OrdersPostResponse_order_itemsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(OrdersPostResponse_order_itemsable)
				}
			}
			m.SetOrderItems(res)
		}
		return nil
	}
	res["payment_plan"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateOrdersPostResponse_payment_planFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPaymentPlan(val.(OrdersPostResponse_payment_planable))
		}
		return nil
	}
	res["refund_total"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetFloat64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRefundTotal(val)
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
	res["shipping_information"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateOrdersPostResponse_shipping_informationFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetShippingInformation(val.(OrdersPostResponse_shipping_informationable))
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
	res["title"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTitle(val)
		}
		return nil
	}
	res["total"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetFloat64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTotal(val)
		}
		return nil
	}
	res["total_due"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetFloat64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTotalDue(val)
		}
		return nil
	}
	res["total_paid"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetFloat64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTotalPaid(val)
		}
		return nil
	}
	return res
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *OrdersPostResponse) GetId() *int64 {
	return m.id
}

// GetLeadAffiliateId gets the lead_affiliate_id property value. The lead_affiliate_id property
// returns a *int64 when successful
func (m *OrdersPostResponse) GetLeadAffiliateId() *int64 {
	return m.lead_affiliate_id
}

// GetModificationDate gets the modification_date property value. The modification_date property
// returns a *string when successful
func (m *OrdersPostResponse) GetModificationDate() *string {
	return m.modification_date
}

// GetNotes gets the notes property value. The notes property
// returns a *string when successful
func (m *OrdersPostResponse) GetNotes() *string {
	return m.notes
}

// GetOrderDate gets the order_date property value. The order_date property
// returns a *string when successful
func (m *OrdersPostResponse) GetOrderDate() *string {
	return m.order_date
}

// GetOrderItems gets the order_items property value. The order_items property
// returns a []OrdersPostResponse_order_itemsable when successful
func (m *OrdersPostResponse) GetOrderItems() []OrdersPostResponse_order_itemsable {
	return m.order_items
}

// GetPaymentPlan gets the payment_plan property value. The payment_plan property
// returns a OrdersPostResponse_payment_planable when successful
func (m *OrdersPostResponse) GetPaymentPlan() OrdersPostResponse_payment_planable {
	return m.payment_plan
}

// GetRefundTotal gets the refund_total property value. The refund_total property
// returns a *float64 when successful
func (m *OrdersPostResponse) GetRefundTotal() *float64 {
	return m.refund_total
}

// GetSalesAffiliateId gets the sales_affiliate_id property value. The sales_affiliate_id property
// returns a *int64 when successful
func (m *OrdersPostResponse) GetSalesAffiliateId() *int64 {
	return m.sales_affiliate_id
}

// GetShippingInformation gets the shipping_information property value. The shipping_information property
// returns a OrdersPostResponse_shipping_informationable when successful
func (m *OrdersPostResponse) GetShippingInformation() OrdersPostResponse_shipping_informationable {
	return m.shipping_information
}

// GetStatus gets the status property value. The status property
// returns a *string when successful
func (m *OrdersPostResponse) GetStatus() *string {
	return m.status
}

// GetTitle gets the title property value. The title property
// returns a *string when successful
func (m *OrdersPostResponse) GetTitle() *string {
	return m.title
}

// GetTotal gets the total property value. The total property
// returns a *float64 when successful
func (m *OrdersPostResponse) GetTotal() *float64 {
	return m.total
}

// GetTotalDue gets the total_due property value. The total_due property
// returns a *float64 when successful
func (m *OrdersPostResponse) GetTotalDue() *float64 {
	return m.total_due
}

// GetTotalPaid gets the total_paid property value. The total_paid property
// returns a *float64 when successful
func (m *OrdersPostResponse) GetTotalPaid() *float64 {
	return m.total_paid
}

// Serialize serializes information the current object
func (m *OrdersPostResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("contact", m.GetContact())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("creation_date", m.GetCreationDate())
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
		err := writer.WriteInt64Value("lead_affiliate_id", m.GetLeadAffiliateId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("modification_date", m.GetModificationDate())
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
		err := writer.WriteStringValue("order_date", m.GetOrderDate())
		if err != nil {
			return err
		}
	}
	if m.GetOrderItems() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOrderItems()))
		for i, v := range m.GetOrderItems() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("order_items", cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("payment_plan", m.GetPaymentPlan())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteFloat64Value("refund_total", m.GetRefundTotal())
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
		err := writer.WriteObjectValue("shipping_information", m.GetShippingInformation())
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
		err := writer.WriteStringValue("title", m.GetTitle())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteFloat64Value("total", m.GetTotal())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteFloat64Value("total_due", m.GetTotalDue())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteFloat64Value("total_paid", m.GetTotalPaid())
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
func (m *OrdersPostResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetContact sets the contact property value. The contact property
func (m *OrdersPostResponse) SetContact(value OrdersPostResponse_contactable) {
	m.contact = value
}

// SetCreationDate sets the creation_date property value. The creation_date property
func (m *OrdersPostResponse) SetCreationDate(value *string) {
	m.creation_date = value
}

// SetId sets the id property value. The id property
func (m *OrdersPostResponse) SetId(value *int64) {
	m.id = value
}

// SetLeadAffiliateId sets the lead_affiliate_id property value. The lead_affiliate_id property
func (m *OrdersPostResponse) SetLeadAffiliateId(value *int64) {
	m.lead_affiliate_id = value
}

// SetModificationDate sets the modification_date property value. The modification_date property
func (m *OrdersPostResponse) SetModificationDate(value *string) {
	m.modification_date = value
}

// SetNotes sets the notes property value. The notes property
func (m *OrdersPostResponse) SetNotes(value *string) {
	m.notes = value
}

// SetOrderDate sets the order_date property value. The order_date property
func (m *OrdersPostResponse) SetOrderDate(value *string) {
	m.order_date = value
}

// SetOrderItems sets the order_items property value. The order_items property
func (m *OrdersPostResponse) SetOrderItems(value []OrdersPostResponse_order_itemsable) {
	m.order_items = value
}

// SetPaymentPlan sets the payment_plan property value. The payment_plan property
func (m *OrdersPostResponse) SetPaymentPlan(value OrdersPostResponse_payment_planable) {
	m.payment_plan = value
}

// SetRefundTotal sets the refund_total property value. The refund_total property
func (m *OrdersPostResponse) SetRefundTotal(value *float64) {
	m.refund_total = value
}

// SetSalesAffiliateId sets the sales_affiliate_id property value. The sales_affiliate_id property
func (m *OrdersPostResponse) SetSalesAffiliateId(value *int64) {
	m.sales_affiliate_id = value
}

// SetShippingInformation sets the shipping_information property value. The shipping_information property
func (m *OrdersPostResponse) SetShippingInformation(value OrdersPostResponse_shipping_informationable) {
	m.shipping_information = value
}

// SetStatus sets the status property value. The status property
func (m *OrdersPostResponse) SetStatus(value *string) {
	m.status = value
}

// SetTitle sets the title property value. The title property
func (m *OrdersPostResponse) SetTitle(value *string) {
	m.title = value
}

// SetTotal sets the total property value. The total property
func (m *OrdersPostResponse) SetTotal(value *float64) {
	m.total = value
}

// SetTotalDue sets the total_due property value. The total_due property
func (m *OrdersPostResponse) SetTotalDue(value *float64) {
	m.total_due = value
}

// SetTotalPaid sets the total_paid property value. The total_paid property
func (m *OrdersPostResponse) SetTotalPaid(value *float64) {
	m.total_paid = value
}

type OrdersPostResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetContact() OrdersPostResponse_contactable
	GetCreationDate() *string
	GetId() *int64
	GetLeadAffiliateId() *int64
	GetModificationDate() *string
	GetNotes() *string
	GetOrderDate() *string
	GetOrderItems() []OrdersPostResponse_order_itemsable
	GetPaymentPlan() OrdersPostResponse_payment_planable
	GetRefundTotal() *float64
	GetSalesAffiliateId() *int64
	GetShippingInformation() OrdersPostResponse_shipping_informationable
	GetStatus() *string
	GetTitle() *string
	GetTotal() *float64
	GetTotalDue() *float64
	GetTotalPaid() *float64
	SetContact(value OrdersPostResponse_contactable)
	SetCreationDate(value *string)
	SetId(value *int64)
	SetLeadAffiliateId(value *int64)
	SetModificationDate(value *string)
	SetNotes(value *string)
	SetOrderDate(value *string)
	SetOrderItems(value []OrdersPostResponse_order_itemsable)
	SetPaymentPlan(value OrdersPostResponse_payment_planable)
	SetRefundTotal(value *float64)
	SetSalesAffiliateId(value *int64)
	SetShippingInformation(value OrdersPostResponse_shipping_informationable)
	SetStatus(value *string)
	SetTitle(value *string)
	SetTotal(value *float64)
	SetTotalDue(value *float64)
	SetTotalPaid(value *float64)
}
