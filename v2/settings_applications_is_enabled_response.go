package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use SettingsApplicationsIsEnabledGetResponseable instead.
type SettingsApplicationsIsEnabledResponse struct {
	SettingsApplicationsIsEnabledGetResponse
}

// NewSettingsApplicationsIsEnabledResponse instantiates a new SettingsApplicationsIsEnabledResponse and sets the default values.
func NewSettingsApplicationsIsEnabledResponse() *SettingsApplicationsIsEnabledResponse {
	m := &SettingsApplicationsIsEnabledResponse{
		SettingsApplicationsIsEnabledGetResponse: *NewSettingsApplicationsIsEnabledGetResponse(),
	}
	return m
}

// CreateSettingsApplicationsIsEnabledResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSettingsApplicationsIsEnabledResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSettingsApplicationsIsEnabledResponse(), nil
}

// Deprecated: This class is obsolete. Use SettingsApplicationsIsEnabledGetResponseable instead.
type SettingsApplicationsIsEnabledResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	SettingsApplicationsIsEnabledGetResponseable
}