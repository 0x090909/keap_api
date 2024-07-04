package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ProductsItemWithProductGetResponseable instead.
type ProductsItemWithProductResponse struct {
	ProductsItemWithProductGetResponse
}

// NewProductsItemWithProductResponse instantiates a new ProductsItemWithProductResponse and sets the default values.
func NewProductsItemWithProductResponse() *ProductsItemWithProductResponse {
	m := &ProductsItemWithProductResponse{
		ProductsItemWithProductGetResponse: *NewProductsItemWithProductGetResponse(),
	}
	return m
}

// CreateProductsItemWithProductResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProductsItemWithProductResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewProductsItemWithProductResponse(), nil
}

// Deprecated: This class is obsolete. Use ProductsItemWithProductGetResponseable instead.
type ProductsItemWithProductResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	ProductsItemWithProductGetResponseable
}
