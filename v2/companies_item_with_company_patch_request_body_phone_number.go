package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CompaniesItemWithCompany_PatchRequestBody_phone_number struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The extension property
	extension *string
	// The number property
	number *string
	// The type property
	typeEscaped *string
}

// NewCompaniesItemWithCompany_PatchRequestBody_phone_number instantiates a new CompaniesItemWithCompany_PatchRequestBody_phone_number and sets the default values.
func NewCompaniesItemWithCompany_PatchRequestBody_phone_number() *CompaniesItemWithCompany_PatchRequestBody_phone_number {
	m := &CompaniesItemWithCompany_PatchRequestBody_phone_number{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateCompaniesItemWithCompany_PatchRequestBody_phone_numberFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCompaniesItemWithCompany_PatchRequestBody_phone_numberFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCompaniesItemWithCompany_PatchRequestBody_phone_number(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CompaniesItemWithCompany_PatchRequestBody_phone_number) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetExtension gets the extension property value. The extension property
// returns a *string when successful
func (m *CompaniesItemWithCompany_PatchRequestBody_phone_number) GetExtension() *string {
	return m.extension
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CompaniesItemWithCompany_PatchRequestBody_phone_number) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["extension"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetExtension(val)
		}
		return nil
	}
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
func (m *CompaniesItemWithCompany_PatchRequestBody_phone_number) GetNumber() *string {
	return m.number
}

// GetTypeEscaped gets the type property value. The type property
// returns a *string when successful
func (m *CompaniesItemWithCompany_PatchRequestBody_phone_number) GetTypeEscaped() *string {
	return m.typeEscaped
}

// Serialize serializes information the current object
func (m *CompaniesItemWithCompany_PatchRequestBody_phone_number) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("extension", m.GetExtension())
		if err != nil {
			return err
		}
	}
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
func (m *CompaniesItemWithCompany_PatchRequestBody_phone_number) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetExtension sets the extension property value. The extension property
func (m *CompaniesItemWithCompany_PatchRequestBody_phone_number) SetExtension(value *string) {
	m.extension = value
}

// SetNumber sets the number property value. The number property
func (m *CompaniesItemWithCompany_PatchRequestBody_phone_number) SetNumber(value *string) {
	m.number = value
}

// SetTypeEscaped sets the type property value. The type property
func (m *CompaniesItemWithCompany_PatchRequestBody_phone_number) SetTypeEscaped(value *string) {
	m.typeEscaped = value
}

type CompaniesItemWithCompany_PatchRequestBody_phone_numberable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetExtension() *string
	GetNumber() *string
	GetTypeEscaped() *string
	SetExtension(value *string)
	SetNumber(value *string)
	SetTypeEscaped(value *string)
}
