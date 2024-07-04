package v1

import (
	"context"
	i5d514860bbee9da2d4c2a51dfefb536627322519e091b28af40f7e6aa46ccf41 "github.com/0x090909/keap_api/v1/orders"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	log "github.com/sirupsen/logrus"
	i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
)

// OrdersRequestBuilder builds and executes requests for operations under \v1\orders
type OrdersRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// OrdersRequestBuilderGetQueryParameters retrieves a list of all orders using the specified search criteria. Each order may or may not have items.Potential values for order status:`DRAFT`, `SENT`, `VIEWED`, `PAID`
type OrdersRequestBuilderGetQueryParameters struct {
	// Returns orders for the provided contact id
	Contact_id *int64 `uriparametername:"contact_id"`
	// Sets a total of items to return
	Limit *int32 `uriparametername:"limit"`
	// Sets a beginning range of items to return
	Offset *int32 `uriparametername:"offset"`
	// Attribute to order items by. Dates are ordered by most recent at the top. Default is creation_date.
	// Deprecated: This property is deprecated, use OrderAsGetOrderQueryParameterType instead
	Order *string `uriparametername:"order"`
	// Attribute to order items by. Dates are ordered by most recent at the top. Default is creation_date.
	OrderAsGetOrderQueryParameterType *i5d514860bbee9da2d4c2a51dfefb536627322519e091b28af40f7e6aa46ccf41.GetOrderQueryParameterType `uriparametername:"order"`
	// Sets paid status of items to return
	Paid *bool `uriparametername:"paid"`
	// Returns orders containing the provided product id
	Product_id *int64 `uriparametername:"product_id"`
	// Date to start searching from ex. `2017-01-01T22:17:59.039Z`
	Since *string `uriparametername:"since"`
	// Date to search to ex. `2017-01-01T22:17:59.039Z`
	Until *string `uriparametername:"until"`
}

// OrdersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OrdersRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *OrdersRequestBuilderGetQueryParameters
}

// OrdersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OrdersRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ByOrderId gets an item from the github.com/0x090909/keap_api.v1.orders.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *OrdersWithOrderItemRequestBuilder when successful
func (m *OrdersRequestBuilder) ByOrderId(orderId string) *OrdersWithOrderItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if orderId != "" {
		urlTplParams["orderId"] = orderId
	}
	return NewOrdersWithOrderItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// ByOrderIdInt64 gets an item from the github.com/0x090909/keap_api.v1.orders.item collection
// returns a *OrdersWithOrderItemRequestBuilder when successful
func (m *OrdersRequestBuilder) ByOrderIdInt64(orderId int64) *OrdersWithOrderItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	urlTplParams["orderId"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(orderId, 10)
	return NewOrdersWithOrderItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewOrdersRequestBuilderInternal instantiates a new OrdersRequestBuilder and sets the default values.
func NewOrdersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *OrdersRequestBuilder {
	m := &OrdersRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/orders{?contact_id*,limit*,offset*,order*,paid*,product_id*,since*,until*}", pathParameters),
	}
	return m
}

// NewOrdersRequestBuilder instantiates a new OrdersRequestBuilder and sets the default values.
func NewOrdersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *OrdersRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewOrdersRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieves a list of all orders using the specified search criteria. Each order may or may not have items.Potential values for order status:`DRAFT`, `SENT`, `VIEWED`, `PAID`
// Deprecated: This method is obsolete. Use GetAsOrdersGetResponse instead.
// returns a OrdersResponseable when successful
// returns a Orders401Error error when the service returns a 401 status code
// returns a Orders403Error error when the service returns a 403 status code
// returns a Orders404Error error when the service returns a 404 status code
func (m *OrdersRequestBuilder) Get(ctx context.Context, requestConfiguration *OrdersRequestBuilderGetRequestConfiguration) (OrdersResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateOrders401ErrorFromDiscriminatorValue,
		"403": CreateOrders403ErrorFromDiscriminatorValue,
		"404": CreateOrders404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOrdersResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(OrdersResponseable), nil
}

// GetAsOrdersGetResponse retrieves a list of all orders using the specified search criteria. Each order may or may not have items.Potential values for order status:`DRAFT`, `SENT`, `VIEWED`, `PAID`
// returns a OrdersGetResponseable when successful
// returns a Orders401Error error when the service returns a 401 status code
// returns a Orders403Error error when the service returns a 403 status code
// returns a Orders404Error error when the service returns a 404 status code
func (m *OrdersRequestBuilder) GetAsOrdersGetResponse(ctx context.Context, requestConfiguration *OrdersRequestBuilderGetRequestConfiguration) (OrdersGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateOrders401ErrorFromDiscriminatorValue,
		"403": CreateOrders403ErrorFromDiscriminatorValue,
		"404": CreateOrders404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOrdersGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(OrdersGetResponseable), nil
}

// Model the model property
// returns a *OrdersModelRequestBuilder when successful
func (m *OrdersRequestBuilder) Model() *OrdersModelRequestBuilder {
	return NewOrdersModelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Post create a one time order with order items.
// Deprecated: This method is obsolete. Use PostAsOrdersPostResponse instead.
// returns a OrdersResponseable when successful
// returns a Orders401Error error when the service returns a 401 status code
// returns a Orders403Error error when the service returns a 403 status code
func (m *OrdersRequestBuilder) Post(ctx context.Context, body OrdersPostRequestBodyable, requestConfiguration *OrdersRequestBuilderPostRequestConfiguration) (OrdersResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateOrders401ErrorFromDiscriminatorValue,
		"403": CreateOrders403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOrdersResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		log.Error(res)
		return nil, err
	}
	if res == nil {
		log.Error(res)
		return nil, nil
	}
	return res.(OrdersResponseable), nil
}

// PostAsOrdersPostResponse create a one time order with order items.
// returns a OrdersPostResponseable when successful
// returns a Orders401Error error when the service returns a 401 status code
// returns a Orders403Error error when the service returns a 403 status code
func (m *OrdersRequestBuilder) PostAsOrdersPostResponse(ctx context.Context, body OrdersPostRequestBodyable, requestConfiguration *OrdersRequestBuilderPostRequestConfiguration) (OrdersPostResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateOrders401ErrorFromDiscriminatorValue,
		"403": CreateOrders403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOrdersPostResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(OrdersPostResponseable), nil
}

// ToGetRequestInformation retrieves a list of all orders using the specified search criteria. Each order may or may not have items.Potential values for order status:`DRAFT`, `SENT`, `VIEWED`, `PAID`
// returns a *RequestInformation when successful
func (m *OrdersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OrdersRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		if requestConfiguration.QueryParameters != nil {
			requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
		}
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// ToPostRequestInformation create a one time order with order items.
// returns a *RequestInformation when successful
func (m *OrdersRequestBuilder) ToPostRequestInformation(ctx context.Context, body OrdersPostRequestBodyable, requestConfiguration *OrdersRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/v1/orders", m.BaseRequestBuilder.PathParameters)
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
// returns a *OrdersRequestBuilder when successful
func (m *OrdersRequestBuilder) WithUrl(rawUrl string) *OrdersRequestBuilder {
	return NewOrdersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
