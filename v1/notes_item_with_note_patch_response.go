package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
)

type NotesItemWithNotePatchResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The body property
	body *string
	// The contact_id property
	contact_id *int64
	// The date_created property
	date_created *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	// The id property
	id *int64
	// The last_updated property
	last_updated *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	// The last_updated_by property
	last_updated_by NotesItemWithNotePatchResponse_last_updated_byable
	// The title property
	title *string
	// The user_id property
	user_id *int64
}

// NewNotesItemWithNotePatchResponse instantiates a new NotesItemWithNotePatchResponse and sets the default values.
func NewNotesItemWithNotePatchResponse() *NotesItemWithNotePatchResponse {
	m := &NotesItemWithNotePatchResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateNotesItemWithNotePatchResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNotesItemWithNotePatchResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewNotesItemWithNotePatchResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NotesItemWithNotePatchResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetBody gets the body property value. The body property
// returns a *string when successful
func (m *NotesItemWithNotePatchResponse) GetBody() *string {
	return m.body
}

// GetContactId gets the contact_id property value. The contact_id property
// returns a *int64 when successful
func (m *NotesItemWithNotePatchResponse) GetContactId() *int64 {
	return m.contact_id
}

// GetDateCreated gets the date_created property value. The date_created property
// returns a *Time when successful
func (m *NotesItemWithNotePatchResponse) GetDateCreated() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.date_created
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NotesItemWithNotePatchResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["body"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetBody(val)
		}
		return nil
	}
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
	res["date_created"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetTimeValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDateCreated(val)
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
	res["last_updated"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetTimeValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLastUpdated(val)
		}
		return nil
	}
	res["last_updated_by"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateNotesItemWithNotePatchResponse_last_updated_byFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLastUpdatedBy(val.(NotesItemWithNotePatchResponse_last_updated_byable))
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

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *NotesItemWithNotePatchResponse) GetId() *int64 {
	return m.id
}

// GetLastUpdated gets the last_updated property value. The last_updated property
// returns a *Time when successful
func (m *NotesItemWithNotePatchResponse) GetLastUpdated() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.last_updated
}

// GetLastUpdatedBy gets the last_updated_by property value. The last_updated_by property
// returns a NotesItemWithNotePatchResponse_last_updated_byable when successful
func (m *NotesItemWithNotePatchResponse) GetLastUpdatedBy() NotesItemWithNotePatchResponse_last_updated_byable {
	return m.last_updated_by
}

// GetTitle gets the title property value. The title property
// returns a *string when successful
func (m *NotesItemWithNotePatchResponse) GetTitle() *string {
	return m.title
}

// GetUserId gets the user_id property value. The user_id property
// returns a *int64 when successful
func (m *NotesItemWithNotePatchResponse) GetUserId() *int64 {
	return m.user_id
}

// Serialize serializes information the current object
func (m *NotesItemWithNotePatchResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("body", m.GetBody())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("contact_id", m.GetContactId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteTimeValue("date_created", m.GetDateCreated())
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
		err := writer.WriteTimeValue("last_updated", m.GetLastUpdated())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("last_updated_by", m.GetLastUpdatedBy())
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
func (m *NotesItemWithNotePatchResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetBody sets the body property value. The body property
func (m *NotesItemWithNotePatchResponse) SetBody(value *string) {
	m.body = value
}

// SetContactId sets the contact_id property value. The contact_id property
func (m *NotesItemWithNotePatchResponse) SetContactId(value *int64) {
	m.contact_id = value
}

// SetDateCreated sets the date_created property value. The date_created property
func (m *NotesItemWithNotePatchResponse) SetDateCreated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	m.date_created = value
}

// SetId sets the id property value. The id property
func (m *NotesItemWithNotePatchResponse) SetId(value *int64) {
	m.id = value
}

// SetLastUpdated sets the last_updated property value. The last_updated property
func (m *NotesItemWithNotePatchResponse) SetLastUpdated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	m.last_updated = value
}

// SetLastUpdatedBy sets the last_updated_by property value. The last_updated_by property
func (m *NotesItemWithNotePatchResponse) SetLastUpdatedBy(value NotesItemWithNotePatchResponse_last_updated_byable) {
	m.last_updated_by = value
}

// SetTitle sets the title property value. The title property
func (m *NotesItemWithNotePatchResponse) SetTitle(value *string) {
	m.title = value
}

// SetUserId sets the user_id property value. The user_id property
func (m *NotesItemWithNotePatchResponse) SetUserId(value *int64) {
	m.user_id = value
}

type NotesItemWithNotePatchResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetBody() *string
	GetContactId() *int64
	GetDateCreated() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetId() *int64
	GetLastUpdated() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetLastUpdatedBy() NotesItemWithNotePatchResponse_last_updated_byable
	GetTitle() *string
	GetUserId() *int64
	SetBody(value *string)
	SetContactId(value *int64)
	SetDateCreated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	SetId(value *int64)
	SetLastUpdated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	SetLastUpdatedBy(value NotesItemWithNotePatchResponse_last_updated_byable)
	SetTitle(value *string)
	SetUserId(value *int64)
}
