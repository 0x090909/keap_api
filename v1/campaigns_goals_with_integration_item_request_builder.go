package v1

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CampaignsGoalsWithIntegrationItemRequestBuilder builds and executes requests for operations under \v1\campaigns\goals\{integration}
type CampaignsGoalsWithIntegrationItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ByCallName gets an item from the github.com/0x090909/keap_api.v1.campaigns.goals.item.item collection
// returns a *CampaignsGoalsItemWithCallNameItemRequestBuilder when successful
func (m *CampaignsGoalsWithIntegrationItemRequestBuilder) ByCallName(callName string) *CampaignsGoalsItemWithCallNameItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if callName != "" {
		urlTplParams["callName"] = callName
	}
	return NewCampaignsGoalsItemWithCallNameItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewCampaignsGoalsWithIntegrationItemRequestBuilderInternal instantiates a new CampaignsGoalsWithIntegrationItemRequestBuilder and sets the default values.
func NewCampaignsGoalsWithIntegrationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CampaignsGoalsWithIntegrationItemRequestBuilder {
	m := &CampaignsGoalsWithIntegrationItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/campaigns/goals/{integration}", pathParameters),
	}
	return m
}

// NewCampaignsGoalsWithIntegrationItemRequestBuilder instantiates a new CampaignsGoalsWithIntegrationItemRequestBuilder and sets the default values.
func NewCampaignsGoalsWithIntegrationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CampaignsGoalsWithIntegrationItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewCampaignsGoalsWithIntegrationItemRequestBuilderInternal(urlParams, requestAdapter)
}
