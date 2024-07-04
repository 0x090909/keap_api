package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// OrdersItemPaymentPlanRequestBuilder builds and executes requests for operations under \v1\orders\{orderId}\paymentPlan
type OrdersItemPaymentPlanRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// OrdersItemPaymentPlanRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OrdersItemPaymentPlanRequestBuilderPutRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewOrdersItemPaymentPlanRequestBuilderInternal instantiates a new OrdersItemPaymentPlanRequestBuilder and sets the default values.
func NewOrdersItemPaymentPlanRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *OrdersItemPaymentPlanRequestBuilder {
	m := &OrdersItemPaymentPlanRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/orders/{orderId}/paymentPlan", pathParameters),
	}
	return m
}

// NewOrdersItemPaymentPlanRequestBuilder instantiates a new OrdersItemPaymentPlanRequestBuilder and sets the default values.
func NewOrdersItemPaymentPlanRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *OrdersItemPaymentPlanRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewOrdersItemPaymentPlanRequestBuilderInternal(urlParams, requestAdapter)
}

// Put replaces the order's payment plan with the given values.
// Deprecated: This method is obsolete. Use PutAsPaymentPlanPutResponse instead.
// returns a OrdersItemPaymentPlanResponseable when successful
// returns a OrdersItemPaymentPlan401Error error when the service returns a 401 status code
// returns a OrdersItemPaymentPlan403Error error when the service returns a 403 status code
// returns a OrdersItemPaymentPlan404Error error when the service returns a 404 status code
func (m *OrdersItemPaymentPlanRequestBuilder) Put(ctx context.Context, body OrdersItemPaymentPlanPutRequestBodyable, requestConfiguration *OrdersItemPaymentPlanRequestBuilderPutRequestConfiguration) (OrdersItemPaymentPlanResponseable, error) {
	requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateOrdersItemPaymentPlan401ErrorFromDiscriminatorValue,
		"403": CreateOrdersItemPaymentPlan403ErrorFromDiscriminatorValue,
		"404": CreateOrdersItemPaymentPlan404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOrdersItemPaymentPlanResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(OrdersItemPaymentPlanResponseable), nil
}

// PutAsPaymentPlanPutResponse replaces the order's payment plan with the given values.
// returns a OrdersItemPaymentPlanPutResponseable when successful
// returns a OrdersItemPaymentPlan401Error error when the service returns a 401 status code
// returns a OrdersItemPaymentPlan403Error error when the service returns a 403 status code
// returns a OrdersItemPaymentPlan404Error error when the service returns a 404 status code
func (m *OrdersItemPaymentPlanRequestBuilder) PutAsPaymentPlanPutResponse(ctx context.Context, body OrdersItemPaymentPlanPutRequestBodyable, requestConfiguration *OrdersItemPaymentPlanRequestBuilderPutRequestConfiguration) (OrdersItemPaymentPlanPutResponseable, error) {
	requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateOrdersItemPaymentPlan401ErrorFromDiscriminatorValue,
		"403": CreateOrdersItemPaymentPlan403ErrorFromDiscriminatorValue,
		"404": CreateOrdersItemPaymentPlan404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOrdersItemPaymentPlanPutResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(OrdersItemPaymentPlanPutResponseable), nil
}

// ToPutRequestInformation replaces the order's payment plan with the given values.
// returns a *RequestInformation when successful
func (m *OrdersItemPaymentPlanRequestBuilder) ToPutRequestInformation(ctx context.Context, body OrdersItemPaymentPlanPutRequestBodyable, requestConfiguration *OrdersItemPaymentPlanRequestBuilderPutRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *OrdersItemPaymentPlanRequestBuilder when successful
func (m *OrdersItemPaymentPlanRequestBuilder) WithUrl(rawUrl string) *OrdersItemPaymentPlanRequestBuilder {
	return NewOrdersItemPaymentPlanRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
