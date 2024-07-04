package v1

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// EmailAddressesRequestBuilder builds and executes requests for operations under \v1\emailAddresses
type EmailAddressesRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ByEmail gets an item from the github.com/0x090909/keap_api.v1.emailAddresses.item collection
// returns a *EmailAddressesWithEmailItemRequestBuilder when successful
func (m *EmailAddressesRequestBuilder) ByEmail(email string) *EmailAddressesWithEmailItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if email != "" {
		urlTplParams["email"] = email
	}
	return NewEmailAddressesWithEmailItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewEmailAddressesRequestBuilderInternal instantiates a new EmailAddressesRequestBuilder and sets the default values.
func NewEmailAddressesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *EmailAddressesRequestBuilder {
	m := &EmailAddressesRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/emailAddresses", pathParameters),
	}
	return m
}

// NewEmailAddressesRequestBuilder instantiates a new EmailAddressesRequestBuilder and sets the default values.
func NewEmailAddressesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *EmailAddressesRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewEmailAddressesRequestBuilderInternal(urlParams, requestAdapter)
}
