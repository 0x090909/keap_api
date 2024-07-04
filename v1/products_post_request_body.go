package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProductsPostRequestBody struct {
	// The active property
	active *bool
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The product_desc property
	product_desc *string
	// The product_name property
	product_name *string
	// The product_price property
	product_price *float64
	// The product_short_desc property
	product_short_desc *string
	// The sku property
	sku *string
	// The subscription_only property
	subscription_only *bool
}

// NewProductsPostRequestBody instantiates a new ProductsPostRequestBody and sets the default values.
func NewProductsPostRequestBody() *ProductsPostRequestBody {
	m := &ProductsPostRequestBody{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateProductsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProductsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewProductsPostRequestBody(), nil
}

// GetActive gets the active property value. The active property
// returns a *bool when successful
func (m *ProductsPostRequestBody) GetActive() *bool {
	return m.active
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ProductsPostRequestBody) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProductsPostRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
	return res
}

// GetProductDesc gets the product_desc property value. The product_desc property
// returns a *string when successful
func (m *ProductsPostRequestBody) GetProductDesc() *string {
	return m.product_desc
}

// GetProductName gets the product_name property value. The product_name property
// returns a *string when successful
func (m *ProductsPostRequestBody) GetProductName() *string {
	return m.product_name
}

// GetProductPrice gets the product_price property value. The product_price property
// returns a *float64 when successful
func (m *ProductsPostRequestBody) GetProductPrice() *float64 {
	return m.product_price
}

// GetProductShortDesc gets the product_short_desc property value. The product_short_desc property
// returns a *string when successful
func (m *ProductsPostRequestBody) GetProductShortDesc() *string {
	return m.product_short_desc
}

// GetSku gets the sku property value. The sku property
// returns a *string when successful
func (m *ProductsPostRequestBody) GetSku() *string {
	return m.sku
}

// GetSubscriptionOnly gets the subscription_only property value. The subscription_only property
// returns a *bool when successful
func (m *ProductsPostRequestBody) GetSubscriptionOnly() *bool {
	return m.subscription_only
}

// Serialize serializes information the current object
func (m *ProductsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteBoolValue("active", m.GetActive())
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
		err := writer.WriteBoolValue("subscription_only", m.GetSubscriptionOnly())
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
func (m *ProductsPostRequestBody) SetActive(value *bool) {
	m.active = value
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ProductsPostRequestBody) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetProductDesc sets the product_desc property value. The product_desc property
func (m *ProductsPostRequestBody) SetProductDesc(value *string) {
	m.product_desc = value
}

// SetProductName sets the product_name property value. The product_name property
func (m *ProductsPostRequestBody) SetProductName(value *string) {
	m.product_name = value
}

// SetProductPrice sets the product_price property value. The product_price property
func (m *ProductsPostRequestBody) SetProductPrice(value *float64) {
	m.product_price = value
}

// SetProductShortDesc sets the product_short_desc property value. The product_short_desc property
func (m *ProductsPostRequestBody) SetProductShortDesc(value *string) {
	m.product_short_desc = value
}

// SetSku sets the sku property value. The sku property
func (m *ProductsPostRequestBody) SetSku(value *string) {
	m.sku = value
}

// SetSubscriptionOnly sets the subscription_only property value. The subscription_only property
func (m *ProductsPostRequestBody) SetSubscriptionOnly(value *bool) {
	m.subscription_only = value
}

type ProductsPostRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetActive() *bool
	GetProductDesc() *string
	GetProductName() *string
	GetProductPrice() *float64
	GetProductShortDesc() *string
	GetSku() *string
	GetSubscriptionOnly() *bool
	SetActive(value *bool)
	SetProductDesc(value *string)
	SetProductName(value *string)
	SetProductPrice(value *float64)
	SetProductShortDesc(value *string)
	SetSku(value *string)
	SetSubscriptionOnly(value *bool)
}
