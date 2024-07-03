package v2

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AffiliatesAffiliate_idItemRequestBuilder builds and executes requests for operations under \v2\affiliates\{affiliate_id-id}
type AffiliatesAffiliate_idItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// AffiliatesAffiliate_idItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AffiliatesAffiliate_idItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// AffiliatesAffiliate_idItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AffiliatesAffiliate_idItemRequestBuilderPatchRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// CommissionTotal the commissionTotal property
// returns a *AffiliatesItemCommissionTotalRequestBuilder when successful
func (m *AffiliatesAffiliate_idItemRequestBuilder) CommissionTotal() *AffiliatesItemCommissionTotalRequestBuilder {
	return NewAffiliatesItemCommissionTotalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// NewAffiliatesAffiliate_idItemRequestBuilderInternal instantiates a new AffiliatesAffiliate_idItemRequestBuilder and sets the default values.
func NewAffiliatesAffiliate_idItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesAffiliate_idItemRequestBuilder {
	m := &AffiliatesAffiliate_idItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/affiliates/{affiliate_id%2Did}", pathParameters),
	}
	return m
}

// NewAffiliatesAffiliate_idItemRequestBuilder instantiates a new AffiliatesAffiliate_idItemRequestBuilder and sets the default values.
func NewAffiliatesAffiliate_idItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesAffiliate_idItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewAffiliatesAffiliate_idItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieves a single Affiliate
// Deprecated: This method is obsolete. Use GetAsAffiliate_idGetResponse instead.
// returns a AffiliatesItemAffiliate_idResponseable when successful
// returns a AffiliatesItemAffiliate_id401Error error when the service returns a 401 status code
// returns a AffiliatesItemAffiliate_id403Error error when the service returns a 403 status code
// returns a AffiliatesItemAffiliate_id404Error error when the service returns a 404 status code
func (m *AffiliatesAffiliate_idItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AffiliatesAffiliate_idItemRequestBuilderGetRequestConfiguration) (AffiliatesItemAffiliate_idResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAffiliatesItemAffiliate_id401ErrorFromDiscriminatorValue,
		"403": CreateAffiliatesItemAffiliate_id403ErrorFromDiscriminatorValue,
		"404": CreateAffiliatesItemAffiliate_id404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAffiliatesItemAffiliate_idResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AffiliatesItemAffiliate_idResponseable), nil
}

// GetAsAffiliate_idGetResponse retrieves a single Affiliate
// returns a AffiliatesItemAffiliate_idGetResponseable when successful
// returns a AffiliatesItemAffiliate_id401Error error when the service returns a 401 status code
// returns a AffiliatesItemAffiliate_id403Error error when the service returns a 403 status code
// returns a AffiliatesItemAffiliate_id404Error error when the service returns a 404 status code
func (m *AffiliatesAffiliate_idItemRequestBuilder) GetAsAffiliate_idGetResponse(ctx context.Context, requestConfiguration *AffiliatesAffiliate_idItemRequestBuilderGetRequestConfiguration) (AffiliatesItemAffiliate_idGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAffiliatesItemAffiliate_id401ErrorFromDiscriminatorValue,
		"403": CreateAffiliatesItemAffiliate_id403ErrorFromDiscriminatorValue,
		"404": CreateAffiliatesItemAffiliate_id404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAffiliatesItemAffiliate_idGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AffiliatesItemAffiliate_idGetResponseable), nil
}

// Patch updates a single Affiliate
// Deprecated: This method is obsolete. Use PatchAsAffiliate_idPatchResponse instead.
// returns a AffiliatesItemAffiliate_idResponseable when successful
// returns a AffiliatesItemAffiliate_id401Error error when the service returns a 401 status code
// returns a AffiliatesItemAffiliate_id403Error error when the service returns a 403 status code
// returns a AffiliatesItemAffiliate_id404Error error when the service returns a 404 status code
func (m *AffiliatesAffiliate_idItemRequestBuilder) Patch(ctx context.Context, body AffiliatesItemAffiliate_idPatchRequestBodyable, requestConfiguration *AffiliatesAffiliate_idItemRequestBuilderPatchRequestConfiguration) (AffiliatesItemAffiliate_idResponseable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAffiliatesItemAffiliate_id401ErrorFromDiscriminatorValue,
		"403": CreateAffiliatesItemAffiliate_id403ErrorFromDiscriminatorValue,
		"404": CreateAffiliatesItemAffiliate_id404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAffiliatesItemAffiliate_idResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AffiliatesItemAffiliate_idResponseable), nil
}

// PatchAsAffiliate_idPatchResponse updates a single Affiliate
// returns a AffiliatesItemAffiliate_idPatchResponseable when successful
// returns a AffiliatesItemAffiliate_id401Error error when the service returns a 401 status code
// returns a AffiliatesItemAffiliate_id403Error error when the service returns a 403 status code
// returns a AffiliatesItemAffiliate_id404Error error when the service returns a 404 status code
func (m *AffiliatesAffiliate_idItemRequestBuilder) PatchAsAffiliate_idPatchResponse(ctx context.Context, body AffiliatesItemAffiliate_idPatchRequestBodyable, requestConfiguration *AffiliatesAffiliate_idItemRequestBuilderPatchRequestConfiguration) (AffiliatesItemAffiliate_idPatchResponseable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAffiliatesItemAffiliate_id401ErrorFromDiscriminatorValue,
		"403": CreateAffiliatesItemAffiliate_id403ErrorFromDiscriminatorValue,
		"404": CreateAffiliatesItemAffiliate_id404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAffiliatesItemAffiliate_idPatchResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AffiliatesItemAffiliate_idPatchResponseable), nil
}

// ToGetRequestInformation retrieves a single Affiliate
// returns a *RequestInformation when successful
func (m *AffiliatesAffiliate_idItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AffiliatesAffiliate_idItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// ToPatchRequestInformation updates a single Affiliate
// returns a *RequestInformation when successful
func (m *AffiliatesAffiliate_idItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body AffiliatesItemAffiliate_idPatchRequestBodyable, requestConfiguration *AffiliatesAffiliate_idItemRequestBuilderPatchRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
	if err != nil {
		return nil, err
	}
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *AffiliatesAffiliate_idItemRequestBuilder when successful
func (m *AffiliatesAffiliate_idItemRequestBuilder) WithUrl(rawUrl string) *AffiliatesAffiliate_idItemRequestBuilder {
	return NewAffiliatesAffiliate_idItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
