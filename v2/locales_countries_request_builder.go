package v2

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// LocalesCountriesRequestBuilder builds and executes requests for operations under \v2\locales\countries
type LocalesCountriesRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ByCountry_code gets an item from the github.com/0x090909/keap_api.v2.locales.countries.item collection
// returns a *LocalesCountriesWithCountry_codeItemRequestBuilder when successful
func (m *LocalesCountriesRequestBuilder) ByCountry_code(country_code string) *LocalesCountriesWithCountry_codeItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if country_code != "" {
		urlTplParams["country_code"] = country_code
	}
	return NewLocalesCountriesWithCountry_codeItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewLocalesCountriesRequestBuilderInternal instantiates a new LocalesCountriesRequestBuilder and sets the default values.
func NewLocalesCountriesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *LocalesCountriesRequestBuilder {
	m := &LocalesCountriesRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/locales/countries", pathParameters),
	}
	return m
}

// NewLocalesCountriesRequestBuilder instantiates a new LocalesCountriesRequestBuilder and sets the default values.
func NewLocalesCountriesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *LocalesCountriesRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewLocalesCountriesRequestBuilderInternal(urlParams, requestAdapter)
}
