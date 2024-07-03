package v2

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// PaymentMethodConfigsRequestBuilder builds and executes requests for operations under \v2\paymentMethodConfigs
type PaymentMethodConfigsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// PaymentMethodConfigsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PaymentMethodConfigsRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewPaymentMethodConfigsRequestBuilderInternal instantiates a new PaymentMethodConfigsRequestBuilder and sets the default values.
func NewPaymentMethodConfigsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *PaymentMethodConfigsRequestBuilder {
	m := &PaymentMethodConfigsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/paymentMethodConfigs", pathParameters),
	}
	return m
}

// NewPaymentMethodConfigsRequestBuilder instantiates a new PaymentMethodConfigsRequestBuilder and sets the default values.
func NewPaymentMethodConfigsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *PaymentMethodConfigsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewPaymentMethodConfigsRequestBuilderInternal(urlParams, requestAdapter)
}

// Post creates configuration details for rendering payment method components
// Deprecated: This method is obsolete. Use PostAsPaymentMethodConfigsPostResponse instead.
// returns a PaymentMethodConfigsResponseable when successful
// returns a PaymentMethodConfigs401Error error when the service returns a 401 status code
// returns a PaymentMethodConfigs403Error error when the service returns a 403 status code
func (m *PaymentMethodConfigsRequestBuilder) Post(ctx context.Context, body PaymentMethodConfigsPostRequestBodyable, requestConfiguration *PaymentMethodConfigsRequestBuilderPostRequestConfiguration) (PaymentMethodConfigsResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreatePaymentMethodConfigs401ErrorFromDiscriminatorValue,
		"403": CreatePaymentMethodConfigs403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreatePaymentMethodConfigsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(PaymentMethodConfigsResponseable), nil
}

// PostAsPaymentMethodConfigsPostResponse creates configuration details for rendering payment method components
// returns a PaymentMethodConfigsPostResponseable when successful
// returns a PaymentMethodConfigs401Error error when the service returns a 401 status code
// returns a PaymentMethodConfigs403Error error when the service returns a 403 status code
func (m *PaymentMethodConfigsRequestBuilder) PostAsPaymentMethodConfigsPostResponse(ctx context.Context, body PaymentMethodConfigsPostRequestBodyable, requestConfiguration *PaymentMethodConfigsRequestBuilderPostRequestConfiguration) (PaymentMethodConfigsPostResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreatePaymentMethodConfigs401ErrorFromDiscriminatorValue,
		"403": CreatePaymentMethodConfigs403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreatePaymentMethodConfigsPostResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(PaymentMethodConfigsPostResponseable), nil
}

// ToPostRequestInformation creates configuration details for rendering payment method components
// returns a *RequestInformation when successful
func (m *PaymentMethodConfigsRequestBuilder) ToPostRequestInformation(ctx context.Context, body PaymentMethodConfigsPostRequestBodyable, requestConfiguration *PaymentMethodConfigsRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *PaymentMethodConfigsRequestBuilder when successful
func (m *PaymentMethodConfigsRequestBuilder) WithUrl(rawUrl string) *PaymentMethodConfigsRequestBuilder {
	return NewPaymentMethodConfigsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
