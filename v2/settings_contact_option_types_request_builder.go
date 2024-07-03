package v2

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SettingsContactOptionTypesRequestBuilder builds and executes requests for operations under \v2\settings\contactOptionTypes
type SettingsContactOptionTypesRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// SettingsContactOptionTypesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SettingsContactOptionTypesRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewSettingsContactOptionTypesRequestBuilderInternal instantiates a new SettingsContactOptionTypesRequestBuilder and sets the default values.
func NewSettingsContactOptionTypesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *SettingsContactOptionTypesRequestBuilder {
	m := &SettingsContactOptionTypesRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/settings/contactOptionTypes", pathParameters),
	}
	return m
}

// NewSettingsContactOptionTypesRequestBuilder instantiates a new SettingsContactOptionTypesRequestBuilder and sets the default values.
func NewSettingsContactOptionTypesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *SettingsContactOptionTypesRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewSettingsContactOptionTypesRequestBuilderInternal(urlParams, requestAdapter)
}

// Get gets a list of Contact Option types.
// Deprecated: This method is obsolete. Use GetAsContactOptionTypesGetResponse instead.
// returns a SettingsContactOptionTypesResponseable when successful
// returns a SettingsContactOptionTypes401Error error when the service returns a 401 status code
// returns a SettingsContactOptionTypes403Error error when the service returns a 403 status code
// returns a SettingsContactOptionTypes404Error error when the service returns a 404 status code
func (m *SettingsContactOptionTypesRequestBuilder) Get(ctx context.Context, requestConfiguration *SettingsContactOptionTypesRequestBuilderGetRequestConfiguration) (SettingsContactOptionTypesResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateSettingsContactOptionTypes401ErrorFromDiscriminatorValue,
		"403": CreateSettingsContactOptionTypes403ErrorFromDiscriminatorValue,
		"404": CreateSettingsContactOptionTypes404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSettingsContactOptionTypesResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(SettingsContactOptionTypesResponseable), nil
}

// GetAsContactOptionTypesGetResponse gets a list of Contact Option types.
// returns a SettingsContactOptionTypesGetResponseable when successful
// returns a SettingsContactOptionTypes401Error error when the service returns a 401 status code
// returns a SettingsContactOptionTypes403Error error when the service returns a 403 status code
// returns a SettingsContactOptionTypes404Error error when the service returns a 404 status code
func (m *SettingsContactOptionTypesRequestBuilder) GetAsContactOptionTypesGetResponse(ctx context.Context, requestConfiguration *SettingsContactOptionTypesRequestBuilderGetRequestConfiguration) (SettingsContactOptionTypesGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateSettingsContactOptionTypes401ErrorFromDiscriminatorValue,
		"403": CreateSettingsContactOptionTypes403ErrorFromDiscriminatorValue,
		"404": CreateSettingsContactOptionTypes404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSettingsContactOptionTypesGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(SettingsContactOptionTypesGetResponseable), nil
}

// ToGetRequestInformation gets a list of Contact Option types.
// returns a *RequestInformation when successful
func (m *SettingsContactOptionTypesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SettingsContactOptionTypesRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *SettingsContactOptionTypesRequestBuilder when successful
func (m *SettingsContactOptionTypesRequestBuilder) WithUrl(rawUrl string) *SettingsContactOptionTypesRequestBuilder {
	return NewSettingsContactOptionTypesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
