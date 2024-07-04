package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AffiliatesModelRequestBuilder builds and executes requests for operations under \v1\affiliates\model
type AffiliatesModelRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// AffiliatesModelRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AffiliatesModelRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewAffiliatesModelRequestBuilderInternal instantiates a new AffiliatesModelRequestBuilder and sets the default values.
func NewAffiliatesModelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesModelRequestBuilder {
	m := &AffiliatesModelRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/affiliates/model", pathParameters),
	}
	return m
}

// NewAffiliatesModelRequestBuilder instantiates a new AffiliatesModelRequestBuilder and sets the default values.
func NewAffiliatesModelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesModelRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewAffiliatesModelRequestBuilderInternal(urlParams, requestAdapter)
}

// Get get the custom fields for the Affiliate object
// Deprecated: This method is obsolete. Use GetAsModelGetResponse instead.
// returns a AffiliatesModelResponseable when successful
// returns a AffiliatesModel401Error error when the service returns a 401 status code
// returns a AffiliatesModel403Error error when the service returns a 403 status code
// returns a AffiliatesModel404Error error when the service returns a 404 status code
func (m *AffiliatesModelRequestBuilder) Get(ctx context.Context, requestConfiguration *AffiliatesModelRequestBuilderGetRequestConfiguration) (AffiliatesModelResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAffiliatesModel401ErrorFromDiscriminatorValue,
		"403": CreateAffiliatesModel403ErrorFromDiscriminatorValue,
		"404": CreateAffiliatesModel404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAffiliatesModelResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AffiliatesModelResponseable), nil
}

// GetAsModelGetResponse get the custom fields for the Affiliate object
// returns a AffiliatesModelGetResponseable when successful
// returns a AffiliatesModel401Error error when the service returns a 401 status code
// returns a AffiliatesModel403Error error when the service returns a 403 status code
// returns a AffiliatesModel404Error error when the service returns a 404 status code
func (m *AffiliatesModelRequestBuilder) GetAsModelGetResponse(ctx context.Context, requestConfiguration *AffiliatesModelRequestBuilderGetRequestConfiguration) (AffiliatesModelGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAffiliatesModel401ErrorFromDiscriminatorValue,
		"403": CreateAffiliatesModel403ErrorFromDiscriminatorValue,
		"404": CreateAffiliatesModel404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAffiliatesModelGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AffiliatesModelGetResponseable), nil
}

// ToGetRequestInformation get the custom fields for the Affiliate object
// returns a *RequestInformation when successful
func (m *AffiliatesModelRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AffiliatesModelRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *AffiliatesModelRequestBuilder when successful
func (m *AffiliatesModelRequestBuilder) WithUrl(rawUrl string) *AffiliatesModelRequestBuilder {
	return NewAffiliatesModelRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
