package v1

import (
	"context"
	ia420c74c00a76ace443752956756f82a4eb4ecbfd9bca5faaabf629e87812c7a "github.com/0x090909/keap_api/v1/affiliates/summaries"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AffiliatesSummariesRequestBuilder builds and executes requests for operations under \v1\affiliates\summaries
type AffiliatesSummariesRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// AffiliatesSummariesRequestBuilderGetQueryParameters retrieve a list of affiliate summaries
type AffiliatesSummariesRequestBuilderGetQueryParameters struct {
	// Ids of the affiliates you would like a summary for
	Affiliate_ids []int64 `uriparametername:"affiliate_ids"`
	// Attribute to order items by
	// Deprecated: This property is deprecated, use OrderAsGetOrderQueryParameterType instead
	Order *string `uriparametername:"order"`
	// How to order the data i.e. ascending (A-Z) or descending (Z-A)
	// Deprecated: This property is deprecated, use Order_directionAsGetOrder_directionQueryParameterType instead
	Order_direction *string `uriparametername:"order_direction"`
	// How to order the data i.e. ascending (A-Z) or descending (Z-A)
	Order_directionAsGetOrder_directionQueryParameterType *ia420c74c00a76ace443752956756f82a4eb4ecbfd9bca5faaabf629e87812c7a.GetOrder_directionQueryParameterType `uriparametername:"order_direction"`
	// Attribute to order items by
	OrderAsGetOrderQueryParameterType *ia420c74c00a76ace443752956756f82a4eb4ecbfd9bca5faaabf629e87812c7a.GetOrderQueryParameterType `uriparametername:"order"`
}

// AffiliatesSummariesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AffiliatesSummariesRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *AffiliatesSummariesRequestBuilderGetQueryParameters
}

// NewAffiliatesSummariesRequestBuilderInternal instantiates a new AffiliatesSummariesRequestBuilder and sets the default values.
func NewAffiliatesSummariesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesSummariesRequestBuilder {
	m := &AffiliatesSummariesRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/affiliates/summaries{?affiliate_ids*,order*,order_direction*}", pathParameters),
	}
	return m
}

// NewAffiliatesSummariesRequestBuilder instantiates a new AffiliatesSummariesRequestBuilder and sets the default values.
func NewAffiliatesSummariesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesSummariesRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewAffiliatesSummariesRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieve a list of affiliate summaries
// Deprecated: This method is obsolete. Use GetAsSummariesGetResponse instead.
// returns a AffiliatesSummariesResponseable when successful
// returns a AffiliatesSummaries401Error error when the service returns a 401 status code
// returns a AffiliatesSummaries403Error error when the service returns a 403 status code
// returns a AffiliatesSummaries404Error error when the service returns a 404 status code
func (m *AffiliatesSummariesRequestBuilder) Get(ctx context.Context, requestConfiguration *AffiliatesSummariesRequestBuilderGetRequestConfiguration) (AffiliatesSummariesResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAffiliatesSummaries401ErrorFromDiscriminatorValue,
		"403": CreateAffiliatesSummaries403ErrorFromDiscriminatorValue,
		"404": CreateAffiliatesSummaries404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAffiliatesSummariesResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AffiliatesSummariesResponseable), nil
}

// GetAsSummariesGetResponse retrieve a list of affiliate summaries
// returns a AffiliatesSummariesGetResponseable when successful
// returns a AffiliatesSummaries401Error error when the service returns a 401 status code
// returns a AffiliatesSummaries403Error error when the service returns a 403 status code
// returns a AffiliatesSummaries404Error error when the service returns a 404 status code
func (m *AffiliatesSummariesRequestBuilder) GetAsSummariesGetResponse(ctx context.Context, requestConfiguration *AffiliatesSummariesRequestBuilderGetRequestConfiguration) (AffiliatesSummariesGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAffiliatesSummaries401ErrorFromDiscriminatorValue,
		"403": CreateAffiliatesSummaries403ErrorFromDiscriminatorValue,
		"404": CreateAffiliatesSummaries404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAffiliatesSummariesGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AffiliatesSummariesGetResponseable), nil
}

// ToGetRequestInformation retrieve a list of affiliate summaries
// returns a *RequestInformation when successful
func (m *AffiliatesSummariesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AffiliatesSummariesRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AffiliatesSummariesRequestBuilder when successful
func (m *AffiliatesSummariesRequestBuilder) WithUrl(rawUrl string) *AffiliatesSummariesRequestBuilder {
	return NewAffiliatesSummariesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
