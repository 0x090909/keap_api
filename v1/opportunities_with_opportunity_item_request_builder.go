package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// OpportunitiesWithOpportunityItemRequestBuilder builds and executes requests for operations under \v1\opportunities\{opportunityId}
type OpportunitiesWithOpportunityItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// OpportunitiesWithOpportunityItemRequestBuilderGetQueryParameters retrives a single opportunity
type OpportunitiesWithOpportunityItemRequestBuilderGetQueryParameters struct {
	// Comma-delimited list of Opportunity properties to include in the response. (Some fields such as `custom_fields` aren't included, by default.)
	Optional_properties []string `uriparametername:"optional_properties"`
}

// OpportunitiesWithOpportunityItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OpportunitiesWithOpportunityItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *OpportunitiesWithOpportunityItemRequestBuilderGetQueryParameters
}

// OpportunitiesWithOpportunityItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OpportunitiesWithOpportunityItemRequestBuilderPatchRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewOpportunitiesWithOpportunityItemRequestBuilderInternal instantiates a new OpportunitiesWithOpportunityItemRequestBuilder and sets the default values.
func NewOpportunitiesWithOpportunityItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *OpportunitiesWithOpportunityItemRequestBuilder {
	m := &OpportunitiesWithOpportunityItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/opportunities/{opportunityId}{?optional_properties*}", pathParameters),
	}
	return m
}

// NewOpportunitiesWithOpportunityItemRequestBuilder instantiates a new OpportunitiesWithOpportunityItemRequestBuilder and sets the default values.
func NewOpportunitiesWithOpportunityItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *OpportunitiesWithOpportunityItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewOpportunitiesWithOpportunityItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrives a single opportunity
// Deprecated: This method is obsolete. Use GetAsWithOpportunityGetResponse instead.
// returns a OpportunitiesItemWithOpportunityResponseable when successful
// returns a OpportunitiesItemWithOpportunity401Error error when the service returns a 401 status code
// returns a OpportunitiesItemWithOpportunity403Error error when the service returns a 403 status code
// returns a OpportunitiesItemWithOpportunity404Error error when the service returns a 404 status code
func (m *OpportunitiesWithOpportunityItemRequestBuilder) Get(ctx context.Context, requestConfiguration *OpportunitiesWithOpportunityItemRequestBuilderGetRequestConfiguration) (OpportunitiesItemWithOpportunityResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateOpportunitiesItemWithOpportunity401ErrorFromDiscriminatorValue,
		"403": CreateOpportunitiesItemWithOpportunity403ErrorFromDiscriminatorValue,
		"404": CreateOpportunitiesItemWithOpportunity404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOpportunitiesItemWithOpportunityResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(OpportunitiesItemWithOpportunityResponseable), nil
}

// GetAsWithOpportunityGetResponse retrives a single opportunity
// returns a OpportunitiesItemWithOpportunityGetResponseable when successful
// returns a OpportunitiesItemWithOpportunity401Error error when the service returns a 401 status code
// returns a OpportunitiesItemWithOpportunity403Error error when the service returns a 403 status code
// returns a OpportunitiesItemWithOpportunity404Error error when the service returns a 404 status code
func (m *OpportunitiesWithOpportunityItemRequestBuilder) GetAsWithOpportunityGetResponse(ctx context.Context, requestConfiguration *OpportunitiesWithOpportunityItemRequestBuilderGetRequestConfiguration) (OpportunitiesItemWithOpportunityGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateOpportunitiesItemWithOpportunity401ErrorFromDiscriminatorValue,
		"403": CreateOpportunitiesItemWithOpportunity403ErrorFromDiscriminatorValue,
		"404": CreateOpportunitiesItemWithOpportunity404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOpportunitiesItemWithOpportunityGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(OpportunitiesItemWithOpportunityGetResponseable), nil
}

// Patch updates an opportunity with only the values provided in the request.    *Note:* The POST payload for this request is identical to ['Replace an Opportunity'](https://developer.infusionsoft.com/docs/rest/#operation/updateOpportunityUsingPUT) operation.
// Deprecated: This method is obsolete. Use PatchAsWithOpportunityPatchResponse instead.
// returns a OpportunitiesItemWithOpportunityResponseable when successful
// returns a OpportunitiesItemWithOpportunity401Error error when the service returns a 401 status code
// returns a OpportunitiesItemWithOpportunity403Error error when the service returns a 403 status code
// returns a OpportunitiesItemWithOpportunity404Error error when the service returns a 404 status code
func (m *OpportunitiesWithOpportunityItemRequestBuilder) Patch(ctx context.Context, requestConfiguration *OpportunitiesWithOpportunityItemRequestBuilderPatchRequestConfiguration) (OpportunitiesItemWithOpportunityResponseable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateOpportunitiesItemWithOpportunity401ErrorFromDiscriminatorValue,
		"403": CreateOpportunitiesItemWithOpportunity403ErrorFromDiscriminatorValue,
		"404": CreateOpportunitiesItemWithOpportunity404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOpportunitiesItemWithOpportunityResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(OpportunitiesItemWithOpportunityResponseable), nil
}

// PatchAsWithOpportunityPatchResponse updates an opportunity with only the values provided in the request.    *Note:* The POST payload for this request is identical to ['Replace an Opportunity'](https://developer.infusionsoft.com/docs/rest/#operation/updateOpportunityUsingPUT) operation.
// returns a OpportunitiesItemWithOpportunityPatchResponseable when successful
// returns a OpportunitiesItemWithOpportunity401Error error when the service returns a 401 status code
// returns a OpportunitiesItemWithOpportunity403Error error when the service returns a 403 status code
// returns a OpportunitiesItemWithOpportunity404Error error when the service returns a 404 status code
func (m *OpportunitiesWithOpportunityItemRequestBuilder) PatchAsWithOpportunityPatchResponse(ctx context.Context, requestConfiguration *OpportunitiesWithOpportunityItemRequestBuilderPatchRequestConfiguration) (OpportunitiesItemWithOpportunityPatchResponseable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateOpportunitiesItemWithOpportunity401ErrorFromDiscriminatorValue,
		"403": CreateOpportunitiesItemWithOpportunity403ErrorFromDiscriminatorValue,
		"404": CreateOpportunitiesItemWithOpportunity404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOpportunitiesItemWithOpportunityPatchResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(OpportunitiesItemWithOpportunityPatchResponseable), nil
}

// ToGetRequestInformation retrives a single opportunity
// returns a *RequestInformation when successful
func (m *OpportunitiesWithOpportunityItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OpportunitiesWithOpportunityItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToPatchRequestInformation updates an opportunity with only the values provided in the request.    *Note:* The POST payload for this request is identical to ['Replace an Opportunity'](https://developer.infusionsoft.com/docs/rest/#operation/updateOpportunityUsingPUT) operation.
// returns a *RequestInformation when successful
func (m *OpportunitiesWithOpportunityItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, requestConfiguration *OpportunitiesWithOpportunityItemRequestBuilderPatchRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/v1/opportunities/{opportunityId}", m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *OpportunitiesWithOpportunityItemRequestBuilder when successful
func (m *OpportunitiesWithOpportunityItemRequestBuilder) WithUrl(rawUrl string) *OpportunitiesWithOpportunityItemRequestBuilder {
	return NewOpportunitiesWithOpportunityItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
