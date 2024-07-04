package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// HooksWithKeyItemRequestBuilder builds and executes requests for operations under \v1\hooks\{key}
type HooksWithKeyItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// HooksWithKeyItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HooksWithKeyItemRequestBuilderDeleteRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// HooksWithKeyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HooksWithKeyItemRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// HooksWithKeyItemRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HooksWithKeyItemRequestBuilderPutRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewHooksWithKeyItemRequestBuilderInternal instantiates a new HooksWithKeyItemRequestBuilder and sets the default values.
func NewHooksWithKeyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *HooksWithKeyItemRequestBuilder {
	m := &HooksWithKeyItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/hooks/{key}", pathParameters),
	}
	return m
}

// NewHooksWithKeyItemRequestBuilder instantiates a new HooksWithKeyItemRequestBuilder and sets the default values.
func NewHooksWithKeyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *HooksWithKeyItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewHooksWithKeyItemRequestBuilderInternal(urlParams, requestAdapter)
}

// DelayedVerify the delayedVerify property
// returns a *HooksItemDelayedVerifyRequestBuilder when successful
func (m *HooksWithKeyItemRequestBuilder) DelayedVerify() *HooksItemDelayedVerifyRequestBuilder {
	return NewHooksItemDelayedVerifyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// Delete stop receiving hooks by deleting an existing hook subscription.
// returns a HooksItemWithKey401Error error when the service returns a 401 status code
// returns a HooksItemWithKey403Error error when the service returns a 403 status code
// returns a HooksItemWithKey404Error error when the service returns a 404 status code
func (m *HooksWithKeyItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *HooksWithKeyItemRequestBuilderDeleteRequestConfiguration) error {
	requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateHooksItemWithKey401ErrorFromDiscriminatorValue,
		"403": CreateHooksItemWithKey403ErrorFromDiscriminatorValue,
		"404": CreateHooksItemWithKey404ErrorFromDiscriminatorValue,
	}
	err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
	if err != nil {
		return err
	}
	return nil
}

// Get retrieves an existing hook subscription and its status.If your hook subscription becomes inactive, you may request an activation attempt via [Verify a Hook Subscription](#!/REST_Hooks/verify_a_hook_subscription).
// Deprecated: This method is obsolete. Use GetAsWithKeyGetResponse instead.
// returns a HooksItemWithKeyResponseable when successful
// returns a HooksItemWithKey401Error error when the service returns a 401 status code
// returns a HooksItemWithKey403Error error when the service returns a 403 status code
// returns a HooksItemWithKey404Error error when the service returns a 404 status code
func (m *HooksWithKeyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *HooksWithKeyItemRequestBuilderGetRequestConfiguration) (HooksItemWithKeyResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateHooksItemWithKey401ErrorFromDiscriminatorValue,
		"403": CreateHooksItemWithKey403ErrorFromDiscriminatorValue,
		"404": CreateHooksItemWithKey404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateHooksItemWithKeyResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(HooksItemWithKeyResponseable), nil
}

// GetAsWithKeyGetResponse retrieves an existing hook subscription and its status.If your hook subscription becomes inactive, you may request an activation attempt via [Verify a Hook Subscription](#!/REST_Hooks/verify_a_hook_subscription).
// returns a HooksItemWithKeyGetResponseable when successful
// returns a HooksItemWithKey401Error error when the service returns a 401 status code
// returns a HooksItemWithKey403Error error when the service returns a 403 status code
// returns a HooksItemWithKey404Error error when the service returns a 404 status code
func (m *HooksWithKeyItemRequestBuilder) GetAsWithKeyGetResponse(ctx context.Context, requestConfiguration *HooksWithKeyItemRequestBuilderGetRequestConfiguration) (HooksItemWithKeyGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateHooksItemWithKey401ErrorFromDiscriminatorValue,
		"403": CreateHooksItemWithKey403ErrorFromDiscriminatorValue,
		"404": CreateHooksItemWithKey404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateHooksItemWithKeyGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(HooksItemWithKeyGetResponseable), nil
}

// Put modify an existing hook subscription using the provided values.
// Deprecated: This method is obsolete. Use PutAsWithKeyPutResponse instead.
// returns a HooksItemWithKeyResponseable when successful
// returns a HooksItemWithKey401Error error when the service returns a 401 status code
// returns a HooksItemWithKey403Error error when the service returns a 403 status code
// returns a HooksItemWithKey404Error error when the service returns a 404 status code
func (m *HooksWithKeyItemRequestBuilder) Put(ctx context.Context, body HooksItemWithKeyPutRequestBodyable, requestConfiguration *HooksWithKeyItemRequestBuilderPutRequestConfiguration) (HooksItemWithKeyResponseable, error) {
	requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateHooksItemWithKey401ErrorFromDiscriminatorValue,
		"403": CreateHooksItemWithKey403ErrorFromDiscriminatorValue,
		"404": CreateHooksItemWithKey404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateHooksItemWithKeyResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(HooksItemWithKeyResponseable), nil
}

// PutAsWithKeyPutResponse modify an existing hook subscription using the provided values.
// returns a HooksItemWithKeyPutResponseable when successful
// returns a HooksItemWithKey401Error error when the service returns a 401 status code
// returns a HooksItemWithKey403Error error when the service returns a 403 status code
// returns a HooksItemWithKey404Error error when the service returns a 404 status code
func (m *HooksWithKeyItemRequestBuilder) PutAsWithKeyPutResponse(ctx context.Context, body HooksItemWithKeyPutRequestBodyable, requestConfiguration *HooksWithKeyItemRequestBuilderPutRequestConfiguration) (HooksItemWithKeyPutResponseable, error) {
	requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateHooksItemWithKey401ErrorFromDiscriminatorValue,
		"403": CreateHooksItemWithKey403ErrorFromDiscriminatorValue,
		"404": CreateHooksItemWithKey404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateHooksItemWithKeyPutResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(HooksItemWithKeyPutResponseable), nil
}

// ToDeleteRequestInformation stop receiving hooks by deleting an existing hook subscription.
// returns a *RequestInformation when successful
func (m *HooksWithKeyItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *HooksWithKeyItemRequestBuilderDeleteRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// ToGetRequestInformation retrieves an existing hook subscription and its status.If your hook subscription becomes inactive, you may request an activation attempt via [Verify a Hook Subscription](#!/REST_Hooks/verify_a_hook_subscription).
// returns a *RequestInformation when successful
func (m *HooksWithKeyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *HooksWithKeyItemRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// ToPutRequestInformation modify an existing hook subscription using the provided values.
// returns a *RequestInformation when successful
func (m *HooksWithKeyItemRequestBuilder) ToPutRequestInformation(ctx context.Context, body HooksItemWithKeyPutRequestBodyable, requestConfiguration *HooksWithKeyItemRequestBuilderPutRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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

// Verify the verify property
// returns a *HooksItemVerifyRequestBuilder when successful
func (m *HooksWithKeyItemRequestBuilder) Verify() *HooksItemVerifyRequestBuilder {
	return NewHooksItemVerifyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *HooksWithKeyItemRequestBuilder when successful
func (m *HooksWithKeyItemRequestBuilder) WithUrl(rawUrl string) *HooksWithKeyItemRequestBuilder {
	return NewHooksWithKeyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
