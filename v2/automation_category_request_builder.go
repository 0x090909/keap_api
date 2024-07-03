package v2

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AutomationCategoryRequestBuilder builds and executes requests for operations under \v2\automationCategory
type AutomationCategoryRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// AutomationCategoryRequestBuilderDeleteQueryParameters deletes one or more automation categories based on the request parameters
type AutomationCategoryRequestBuilderDeleteQueryParameters struct {
	// ids
	Ids []int64 `uriparametername:"ids"`
}

// AutomationCategoryRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AutomationCategoryRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *AutomationCategoryRequestBuilderDeleteQueryParameters
}

// AutomationCategoryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AutomationCategoryRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// AutomationCategoryRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AutomationCategoryRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// AutomationCategoryRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AutomationCategoryRequestBuilderPutRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewAutomationCategoryRequestBuilderInternal instantiates a new AutomationCategoryRequestBuilder and sets the default values.
func NewAutomationCategoryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AutomationCategoryRequestBuilder {
	m := &AutomationCategoryRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/automationCategory", pathParameters),
	}
	return m
}

// NewAutomationCategoryRequestBuilder instantiates a new AutomationCategoryRequestBuilder and sets the default values.
func NewAutomationCategoryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AutomationCategoryRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewAutomationCategoryRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete deletes one or more automation categories based on the request parameters
// returns a AutomationCategory401Error error when the service returns a 401 status code
// returns a AutomationCategory403Error error when the service returns a 403 status code
// returns a AutomationCategory404Error error when the service returns a 404 status code
func (m *AutomationCategoryRequestBuilder) Delete(ctx context.Context, requestConfiguration *AutomationCategoryRequestBuilderDeleteRequestConfiguration) error {
	requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAutomationCategory401ErrorFromDiscriminatorValue,
		"403": CreateAutomationCategory403ErrorFromDiscriminatorValue,
		"404": CreateAutomationCategory404ErrorFromDiscriminatorValue,
	}
	err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
	if err != nil {
		return err
	}
	return nil
}

// Get lists all automation categories based on the request parameters
// Deprecated: This method is obsolete. Use GetAsAutomationCategoryGetResponse instead.
// returns a AutomationCategoryResponseable when successful
// returns a AutomationCategory401Error error when the service returns a 401 status code
// returns a AutomationCategory403Error error when the service returns a 403 status code
// returns a AutomationCategory404Error error when the service returns a 404 status code
func (m *AutomationCategoryRequestBuilder) Get(ctx context.Context, requestConfiguration *AutomationCategoryRequestBuilderGetRequestConfiguration) (AutomationCategoryResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAutomationCategory401ErrorFromDiscriminatorValue,
		"403": CreateAutomationCategory403ErrorFromDiscriminatorValue,
		"404": CreateAutomationCategory404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAutomationCategoryResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AutomationCategoryResponseable), nil
}

// GetAsAutomationCategoryGetResponse lists all automation categories based on the request parameters
// returns a AutomationCategoryGetResponseable when successful
// returns a AutomationCategory401Error error when the service returns a 401 status code
// returns a AutomationCategory403Error error when the service returns a 403 status code
// returns a AutomationCategory404Error error when the service returns a 404 status code
func (m *AutomationCategoryRequestBuilder) GetAsAutomationCategoryGetResponse(ctx context.Context, requestConfiguration *AutomationCategoryRequestBuilderGetRequestConfiguration) (AutomationCategoryGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAutomationCategory401ErrorFromDiscriminatorValue,
		"403": CreateAutomationCategory403ErrorFromDiscriminatorValue,
		"404": CreateAutomationCategory404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAutomationCategoryGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AutomationCategoryGetResponseable), nil
}

// Post creates a single automation category
// Deprecated: This method is obsolete. Use PostAsAutomationCategoryPostResponse instead.
// returns a AutomationCategoryResponseable when successful
// returns a AutomationCategory401Error error when the service returns a 401 status code
// returns a AutomationCategory403Error error when the service returns a 403 status code
func (m *AutomationCategoryRequestBuilder) Post(ctx context.Context, body AutomationCategoryPostRequestBodyable, requestConfiguration *AutomationCategoryRequestBuilderPostRequestConfiguration) (AutomationCategoryResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAutomationCategory401ErrorFromDiscriminatorValue,
		"403": CreateAutomationCategory403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAutomationCategoryResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AutomationCategoryResponseable), nil
}

