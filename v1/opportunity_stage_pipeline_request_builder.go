package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// OpportunityStage_pipelineRequestBuilder builds and executes requests for operations under \v1\opportunity\stage_pipeline
type OpportunityStage_pipelineRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// OpportunityStage_pipelineRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OpportunityStage_pipelineRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewOpportunityStage_pipelineRequestBuilderInternal instantiates a new OpportunityStage_pipelineRequestBuilder and sets the default values.
func NewOpportunityStage_pipelineRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *OpportunityStage_pipelineRequestBuilder {
	m := &OpportunityStage_pipelineRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/opportunity/stage_pipeline", pathParameters),
	}
	return m
}

// NewOpportunityStage_pipelineRequestBuilder instantiates a new OpportunityStage_pipelineRequestBuilder and sets the default values.
func NewOpportunityStage_pipelineRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *OpportunityStage_pipelineRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewOpportunityStage_pipelineRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieves a list of all opportunity stages with pipeline details
// returns a []OpportunityStage_pipelineable when successful
// returns a OpportunityStage_pipeline401Error error when the service returns a 401 status code
// returns a OpportunityStage_pipeline403Error error when the service returns a 403 status code
// returns a OpportunityStage_pipeline404Error error when the service returns a 404 status code
func (m *OpportunityStage_pipelineRequestBuilder) Get(ctx context.Context, requestConfiguration *OpportunityStage_pipelineRequestBuilderGetRequestConfiguration) ([]OpportunityStage_pipelineable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateOpportunityStage_pipeline401ErrorFromDiscriminatorValue,
		"403": CreateOpportunityStage_pipeline403ErrorFromDiscriminatorValue,
		"404": CreateOpportunityStage_pipeline404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, CreateOpportunityStage_pipelineFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	val := make([]OpportunityStage_pipelineable, len(res))
	for i, v := range res {
		if v != nil {
			val[i] = v.(OpportunityStage_pipelineable)
		}
	}
	return val, nil
}

// ToGetRequestInformation retrieves a list of all opportunity stages with pipeline details
// returns a *RequestInformation when successful
func (m *OpportunityStage_pipelineRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OpportunityStage_pipelineRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *OpportunityStage_pipelineRequestBuilder when successful
func (m *OpportunityStage_pipelineRequestBuilder) WithUrl(rawUrl string) *OpportunityStage_pipelineRequestBuilder {
	return NewOpportunityStage_pipelineRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
