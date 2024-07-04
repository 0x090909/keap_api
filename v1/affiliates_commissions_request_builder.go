package v1

import (
	"context"
	i88b4ee534beb69eba8d14dbf18bce87a23c5e7447a9644702223eb405798b7c2 "github.com/0x090909/keap_api/v1/affiliates/commissions"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AffiliatesCommissionsRequestBuilder builds and executes requests for operations under \v1\affiliates\commissions
type AffiliatesCommissionsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// AffiliatesCommissionsRequestBuilderGetQueryParameters retrieve a list of Commissions based on Affiliate or Date Range
type AffiliatesCommissionsRequestBuilderGetQueryParameters struct {
	// Affiliate to retrieve commissions for
	AffiliateId *int64 `uriparametername:"affiliateId"`
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
	Order_directionAsGetOrder_directionQueryParameterType *i88b4ee534beb69eba8d14dbf18bce87a23c5e7447a9644702223eb405798b7c2.GetOrder_directionQueryParameterType `uriparametername:"order_direction"`
	// Attribute to order items by
	OrderAsGetOrderQueryParameterType *i88b4ee534beb69eba8d14dbf18bce87a23c5e7447a9644702223eb405798b7c2.GetOrderQueryParameterType `uriparametername:"order"`
	// Date to start searching from ex. `2017-01-01T22:17:59.039Z`
	Since *string `uriparametername:"since"`
	// Date to search to ex. `2017-01-01T22:17:59.039Z`
	Until *string `uriparametername:"until"`
}

// AffiliatesCommissionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AffiliatesCommissionsRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *AffiliatesCommissionsRequestBuilderGetQueryParameters
}

// NewAffiliatesCommissionsRequestBuilderInternal instantiates a new AffiliatesCommissionsRequestBuilder and sets the default values.
func NewAffiliatesCommissionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesCommissionsRequestBuilder {
	m := &AffiliatesCommissionsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/affiliates/commissions{?affiliateId*,limit*,offset*,order*,order_direction*,since*,until*}", pathParameters),
	}
	return m
}

// NewAffiliatesCommissionsRequestBuilder instantiates a new AffiliatesCommissionsRequestBuilder and sets the default values.
func NewAffiliatesCommissionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesCommissionsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewAffiliatesCommissionsRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieve a list of Commissions based on Affiliate or Date Range
// Deprecated: This method is obsolete. Use GetAsCommissionsGetResponse instead.
// returns a AffiliatesCommissionsResponseable when successful
// returns a AffiliatesCommissions401Error error when the service returns a 401 status code
// returns a AffiliatesCommissions403Error error when the service returns a 403 status code
// returns a AffiliatesCommissions404Error error when the service returns a 404 status code
func (m *AffiliatesCommissionsRequestBuilder) Get(ctx context.Context, requestConfiguration *AffiliatesCommissionsRequestBuilderGetRequestConfiguration) (AffiliatesCommissionsResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAffiliatesCommissions401ErrorFromDiscriminatorValue,
		"403": CreateAffiliatesCommissions403ErrorFromDiscriminatorValue,
		"404": CreateAffiliatesCommissions404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAffiliatesCommissionsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AffiliatesCommissionsResponseable), nil
}

// GetAsCommissionsGetResponse retrieve a list of Commissions based on Affiliate or Date Range
// returns a AffiliatesCommissionsGetResponseable when successful
// returns a AffiliatesCommissions401Error error when the service returns a 401 status code
// returns a AffiliatesCommissions403Error error when the service returns a 403 status code
// returns a AffiliatesCommissions404Error error when the service returns a 404 status code
func (m *AffiliatesCommissionsRequestBuilder) GetAsCommissionsGetResponse(ctx context.Context, requestConfiguration *AffiliatesCommissionsRequestBuilderGetRequestConfiguration) (AffiliatesCommissionsGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAffiliatesCommissions401ErrorFromDiscriminatorValue,
		"403": CreateAffiliatesCommissions403ErrorFromDiscriminatorValue,
		"404": CreateAffiliatesCommissions404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAffiliatesCommissionsGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AffiliatesCommissionsGetResponseable), nil
}

// ToGetRequestInformation retrieve a list of Commissions based on Affiliate or Date Range
// returns a *RequestInformation when successful
func (m *AffiliatesCommissionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AffiliatesCommissionsRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AffiliatesCommissionsRequestBuilder when successful
func (m *AffiliatesCommissionsRequestBuilder) WithUrl(rawUrl string) *AffiliatesCommissionsRequestBuilder {
	return NewAffiliatesCommissionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
