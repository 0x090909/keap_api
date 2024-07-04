package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// NotesModelCustomFieldsRequestBuilder builds and executes requests for operations under \v1\notes\model\customFields
type NotesModelCustomFieldsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// NotesModelCustomFieldsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type NotesModelCustomFieldsRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewNotesModelCustomFieldsRequestBuilderInternal instantiates a new NotesModelCustomFieldsRequestBuilder and sets the default values.
func NewNotesModelCustomFieldsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *NotesModelCustomFieldsRequestBuilder {
	m := &NotesModelCustomFieldsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/notes/model/customFields", pathParameters),
	}
	return m
}

// NewNotesModelCustomFieldsRequestBuilder instantiates a new NotesModelCustomFieldsRequestBuilder and sets the default values.
func NewNotesModelCustomFieldsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *NotesModelCustomFieldsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewNotesModelCustomFieldsRequestBuilderInternal(urlParams, requestAdapter)
}

// Post adds a custom field of the specified type and options to the Note object.
// Deprecated: This method is obsolete. Use PostAsCustomFieldsPostResponse instead.
// returns a NotesModelCustomFieldsResponseable when successful
// returns a NotesModelCustomFields401Error error when the service returns a 401 status code
// returns a NotesModelCustomFields403Error error when the service returns a 403 status code
func (m *NotesModelCustomFieldsRequestBuilder) Post(ctx context.Context, body NotesModelCustomFieldsPostRequestBodyable, requestConfiguration *NotesModelCustomFieldsRequestBuilderPostRequestConfiguration) (NotesModelCustomFieldsResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateNotesModelCustomFields401ErrorFromDiscriminatorValue,
		"403": CreateNotesModelCustomFields403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateNotesModelCustomFieldsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(NotesModelCustomFieldsResponseable), nil
}

// PostAsCustomFieldsPostResponse adds a custom field of the specified type and options to the Note object.
// returns a NotesModelCustomFieldsPostResponseable when successful
// returns a NotesModelCustomFields401Error error when the service returns a 401 status code
// returns a NotesModelCustomFields403Error error when the service returns a 403 status code
func (m *NotesModelCustomFieldsRequestBuilder) PostAsCustomFieldsPostResponse(ctx context.Context, body NotesModelCustomFieldsPostRequestBodyable, requestConfiguration *NotesModelCustomFieldsRequestBuilderPostRequestConfiguration) (NotesModelCustomFieldsPostResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateNotesModelCustomFields401ErrorFromDiscriminatorValue,
		"403": CreateNotesModelCustomFields403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateNotesModelCustomFieldsPostResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(NotesModelCustomFieldsPostResponseable), nil
}

// ToPostRequestInformation adds a custom field of the specified type and options to the Note object.
// returns a *RequestInformation when successful
func (m *NotesModelCustomFieldsRequestBuilder) ToPostRequestInformation(ctx context.Context, body NotesModelCustomFieldsPostRequestBodyable, requestConfiguration *NotesModelCustomFieldsRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *NotesModelCustomFieldsRequestBuilder when successful
func (m *NotesModelCustomFieldsRequestBuilder) WithUrl(rawUrl string) *NotesModelCustomFieldsRequestBuilder {
	return NewNotesModelCustomFieldsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
