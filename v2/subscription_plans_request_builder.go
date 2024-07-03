package v2

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SubscriptionPlansRequestBuilder builds and executes requests for operations under \v2\subscriptionPlans
type SubscriptionPlansRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// NewSubscriptionPlansRequestBuilderInternal instantiates a new SubscriptionPlansRequestBuilder and sets the default values.
func NewSubscriptionPlansRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *SubscriptionPlansRequestBuilder {
	m := &SubscriptionPlansRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/subscriptionPlans", pathParameters),
	}
	return m
}

// NewSubscriptionPlansRequestBuilder instantiates a new SubscriptionPlansRequestBuilder and sets the default values.
func NewSubscriptionPlansRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *SubscriptionPlansRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewSubscriptionPlansRequestBuilderInternal(urlParams, requestAdapter)
}
