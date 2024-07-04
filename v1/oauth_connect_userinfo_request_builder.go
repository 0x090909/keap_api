package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// OauthConnectUserinfoRequestBuilder builds and executes requests for operations under \v1\oauth\connect\userinfo
type OauthConnectUserinfoRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// OauthConnectUserinfoRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type OauthConnectUserinfoRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewOauthConnectUserinfoRequestBuilderInternal instantiates a new OauthConnectUserinfoRequestBuilder and sets the default values.
func NewOauthConnectUserinfoRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *OauthConnectUserinfoRequestBuilder {
	m := &OauthConnectUserinfoRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/oauth/connect/userinfo", pathParameters),
	}
	return m
}

// NewOauthConnectUserinfoRequestBuilder instantiates a new OauthConnectUserinfoRequestBuilder and sets the default values.
func NewOauthConnectUserinfoRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *OauthConnectUserinfoRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewOauthConnectUserinfoRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieves information for the current authenticated end-user, as outlined by the [OpenID Connect specification](http://openid.net/specs/openid-connect-core-1_0.html#UserInfo).
// Deprecated: This method is obsolete. Use GetAsUserinfoGetResponse instead.
// returns a OauthConnectUserinfoResponseable when successful
// returns a OauthConnectUserinfo401Error error when the service returns a 401 status code
// returns a OauthConnectUserinfo403Error error when the service returns a 403 status code
// returns a OauthConnectUserinfo404Error error when the service returns a 404 status code
func (m *OauthConnectUserinfoRequestBuilder) Get(ctx context.Context, requestConfiguration *OauthConnectUserinfoRequestBuilderGetRequestConfiguration) (OauthConnectUserinfoResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateOauthConnectUserinfo401ErrorFromDiscriminatorValue,
		"403": CreateOauthConnectUserinfo403ErrorFromDiscriminatorValue,
		"404": CreateOauthConnectUserinfo404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOauthConnectUserinfoResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(OauthConnectUserinfoResponseable), nil
}

// GetAsUserinfoGetResponse retrieves information for the current authenticated end-user, as outlined by the [OpenID Connect specification](http://openid.net/specs/openid-connect-core-1_0.html#UserInfo).
// returns a OauthConnectUserinfoGetResponseable when successful
// returns a OauthConnectUserinfo401Error error when the service returns a 401 status code
// returns a OauthConnectUserinfo403Error error when the service returns a 403 status code
// returns a OauthConnectUserinfo404Error error when the service returns a 404 status code
func (m *OauthConnectUserinfoRequestBuilder) GetAsUserinfoGetResponse(ctx context.Context, requestConfiguration *OauthConnectUserinfoRequestBuilderGetRequestConfiguration) (OauthConnectUserinfoGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateOauthConnectUserinfo401ErrorFromDiscriminatorValue,
		"403": CreateOauthConnectUserinfo403ErrorFromDiscriminatorValue,
		"404": CreateOauthConnectUserinfo404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateOauthConnectUserinfoGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(OauthConnectUserinfoGetResponseable), nil
}

// ToGetRequestInformation retrieves information for the current authenticated end-user, as outlined by the [OpenID Connect specification](http://openid.net/specs/openid-connect-core-1_0.html#UserInfo).
// returns a *RequestInformation when successful
func (m *OauthConnectUserinfoRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *OauthConnectUserinfoRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *OauthConnectUserinfoRequestBuilder when successful
func (m *OauthConnectUserinfoRequestBuilder) WithUrl(rawUrl string) *OauthConnectUserinfoRequestBuilder {
	return NewOauthConnectUserinfoRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
