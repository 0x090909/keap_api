package v1

import (
	"context"
	i7000a04bcc287966e11f9bf4b54e0affbb4df21312b21d781999ce245c6abcfe "github.com/0x090909/keap_api/v1/companies"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
)

// CompaniesRequestBuilder builds and executes requests for operations under \v1\companies
type CompaniesRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// CompaniesRequestBuilderGetQueryParameters retrieves a list of all companies
type CompaniesRequestBuilderGetQueryParameters struct {
	// Optional company name to query on
	Company_name *string `uriparametername:"company_name"`
	// Sets a total of items to return
	Limit *int32 `uriparametername:"limit"`
	// Sets a beginning range of items to return
	Offset *int32 `uriparametername:"offset"`
	// Comma-delimited list of Company properties to include in the response. (Fields such as `notes`, `fax_number` and `custom_fields` aren't included, by default.)
	Optional_properties []string `uriparametername:"optional_properties"`
	// Attribute to order items by
	// Deprecated: This property is deprecated, use OrderAsGetOrderQueryParameterType instead
	Order *string `uriparametername:"order"`
	// How to order the data i.e. ascending (A-Z) or descending (Z-A)
	// Deprecated: This property is deprecated, use Order_directionAsGetOrder_directionQueryParameterType instead
	Order_direction *string `uriparametername:"order_direction"`
	// How to order the data i.e. ascending (A-Z) or descending (Z-A)
	Order_directionAsGetOrder_directionQueryParameterType *i7000a04bcc287966e11f9bf4b54e0affbb4df21312b21d781999ce245c6abcfe.GetOrder_directionQueryParameterType `uriparametername:"order_direction"`
	// Attribute to order items by
	OrderAsGetOrderQueryParameterType *i7000a04bcc287966e11f9bf4b54e0affbb4df21312b21d781999ce245c6abcfe.GetOrderQueryParameterType `uriparametername:"order"`
}

// CompaniesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *CompaniesRequestBuilderGetQueryParameters
}

// CompaniesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ByCompanyId gets an item from the github.com/0x090909/keap_api.v1.companies.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *CompaniesWithCompanyItemRequestBuilder when successful
func (m *CompaniesRequestBuilder) ByCompanyId(companyId string) *CompaniesWithCompanyItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if companyId != "" {
		urlTplParams["companyId"] = companyId
	}
	return NewCompaniesWithCompanyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// ByCompanyIdInt64 gets an item from the github.com/0x090909/keap_api.v1.companies.item collection
// returns a *CompaniesWithCompanyItemRequestBuilder when successful
func (m *CompaniesRequestBuilder) ByCompanyIdInt64(companyId int64) *CompaniesWithCompanyItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	urlTplParams["companyId"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(companyId, 10)
	return NewCompaniesWithCompanyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewCompaniesRequestBuilderInternal instantiates a new CompaniesRequestBuilder and sets the default values.
func NewCompaniesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CompaniesRequestBuilder {
	m := &CompaniesRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/companies{?company_name*,limit*,offset*,optional_properties*,order*,order_direction*}", pathParameters),
	}
	return m
}

// NewCompaniesRequestBuilder instantiates a new CompaniesRequestBuilder and sets the default values.
func NewCompaniesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CompaniesRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewCompaniesRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieves a list of all companies
// Deprecated: This method is obsolete. Use GetAsCompaniesGetResponse instead.
// returns a CompaniesResponseable when successful
// returns a Companies401Error error when the service returns a 401 status code
// returns a Companies403Error error when the service returns a 403 status code
// returns a Companies404Error error when the service returns a 404 status code
func (m *CompaniesRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesRequestBuilderGetRequestConfiguration) (CompaniesResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateCompanies401ErrorFromDiscriminatorValue,
		"403": CreateCompanies403ErrorFromDiscriminatorValue,
		"404": CreateCompanies404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCompaniesResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(CompaniesResponseable), nil
}

// GetAsCompaniesGetResponse retrieves a list of all companies
// returns a CompaniesGetResponseable when successful
// returns a Companies401Error error when the service returns a 401 status code
// returns a Companies403Error error when the service returns a 403 status code
// returns a Companies404Error error when the service returns a 404 status code
func (m *CompaniesRequestBuilder) GetAsCompaniesGetResponse(ctx context.Context, requestConfiguration *CompaniesRequestBuilderGetRequestConfiguration) (CompaniesGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateCompanies401ErrorFromDiscriminatorValue,
		"403": CreateCompanies403ErrorFromDiscriminatorValue,
		"404": CreateCompanies404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCompaniesGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(CompaniesGetResponseable), nil
}

