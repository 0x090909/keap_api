package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ContactsItemWithContact_GetResponse_email_addresses struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The email property
	email *string
	// The is_opt_in property
	is_opt_in *bool
	// The opt_in_reason property
	opt_in_reason *string
}

// NewContactsItemWithContact_GetResponse_email_addresses instantiates a new ContactsItemWithContact_GetResponse_email_addresses and sets the default values.
func NewContactsItemWithContact_GetResponse_email_addresses() *ContactsItemWithContact_GetResponse_email_addresses {
	m := &ContactsItemWithContact_GetResponse_email_addresses{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateContactsItemWithContact_GetResponse_email_addressesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsItemWithContact_GetResponse_email_addressesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsItemWithContact_GetResponse_email_addresses(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContactsItemWithContact_GetResponse_email_addresses) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetEmail gets the email property value. The email property
// returns a *string when successful
func (m *ContactsItemWithContact_GetResponse_email_addresses) GetEmail() *string {
	return m.email
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContactsItemWithContact_GetResponse_email_addresses) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
	res["is_opt_in"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetIsOptIn(val)
		}
		return nil
	}
	res["opt_in_reason"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetOptInReason(val)
		}
		return nil
	}
	return res
}

// GetIsOptIn gets the is_opt_in property value. The is_opt_in property
// returns a *bool when successful
func (m *ContactsItemWithContact_GetResponse_email_addresses) GetIsOptIn() *bool {
	return m.is_opt_in
}

// GetOptInReason gets the opt_in_reason property value. The opt_in_reason property
// returns a *string when successful
func (m *ContactsItemWithContact_GetResponse_email_addresses) GetOptInReason() *string {
	return m.opt_in_reason
}

// Serialize serializes information the current object
func (m *ContactsItemWithContact_GetResponse_email_addresses) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("email", m.GetEmail())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("is_opt_in", m.GetIsOptIn())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("opt_in_reason", m.GetOptInReason())
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
func (m *ContactsItemWithContact_GetResponse_email_addresses) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetEmail sets the email property value. The email property
func (m *ContactsItemWithContact_GetResponse_email_addresses) SetEmail(value *string) {
	m.email = value
}

// SetIsOptIn sets the is_opt_in property value. The is_opt_in property
func (m *ContactsItemWithContact_GetResponse_email_addresses) SetIsOptIn(value *bool) {
	m.is_opt_in = value
}

// SetOptInReason sets the opt_in_reason property value. The opt_in_reason property
func (m *ContactsItemWithContact_GetResponse_email_addresses) SetOptInReason(value *string) {
	m.opt_in_reason = value
}

type ContactsItemWithContact_GetResponse_email_addressesable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetEmail() *string
	GetIsOptIn() *bool
	GetOptInReason() *string
	SetEmail(value *string)
	SetIsOptIn(value *bool)
	SetOptInReason(value *string)
}
