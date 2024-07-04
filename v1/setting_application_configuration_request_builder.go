package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SettingApplicationConfigurationRequestBuilder builds and executes requests for operations under \v1\setting\application\configuration
type SettingApplicationConfigurationRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// SettingApplicationConfigurationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SettingApplicationConfigurationRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewSettingApplicationConfigurationRequestBuilderInternal instantiates a new SettingApplicationConfigurationRequestBuilder and sets the default values.
func NewSettingApplicationConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *SettingApplicationConfigurationRequestBuilder {
	m := &SettingApplicationConfigurationRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/setting/application/configuration", pathParameters),
	}
	return m
}

// NewSettingApplicationConfigurationRequestBuilder instantiates a new SettingApplicationConfigurationRequestBuilder and sets the default values.
func NewSettingApplicationConfigurationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *SettingApplicationConfigurationRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewSettingApplicationConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}

// Get get the properties for the current application's configuration
// Deprecated: This method is obsolete. Use GetAsConfigurationGetResponse instead.
// returns a SettingApplicationConfigurationResponseable when successful
// returns a SettingApplicationConfiguration401Error error when the service returns a 401 status code
// returns a SettingApplicationConfiguration403Error error when the service returns a 403 status code
// returns a SettingApplicationConfiguration404Error error when the service returns a 404 status code
func (m *SettingApplicationConfigurationRequestBuilder) Get(ctx context.Context, requestConfiguration *SettingApplicationConfigurationRequestBuilderGetRequestConfiguration) (SettingApplicationConfigurationResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateSettingApplicationConfiguration401ErrorFromDiscriminatorValue,
		"403": CreateSettingApplicationConfiguration403ErrorFromDiscriminatorValue,
		"404": CreateSettingApplicationConfiguration404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSettingApplicationConfigurationResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(SettingApplicationConfigurationResponseable), nil
}

// GetAsConfigurationGetResponse get the properties for the current application's configuration
// returns a SettingApplicationConfigurationGetResponseable when successful
// returns a SettingApplicationConfiguration401Error error when the service returns a 401 status code
// returns a SettingApplicationConfiguration403Error error when the service returns a 403 status code
// returns a SettingApplicationConfiguration404Error error when the service returns a 404 status code
func (m *SettingApplicationConfigurationRequestBuilder) GetAsConfigurationGetResponse(ctx context.Context, requestConfiguration *SettingApplicationConfigurationRequestBuilderGetRequestConfiguration) (SettingApplicationConfigurationGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateSettingApplicationConfiguration401ErrorFromDiscriminatorValue,
		"403": CreateSettingApplicationConfiguration403ErrorFromDiscriminatorValue,
		"404": CreateSettingApplicationConfiguration404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSettingApplicationConfigurationGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(SettingApplicationConfigurationGetResponseable), nil
}

// ToGetRequestInformation get the properties for the current application's configuration
// returns a *RequestInformation when successful
func (m *SettingApplicationConfigurationRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SettingApplicationConfigurationRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *SettingApplicationConfigurationRequestBuilder when successful
func (m *SettingApplicationConfigurationRequestBuilder) WithUrl(rawUrl string) *SettingApplicationConfigurationRequestBuilder {
	return NewSettingApplicationConfigurationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
