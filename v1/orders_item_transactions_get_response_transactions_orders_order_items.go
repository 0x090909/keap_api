package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OrdersItemTransactionsGetResponse_transactions_orders_order_items struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The cost property
	cost *float64
	// The description property
	description *string
	// The discount property
	discount *float64
	// The id property
	id *int64
	// The jobRecurringId property
	jobRecurringId *int64
	// The name property
	name *string
	// The notes property
	notes *string
	// The price property
	price *float64
	// The product property
	product OrdersItemTransactionsGetResponse_transactions_orders_order_items_productable
	// The quantity property
	quantity *int32
	// The specialAmount property
	specialAmount *float64
	// The specialId property
	specialId *int64
	// The specialPctOrAmt property
	specialPctOrAmt *int32
	// The subscriptionPlan property
	subscriptionPlan OrdersItemTransactionsGetResponse_transactions_orders_order_items_subscriptionPlanable
	// The type property
	typeEscaped *string
}

// NewOrdersItemTransactionsGetResponse_transactions_orders_order_items instantiates a new OrdersItemTransactionsGetResponse_transactions_orders_order_items and sets the default values.
func NewOrdersItemTransactionsGetResponse_transactions_orders_order_items() *OrdersItemTransactionsGetResponse_transactions_orders_order_items {
	m := &OrdersItemTransactionsGetResponse_transactions_orders_order_items{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateOrdersItemTransactionsGetResponse_transactions_orders_order_itemsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrdersItemTransactionsGetResponse_transactions_orders_order_itemsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewOrdersItemTransactionsGetResponse_transactions_orders_order_items(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OrdersItemTransactionsGetResponse_transactions_orders_order_items) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCost gets the cost property value. The cost property
// returns a *float64 when successful
func (m *OrdersItemTransactionsGetResponse_transactions_orders_order_items) GetCost() *float64 {
	return m.cost
}

// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *OrdersItemTransactionsGetResponse_transactions_orders_order_items) GetDescription() *string {
	return m.description
}

// GetDiscount gets the discount property value. The discount property
// returns a *float64 when successful
func (m *OrdersItemTransactionsGetResponse_transactions_orders_order_items) GetDiscount() *float64 {
	return m.discount
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrdersItemTransactionsGetResponse_transactions_orders_order_items) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["cost"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetFloat64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCost(val)
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
	res["discount"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetFloat64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDiscount(val)
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
	res["jobRecurringId"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetJobRecurringId(val)
		}
		return nil
	}
	res["name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetName(val)
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
	res["price"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetFloat64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPrice(val)
		}
		return nil
	}
	res["product"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateOrdersItemTransactionsGetResponse_transactions_orders_order_items_productFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetProduct(val.(OrdersItemTransactionsGetResponse_transactions_orders_order_items_productable))
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
	res["specialAmount"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetFloat64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSpecialAmount(val)
		}
		return nil
	}
	res["specialId"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSpecialId(val)
		}
		return nil
	}
	res["specialPctOrAmt"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSpecialPctOrAmt(val)
		}
		return nil
	}
	res["subscriptionPlan"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateOrdersItemTransactionsGetResponse_transactions_orders_order_items_subscriptionPlanFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSubscriptionPlan(val.(OrdersItemTransactionsGetResponse_transactions_orders_order_items_subscriptionPlanable))
		}
		return nil
	}
	res["type"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTypeEscaped(val)
		}
		return nil
	}
	return res
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *OrdersItemTransactionsGetResponse_transactions_orders_order_items) GetId() *int64 {
	return m.id
}

// GetJobRecurringId gets the jobRecurringId property value. The jobRecurringId property
// returns a *int64 when successful
func (m *OrdersItemTransactionsGetResponse_transactions_orders_order_items) GetJobRecurringId() *int64 {
	return m.jobRecurringId
}

// GetName gets the name property value. The name property
// returns a *string when successful
func (m *OrdersItemTransactionsGetResponse_transactions_orders_order_items) GetName() *string {
	return m.name
}

