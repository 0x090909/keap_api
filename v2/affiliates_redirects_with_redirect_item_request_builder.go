package v2

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AffiliatesRedirectsWithRedirect_ItemRequestBuilder builds and executes requests for operations under \v2\affiliates\redirects\{redirect_id}
type AffiliatesRedirectsWithRedirect_ItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// NewAffiliatesRedirectsWithRedirect_ItemRequestBuilderInternal instantiates a new AffiliatesRedirectsWithRedirect_ItemRequestBuilder and sets the default values.
func NewAffiliatesRedirectsWithRedirect_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesRedirectsWithRedirect_ItemRequestBuilder {
	m := &AffiliatesRedirectsWithRedirect_ItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/affiliates/redirects/{redirect_id}", pathParameters),
	}
	return m
}

// NewAffiliatesRedirectsWithRedirect_ItemRequestBuilder instantiates a new AffiliatesRedirectsWithRedirect_ItemRequestBuilder and sets the default values.
func NewAffiliatesRedirectsWithRedirect_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesRedirectsWithRedirect_ItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewAffiliatesRedirectsWithRedirect_ItemRequestBuilderInternal(urlParams, requestAdapter)
}
