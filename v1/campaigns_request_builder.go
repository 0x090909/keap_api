package v1

import (
	"context"
	ic07fa4cb1af95394c3cca5ec426de6ee372cdfb80683fe6ef39419b451b07c89 "github.com/0x090909/keap_api/v1/campaigns"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
)

// CampaignsRequestBuilder builds and executes requests for operations under \v1\campaigns
type CampaignsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// CampaignsRequestBuilderGetQueryParameters retrieves all campaigns for the authenticated user
type CampaignsRequestBuilderGetQueryParameters struct {
	// Sets a total of items to return
	Limit *int32 `uriparametername:"limit"`
	// Sets a beginning range of items to return
	Offset *int32 `uriparametername:"offset"`
	// Attribute to order items by
	// Deprecated: This property is deprecated, use OrderAsGetOrderQueryParameterType instead
	Order *string `uriparametername:"order"`
	// How to order the data i.e. ascending (A-Z) or descending (Z-A)
	// Deprecated: This property is deprecated, use Order_directionAsGetOrder_directionQueryParameterType instead
	Order_direction *string `uriparametername:"order_direction"`
	// How to order the data i.e. ascending (A-Z) or descending (Z-A)
	Order_directionAsGetOrder_directionQueryParameterType *ic07fa4cb1af95394c3cca5ec426de6ee372cdfb80683fe6ef39419b451b07c89.GetOrder_directionQueryParameterType `uriparametername:"order_direction"`
	// Attribute to order items by
	OrderAsGetOrderQueryParameterType *ic07fa4cb1af95394c3cca5ec426de6ee372cdfb80683fe6ef39419b451b07c89.GetOrderQueryParameterType `uriparametername:"order"`
	// Optional text to search
	Search_text *string `uriparametername:"search_text"`
}

// CampaignsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CampaignsRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *CampaignsRequestBuilderGetQueryParameters
}

// ByCampaignId gets an item from the github.com/0x090909/keap_api.v1.campaigns.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *CampaignsWithCampaignItemRequestBuilder when successful
func (m *CampaignsRequestBuilder) ByCampaignId(campaignId string) *CampaignsWithCampaignItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if campaignId != "" {
		urlTplParams["campaignId"] = campaignId
	}
	return NewCampaignsWithCampaignItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// ByCampaignIdInt64 gets an item from the github.com/0x090909/keap_api.v1.campaigns.item collection
// returns a *CampaignsWithCampaignItemRequestBuilder when successful
func (m *CampaignsRequestBuilder) ByCampaignIdInt64(campaignId int64) *CampaignsWithCampaignItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	urlTplParams["campaignId"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(campaignId, 10)
	return NewCampaignsWithCampaignItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewCampaignsRequestBuilderInternal instantiates a new CampaignsRequestBuilder and sets the default values.
func NewCampaignsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CampaignsRequestBuilder {
	m := &CampaignsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/campaigns{?limit*,offset*,order*,order_direction*,search_text*}", pathParameters),
	}
	return m
}

// NewCampaignsRequestBuilder instantiates a new CampaignsRequestBuilder and sets the default values.
func NewCampaignsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CampaignsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewCampaignsRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieves all campaigns for the authenticated user
// Deprecated: This method is obsolete. Use GetAsCampaignsGetResponse instead.
// returns a CampaignsResponseable when successful
// returns a Campaigns401Error error when the service returns a 401 status code
// returns a Campaigns403Error error when the service returns a 403 status code
// returns a Campaigns404Error error when the service returns a 404 status code
func (m *CampaignsRequestBuilder) Get(ctx context.Context, requestConfiguration *CampaignsRequestBuilderGetRequestConfiguration) (CampaignsResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateCampaigns401ErrorFromDiscriminatorValue,
		"403": CreateCampaigns403ErrorFromDiscriminatorValue,
		"404": CreateCampaigns404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCampaignsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(CampaignsResponseable), nil
}

// GetAsCampaignsGetResponse retrieves all campaigns for the authenticated user
// returns a CampaignsGetResponseable when successful
// returns a Campaigns401Error error when the service returns a 401 status code
// returns a Campaigns403Error error when the service returns a 403 status code
// returns a Campaigns404Error error when the service returns a 404 status code
func (m *CampaignsRequestBuilder) GetAsCampaignsGetResponse(ctx context.Context, requestConfiguration *CampaignsRequestBuilderGetRequestConfiguration) (CampaignsGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateCampaigns401ErrorFromDiscriminatorValue,
		"403": CreateCampaigns403ErrorFromDiscriminatorValue,
		"404": CreateCampaigns404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCampaignsGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(CampaignsGetResponseable), nil
}

// Goals the goals property
// returns a *CampaignsGoalsRequestBuilder when successful
func (m *CampaignsRequestBuilder) Goals() *CampaignsGoalsRequestBuilder {
	return NewCampaignsGoalsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// ToGetRequestInformation retrieves all campaigns for the authenticated user
// returns a *RequestInformation when successful
func (m *CampaignsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CampaignsRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CampaignsRequestBuilder when successful
func (m *CampaignsRequestBuilder) WithUrl(rawUrl string) *CampaignsRequestBuilder {
	return NewCampaignsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
