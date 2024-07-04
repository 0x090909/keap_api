package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SubscriptionsRequestBuilder builds and executes requests for operations under \v1\subscriptions
type SubscriptionsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// SubscriptionsRequestBuilderGetQueryParameters retrieves a list of all subcriptions using the specified search criteria.
type SubscriptionsRequestBuilderGetQueryParameters struct {
	// Returns subscriptions for the provided contact id
	Contact_id *int64 `uriparametername:"contact_id"`
	// Sets a total of items to return
	Limit *int32 `uriparametername:"limit"`
	// Sets a beginning range of items to return
	Offset *int32 `uriparametername:"offset"`
}

// SubscriptionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SubscriptionsRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *SubscriptionsRequestBuilderGetQueryParameters
}

// SubscriptionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SubscriptionsRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewSubscriptionsRequestBuilderInternal instantiates a new SubscriptionsRequestBuilder and sets the default values.
func NewSubscriptionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *SubscriptionsRequestBuilder {
	m := &SubscriptionsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/subscriptions{?contact_id*,limit*,offset*}", pathParameters),
	}
	return m
}

// NewSubscriptionsRequestBuilder instantiates a new SubscriptionsRequestBuilder and sets the default values.
func NewSubscriptionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *SubscriptionsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewSubscriptionsRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieves a list of all subcriptions using the specified search criteria.
// Deprecated: This method is obsolete. Use GetAsSubscriptionsGetResponse instead.
// returns a SubscriptionsResponseable when successful
// returns a Subscriptions401Error error when the service returns a 401 status code
// returns a Subscriptions403Error error when the service returns a 403 status code
// returns a Subscriptions404Error error when the service returns a 404 status code
func (m *SubscriptionsRequestBuilder) Get(ctx context.Context, requestConfiguration *SubscriptionsRequestBuilderGetRequestConfiguration) (SubscriptionsResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateSubscriptions401ErrorFromDiscriminatorValue,
		"403": CreateSubscriptions403ErrorFromDiscriminatorValue,
		"404": CreateSubscriptions404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSubscriptionsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(SubscriptionsResponseable), nil
}

// GetAsSubscriptionsGetResponse retrieves a list of all subcriptions using the specified search criteria.
// returns a SubscriptionsGetResponseable when successful
// returns a Subscriptions401Error error when the service returns a 401 status code
// returns a Subscriptions403Error error when the service returns a 403 status code
// returns a Subscriptions404Error error when the service returns a 404 status code
func (m *SubscriptionsRequestBuilder) GetAsSubscriptionsGetResponse(ctx context.Context, requestConfiguration *SubscriptionsRequestBuilderGetRequestConfiguration) (SubscriptionsGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateSubscriptions401ErrorFromDiscriminatorValue,
		"403": CreateSubscriptions403ErrorFromDiscriminatorValue,
		"404": CreateSubscriptions404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSubscriptionsGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(SubscriptionsGetResponseable), nil
}

// Model the model property
// returns a *SubscriptionsModelRequestBuilder when successful
func (m *SubscriptionsRequestBuilder) Model() *SubscriptionsModelRequestBuilder {
	return NewSubscriptionsModelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Post creates a subscription with the specified product and product subscription id.
// Deprecated: This method is obsolete. Use PostAsSubscriptionsPostResponse instead.
// returns a SubscriptionsResponseable when successful
// returns a Subscriptions401Error error when the service returns a 401 status code
// returns a Subscriptions403Error error when the service returns a 403 status code
func (m *SubscriptionsRequestBuilder) Post(ctx context.Context, body SubscriptionsPostRequestBodyable, requestConfiguration *SubscriptionsRequestBuilderPostRequestConfiguration) (SubscriptionsResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateSubscriptions401ErrorFromDiscriminatorValue,
		"403": CreateSubscriptions403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSubscriptionsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(SubscriptionsResponseable), nil
}

// PostAsSubscriptionsPostResponse creates a subscription with the specified product and product subscription id.
// returns a SubscriptionsPostResponseable when successful
// returns a Subscriptions401Error error when the service returns a 401 status code
// returns a Subscriptions403Error error when the service returns a 403 status code
func (m *SubscriptionsRequestBuilder) PostAsSubscriptionsPostResponse(ctx context.Context, body SubscriptionsPostRequestBodyable, requestConfiguration *SubscriptionsRequestBuilderPostRequestConfiguration) (SubscriptionsPostResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateSubscriptions401ErrorFromDiscriminatorValue,
		"403": CreateSubscriptions403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSubscriptionsPostResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(SubscriptionsPostResponseable), nil
}

// ToGetRequestInformation retrieves a list of all subcriptions using the specified search criteria.
// returns a *RequestInformation when successful
func (m *SubscriptionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SubscriptionsRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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

// ToPostRequestInformation creates a subscription with the specified product and product subscription id.
// returns a *RequestInformation when successful
func (m *SubscriptionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body SubscriptionsPostRequestBodyable, requestConfiguration *SubscriptionsRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/v1/subscriptions", m.BaseRequestBuilder.PathParameters)
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
// returns a *SubscriptionsRequestBuilder when successful
func (m *SubscriptionsRequestBuilder) WithUrl(rawUrl string) *SubscriptionsRequestBuilder {
	return NewSubscriptionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
