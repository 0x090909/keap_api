package v1

import (
	"context"
	i211b9bfbed966901f2709505f9940dc5d7da8e29cb5287c98a3cd829ff64593e "github.com/0x090909/keap_api/v1/affiliates"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
)

// AffiliatesRequestBuilder builds and executes requests for operations under \v1\affiliates
type AffiliatesRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// AffiliatesRequestBuilderGetQueryParameters retrieves a list of all affiliates
type AffiliatesRequestBuilderGetQueryParameters struct {
	// Optional affiliate parent ID to query on
	Code *string `uriparametername:"code"`
	// Optional contact ID to query on
	Contact_id *string `uriparametername:"contact_id"`
	// Sets a beginning range of items to return
	Limit *int32 `uriparametername:"limit"`
	// Optional affiliate name to query on
	Name *string `uriparametername:"name"`
	// Count to offset the returned results by
	Offset *int32 `uriparametername:"offset"`
	// Attribute to order items by
	// Deprecated: This property is deprecated, use OrderAsGetOrderQueryParameterType instead
	Order *string `uriparametername:"order"`
	// How to order the data i.e. ascending (A-Z) or descending (Z-A)
	// Deprecated: This property is deprecated, use Order_directionAsGetOrder_directionQueryParameterType instead
	Order_direction *string `uriparametername:"order_direction"`
	// How to order the data i.e. ascending (A-Z) or descending (Z-A)
	Order_directionAsGetOrder_directionQueryParameterType *i211b9bfbed966901f2709505f9940dc5d7da8e29cb5287c98a3cd829ff64593e.GetOrder_directionQueryParameterType `uriparametername:"order_direction"`
	// Attribute to order items by
	OrderAsGetOrderQueryParameterType *i211b9bfbed966901f2709505f9940dc5d7da8e29cb5287c98a3cd829ff64593e.GetOrderQueryParameterType `uriparametername:"order"`
	// Optional affiliate code to query on
	Parent_id *string `uriparametername:"parent_id"`
	// Optional affiliate program ID to query on
	Program_id *string `uriparametername:"program_id"`
	// Optional affiliate status to query on
	// Deprecated: This property is deprecated, use StatusAsGetStatusQueryParameterType instead
	Status *string `uriparametername:"status"`
	// Optional affiliate status to query on
	StatusAsGetStatusQueryParameterType *i211b9bfbed966901f2709505f9940dc5d7da8e29cb5287c98a3cd829ff64593e.GetStatusQueryParameterType `uriparametername:"status"`
}

// AffiliatesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AffiliatesRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *AffiliatesRequestBuilderGetQueryParameters
}

// AffiliatesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AffiliatesRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ByAffiliateIdId gets an item from the github.com/0x090909/keap_api.v1.affiliates.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *AffiliatesAffiliateIdItemRequestBuilder when successful
func (m *AffiliatesRequestBuilder) ByAffiliateIdId(affiliateIdId string) *AffiliatesAffiliateIdItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if affiliateIdId != "" {
		urlTplParams["affiliateId%2Did"] = affiliateIdId
	}
	return NewAffiliatesAffiliateIdItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// ByAffiliateIdIdInt64 gets an item from the github.com/0x090909/keap_api.v1.affiliates.item collection
// returns a *AffiliatesAffiliateIdItemRequestBuilder when successful
func (m *AffiliatesRequestBuilder) ByAffiliateIdIdInt64(affiliateIdId int64) *AffiliatesAffiliateIdItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	urlTplParams["affiliateId%2Did"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(affiliateIdId, 10)
	return NewAffiliatesAffiliateIdItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// Commissions the commissions property
// returns a *AffiliatesCommissionsRequestBuilder when successful
func (m *AffiliatesRequestBuilder) Commissions() *AffiliatesCommissionsRequestBuilder {
	return NewAffiliatesCommissionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// NewAffiliatesRequestBuilderInternal instantiates a new AffiliatesRequestBuilder and sets the default values.
func NewAffiliatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesRequestBuilder {
	m := &AffiliatesRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/affiliates{?code*,contact_id*,limit*,name*,offset*,order*,order_direction*,parent_id*,program_id*,status*}", pathParameters),
	}
	return m
}

// NewAffiliatesRequestBuilder instantiates a new AffiliatesRequestBuilder and sets the default values.
func NewAffiliatesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewAffiliatesRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieves a list of all affiliates
// Deprecated: This method is obsolete. Use GetAsAffiliatesGetResponse instead.
// returns a AffiliatesResponseable when successful
// returns a Affiliates401Error error when the service returns a 401 status code
// returns a Affiliates403Error error when the service returns a 403 status code
// returns a Affiliates404Error error when the service returns a 404 status code
func (m *AffiliatesRequestBuilder) Get(ctx context.Context, requestConfiguration *AffiliatesRequestBuilderGetRequestConfiguration) (AffiliatesResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAffiliates401ErrorFromDiscriminatorValue,
		"403": CreateAffiliates403ErrorFromDiscriminatorValue,
		"404": CreateAffiliates404ErrorFromDiscriminatorValue,
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

// GetAsAffiliatesGetResponse retrieves a list of all affiliates
// returns a AffiliatesGetResponseable when successful
// returns a Affiliates401Error error when the service returns a 401 status code
// returns a Affiliates403Error error when the service returns a 403 status code
// returns a Affiliates404Error error when the service returns a 404 status code
func (m *AffiliatesRequestBuilder) GetAsAffiliatesGetResponse(ctx context.Context, requestConfiguration *AffiliatesRequestBuilderGetRequestConfiguration) (AffiliatesGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAffiliates401ErrorFromDiscriminatorValue,
		"403": CreateAffiliates403ErrorFromDiscriminatorValue,
		"404": CreateAffiliates404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAffiliatesGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AffiliatesGetResponseable), nil
}

// Model the model property
// returns a *AffiliatesModelRequestBuilder when successful
func (m *AffiliatesRequestBuilder) Model() *AffiliatesModelRequestBuilder {
	return NewAffiliatesModelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Post create a single affiliate
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

// PostAsAffiliatesPostResponse create a single affiliate
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

// Programs the programs property
// returns a *AffiliatesProgramsRequestBuilder when successful
func (m *AffiliatesRequestBuilder) Programs() *AffiliatesProgramsRequestBuilder {
	return NewAffiliatesProgramsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Redirectlinks the redirectlinks property
// returns a *AffiliatesRedirectlinksRequestBuilder when successful
func (m *AffiliatesRequestBuilder) Redirectlinks() *AffiliatesRedirectlinksRequestBuilder {
	return NewAffiliatesRedirectlinksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Summaries the summaries property
// returns a *AffiliatesSummariesRequestBuilder when successful
func (m *AffiliatesRequestBuilder) Summaries() *AffiliatesSummariesRequestBuilder {
	return NewAffiliatesSummariesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// ToGetRequestInformation retrieves a list of all affiliates
// returns a *RequestInformation when successful
func (m *AffiliatesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AffiliatesRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToPostRequestInformation create a single affiliate
// returns a *RequestInformation when successful
func (m *AffiliatesRequestBuilder) ToPostRequestInformation(ctx context.Context, body AffiliatesPostRequestBodyable, requestConfiguration *AffiliatesRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/v1/affiliates", m.BaseRequestBuilder.PathParameters)
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
// returns a *AffiliatesRequestBuilder when successful
func (m *AffiliatesRequestBuilder) WithUrl(rawUrl string) *AffiliatesRequestBuilder {
	return NewAffiliatesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
