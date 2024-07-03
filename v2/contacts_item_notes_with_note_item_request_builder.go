package v2

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ContactsItemNotesWithNote_ItemRequestBuilder builds and executes requests for operations under \v2\contacts\{contact_id}\notes\{note_id}
type ContactsItemNotesWithNote_ItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ContactsItemNotesWithNote_ItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContactsItemNotesWithNote_ItemRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ContactsItemNotesWithNote_ItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContactsItemNotesWithNote_ItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ContactsItemNotesWithNote_ItemRequestBuilderPatchQueryParameters updates a Note for a Contact
type ContactsItemNotesWithNote_ItemRequestBuilderPatchQueryParameters struct {
	// An optional list of fields to be updated. If set, only the provided properties will be updated and others will be skipped.
	Update_mask []string `uriparametername:"update_mask"`
}

// ContactsItemNotesWithNote_ItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContactsItemNotesWithNote_ItemRequestBuilderPatchRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ContactsItemNotesWithNote_ItemRequestBuilderPatchQueryParameters
}

// NewContactsItemNotesWithNote_ItemRequestBuilderInternal instantiates a new ContactsItemNotesWithNote_ItemRequestBuilder and sets the default values.
func NewContactsItemNotesWithNote_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ContactsItemNotesWithNote_ItemRequestBuilder {
	m := &ContactsItemNotesWithNote_ItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/contacts/{contact_id}/notes/{note_id}", pathParameters),
	}
	return m
}

// NewContactsItemNotesWithNote_ItemRequestBuilder instantiates a new ContactsItemNotesWithNote_ItemRequestBuilder and sets the default values.
func NewContactsItemNotesWithNote_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ContactsItemNotesWithNote_ItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewContactsItemNotesWithNote_ItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete deletes the specified Note
// returns a ContactsItemNotesItemWithNote_401Error error when the service returns a 401 status code
// returns a ContactsItemNotesItemWithNote_403Error error when the service returns a 403 status code
// returns a ContactsItemNotesItemWithNote_404Error error when the service returns a 404 status code
func (m *ContactsItemNotesWithNote_ItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ContactsItemNotesWithNote_ItemRequestBuilderDeleteRequestConfiguration) error {
	requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsItemNotesItemWithNote_401ErrorFromDiscriminatorValue,
		"403": CreateContactsItemNotesItemWithNote_403ErrorFromDiscriminatorValue,
		"404": CreateContactsItemNotesItemWithNote_404ErrorFromDiscriminatorValue,
	}
	err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
	if err != nil {
		return err
	}
	return nil
}

// Get retrieves the specified Note
// Deprecated: This method is obsolete. Use GetAsWithNote_GetResponse instead.
// returns a ContactsItemNotesItemWithNote_Responseable when successful
// returns a ContactsItemNotesItemWithNote_401Error error when the service returns a 401 status code
// returns a ContactsItemNotesItemWithNote_403Error error when the service returns a 403 status code
// returns a ContactsItemNotesItemWithNote_404Error error when the service returns a 404 status code
func (m *ContactsItemNotesWithNote_ItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ContactsItemNotesWithNote_ItemRequestBuilderGetRequestConfiguration) (ContactsItemNotesItemWithNote_Responseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsItemNotesItemWithNote_401ErrorFromDiscriminatorValue,
		"403": CreateContactsItemNotesItemWithNote_403ErrorFromDiscriminatorValue,
		"404": CreateContactsItemNotesItemWithNote_404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsItemNotesItemWithNote_ResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsItemNotesItemWithNote_Responseable), nil
}

