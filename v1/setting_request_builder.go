package v1

import (
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SettingRequestBuilder builds and executes requests for operations under \v1\setting
type SettingRequestBuilder struct {
	i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}

// Application the application property
// returns a *SettingApplicationRequestBuilder when successful
func (m *SettingRequestBuilder) Application() *SettingApplicationRequestBuilder {
	return NewSettingApplicationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}

// NewSettingRequestBuilderInternal instantiates a new SettingRequestBuilder and sets the default values.
func NewSettingRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *SettingRequestBuilder {
	m := &SettingRequestBuilder{
		BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/v1/setting", pathParameters),
	}
	return m
}

// NewSettingRequestBuilder instantiates a new SettingRequestBuilder and sets the default values.
func NewSettingRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter) *SettingRequestBuilder {
	urlParams := make(map[string]string)
	urlParams["request-raw-url"] = rawUrl
	return NewSettingRequestBuilderInternal(urlParams, requestAdapter)
}

// Contact the contact property
// returns a *SettingContactRequestBuilder when successful
func (m *SettingRequestBuilder) Contact() *SettingContactRequestBuilder {
	return NewSettingContactRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
