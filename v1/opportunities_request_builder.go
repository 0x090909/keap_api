package v1

import (
	"context"
	i26bd8cb7fd079cc0cdb3d7a4d285abb61d385f288ccad7fd6743f3569c4527fa "github.com/0x090909/keap_api/v1/opportunities"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
)

// OpportunitiesRequestBuilder builds and executes requests for operations under \v1\opportunities
type OpportunitiesRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// OpportunitiesRequestBuilderGetQueryParameters retrieves a list of all opportunities.Please note: the sample response erroneously shows properties, such as _stage reasons_, that are unavailable through the list endpoint. Such properties are only available through the retrieve operation. Future versions of the Opportunity resource will correct the oversight.
type OpportunitiesRequestBuilderGetQueryParameters struct {
	// Sets a total of items to return
	Limit *int32 `uriparametername:"limit"`
	// Sets a beginning range of items to return
	Offset *int32 `uriparametername:"offset"`
	// Attribute to order items by
	// Deprecated: This property is deprecated, use OrderAsGetOrderQueryParameterType instead
	Order *string `uriparametername:"order"`
	// Attribute to order items by
	OrderAsGetOrderQueryParameterType *i26bd8cb7fd079cc0cdb3d7a4d285abb61d385f288ccad7fd6743f3569c4527fa.GetOrderQueryParameterType `uriparametername:"order"`
	// Returns opportunities that match any of the contact's `given_name`, `family_name`, `company_name`, and `email_addresses` (searches `EMAIL1` only) fields as well as `opportunity_title`
	Search_term *string `uriparametername:"search_term"`
	// Returns opportunities for the provided stage id
	Stage_id *int64 `uriparametername:"stage_id"`
	// Returns opportunities for the provided user id
	User_id *int64 `uriparametername:"user_id"`
}

// OpportunitiesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OpportunitiesRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *OpportunitiesRequestBuilderGetQueryParameters
}

// OpportunitiesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OpportunitiesRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// OpportunitiesRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OpportunitiesRequestBuilderPutRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ByOpportunityId gets an item from the github.com/0x090909/keap_api.v1.opportunities.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *OpportunitiesWithOpportunityItemRequestBuilder when successful
func (m *OpportunitiesRequestBuilder) ByOpportunityId(opportunityId string) *OpportunitiesWithOpportunityItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if opportunityId != "" {
		urlTplParams["opportunityId"] = opportunityId
	}
	return NewOpportunitiesWithOpportunityItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// ByOpportunityIdInt64 gets an item from the github.com/0x090909/keap_api.v1.opportunities.item collection
// returns a *OpportunitiesWithOpportunityItemRequestBuilder when successful
func (m *OpportunitiesRequestBuilder) ByOpportunityIdInt64(opportunityId int64) *OpportunitiesWithOpportunityItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	urlTplParams["opportunityId"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(opportunityId, 10)
	return NewOpportunitiesWithOpportunityItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewOpportunitiesRequestBuilderInternal instantiates a new OpportunitiesRequestBuilder and sets the default values.
func NewOpportunitiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *OpportunitiesRequestBuilder {
	m := &OpportunitiesRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/opportunities{?limit*,offset*,order*,search_term*,stage_id*,user_id*}", pathParameters),
	}
	return m
}

// NewOpportunitiesRequestBuilder instantiates a new OpportunitiesRequestBuilder and sets the default values.
func NewOpportunitiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *OpportunitiesRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewOpportunitiesRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieves a list of all opportunities.Please note: the sample response erroneously shows properties, such as _stage reasons_, that are unavailable through the list endpoint. Such properties are only available through the retrieve operation. Future versions of the Opportunity resource will correct the oversight.
// Deprecated: This method is obsolete. Use GetAsOpportunitiesGetResponse instead.
// returns a OpportunitiesResponseable when successful
// returns a Opportunities401Error error when the service returns a 401 status code
// returns a Opportunities403Error error when the service returns a 403 status code
// returns a Opportunities404Error error when the service returns a 404 status code
func (m *OpportunitiesRequestBuilder) Get(ctx context.Context, requestConfiguration *OpportunitiesRequestBuilderGetRequestConfiguration) (OpportunitiesResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateOpportunities401ErrorFromDiscriminatorValue,
		"403": CreateOpportunities403ErrorFromDiscriminatorValue,
		"404": CreateOpportunities404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOpportunitiesResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(OpportunitiesResponseable), nil
}

// GetAsOpportunitiesGetResponse retrieves a list of all opportunities.Please note: the sample response erroneously shows properties, such as _stage reasons_, that are unavailable through the list endpoint. Such properties are only available through the retrieve operation. Future versions of the Opportunity resource will correct the oversight.
// returns a OpportunitiesGetResponseable when successful
// returns a Opportunities401Error error when the service returns a 401 status code
// returns a Opportunities403Error error when the service returns a 403 status code
// returns a Opportunities404Error error when the service returns a 404 status code
func (m *OpportunitiesRequestBuilder) GetAsOpportunitiesGetResponse(ctx context.Context, requestConfiguration *OpportunitiesRequestBuilderGetRequestConfiguration) (OpportunitiesGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateOpportunities401ErrorFromDiscriminatorValue,
		"403": CreateOpportunities403ErrorFromDiscriminatorValue,
		"404": CreateOpportunities404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOpportunitiesGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(OpportunitiesGetResponseable), nil
}