// GetNotes gets the notes property value. The notes property
// returns a *string when successful
func (m *OrdersItemTransactionsGetResponse_transactions_orders_order_items) GetNotes() *string {
	return m.notes
}

// GetPrice gets the price property value. The price property
// returns a *float64 when successful
func (m *OrdersItemTransactionsGetResponse_transactions_orders_order_items) GetPrice() *float64 {
	return m.price
}

// GetProduct gets the product property value. The product property
// returns a OrdersItemTransactionsGetResponse_transactions_orders_order_items_productable when successful
func (m *OrdersItemTransactionsGetResponse_transactions_orders_order_items) GetProduct() OrdersItemTransactionsGetResponse_transactions_orders_order_items_productable {
	return m.product
}

// GetQuantity gets the quantity property value. The quantity property
// returns a *int32 when successful
func (m *OrdersItemTransactionsGetResponse_transactions_orders_order_items) GetQuantity() *int32 {
	return m.quantity
}

// GetSpecialAmount gets the specialAmount property value. The specialAmount property
// returns a *float64 when successful
func (m *OrdersItemTransactionsGetResponse_transactions_orders_order_items) GetSpecialAmount() *float64 {
	return m.specialAmount
}

// GetSpecialId gets the specialId property value. The specialId property
// returns a *int64 when successful
func (m *OrdersItemTransactionsGetResponse_transactions_orders_order_items) GetSpecialId() *int64 {
	return m.specialId
}

// GetSpecialPctOrAmt gets the specialPctOrAmt property value. The specialPctOrAmt property
// returns a *int32 when successful
func (m *OrdersItemTransactionsGetResponse_transactions_orders_order_items) GetSpecialPctOrAmt() *int32 {
	return m.specialPctOrAmt
}

// GetSubscriptionPlan gets the subscriptionPlan property value. The subscriptionPlan property
// returns a OrdersItemTransactionsGetResponse_transactions_orders_order_items_subscriptionPlanable when successful
func (m *OrdersItemTransactionsGetResponse_transactions_orders_order_items) GetSubscriptionPlan() OrdersItemTransactionsGetResponse_transactions_orders_order_items_subscriptionPlanable {
	return m.subscriptionPlan
}

// GetTypeEscaped gets the type property value. The type property
// returns a *string when successful
func (m *OrdersItemTransactionsGetResponse_transactions_orders_order_items) GetTypeEscaped() *string {
	return m.typeEscaped
}

