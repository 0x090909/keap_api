package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// OrdersWithOrderItemRequestBuilder builds and executes requests for operations under \v1\orders\{orderId}
type OrdersWithOrderItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// OrdersWithOrderItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OrdersWithOrderItemRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// OrdersWithOrderItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OrdersWithOrderItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewOrdersWithOrderItemRequestBuilderInternal instantiates a new OrdersWithOrderItemRequestBuilder and sets the default values.
func NewOrdersWithOrderItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *OrdersWithOrderItemRequestBuilder {
	m := &OrdersWithOrderItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/orders/{orderId}", pathParameters),
	}
	return m
}

// NewOrdersWithOrderItemRequestBuilder instantiates a new OrdersWithOrderItemRequestBuilder and sets the default values.
func NewOrdersWithOrderItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *OrdersWithOrderItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewOrdersWithOrderItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete delete an order that does not have a payment
// returns a OrdersItemWithOrder401Error error when the service returns a 401 status code
// returns a OrdersItemWithOrder403Error error when the service returns a 403 status code
// returns a OrdersItemWithOrder404Error error when the service returns a 404 status code
func (m *OrdersWithOrderItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *OrdersWithOrderItemRequestBuilderDeleteRequestConfiguration) error {
	requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateOrdersItemWithOrder401ErrorFromDiscriminatorValue,
		"403": CreateOrdersItemWithOrder403ErrorFromDiscriminatorValue,
		"404": CreateOrdersItemWithOrder404ErrorFromDiscriminatorValue,
	}
	err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
	if err != nil {
		return err
	}
	return nil
}

// Get retrieves a single order. The order may or may not have items.Potential values for order status:`DRAFT`, `SENT`, `VIEWED`, `PAID`
// Deprecated: This method is obsolete. Use GetAsWithOrderGetResponse instead.
// returns a OrdersItemWithOrderResponseable when successful
// returns a OrdersItemWithOrder401Error error when the service returns a 401 status code
// returns a OrdersItemWithOrder403Error error when the service returns a 403 status code
// returns a OrdersItemWithOrder404Error error when the service returns a 404 status code
func (m *OrdersWithOrderItemRequestBuilder) Get(ctx context.Context, requestConfiguration *OrdersWithOrderItemRequestBuilderGetRequestConfiguration) (OrdersItemWithOrderResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateOrdersItemWithOrder401ErrorFromDiscriminatorValue,
		"403": CreateOrdersItemWithOrder403ErrorFromDiscriminatorValue,
		"404": CreateOrdersItemWithOrder404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOrdersItemWithOrderResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(OrdersItemWithOrderResponseable), nil
}

// GetAsWithOrderGetResponse retrieves a single order. The order may or may not have items.Potential values for order status:`DRAFT`, `SENT`, `VIEWED`, `PAID`
// returns a OrdersItemWithOrderGetResponseable when successful
// returns a OrdersItemWithOrder401Error error when the service returns a 401 status code
// returns a OrdersItemWithOrder403Error error when the service returns a 403 status code
// returns a OrdersItemWithOrder404Error error when the service returns a 404 status code
func (m *OrdersWithOrderItemRequestBuilder) GetAsWithOrderGetResponse(ctx context.Context, requestConfiguration *OrdersWithOrderItemRequestBuilderGetRequestConfiguration) (OrdersItemWithOrderGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateOrdersItemWithOrder401ErrorFromDiscriminatorValue,
		"403": CreateOrdersItemWithOrder403ErrorFromDiscriminatorValue,
		"404": CreateOrdersItemWithOrder404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOrdersItemWithOrderGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(OrdersItemWithOrderGetResponseable), nil
}

// Items the items property
// returns a *OrdersItemItemsRequestBuilder when successful
func (m *OrdersWithOrderItemRequestBuilder) Items() *OrdersItemItemsRequestBuilder {
	return NewOrdersItemItemsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// PaymentPlan the paymentPlan property
// returns a *OrdersItemPaymentPlanRequestBuilder when successful
func (m *OrdersWithOrderItemRequestBuilder) PaymentPlan() *OrdersItemPaymentPlanRequestBuilder {
	return NewOrdersItemPaymentPlanRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Payments the payments property
// returns a *OrdersItemPaymentsRequestBuilder when successful
func (m *OrdersWithOrderItemRequestBuilder) Payments() *OrdersItemPaymentsRequestBuilder {
	return NewOrdersItemPaymentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// ToDeleteRequestInformation delete an order that does not have a payment
// returns a *RequestInformation when successful
func (m *OrdersWithOrderItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *OrdersWithOrderItemRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// ToGetRequestInformation retrieves a single order. The order may or may not have items.Potential values for order status:`DRAFT`, `SENT`, `VIEWED`, `PAID`
// returns a *RequestInformation when successful
func (m *OrdersWithOrderItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OrdersWithOrderItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// Transactions the transactions property
// returns a *OrdersItemTransactionsRequestBuilder when successful
func (m *OrdersWithOrderItemRequestBuilder) Transactions() *OrdersItemTransactionsRequestBuilder {
	return NewOrdersItemTransactionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *OrdersWithOrderItemRequestBuilder when successful
func (m *OrdersWithOrderItemRequestBuilder) WithUrl(rawUrl string) *OrdersWithOrderItemRequestBuilder {
	return NewOrdersWithOrderItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
