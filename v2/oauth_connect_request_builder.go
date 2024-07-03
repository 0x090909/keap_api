package v2

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// OauthConnectRequestBuilder builds and executes requests for operations under \v2\oauth\connect
type OauthConnectRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// NewOauthConnectRequestBuilderInternal instantiates a new OauthConnectRequestBuilder and sets the default values.
func NewOauthConnectRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *OauthConnectRequestBuilder {
	m := &OauthConnectRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/oauth/connect", pathParameters),
	}
	return m
}

// NewOauthConnectRequestBuilder instantiates a new OauthConnectRequestBuilder and sets the default values.
func NewOauthConnectRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *OauthConnectRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewOauthConnectRequestBuilderInternal(urlParams, requestAdapter)
}

// Userinfo the userinfo property
// returns a *OauthConnectUserinfoRequestBuilder when successful
func (m *OauthConnectRequestBuilder) Userinfo() *OauthConnectUserinfoRequestBuilder {
	return NewOauthConnectUserinfoRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
