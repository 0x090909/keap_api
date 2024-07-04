package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// HooksRequestBuilder builds and executes requests for operations under \v1\hooks
type HooksRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// HooksRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HooksRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// HooksRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HooksRequestBuilderPostRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ByKey gets an item from the github.com/0x090909/keap_api.v1.hooks.item collection
// returns a *HooksWithKeyItemRequestBuilder when successful
func (m *HooksRequestBuilder) ByKey(key string) *HooksWithKeyItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if key != "" {
		urlTplParams["key"] = key
	}
	return NewHooksWithKeyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewHooksRequestBuilderInternal instantiates a new HooksRequestBuilder and sets the default values.
func NewHooksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *HooksRequestBuilder {
	m := &HooksRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/hooks", pathParameters),
	}
	return m
}

// NewHooksRequestBuilder instantiates a new HooksRequestBuilder and sets the default values.
func NewHooksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *HooksRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewHooksRequestBuilderInternal(urlParams, requestAdapter)
}

// Event_keys the event_keys property
// returns a *HooksEvent_keysRequestBuilder when successful
func (m *HooksRequestBuilder) Event_keys() *HooksEvent_keysRequestBuilder {
	return NewHooksEvent_keysRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Get lists your hook subscriptions.
// returns a []Hooksable when successful
// returns a Hooks401Error error when the service returns a 401 status code
// returns a Hooks403Error error when the service returns a 403 status code
// returns a Hooks404Error error when the service returns a 404 status code
func (m *HooksRequestBuilder) Get(ctx context.Context, requestConfiguration *HooksRequestBuilderGetRequestConfiguration) ([]Hooksable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateHooks401ErrorFromDiscriminatorValue,
		"403": CreateHooks403ErrorFromDiscriminatorValue,
		"404": CreateHooks404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.SendCollection(ctx, requestInfo, CreateHooksFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	val := make([]Hooksable, len(res))
	for i, v := range res {
		if v != nil {
			val[i] = v.(Hooksable)
		}
	}
	return val, nil
}

// Post to receive hooks, you'll first need to subscribe to events one at a time *and* individually verify each subscription.This operation is used to create hook subscriptions. During this process, Infusionsoft will attempt to verify your subscription. Continue reading to learn how that works.To verify or reactivate a hook subscription, Infusionsoft will make a `POST` request with a temporary secret to the `hookUrl` you provided during creation. The secret is passed as the value of a header named `X-Hook-Secret`. Your response to that `POST` *must* have a status code of `200` and return the same `X-Hook-Secret` header and value pair. Once you've done that, you'll begin receiving hooks. Don't worry if verification fails; you can always [request another verification attempt](#!/REST_Hooks/verify_a_hook_subscription).NB: You will not receive events until the subscription is verified.If the verification process seems confusing, head over to [RESTHooks.org](http://resthooks.org/docs/security/) for more on the concept.
// Deprecated: This method is obsolete. Use PostAsHooksPostResponse instead.
// returns a HooksResponseable when successful
// returns a Hooks401Error error when the service returns a 401 status code
// returns a Hooks403Error error when the service returns a 403 status code
func (m *HooksRequestBuilder) Post(ctx context.Context, body HooksPostRequestBodyable, requestConfiguration *HooksRequestBuilderPostRequestConfiguration) (HooksResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateHooks401ErrorFromDiscriminatorValue,
		"403": CreateHooks403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateHooksResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(HooksResponseable), nil
}

// PostAsHooksPostResponse to receive hooks, you'll first need to subscribe to events one at a time *and* individually verify each subscription.This operation is used to create hook subscriptions. During this process, Infusionsoft will attempt to verify your subscription. Continue reading to learn how that works.To verify or reactivate a hook subscription, Infusionsoft will make a `POST` request with a temporary secret to the `hookUrl` you provided during creation. The secret is passed as the value of a header named `X-Hook-Secret`. Your response to that `POST` *must* have a status code of `200` and return the same `X-Hook-Secret` header and value pair. Once you've done that, you'll begin receiving hooks. Don't worry if verification fails; you can always [request another verification attempt](#!/REST_Hooks/verify_a_hook_subscription).NB: You will not receive events until the subscription is verified.If the verification process seems confusing, head over to [RESTHooks.org](http://resthooks.org/docs/security/) for more on the concept.
// returns a HooksPostResponseable when successful
// returns a Hooks401Error error when the service returns a 401 status code
// returns a Hooks403Error error when the service returns a 403 status code
func (m *HooksRequestBuilder) PostAsHooksPostResponse(ctx context.Context, body HooksPostRequestBodyable, requestConfiguration *HooksRequestBuilderPostRequestConfiguration) (HooksPostResponseable, error) {
	requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateHooks401ErrorFromDiscriminatorValue,
		"403": CreateHooks403ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateHooksPostResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(HooksPostResponseable), nil
}

// ToGetRequestInformation lists your hook subscriptions.
// returns a *RequestInformation when successful
func (m *HooksRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *HooksRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// ToPostRequestInformation to receive hooks, you'll first need to subscribe to events one at a time *and* individually verify each subscription.This operation is used to create hook subscriptions. During this process, Infusionsoft will attempt to verify your subscription. Continue reading to learn how that works.To verify or reactivate a hook subscription, Infusionsoft will make a `POST` request with a temporary secret to the `hookUrl` you provided during creation. The secret is passed as the value of a header named `X-Hook-Secret`. Your response to that `POST` *must* have a status code of `200` and return the same `X-Hook-Secret` header and value pair. Once you've done that, you'll begin receiving hooks. Don't worry if verification fails; you can always [request another verification attempt](#!/REST_Hooks/verify_a_hook_subscription).NB: You will not receive events until the subscription is verified.If the verification process seems confusing, head over to [RESTHooks.org](http://resthooks.org/docs/security/) for more on the concept.
// returns a *RequestInformation when successful
func (m *HooksRequestBuilder) ToPostRequestInformation(ctx context.Context, body HooksPostRequestBodyable, requestConfiguration *HooksRequestBuilderPostRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *HooksRequestBuilder when successful
func (m *HooksRequestBuilder) WithUrl(rawUrl string) *HooksRequestBuilder {
	return NewHooksRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