// Model the model property
// returns a *CompaniesModelRequestBuilder when successful
func (m *CompaniesRequestBuilder) Model() *CompaniesModelRequestBuilder {
	return NewCompaniesModelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Post creates a new company as the authenticated user. NB: Company must contain at least the company name. The `country_code` is required if `region` is specified.You may opt-in or mark a Company as _Marketable_ by including the following field in the request JSON with an opt-in reason. (This field is also shown in the complete request body sample.) The reason you provide here will help with compliance. Example reasons: "Customer opted-in through webform", "Company gave explicit permission."```json"opt_in_reason": "your reason for opt-in"```Note that the email address status will only be updated to unconfirmed (marketable) for email addresses that are currently in the following states:Unengaged MarketableUnengaged Non-MarketableNon-MarketableOpt-Out: ManualAll other existing statuses e.g. List Unsubscribe, Opt-Out System etc will remain non-marketable and in their existing state.
// Deprecated: This method is obsolete. Use PostAsCompaniesPostResponse instead.
// returns a CompaniesResponseable when successful
// returns a Companies401Error error when the service returns a 401 status code
// returns a Companies403Error error when the service returns a 403 status code
func (m *CompaniesRequestBuilder) Post(ctx context.Context, body CompaniesPostRequestBodyable, requestConfiguration *CompaniesRequestBuilderPostRequestConfiguration) (CompaniesResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateCompanies401ErrorFromDiscriminatorValue,
		"403": CreateCompanies403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCompaniesResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(CompaniesResponseable), nil
}

// PostAsCompaniesPostResponse creates a new company as the authenticated user. NB: Company must contain at least the company name. The `country_code` is required if `region` is specified.You may opt-in or mark a Company as _Marketable_ by including the following field in the request JSON with an opt-in reason. (This field is also shown in the complete request body sample.) The reason you provide here will help with compliance. Example reasons: "Customer opted-in through webform", "Company gave explicit permission."```json"opt_in_reason": "your reason for opt-in"```Note that the email address status will only be updated to unconfirmed (marketable) for email addresses that are currently in the following states:Unengaged MarketableUnengaged Non-MarketableNon-MarketableOpt-Out: ManualAll other existing statuses e.g. List Unsubscribe, Opt-Out System etc will remain non-marketable and in their existing state.
// returns a CompaniesPostResponseable when successful
// returns a Companies401Error error when the service returns a 401 status code
// returns a Companies403Error error when the service returns a 403 status code
func (m *CompaniesRequestBuilder) PostAsCompaniesPostResponse(ctx context.Context, body CompaniesPostRequestBodyable, requestConfiguration *CompaniesRequestBuilderPostRequestConfiguration) (CompaniesPostResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateCompanies401ErrorFromDiscriminatorValue,
		"403": CreateCompanies403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCompaniesPostResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(CompaniesPostResponseable), nil
}

// ToGetRequestInformation retrieves a list of all companies
// returns a *RequestInformation when successful
func (m *CompaniesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToPostRequestInformation creates a new company as the authenticated user. NB: Company must contain at least the company name. The `country_code` is required if `region` is specified.You may opt-in or mark a Company as _Marketable_ by including the following field in the request JSON with an opt-in reason. (This field is also shown in the complete request body sample.) The reason you provide here will help with compliance. Example reasons: "Customer opted-in through webform", "Company gave explicit permission."```json"opt_in_reason": "your reason for opt-in"```Note that the email address status will only be updated to unconfirmed (marketable) for email addresses that are currently in the following states:Unengaged MarketableUnengaged Non-MarketableNon-MarketableOpt-Out: ManualAll other existing statuses e.g. List Unsubscribe, Opt-Out System etc will remain non-marketable and in their existing state.
// returns a *RequestInformation when successful
func (m *CompaniesRequestBuilder) ToPostRequestInformation(ctx context.Context, body CompaniesPostRequestBodyable, requestConfiguration *CompaniesRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/v1/companies", m.BaseRequestBuilder.PathParameters)
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
// returns a *CompaniesRequestBuilder when successful
func (m *CompaniesRequestBuilder) WithUrl(rawUrl string) *CompaniesRequestBuilder {
	return NewCompaniesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
