package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CampaignsItemSequencesWithSequence_idRemoveContactsPostResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The remove_from_sequence_results property
	remove_from_sequence_results CampaignsItemSequencesWithSequence_idRemoveContactsPostResponse_remove_from_sequence_resultsable
}

// NewCampaignsItemSequencesWithSequence_idRemoveContactsPostResponse instantiates a new CampaignsItemSequencesWithSequence_idRemoveContactsPostResponse and sets the default values.
func NewCampaignsItemSequencesWithSequence_idRemoveContactsPostResponse() *CampaignsItemSequencesWithSequence_idRemoveContactsPostResponse {
	m := &CampaignsItemSequencesWithSequence_idRemoveContactsPostResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateCampaignsItemSequencesWithSequence_idRemoveContactsPostResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCampaignsItemSequencesWithSequence_idRemoveContactsPostResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCampaignsItemSequencesWithSequence_idRemoveContactsPostResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CampaignsItemSequencesWithSequence_idRemoveContactsPostResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CampaignsItemSequencesWithSequence_idRemoveContactsPostResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["remove_from_sequence_results"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateCampaignsItemSequencesWithSequence_idRemoveContactsPostResponse_remove_from_sequence_resultsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRemoveFromSequenceResults(val.(CampaignsItemSequencesWithSequence_idRemoveContactsPostResponse_remove_from_sequence_resultsable))
		}
		return nil
	}
	return res
}

// GetRemoveFromSequenceResults gets the remove_from_sequence_results property value. The remove_from_sequence_results property
// returns a CampaignsItemSequencesWithSequence_idRemoveContactsPostResponse_remove_from_sequence_resultsable when successful
func (m *CampaignsItemSequencesWithSequence_idRemoveContactsPostResponse) GetRemoveFromSequenceResults() CampaignsItemSequencesWithSequence_idRemoveContactsPostResponse_remove_from_sequence_resultsable {
	return m.remove_from_sequence_results
}

// Serialize serializes information the current object
func (m *CampaignsItemSequencesWithSequence_idRemoveContactsPostResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("remove_from_sequence_results", m.GetRemoveFromSequenceResults())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteAdditionalData(m.GetAdditionalData())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CampaignsItemSequencesWithSequence_idRemoveContactsPostResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetRemoveFromSequenceResults sets the remove_from_sequence_results property value. The remove_from_sequence_results property
func (m *CampaignsItemSequencesWithSequence_idRemoveContactsPostResponse) SetRemoveFromSequenceResults(value CampaignsItemSequencesWithSequence_idRemoveContactsPostResponse_remove_from_sequence_resultsable) {
	m.remove_from_sequence_results = value
}

type CampaignsItemSequencesWithSequence_idRemoveContactsPostResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetRemoveFromSequenceResults() CampaignsItemSequencesWithSequence_idRemoveContactsPostResponse_remove_from_sequence_resultsable
	SetRemoveFromSequenceResults(value CampaignsItemSequencesWithSequence_idRemoveContactsPostResponse_remove_from_sequence_resultsable)
}
