package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FilesPostRequestBody struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The contact_id property
	contact_id *int64
	// The file_data property
	file_data *string
	// The file_name property
	file_name *string
	// The is_public property
	is_public *bool
	// The public property
	public *bool
}

// NewFilesPostRequestBody instantiates a new FilesPostRequestBody and sets the default values.
func NewFilesPostRequestBody() *FilesPostRequestBody {
	m := &FilesPostRequestBody{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateFilesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFilesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewFilesPostRequestBody(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *FilesPostRequestBody) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetContactId gets the contact_id property value. The contact_id property
// returns a *int64 when successful
func (m *FilesPostRequestBody) GetContactId() *int64 {
	return m.contact_id
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FilesPostRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
	res["file_data"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFileData(val)
		}
		return nil
	}
	res["file_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFileName(val)
		}
		return nil
	}
	res["is_public"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetIsPublic(val)
		}
		return nil
	}
	res["public"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPublic(val)
		}
		return nil
	}
	return res
}

// GetFileData gets the file_data property value. The file_data property
// returns a *string when successful
func (m *FilesPostRequestBody) GetFileData() *string {
	return m.file_data
}

// GetFileName gets the file_name property value. The file_name property
// returns a *string when successful
func (m *FilesPostRequestBody) GetFileName() *string {
	return m.file_name
}

// GetIsPublic gets the is_public property value. The is_public property
// returns a *bool when successful
func (m *FilesPostRequestBody) GetIsPublic() *bool {
	return m.is_public
}

// GetPublic gets the public property value. The public property
// returns a *bool when successful
func (m *FilesPostRequestBody) GetPublic() *bool {
	return m.public
}

// Serialize serializes information the current object
func (m *FilesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt64Value("contact_id", m.GetContactId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("file_data", m.GetFileData())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("file_name", m.GetFileName())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("is_public", m.GetIsPublic())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("public", m.GetPublic())
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
func (m *FilesPostRequestBody) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetContactId sets the contact_id property value. The contact_id property
func (m *FilesPostRequestBody) SetContactId(value *int64) {
	m.contact_id = value
}

// SetFileData sets the file_data property value. The file_data property
func (m *FilesPostRequestBody) SetFileData(value *string) {
	m.file_data = value
}

// SetFileName sets the file_name property value. The file_name property
func (m *FilesPostRequestBody) SetFileName(value *string) {
	m.file_name = value
}

// SetIsPublic sets the is_public property value. The is_public property
func (m *FilesPostRequestBody) SetIsPublic(value *bool) {
	m.is_public = value
}

// SetPublic sets the public property value. The public property
func (m *FilesPostRequestBody) SetPublic(value *bool) {
	m.public = value
}

type FilesPostRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetContactId() *int64
	GetFileData() *string
	GetFileName() *string
	GetIsPublic() *bool
	GetPublic() *bool
	SetContactId(value *int64)
	SetFileData(value *string)
	SetFileName(value *string)
	SetIsPublic(value *bool)
	SetPublic(value *bool)
}
