package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
)

type TasksItemWithTaskPutRequestBody struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The completed property
	completed *bool
	// The completion_date property
	completion_date *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	// The contact property
	contact TasksItemWithTaskPutRequestBody_contactable
	// The creation_date property
	creation_date *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	// The description property
	description *string
	// The due_date property
	due_date *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	// The funnel_id property
	funnel_id *int64
	// The jgraph_id property
	jgraph_id *int64
	// The modification_date property
	modification_date *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	// The priority property
	priority *int32
	// Value in minutes before start_date to show pop-up reminder.  Acceptable values are in [`5`,`10`,`15`,`30`,`60`,`120`,`240`,`480`,`1440`,`2880`]
	remind_time *int32
	// The title property
	title *string
	// The type property
	typeEscaped *string
	// The url property
	url *string
	// The user_id property
	user_id *int64
}

// NewTasksItemWithTaskPutRequestBody instantiates a new TasksItemWithTaskPutRequestBody and sets the default values.
func NewTasksItemWithTaskPutRequestBody() *TasksItemWithTaskPutRequestBody {
	m := &TasksItemWithTaskPutRequestBody{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateTasksItemWithTaskPutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTasksItemWithTaskPutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTasksItemWithTaskPutRequestBody(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TasksItemWithTaskPutRequestBody) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCompleted gets the completed property value. The completed property
// returns a *bool when successful
func (m *TasksItemWithTaskPutRequestBody) GetCompleted() *bool {
	return m.completed
}

// GetCompletionDate gets the completion_date property value. The completion_date property
// returns a *Time when successful
func (m *TasksItemWithTaskPutRequestBody) GetCompletionDate() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.completion_date
}

// GetContact gets the contact property value. The contact property
// returns a TasksItemWithTaskPutRequestBody_contactable when successful
func (m *TasksItemWithTaskPutRequestBody) GetContact() TasksItemWithTaskPutRequestBody_contactable {
	return m.contact
}

// GetCreationDate gets the creation_date property value. The creation_date property
// returns a *Time when successful
func (m *TasksItemWithTaskPutRequestBody) GetCreationDate() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.creation_date
}

// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *TasksItemWithTaskPutRequestBody) GetDescription() *string {
	return m.description
}

// GetDueDate gets the due_date property value. The due_date property
// returns a *Time when successful
func (m *TasksItemWithTaskPutRequestBody) GetDueDate() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.due_date
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TasksItemWithTaskPutRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["completed"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCompleted(val)
		}
		return nil
	}
	res["completion_date"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetTimeValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCompletionDate(val)
		}
		return nil
	}
	res["contact"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateTasksItemWithTaskPutRequestBody_contactFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetContact(val.(TasksItemWithTaskPutRequestBody_contactable))
		}
		return nil
	}
	res["creation_date"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetTimeValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCreationDate(val)
		}
		return nil
	}
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
	res["due_date"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetTimeValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDueDate(val)
		}
		return nil
	}
	res["funnel_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFunnelId(val)
		}
		return nil
	}
	res["jgraph_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetJgraphId(val)
		}
		return nil
	}
	res["modification_date"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetTimeValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetModificationDate(val)
		}
		return nil
	}
	res["priority"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPriority(val)
		}
		return nil
	}
	res["remind_time"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRemindTime(val)
		}
		return nil
	}
	res["title"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTitle(val)
		}
		return nil
	}
	res["type"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTypeEscaped(val)
		}
		return nil
	}
	res["url"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetUrl(val)
		}
		return nil
	}
	res["user_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetUserId(val)
		}
		return nil
	}
	return res
}

// GetFunnelId gets the funnel_id property value. The funnel_id property
// returns a *int64 when successful
func (m *TasksItemWithTaskPutRequestBody) GetFunnelId() *int64 {
	return m.funnel_id
}

// GetJgraphId gets the jgraph_id property value. The jgraph_id property
// returns a *int64 when successful
func (m *TasksItemWithTaskPutRequestBody) GetJgraphId() *int64 {
	return m.jgraph_id
}

// GetModificationDate gets the modification_date property value. The modification_date property
// returns a *Time when successful
func (m *TasksItemWithTaskPutRequestBody) GetModificationDate() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.modification_date
}

// GetPriority gets the priority property value. The priority property
// returns a *int32 when successful
func (m *TasksItemWithTaskPutRequestBody) GetPriority() *int32 {
	return m.priority
}

// GetRemindTime gets the remind_time property value. Value in minutes before start_date to show pop-up reminder.  Acceptable values are in [`5`,`10`,`15`,`30`,`60`,`120`,`240`,`480`,`1440`,`2880`]
// returns a *int32 when successful
func (m *TasksItemWithTaskPutRequestBody) GetRemindTime() *int32 {
	return m.remind_time
}

// GetTitle gets the title property value. The title property
// returns a *string when successful
func (m *TasksItemWithTaskPutRequestBody) GetTitle() *string {
	return m.title
}

