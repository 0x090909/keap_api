package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ProductsItemSubscriptionsPostResponseable instead.
type ProductsItemSubscriptionsResponse struct {
	ProductsItemSubscriptionsPostResponse
}

// NewProductsItemSubscriptionsResponse instantiates a new ProductsItemSubscriptionsResponse and sets the default values.
func NewProductsItemSubscriptionsResponse() *ProductsItemSubscriptionsResponse {
	m := &ProductsItemSubscriptionsResponse{
		ProductsItemSubscriptionsPostResponse: *NewProductsItemSubscriptionsPostResponse(),
	}
	return m
}

// CreateProductsItemSubscriptionsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProductsItemSubscriptionsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewProductsItemSubscriptionsResponse(), nil
}

// Deprecated: This class is obsolete. Use ProductsItemSubscriptionsPostResponseable instead.
type ProductsItemSubscriptionsResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	ProductsItemSubscriptionsPostResponseable
}
