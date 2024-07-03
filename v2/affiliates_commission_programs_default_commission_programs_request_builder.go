package v2

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AffiliatesCommissionProgramsDefaultCommissionProgramsRequestBuilder builds and executes requests for operations under \v2\affiliates\commissionPrograms\defaultCommissionPrograms
type AffiliatesCommissionProgramsDefaultCommissionProgramsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ByCommission_program_id gets an item from the github.com/0x090909/keap_api.v2.affiliates.commissionPrograms.defaultCommissionPrograms.item collection
// returns a *AffiliatesCommissionProgramsDefaultCommissionProgramsWithCommission_program_ItemRequestBuilder when successful
func (m *AffiliatesCommissionProgramsDefaultCommissionProgramsRequestBuilder) ByCommission_program_id(commission_program_id string) *AffiliatesCommissionProgramsDefaultCommissionProgramsWithCommission_program_ItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if commission_program_id != "" {
		urlTplParams["commission_program_id"] = commission_program_id
	}
	return NewAffiliatesCommissionProgramsDefaultCommissionProgramsWithCommission_program_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewAffiliatesCommissionProgramsDefaultCommissionProgramsRequestBuilderInternal instantiates a new AffiliatesCommissionProgramsDefaultCommissionProgramsRequestBuilder and sets the default values.
func NewAffiliatesCommissionProgramsDefaultCommissionProgramsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesCommissionProgramsDefaultCommissionProgramsRequestBuilder {
	m := &AffiliatesCommissionProgramsDefaultCommissionProgramsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/affiliates/commissionPrograms/defaultCommissionPrograms", pathParameters),
	}
	return m
}

// NewAffiliatesCommissionProgramsDefaultCommissionProgramsRequestBuilder instantiates a new AffiliatesCommissionProgramsDefaultCommissionProgramsRequestBuilder and sets the default values.
func NewAffiliatesCommissionProgramsDefaultCommissionProgramsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesCommissionProgramsDefaultCommissionProgramsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewAffiliatesCommissionProgramsDefaultCommissionProgramsRequestBuilderInternal(urlParams, requestAdapter)
}
