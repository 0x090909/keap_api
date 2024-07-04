package v1

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CampaignsGoalsRequestBuilder builds and executes requests for operations under \v1\campaigns\goals
type CampaignsGoalsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ByIntegration gets an item from the github.com/0x090909/keap_api.v1.campaigns.goals.item collection
// returns a *CampaignsGoalsWithIntegrationItemRequestBuilder when successful
func (m *CampaignsGoalsRequestBuilder) ByIntegration(integration string) *CampaignsGoalsWithIntegrationItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if integration != "" {
		urlTplParams["integration"] = integration
	}
	return NewCampaignsGoalsWithIntegrationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewCampaignsGoalsRequestBuilderInternal instantiates a new CampaignsGoalsRequestBuilder and sets the default values.
func NewCampaignsGoalsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CampaignsGoalsRequestBuilder {
	m := &CampaignsGoalsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/campaigns/goals", pathParameters),
	}
	return m
}

// NewCampaignsGoalsRequestBuilder instantiates a new CampaignsGoalsRequestBuilder and sets the default values.
func NewCampaignsGoalsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CampaignsGoalsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewCampaignsGoalsRequestBuilderInternal(urlParams, requestAdapter)
}
