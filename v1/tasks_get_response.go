package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TasksGetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The count property
	count *int32
	// The next property
	next *string
	// The previous property
	previous *string
	// The sync_token property
	sync_token *string
	// The tasks property
	tasks []TasksGetResponse_tasksable
}

// NewTasksGetResponse instantiates a new TasksGetResponse and sets the default values.
func NewTasksGetResponse() *TasksGetResponse {
	m := &TasksGetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateTasksGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTasksGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTasksGetResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TasksGetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCount gets the count property value. The count property
// returns a *int32 when successful
func (m *TasksGetResponse) GetCount() *int32 {
	return m.count
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TasksGetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["count"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCount(val)
		}
		return nil
	}
	res["next"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNext(val)
		}
		return nil
	}
	res["previous"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPrevious(val)
		}
		return nil
	}
	res["sync_token"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSyncToken(val)
		}
		return nil
	}
	res["tasks"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateTasksGetResponse_tasksFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]TasksGetResponse_tasksable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(TasksGetResponse_tasksable)
				}
			}
			m.SetTasks(res)
		}
		return nil
	}
	return res
}

// GetNext gets the next property value. The next property
// returns a *string when successful
func (m *TasksGetResponse) GetNext() *string {
	return m.next
}

// GetPrevious gets the previous property value. The previous property
// returns a *string when successful
func (m *TasksGetResponse) GetPrevious() *string {
	return m.previous
}

// GetSyncToken gets the sync_token property value. The sync_token property
// returns a *string when successful
func (m *TasksGetResponse) GetSyncToken() *string {
	return m.sync_token
}

// GetTasks gets the tasks property value. The tasks property
// returns a []TasksGetResponse_tasksable when successful
func (m *TasksGetResponse) GetTasks() []TasksGetResponse_tasksable {
	return m.tasks
}

// Serialize serializes information the current object
func (m *TasksGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt32Value("count", m.GetCount())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("next", m.GetNext())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("previous", m.GetPrevious())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("sync_token", m.GetSyncToken())
		if err != nil {
			return err
		}
	}
	if m.GetTasks() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTasks()))
		for i, v := range m.GetTasks() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("tasks", cast)
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
func (m *TasksGetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCount sets the count property value. The count property
func (m *TasksGetResponse) SetCount(value *int32) {
	m.count = value
}

// SetNext sets the next property value. The next property
func (m *TasksGetResponse) SetNext(value *string) {
	m.next = value
}

// SetPrevious sets the previous property value. The previous property
func (m *TasksGetResponse) SetPrevious(value *string) {
	m.previous = value
}

// SetSyncToken sets the sync_token property value. The sync_token property
func (m *TasksGetResponse) SetSyncToken(value *string) {
	m.sync_token = value
}

// SetTasks sets the tasks property value. The tasks property
func (m *TasksGetResponse) SetTasks(value []TasksGetResponse_tasksable) {
	m.tasks = value
}

type TasksGetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCount() *int32
	GetNext() *string
	GetPrevious() *string
	GetSyncToken() *string
	GetTasks() []TasksGetResponse_tasksable
	SetCount(value *int32)
	SetNext(value *string)
	SetPrevious(value *string)
	SetSyncToken(value *string)
	SetTasks(value []TasksGetResponse_tasksable)
}
