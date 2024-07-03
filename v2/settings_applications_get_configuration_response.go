package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use SettingsApplicationsGetConfigurationGetResponseable instead.
type SettingsApplicationsGetConfigurationResponse struct {
	SettingsApplicationsGetConfigurationGetResponse
}

// NewSettingsApplicationsGetConfigurationResponse instantiates a new SettingsApplicationsGetConfigurationResponse and sets the default values.
func NewSettingsApplicationsGetConfigurationResponse() *SettingsApplicationsGetConfigurationResponse {
	m := &SettingsApplicationsGetConfigurationResponse{
		SettingsApplicationsGetConfigurationGetResponse: *NewSettingsApplicationsGetConfigurationGetResponse(),
	}
	return m
}

// CreateSettingsApplicationsGetConfigurationResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSettingsApplicationsGetConfigurationResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSettingsApplicationsGetConfigurationResponse(), nil
}

// Deprecated: This class is obsolete. Use SettingsApplicationsGetConfigurationGetResponseable instead.
type SettingsApplicationsGetConfigurationResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	SettingsApplicationsGetConfigurationGetResponseable
}
