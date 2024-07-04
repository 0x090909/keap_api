package v1

import (
	"context"
	i95833f6c8aad38d29938cc30301705eea017611ed729b69ffa00f8f7b008909d "github.com/0x090909/keap_api/v1/affiliates/item/payments"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AffiliatesItemPaymentsRequestBuilder builds and executes requests for operations under \v1\affiliates\{affiliateId-id}\payments
type AffiliatesItemPaymentsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// AffiliatesItemPaymentsRequestBuilderGetQueryParameters retrieves a list of all affiliate payments
type AffiliatesItemPaymentsRequestBuilderGetQueryParameters struct {
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
	Order_directionAsGetOrder_directionQueryParameterType *i95833f6c8aad38d29938cc30301705eea017611ed729b69ffa00f8f7b008909d.GetOrder_directionQueryParameterType `uriparametername:"order_direction"`
	// Attribute to order items by
	OrderAsGetOrderQueryParameterType *i95833f6c8aad38d29938cc30301705eea017611ed729b69ffa00f8f7b008909d.GetOrderQueryParameterType `uriparametername:"order"`
	// Inclusive start date
	Since *string `uriparametername:"since"`
	// Inclusive end date
	Until *string `uriparametername:"until"`
}

// AffiliatesItemPaymentsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AffiliatesItemPaymentsRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *AffiliatesItemPaymentsRequestBuilderGetQueryParameters
}

// NewAffiliatesItemPaymentsRequestBuilderInternal instantiates a new AffiliatesItemPaymentsRequestBuilder and sets the default values.
func NewAffiliatesItemPaymentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesItemPaymentsRequestBuilder {
	m := &AffiliatesItemPaymentsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/affiliates/{affiliateId%2Did}/payments{?limit*,offset*,order*,order_direction*,since*,until*}", pathParameters),
	}
	return m
}

// NewAffiliatesItemPaymentsRequestBuilder instantiates a new AffiliatesItemPaymentsRequestBuilder and sets the default values.
func NewAffiliatesItemPaymentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesItemPaymentsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewAffiliatesItemPaymentsRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieves a list of all affiliate payments
// Deprecated: This method is obsolete. Use GetAsPaymentsGetResponse instead.
// returns a AffiliatesItemPaymentsResponseable when successful
// returns a AffiliatesItemPayments401Error error when the service returns a 401 status code
// returns a AffiliatesItemPayments403Error error when the service returns a 403 status code
// returns a AffiliatesItemPayments404Error error when the service returns a 404 status code
func (m *AffiliatesItemPaymentsRequestBuilder) Get(ctx context.Context, requestConfiguration *AffiliatesItemPaymentsRequestBuilderGetRequestConfiguration) (AffiliatesItemPaymentsResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAffiliatesItemPayments401ErrorFromDiscriminatorValue,
		"403": CreateAffiliatesItemPayments403ErrorFromDiscriminatorValue,
		"404": CreateAffiliatesItemPayments404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAffiliatesItemPaymentsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AffiliatesItemPaymentsResponseable), nil
}

// GetAsPaymentsGetResponse retrieves a list of all affiliate payments
// returns a AffiliatesItemPaymentsGetResponseable when successful
// returns a AffiliatesItemPayments401Error error when the service returns a 401 status code
// returns a AffiliatesItemPayments403Error error when the service returns a 403 status code
// returns a AffiliatesItemPayments404Error error when the service returns a 404 status code
func (m *AffiliatesItemPaymentsRequestBuilder) GetAsPaymentsGetResponse(ctx context.Context, requestConfiguration *AffiliatesItemPaymentsRequestBuilderGetRequestConfiguration) (AffiliatesItemPaymentsGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAffiliatesItemPayments401ErrorFromDiscriminatorValue,
		"403": CreateAffiliatesItemPayments403ErrorFromDiscriminatorValue,
		"404": CreateAffiliatesItemPayments404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAffiliatesItemPaymentsGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AffiliatesItemPaymentsGetResponseable), nil
}

// ToGetRequestInformation retrieves a list of all affiliate payments
// returns a *RequestInformation when successful
func (m *AffiliatesItemPaymentsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AffiliatesItemPaymentsRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AffiliatesItemPaymentsRequestBuilder when successful
func (m *AffiliatesItemPaymentsRequestBuilder) WithUrl(rawUrl string) *AffiliatesItemPaymentsRequestBuilder {
	return NewAffiliatesItemPaymentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
