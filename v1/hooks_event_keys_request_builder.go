package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// HooksEvent_keysRequestBuilder builds and executes requests for operations under \v1\hooks\event_keys
type HooksEvent_keysRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// HooksEvent_keysRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HooksEvent_keysRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewHooksEvent_keysRequestBuilderInternal instantiates a new HooksEvent_keysRequestBuilder and sets the default values.
func NewHooksEvent_keysRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *HooksEvent_keysRequestBuilder {
	m := &HooksEvent_keysRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/hooks/event_keys", pathParameters),
	}
	return m
}

// NewHooksEvent_keysRequestBuilder instantiates a new HooksEvent_keysRequestBuilder and sets the default values.
func NewHooksEvent_keysRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *HooksEvent_keysRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewHooksEvent_keysRequestBuilderInternal(urlParams, requestAdapter)
}

// Get list the available types of Events that can be listened to
// returns a []string when successful
// returns a HooksEvent_keys401Error error when the service returns a 401 status code
// returns a HooksEvent_keys403Error error when the service returns a 403 status code
// returns a HooksEvent_keys404Error error when the service returns a 404 status code
func (m *HooksEvent_keysRequestBuilder) Get(ctx context.Context, requestConfiguration *HooksEvent_keysRequestBuilderGetRequestConfiguration) ([]string, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateHooksEvent_keys401ErrorFromDiscriminatorValue,
		"403": CreateHooksEvent_keys403ErrorFromDiscriminatorValue,
		"404": CreateHooksEvent_keys404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitiveCollection(ctx, requestInfo, "string", errorMapping)
	if err != nil {
		return nil, err
	}
	val := make([]string, len(res))
	for i, v := range res {
		if v != nil {
			val[i] = *(v.(*string))
		}
	}
	return val, nil
}

// ToGetRequestInformation list the available types of Events that can be listened to
// returns a *RequestInformation when successful
func (m *HooksEvent_keysRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *HooksEvent_keysRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *HooksEvent_keysRequestBuilder when successful
func (m *HooksEvent_keysRequestBuilder) WithUrl(rawUrl string) *HooksEvent_keysRequestBuilder {
	return NewHooksEvent_keysRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
