package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type UsersGetResponse_users_fax_numbers struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The number property
	number *string
	// The type property
	typeEscaped *string
}

// NewUsersGetResponse_users_fax_numbers instantiates a new UsersGetResponse_users_fax_numbers and sets the default values.
func NewUsersGetResponse_users_fax_numbers() *UsersGetResponse_users_fax_numbers {
	m := &UsersGetResponse_users_fax_numbers{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateUsersGetResponse_users_fax_numbersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUsersGetResponse_users_fax_numbersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewUsersGetResponse_users_fax_numbers(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *UsersGetResponse_users_fax_numbers) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UsersGetResponse_users_fax_numbers) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["number"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNumber(val)
		}
		return nil
	}
	res["type"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTypeEscaped(val)
		}
		return nil
	}
	return res
}

// GetNumber gets the number property value. The number property
// returns a *string when successful
func (m *UsersGetResponse_users_fax_numbers) GetNumber() *string {
	return m.number
}

// GetTypeEscaped gets the type property value. The type property
// returns a *string when successful
func (m *UsersGetResponse_users_fax_numbers) GetTypeEscaped() *string {
	return m.typeEscaped
}

// Serialize serializes information the current object
func (m *UsersGetResponse_users_fax_numbers) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("number", m.GetNumber())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("type", m.GetTypeEscaped())
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
func (m *UsersGetResponse_users_fax_numbers) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetNumber sets the number property value. The number property
func (m *UsersGetResponse_users_fax_numbers) SetNumber(value *string) {
	m.number = value
}

// SetTypeEscaped sets the type property value. The type property
func (m *UsersGetResponse_users_fax_numbers) SetTypeEscaped(value *string) {
	m.typeEscaped = value
}

type UsersGetResponse_users_fax_numbersable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetNumber() *string
	GetTypeEscaped() *string
	SetNumber(value *string)
	SetTypeEscaped(value *string)
}
