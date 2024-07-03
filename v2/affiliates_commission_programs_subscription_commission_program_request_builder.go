package v2

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AffiliatesCommissionProgramsSubscriptionCommissionProgramRequestBuilder builds and executes requests for operations under \v2\affiliates\commissionPrograms\subscriptionCommissionProgram
type AffiliatesCommissionProgramsSubscriptionCommissionProgramRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ByCommission_program_id gets an item from the github.com/0x090909/keap_api.v2.affiliates.commissionPrograms.subscriptionCommissionProgram.item collection
// returns a *AffiliatesCommissionProgramsSubscriptionCommissionProgramWithCommission_program_ItemRequestBuilder when successful
func (m *AffiliatesCommissionProgramsSubscriptionCommissionProgramRequestBuilder) ByCommission_program_id(commission_program_id string) *AffiliatesCommissionProgramsSubscriptionCommissionProgramWithCommission_program_ItemRequestBuilder {
	urlTplParams := make(map[string]string)
	for idx, item := range m.BaseRequestBuilder.PathParameters {
		urlTplParams[idx] = item
	}
	if commission_program_id != "" {
		urlTplParams["commission_program_id"] = commission_program_id
	}
	return NewAffiliatesCommissionProgramsSubscriptionCommissionProgramWithCommission_program_ItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}

// NewAffiliatesCommissionProgramsSubscriptionCommissionProgramRequestBuilderInternal instantiates a new AffiliatesCommissionProgramsSubscriptionCommissionProgramRequestBuilder and sets the default values.
func NewAffiliatesCommissionProgramsSubscriptionCommissionProgramRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesCommissionProgramsSubscriptionCommissionProgramRequestBuilder {
	m := &AffiliatesCommissionProgramsSubscriptionCommissionProgramRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/affiliates/commissionPrograms/subscriptionCommissionProgram", pathParameters),
	}
	return m
}

// NewAffiliatesCommissionProgramsSubscriptionCommissionProgramRequestBuilder instantiates a new AffiliatesCommissionProgramsSubscriptionCommissionProgramRequestBuilder and sets the default values.
func NewAffiliatesCommissionProgramsSubscriptionCommissionProgramRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesCommissionProgramsSubscriptionCommissionProgramRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewAffiliatesCommissionProgramsSubscriptionCommissionProgramRequestBuilderInternal(urlParams, requestAdapter)
}