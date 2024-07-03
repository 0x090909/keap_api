package v2

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// OauthConnectUserinfoRequestBuilder builds and executes requests for operations under \v2\oauth\connect\userinfo
type OauthConnectUserinfoRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// NewOauthConnectUserinfoRequestBuilderInternal instantiates a new OauthConnectUserinfoRequestBuilder and sets the default values.
func NewOauthConnectUserinfoRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *OauthConnectUserinfoRequestBuilder {
	m := &OauthConnectUserinfoRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/oauth/connect/userinfo", pathParameters),
	}
	return m
}

// NewOauthConnectUserinfoRequestBuilder instantiates a new OauthConnectUserinfoRequestBuilder and sets the default values.
func NewOauthConnectUserinfoRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *OauthConnectUserinfoRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewOauthConnectUserinfoRequestBuilderInternal(urlParams, requestAdapter)
}
