package v2

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// OrdersWithOrder_ItemRequestBuilder builds and executes requests for operations under \v2\orders\{order_id}
type OrdersWithOrder_ItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// NewOrdersWithOrder_ItemRequestBuilderInternal instantiates a new OrdersWithOrder_ItemRequestBuilder and sets the default values.
func NewOrdersWithOrder_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *OrdersWithOrder_ItemRequestBuilder {
	m := &OrdersWithOrder_ItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/orders/{order_id}", pathParameters),
	}
	return m
}

// NewOrdersWithOrder_ItemRequestBuilder instantiates a new OrdersWithOrder_ItemRequestBuilder and sets the default values.
func NewOrdersWithOrder_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *OrdersWithOrder_ItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewOrdersWithOrder_ItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Payments the payments property
// returns a *OrdersItemPaymentsRequestBuilder when successful
func (m *OrdersWithOrder_ItemRequestBuilder) Payments() *OrdersItemPaymentsRequestBuilder {
	return NewOrdersItemPaymentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
