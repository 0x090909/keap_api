package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ContactsItemNotesItemWithNote_PatchRequestBody struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The contact_id property
	contact_id *string
	// The text property
	text *string
	// A value for either `title` or `type` is required.
	title *string
	// A value for either `title` or `type` is required. The value may be one of `Appointment`, `Call`, `Email`, `Fax`, `Letter` or `Other` in Keap Max/Pro, or an admin-configured value in Classic.
	typeEscaped *string
	// The user_id property
	user_id *string
}

// NewContactsItemNotesItemWithNote_PatchRequestBody instantiates a new ContactsItemNotesItemWithNote_PatchRequestBody and sets the default values.
func NewContactsItemNotesItemWithNote_PatchRequestBody() *ContactsItemNotesItemWithNote_PatchRequestBody {
	m := &ContactsItemNotesItemWithNote_PatchRequestBody{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateContactsItemNotesItemWithNote_PatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsItemNotesItemWithNote_PatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsItemNotesItemWithNote_PatchRequestBody(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContactsItemNotesItemWithNote_PatchRequestBody) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetContactId gets the contact_id property value. The contact_id property
// returns a *string when successful
func (m *ContactsItemNotesItemWithNote_PatchRequestBody) GetContactId() *string {
	return m.contact_id
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContactsItemNotesItemWithNote_PatchRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
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
	res["user_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
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

// GetText gets the text property value. The text property
// returns a *string when successful
func (m *ContactsItemNotesItemWithNote_PatchRequestBody) GetText() *string {
	return m.text
}

// GetTitle gets the title property value. A value for either `title` or `type` is required.
// returns a *string when successful
func (m *ContactsItemNotesItemWithNote_PatchRequestBody) GetTitle() *string {
	return m.title
}

// GetTypeEscaped gets the type property value. A value for either `title` or `type` is required. The value may be one of `Appointment`, `Call`, `Email`, `Fax`, `Letter` or `Other` in Keap Max/Pro, or an admin-configured value in Classic.
// returns a *string when successful
func (m *ContactsItemNotesItemWithNote_PatchRequestBody) GetTypeEscaped() *string {
	return m.typeEscaped
}

// GetUserId gets the user_id property value. The user_id property
// returns a *string when successful
func (m *ContactsItemNotesItemWithNote_PatchRequestBody) GetUserId() *string {
	return m.user_id
}

// Serialize serializes information the current object
func (m *ContactsItemNotesItemWithNote_PatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("contact_id", m.GetContactId())
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
		err := writer.WriteStringValue("user_id", m.GetUserId())
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
func (m *ContactsItemNotesItemWithNote_PatchRequestBody) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetContactId sets the contact_id property value. The contact_id property
func (m *ContactsItemNotesItemWithNote_PatchRequestBody) SetContactId(value *string) {
	m.contact_id = value
}

// SetText sets the text property value. The text property
func (m *ContactsItemNotesItemWithNote_PatchRequestBody) SetText(value *string) {
	m.text = value
}

// SetTitle sets the title property value. A value for either `title` or `type` is required.
func (m *ContactsItemNotesItemWithNote_PatchRequestBody) SetTitle(value *string) {
	m.title = value
}

// SetTypeEscaped sets the type property value. A value for either `title` or `type` is required. The value may be one of `Appointment`, `Call`, `Email`, `Fax`, `Letter` or `Other` in Keap Max/Pro, or an admin-configured value in Classic.
func (m *ContactsItemNotesItemWithNote_PatchRequestBody) SetTypeEscaped(value *string) {
	m.typeEscaped = value
}

// SetUserId sets the user_id property value. The user_id property
func (m *ContactsItemNotesItemWithNote_PatchRequestBody) SetUserId(value *string) {
	m.user_id = value
}

type ContactsItemNotesItemWithNote_PatchRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetContactId() *string
	GetText() *string
	GetTitle() *string
	GetTypeEscaped() *string
	GetUserId() *string
	SetContactId(value *string)
	SetText(value *string)
	SetTitle(value *string)
	SetTypeEscaped(value *string)
	SetUserId(value *string)
}
