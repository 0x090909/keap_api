package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NotesItemWithNotePutResponse_last_updated_by struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The family_name property
	family_name *string
	// The given_name property
	given_name *string
	// The user_id property
	user_id *int64
}

// NewNotesItemWithNotePutResponse_last_updated_by instantiates a new NotesItemWithNotePutResponse_last_updated_by and sets the default values.
func NewNotesItemWithNotePutResponse_last_updated_by() *NotesItemWithNotePutResponse_last_updated_by {
	m := &NotesItemWithNotePutResponse_last_updated_by{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateNotesItemWithNotePutResponse_last_updated_byFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNotesItemWithNotePutResponse_last_updated_byFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewNotesItemWithNotePutResponse_last_updated_by(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *NotesItemWithNotePutResponse_last_updated_by) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFamilyName gets the family_name property value. The family_name property
// returns a *string when successful
func (m *NotesItemWithNotePutResponse_last_updated_by) GetFamilyName() *string {
	return m.family_name
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NotesItemWithNotePutResponse_last_updated_by) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["family_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFamilyName(val)
		}
		return nil
	}
	res["given_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetGivenName(val)
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

// GetGivenName gets the given_name property value. The given_name property
// returns a *string when successful
func (m *NotesItemWithNotePutResponse_last_updated_by) GetGivenName() *string {
	return m.given_name
}

// GetUserId gets the user_id property value. The user_id property
// returns a *int64 when successful
func (m *NotesItemWithNotePutResponse_last_updated_by) GetUserId() *int64 {
	return m.user_id
}

// Serialize serializes information the current object
func (m *NotesItemWithNotePutResponse_last_updated_by) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("family_name", m.GetFamilyName())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("given_name", m.GetGivenName())
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
func (m *NotesItemWithNotePutResponse_last_updated_by) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetFamilyName sets the family_name property value. The family_name property
func (m *NotesItemWithNotePutResponse_last_updated_by) SetFamilyName(value *string) {
	m.family_name = value
}

// SetGivenName sets the given_name property value. The given_name property
func (m *NotesItemWithNotePutResponse_last_updated_by) SetGivenName(value *string) {
	m.given_name = value
}

// SetUserId sets the user_id property value. The user_id property
func (m *NotesItemWithNotePutResponse_last_updated_by) SetUserId(value *int64) {
	m.user_id = value
}

type NotesItemWithNotePutResponse_last_updated_byable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetFamilyName() *string
	GetGivenName() *string
	GetUserId() *int64
	SetFamilyName(value *string)
	SetGivenName(value *string)
	SetUserId(value *int64)
}
