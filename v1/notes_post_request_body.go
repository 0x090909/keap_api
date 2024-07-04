package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NotesPostRequestBody struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The body property
	body *string
	// The contact_id property
	contact_id *int64
	// The title property
	title *string
	// The user_id property
	user_id *int64
}

// NewNotesPostRequestBody instantiates a new NotesPostRequestBody and sets the default values.
func NewNotesPostRequestBody() *NotesPostRequestBody {
	m := &NotesPostRequestBody{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateNotesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNotesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewNotesPostRequestBody(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NotesPostRequestBody) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetBody gets the body property value. The body property
// returns a *string when successful
func (m *NotesPostRequestBody) GetBody() *string {
	return m.body
}

// GetContactId gets the contact_id property value. The contact_id property
// returns a *int64 when successful
func (m *NotesPostRequestBody) GetContactId() *int64 {
	return m.contact_id
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NotesPostRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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

// GetTitle gets the title property value. The title property
// returns a *string when successful
func (m *NotesPostRequestBody) GetTitle() *string {
	return m.title
}

// GetUserId gets the user_id property value. The user_id property
// returns a *int64 when successful
func (m *NotesPostRequestBody) GetUserId() *int64 {
	return m.user_id
}

// Serialize serializes information the current object
func (m *NotesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *NotesPostRequestBody) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetBody sets the body property value. The body property
func (m *NotesPostRequestBody) SetBody(value *string) {
	m.body = value
}

// SetContactId sets the contact_id property value. The contact_id property
func (m *NotesPostRequestBody) SetContactId(value *int64) {
	m.contact_id = value
}

// SetTitle sets the title property value. The title property
func (m *NotesPostRequestBody) SetTitle(value *string) {
	m.title = value
}

// SetUserId sets the user_id property value. The user_id property
func (m *NotesPostRequestBody) SetUserId(value *int64) {
	m.user_id = value
}

type NotesPostRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetBody() *string
	GetContactId() *int64
	GetTitle() *string
	GetUserId() *int64
	SetBody(value *string)
	SetContactId(value *int64)
	SetTitle(value *string)
	SetUserId(value *int64)
}
