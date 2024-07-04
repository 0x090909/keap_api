package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EmailAddressesItemWithEmailPutRequestBody struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The opted_in property
	opted_in *bool
	// The reason property
	reason *string
}

// NewEmailAddressesItemWithEmailPutRequestBody instantiates a new EmailAddressesItemWithEmailPutRequestBody and sets the default values.
func NewEmailAddressesItemWithEmailPutRequestBody() *EmailAddressesItemWithEmailPutRequestBody {
	m := &EmailAddressesItemWithEmailPutRequestBody{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateEmailAddressesItemWithEmailPutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEmailAddressesItemWithEmailPutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewEmailAddressesItemWithEmailPutRequestBody(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EmailAddressesItemWithEmailPutRequestBody) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EmailAddressesItemWithEmailPutRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["opted_in"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetOptedIn(val)
		}
		return nil
	}
	res["reason"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetReason(val)
		}
		return nil
	}
	return res
}

// GetOptedIn gets the opted_in property value. The opted_in property
// returns a *bool when successful
func (m *EmailAddressesItemWithEmailPutRequestBody) GetOptedIn() *bool {
	return m.opted_in
}

// GetReason gets the reason property value. The reason property
// returns a *string when successful
func (m *EmailAddressesItemWithEmailPutRequestBody) GetReason() *string {
	return m.reason
}

// Serialize serializes information the current object
func (m *EmailAddressesItemWithEmailPutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteBoolValue("opted_in", m.GetOptedIn())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("reason", m.GetReason())
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
func (m *EmailAddressesItemWithEmailPutRequestBody) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetOptedIn sets the opted_in property value. The opted_in property
func (m *EmailAddressesItemWithEmailPutRequestBody) SetOptedIn(value *bool) {
	m.opted_in = value
}

// SetReason sets the reason property value. The reason property
func (m *EmailAddressesItemWithEmailPutRequestBody) SetReason(value *string) {
	m.reason = value
}

type EmailAddressesItemWithEmailPutRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetOptedIn() *bool
	GetReason() *string
	SetOptedIn(value *bool)
	SetReason(value *string)
}
