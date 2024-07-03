package v2

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SettingsApplicationsIsEnabledRequestBuilder builds and executes requests for operations under \v2\settings\applications:isEnabled
type SettingsApplicationsIsEnabledRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// SettingsApplicationsIsEnabledRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SettingsApplicationsIsEnabledRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewSettingsApplicationsIsEnabledRequestBuilderInternal instantiates a new SettingsApplicationsIsEnabledRequestBuilder and sets the default values.
func NewSettingsApplicationsIsEnabledRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *SettingsApplicationsIsEnabledRequestBuilder {
	m := &SettingsApplicationsIsEnabledRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/settings/applications:isEnabled", pathParameters),
	}
	return m
}

// NewSettingsApplicationsIsEnabledRequestBuilder instantiates a new SettingsApplicationsIsEnabledRequestBuilder and sets the default values.
func NewSettingsApplicationsIsEnabledRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *SettingsApplicationsIsEnabledRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewSettingsApplicationsIsEnabledRequestBuilderInternal(urlParams, requestAdapter)
}

// Get check if the application is enabled or not
// Deprecated: This method is obsolete. Use GetAsApplicationsIsEnabledGetResponse instead.
// returns a SettingsApplicationsIsEnabledResponseable when successful
// returns a SettingsApplicationsIsEnabled401Error error when the service returns a 401 status code
// returns a SettingsApplicationsIsEnabled403Error error when the service returns a 403 status code
// returns a SettingsApplicationsIsEnabled404Error error when the service returns a 404 status code
func (m *SettingsApplicationsIsEnabledRequestBuilder) Get(ctx context.Context, requestConfiguration *SettingsApplicationsIsEnabledRequestBuilderGetRequestConfiguration) (SettingsApplicationsIsEnabledResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateSettingsApplicationsIsEnabled401ErrorFromDiscriminatorValue,
		"403": CreateSettingsApplicationsIsEnabled403ErrorFromDiscriminatorValue,
		"404": CreateSettingsApplicationsIsEnabled404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSettingsApplicationsIsEnabledResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(SettingsApplicationsIsEnabledResponseable), nil
}

// GetAsApplicationsIsEnabledGetResponse check if the application is enabled or not
// returns a SettingsApplicationsIsEnabledGetResponseable when successful
// returns a SettingsApplicationsIsEnabled401Error error when the service returns a 401 status code
// returns a SettingsApplicationsIsEnabled403Error error when the service returns a 403 status code
// returns a SettingsApplicationsIsEnabled404Error error when the service returns a 404 status code
func (m *SettingsApplicationsIsEnabledRequestBuilder) GetAsApplicationsIsEnabledGetResponse(ctx context.Context, requestConfiguration *SettingsApplicationsIsEnabledRequestBuilderGetRequestConfiguration) (SettingsApplicationsIsEnabledGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateSettingsApplicationsIsEnabled401ErrorFromDiscriminatorValue,
		"403": CreateSettingsApplicationsIsEnabled403ErrorFromDiscriminatorValue,
		"404": CreateSettingsApplicationsIsEnabled404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSettingsApplicationsIsEnabledGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(SettingsApplicationsIsEnabledGetResponseable), nil
}

// ToGetRequestInformation check if the application is enabled or not
// returns a *RequestInformation when successful
func (m *SettingsApplicationsIsEnabledRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SettingsApplicationsIsEnabledRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *SettingsApplicationsIsEnabledRequestBuilder when successful
func (m *SettingsApplicationsIsEnabledRequestBuilder) WithUrl(rawUrl string) *SettingsApplicationsIsEnabledRequestBuilder {
	return NewSettingsApplicationsIsEnabledRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
