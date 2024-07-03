package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ContactsItemNotesItemWithNote_PatchResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The assigned_to_user property
	assigned_to_user ContactsItemNotesItemWithNote_PatchResponse_assigned_to_userable
	// The contact_id property
	contact_id *string
	// The create_time property
	create_time *string
	// The id property
	id *string
	// The last_updated_by_user_id property
	last_updated_by_user_id *string
	// The text property
	text *string
	// The title property
	title *string
	// The type property
	typeEscaped *string
	// The update_time property
	update_time *string
}

// NewContactsItemNotesItemWithNote_PatchResponse instantiates a new ContactsItemNotesItemWithNote_PatchResponse and sets the default values.
func NewContactsItemNotesItemWithNote_PatchResponse() *ContactsItemNotesItemWithNote_PatchResponse {
	m := &ContactsItemNotesItemWithNote_PatchResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateContactsItemNotesItemWithNote_PatchResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsItemNotesItemWithNote_PatchResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsItemNotesItemWithNote_PatchResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContactsItemNotesItemWithNote_PatchResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAssignedToUser gets the assigned_to_user property value. The assigned_to_user property
// returns a ContactsItemNotesItemWithNote_PatchResponse_assigned_to_userable when successful
func (m *ContactsItemNotesItemWithNote_PatchResponse) GetAssignedToUser() ContactsItemNotesItemWithNote_PatchResponse_assigned_to_userable {
	return m.assigned_to_user
}

// GetContactId gets the contact_id property value. The contact_id property
// returns a *string when successful
func (m *ContactsItemNotesItemWithNote_PatchResponse) GetContactId() *string {
	return m.contact_id
}

// GetCreateTime gets the create_time property value. The create_time property
// returns a *string when successful
func (m *ContactsItemNotesItemWithNote_PatchResponse) GetCreateTime() *string {
	return m.create_time
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContactsItemNotesItemWithNote_PatchResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["assigned_to_user"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateContactsItemNotesItemWithNote_PatchResponse_assigned_to_userFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAssignedToUser(val.(ContactsItemNotesItemWithNote_PatchResponse_assigned_to_userable))
		}
		return nil
	}
	res["contact_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetContactId(val)
		}
		return nil
	}
	res["create_time"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCreateTime(val)
		}
		return nil
	}
	res["id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetId(val)
		}
		return nil
	}
	res["last_updated_by_user_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLastUpdatedByUserId(val)
		}
		return nil
	}
	res["text"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetText(val)
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
	res["update_time"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetUpdateTime(val)
		}
		return nil
	}
	return res
}

// GetId gets the id property value. The id property
// returns a *string when successful
func (m *ContactsItemNotesItemWithNote_PatchResponse) GetId() *string {
	return m.id
}

// GetLastUpdatedByUserId gets the last_updated_by_user_id property value. The last_updated_by_user_id property
// returns a *string when successful
func (m *ContactsItemNotesItemWithNote_PatchResponse) GetLastUpdatedByUserId() *string {
	return m.last_updated_by_user_id
}

// GetText gets the text property value. The text property
// returns a *string when successful
func (m *ContactsItemNotesItemWithNote_PatchResponse) GetText() *string {
	return m.text
}

// GetTitle gets the title property value. The title property
// returns a *string when successful
func (m *ContactsItemNotesItemWithNote_PatchResponse) GetTitle() *string {
	return m.title
}

// GetTypeEscaped gets the type property value. The type property
// returns a *string when successful
func (m *ContactsItemNotesItemWithNote_PatchResponse) GetTypeEscaped() *string {
	return m.typeEscaped
}

// GetUpdateTime gets the update_time property value. The update_time property
// returns a *string when successful
func (m *ContactsItemNotesItemWithNote_PatchResponse) GetUpdateTime() *string {
	return m.update_time
}

// Serialize serializes information the current object
func (m *ContactsItemNotesItemWithNote_PatchResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("assigned_to_user", m.GetAssignedToUser())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("contact_id", m.GetContactId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("create_time", m.GetCreateTime())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("id", m.GetId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("last_updated_by_user_id", m.GetLastUpdatedByUserId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("text", m.GetText())
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
		err := writer.WriteStringValue("update_time", m.GetUpdateTime())
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
func (m *ContactsItemNotesItemWithNote_PatchResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAssignedToUser sets the assigned_to_user property value. The assigned_to_user property
func (m *ContactsItemNotesItemWithNote_PatchResponse) SetAssignedToUser(value ContactsItemNotesItemWithNote_PatchResponse_assigned_to_userable) {
	m.assigned_to_user = value
}

// SetContactId sets the contact_id property value. The contact_id property
func (m *ContactsItemNotesItemWithNote_PatchResponse) SetContactId(value *string) {
	m.contact_id = value
}

// SetCreateTime sets the create_time property value. The create_time property
func (m *ContactsItemNotesItemWithNote_PatchResponse) SetCreateTime(value *string) {
	m.create_time = value
}

// SetId sets the id property value. The id property
func (m *ContactsItemNotesItemWithNote_PatchResponse) SetId(value *string) {
	m.id = value
}

// SetLastUpdatedByUserId sets the last_updated_by_user_id property value. The last_updated_by_user_id property
func (m *ContactsItemNotesItemWithNote_PatchResponse) SetLastUpdatedByUserId(value *string) {
	m.last_updated_by_user_id = value
}

// SetText sets the text property value. The text property
func (m *ContactsItemNotesItemWithNote_PatchResponse) SetText(value *string) {
	m.text = value
}

// SetTitle sets the title property value. The title property
func (m *ContactsItemNotesItemWithNote_PatchResponse) SetTitle(value *string) {
	m.title = value
}

// SetTypeEscaped sets the type property value. The type property
func (m *ContactsItemNotesItemWithNote_PatchResponse) SetTypeEscaped(value *string) {
	m.typeEscaped = value
}

// SetUpdateTime sets the update_time property value. The update_time property
func (m *ContactsItemNotesItemWithNote_PatchResponse) SetUpdateTime(value *string) {
	m.update_time = value
}

type ContactsItemNotesItemWithNote_PatchResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAssignedToUser() ContactsItemNotesItemWithNote_PatchResponse_assigned_to_userable
	GetContactId() *string
	GetCreateTime() *string
	GetId() *string
	GetLastUpdatedByUserId() *string
	GetText() *string
	GetTitle() *string
	GetTypeEscaped() *string
	GetUpdateTime() *string
	SetAssignedToUser(value ContactsItemNotesItemWithNote_PatchResponse_assigned_to_userable)
	SetContactId(value *string)
	SetCreateTime(value *string)
	SetId(value *string)
	SetLastUpdatedByUserId(value *string)
	SetText(value *string)
	SetTitle(value *string)
	SetTypeEscaped(value *string)
	SetUpdateTime(value *string)
}
