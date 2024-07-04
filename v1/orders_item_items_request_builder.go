package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
)

// OrdersItemItemsRequestBuilder builds and executes requests for operations under \v1\orders\{orderId}\items
type OrdersItemItemsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// OrdersItemItemsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OrdersItemItemsRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ByOrderItemId gets an item from the github.com/0x090909/keap_api.v1.orders.item.items.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *OrdersItemItemsWithOrderItemItemRequestBuilder when successful
func (m *OrdersItemItemsRequestBuilder) ByOrderItemId(orderItemId string) *OrdersItemItemsWithOrderItemItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if orderItemId != "" {
		urlTplParams["orderItemId"] = orderItemId
	}
	return NewOrdersItemItemsWithOrderItemItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// ByOrderItemIdInt64 gets an item from the github.com/0x090909/keap_api.v1.orders.item.items.item collection
// returns a *OrdersItemItemsWithOrderItemItemRequestBuilder when successful
func (m *OrdersItemItemsRequestBuilder) ByOrderItemIdInt64(orderItemId int64) *OrdersItemItemsWithOrderItemItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	urlTplParams["orderItemId"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(orderItemId, 10)
	return NewOrdersItemItemsWithOrderItemItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewOrdersItemItemsRequestBuilderInternal instantiates a new OrdersItemItemsRequestBuilder and sets the default values.
func NewOrdersItemItemsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *OrdersItemItemsRequestBuilder {
	m := &OrdersItemItemsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/orders/{orderId}/items", pathParameters),
	}
	return m
}

// NewOrdersItemItemsRequestBuilder instantiates a new OrdersItemItemsRequestBuilder and sets the default values.
func NewOrdersItemItemsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *OrdersItemItemsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewOrdersItemItemsRequestBuilderInternal(urlParams, requestAdapter)
}

// Post creates an order item on an existing order.
// Deprecated: This method is obsolete. Use PostAsItemsPostResponse instead.
// returns a OrdersItemItemsResponseable when successful
// returns a OrdersItemItems401Error error when the service returns a 401 status code
// returns a OrdersItemItems403Error error when the service returns a 403 status code
func (m *OrdersItemItemsRequestBuilder) Post(ctx context.Context, body OrdersItemItemsPostRequestBodyable, requestConfiguration *OrdersItemItemsRequestBuilderPostRequestConfiguration) (OrdersItemItemsResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateOrdersItemItems401ErrorFromDiscriminatorValue,
		"403": CreateOrdersItemItems403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOrdersItemItemsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(OrdersItemItemsResponseable), nil
}

// PostAsItemsPostResponse creates an order item on an existing order.
// returns a OrdersItemItemsPostResponseable when successful
// returns a OrdersItemItems401Error error when the service returns a 401 status code
// returns a OrdersItemItems403Error error when the service returns a 403 status code
func (m *OrdersItemItemsRequestBuilder) PostAsItemsPostResponse(ctx context.Context, body OrdersItemItemsPostRequestBodyable, requestConfiguration *OrdersItemItemsRequestBuilderPostRequestConfiguration) (OrdersItemItemsPostResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateOrdersItemItems401ErrorFromDiscriminatorValue,
		"403": CreateOrdersItemItems403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOrdersItemItemsPostResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(OrdersItemItemsPostResponseable), nil
}

// ToPostRequestInformation creates an order item on an existing order.
// returns a *RequestInformation when successful
func (m *OrdersItemItemsRequestBuilder) ToPostRequestInformation(ctx context.Context, body OrdersItemItemsPostRequestBodyable, requestConfiguration *OrdersItemItemsRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *OrdersItemItemsRequestBuilder when successful
func (m *OrdersItemItemsRequestBuilder) WithUrl(rawUrl string) *OrdersItemItemsRequestBuilder {
	return NewOrdersItemItemsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
