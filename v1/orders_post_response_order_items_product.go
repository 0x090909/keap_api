package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OrdersPostResponse_order_items_product struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The id property
	id *int64
	// The product_desc property
	product_desc *string
	// The product_name property
	product_name *string
	// The product_options property
	product_options []OrdersPostResponse_order_items_product_product_optionsable
	// The product_price property
	product_price *float64
	// The product_short_desc property
	product_short_desc *string
	// The sku property
	sku *string
	// The status property
	status *int32
	// The subscription_only property
	subscription_only *bool
	// The subscription_plans property
	subscription_plans []OrdersPostResponse_order_items_product_subscription_plansable
	// The url property
	url *string
}

// NewOrdersPostResponse_order_items_product instantiates a new OrdersPostResponse_order_items_product and sets the default values.
func NewOrdersPostResponse_order_items_product() *OrdersPostResponse_order_items_product {
	m := &OrdersPostResponse_order_items_product{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateOrdersPostResponse_order_items_productFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrdersPostResponse_order_items_productFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewOrdersPostResponse_order_items_product(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OrdersPostResponse_order_items_product) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrdersPostResponse_order_items_product) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
	res["product_desc"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetProductDesc(val)
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
	res["product_options"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateOrdersPostResponse_order_items_product_product_optionsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]OrdersPostResponse_order_items_product_product_optionsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(OrdersPostResponse_order_items_product_product_optionsable)
				}
			}
			m.SetProductOptions(res)
		}
		return nil
	}
	res["product_price"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetFloat64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetProductPrice(val)
		}
		return nil
	}
	res["product_short_desc"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetProductShortDesc(val)
		}
		return nil
	}
	res["sku"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSku(val)
		}
		return nil
	}
	res["status"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetStatus(val)
		}
		return nil
	}
	res["subscription_only"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSubscriptionOnly(val)
		}
		return nil
	}
	res["subscription_plans"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateOrdersPostResponse_order_items_product_subscription_plansFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]OrdersPostResponse_order_items_product_subscription_plansable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(OrdersPostResponse_order_items_product_subscription_plansable)
				}
			}
			m.SetSubscriptionPlans(res)
		}
		return nil
	}
	res["url"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetUrl(val)
		}
		return nil
	}
	return res
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *OrdersPostResponse_order_items_product) GetId() *int64 {
	return m.id
}

// GetProductDesc gets the product_desc property value. The product_desc property
// returns a *string when successful
func (m *OrdersPostResponse_order_items_product) GetProductDesc() *string {
	return m.product_desc
}

// GetProductName gets the product_name property value. The product_name property
// returns a *string when successful
func (m *OrdersPostResponse_order_items_product) GetProductName() *string {
	return m.product_name
}

// GetProductOptions gets the product_options property value. The product_options property
// returns a []OrdersPostResponse_order_items_product_product_optionsable when successful
func (m *OrdersPostResponse_order_items_product) GetProductOptions() []OrdersPostResponse_order_items_product_product_optionsable {
	return m.product_options
}

// GetProductPrice gets the product_price property value. The product_price property
// returns a *float64 when successful
func (m *OrdersPostResponse_order_items_product) GetProductPrice() *float64 {
	return m.product_price
}

// GetProductShortDesc gets the product_short_desc property value. The product_short_desc property
// returns a *string when successful
func (m *OrdersPostResponse_order_items_product) GetProductShortDesc() *string {
	return m.product_short_desc
}

// GetSku gets the sku property value. The sku property
// returns a *string when successful
func (m *OrdersPostResponse_order_items_product) GetSku() *string {
	return m.sku
}

// GetStatus gets the status property value. The status property
// returns a *int32 when successful
func (m *OrdersPostResponse_order_items_product) GetStatus() *int32 {
	return m.status
}

// GetSubscriptionOnly gets the subscription_only property value. The subscription_only property
// returns a *bool when successful
func (m *OrdersPostResponse_order_items_product) GetSubscriptionOnly() *bool {
	return m.subscription_only
}

