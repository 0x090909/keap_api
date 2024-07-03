package v2

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CompaniesRequestBuilder builds and executes requests for operations under \v2\companies
type CompaniesRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// CompaniesRequestBuilderGetQueryParameters retrieves a list of all Companies
type CompaniesRequestBuilderGetQueryParameters struct {
	// Comma-delimited list of Company properties to include in the response. (Fields such as `notes`, `fax_number` and `custom_fields` aren't included, by default.)
	Fields []string `uriparametername:"fields"`
	// Search filter to apply to results
	Filter *string `uriparametername:"filter"`
	// Attribute and direction to order items by. E.g. `given_name desc`
	Order_by *string `uriparametername:"order_by"`
	// Total number of items to return per page
	Page_size *int32 `uriparametername:"page_size"`
	// Page token
	Page_token *string `uriparametername:"page_token"`
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

// ByCompany_id gets an item from the github.com/0x090909/keap_api.v2.companies.item collection
// returns a *CompaniesWithCompany_ItemRequestBuilder when successful
func (m *CompaniesRequestBuilder) ByCompany_id(company_id string) *CompaniesWithCompany_ItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if company_id != "" {
		urlTplParams["company_id"] = company_id
	}
	return NewCompaniesWithCompany_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewCompaniesRequestBuilderInternal instantiates a new CompaniesRequestBuilder and sets the default values.
func NewCompaniesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CompaniesRequestBuilder {
	m := &CompaniesRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/companies{?fields*,filter*,order_by*,page_size*,page_token*}", pathParameters),
	}
	return m
}

// NewCompaniesRequestBuilder instantiates a new CompaniesRequestBuilder and sets the default values.
func NewCompaniesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CompaniesRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewCompaniesRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieves a list of all Companies
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

// GetAsCompaniesGetResponse retrieves a list of all Companies
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

// Post creates a new Company.`country_code` is required if `region` is specified.
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

// PostAsCompaniesPostResponse creates a new Company.`country_code` is required if `region` is specified.
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

// ToGetRequestInformation retrieves a list of all Companies
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

// ToPostRequestInformation creates a new Company.`country_code` is required if `region` is specified.
// returns a *RequestInformation when successful
func (m *CompaniesRequestBuilder) ToPostRequestInformation(ctx context.Context, body CompaniesPostRequestBodyable, requestConfiguration *CompaniesRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/v2/companies", m.BaseRequestBuilder.PathParameters)
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
