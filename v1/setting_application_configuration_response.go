package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use SettingApplicationConfigurationGetResponseable instead.
type SettingApplicationConfigurationResponse struct {
	SettingApplicationConfigurationGetResponse
}

// NewSettingApplicationConfigurationResponse instantiates a new SettingApplicationConfigurationResponse and sets the default values.
func NewSettingApplicationConfigurationResponse() *SettingApplicationConfigurationResponse {
	m := &SettingApplicationConfigurationResponse{
		SettingApplicationConfigurationGetResponse: *NewSettingApplicationConfigurationGetResponse(),
	}
	return m
}

// CreateSettingApplicationConfigurationResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSettingApplicationConfigurationResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSettingApplicationConfigurationResponse(), nil
}

// Deprecated: This class is obsolete. Use SettingApplicationConfigurationGetResponseable instead.
type SettingApplicationConfigurationResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	SettingApplicationConfigurationGetResponseable
}
