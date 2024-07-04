package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OrdersItemItemsPostResponse_orderItemTaxes struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The amount property
	amount *float64
	// The id property
	id *int64
	// The name property
	name *string
	// The orderItemId property
	orderItemId *int64
	// The rate property
	rate *float64
	// The taxTemplateId property
	taxTemplateId *int64
}

// NewOrdersItemItemsPostResponse_orderItemTaxes instantiates a new OrdersItemItemsPostResponse_orderItemTaxes and sets the default values.
func NewOrdersItemItemsPostResponse_orderItemTaxes() *OrdersItemItemsPostResponse_orderItemTaxes {
	m := &OrdersItemItemsPostResponse_orderItemTaxes{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateOrdersItemItemsPostResponse_orderItemTaxesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrdersItemItemsPostResponse_orderItemTaxesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewOrdersItemItemsPostResponse_orderItemTaxes(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OrdersItemItemsPostResponse_orderItemTaxes) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAmount gets the amount property value. The amount property
// returns a *float64 when successful
func (m *OrdersItemItemsPostResponse_orderItemTaxes) GetAmount() *float64 {
	return m.amount
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrdersItemItemsPostResponse_orderItemTaxes) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
	res["orderItemId"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetOrderItemId(val)
		}
		return nil
	}
	res["rate"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetFloat64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRate(val)
		}
		return nil
	}
	res["taxTemplateId"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTaxTemplateId(val)
		}
		return nil
	}
	return res
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *OrdersItemItemsPostResponse_orderItemTaxes) GetId() *int64 {
	return m.id
}

// GetName gets the name property value. The name property
// returns a *string when successful
func (m *OrdersItemItemsPostResponse_orderItemTaxes) GetName() *string {
	return m.name
}

// GetOrderItemId gets the orderItemId property value. The orderItemId property
// returns a *int64 when successful
func (m *OrdersItemItemsPostResponse_orderItemTaxes) GetOrderItemId() *int64 {
	return m.orderItemId
}

// GetRate gets the rate property value. The rate property
// returns a *float64 when successful
func (m *OrdersItemItemsPostResponse_orderItemTaxes) GetRate() *float64 {
	return m.rate
}

// GetTaxTemplateId gets the taxTemplateId property value. The taxTemplateId property
// returns a *int64 when successful
func (m *OrdersItemItemsPostResponse_orderItemTaxes) GetTaxTemplateId() *int64 {
	return m.taxTemplateId
}

// Serialize serializes information the current object
func (m *OrdersItemItemsPostResponse_orderItemTaxes) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteFloat64Value("amount", m.GetAmount())
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
		err := writer.WriteStringValue("name", m.GetName())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("orderItemId", m.GetOrderItemId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteFloat64Value("rate", m.GetRate())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("taxTemplateId", m.GetTaxTemplateId())
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
func (m *OrdersItemItemsPostResponse_orderItemTaxes) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAmount sets the amount property value. The amount property
func (m *OrdersItemItemsPostResponse_orderItemTaxes) SetAmount(value *float64) {
	m.amount = value
}

// SetId sets the id property value. The id property
func (m *OrdersItemItemsPostResponse_orderItemTaxes) SetId(value *int64) {
	m.id = value
}

// SetName sets the name property value. The name property
func (m *OrdersItemItemsPostResponse_orderItemTaxes) SetName(value *string) {
	m.name = value
}

// SetOrderItemId sets the orderItemId property value. The orderItemId property
func (m *OrdersItemItemsPostResponse_orderItemTaxes) SetOrderItemId(value *int64) {
	m.orderItemId = value
}

// SetRate sets the rate property value. The rate property
func (m *OrdersItemItemsPostResponse_orderItemTaxes) SetRate(value *float64) {
	m.rate = value
}

// SetTaxTemplateId sets the taxTemplateId property value. The taxTemplateId property
func (m *OrdersItemItemsPostResponse_orderItemTaxes) SetTaxTemplateId(value *int64) {
	m.taxTemplateId = value
}

type OrdersItemItemsPostResponse_orderItemTaxesable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAmount() *float64
	GetId() *int64
	GetName() *string
	GetOrderItemId() *int64
	GetRate() *float64
	GetTaxTemplateId() *int64
	SetAmount(value *float64)
	SetId(value *int64)
	SetName(value *string)
	SetOrderItemId(value *int64)
	SetRate(value *float64)
	SetTaxTemplateId(value *int64)
}
