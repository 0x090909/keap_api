package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
)

// ProductsItemSubscriptionsRequestBuilder builds and executes requests for operations under \v1\products\{productId}\subscriptions
type ProductsItemSubscriptionsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ProductsItemSubscriptionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ProductsItemSubscriptionsRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// BySubscriptionId gets an item from the github.com/0x090909/keap_api.v1.products.item.subscriptions.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *ProductsItemSubscriptionsWithSubscriptionItemRequestBuilder when successful
func (m *ProductsItemSubscriptionsRequestBuilder) BySubscriptionId(subscriptionId string) *ProductsItemSubscriptionsWithSubscriptionItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if subscriptionId != "" {
		urlTplParams["subscriptionId"] = subscriptionId
	}
	return NewProductsItemSubscriptionsWithSubscriptionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// BySubscriptionIdInt64 gets an item from the github.com/0x090909/keap_api.v1.products.item.subscriptions.item collection
// returns a *ProductsItemSubscriptionsWithSubscriptionItemRequestBuilder when successful
func (m *ProductsItemSubscriptionsRequestBuilder) BySubscriptionIdInt64(subscriptionId int64) *ProductsItemSubscriptionsWithSubscriptionItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	urlTplParams["subscriptionId"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(subscriptionId, 10)
	return NewProductsItemSubscriptionsWithSubscriptionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewProductsItemSubscriptionsRequestBuilderInternal instantiates a new ProductsItemSubscriptionsRequestBuilder and sets the default values.
func NewProductsItemSubscriptionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ProductsItemSubscriptionsRequestBuilder {
	m := &ProductsItemSubscriptionsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/products/{productId}/subscriptions", pathParameters),
	}
	return m
}

// NewProductsItemSubscriptionsRequestBuilder instantiates a new ProductsItemSubscriptionsRequestBuilder and sets the default values.
func NewProductsItemSubscriptionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ProductsItemSubscriptionsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewProductsItemSubscriptionsRequestBuilderInternal(urlParams, requestAdapter)
}

// Post creates a new product subscription
// Deprecated: This method is obsolete. Use PostAsSubscriptionsPostResponse instead.
// returns a ProductsItemSubscriptionsResponseable when successful
// returns a ProductsItemSubscriptions401Error error when the service returns a 401 status code
// returns a ProductsItemSubscriptions403Error error when the service returns a 403 status code
func (m *ProductsItemSubscriptionsRequestBuilder) Post(ctx context.Context, body ProductsItemSubscriptionsPostRequestBodyable, requestConfiguration *ProductsItemSubscriptionsRequestBuilderPostRequestConfiguration) (ProductsItemSubscriptionsResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateProductsItemSubscriptions401ErrorFromDiscriminatorValue,
		"403": CreateProductsItemSubscriptions403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateProductsItemSubscriptionsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ProductsItemSubscriptionsResponseable), nil
}

// PostAsSubscriptionsPostResponse creates a new product subscription
// returns a ProductsItemSubscriptionsPostResponseable when successful
// returns a ProductsItemSubscriptions401Error error when the service returns a 401 status code
// returns a ProductsItemSubscriptions403Error error when the service returns a 403 status code
func (m *ProductsItemSubscriptionsRequestBuilder) PostAsSubscriptionsPostResponse(ctx context.Context, body ProductsItemSubscriptionsPostRequestBodyable, requestConfiguration *ProductsItemSubscriptionsRequestBuilderPostRequestConfiguration) (ProductsItemSubscriptionsPostResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateProductsItemSubscriptions401ErrorFromDiscriminatorValue,
		"403": CreateProductsItemSubscriptions403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateProductsItemSubscriptionsPostResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ProductsItemSubscriptionsPostResponseable), nil
}

// ToPostRequestInformation creates a new product subscription
// returns a *RequestInformation when successful
func (m *ProductsItemSubscriptionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ProductsItemSubscriptionsPostRequestBodyable, requestConfiguration *ProductsItemSubscriptionsRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ProductsItemSubscriptionsRequestBuilder when successful
func (m *ProductsItemSubscriptionsRequestBuilder) WithUrl(rawUrl string) *ProductsItemSubscriptionsRequestBuilder {
	return NewProductsItemSubscriptionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
