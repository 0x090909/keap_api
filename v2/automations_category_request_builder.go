package v2

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AutomationsCategoryRequestBuilder builds and executes requests for operations under \v2\automations\category
type AutomationsCategoryRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// AutomationsCategoryRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AutomationsCategoryRequestBuilderPutRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewAutomationsCategoryRequestBuilderInternal instantiates a new AutomationsCategoryRequestBuilder and sets the default values.
func NewAutomationsCategoryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AutomationsCategoryRequestBuilder {
	m := &AutomationsCategoryRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/automations/category", pathParameters),
	}
	return m
}

// NewAutomationsCategoryRequestBuilder instantiates a new AutomationsCategoryRequestBuilder and sets the default values.
func NewAutomationsCategoryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AutomationsCategoryRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewAutomationsCategoryRequestBuilderInternal(urlParams, requestAdapter)
}

// Put updates the category of one or more automations
// returns a []byte when successful
// returns a AutomationsCategory401Error error when the service returns a 401 status code
// returns a AutomationsCategory403Error error when the service returns a 403 status code
// returns a AutomationsCategory404Error error when the service returns a 404 status code
func (m *AutomationsCategoryRequestBuilder) Put(ctx context.Context, body AutomationsCategoryPutRequestBodyable, requestConfiguration *AutomationsCategoryRequestBuilderPutRequestConfiguration) ([]byte, error) {
	requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAutomationsCategory401ErrorFromDiscriminatorValue,
		"403": CreateAutomationsCategory403ErrorFromDiscriminatorValue,
		"404": CreateAutomationsCategory404ErrorFromDiscriminatorValue,
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

// ToPutRequestInformation updates the category of one or more automations
// returns a *RequestInformation when successful
func (m *AutomationsCategoryRequestBuilder) ToPutRequestInformation(ctx context.Context, body AutomationsCategoryPutRequestBodyable, requestConfiguration *AutomationsCategoryRequestBuilderPutRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AutomationsCategoryRequestBuilder when successful
func (m *AutomationsCategoryRequestBuilder) WithUrl(rawUrl string) *AutomationsCategoryRequestBuilder {
	return NewAutomationsCategoryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
