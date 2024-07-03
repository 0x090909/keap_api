package v2

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CompaniesWithCompany_ItemRequestBuilder builds and executes requests for operations under \v2\companies\{company_id}
type CompaniesWithCompany_ItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// CompaniesWithCompany_ItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesWithCompany_ItemRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// CompaniesWithCompany_ItemRequestBuilderGetQueryParameters retrieves a single Company
type CompaniesWithCompany_ItemRequestBuilderGetQueryParameters struct {
	// Comma-delimited list of Company properties to include in the response. (Available fields are: `company_name`, `address`, `custom_fields`, `email_address`, `fax_number`, `phone_number`, `website`, `notes`)
	Fields []string `uriparametername:"fields"`
}

// CompaniesWithCompany_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesWithCompany_ItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *CompaniesWithCompany_ItemRequestBuilderGetQueryParameters
}

// CompaniesWithCompany_ItemRequestBuilderPatchQueryParameters updates a Company with the values provided in the request
type CompaniesWithCompany_ItemRequestBuilderPatchQueryParameters struct {
	// An optional list of properties to be updated. If set, only the provided properties will be updated and others will be skipped.
	Update_mask []string `uriparametername:"update_mask"`
}

// CompaniesWithCompany_ItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesWithCompany_ItemRequestBuilderPatchRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *CompaniesWithCompany_ItemRequestBuilderPatchQueryParameters
}

// NewCompaniesWithCompany_ItemRequestBuilderInternal instantiates a new CompaniesWithCompany_ItemRequestBuilder and sets the default values.
func NewCompaniesWithCompany_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CompaniesWithCompany_ItemRequestBuilder {
	m := &CompaniesWithCompany_ItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/companies/{company_id}{?fields*}", pathParameters),
	}
	return m
}

// NewCompaniesWithCompany_ItemRequestBuilder instantiates a new CompaniesWithCompany_ItemRequestBuilder and sets the default values.
func NewCompaniesWithCompany_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CompaniesWithCompany_ItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewCompaniesWithCompany_ItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete deletes the specified Company
// returns a CompaniesItemWithCompany_401Error error when the service returns a 401 status code
// returns a CompaniesItemWithCompany_403Error error when the service returns a 403 status code
// returns a CompaniesItemWithCompany_404Error error when the service returns a 404 status code
func (m *CompaniesWithCompany_ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CompaniesWithCompany_ItemRequestBuilderDeleteRequestConfiguration) error {
	requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateCompaniesItemWithCompany_401ErrorFromDiscriminatorValue,
		"403": CreateCompaniesItemWithCompany_403ErrorFromDiscriminatorValue,
		"404": CreateCompaniesItemWithCompany_404ErrorFromDiscriminatorValue,
	}
	err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
	if err != nil {
		return err
	}
	return nil
}

// Get retrieves a single Company
// Deprecated: This method is obsolete. Use GetAsWithCompany_GetResponse instead.
// returns a CompaniesItemWithCompany_Responseable when successful
// returns a CompaniesItemWithCompany_401Error error when the service returns a 401 status code
// returns a CompaniesItemWithCompany_403Error error when the service returns a 403 status code
// returns a CompaniesItemWithCompany_404Error error when the service returns a 404 status code
func (m *CompaniesWithCompany_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesWithCompany_ItemRequestBuilderGetRequestConfiguration) (CompaniesItemWithCompany_Responseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateCompaniesItemWithCompany_401ErrorFromDiscriminatorValue,
		"403": CreateCompaniesItemWithCompany_403ErrorFromDiscriminatorValue,
		"404": CreateCompaniesItemWithCompany_404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCompaniesItemWithCompany_ResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(CompaniesItemWithCompany_Responseable), nil
}

