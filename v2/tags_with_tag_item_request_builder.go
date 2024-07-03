package v2

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// TagsWithTag_ItemRequestBuilder builds and executes requests for operations under \v2\tags\{tag_id}
type TagsWithTag_ItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// TagsWithTag_ItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TagsWithTag_ItemRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// TagsWithTag_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TagsWithTag_ItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// TagsWithTag_ItemRequestBuilderPatchQueryParameters updates a Tag with only the values provided in the request
type TagsWithTag_ItemRequestBuilderPatchQueryParameters struct {
	// An optional list of fields to be updated. If set, only the fields provided in the update_mask will be updated and others will be skipped.
	Update_mask []string `uriparametername:"update_mask"`
}

// TagsWithTag_ItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TagsWithTag_ItemRequestBuilderPatchRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *TagsWithTag_ItemRequestBuilderPatchQueryParameters
}

// Companies the companies property
// returns a *TagsItemCompaniesRequestBuilder when successful
func (m *TagsWithTag_ItemRequestBuilder) Companies() *TagsItemCompaniesRequestBuilder {
	return NewTagsItemCompaniesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// NewTagsWithTag_ItemRequestBuilderInternal instantiates a new TagsWithTag_ItemRequestBuilder and sets the default values.
func NewTagsWithTag_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TagsWithTag_ItemRequestBuilder {
	m := &TagsWithTag_ItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/tags/{tag_id}", pathParameters),
	}
	return m
}

// NewTagsWithTag_ItemRequestBuilder instantiates a new TagsWithTag_ItemRequestBuilder and sets the default values.
func NewTagsWithTag_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *TagsWithTag_ItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewTagsWithTag_ItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Contacts the contacts property
// returns a *TagsItemContactsRequestBuilder when successful
func (m *TagsWithTag_ItemRequestBuilder) Contacts() *TagsItemContactsRequestBuilder {
	return NewTagsItemContactsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// ContactsApplyTags the contactsApplyTags property
// returns a *TagsItemContactsApplyTagsRequestBuilder when successful
func (m *TagsWithTag_ItemRequestBuilder) ContactsApplyTags() *TagsItemContactsApplyTagsRequestBuilder {
	return NewTagsItemContactsApplyTagsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// ContactsRemoveTags the contactsRemoveTags property
// returns a *TagsItemContactsRemoveTagsRequestBuilder when successful
func (m *TagsWithTag_ItemRequestBuilder) ContactsRemoveTags() *TagsItemContactsRemoveTagsRequestBuilder {
	return NewTagsItemContactsRemoveTagsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Delete deletes a Tag.
// returns a TagsItemWithTag_401Error error when the service returns a 401 status code
// returns a TagsItemWithTag_403Error error when the service returns a 403 status code
// returns a TagsItemWithTag_404Error error when the service returns a 404 status code
func (m *TagsWithTag_ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TagsWithTag_ItemRequestBuilderDeleteRequestConfiguration) error {
	requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTagsItemWithTag_401ErrorFromDiscriminatorValue,
		"403": CreateTagsItemWithTag_403ErrorFromDiscriminatorValue,
		"404": CreateTagsItemWithTag_404ErrorFromDiscriminatorValue,
	}
	err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
	if err != nil {
		return err
	}
	return nil
}

// Get retrieves information about the specified Tag
// Deprecated: This method is obsolete. Use GetAsWithTag_GetResponse instead.
// returns a TagsItemWithTag_Responseable when successful
// returns a TagsItemWithTag_401Error error when the service returns a 401 status code
// returns a TagsItemWithTag_403Error error when the service returns a 403 status code
// returns a TagsItemWithTag_404Error error when the service returns a 404 status code
func (m *TagsWithTag_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TagsWithTag_ItemRequestBuilderGetRequestConfiguration) (TagsItemWithTag_Responseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTagsItemWithTag_401ErrorFromDiscriminatorValue,
		"403": CreateTagsItemWithTag_403ErrorFromDiscriminatorValue,
		"404": CreateTagsItemWithTag_404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTagsItemWithTag_ResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TagsItemWithTag_Responseable), nil
}

// GetAsWithTag_GetResponse retrieves information about the specified Tag
// returns a TagsItemWithTag_GetResponseable when successful
// returns a TagsItemWithTag_401Error error when the service returns a 401 status code
// returns a TagsItemWithTag_403Error error when the service returns a 403 status code
// returns a TagsItemWithTag_404Error error when the service returns a 404 status code
func (m *TagsWithTag_ItemRequestBuilder) GetAsWithTag_GetResponse(ctx context.Context, requestConfiguration *TagsWithTag_ItemRequestBuilderGetRequestConfiguration) (TagsItemWithTag_GetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTagsItemWithTag_401ErrorFromDiscriminatorValue,
		"403": CreateTagsItemWithTag_403ErrorFromDiscriminatorValue,
		"404": CreateTagsItemWithTag_404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTagsItemWithTag_GetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TagsItemWithTag_GetResponseable), nil
}

