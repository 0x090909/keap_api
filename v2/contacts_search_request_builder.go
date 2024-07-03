package v2

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ContactsSearchRequestBuilder builds and executes requests for operations under \v2\contacts:search
type ContactsSearchRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// NewContactsSearchRequestBuilderInternal instantiates a new ContactsSearchRequestBuilder and sets the default values.
func NewContactsSearchRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ContactsSearchRequestBuilder {
	m := &ContactsSearchRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/contacts:search", pathParameters),
	}
	return m
}

// NewContactsSearchRequestBuilder instantiates a new ContactsSearchRequestBuilder and sets the default values.
func NewContactsSearchRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ContactsSearchRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewContactsSearchRequestBuilderInternal(urlParams, requestAdapter)
}
