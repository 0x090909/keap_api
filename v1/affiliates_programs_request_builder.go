package v1

import (
	"context"
	ieb560f29f4c0ba79720ce4afb2a705967c63dfde524ca674cb5e40a406ba761e "github.com/0x090909/keap_api/v1/affiliates/programs"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AffiliatesProgramsRequestBuilder builds and executes requests for operations under \v1\affiliates\programs
type AffiliatesProgramsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// AffiliatesProgramsRequestBuilderGetQueryParameters retrieve a list of Commission Programs
type AffiliatesProgramsRequestBuilderGetQueryParameters struct {
	// Id of the affiliate you would like programs for
	Affiliate_id *int64 `uriparametername:"affiliate_id"`
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
	Order_directionAsGetOrder_directionQueryParameterType *ieb560f29f4c0ba79720ce4afb2a705967c63dfde524ca674cb5e40a406ba761e.GetOrder_directionQueryParameterType `uriparametername:"order_direction"`
	// Attribute to order items by
	OrderAsGetOrderQueryParameterType *ieb560f29f4c0ba79720ce4afb2a705967c63dfde524ca674cb5e40a406ba761e.GetOrderQueryParameterType `uriparametername:"order"`
}

// AffiliatesProgramsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AffiliatesProgramsRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *AffiliatesProgramsRequestBuilderGetQueryParameters
}

// NewAffiliatesProgramsRequestBuilderInternal instantiates a new AffiliatesProgramsRequestBuilder and sets the default values.
func NewAffiliatesProgramsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesProgramsRequestBuilder {
	m := &AffiliatesProgramsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/affiliates/programs{?affiliate_id*,limit*,offset*,order*,order_direction*}", pathParameters),
	}
	return m
}

// NewAffiliatesProgramsRequestBuilder instantiates a new AffiliatesProgramsRequestBuilder and sets the default values.
func NewAffiliatesProgramsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesProgramsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewAffiliatesProgramsRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieve a list of Commission Programs
// Deprecated: This method is obsolete. Use GetAsProgramsGetResponse instead.
// returns a AffiliatesProgramsResponseable when successful
// returns a AffiliatesPrograms401Error error when the service returns a 401 status code
// returns a AffiliatesPrograms403Error error when the service returns a 403 status code
// returns a AffiliatesPrograms404Error error when the service returns a 404 status code
func (m *AffiliatesProgramsRequestBuilder) Get(ctx context.Context, requestConfiguration *AffiliatesProgramsRequestBuilderGetRequestConfiguration) (AffiliatesProgramsResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAffiliatesPrograms401ErrorFromDiscriminatorValue,
		"403": CreateAffiliatesPrograms403ErrorFromDiscriminatorValue,
		"404": CreateAffiliatesPrograms404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAffiliatesProgramsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AffiliatesProgramsResponseable), nil
}

// GetAsProgramsGetResponse retrieve a list of Commission Programs
// returns a AffiliatesProgramsGetResponseable when successful
// returns a AffiliatesPrograms401Error error when the service returns a 401 status code
// returns a AffiliatesPrograms403Error error when the service returns a 403 status code
// returns a AffiliatesPrograms404Error error when the service returns a 404 status code
func (m *AffiliatesProgramsRequestBuilder) GetAsProgramsGetResponse(ctx context.Context, requestConfiguration *AffiliatesProgramsRequestBuilderGetRequestConfiguration) (AffiliatesProgramsGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAffiliatesPrograms401ErrorFromDiscriminatorValue,
		"403": CreateAffiliatesPrograms403ErrorFromDiscriminatorValue,
		"404": CreateAffiliatesPrograms404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAffiliatesProgramsGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AffiliatesProgramsGetResponseable), nil
}

// ToGetRequestInformation retrieve a list of Commission Programs
// returns a *RequestInformation when successful
func (m *AffiliatesProgramsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AffiliatesProgramsRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AffiliatesProgramsRequestBuilder when successful
func (m *AffiliatesProgramsRequestBuilder) WithUrl(rawUrl string) *AffiliatesProgramsRequestBuilder {
	return NewAffiliatesProgramsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
