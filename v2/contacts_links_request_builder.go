package v2

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ContactsLinksRequestBuilder builds and executes requests for operations under \v2\contacts\links
type ContactsLinksRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// NewContactsLinksRequestBuilderInternal instantiates a new ContactsLinksRequestBuilder and sets the default values.
func NewContactsLinksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ContactsLinksRequestBuilder {
	m := &ContactsLinksRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/contacts/links", pathParameters),
	}
	return m
}

// NewContactsLinksRequestBuilder instantiates a new ContactsLinksRequestBuilder and sets the default values.
func NewContactsLinksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ContactsLinksRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewContactsLinksRequestBuilderInternal(urlParams, requestAdapter)
}

// Types the types property
// returns a *ContactsLinksTypesRequestBuilder when successful
func (m *ContactsLinksRequestBuilder) Types() *ContactsLinksTypesRequestBuilder {
	return NewContactsLinksTypesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
