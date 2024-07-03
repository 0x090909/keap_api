package v2

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AffiliatesWithIdRemoveFromProgramRequestBuilder builds and executes requests for operations under \v2\affiliates\{id}:removeFromProgram
type AffiliatesWithIdRemoveFromProgramRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// NewAffiliatesWithIdRemoveFromProgramRequestBuilderInternal instantiates a new AffiliatesWithIdRemoveFromProgramRequestBuilder and sets the default values.
func NewAffiliatesWithIdRemoveFromProgramRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesWithIdRemoveFromProgramRequestBuilder {
	m := &AffiliatesWithIdRemoveFromProgramRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/affiliates/{id}:removeFromProgram", pathParameters),
	}
	return m
}

// NewAffiliatesWithIdRemoveFromProgramRequestBuilder instantiates a new AffiliatesWithIdRemoveFromProgramRequestBuilder and sets the default values.
func NewAffiliatesWithIdRemoveFromProgramRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *AffiliatesWithIdRemoveFromProgramRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewAffiliatesWithIdRemoveFromProgramRequestBuilderInternal(urlParams, requestAdapter)
}
