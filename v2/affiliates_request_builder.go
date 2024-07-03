package v2

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AffiliatesRequestBuilder builds and executes requests for operations under \v2\affiliates
type AffiliatesRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// AffiliatesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AffiliatesRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ByAffiliate_idId gets an item from the github.com/0x090909/keap_api.v2.affiliates.item collection
// returns a *AffiliatesAffiliate_idItemRequestBuilder when successful
func (m *AffiliatesRequestBuilder) ByAffiliate_idId(affiliate_idId string) *AffiliatesAffiliate_idItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if affiliate_idId != "" {
		urlTplParams["affiliate_id%2Did"] = affiliate_idId
	}
	return NewAffiliatesAffiliate_idItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// CommissionPrograms the commissionPrograms property
// returns a *AffiliatesCommissionProgramsRequestBuilder when successful
func (m *AffiliatesRequestBuilder) CommissionPrograms() *AffiliatesCommissionProgramsRequestBuilder {
	return NewAffiliatesCommissionProgramsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// NewAffiliatesRequestBuilderInternal instantiates a new AffiliatesRequestBuilder and sets the default values.
func NewAffiliatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesRequestBuilder {
	m := &AffiliatesRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/affiliates", pathParameters),
	}
	return m
}

// NewAffiliatesRequestBuilder instantiates a new AffiliatesRequestBuilder and sets the default values.
func NewAffiliatesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewAffiliatesRequestBuilderInternal(urlParams, requestAdapter)
}

// Post creates a single affiliate
// Deprecated: This method is obsolete. Use PostAsAffiliatesPostResponse instead.
// returns a AffiliatesResponseable when successful
// returns a Affiliates401Error error when the service returns a 401 status code
// returns a Affiliates403Error error when the service returns a 403 status code
func (m *AffiliatesRequestBuilder) Post(ctx context.Context, body AffiliatesPostRequestBodyable, requestConfiguration *AffiliatesRequestBuilderPostRequestConfiguration) (AffiliatesResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAffiliates401ErrorFromDiscriminatorValue,
		"403": CreateAffiliates403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAffiliatesResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AffiliatesResponseable), nil
}

// PostAsAffiliatesPostResponse creates a single affiliate
// returns a AffiliatesPostResponseable when successful
// returns a Affiliates401Error error when the service returns a 401 status code
// returns a Affiliates403Error error when the service returns a 403 status code
func (m *AffiliatesRequestBuilder) PostAsAffiliatesPostResponse(ctx context.Context, body AffiliatesPostRequestBodyable, requestConfiguration *AffiliatesRequestBuilderPostRequestConfiguration) (AffiliatesPostResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAffiliates401ErrorFromDiscriminatorValue,
		"403": CreateAffiliates403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAffiliatesPostResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AffiliatesPostResponseable), nil
}

// Redirects the redirects property
// returns a *AffiliatesRedirectsRequestBuilder when successful
func (m *AffiliatesRequestBuilder) Redirects() *AffiliatesRedirectsRequestBuilder {
	return NewAffiliatesRedirectsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// ToPostRequestInformation creates a single affiliate
// returns a *RequestInformation when successful
func (m *AffiliatesRequestBuilder) ToPostRequestInformation(ctx context.Context, body AffiliatesPostRequestBodyable, requestConfiguration *AffiliatesRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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

// WithAffiliate_idCommissions builds and executes requests for operations under \v2\affiliates\{affiliate_id}:commissions
// returns a *AffiliatesWithAffiliate_idCommissionsRequestBuilder when successful
func (m *AffiliatesRequestBuilder) WithAffiliate_idCommissions() *AffiliatesWithAffiliate_idCommissionsRequestBuilder {
	return NewAffiliatesWithAffiliate_idCommissionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// WithIdAssignToProgram builds and executes requests for operations under \v2\affiliates\{id}:assignToProgram
// returns a *AffiliatesWithIdAssignToProgramRequestBuilder when successful
func (m *AffiliatesRequestBuilder) WithIdAssignToProgram() *AffiliatesWithIdAssignToProgramRequestBuilder {
	return NewAffiliatesWithIdAssignToProgramRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// WithIdRemoveFromProgram builds and executes requests for operations under \v2\affiliates\{id}:removeFromProgram
// returns a *AffiliatesWithIdRemoveFromProgramRequestBuilder when successful
func (m *AffiliatesRequestBuilder) WithIdRemoveFromProgram() *AffiliatesWithIdRemoveFromProgramRequestBuilder {
	return NewAffiliatesWithIdRemoveFromProgramRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *AffiliatesRequestBuilder when successful
func (m *AffiliatesRequestBuilder) WithUrl(rawUrl string) *AffiliatesRequestBuilder {
	return NewAffiliatesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
