package v1

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CampaignsItemSequencesWithSequenceItemRequestBuilder builds and executes requests for operations under \v1\campaigns\{campaignId}\sequences\{sequenceId}
type CampaignsItemSequencesWithSequenceItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// NewCampaignsItemSequencesWithSequenceItemRequestBuilderInternal instantiates a new CampaignsItemSequencesWithSequenceItemRequestBuilder and sets the default values.
func NewCampaignsItemSequencesWithSequenceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CampaignsItemSequencesWithSequenceItemRequestBuilder {
	m := &CampaignsItemSequencesWithSequenceItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/campaigns/{campaignId}/sequences/{sequenceId}", pathParameters),
	}
	return m
}

// NewCampaignsItemSequencesWithSequenceItemRequestBuilder instantiates a new CampaignsItemSequencesWithSequenceItemRequestBuilder and sets the default values.
func NewCampaignsItemSequencesWithSequenceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CampaignsItemSequencesWithSequenceItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewCampaignsItemSequencesWithSequenceItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Contacts the contacts property
// returns a *CampaignsItemSequencesItemContactsRequestBuilder when successful
func (m *CampaignsItemSequencesWithSequenceItemRequestBuilder) Contacts() *CampaignsItemSequencesItemContactsRequestBuilder {
	return NewCampaignsItemSequencesItemContactsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
