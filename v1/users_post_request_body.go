package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type UsersPostRequestBody struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The admin property
	admin *bool
	// The email property
	email *string
	// The given_name property
	given_name *string
	// The partner property
	partner *bool
}

// NewUsersPostRequestBody instantiates a new UsersPostRequestBody and sets the default values.
func NewUsersPostRequestBody() *UsersPostRequestBody {
	m := &UsersPostRequestBody{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateUsersPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUsersPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewUsersPostRequestBody(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *UsersPostRequestBody) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAdmin gets the admin property value. The admin property
// returns a *bool when successful
func (m *UsersPostRequestBody) GetAdmin() *bool {
	return m.admin
}

// GetEmail gets the email property value. The email property
// returns a *string when successful
func (m *UsersPostRequestBody) GetEmail() *string {
	return m.email
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UsersPostRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["admin"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAdmin(val)
		}
		return nil
	}
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
	res["partner"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPartner(val)
		}
		return nil
	}
	return res
}

// GetGivenName gets the given_name property value. The given_name property
// returns a *string when successful
func (m *UsersPostRequestBody) GetGivenName() *string {
	return m.given_name
}

// GetPartner gets the partner property value. The partner property
// returns a *bool when successful
func (m *UsersPostRequestBody) GetPartner() *bool {
	return m.partner
}

// Serialize serializes information the current object
func (m *UsersPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteBoolValue("admin", m.GetAdmin())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("email", m.GetEmail())
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
		err := writer.WriteBoolValue("partner", m.GetPartner())
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
func (m *UsersPostRequestBody) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAdmin sets the admin property value. The admin property
func (m *UsersPostRequestBody) SetAdmin(value *bool) {
	m.admin = value
}

// SetEmail sets the email property value. The email property
func (m *UsersPostRequestBody) SetEmail(value *string) {
	m.email = value
}

// SetGivenName sets the given_name property value. The given_name property
func (m *UsersPostRequestBody) SetGivenName(value *string) {
	m.given_name = value
}

// SetPartner sets the partner property value. The partner property
func (m *UsersPostRequestBody) SetPartner(value *bool) {
	m.partner = value
}

type UsersPostRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAdmin() *bool
	GetEmail() *string
	GetGivenName() *string
	GetPartner() *bool
	SetAdmin(value *bool)
	SetEmail(value *string)
	SetGivenName(value *string)
	SetPartner(value *bool)
}
