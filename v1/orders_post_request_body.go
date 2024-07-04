package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OrdersPostRequestBody struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The contact_id property
	contact_id *int64
	// The lead_affiliate_id property
	lead_affiliate_id *int64
	// The order_date property
	order_date *string
	// The order_items property
	order_items []OrdersPostRequestBody_order_itemsable
	// The order_title property
	order_title *string
	// Uses multiple strings as promo codes. The corresponding discount will be applied to the order.
	promo_codes []string
	// The sales_affiliate_id property
	sales_affiliate_id *int64
	// The shipping_address property
	shipping_address OrdersPostRequestBody_shipping_addressable
}

// NewOrdersPostRequestBody instantiates a new OrdersPostRequestBody and sets the default values.
func NewOrdersPostRequestBody() *OrdersPostRequestBody {
	m := &OrdersPostRequestBody{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateOrdersPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrdersPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewOrdersPostRequestBody(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OrdersPostRequestBody) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetContactId gets the contact_id property value. The contact_id property
// returns a *int64 when successful
func (m *OrdersPostRequestBody) GetContactId() *int64 {
	return m.contact_id
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrdersPostRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
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
		val, err := n.GetCollectionOfObjectValues(CreateOrdersPostRequestBody_order_itemsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]OrdersPostRequestBody_order_itemsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(OrdersPostRequestBody_order_itemsable)
				}
			}
			m.SetOrderItems(res)
		}
		return nil
	}
	res["order_title"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetOrderTitle(val)
		}
		return nil
	}
	res["promo_codes"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfPrimitiveValues("string")
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]string, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = *(v.(*string))
				}
			}
			m.SetPromoCodes(res)
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
	res["shipping_address"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateOrdersPostRequestBody_shipping_addressFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetShippingAddress(val.(OrdersPostRequestBody_shipping_addressable))
		}
		return nil
	}
	return res
}

// GetLeadAffiliateId gets the lead_affiliate_id property value. The lead_affiliate_id property
// returns a *int64 when successful
func (m *OrdersPostRequestBody) GetLeadAffiliateId() *int64 {
	return m.lead_affiliate_id
}

// GetOrderDate gets the order_date property value. The order_date property
// returns a *Time when successful
func (m *OrdersPostRequestBody) GetOrderDate() *string {
	return m.order_date
}

// GetOrderItems gets the order_items property value. The order_items property
// returns a []OrdersPostRequestBody_order_itemsable when successful
func (m *OrdersPostRequestBody) GetOrderItems() []OrdersPostRequestBody_order_itemsable {
	return m.order_items
}

// GetOrderTitle gets the order_title property value. The order_title property
// returns a *string when successful
func (m *OrdersPostRequestBody) GetOrderTitle() *string {
	return m.order_title
}

// GetPromoCodes gets the promo_codes property value. Uses multiple strings as promo codes. The corresponding discount will be applied to the order.
// returns a []string when successful
func (m *OrdersPostRequestBody) GetPromoCodes() []string {
	return m.promo_codes
}

// GetSalesAffiliateId gets the sales_affiliate_id property value. The sales_affiliate_id property
// returns a *int64 when successful
func (m *OrdersPostRequestBody) GetSalesAffiliateId() *int64 {
	return m.sales_affiliate_id
}

// GetShippingAddress gets the shipping_address property value. The shipping_address property
// returns a OrdersPostRequestBody_shipping_addressable when successful
func (m *OrdersPostRequestBody) GetShippingAddress() OrdersPostRequestBody_shipping_addressable {
	return m.shipping_address
}

// Serialize serializes information the current object
func (m *OrdersPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt64Value("contact_id", m.GetContactId())
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
		err := writer.WriteStringValue("order_title", m.GetOrderTitle())
		if err != nil {
			return err
		}
	}
	if m.GetPromoCodes() != nil {
		err := writer.WriteCollectionOfStringValues("promo_codes", m.GetPromoCodes())
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
		err := writer.WriteObjectValue("shipping_address", m.GetShippingAddress())
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
func (m *OrdersPostRequestBody) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetContactId sets the contact_id property value. The contact_id property
func (m *OrdersPostRequestBody) SetContactId(value *int64) {
	m.contact_id = value
}

// SetLeadAffiliateId sets the lead_affiliate_id property value. The lead_affiliate_id property
func (m *OrdersPostRequestBody) SetLeadAffiliateId(value *int64) {
	m.lead_affiliate_id = value
}

// SetOrderDate sets the order_date property value. The order_date property
func (m *OrdersPostRequestBody) SetOrderDate(value *string) {
	m.order_date = value
}

// SetOrderItems sets the order_items property value. The order_items property
func (m *OrdersPostRequestBody) SetOrderItems(value []OrdersPostRequestBody_order_itemsable) {
	m.order_items = value
}

// SetOrderTitle sets the order_title property value. The order_title property
func (m *OrdersPostRequestBody) SetOrderTitle(value *string) {
	m.order_title = value
}

// SetPromoCodes sets the promo_codes property value. Uses multiple strings as promo codes. The corresponding discount will be applied to the order.
func (m *OrdersPostRequestBody) SetPromoCodes(value []string) {
	m.promo_codes = value
}

// SetSalesAffiliateId sets the sales_affiliate_id property value. The sales_affiliate_id property
func (m *OrdersPostRequestBody) SetSalesAffiliateId(value *int64) {
	m.sales_affiliate_id = value
}

// SetShippingAddress sets the shipping_address property value. The shipping_address property
func (m *OrdersPostRequestBody) SetShippingAddress(value OrdersPostRequestBody_shipping_addressable) {
	m.shipping_address = value
}

type OrdersPostRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetContactId() *int64
	GetLeadAffiliateId() *int64
	GetOrderDate() *string
	GetOrderItems() []OrdersPostRequestBody_order_itemsable
	GetOrderTitle() *string
	GetPromoCodes() []string
	GetSalesAffiliateId() *int64
	GetShippingAddress() OrdersPostRequestBody_shipping_addressable
	SetContactId(value *int64)
	SetLeadAffiliateId(value *int64)
	SetOrderDate(value *string)
	SetOrderItems(value []OrdersPostRequestBody_order_itemsable)
	SetOrderTitle(value *string)
	SetPromoCodes(value []string)
	SetSalesAffiliateId(value *int64)
	SetShippingAddress(value OrdersPostRequestBody_shipping_addressable)
}
