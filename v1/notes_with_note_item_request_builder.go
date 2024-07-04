package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// NotesWithNoteItemRequestBuilder builds and executes requests for operations under \v1\notes\{noteId}
type NotesWithNoteItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// NotesWithNoteItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type NotesWithNoteItemRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NotesWithNoteItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type NotesWithNoteItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NotesWithNoteItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type NotesWithNoteItemRequestBuilderPatchRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NotesWithNoteItemRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type NotesWithNoteItemRequestBuilderPutRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewNotesWithNoteItemRequestBuilderInternal instantiates a new NotesWithNoteItemRequestBuilder and sets the default values.
func NewNotesWithNoteItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *NotesWithNoteItemRequestBuilder {
	m := &NotesWithNoteItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/notes/{noteId}", pathParameters),
	}
	return m
}

// NewNotesWithNoteItemRequestBuilder instantiates a new NotesWithNoteItemRequestBuilder and sets the default values.
func NewNotesWithNoteItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *NotesWithNoteItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewNotesWithNoteItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete deletes a note
// returns a NotesItemWithNote401Error error when the service returns a 401 status code
// returns a NotesItemWithNote403Error error when the service returns a 403 status code
// returns a NotesItemWithNote404Error error when the service returns a 404 status code
func (m *NotesWithNoteItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *NotesWithNoteItemRequestBuilderDeleteRequestConfiguration) error {
	requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateNotesItemWithNote401ErrorFromDiscriminatorValue,
		"403": CreateNotesItemWithNote403ErrorFromDiscriminatorValue,
		"404": CreateNotesItemWithNote404ErrorFromDiscriminatorValue,
	}
	err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
	if err != nil {
		return err
	}
	return nil
}

// Get retrieves a single note
// Deprecated: This method is obsolete. Use GetAsWithNoteGetResponse instead.
// returns a NotesItemWithNoteResponseable when successful
// returns a NotesItemWithNote401Error error when the service returns a 401 status code
// returns a NotesItemWithNote403Error error when the service returns a 403 status code
// returns a NotesItemWithNote404Error error when the service returns a 404 status code
func (m *NotesWithNoteItemRequestBuilder) Get(ctx context.Context, requestConfiguration *NotesWithNoteItemRequestBuilderGetRequestConfiguration) (NotesItemWithNoteResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateNotesItemWithNote401ErrorFromDiscriminatorValue,
		"403": CreateNotesItemWithNote403ErrorFromDiscriminatorValue,
		"404": CreateNotesItemWithNote404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateNotesItemWithNoteResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(NotesItemWithNoteResponseable), nil
}

// GetAsWithNoteGetResponse retrieves a single note
// returns a NotesItemWithNoteGetResponseable when successful
// returns a NotesItemWithNote401Error error when the service returns a 401 status code
// returns a NotesItemWithNote403Error error when the service returns a 403 status code
// returns a NotesItemWithNote404Error error when the service returns a 404 status code
func (m *NotesWithNoteItemRequestBuilder) GetAsWithNoteGetResponse(ctx context.Context, requestConfiguration *NotesWithNoteItemRequestBuilderGetRequestConfiguration) (NotesItemWithNoteGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateNotesItemWithNote401ErrorFromDiscriminatorValue,
		"403": CreateNotesItemWithNote403ErrorFromDiscriminatorValue,
		"404": CreateNotesItemWithNote404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateNotesItemWithNoteGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(NotesItemWithNoteGetResponseable), nil
}

// Patch updates a note with only the values provided in the request
// Deprecated: This method is obsolete. Use PatchAsWithNotePatchResponse instead.
// returns a NotesItemWithNoteResponseable when successful
// returns a NotesItemWithNote401Error error when the service returns a 401 status code
// returns a NotesItemWithNote403Error error when the service returns a 403 status code
// returns a NotesItemWithNote404Error error when the service returns a 404 status code
func (m *NotesWithNoteItemRequestBuilder) Patch(ctx context.Context, body NotesItemWithNotePatchRequestBodyable, requestConfiguration *NotesWithNoteItemRequestBuilderPatchRequestConfiguration) (NotesItemWithNoteResponseable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateNotesItemWithNote401ErrorFromDiscriminatorValue,
		"403": CreateNotesItemWithNote403ErrorFromDiscriminatorValue,
		"404": CreateNotesItemWithNote404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateNotesItemWithNoteResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(NotesItemWithNoteResponseable), nil
}

