package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SettingApplicationEnabledRequestBuilder builds and executes requests for operations under \v1\setting\application\enabled
type SettingApplicationEnabledRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// SettingApplicationEnabledRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SettingApplicationEnabledRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewSettingApplicationEnabledRequestBuilderInternal instantiates a new SettingApplicationEnabledRequestBuilder and sets the default values.
func NewSettingApplicationEnabledRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *SettingApplicationEnabledRequestBuilder {
	m := &SettingApplicationEnabledRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/setting/application/enabled", pathParameters),
	}
	return m
}

// NewSettingApplicationEnabledRequestBuilder instantiates a new SettingApplicationEnabledRequestBuilder and sets the default values.
func NewSettingApplicationEnabledRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *SettingApplicationEnabledRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewSettingApplicationEnabledRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieves whether the application is enabled
// Deprecated: This method is obsolete. Use GetAsEnabledGetResponse instead.
// returns a SettingApplicationEnabledResponseable when successful
// returns a SettingApplicationEnabled401Error error when the service returns a 401 status code
// returns a SettingApplicationEnabled403Error error when the service returns a 403 status code
// returns a SettingApplicationEnabled404Error error when the service returns a 404 status code
func (m *SettingApplicationEnabledRequestBuilder) Get(ctx context.Context, requestConfiguration *SettingApplicationEnabledRequestBuilderGetRequestConfiguration) (SettingApplicationEnabledResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateSettingApplicationEnabled401ErrorFromDiscriminatorValue,
		"403": CreateSettingApplicationEnabled403ErrorFromDiscriminatorValue,
		"404": CreateSettingApplicationEnabled404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSettingApplicationEnabledResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(SettingApplicationEnabledResponseable), nil
}

// GetAsEnabledGetResponse retrieves whether the application is enabled
// returns a SettingApplicationEnabledGetResponseable when successful
// returns a SettingApplicationEnabled401Error error when the service returns a 401 status code
// returns a SettingApplicationEnabled403Error error when the service returns a 403 status code
// returns a SettingApplicationEnabled404Error error when the service returns a 404 status code
func (m *SettingApplicationEnabledRequestBuilder) GetAsEnabledGetResponse(ctx context.Context, requestConfiguration *SettingApplicationEnabledRequestBuilderGetRequestConfiguration) (SettingApplicationEnabledGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateSettingApplicationEnabled401ErrorFromDiscriminatorValue,
		"403": CreateSettingApplicationEnabled403ErrorFromDiscriminatorValue,
		"404": CreateSettingApplicationEnabled404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSettingApplicationEnabledGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(SettingApplicationEnabledGetResponseable), nil
}

// ToGetRequestInformation retrieves whether the application is enabled
// returns a *RequestInformation when successful
func (m *SettingApplicationEnabledRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SettingApplicationEnabledRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *SettingApplicationEnabledRequestBuilder when successful
func (m *SettingApplicationEnabledRequestBuilder) WithUrl(rawUrl string) *SettingApplicationEnabledRequestBuilder {
	return NewSettingApplicationEnabledRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