// GetAsWithNote_GetResponse retrieves the specified Note
// returns a ContactsItemNotesItemWithNote_GetResponseable when successful
// returns a ContactsItemNotesItemWithNote_401Error error when the service returns a 401 status code
// returns a ContactsItemNotesItemWithNote_403Error error when the service returns a 403 status code
// returns a ContactsItemNotesItemWithNote_404Error error when the service returns a 404 status code
func (m *ContactsItemNotesWithNote_ItemRequestBuilder) GetAsWithNote_GetResponse(ctx context.Context, requestConfiguration *ContactsItemNotesWithNote_ItemRequestBuilderGetRequestConfiguration) (ContactsItemNotesItemWithNote_GetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsItemNotesItemWithNote_401ErrorFromDiscriminatorValue,
		"403": CreateContactsItemNotesItemWithNote_403ErrorFromDiscriminatorValue,
		"404": CreateContactsItemNotesItemWithNote_404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsItemNotesItemWithNote_GetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsItemNotesItemWithNote_GetResponseable), nil
}

// Patch updates a Note for a Contact
// Deprecated: This method is obsolete. Use PatchAsWithNote_PatchResponse instead.
// returns a ContactsItemNotesItemWithNote_Responseable when successful
// returns a ContactsItemNotesItemWithNote_401Error error when the service returns a 401 status code
// returns a ContactsItemNotesItemWithNote_403Error error when the service returns a 403 status code
// returns a ContactsItemNotesItemWithNote_404Error error when the service returns a 404 status code
func (m *ContactsItemNotesWithNote_ItemRequestBuilder) Patch(ctx context.Context, body ContactsItemNotesItemWithNote_PatchRequestBodyable, requestConfiguration *ContactsItemNotesWithNote_ItemRequestBuilderPatchRequestConfiguration) (ContactsItemNotesItemWithNote_Responseable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsItemNotesItemWithNote_401ErrorFromDiscriminatorValue,
		"403": CreateContactsItemNotesItemWithNote_403ErrorFromDiscriminatorValue,
		"404": CreateContactsItemNotesItemWithNote_404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsItemNotesItemWithNote_ResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsItemNotesItemWithNote_Responseable), nil
}

// PatchAsWithNote_PatchResponse updates a Note for a Contact
// returns a ContactsItemNotesItemWithNote_PatchResponseable when successful
// returns a ContactsItemNotesItemWithNote_401Error error when the service returns a 401 status code
// returns a ContactsItemNotesItemWithNote_403Error error when the service returns a 403 status code
// returns a ContactsItemNotesItemWithNote_404Error error when the service returns a 404 status code
func (m *ContactsItemNotesWithNote_ItemRequestBuilder) PatchAsWithNote_PatchResponse(ctx context.Context, body ContactsItemNotesItemWithNote_PatchRequestBodyable, requestConfiguration *ContactsItemNotesWithNote_ItemRequestBuilderPatchRequestConfiguration) (ContactsItemNotesItemWithNote_PatchResponseable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsItemNotesItemWithNote_401ErrorFromDiscriminatorValue,
		"403": CreateContactsItemNotesItemWithNote_403ErrorFromDiscriminatorValue,
		"404": CreateContactsItemNotesItemWithNote_404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsItemNotesItemWithNote_PatchResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsItemNotesItemWithNote_PatchResponseable), nil
}

// ToDeleteRequestInformation deletes the specified Note
// returns a *RequestInformation when successful
func (m *ContactsItemNotesWithNote_ItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ContactsItemNotesWithNote_ItemRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// ToGetRequestInformation retrieves the specified Note
// returns a *RequestInformation when successful
func (m *ContactsItemNotesWithNote_ItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ContactsItemNotesWithNote_ItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// ToPatchRequestInformation updates a Note for a Contact
// returns a *RequestInformation when successful
func (m *ContactsItemNotesWithNote_ItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ContactsItemNotesItemWithNote_PatchRequestBodyable, requestConfiguration *ContactsItemNotesWithNote_ItemRequestBuilderPatchRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/v2/contacts/{contact_id}/notes/{note_id}{?update_mask*}", m.BaseRequestBuilder.PathParameters)
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
// returns a *ContactsItemNotesWithNote_ItemRequestBuilder when successful
func (m *ContactsItemNotesWithNote_ItemRequestBuilder) WithUrl(rawUrl string) *ContactsItemNotesWithNote_ItemRequestBuilder {
	return NewContactsItemNotesWithNote_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
