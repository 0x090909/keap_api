package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ContactsItemCreditCardsPostRequestBody struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The address property
	address ContactsItemCreditCardsPostRequestBody_addressable
	// The card_number property
	card_number *string
	// The card_type property
	card_type *string
	// The email_address property
	email_address *string
	// The expiration_month property
	expiration_month *string
	// The expiration_year property
	expiration_year *string
	// The maestro_issue_number property
	maestro_issue_number *string
	// The maestro_start_date_month property
	maestro_start_date_month *string
	// The maestro_start_date_year property
	maestro_start_date_year *string
	// The name_on_card property
	name_on_card *string
	// The verification_code property
	verification_code *string
}

// NewContactsItemCreditCardsPostRequestBody instantiates a new ContactsItemCreditCardsPostRequestBody and sets the default values.
func NewContactsItemCreditCardsPostRequestBody() *ContactsItemCreditCardsPostRequestBody {
	m := &ContactsItemCreditCardsPostRequestBody{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateContactsItemCreditCardsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsItemCreditCardsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsItemCreditCardsPostRequestBody(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContactsItemCreditCardsPostRequestBody) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAddress gets the address property value. The address property
// returns a ContactsItemCreditCardsPostRequestBody_addressable when successful
func (m *ContactsItemCreditCardsPostRequestBody) GetAddress() ContactsItemCreditCardsPostRequestBody_addressable {
	return m.address
}

// GetCardNumber gets the card_number property value. The card_number property
// returns a *string when successful
func (m *ContactsItemCreditCardsPostRequestBody) GetCardNumber() *string {
	return m.card_number
}

// GetCardType gets the card_type property value. The card_type property
// returns a *string when successful
func (m *ContactsItemCreditCardsPostRequestBody) GetCardType() *string {
	return m.card_type
}

// GetEmailAddress gets the email_address property value. The email_address property
// returns a *string when successful
func (m *ContactsItemCreditCardsPostRequestBody) GetEmailAddress() *string {
	return m.email_address
}

// GetExpirationMonth gets the expiration_month property value. The expiration_month property
// returns a *string when successful
func (m *ContactsItemCreditCardsPostRequestBody) GetExpirationMonth() *string {
	return m.expiration_month
}

// GetExpirationYear gets the expiration_year property value. The expiration_year property
// returns a *string when successful
func (m *ContactsItemCreditCardsPostRequestBody) GetExpirationYear() *string {
	return m.expiration_year
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContactsItemCreditCardsPostRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["address"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateContactsItemCreditCardsPostRequestBody_addressFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAddress(val.(ContactsItemCreditCardsPostRequestBody_addressable))
		}
		return nil
	}
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
	res["email_address"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetEmailAddress(val)
		}
		return nil
	}
	res["expiration_month"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetExpirationMonth(val)
		}
		return nil
	}
	res["expiration_year"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetExpirationYear(val)
		}
		return nil
	}
	res["maestro_issue_number"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetMaestroIssueNumber(val)
		}
		return nil
	}
	res["maestro_start_date_month"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetMaestroStartDateMonth(val)
		}
		return nil
	}
	res["maestro_start_date_year"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetMaestroStartDateYear(val)
		}
		return nil
	}
	res["name_on_card"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNameOnCard(val)
		}
		return nil
	}
	res["verification_code"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetVerificationCode(val)
		}
		return nil
	}
	return res
}

// GetMaestroIssueNumber gets the maestro_issue_number property value. The maestro_issue_number property
// returns a *string when successful
func (m *ContactsItemCreditCardsPostRequestBody) GetMaestroIssueNumber() *string {
	return m.maestro_issue_number
}

// GetMaestroStartDateMonth gets the maestro_start_date_month property value. The maestro_start_date_month property
// returns a *string when successful
func (m *ContactsItemCreditCardsPostRequestBody) GetMaestroStartDateMonth() *string {
	return m.maestro_start_date_month
}

// GetMaestroStartDateYear gets the maestro_start_date_year property value. The maestro_start_date_year property
// returns a *string when successful
func (m *ContactsItemCreditCardsPostRequestBody) GetMaestroStartDateYear() *string {
	return m.maestro_start_date_year
}

