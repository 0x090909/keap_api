package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// OrdersItemItemsWithOrderItemItemRequestBuilder builds and executes requests for operations under \v1\orders\{orderId}\items\{orderItemId}
type OrdersItemItemsWithOrderItemItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// OrdersItemItemsWithOrderItemItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OrdersItemItemsWithOrderItemItemRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewOrdersItemItemsWithOrderItemItemRequestBuilderInternal instantiates a new OrdersItemItemsWithOrderItemItemRequestBuilder and sets the default values.
func NewOrdersItemItemsWithOrderItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *OrdersItemItemsWithOrderItemItemRequestBuilder {
	m := &OrdersItemItemsWithOrderItemItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/orders/{orderId}/items/{orderItemId}", pathParameters),
	}
	return m
}

// NewOrdersItemItemsWithOrderItemItemRequestBuilder instantiates a new OrdersItemItemsWithOrderItemItemRequestBuilder and sets the default values.
func NewOrdersItemItemsWithOrderItemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *OrdersItemItemsWithOrderItemItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewOrdersItemItemsWithOrderItemItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete delete an order Item on specified Order
// returns a OrdersItemItemsWithOrderItem401Error error when the service returns a 401 status code
// returns a OrdersItemItemsWithOrderItem403Error error when the service returns a 403 status code
// returns a OrdersItemItemsWithOrderItem404Error error when the service returns a 404 status code
func (m *OrdersItemItemsWithOrderItemItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *OrdersItemItemsWithOrderItemItemRequestBuilderDeleteRequestConfiguration) error {
	requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateOrdersItemItemsWithOrderItem401ErrorFromDiscriminatorValue,
		"403": CreateOrdersItemItemsWithOrderItem403ErrorFromDiscriminatorValue,
		"404": CreateOrdersItemItemsWithOrderItem404ErrorFromDiscriminatorValue,
	}
	err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
	if err != nil {
		return err
	}
	return nil
}

// ToDeleteRequestInformation delete an order Item on specified Order
// returns a *RequestInformation when successful
func (m *OrdersItemItemsWithOrderItemItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *OrdersItemItemsWithOrderItemItemRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *OrdersItemItemsWithOrderItemItemRequestBuilder when successful
func (m *OrdersItemItemsWithOrderItemItemRequestBuilder) WithUrl(rawUrl string) *OrdersItemItemsWithOrderItemItemRequestBuilder {
	return NewOrdersItemItemsWithOrderItemItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
