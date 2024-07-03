package v2

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CampaignsItemSequencesWithSequence_idRemoveContactsRequestBuilder builds and executes requests for operations under \v2\campaigns\{campaign_id}\sequences\{sequence_id}:removeContacts
type CampaignsItemSequencesWithSequence_idRemoveContactsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// CampaignsItemSequencesWithSequence_idRemoveContactsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CampaignsItemSequencesWithSequence_idRemoveContactsRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewCampaignsItemSequencesWithSequence_idRemoveContactsRequestBuilderInternal instantiates a new CampaignsItemSequencesWithSequence_idRemoveContactsRequestBuilder and sets the default values.
func NewCampaignsItemSequencesWithSequence_idRemoveContactsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, sequence_id *string) *CampaignsItemSequencesWithSequence_idRemoveContactsRequestBuilder {
	m := &CampaignsItemSequencesWithSequence_idRemoveContactsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/campaigns/{campaign_id}/sequences/{sequence_id}:removeContacts", pathParameters),
	}
	if sequence_id != nil {
		m.BaseRequestBuilder.PathParameters["sequence_id"] = *sequence_id
	}
	return m
}

// NewCampaignsItemSequencesWithSequence_idRemoveContactsRequestBuilder instantiates a new CampaignsItemSequencesWithSequence_idRemoveContactsRequestBuilder and sets the default values.
func NewCampaignsItemSequencesWithSequence_idRemoveContactsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CampaignsItemSequencesWithSequence_idRemoveContactsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewCampaignsItemSequencesWithSequence_idRemoveContactsRequestBuilderInternal(urlParams, requestAdapter, nil)
}

// Post removes a list of contacts from a campaign sequenceResponse contains a map of the provided list of Contact Ids related to their individual result.
// Deprecated: This method is obsolete. Use PostAsWithSequence_idRemoveContactsPostResponse instead.
// returns a CampaignsItemSequencesWithSequence_idRemoveContactsResponseable when successful
// returns a CampaignsItemSequencesWithSequence_idRemoveContacts401Error error when the service returns a 401 status code
// returns a CampaignsItemSequencesWithSequence_idRemoveContacts403Error error when the service returns a 403 status code
func (m *CampaignsItemSequencesWithSequence_idRemoveContactsRequestBuilder) Post(ctx context.Context, body CampaignsItemSequencesWithSequence_idRemoveContactsPostRequestBodyable, requestConfiguration *CampaignsItemSequencesWithSequence_idRemoveContactsRequestBuilderPostRequestConfiguration) (CampaignsItemSequencesWithSequence_idRemoveContactsResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateCampaignsItemSequencesWithSequence_idRemoveContacts401ErrorFromDiscriminatorValue,
		"403": CreateCampaignsItemSequencesWithSequence_idRemoveContacts403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCampaignsItemSequencesWithSequence_idRemoveContactsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(CampaignsItemSequencesWithSequence_idRemoveContactsResponseable), nil
}

// PostAsWithSequence_idRemoveContactsPostResponse removes a list of contacts from a campaign sequenceResponse contains a map of the provided list of Contact Ids related to their individual result.
// returns a CampaignsItemSequencesWithSequence_idRemoveContactsPostResponseable when successful
// returns a CampaignsItemSequencesWithSequence_idRemoveContacts401Error error when the service returns a 401 status code
// returns a CampaignsItemSequencesWithSequence_idRemoveContacts403Error error when the service returns a 403 status code
func (m *CampaignsItemSequencesWithSequence_idRemoveContactsRequestBuilder) PostAsWithSequence_idRemoveContactsPostResponse(ctx context.Context, body CampaignsItemSequencesWithSequence_idRemoveContactsPostRequestBodyable, requestConfiguration *CampaignsItemSequencesWithSequence_idRemoveContactsRequestBuilderPostRequestConfiguration) (CampaignsItemSequencesWithSequence_idRemoveContactsPostResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateCampaignsItemSequencesWithSequence_idRemoveContacts401ErrorFromDiscriminatorValue,
		"403": CreateCampaignsItemSequencesWithSequence_idRemoveContacts403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCampaignsItemSequencesWithSequence_idRemoveContactsPostResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(CampaignsItemSequencesWithSequence_idRemoveContactsPostResponseable), nil
}

// ToPostRequestInformation removes a list of contacts from a campaign sequenceResponse contains a map of the provided list of Contact Ids related to their individual result.
// returns a *RequestInformation when successful
func (m *CampaignsItemSequencesWithSequence_idRemoveContactsRequestBuilder) ToPostRequestInformation(ctx context.Context, body CampaignsItemSequencesWithSequence_idRemoveContactsPostRequestBodyable, requestConfiguration *CampaignsItemSequencesWithSequence_idRemoveContactsRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CampaignsItemSequencesWithSequence_idRemoveContactsRequestBuilder when successful
func (m *CampaignsItemSequencesWithSequence_idRemoveContactsRequestBuilder) WithUrl(rawUrl string) *CampaignsItemSequencesWithSequence_idRemoveContactsRequestBuilder {
	return NewCampaignsItemSequencesWithSequence_idRemoveContactsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
