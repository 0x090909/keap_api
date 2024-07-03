package v2

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CampaignsRequestBuilder builds and executes requests for operations under \v2\campaigns
type CampaignsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// CampaignsRequestBuilderGetQueryParameters retrieves all campaigns for the authenticated user
type CampaignsRequestBuilderGetQueryParameters struct {
	// Filter to apply, allowed fields are:  - (String) name  You will need to apply the `==` operator to check the equality of the filter with your searched text, in the encoded form `%3D%3D`. The search will look for the text anywhere in the campaign name.  - `filter=name%3D%3DSpring Campaign` - `filter=name%3D%3DTag New Contacts`
	Filter *string `uriparametername:"filter"`
	// Attribute and direction to order items.  One of the following fields:  - name  - publisheddate  - id  - completedContactCount  - activeContacts  - datecreated  - lastupdated  One of the following directions:  - asc  - desc
	Order_by *string `uriparametername:"order_by"`
	// Total number of items to return per page
	Page_size *int32 `uriparametername:"page_size"`
	// Page token
	Page_token *string `uriparametername:"page_token"`
	Stats      *bool   `uriparametername:"stats"`
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

// ByCampaign_id gets an item from the keapapi.v2.campaigns.item collection
// returns a *CampaignsWithCampaign_ItemRequestBuilder when successful
func (m *CampaignsRequestBuilder) ByCampaign_id(campaign_id string) *CampaignsWithCampaign_ItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if campaign_id != "" {
		urlTplParams["campaign_id"] = campaign_id
	}
	return NewCampaignsWithCampaign_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewCampaignsRequestBuilderInternal instantiates a new CampaignsRequestBuilder and sets the default values.
func NewCampaignsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CampaignsRequestBuilder {
	m := &CampaignsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/campaigns{?filter*,order_by*,page_size*,page_token*,stats*}", pathParameters),
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
