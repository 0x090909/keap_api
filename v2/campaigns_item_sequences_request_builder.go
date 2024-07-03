package v2

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CampaignsItemSequencesRequestBuilder builds and executes requests for operations under \v2\campaigns\{campaign_id}\sequences
type CampaignsItemSequencesRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// NewCampaignsItemSequencesRequestBuilderInternal instantiates a new CampaignsItemSequencesRequestBuilder and sets the default values.
func NewCampaignsItemSequencesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CampaignsItemSequencesRequestBuilder {
	m := &CampaignsItemSequencesRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/campaigns/{campaign_id}/sequences", pathParameters),
	}
	return m
}

// NewCampaignsItemSequencesRequestBuilder instantiates a new CampaignsItemSequencesRequestBuilder and sets the default values.
func NewCampaignsItemSequencesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CampaignsItemSequencesRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewCampaignsItemSequencesRequestBuilderInternal(urlParams, requestAdapter)
}

// WithSequence_idAddContacts builds and executes requests for operations under \v2\campaigns\{campaign_id}\sequences\{sequence_id}:addContacts
// returns a *CampaignsItemSequencesWithSequence_idAddContactsRequestBuilder when successful
func (m *CampaignsItemSequencesRequestBuilder) WithSequence_idAddContacts(sequence_id *string) *CampaignsItemSequencesWithSequence_idAddContactsRequestBuilder {
	return NewCampaignsItemSequencesWithSequence_idAddContactsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, sequence_id)
}

// WithSequence_idRemoveContacts builds and executes requests for operations under \v2\campaigns\{campaign_id}\sequences\{sequence_id}:removeContacts
// returns a *CampaignsItemSequencesWithSequence_idRemoveContactsRequestBuilder when successful
func (m *CampaignsItemSequencesRequestBuilder) WithSequence_idRemoveContacts(sequence_id *string) *CampaignsItemSequencesWithSequence_idRemoveContactsRequestBuilder {
	return NewCampaignsItemSequencesWithSequence_idRemoveContactsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, sequence_id)
}
