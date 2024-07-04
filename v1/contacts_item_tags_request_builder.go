package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
)

// ContactsItemTagsRequestBuilder builds and executes requests for operations under \v1\contacts\{contactId-id}\tags
type ContactsItemTagsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ContactsItemTagsRequestBuilderDeleteQueryParameters removes a list of tags from the given contact. Provide one or more tag ids in the querystring as a comma-separated URIencoded list (%2C is a comma). E.g. DELETE /contacts/{contact_id}/tags?ids=1%2C2%2C3
type ContactsItemTagsRequestBuilderDeleteQueryParameters struct {
	// ids
	Ids *string `uriparametername:"ids"`
}

// ContactsItemTagsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContactsItemTagsRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ContactsItemTagsRequestBuilderDeleteQueryParameters
}

// ContactsItemTagsRequestBuilderGetQueryParameters retrieves a list of tags applied to a given contact
type ContactsItemTagsRequestBuilderGetQueryParameters struct {
	// Sets a total of items to return
	Limit *int32 `uriparametername:"limit"`
	// Sets a beginning range of items to return
	Offset *int32 `uriparametername:"offset"`
}

// ContactsItemTagsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContactsItemTagsRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ContactsItemTagsRequestBuilderGetQueryParameters
}

// ContactsItemTagsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContactsItemTagsRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ByTagId gets an item from the github.com/0x090909/keap_api.v1.contacts.item.tags.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *ContactsItemTagsWithTagItemRequestBuilder when successful
func (m *ContactsItemTagsRequestBuilder) ByTagId(tagId string) *ContactsItemTagsWithTagItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if tagId != "" {
		urlTplParams["tagId"] = tagId
	}
	return NewContactsItemTagsWithTagItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// ByTagIdInt64 gets an item from the github.com/0x090909/keap_api.v1.contacts.item.tags.item collection
// returns a *ContactsItemTagsWithTagItemRequestBuilder when successful
func (m *ContactsItemTagsRequestBuilder) ByTagIdInt64(tagId int64) *ContactsItemTagsWithTagItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	urlTplParams["tagId"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(tagId, 10)
	return NewContactsItemTagsWithTagItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewContactsItemTagsRequestBuilderInternal instantiates a new ContactsItemTagsRequestBuilder and sets the default values.
func NewContactsItemTagsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ContactsItemTagsRequestBuilder {
	m := &ContactsItemTagsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/contacts/{contactId%2Did}/tags{?limit*,offset*}", pathParameters),
	}
	return m
}

// NewContactsItemTagsRequestBuilder instantiates a new ContactsItemTagsRequestBuilder and sets the default values.
func NewContactsItemTagsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ContactsItemTagsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewContactsItemTagsRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete removes a list of tags from the given contact. Provide one or more tag ids in the querystring as a comma-separated URIencoded list (%2C is a comma). E.g. DELETE /contacts/{contact_id}/tags?ids=1%2C2%2C3
// returns a ContactsItemTags401Error error when the service returns a 401 status code
// returns a ContactsItemTags403Error error when the service returns a 403 status code
// returns a ContactsItemTags404Error error when the service returns a 404 status code
func (m *ContactsItemTagsRequestBuilder) Delete(ctx context.Context, requestConfiguration *ContactsItemTagsRequestBuilderDeleteRequestConfiguration) error {
	requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsItemTags401ErrorFromDiscriminatorValue,
		"403": CreateContactsItemTags403ErrorFromDiscriminatorValue,
		"404": CreateContactsItemTags404ErrorFromDiscriminatorValue,
	}
	err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
	if err != nil {
		return err
	}
	return nil
}

// Get retrieves a list of tags applied to a given contact
// Deprecated: This method is obsolete. Use GetAsTagsGetResponse instead.
// returns a ContactsItemTagsResponseable when successful
// returns a ContactsItemTags401Error error when the service returns a 401 status code
// returns a ContactsItemTags403Error error when the service returns a 403 status code
// returns a ContactsItemTags404Error error when the service returns a 404 status code
func (m *ContactsItemTagsRequestBuilder) Get(ctx context.Context, requestConfiguration *ContactsItemTagsRequestBuilderGetRequestConfiguration) (ContactsItemTagsResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsItemTags401ErrorFromDiscriminatorValue,
		"403": CreateContactsItemTags403ErrorFromDiscriminatorValue,
		"404": CreateContactsItemTags404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsItemTagsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsItemTagsResponseable), nil
}

// GetAsTagsGetResponse retrieves a list of tags applied to a given contact
// returns a ContactsItemTagsGetResponseable when successful
// returns a ContactsItemTags401Error error when the service returns a 401 status code
// returns a ContactsItemTags403Error error when the service returns a 403 status code
// returns a ContactsItemTags404Error error when the service returns a 404 status code
func (m *ContactsItemTagsRequestBuilder) GetAsTagsGetResponse(ctx context.Context, requestConfiguration *ContactsItemTagsRequestBuilderGetRequestConfiguration) (ContactsItemTagsGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsItemTags401ErrorFromDiscriminatorValue,
		"403": CreateContactsItemTags403ErrorFromDiscriminatorValue,
		"404": CreateContactsItemTags404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsItemTagsGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsItemTagsGetResponseable), nil
}

// Post apply a list of tags to a given contact record
// returns a []ContactsItemTagsable when successful
// returns a ContactsItemTags401Error error when the service returns a 401 status code
// returns a ContactsItemTags403Error error when the service returns a 403 status code
func (m *ContactsItemTagsRequestBuilder) Post(ctx context.Context, body ContactsItemTagsPostRequestBodyable, requestConfiguration *ContactsItemTagsRequestBuilderPostRequestConfiguration) ([]ContactsItemTagsable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsItemTags401ErrorFromDiscriminatorValue,
		"403": CreateContactsItemTags403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, CreateContactsItemTagsFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	val := make([]ContactsItemTagsable, len(res))
	for i, v := range res {
		if v != nil {
			val[i] = v.(ContactsItemTagsable)
		}
	}
	return val, nil
}

// ToDeleteRequestInformation removes a list of tags from the given contact. Provide one or more tag ids in the querystring as a comma-separated URIencoded list (%2C is a comma). E.g. DELETE /contacts/{contact_id}/tags?ids=1%2C2%2C3
// returns a *RequestInformation when successful
func (m *ContactsItemTagsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ContactsItemTagsRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/v1/contacts/{contactId%2Did}/tags?ids={ids}", m.BaseRequestBuilder.PathParameters)
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

// ToGetRequestInformation retrieves a list of tags applied to a given contact
// returns a *RequestInformation when successful
func (m *ContactsItemTagsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ContactsItemTagsRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToPostRequestInformation apply a list of tags to a given contact record
// returns a *RequestInformation when successful
func (m *ContactsItemTagsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ContactsItemTagsPostRequestBodyable, requestConfiguration *ContactsItemTagsRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/v1/contacts/{contactId%2Did}/tags", m.BaseRequestBuilder.PathParameters)
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
// returns a *ContactsItemTagsRequestBuilder when successful
func (m *ContactsItemTagsRequestBuilder) WithUrl(rawUrl string) *ContactsItemTagsRequestBuilder {
	return NewContactsItemTagsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
