package v1

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// LocalesRequestBuilder builds and executes requests for operations under \v1\locales
type LocalesRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// NewLocalesRequestBuilderInternal instantiates a new LocalesRequestBuilder and sets the default values.
func NewLocalesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *LocalesRequestBuilder {
	m := &LocalesRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/locales", pathParameters),
	}
	return m
}

// NewLocalesRequestBuilder instantiates a new LocalesRequestBuilder and sets the default values.
func NewLocalesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *LocalesRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewLocalesRequestBuilderInternal(urlParams, requestAdapter)
}

// Countries the countries property
// returns a *LocalesCountriesRequestBuilder when successful
func (m *LocalesRequestBuilder) Countries() *LocalesCountriesRequestBuilder {
	return NewLocalesCountriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// DefaultOptions the defaultOptions property
// returns a *LocalesDefaultOptionsRequestBuilder when successful
func (m *LocalesRequestBuilder) DefaultOptions() *LocalesDefaultOptionsRequestBuilder {
	return NewLocalesDefaultOptionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
