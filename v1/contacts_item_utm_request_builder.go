package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ContactsItemUtmRequestBuilder builds and executes requests for operations under \v1\contacts\{contactId-id}\utm
type ContactsItemUtmRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ContactsItemUtmRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContactsItemUtmRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewContactsItemUtmRequestBuilderInternal instantiates a new ContactsItemUtmRequestBuilder and sets the default values.
func NewContactsItemUtmRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ContactsItemUtmRequestBuilder {
	m := &ContactsItemUtmRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/contacts/{contactId%2Did}/utm", pathParameters),
	}
	return m
}

// NewContactsItemUtmRequestBuilder instantiates a new ContactsItemUtmRequestBuilder and sets the default values.
func NewContactsItemUtmRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ContactsItemUtmRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewContactsItemUtmRequestBuilderInternal(urlParams, requestAdapter)
}

// Post inserts UTMs for the contact specified by the contactId. The authenticated user must have permission to modify the contact utm info
// Deprecated: This method is obsolete. Use PostAsUtmPostResponse instead.
// returns a ContactsItemUtmResponseable when successful
// returns a ContactsItemUtm401Error error when the service returns a 401 status code
// returns a ContactsItemUtm403Error error when the service returns a 403 status code
func (m *ContactsItemUtmRequestBuilder) Post(ctx context.Context, body ContactsItemUtmPostRequestBodyable, requestConfiguration *ContactsItemUtmRequestBuilderPostRequestConfiguration) (ContactsItemUtmResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsItemUtm401ErrorFromDiscriminatorValue,
		"403": CreateContactsItemUtm403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsItemUtmResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsItemUtmResponseable), nil
}

// PostAsUtmPostResponse inserts UTMs for the contact specified by the contactId. The authenticated user must have permission to modify the contact utm info
// returns a ContactsItemUtmPostResponseable when successful
// returns a ContactsItemUtm401Error error when the service returns a 401 status code
// returns a ContactsItemUtm403Error error when the service returns a 403 status code
func (m *ContactsItemUtmRequestBuilder) PostAsUtmPostResponse(ctx context.Context, body ContactsItemUtmPostRequestBodyable, requestConfiguration *ContactsItemUtmRequestBuilderPostRequestConfiguration) (ContactsItemUtmPostResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsItemUtm401ErrorFromDiscriminatorValue,
		"403": CreateContactsItemUtm403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsItemUtmPostResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsItemUtmPostResponseable), nil
}

// ToPostRequestInformation inserts UTMs for the contact specified by the contactId. The authenticated user must have permission to modify the contact utm info
// returns a *RequestInformation when successful
func (m *ContactsItemUtmRequestBuilder) ToPostRequestInformation(ctx context.Context, body ContactsItemUtmPostRequestBodyable, requestConfiguration *ContactsItemUtmRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ContactsItemUtmRequestBuilder when successful
func (m *ContactsItemUtmRequestBuilder) WithUrl(rawUrl string) *ContactsItemUtmRequestBuilder {
	return NewContactsItemUtmRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
