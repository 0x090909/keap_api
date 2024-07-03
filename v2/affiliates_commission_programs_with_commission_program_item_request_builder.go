package v2

import (
	"context"
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AffiliatesCommissionProgramsWithCommission_program_ItemRequestBuilder builds and executes requests for operations under \v2\affiliates\commissionPrograms\{commission_program_id}
type AffiliatesCommissionProgramsWithCommission_program_ItemRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// AffiliatesCommissionProgramsWithCommission_program_ItemRequestBuilderPatchQueryParameters updates the properties of an Affiliate Commission Program
type AffiliatesCommissionProgramsWithCommission_program_ItemRequestBuilderPatchQueryParameters struct {
	// An optional list of properties to be updated. If set, only the provided properties will be updated and others will be skipped.
	Update_mask []string `uriparametername:"update_mask"`
}

// AffiliatesCommissionProgramsWithCommission_program_ItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AffiliatesCommissionProgramsWithCommission_program_ItemRequestBuilderPatchRequestConfiguration struct {
	// Request headers
	Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
	// Request options
	Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
	// Request query parameters
	QueryParameters *AffiliatesCommissionProgramsWithCommission_program_ItemRequestBuilderPatchQueryParameters
}

// NewAffiliatesCommissionProgramsWithCommission_program_ItemRequestBuilderInternal instantiates a new AffiliatesCommissionProgramsWithCommission_program_ItemRequestBuilder and sets the default values.
func NewAffiliatesCommissionProgramsWithCommission_program_ItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesCommissionProgramsWithCommission_program_ItemRequestBuilder {
	m := &AffiliatesCommissionProgramsWithCommission_program_ItemRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/affiliates/commissionPrograms/{commission_program_id}{?update_mask*}", pathParameters),
	}
	return m
}

// NewAffiliatesCommissionProgramsWithCommission_program_ItemRequestBuilder instantiates a new AffiliatesCommissionProgramsWithCommission_program_ItemRequestBuilder and sets the default values.
func NewAffiliatesCommissionProgramsWithCommission_program_ItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesCommissionProgramsWithCommission_program_ItemRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewAffiliatesCommissionProgramsWithCommission_program_ItemRequestBuilderInternal(urlParams, requestAdapter)
}

// Patch updates the properties of an Affiliate Commission Program
// Deprecated: This method is obsolete. Use PatchAsWithCommission_program_PatchResponse instead.
// returns a AffiliatesCommissionProgramsItemWithCommission_program_Responseable when successful
// returns a AffiliatesCommissionProgramsItemWithCommission_program_401Error error when the service returns a 401 status code
// returns a AffiliatesCommissionProgramsItemWithCommission_program_403Error error when the service returns a 403 status code
// returns a AffiliatesCommissionProgramsItemWithCommission_program_404Error error when the service returns a 404 status code
func (m *AffiliatesCommissionProgramsWithCommission_program_ItemRequestBuilder) Patch(ctx context.Context, body AffiliatesCommissionProgramsItemWithCommission_program_PatchRequestBodyable, requestConfiguration *AffiliatesCommissionProgramsWithCommission_program_ItemRequestBuilderPatchRequestConfiguration) (AffiliatesCommissionProgramsItemWithCommission_program_Responseable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAffiliatesCommissionProgramsItemWithCommission_program_401ErrorFromDiscriminatorValue,
		"403": CreateAffiliatesCommissionProgramsItemWithCommission_program_403ErrorFromDiscriminatorValue,
		"404": CreateAffiliatesCommissionProgramsItemWithCommission_program_404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAffiliatesCommissionProgramsItemWithCommission_program_ResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AffiliatesCommissionProgramsItemWithCommission_program_Responseable), nil
}

// PatchAsWithCommission_program_PatchResponse updates the properties of an Affiliate Commission Program
// returns a AffiliatesCommissionProgramsItemWithCommission_program_PatchResponseable when successful
// returns a AffiliatesCommissionProgramsItemWithCommission_program_401Error error when the service returns a 401 status code
// returns a AffiliatesCommissionProgramsItemWithCommission_program_403Error error when the service returns a 403 status code
// returns a AffiliatesCommissionProgramsItemWithCommission_program_404Error error when the service returns a 404 status code
func (m *AffiliatesCommissionProgramsWithCommission_program_ItemRequestBuilder) PatchAsWithCommission_program_PatchResponse(ctx context.Context, body AffiliatesCommissionProgramsItemWithCommission_program_PatchRequestBodyable, requestConfiguration *AffiliatesCommissionProgramsWithCommission_program_ItemRequestBuilderPatchRequestConfiguration) (AffiliatesCommissionProgramsItemWithCommission_program_PatchResponseable, error) {
	requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration)
	if err != nil {
		return nil, err
	}
	errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings{
		"401": CreateAffiliatesCommissionProgramsItemWithCommission_program_401ErrorFromDiscriminatorValue,
		"403": CreateAffiliatesCommissionProgramsItemWithCommission_program_403ErrorFromDiscriminatorValue,
		"404": CreateAffiliatesCommissionProgramsItemWithCommission_program_404ErrorFromDiscriminatorValue,
	}
	res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateAffiliatesCommissionProgramsItemWithCommission_program_PatchResponseFromDiscriminatorValue, errorMapping)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(AffiliatesCommissionProgramsItemWithCommission_program_PatchResponseable), nil
}

// ToPatchRequestInformation updates the properties of an Affiliate Commission Program
// returns a *RequestInformation when successful
func (m *AffiliatesCommissionProgramsWithCommission_program_ItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body AffiliatesCommissionProgramsItemWithCommission_program_PatchRequestBodyable, requestConfiguration *AffiliatesCommissionProgramsWithCommission_program_ItemRequestBuilderPatchRequestConfiguration) (*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
	requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
	if requestConfiguration != nil {
		if requestConfiguration.QueryParameters != nil {
			requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
		}
		requestInfo.Headers.AddAll(requestConfiguration.Headers)
		requestInfo.AddRequestOptions(requestConfiguration.Options)
	}
	requestInfo.Headers.TryAdd("Accept", "application/json")
	err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
	if err != nil {
		return nil, err
	}
	return requestInfo, nil
}

// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *AffiliatesCommissionProgramsWithCommission_program_ItemRequestBuilder when successful
func (m *AffiliatesCommissionProgramsWithCommission_program_ItemRequestBuilder) WithUrl(rawUrl string) *AffiliatesCommissionProgramsWithCommission_program_ItemRequestBuilder {
	return NewAffiliatesCommissionProgramsWithCommission_program_ItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter)
}
