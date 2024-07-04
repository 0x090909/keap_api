package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AffiliatesAffiliateIdItemRequestBuilder builds and executes requests for operations under \v1\affiliates\{affiliateId-id}
type AffiliatesAffiliateIdItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// AffiliatesAffiliateIdItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AffiliatesAffiliateIdItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// Clawbacks the clawbacks property
// returns a *AffiliatesItemClawbacksRequestBuilder when successful
func (m *AffiliatesAffiliateIdItemRequestBuilder) Clawbacks() *AffiliatesItemClawbacksRequestBuilder {
	return NewAffiliatesItemClawbacksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// NewAffiliatesAffiliateIdItemRequestBuilderInternal instantiates a new AffiliatesAffiliateIdItemRequestBuilder and sets the default values.
func NewAffiliatesAffiliateIdItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesAffiliateIdItemRequestBuilder {
	m := &AffiliatesAffiliateIdItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/affiliates/{affiliateId%2Did}", pathParameters),
	}
	return m
}

// NewAffiliatesAffiliateIdItemRequestBuilder instantiates a new AffiliatesAffiliateIdItemRequestBuilder and sets the default values.
func NewAffiliatesAffiliateIdItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesAffiliateIdItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewAffiliatesAffiliateIdItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieve a single affiliate
// Deprecated: This method is obsolete. Use GetAsAffiliateIdGetResponse instead.
// returns a AffiliatesItemAffiliateIdResponseable when successful
// returns a AffiliatesItemAffiliateId401Error error when the service returns a 401 status code
// returns a AffiliatesItemAffiliateId403Error error when the service returns a 403 status code
// returns a AffiliatesItemAffiliateId404Error error when the service returns a 404 status code
func (m *AffiliatesAffiliateIdItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AffiliatesAffiliateIdItemRequestBuilderGetRequestConfiguration) (AffiliatesItemAffiliateIdResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAffiliatesItemAffiliateId401ErrorFromDiscriminatorValue,
		"403": CreateAffiliatesItemAffiliateId403ErrorFromDiscriminatorValue,
		"404": CreateAffiliatesItemAffiliateId404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAffiliatesItemAffiliateIdResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AffiliatesItemAffiliateIdResponseable), nil
}

// GetAsAffiliateIdGetResponse retrieve a single affiliate
// returns a AffiliatesItemAffiliateIdGetResponseable when successful
// returns a AffiliatesItemAffiliateId401Error error when the service returns a 401 status code
// returns a AffiliatesItemAffiliateId403Error error when the service returns a 403 status code
// returns a AffiliatesItemAffiliateId404Error error when the service returns a 404 status code
func (m *AffiliatesAffiliateIdItemRequestBuilder) GetAsAffiliateIdGetResponse(ctx context.Context, requestConfiguration *AffiliatesAffiliateIdItemRequestBuilderGetRequestConfiguration) (AffiliatesItemAffiliateIdGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAffiliatesItemAffiliateId401ErrorFromDiscriminatorValue,
		"403": CreateAffiliatesItemAffiliateId403ErrorFromDiscriminatorValue,
		"404": CreateAffiliatesItemAffiliateId404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAffiliatesItemAffiliateIdGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AffiliatesItemAffiliateIdGetResponseable), nil
}

// Payments the payments property
// returns a *AffiliatesItemPaymentsRequestBuilder when successful
func (m *AffiliatesAffiliateIdItemRequestBuilder) Payments() *AffiliatesItemPaymentsRequestBuilder {
	return NewAffiliatesItemPaymentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// ToGetRequestInformation retrieve a single affiliate
// returns a *RequestInformation when successful
func (m *AffiliatesAffiliateIdItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AffiliatesAffiliateIdItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *AffiliatesAffiliateIdItemRequestBuilder when successful
func (m *AffiliatesAffiliateIdItemRequestBuilder) WithUrl(rawUrl string) *AffiliatesAffiliateIdItemRequestBuilder {
	return NewAffiliatesAffiliateIdItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
