package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// TagsCategoriesRequestBuilder builds and executes requests for operations under \v1\tags\categories
type TagsCategoriesRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// TagsCategoriesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TagsCategoriesRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewTagsCategoriesRequestBuilderInternal instantiates a new TagsCategoriesRequestBuilder and sets the default values.
func NewTagsCategoriesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TagsCategoriesRequestBuilder {
	m := &TagsCategoriesRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/tags/categories", pathParameters),
	}
	return m
}

// NewTagsCategoriesRequestBuilder instantiates a new TagsCategoriesRequestBuilder and sets the default values.
func NewTagsCategoriesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TagsCategoriesRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewTagsCategoriesRequestBuilderInternal(urlParams, requestAdapter)
}

// Post create a new tag category
// Deprecated: This method is obsolete. Use PostAsCategoriesPostResponse instead.
// returns a TagsCategoriesResponseable when successful
// returns a TagsCategories401Error error when the service returns a 401 status code
// returns a TagsCategories403Error error when the service returns a 403 status code
func (m *TagsCategoriesRequestBuilder) Post(ctx context.Context, body TagsCategoriesPostRequestBodyable, requestConfiguration *TagsCategoriesRequestBuilderPostRequestConfiguration) (TagsCategoriesResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTagsCategories401ErrorFromDiscriminatorValue,
		"403": CreateTagsCategories403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTagsCategoriesResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TagsCategoriesResponseable), nil
}

// PostAsCategoriesPostResponse create a new tag category
// returns a TagsCategoriesPostResponseable when successful
// returns a TagsCategories401Error error when the service returns a 401 status code
// returns a TagsCategories403Error error when the service returns a 403 status code
func (m *TagsCategoriesRequestBuilder) PostAsCategoriesPostResponse(ctx context.Context, body TagsCategoriesPostRequestBodyable, requestConfiguration *TagsCategoriesRequestBuilderPostRequestConfiguration) (TagsCategoriesPostResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTagsCategories401ErrorFromDiscriminatorValue,
		"403": CreateTagsCategories403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTagsCategoriesPostResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TagsCategoriesPostResponseable), nil
}

// ToPostRequestInformation create a new tag category
// returns a *RequestInformation when successful
func (m *TagsCategoriesRequestBuilder) ToPostRequestInformation(ctx context.Context, body TagsCategoriesPostRequestBodyable, requestConfiguration *TagsCategoriesRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *TagsCategoriesRequestBuilder when successful
func (m *TagsCategoriesRequestBuilder) WithUrl(rawUrl string) *TagsCategoriesRequestBuilder {
	return NewTagsCategoriesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
