package v1

import (
	"context"
	ie80f4c9c7b122b60ab61bbd04026e6ff16a44c89caa0e21a7bc142d0db06a6ff "github.com/0x090909/keap_api/v1/affiliates/redirectlinks"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AffiliatesRedirectlinksRequestBuilder builds and executes requests for operations under \v1\affiliates\redirectlinks
type AffiliatesRedirectlinksRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// AffiliatesRedirectlinksRequestBuilderGetQueryParameters retrieves a list of all affiliate redirects
type AffiliatesRedirectlinksRequestBuilderGetQueryParameters struct {
	// Affiliate ID to search by
	Affiliate_id *string `uriparametername:"affiliate_id"`
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
	Order_directionAsGetOrder_directionQueryParameterType *ie80f4c9c7b122b60ab61bbd04026e6ff16a44c89caa0e21a7bc142d0db06a6ff.GetOrder_directionQueryParameterType `uriparametername:"order_direction"`
	// Attribute to order items by
	OrderAsGetOrderQueryParameterType *ie80f4c9c7b122b60ab61bbd04026e6ff16a44c89caa0e21a7bc142d0db06a6ff.GetOrderQueryParameterType `uriparametername:"order"`
}

// AffiliatesRedirectlinksRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AffiliatesRedirectlinksRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *AffiliatesRedirectlinksRequestBuilderGetQueryParameters
}

// NewAffiliatesRedirectlinksRequestBuilderInternal instantiates a new AffiliatesRedirectlinksRequestBuilder and sets the default values.
func NewAffiliatesRedirectlinksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesRedirectlinksRequestBuilder {
	m := &AffiliatesRedirectlinksRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/affiliates/redirectlinks{?affiliate_id*,limit*,offset*,order*,order_direction*}", pathParameters),
	}
	return m
}

// NewAffiliatesRedirectlinksRequestBuilder instantiates a new AffiliatesRedirectlinksRequestBuilder and sets the default values.
func NewAffiliatesRedirectlinksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesRedirectlinksRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewAffiliatesRedirectlinksRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieves a list of all affiliate redirects
// Deprecated: This method is obsolete. Use GetAsRedirectlinksGetResponse instead.
// returns a AffiliatesRedirectlinksResponseable when successful
// returns a AffiliatesRedirectlinks401Error error when the service returns a 401 status code
// returns a AffiliatesRedirectlinks403Error error when the service returns a 403 status code
// returns a AffiliatesRedirectlinks404Error error when the service returns a 404 status code
func (m *AffiliatesRedirectlinksRequestBuilder) Get(ctx context.Context, requestConfiguration *AffiliatesRedirectlinksRequestBuilderGetRequestConfiguration) (AffiliatesRedirectlinksResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAffiliatesRedirectlinks401ErrorFromDiscriminatorValue,
		"403": CreateAffiliatesRedirectlinks403ErrorFromDiscriminatorValue,
		"404": CreateAffiliatesRedirectlinks404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAffiliatesRedirectlinksResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AffiliatesRedirectlinksResponseable), nil
}

// GetAsRedirectlinksGetResponse retrieves a list of all affiliate redirects
// returns a AffiliatesRedirectlinksGetResponseable when successful
// returns a AffiliatesRedirectlinks401Error error when the service returns a 401 status code
// returns a AffiliatesRedirectlinks403Error error when the service returns a 403 status code
// returns a AffiliatesRedirectlinks404Error error when the service returns a 404 status code
func (m *AffiliatesRedirectlinksRequestBuilder) GetAsRedirectlinksGetResponse(ctx context.Context, requestConfiguration *AffiliatesRedirectlinksRequestBuilderGetRequestConfiguration) (AffiliatesRedirectlinksGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAffiliatesRedirectlinks401ErrorFromDiscriminatorValue,
		"403": CreateAffiliatesRedirectlinks403ErrorFromDiscriminatorValue,
		"404": CreateAffiliatesRedirectlinks404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAffiliatesRedirectlinksGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AffiliatesRedirectlinksGetResponseable), nil
}

// ToGetRequestInformation retrieves a list of all affiliate redirects
// returns a *RequestInformation when successful
func (m *AffiliatesRedirectlinksRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AffiliatesRedirectlinksRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AffiliatesRedirectlinksRequestBuilder when successful
func (m *AffiliatesRedirectlinksRequestBuilder) WithUrl(rawUrl string) *AffiliatesRedirectlinksRequestBuilder {
	return NewAffiliatesRedirectlinksRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
