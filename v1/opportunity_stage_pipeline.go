package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OpportunityStage_pipeline struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The end_stage property
	end_stage *bool
	// The is_default property
	is_default *bool
	// The stage_count property
	stage_count *int32
	// The stage_id property
	stage_id *int64
	// The stage_name property
	stage_name *string
	// The stage_order property
	stage_order *int32
}

// NewOpportunityStage_pipeline instantiates a new OpportunityStage_pipeline and sets the default values.
func NewOpportunityStage_pipeline() *OpportunityStage_pipeline {
	m := &OpportunityStage_pipeline{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateOpportunityStage_pipelineFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOpportunityStage_pipelineFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewOpportunityStage_pipeline(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OpportunityStage_pipeline) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetEndStage gets the end_stage property value. The end_stage property
// returns a *bool when successful
func (m *OpportunityStage_pipeline) GetEndStage() *bool {
	return m.end_stage
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OpportunityStage_pipeline) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["end_stage"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetEndStage(val)
		}
		return nil
	}
	res["is_default"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetIsDefault(val)
		}
		return nil
	}
	res["stage_count"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetStageCount(val)
		}
		return nil
	}
	res["stage_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetStageId(val)
		}
		return nil
	}
	res["stage_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetStageName(val)
		}
		return nil
	}
	res["stage_order"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetStageOrder(val)
		}
		return nil
	}
	return res
}

// GetIsDefault gets the is_default property value. The is_default property
// returns a *bool when successful
func (m *OpportunityStage_pipeline) GetIsDefault() *bool {
	return m.is_default
}

// GetStageCount gets the stage_count property value. The stage_count property
// returns a *int32 when successful
func (m *OpportunityStage_pipeline) GetStageCount() *int32 {
	return m.stage_count
}

// GetStageId gets the stage_id property value. The stage_id property
// returns a *int64 when successful
func (m *OpportunityStage_pipeline) GetStageId() *int64 {
	return m.stage_id
}

// GetStageName gets the stage_name property value. The stage_name property
// returns a *string when successful
func (m *OpportunityStage_pipeline) GetStageName() *string {
	return m.stage_name
}

// GetStageOrder gets the stage_order property value. The stage_order property
// returns a *int32 when successful
func (m *OpportunityStage_pipeline) GetStageOrder() *int32 {
	return m.stage_order
}

// Serialize serializes information the current object
func (m *OpportunityStage_pipeline) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteBoolValue("end_stage", m.GetEndStage())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("is_default", m.GetIsDefault())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("stage_count", m.GetStageCount())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("stage_id", m.GetStageId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("stage_name", m.GetStageName())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("stage_order", m.GetStageOrder())
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
func (m *OpportunityStage_pipeline) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetEndStage sets the end_stage property value. The end_stage property
func (m *OpportunityStage_pipeline) SetEndStage(value *bool) {
	m.end_stage = value
}

// SetIsDefault sets the is_default property value. The is_default property
func (m *OpportunityStage_pipeline) SetIsDefault(value *bool) {
	m.is_default = value
}

// SetStageCount sets the stage_count property value. The stage_count property
func (m *OpportunityStage_pipeline) SetStageCount(value *int32) {
	m.stage_count = value
}

// SetStageId sets the stage_id property value. The stage_id property
func (m *OpportunityStage_pipeline) SetStageId(value *int64) {
	m.stage_id = value
}

// SetStageName sets the stage_name property value. The stage_name property
func (m *OpportunityStage_pipeline) SetStageName(value *string) {
	m.stage_name = value
}

// SetStageOrder sets the stage_order property value. The stage_order property
func (m *OpportunityStage_pipeline) SetStageOrder(value *int32) {
	m.stage_order = value
}

type OpportunityStage_pipelineable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetEndStage() *bool
	GetIsDefault() *bool
	GetStageCount() *int32
	GetStageId() *int64
	GetStageName() *string
	GetStageOrder() *int32
	SetEndStage(value *bool)
	SetIsDefault(value *bool)
	SetStageCount(value *int32)
	SetStageId(value *int64)
	SetStageName(value *string)
	SetStageOrder(value *int32)
}
