package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OrdersGetResponse_orders_order_items_subscriptionPlan struct {
	// The active property
	active *bool
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The cycle property
	cycle *int32
	// The frequency property
	frequency *int32
	// The id property
	id *int64
	// The number_of_cycles property
	number_of_cycles *int32
	// The plan_price property
	plan_price *float64
	// The subscription_plan_index property
	subscription_plan_index *int32
	// The subscription_plan_name property
	subscription_plan_name *string
}

// NewOrdersGetResponse_orders_order_items_subscriptionPlan instantiates a new OrdersGetResponse_orders_order_items_subscriptionPlan and sets the default values.
func NewOrdersGetResponse_orders_order_items_subscriptionPlan() *OrdersGetResponse_orders_order_items_subscriptionPlan {
	m := &OrdersGetResponse_orders_order_items_subscriptionPlan{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateOrdersGetResponse_orders_order_items_subscriptionPlanFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrdersGetResponse_orders_order_items_subscriptionPlanFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewOrdersGetResponse_orders_order_items_subscriptionPlan(), nil
}

// GetActive gets the active property value. The active property
// returns a *bool when successful
func (m *OrdersGetResponse_orders_order_items_subscriptionPlan) GetActive() *bool {
	return m.active
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OrdersGetResponse_orders_order_items_subscriptionPlan) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCycle gets the cycle property value. The cycle property
// returns a *int32 when successful
func (m *OrdersGetResponse_orders_order_items_subscriptionPlan) GetCycle() *int32 {
	return m.cycle
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrdersGetResponse_orders_order_items_subscriptionPlan) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
	res["cycle"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCycle(val)
		}
		return nil
	}
	res["frequency"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFrequency(val)
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
	res["number_of_cycles"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNumberOfCycles(val)
		}
		return nil
	}
	res["plan_price"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetFloat64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPlanPrice(val)
		}
		return nil
	}
	res["subscription_plan_index"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSubscriptionPlanIndex(val)
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

// GetFrequency gets the frequency property value. The frequency property
// returns a *int32 when successful
func (m *OrdersGetResponse_orders_order_items_subscriptionPlan) GetFrequency() *int32 {
	return m.frequency
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *OrdersGetResponse_orders_order_items_subscriptionPlan) GetId() *int64 {
	return m.id
}

// GetNumberOfCycles gets the number_of_cycles property value. The number_of_cycles property
// returns a *int32 when successful
func (m *OrdersGetResponse_orders_order_items_subscriptionPlan) GetNumberOfCycles() *int32 {
	return m.number_of_cycles
}

// GetPlanPrice gets the plan_price property value. The plan_price property
// returns a *float64 when successful
func (m *OrdersGetResponse_orders_order_items_subscriptionPlan) GetPlanPrice() *float64 {
	return m.plan_price
}

// GetSubscriptionPlanIndex gets the subscription_plan_index property value. The subscription_plan_index property
// returns a *int32 when successful
func (m *OrdersGetResponse_orders_order_items_subscriptionPlan) GetSubscriptionPlanIndex() *int32 {
	return m.subscription_plan_index
}

// GetSubscriptionPlanName gets the subscription_plan_name property value. The subscription_plan_name property
// returns a *string when successful
func (m *OrdersGetResponse_orders_order_items_subscriptionPlan) GetSubscriptionPlanName() *string {
	return m.subscription_plan_name
}

// Serialize serializes information the current object
func (m *OrdersGetResponse_orders_order_items_subscriptionPlan) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteBoolValue("active", m.GetActive())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("cycle", m.GetCycle())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("frequency", m.GetFrequency())
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
		err := writer.WriteInt32Value("number_of_cycles", m.GetNumberOfCycles())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteFloat64Value("plan_price", m.GetPlanPrice())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("subscription_plan_index", m.GetSubscriptionPlanIndex())
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

// SetActive sets the active property value. The active property
func (m *OrdersGetResponse_orders_order_items_subscriptionPlan) SetActive(value *bool) {
	m.active = value
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OrdersGetResponse_orders_order_items_subscriptionPlan) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCycle sets the cycle property value. The cycle property
func (m *OrdersGetResponse_orders_order_items_subscriptionPlan) SetCycle(value *int32) {
	m.cycle = value
}

// SetFrequency sets the frequency property value. The frequency property
func (m *OrdersGetResponse_orders_order_items_subscriptionPlan) SetFrequency(value *int32) {
	m.frequency = value
}

// SetId sets the id property value. The id property
func (m *OrdersGetResponse_orders_order_items_subscriptionPlan) SetId(value *int64) {
	m.id = value
}

// SetNumberOfCycles sets the number_of_cycles property value. The number_of_cycles property
func (m *OrdersGetResponse_orders_order_items_subscriptionPlan) SetNumberOfCycles(value *int32) {
	m.number_of_cycles = value
}

// SetPlanPrice sets the plan_price property value. The plan_price property
func (m *OrdersGetResponse_orders_order_items_subscriptionPlan) SetPlanPrice(value *float64) {
	m.plan_price = value
}

// SetSubscriptionPlanIndex sets the subscription_plan_index property value. The subscription_plan_index property
func (m *OrdersGetResponse_orders_order_items_subscriptionPlan) SetSubscriptionPlanIndex(value *int32) {
	m.subscription_plan_index = value
}

// SetSubscriptionPlanName sets the subscription_plan_name property value. The subscription_plan_name property
func (m *OrdersGetResponse_orders_order_items_subscriptionPlan) SetSubscriptionPlanName(value *string) {
	m.subscription_plan_name = value
}

type OrdersGetResponse_orders_order_items_subscriptionPlanable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetActive() *bool
	GetCycle() *int32
	GetFrequency() *int32
	GetId() *int64
	GetNumberOfCycles() *int32
	GetPlanPrice() *float64
	GetSubscriptionPlanIndex() *int32
	GetSubscriptionPlanName() *string
	SetActive(value *bool)
	SetCycle(value *int32)
	SetFrequency(value *int32)
	SetId(value *int64)
	SetNumberOfCycles(value *int32)
	SetPlanPrice(value *float64)
	SetSubscriptionPlanIndex(value *int32)
	SetSubscriptionPlanName(value *string)
}
