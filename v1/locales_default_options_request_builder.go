package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// LocalesDefaultOptionsRequestBuilder builds and executes requests for operations under \v1\locales\defaultOptions
type LocalesDefaultOptionsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// LocalesDefaultOptionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LocalesDefaultOptionsRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewLocalesDefaultOptionsRequestBuilderInternal instantiates a new LocalesDefaultOptionsRequestBuilder and sets the default values.
func NewLocalesDefaultOptionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *LocalesDefaultOptionsRequestBuilder {
	m := &LocalesDefaultOptionsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/locales/defaultOptions", pathParameters),
	}
	return m
}

// NewLocalesDefaultOptionsRequestBuilder instantiates a new LocalesDefaultOptionsRequestBuilder and sets the default values.
func NewLocalesDefaultOptionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *LocalesDefaultOptionsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewLocalesDefaultOptionsRequestBuilderInternal(urlParams, requestAdapter)
}

// Get list dropdown default options
// Deprecated: This method is obsolete. Use GetAsDefaultOptionsGetResponse instead.
// returns a LocalesDefaultOptionsResponseable when successful
// returns a LocalesDefaultOptions401Error error when the service returns a 401 status code
// returns a LocalesDefaultOptions403Error error when the service returns a 403 status code
// returns a LocalesDefaultOptions404Error error when the service returns a 404 status code
func (m *LocalesDefaultOptionsRequestBuilder) Get(ctx context.Context, requestConfiguration *LocalesDefaultOptionsRequestBuilderGetRequestConfiguration) (LocalesDefaultOptionsResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateLocalesDefaultOptions401ErrorFromDiscriminatorValue,
		"403": CreateLocalesDefaultOptions403ErrorFromDiscriminatorValue,
		"404": CreateLocalesDefaultOptions404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateLocalesDefaultOptionsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(LocalesDefaultOptionsResponseable), nil
}

// GetAsDefaultOptionsGetResponse list dropdown default options
// returns a LocalesDefaultOptionsGetResponseable when successful
// returns a LocalesDefaultOptions401Error error when the service returns a 401 status code
// returns a LocalesDefaultOptions403Error error when the service returns a 403 status code
// returns a LocalesDefaultOptions404Error error when the service returns a 404 status code
func (m *LocalesDefaultOptionsRequestBuilder) GetAsDefaultOptionsGetResponse(ctx context.Context, requestConfiguration *LocalesDefaultOptionsRequestBuilderGetRequestConfiguration) (LocalesDefaultOptionsGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateLocalesDefaultOptions401ErrorFromDiscriminatorValue,
		"403": CreateLocalesDefaultOptions403ErrorFromDiscriminatorValue,
		"404": CreateLocalesDefaultOptions404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateLocalesDefaultOptionsGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(LocalesDefaultOptionsGetResponseable), nil
}

// ToGetRequestInformation list dropdown default options
// returns a *RequestInformation when successful
func (m *LocalesDefaultOptionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LocalesDefaultOptionsRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *LocalesDefaultOptionsRequestBuilder when successful
func (m *LocalesDefaultOptionsRequestBuilder) WithUrl(rawUrl string) *LocalesDefaultOptionsRequestBuilder {
	return NewLocalesDefaultOptionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
