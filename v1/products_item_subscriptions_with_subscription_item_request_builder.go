package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ProductsItemSubscriptionsWithSubscriptionItemRequestBuilder builds and executes requests for operations under \v1\products\{productId}\subscriptions\{subscriptionId}
type ProductsItemSubscriptionsWithSubscriptionItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ProductsItemSubscriptionsWithSubscriptionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ProductsItemSubscriptionsWithSubscriptionItemRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ProductsItemSubscriptionsWithSubscriptionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ProductsItemSubscriptionsWithSubscriptionItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewProductsItemSubscriptionsWithSubscriptionItemRequestBuilderInternal instantiates a new ProductsItemSubscriptionsWithSubscriptionItemRequestBuilder and sets the default values.
func NewProductsItemSubscriptionsWithSubscriptionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ProductsItemSubscriptionsWithSubscriptionItemRequestBuilder {
	m := &ProductsItemSubscriptionsWithSubscriptionItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/products/{productId}/subscriptions/{subscriptionId}", pathParameters),
	}
	return m
}

// NewProductsItemSubscriptionsWithSubscriptionItemRequestBuilder instantiates a new ProductsItemSubscriptionsWithSubscriptionItemRequestBuilder and sets the default values.
func NewProductsItemSubscriptionsWithSubscriptionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ProductsItemSubscriptionsWithSubscriptionItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewProductsItemSubscriptionsWithSubscriptionItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete deletes a product subscription
// returns a ProductsItemSubscriptionsItemWithSubscription401Error error when the service returns a 401 status code
// returns a ProductsItemSubscriptionsItemWithSubscription403Error error when the service returns a 403 status code
// returns a ProductsItemSubscriptionsItemWithSubscription404Error error when the service returns a 404 status code
func (m *ProductsItemSubscriptionsWithSubscriptionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ProductsItemSubscriptionsWithSubscriptionItemRequestBuilderDeleteRequestConfiguration) error {
	requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateProductsItemSubscriptionsItemWithSubscription401ErrorFromDiscriminatorValue,
		"403": CreateProductsItemSubscriptionsItemWithSubscription403ErrorFromDiscriminatorValue,
		"404": CreateProductsItemSubscriptionsItemWithSubscription404ErrorFromDiscriminatorValue,
	}
	err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
	if err != nil {
		return err
	}
	return nil
}

// Get retrieve a Product Subscription
// Deprecated: This method is obsolete. Use GetAsWithSubscriptionGetResponse instead.
// returns a ProductsItemSubscriptionsItemWithSubscriptionResponseable when successful
// returns a ProductsItemSubscriptionsItemWithSubscription401Error error when the service returns a 401 status code
// returns a ProductsItemSubscriptionsItemWithSubscription403Error error when the service returns a 403 status code
// returns a ProductsItemSubscriptionsItemWithSubscription404Error error when the service returns a 404 status code
func (m *ProductsItemSubscriptionsWithSubscriptionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ProductsItemSubscriptionsWithSubscriptionItemRequestBuilderGetRequestConfiguration) (ProductsItemSubscriptionsItemWithSubscriptionResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateProductsItemSubscriptionsItemWithSubscription401ErrorFromDiscriminatorValue,
		"403": CreateProductsItemSubscriptionsItemWithSubscription403ErrorFromDiscriminatorValue,
		"404": CreateProductsItemSubscriptionsItemWithSubscription404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateProductsItemSubscriptionsItemWithSubscriptionResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ProductsItemSubscriptionsItemWithSubscriptionResponseable), nil
}

// GetAsWithSubscriptionGetResponse retrieve a Product Subscription
// returns a ProductsItemSubscriptionsItemWithSubscriptionGetResponseable when successful
// returns a ProductsItemSubscriptionsItemWithSubscription401Error error when the service returns a 401 status code
// returns a ProductsItemSubscriptionsItemWithSubscription403Error error when the service returns a 403 status code
// returns a ProductsItemSubscriptionsItemWithSubscription404Error error when the service returns a 404 status code
func (m *ProductsItemSubscriptionsWithSubscriptionItemRequestBuilder) GetAsWithSubscriptionGetResponse(ctx context.Context, requestConfiguration *ProductsItemSubscriptionsWithSubscriptionItemRequestBuilderGetRequestConfiguration) (ProductsItemSubscriptionsItemWithSubscriptionGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateProductsItemSubscriptionsItemWithSubscription401ErrorFromDiscriminatorValue,
		"403": CreateProductsItemSubscriptionsItemWithSubscription403ErrorFromDiscriminatorValue,
		"404": CreateProductsItemSubscriptionsItemWithSubscription404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateProductsItemSubscriptionsItemWithSubscriptionGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ProductsItemSubscriptionsItemWithSubscriptionGetResponseable), nil
}

// ToDeleteRequestInformation deletes a product subscription
// returns a *RequestInformation when successful
func (m *ProductsItemSubscriptionsWithSubscriptionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ProductsItemSubscriptionsWithSubscriptionItemRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// ToGetRequestInformation retrieve a Product Subscription
// returns a *RequestInformation when successful
func (m *ProductsItemSubscriptionsWithSubscriptionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ProductsItemSubscriptionsWithSubscriptionItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ProductsItemSubscriptionsWithSubscriptionItemRequestBuilder when successful
func (m *ProductsItemSubscriptionsWithSubscriptionItemRequestBuilder) WithUrl(rawUrl string) *ProductsItemSubscriptionsWithSubscriptionItemRequestBuilder {
	return NewProductsItemSubscriptionsWithSubscriptionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
