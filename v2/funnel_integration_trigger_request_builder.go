package v2

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// FunnelIntegrationTriggerRequestBuilder builds and executes requests for operations under \v2\funnelIntegration\trigger
type FunnelIntegrationTriggerRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// FunnelIntegrationTriggerRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FunnelIntegrationTriggerRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewFunnelIntegrationTriggerRequestBuilderInternal instantiates a new FunnelIntegrationTriggerRequestBuilder and sets the default values.
func NewFunnelIntegrationTriggerRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *FunnelIntegrationTriggerRequestBuilder {
	m := &FunnelIntegrationTriggerRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/funnelIntegration/trigger", pathParameters),
	}
	return m
}

// NewFunnelIntegrationTriggerRequestBuilder instantiates a new FunnelIntegrationTriggerRequestBuilder and sets the default values.
func NewFunnelIntegrationTriggerRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *FunnelIntegrationTriggerRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewFunnelIntegrationTriggerRequestBuilderInternal(urlParams, requestAdapter)
}

// Post achieve Funnel Integration Trigger Goal
// returns a []FunnelIntegrationTriggerable when successful
// returns a FunnelIntegrationTrigger401Error error when the service returns a 401 status code
// returns a FunnelIntegrationTrigger403Error error when the service returns a 403 status code
func (m *FunnelIntegrationTriggerRequestBuilder) Post(ctx context.Context, body FunnelIntegrationTriggerPostRequestBodyable, requestConfiguration *FunnelIntegrationTriggerRequestBuilderPostRequestConfiguration) ([]FunnelIntegrationTriggerable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateFunnelIntegrationTrigger401ErrorFromDiscriminatorValue,
		"403": CreateFunnelIntegrationTrigger403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, CreateFunnelIntegrationTriggerFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	val := make([]FunnelIntegrationTriggerable, len(res))
	for i, v := range res {
		if v != nil {
			val[i] = v.(FunnelIntegrationTriggerable)
		}
	}
	return val, nil
}

// ToPostRequestInformation achieve Funnel Integration Trigger Goal
// returns a *RequestInformation when successful
func (m *FunnelIntegrationTriggerRequestBuilder) ToPostRequestInformation(ctx context.Context, body FunnelIntegrationTriggerPostRequestBodyable, requestConfiguration *FunnelIntegrationTriggerRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FunnelIntegrationTriggerRequestBuilder when successful
func (m *FunnelIntegrationTriggerRequestBuilder) WithUrl(rawUrl string) *FunnelIntegrationTriggerRequestBuilder {
	return NewFunnelIntegrationTriggerRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
