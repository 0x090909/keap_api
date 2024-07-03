package v2

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// LeadsourcesRequestBuilder builds and executes requests for operations under \v2\leadsources
type LeadsourcesRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// NewLeadsourcesRequestBuilderInternal instantiates a new LeadsourcesRequestBuilder and sets the default values.
func NewLeadsourcesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *LeadsourcesRequestBuilder {
	m := &LeadsourcesRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/leadsources", pathParameters),
	}
	return m
}

// NewLeadsourcesRequestBuilder instantiates a new LeadsourcesRequestBuilder and sets the default values.
func NewLeadsourcesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *LeadsourcesRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewLeadsourcesRequestBuilderInternal(urlParams, requestAdapter)
}
