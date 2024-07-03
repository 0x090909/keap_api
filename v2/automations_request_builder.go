package v2

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AutomationsRequestBuilder builds and executes requests for operations under \v2\automations
type AutomationsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// AutomationsRequestBuilderDeleteQueryParameters deletes a single automation
type AutomationsRequestBuilderDeleteQueryParameters struct {
	// automation_ids
	Automation_ids []int64 `uriparametername:"automation_ids"`
}

// AutomationsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AutomationsRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *AutomationsRequestBuilderDeleteQueryParameters
}

// AutomationsRequestBuilderGetQueryParameters lists all automations based on the request parameters.
type AutomationsRequestBuilderGetQueryParameters struct {
	// Filter to apply, allowed fields are:  - (String) name  You will need to apply the `==` operator to check the equality of the filter with your searched text, in the encoded form `%3D%3D`. The search will look for the text anywhere in the automation name.  - `filter=name%3D%3DSpring Automation` - `filter=name%3D%3DTag New Contacts`
	Filter *string `uriparametername:"filter"`
	// Attribute and direction to order items.  One of the following fields:  - name  - category  - activeContacts  - publishedDate  One of the following directions:  - asc  - desc
	Order_by *string `uriparametername:"order_by"`
	// Total number of items to return per page
	Page_size *int32 `uriparametername:"page_size"`
	// Page token
	Page_token *string `uriparametername:"page_token"`
	Stats      *bool   `uriparametername:"stats"`
}

// AutomationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AutomationsRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *AutomationsRequestBuilderGetQueryParameters
}

// ByAutomation_id gets an item from the keapapi.v2.automations.item collection
// returns a *AutomationsWithAutomation_ItemRequestBuilder when successful
func (m *AutomationsRequestBuilder) ByAutomation_id(automation_id string) *AutomationsWithAutomation_ItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if automation_id != "" {
		urlTplParams["automation_id"] = automation_id
	}
	return NewAutomationsWithAutomation_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// Category the category property
// returns a *AutomationsCategoryRequestBuilder when successful
func (m *AutomationsRequestBuilder) Category() *AutomationsCategoryRequestBuilder {
	return NewAutomationsCategoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// NewAutomationsRequestBuilderInternal instantiates a new AutomationsRequestBuilder and sets the default values.
func NewAutomationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AutomationsRequestBuilder {
	m := &AutomationsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/automations{?filter*,order_by*,page_size*,page_token*,stats*}", pathParameters),
	}
	return m
}

// NewAutomationsRequestBuilder instantiates a new AutomationsRequestBuilder and sets the default values.
func NewAutomationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AutomationsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewAutomationsRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete deletes a single automation
// returns a Automations401Error error when the service returns a 401 status code
// returns a Automations403Error error when the service returns a 403 status code
// returns a Automations404Error error when the service returns a 404 status code
func (m *AutomationsRequestBuilder) Delete(ctx context.Context, requestConfiguration *AutomationsRequestBuilderDeleteRequestConfiguration) error {
	requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAutomations401ErrorFromDiscriminatorValue,
		"403": CreateAutomations403ErrorFromDiscriminatorValue,
		"404": CreateAutomations404ErrorFromDiscriminatorValue,
	}
	err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
	if err != nil {
		return err
	}
	return nil
}

// Get lists all automations based on the request parameters.
// Deprecated: This method is obsolete. Use GetAsAutomationsGetResponse instead.
// returns a AutomationsResponseable when successful
// returns a Automations401Error error when the service returns a 401 status code
// returns a Automations403Error error when the service returns a 403 status code
// returns a Automations404Error error when the service returns a 404 status code
func (m *AutomationsRequestBuilder) Get(ctx context.Context, requestConfiguration *AutomationsRequestBuilderGetRequestConfiguration) (AutomationsResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAutomations401ErrorFromDiscriminatorValue,
		"403": CreateAutomations403ErrorFromDiscriminatorValue,
		"404": CreateAutomations404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAutomationsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AutomationsResponseable), nil
}

// GetAsAutomationsGetResponse lists all automations based on the request parameters.
// returns a AutomationsGetResponseable when successful
// returns a Automations401Error error when the service returns a 401 status code
// returns a Automations403Error error when the service returns a 403 status code
// returns a Automations404Error error when the service returns a 404 status code
func (m *AutomationsRequestBuilder) GetAsAutomationsGetResponse(ctx context.Context, requestConfiguration *AutomationsRequestBuilderGetRequestConfiguration) (AutomationsGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAutomations401ErrorFromDiscriminatorValue,
		"403": CreateAutomations403ErrorFromDiscriminatorValue,
		"404": CreateAutomations404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAutomationsGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AutomationsGetResponseable), nil
}

// ToDeleteRequestInformation deletes a single automation
// returns a *RequestInformation when successful
func (m *AutomationsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AutomationsRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/v2/automations?automation_ids={automation_ids}", m.BaseRequestBuilder.PathParameters)
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

// ToGetRequestInformation lists all automations based on the request parameters.
// returns a *RequestInformation when successful
func (m *AutomationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AutomationsRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AutomationsRequestBuilder when successful
func (m *AutomationsRequestBuilder) WithUrl(rawUrl string) *AutomationsRequestBuilder {
	return NewAutomationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
