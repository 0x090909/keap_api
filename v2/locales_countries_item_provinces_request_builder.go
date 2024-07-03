package v2

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// LocalesCountriesItemProvincesRequestBuilder builds and executes requests for operations under \v2\locales\countries\{country_code}\provinces
type LocalesCountriesItemProvincesRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// NewLocalesCountriesItemProvincesRequestBuilderInternal instantiates a new LocalesCountriesItemProvincesRequestBuilder and sets the default values.
func NewLocalesCountriesItemProvincesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *LocalesCountriesItemProvincesRequestBuilder {
	m := &LocalesCountriesItemProvincesRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/locales/countries/{country_code}/provinces", pathParameters),
	}
	return m
}

// NewLocalesCountriesItemProvincesRequestBuilder instantiates a new LocalesCountriesItemProvincesRequestBuilder and sets the default values.
func NewLocalesCountriesItemProvincesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *LocalesCountriesItemProvincesRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewLocalesCountriesItemProvincesRequestBuilderInternal(urlParams, requestAdapter)
}
