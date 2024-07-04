package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// OpportunitiesModelRequestBuilder builds and executes requests for operations under \v1\opportunities\model
type OpportunitiesModelRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// OpportunitiesModelRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OpportunitiesModelRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewOpportunitiesModelRequestBuilderInternal instantiates a new OpportunitiesModelRequestBuilder and sets the default values.
func NewOpportunitiesModelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *OpportunitiesModelRequestBuilder {
	m := &OpportunitiesModelRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/opportunities/model", pathParameters),
	}
	return m
}

// NewOpportunitiesModelRequestBuilder instantiates a new OpportunitiesModelRequestBuilder and sets the default values.
func NewOpportunitiesModelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *OpportunitiesModelRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewOpportunitiesModelRequestBuilderInternal(urlParams, requestAdapter)
}

// Get get the custom fields for the Opportunity object
// Deprecated: This method is obsolete. Use GetAsModelGetResponse instead.
// returns a OpportunitiesModelResponseable when successful
// returns a OpportunitiesModel401Error error when the service returns a 401 status code
// returns a OpportunitiesModel403Error error when the service returns a 403 status code
// returns a OpportunitiesModel404Error error when the service returns a 404 status code
func (m *OpportunitiesModelRequestBuilder) Get(ctx context.Context, requestConfiguration *OpportunitiesModelRequestBuilderGetRequestConfiguration) (OpportunitiesModelResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateOpportunitiesModel401ErrorFromDiscriminatorValue,
		"403": CreateOpportunitiesModel403ErrorFromDiscriminatorValue,
		"404": CreateOpportunitiesModel404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOpportunitiesModelResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(OpportunitiesModelResponseable), nil
}

// GetAsModelGetResponse get the custom fields for the Opportunity object
// returns a OpportunitiesModelGetResponseable when successful
// returns a OpportunitiesModel401Error error when the service returns a 401 status code
// returns a OpportunitiesModel403Error error when the service returns a 403 status code
// returns a OpportunitiesModel404Error error when the service returns a 404 status code
func (m *OpportunitiesModelRequestBuilder) GetAsModelGetResponse(ctx context.Context, requestConfiguration *OpportunitiesModelRequestBuilderGetRequestConfiguration) (OpportunitiesModelGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateOpportunitiesModel401ErrorFromDiscriminatorValue,
		"403": CreateOpportunitiesModel403ErrorFromDiscriminatorValue,
		"404": CreateOpportunitiesModel404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOpportunitiesModelGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(OpportunitiesModelGetResponseable), nil
}

// ToGetRequestInformation get the custom fields for the Opportunity object
// returns a *RequestInformation when successful
func (m *OpportunitiesModelRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OpportunitiesModelRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *OpportunitiesModelRequestBuilder when successful
func (m *OpportunitiesModelRequestBuilder) WithUrl(rawUrl string) *OpportunitiesModelRequestBuilder {
	return NewOpportunitiesModelRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
