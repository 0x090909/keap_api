package v2

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// LocalesCountriesWithCountry_codeItemRequestBuilder builds and executes requests for operations under \v2\locales\countries\{country_code}
type LocalesCountriesWithCountry_codeItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// NewLocalesCountriesWithCountry_codeItemRequestBuilderInternal instantiates a new LocalesCountriesWithCountry_codeItemRequestBuilder and sets the default values.
func NewLocalesCountriesWithCountry_codeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *LocalesCountriesWithCountry_codeItemRequestBuilder {
	m := &LocalesCountriesWithCountry_codeItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/locales/countries/{country_code}", pathParameters),
	}
	return m
}

// NewLocalesCountriesWithCountry_codeItemRequestBuilder instantiates a new LocalesCountriesWithCountry_codeItemRequestBuilder and sets the default values.
func NewLocalesCountriesWithCountry_codeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *LocalesCountriesWithCountry_codeItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewLocalesCountriesWithCountry_codeItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Provinces the provinces property
// returns a *LocalesCountriesItemProvincesRequestBuilder when successful
func (m *LocalesCountriesWithCountry_codeItemRequestBuilder) Provinces() *LocalesCountriesItemProvincesRequestBuilder {
	return NewLocalesCountriesItemProvincesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
