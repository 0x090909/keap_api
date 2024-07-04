package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CampaignsGetResponse_campaigns_sequences_paths_items struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The id property
	id *int64
	// The name property
	name *string
	// The next_item_id property
	next_item_id *int64
	// The previous_item_id property
	previous_item_id *int64
}

// NewCampaignsGetResponse_campaigns_sequences_paths_items instantiates a new CampaignsGetResponse_campaigns_sequences_paths_items and sets the default values.
func NewCampaignsGetResponse_campaigns_sequences_paths_items() *CampaignsGetResponse_campaigns_sequences_paths_items {
	m := &CampaignsGetResponse_campaigns_sequences_paths_items{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateCampaignsGetResponse_campaigns_sequences_paths_itemsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCampaignsGetResponse_campaigns_sequences_paths_itemsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCampaignsGetResponse_campaigns_sequences_paths_items(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CampaignsGetResponse_campaigns_sequences_paths_items) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CampaignsGetResponse_campaigns_sequences_paths_items) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
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
	res["next_item_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNextItemId(val)
		}
		return nil
	}
	res["previous_item_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPreviousItemId(val)
		}
		return nil
	}
	return res
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *CampaignsGetResponse_campaigns_sequences_paths_items) GetId() *int64 {
	return m.id
}

// GetName gets the name property value. The name property
// returns a *string when successful
func (m *CampaignsGetResponse_campaigns_sequences_paths_items) GetName() *string {
	return m.name
}

// GetNextItemId gets the next_item_id property value. The next_item_id property
// returns a *int64 when successful
func (m *CampaignsGetResponse_campaigns_sequences_paths_items) GetNextItemId() *int64 {
	return m.next_item_id
}

// GetPreviousItemId gets the previous_item_id property value. The previous_item_id property
// returns a *int64 when successful
func (m *CampaignsGetResponse_campaigns_sequences_paths_items) GetPreviousItemId() *int64 {
	return m.previous_item_id
}

// Serialize serializes information the current object
func (m *CampaignsGetResponse_campaigns_sequences_paths_items) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
	{
		err := writer.WriteInt64Value("next_item_id", m.GetNextItemId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("previous_item_id", m.GetPreviousItemId())
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
func (m *CampaignsGetResponse_campaigns_sequences_paths_items) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetId sets the id property value. The id property
func (m *CampaignsGetResponse_campaigns_sequences_paths_items) SetId(value *int64) {
	m.id = value
}

// SetName sets the name property value. The name property
func (m *CampaignsGetResponse_campaigns_sequences_paths_items) SetName(value *string) {
	m.name = value
}

// SetNextItemId sets the next_item_id property value. The next_item_id property
func (m *CampaignsGetResponse_campaigns_sequences_paths_items) SetNextItemId(value *int64) {
	m.next_item_id = value
}

// SetPreviousItemId sets the previous_item_id property value. The previous_item_id property
func (m *CampaignsGetResponse_campaigns_sequences_paths_items) SetPreviousItemId(value *int64) {
	m.previous_item_id = value
}

type CampaignsGetResponse_campaigns_sequences_paths_itemsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetId() *int64
	GetName() *string
	GetNextItemId() *int64
	GetPreviousItemId() *int64
	SetId(value *int64)
	SetName(value *string)
	SetNextItemId(value *int64)
	SetPreviousItemId(value *int64)
}
