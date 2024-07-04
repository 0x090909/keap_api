package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// NotesModelRequestBuilder builds and executes requests for operations under \v1\notes\model
type NotesModelRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// NotesModelRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type NotesModelRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewNotesModelRequestBuilderInternal instantiates a new NotesModelRequestBuilder and sets the default values.
func NewNotesModelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *NotesModelRequestBuilder {
	m := &NotesModelRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/notes/model", pathParameters),
	}
	return m
}

// NewNotesModelRequestBuilder instantiates a new NotesModelRequestBuilder and sets the default values.
func NewNotesModelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *NotesModelRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewNotesModelRequestBuilderInternal(urlParams, requestAdapter)
}

// CustomFields the customFields property
// returns a *NotesModelCustomFieldsRequestBuilder when successful
func (m *NotesModelRequestBuilder) CustomFields() *NotesModelCustomFieldsRequestBuilder {
	return NewNotesModelCustomFieldsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Get get the custom fields for the Note object
// Deprecated: This method is obsolete. Use GetAsModelGetResponse instead.
// returns a NotesModelResponseable when successful
// returns a NotesModel401Error error when the service returns a 401 status code
// returns a NotesModel403Error error when the service returns a 403 status code
// returns a NotesModel404Error error when the service returns a 404 status code
func (m *NotesModelRequestBuilder) Get(ctx context.Context, requestConfiguration *NotesModelRequestBuilderGetRequestConfiguration) (NotesModelResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateNotesModel401ErrorFromDiscriminatorValue,
		"403": CreateNotesModel403ErrorFromDiscriminatorValue,
		"404": CreateNotesModel404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateNotesModelResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(NotesModelResponseable), nil
}

// GetAsModelGetResponse get the custom fields for the Note object
// returns a NotesModelGetResponseable when successful
// returns a NotesModel401Error error when the service returns a 401 status code
// returns a NotesModel403Error error when the service returns a 403 status code
// returns a NotesModel404Error error when the service returns a 404 status code
func (m *NotesModelRequestBuilder) GetAsModelGetResponse(ctx context.Context, requestConfiguration *NotesModelRequestBuilderGetRequestConfiguration) (NotesModelGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateNotesModel401ErrorFromDiscriminatorValue,
		"403": CreateNotesModel403ErrorFromDiscriminatorValue,
		"404": CreateNotesModel404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateNotesModelGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(NotesModelGetResponseable), nil
}

// ToGetRequestInformation get the custom fields for the Note object
// returns a *RequestInformation when successful
func (m *NotesModelRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *NotesModelRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *NotesModelRequestBuilder when successful
func (m *NotesModelRequestBuilder) WithUrl(rawUrl string) *NotesModelRequestBuilder {
	return NewNotesModelRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
