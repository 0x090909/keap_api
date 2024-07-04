package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProductsSyncGetResponse_product_statuses struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The product property
	product ProductsSyncGetResponse_product_statuses_productable
}

// NewProductsSyncGetResponse_product_statuses instantiates a new ProductsSyncGetResponse_product_statuses and sets the default values.
func NewProductsSyncGetResponse_product_statuses() *ProductsSyncGetResponse_product_statuses {
	m := &ProductsSyncGetResponse_product_statuses{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateProductsSyncGetResponse_product_statusesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProductsSyncGetResponse_product_statusesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewProductsSyncGetResponse_product_statuses(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ProductsSyncGetResponse_product_statuses) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProductsSyncGetResponse_product_statuses) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["product"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateProductsSyncGetResponse_product_statuses_productFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetProduct(val.(ProductsSyncGetResponse_product_statuses_productable))
		}
		return nil
	}
	return res
}

// GetProduct gets the product property value. The product property
// returns a ProductsSyncGetResponse_product_statuses_productable when successful
func (m *ProductsSyncGetResponse_product_statuses) GetProduct() ProductsSyncGetResponse_product_statuses_productable {
	return m.product
}

// Serialize serializes information the current object
func (m *ProductsSyncGetResponse_product_statuses) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("product", m.GetProduct())
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
func (m *ProductsSyncGetResponse_product_statuses) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetProduct sets the product property value. The product property
func (m *ProductsSyncGetResponse_product_statuses) SetProduct(value ProductsSyncGetResponse_product_statuses_productable) {
	m.product = value
}

type ProductsSyncGetResponse_product_statusesable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetProduct() ProductsSyncGetResponse_product_statuses_productable
	SetProduct(value ProductsSyncGetResponse_product_statuses_productable)
}
