package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CompaniesPostRequestBody_fax_number struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The number property
	number *string
	// The type property
	typeEscaped *string
}

// NewCompaniesPostRequestBody_fax_number instantiates a new CompaniesPostRequestBody_fax_number and sets the default values.
func NewCompaniesPostRequestBody_fax_number() *CompaniesPostRequestBody_fax_number {
	m := &CompaniesPostRequestBody_fax_number{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateCompaniesPostRequestBody_fax_numberFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCompaniesPostRequestBody_fax_numberFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCompaniesPostRequestBody_fax_number(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CompaniesPostRequestBody_fax_number) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CompaniesPostRequestBody_fax_number) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
func (m *CompaniesPostRequestBody_fax_number) GetNumber() *string {
	return m.number
}

// GetTypeEscaped gets the type property value. The type property
// returns a *string when successful
func (m *CompaniesPostRequestBody_fax_number) GetTypeEscaped() *string {
	return m.typeEscaped
}

// Serialize serializes information the current object
func (m *CompaniesPostRequestBody_fax_number) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *CompaniesPostRequestBody_fax_number) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetNumber sets the number property value. The number property
func (m *CompaniesPostRequestBody_fax_number) SetNumber(value *string) {
	m.number = value
}

// SetTypeEscaped sets the type property value. The type property
func (m *CompaniesPostRequestBody_fax_number) SetTypeEscaped(value *string) {
	m.typeEscaped = value
}

type CompaniesPostRequestBody_fax_numberable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetNumber() *string
	GetTypeEscaped() *string
	SetNumber(value *string)
	SetTypeEscaped(value *string)
}