// PatchAsWithNotePatchResponse updates a note with only the values provided in the request
// returns a NotesItemWithNotePatchResponseable when successful
// returns a NotesItemWithNote401Error error when the service returns a 401 status code
// returns a NotesItemWithNote403Error error when the service returns a 403 status code
// returns a NotesItemWithNote404Error error when the service returns a 404 status code
func (m *NotesWithNoteItemRequestBuilder) PatchAsWithNotePatchResponse(ctx context.Context, body NotesItemWithNotePatchRequestBodyable, requestConfiguration *NotesWithNoteItemRequestBuilderPatchRequestConfiguration) (NotesItemWithNotePatchResponseable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateNotesItemWithNote401ErrorFromDiscriminatorValue,
		"403": CreateNotesItemWithNote403ErrorFromDiscriminatorValue,
		"404": CreateNotesItemWithNote404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateNotesItemWithNotePatchResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(NotesItemWithNotePatchResponseable), nil
}

// Put replaces all values of a given note
// Deprecated: This method is obsolete. Use PutAsWithNotePutResponse instead.
// returns a NotesItemWithNoteResponseable when successful
// returns a NotesItemWithNote401Error error when the service returns a 401 status code
// returns a NotesItemWithNote403Error error when the service returns a 403 status code
// returns a NotesItemWithNote404Error error when the service returns a 404 status code
func (m *NotesWithNoteItemRequestBuilder) Put(ctx context.Context, body NotesItemWithNotePutRequestBodyable, requestConfiguration *NotesWithNoteItemRequestBuilderPutRequestConfiguration) (NotesItemWithNoteResponseable, error) {
	requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateNotesItemWithNote401ErrorFromDiscriminatorValue,
		"403": CreateNotesItemWithNote403ErrorFromDiscriminatorValue,
		"404": CreateNotesItemWithNote404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateNotesItemWithNoteResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(NotesItemWithNoteResponseable), nil
}

// PutAsWithNotePutResponse replaces all values of a given note
// returns a NotesItemWithNotePutResponseable when successful
// returns a NotesItemWithNote401Error error when the service returns a 401 status code
// returns a NotesItemWithNote403Error error when the service returns a 403 status code
// returns a NotesItemWithNote404Error error when the service returns a 404 status code
func (m *NotesWithNoteItemRequestBuilder) PutAsWithNotePutResponse(ctx context.Context, body NotesItemWithNotePutRequestBodyable, requestConfiguration *NotesWithNoteItemRequestBuilderPutRequestConfiguration) (NotesItemWithNotePutResponseable, error) {
	requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateNotesItemWithNote401ErrorFromDiscriminatorValue,
		"403": CreateNotesItemWithNote403ErrorFromDiscriminatorValue,
		"404": CreateNotesItemWithNote404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateNotesItemWithNotePutResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(NotesItemWithNotePutResponseable), nil
}

// ToDeleteRequestInformation deletes a note
// returns a *RequestInformation when successful
func (m *NotesWithNoteItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *NotesWithNoteItemRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// ToGetRequestInformation retrieves a single note
// returns a *RequestInformation when successful
func (m *NotesWithNoteItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *NotesWithNoteItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// ToPatchRequestInformation updates a note with only the values provided in the request
// returns a *RequestInformation when successful
func (m *NotesWithNoteItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body NotesItemWithNotePatchRequestBodyable, requestConfiguration *NotesWithNoteItemRequestBuilderPatchRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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

// ToPutRequestInformation replaces all values of a given note
// returns a *RequestInformation when successful
func (m *NotesWithNoteItemRequestBuilder) ToPutRequestInformation(ctx context.Context, body NotesItemWithNotePutRequestBodyable, requestConfiguration *NotesWithNoteItemRequestBuilderPutRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *NotesWithNoteItemRequestBuilder when successful
func (m *NotesWithNoteItemRequestBuilder) WithUrl(rawUrl string) *NotesWithNoteItemRequestBuilder {
	return NewNotesWithNoteItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
