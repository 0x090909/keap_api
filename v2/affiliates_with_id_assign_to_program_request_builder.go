package v2

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AffiliatesWithIdAssignToProgramRequestBuilder builds and executes requests for operations under \v2\affiliates\{id}:assignToProgram
type AffiliatesWithIdAssignToProgramRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// NewAffiliatesWithIdAssignToProgramRequestBuilderInternal instantiates a new AffiliatesWithIdAssignToProgramRequestBuilder and sets the default values.
func NewAffiliatesWithIdAssignToProgramRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesWithIdAssignToProgramRequestBuilder {
	m := &AffiliatesWithIdAssignToProgramRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/affiliates/{id}:assignToProgram", pathParameters),
	}
	return m
}

// NewAffiliatesWithIdAssignToProgramRequestBuilder instantiates a new AffiliatesWithIdAssignToProgramRequestBuilder and sets the default values.
func NewAffiliatesWithIdAssignToProgramRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesWithIdAssignToProgramRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewAffiliatesWithIdAssignToProgramRequestBuilderInternal(urlParams, requestAdapter)
}
