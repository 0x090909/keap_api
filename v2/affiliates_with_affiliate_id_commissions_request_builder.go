package v2

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AffiliatesWithAffiliate_idCommissionsRequestBuilder builds and executes requests for operations under \v2\affiliates\{affiliate_id}:commissions
type AffiliatesWithAffiliate_idCommissionsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// NewAffiliatesWithAffiliate_idCommissionsRequestBuilderInternal instantiates a new AffiliatesWithAffiliate_idCommissionsRequestBuilder and sets the default values.
func NewAffiliatesWithAffiliate_idCommissionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesWithAffiliate_idCommissionsRequestBuilder {
	m := &AffiliatesWithAffiliate_idCommissionsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/affiliates/{affiliate_id}:commissions", pathParameters),
	}
	return m
}

// NewAffiliatesWithAffiliate_idCommissionsRequestBuilder instantiates a new AffiliatesWithAffiliate_idCommissionsRequestBuilder and sets the default values.
func NewAffiliatesWithAffiliate_idCommissionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesWithAffiliate_idCommissionsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewAffiliatesWithAffiliate_idCommissionsRequestBuilderInternal(urlParams, requestAdapter)
}
