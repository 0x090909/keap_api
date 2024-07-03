package v2

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// TagsCategoriesRequestBuilder builds and executes requests for operations under \v2\tags\categories
type TagsCategoriesRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// TagsCategoriesRequestBuilderGetQueryParameters retrieves the list of defined Tag CategoriesTo search for `null` or empty fields use `filter==NONE`
type TagsCategoriesRequestBuilderGetQueryParameters struct {
	// Search filter to apply to results
	Filter *string `uriparametername:"filter"`
	// Attribute and direction to order items by. E.g. `given_name desc`
	Order_by *string `uriparametername:"order_by"`
	// Total number of items to return per page
	Page_size *int32 `uriparametername:"page_size"`
	// Page token
	Page_token *string `uriparametername:"page_token"`
}

// TagsCategoriesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TagsCategoriesRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *TagsCategoriesRequestBuilderGetQueryParameters
}

// TagsCategoriesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TagsCategoriesRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ByTag_category_id gets an item from the keapapi.v2.tags.categories.item collection
// returns a *TagsCategoriesWithTag_category_ItemRequestBuilder when successful
func (m *TagsCategoriesRequestBuilder) ByTag_category_id(tag_category_id string) *TagsCategoriesWithTag_category_ItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if tag_category_id != "" {
		urlTplParams["tag_category_id"] = tag_category_id
	}
	return NewTagsCategoriesWithTag_category_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewTagsCategoriesRequestBuilderInternal instantiates a new TagsCategoriesRequestBuilder and sets the default values.
func NewTagsCategoriesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TagsCategoriesRequestBuilder {
	m := &TagsCategoriesRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/tags/categories{?filter*,order_by*,page_size*,page_token*}", pathParameters),
	}
	return m
}

// NewTagsCategoriesRequestBuilder instantiates a new TagsCategoriesRequestBuilder and sets the default values.
func NewTagsCategoriesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TagsCategoriesRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewTagsCategoriesRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieves the list of defined Tag CategoriesTo search for `null` or empty fields use `filter==NONE`
// Deprecated: This method is obsolete. Use GetAsCategoriesGetResponse instead.
// returns a TagsCategoriesResponseable when successful
// returns a TagsCategories401Error error when the service returns a 401 status code
// returns a TagsCategories403Error error when the service returns a 403 status code
// returns a TagsCategories404Error error when the service returns a 404 status code
func (m *TagsCategoriesRequestBuilder) Get(ctx context.Context, requestConfiguration *TagsCategoriesRequestBuilderGetRequestConfiguration) (TagsCategoriesResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTagsCategories401ErrorFromDiscriminatorValue,
		"403": CreateTagsCategories403ErrorFromDiscriminatorValue,
		"404": CreateTagsCategories404ErrorFromDiscriminatorValue,
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

// GetAsCategoriesGetResponse retrieves the list of defined Tag CategoriesTo search for `null` or empty fields use `filter==NONE`
// returns a TagsCategoriesGetResponseable when successful
// returns a TagsCategories401Error error when the service returns a 401 status code
// returns a TagsCategories403Error error when the service returns a 403 status code
// returns a TagsCategories404Error error when the service returns a 404 status code
func (m *TagsCategoriesRequestBuilder) GetAsCategoriesGetResponse(ctx context.Context, requestConfiguration *TagsCategoriesRequestBuilderGetRequestConfiguration) (TagsCategoriesGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTagsCategories401ErrorFromDiscriminatorValue,
		"403": CreateTagsCategories403ErrorFromDiscriminatorValue,
		"404": CreateTagsCategories404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTagsCategoriesGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TagsCategoriesGetResponseable), nil
}

// Post creates a new Tag Category.
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

// PostAsCategoriesPostResponse creates a new Tag Category.
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

// ToGetRequestInformation retrieves the list of defined Tag CategoriesTo search for `null` or empty fields use `filter==NONE`
// returns a *RequestInformation when successful
func (m *TagsCategoriesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TagsCategoriesRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToPostRequestInformation creates a new Tag Category.
// returns a *RequestInformation when successful
func (m *TagsCategoriesRequestBuilder) ToPostRequestInformation(ctx context.Context, body TagsCategoriesPostRequestBodyable, requestConfiguration *TagsCategoriesRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/v2/tags/categories", m.BaseRequestBuilder.PathParameters)
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
