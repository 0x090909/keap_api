package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AccountProfileRequestBuilder builds and executes requests for operations under \v1\account\profile
type AccountProfileRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// AccountProfileRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccountProfileRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// AccountProfileRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccountProfileRequestBuilderPutRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewAccountProfileRequestBuilderInternal instantiates a new AccountProfileRequestBuilder and sets the default values.
func NewAccountProfileRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AccountProfileRequestBuilder {
	m := &AccountProfileRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/account/profile", pathParameters),
	}
	return m
}

// NewAccountProfileRequestBuilder instantiates a new AccountProfileRequestBuilder and sets the default values.
func NewAccountProfileRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AccountProfileRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewAccountProfileRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieves profile/company info for an account.
// Deprecated: This method is obsolete. Use GetAsProfileGetResponse instead.
// returns a AccountProfileResponseable when successful
// returns a AccountProfile401Error error when the service returns a 401 status code
// returns a AccountProfile403Error error when the service returns a 403 status code
// returns a AccountProfile404Error error when the service returns a 404 status code
func (m *AccountProfileRequestBuilder) Get(ctx context.Context, requestConfiguration *AccountProfileRequestBuilderGetRequestConfiguration) (AccountProfileResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAccountProfile401ErrorFromDiscriminatorValue,
		"403": CreateAccountProfile403ErrorFromDiscriminatorValue,
		"404": CreateAccountProfile404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAccountProfileResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AccountProfileResponseable), nil
}

// GetAsProfileGetResponse retrieves profile/company info for an account.
// returns a AccountProfileGetResponseable when successful
// returns a AccountProfile401Error error when the service returns a 401 status code
// returns a AccountProfile403Error error when the service returns a 403 status code
// returns a AccountProfile404Error error when the service returns a 404 status code
func (m *AccountProfileRequestBuilder) GetAsProfileGetResponse(ctx context.Context, requestConfiguration *AccountProfileRequestBuilderGetRequestConfiguration) (AccountProfileGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAccountProfile401ErrorFromDiscriminatorValue,
		"403": CreateAccountProfile403ErrorFromDiscriminatorValue,
		"404": CreateAccountProfile404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAccountProfileGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AccountProfileGetResponseable), nil
}

// Put updates profile/company info for an account.
// Deprecated: This method is obsolete. Use PutAsProfilePutResponse instead.
// returns a AccountProfileResponseable when successful
// returns a AccountProfile401Error error when the service returns a 401 status code
// returns a AccountProfile403Error error when the service returns a 403 status code
// returns a AccountProfile404Error error when the service returns a 404 status code
func (m *AccountProfileRequestBuilder) Put(ctx context.Context, body AccountProfilePutRequestBodyable, requestConfiguration *AccountProfileRequestBuilderPutRequestConfiguration) (AccountProfileResponseable, error) {
	requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAccountProfile401ErrorFromDiscriminatorValue,
		"403": CreateAccountProfile403ErrorFromDiscriminatorValue,
		"404": CreateAccountProfile404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAccountProfileResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AccountProfileResponseable), nil
}

// PutAsProfilePutResponse updates profile/company info for an account.
// returns a AccountProfilePutResponseable when successful
// returns a AccountProfile401Error error when the service returns a 401 status code
// returns a AccountProfile403Error error when the service returns a 403 status code
// returns a AccountProfile404Error error when the service returns a 404 status code
func (m *AccountProfileRequestBuilder) PutAsProfilePutResponse(ctx context.Context, body AccountProfilePutRequestBodyable, requestConfiguration *AccountProfileRequestBuilderPutRequestConfiguration) (AccountProfilePutResponseable, error) {
	requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAccountProfile401ErrorFromDiscriminatorValue,
		"403": CreateAccountProfile403ErrorFromDiscriminatorValue,
		"404": CreateAccountProfile404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAccountProfilePutResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AccountProfilePutResponseable), nil
}

// ToGetRequestInformation retrieves profile/company info for an account.
// returns a *RequestInformation when successful
func (m *AccountProfileRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AccountProfileRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// ToPutRequestInformation updates profile/company info for an account.
// returns a *RequestInformation when successful
func (m *AccountProfileRequestBuilder) ToPutRequestInformation(ctx context.Context, body AccountProfilePutRequestBodyable, requestConfiguration *AccountProfileRequestBuilderPutRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AccountProfileRequestBuilder when successful
func (m *AccountProfileRequestBuilder) WithUrl(rawUrl string) *AccountProfileRequestBuilder {
	return NewAccountProfileRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
