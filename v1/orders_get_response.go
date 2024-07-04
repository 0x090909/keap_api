package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OrdersGetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The count property
	count *int32
	// The next property
	next *string
	// The orders property
	orders []OrdersGetResponse_ordersable
	// The previous property
	previous *string
}

// NewOrdersGetResponse instantiates a new OrdersGetResponse and sets the default values.
func NewOrdersGetResponse() *OrdersGetResponse {
	m := &OrdersGetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateOrdersGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrdersGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewOrdersGetResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OrdersGetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCount gets the count property value. The count property
// returns a *int32 when successful
func (m *OrdersGetResponse) GetCount() *int32 {
	return m.count
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrdersGetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["count"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCount(val)
		}
		return nil
	}
	res["next"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNext(val)
		}
		return nil
	}
	res["orders"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateOrdersGetResponse_ordersFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]OrdersGetResponse_ordersable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(OrdersGetResponse_ordersable)
				}
			}
			m.SetOrders(res)
		}
		return nil
	}
	res["previous"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPrevious(val)
		}
		return nil
	}
	return res
}

// GetNext gets the next property value. The next property
// returns a *string when successful
func (m *OrdersGetResponse) GetNext() *string {
	return m.next
}

// GetOrders gets the orders property value. The orders property
// returns a []OrdersGetResponse_ordersable when successful
func (m *OrdersGetResponse) GetOrders() []OrdersGetResponse_ordersable {
	return m.orders
}

// GetPrevious gets the previous property value. The previous property
// returns a *string when successful
func (m *OrdersGetResponse) GetPrevious() *string {
	return m.previous
}

// Serialize serializes information the current object
func (m *OrdersGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt32Value("count", m.GetCount())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("next", m.GetNext())
		if err != nil {
			return err
		}
	}
	if m.GetOrders() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOrders()))
		for i, v := range m.GetOrders() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("orders", cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("previous", m.GetPrevious())
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
func (m *OrdersGetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCount sets the count property value. The count property
func (m *OrdersGetResponse) SetCount(value *int32) {
	m.count = value
}

// SetNext sets the next property value. The next property
func (m *OrdersGetResponse) SetNext(value *string) {
	m.next = value
}

// SetOrders sets the orders property value. The orders property
func (m *OrdersGetResponse) SetOrders(value []OrdersGetResponse_ordersable) {
	m.orders = value
}

// SetPrevious sets the previous property value. The previous property
func (m *OrdersGetResponse) SetPrevious(value *string) {
	m.previous = value
}

type OrdersGetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCount() *int32
	GetNext() *string
	GetOrders() []OrdersGetResponse_ordersable
	GetPrevious() *string
	SetCount(value *int32)
	SetNext(value *string)
	SetOrders(value []OrdersGetResponse_ordersable)
	SetPrevious(value *string)
}
