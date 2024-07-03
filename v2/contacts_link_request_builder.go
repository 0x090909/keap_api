package v2

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ContactsLinkRequestBuilder builds and executes requests for operations under \v2\contacts:link
type ContactsLinkRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ContactsLinkRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContactsLinkRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewContactsLinkRequestBuilderInternal instantiates a new ContactsLinkRequestBuilder and sets the default values.
func NewContactsLinkRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ContactsLinkRequestBuilder {
	m := &ContactsLinkRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/contacts:link", pathParameters),
	}
	return m
}

// NewContactsLinkRequestBuilder instantiates a new ContactsLinkRequestBuilder and sets the default values.
func NewContactsLinkRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ContactsLinkRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewContactsLinkRequestBuilderInternal(urlParams, requestAdapter)
}

// Post links two Contacts together using the provided Link type
// Deprecated: This method is obsolete. Use PostAsContactsLinkPostResponse instead.
// returns a ContactsLinkResponseable when successful
// returns a ContactsLink401Error error when the service returns a 401 status code
// returns a ContactsLink403Error error when the service returns a 403 status code
func (m *ContactsLinkRequestBuilder) Post(ctx context.Context, body ContactsLinkPostRequestBodyable, requestConfiguration *ContactsLinkRequestBuilderPostRequestConfiguration) (ContactsLinkResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsLink401ErrorFromDiscriminatorValue,
		"403": CreateContactsLink403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsLinkResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsLinkResponseable), nil
}

// PostAsContactsLinkPostResponse links two Contacts together using the provided Link type
// returns a ContactsLinkPostResponseable when successful
// returns a ContactsLink401Error error when the service returns a 401 status code
// returns a ContactsLink403Error error when the service returns a 403 status code
func (m *ContactsLinkRequestBuilder) PostAsContactsLinkPostResponse(ctx context.Context, body ContactsLinkPostRequestBodyable, requestConfiguration *ContactsLinkRequestBuilderPostRequestConfiguration) (ContactsLinkPostResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsLink401ErrorFromDiscriminatorValue,
		"403": CreateContactsLink403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsLinkPostResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsLinkPostResponseable), nil
}

// ToPostRequestInformation links two Contacts together using the provided Link type
// returns a *RequestInformation when successful
func (m *ContactsLinkRequestBuilder) ToPostRequestInformation(ctx context.Context, body ContactsLinkPostRequestBodyable, requestConfiguration *ContactsLinkRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ContactsLinkRequestBuilder when successful
func (m *ContactsLinkRequestBuilder) WithUrl(rawUrl string) *ContactsLinkRequestBuilder {
	return NewContactsLinkRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
