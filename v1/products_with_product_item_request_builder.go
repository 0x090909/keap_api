package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ProductsWithProductItemRequestBuilder builds and executes requests for operations under \v1\products\{productId}
type ProductsWithProductItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ProductsWithProductItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ProductsWithProductItemRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ProductsWithProductItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ProductsWithProductItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ProductsWithProductItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ProductsWithProductItemRequestBuilderPatchRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewProductsWithProductItemRequestBuilderInternal instantiates a new ProductsWithProductItemRequestBuilder and sets the default values.
func NewProductsWithProductItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ProductsWithProductItemRequestBuilder {
	m := &ProductsWithProductItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/products/{productId}", pathParameters),
	}
	return m
}

// NewProductsWithProductItemRequestBuilder instantiates a new ProductsWithProductItemRequestBuilder and sets the default values.
func NewProductsWithProductItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ProductsWithProductItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewProductsWithProductItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete deletes a product and its subscriptions
// returns a ProductsItemWithProduct401Error error when the service returns a 401 status code
// returns a ProductsItemWithProduct403Error error when the service returns a 403 status code
// returns a ProductsItemWithProduct404Error error when the service returns a 404 status code
func (m *ProductsWithProductItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ProductsWithProductItemRequestBuilderDeleteRequestConfiguration) error {
	requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateProductsItemWithProduct401ErrorFromDiscriminatorValue,
		"403": CreateProductsItemWithProduct403ErrorFromDiscriminatorValue,
		"404": CreateProductsItemWithProduct404ErrorFromDiscriminatorValue,
	}
	err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
	if err != nil {
		return err
	}
	return nil
}

// Get retrieve a Product
// Deprecated: This method is obsolete. Use GetAsWithProductGetResponse instead.
// returns a ProductsItemWithProductResponseable when successful
// returns a ProductsItemWithProduct401Error error when the service returns a 401 status code
// returns a ProductsItemWithProduct403Error error when the service returns a 403 status code
// returns a ProductsItemWithProduct404Error error when the service returns a 404 status code
func (m *ProductsWithProductItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ProductsWithProductItemRequestBuilderGetRequestConfiguration) (ProductsItemWithProductResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateProductsItemWithProduct401ErrorFromDiscriminatorValue,
		"403": CreateProductsItemWithProduct403ErrorFromDiscriminatorValue,
		"404": CreateProductsItemWithProduct404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateProductsItemWithProductResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ProductsItemWithProductResponseable), nil
}

// GetAsWithProductGetResponse retrieve a Product
// returns a ProductsItemWithProductGetResponseable when successful
// returns a ProductsItemWithProduct401Error error when the service returns a 401 status code
// returns a ProductsItemWithProduct403Error error when the service returns a 403 status code
// returns a ProductsItemWithProduct404Error error when the service returns a 404 status code
func (m *ProductsWithProductItemRequestBuilder) GetAsWithProductGetResponse(ctx context.Context, requestConfiguration *ProductsWithProductItemRequestBuilderGetRequestConfiguration) (ProductsItemWithProductGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateProductsItemWithProduct401ErrorFromDiscriminatorValue,
		"403": CreateProductsItemWithProduct403ErrorFromDiscriminatorValue,
		"404": CreateProductsItemWithProduct404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateProductsItemWithProductGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ProductsItemWithProductGetResponseable), nil
}

// Image the image property
// returns a *ProductsItemImageRequestBuilder when successful
func (m *ProductsWithProductItemRequestBuilder) Image() *ProductsItemImageRequestBuilder {
	return NewProductsItemImageRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Patch updates a Product with only the values provided in the request.
// Deprecated: This method is obsolete. Use PatchAsWithProductPatchResponse instead.
// returns a ProductsItemWithProductResponseable when successful
// returns a ProductsItemWithProduct401Error error when the service returns a 401 status code
// returns a ProductsItemWithProduct403Error error when the service returns a 403 status code
// returns a ProductsItemWithProduct404Error error when the service returns a 404 status code
func (m *ProductsWithProductItemRequestBuilder) Patch(ctx context.Context, body ProductsItemWithProductPatchRequestBodyable, requestConfiguration *ProductsWithProductItemRequestBuilderPatchRequestConfiguration) (ProductsItemWithProductResponseable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateProductsItemWithProduct401ErrorFromDiscriminatorValue,
		"403": CreateProductsItemWithProduct403ErrorFromDiscriminatorValue,
		"404": CreateProductsItemWithProduct404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateProductsItemWithProductResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ProductsItemWithProductResponseable), nil
}

// PatchAsWithProductPatchResponse updates a Product with only the values provided in the request.
// returns a ProductsItemWithProductPatchResponseable when successful
// returns a ProductsItemWithProduct401Error error when the service returns a 401 status code
// returns a ProductsItemWithProduct403Error error when the service returns a 403 status code
// returns a ProductsItemWithProduct404Error error when the service returns a 404 status code
func (m *ProductsWithProductItemRequestBuilder) PatchAsWithProductPatchResponse(ctx context.Context, body ProductsItemWithProductPatchRequestBodyable, requestConfiguration *ProductsWithProductItemRequestBuilderPatchRequestConfiguration) (ProductsItemWithProductPatchResponseable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateProductsItemWithProduct401ErrorFromDiscriminatorValue,
		"403": CreateProductsItemWithProduct403ErrorFromDiscriminatorValue,
		"404": CreateProductsItemWithProduct404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateProductsItemWithProductPatchResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ProductsItemWithProductPatchResponseable), nil
}

// Subscriptions the subscriptions property
// returns a *ProductsItemSubscriptionsRequestBuilder when successful
func (m *ProductsWithProductItemRequestBuilder) Subscriptions() *ProductsItemSubscriptionsRequestBuilder {
	return NewProductsItemSubscriptionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// ToDeleteRequestInformation deletes a product and its subscriptions
// returns a *RequestInformation when successful
func (m *ProductsWithProductItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ProductsWithProductItemRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// ToGetRequestInformation retrieve a Product
// returns a *RequestInformation when successful
func (m *ProductsWithProductItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ProductsWithProductItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// ToPatchRequestInformation updates a Product with only the values provided in the request.
// returns a *RequestInformation when successful
func (m *ProductsWithProductItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ProductsItemWithProductPatchRequestBodyable, requestConfiguration *ProductsWithProductItemRequestBuilderPatchRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ProductsWithProductItemRequestBuilder when successful
func (m *ProductsWithProductItemRequestBuilder) WithUrl(rawUrl string) *ProductsWithProductItemRequestBuilder {
	return NewProductsWithProductItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
