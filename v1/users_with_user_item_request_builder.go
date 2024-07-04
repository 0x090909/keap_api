package v1

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// UsersWithUserItemRequestBuilder builds and executes requests for operations under \v1\users\{userId}
type UsersWithUserItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// NewUsersWithUserItemRequestBuilderInternal instantiates a new UsersWithUserItemRequestBuilder and sets the default values.
func NewUsersWithUserItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *UsersWithUserItemRequestBuilder {
	m := &UsersWithUserItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/users/{userId}", pathParameters),
	}
	return m
}

// NewUsersWithUserItemRequestBuilder instantiates a new UsersWithUserItemRequestBuilder and sets the default values.
func NewUsersWithUserItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *UsersWithUserItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewUsersWithUserItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Signature the signature property
// returns a *UsersItemSignatureRequestBuilder when successful
func (m *UsersWithUserItemRequestBuilder) Signature() *UsersItemSignatureRequestBuilder {
	return NewUsersItemSignatureRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
