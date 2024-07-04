package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CampaignsGoalsItemItemWithCallName struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The campaign_id property
	campaign_id *int64
	// The goal_id property
	goal_id *int64
	// The message property
	message *string
	// The success property
	success *bool
}

// NewCampaignsGoalsItemItemWithCallName instantiates a new CampaignsGoalsItemItemWithCallName and sets the default values.
func NewCampaignsGoalsItemItemWithCallName() *CampaignsGoalsItemItemWithCallName {
	m := &CampaignsGoalsItemItemWithCallName{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateCampaignsGoalsItemItemWithCallNameFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCampaignsGoalsItemItemWithCallNameFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCampaignsGoalsItemItemWithCallName(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CampaignsGoalsItemItemWithCallName) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCampaignId gets the campaign_id property value. The campaign_id property
// returns a *int64 when successful
func (m *CampaignsGoalsItemItemWithCallName) GetCampaignId() *int64 {
	return m.campaign_id
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CampaignsGoalsItemItemWithCallName) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["campaign_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCampaignId(val)
		}
		return nil
	}
	res["goal_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetGoalId(val)
		}
		return nil
	}
	res["message"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetMessage(val)
		}
		return nil
	}
	res["success"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSuccess(val)
		}
		return nil
	}
	return res
}

// GetGoalId gets the goal_id property value. The goal_id property
// returns a *int64 when successful
func (m *CampaignsGoalsItemItemWithCallName) GetGoalId() *int64 {
	return m.goal_id
}

// GetMessage gets the message property value. The message property
// returns a *string when successful
func (m *CampaignsGoalsItemItemWithCallName) GetMessage() *string {
	return m.message
}

// GetSuccess gets the success property value. The success property
// returns a *bool when successful
func (m *CampaignsGoalsItemItemWithCallName) GetSuccess() *bool {
	return m.success
}

// Serialize serializes information the current object
func (m *CampaignsGoalsItemItemWithCallName) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt64Value("campaign_id", m.GetCampaignId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("goal_id", m.GetGoalId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("message", m.GetMessage())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("success", m.GetSuccess())
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
func (m *CampaignsGoalsItemItemWithCallName) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCampaignId sets the campaign_id property value. The campaign_id property
func (m *CampaignsGoalsItemItemWithCallName) SetCampaignId(value *int64) {
	m.campaign_id = value
}

// SetGoalId sets the goal_id property value. The goal_id property
func (m *CampaignsGoalsItemItemWithCallName) SetGoalId(value *int64) {
	m.goal_id = value
}

// SetMessage sets the message property value. The message property
func (m *CampaignsGoalsItemItemWithCallName) SetMessage(value *string) {
	m.message = value
}

// SetSuccess sets the success property value. The success property
func (m *CampaignsGoalsItemItemWithCallName) SetSuccess(value *bool) {
	m.success = value
}

type CampaignsGoalsItemItemWithCallNameable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCampaignId() *int64
	GetGoalId() *int64
	GetMessage() *string
	GetSuccess() *bool
	SetCampaignId(value *int64)
	SetGoalId(value *int64)
	SetMessage(value *string)
	SetSuccess(value *bool)
}
