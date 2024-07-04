package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
	i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
)

// UsersRequestBuilder builds and executes requests for operations under \v1\users
type UsersRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// UsersRequestBuilderGetQueryParameters retrieves a list of all users
type UsersRequestBuilderGetQueryParameters struct {
	// Include users that are Inactive in results, defaults to TRUE
	Include_inactive *bool `uriparametername:"include_inactive"`
	// Include partner users in results, defaults to TRUE
	Include_partners *bool `uriparametername:"include_partners"`
	// Sets a total of items to return
	Limit *int32 `uriparametername:"limit"`
	// Sets a beginning range of items to return
	Offset *int32 `uriparametername:"offset"`
}

// UsersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *UsersRequestBuilderGetQueryParameters
}

// UsersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ByUserId gets an item from the github.com/0x090909/keap_api.v1.users.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *UsersWithUserItemRequestBuilder when successful
func (m *UsersRequestBuilder) ByUserId(userId string) *UsersWithUserItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if userId != "" {
		urlTplParams["userId"] = userId
	}
	return NewUsersWithUserItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// ByUserIdInt64 gets an item from the github.com/0x090909/keap_api.v1.users.item collection
// returns a *UsersWithUserItemRequestBuilder when successful
func (m *UsersRequestBuilder) ByUserIdInt64(userId int64) *UsersWithUserItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	urlTplParams["userId"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(userId, 10)
	return NewUsersWithUserItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewUsersRequestBuilderInternal instantiates a new UsersRequestBuilder and sets the default values.
func NewUsersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *UsersRequestBuilder {
	m := &UsersRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/users{?include_inactive*,include_partners*,limit*,offset*}", pathParameters),
	}
	return m
}

// NewUsersRequestBuilder instantiates a new UsersRequestBuilder and sets the default values.
func NewUsersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *UsersRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewUsersRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieves a list of all users
// Deprecated: This method is obsolete. Use GetAsUsersGetResponse instead.
// returns a UsersResponseable when successful
// returns a Users401Error error when the service returns a 401 status code
// returns a Users403Error error when the service returns a 403 status code
// returns a Users404Error error when the service returns a 404 status code
func (m *UsersRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersRequestBuilderGetRequestConfiguration) (UsersResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateUsers401ErrorFromDiscriminatorValue,
		"403": CreateUsers403ErrorFromDiscriminatorValue,
		"404": CreateUsers404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateUsersResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(UsersResponseable), nil
}

// GetAsUsersGetResponse retrieves a list of all users
// returns a UsersGetResponseable when successful
// returns a Users401Error error when the service returns a 401 status code
// returns a Users403Error error when the service returns a 403 status code
// returns a Users404Error error when the service returns a 404 status code
func (m *UsersRequestBuilder) GetAsUsersGetResponse(ctx context.Context, requestConfiguration *UsersRequestBuilderGetRequestConfiguration) (UsersGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateUsers401ErrorFromDiscriminatorValue,
		"403": CreateUsers403ErrorFromDiscriminatorValue,
		"404": CreateUsers404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateUsersGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(UsersGetResponseable), nil
}

// Post creates a new user record. NB: Users will be invited to the application and remain in the "Invited" status until the user accepts the invite. "Inactive" users will not take up a user license.
// Deprecated: This method is obsolete. Use PostAsUsersPostResponse instead.
// returns a UsersResponseable when successful
// returns a Users401Error error when the service returns a 401 status code
// returns a Users403Error error when the service returns a 403 status code
func (m *UsersRequestBuilder) Post(ctx context.Context, body UsersPostRequestBodyable, requestConfiguration *UsersRequestBuilderPostRequestConfiguration) (UsersResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateUsers401ErrorFromDiscriminatorValue,
		"403": CreateUsers403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateUsersResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(UsersResponseable), nil
}

// PostAsUsersPostResponse creates a new user record. NB: Users will be invited to the application and remain in the "Invited" status until the user accepts the invite. "Inactive" users will not take up a user license.
// returns a UsersPostResponseable when successful
// returns a Users401Error error when the service returns a 401 status code
// returns a Users403Error error when the service returns a 403 status code
func (m *UsersRequestBuilder) PostAsUsersPostResponse(ctx context.Context, body UsersPostRequestBodyable, requestConfiguration *UsersRequestBuilderPostRequestConfiguration) (UsersPostResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateUsers401ErrorFromDiscriminatorValue,
		"403": CreateUsers403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateUsersPostResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(UsersPostResponseable), nil
}

// ToGetRequestInformation retrieves a list of all users
// returns a *RequestInformation when successful
func (m *UsersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UsersRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		if requestConfiguration.QueryParameters != nil {
			requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
		}
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// ToPostRequestInformation creates a new user record. NB: Users will be invited to the application and remain in the "Invited" status until the user accepts the invite. "Inactive" users will not take up a user license.
// returns a *RequestInformation when successful
func (m *UsersRequestBuilder) ToPostRequestInformation(ctx context.Context, body UsersPostRequestBodyable, requestConfiguration *UsersRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/v1/users", m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
	if err != nil {
		return nil, err
	}
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *UsersRequestBuilder when successful
func (m *UsersRequestBuilder) WithUrl(rawUrl string) *UsersRequestBuilder {
	return NewUsersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
