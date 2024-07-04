package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SubscriptionsModelRequestBuilder builds and executes requests for operations under \v1\subscriptions\model
type SubscriptionsModelRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// SubscriptionsModelRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SubscriptionsModelRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewSubscriptionsModelRequestBuilderInternal instantiates a new SubscriptionsModelRequestBuilder and sets the default values.
func NewSubscriptionsModelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *SubscriptionsModelRequestBuilder {
	m := &SubscriptionsModelRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/subscriptions/model", pathParameters),
	}
	return m
}

// NewSubscriptionsModelRequestBuilder instantiates a new SubscriptionsModelRequestBuilder and sets the default values.
func NewSubscriptionsModelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *SubscriptionsModelRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewSubscriptionsModelRequestBuilderInternal(urlParams, requestAdapter)
}

// Get get the custom fields for the Subscription object
// Deprecated: This method is obsolete. Use GetAsModelGetResponse instead.
// returns a SubscriptionsModelResponseable when successful
// returns a SubscriptionsModel401Error error when the service returns a 401 status code
// returns a SubscriptionsModel403Error error when the service returns a 403 status code
// returns a SubscriptionsModel404Error error when the service returns a 404 status code
func (m *SubscriptionsModelRequestBuilder) Get(ctx context.Context, requestConfiguration *SubscriptionsModelRequestBuilderGetRequestConfiguration) (SubscriptionsModelResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateSubscriptionsModel401ErrorFromDiscriminatorValue,
		"403": CreateSubscriptionsModel403ErrorFromDiscriminatorValue,
		"404": CreateSubscriptionsModel404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSubscriptionsModelResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(SubscriptionsModelResponseable), nil
}

// GetAsModelGetResponse get the custom fields for the Subscription object
// returns a SubscriptionsModelGetResponseable when successful
// returns a SubscriptionsModel401Error error when the service returns a 401 status code
// returns a SubscriptionsModel403Error error when the service returns a 403 status code
// returns a SubscriptionsModel404Error error when the service returns a 404 status code
func (m *SubscriptionsModelRequestBuilder) GetAsModelGetResponse(ctx context.Context, requestConfiguration *SubscriptionsModelRequestBuilderGetRequestConfiguration) (SubscriptionsModelGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateSubscriptionsModel401ErrorFromDiscriminatorValue,
		"403": CreateSubscriptionsModel403ErrorFromDiscriminatorValue,
		"404": CreateSubscriptionsModel404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateSubscriptionsModelGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(SubscriptionsModelGetResponseable), nil
}

// ToGetRequestInformation get the custom fields for the Subscription object
// returns a *RequestInformation when successful
func (m *SubscriptionsModelRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SubscriptionsModelRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *SubscriptionsModelRequestBuilder when successful
func (m *SubscriptionsModelRequestBuilder) WithUrl(rawUrl string) *SubscriptionsModelRequestBuilder {
	return NewSubscriptionsModelRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
