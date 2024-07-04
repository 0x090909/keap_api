package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use SettingApplicationEnabledGetResponseable instead.
type SettingApplicationEnabledResponse struct {
	SettingApplicationEnabledGetResponse
}

// NewSettingApplicationEnabledResponse instantiates a new SettingApplicationEnabledResponse and sets the default values.
func NewSettingApplicationEnabledResponse() *SettingApplicationEnabledResponse {
	m := &SettingApplicationEnabledResponse{
		SettingApplicationEnabledGetResponse: *NewSettingApplicationEnabledGetResponse(),
	}
	return m
}

// CreateSettingApplicationEnabledResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSettingApplicationEnabledResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSettingApplicationEnabledResponse(), nil
}

// Deprecated: This class is obsolete. Use SettingApplicationEnabledGetResponseable instead.
type SettingApplicationEnabledResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	SettingApplicationEnabledGetResponseable
}
