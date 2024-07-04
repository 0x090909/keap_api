package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
)

// TagsItemContactsRequestBuilder builds and executes requests for operations under \v1\tags\{id-id}\contacts
type TagsItemContactsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// TagsItemContactsRequestBuilderDeleteQueryParameters remove a tag from a list of contacts
type TagsItemContactsRequestBuilderDeleteQueryParameters struct {
	// ids
	Ids []int64 `uriparametername:"ids"`
}

// TagsItemContactsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TagsItemContactsRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *TagsItemContactsRequestBuilderDeleteQueryParameters
}

// TagsItemContactsRequestBuilderGetQueryParameters retrieves a list of contacts that have the given tag applied
type TagsItemContactsRequestBuilderGetQueryParameters struct {
	// Sets a total of items to return
	Limit *int32 `uriparametername:"limit"`
	// Sets a beginning range of items to return
	Offset *int32 `uriparametername:"offset"`
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

// TagsItemContactsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TagsItemContactsRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ByContactId gets an item from the github.com/0x090909/keap_api.v1.tags.item.contacts.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *TagsItemContactsWithContactItemRequestBuilder when successful
func (m *TagsItemContactsRequestBuilder) ByContactId(contactId string) *TagsItemContactsWithContactItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if contactId != "" {
		urlTplParams["contactId"] = contactId
	}
	return NewTagsItemContactsWithContactItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// ByContactIdInt64 gets an item from the github.com/0x090909/keap_api.v1.tags.item.contacts.item collection
// returns a *TagsItemContactsWithContactItemRequestBuilder when successful
func (m *TagsItemContactsRequestBuilder) ByContactIdInt64(contactId int64) *TagsItemContactsWithContactItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	urlTplParams["contactId"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(contactId, 10)
	return NewTagsItemContactsWithContactItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewTagsItemContactsRequestBuilderInternal instantiates a new TagsItemContactsRequestBuilder and sets the default values.
func NewTagsItemContactsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TagsItemContactsRequestBuilder {
	m := &TagsItemContactsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/tags/{id%2Did}/contacts{?limit*,offset*}", pathParameters),
	}
	return m
}

// NewTagsItemContactsRequestBuilder instantiates a new TagsItemContactsRequestBuilder and sets the default values.
func NewTagsItemContactsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TagsItemContactsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewTagsItemContactsRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete remove a tag from a list of contacts
// returns a TagsItemContacts401Error error when the service returns a 401 status code
// returns a TagsItemContacts403Error error when the service returns a 403 status code
// returns a TagsItemContacts404Error error when the service returns a 404 status code
func (m *TagsItemContactsRequestBuilder) Delete(ctx context.Context, requestConfiguration *TagsItemContactsRequestBuilderDeleteRequestConfiguration) error {
	requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTagsItemContacts401ErrorFromDiscriminatorValue,
		"403": CreateTagsItemContacts403ErrorFromDiscriminatorValue,
		"404": CreateTagsItemContacts404ErrorFromDiscriminatorValue,
	}
	err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
	if err != nil {
		return err
	}
	return nil
}

// Get retrieves a list of contacts that have the given tag applied
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

// GetAsContactsGetResponse retrieves a list of contacts that have the given tag applied
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

// Post apply a tag to a list of contacts
// returns a []TagsItemContactsable when successful
// returns a TagsItemContacts401Error error when the service returns a 401 status code
// returns a TagsItemContacts403Error error when the service returns a 403 status code
func (m *TagsItemContactsRequestBuilder) Post(ctx context.Context, body TagsItemContactsPostRequestBodyable, requestConfiguration *TagsItemContactsRequestBuilderPostRequestConfiguration) ([]TagsItemContactsable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTagsItemContacts401ErrorFromDiscriminatorValue,
		"403": CreateTagsItemContacts403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, CreateTagsItemContactsFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	val := make([]TagsItemContactsable, len(res))
	for i, v := range res {
		if v != nil {
			val[i] = v.(TagsItemContactsable)
		}
	}
	return val, nil
}

// ToDeleteRequestInformation remove a tag from a list of contacts
// returns a *RequestInformation when successful
func (m *TagsItemContactsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TagsItemContactsRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/v1/tags/{id%2Did}/contacts?ids={ids}", m.BaseRequestBuilder.PathParameters)
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

// ToGetRequestInformation retrieves a list of contacts that have the given tag applied
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

// ToPostRequestInformation apply a tag to a list of contacts
// returns a *RequestInformation when successful
func (m *TagsItemContactsRequestBuilder) ToPostRequestInformation(ctx context.Context, body TagsItemContactsPostRequestBodyable, requestConfiguration *TagsItemContactsRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/v1/tags/{id%2Did}/contacts", m.BaseRequestBuilder.PathParameters)
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
// returns a *TagsItemContactsRequestBuilder when successful
func (m *TagsItemContactsRequestBuilder) WithUrl(rawUrl string) *TagsItemContactsRequestBuilder {
	return NewTagsItemContactsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
