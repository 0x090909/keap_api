package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NotesModelCustomFieldsPostRequestBody struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// An optional tab group to place the field under in the interface.  If not specified, will default to the 'Custom Fields' tab.
	group_id *int64
	// The label property
	label *string
	// The options property
	options []NotesModelCustomFieldsPostRequestBody_optionsable
	// An optional user group to choose from when selecting values for User or UserListBox fields.
	user_group_id *int64
}

// NewNotesModelCustomFieldsPostRequestBody instantiates a new NotesModelCustomFieldsPostRequestBody and sets the default values.
func NewNotesModelCustomFieldsPostRequestBody() *NotesModelCustomFieldsPostRequestBody {
	m := &NotesModelCustomFieldsPostRequestBody{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateNotesModelCustomFieldsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNotesModelCustomFieldsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewNotesModelCustomFieldsPostRequestBody(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NotesModelCustomFieldsPostRequestBody) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NotesModelCustomFieldsPostRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["group_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetGroupId(val)
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
		val, err := n.GetCollectionOfObjectValues(CreateNotesModelCustomFieldsPostRequestBody_optionsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]NotesModelCustomFieldsPostRequestBody_optionsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(NotesModelCustomFieldsPostRequestBody_optionsable)
				}
			}
			m.SetOptions(res)
		}
		return nil
	}
	res["user_group_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetUserGroupId(val)
		}
		return nil
	}
	return res
}

// GetGroupId gets the group_id property value. An optional tab group to place the field under in the interface.  If not specified, will default to the 'Custom Fields' tab.
// returns a *int64 when successful
func (m *NotesModelCustomFieldsPostRequestBody) GetGroupId() *int64 {
	return m.group_id
}

// GetLabel gets the label property value. The label property
// returns a *string when successful
func (m *NotesModelCustomFieldsPostRequestBody) GetLabel() *string {
	return m.label
}

// GetOptions gets the options property value. The options property
// returns a []NotesModelCustomFieldsPostRequestBody_optionsable when successful
func (m *NotesModelCustomFieldsPostRequestBody) GetOptions() []NotesModelCustomFieldsPostRequestBody_optionsable {
	return m.options
}

// GetUserGroupId gets the user_group_id property value. An optional user group to choose from when selecting values for User or UserListBox fields.
// returns a *int64 when successful
func (m *NotesModelCustomFieldsPostRequestBody) GetUserGroupId() *int64 {
	return m.user_group_id
}

// Serialize serializes information the current object
func (m *NotesModelCustomFieldsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt64Value("group_id", m.GetGroupId())
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
		err := writer.WriteInt64Value("user_group_id", m.GetUserGroupId())
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
func (m *NotesModelCustomFieldsPostRequestBody) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetGroupId sets the group_id property value. An optional tab group to place the field under in the interface.  If not specified, will default to the 'Custom Fields' tab.
func (m *NotesModelCustomFieldsPostRequestBody) SetGroupId(value *int64) {
	m.group_id = value
}

// SetLabel sets the label property value. The label property
func (m *NotesModelCustomFieldsPostRequestBody) SetLabel(value *string) {
	m.label = value
}

// SetOptions sets the options property value. The options property
func (m *NotesModelCustomFieldsPostRequestBody) SetOptions(value []NotesModelCustomFieldsPostRequestBody_optionsable) {
	m.options = value
}

// SetUserGroupId sets the user_group_id property value. An optional user group to choose from when selecting values for User or UserListBox fields.
func (m *NotesModelCustomFieldsPostRequestBody) SetUserGroupId(value *int64) {
	m.user_group_id = value
}

type NotesModelCustomFieldsPostRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetGroupId() *int64
	GetLabel() *string
	GetOptions() []NotesModelCustomFieldsPostRequestBody_optionsable
	GetUserGroupId() *int64
	SetGroupId(value *int64)
	SetLabel(value *string)
	SetOptions(value []NotesModelCustomFieldsPostRequestBody_optionsable)
	SetUserGroupId(value *int64)
}
