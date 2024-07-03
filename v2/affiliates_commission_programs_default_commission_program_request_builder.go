package v2

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AffiliatesCommissionProgramsDefaultCommissionProgramRequestBuilder builds and executes requests for operations under \v2\affiliates\commissionPrograms\defaultCommissionProgram
type AffiliatesCommissionProgramsDefaultCommissionProgramRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ByCommission_program_id gets an item from the github.com/0x090909/keap_api.v2.affiliates.commissionPrograms.defaultCommissionProgram.item collection
// returns a *AffiliatesCommissionProgramsDefaultCommissionProgramWithCommission_program_ItemRequestBuilder when successful
func (m *AffiliatesCommissionProgramsDefaultCommissionProgramRequestBuilder) ByCommission_program_id(commission_program_id string) *AffiliatesCommissionProgramsDefaultCommissionProgramWithCommission_program_ItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if commission_program_id != "" {
		urlTplParams["commission_program_id"] = commission_program_id
	}
	return NewAffiliatesCommissionProgramsDefaultCommissionProgramWithCommission_program_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewAffiliatesCommissionProgramsDefaultCommissionProgramRequestBuilderInternal instantiates a new AffiliatesCommissionProgramsDefaultCommissionProgramRequestBuilder and sets the default values.
func NewAffiliatesCommissionProgramsDefaultCommissionProgramRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesCommissionProgramsDefaultCommissionProgramRequestBuilder {
	m := &AffiliatesCommissionProgramsDefaultCommissionProgramRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/affiliates/commissionPrograms/defaultCommissionProgram", pathParameters),
	}
	return m
}

// NewAffiliatesCommissionProgramsDefaultCommissionProgramRequestBuilder instantiates a new AffiliatesCommissionProgramsDefaultCommissionProgramRequestBuilder and sets the default values.
func NewAffiliatesCommissionProgramsDefaultCommissionProgramRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesCommissionProgramsDefaultCommissionProgramRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewAffiliatesCommissionProgramsDefaultCommissionProgramRequestBuilderInternal(urlParams, requestAdapter)
}
