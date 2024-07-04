package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ContactsItemCreditCardsPostResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The address property
	address ContactsItemCreditCardsPostResponse_addressable
	// The card_type property
	card_type *string
	// The contact_id property
	contact_id *int64
	// The email_address property
	email_address *string
	// The expiration_month property
	expiration_month *string
	// The expiration_year property
	expiration_year *string
	// The id property
	id *int64
	// The maestro_issue_number property
	maestro_issue_number *string
	// The maestro_start_date_month property
	maestro_start_date_month *string
	// The maestro_start_date_year property
	maestro_start_date_year *string
	// The name_on_card property
	name_on_card *string
	// The validation_message property
	validation_message *string
	// The validation_status property
	validation_status *string
}

// NewContactsItemCreditCardsPostResponse instantiates a new ContactsItemCreditCardsPostResponse and sets the default values.
func NewContactsItemCreditCardsPostResponse() *ContactsItemCreditCardsPostResponse {
	m := &ContactsItemCreditCardsPostResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateContactsItemCreditCardsPostResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsItemCreditCardsPostResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsItemCreditCardsPostResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContactsItemCreditCardsPostResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAddress gets the address property value. The address property
// returns a ContactsItemCreditCardsPostResponse_addressable when successful
func (m *ContactsItemCreditCardsPostResponse) GetAddress() ContactsItemCreditCardsPostResponse_addressable {
	return m.address
}

// GetCardType gets the card_type property value. The card_type property
// returns a *string when successful
func (m *ContactsItemCreditCardsPostResponse) GetCardType() *string {
	return m.card_type
}

// GetContactId gets the contact_id property value. The contact_id property
// returns a *int64 when successful
func (m *ContactsItemCreditCardsPostResponse) GetContactId() *int64 {
	return m.contact_id
}

// GetEmailAddress gets the email_address property value. The email_address property
// returns a *string when successful
func (m *ContactsItemCreditCardsPostResponse) GetEmailAddress() *string {
	return m.email_address
}

// GetExpirationMonth gets the expiration_month property value. The expiration_month property
// returns a *string when successful
func (m *ContactsItemCreditCardsPostResponse) GetExpirationMonth() *string {
	return m.expiration_month
}

// GetExpirationYear gets the expiration_year property value. The expiration_year property
// returns a *string when successful
func (m *ContactsItemCreditCardsPostResponse) GetExpirationYear() *string {
	return m.expiration_year
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContactsItemCreditCardsPostResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["address"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateContactsItemCreditCardsPostResponse_addressFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAddress(val.(ContactsItemCreditCardsPostResponse_addressable))
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
	res["contact_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetContactId(val)
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
	res["validation_message"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetValidationMessage(val)
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
func (m *ContactsItemCreditCardsPostResponse) GetId() *int64 {
	return m.id
}

// GetMaestroIssueNumber gets the maestro_issue_number property value. The maestro_issue_number property
// returns a *string when successful
func (m *ContactsItemCreditCardsPostResponse) GetMaestroIssueNumber() *string {
	return m.maestro_issue_number
}

// GetMaestroStartDateMonth gets the maestro_start_date_month property value. The maestro_start_date_month property
// returns a *string when successful
func (m *ContactsItemCreditCardsPostResponse) GetMaestroStartDateMonth() *string {
	return m.maestro_start_date_month
}

// GetMaestroStartDateYear gets the maestro_start_date_year property value. The maestro_start_date_year property
// returns a *string when successful
func (m *ContactsItemCreditCardsPostResponse) GetMaestroStartDateYear() *string {
	return m.maestro_start_date_year
}

// GetNameOnCard gets the name_on_card property value. The name_on_card property
// returns a *string when successful
func (m *ContactsItemCreditCardsPostResponse) GetNameOnCard() *string {
	return m.name_on_card
}

// GetValidationMessage gets the validation_message property value. The validation_message property
// returns a *string when successful
func (m *ContactsItemCreditCardsPostResponse) GetValidationMessage() *string {
	return m.validation_message
}

// GetValidationStatus gets the validation_status property value. The validation_status property
// returns a *string when successful
func (m *ContactsItemCreditCardsPostResponse) GetValidationStatus() *string {
	return m.validation_status
}

// Serialize serializes information the current object
func (m *ContactsItemCreditCardsPostResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("address", m.GetAddress())
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
		err := writer.WriteInt64Value("contact_id", m.GetContactId())
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
		err := writer.WriteInt64Value("id", m.GetId())
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
		err := writer.WriteStringValue("validation_message", m.GetValidationMessage())
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
func (m *ContactsItemCreditCardsPostResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAddress sets the address property value. The address property
func (m *ContactsItemCreditCardsPostResponse) SetAddress(value ContactsItemCreditCardsPostResponse_addressable) {
	m.address = value
}

// SetCardType sets the card_type property value. The card_type property
func (m *ContactsItemCreditCardsPostResponse) SetCardType(value *string) {
	m.card_type = value
}

// SetContactId sets the contact_id property value. The contact_id property
func (m *ContactsItemCreditCardsPostResponse) SetContactId(value *int64) {
	m.contact_id = value
}

// SetEmailAddress sets the email_address property value. The email_address property
func (m *ContactsItemCreditCardsPostResponse) SetEmailAddress(value *string) {
	m.email_address = value
}

// SetExpirationMonth sets the expiration_month property value. The expiration_month property
func (m *ContactsItemCreditCardsPostResponse) SetExpirationMonth(value *string) {
	m.expiration_month = value
}

// SetExpirationYear sets the expiration_year property value. The expiration_year property
func (m *ContactsItemCreditCardsPostResponse) SetExpirationYear(value *string) {
	m.expiration_year = value
}

// SetId sets the id property value. The id property
func (m *ContactsItemCreditCardsPostResponse) SetId(value *int64) {
	m.id = value
}

// SetMaestroIssueNumber sets the maestro_issue_number property value. The maestro_issue_number property
func (m *ContactsItemCreditCardsPostResponse) SetMaestroIssueNumber(value *string) {
	m.maestro_issue_number = value
}

// SetMaestroStartDateMonth sets the maestro_start_date_month property value. The maestro_start_date_month property
func (m *ContactsItemCreditCardsPostResponse) SetMaestroStartDateMonth(value *string) {
	m.maestro_start_date_month = value
}

// SetMaestroStartDateYear sets the maestro_start_date_year property value. The maestro_start_date_year property
func (m *ContactsItemCreditCardsPostResponse) SetMaestroStartDateYear(value *string) {
	m.maestro_start_date_year = value
}

// SetNameOnCard sets the name_on_card property value. The name_on_card property
func (m *ContactsItemCreditCardsPostResponse) SetNameOnCard(value *string) {
	m.name_on_card = value
}

// SetValidationMessage sets the validation_message property value. The validation_message property
func (m *ContactsItemCreditCardsPostResponse) SetValidationMessage(value *string) {
	m.validation_message = value
}

// SetValidationStatus sets the validation_status property value. The validation_status property
func (m *ContactsItemCreditCardsPostResponse) SetValidationStatus(value *string) {
	m.validation_status = value
}

type ContactsItemCreditCardsPostResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAddress() ContactsItemCreditCardsPostResponse_addressable
	GetCardType() *string
	GetContactId() *int64
	GetEmailAddress() *string
	GetExpirationMonth() *string
	GetExpirationYear() *string
	GetId() *int64
	GetMaestroIssueNumber() *string
	GetMaestroStartDateMonth() *string
	GetMaestroStartDateYear() *string
	GetNameOnCard() *string
	GetValidationMessage() *string
	GetValidationStatus() *string
	SetAddress(value ContactsItemCreditCardsPostResponse_addressable)
	SetCardType(value *string)
	SetContactId(value *int64)
	SetEmailAddress(value *string)
	SetExpirationMonth(value *string)
	SetExpirationYear(value *string)
	SetId(value *int64)
	SetMaestroIssueNumber(value *string)
	SetMaestroStartDateMonth(value *string)
	SetMaestroStartDateYear(value *string)
	SetNameOnCard(value *string)
	SetValidationMessage(value *string)
	SetValidationStatus(value *string)
}
