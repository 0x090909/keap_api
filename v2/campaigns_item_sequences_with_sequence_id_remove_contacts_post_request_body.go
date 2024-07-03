package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CampaignsItemSequencesWithSequence_idRemoveContactsPostRequestBody struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The contact_ids property
	contact_ids []string
}

// NewCampaignsItemSequencesWithSequence_idRemoveContactsPostRequestBody instantiates a new CampaignsItemSequencesWithSequence_idRemoveContactsPostRequestBody and sets the default values.
func NewCampaignsItemSequencesWithSequence_idRemoveContactsPostRequestBody() *CampaignsItemSequencesWithSequence_idRemoveContactsPostRequestBody {
	m := &CampaignsItemSequencesWithSequence_idRemoveContactsPostRequestBody{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateCampaignsItemSequencesWithSequence_idRemoveContactsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCampaignsItemSequencesWithSequence_idRemoveContactsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCampaignsItemSequencesWithSequence_idRemoveContactsPostRequestBody(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CampaignsItemSequencesWithSequence_idRemoveContactsPostRequestBody) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetContactIds gets the contact_ids property value. The contact_ids property
// returns a []string when successful
func (m *CampaignsItemSequencesWithSequence_idRemoveContactsPostRequestBody) GetContactIds() []string {
	return m.contact_ids
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CampaignsItemSequencesWithSequence_idRemoveContactsPostRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["contact_ids"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfPrimitiveValues("string")
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]string, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = *(v.(*string))
				}
			}
			m.SetContactIds(res)
		}
		return nil
	}
	return res
}

// Serialize serializes information the current object
func (m *CampaignsItemSequencesWithSequence_idRemoveContactsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetContactIds() != nil {
		err := writer.WriteCollectionOfStringValues("contact_ids", m.GetContactIds())
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
func (m *CampaignsItemSequencesWithSequence_idRemoveContactsPostRequestBody) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetContactIds sets the contact_ids property value. The contact_ids property
func (m *CampaignsItemSequencesWithSequence_idRemoveContactsPostRequestBody) SetContactIds(value []string) {
	m.contact_ids = value
}

type CampaignsItemSequencesWithSequence_idRemoveContactsPostRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetContactIds() []string
	SetContactIds(value []string)
}