// GetNameOnCard gets the name_on_card property value. The name_on_card property
// returns a *string when successful
func (m *ContactsItemCreditCardsPostRequestBody) GetNameOnCard() *string {
	return m.name_on_card
}

// GetVerificationCode gets the verification_code property value. The verification_code property
// returns a *string when successful
func (m *ContactsItemCreditCardsPostRequestBody) GetVerificationCode() *string {
	return m.verification_code
}

// Serialize serializes information the current object
func (m *ContactsItemCreditCardsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("address", m.GetAddress())
		if err != nil {
			return err
		}
	}
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
		err := writer.WriteStringValue("email_address", m.GetEmailAddress())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("expiration_month", m.GetExpirationMonth())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("expiration_year", m.GetExpirationYear())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("maestro_issue_number", m.GetMaestroIssueNumber())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("maestro_start_date_month", m.GetMaestroStartDateMonth())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("maestro_start_date_year", m.GetMaestroStartDateYear())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("name_on_card", m.GetNameOnCard())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("verification_code", m.GetVerificationCode())
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
func (m *ContactsItemCreditCardsPostRequestBody) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAddress sets the address property value. The address property
func (m *ContactsItemCreditCardsPostRequestBody) SetAddress(value ContactsItemCreditCardsPostRequestBody_addressable) {
	m.address = value
}

// SetCardNumber sets the card_number property value. The card_number property
func (m *ContactsItemCreditCardsPostRequestBody) SetCardNumber(value *string) {
	m.card_number = value
}

// SetCardType sets the card_type property value. The card_type property
func (m *ContactsItemCreditCardsPostRequestBody) SetCardType(value *string) {
	m.card_type = value
}

// SetEmailAddress sets the email_address property value. The email_address property
func (m *ContactsItemCreditCardsPostRequestBody) SetEmailAddress(value *string) {
	m.email_address = value
}

// SetExpirationMonth sets the expiration_month property value. The expiration_month property
func (m *ContactsItemCreditCardsPostRequestBody) SetExpirationMonth(value *string) {
	m.expiration_month = value
}

// SetExpirationYear sets the expiration_year property value. The expiration_year property
func (m *ContactsItemCreditCardsPostRequestBody) SetExpirationYear(value *string) {
	m.expiration_year = value
}

// SetMaestroIssueNumber sets the maestro_issue_number property value. The maestro_issue_number property
func (m *ContactsItemCreditCardsPostRequestBody) SetMaestroIssueNumber(value *string) {
	m.maestro_issue_number = value
}

// SetMaestroStartDateMonth sets the maestro_start_date_month property value. The maestro_start_date_month property
func (m *ContactsItemCreditCardsPostRequestBody) SetMaestroStartDateMonth(value *string) {
	m.maestro_start_date_month = value
}

// SetMaestroStartDateYear sets the maestro_start_date_year property value. The maestro_start_date_year property
func (m *ContactsItemCreditCardsPostRequestBody) SetMaestroStartDateYear(value *string) {
	m.maestro_start_date_year = value
}

// SetNameOnCard sets the name_on_card property value. The name_on_card property
func (m *ContactsItemCreditCardsPostRequestBody) SetNameOnCard(value *string) {
	m.name_on_card = value
}

// SetVerificationCode sets the verification_code property value. The verification_code property
func (m *ContactsItemCreditCardsPostRequestBody) SetVerificationCode(value *string) {
	m.verification_code = value
}

type ContactsItemCreditCardsPostRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAddress() ContactsItemCreditCardsPostRequestBody_addressable
	GetCardNumber() *string
	GetCardType() *string
	GetEmailAddress() *string
	GetExpirationMonth() *string
	GetExpirationYear() *string
	GetMaestroIssueNumber() *string
	GetMaestroStartDateMonth() *string
	GetMaestroStartDateYear() *string
	GetNameOnCard() *string
	GetVerificationCode() *string
	SetAddress(value ContactsItemCreditCardsPostRequestBody_addressable)
	SetCardNumber(value *string)
	SetCardType(value *string)
	SetEmailAddress(value *string)
	SetExpirationMonth(value *string)
	SetExpirationYear(value *string)
	SetMaestroIssueNumber(value *string)
	SetMaestroStartDateMonth(value *string)
	SetMaestroStartDateYear(value *string)
	SetNameOnCard(value *string)
	SetVerificationCode(value *string)
}
