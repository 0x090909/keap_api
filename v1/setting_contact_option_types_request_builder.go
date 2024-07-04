package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SettingContactOptionTypesRequestBuilder builds and executes requests for operations under \v1\setting\contact\optionTypes
type SettingContactOptionTypesRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// SettingContactOptionTypesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SettingContactOptionTypesRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewSettingContactOptionTypesRequestBuilderInternal instantiates a new SettingContactOptionTypesRequestBuilder and sets the default values.
func NewSettingContactOptionTypesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *SettingContactOptionTypesRequestBuilder {
	m := &SettingContactOptionTypesRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/setting/contact/optionTypes", pathParameters),
	}
	return m
}

// NewSettingContactOptionTypesRequestBuilder instantiates a new SettingContactOptionTypesRequestBuilder and sets the default values.
func NewSettingContactOptionTypesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *SettingContactOptionTypesRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewSettingContactOptionTypesRequestBuilderInternal(urlParams, requestAdapter)
}

// Get lists the Contact types in a comma-separated list.  *Note:* This is now provided by GET /setting/application/configuration
// Deprecated: This method is obsolete. Use GetAsOptionTypesGetResponse instead.
// returns a SettingContactOptionTypesResponseable when successful
// returns a SettingContactOptionTypes401Error error when the service returns a 401 status code
// returns a SettingContactOptionTypes403Error error when the service returns a 403 status code
// returns a SettingContactOptionTypes404Error error when the service returns a 404 status code
func (m *SettingContactOptionTypesRequestBuilder) Get(ctx context.Context, requestConfiguration *SettingContactOptionTypesRequestBuilderGetRequestConfiguration) (SettingContactOptionTypesResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateSettingContactOptionTypes401ErrorFromDiscriminatorValue,
		"403": CreateSettingContactOptionTypes403ErrorFromDiscriminatorValue,
		"404": CreateSettingContactOptionTypes404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSettingContactOptionTypesResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(SettingContactOptionTypesResponseable), nil
}

// GetAsOptionTypesGetResponse lists the Contact types in a comma-separated list.  *Note:* This is now provided by GET /setting/application/configuration
// Deprecated:
// returns a SettingContactOptionTypesGetResponseable when successful
// returns a SettingContactOptionTypes401Error error when the service returns a 401 status code
// returns a SettingContactOptionTypes403Error error when the service returns a 403 status code
// returns a SettingContactOptionTypes404Error error when the service returns a 404 status code
func (m *SettingContactOptionTypesRequestBuilder) GetAsOptionTypesGetResponse(ctx context.Context, requestConfiguration *SettingContactOptionTypesRequestBuilderGetRequestConfiguration) (SettingContactOptionTypesGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateSettingContactOptionTypes401ErrorFromDiscriminatorValue,
		"403": CreateSettingContactOptionTypes403ErrorFromDiscriminatorValue,
		"404": CreateSettingContactOptionTypes404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSettingContactOptionTypesGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(SettingContactOptionTypesGetResponseable), nil
}

// ToGetRequestInformation lists the Contact types in a comma-separated list.  *Note:* This is now provided by GET /setting/application/configuration
// Deprecated:
// returns a *RequestInformation when successful
func (m *SettingContactOptionTypesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SettingContactOptionTypesRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated:
// returns a *SettingContactOptionTypesRequestBuilder when successful
func (m *SettingContactOptionTypesRequestBuilder) WithUrl(rawUrl string) *SettingContactOptionTypesRequestBuilder {
	return NewSettingContactOptionTypesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
