package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TransactionsItemWithTransactionGetResponse_orders_order_items_product struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The description property
	description *string
	// The id property
	id *int64
	// The name property
	name *string
	// The shippable property
	shippable *bool
	// The sku property
	sku *string
	// The taxable property
	taxable *bool
}

// NewTransactionsItemWithTransactionGetResponse_orders_order_items_product instantiates a new TransactionsItemWithTransactionGetResponse_orders_order_items_product and sets the default values.
func NewTransactionsItemWithTransactionGetResponse_orders_order_items_product() *TransactionsItemWithTransactionGetResponse_orders_order_items_product {
	m := &TransactionsItemWithTransactionGetResponse_orders_order_items_product{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateTransactionsItemWithTransactionGetResponse_orders_order_items_productFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionsItemWithTransactionGetResponse_orders_order_items_productFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTransactionsItemWithTransactionGetResponse_orders_order_items_product(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransactionsItemWithTransactionGetResponse_orders_order_items_product) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *TransactionsItemWithTransactionGetResponse_orders_order_items_product) GetDescription() *string {
	return m.description
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransactionsItemWithTransactionGetResponse_orders_order_items_product) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
	res["shippable"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetShippable(val)
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
	res["taxable"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTaxable(val)
		}
		return nil
	}
	return res
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *TransactionsItemWithTransactionGetResponse_orders_order_items_product) GetId() *int64 {
	return m.id
}

// GetName gets the name property value. The name property
// returns a *string when successful
func (m *TransactionsItemWithTransactionGetResponse_orders_order_items_product) GetName() *string {
	return m.name
}

// GetShippable gets the shippable property value. The shippable property
// returns a *bool when successful
func (m *TransactionsItemWithTransactionGetResponse_orders_order_items_product) GetShippable() *bool {
	return m.shippable
}

// GetSku gets the sku property value. The sku property
// returns a *string when successful
func (m *TransactionsItemWithTransactionGetResponse_orders_order_items_product) GetSku() *string {
	return m.sku
}

// GetTaxable gets the taxable property value. The taxable property
// returns a *bool when successful
func (m *TransactionsItemWithTransactionGetResponse_orders_order_items_product) GetTaxable() *bool {
	return m.taxable
}

// Serialize serializes information the current object
func (m *TransactionsItemWithTransactionGetResponse_orders_order_items_product) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("description", m.GetDescription())
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
		err := writer.WriteBoolValue("shippable", m.GetShippable())
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
		err := writer.WriteBoolValue("taxable", m.GetTaxable())
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
func (m *TransactionsItemWithTransactionGetResponse_orders_order_items_product) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetDescription sets the description property value. The description property
func (m *TransactionsItemWithTransactionGetResponse_orders_order_items_product) SetDescription(value *string) {
	m.description = value
}

// SetId sets the id property value. The id property
func (m *TransactionsItemWithTransactionGetResponse_orders_order_items_product) SetId(value *int64) {
	m.id = value
}

// SetName sets the name property value. The name property
func (m *TransactionsItemWithTransactionGetResponse_orders_order_items_product) SetName(value *string) {
	m.name = value
}

// SetShippable sets the shippable property value. The shippable property
func (m *TransactionsItemWithTransactionGetResponse_orders_order_items_product) SetShippable(value *bool) {
	m.shippable = value
}

// SetSku sets the sku property value. The sku property
func (m *TransactionsItemWithTransactionGetResponse_orders_order_items_product) SetSku(value *string) {
	m.sku = value
}

// SetTaxable sets the taxable property value. The taxable property
func (m *TransactionsItemWithTransactionGetResponse_orders_order_items_product) SetTaxable(value *bool) {
	m.taxable = value
}

type TransactionsItemWithTransactionGetResponse_orders_order_items_productable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetDescription() *string
	GetId() *int64
	GetName() *string
	GetShippable() *bool
	GetSku() *string
	GetTaxable() *bool
	SetDescription(value *string)
	SetId(value *int64)
	SetName(value *string)
	SetShippable(value *bool)
	SetSku(value *string)
	SetTaxable(value *bool)
}
