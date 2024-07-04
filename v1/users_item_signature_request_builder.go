package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// UsersItemSignatureRequestBuilder builds and executes requests for operations under \v1\users\{userId}\signature
type UsersItemSignatureRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// UsersItemSignatureRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemSignatureRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewUsersItemSignatureRequestBuilderInternal instantiates a new UsersItemSignatureRequestBuilder and sets the default values.
func NewUsersItemSignatureRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *UsersItemSignatureRequestBuilder {
	m := &UsersItemSignatureRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/users/{userId}/signature", pathParameters),
	}
	return m
}

// NewUsersItemSignatureRequestBuilder instantiates a new UsersItemSignatureRequestBuilder and sets the default values.
func NewUsersItemSignatureRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *UsersItemSignatureRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewUsersItemSignatureRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieves a HTML snippet that contains the user's email signature.
// returns a *string when successful
// returns a UsersItemSignature401Error error when the service returns a 401 status code
// returns a UsersItemSignature403Error error when the service returns a 403 status code
// returns a UsersItemSignature404Error error when the service returns a 404 status code
func (m *UsersItemSignatureRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemSignatureRequestBuilderGetRequestConfiguration) (*string, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateUsersItemSignature401ErrorFromDiscriminatorValue,
		"403": CreateUsersItemSignature403ErrorFromDiscriminatorValue,
		"404": CreateUsersItemSignature404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "string", errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(*string), nil
}

// ToGetRequestInformation retrieves a HTML snippet that contains the user's email signature.
// returns a *RequestInformation when successful
func (m *UsersItemSignatureRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemSignatureRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *UsersItemSignatureRequestBuilder when successful
func (m *UsersItemSignatureRequestBuilder) WithUrl(rawUrl string) *UsersItemSignatureRequestBuilder {
	return NewUsersItemSignatureRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
