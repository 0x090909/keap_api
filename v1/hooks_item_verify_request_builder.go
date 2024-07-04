package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// HooksItemVerifyRequestBuilder builds and executes requests for operations under \v1\hooks\{key}\verify
type HooksItemVerifyRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// HooksItemVerifyRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HooksItemVerifyRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewHooksItemVerifyRequestBuilderInternal instantiates a new HooksItemVerifyRequestBuilder and sets the default values.
func NewHooksItemVerifyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *HooksItemVerifyRequestBuilder {
	m := &HooksItemVerifyRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/hooks/{key}/verify", pathParameters),
	}
	return m
}

// NewHooksItemVerifyRequestBuilder instantiates a new HooksItemVerifyRequestBuilder and sets the default values.
func NewHooksItemVerifyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *HooksItemVerifyRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewHooksItemVerifyRequestBuilderInternal(urlParams, requestAdapter)
}

// Post this operation is used to verify or reactivate a hook subscription using RESTHooks.org's [Immediate Confirmation](http://resthooks.org/docs/security/) pattern.To verify or reactivate a hook subscription, Infusionsoft will make a `POST` request with a temporary secret to the `hookUrl` you provided during creation. The secret is passed as the value of a header named `X-Hook-Secret`. Your response to that `POST` *must* have a status code of `200` and return the same `X-Hook-Secret` header and value pair. Once you've done that, you'll begin receiving hooks. Don't worry if verification fails; you can always [request another verification attempt](#!/REST_Hooks/verify_a_hook_subscription).NB: You will not receive events until the subscription is verified.If the verification process seems confusing, head over to [RESTHooks.org](http://resthooks.org/docs/security/) for more on the concept.
// Deprecated: This method is obsolete. Use PostAsVerifyPostResponse instead.
// returns a HooksItemVerifyResponseable when successful
// returns a HooksItemVerify401Error error when the service returns a 401 status code
// returns a HooksItemVerify403Error error when the service returns a 403 status code
func (m *HooksItemVerifyRequestBuilder) Post(ctx context.Context, requestConfiguration *HooksItemVerifyRequestBuilderPostRequestConfiguration) (HooksItemVerifyResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateHooksItemVerify401ErrorFromDiscriminatorValue,
		"403": CreateHooksItemVerify403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateHooksItemVerifyResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(HooksItemVerifyResponseable), nil
}

// PostAsVerifyPostResponse this operation is used to verify or reactivate a hook subscription using RESTHooks.org's [Immediate Confirmation](http://resthooks.org/docs/security/) pattern.To verify or reactivate a hook subscription, Infusionsoft will make a `POST` request with a temporary secret to the `hookUrl` you provided during creation. The secret is passed as the value of a header named `X-Hook-Secret`. Your response to that `POST` *must* have a status code of `200` and return the same `X-Hook-Secret` header and value pair. Once you've done that, you'll begin receiving hooks. Don't worry if verification fails; you can always [request another verification attempt](#!/REST_Hooks/verify_a_hook_subscription).NB: You will not receive events until the subscription is verified.If the verification process seems confusing, head over to [RESTHooks.org](http://resthooks.org/docs/security/) for more on the concept.
// returns a HooksItemVerifyPostResponseable when successful
// returns a HooksItemVerify401Error error when the service returns a 401 status code
// returns a HooksItemVerify403Error error when the service returns a 403 status code
func (m *HooksItemVerifyRequestBuilder) PostAsVerifyPostResponse(ctx context.Context, requestConfiguration *HooksItemVerifyRequestBuilderPostRequestConfiguration) (HooksItemVerifyPostResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateHooksItemVerify401ErrorFromDiscriminatorValue,
		"403": CreateHooksItemVerify403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateHooksItemVerifyPostResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(HooksItemVerifyPostResponseable), nil
}

// ToPostRequestInformation this operation is used to verify or reactivate a hook subscription using RESTHooks.org's [Immediate Confirmation](http://resthooks.org/docs/security/) pattern.To verify or reactivate a hook subscription, Infusionsoft will make a `POST` request with a temporary secret to the `hookUrl` you provided during creation. The secret is passed as the value of a header named `X-Hook-Secret`. Your response to that `POST` *must* have a status code of `200` and return the same `X-Hook-Secret` header and value pair. Once you've done that, you'll begin receiving hooks. Don't worry if verification fails; you can always [request another verification attempt](#!/REST_Hooks/verify_a_hook_subscription).NB: You will not receive events until the subscription is verified.If the verification process seems confusing, head over to [RESTHooks.org](http://resthooks.org/docs/security/) for more on the concept.
// returns a *RequestInformation when successful
func (m *HooksItemVerifyRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *HooksItemVerifyRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *HooksItemVerifyRequestBuilder when successful
func (m *HooksItemVerifyRequestBuilder) WithUrl(rawUrl string) *HooksItemVerifyRequestBuilder {
	return NewHooksItemVerifyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
