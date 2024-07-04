package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ContactsModelCustomFieldsRequestBuilder builds and executes requests for operations under \v1\contacts\model\customFields
type ContactsModelCustomFieldsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ContactsModelCustomFieldsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContactsModelCustomFieldsRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewContactsModelCustomFieldsRequestBuilderInternal instantiates a new ContactsModelCustomFieldsRequestBuilder and sets the default values.
func NewContactsModelCustomFieldsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ContactsModelCustomFieldsRequestBuilder {
	m := &ContactsModelCustomFieldsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/contacts/model/customFields", pathParameters),
	}
	return m
}

// NewContactsModelCustomFieldsRequestBuilder instantiates a new ContactsModelCustomFieldsRequestBuilder and sets the default values.
func NewContactsModelCustomFieldsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ContactsModelCustomFieldsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewContactsModelCustomFieldsRequestBuilderInternal(urlParams, requestAdapter)
}

// Post adds a custom field of the specified type and options to the Contact object.
// Deprecated: This method is obsolete. Use PostAsCustomFieldsPostResponse instead.
// returns a ContactsModelCustomFieldsResponseable when successful
// returns a ContactsModelCustomFields401Error error when the service returns a 401 status code
// returns a ContactsModelCustomFields403Error error when the service returns a 403 status code
func (m *ContactsModelCustomFieldsRequestBuilder) Post(ctx context.Context, body ContactsModelCustomFieldsPostRequestBodyable, requestConfiguration *ContactsModelCustomFieldsRequestBuilderPostRequestConfiguration) (ContactsModelCustomFieldsResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsModelCustomFields401ErrorFromDiscriminatorValue,
		"403": CreateContactsModelCustomFields403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsModelCustomFieldsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsModelCustomFieldsResponseable), nil
}

// PostAsCustomFieldsPostResponse adds a custom field of the specified type and options to the Contact object.
// returns a ContactsModelCustomFieldsPostResponseable when successful
// returns a ContactsModelCustomFields401Error error when the service returns a 401 status code
// returns a ContactsModelCustomFields403Error error when the service returns a 403 status code
func (m *ContactsModelCustomFieldsRequestBuilder) PostAsCustomFieldsPostResponse(ctx context.Context, body ContactsModelCustomFieldsPostRequestBodyable, requestConfiguration *ContactsModelCustomFieldsRequestBuilderPostRequestConfiguration) (ContactsModelCustomFieldsPostResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsModelCustomFields401ErrorFromDiscriminatorValue,
		"403": CreateContactsModelCustomFields403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateContactsModelCustomFieldsPostResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(ContactsModelCustomFieldsPostResponseable), nil
}

// ToPostRequestInformation adds a custom field of the specified type and options to the Contact object.
// returns a *RequestInformation when successful
func (m *ContactsModelCustomFieldsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ContactsModelCustomFieldsPostRequestBodyable, requestConfiguration *ContactsModelCustomFieldsRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ContactsModelCustomFieldsRequestBuilder when successful
func (m *ContactsModelCustomFieldsRequestBuilder) WithUrl(rawUrl string) *ContactsModelCustomFieldsRequestBuilder {
	return NewContactsModelCustomFieldsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
