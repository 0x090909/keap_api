package v2

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SalesMerchantsRequestBuilder builds and executes requests for operations under \v2\sales\merchants
type SalesMerchantsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// NewSalesMerchantsRequestBuilderInternal instantiates a new SalesMerchantsRequestBuilder and sets the default values.
func NewSalesMerchantsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *SalesMerchantsRequestBuilder {
	m := &SalesMerchantsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/sales/merchants", pathParameters),
	}
	return m
}

// NewSalesMerchantsRequestBuilder instantiates a new SalesMerchantsRequestBuilder and sets the default values.
func NewSalesMerchantsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *SalesMerchantsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewSalesMerchantsRequestBuilderInternal(urlParams, requestAdapter)
}

// WithIdSetDefault builds and executes requests for operations under \v2\sales\merchants\{id}:setDefault
// returns a *SalesMerchantsWithIdSetDefaultRequestBuilder when successful
func (m *SalesMerchantsRequestBuilder) WithIdSetDefault(id *int64) *SalesMerchantsWithIdSetDefaultRequestBuilder {
	return NewSalesMerchantsWithIdSetDefaultRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, id)
}
