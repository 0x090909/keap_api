package v1

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// OauthRequestBuilder builds and executes requests for operations under \v1\oauth
type OauthRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// Connect the connect property
// returns a *OauthConnectRequestBuilder when successful
func (m *OauthRequestBuilder) Connect() *OauthConnectRequestBuilder {
	return NewOauthConnectRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// NewOauthRequestBuilderInternal instantiates a new OauthRequestBuilder and sets the default values.
func NewOauthRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *OauthRequestBuilder {
	m := &OauthRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/oauth", pathParameters),
	}
	return m
}

// NewOauthRequestBuilder instantiates a new OauthRequestBuilder and sets the default values.
func NewOauthRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *OauthRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewOauthRequestBuilderInternal(urlParams, requestAdapter)
}