// GetTypeEscaped gets the type property value. The type property
// returns a *string when successful
func (m *TasksItemWithTaskPutRequestBody) GetTypeEscaped() *string {
	return m.typeEscaped
}

// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *TasksItemWithTaskPutRequestBody) GetUrl() *string {
	return m.url
}

// GetUserId gets the user_id property value. The user_id property
// returns a *int64 when successful
func (m *TasksItemWithTaskPutRequestBody) GetUserId() *int64 {
	return m.user_id
}

// Serialize serializes information the current object
func (m *TasksItemWithTaskPutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteBoolValue("completed", m.GetCompleted())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteTimeValue("completion_date", m.GetCompletionDate())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("contact", m.GetContact())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteTimeValue("creation_date", m.GetCreationDate())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("description", m.GetDescription())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteTimeValue("due_date", m.GetDueDate())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("funnel_id", m.GetFunnelId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("jgraph_id", m.GetJgraphId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteTimeValue("modification_date", m.GetModificationDate())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("priority", m.GetPriority())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("remind_time", m.GetRemindTime())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("title", m.GetTitle())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("type", m.GetTypeEscaped())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("url", m.GetUrl())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("user_id", m.GetUserId())
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
func (m *TasksItemWithTaskPutRequestBody) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCompleted sets the completed property value. The completed property
func (m *TasksItemWithTaskPutRequestBody) SetCompleted(value *bool) {
	m.completed = value
}

// SetCompletionDate sets the completion_date property value. The completion_date property
func (m *TasksItemWithTaskPutRequestBody) SetCompletionDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	m.completion_date = value
}

// SetContact sets the contact property value. The contact property
func (m *TasksItemWithTaskPutRequestBody) SetContact(value TasksItemWithTaskPutRequestBody_contactable) {
	m.contact = value
}

// SetCreationDate sets the creation_date property value. The creation_date property
func (m *TasksItemWithTaskPutRequestBody) SetCreationDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	m.creation_date = value
}

// SetDescription sets the description property value. The description property
func (m *TasksItemWithTaskPutRequestBody) SetDescription(value *string) {
	m.description = value
}

// SetDueDate sets the due_date property value. The due_date property
func (m *TasksItemWithTaskPutRequestBody) SetDueDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	m.due_date = value
}

// SetFunnelId sets the funnel_id property value. The funnel_id property
func (m *TasksItemWithTaskPutRequestBody) SetFunnelId(value *int64) {
	m.funnel_id = value
}

// SetJgraphId sets the jgraph_id property value. The jgraph_id property
func (m *TasksItemWithTaskPutRequestBody) SetJgraphId(value *int64) {
	m.jgraph_id = value
}

// SetModificationDate sets the modification_date property value. The modification_date property
func (m *TasksItemWithTaskPutRequestBody) SetModificationDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	m.modification_date = value
}

// SetPriority sets the priority property value. The priority property
func (m *TasksItemWithTaskPutRequestBody) SetPriority(value *int32) {
	m.priority = value
}

// SetRemindTime sets the remind_time property value. Value in minutes before start_date to show pop-up reminder.  Acceptable values are in [`5`,`10`,`15`,`30`,`60`,`120`,`240`,`480`,`1440`,`2880`]
func (m *TasksItemWithTaskPutRequestBody) SetRemindTime(value *int32) {
	m.remind_time = value
}

// SetTitle sets the title property value. The title property
func (m *TasksItemWithTaskPutRequestBody) SetTitle(value *string) {
	m.title = value
}

// SetTypeEscaped sets the type property value. The type property
func (m *TasksItemWithTaskPutRequestBody) SetTypeEscaped(value *string) {
	m.typeEscaped = value
}

// SetUrl sets the url property value. The url property
func (m *TasksItemWithTaskPutRequestBody) SetUrl(value *string) {
	m.url = value
}

// SetUserId sets the user_id property value. The user_id property
func (m *TasksItemWithTaskPutRequestBody) SetUserId(value *int64) {
	m.user_id = value
}

type TasksItemWithTaskPutRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCompleted() *bool
	GetCompletionDate() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetContact() TasksItemWithTaskPutRequestBody_contactable
	GetCreationDate() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetDescription() *string
	GetDueDate() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetFunnelId() *int64
	GetJgraphId() *int64
	GetModificationDate() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetPriority() *int32
	GetRemindTime() *int32
	GetTitle() *string
	GetTypeEscaped() *string
	GetUrl() *string
	GetUserId() *int64
	SetCompleted(value *bool)
	SetCompletionDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	SetContact(value TasksItemWithTaskPutRequestBody_contactable)
	SetCreationDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	SetDescription(value *string)
	SetDueDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	SetFunnelId(value *int64)
	SetJgraphId(value *int64)
	SetModificationDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	SetPriority(value *int32)
	SetRemindTime(value *int32)
	SetTitle(value *string)
	SetTypeEscaped(value *string)
	SetUrl(value *string)
	SetUserId(value *int64)
}
