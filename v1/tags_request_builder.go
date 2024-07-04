package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
)

// TagsRequestBuilder builds and executes requests for operations under \v1\tags
type TagsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// TagsRequestBuilderGetQueryParameters retrieve a list of tags defined in the application
type TagsRequestBuilderGetQueryParameters struct {
	// Category Id of tags to filter by
	Category *int64 `uriparametername:"category"`
	// Sets a total of items to return
	Limit *int32 `uriparametername:"limit"`
	// Filter for tags with a specific name
	Name *string `uriparametername:"name"`
	// Sets a beginning range of items to return
	Offset *int32 `uriparametername:"offset"`
}

// TagsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TagsRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *TagsRequestBuilderGetQueryParameters
}

// TagsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TagsRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ByIdId gets an item from the github.com/0x090909/keap_api.v1.tags.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *TagsIdItemRequestBuilder when successful
func (m *TagsRequestBuilder) ByIdId(idId string) *TagsIdItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if idId != "" {
		urlTplParams["id%2Did"] = idId
	}
	return NewTagsIdItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// ByIdIdInt64 gets an item from the github.com/0x090909/keap_api.v1.tags.item collection
// returns a *TagsIdItemRequestBuilder when successful
func (m *TagsRequestBuilder) ByIdIdInt64(idId int64) *TagsIdItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	urlTplParams["id%2Did"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(idId, 10)
	return NewTagsIdItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// Categories the categories property
// returns a *TagsCategoriesRequestBuilder when successful
func (m *TagsRequestBuilder) Categories() *TagsCategoriesRequestBuilder {
	return NewTagsCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// NewTagsRequestBuilderInternal instantiates a new TagsRequestBuilder and sets the default values.
func NewTagsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TagsRequestBuilder {
	m := &TagsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/tags{?category*,limit*,name*,offset*}", pathParameters),
	}
	return m
}

// NewTagsRequestBuilder instantiates a new TagsRequestBuilder and sets the default values.
func NewTagsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TagsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewTagsRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieve a list of tags defined in the application
// Deprecated: This method is obsolete. Use GetAsTagsGetResponse instead.
// returns a TagsResponseable when successful
// returns a Tags401Error error when the service returns a 401 status code
// returns a Tags403Error error when the service returns a 403 status code
// returns a Tags404Error error when the service returns a 404 status code
func (m *TagsRequestBuilder) Get(ctx context.Context, requestConfiguration *TagsRequestBuilderGetRequestConfiguration) (TagsResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTags401ErrorFromDiscriminatorValue,
		"403": CreateTags403ErrorFromDiscriminatorValue,
		"404": CreateTags404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTagsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TagsResponseable), nil
}

// GetAsTagsGetResponse retrieve a list of tags defined in the application
// returns a TagsGetResponseable when successful
// returns a Tags401Error error when the service returns a 401 status code
// returns a Tags403Error error when the service returns a 403 status code
// returns a Tags404Error error when the service returns a 404 status code
func (m *TagsRequestBuilder) GetAsTagsGetResponse(ctx context.Context, requestConfiguration *TagsRequestBuilderGetRequestConfiguration) (TagsGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTags401ErrorFromDiscriminatorValue,
		"403": CreateTags403ErrorFromDiscriminatorValue,
		"404": CreateTags404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTagsGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TagsGetResponseable), nil
}

// Post create a new tag
// Deprecated: This method is obsolete. Use PostAsTagsPostResponse instead.
// returns a TagsResponseable when successful
// returns a Tags401Error error when the service returns a 401 status code
// returns a Tags403Error error when the service returns a 403 status code
func (m *TagsRequestBuilder) Post(ctx context.Context, body TagsPostRequestBodyable, requestConfiguration *TagsRequestBuilderPostRequestConfiguration) (TagsResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTags401ErrorFromDiscriminatorValue,
		"403": CreateTags403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTagsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TagsResponseable), nil
}

// PostAsTagsPostResponse create a new tag
// returns a TagsPostResponseable when successful
// returns a Tags401Error error when the service returns a 401 status code
// returns a Tags403Error error when the service returns a 403 status code
func (m *TagsRequestBuilder) PostAsTagsPostResponse(ctx context.Context, body TagsPostRequestBodyable, requestConfiguration *TagsRequestBuilderPostRequestConfiguration) (TagsPostResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTags401ErrorFromDiscriminatorValue,
		"403": CreateTags403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTagsPostResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TagsPostResponseable), nil
}

// ToGetRequestInformation retrieve a list of tags defined in the application
// returns a *RequestInformation when successful
func (m *TagsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TagsRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToPostRequestInformation create a new tag
// returns a *RequestInformation when successful
func (m *TagsRequestBuilder) ToPostRequestInformation(ctx context.Context, body TagsPostRequestBodyable, requestConfiguration *TagsRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/v1/tags", m.BaseRequestBuilder.PathParameters)
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
// returns a *TagsRequestBuilder when successful
func (m *TagsRequestBuilder) WithUrl(rawUrl string) *TagsRequestBuilder {
	return NewTagsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
