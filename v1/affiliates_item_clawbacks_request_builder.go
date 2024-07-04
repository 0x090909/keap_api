package v1

import (
	"context"
	i591ed3eae84e6b84dfb7e65ef05d8057f5c84c63892dbfd55db679eab76dc7ba "github.com/0x090909/keap_api/v1/affiliates/item/clawbacks"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AffiliatesItemClawbacksRequestBuilder builds and executes requests for operations under \v1\affiliates\{affiliateId-id}\clawbacks
type AffiliatesItemClawbacksRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// AffiliatesItemClawbacksRequestBuilderGetQueryParameters retrieves a list of all affiliate clawbacks
type AffiliatesItemClawbacksRequestBuilderGetQueryParameters struct {
	// Sets a beginning range of items to return
	Limit *int32 `uriparametername:"limit"`
	// Count to offset the returned results by
	Offset *int32 `uriparametername:"offset"`
	// Attribute to order items by
	// Deprecated: This property is deprecated, use OrderAsGetOrderQueryParameterType instead
	Order *string `uriparametername:"order"`
	// How to order the data i.e. ascending (A-Z) or descending (Z-A)
	// Deprecated: This property is deprecated, use Order_directionAsGetOrder_directionQueryParameterType instead
	Order_direction *string `uriparametername:"order_direction"`
	// How to order the data i.e. ascending (A-Z) or descending (Z-A)
	Order_directionAsGetOrder_directionQueryParameterType *i591ed3eae84e6b84dfb7e65ef05d8057f5c84c63892dbfd55db679eab76dc7ba.GetOrder_directionQueryParameterType `uriparametername:"order_direction"`
	// Attribute to order items by
	OrderAsGetOrderQueryParameterType *i591ed3eae84e6b84dfb7e65ef05d8057f5c84c63892dbfd55db679eab76dc7ba.GetOrderQueryParameterType `uriparametername:"order"`
	// Optional inclusive start date
	Since *string `uriparametername:"since"`
	// Optional inclusive end date
	Until *string `uriparametername:"until"`
}

// AffiliatesItemClawbacksRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AffiliatesItemClawbacksRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *AffiliatesItemClawbacksRequestBuilderGetQueryParameters
}

// NewAffiliatesItemClawbacksRequestBuilderInternal instantiates a new AffiliatesItemClawbacksRequestBuilder and sets the default values.
func NewAffiliatesItemClawbacksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesItemClawbacksRequestBuilder {
	m := &AffiliatesItemClawbacksRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/affiliates/{affiliateId%2Did}/clawbacks{?limit*,offset*,order*,order_direction*,since*,until*}", pathParameters),
	}
	return m
}

// NewAffiliatesItemClawbacksRequestBuilder instantiates a new AffiliatesItemClawbacksRequestBuilder and sets the default values.
func NewAffiliatesItemClawbacksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesItemClawbacksRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewAffiliatesItemClawbacksRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieves a list of all affiliate clawbacks
// Deprecated: This method is obsolete. Use GetAsClawbacksGetResponse instead.
// returns a AffiliatesItemClawbacksResponseable when successful
// returns a AffiliatesItemClawbacks401Error error when the service returns a 401 status code
// returns a AffiliatesItemClawbacks403Error error when the service returns a 403 status code
// returns a AffiliatesItemClawbacks404Error error when the service returns a 404 status code
func (m *AffiliatesItemClawbacksRequestBuilder) Get(ctx context.Context, requestConfiguration *AffiliatesItemClawbacksRequestBuilderGetRequestConfiguration) (AffiliatesItemClawbacksResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAffiliatesItemClawbacks401ErrorFromDiscriminatorValue,
		"403": CreateAffiliatesItemClawbacks403ErrorFromDiscriminatorValue,
		"404": CreateAffiliatesItemClawbacks404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAffiliatesItemClawbacksResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AffiliatesItemClawbacksResponseable), nil
}

// GetAsClawbacksGetResponse retrieves a list of all affiliate clawbacks
// returns a AffiliatesItemClawbacksGetResponseable when successful
// returns a AffiliatesItemClawbacks401Error error when the service returns a 401 status code
// returns a AffiliatesItemClawbacks403Error error when the service returns a 403 status code
// returns a AffiliatesItemClawbacks404Error error when the service returns a 404 status code
func (m *AffiliatesItemClawbacksRequestBuilder) GetAsClawbacksGetResponse(ctx context.Context, requestConfiguration *AffiliatesItemClawbacksRequestBuilderGetRequestConfiguration) (AffiliatesItemClawbacksGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAffiliatesItemClawbacks401ErrorFromDiscriminatorValue,
		"403": CreateAffiliatesItemClawbacks403ErrorFromDiscriminatorValue,
		"404": CreateAffiliatesItemClawbacks404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAffiliatesItemClawbacksGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AffiliatesItemClawbacksGetResponseable), nil
}

// ToGetRequestInformation retrieves a list of all affiliate clawbacks
// returns a *RequestInformation when successful
func (m *AffiliatesItemClawbacksRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AffiliatesItemClawbacksRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AffiliatesItemClawbacksRequestBuilder when successful
func (m *AffiliatesItemClawbacksRequestBuilder) WithUrl(rawUrl string) *AffiliatesItemClawbacksRequestBuilder {
	return NewAffiliatesItemClawbacksRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
