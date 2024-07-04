package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
)

type AppointmentsItemWithAppointmentPatchRequestBody struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// Required for pop-up reminders
	contact_id *int64
	// The description property
	description *string
	// The end_date property
	end_date *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	// The location property
	location *string
	// Value in minutes before start_date to show pop-up reminder.  Acceptable values are [5,10,15,30,60,120,240,480,1440,2880]
	remind_time *int32
	// The start_date property
	start_date *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	// The title property
	title *string
	// Required only for pop-up reminders
	user *int64
}

// NewAppointmentsItemWithAppointmentPatchRequestBody instantiates a new AppointmentsItemWithAppointmentPatchRequestBody and sets the default values.
func NewAppointmentsItemWithAppointmentPatchRequestBody() *AppointmentsItemWithAppointmentPatchRequestBody {
	m := &AppointmentsItemWithAppointmentPatchRequestBody{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateAppointmentsItemWithAppointmentPatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAppointmentsItemWithAppointmentPatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAppointmentsItemWithAppointmentPatchRequestBody(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AppointmentsItemWithAppointmentPatchRequestBody) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetContactId gets the contact_id property value. Required for pop-up reminders
// returns a *int64 when successful
func (m *AppointmentsItemWithAppointmentPatchRequestBody) GetContactId() *int64 {
	return m.contact_id
}

// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *AppointmentsItemWithAppointmentPatchRequestBody) GetDescription() *string {
	return m.description
}

// GetEndDate gets the end_date property value. The end_date property
// returns a *Time when successful
func (m *AppointmentsItemWithAppointmentPatchRequestBody) GetEndDate() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.end_date
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AppointmentsItemWithAppointmentPatchRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["contact_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetContactId(val)
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
	res["end_date"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetTimeValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetEndDate(val)
		}
		return nil
	}
	res["location"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLocation(val)
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
	res["start_date"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetTimeValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetStartDate(val)
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
	res["user"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetUser(val)
		}
		return nil
	}
	return res
}

// GetLocation gets the location property value. The location property
// returns a *string when successful
func (m *AppointmentsItemWithAppointmentPatchRequestBody) GetLocation() *string {
	return m.location
}

// GetRemindTime gets the remind_time property value. Value in minutes before start_date to show pop-up reminder.  Acceptable values are [5,10,15,30,60,120,240,480,1440,2880]
// returns a *int32 when successful
func (m *AppointmentsItemWithAppointmentPatchRequestBody) GetRemindTime() *int32 {
	return m.remind_time
}

// GetStartDate gets the start_date property value. The start_date property
// returns a *Time when successful
func (m *AppointmentsItemWithAppointmentPatchRequestBody) GetStartDate() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.start_date
}

// GetTitle gets the title property value. The title property
// returns a *string when successful
func (m *AppointmentsItemWithAppointmentPatchRequestBody) GetTitle() *string {
	return m.title
}

// GetUser gets the user property value. Required only for pop-up reminders
// returns a *int64 when successful
func (m *AppointmentsItemWithAppointmentPatchRequestBody) GetUser() *int64 {
	return m.user
}

// Serialize serializes information the current object
func (m *AppointmentsItemWithAppointmentPatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt64Value("contact_id", m.GetContactId())
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
		err := writer.WriteTimeValue("end_date", m.GetEndDate())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("location", m.GetLocation())
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
		err := writer.WriteTimeValue("start_date", m.GetStartDate())
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
		err := writer.WriteInt64Value("user", m.GetUser())
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
func (m *AppointmentsItemWithAppointmentPatchRequestBody) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetContactId sets the contact_id property value. Required for pop-up reminders
func (m *AppointmentsItemWithAppointmentPatchRequestBody) SetContactId(value *int64) {
	m.contact_id = value
}

// SetDescription sets the description property value. The description property
func (m *AppointmentsItemWithAppointmentPatchRequestBody) SetDescription(value *string) {
	m.description = value
}

// SetEndDate sets the end_date property value. The end_date property
func (m *AppointmentsItemWithAppointmentPatchRequestBody) SetEndDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	m.end_date = value
}

// SetLocation sets the location property value. The location property
func (m *AppointmentsItemWithAppointmentPatchRequestBody) SetLocation(value *string) {
	m.location = value
}

// SetRemindTime sets the remind_time property value. Value in minutes before start_date to show pop-up reminder.  Acceptable values are [5,10,15,30,60,120,240,480,1440,2880]
func (m *AppointmentsItemWithAppointmentPatchRequestBody) SetRemindTime(value *int32) {
	m.remind_time = value
}

// SetStartDate sets the start_date property value. The start_date property
func (m *AppointmentsItemWithAppointmentPatchRequestBody) SetStartDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	m.start_date = value
}

// SetTitle sets the title property value. The title property
func (m *AppointmentsItemWithAppointmentPatchRequestBody) SetTitle(value *string) {
	m.title = value
}

// SetUser sets the user property value. Required only for pop-up reminders
func (m *AppointmentsItemWithAppointmentPatchRequestBody) SetUser(value *int64) {
	m.user = value
}

type AppointmentsItemWithAppointmentPatchRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetContactId() *int64
	GetDescription() *string
	GetEndDate() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetLocation() *string
	GetRemindTime() *int32
	GetStartDate() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetTitle() *string
	GetUser() *int64
	SetContactId(value *int64)
	SetDescription(value *string)
	SetEndDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	SetLocation(value *string)
	SetRemindTime(value *int32)
	SetStartDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	SetTitle(value *string)
	SetUser(value *int64)
}
