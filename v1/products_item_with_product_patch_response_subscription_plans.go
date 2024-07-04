package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProductsItemWithProductPatchResponse_subscription_plans struct {
	// The active property
	active *bool
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
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
	// The url property
	url *string
}

// NewProductsItemWithProductPatchResponse_subscription_plans instantiates a new ProductsItemWithProductPatchResponse_subscription_plans and sets the default values.
func NewProductsItemWithProductPatchResponse_subscription_plans() *ProductsItemWithProductPatchResponse_subscription_plans {
	m := &ProductsItemWithProductPatchResponse_subscription_plans{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateProductsItemWithProductPatchResponse_subscription_plansFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProductsItemWithProductPatchResponse_subscription_plansFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewProductsItemWithProductPatchResponse_subscription_plans(), nil
}

// GetActive gets the active property value. The active property
// returns a *bool when successful
func (m *ProductsItemWithProductPatchResponse_subscription_plans) GetActive() *bool {
	return m.active
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ProductsItemWithProductPatchResponse_subscription_plans) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProductsItemWithProductPatchResponse_subscription_plans) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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

// GetFrequency gets the frequency property value. The frequency property
// returns a *int32 when successful
func (m *ProductsItemWithProductPatchResponse_subscription_plans) GetFrequency() *int32 {
	return m.frequency
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *ProductsItemWithProductPatchResponse_subscription_plans) GetId() *int64 {
	return m.id
}

// GetNumberOfCycles gets the number_of_cycles property value. The number_of_cycles property
// returns a *int32 when successful
func (m *ProductsItemWithProductPatchResponse_subscription_plans) GetNumberOfCycles() *int32 {
	return m.number_of_cycles
}

// GetPlanPrice gets the plan_price property value. The plan_price property
// returns a *float64 when successful
func (m *ProductsItemWithProductPatchResponse_subscription_plans) GetPlanPrice() *float64 {
	return m.plan_price
}

// GetSubscriptionPlanIndex gets the subscription_plan_index property value. The subscription_plan_index property
// returns a *int32 when successful
func (m *ProductsItemWithProductPatchResponse_subscription_plans) GetSubscriptionPlanIndex() *int32 {
	return m.subscription_plan_index
}

// GetSubscriptionPlanName gets the subscription_plan_name property value. The subscription_plan_name property
// returns a *string when successful
func (m *ProductsItemWithProductPatchResponse_subscription_plans) GetSubscriptionPlanName() *string {
	return m.subscription_plan_name
}

// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *ProductsItemWithProductPatchResponse_subscription_plans) GetUrl() *string {
	return m.url
}

// Serialize serializes information the current object
func (m *ProductsItemWithProductPatchResponse_subscription_plans) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteBoolValue("active", m.GetActive())
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

// SetActive sets the active property value. The active property
func (m *ProductsItemWithProductPatchResponse_subscription_plans) SetActive(value *bool) {
	m.active = value
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ProductsItemWithProductPatchResponse_subscription_plans) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetFrequency sets the frequency property value. The frequency property
func (m *ProductsItemWithProductPatchResponse_subscription_plans) SetFrequency(value *int32) {
	m.frequency = value
}

// SetId sets the id property value. The id property
func (m *ProductsItemWithProductPatchResponse_subscription_plans) SetId(value *int64) {
	m.id = value
}

// SetNumberOfCycles sets the number_of_cycles property value. The number_of_cycles property
func (m *ProductsItemWithProductPatchResponse_subscription_plans) SetNumberOfCycles(value *int32) {
	m.number_of_cycles = value
}

// SetPlanPrice sets the plan_price property value. The plan_price property
func (m *ProductsItemWithProductPatchResponse_subscription_plans) SetPlanPrice(value *float64) {
	m.plan_price = value
}

// SetSubscriptionPlanIndex sets the subscription_plan_index property value. The subscription_plan_index property
func (m *ProductsItemWithProductPatchResponse_subscription_plans) SetSubscriptionPlanIndex(value *int32) {
	m.subscription_plan_index = value
}

// SetSubscriptionPlanName sets the subscription_plan_name property value. The subscription_plan_name property
func (m *ProductsItemWithProductPatchResponse_subscription_plans) SetSubscriptionPlanName(value *string) {
	m.subscription_plan_name = value
}

// SetUrl sets the url property value. The url property
func (m *ProductsItemWithProductPatchResponse_subscription_plans) SetUrl(value *string) {
	m.url = value
}

type ProductsItemWithProductPatchResponse_subscription_plansable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetActive() *bool
	GetFrequency() *int32
	GetId() *int64
	GetNumberOfCycles() *int32
	GetPlanPrice() *float64
	GetSubscriptionPlanIndex() *int32
	GetSubscriptionPlanName() *string
	GetUrl() *string
	SetActive(value *bool)
	SetFrequency(value *int32)
	SetId(value *int64)
	SetNumberOfCycles(value *int32)
	SetPlanPrice(value *float64)
	SetSubscriptionPlanIndex(value *int32)
	SetSubscriptionPlanName(value *string)
	SetUrl(value *string)
}
