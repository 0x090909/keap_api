package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EmailsBatchAddPostResponse_errors struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The email property
	email EmailsBatchAddPostResponse_errors_emailable
	// The error_message property
	error_message *string
}

// NewEmailsBatchAddPostResponse_errors instantiates a new EmailsBatchAddPostResponse_errors and sets the default values.
func NewEmailsBatchAddPostResponse_errors() *EmailsBatchAddPostResponse_errors {
	m := &EmailsBatchAddPostResponse_errors{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateEmailsBatchAddPostResponse_errorsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEmailsBatchAddPostResponse_errorsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewEmailsBatchAddPostResponse_errors(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EmailsBatchAddPostResponse_errors) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetEmail gets the email property value. The email property
// returns a EmailsBatchAddPostResponse_errors_emailable when successful
func (m *EmailsBatchAddPostResponse_errors) GetEmail() EmailsBatchAddPostResponse_errors_emailable {
	return m.email
}

// GetErrorMessage gets the error_message property value. The error_message property
// returns a *string when successful
func (m *EmailsBatchAddPostResponse_errors) GetErrorMessage() *string {
	return m.error_message
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EmailsBatchAddPostResponse_errors) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["email"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateEmailsBatchAddPostResponse_errors_emailFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetEmail(val.(EmailsBatchAddPostResponse_errors_emailable))
		}
		return nil
	}
	res["error_message"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetErrorMessage(val)
		}
		return nil
	}
	return res
}

// Serialize serializes information the current object
func (m *EmailsBatchAddPostResponse_errors) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("email", m.GetEmail())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("error_message", m.GetErrorMessage())
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
func (m *EmailsBatchAddPostResponse_errors) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetEmail sets the email property value. The email property
func (m *EmailsBatchAddPostResponse_errors) SetEmail(value EmailsBatchAddPostResponse_errors_emailable) {
	m.email = value
}

// SetErrorMessage sets the error_message property value. The error_message property
func (m *EmailsBatchAddPostResponse_errors) SetErrorMessage(value *string) {
	m.error_message = value
}

type EmailsBatchAddPostResponse_errorsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetEmail() EmailsBatchAddPostResponse_errors_emailable
	GetErrorMessage() *string
	SetEmail(value EmailsBatchAddPostResponse_errors_emailable)
	SetErrorMessage(value *string)
}
