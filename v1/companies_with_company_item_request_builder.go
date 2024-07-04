package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CompaniesWithCompanyItemRequestBuilder builds and executes requests for operations under \v1\companies\{companyId}
type CompaniesWithCompanyItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// CompaniesWithCompanyItemRequestBuilderGetQueryParameters retrieves a single company
type CompaniesWithCompanyItemRequestBuilderGetQueryParameters struct {
	// Comma-delimited list of Company properties to include in the response. (Fields such as `notes`, `fax_number` and `custom_fields` aren't included, by default.)
	Optional_properties []string `uriparametername:"optional_properties"`
}

// CompaniesWithCompanyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesWithCompanyItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *CompaniesWithCompanyItemRequestBuilderGetQueryParameters
}

// CompaniesWithCompanyItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesWithCompanyItemRequestBuilderPatchRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewCompaniesWithCompanyItemRequestBuilderInternal instantiates a new CompaniesWithCompanyItemRequestBuilder and sets the default values.
func NewCompaniesWithCompanyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CompaniesWithCompanyItemRequestBuilder {
	m := &CompaniesWithCompanyItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/companies/{companyId}{?optional_properties*}", pathParameters),
	}
	return m
}

// NewCompaniesWithCompanyItemRequestBuilder instantiates a new CompaniesWithCompanyItemRequestBuilder and sets the default values.
func NewCompaniesWithCompanyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CompaniesWithCompanyItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewCompaniesWithCompanyItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieves a single company
// Deprecated: This method is obsolete. Use GetAsWithCompanyGetResponse instead.
// returns a CompaniesItemWithCompanyResponseable when successful
// returns a CompaniesItemWithCompany401Error error when the service returns a 401 status code
// returns a CompaniesItemWithCompany403Error error when the service returns a 403 status code
// returns a CompaniesItemWithCompany404Error error when the service returns a 404 status code
func (m *CompaniesWithCompanyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesWithCompanyItemRequestBuilderGetRequestConfiguration) (CompaniesItemWithCompanyResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateCompaniesItemWithCompany401ErrorFromDiscriminatorValue,
		"403": CreateCompaniesItemWithCompany403ErrorFromDiscriminatorValue,
		"404": CreateCompaniesItemWithCompany404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCompaniesItemWithCompanyResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(CompaniesItemWithCompanyResponseable), nil
}

// GetAsWithCompanyGetResponse retrieves a single company
// returns a CompaniesItemWithCompanyGetResponseable when successful
// returns a CompaniesItemWithCompany401Error error when the service returns a 401 status code
// returns a CompaniesItemWithCompany403Error error when the service returns a 403 status code
// returns a CompaniesItemWithCompany404Error error when the service returns a 404 status code
func (m *CompaniesWithCompanyItemRequestBuilder) GetAsWithCompanyGetResponse(ctx context.Context, requestConfiguration *CompaniesWithCompanyItemRequestBuilderGetRequestConfiguration) (CompaniesItemWithCompanyGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateCompaniesItemWithCompany401ErrorFromDiscriminatorValue,
		"403": CreateCompaniesItemWithCompany403ErrorFromDiscriminatorValue,
		"404": CreateCompaniesItemWithCompany404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCompaniesItemWithCompanyGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(CompaniesItemWithCompanyGetResponseable), nil
}

