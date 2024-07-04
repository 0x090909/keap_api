package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CampaignsItemWithCampaignGetResponse_sequences struct {
	// The active_contact_count property
	active_contact_count *int32
	// The active_contact_count_completed property
	active_contact_count_completed *int32
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The historical_contact_count property
	historical_contact_count CampaignsItemWithCampaignGetResponse_sequences_historical_contact_countable
	// The id property
	id *int64
	// The name property
	name *string
	// The paths property
	paths []CampaignsItemWithCampaignGetResponse_sequences_pathsable
}

// NewCampaignsItemWithCampaignGetResponse_sequences instantiates a new CampaignsItemWithCampaignGetResponse_sequences and sets the default values.
func NewCampaignsItemWithCampaignGetResponse_sequences() *CampaignsItemWithCampaignGetResponse_sequences {
	m := &CampaignsItemWithCampaignGetResponse_sequences{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateCampaignsItemWithCampaignGetResponse_sequencesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCampaignsItemWithCampaignGetResponse_sequencesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCampaignsItemWithCampaignGetResponse_sequences(), nil
}

// GetActiveContactCount gets the active_contact_count property value. The active_contact_count property
// returns a *int32 when successful
func (m *CampaignsItemWithCampaignGetResponse_sequences) GetActiveContactCount() *int32 {
	return m.active_contact_count
}

// GetActiveContactCountCompleted gets the active_contact_count_completed property value. The active_contact_count_completed property
// returns a *int32 when successful
func (m *CampaignsItemWithCampaignGetResponse_sequences) GetActiveContactCountCompleted() *int32 {
	return m.active_contact_count_completed
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CampaignsItemWithCampaignGetResponse_sequences) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CampaignsItemWithCampaignGetResponse_sequences) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["active_contact_count"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetActiveContactCount(val)
		}
		return nil
	}
	res["active_contact_count_completed"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetActiveContactCountCompleted(val)
		}
		return nil
	}
	res["historical_contact_count"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateCampaignsItemWithCampaignGetResponse_sequences_historical_contact_countFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetHistoricalContactCount(val.(CampaignsItemWithCampaignGetResponse_sequences_historical_contact_countable))
		}
		return nil
	}
	res["id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
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
	res["paths"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateCampaignsItemWithCampaignGetResponse_sequences_pathsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]CampaignsItemWithCampaignGetResponse_sequences_pathsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(CampaignsItemWithCampaignGetResponse_sequences_pathsable)
				}
			}
			m.SetPaths(res)
		}
		return nil
	}
	return res
}

// GetHistoricalContactCount gets the historical_contact_count property value. The historical_contact_count property
// returns a CampaignsItemWithCampaignGetResponse_sequences_historical_contact_countable when successful
func (m *CampaignsItemWithCampaignGetResponse_sequences) GetHistoricalContactCount() CampaignsItemWithCampaignGetResponse_sequences_historical_contact_countable {
	return m.historical_contact_count
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *CampaignsItemWithCampaignGetResponse_sequences) GetId() *int64 {
	return m.id
}

// GetName gets the name property value. The name property
// returns a *string when successful
func (m *CampaignsItemWithCampaignGetResponse_sequences) GetName() *string {
	return m.name
}

// GetPaths gets the paths property value. The paths property
// returns a []CampaignsItemWithCampaignGetResponse_sequences_pathsable when successful
func (m *CampaignsItemWithCampaignGetResponse_sequences) GetPaths() []CampaignsItemWithCampaignGetResponse_sequences_pathsable {
	return m.paths
}

// Serialize serializes information the current object
func (m *CampaignsItemWithCampaignGetResponse_sequences) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt32Value("active_contact_count", m.GetActiveContactCount())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("active_contact_count_completed", m.GetActiveContactCountCompleted())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("historical_contact_count", m.GetHistoricalContactCount())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("id", m.GetId())
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
	if m.GetPaths() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPaths()))
		for i, v := range m.GetPaths() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("paths", cast)
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

// SetActiveContactCount sets the active_contact_count property value. The active_contact_count property
func (m *CampaignsItemWithCampaignGetResponse_sequences) SetActiveContactCount(value *int32) {
	m.active_contact_count = value
}

// SetActiveContactCountCompleted sets the active_contact_count_completed property value. The active_contact_count_completed property
func (m *CampaignsItemWithCampaignGetResponse_sequences) SetActiveContactCountCompleted(value *int32) {
	m.active_contact_count_completed = value
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CampaignsItemWithCampaignGetResponse_sequences) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetHistoricalContactCount sets the historical_contact_count property value. The historical_contact_count property
func (m *CampaignsItemWithCampaignGetResponse_sequences) SetHistoricalContactCount(value CampaignsItemWithCampaignGetResponse_sequences_historical_contact_countable) {
	m.historical_contact_count = value
}

// SetId sets the id property value. The id property
func (m *CampaignsItemWithCampaignGetResponse_sequences) SetId(value *int64) {
	m.id = value
}

// SetName sets the name property value. The name property
func (m *CampaignsItemWithCampaignGetResponse_sequences) SetName(value *string) {
	m.name = value
}

// SetPaths sets the paths property value. The paths property
func (m *CampaignsItemWithCampaignGetResponse_sequences) SetPaths(value []CampaignsItemWithCampaignGetResponse_sequences_pathsable) {
	m.paths = value
}

type CampaignsItemWithCampaignGetResponse_sequencesable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetActiveContactCount() *int32
	GetActiveContactCountCompleted() *int32
	GetHistoricalContactCount() CampaignsItemWithCampaignGetResponse_sequences_historical_contact_countable
	GetId() *int64
	GetName() *string
	GetPaths() []CampaignsItemWithCampaignGetResponse_sequences_pathsable
	SetActiveContactCount(value *int32)
	SetActiveContactCountCompleted(value *int32)
	SetHistoricalContactCount(value CampaignsItemWithCampaignGetResponse_sequences_historical_contact_countable)
	SetId(value *int64)
	SetName(value *string)
	SetPaths(value []CampaignsItemWithCampaignGetResponse_sequences_pathsable)
}
