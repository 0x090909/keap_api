package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// MerchantsRequestBuilder builds and executes requests for operations under \v1\merchants
type MerchantsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// MerchantsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MerchantsRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewMerchantsRequestBuilderInternal instantiates a new MerchantsRequestBuilder and sets the default values.
func NewMerchantsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *MerchantsRequestBuilder {
	m := &MerchantsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/merchants", pathParameters),
	}
	return m
}

// NewMerchantsRequestBuilder instantiates a new MerchantsRequestBuilder and sets the default values.
func NewMerchantsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *MerchantsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewMerchantsRequestBuilderInternal(urlParams, requestAdapter)
}

// Get retrieves a list of all merchant accounts
// Deprecated: This method is obsolete. Use GetAsMerchantsGetResponse instead.
// returns a MerchantsResponseable when successful
// returns a Merchants401Error error when the service returns a 401 status code
// returns a Merchants403Error error when the service returns a 403 status code
// returns a Merchants404Error error when the service returns a 404 status code
func (m *MerchantsRequestBuilder) Get(ctx context.Context, requestConfiguration *MerchantsRequestBuilderGetRequestConfiguration) (MerchantsResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateMerchants401ErrorFromDiscriminatorValue,
		"403": CreateMerchants403ErrorFromDiscriminatorValue,
		"404": CreateMerchants404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateMerchantsResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(MerchantsResponseable), nil
}

// GetAsMerchantsGetResponse retrieves a list of all merchant accounts
// returns a MerchantsGetResponseable when successful
// returns a Merchants401Error error when the service returns a 401 status code
// returns a Merchants403Error error when the service returns a 403 status code
// returns a Merchants404Error error when the service returns a 404 status code
func (m *MerchantsRequestBuilder) GetAsMerchantsGetResponse(ctx context.Context, requestConfiguration *MerchantsRequestBuilderGetRequestConfiguration) (MerchantsGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateMerchants401ErrorFromDiscriminatorValue,
		"403": CreateMerchants403ErrorFromDiscriminatorValue,
		"404": CreateMerchants404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateMerchantsGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(MerchantsGetResponseable), nil
}

// ToGetRequestInformation retrieves a list of all merchant accounts
// returns a *RequestInformation when successful
func (m *MerchantsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MerchantsRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *MerchantsRequestBuilder when successful
func (m *MerchantsRequestBuilder) WithUrl(rawUrl string) *MerchantsRequestBuilder {
	return NewMerchantsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
