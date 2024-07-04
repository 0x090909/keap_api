package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// OrdersModelRequestBuilder builds and executes requests for operations under \v1\orders\model
type OrdersModelRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// OrdersModelRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OrdersModelRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewOrdersModelRequestBuilderInternal instantiates a new OrdersModelRequestBuilder and sets the default values.
func NewOrdersModelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *OrdersModelRequestBuilder {
	m := &OrdersModelRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/orders/model", pathParameters),
	}
	return m
}

// NewOrdersModelRequestBuilder instantiates a new OrdersModelRequestBuilder and sets the default values.
func NewOrdersModelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *OrdersModelRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewOrdersModelRequestBuilderInternal(urlParams, requestAdapter)
}

// Get get the custom fields for the Order object
// Deprecated: This method is obsolete. Use GetAsModelGetResponse instead.
// returns a OrdersModelResponseable when successful
// returns a OrdersModel401Error error when the service returns a 401 status code
// returns a OrdersModel403Error error when the service returns a 403 status code
// returns a OrdersModel404Error error when the service returns a 404 status code
func (m *OrdersModelRequestBuilder) Get(ctx context.Context, requestConfiguration *OrdersModelRequestBuilderGetRequestConfiguration) (OrdersModelResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateOrdersModel401ErrorFromDiscriminatorValue,
		"403": CreateOrdersModel403ErrorFromDiscriminatorValue,
		"404": CreateOrdersModel404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOrdersModelResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(OrdersModelResponseable), nil
}

// GetAsModelGetResponse get the custom fields for the Order object
// returns a OrdersModelGetResponseable when successful
// returns a OrdersModel401Error error when the service returns a 401 status code
// returns a OrdersModel403Error error when the service returns a 403 status code
// returns a OrdersModel404Error error when the service returns a 404 status code
func (m *OrdersModelRequestBuilder) GetAsModelGetResponse(ctx context.Context, requestConfiguration *OrdersModelRequestBuilderGetRequestConfiguration) (OrdersModelGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateOrdersModel401ErrorFromDiscriminatorValue,
		"403": CreateOrdersModel403ErrorFromDiscriminatorValue,
		"404": CreateOrdersModel404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOrdersModelGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(OrdersModelGetResponseable), nil
}

// ToGetRequestInformation get the custom fields for the Order object
// returns a *RequestInformation when successful
func (m *OrdersModelRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OrdersModelRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *OrdersModelRequestBuilder when successful
func (m *OrdersModelRequestBuilder) WithUrl(rawUrl string) *OrdersModelRequestBuilder {
	return NewOrdersModelRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
