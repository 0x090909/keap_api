package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use CampaignsItemWithCampaignGetResponseable instead.
type CampaignsItemWithCampaignResponse struct {
	CampaignsItemWithCampaignGetResponse
}

// NewCampaignsItemWithCampaignResponse instantiates a new CampaignsItemWithCampaignResponse and sets the default values.
func NewCampaignsItemWithCampaignResponse() *CampaignsItemWithCampaignResponse {
	m := &CampaignsItemWithCampaignResponse{
		CampaignsItemWithCampaignGetResponse: *NewCampaignsItemWithCampaignGetResponse(),
	}
	return m
}

// CreateCampaignsItemWithCampaignResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCampaignsItemWithCampaignResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCampaignsItemWithCampaignResponse(), nil
}

// Deprecated: This class is obsolete. Use CampaignsItemWithCampaignGetResponseable instead.
type CampaignsItemWithCampaignResponseable interface {
	CampaignsItemWithCampaignGetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
