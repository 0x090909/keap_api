package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use CampaignsItemSequencesWithSequence_idAddContactsPostResponseable instead.
type CampaignsItemSequencesWithSequence_idAddContactsResponse struct {
	CampaignsItemSequencesWithSequence_idAddContactsPostResponse
}

// NewCampaignsItemSequencesWithSequence_idAddContactsResponse instantiates a new CampaignsItemSequencesWithSequence_idAddContactsResponse and sets the default values.
func NewCampaignsItemSequencesWithSequence_idAddContactsResponse() *CampaignsItemSequencesWithSequence_idAddContactsResponse {
	m := &CampaignsItemSequencesWithSequence_idAddContactsResponse{
		CampaignsItemSequencesWithSequence_idAddContactsPostResponse: *NewCampaignsItemSequencesWithSequence_idAddContactsPostResponse(),
	}
	return m
}

// CreateCampaignsItemSequencesWithSequence_idAddContactsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCampaignsItemSequencesWithSequence_idAddContactsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCampaignsItemSequencesWithSequence_idAddContactsResponse(), nil
}

// Deprecated: This class is obsolete. Use CampaignsItemSequencesWithSequence_idAddContactsPostResponseable instead.
type CampaignsItemSequencesWithSequence_idAddContactsResponseable interface {
	CampaignsItemSequencesWithSequence_idAddContactsPostResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}