// Patch updates a Tag with only the values provided in the request
// Deprecated: This method is obsolete. Use PatchAsWithTag_PatchResponse instead.
// returns a TagsItemWithTag_Responseable when successful
// returns a TagsItemWithTag_401Error error when the service returns a 401 status code
// returns a TagsItemWithTag_403Error error when the service returns a 403 status code
// returns a TagsItemWithTag_404Error error when the service returns a 404 status code
func (m *TagsWithTag_ItemRequestBuilder) Patch(ctx context.Context, body TagsItemWithTag_PatchRequestBodyable, requestConfiguration *TagsWithTag_ItemRequestBuilderPatchRequestConfiguration) (TagsItemWithTag_Responseable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTagsItemWithTag_401ErrorFromDiscriminatorValue,
		"403": CreateTagsItemWithTag_403ErrorFromDiscriminatorValue,
		"404": CreateTagsItemWithTag_404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTagsItemWithTag_ResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TagsItemWithTag_Responseable), nil
}

// PatchAsWithTag_PatchResponse updates a Tag with only the values provided in the request
// returns a TagsItemWithTag_PatchResponseable when successful
// returns a TagsItemWithTag_401Error error when the service returns a 401 status code
// returns a TagsItemWithTag_403Error error when the service returns a 403 status code
// returns a TagsItemWithTag_404Error error when the service returns a 404 status code
func (m *TagsWithTag_ItemRequestBuilder) PatchAsWithTag_PatchResponse(ctx context.Context, body TagsItemWithTag_PatchRequestBodyable, requestConfiguration *TagsWithTag_ItemRequestBuilderPatchRequestConfiguration) (TagsItemWithTag_PatchResponseable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateTagsItemWithTag_401ErrorFromDiscriminatorValue,
		"403": CreateTagsItemWithTag_403ErrorFromDiscriminatorValue,
		"404": CreateTagsItemWithTag_404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateTagsItemWithTag_PatchResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(TagsItemWithTag_PatchResponseable), nil
}

// ToDeleteRequestInformation deletes a Tag.
// returns a *RequestInformation when successful
func (m *TagsWithTag_ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TagsWithTag_ItemRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// ToGetRequestInformation retrieves information about the specified Tag
// returns a *RequestInformation when successful
func (m *TagsWithTag_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TagsWithTag_ItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// ToPatchRequestInformation updates a Tag with only the values provided in the request
// returns a *RequestInformation when successful
func (m *TagsWithTag_ItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body TagsItemWithTag_PatchRequestBodyable, requestConfiguration *TagsWithTag_ItemRequestBuilderPatchRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/v2/tags/{tag_id}{?update_mask*}", m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		if requestConfiguration.QueryParameters != nil {
			requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
		}
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
// returns a *TagsWithTag_ItemRequestBuilder when successful
func (m *TagsWithTag_ItemRequestBuilder) WithUrl(rawUrl string) *TagsWithTag_ItemRequestBuilder {
	return NewTagsWithTag_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
