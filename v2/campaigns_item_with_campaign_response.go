package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use CampaignsItemWithCampaign_GetResponseable instead.
type CampaignsItemWithCampaign_Response struct {
	CampaignsItemWithCampaign_GetResponse
}

// NewCampaignsItemWithCampaign_Response instantiates a new CampaignsItemWithCampaign_Response and sets the default values.
func NewCampaignsItemWithCampaign_Response() *CampaignsItemWithCampaign_Response {
	m := &CampaignsItemWithCampaign_Response{
		CampaignsItemWithCampaign_GetResponse: *NewCampaignsItemWithCampaign_GetResponse(),
	}
	return m
}

// CreateCampaignsItemWithCampaign_ResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCampaignsItemWithCampaign_ResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCampaignsItemWithCampaign_Response(), nil
}

// Deprecated: This class is obsolete. Use CampaignsItemWithCampaign_GetResponseable instead.
type CampaignsItemWithCampaign_Responseable interface {
	CampaignsItemWithCampaign_GetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
