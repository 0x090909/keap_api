package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// TransactionsWithTransactionItemRequestBuilder builds and executes requests for operations under \v1\transactions\{transactionId}
type TransactionsWithTransactionItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// TransactionsWithTransactionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TransactionsWithTransactionItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewTransactionsWithTransactionItemRequestBuilderInternal instantiates a new TransactionsWithTransactionItemRequestBuilder and sets the default values.
func NewTransactionsWithTransactionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TransactionsWithTransactionItemRequestBuilder {
	m := &TransactionsWithTransactionItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/transactions/{transactionId}", pathParameters),
	}
	return m
}

// NewTransactionsWithTransactionItemRequestBuilder instantiates a new TransactionsWithTransactionItemRequestBuilder and sets the default values.
func NewTransactionsWithTransactionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TransactionsWithTransactionItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewTransactionsWithTransactionItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieves a single transaction
// Deprecated: This method is obsolete. Use GetAsWithTransactionGetResponse instead.
// returns a TransactionsItemWithTransactionResponseable when successful
// returns a TransactionsItemWithTransaction401Error error when the service returns a 401 status code
// returns a TransactionsItemWithTransaction403Error error when the service returns a 403 status code
// returns a TransactionsItemWithTransaction404Error error when the service returns a 404 status code
func (m *TransactionsWithTransactionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TransactionsWithTransactionItemRequestBuilderGetRequestConfiguration) (TransactionsItemWithTransactionResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTransactionsItemWithTransaction401ErrorFromDiscriminatorValue,
		"403": CreateTransactionsItemWithTransaction403ErrorFromDiscriminatorValue,
		"404": CreateTransactionsItemWithTransaction404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTransactionsItemWithTransactionResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TransactionsItemWithTransactionResponseable), nil
}

// GetAsWithTransactionGetResponse retrieves a single transaction
// returns a TransactionsItemWithTransactionGetResponseable when successful
// returns a TransactionsItemWithTransaction401Error error when the service returns a 401 status code
// returns a TransactionsItemWithTransaction403Error error when the service returns a 403 status code
// returns a TransactionsItemWithTransaction404Error error when the service returns a 404 status code
func (m *TransactionsWithTransactionItemRequestBuilder) GetAsWithTransactionGetResponse(ctx context.Context, requestConfiguration *TransactionsWithTransactionItemRequestBuilderGetRequestConfiguration) (TransactionsItemWithTransactionGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTransactionsItemWithTransaction401ErrorFromDiscriminatorValue,
		"403": CreateTransactionsItemWithTransaction403ErrorFromDiscriminatorValue,
		"404": CreateTransactionsItemWithTransaction404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTransactionsItemWithTransactionGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TransactionsItemWithTransactionGetResponseable), nil
}

// ToGetRequestInformation retrieves a single transaction
// returns a *RequestInformation when successful
func (m *TransactionsWithTransactionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TransactionsWithTransactionItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TransactionsWithTransactionItemRequestBuilder when successful
func (m *TransactionsWithTransactionItemRequestBuilder) WithUrl(rawUrl string) *TransactionsWithTransactionItemRequestBuilder {
	return NewTransactionsWithTransactionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
