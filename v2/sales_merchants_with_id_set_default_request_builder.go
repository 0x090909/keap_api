package v2

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
)

// SalesMerchantsWithIdSetDefaultRequestBuilder builds and executes requests for operations under \v2\sales\merchants\{id}:setDefault
type SalesMerchantsWithIdSetDefaultRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// SalesMerchantsWithIdSetDefaultRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SalesMerchantsWithIdSetDefaultRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewSalesMerchantsWithIdSetDefaultRequestBuilderInternal instantiates a new SalesMerchantsWithIdSetDefaultRequestBuilder and sets the default values.
func NewSalesMerchantsWithIdSetDefaultRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, id *int64) *SalesMerchantsWithIdSetDefaultRequestBuilder {
	m := &SalesMerchantsWithIdSetDefaultRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/sales/merchants/{id}:setDefault", pathParameters),
	}
	if id != nil {
		m.BaseRequestBuilder.PathParameters["id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(*id, 10)
	}
	return m
}

// NewSalesMerchantsWithIdSetDefaultRequestBuilder instantiates a new SalesMerchantsWithIdSetDefaultRequestBuilder and sets the default values.
func NewSalesMerchantsWithIdSetDefaultRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *SalesMerchantsWithIdSetDefaultRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewSalesMerchantsWithIdSetDefaultRequestBuilderInternal(urlParams, requestAdapter, nil)
}

// Post sets the specified Merchant Account as the default Merchant Account.
// returns a []byte when successful
// returns a SalesMerchantsWithIdSetDefault401Error error when the service returns a 401 status code
// returns a SalesMerchantsWithIdSetDefault403Error error when the service returns a 403 status code
func (m *SalesMerchantsWithIdSetDefaultRequestBuilder) Post(ctx context.Context, requestConfiguration *SalesMerchantsWithIdSetDefaultRequestBuilderPostRequestConfiguration) ([]byte, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateSalesMerchantsWithIdSetDefault401ErrorFromDiscriminatorValue,
		"403": CreateSalesMerchantsWithIdSetDefault403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.([]byte), nil
}

// ToPostRequestInformation sets the specified Merchant Account as the default Merchant Account.
// returns a *RequestInformation when successful
func (m *SalesMerchantsWithIdSetDefaultRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *SalesMerchantsWithIdSetDefaultRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *SalesMerchantsWithIdSetDefaultRequestBuilder when successful
func (m *SalesMerchantsWithIdSetDefaultRequestBuilder) WithUrl(rawUrl string) *SalesMerchantsWithIdSetDefaultRequestBuilder {
	return NewSalesMerchantsWithIdSetDefaultRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
