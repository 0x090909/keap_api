package v2

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ContactsItemNotesRequestBuilder builds and executes requests for operations under \v2\contacts\{contact_id}\notes
type ContactsItemNotesRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ContactsItemNotesRequestBuilderGetQueryParameters retrieves a list of Notes
type ContactsItemNotesRequestBuilderGetQueryParameters struct {
	// Search filter to apply to results
	Filter *string `uriparametername:"filter"`
	// Attribute and direction to order items by. E.g. `given_name desc`
	Order_by *string `uriparametername:"order_by"`
	// Total number of items to return per page
	Page_size *int32 `uriparametername:"page_size"`
	// Page token
	Page_token *string `uriparametername:"page_token"`
}

// ContactsItemNotesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContactsItemNotesRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *ContactsItemNotesRequestBuilderGetQueryParameters
}

// ContactsItemNotesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContactsItemNotesRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ByNote_id gets an item from the keapapi.v2.contacts.item.notes.item collection
// returns a *ContactsItemNotesWithNote_ItemRequestBuilder when successful
func (m *ContactsItemNotesRequestBuilder) ByNote_id(note_id string) *ContactsItemNotesWithNote_ItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if note_id != "" {
		urlTplParams["note_id"] = note_id
	}
	return NewContactsItemNotesWithNote_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewContactsItemNotesRequestBuilderInternal instantiates a new ContactsItemNotesRequestBuilder and sets the default values.
func NewContactsItemNotesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ContactsItemNotesRequestBuilder {
	m := &ContactsItemNotesRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/contacts/{contact_id}/notes{?filter*,order_by*,page_size*,page_token*}", pathParameters),
	}
	return m
}

// NewContactsItemNotesRequestBuilder instantiates a new ContactsItemNotesRequestBuilder and sets the default values.
func NewContactsItemNotesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ContactsItemNotesRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewContactsItemNotesRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieves a list of Notes
// Deprecated: This method is obsolete. Use GetAsNotesGetResponse instead.
// returns a ContactsItemNotesResponseable when successful
// returns a ContactsItemNotes401Error error when the service returns a 401 status code
// returns a ContactsItemNotes403Error error when the service returns a 403 status code
// returns a ContactsItemNotes404Error error when the service returns a 404 status code
func (m *ContactsItemNotesRequestBuilder) Get(ctx context.Context, requestConfiguration *ContactsItemNotesRequestBuilderGetRequestConfiguration) (ContactsItemNotesResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsItemNotes401ErrorFromDiscriminatorValue,
		"403": CreateContactsItemNotes403ErrorFromDiscriminatorValue,
		"404": CreateContactsItemNotes404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsItemNotesResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsItemNotesResponseable), nil
}

// GetAsNotesGetResponse retrieves a list of Notes
// returns a ContactsItemNotesGetResponseable when successful
// returns a ContactsItemNotes401Error error when the service returns a 401 status code
// returns a ContactsItemNotes403Error error when the service returns a 403 status code
// returns a ContactsItemNotes404Error error when the service returns a 404 status code
func (m *ContactsItemNotesRequestBuilder) GetAsNotesGetResponse(ctx context.Context, requestConfiguration *ContactsItemNotesRequestBuilderGetRequestConfiguration) (ContactsItemNotesGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsItemNotes401ErrorFromDiscriminatorValue,
		"403": CreateContactsItemNotes403ErrorFromDiscriminatorValue,
		"404": CreateContactsItemNotes404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsItemNotesGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsItemNotesGetResponseable), nil
}

// Post creates a new Note.
// Deprecated: This method is obsolete. Use PostAsNotesPostResponse instead.
// returns a ContactsItemNotesResponseable when successful
// returns a ContactsItemNotes401Error error when the service returns a 401 status code
// returns a ContactsItemNotes403Error error when the service returns a 403 status code
func (m *ContactsItemNotesRequestBuilder) Post(ctx context.Context, body ContactsItemNotesPostRequestBodyable, requestConfiguration *ContactsItemNotesRequestBuilderPostRequestConfiguration) (ContactsItemNotesResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsItemNotes401ErrorFromDiscriminatorValue,
		"403": CreateContactsItemNotes403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsItemNotesResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsItemNotesResponseable), nil
}

// PostAsNotesPostResponse creates a new Note.
// returns a ContactsItemNotesPostResponseable when successful
// returns a ContactsItemNotes401Error error when the service returns a 401 status code
// returns a ContactsItemNotes403Error error when the service returns a 403 status code
func (m *ContactsItemNotesRequestBuilder) PostAsNotesPostResponse(ctx context.Context, body ContactsItemNotesPostRequestBodyable, requestConfiguration *ContactsItemNotesRequestBuilderPostRequestConfiguration) (ContactsItemNotesPostResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsItemNotes401ErrorFromDiscriminatorValue,
		"403": CreateContactsItemNotes403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsItemNotesPostResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsItemNotesPostResponseable), nil
}

// ToGetRequestInformation retrieves a list of Notes
// returns a *RequestInformation when successful
func (m *ContactsItemNotesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ContactsItemNotesRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToPostRequestInformation creates a new Note.
// returns a *RequestInformation when successful
func (m *ContactsItemNotesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ContactsItemNotesPostRequestBodyable, requestConfiguration *ContactsItemNotesRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/v2/contacts/{contact_id}/notes", m.BaseRequestBuilder.PathParameters)
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
// returns a *ContactsItemNotesRequestBuilder when successful
func (m *ContactsItemNotesRequestBuilder) WithUrl(rawUrl string) *ContactsItemNotesRequestBuilder {
	return NewContactsItemNotesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
