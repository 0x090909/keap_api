package v2

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SalesRequestBuilder builds and executes requests for operations under \v2\sales
type SalesRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// NewSalesRequestBuilderInternal instantiates a new SalesRequestBuilder and sets the default values.
func NewSalesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *SalesRequestBuilder {
	m := &SalesRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/sales", pathParameters),
	}
	return m
}

// NewSalesRequestBuilder instantiates a new SalesRequestBuilder and sets the default values.
func NewSalesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *SalesRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewSalesRequestBuilderInternal(urlParams, requestAdapter)
}

// Merchants the merchants property
// returns a *SalesMerchantsRequestBuilder when successful
func (m *SalesRequestBuilder) Merchants() *SalesMerchantsRequestBuilder {
	return NewSalesMerchantsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
