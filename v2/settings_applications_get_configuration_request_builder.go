package v2

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	i718c05a74102a917dcb5d1b3c11debed2a54df2782be8906687e81514800fee7 "keapapi/v2/settings/applicationsgetconfiguration"
)

// SettingsApplicationsGetConfigurationRequestBuilder builds and executes requests for operations under \v2\settings\applications:getConfiguration
type SettingsApplicationsGetConfigurationRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// SettingsApplicationsGetConfigurationRequestBuilderGetQueryParameters get configuration values for the application instance.
type SettingsApplicationsGetConfigurationRequestBuilderGetQueryParameters struct {
	// By default only application data is returned. In addition to that, data is returned for the fields that are mentioned in the query.
	// Deprecated: This property is deprecated, use FieldsAsGetFieldsQueryParameterType instead
	Fields []string `uriparametername:"fields"`
	// By default only application data is returned. In addition to that, data is returned for the fields that are mentioned in the query.
	FieldsAsGetFieldsQueryParameterType []i718c05a74102a917dcb5d1b3c11debed2a54df2782be8906687e81514800fee7.GetFieldsQueryParameterType `uriparametername:"fields"`
}

// SettingsApplicationsGetConfigurationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SettingsApplicationsGetConfigurationRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *SettingsApplicationsGetConfigurationRequestBuilderGetQueryParameters
}

// NewSettingsApplicationsGetConfigurationRequestBuilderInternal instantiates a new SettingsApplicationsGetConfigurationRequestBuilder and sets the default values.
func NewSettingsApplicationsGetConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *SettingsApplicationsGetConfigurationRequestBuilder {
	m := &SettingsApplicationsGetConfigurationRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/settings/applications:getConfiguration{?fields*}", pathParameters),
	}
	return m
}

// NewSettingsApplicationsGetConfigurationRequestBuilder instantiates a new SettingsApplicationsGetConfigurationRequestBuilder and sets the default values.
func NewSettingsApplicationsGetConfigurationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *SettingsApplicationsGetConfigurationRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewSettingsApplicationsGetConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}

// Get get configuration values for the application instance.
// Deprecated: This method is obsolete. Use GetAsApplicationsGetConfigurationGetResponse instead.
// returns a SettingsApplicationsGetConfigurationResponseable when successful
// returns a SettingsApplicationsGetConfiguration401Error error when the service returns a 401 status code
// returns a SettingsApplicationsGetConfiguration403Error error when the service returns a 403 status code
// returns a SettingsApplicationsGetConfiguration404Error error when the service returns a 404 status code
func (m *SettingsApplicationsGetConfigurationRequestBuilder) Get(ctx context.Context, requestConfiguration *SettingsApplicationsGetConfigurationRequestBuilderGetRequestConfiguration) (SettingsApplicationsGetConfigurationResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateSettingsApplicationsGetConfiguration401ErrorFromDiscriminatorValue,
		"403": CreateSettingsApplicationsGetConfiguration403ErrorFromDiscriminatorValue,
		"404": CreateSettingsApplicationsGetConfiguration404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSettingsApplicationsGetConfigurationResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(SettingsApplicationsGetConfigurationResponseable), nil
}

// GetAsApplicationsGetConfigurationGetResponse get configuration values for the application instance.
// returns a SettingsApplicationsGetConfigurationGetResponseable when successful
// returns a SettingsApplicationsGetConfiguration401Error error when the service returns a 401 status code
// returns a SettingsApplicationsGetConfiguration403Error error when the service returns a 403 status code
// returns a SettingsApplicationsGetConfiguration404Error error when the service returns a 404 status code
func (m *SettingsApplicationsGetConfigurationRequestBuilder) GetAsApplicationsGetConfigurationGetResponse(ctx context.Context, requestConfiguration *SettingsApplicationsGetConfigurationRequestBuilderGetRequestConfiguration) (SettingsApplicationsGetConfigurationGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateSettingsApplicationsGetConfiguration401ErrorFromDiscriminatorValue,
		"403": CreateSettingsApplicationsGetConfiguration403ErrorFromDiscriminatorValue,
		"404": CreateSettingsApplicationsGetConfiguration404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSettingsApplicationsGetConfigurationGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(SettingsApplicationsGetConfigurationGetResponseable), nil
}

// ToGetRequestInformation get configuration values for the application instance.
// returns a *RequestInformation when successful
func (m *SettingsApplicationsGetConfigurationRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SettingsApplicationsGetConfigurationRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *SettingsApplicationsGetConfigurationRequestBuilder when successful
func (m *SettingsApplicationsGetConfigurationRequestBuilder) WithUrl(rawUrl string) *SettingsApplicationsGetConfigurationRequestBuilder {
	return NewSettingsApplicationsGetConfigurationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
