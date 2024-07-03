package v2

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ReferralsRequestBuilder builds and executes requests for operations under \v2\referrals
type ReferralsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// NewReferralsRequestBuilderInternal instantiates a new ReferralsRequestBuilder and sets the default values.
func NewReferralsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ReferralsRequestBuilder {
	m := &ReferralsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/referrals", pathParameters),
	}
	return m
}

// NewReferralsRequestBuilder instantiates a new ReferralsRequestBuilder and sets the default values.
func NewReferralsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *ReferralsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewReferralsRequestBuilderInternal(urlParams, requestAdapter)
}
