package v1

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SettingApplicationRequestBuilder builds and executes requests for operations under \v1\setting\application
type SettingApplicationRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// Configuration the configuration property
// returns a *SettingApplicationConfigurationRequestBuilder when successful
func (m *SettingApplicationRequestBuilder) Configuration() *SettingApplicationConfigurationRequestBuilder {
	return NewSettingApplicationConfigurationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// NewSettingApplicationRequestBuilderInternal instantiates a new SettingApplicationRequestBuilder and sets the default values.
func NewSettingApplicationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *SettingApplicationRequestBuilder {
	m := &SettingApplicationRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/setting/application", pathParameters),
	}
	return m
}

// NewSettingApplicationRequestBuilder instantiates a new SettingApplicationRequestBuilder and sets the default values.
func NewSettingApplicationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *SettingApplicationRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewSettingApplicationRequestBuilderInternal(urlParams, requestAdapter)
}

// Enabled the enabled property
// returns a *SettingApplicationEnabledRequestBuilder when successful
func (m *SettingApplicationRequestBuilder) Enabled() *SettingApplicationEnabledRequestBuilder {
	return NewSettingApplicationEnabledRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
