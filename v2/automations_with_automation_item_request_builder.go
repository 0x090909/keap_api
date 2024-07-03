package v2

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AutomationsWithAutomation_ItemRequestBuilder builds and executes requests for operations under \v2\automations\{automation_id}
type AutomationsWithAutomation_ItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// AutomationsWithAutomation_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AutomationsWithAutomation_ItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewAutomationsWithAutomation_ItemRequestBuilderInternal instantiates a new AutomationsWithAutomation_ItemRequestBuilder and sets the default values.
func NewAutomationsWithAutomation_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AutomationsWithAutomation_ItemRequestBuilder {
	m := &AutomationsWithAutomation_ItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/automations/{automation_id}", pathParameters),
	}
	return m
}

// NewAutomationsWithAutomation_ItemRequestBuilder instantiates a new AutomationsWithAutomation_ItemRequestBuilder and sets the default values.
func NewAutomationsWithAutomation_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AutomationsWithAutomation_ItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewAutomationsWithAutomation_ItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieves a single automation
// Deprecated: This method is obsolete. Use GetAsWithAutomation_GetResponse instead.
// returns a AutomationsItemWithAutomation_Responseable when successful
// returns a AutomationsItemWithAutomation_401Error error when the service returns a 401 status code
// returns a AutomationsItemWithAutomation_403Error error when the service returns a 403 status code
// returns a AutomationsItemWithAutomation_404Error error when the service returns a 404 status code
func (m *AutomationsWithAutomation_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AutomationsWithAutomation_ItemRequestBuilderGetRequestConfiguration) (AutomationsItemWithAutomation_Responseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAutomationsItemWithAutomation_401ErrorFromDiscriminatorValue,
		"403": CreateAutomationsItemWithAutomation_403ErrorFromDiscriminatorValue,
		"404": CreateAutomationsItemWithAutomation_404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAutomationsItemWithAutomation_ResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AutomationsItemWithAutomation_Responseable), nil
}

// GetAsWithAutomation_GetResponse retrieves a single automation
// returns a AutomationsItemWithAutomation_GetResponseable when successful
// returns a AutomationsItemWithAutomation_401Error error when the service returns a 401 status code
// returns a AutomationsItemWithAutomation_403Error error when the service returns a 403 status code
// returns a AutomationsItemWithAutomation_404Error error when the service returns a 404 status code
func (m *AutomationsWithAutomation_ItemRequestBuilder) GetAsWithAutomation_GetResponse(ctx context.Context, requestConfiguration *AutomationsWithAutomation_ItemRequestBuilderGetRequestConfiguration) (AutomationsItemWithAutomation_GetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAutomationsItemWithAutomation_401ErrorFromDiscriminatorValue,
		"403": CreateAutomationsItemWithAutomation_403ErrorFromDiscriminatorValue,
		"404": CreateAutomationsItemWithAutomation_404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAutomationsItemWithAutomation_GetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AutomationsItemWithAutomation_GetResponseable), nil
}

// ToGetRequestInformation retrieves a single automation
// returns a *RequestInformation when successful
func (m *AutomationsWithAutomation_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AutomationsWithAutomation_ItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *AutomationsWithAutomation_ItemRequestBuilder when successful
func (m *AutomationsWithAutomation_ItemRequestBuilder) WithUrl(rawUrl string) *AutomationsWithAutomation_ItemRequestBuilder {
	return NewAutomationsWithAutomation_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
