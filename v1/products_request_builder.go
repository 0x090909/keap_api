package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
)

// ProductsRequestBuilder builds and executes requests for operations under \v1\products
type ProductsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ProductsRequestBuilderGetQueryParameters retrieves a list of all products
type ProductsRequestBuilderGetQueryParameters struct {
	// Sets status of items to return
	Active *bool `uriparametername:"active"`
	// Sets a total of items to return
	Limit *int32 `uriparametername:"limit"`
	// Sets a beginning range of items to return
	Offset *int32 `uriparametername:"offset"`
}

// ProductsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ProductsRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ProductsRequestBuilderGetQueryParameters
}

// ProductsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ProductsRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ByProductId gets an item from the github.com/0x090909/keap_api.v1.products.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *ProductsWithProductItemRequestBuilder when successful
func (m *ProductsRequestBuilder) ByProductId(productId string) *ProductsWithProductItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if productId != "" {
		urlTplParams["productId"] = productId
	}
	return NewProductsWithProductItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// ByProductIdInt64 gets an item from the github.com/0x090909/keap_api.v1.products.item collection
// returns a *ProductsWithProductItemRequestBuilder when successful
func (m *ProductsRequestBuilder) ByProductIdInt64(productId int64) *ProductsWithProductItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	urlTplParams["productId"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(productId, 10)
	return NewProductsWithProductItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewProductsRequestBuilderInternal instantiates a new ProductsRequestBuilder and sets the default values.
func NewProductsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ProductsRequestBuilder {
	m := &ProductsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/products{?active*,limit*,offset*}", pathParameters),
	}
	return m
}

// NewProductsRequestBuilder instantiates a new ProductsRequestBuilder and sets the default values.
func NewProductsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ProductsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewProductsRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieves a list of all products
// Deprecated: This method is obsolete. Use GetAsProductsGetResponse instead.
// returns a ProductsResponseable when successful
// returns a Products401Error error when the service returns a 401 status code
// returns a Products403Error error when the service returns a 403 status code
// returns a Products404Error error when the service returns a 404 status code
func (m *ProductsRequestBuilder) Get(ctx context.Context, requestConfiguration *ProductsRequestBuilderGetRequestConfiguration) (ProductsResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateProducts401ErrorFromDiscriminatorValue,
		"403": CreateProducts403ErrorFromDiscriminatorValue,
		"404": CreateProducts404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateProductsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ProductsResponseable), nil
}

// GetAsProductsGetResponse retrieves a list of all products
// returns a ProductsGetResponseable when successful
// returns a Products401Error error when the service returns a 401 status code
// returns a Products403Error error when the service returns a 403 status code
// returns a Products404Error error when the service returns a 404 status code
func (m *ProductsRequestBuilder) GetAsProductsGetResponse(ctx context.Context, requestConfiguration *ProductsRequestBuilderGetRequestConfiguration) (ProductsGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateProducts401ErrorFromDiscriminatorValue,
		"403": CreateProducts403ErrorFromDiscriminatorValue,
		"404": CreateProducts404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateProductsGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ProductsGetResponseable), nil
}

// Post creates a new product
// Deprecated: This method is obsolete. Use PostAsProductsPostResponse instead.
// returns a ProductsResponseable when successful
// returns a Products401Error error when the service returns a 401 status code
// returns a Products403Error error when the service returns a 403 status code
func (m *ProductsRequestBuilder) Post(ctx context.Context, body ProductsPostRequestBodyable, requestConfiguration *ProductsRequestBuilderPostRequestConfiguration) (ProductsResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateProducts401ErrorFromDiscriminatorValue,
		"403": CreateProducts403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateProductsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ProductsResponseable), nil
}

// PostAsProductsPostResponse creates a new product
// returns a ProductsPostResponseable when successful
// returns a Products401Error error when the service returns a 401 status code
// returns a Products403Error error when the service returns a 403 status code
func (m *ProductsRequestBuilder) PostAsProductsPostResponse(ctx context.Context, body ProductsPostRequestBodyable, requestConfiguration *ProductsRequestBuilderPostRequestConfiguration) (ProductsPostResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateProducts401ErrorFromDiscriminatorValue,
		"403": CreateProducts403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateProductsPostResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ProductsPostResponseable), nil
}

// Sync the sync property
// returns a *ProductsSyncRequestBuilder when successful
func (m *ProductsRequestBuilder) Sync() *ProductsSyncRequestBuilder {
	return NewProductsSyncRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// ToGetRequestInformation retrieves a list of all products
// returns a *RequestInformation when successful
func (m *ProductsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ProductsRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToPostRequestInformation creates a new product
// returns a *RequestInformation when successful
func (m *ProductsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ProductsPostRequestBodyable, requestConfiguration *ProductsRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/v1/products", m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
	if err != nil {
		return nil, err
	}
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ProductsRequestBuilder when successful
func (m *ProductsRequestBuilder) WithUrl(rawUrl string) *ProductsRequestBuilder {
	return NewProductsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
