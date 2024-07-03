package v2

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
)

// OrdersRequestBuilder builds and executes requests for operations under \v2\orders
type OrdersRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ByOrder_id gets an item from the github.com/0x090909/keap_api.v2.orders.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *OrdersWithOrder_ItemRequestBuilder when successful
func (m *OrdersRequestBuilder) ByOrder_id(order_id string) *OrdersWithOrder_ItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if order_id != "" {
		urlTplParams["order_id"] = order_id
	}
	return NewOrdersWithOrder_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// ByOrder_idInt64 gets an item from the github.com/0x090909/keap_api.v2.orders.item collection
// returns a *OrdersWithOrder_ItemRequestBuilder when successful
func (m *OrdersRequestBuilder) ByOrder_idInt64(order_id int64) *OrdersWithOrder_ItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	urlTplParams["order_id"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(order_id, 10)
	return NewOrdersWithOrder_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewOrdersRequestBuilderInternal instantiates a new OrdersRequestBuilder and sets the default values.
func NewOrdersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *OrdersRequestBuilder {
	m := &OrdersRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/orders", pathParameters),
	}
	return m
}

// NewOrdersRequestBuilder instantiates a new OrdersRequestBuilder and sets the default values.
func NewOrdersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *OrdersRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewOrdersRequestBuilderInternal(urlParams, requestAdapter)
}
