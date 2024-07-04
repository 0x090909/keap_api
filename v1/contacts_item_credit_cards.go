package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ContactsItemCreditCards struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The card_number property
	card_number *string
	// The card_type property
	card_type *string
	// The id property
	id *int64
	// The validation_status property
	validation_status *string
}

// NewContactsItemCreditCards instantiates a new ContactsItemCreditCards and sets the default values.
func NewContactsItemCreditCards() *ContactsItemCreditCards {
	m := &ContactsItemCreditCards{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateContactsItemCreditCardsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsItemCreditCardsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsItemCreditCards(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContactsItemCreditCards) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCardNumber gets the card_number property value. The card_number property
// returns a *string when successful
func (m *ContactsItemCreditCards) GetCardNumber() *string {
	return m.card_number
}

// GetCardType gets the card_type property value. The card_type property
// returns a *string when successful
func (m *ContactsItemCreditCards) GetCardType() *string {
	return m.card_type
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContactsItemCreditCards) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["card_number"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCardNumber(val)
		}
		return nil
	}
	res["card_type"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCardType(val)
		}
		return nil
	}
	res["id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetId(val)
		}
		return nil
	}
	res["validation_status"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetValidationStatus(val)
		}
		return nil
	}
	return res
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *ContactsItemCreditCards) GetId() *int64 {
	return m.id
}

// GetValidationStatus gets the validation_status property value. The validation_status property
// returns a *string when successful
func (m *ContactsItemCreditCards) GetValidationStatus() *string {
	return m.validation_status
}

// Serialize serializes information the current object
func (m *ContactsItemCreditCards) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("card_number", m.GetCardNumber())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("card_type", m.GetCardType())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("id", m.GetId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("validation_status", m.GetValidationStatus())
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
func (m *ContactsItemCreditCards) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCardNumber sets the card_number property value. The card_number property
func (m *ContactsItemCreditCards) SetCardNumber(value *string) {
	m.card_number = value
}

// SetCardType sets the card_type property value. The card_type property
func (m *ContactsItemCreditCards) SetCardType(value *string) {
	m.card_type = value
}

// SetId sets the id property value. The id property
func (m *ContactsItemCreditCards) SetId(value *int64) {
	m.id = value
}

// SetValidationStatus sets the validation_status property value. The validation_status property
func (m *ContactsItemCreditCards) SetValidationStatus(value *string) {
	m.validation_status = value
}

type ContactsItemCreditCardsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCardNumber() *string
	GetCardType() *string
	GetId() *int64
	GetValidationStatus() *string
	SetCardNumber(value *string)
	SetCardType(value *string)
	SetId(value *int64)
	SetValidationStatus(value *string)
}
