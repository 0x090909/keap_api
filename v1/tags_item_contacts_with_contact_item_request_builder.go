package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// TagsItemContactsWithContactItemRequestBuilder builds and executes requests for operations under \v1\tags\{id-id}\contacts\{contactId}
type TagsItemContactsWithContactItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// TagsItemContactsWithContactItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TagsItemContactsWithContactItemRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewTagsItemContactsWithContactItemRequestBuilderInternal instantiates a new TagsItemContactsWithContactItemRequestBuilder and sets the default values.
func NewTagsItemContactsWithContactItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TagsItemContactsWithContactItemRequestBuilder {
	m := &TagsItemContactsWithContactItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/tags/{id%2Did}/contacts/{contactId}", pathParameters),
	}
	return m
}

// NewTagsItemContactsWithContactItemRequestBuilder instantiates a new TagsItemContactsWithContactItemRequestBuilder and sets the default values.
func NewTagsItemContactsWithContactItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TagsItemContactsWithContactItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewTagsItemContactsWithContactItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete remove a tag from a Contact
// returns a TagsItemContactsItemWithContact401Error error when the service returns a 401 status code
// returns a TagsItemContactsItemWithContact403Error error when the service returns a 403 status code
// returns a TagsItemContactsItemWithContact404Error error when the service returns a 404 status code
func (m *TagsItemContactsWithContactItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TagsItemContactsWithContactItemRequestBuilderDeleteRequestConfiguration) error {
	requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTagsItemContactsItemWithContact401ErrorFromDiscriminatorValue,
		"403": CreateTagsItemContactsItemWithContact403ErrorFromDiscriminatorValue,
		"404": CreateTagsItemContactsItemWithContact404ErrorFromDiscriminatorValue,
	}
	err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
	if err != nil {
		return err
	}
	return nil
}

// ToDeleteRequestInformation remove a tag from a Contact
// returns a *RequestInformation when successful
func (m *TagsItemContactsWithContactItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TagsItemContactsWithContactItemRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TagsItemContactsWithContactItemRequestBuilder when successful
func (m *TagsItemContactsWithContactItemRequestBuilder) WithUrl(rawUrl string) *TagsItemContactsWithContactItemRequestBuilder {
	return NewTagsItemContactsWithContactItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
