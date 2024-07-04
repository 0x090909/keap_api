package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
)

// NotesRequestBuilder builds and executes requests for operations under \v1\notes
type NotesRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// NotesRequestBuilderGetQueryParameters retrieves a list of all notes
type NotesRequestBuilderGetQueryParameters struct {
	// Filter based on the contact id assigned to the note.
	Contact_id *int64 `uriparametername:"contact_id"`
	// Sets a total of items to return
	Limit *int32 `uriparametername:"limit"`
	// Sets a beginning range of items to return
	Offset *int32 `uriparametername:"offset"`
	// Filter based on the user id assigned to the note.
	User_id *int64 `uriparametername:"user_id"`
}

// NotesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type NotesRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *NotesRequestBuilderGetQueryParameters
}

// NotesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type NotesRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ByNoteId gets an item from the github.com/0x090909/keap_api.v1.notes.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *NotesWithNoteItemRequestBuilder when successful
func (m *NotesRequestBuilder) ByNoteId(noteId string) *NotesWithNoteItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if noteId != "" {
		urlTplParams["noteId"] = noteId
	}
	return NewNotesWithNoteItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// ByNoteIdInt64 gets an item from the github.com/0x090909/keap_api.v1.notes.item collection
// returns a *NotesWithNoteItemRequestBuilder when successful
func (m *NotesRequestBuilder) ByNoteIdInt64(noteId int64) *NotesWithNoteItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	urlTplParams["noteId"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(noteId, 10)
	return NewNotesWithNoteItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewNotesRequestBuilderInternal instantiates a new NotesRequestBuilder and sets the default values.
func NewNotesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *NotesRequestBuilder {
	m := &NotesRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/notes{?contact_id*,limit*,offset*,user_id*}", pathParameters),
	}
	return m
}

// NewNotesRequestBuilder instantiates a new NotesRequestBuilder and sets the default values.
func NewNotesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *NotesRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewNotesRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieves a list of all notes
// Deprecated: This method is obsolete. Use GetAsNotesGetResponse instead.
// returns a NotesResponseable when successful
// returns a Notes401Error error when the service returns a 401 status code
// returns a Notes403Error error when the service returns a 403 status code
// returns a Notes404Error error when the service returns a 404 status code
func (m *NotesRequestBuilder) Get(ctx context.Context, requestConfiguration *NotesRequestBuilderGetRequestConfiguration) (NotesResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateNotes401ErrorFromDiscriminatorValue,
		"403": CreateNotes403ErrorFromDiscriminatorValue,
		"404": CreateNotes404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateNotesResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(NotesResponseable), nil
}

// GetAsNotesGetResponse retrieves a list of all notes
// returns a NotesGetResponseable when successful
// returns a Notes401Error error when the service returns a 401 status code
// returns a Notes403Error error when the service returns a 403 status code
// returns a Notes404Error error when the service returns a 404 status code
func (m *NotesRequestBuilder) GetAsNotesGetResponse(ctx context.Context, requestConfiguration *NotesRequestBuilderGetRequestConfiguration) (NotesGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateNotes401ErrorFromDiscriminatorValue,
		"403": CreateNotes403ErrorFromDiscriminatorValue,
		"404": CreateNotes404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateNotesGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(NotesGetResponseable), nil
}

// Model the model property
// returns a *NotesModelRequestBuilder when successful
func (m *NotesRequestBuilder) Model() *NotesModelRequestBuilder {
	return NewNotesModelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Post creates a new note as the authenticated user. Either a "title" or "body" is required
// Deprecated: This method is obsolete. Use PostAsNotesPostResponse instead.
// returns a NotesResponseable when successful
// returns a Notes401Error error when the service returns a 401 status code
// returns a Notes403Error error when the service returns a 403 status code
func (m *NotesRequestBuilder) Post(ctx context.Context, body NotesPostRequestBodyable, requestConfiguration *NotesRequestBuilderPostRequestConfiguration) (NotesResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateNotes401ErrorFromDiscriminatorValue,
		"403": CreateNotes403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateNotesResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(NotesResponseable), nil
}

// PostAsNotesPostResponse creates a new note as the authenticated user. Either a "title" or "body" is required
// returns a NotesPostResponseable when successful
// returns a Notes401Error error when the service returns a 401 status code
// returns a Notes403Error error when the service returns a 403 status code
func (m *NotesRequestBuilder) PostAsNotesPostResponse(ctx context.Context, body NotesPostRequestBodyable, requestConfiguration *NotesRequestBuilderPostRequestConfiguration) (NotesPostResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateNotes401ErrorFromDiscriminatorValue,
		"403": CreateNotes403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateNotesPostResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(NotesPostResponseable), nil
}

// ToGetRequestInformation retrieves a list of all notes
// returns a *RequestInformation when successful
func (m *NotesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *NotesRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToPostRequestInformation creates a new note as the authenticated user. Either a "title" or "body" is required
// returns a *RequestInformation when successful
func (m *NotesRequestBuilder) ToPostRequestInformation(ctx context.Context, body NotesPostRequestBodyable, requestConfiguration *NotesRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/v1/notes", m.BaseRequestBuilder.PathParameters)
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
// returns a *NotesRequestBuilder when successful
func (m *NotesRequestBuilder) WithUrl(rawUrl string) *NotesRequestBuilder {
	return NewNotesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
