package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CampaignsWithCampaignItemRequestBuilder builds and executes requests for operations under \v1\campaigns\{campaignId}
type CampaignsWithCampaignItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// CampaignsWithCampaignItemRequestBuilderGetQueryParameters retrieves a single campaign
type CampaignsWithCampaignItemRequestBuilderGetQueryParameters struct {
	// Comma-delimited list of Campaign properties to include in the response. (The fields `goals` and `sequences` aren't included, by default.)
	Optional_properties []string `uriparametername:"optional_properties"`
}

// CampaignsWithCampaignItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CampaignsWithCampaignItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *CampaignsWithCampaignItemRequestBuilderGetQueryParameters
}

// NewCampaignsWithCampaignItemRequestBuilderInternal instantiates a new CampaignsWithCampaignItemRequestBuilder and sets the default values.
func NewCampaignsWithCampaignItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CampaignsWithCampaignItemRequestBuilder {
	m := &CampaignsWithCampaignItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/campaigns/{campaignId}{?optional_properties*}", pathParameters),
	}
	return m
}

// NewCampaignsWithCampaignItemRequestBuilder instantiates a new CampaignsWithCampaignItemRequestBuilder and sets the default values.
func NewCampaignsWithCampaignItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CampaignsWithCampaignItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewCampaignsWithCampaignItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieves a single campaign
// Deprecated: This method is obsolete. Use GetAsWithCampaignGetResponse instead.
// returns a CampaignsItemWithCampaignResponseable when successful
// returns a CampaignsItemWithCampaign401Error error when the service returns a 401 status code
// returns a CampaignsItemWithCampaign403Error error when the service returns a 403 status code
// returns a CampaignsItemWithCampaign404Error error when the service returns a 404 status code
func (m *CampaignsWithCampaignItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CampaignsWithCampaignItemRequestBuilderGetRequestConfiguration) (CampaignsItemWithCampaignResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateCampaignsItemWithCampaign401ErrorFromDiscriminatorValue,
		"403": CreateCampaignsItemWithCampaign403ErrorFromDiscriminatorValue,
		"404": CreateCampaignsItemWithCampaign404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCampaignsItemWithCampaignResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(CampaignsItemWithCampaignResponseable), nil
}

// GetAsWithCampaignGetResponse retrieves a single campaign
// returns a CampaignsItemWithCampaignGetResponseable when successful
// returns a CampaignsItemWithCampaign401Error error when the service returns a 401 status code
// returns a CampaignsItemWithCampaign403Error error when the service returns a 403 status code
// returns a CampaignsItemWithCampaign404Error error when the service returns a 404 status code
func (m *CampaignsWithCampaignItemRequestBuilder) GetAsWithCampaignGetResponse(ctx context.Context, requestConfiguration *CampaignsWithCampaignItemRequestBuilderGetRequestConfiguration) (CampaignsItemWithCampaignGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateCampaignsItemWithCampaign401ErrorFromDiscriminatorValue,
		"403": CreateCampaignsItemWithCampaign403ErrorFromDiscriminatorValue,
		"404": CreateCampaignsItemWithCampaign404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCampaignsItemWithCampaignGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(CampaignsItemWithCampaignGetResponseable), nil
}

// Sequences the sequences property
// returns a *CampaignsItemSequencesRequestBuilder when successful
func (m *CampaignsWithCampaignItemRequestBuilder) Sequences() *CampaignsItemSequencesRequestBuilder {
	return NewCampaignsItemSequencesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// ToGetRequestInformation retrieves a single campaign
// returns a *RequestInformation when successful
func (m *CampaignsWithCampaignItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CampaignsWithCampaignItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		if requestConfiguration.QueryParameters != nil {
			requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
		}
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CampaignsWithCampaignItemRequestBuilder when successful
func (m *CampaignsWithCampaignItemRequestBuilder) WithUrl(rawUrl string) *CampaignsWithCampaignItemRequestBuilder {
	return NewCampaignsWithCampaignItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
