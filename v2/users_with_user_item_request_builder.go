package v2

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// UsersWithUser_ItemRequestBuilder builds and executes requests for operations under \v2\users\{user_id}
type UsersWithUser_ItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// NewUsersWithUser_ItemRequestBuilderInternal instantiates a new UsersWithUser_ItemRequestBuilder and sets the default values.
func NewUsersWithUser_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *UsersWithUser_ItemRequestBuilder {
	m := &UsersWithUser_ItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/users/{user_id}", pathParameters),
	}
	return m
}

// NewUsersWithUser_ItemRequestBuilder instantiates a new UsersWithUser_ItemRequestBuilder and sets the default values.
func NewUsersWithUser_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *UsersWithUser_ItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewUsersWithUser_ItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Signature the signature property
// returns a *UsersItemSignatureRequestBuilder when successful
func (m *UsersWithUser_ItemRequestBuilder) Signature() *UsersItemSignatureRequestBuilder {
	return NewUsersItemSignatureRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
