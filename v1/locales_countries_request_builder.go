package v1

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// LocalesCountriesRequestBuilder builds and executes requests for operations under \v1\locales\countries
type LocalesCountriesRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// LocalesCountriesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LocalesCountriesRequestBuilderGetRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}

// ByCountryCode gets an item from the github.com/0x090909/keap_api.v1.locales.countries.item collection
// returns a *LocalesCountriesWithCountryCodeItemRequestBuilder when successful
func (m *LocalesCountriesRequestBuilder) ByCountryCode(countryCode string) *LocalesCountriesWithCountryCodeItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if countryCode != "" {
		urlTplParams["countryCode"] = countryCode
	}
	return NewLocalesCountriesWithCountryCodeItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewLocalesCountriesRequestBuilderInternal instantiates a new LocalesCountriesRequestBuilder and sets the default values.
func NewLocalesCountriesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *LocalesCountriesRequestBuilder {
	m := &LocalesCountriesRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/locales/countries", pathParameters),
	}
	return m
}

// NewLocalesCountriesRequestBuilder instantiates a new LocalesCountriesRequestBuilder and sets the default values.
func NewLocalesCountriesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *LocalesCountriesRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewLocalesCountriesRequestBuilderInternal(urlParams, requestAdapter)
}

// Get list Countries
// Deprecated: This method is obsolete. Use GetAsCountriesGetResponse instead.
// returns a LocalesCountriesResponseable when successful
// returns a LocalesCountries401Error error when the service returns a 401 status code
// returns a LocalesCountries403Error error when the service returns a 403 status code
// returns a LocalesCountries404Error error when the service returns a 404 status code
func (m *LocalesCountriesRequestBuilder) Get(ctx context.Context, requestConfiguration *LocalesCountriesRequestBuilderGetRequestConfiguration) (LocalesCountriesResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateLocalesCountries401ErrorFromDiscriminatorValue,
		"403": CreateLocalesCountries403ErrorFromDiscriminatorValue,
		"404": CreateLocalesCountries404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateLocalesCountriesResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(LocalesCountriesResponseable), nil
}

// GetAsCountriesGetResponse list Countries
// returns a LocalesCountriesGetResponseable when successful
// returns a LocalesCountries401Error error when the service returns a 401 status code
// returns a LocalesCountries403Error error when the service returns a 403 status code
// returns a LocalesCountries404Error error when the service returns a 404 status code
func (m *LocalesCountriesRequestBuilder) GetAsCountriesGetResponse(ctx context.Context, requestConfiguration *LocalesCountriesRequestBuilderGetRequestConfiguration) (LocalesCountriesGetResponseable, error) {
	requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateLocalesCountries401ErrorFromDiscriminatorValue,
		"403": CreateLocalesCountries403ErrorFromDiscriminatorValue,
		"404": CreateLocalesCountries404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateLocalesCountriesGetResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(LocalesCountriesGetResponseable), nil
}

// ToGetRequestInformation list Countries
// returns a *RequestInformation when successful
func (m *LocalesCountriesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LocalesCountriesRequestBuilderGetRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *LocalesCountriesRequestBuilder when successful
func (m *LocalesCountriesRequestBuilder) WithUrl(rawUrl string) *LocalesCountriesRequestBuilder {
	return NewLocalesCountriesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
