package v1

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
)

// CampaignsItemSequencesRequestBuilder builds and executes requests for operations under \v1\campaigns\{campaignId}\sequences
type CampaignsItemSequencesRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// BySequenceId gets an item from the github.com/0x090909/keap_api.v1.campaigns.item.sequences.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *CampaignsItemSequencesWithSequenceItemRequestBuilder when successful
func (m *CampaignsItemSequencesRequestBuilder) BySequenceId(sequenceId string) *CampaignsItemSequencesWithSequenceItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if sequenceId != "" {
		urlTplParams["sequenceId"] = sequenceId
	}
	return NewCampaignsItemSequencesWithSequenceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// BySequenceIdInt64 gets an item from the github.com/0x090909/keap_api.v1.campaigns.item.sequences.item collection
// returns a *CampaignsItemSequencesWithSequenceItemRequestBuilder when successful
func (m *CampaignsItemSequencesRequestBuilder) BySequenceIdInt64(sequenceId int64) *CampaignsItemSequencesWithSequenceItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	urlTplParams["sequenceId"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(sequenceId, 10)
	return NewCampaignsItemSequencesWithSequenceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewCampaignsItemSequencesRequestBuilderInternal instantiates a new CampaignsItemSequencesRequestBuilder and sets the default values.
func NewCampaignsItemSequencesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CampaignsItemSequencesRequestBuilder {
	m := &CampaignsItemSequencesRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/campaigns/{campaignId}/sequences", pathParameters),
	}
	return m
}

// NewCampaignsItemSequencesRequestBuilder instantiates a new CampaignsItemSequencesRequestBuilder and sets the default values.
func NewCampaignsItemSequencesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CampaignsItemSequencesRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewCampaignsItemSequencesRequestBuilderInternal(urlParams, requestAdapter)
}