// GetAsWithCompany_GetResponse retrieves a single Company
// returns a CompaniesItemWithCompany_GetResponseable when successful
// returns a CompaniesItemWithCompany_401Error error when the service returns a 401 status code
// returns a CompaniesItemWithCompany_403Error error when the service returns a 403 status code
// returns a CompaniesItemWithCompany_404Error error when the service returns a 404 status code
func (m *CompaniesWithCompany_ItemRequestBuilder) GetAsWithCompany_GetResponse(ctx context.Context, requestConfiguration *CompaniesWithCompany_ItemRequestBuilderGetRequestConfiguration) (CompaniesItemWithCompany_GetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateCompaniesItemWithCompany_401ErrorFromDiscriminatorValue,
		"403": CreateCompaniesItemWithCompany_403ErrorFromDiscriminatorValue,
		"404": CreateCompaniesItemWithCompany_404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCompaniesItemWithCompany_GetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(CompaniesItemWithCompany_GetResponseable), nil
}

// Patch updates a Company with the values provided in the request
// Deprecated: This method is obsolete. Use PatchAsWithCompany_PatchResponse instead.
// returns a CompaniesItemWithCompany_Responseable when successful
// returns a CompaniesItemWithCompany_401Error error when the service returns a 401 status code
// returns a CompaniesItemWithCompany_403Error error when the service returns a 403 status code
// returns a CompaniesItemWithCompany_404Error error when the service returns a 404 status code
func (m *CompaniesWithCompany_ItemRequestBuilder) Patch(ctx context.Context, body CompaniesItemWithCompany_PatchRequestBodyable, requestConfiguration *CompaniesWithCompany_ItemRequestBuilderPatchRequestConfiguration) (CompaniesItemWithCompany_Responseable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateCompaniesItemWithCompany_401ErrorFromDiscriminatorValue,
		"403": CreateCompaniesItemWithCompany_403ErrorFromDiscriminatorValue,
		"404": CreateCompaniesItemWithCompany_404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCompaniesItemWithCompany_ResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(CompaniesItemWithCompany_Responseable), nil
}

// PatchAsWithCompany_PatchResponse updates a Company with the values provided in the request
// returns a CompaniesItemWithCompany_PatchResponseable when successful
// returns a CompaniesItemWithCompany_401Error error when the service returns a 401 status code
// returns a CompaniesItemWithCompany_403Error error when the service returns a 403 status code
// returns a CompaniesItemWithCompany_404Error error when the service returns a 404 status code
func (m *CompaniesWithCompany_ItemRequestBuilder) PatchAsWithCompany_PatchResponse(ctx context.Context, body CompaniesItemWithCompany_PatchRequestBodyable, requestConfiguration *CompaniesWithCompany_ItemRequestBuilderPatchRequestConfiguration) (CompaniesItemWithCompany_PatchResponseable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateCompaniesItemWithCompany_401ErrorFromDiscriminatorValue,
		"403": CreateCompaniesItemWithCompany_403ErrorFromDiscriminatorValue,
		"404": CreateCompaniesItemWithCompany_404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCompaniesItemWithCompany_PatchResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(CompaniesItemWithCompany_PatchResponseable), nil
}

// ToDeleteRequestInformation deletes the specified Company
// returns a *RequestInformation when successful
func (m *CompaniesWithCompany_ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CompaniesWithCompany_ItemRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/v2/companies/{company_id}", m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// ToGetRequestInformation retrieves a single Company
// returns a *RequestInformation when successful
func (m *CompaniesWithCompany_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesWithCompany_ItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToPatchRequestInformation updates a Company with the values provided in the request
// returns a *RequestInformation when successful
func (m *CompaniesWithCompany_ItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body CompaniesItemWithCompany_PatchRequestBodyable, requestConfiguration *CompaniesWithCompany_ItemRequestBuilderPatchRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/v2/companies/{company_id}{?update_mask*}", m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		if requestConfiguration.QueryParameters != nil {
			requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
		}
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
// returns a *CompaniesWithCompany_ItemRequestBuilder when successful
func (m *CompaniesWithCompany_ItemRequestBuilder) WithUrl(rawUrl string) *CompaniesWithCompany_ItemRequestBuilder {
	return NewCompaniesWithCompany_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
