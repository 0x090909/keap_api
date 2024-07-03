package v2

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// FunnelIntegrationRequestBuilder builds and executes requests for operations under \v2\funnelIntegration
type FunnelIntegrationRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// FunnelIntegrationRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FunnelIntegrationRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// FunnelIntegrationRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FunnelIntegrationRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewFunnelIntegrationRequestBuilderInternal instantiates a new FunnelIntegrationRequestBuilder and sets the default values.
func NewFunnelIntegrationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *FunnelIntegrationRequestBuilder {
	m := &FunnelIntegrationRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/funnelIntegration", pathParameters),
	}
	return m
}

// NewFunnelIntegrationRequestBuilder instantiates a new FunnelIntegrationRequestBuilder and sets the default values.
func NewFunnelIntegrationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *FunnelIntegrationRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewFunnelIntegrationRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete deletes all triggers / goals, and actions / sequence items for the given funnel integration
// returns a []byte when successful
// returns a FunnelIntegration401Error error when the service returns a 401 status code
// returns a FunnelIntegration403Error error when the service returns a 403 status code
// returns a FunnelIntegration404Error error when the service returns a 404 status code
func (m *FunnelIntegrationRequestBuilder) Delete(ctx context.Context, body FunnelIntegrationDeleteRequestBodyable, requestConfiguration *FunnelIntegrationRequestBuilderDeleteRequestConfiguration) ([]byte, error) {
	requestInfo, err := m.ToDeleteRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateFunnelIntegration401ErrorFromDiscriminatorValue,
		"403": CreateFunnelIntegration403ErrorFromDiscriminatorValue,
		"404": CreateFunnelIntegration404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.([]byte), nil
}

// Post allows a list of both triggers / goals, and actions / sequence items to be installed at the same time.
// returns a []byte when successful
// returns a FunnelIntegration401Error error when the service returns a 401 status code
// returns a FunnelIntegration403Error error when the service returns a 403 status code
func (m *FunnelIntegrationRequestBuilder) Post(ctx context.Context, body FunnelIntegrationPostRequestBodyable, requestConfiguration *FunnelIntegrationRequestBuilderPostRequestConfiguration) ([]byte, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateFunnelIntegration401ErrorFromDiscriminatorValue,
		"403": CreateFunnelIntegration403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.([]byte), nil
}

// ToDeleteRequestInformation deletes all triggers / goals, and actions / sequence items for the given funnel integration
// returns a *RequestInformation when successful
func (m *FunnelIntegrationRequestBuilder) ToDeleteRequestInformation(ctx context.Context, body FunnelIntegrationDeleteRequestBodyable, requestConfiguration *FunnelIntegrationRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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

// ToPostRequestInformation allows a list of both triggers / goals, and actions / sequence items to be installed at the same time.
// returns a *RequestInformation when successful
func (m *FunnelIntegrationRequestBuilder) ToPostRequestInformation(ctx context.Context, body FunnelIntegrationPostRequestBodyable, requestConfiguration *FunnelIntegrationRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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

// Trigger the trigger property
// returns a *FunnelIntegrationTriggerRequestBuilder when successful
func (m *FunnelIntegrationRequestBuilder) Trigger() *FunnelIntegrationTriggerRequestBuilder {
	return NewFunnelIntegrationTriggerRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FunnelIntegrationRequestBuilder when successful
func (m *FunnelIntegrationRequestBuilder) WithUrl(rawUrl string) *FunnelIntegrationRequestBuilder {
	return NewFunnelIntegrationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
