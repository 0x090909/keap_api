package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use CampaignsGetResponseable instead.
type CampaignsResponse struct {
	CampaignsGetResponse
}

// NewCampaignsResponse instantiates a new CampaignsResponse and sets the default values.
func NewCampaignsResponse() *CampaignsResponse {
	m := &CampaignsResponse{
		CampaignsGetResponse: *NewCampaignsGetResponse(),
	}
	return m
}

// CreateCampaignsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCampaignsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCampaignsResponse(), nil
}

// Deprecated: This class is obsolete. Use CampaignsGetResponseable instead.
type CampaignsResponseable interface {
	CampaignsGetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
