package v2

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// TagsCategoriesWithTag_category_ItemRequestBuilder builds and executes requests for operations under \v2\tags\categories\{tag_category_id}
type TagsCategoriesWithTag_category_ItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// TagsCategoriesWithTag_category_ItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TagsCategoriesWithTag_category_ItemRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// TagsCategoriesWithTag_category_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TagsCategoriesWithTag_category_ItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// TagsCategoriesWithTag_category_ItemRequestBuilderPatchQueryParameters updates a Tag Category with only the values provided in the request
type TagsCategoriesWithTag_category_ItemRequestBuilderPatchQueryParameters struct {
	// An optional list of fields to be updated. If set, only the fields provided in the update_mask will be updated and others will be skipped.
	Update_mask []string `uriparametername:"update_mask"`
}

// TagsCategoriesWithTag_category_ItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TagsCategoriesWithTag_category_ItemRequestBuilderPatchRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *TagsCategoriesWithTag_category_ItemRequestBuilderPatchQueryParameters
}

// NewTagsCategoriesWithTag_category_ItemRequestBuilderInternal instantiates a new TagsCategoriesWithTag_category_ItemRequestBuilder and sets the default values.
func NewTagsCategoriesWithTag_category_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TagsCategoriesWithTag_category_ItemRequestBuilder {
	m := &TagsCategoriesWithTag_category_ItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/tags/categories/{tag_category_id}", pathParameters),
	}
	return m
}

// NewTagsCategoriesWithTag_category_ItemRequestBuilder instantiates a new TagsCategoriesWithTag_category_ItemRequestBuilder and sets the default values.
func NewTagsCategoriesWithTag_category_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TagsCategoriesWithTag_category_ItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewTagsCategoriesWithTag_category_ItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete deletes the specified Tag Category
// returns a TagsCategoriesItemWithTag_category_401Error error when the service returns a 401 status code
// returns a TagsCategoriesItemWithTag_category_403Error error when the service returns a 403 status code
// returns a TagsCategoriesItemWithTag_category_404Error error when the service returns a 404 status code
func (m *TagsCategoriesWithTag_category_ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TagsCategoriesWithTag_category_ItemRequestBuilderDeleteRequestConfiguration) error {
	requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTagsCategoriesItemWithTag_category_401ErrorFromDiscriminatorValue,
		"403": CreateTagsCategoriesItemWithTag_category_403ErrorFromDiscriminatorValue,
		"404": CreateTagsCategoriesItemWithTag_category_404ErrorFromDiscriminatorValue,
	}
	err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
	if err != nil {
		return err
	}
	return nil
}

// Get returns information about the specified Tag Category
// Deprecated: This method is obsolete. Use GetAsWithTag_category_GetResponse instead.
// returns a TagsCategoriesItemWithTag_category_Responseable when successful
// returns a TagsCategoriesItemWithTag_category_401Error error when the service returns a 401 status code
// returns a TagsCategoriesItemWithTag_category_403Error error when the service returns a 403 status code
// returns a TagsCategoriesItemWithTag_category_404Error error when the service returns a 404 status code
func (m *TagsCategoriesWithTag_category_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TagsCategoriesWithTag_category_ItemRequestBuilderGetRequestConfiguration) (TagsCategoriesItemWithTag_category_Responseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTagsCategoriesItemWithTag_category_401ErrorFromDiscriminatorValue,
		"403": CreateTagsCategoriesItemWithTag_category_403ErrorFromDiscriminatorValue,
		"404": CreateTagsCategoriesItemWithTag_category_404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTagsCategoriesItemWithTag_category_ResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TagsCategoriesItemWithTag_category_Responseable), nil
}

// GetAsWithTag_category_GetResponse returns information about the specified Tag Category
// returns a TagsCategoriesItemWithTag_category_GetResponseable when successful
// returns a TagsCategoriesItemWithTag_category_401Error error when the service returns a 401 status code
// returns a TagsCategoriesItemWithTag_category_403Error error when the service returns a 403 status code
// returns a TagsCategoriesItemWithTag_category_404Error error when the service returns a 404 status code
func (m *TagsCategoriesWithTag_category_ItemRequestBuilder) GetAsWithTag_category_GetResponse(ctx context.Context, requestConfiguration *TagsCategoriesWithTag_category_ItemRequestBuilderGetRequestConfiguration) (TagsCategoriesItemWithTag_category_GetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTagsCategoriesItemWithTag_category_401ErrorFromDiscriminatorValue,
		"403": CreateTagsCategoriesItemWithTag_category_403ErrorFromDiscriminatorValue,
		"404": CreateTagsCategoriesItemWithTag_category_404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTagsCategoriesItemWithTag_category_GetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TagsCategoriesItemWithTag_category_GetResponseable), nil
}

