package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EmailsQueuePostRequestBody_attachments struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The content of the attachment, encoded in Base64
	file_data *string
	// The filename of the attached file, including extension
	file_name *string
}

// NewEmailsQueuePostRequestBody_attachments instantiates a new EmailsQueuePostRequestBody_attachments and sets the default values.
func NewEmailsQueuePostRequestBody_attachments() *EmailsQueuePostRequestBody_attachments {
	m := &EmailsQueuePostRequestBody_attachments{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateEmailsQueuePostRequestBody_attachmentsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEmailsQueuePostRequestBody_attachmentsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewEmailsQueuePostRequestBody_attachments(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EmailsQueuePostRequestBody_attachments) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EmailsQueuePostRequestBody_attachments) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
	return res
}

// GetFileData gets the file_data property value. The content of the attachment, encoded in Base64
// returns a *string when successful
func (m *EmailsQueuePostRequestBody_attachments) GetFileData() *string {
	return m.file_data
}

// GetFileName gets the file_name property value. The filename of the attached file, including extension
// returns a *string when successful
func (m *EmailsQueuePostRequestBody_attachments) GetFileName() *string {
	return m.file_name
}

// Serialize serializes information the current object
func (m *EmailsQueuePostRequestBody_attachments) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
		err := writer.WriteAdditionalData(m.GetAdditionalData())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EmailsQueuePostRequestBody_attachments) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetFileData sets the file_data property value. The content of the attachment, encoded in Base64
func (m *EmailsQueuePostRequestBody_attachments) SetFileData(value *string) {
	m.file_data = value
}

// SetFileName sets the file_name property value. The filename of the attached file, including extension
func (m *EmailsQueuePostRequestBody_attachments) SetFileName(value *string) {
	m.file_name = value
}

type EmailsQueuePostRequestBody_attachmentsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetFileData() *string
	GetFileName() *string
	SetFileData(value *string)
	SetFileName(value *string)
}