// PostAsAutomationCategoryPostResponse creates a single automation category
// returns a AutomationCategoryPostResponseable when successful
// returns a AutomationCategory401Error error when the service returns a 401 status code
// returns a AutomationCategory403Error error when the service returns a 403 status code
func (m *AutomationCategoryRequestBuilder) PostAsAutomationCategoryPostResponse(ctx context.Context, body AutomationCategoryPostRequestBodyable, requestConfiguration *AutomationCategoryRequestBuilderPostRequestConfiguration) (AutomationCategoryPostResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAutomationCategory401ErrorFromDiscriminatorValue,
		"403": CreateAutomationCategory403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAutomationCategoryPostResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AutomationCategoryPostResponseable), nil
}

// Put creates or updates a single automation category
// Deprecated: This method is obsolete. Use PutAsAutomationCategoryPutResponse instead.
// returns a AutomationCategoryResponseable when successful
// returns a AutomationCategory401Error error when the service returns a 401 status code
// returns a AutomationCategory403Error error when the service returns a 403 status code
// returns a AutomationCategory404Error error when the service returns a 404 status code
func (m *AutomationCategoryRequestBuilder) Put(ctx context.Context, body AutomationCategoryPutRequestBodyable, requestConfiguration *AutomationCategoryRequestBuilderPutRequestConfiguration) (AutomationCategoryResponseable, error) {
	requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAutomationCategory401ErrorFromDiscriminatorValue,
		"403": CreateAutomationCategory403ErrorFromDiscriminatorValue,
		"404": CreateAutomationCategory404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAutomationCategoryResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AutomationCategoryResponseable), nil
}

// PutAsAutomationCategoryPutResponse creates or updates a single automation category
// returns a AutomationCategoryPutResponseable when successful
// returns a AutomationCategory401Error error when the service returns a 401 status code
// returns a AutomationCategory403Error error when the service returns a 403 status code
// returns a AutomationCategory404Error error when the service returns a 404 status code
func (m *AutomationCategoryRequestBuilder) PutAsAutomationCategoryPutResponse(ctx context.Context, body AutomationCategoryPutRequestBodyable, requestConfiguration *AutomationCategoryRequestBuilderPutRequestConfiguration) (AutomationCategoryPutResponseable, error) {
	requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAutomationCategory401ErrorFromDiscriminatorValue,
		"403": CreateAutomationCategory403ErrorFromDiscriminatorValue,
		"404": CreateAutomationCategory404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAutomationCategoryPutResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AutomationCategoryPutResponseable), nil
}

// ToDeleteRequestInformation deletes one or more automation categories based on the request parameters
// returns a *RequestInformation when successful
func (m *AutomationCategoryRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AutomationCategoryRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/v2/automationCategory?ids={ids}", m.BaseRequestBuilder.PathParameters)
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

// ToGetRequestInformation lists all automation categories based on the request parameters
// returns a *RequestInformation when successful
func (m *AutomationCategoryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AutomationCategoryRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// ToPostRequestInformation creates a single automation category
// returns a *RequestInformation when successful
func (m *AutomationCategoryRequestBuilder) ToPostRequestInformation(ctx context.Context, body AutomationCategoryPostRequestBodyable, requestConfiguration *AutomationCategoryRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToPutRequestInformation creates or updates a single automation category
// returns a *RequestInformation when successful
func (m *AutomationCategoryRequestBuilder) ToPutRequestInformation(ctx context.Context, body AutomationCategoryPutRequestBodyable, requestConfiguration *AutomationCategoryRequestBuilderPutRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AutomationCategoryRequestBuilder when successful
func (m *AutomationCategoryRequestBuilder) WithUrl(rawUrl string) *AutomationCategoryRequestBuilder {
	return NewAutomationCategoryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
