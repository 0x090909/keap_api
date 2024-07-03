package v2

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CampaignsWithCampaign_ItemRequestBuilder builds and executes requests for operations under \v2\campaigns\{campaign_id}
type CampaignsWithCampaign_ItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// CampaignsWithCampaign_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CampaignsWithCampaign_ItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewCampaignsWithCampaign_ItemRequestBuilderInternal instantiates a new CampaignsWithCampaign_ItemRequestBuilder and sets the default values.
func NewCampaignsWithCampaign_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CampaignsWithCampaign_ItemRequestBuilder {
	m := &CampaignsWithCampaign_ItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/campaigns/{campaign_id}", pathParameters),
	}
	return m
}

// NewCampaignsWithCampaign_ItemRequestBuilder instantiates a new CampaignsWithCampaign_ItemRequestBuilder and sets the default values.
func NewCampaignsWithCampaign_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CampaignsWithCampaign_ItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewCampaignsWithCampaign_ItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieves a single campaign
// Deprecated: This method is obsolete. Use GetAsWithCampaign_GetResponse instead.
// returns a CampaignsItemWithCampaign_Responseable when successful
// returns a CampaignsItemWithCampaign_401Error error when the service returns a 401 status code
// returns a CampaignsItemWithCampaign_403Error error when the service returns a 403 status code
// returns a CampaignsItemWithCampaign_404Error error when the service returns a 404 status code
func (m *CampaignsWithCampaign_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CampaignsWithCampaign_ItemRequestBuilderGetRequestConfiguration) (CampaignsItemWithCampaign_Responseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateCampaignsItemWithCampaign_401ErrorFromDiscriminatorValue,
		"403": CreateCampaignsItemWithCampaign_403ErrorFromDiscriminatorValue,
		"404": CreateCampaignsItemWithCampaign_404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCampaignsItemWithCampaign_ResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(CampaignsItemWithCampaign_Responseable), nil
}

// GetAsWithCampaign_GetResponse retrieves a single campaign
// returns a CampaignsItemWithCampaign_GetResponseable when successful
// returns a CampaignsItemWithCampaign_401Error error when the service returns a 401 status code
// returns a CampaignsItemWithCampaign_403Error error when the service returns a 403 status code
// returns a CampaignsItemWithCampaign_404Error error when the service returns a 404 status code
func (m *CampaignsWithCampaign_ItemRequestBuilder) GetAsWithCampaign_GetResponse(ctx context.Context, requestConfiguration *CampaignsWithCampaign_ItemRequestBuilderGetRequestConfiguration) (CampaignsItemWithCampaign_GetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateCampaignsItemWithCampaign_401ErrorFromDiscriminatorValue,
		"403": CreateCampaignsItemWithCampaign_403ErrorFromDiscriminatorValue,
		"404": CreateCampaignsItemWithCampaign_404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCampaignsItemWithCampaign_GetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(CampaignsItemWithCampaign_GetResponseable), nil
}

// Sequences the sequences property
// returns a *CampaignsItemSequencesRequestBuilder when successful
func (m *CampaignsWithCampaign_ItemRequestBuilder) Sequences() *CampaignsItemSequencesRequestBuilder {
	return NewCampaignsItemSequencesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// ToGetRequestInformation retrieves a single campaign
// returns a *RequestInformation when successful
func (m *CampaignsWithCampaign_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CampaignsWithCampaign_ItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CampaignsWithCampaign_ItemRequestBuilder when successful
func (m *CampaignsWithCampaign_ItemRequestBuilder) WithUrl(rawUrl string) *CampaignsWithCampaign_ItemRequestBuilder {
	return NewCampaignsWithCampaign_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
