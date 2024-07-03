package v2

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AffiliatesItemCommissionTotalRequestBuilder builds and executes requests for operations under \v2\affiliates\{affiliate_id-id}\commissionTotal
type AffiliatesItemCommissionTotalRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// NewAffiliatesItemCommissionTotalRequestBuilderInternal instantiates a new AffiliatesItemCommissionTotalRequestBuilder and sets the default values.
func NewAffiliatesItemCommissionTotalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesItemCommissionTotalRequestBuilder {
	m := &AffiliatesItemCommissionTotalRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/affiliates/{affiliate_id%2Did}/commissionTotal", pathParameters),
	}
	return m
}

// NewAffiliatesItemCommissionTotalRequestBuilder instantiates a new AffiliatesItemCommissionTotalRequestBuilder and sets the default values.
func NewAffiliatesItemCommissionTotalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesItemCommissionTotalRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewAffiliatesItemCommissionTotalRequestBuilderInternal(urlParams, requestAdapter)
}
