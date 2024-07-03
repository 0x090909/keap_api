package v2

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AffiliatesCommissionProgramsProductCommissionProgramsRequestBuilder builds and executes requests for operations under \v2\affiliates\commissionPrograms\productCommissionPrograms
type AffiliatesCommissionProgramsProductCommissionProgramsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ByCommission_program_id gets an item from the github.com/0x090909/keap_api.v2.affiliates.commissionPrograms.productCommissionPrograms.item collection
// returns a *AffiliatesCommissionProgramsProductCommissionProgramsWithCommission_program_ItemRequestBuilder when successful
func (m *AffiliatesCommissionProgramsProductCommissionProgramsRequestBuilder) ByCommission_program_id(commission_program_id string) *AffiliatesCommissionProgramsProductCommissionProgramsWithCommission_program_ItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if commission_program_id != "" {
		urlTplParams["commission_program_id"] = commission_program_id
	}
	return NewAffiliatesCommissionProgramsProductCommissionProgramsWithCommission_program_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewAffiliatesCommissionProgramsProductCommissionProgramsRequestBuilderInternal instantiates a new AffiliatesCommissionProgramsProductCommissionProgramsRequestBuilder and sets the default values.
func NewAffiliatesCommissionProgramsProductCommissionProgramsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesCommissionProgramsProductCommissionProgramsRequestBuilder {
	m := &AffiliatesCommissionProgramsProductCommissionProgramsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/affiliates/commissionPrograms/productCommissionPrograms", pathParameters),
	}
	return m
}

// NewAffiliatesCommissionProgramsProductCommissionProgramsRequestBuilder instantiates a new AffiliatesCommissionProgramsProductCommissionProgramsRequestBuilder and sets the default values.
func NewAffiliatesCommissionProgramsProductCommissionProgramsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesCommissionProgramsProductCommissionProgramsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewAffiliatesCommissionProgramsProductCommissionProgramsRequestBuilderInternal(urlParams, requestAdapter)
}
