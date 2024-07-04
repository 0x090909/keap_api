package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CampaignsItemSequencesItemContactsWithContactItemRequestBuilder builds and executes requests for operations under \v1\campaigns\{campaignId}\sequences\{sequenceId}\contacts\{contactId}
type CampaignsItemSequencesItemContactsWithContactItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// CampaignsItemSequencesItemContactsWithContactItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CampaignsItemSequencesItemContactsWithContactItemRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// CampaignsItemSequencesItemContactsWithContactItemRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CampaignsItemSequencesItemContactsWithContactItemRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewCampaignsItemSequencesItemContactsWithContactItemRequestBuilderInternal instantiates a new CampaignsItemSequencesItemContactsWithContactItemRequestBuilder and sets the default values.
func NewCampaignsItemSequencesItemContactsWithContactItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CampaignsItemSequencesItemContactsWithContactItemRequestBuilder {
	m := &CampaignsItemSequencesItemContactsWithContactItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/campaigns/{campaignId}/sequences/{sequenceId}/contacts/{contactId}", pathParameters),
	}
	return m
}

// NewCampaignsItemSequencesItemContactsWithContactItemRequestBuilder instantiates a new CampaignsItemSequencesItemContactsWithContactItemRequestBuilder and sets the default values.
func NewCampaignsItemSequencesItemContactsWithContactItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CampaignsItemSequencesItemContactsWithContactItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewCampaignsItemSequencesItemContactsWithContactItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete removes a single contact from a campaign sequence
// returns a CampaignsItemSequencesItemContactsItemWithContact401Error error when the service returns a 401 status code
// returns a CampaignsItemSequencesItemContactsItemWithContact403Error error when the service returns a 403 status code
// returns a CampaignsItemSequencesItemContactsItemWithContact404Error error when the service returns a 404 status code
func (m *CampaignsItemSequencesItemContactsWithContactItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CampaignsItemSequencesItemContactsWithContactItemRequestBuilderDeleteRequestConfiguration) error {
	requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateCampaignsItemSequencesItemContactsItemWithContact401ErrorFromDiscriminatorValue,
		"403": CreateCampaignsItemSequencesItemContactsItemWithContact403ErrorFromDiscriminatorValue,
		"404": CreateCampaignsItemSequencesItemContactsItemWithContact404ErrorFromDiscriminatorValue,
	}
	err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
	if err != nil {
		return err
	}
	return nil
}

// Post adds a single contact to a campaign sequence
// returns a CampaignsItemSequencesItemContactsItemWithContact401Error error when the service returns a 401 status code
// returns a CampaignsItemSequencesItemContactsItemWithContact403Error error when the service returns a 403 status code
func (m *CampaignsItemSequencesItemContactsWithContactItemRequestBuilder) Post(ctx context.Context, requestConfiguration *CampaignsItemSequencesItemContactsWithContactItemRequestBuilderPostRequestConfiguration) error {
	requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateCampaignsItemSequencesItemContactsItemWithContact401ErrorFromDiscriminatorValue,
		"403": CreateCampaignsItemSequencesItemContactsItemWithContact403ErrorFromDiscriminatorValue,
	}
	err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
	if err != nil {
		return err
	}
	return nil
}

// ToDeleteRequestInformation removes a single contact from a campaign sequence
// returns a *RequestInformation when successful
func (m *CampaignsItemSequencesItemContactsWithContactItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CampaignsItemSequencesItemContactsWithContactItemRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// ToPostRequestInformation adds a single contact to a campaign sequence
// returns a *RequestInformation when successful
func (m *CampaignsItemSequencesItemContactsWithContactItemRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *CampaignsItemSequencesItemContactsWithContactItemRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CampaignsItemSequencesItemContactsWithContactItemRequestBuilder when successful
func (m *CampaignsItemSequencesItemContactsWithContactItemRequestBuilder) WithUrl(rawUrl string) *CampaignsItemSequencesItemContactsWithContactItemRequestBuilder {
	return NewCampaignsItemSequencesItemContactsWithContactItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
