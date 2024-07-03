package v2

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// TagsItemContactsApplyTagsRequestBuilder builds and executes requests for operations under \v2\tags\{tag_id}\contacts:applyTags
type TagsItemContactsApplyTagsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// TagsItemContactsApplyTagsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TagsItemContactsApplyTagsRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewTagsItemContactsApplyTagsRequestBuilderInternal instantiates a new TagsItemContactsApplyTagsRequestBuilder and sets the default values.
func NewTagsItemContactsApplyTagsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TagsItemContactsApplyTagsRequestBuilder {
	m := &TagsItemContactsApplyTagsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/tags/{tag_id}/contacts:applyTags", pathParameters),
	}
	return m
}

// NewTagsItemContactsApplyTagsRequestBuilder instantiates a new TagsItemContactsApplyTagsRequestBuilder and sets the default values.
func NewTagsItemContactsApplyTagsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TagsItemContactsApplyTagsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewTagsItemContactsApplyTagsRequestBuilderInternal(urlParams, requestAdapter)
}

// Post applies a Tag to a list of Contacts.
// Deprecated: This method is obsolete. Use PostAsContactsApplyTagsPostResponse instead.
// returns a TagsItemContactsApplyTagsResponseable when successful
// returns a TagsItemContactsApplyTags401Error error when the service returns a 401 status code
// returns a TagsItemContactsApplyTags403Error error when the service returns a 403 status code
func (m *TagsItemContactsApplyTagsRequestBuilder) Post(ctx context.Context, body TagsItemContactsApplyTagsPostRequestBodyable, requestConfiguration *TagsItemContactsApplyTagsRequestBuilderPostRequestConfiguration) (TagsItemContactsApplyTagsResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTagsItemContactsApplyTags401ErrorFromDiscriminatorValue,
		"403": CreateTagsItemContactsApplyTags403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTagsItemContactsApplyTagsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TagsItemContactsApplyTagsResponseable), nil
}

// PostAsContactsApplyTagsPostResponse applies a Tag to a list of Contacts.
// returns a TagsItemContactsApplyTagsPostResponseable when successful
// returns a TagsItemContactsApplyTags401Error error when the service returns a 401 status code
// returns a TagsItemContactsApplyTags403Error error when the service returns a 403 status code
func (m *TagsItemContactsApplyTagsRequestBuilder) PostAsContactsApplyTagsPostResponse(ctx context.Context, body TagsItemContactsApplyTagsPostRequestBodyable, requestConfiguration *TagsItemContactsApplyTagsRequestBuilderPostRequestConfiguration) (TagsItemContactsApplyTagsPostResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTagsItemContactsApplyTags401ErrorFromDiscriminatorValue,
		"403": CreateTagsItemContactsApplyTags403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTagsItemContactsApplyTagsPostResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TagsItemContactsApplyTagsPostResponseable), nil
}

// ToPostRequestInformation applies a Tag to a list of Contacts.
// returns a *RequestInformation when successful
func (m *TagsItemContactsApplyTagsRequestBuilder) ToPostRequestInformation(ctx context.Context, body TagsItemContactsApplyTagsPostRequestBodyable, requestConfiguration *TagsItemContactsApplyTagsRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TagsItemContactsApplyTagsRequestBuilder when successful
func (m *TagsItemContactsApplyTagsRequestBuilder) WithUrl(rawUrl string) *TagsItemContactsApplyTagsRequestBuilder {
	return NewTagsItemContactsApplyTagsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
