package v1

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// LocalesCountriesWithCountryCodeItemRequestBuilder builds and executes requests for operations under \v1\locales\countries\{countryCode}
type LocalesCountriesWithCountryCodeItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// NewLocalesCountriesWithCountryCodeItemRequestBuilderInternal instantiates a new LocalesCountriesWithCountryCodeItemRequestBuilder and sets the default values.
func NewLocalesCountriesWithCountryCodeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *LocalesCountriesWithCountryCodeItemRequestBuilder {
	m := &LocalesCountriesWithCountryCodeItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/locales/countries/{countryCode}", pathParameters),
	}
	return m
}

// NewLocalesCountriesWithCountryCodeItemRequestBuilder instantiates a new LocalesCountriesWithCountryCodeItemRequestBuilder and sets the default values.
func NewLocalesCountriesWithCountryCodeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *LocalesCountriesWithCountryCodeItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewLocalesCountriesWithCountryCodeItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Provinces the provinces property
// returns a *LocalesCountriesItemProvincesRequestBuilder when successful
func (m *LocalesCountriesWithCountryCodeItemRequestBuilder) Provinces() *LocalesCountriesItemProvincesRequestBuilder {
	return NewLocalesCountriesItemProvincesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
