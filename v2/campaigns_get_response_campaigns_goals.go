package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CampaignsGetResponse_campaigns_goals struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The historical_contact_counts property
	historical_contact_counts CampaignsGetResponse_campaigns_goals_historical_contact_countsable
	// The id property
	id *string
	// The name property
	name *string
	// The next_sequence_ids property
	next_sequence_ids []string
	// The previous_sequence_ids property
	previous_sequence_ids []string
}

// NewCampaignsGetResponse_campaigns_goals instantiates a new CampaignsGetResponse_campaigns_goals and sets the default values.
func NewCampaignsGetResponse_campaigns_goals() *CampaignsGetResponse_campaigns_goals {
	m := &CampaignsGetResponse_campaigns_goals{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateCampaignsGetResponse_campaigns_goalsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCampaignsGetResponse_campaigns_goalsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCampaignsGetResponse_campaigns_goals(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CampaignsGetResponse_campaigns_goals) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CampaignsGetResponse_campaigns_goals) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["historical_contact_counts"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateCampaignsGetResponse_campaigns_goals_historical_contact_countsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetHistoricalContactCounts(val.(CampaignsGetResponse_campaigns_goals_historical_contact_countsable))
		}
		return nil
	}
	res["id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetId(val)
		}
		return nil
	}
	res["name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetName(val)
		}
		return nil
	}
	res["next_sequence_ids"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
			m.SetNextSequenceIds(res)
		}
		return nil
	}
	res["previous_sequence_ids"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
			m.SetPreviousSequenceIds(res)
		}
		return nil
	}
	return res
}

// GetHistoricalContactCounts gets the historical_contact_counts property value. The historical_contact_counts property
// returns a CampaignsGetResponse_campaigns_goals_historical_contact_countsable when successful
func (m *CampaignsGetResponse_campaigns_goals) GetHistoricalContactCounts() CampaignsGetResponse_campaigns_goals_historical_contact_countsable {
	return m.historical_contact_counts
}

// GetId gets the id property value. The id property
// returns a *string when successful
func (m *CampaignsGetResponse_campaigns_goals) GetId() *string {
	return m.id
}

// GetName gets the name property value. The name property
// returns a *string when successful
func (m *CampaignsGetResponse_campaigns_goals) GetName() *string {
	return m.name
}

// GetNextSequenceIds gets the next_sequence_ids property value. The next_sequence_ids property
// returns a []string when successful
func (m *CampaignsGetResponse_campaigns_goals) GetNextSequenceIds() []string {
	return m.next_sequence_ids
}

// GetPreviousSequenceIds gets the previous_sequence_ids property value. The previous_sequence_ids property
// returns a []string when successful
func (m *CampaignsGetResponse_campaigns_goals) GetPreviousSequenceIds() []string {
	return m.previous_sequence_ids
}

// Serialize serializes information the current object
func (m *CampaignsGetResponse_campaigns_goals) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("historical_contact_counts", m.GetHistoricalContactCounts())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("id", m.GetId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("name", m.GetName())
		if err != nil {
			return err
		}
	}
	if m.GetNextSequenceIds() != nil {
		err := writer.WriteCollectionOfStringValues("next_sequence_ids", m.GetNextSequenceIds())
		if err != nil {
			return err
		}
	}
	if m.GetPreviousSequenceIds() != nil {
		err := writer.WriteCollectionOfStringValues("previous_sequence_ids", m.GetPreviousSequenceIds())
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
func (m *CampaignsGetResponse_campaigns_goals) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetHistoricalContactCounts sets the historical_contact_counts property value. The historical_contact_counts property
func (m *CampaignsGetResponse_campaigns_goals) SetHistoricalContactCounts(value CampaignsGetResponse_campaigns_goals_historical_contact_countsable) {
	m.historical_contact_counts = value
}

// SetId sets the id property value. The id property
func (m *CampaignsGetResponse_campaigns_goals) SetId(value *string) {
	m.id = value
}

// SetName sets the name property value. The name property
func (m *CampaignsGetResponse_campaigns_goals) SetName(value *string) {
	m.name = value
}

// SetNextSequenceIds sets the next_sequence_ids property value. The next_sequence_ids property
func (m *CampaignsGetResponse_campaigns_goals) SetNextSequenceIds(value []string) {
	m.next_sequence_ids = value
}

// SetPreviousSequenceIds sets the previous_sequence_ids property value. The previous_sequence_ids property
func (m *CampaignsGetResponse_campaigns_goals) SetPreviousSequenceIds(value []string) {
	m.previous_sequence_ids = value
}

type CampaignsGetResponse_campaigns_goalsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetHistoricalContactCounts() CampaignsGetResponse_campaigns_goals_historical_contact_countsable
	GetId() *string
	GetName() *string
	GetNextSequenceIds() []string
	GetPreviousSequenceIds() []string
	SetHistoricalContactCounts(value CampaignsGetResponse_campaigns_goals_historical_contact_countsable)
	SetId(value *string)
	SetName(value *string)
	SetNextSequenceIds(value []string)
	SetPreviousSequenceIds(value []string)
}
