package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ContactsItemWithContact_PatchResponse_phone_numbers struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The extension property
	extension *string
	// The number property
	number *string
	// The type property
	typeEscaped *string
}

// NewContactsItemWithContact_PatchResponse_phone_numbers instantiates a new ContactsItemWithContact_PatchResponse_phone_numbers and sets the default values.
func NewContactsItemWithContact_PatchResponse_phone_numbers() *ContactsItemWithContact_PatchResponse_phone_numbers {
	m := &ContactsItemWithContact_PatchResponse_phone_numbers{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateContactsItemWithContact_PatchResponse_phone_numbersFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsItemWithContact_PatchResponse_phone_numbersFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsItemWithContact_PatchResponse_phone_numbers(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContactsItemWithContact_PatchResponse_phone_numbers) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetExtension gets the extension property value. The extension property
// returns a *string when successful
func (m *ContactsItemWithContact_PatchResponse_phone_numbers) GetExtension() *string {
	return m.extension
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContactsItemWithContact_PatchResponse_phone_numbers) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
func (m *ContactsItemWithContact_PatchResponse_phone_numbers) GetNumber() *string {
	return m.number
}

// GetTypeEscaped gets the type property value. The type property
// returns a *string when successful
func (m *ContactsItemWithContact_PatchResponse_phone_numbers) GetTypeEscaped() *string {
	return m.typeEscaped
}

// Serialize serializes information the current object
func (m *ContactsItemWithContact_PatchResponse_phone_numbers) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *ContactsItemWithContact_PatchResponse_phone_numbers) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetExtension sets the extension property value. The extension property
func (m *ContactsItemWithContact_PatchResponse_phone_numbers) SetExtension(value *string) {
	m.extension = value
}

// SetNumber sets the number property value. The number property
func (m *ContactsItemWithContact_PatchResponse_phone_numbers) SetNumber(value *string) {
	m.number = value
}

// SetTypeEscaped sets the type property value. The type property
func (m *ContactsItemWithContact_PatchResponse_phone_numbers) SetTypeEscaped(value *string) {
	m.typeEscaped = value
}

type ContactsItemWithContact_PatchResponse_phone_numbersable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetExtension() *string
	GetNumber() *string
	GetTypeEscaped() *string
	SetExtension(value *string)
	SetNumber(value *string)
	SetTypeEscaped(value *string)
}