// Patch updates a Tag Category with only the values provided in the request
// Deprecated: This method is obsolete. Use PatchAsWithTag_category_PatchResponse instead.
// returns a TagsCategoriesItemWithTag_category_Responseable when successful
// returns a TagsCategoriesItemWithTag_category_401Error error when the service returns a 401 status code
// returns a TagsCategoriesItemWithTag_category_403Error error when the service returns a 403 status code
// returns a TagsCategoriesItemWithTag_category_404Error error when the service returns a 404 status code
func (m *TagsCategoriesWithTag_category_ItemRequestBuilder) Patch(ctx context.Context, body TagsCategoriesItemWithTag_category_PatchRequestBodyable, requestConfiguration *TagsCategoriesWithTag_category_ItemRequestBuilderPatchRequestConfiguration) (TagsCategoriesItemWithTag_category_Responseable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTagsCategoriesItemWithTag_category_401ErrorFromDiscriminatorValue,
		"403": CreateTagsCategoriesItemWithTag_category_403ErrorFromDiscriminatorValue,
		"404": CreateTagsCategoriesItemWithTag_category_404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTagsCategoriesItemWithTag_category_ResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TagsCategoriesItemWithTag_category_Responseable), nil
}

// PatchAsWithTag_category_PatchResponse updates a Tag Category with only the values provided in the request
// returns a TagsCategoriesItemWithTag_category_PatchResponseable when successful
// returns a TagsCategoriesItemWithTag_category_401Error error when the service returns a 401 status code
// returns a TagsCategoriesItemWithTag_category_403Error error when the service returns a 403 status code
// returns a TagsCategoriesItemWithTag_category_404Error error when the service returns a 404 status code
func (m *TagsCategoriesWithTag_category_ItemRequestBuilder) PatchAsWithTag_category_PatchResponse(ctx context.Context, body TagsCategoriesItemWithTag_category_PatchRequestBodyable, requestConfiguration *TagsCategoriesWithTag_category_ItemRequestBuilderPatchRequestConfiguration) (TagsCategoriesItemWithTag_category_PatchResponseable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTagsCategoriesItemWithTag_category_401ErrorFromDiscriminatorValue,
		"403": CreateTagsCategoriesItemWithTag_category_403ErrorFromDiscriminatorValue,
		"404": CreateTagsCategoriesItemWithTag_category_404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTagsCategoriesItemWithTag_category_PatchResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TagsCategoriesItemWithTag_category_PatchResponseable), nil
}

// ToDeleteRequestInformation deletes the specified Tag Category
// returns a *RequestInformation when successful
func (m *TagsCategoriesWithTag_category_ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TagsCategoriesWithTag_category_ItemRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// ToGetRequestInformation returns information about the specified Tag Category
// returns a *RequestInformation when successful
func (m *TagsCategoriesWithTag_category_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TagsCategoriesWithTag_category_ItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// ToPatchRequestInformation updates a Tag Category with only the values provided in the request
// returns a *RequestInformation when successful
func (m *TagsCategoriesWithTag_category_ItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body TagsCategoriesItemWithTag_category_PatchRequestBodyable, requestConfiguration *TagsCategoriesWithTag_category_ItemRequestBuilderPatchRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/v2/tags/categories/{tag_category_id}{?update_mask*}", m.BaseRequestBuilder.PathParameters)
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
// returns a *TagsCategoriesWithTag_category_ItemRequestBuilder when successful
func (m *TagsCategoriesWithTag_category_ItemRequestBuilder) WithUrl(rawUrl string) *TagsCategoriesWithTag_category_ItemRequestBuilder {
	return NewTagsCategoriesWithTag_category_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
