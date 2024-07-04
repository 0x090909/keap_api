package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ContactsItemTagsWithTagItemRequestBuilder builds and executes requests for operations under \v1\contacts\{contactId-id}\tags\{tagId}
type ContactsItemTagsWithTagItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ContactsItemTagsWithTagItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContactsItemTagsWithTagItemRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewContactsItemTagsWithTagItemRequestBuilderInternal instantiates a new ContactsItemTagsWithTagItemRequestBuilder and sets the default values.
func NewContactsItemTagsWithTagItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ContactsItemTagsWithTagItemRequestBuilder {
	m := &ContactsItemTagsWithTagItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/contacts/{contactId%2Did}/tags/{tagId}", pathParameters),
	}
	return m
}

// NewContactsItemTagsWithTagItemRequestBuilder instantiates a new ContactsItemTagsWithTagItemRequestBuilder and sets the default values.
func NewContactsItemTagsWithTagItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ContactsItemTagsWithTagItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewContactsItemTagsWithTagItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Delete removes a tag from the given contact
// returns a ContactsItemTagsItemWithTag401Error error when the service returns a 401 status code
// returns a ContactsItemTagsItemWithTag403Error error when the service returns a 403 status code
// returns a ContactsItemTagsItemWithTag404Error error when the service returns a 404 status code
func (m *ContactsItemTagsWithTagItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ContactsItemTagsWithTagItemRequestBuilderDeleteRequestConfiguration) error {
	requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateContactsItemTagsItemWithTag401ErrorFromDiscriminatorValue,
		"403": CreateContactsItemTagsItemWithTag403ErrorFromDiscriminatorValue,
		"404": CreateContactsItemTagsItemWithTag404ErrorFromDiscriminatorValue,
	}
	err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
	if err != nil {
		return err
	}
	return nil
}

// ToDeleteRequestInformation removes a tag from the given contact
// returns a *RequestInformation when successful
func (m *ContactsItemTagsWithTagItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ContactsItemTagsWithTagItemRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ContactsItemTagsWithTagItemRequestBuilder when successful
func (m *ContactsItemTagsWithTagItemRequestBuilder) WithUrl(rawUrl string) *ContactsItemTagsWithTagItemRequestBuilder {
	return NewContactsItemTagsWithTagItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
