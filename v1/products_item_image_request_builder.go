package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ProductsItemImageRequestBuilder builds and executes requests for operations under \v1\products\{productId}\image
type ProductsItemImageRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ProductsItemImageRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ProductsItemImageRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ProductsItemImageRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ProductsItemImageRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewProductsItemImageRequestBuilderInternal instantiates a new ProductsItemImageRequestBuilder and sets the default values.
func NewProductsItemImageRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ProductsItemImageRequestBuilder {
	m := &ProductsItemImageRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/products/{productId}/image", pathParameters),
	}
	return m
}

// NewProductsItemImageRequestBuilder instantiates a new ProductsItemImageRequestBuilder and sets the default values.
func NewProductsItemImageRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ProductsItemImageRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewProductsItemImageRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete delete a product image
// returns a ProductsItemImage401Error error when the service returns a 401 status code
// returns a ProductsItemImage403Error error when the service returns a 403 status code
// returns a ProductsItemImage404Error error when the service returns a 404 status code
func (m *ProductsItemImageRequestBuilder) Delete(ctx context.Context, requestConfiguration *ProductsItemImageRequestBuilderDeleteRequestConfiguration) error {
	requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateProductsItemImage401ErrorFromDiscriminatorValue,
		"403": CreateProductsItemImage403ErrorFromDiscriminatorValue,
		"404": CreateProductsItemImage404ErrorFromDiscriminatorValue,
	}
	err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
	if err != nil {
		return err
	}
	return nil
}

// Post max payload 3 megabytes, the `file_data` is base64 encoded.
// returns a ProductsItemImage401Error error when the service returns a 401 status code
// returns a ProductsItemImage403Error error when the service returns a 403 status code
func (m *ProductsItemImageRequestBuilder) Post(ctx context.Context, body ProductsItemImagePostRequestBodyable, requestConfiguration *ProductsItemImageRequestBuilderPostRequestConfiguration) error {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateProductsItemImage401ErrorFromDiscriminatorValue,
		"403": CreateProductsItemImage403ErrorFromDiscriminatorValue,
	}
	err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
	if err != nil {
		return err
	}
	return nil
}

// ToDeleteRequestInformation delete a product image
// returns a *RequestInformation when successful
func (m *ProductsItemImageRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ProductsItemImageRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// ToPostRequestInformation max payload 3 megabytes, the `file_data` is base64 encoded.
// returns a *RequestInformation when successful
func (m *ProductsItemImageRequestBuilder) ToPostRequestInformation(ctx context.Context, body ProductsItemImagePostRequestBodyable, requestConfiguration *ProductsItemImageRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ProductsItemImageRequestBuilder when successful
func (m *ProductsItemImageRequestBuilder) WithUrl(rawUrl string) *ProductsItemImageRequestBuilder {
	return NewProductsItemImageRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
