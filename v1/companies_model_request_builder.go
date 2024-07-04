package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// CompaniesModelRequestBuilder builds and executes requests for operations under \v1\companies\model
type CompaniesModelRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// CompaniesModelRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompaniesModelRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewCompaniesModelRequestBuilderInternal instantiates a new CompaniesModelRequestBuilder and sets the default values.
func NewCompaniesModelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CompaniesModelRequestBuilder {
	m := &CompaniesModelRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/companies/model", pathParameters),
	}
	return m
}

// NewCompaniesModelRequestBuilder instantiates a new CompaniesModelRequestBuilder and sets the default values.
func NewCompaniesModelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *CompaniesModelRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewCompaniesModelRequestBuilderInternal(urlParams, requestAdapter)
}

// Get get the custom fields and optional properties for the Company object
// Deprecated: This method is obsolete. Use GetAsModelGetResponse instead.
// returns a CompaniesModelResponseable when successful
// returns a CompaniesModel401Error error when the service returns a 401 status code
// returns a CompaniesModel403Error error when the service returns a 403 status code
// returns a CompaniesModel404Error error when the service returns a 404 status code
func (m *CompaniesModelRequestBuilder) Get(ctx context.Context, requestConfiguration *CompaniesModelRequestBuilderGetRequestConfiguration) (CompaniesModelResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateCompaniesModel401ErrorFromDiscriminatorValue,
		"403": CreateCompaniesModel403ErrorFromDiscriminatorValue,
		"404": CreateCompaniesModel404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCompaniesModelResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(CompaniesModelResponseable), nil
}

// GetAsModelGetResponse get the custom fields and optional properties for the Company object
// returns a CompaniesModelGetResponseable when successful
// returns a CompaniesModel401Error error when the service returns a 401 status code
// returns a CompaniesModel403Error error when the service returns a 403 status code
// returns a CompaniesModel404Error error when the service returns a 404 status code
func (m *CompaniesModelRequestBuilder) GetAsModelGetResponse(ctx context.Context, requestConfiguration *CompaniesModelRequestBuilderGetRequestConfiguration) (CompaniesModelGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateCompaniesModel401ErrorFromDiscriminatorValue,
		"403": CreateCompaniesModel403ErrorFromDiscriminatorValue,
		"404": CreateCompaniesModel404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCompaniesModelGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(CompaniesModelGetResponseable), nil
}

// ToGetRequestInformation get the custom fields and optional properties for the Company object
// returns a *RequestInformation when successful
func (m *CompaniesModelRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompaniesModelRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CompaniesModelRequestBuilder when successful
func (m *CompaniesModelRequestBuilder) WithUrl(rawUrl string) *CompaniesModelRequestBuilder {
	return NewCompaniesModelRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
