package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OrdersItemItemsPostRequestBody struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The description property
	description *string
	// Overridable price of the product, if not specified, the default will be used. Must be greater than or equal to 0.
	price *string
	// The id of the product to be added to the order.
	product_id *int64
	// Quantity must be greater than or equal to 1
	quantity *int32
}

// NewOrdersItemItemsPostRequestBody instantiates a new OrdersItemItemsPostRequestBody and sets the default values.
func NewOrdersItemItemsPostRequestBody() *OrdersItemItemsPostRequestBody {
	m := &OrdersItemItemsPostRequestBody{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateOrdersItemItemsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrdersItemItemsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewOrdersItemItemsPostRequestBody(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OrdersItemItemsPostRequestBody) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *OrdersItemItemsPostRequestBody) GetDescription() *string {
	return m.description
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrdersItemItemsPostRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
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
	res["price"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPrice(val)
		}
		return nil
	}
	res["product_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetProductId(val)
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
	return res
}

// GetPrice gets the price property value. Overridable price of the product, if not specified, the default will be used. Must be greater than or equal to 0.
// returns a *string when successful
func (m *OrdersItemItemsPostRequestBody) GetPrice() *string {
	return m.price
}

// GetProductId gets the product_id property value. The id of the product to be added to the order.
// returns a *int64 when successful
func (m *OrdersItemItemsPostRequestBody) GetProductId() *int64 {
	return m.product_id
}

// GetQuantity gets the quantity property value. Quantity must be greater than or equal to 1
// returns a *int32 when successful
func (m *OrdersItemItemsPostRequestBody) GetQuantity() *int32 {
	return m.quantity
}

// Serialize serializes information the current object
func (m *OrdersItemItemsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("description", m.GetDescription())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("price", m.GetPrice())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("product_id", m.GetProductId())
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
		err := writer.WriteAdditionalData(m.GetAdditionalData())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OrdersItemItemsPostRequestBody) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetDescription sets the description property value. The description property
func (m *OrdersItemItemsPostRequestBody) SetDescription(value *string) {
	m.description = value
}

// SetPrice sets the price property value. Overridable price of the product, if not specified, the default will be used. Must be greater than or equal to 0.
func (m *OrdersItemItemsPostRequestBody) SetPrice(value *string) {
	m.price = value
}

// SetProductId sets the product_id property value. The id of the product to be added to the order.
func (m *OrdersItemItemsPostRequestBody) SetProductId(value *int64) {
	m.product_id = value
}

// SetQuantity sets the quantity property value. Quantity must be greater than or equal to 1
func (m *OrdersItemItemsPostRequestBody) SetQuantity(value *int32) {
	m.quantity = value
}

type OrdersItemItemsPostRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetDescription() *string
	GetPrice() *string
	GetProductId() *int64
	GetQuantity() *int32
	SetDescription(value *string)
	SetPrice(value *string)
	SetProductId(value *int64)
	SetQuantity(value *int32)
}
