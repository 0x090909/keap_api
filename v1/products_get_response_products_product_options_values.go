package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProductsGetResponse_products_product_options_values struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The id property
	id *int64
	// The index property
	index *int64
	// The is_default property
	is_default *bool
	// The label property
	label *string
	// The price_adjustment property
	price_adjustment *float64
	// The sku property
	sku *string
}

// NewProductsGetResponse_products_product_options_values instantiates a new ProductsGetResponse_products_product_options_values and sets the default values.
func NewProductsGetResponse_products_product_options_values() *ProductsGetResponse_products_product_options_values {
	m := &ProductsGetResponse_products_product_options_values{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateProductsGetResponse_products_product_options_valuesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProductsGetResponse_products_product_options_valuesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewProductsGetResponse_products_product_options_values(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ProductsGetResponse_products_product_options_values) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProductsGetResponse_products_product_options_values) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
	res["index"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetIndex(val)
		}
		return nil
	}
	res["is_default"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetIsDefault(val)
		}
		return nil
	}
	res["label"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLabel(val)
		}
		return nil
	}
	res["price_adjustment"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetFloat64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPriceAdjustment(val)
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
	return res
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *ProductsGetResponse_products_product_options_values) GetId() *int64 {
	return m.id
}

// GetIndex gets the index property value. The index property
// returns a *int64 when successful
func (m *ProductsGetResponse_products_product_options_values) GetIndex() *int64 {
	return m.index
}

// GetIsDefault gets the is_default property value. The is_default property
// returns a *bool when successful
func (m *ProductsGetResponse_products_product_options_values) GetIsDefault() *bool {
	return m.is_default
}

// GetLabel gets the label property value. The label property
// returns a *string when successful
func (m *ProductsGetResponse_products_product_options_values) GetLabel() *string {
	return m.label
}

// GetPriceAdjustment gets the price_adjustment property value. The price_adjustment property
// returns a *float64 when successful
func (m *ProductsGetResponse_products_product_options_values) GetPriceAdjustment() *float64 {
	return m.price_adjustment
}

// GetSku gets the sku property value. The sku property
// returns a *string when successful
func (m *ProductsGetResponse_products_product_options_values) GetSku() *string {
	return m.sku
}

// Serialize serializes information the current object
func (m *ProductsGetResponse_products_product_options_values) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt64Value("id", m.GetId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("index", m.GetIndex())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("is_default", m.GetIsDefault())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("label", m.GetLabel())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteFloat64Value("price_adjustment", m.GetPriceAdjustment())
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
		err := writer.WriteAdditionalData(m.GetAdditionalData())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ProductsGetResponse_products_product_options_values) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetId sets the id property value. The id property
func (m *ProductsGetResponse_products_product_options_values) SetId(value *int64) {
	m.id = value
}

// SetIndex sets the index property value. The index property
func (m *ProductsGetResponse_products_product_options_values) SetIndex(value *int64) {
	m.index = value
}

// SetIsDefault sets the is_default property value. The is_default property
func (m *ProductsGetResponse_products_product_options_values) SetIsDefault(value *bool) {
	m.is_default = value
}

// SetLabel sets the label property value. The label property
func (m *ProductsGetResponse_products_product_options_values) SetLabel(value *string) {
	m.label = value
}

// SetPriceAdjustment sets the price_adjustment property value. The price_adjustment property
func (m *ProductsGetResponse_products_product_options_values) SetPriceAdjustment(value *float64) {
	m.price_adjustment = value
}

// SetSku sets the sku property value. The sku property
func (m *ProductsGetResponse_products_product_options_values) SetSku(value *string) {
	m.sku = value
}

type ProductsGetResponse_products_product_options_valuesable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetId() *int64
	GetIndex() *int64
	GetIsDefault() *bool
	GetLabel() *string
	GetPriceAdjustment() *float64
	GetSku() *string
	SetId(value *int64)
	SetIndex(value *int64)
	SetIsDefault(value *bool)
	SetLabel(value *string)
	SetPriceAdjustment(value *float64)
	SetSku(value *string)
}
