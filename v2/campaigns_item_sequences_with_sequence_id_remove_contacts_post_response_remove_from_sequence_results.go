package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CampaignsItemSequencesWithSequence_idRemoveContactsPostResponse_remove_from_sequence_results struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
}

// NewCampaignsItemSequencesWithSequence_idRemoveContactsPostResponse_remove_from_sequence_results instantiates a new CampaignsItemSequencesWithSequence_idRemoveContactsPostResponse_remove_from_sequence_results and sets the default values.
func NewCampaignsItemSequencesWithSequence_idRemoveContactsPostResponse_remove_from_sequence_results() *CampaignsItemSequencesWithSequence_idRemoveContactsPostResponse_remove_from_sequence_results {
	m := &CampaignsItemSequencesWithSequence_idRemoveContactsPostResponse_remove_from_sequence_results{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateCampaignsItemSequencesWithSequence_idRemoveContactsPostResponse_remove_from_sequence_resultsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCampaignsItemSequencesWithSequence_idRemoveContactsPostResponse_remove_from_sequence_resultsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCampaignsItemSequencesWithSequence_idRemoveContactsPostResponse_remove_from_sequence_results(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CampaignsItemSequencesWithSequence_idRemoveContactsPostResponse_remove_from_sequence_results) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CampaignsItemSequencesWithSequence_idRemoveContactsPostResponse_remove_from_sequence_results) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	return res
}

// Serialize serializes information the current object
func (m *CampaignsItemSequencesWithSequence_idRemoveContactsPostResponse_remove_from_sequence_results) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteAdditionalData(m.GetAdditionalData())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CampaignsItemSequencesWithSequence_idRemoveContactsPostResponse_remove_from_sequence_results) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

type CampaignsItemSequencesWithSequence_idRemoveContactsPostResponse_remove_from_sequence_resultsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
