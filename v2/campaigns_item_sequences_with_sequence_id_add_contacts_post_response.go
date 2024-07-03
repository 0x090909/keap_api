package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CampaignsItemSequencesWithSequence_idAddContactsPostResponse struct {
	// The add_to_sequence_results property
	add_to_sequence_results CampaignsItemSequencesWithSequence_idAddContactsPostResponse_add_to_sequence_resultsable
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
}

// NewCampaignsItemSequencesWithSequence_idAddContactsPostResponse instantiates a new CampaignsItemSequencesWithSequence_idAddContactsPostResponse and sets the default values.
func NewCampaignsItemSequencesWithSequence_idAddContactsPostResponse() *CampaignsItemSequencesWithSequence_idAddContactsPostResponse {
	m := &CampaignsItemSequencesWithSequence_idAddContactsPostResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateCampaignsItemSequencesWithSequence_idAddContactsPostResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCampaignsItemSequencesWithSequence_idAddContactsPostResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCampaignsItemSequencesWithSequence_idAddContactsPostResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CampaignsItemSequencesWithSequence_idAddContactsPostResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAddToSequenceResults gets the add_to_sequence_results property value. The add_to_sequence_results property
// returns a CampaignsItemSequencesWithSequence_idAddContactsPostResponse_add_to_sequence_resultsable when successful
func (m *CampaignsItemSequencesWithSequence_idAddContactsPostResponse) GetAddToSequenceResults() CampaignsItemSequencesWithSequence_idAddContactsPostResponse_add_to_sequence_resultsable {
	return m.add_to_sequence_results
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CampaignsItemSequencesWithSequence_idAddContactsPostResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["add_to_sequence_results"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateCampaignsItemSequencesWithSequence_idAddContactsPostResponse_add_to_sequence_resultsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAddToSequenceResults(val.(CampaignsItemSequencesWithSequence_idAddContactsPostResponse_add_to_sequence_resultsable))
		}
		return nil
	}
	return res
}

// Serialize serializes information the current object
func (m *CampaignsItemSequencesWithSequence_idAddContactsPostResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("add_to_sequence_results", m.GetAddToSequenceResults())
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
func (m *CampaignsItemSequencesWithSequence_idAddContactsPostResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAddToSequenceResults sets the add_to_sequence_results property value. The add_to_sequence_results property
func (m *CampaignsItemSequencesWithSequence_idAddContactsPostResponse) SetAddToSequenceResults(value CampaignsItemSequencesWithSequence_idAddContactsPostResponse_add_to_sequence_resultsable) {
	m.add_to_sequence_results = value
}

type CampaignsItemSequencesWithSequence_idAddContactsPostResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAddToSequenceResults() CampaignsItemSequencesWithSequence_idAddContactsPostResponse_add_to_sequence_resultsable
	SetAddToSequenceResults(value CampaignsItemSequencesWithSequence_idAddContactsPostResponse_add_to_sequence_resultsable)
}
