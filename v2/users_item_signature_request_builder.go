package v2

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// UsersItemSignatureRequestBuilder builds and executes requests for operations under \v2\users\{user_id}\signature
type UsersItemSignatureRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// NewUsersItemSignatureRequestBuilderInternal instantiates a new UsersItemSignatureRequestBuilder and sets the default values.
func NewUsersItemSignatureRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *UsersItemSignatureRequestBuilder {
	m := &UsersItemSignatureRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/users/{user_id}/signature", pathParameters),
	}
	return m
}

// NewUsersItemSignatureRequestBuilder instantiates a new UsersItemSignatureRequestBuilder and sets the default values.
func NewUsersItemSignatureRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *UsersItemSignatureRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewUsersItemSignatureRequestBuilderInternal(urlParams, requestAdapter)
}
