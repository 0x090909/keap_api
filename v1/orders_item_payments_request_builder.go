package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// OrdersItemPaymentsRequestBuilder builds and executes requests for operations under \v1\orders\{orderId}\payments
type OrdersItemPaymentsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// OrdersItemPaymentsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OrdersItemPaymentsRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// OrdersItemPaymentsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OrdersItemPaymentsRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewOrdersItemPaymentsRequestBuilderInternal instantiates a new OrdersItemPaymentsRequestBuilder and sets the default values.
func NewOrdersItemPaymentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *OrdersItemPaymentsRequestBuilder {
	m := &OrdersItemPaymentsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/orders/{orderId}/payments", pathParameters),
	}
	return m
}

// NewOrdersItemPaymentsRequestBuilder instantiates a new OrdersItemPaymentsRequestBuilder and sets the default values.
func NewOrdersItemPaymentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *OrdersItemPaymentsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewOrdersItemPaymentsRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieves a list of payments made against a given order, including historical or external payments of cash or credit card.
// returns a []OrdersItemPaymentsable when successful
// returns a OrdersItemPayments401Error error when the service returns a 401 status code
// returns a OrdersItemPayments403Error error when the service returns a 403 status code
// returns a OrdersItemPayments404Error error when the service returns a 404 status code
func (m *OrdersItemPaymentsRequestBuilder) Get(ctx context.Context, requestConfiguration *OrdersItemPaymentsRequestBuilderGetRequestConfiguration) ([]OrdersItemPaymentsable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateOrdersItemPayments401ErrorFromDiscriminatorValue,
		"403": CreateOrdersItemPayments403ErrorFromDiscriminatorValue,
		"404": CreateOrdersItemPayments404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, CreateOrdersItemPaymentsFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	val := make([]OrdersItemPaymentsable, len(res))
	for i, v := range res {
		if v != nil {
			val[i] = v.(OrdersItemPaymentsable)
		}
	}
	return val, nil
}

// Post creates a payment record that can charge a credit card. Alternatively, adds a record of historical or external payment of cash or credit card.
// Deprecated: This method is obsolete. Use PostAsPaymentsPostResponse instead.
// returns a OrdersItemPaymentsResponseable when successful
// returns a OrdersItemPayments401Error error when the service returns a 401 status code
// returns a OrdersItemPayments403Error error when the service returns a 403 status code
func (m *OrdersItemPaymentsRequestBuilder) Post(ctx context.Context, body OrdersItemPaymentsPostRequestBodyable, requestConfiguration *OrdersItemPaymentsRequestBuilderPostRequestConfiguration) (OrdersItemPaymentsResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateOrdersItemPayments401ErrorFromDiscriminatorValue,
		"403": CreateOrdersItemPayments403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOrdersItemPaymentsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(OrdersItemPaymentsResponseable), nil
}

// PostAsPaymentsPostResponse creates a payment record that can charge a credit card. Alternatively, adds a record of historical or external payment of cash or credit card.
// returns a OrdersItemPaymentsPostResponseable when successful
// returns a OrdersItemPayments401Error error when the service returns a 401 status code
// returns a OrdersItemPayments403Error error when the service returns a 403 status code
func (m *OrdersItemPaymentsRequestBuilder) PostAsPaymentsPostResponse(ctx context.Context, body OrdersItemPaymentsPostRequestBodyable, requestConfiguration *OrdersItemPaymentsRequestBuilderPostRequestConfiguration) (OrdersItemPaymentsPostResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateOrdersItemPayments401ErrorFromDiscriminatorValue,
		"403": CreateOrdersItemPayments403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOrdersItemPaymentsPostResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(OrdersItemPaymentsPostResponseable), nil
}

// ToGetRequestInformation retrieves a list of payments made against a given order, including historical or external payments of cash or credit card.
// returns a *RequestInformation when successful
func (m *OrdersItemPaymentsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OrdersItemPaymentsRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// ToPostRequestInformation creates a payment record that can charge a credit card. Alternatively, adds a record of historical or external payment of cash or credit card.
// returns a *RequestInformation when successful
func (m *OrdersItemPaymentsRequestBuilder) ToPostRequestInformation(ctx context.Context, body OrdersItemPaymentsPostRequestBodyable, requestConfiguration *OrdersItemPaymentsRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *OrdersItemPaymentsRequestBuilder when successful
func (m *OrdersItemPaymentsRequestBuilder) WithUrl(rawUrl string) *OrdersItemPaymentsRequestBuilder {
	return NewOrdersItemPaymentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
