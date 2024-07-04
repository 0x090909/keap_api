package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CampaignsGoalsItemWithCallNameItemRequestBuilder builds and executes requests for operations under \v1\campaigns\goals\{integration}\{callName}
type CampaignsGoalsItemWithCallNameItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// CampaignsGoalsItemWithCallNameItemRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CampaignsGoalsItemWithCallNameItemRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewCampaignsGoalsItemWithCallNameItemRequestBuilderInternal instantiates a new CampaignsGoalsItemWithCallNameItemRequestBuilder and sets the default values.
func NewCampaignsGoalsItemWithCallNameItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CampaignsGoalsItemWithCallNameItemRequestBuilder {
	m := &CampaignsGoalsItemWithCallNameItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/campaigns/goals/{integration}/{callName}", pathParameters),
	}
	return m
}

// NewCampaignsGoalsItemWithCallNameItemRequestBuilder instantiates a new CampaignsGoalsItemWithCallNameItemRequestBuilder and sets the default values.
func NewCampaignsGoalsItemWithCallNameItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CampaignsGoalsItemWithCallNameItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewCampaignsGoalsItemWithCallNameItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Post achieves API goal for campaigns with matching integration, callName for a given contactId
// returns a []CampaignsGoalsItemItemWithCallNameable when successful
// returns a CampaignsGoalsItemItemWithCallName401Error error when the service returns a 401 status code
// returns a CampaignsGoalsItemItemWithCallName403Error error when the service returns a 403 status code
func (m *CampaignsGoalsItemWithCallNameItemRequestBuilder) Post(ctx context.Context, body CampaignsGoalsItemItemWithCallNamePostRequestBodyable, requestConfiguration *CampaignsGoalsItemWithCallNameItemRequestBuilderPostRequestConfiguration) ([]CampaignsGoalsItemItemWithCallNameable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateCampaignsGoalsItemItemWithCallName401ErrorFromDiscriminatorValue,
		"403": CreateCampaignsGoalsItemItemWithCallName403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, CreateCampaignsGoalsItemItemWithCallNameFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	val := make([]CampaignsGoalsItemItemWithCallNameable, len(res))
	for i, v := range res {
		if v != nil {
			val[i] = v.(CampaignsGoalsItemItemWithCallNameable)
		}
	}
	return val, nil
}

// ToPostRequestInformation achieves API goal for campaigns with matching integration, callName for a given contactId
// returns a *RequestInformation when successful
func (m *CampaignsGoalsItemWithCallNameItemRequestBuilder) ToPostRequestInformation(ctx context.Context, body CampaignsGoalsItemItemWithCallNamePostRequestBodyable, requestConfiguration *CampaignsGoalsItemWithCallNameItemRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CampaignsGoalsItemWithCallNameItemRequestBuilder when successful
func (m *CampaignsGoalsItemWithCallNameItemRequestBuilder) WithUrl(rawUrl string) *CampaignsGoalsItemWithCallNameItemRequestBuilder {
	return NewCampaignsGoalsItemWithCallNameItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
