package v2

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AffiliatesCommissionProgramsSubscriptionCommissionProgramWithCommission_program_ItemRequestBuilder builds and executes requests for operations under \v2\affiliates\commissionPrograms\subscriptionCommissionProgram\{commission_program_id}
type AffiliatesCommissionProgramsSubscriptionCommissionProgramWithCommission_program_ItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// NewAffiliatesCommissionProgramsSubscriptionCommissionProgramWithCommission_program_ItemRequestBuilderInternal instantiates a new AffiliatesCommissionProgramsSubscriptionCommissionProgramWithCommission_program_ItemRequestBuilder and sets the default values.
func NewAffiliatesCommissionProgramsSubscriptionCommissionProgramWithCommission_program_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesCommissionProgramsSubscriptionCommissionProgramWithCommission_program_ItemRequestBuilder {
	m := &AffiliatesCommissionProgramsSubscriptionCommissionProgramWithCommission_program_ItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/affiliates/commissionPrograms/subscriptionCommissionProgram/{commission_program_id}", pathParameters),
	}
	return m
}

// NewAffiliatesCommissionProgramsSubscriptionCommissionProgramWithCommission_program_ItemRequestBuilder instantiates a new AffiliatesCommissionProgramsSubscriptionCommissionProgramWithCommission_program_ItemRequestBuilder and sets the default values.
func NewAffiliatesCommissionProgramsSubscriptionCommissionProgramWithCommission_program_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesCommissionProgramsSubscriptionCommissionProgramWithCommission_program_ItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewAffiliatesCommissionProgramsSubscriptionCommissionProgramWithCommission_program_ItemRequestBuilderInternal(urlParams, requestAdapter)
}