package v2

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// LandingpagesRequestBuilder builds and executes requests for operations under \v2\landingpages
type LandingpagesRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// NewLandingpagesRequestBuilderInternal instantiates a new LandingpagesRequestBuilder and sets the default values.
func NewLandingpagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *LandingpagesRequestBuilder {
	m := &LandingpagesRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/landingpages", pathParameters),
	}
	return m
}

// NewLandingpagesRequestBuilder instantiates a new LandingpagesRequestBuilder and sets the default values.
func NewLandingpagesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *LandingpagesRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewLandingpagesRequestBuilderInternal(urlParams, requestAdapter)
}
