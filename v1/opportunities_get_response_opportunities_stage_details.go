package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OpportunitiesGetResponse_opportunities_stage_details struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The check_list_items property
	check_list_items []OpportunitiesGetResponse_opportunities_stage_details_check_list_itemsable
	// The probability property
	probability *int32
	// The stage_order property
	stage_order *int32
	// The target_num_days property
	target_num_days *int32
}

// NewOpportunitiesGetResponse_opportunities_stage_details instantiates a new OpportunitiesGetResponse_opportunities_stage_details and sets the default values.
func NewOpportunitiesGetResponse_opportunities_stage_details() *OpportunitiesGetResponse_opportunities_stage_details {
	m := &OpportunitiesGetResponse_opportunities_stage_details{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateOpportunitiesGetResponse_opportunities_stage_detailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOpportunitiesGetResponse_opportunities_stage_detailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewOpportunitiesGetResponse_opportunities_stage_details(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OpportunitiesGetResponse_opportunities_stage_details) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCheckListItems gets the check_list_items property value. The check_list_items property
// returns a []OpportunitiesGetResponse_opportunities_stage_details_check_list_itemsable when successful
func (m *OpportunitiesGetResponse_opportunities_stage_details) GetCheckListItems() []OpportunitiesGetResponse_opportunities_stage_details_check_list_itemsable {
	return m.check_list_items
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OpportunitiesGetResponse_opportunities_stage_details) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["check_list_items"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateOpportunitiesGetResponse_opportunities_stage_details_check_list_itemsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]OpportunitiesGetResponse_opportunities_stage_details_check_list_itemsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(OpportunitiesGetResponse_opportunities_stage_details_check_list_itemsable)
				}
			}
			m.SetCheckListItems(res)
		}
		return nil
	}
	res["probability"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetProbability(val)
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
	res["target_num_days"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTargetNumDays(val)
		}
		return nil
	}
	return res
}

// GetProbability gets the probability property value. The probability property
// returns a *int32 when successful
func (m *OpportunitiesGetResponse_opportunities_stage_details) GetProbability() *int32 {
	return m.probability
}

// GetStageOrder gets the stage_order property value. The stage_order property
// returns a *int32 when successful
func (m *OpportunitiesGetResponse_opportunities_stage_details) GetStageOrder() *int32 {
	return m.stage_order
}

// GetTargetNumDays gets the target_num_days property value. The target_num_days property
// returns a *int32 when successful
func (m *OpportunitiesGetResponse_opportunities_stage_details) GetTargetNumDays() *int32 {
	return m.target_num_days
}

// Serialize serializes information the current object
func (m *OpportunitiesGetResponse_opportunities_stage_details) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetCheckListItems() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCheckListItems()))
		for i, v := range m.GetCheckListItems() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("check_list_items", cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("probability", m.GetProbability())
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
		err := writer.WriteInt32Value("target_num_days", m.GetTargetNumDays())
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
func (m *OpportunitiesGetResponse_opportunities_stage_details) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCheckListItems sets the check_list_items property value. The check_list_items property
func (m *OpportunitiesGetResponse_opportunities_stage_details) SetCheckListItems(value []OpportunitiesGetResponse_opportunities_stage_details_check_list_itemsable) {
	m.check_list_items = value
}

// SetProbability sets the probability property value. The probability property
func (m *OpportunitiesGetResponse_opportunities_stage_details) SetProbability(value *int32) {
	m.probability = value
}

// SetStageOrder sets the stage_order property value. The stage_order property
func (m *OpportunitiesGetResponse_opportunities_stage_details) SetStageOrder(value *int32) {
	m.stage_order = value
}

// SetTargetNumDays sets the target_num_days property value. The target_num_days property
func (m *OpportunitiesGetResponse_opportunities_stage_details) SetTargetNumDays(value *int32) {
	m.target_num_days = value
}

type OpportunitiesGetResponse_opportunities_stage_detailsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCheckListItems() []OpportunitiesGetResponse_opportunities_stage_details_check_list_itemsable
	GetProbability() *int32
	GetStageOrder() *int32
	GetTargetNumDays() *int32
	SetCheckListItems(value []OpportunitiesGetResponse_opportunities_stage_details_check_list_itemsable)
	SetProbability(value *int32)
	SetStageOrder(value *int32)
	SetTargetNumDays(value *int32)
}
