package v1

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// OpportunityRequestBuilder builds and executes requests for operations under \v1\opportunity
type OpportunityRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// NewOpportunityRequestBuilderInternal instantiates a new OpportunityRequestBuilder and sets the default values.
func NewOpportunityRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *OpportunityRequestBuilder {
	m := &OpportunityRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/opportunity", pathParameters),
	}
	return m
}

// NewOpportunityRequestBuilder instantiates a new OpportunityRequestBuilder and sets the default values.
func NewOpportunityRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *OpportunityRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewOpportunityRequestBuilderInternal(urlParams, requestAdapter)
}

// Stage_pipeline the stage_pipeline property
// returns a *OpportunityStage_pipelineRequestBuilder when successful
func (m *OpportunityRequestBuilder) Stage_pipeline() *OpportunityStage_pipelineRequestBuilder {
	return NewOpportunityStage_pipelineRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
