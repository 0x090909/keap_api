package v2

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// TagsItemContactsRequestBuilder builds and executes requests for operations under \v2\tags\{tag_id}\contacts
type TagsItemContactsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// TagsItemContactsRequestBuilderGetQueryParameters retrieve a list of Contacts with the specified Tag.To search for `null` or empty fields use `filter==NONE`
type TagsItemContactsRequestBuilderGetQueryParameters struct {
	// Search filter to apply to results
	Filter *string `uriparametername:"filter"`
	// Attribute and direction to order items by. E.g. `given_name desc`
	Order_by *string `uriparametername:"order_by"`
	// Total number of items to return per page
	Page_size *int32 `uriparametername:"page_size"`
	// Page token
	Page_token *string `uriparametername:"page_token"`
}

// TagsItemContactsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TagsItemContactsRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *TagsItemContactsRequestBuilderGetQueryParameters
}

// NewTagsItemContactsRequestBuilderInternal instantiates a new TagsItemContactsRequestBuilder and sets the default values.
func NewTagsItemContactsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TagsItemContactsRequestBuilder {
	m := &TagsItemContactsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/tags/{tag_id}/contacts{?filter*,order_by*,page_size*,page_token*}", pathParameters),
	}
	return m
}

// NewTagsItemContactsRequestBuilder instantiates a new TagsItemContactsRequestBuilder and sets the default values.
func NewTagsItemContactsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TagsItemContactsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewTagsItemContactsRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieve a list of Contacts with the specified Tag.To search for `null` or empty fields use `filter==NONE`
// Deprecated: This method is obsolete. Use GetAsContactsGetResponse instead.
// returns a TagsItemContactsResponseable when successful
// returns a TagsItemContacts401Error error when the service returns a 401 status code
// returns a TagsItemContacts403Error error when the service returns a 403 status code
// returns a TagsItemContacts404Error error when the service returns a 404 status code
func (m *TagsItemContactsRequestBuilder) Get(ctx context.Context, requestConfiguration *TagsItemContactsRequestBuilderGetRequestConfiguration) (TagsItemContactsResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTagsItemContacts401ErrorFromDiscriminatorValue,
		"403": CreateTagsItemContacts403ErrorFromDiscriminatorValue,
		"404": CreateTagsItemContacts404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTagsItemContactsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TagsItemContactsResponseable), nil
}

// GetAsContactsGetResponse retrieve a list of Contacts with the specified Tag.To search for `null` or empty fields use `filter==NONE`
// returns a TagsItemContactsGetResponseable when successful
// returns a TagsItemContacts401Error error when the service returns a 401 status code
// returns a TagsItemContacts403Error error when the service returns a 403 status code
// returns a TagsItemContacts404Error error when the service returns a 404 status code
func (m *TagsItemContactsRequestBuilder) GetAsContactsGetResponse(ctx context.Context, requestConfiguration *TagsItemContactsRequestBuilderGetRequestConfiguration) (TagsItemContactsGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTagsItemContacts401ErrorFromDiscriminatorValue,
		"403": CreateTagsItemContacts403ErrorFromDiscriminatorValue,
		"404": CreateTagsItemContacts404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTagsItemContactsGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TagsItemContactsGetResponseable), nil
}

// ToGetRequestInformation retrieve a list of Contacts with the specified Tag.To search for `null` or empty fields use `filter==NONE`
// returns a *RequestInformation when successful
func (m *TagsItemContactsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TagsItemContactsRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TagsItemContactsRequestBuilder when successful
func (m *TagsItemContactsRequestBuilder) WithUrl(rawUrl string) *TagsItemContactsRequestBuilder {
	return NewTagsItemContactsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
