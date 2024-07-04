package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
)

// TransactionsRequestBuilder builds and executes requests for operations under \v1\transactions
type TransactionsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// TransactionsRequestBuilderGetQueryParameters retrieves a list transactions for a given contact
type TransactionsRequestBuilderGetQueryParameters struct {
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

// TransactionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TransactionsRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *TransactionsRequestBuilderGetQueryParameters
}

// ByTransactionId gets an item from the github.com/0x090909/keap_api.v1.transactions.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *TransactionsWithTransactionItemRequestBuilder when successful
func (m *TransactionsRequestBuilder) ByTransactionId(transactionId string) *TransactionsWithTransactionItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if transactionId != "" {
		urlTplParams["transactionId"] = transactionId
	}
	return NewTransactionsWithTransactionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// ByTransactionIdInt64 gets an item from the github.com/0x090909/keap_api.v1.transactions.item collection
// returns a *TransactionsWithTransactionItemRequestBuilder when successful
func (m *TransactionsRequestBuilder) ByTransactionIdInt64(transactionId int64) *TransactionsWithTransactionItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	urlTplParams["transactionId"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(transactionId, 10)
	return NewTransactionsWithTransactionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewTransactionsRequestBuilderInternal instantiates a new TransactionsRequestBuilder and sets the default values.
func NewTransactionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TransactionsRequestBuilder {
	m := &TransactionsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/transactions{?contact_id*,limit*,offset*,since*,until*}", pathParameters),
	}
	return m
}

// NewTransactionsRequestBuilder instantiates a new TransactionsRequestBuilder and sets the default values.
func NewTransactionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TransactionsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewTransactionsRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieves a list transactions for a given contact
// Deprecated: This method is obsolete. Use GetAsTransactionsGetResponse instead.
// returns a TransactionsResponseable when successful
// returns a Transactions401Error error when the service returns a 401 status code
// returns a Transactions403Error error when the service returns a 403 status code
// returns a Transactions404Error error when the service returns a 404 status code
func (m *TransactionsRequestBuilder) Get(ctx context.Context, requestConfiguration *TransactionsRequestBuilderGetRequestConfiguration) (TransactionsResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTransactions401ErrorFromDiscriminatorValue,
		"403": CreateTransactions403ErrorFromDiscriminatorValue,
		"404": CreateTransactions404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTransactionsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TransactionsResponseable), nil
}

// GetAsTransactionsGetResponse retrieves a list transactions for a given contact
// returns a TransactionsGetResponseable when successful
// returns a Transactions401Error error when the service returns a 401 status code
// returns a Transactions403Error error when the service returns a 403 status code
// returns a Transactions404Error error when the service returns a 404 status code
func (m *TransactionsRequestBuilder) GetAsTransactionsGetResponse(ctx context.Context, requestConfiguration *TransactionsRequestBuilderGetRequestConfiguration) (TransactionsGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTransactions401ErrorFromDiscriminatorValue,
		"403": CreateTransactions403ErrorFromDiscriminatorValue,
		"404": CreateTransactions404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTransactionsGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TransactionsGetResponseable), nil
}

// ToGetRequestInformation retrieves a list transactions for a given contact
// returns a *RequestInformation when successful
func (m *TransactionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TransactionsRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TransactionsRequestBuilder when successful
func (m *TransactionsRequestBuilder) WithUrl(rawUrl string) *TransactionsRequestBuilder {
	return NewTransactionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