// Serialize serializes information the current object
func (m *OrdersItemTransactionsGetResponse_transactions_orders_order_items) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteFloat64Value("cost", m.GetCost())
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
		err := writer.WriteFloat64Value("discount", m.GetDiscount())
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
		err := writer.WriteInt64Value("jobRecurringId", m.GetJobRecurringId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("name", m.GetName())
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
		err := writer.WriteFloat64Value("price", m.GetPrice())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("product", m.GetProduct())
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
		err := writer.WriteFloat64Value("specialAmount", m.GetSpecialAmount())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("specialId", m.GetSpecialId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("specialPctOrAmt", m.GetSpecialPctOrAmt())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("subscriptionPlan", m.GetSubscriptionPlan())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("type", m.GetTypeEscaped())
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
func (m *OrdersItemTransactionsGetResponse_transactions_orders_order_items) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCost sets the cost property value. The cost property
func (m *OrdersItemTransactionsGetResponse_transactions_orders_order_items) SetCost(value *float64) {
	m.cost = value
}

// SetDescription sets the description property value. The description property
func (m *OrdersItemTransactionsGetResponse_transactions_orders_order_items) SetDescription(value *string) {
	m.description = value
}

// SetDiscount sets the discount property value. The discount property
func (m *OrdersItemTransactionsGetResponse_transactions_orders_order_items) SetDiscount(value *float64) {
	m.discount = value
}

// SetId sets the id property value. The id property
func (m *OrdersItemTransactionsGetResponse_transactions_orders_order_items) SetId(value *int64) {
	m.id = value
}

// SetJobRecurringId sets the jobRecurringId property value. The jobRecurringId property
func (m *OrdersItemTransactionsGetResponse_transactions_orders_order_items) SetJobRecurringId(value *int64) {
	m.jobRecurringId = value
}

// SetName sets the name property value. The name property
func (m *OrdersItemTransactionsGetResponse_transactions_orders_order_items) SetName(value *string) {
	m.name = value
}

// SetNotes sets the notes property value. The notes property
func (m *OrdersItemTransactionsGetResponse_transactions_orders_order_items) SetNotes(value *string) {
	m.notes = value
}

// SetPrice sets the price property value. The price property
func (m *OrdersItemTransactionsGetResponse_transactions_orders_order_items) SetPrice(value *float64) {
	m.price = value
}

// SetProduct sets the product property value. The product property
func (m *OrdersItemTransactionsGetResponse_transactions_orders_order_items) SetProduct(value OrdersItemTransactionsGetResponse_transactions_orders_order_items_productable) {
	m.product = value
}

// SetQuantity sets the quantity property value. The quantity property
func (m *OrdersItemTransactionsGetResponse_transactions_orders_order_items) SetQuantity(value *int32) {
	m.quantity = value
}

// SetSpecialAmount sets the specialAmount property value. The specialAmount property
func (m *OrdersItemTransactionsGetResponse_transactions_orders_order_items) SetSpecialAmount(value *float64) {
	m.specialAmount = value
}

// SetSpecialId sets the specialId property value. The specialId property
func (m *OrdersItemTransactionsGetResponse_transactions_orders_order_items) SetSpecialId(value *int64) {
	m.specialId = value
}

// SetSpecialPctOrAmt sets the specialPctOrAmt property value. The specialPctOrAmt property
func (m *OrdersItemTransactionsGetResponse_transactions_orders_order_items) SetSpecialPctOrAmt(value *int32) {
	m.specialPctOrAmt = value
}

// SetSubscriptionPlan sets the subscriptionPlan property value. The subscriptionPlan property
func (m *OrdersItemTransactionsGetResponse_transactions_orders_order_items) SetSubscriptionPlan(value OrdersItemTransactionsGetResponse_transactions_orders_order_items_subscriptionPlanable) {
	m.subscriptionPlan = value
}

// SetTypeEscaped sets the type property value. The type property
func (m *OrdersItemTransactionsGetResponse_transactions_orders_order_items) SetTypeEscaped(value *string) {
	m.typeEscaped = value
}

type OrdersItemTransactionsGetResponse_transactions_orders_order_itemsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCost() *float64
	GetDescription() *string
	GetDiscount() *float64
	GetId() *int64
	GetJobRecurringId() *int64
	GetName() *string
	GetNotes() *string
	GetPrice() *float64
	GetProduct() OrdersItemTransactionsGetResponse_transactions_orders_order_items_productable
	GetQuantity() *int32
	GetSpecialAmount() *float64
	GetSpecialId() *int64
	GetSpecialPctOrAmt() *int32
	GetSubscriptionPlan() OrdersItemTransactionsGetResponse_transactions_orders_order_items_subscriptionPlanable
	GetTypeEscaped() *string
	SetCost(value *float64)
	SetDescription(value *string)
	SetDiscount(value *float64)
	SetId(value *int64)
	SetJobRecurringId(value *int64)
	SetName(value *string)
	SetNotes(value *string)
	SetPrice(value *float64)
	SetProduct(value OrdersItemTransactionsGetResponse_transactions_orders_order_items_productable)
	SetQuantity(value *int32)
	SetSpecialAmount(value *float64)
	SetSpecialId(value *int64)
	SetSpecialPctOrAmt(value *int32)
	SetSubscriptionPlan(value OrdersItemTransactionsGetResponse_transactions_orders_order_items_subscriptionPlanable)
	SetTypeEscaped(value *string)
}
