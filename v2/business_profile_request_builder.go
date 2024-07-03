package v2

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// BusinessProfileRequestBuilder builds and executes requests for operations under \v2\businessProfile
type BusinessProfileRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// BusinessProfileRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BusinessProfileRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// BusinessProfileRequestBuilderPatchQueryParameters updates Business Profile information.
type BusinessProfileRequestBuilderPatchQueryParameters struct {
	// An optional list of fields to be updated. If set, only the fields provided in the update_mask will be updated and others will be skipped.
	Update_mask []string `uriparametername:"update_mask"`
}

// BusinessProfileRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BusinessProfileRequestBuilderPatchRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *BusinessProfileRequestBuilderPatchQueryParameters
}

// NewBusinessProfileRequestBuilderInternal instantiates a new BusinessProfileRequestBuilder and sets the default values.
func NewBusinessProfileRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *BusinessProfileRequestBuilder {
	m := &BusinessProfileRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/businessProfile", pathParameters),
	}
	return m
}

// NewBusinessProfileRequestBuilder instantiates a new BusinessProfileRequestBuilder and sets the default values.
func NewBusinessProfileRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *BusinessProfileRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewBusinessProfileRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieves Business Profile information.
// Deprecated: This method is obsolete. Use GetAsBusinessProfileGetResponse instead.
// returns a BusinessProfileResponseable when successful
// returns a BusinessProfile401Error error when the service returns a 401 status code
// returns a BusinessProfile403Error error when the service returns a 403 status code
// returns a BusinessProfile404Error error when the service returns a 404 status code
func (m *BusinessProfileRequestBuilder) Get(ctx context.Context, requestConfiguration *BusinessProfileRequestBuilderGetRequestConfiguration) (BusinessProfileResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateBusinessProfile401ErrorFromDiscriminatorValue,
		"403": CreateBusinessProfile403ErrorFromDiscriminatorValue,
		"404": CreateBusinessProfile404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateBusinessProfileResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(BusinessProfileResponseable), nil
}

// GetAsBusinessProfileGetResponse retrieves Business Profile information.
// returns a BusinessProfileGetResponseable when successful
// returns a BusinessProfile401Error error when the service returns a 401 status code
// returns a BusinessProfile403Error error when the service returns a 403 status code
// returns a BusinessProfile404Error error when the service returns a 404 status code
func (m *BusinessProfileRequestBuilder) GetAsBusinessProfileGetResponse(ctx context.Context, requestConfiguration *BusinessProfileRequestBuilderGetRequestConfiguration) (BusinessProfileGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateBusinessProfile401ErrorFromDiscriminatorValue,
		"403": CreateBusinessProfile403ErrorFromDiscriminatorValue,
		"404": CreateBusinessProfile404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateBusinessProfileGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(BusinessProfileGetResponseable), nil
}

// Patch updates Business Profile information.
// Deprecated: This method is obsolete. Use PatchAsBusinessProfilePatchResponse instead.
// returns a BusinessProfileResponseable when successful
// returns a BusinessProfile401Error error when the service returns a 401 status code
// returns a BusinessProfile403Error error when the service returns a 403 status code
// returns a BusinessProfile404Error error when the service returns a 404 status code
func (m *BusinessProfileRequestBuilder) Patch(ctx context.Context, body BusinessProfilePatchRequestBodyable, requestConfiguration *BusinessProfileRequestBuilderPatchRequestConfiguration) (BusinessProfileResponseable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateBusinessProfile401ErrorFromDiscriminatorValue,
		"403": CreateBusinessProfile403ErrorFromDiscriminatorValue,
		"404": CreateBusinessProfile404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateBusinessProfileResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(BusinessProfileResponseable), nil
}

// PatchAsBusinessProfilePatchResponse updates Business Profile information.
// returns a BusinessProfilePatchResponseable when successful
// returns a BusinessProfile401Error error when the service returns a 401 status code
// returns a BusinessProfile403Error error when the service returns a 403 status code
// returns a BusinessProfile404Error error when the service returns a 404 status code
func (m *BusinessProfileRequestBuilder) PatchAsBusinessProfilePatchResponse(ctx context.Context, body BusinessProfilePatchRequestBodyable, requestConfiguration *BusinessProfileRequestBuilderPatchRequestConfiguration) (BusinessProfilePatchResponseable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateBusinessProfile401ErrorFromDiscriminatorValue,
		"403": CreateBusinessProfile403ErrorFromDiscriminatorValue,
		"404": CreateBusinessProfile404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateBusinessProfilePatchResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(BusinessProfilePatchResponseable), nil
}

// ToGetRequestInformation retrieves Business Profile information.
// returns a *RequestInformation when successful
func (m *BusinessProfileRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BusinessProfileRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// ToPatchRequestInformation updates Business Profile information.
// returns a *RequestInformation when successful
func (m *BusinessProfileRequestBuilder) ToPatchRequestInformation(ctx context.Context, body BusinessProfilePatchRequestBodyable, requestConfiguration *BusinessProfileRequestBuilderPatchRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/v2/businessProfile{?update_mask*}", m.BaseRequestBuilder.PathParameters)
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
// returns a *BusinessProfileRequestBuilder when successful
func (m *BusinessProfileRequestBuilder) WithUrl(rawUrl string) *BusinessProfileRequestBuilder {
	return NewBusinessProfileRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
