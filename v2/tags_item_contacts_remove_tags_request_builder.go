package v2

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// TagsItemContactsRemoveTagsRequestBuilder builds and executes requests for operations under \v2\tags\{tag_id}\contacts:removeTags
type TagsItemContactsRemoveTagsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// TagsItemContactsRemoveTagsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TagsItemContactsRemoveTagsRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewTagsItemContactsRemoveTagsRequestBuilderInternal instantiates a new TagsItemContactsRemoveTagsRequestBuilder and sets the default values.
func NewTagsItemContactsRemoveTagsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TagsItemContactsRemoveTagsRequestBuilder {
	m := &TagsItemContactsRemoveTagsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/tags/{tag_id}/contacts:removeTags", pathParameters),
	}
	return m
}

// NewTagsItemContactsRemoveTagsRequestBuilder instantiates a new TagsItemContactsRemoveTagsRequestBuilder and sets the default values.
func NewTagsItemContactsRemoveTagsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TagsItemContactsRemoveTagsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewTagsItemContactsRemoveTagsRequestBuilderInternal(urlParams, requestAdapter)
}

// Post removes a Tag from a list of Contacts.
// returns a TagsItemContactsRemoveTags401Error error when the service returns a 401 status code
// returns a TagsItemContactsRemoveTags403Error error when the service returns a 403 status code
func (m *TagsItemContactsRemoveTagsRequestBuilder) Post(ctx context.Context, body TagsItemContactsRemoveTagsPostRequestBodyable, requestConfiguration *TagsItemContactsRemoveTagsRequestBuilderPostRequestConfiguration) error {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTagsItemContactsRemoveTags401ErrorFromDiscriminatorValue,
		"403": CreateTagsItemContactsRemoveTags403ErrorFromDiscriminatorValue,
	}
	err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
	if err != nil {
		return err
	}
	return nil
}

// ToPostRequestInformation removes a Tag from a list of Contacts.
// returns a *RequestInformation when successful
func (m *TagsItemContactsRemoveTagsRequestBuilder) ToPostRequestInformation(ctx context.Context, body TagsItemContactsRemoveTagsPostRequestBodyable, requestConfiguration *TagsItemContactsRemoveTagsRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TagsItemContactsRemoveTagsRequestBuilder when successful
func (m *TagsItemContactsRemoveTagsRequestBuilder) WithUrl(rawUrl string) *TagsItemContactsRemoveTagsRequestBuilder {
	return NewTagsItemContactsRemoveTagsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
