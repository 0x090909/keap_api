package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TasksModelCustomFieldsPostResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The default_value property
	default_value *string
	// The field_name property
	field_name *string
	// The id property
	id *int64
	// The label property
	label *string
	// The options property
	options []TasksModelCustomFieldsPostResponse_optionsable
}

// NewTasksModelCustomFieldsPostResponse instantiates a new TasksModelCustomFieldsPostResponse and sets the default values.
func NewTasksModelCustomFieldsPostResponse() *TasksModelCustomFieldsPostResponse {
	m := &TasksModelCustomFieldsPostResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateTasksModelCustomFieldsPostResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTasksModelCustomFieldsPostResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTasksModelCustomFieldsPostResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TasksModelCustomFieldsPostResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetDefaultValue gets the default_value property value. The default_value property
// returns a *string when successful
func (m *TasksModelCustomFieldsPostResponse) GetDefaultValue() *string {
	return m.default_value
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TasksModelCustomFieldsPostResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["default_value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDefaultValue(val)
		}
		return nil
	}
	res["field_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFieldName(val)
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
	res["label"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLabel(val)
		}
		return nil
	}
	res["options"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateTasksModelCustomFieldsPostResponse_optionsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]TasksModelCustomFieldsPostResponse_optionsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(TasksModelCustomFieldsPostResponse_optionsable)
				}
			}
			m.SetOptions(res)
		}
		return nil
	}
	return res
}

// GetFieldName gets the field_name property value. The field_name property
// returns a *string when successful
func (m *TasksModelCustomFieldsPostResponse) GetFieldName() *string {
	return m.field_name
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *TasksModelCustomFieldsPostResponse) GetId() *int64 {
	return m.id
}

// GetLabel gets the label property value. The label property
// returns a *string when successful
func (m *TasksModelCustomFieldsPostResponse) GetLabel() *string {
	return m.label
}

// GetOptions gets the options property value. The options property
// returns a []TasksModelCustomFieldsPostResponse_optionsable when successful
func (m *TasksModelCustomFieldsPostResponse) GetOptions() []TasksModelCustomFieldsPostResponse_optionsable {
	return m.options
}

// Serialize serializes information the current object
func (m *TasksModelCustomFieldsPostResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("default_value", m.GetDefaultValue())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("field_name", m.GetFieldName())
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
		err := writer.WriteStringValue("label", m.GetLabel())
		if err != nil {
			return err
		}
	}
	if m.GetOptions() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOptions()))
		for i, v := range m.GetOptions() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("options", cast)
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
func (m *TasksModelCustomFieldsPostResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetDefaultValue sets the default_value property value. The default_value property
func (m *TasksModelCustomFieldsPostResponse) SetDefaultValue(value *string) {
	m.default_value = value
}

// SetFieldName sets the field_name property value. The field_name property
func (m *TasksModelCustomFieldsPostResponse) SetFieldName(value *string) {
	m.field_name = value
}

// SetId sets the id property value. The id property
func (m *TasksModelCustomFieldsPostResponse) SetId(value *int64) {
	m.id = value
}

// SetLabel sets the label property value. The label property
func (m *TasksModelCustomFieldsPostResponse) SetLabel(value *string) {
	m.label = value
}

// SetOptions sets the options property value. The options property
func (m *TasksModelCustomFieldsPostResponse) SetOptions(value []TasksModelCustomFieldsPostResponse_optionsable) {
	m.options = value
}

type TasksModelCustomFieldsPostResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetDefaultValue() *string
	GetFieldName() *string
	GetId() *int64
	GetLabel() *string
	GetOptions() []TasksModelCustomFieldsPostResponse_optionsable
	SetDefaultValue(value *string)
	SetFieldName(value *string)
	SetId(value *int64)
	SetLabel(value *string)
	SetOptions(value []TasksModelCustomFieldsPostResponse_optionsable)
}
