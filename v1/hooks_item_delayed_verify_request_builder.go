package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// HooksItemDelayedVerifyRequestBuilder builds and executes requests for operations under \v1\hooks\{key}\delayedVerify
type HooksItemDelayedVerifyRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// HooksItemDelayedVerifyRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HooksItemDelayedVerifyRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewHooksItemDelayedVerifyRequestBuilderInternal instantiates a new HooksItemDelayedVerifyRequestBuilder and sets the default values.
func NewHooksItemDelayedVerifyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *HooksItemDelayedVerifyRequestBuilder {
	m := &HooksItemDelayedVerifyRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/hooks/{key}/delayedVerify", pathParameters),
	}
	return m
}

// NewHooksItemDelayedVerifyRequestBuilder instantiates a new HooksItemDelayedVerifyRequestBuilder and sets the default values.
func NewHooksItemDelayedVerifyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *HooksItemDelayedVerifyRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewHooksItemDelayedVerifyRequestBuilderInternal(urlParams, requestAdapter)
}

// Post this operation is used to verify or reactivate a hook subscription out-of-band using RESTHooks.org's [Delayed Confirmation](http://resthooks.org/docs/security/) pattern.Use this verification method if you're not able to use the Immediate Confirmation provided through [Create a Hook Subscription](#!/REST_Hooks/create_a_hook_subscription) or [Verify a Hook Subscription](#!/REST_Hooks/verify_a_hook_subscription). This operation allows you to confirm a subscription by manually sending us the `X-Hook-Secret` you received.NB: **The `X-Hook-Secret` _must_ be passed as a _header_.**Don't worry if verification fails; you can always [request another verification attempt](#!/REST_Hooks/verify_a_hook_subscription).NB: You will not receive events until the subscription is verified.If the verification process seems confusing, head over to [RESTHooks.org](http://resthooks.org/docs/security/) for more on the concept.
// Deprecated: This method is obsolete. Use PostAsDelayedVerifyPostResponse instead.
// returns a HooksItemDelayedVerifyResponseable when successful
// returns a HooksItemDelayedVerify401Error error when the service returns a 401 status code
// returns a HooksItemDelayedVerify403Error error when the service returns a 403 status code
func (m *HooksItemDelayedVerifyRequestBuilder) Post(ctx context.Context, requestConfiguration *HooksItemDelayedVerifyRequestBuilderPostRequestConfiguration) (HooksItemDelayedVerifyResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateHooksItemDelayedVerify401ErrorFromDiscriminatorValue,
		"403": CreateHooksItemDelayedVerify403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateHooksItemDelayedVerifyResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(HooksItemDelayedVerifyResponseable), nil
}

// PostAsDelayedVerifyPostResponse this operation is used to verify or reactivate a hook subscription out-of-band using RESTHooks.org's [Delayed Confirmation](http://resthooks.org/docs/security/) pattern.Use this verification method if you're not able to use the Immediate Confirmation provided through [Create a Hook Subscription](#!/REST_Hooks/create_a_hook_subscription) or [Verify a Hook Subscription](#!/REST_Hooks/verify_a_hook_subscription). This operation allows you to confirm a subscription by manually sending us the `X-Hook-Secret` you received.NB: **The `X-Hook-Secret` _must_ be passed as a _header_.**Don't worry if verification fails; you can always [request another verification attempt](#!/REST_Hooks/verify_a_hook_subscription).NB: You will not receive events until the subscription is verified.If the verification process seems confusing, head over to [RESTHooks.org](http://resthooks.org/docs/security/) for more on the concept.
// returns a HooksItemDelayedVerifyPostResponseable when successful
// returns a HooksItemDelayedVerify401Error error when the service returns a 401 status code
// returns a HooksItemDelayedVerify403Error error when the service returns a 403 status code
func (m *HooksItemDelayedVerifyRequestBuilder) PostAsDelayedVerifyPostResponse(ctx context.Context, requestConfiguration *HooksItemDelayedVerifyRequestBuilderPostRequestConfiguration) (HooksItemDelayedVerifyPostResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateHooksItemDelayedVerify401ErrorFromDiscriminatorValue,
		"403": CreateHooksItemDelayedVerify403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateHooksItemDelayedVerifyPostResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(HooksItemDelayedVerifyPostResponseable), nil
}

// ToPostRequestInformation this operation is used to verify or reactivate a hook subscription out-of-band using RESTHooks.org's [Delayed Confirmation](http://resthooks.org/docs/security/) pattern.Use this verification method if you're not able to use the Immediate Confirmation provided through [Create a Hook Subscription](#!/REST_Hooks/create_a_hook_subscription) or [Verify a Hook Subscription](#!/REST_Hooks/verify_a_hook_subscription). This operation allows you to confirm a subscription by manually sending us the `X-Hook-Secret` you received.NB: **The `X-Hook-Secret` _must_ be passed as a _header_.**Don't worry if verification fails; you can always [request another verification attempt](#!/REST_Hooks/verify_a_hook_subscription).NB: You will not receive events until the subscription is verified.If the verification process seems confusing, head over to [RESTHooks.org](http://resthooks.org/docs/security/) for more on the concept.
// returns a *RequestInformation when successful
func (m *HooksItemDelayedVerifyRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *HooksItemDelayedVerifyRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *HooksItemDelayedVerifyRequestBuilder when successful
func (m *HooksItemDelayedVerifyRequestBuilder) WithUrl(rawUrl string) *HooksItemDelayedVerifyRequestBuilder {
	return NewHooksItemDelayedVerifyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
