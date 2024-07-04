package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// TagsIdItemRequestBuilder builds and executes requests for operations under \v1\tags\{id-id}
type TagsIdItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// TagsIdItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TagsIdItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// Companies the companies property
// returns a *TagsItemCompaniesRequestBuilder when successful
func (m *TagsIdItemRequestBuilder) Companies() *TagsItemCompaniesRequestBuilder {
	return NewTagsItemCompaniesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// NewTagsIdItemRequestBuilderInternal instantiates a new TagsIdItemRequestBuilder and sets the default values.
func NewTagsIdItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TagsIdItemRequestBuilder {
	m := &TagsIdItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/tags/{id%2Did}", pathParameters),
	}
	return m
}

// NewTagsIdItemRequestBuilder instantiates a new TagsIdItemRequestBuilder and sets the default values.
func NewTagsIdItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TagsIdItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewTagsIdItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Contacts the contacts property
// returns a *TagsItemContactsRequestBuilder when successful
func (m *TagsIdItemRequestBuilder) Contacts() *TagsItemContactsRequestBuilder {
	return NewTagsItemContactsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Get retrieves a single tag
// Deprecated: This method is obsolete. Use GetAsIdGetResponse instead.
// returns a TagsItemIdResponseable when successful
// returns a TagsItemId401Error error when the service returns a 401 status code
// returns a TagsItemId403Error error when the service returns a 403 status code
// returns a TagsItemId404Error error when the service returns a 404 status code
func (m *TagsIdItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TagsIdItemRequestBuilderGetRequestConfiguration) (TagsItemIdResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTagsItemId401ErrorFromDiscriminatorValue,
		"403": CreateTagsItemId403ErrorFromDiscriminatorValue,
		"404": CreateTagsItemId404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTagsItemIdResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TagsItemIdResponseable), nil
}

// GetAsIdGetResponse retrieves a single tag
// returns a TagsItemIdGetResponseable when successful
// returns a TagsItemId401Error error when the service returns a 401 status code
// returns a TagsItemId403Error error when the service returns a 403 status code
// returns a TagsItemId404Error error when the service returns a 404 status code
func (m *TagsIdItemRequestBuilder) GetAsIdGetResponse(ctx context.Context, requestConfiguration *TagsIdItemRequestBuilderGetRequestConfiguration) (TagsItemIdGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTagsItemId401ErrorFromDiscriminatorValue,
		"403": CreateTagsItemId403ErrorFromDiscriminatorValue,
		"404": CreateTagsItemId404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTagsItemIdGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TagsItemIdGetResponseable), nil
}

// ToGetRequestInformation retrieves a single tag
// returns a *RequestInformation when successful
func (m *TagsIdItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TagsIdItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TagsIdItemRequestBuilder when successful
func (m *TagsIdItemRequestBuilder) WithUrl(rawUrl string) *TagsIdItemRequestBuilder {
	return NewTagsIdItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
