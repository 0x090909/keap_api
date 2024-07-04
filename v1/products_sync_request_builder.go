package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ProductsSyncRequestBuilder builds and executes requests for operations under \v1\products\sync
type ProductsSyncRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ProductsSyncRequestBuilderGetQueryParameters the Sync endpoint returns a set of products that have been updated or created since the last result set was retrieved, minus any products that have been deleted.
type ProductsSyncRequestBuilderGetQueryParameters struct {
	// Sets a total of items to return
	Limit *int32 `uriparametername:"limit"`
	// Sets a beginning range of items to return
	Offset *int32 `uriparametername:"offset"`
	// sync_token
	Sync_token *string `uriparametername:"sync_token"`
}

// ProductsSyncRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ProductsSyncRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ProductsSyncRequestBuilderGetQueryParameters
}

// NewProductsSyncRequestBuilderInternal instantiates a new ProductsSyncRequestBuilder and sets the default values.
func NewProductsSyncRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ProductsSyncRequestBuilder {
	m := &ProductsSyncRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/products/sync{?limit*,offset*,sync_token*}", pathParameters),
	}
	return m
}

// NewProductsSyncRequestBuilder instantiates a new ProductsSyncRequestBuilder and sets the default values.
func NewProductsSyncRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ProductsSyncRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewProductsSyncRequestBuilderInternal(urlParams, requestAdapter)
}

// Get the Sync endpoint returns a set of products that have been updated or created since the last result set was retrieved, minus any products that have been deleted.
// Deprecated: This method is obsolete. Use GetAsSyncGetResponse instead.
// returns a ProductsSyncResponseable when successful
// returns a ProductsSync401Error error when the service returns a 401 status code
// returns a ProductsSync403Error error when the service returns a 403 status code
// returns a ProductsSync404Error error when the service returns a 404 status code
func (m *ProductsSyncRequestBuilder) Get(ctx context.Context, requestConfiguration *ProductsSyncRequestBuilderGetRequestConfiguration) (ProductsSyncResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateProductsSync401ErrorFromDiscriminatorValue,
		"403": CreateProductsSync403ErrorFromDiscriminatorValue,
		"404": CreateProductsSync404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateProductsSyncResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ProductsSyncResponseable), nil
}

// GetAsSyncGetResponse the Sync endpoint returns a set of products that have been updated or created since the last result set was retrieved, minus any products that have been deleted.
// Deprecated:
// returns a ProductsSyncGetResponseable when successful
// returns a ProductsSync401Error error when the service returns a 401 status code
// returns a ProductsSync403Error error when the service returns a 403 status code
// returns a ProductsSync404Error error when the service returns a 404 status code
func (m *ProductsSyncRequestBuilder) GetAsSyncGetResponse(ctx context.Context, requestConfiguration *ProductsSyncRequestBuilderGetRequestConfiguration) (ProductsSyncGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateProductsSync401ErrorFromDiscriminatorValue,
		"403": CreateProductsSync403ErrorFromDiscriminatorValue,
		"404": CreateProductsSync404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateProductsSyncGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ProductsSyncGetResponseable), nil
}

// ToGetRequestInformation the Sync endpoint returns a set of products that have been updated or created since the last result set was retrieved, minus any products that have been deleted.
// Deprecated:
// returns a *RequestInformation when successful
func (m *ProductsSyncRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ProductsSyncRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		if requestConfiguration.QueryParameters != nil {
			requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
		}
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated:
// returns a *ProductsSyncRequestBuilder when successful
func (m *ProductsSyncRequestBuilder) WithUrl(rawUrl string) *ProductsSyncRequestBuilder {
	return NewProductsSyncRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
