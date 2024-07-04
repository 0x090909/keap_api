package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
)

type OpportunitiesGetResponse_opportunities_stage_details_check_list_items struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The description property
	description *string
	// The done_date property
	done_date *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	// The id property
	id *int64
	// The instance_id property
	instance_id *int64
	// The item_order property
	item_order *int32
	// The required property
	required *bool
}

// NewOpportunitiesGetResponse_opportunities_stage_details_check_list_items instantiates a new OpportunitiesGetResponse_opportunities_stage_details_check_list_items and sets the default values.
func NewOpportunitiesGetResponse_opportunities_stage_details_check_list_items() *OpportunitiesGetResponse_opportunities_stage_details_check_list_items {
	m := &OpportunitiesGetResponse_opportunities_stage_details_check_list_items{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateOpportunitiesGetResponse_opportunities_stage_details_check_list_itemsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOpportunitiesGetResponse_opportunities_stage_details_check_list_itemsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewOpportunitiesGetResponse_opportunities_stage_details_check_list_items(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OpportunitiesGetResponse_opportunities_stage_details_check_list_items) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *OpportunitiesGetResponse_opportunities_stage_details_check_list_items) GetDescription() *string {
	return m.description
}

// GetDoneDate gets the done_date property value. The done_date property
// returns a *Time when successful
func (m *OpportunitiesGetResponse_opportunities_stage_details_check_list_items) GetDoneDate() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.done_date
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OpportunitiesGetResponse_opportunities_stage_details_check_list_items) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["description"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDescription(val)
		}
		return nil
	}
	res["done_date"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetTimeValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDoneDate(val)
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
	res["instance_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetInstanceId(val)
		}
		return nil
	}
	res["item_order"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetItemOrder(val)
		}
		return nil
	}
	res["required"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRequired(val)
		}
		return nil
	}
	return res
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *OpportunitiesGetResponse_opportunities_stage_details_check_list_items) GetId() *int64 {
	return m.id
}

// GetInstanceId gets the instance_id property value. The instance_id property
// returns a *int64 when successful
func (m *OpportunitiesGetResponse_opportunities_stage_details_check_list_items) GetInstanceId() *int64 {
	return m.instance_id
}

// GetItemOrder gets the item_order property value. The item_order property
// returns a *int32 when successful
func (m *OpportunitiesGetResponse_opportunities_stage_details_check_list_items) GetItemOrder() *int32 {
	return m.item_order
}

// GetRequired gets the required property value. The required property
// returns a *bool when successful
func (m *OpportunitiesGetResponse_opportunities_stage_details_check_list_items) GetRequired() *bool {
	return m.required
}

// Serialize serializes information the current object
func (m *OpportunitiesGetResponse_opportunities_stage_details_check_list_items) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("description", m.GetDescription())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteTimeValue("done_date", m.GetDoneDate())
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
		err := writer.WriteInt64Value("instance_id", m.GetInstanceId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("item_order", m.GetItemOrder())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("required", m.GetRequired())
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
func (m *OpportunitiesGetResponse_opportunities_stage_details_check_list_items) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetDescription sets the description property value. The description property
func (m *OpportunitiesGetResponse_opportunities_stage_details_check_list_items) SetDescription(value *string) {
	m.description = value
}

// SetDoneDate sets the done_date property value. The done_date property
func (m *OpportunitiesGetResponse_opportunities_stage_details_check_list_items) SetDoneDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	m.done_date = value
}

// SetId sets the id property value. The id property
func (m *OpportunitiesGetResponse_opportunities_stage_details_check_list_items) SetId(value *int64) {
	m.id = value
}

// SetInstanceId sets the instance_id property value. The instance_id property
func (m *OpportunitiesGetResponse_opportunities_stage_details_check_list_items) SetInstanceId(value *int64) {
	m.instance_id = value
}

// SetItemOrder sets the item_order property value. The item_order property
func (m *OpportunitiesGetResponse_opportunities_stage_details_check_list_items) SetItemOrder(value *int32) {
	m.item_order = value
}

// SetRequired sets the required property value. The required property
func (m *OpportunitiesGetResponse_opportunities_stage_details_check_list_items) SetRequired(value *bool) {
	m.required = value
}

type OpportunitiesGetResponse_opportunities_stage_details_check_list_itemsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetDescription() *string
	GetDoneDate() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetId() *int64
	GetInstanceId() *int64
	GetItemOrder() *int32
	GetRequired() *bool
	SetDescription(value *string)
	SetDoneDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	SetId(value *int64)
	SetInstanceId(value *int64)
	SetItemOrder(value *int32)
	SetRequired(value *bool)
}
