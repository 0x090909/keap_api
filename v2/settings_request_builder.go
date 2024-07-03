package v2

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SettingsRequestBuilder builds and executes requests for operations under \v2\settings
type SettingsRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// ApplicationsGetConfiguration the applicationsGetConfiguration property
// returns a *SettingsApplicationsGetConfigurationRequestBuilder when successful
func (m *SettingsRequestBuilder) ApplicationsGetConfiguration() *SettingsApplicationsGetConfigurationRequestBuilder {
	return NewSettingsApplicationsGetConfigurationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// ApplicationsIsEnabled the applicationsIsEnabled property
// returns a *SettingsApplicationsIsEnabledRequestBuilder when successful
func (m *SettingsRequestBuilder) ApplicationsIsEnabled() *SettingsApplicationsIsEnabledRequestBuilder {
	return NewSettingsApplicationsIsEnabledRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// NewSettingsRequestBuilderInternal instantiates a new SettingsRequestBuilder and sets the default values.
func NewSettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *SettingsRequestBuilder {
	m := &SettingsRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v2/settings", pathParameters),
	}
	return m
}

// NewSettingsRequestBuilder instantiates a new SettingsRequestBuilder and sets the default values.
func NewSettingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *SettingsRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewSettingsRequestBuilderInternal(urlParams, requestAdapter)
}

// ContactOptionTypes the contactOptionTypes property
// returns a *SettingsContactOptionTypesRequestBuilder when successful
func (m *SettingsRequestBuilder) ContactOptionTypes() *SettingsContactOptionTypesRequestBuilder {
	return NewSettingsContactOptionTypesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