// Patch updates a company with only the values provided in the request.You may opt-in or mark a Company as _Marketable_ by including the following field in the request JSON with an opt-in reason. (This field is also shown in the complete request body sample.) The reason you provide here will help with compliance. Example reasons: "Customer opted-in through webform", "Company gave explicit permission."```json"opt_in_reason": "your reason for opt-in"```Note that the email address status will only be updated to unconfirmed (marketable) for email addresses that are currently in the following states:Unengaged MarketableUnengaged Non-MarketableNon-MarketableOpt-Out: ManualAll other existing statuses e.g. List Unsubscribe, Opt-Out System etc will remain non-marketable and in their existing state.
// Deprecated: This method is obsolete. Use PatchAsWithCompanyPatchResponse instead.
// returns a CompaniesItemWithCompanyResponseable when successful
// returns a CompaniesItemWithCompany401Error error when the service returns a 401 status code
// returns a CompaniesItemWithCompany403Error error when the service returns a 403 status code
// returns a CompaniesItemWithCompany404Error error when the service returns a 404 status code
func (m *CompaniesWithCompanyItemRequestBuilder) Patch(ctx context.Context, body CompaniesItemWithCompanyPatchRequestBodyable, requestConfiguration *CompaniesWithCompanyItemRequestBuilderPatchRequestConfiguration) (CompaniesItemWithCompanyResponseable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateCompaniesItemWithCompany401ErrorFromDiscriminatorValue,
		"403": CreateCompaniesItemWithCompany403ErrorFromDiscriminatorValue,
		"404": CreateCompaniesItemWithCompany404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCompaniesItemWithCompanyResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(CompaniesItemWithCompanyResponseable), nil
}

// PatchAsWithCompanyPatchResponse updates a company with only the values provided in the request.You may opt-in or mark a Company as _Marketable_ by including the following field in the request JSON with an opt-in reason. (This field is also shown in the complete request body sample.) The reason you provide here will help with compliance. Example reasons: "Customer opted-in through webform", "Company gave explicit permission."```json"opt_in_reason": "your reason for opt-in"```Note that the email address status will only be updated to unconfirmed (marketable) for email addresses that are currently in the following states:Unengaged MarketableUnengaged Non-MarketableNon-MarketableOpt-Out: ManualAll other existing statuses e.g. List Unsubscribe, Opt-Out System etc will remain non-marketable and in their existing state.
// returns a CompaniesItemWithCompanyPatchResponseable when successful
// returns a CompaniesItemWithCompany401Error error when the service returns a 401 status code
// returns a CompaniesItemWithCompany403Error error when the service returns a 403 status code
// returns a CompaniesItemWithCompany404Error error when the service returns a 404 status code
func (m *CompaniesWithCompanyItemRequestBuilder) PatchAsWithCompanyPatchResponse(ctx context.Context, body CompaniesItemWithCompanyPatchRequestBodyable, requestConfiguration *CompaniesWithCompanyItemRequestBuilderPatchRequestConfiguration) (CompaniesItemWithCompanyPatchResponseable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateCompaniesItemWithCompany401ErrorFromDiscriminatorValue,
		"403": CreateCompaniesItemWithCompany403ErrorFromDiscriminatorValue,
		"404": CreateCompaniesItemWithCompany404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCompaniesItemWithCompanyPatchResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(CompaniesItemWithCompanyPatchResponseable), nil
}

// ToGetRequestInformation retrieves a single company
// returns a *RequestInformation when successful
func (m *CompaniesWithCompanyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesWithCompanyItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToPatchRequestInformation updates a company with only the values provided in the request.You may opt-in or mark a Company as _Marketable_ by including the following field in the request JSON with an opt-in reason. (This field is also shown in the complete request body sample.) The reason you provide here will help with compliance. Example reasons: "Customer opted-in through webform", "Company gave explicit permission."```json"opt_in_reason": "your reason for opt-in"```Note that the email address status will only be updated to unconfirmed (marketable) for email addresses that are currently in the following states:Unengaged MarketableUnengaged Non-MarketableNon-MarketableOpt-Out: ManualAll other existing statuses e.g. List Unsubscribe, Opt-Out System etc will remain non-marketable and in their existing state.
// returns a *RequestInformation when successful
func (m *CompaniesWithCompanyItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body CompaniesItemWithCompanyPatchRequestBodyable, requestConfiguration *CompaniesWithCompanyItemRequestBuilderPatchRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/v1/companies/{companyId}", m.BaseRequestBuilder.PathParameters)
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
// returns a *CompaniesWithCompanyItemRequestBuilder when successful
func (m *CompaniesWithCompanyItemRequestBuilder) WithUrl(rawUrl string) *CompaniesWithCompanyItemRequestBuilder {
	return NewCompaniesWithCompanyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