// Model the model property
// returns a *OpportunitiesModelRequestBuilder when successful
func (m *OpportunitiesRequestBuilder) Model() *OpportunitiesModelRequestBuilder {
	return NewOpportunitiesModelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Post creates a new opportunity as the authenticated user.
// Deprecated: This method is obsolete. Use PostAsOpportunitiesPostResponse instead.
// returns a OpportunitiesResponseable when successful
// returns a Opportunities401Error error when the service returns a 401 status code
// returns a Opportunities403Error error when the service returns a 403 status code
func (m *OpportunitiesRequestBuilder) Post(ctx context.Context, body OpportunitiesPostRequestBodyable, requestConfiguration *OpportunitiesRequestBuilderPostRequestConfiguration) (OpportunitiesResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateOpportunities401ErrorFromDiscriminatorValue,
		"403": CreateOpportunities403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOpportunitiesResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(OpportunitiesResponseable), nil
}

// PostAsOpportunitiesPostResponse creates a new opportunity as the authenticated user.
// returns a OpportunitiesPostResponseable when successful
// returns a Opportunities401Error error when the service returns a 401 status code
// returns a Opportunities403Error error when the service returns a 403 status code
func (m *OpportunitiesRequestBuilder) PostAsOpportunitiesPostResponse(ctx context.Context, body OpportunitiesPostRequestBodyable, requestConfiguration *OpportunitiesRequestBuilderPostRequestConfiguration) (OpportunitiesPostResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateOpportunities401ErrorFromDiscriminatorValue,
		"403": CreateOpportunities403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOpportunitiesPostResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(OpportunitiesPostResponseable), nil
}

// Put replaces all values of a given opportunity
// Deprecated: This method is obsolete. Use PutAsOpportunitiesPutResponse instead.
// returns a OpportunitiesResponseable when successful
// returns a Opportunities401Error error when the service returns a 401 status code
// returns a Opportunities403Error error when the service returns a 403 status code
// returns a Opportunities404Error error when the service returns a 404 status code
func (m *OpportunitiesRequestBuilder) Put(ctx context.Context, body OpportunitiesPutRequestBodyable, requestConfiguration *OpportunitiesRequestBuilderPutRequestConfiguration) (OpportunitiesResponseable, error) {
	requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateOpportunities401ErrorFromDiscriminatorValue,
		"403": CreateOpportunities403ErrorFromDiscriminatorValue,
		"404": CreateOpportunities404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOpportunitiesResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(OpportunitiesResponseable), nil
}

// PutAsOpportunitiesPutResponse replaces all values of a given opportunity
// returns a OpportunitiesPutResponseable when successful
// returns a Opportunities401Error error when the service returns a 401 status code
// returns a Opportunities403Error error when the service returns a 403 status code
// returns a Opportunities404Error error when the service returns a 404 status code
func (m *OpportunitiesRequestBuilder) PutAsOpportunitiesPutResponse(ctx context.Context, body OpportunitiesPutRequestBodyable, requestConfiguration *OpportunitiesRequestBuilderPutRequestConfiguration) (OpportunitiesPutResponseable, error) {
	requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateOpportunities401ErrorFromDiscriminatorValue,
		"403": CreateOpportunities403ErrorFromDiscriminatorValue,
		"404": CreateOpportunities404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOpportunitiesPutResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(OpportunitiesPutResponseable), nil
}

// ToGetRequestInformation retrieves a list of all opportunities.Please note: the sample response erroneously shows properties, such as _stage reasons_, that are unavailable through the list endpoint. Such properties are only available through the retrieve operation. Future versions of the Opportunity resource will correct the oversight.
// returns a *RequestInformation when successful
func (m *OpportunitiesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OpportunitiesRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToPostRequestInformation creates a new opportunity as the authenticated user.
// returns a *RequestInformation when successful
func (m *OpportunitiesRequestBuilder) ToPostRequestInformation(ctx context.Context, body OpportunitiesPostRequestBodyable, requestConfiguration *OpportunitiesRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/v1/opportunities", m.BaseRequestBuilder.PathParameters)
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

// ToPutRequestInformation replaces all values of a given opportunity
// returns a *RequestInformation when successful
func (m *OpportunitiesRequestBuilder) ToPutRequestInformation(ctx context.Context, body OpportunitiesPutRequestBodyable, requestConfiguration *OpportunitiesRequestBuilderPutRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, "{+baseurl}/v1/opportunities", m.BaseRequestBuilder.PathParameters)
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
// returns a *OpportunitiesRequestBuilder when successful
func (m *OpportunitiesRequestBuilder) WithUrl(rawUrl string) *OpportunitiesRequestBuilder {
	return NewOpportunitiesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
