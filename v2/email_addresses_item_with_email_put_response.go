package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EmailAddressesItemWithEmailPutResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The email property
	email *string
	// The opted_in property
	opted_in *bool
}

// NewEmailAddressesItemWithEmailPutResponse instantiates a new EmailAddressesItemWithEmailPutResponse and sets the default values.
func NewEmailAddressesItemWithEmailPutResponse() *EmailAddressesItemWithEmailPutResponse {
	m := &EmailAddressesItemWithEmailPutResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateEmailAddressesItemWithEmailPutResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEmailAddressesItemWithEmailPutResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewEmailAddressesItemWithEmailPutResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EmailAddressesItemWithEmailPutResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetEmail gets the email property value. The email property
// returns a *string when successful
func (m *EmailAddressesItemWithEmailPutResponse) GetEmail() *string {
	return m.email
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EmailAddressesItemWithEmailPutResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["email"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetEmail(val)
		}
		return nil
	}
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
	return res
}

// GetOptedIn gets the opted_in property value. The opted_in property
// returns a *bool when successful
func (m *EmailAddressesItemWithEmailPutResponse) GetOptedIn() *bool {
	return m.opted_in
}

// Serialize serializes information the current object
func (m *EmailAddressesItemWithEmailPutResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("email", m.GetEmail())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("opted_in", m.GetOptedIn())
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
func (m *EmailAddressesItemWithEmailPutResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetEmail sets the email property value. The email property
func (m *EmailAddressesItemWithEmailPutResponse) SetEmail(value *string) {
	m.email = value
}

// SetOptedIn sets the opted_in property value. The opted_in property
func (m *EmailAddressesItemWithEmailPutResponse) SetOptedIn(value *bool) {
	m.opted_in = value
}

type EmailAddressesItemWithEmailPutResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetEmail() *string
	GetOptedIn() *bool
	SetEmail(value *string)
	SetOptedIn(value *bool)
}
