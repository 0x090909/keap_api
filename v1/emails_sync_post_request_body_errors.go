package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EmailsSyncPostRequestBody_errors struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The email property
	email EmailsSyncPostRequestBody_errors_emailable
	// The error_message property
	error_message *string
}

// NewEmailsSyncPostRequestBody_errors instantiates a new EmailsSyncPostRequestBody_errors and sets the default values.
func NewEmailsSyncPostRequestBody_errors() *EmailsSyncPostRequestBody_errors {
	m := &EmailsSyncPostRequestBody_errors{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateEmailsSyncPostRequestBody_errorsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEmailsSyncPostRequestBody_errorsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewEmailsSyncPostRequestBody_errors(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EmailsSyncPostRequestBody_errors) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetEmail gets the email property value. The email property
// returns a EmailsSyncPostRequestBody_errors_emailable when successful
func (m *EmailsSyncPostRequestBody_errors) GetEmail() EmailsSyncPostRequestBody_errors_emailable {
	return m.email
}

// GetErrorMessage gets the error_message property value. The error_message property
// returns a *string when successful
func (m *EmailsSyncPostRequestBody_errors) GetErrorMessage() *string {
	return m.error_message
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EmailsSyncPostRequestBody_errors) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["email"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateEmailsSyncPostRequestBody_errors_emailFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetEmail(val.(EmailsSyncPostRequestBody_errors_emailable))
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
func (m *EmailsSyncPostRequestBody_errors) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *EmailsSyncPostRequestBody_errors) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetEmail sets the email property value. The email property
func (m *EmailsSyncPostRequestBody_errors) SetEmail(value EmailsSyncPostRequestBody_errors_emailable) {
	m.email = value
}

// SetErrorMessage sets the error_message property value. The error_message property
func (m *EmailsSyncPostRequestBody_errors) SetErrorMessage(value *string) {
	m.error_message = value
}

type EmailsSyncPostRequestBody_errorsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetEmail() EmailsSyncPostRequestBody_errors_emailable
	GetErrorMessage() *string
	SetEmail(value EmailsSyncPostRequestBody_errors_emailable)
	SetErrorMessage(value *string)
}
