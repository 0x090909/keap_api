package v2

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CampaignsItemSequencesWithSequence_idAddContactsRequestBuilder builds and executes requests for operations under \v2\campaigns\{campaign_id}\sequences\{sequence_id}:addContacts
type CampaignsItemSequencesWithSequence_idAddContactsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// CampaignsItemSequencesWithSequence_idAddContactsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CampaignsItemSequencesWithSequence_idAddContactsRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewCampaignsItemSequencesWithSequence_idAddContactsRequestBuilderInternal instantiates a new CampaignsItemSequencesWithSequence_idAddContactsRequestBuilder and sets the default values.
func NewCampaignsItemSequencesWithSequence_idAddContactsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, sequence_id *string) *CampaignsItemSequencesWithSequence_idAddContactsRequestBuilder {
	m := &CampaignsItemSequencesWithSequence_idAddContactsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/campaigns/{campaign_id}/sequences/{sequence_id}:addContacts", pathParameters),
	}
	if sequence_id != nil {
		m.BaseRequestBuilder.PathParameters["sequence_id"] = *sequence_id
	}
	return m
}

// NewCampaignsItemSequencesWithSequence_idAddContactsRequestBuilder instantiates a new CampaignsItemSequencesWithSequence_idAddContactsRequestBuilder and sets the default values.
func NewCampaignsItemSequencesWithSequence_idAddContactsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CampaignsItemSequencesWithSequence_idAddContactsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewCampaignsItemSequencesWithSequence_idAddContactsRequestBuilderInternal(urlParams, requestAdapter, nil)
}

// Post adds a list of contacts to a campaign sequenceResponse contains a map of the provided list of Contact Ids related to their individual result.
// Deprecated: This method is obsolete. Use PostAsWithSequence_idAddContactsPostResponse instead.
// returns a CampaignsItemSequencesWithSequence_idAddContactsResponseable when successful
// returns a CampaignsItemSequencesWithSequence_idAddContacts401Error error when the service returns a 401 status code
// returns a CampaignsItemSequencesWithSequence_idAddContacts403Error error when the service returns a 403 status code
func (m *CampaignsItemSequencesWithSequence_idAddContactsRequestBuilder) Post(ctx context.Context, body CampaignsItemSequencesWithSequence_idAddContactsPostRequestBodyable, requestConfiguration *CampaignsItemSequencesWithSequence_idAddContactsRequestBuilderPostRequestConfiguration) (CampaignsItemSequencesWithSequence_idAddContactsResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateCampaignsItemSequencesWithSequence_idAddContacts401ErrorFromDiscriminatorValue,
		"403": CreateCampaignsItemSequencesWithSequence_idAddContacts403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCampaignsItemSequencesWithSequence_idAddContactsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(CampaignsItemSequencesWithSequence_idAddContactsResponseable), nil
}

// PostAsWithSequence_idAddContactsPostResponse adds a list of contacts to a campaign sequenceResponse contains a map of the provided list of Contact Ids related to their individual result.
// returns a CampaignsItemSequencesWithSequence_idAddContactsPostResponseable when successful
// returns a CampaignsItemSequencesWithSequence_idAddContacts401Error error when the service returns a 401 status code
// returns a CampaignsItemSequencesWithSequence_idAddContacts403Error error when the service returns a 403 status code
func (m *CampaignsItemSequencesWithSequence_idAddContactsRequestBuilder) PostAsWithSequence_idAddContactsPostResponse(ctx context.Context, body CampaignsItemSequencesWithSequence_idAddContactsPostRequestBodyable, requestConfiguration *CampaignsItemSequencesWithSequence_idAddContactsRequestBuilderPostRequestConfiguration) (CampaignsItemSequencesWithSequence_idAddContactsPostResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateCampaignsItemSequencesWithSequence_idAddContacts401ErrorFromDiscriminatorValue,
		"403": CreateCampaignsItemSequencesWithSequence_idAddContacts403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCampaignsItemSequencesWithSequence_idAddContactsPostResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(CampaignsItemSequencesWithSequence_idAddContactsPostResponseable), nil
}

// ToPostRequestInformation adds a list of contacts to a campaign sequenceResponse contains a map of the provided list of Contact Ids related to their individual result.
// returns a *RequestInformation when successful
func (m *CampaignsItemSequencesWithSequence_idAddContactsRequestBuilder) ToPostRequestInformation(ctx context.Context, body CampaignsItemSequencesWithSequence_idAddContactsPostRequestBodyable, requestConfiguration *CampaignsItemSequencesWithSequence_idAddContactsRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CampaignsItemSequencesWithSequence_idAddContactsRequestBuilder when successful
func (m *CampaignsItemSequencesWithSequence_idAddContactsRequestBuilder) WithUrl(rawUrl string) *CampaignsItemSequencesWithSequence_idAddContactsRequestBuilder {
	return NewCampaignsItemSequencesWithSequence_idAddContactsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
