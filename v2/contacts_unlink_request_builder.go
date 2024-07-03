package v2

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ContactsUnlinkRequestBuilder builds and executes requests for operations under \v2\contacts:unlink
type ContactsUnlinkRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ContactsUnlinkRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContactsUnlinkRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewContactsUnlinkRequestBuilderInternal instantiates a new ContactsUnlinkRequestBuilder and sets the default values.
func NewContactsUnlinkRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ContactsUnlinkRequestBuilder {
	m := &ContactsUnlinkRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/contacts:unlink", pathParameters),
	}
	return m
}

// NewContactsUnlinkRequestBuilder instantiates a new ContactsUnlinkRequestBuilder and sets the default values.
func NewContactsUnlinkRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ContactsUnlinkRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewContactsUnlinkRequestBuilderInternal(urlParams, requestAdapter)
}

// Post deletes Link between two Contacts
// returns a ContactsUnlink401Error error when the service returns a 401 status code
// returns a ContactsUnlink403Error error when the service returns a 403 status code
func (m *ContactsUnlinkRequestBuilder) Post(ctx context.Context, body ContactsUnlinkPostRequestBodyable, requestConfiguration *ContactsUnlinkRequestBuilderPostRequestConfiguration) error {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsUnlink401ErrorFromDiscriminatorValue,
		"403": CreateContactsUnlink403ErrorFromDiscriminatorValue,
	}
	err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
	if err != nil {
		return err
	}
	return nil
}

// ToPostRequestInformation deletes Link between two Contacts
// returns a *RequestInformation when successful
func (m *ContactsUnlinkRequestBuilder) ToPostRequestInformation(ctx context.Context, body ContactsUnlinkPostRequestBodyable, requestConfiguration *ContactsUnlinkRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ContactsUnlinkRequestBuilder when successful
func (m *ContactsUnlinkRequestBuilder) WithUrl(rawUrl string) *ContactsUnlinkRequestBuilder {
	return NewContactsUnlinkRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