// GetSubscriptionPlans gets the subscription_plans property value. The subscription_plans property
// returns a []OrdersPostResponse_order_items_product_subscription_plansable when successful
func (m *OrdersPostResponse_order_items_product) GetSubscriptionPlans() []OrdersPostResponse_order_items_product_subscription_plansable {
	return m.subscription_plans
}

// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *OrdersPostResponse_order_items_product) GetUrl() *string {
	return m.url
}

// Serialize serializes information the current object
func (m *OrdersPostResponse_order_items_product) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt64Value("id", m.GetId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("product_desc", m.GetProductDesc())
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
	if m.GetProductOptions() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProductOptions()))
		for i, v := range m.GetProductOptions() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("product_options", cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteFloat64Value("product_price", m.GetProductPrice())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("product_short_desc", m.GetProductShortDesc())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("sku", m.GetSku())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("status", m.GetStatus())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("subscription_only", m.GetSubscriptionOnly())
		if err != nil {
			return err
		}
	}
	if m.GetSubscriptionPlans() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSubscriptionPlans()))
		for i, v := range m.GetSubscriptionPlans() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("subscription_plans", cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("url", m.GetUrl())
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
func (m *OrdersPostResponse_order_items_product) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetId sets the id property value. The id property
func (m *OrdersPostResponse_order_items_product) SetId(value *int64) {
	m.id = value
}

// SetProductDesc sets the product_desc property value. The product_desc property
func (m *OrdersPostResponse_order_items_product) SetProductDesc(value *string) {
	m.product_desc = value
}

// SetProductName sets the product_name property value. The product_name property
func (m *OrdersPostResponse_order_items_product) SetProductName(value *string) {
	m.product_name = value
}

// SetProductOptions sets the product_options property value. The product_options property
func (m *OrdersPostResponse_order_items_product) SetProductOptions(value []OrdersPostResponse_order_items_product_product_optionsable) {
	m.product_options = value
}

// SetProductPrice sets the product_price property value. The product_price property
func (m *OrdersPostResponse_order_items_product) SetProductPrice(value *float64) {
	m.product_price = value
}

// SetProductShortDesc sets the product_short_desc property value. The product_short_desc property
func (m *OrdersPostResponse_order_items_product) SetProductShortDesc(value *string) {
	m.product_short_desc = value
}

// SetSku sets the sku property value. The sku property
func (m *OrdersPostResponse_order_items_product) SetSku(value *string) {
	m.sku = value
}

// SetStatus sets the status property value. The status property
func (m *OrdersPostResponse_order_items_product) SetStatus(value *int32) {
	m.status = value
}

// SetSubscriptionOnly sets the subscription_only property value. The subscription_only property
func (m *OrdersPostResponse_order_items_product) SetSubscriptionOnly(value *bool) {
	m.subscription_only = value
}

// SetSubscriptionPlans sets the subscription_plans property value. The subscription_plans property
func (m *OrdersPostResponse_order_items_product) SetSubscriptionPlans(value []OrdersPostResponse_order_items_product_subscription_plansable) {
	m.subscription_plans = value
}

// SetUrl sets the url property value. The url property
func (m *OrdersPostResponse_order_items_product) SetUrl(value *string) {
	m.url = value
}

type OrdersPostResponse_order_items_productable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetId() *int64
	GetProductDesc() *string
	GetProductName() *string
	GetProductOptions() []OrdersPostResponse_order_items_product_product_optionsable
	GetProductPrice() *float64
	GetProductShortDesc() *string
	GetSku() *string
	GetStatus() *int32
	GetSubscriptionOnly() *bool
	GetSubscriptionPlans() []OrdersPostResponse_order_items_product_subscription_plansable
	GetUrl() *string
	SetId(value *int64)
	SetProductDesc(value *string)
	SetProductName(value *string)
	SetProductOptions(value []OrdersPostResponse_order_items_product_product_optionsable)
	SetProductPrice(value *float64)
	SetProductShortDesc(value *string)
	SetSku(value *string)
	SetStatus(value *int32)
	SetSubscriptionOnly(value *bool)
	SetSubscriptionPlans(value []OrdersPostResponse_order_items_product_subscription_plansable)
	SetUrl(value *string)
}
