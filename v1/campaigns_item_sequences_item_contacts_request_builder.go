package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
)

// CampaignsItemSequencesItemContactsRequestBuilder builds and executes requests for operations under \v1\campaigns\{campaignId}\sequences\{sequenceId}\contacts
type CampaignsItemSequencesItemContactsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// CampaignsItemSequencesItemContactsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CampaignsItemSequencesItemContactsRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// CampaignsItemSequencesItemContactsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CampaignsItemSequencesItemContactsRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ByContactId gets an item from the github.com/0x090909/keap_api.v1.campaigns.item.sequences.item.contacts.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *CampaignsItemSequencesItemContactsWithContactItemRequestBuilder when successful
func (m *CampaignsItemSequencesItemContactsRequestBuilder) ByContactId(contactId string) *CampaignsItemSequencesItemContactsWithContactItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if contactId != "" {
		urlTplParams["contactId"] = contactId
	}
	return NewCampaignsItemSequencesItemContactsWithContactItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// ByContactIdInt64 gets an item from the github.com/0x090909/keap_api.v1.campaigns.item.sequences.item.contacts.item collection
// returns a *CampaignsItemSequencesItemContactsWithContactItemRequestBuilder when successful
func (m *CampaignsItemSequencesItemContactsRequestBuilder) ByContactIdInt64(contactId int64) *CampaignsItemSequencesItemContactsWithContactItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	urlTplParams["contactId"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(contactId, 10)
	return NewCampaignsItemSequencesItemContactsWithContactItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewCampaignsItemSequencesItemContactsRequestBuilderInternal instantiates a new CampaignsItemSequencesItemContactsRequestBuilder and sets the default values.
func NewCampaignsItemSequencesItemContactsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CampaignsItemSequencesItemContactsRequestBuilder {
	m := &CampaignsItemSequencesItemContactsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/campaigns/{campaignId}/sequences/{sequenceId}/contacts", pathParameters),
	}
	return m
}

// NewCampaignsItemSequencesItemContactsRequestBuilder instantiates a new CampaignsItemSequencesItemContactsRequestBuilder and sets the default values.
func NewCampaignsItemSequencesItemContactsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CampaignsItemSequencesItemContactsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewCampaignsItemSequencesItemContactsRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete removes a list of contacts from a campaign sequence
// returns a CampaignsItemSequencesItemContacts401Error error when the service returns a 401 status code
// returns a CampaignsItemSequencesItemContacts403Error error when the service returns a 403 status code
// returns a CampaignsItemSequencesItemContacts404Error error when the service returns a 404 status code
func (m *CampaignsItemSequencesItemContactsRequestBuilder) Delete(ctx context.Context, body CampaignsItemSequencesItemContactsDeleteRequestBodyable, requestConfiguration *CampaignsItemSequencesItemContactsRequestBuilderDeleteRequestConfiguration) error {
	requestInfo, err := m.ToDeleteRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateCampaignsItemSequencesItemContacts401ErrorFromDiscriminatorValue,
		"403": CreateCampaignsItemSequencesItemContacts403ErrorFromDiscriminatorValue,
		"404": CreateCampaignsItemSequencesItemContacts404ErrorFromDiscriminatorValue,
	}
	err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
	if err != nil {
		return err
	}
	return nil
}

// Post adds a list of contacts to a campaign sequenceResponse contains a map of the provided list of Contact Ids related to their individual result.
// Deprecated: This method is obsolete. Use PostAsContactsPostResponse instead.
// returns a CampaignsItemSequencesItemContactsResponseable when successful
// returns a CampaignsItemSequencesItemContacts401Error error when the service returns a 401 status code
// returns a CampaignsItemSequencesItemContacts403Error error when the service returns a 403 status code
func (m *CampaignsItemSequencesItemContactsRequestBuilder) Post(ctx context.Context, body CampaignsItemSequencesItemContactsPostRequestBodyable, requestConfiguration *CampaignsItemSequencesItemContactsRequestBuilderPostRequestConfiguration) (CampaignsItemSequencesItemContactsResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateCampaignsItemSequencesItemContacts401ErrorFromDiscriminatorValue,
		"403": CreateCampaignsItemSequencesItemContacts403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCampaignsItemSequencesItemContactsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(CampaignsItemSequencesItemContactsResponseable), nil
}

// PostAsContactsPostResponse adds a list of contacts to a campaign sequenceResponse contains a map of the provided list of Contact Ids related to their individual result.
// returns a CampaignsItemSequencesItemContactsPostResponseable when successful
// returns a CampaignsItemSequencesItemContacts401Error error when the service returns a 401 status code
// returns a CampaignsItemSequencesItemContacts403Error error when the service returns a 403 status code
func (m *CampaignsItemSequencesItemContactsRequestBuilder) PostAsContactsPostResponse(ctx context.Context, body CampaignsItemSequencesItemContactsPostRequestBodyable, requestConfiguration *CampaignsItemSequencesItemContactsRequestBuilderPostRequestConfiguration) (CampaignsItemSequencesItemContactsPostResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateCampaignsItemSequencesItemContacts401ErrorFromDiscriminatorValue,
		"403": CreateCampaignsItemSequencesItemContacts403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCampaignsItemSequencesItemContactsPostResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(CampaignsItemSequencesItemContactsPostResponseable), nil
}

// ToDeleteRequestInformation removes a list of contacts from a campaign sequence
// returns a *RequestInformation when successful
func (m *CampaignsItemSequencesItemContactsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, body CampaignsItemSequencesItemContactsDeleteRequestBodyable, requestConfiguration *CampaignsItemSequencesItemContactsRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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

// ToPostRequestInformation adds a list of contacts to a campaign sequenceResponse contains a map of the provided list of Contact Ids related to their individual result.
// returns a *RequestInformation when successful
func (m *CampaignsItemSequencesItemContactsRequestBuilder) ToPostRequestInformation(ctx context.Context, body CampaignsItemSequencesItemContactsPostRequestBodyable, requestConfiguration *CampaignsItemSequencesItemContactsRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CampaignsItemSequencesItemContactsRequestBuilder when successful
func (m *CampaignsItemSequencesItemContactsRequestBuilder) WithUrl(rawUrl string) *CampaignsItemSequencesItemContactsRequestBuilder {
	return NewCampaignsItemSequencesItemContactsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
