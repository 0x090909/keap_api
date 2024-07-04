package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// LocalesCountriesItemProvincesRequestBuilder builds and executes requests for operations under \v1\locales\countries\{countryCode}\provinces
type LocalesCountriesItemProvincesRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// LocalesCountriesItemProvincesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LocalesCountriesItemProvincesRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// NewLocalesCountriesItemProvincesRequestBuilderInternal instantiates a new LocalesCountriesItemProvincesRequestBuilder and sets the default values.
func NewLocalesCountriesItemProvincesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *LocalesCountriesItemProvincesRequestBuilder {
	m := &LocalesCountriesItemProvincesRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/locales/countries/{countryCode}/provinces", pathParameters),
	}
	return m
}

// NewLocalesCountriesItemProvincesRequestBuilder instantiates a new LocalesCountriesItemProvincesRequestBuilder and sets the default values.
func NewLocalesCountriesItemProvincesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *LocalesCountriesItemProvincesRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewLocalesCountriesItemProvincesRequestBuilderInternal(urlParams, requestAdapter)
}

// Get list a Country's Provinces
// Deprecated: This method is obsolete. Use GetAsProvincesGetResponse instead.
// returns a LocalesCountriesItemProvincesResponseable when successful
// returns a LocalesCountriesItemProvinces401Error error when the service returns a 401 status code
// returns a LocalesCountriesItemProvinces403Error error when the service returns a 403 status code
// returns a LocalesCountriesItemProvinces404Error error when the service returns a 404 status code
func (m *LocalesCountriesItemProvincesRequestBuilder) Get(ctx context.Context, requestConfiguration *LocalesCountriesItemProvincesRequestBuilderGetRequestConfiguration) (LocalesCountriesItemProvincesResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateLocalesCountriesItemProvinces401ErrorFromDiscriminatorValue,
		"403": CreateLocalesCountriesItemProvinces403ErrorFromDiscriminatorValue,
		"404": CreateLocalesCountriesItemProvinces404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateLocalesCountriesItemProvincesResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(LocalesCountriesItemProvincesResponseable), nil
}

// GetAsProvincesGetResponse list a Country's Provinces
// returns a LocalesCountriesItemProvincesGetResponseable when successful
// returns a LocalesCountriesItemProvinces401Error error when the service returns a 401 status code
// returns a LocalesCountriesItemProvinces403Error error when the service returns a 403 status code
// returns a LocalesCountriesItemProvinces404Error error when the service returns a 404 status code
func (m *LocalesCountriesItemProvincesRequestBuilder) GetAsProvincesGetResponse(ctx context.Context, requestConfiguration *LocalesCountriesItemProvincesRequestBuilderGetRequestConfiguration) (LocalesCountriesItemProvincesGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateLocalesCountriesItemProvinces401ErrorFromDiscriminatorValue,
		"403": CreateLocalesCountriesItemProvinces403ErrorFromDiscriminatorValue,
		"404": CreateLocalesCountriesItemProvinces404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateLocalesCountriesItemProvincesGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(LocalesCountriesItemProvincesGetResponseable), nil
}

// ToGetRequestInformation list a Country's Provinces
// returns a *RequestInformation when successful
func (m *LocalesCountriesItemProvincesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LocalesCountriesItemProvincesRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *LocalesCountriesItemProvincesRequestBuilder when successful
func (m *LocalesCountriesItemProvincesRequestBuilder) WithUrl(rawUrl string) *LocalesCountriesItemProvincesRequestBuilder {
	return NewLocalesCountriesItemProvincesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
