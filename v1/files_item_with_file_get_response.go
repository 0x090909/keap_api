package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FilesItemWithFileGetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The file_data property
	file_data *string
	// The file_descriptor property
	file_descriptor FilesItemWithFileGetResponse_file_descriptorable
}

// NewFilesItemWithFileGetResponse instantiates a new FilesItemWithFileGetResponse and sets the default values.
func NewFilesItemWithFileGetResponse() *FilesItemWithFileGetResponse {
	m := &FilesItemWithFileGetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateFilesItemWithFileGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFilesItemWithFileGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewFilesItemWithFileGetResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *FilesItemWithFileGetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FilesItemWithFileGetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
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
	res["file_descriptor"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateFilesItemWithFileGetResponse_file_descriptorFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFileDescriptor(val.(FilesItemWithFileGetResponse_file_descriptorable))
		}
		return nil
	}
	return res
}

// GetFileData gets the file_data property value. The file_data property
// returns a *string when successful
func (m *FilesItemWithFileGetResponse) GetFileData() *string {
	return m.file_data
}

// GetFileDescriptor gets the file_descriptor property value. The file_descriptor property
// returns a FilesItemWithFileGetResponse_file_descriptorable when successful
func (m *FilesItemWithFileGetResponse) GetFileDescriptor() FilesItemWithFileGetResponse_file_descriptorable {
	return m.file_descriptor
}

// Serialize serializes information the current object
func (m *FilesItemWithFileGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("file_data", m.GetFileData())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("file_descriptor", m.GetFileDescriptor())
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
func (m *FilesItemWithFileGetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetFileData sets the file_data property value. The file_data property
func (m *FilesItemWithFileGetResponse) SetFileData(value *string) {
	m.file_data = value
}

// SetFileDescriptor sets the file_descriptor property value. The file_descriptor property
func (m *FilesItemWithFileGetResponse) SetFileDescriptor(value FilesItemWithFileGetResponse_file_descriptorable) {
	m.file_descriptor = value
}

type FilesItemWithFileGetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetFileData() *string
	GetFileDescriptor() FilesItemWithFileGetResponse_file_descriptorable
	SetFileData(value *string)
	SetFileDescriptor(value FilesItemWithFileGetResponse_file_descriptorable)
}
