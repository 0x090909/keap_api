package v2

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AffiliatesCommissionProgramsProductCommissionProgramRequestBuilder builds and executes requests for operations under \v2\affiliates\commissionPrograms\productCommissionProgram
type AffiliatesCommissionProgramsProductCommissionProgramRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ByCommission_program_id gets an item from the github.com/0x090909/keap_api.v2.affiliates.commissionPrograms.productCommissionProgram.item collection
// returns a *AffiliatesCommissionProgramsProductCommissionProgramWithCommission_program_ItemRequestBuilder when successful
func (m *AffiliatesCommissionProgramsProductCommissionProgramRequestBuilder) ByCommission_program_id(commission_program_id string) *AffiliatesCommissionProgramsProductCommissionProgramWithCommission_program_ItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if commission_program_id != "" {
		urlTplParams["commission_program_id"] = commission_program_id
	}
	return NewAffiliatesCommissionProgramsProductCommissionProgramWithCommission_program_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewAffiliatesCommissionProgramsProductCommissionProgramRequestBuilderInternal instantiates a new AffiliatesCommissionProgramsProductCommissionProgramRequestBuilder and sets the default values.
func NewAffiliatesCommissionProgramsProductCommissionProgramRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesCommissionProgramsProductCommissionProgramRequestBuilder {
	m := &AffiliatesCommissionProgramsProductCommissionProgramRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/affiliates/commissionPrograms/productCommissionProgram", pathParameters),
	}
	return m
}

// NewAffiliatesCommissionProgramsProductCommissionProgramRequestBuilder instantiates a new AffiliatesCommissionProgramsProductCommissionProgramRequestBuilder and sets the default values.
func NewAffiliatesCommissionProgramsProductCommissionProgramRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesCommissionProgramsProductCommissionProgramRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewAffiliatesCommissionProgramsProductCommissionProgramRequestBuilderInternal(urlParams, requestAdapter)
}
