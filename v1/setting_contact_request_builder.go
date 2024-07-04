package v1

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SettingContactRequestBuilder builds and executes requests for operations under \v1\setting\contact
type SettingContactRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// NewSettingContactRequestBuilderInternal instantiates a new SettingContactRequestBuilder and sets the default values.
func NewSettingContactRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *SettingContactRequestBuilder {
	m := &SettingContactRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/setting/contact", pathParameters),
	}
	return m
}

// NewSettingContactRequestBuilder instantiates a new SettingContactRequestBuilder and sets the default values.
func NewSettingContactRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *SettingContactRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewSettingContactRequestBuilderInternal(urlParams, requestAdapter)
}

// OptionTypes the optionTypes property
// returns a *SettingContactOptionTypesRequestBuilder when successful
func (m *SettingContactRequestBuilder) OptionTypes() *SettingContactOptionTypesRequestBuilder {
	return NewSettingContactOptionTypesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
