package v2

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AffiliatesRedirectsRequestBuilder builds and executes requests for operations under \v2\affiliates\redirects
type AffiliatesRedirectsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ByRedirect_id gets an item from the github.com/0x090909/keap_api.v2.affiliates.redirects.item collection
// returns a *AffiliatesRedirectsWithRedirect_ItemRequestBuilder when successful
func (m *AffiliatesRedirectsRequestBuilder) ByRedirect_id(redirect_id string) *AffiliatesRedirectsWithRedirect_ItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if redirect_id != "" {
		urlTplParams["redirect_id"] = redirect_id
	}
	return NewAffiliatesRedirectsWithRedirect_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewAffiliatesRedirectsRequestBuilderInternal instantiates a new AffiliatesRedirectsRequestBuilder and sets the default values.
func NewAffiliatesRedirectsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesRedirectsRequestBuilder {
	m := &AffiliatesRedirectsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/affiliates/redirects", pathParameters),
	}
	return m
}

// NewAffiliatesRedirectsRequestBuilder instantiates a new AffiliatesRedirectsRequestBuilder and sets the default values.
func NewAffiliatesRedirectsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesRedirectsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewAffiliatesRedirectsRequestBuilderInternal(urlParams, requestAdapter)
}
