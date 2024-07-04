package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use OauthConnectUserinfoGetResponseable instead.
type OauthConnectUserinfoResponse struct {
	OauthConnectUserinfoGetResponse
}

// NewOauthConnectUserinfoResponse instantiates a new OauthConnectUserinfoResponse and sets the default values.
func NewOauthConnectUserinfoResponse() *OauthConnectUserinfoResponse {
	m := &OauthConnectUserinfoResponse{
		OauthConnectUserinfoGetResponse: *NewOauthConnectUserinfoGetResponse(),
	}
	return m
}

// CreateOauthConnectUserinfoResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOauthConnectUserinfoResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewOauthConnectUserinfoResponse(), nil
}

// Deprecated: This class is obsolete. Use OauthConnectUserinfoGetResponseable instead.
type OauthConnectUserinfoResponseable interface {
	OauthConnectUserinfoGetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
