package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FilesItemWithFilePutResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The file_data property
	file_data *string
	// The file_descriptor property
	file_descriptor FilesItemWithFilePutResponse_file_descriptorable
}

// NewFilesItemWithFilePutResponse instantiates a new FilesItemWithFilePutResponse and sets the default values.
func NewFilesItemWithFilePutResponse() *FilesItemWithFilePutResponse {
	m := &FilesItemWithFilePutResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateFilesItemWithFilePutResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFilesItemWithFilePutResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewFilesItemWithFilePutResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *FilesItemWithFilePutResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FilesItemWithFilePutResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
		val, err := n.GetObjectValue(CreateFilesItemWithFilePutResponse_file_descriptorFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFileDescriptor(val.(FilesItemWithFilePutResponse_file_descriptorable))
		}
		return nil
	}
	return res
}

// GetFileData gets the file_data property value. The file_data property
// returns a *string when successful
func (m *FilesItemWithFilePutResponse) GetFileData() *string {
	return m.file_data
}

// GetFileDescriptor gets the file_descriptor property value. The file_descriptor property
// returns a FilesItemWithFilePutResponse_file_descriptorable when successful
func (m *FilesItemWithFilePutResponse) GetFileDescriptor() FilesItemWithFilePutResponse_file_descriptorable {
	return m.file_descriptor
}

// Serialize serializes information the current object
func (m *FilesItemWithFilePutResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *FilesItemWithFilePutResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetFileData sets the file_data property value. The file_data property
func (m *FilesItemWithFilePutResponse) SetFileData(value *string) {
	m.file_data = value
}

// SetFileDescriptor sets the file_descriptor property value. The file_descriptor property
func (m *FilesItemWithFilePutResponse) SetFileDescriptor(value FilesItemWithFilePutResponse_file_descriptorable) {
	m.file_descriptor = value
}

type FilesItemWithFilePutResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetFileData() *string
	GetFileDescriptor() FilesItemWithFilePutResponse_file_descriptorable
	SetFileData(value *string)
	SetFileDescriptor(value FilesItemWithFilePutResponse_file_descriptorable)
}
