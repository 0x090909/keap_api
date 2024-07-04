package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// OrdersItemTransactionsRequestBuilder builds and executes requests for operations under \v1\orders\{orderId}\transactions
type OrdersItemTransactionsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// OrdersItemTransactionsRequestBuilderGetQueryParameters retrieves a list of all transactions on a given order using the specified search criteria
type OrdersItemTransactionsRequestBuilderGetQueryParameters struct {
	// Returns transactions for the provided contact id
	Contact_id *int64 `uriparametername:"contact_id"`
	// Sets a total of items to return
	Limit *int32 `uriparametername:"limit"`
	// Sets a beginning range of items to return
	Offset *int32 `uriparametername:"offset"`
	// Date to start searching from ex. `2017-01-01T22:17:59.039Z`
	Since *string `uriparametername:"since"`
	// Date to search to ex. `2017-01-01T22:17:59.039Z`
	Until *string `uriparametername:"until"`
}

// OrdersItemTransactionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OrdersItemTransactionsRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *OrdersItemTransactionsRequestBuilderGetQueryParameters
}

// NewOrdersItemTransactionsRequestBuilderInternal instantiates a new OrdersItemTransactionsRequestBuilder and sets the default values.
func NewOrdersItemTransactionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *OrdersItemTransactionsRequestBuilder {
	m := &OrdersItemTransactionsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/orders/{orderId}/transactions{?contact_id*,limit*,offset*,since*,until*}", pathParameters),
	}
	return m
}

// NewOrdersItemTransactionsRequestBuilder instantiates a new OrdersItemTransactionsRequestBuilder and sets the default values.
func NewOrdersItemTransactionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *OrdersItemTransactionsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewOrdersItemTransactionsRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieves a list of all transactions on a given order using the specified search criteria
// Deprecated: This method is obsolete. Use GetAsTransactionsGetResponse instead.
// returns a OrdersItemTransactionsResponseable when successful
// returns a OrdersItemTransactions401Error error when the service returns a 401 status code
// returns a OrdersItemTransactions403Error error when the service returns a 403 status code
// returns a OrdersItemTransactions404Error error when the service returns a 404 status code
func (m *OrdersItemTransactionsRequestBuilder) Get(ctx context.Context, requestConfiguration *OrdersItemTransactionsRequestBuilderGetRequestConfiguration) (OrdersItemTransactionsResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateOrdersItemTransactions401ErrorFromDiscriminatorValue,
		"403": CreateOrdersItemTransactions403ErrorFromDiscriminatorValue,
		"404": CreateOrdersItemTransactions404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOrdersItemTransactionsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(OrdersItemTransactionsResponseable), nil
}

// GetAsTransactionsGetResponse retrieves a list of all transactions on a given order using the specified search criteria
// returns a OrdersItemTransactionsGetResponseable when successful
// returns a OrdersItemTransactions401Error error when the service returns a 401 status code
// returns a OrdersItemTransactions403Error error when the service returns a 403 status code
// returns a OrdersItemTransactions404Error error when the service returns a 404 status code
func (m *OrdersItemTransactionsRequestBuilder) GetAsTransactionsGetResponse(ctx context.Context, requestConfiguration *OrdersItemTransactionsRequestBuilderGetRequestConfiguration) (OrdersItemTransactionsGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateOrdersItemTransactions401ErrorFromDiscriminatorValue,
		"403": CreateOrdersItemTransactions403ErrorFromDiscriminatorValue,
		"404": CreateOrdersItemTransactions404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOrdersItemTransactionsGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(OrdersItemTransactionsGetResponseable), nil
}

// ToGetRequestInformation retrieves a list of all transactions on a given order using the specified search criteria
// returns a *RequestInformation when successful
func (m *OrdersItemTransactionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OrdersItemTransactionsRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *OrdersItemTransactionsRequestBuilder when successful
func (m *OrdersItemTransactionsRequestBuilder) WithUrl(rawUrl string) *OrdersItemTransactionsRequestBuilder {
	return NewOrdersItemTransactionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
