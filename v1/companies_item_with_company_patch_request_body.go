package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CompaniesItemWithCompanyPatchRequestBody struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The address property
	address CompaniesItemWithCompanyPatchRequestBody_addressable
	// The company_name property
	company_name *string
	// The custom_fields property
	custom_fields []CompaniesItemWithCompanyPatchRequestBody_custom_fieldsable
	// The email_address property
	email_address *string
	// The fax_number property
	fax_number CompaniesItemWithCompanyPatchRequestBody_fax_numberable
	// The notes property
	notes *string
	// The opt_in_reason property
	opt_in_reason *string
	// The phone_number property
	phone_number CompaniesItemWithCompanyPatchRequestBody_phone_numberable
	// The website property
	website *string
}

// NewCompaniesItemWithCompanyPatchRequestBody instantiates a new CompaniesItemWithCompanyPatchRequestBody and sets the default values.
func NewCompaniesItemWithCompanyPatchRequestBody() *CompaniesItemWithCompanyPatchRequestBody {
	m := &CompaniesItemWithCompanyPatchRequestBody{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateCompaniesItemWithCompanyPatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCompaniesItemWithCompanyPatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCompaniesItemWithCompanyPatchRequestBody(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CompaniesItemWithCompanyPatchRequestBody) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAddress gets the address property value. The address property
// returns a CompaniesItemWithCompanyPatchRequestBody_addressable when successful
func (m *CompaniesItemWithCompanyPatchRequestBody) GetAddress() CompaniesItemWithCompanyPatchRequestBody_addressable {
	return m.address
}

// GetCompanyName gets the company_name property value. The company_name property
// returns a *string when successful
func (m *CompaniesItemWithCompanyPatchRequestBody) GetCompanyName() *string {
	return m.company_name
}

// GetCustomFields gets the custom_fields property value. The custom_fields property
// returns a []CompaniesItemWithCompanyPatchRequestBody_custom_fieldsable when successful
func (m *CompaniesItemWithCompanyPatchRequestBody) GetCustomFields() []CompaniesItemWithCompanyPatchRequestBody_custom_fieldsable {
	return m.custom_fields
}

// GetEmailAddress gets the email_address property value. The email_address property
// returns a *string when successful
func (m *CompaniesItemWithCompanyPatchRequestBody) GetEmailAddress() *string {
	return m.email_address
}

// GetFaxNumber gets the fax_number property value. The fax_number property
// returns a CompaniesItemWithCompanyPatchRequestBody_fax_numberable when successful
func (m *CompaniesItemWithCompanyPatchRequestBody) GetFaxNumber() CompaniesItemWithCompanyPatchRequestBody_fax_numberable {
	return m.fax_number
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CompaniesItemWithCompanyPatchRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["address"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateCompaniesItemWithCompanyPatchRequestBody_addressFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAddress(val.(CompaniesItemWithCompanyPatchRequestBody_addressable))
		}
		return nil
	}
	res["company_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCompanyName(val)
		}
		return nil
	}
	res["custom_fields"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateCompaniesItemWithCompanyPatchRequestBody_custom_fieldsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]CompaniesItemWithCompanyPatchRequestBody_custom_fieldsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(CompaniesItemWithCompanyPatchRequestBody_custom_fieldsable)
				}
			}
			m.SetCustomFields(res)
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
	res["fax_number"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateCompaniesItemWithCompanyPatchRequestBody_fax_numberFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFaxNumber(val.(CompaniesItemWithCompanyPatchRequestBody_fax_numberable))
		}
		return nil
	}
	res["notes"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNotes(val)
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
	res["phone_number"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateCompaniesItemWithCompanyPatchRequestBody_phone_numberFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPhoneNumber(val.(CompaniesItemWithCompanyPatchRequestBody_phone_numberable))
		}
		return nil
	}
	res["website"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetWebsite(val)
		}
		return nil
	}
	return res
}

// GetNotes gets the notes property value. The notes property
// returns a *string when successful
func (m *CompaniesItemWithCompanyPatchRequestBody) GetNotes() *string {
	return m.notes
}

// GetOptInReason gets the opt_in_reason property value. The opt_in_reason property
// returns a *string when successful
func (m *CompaniesItemWithCompanyPatchRequestBody) GetOptInReason() *string {
	return m.opt_in_reason
}

// GetPhoneNumber gets the phone_number property value. The phone_number property
// returns a CompaniesItemWithCompanyPatchRequestBody_phone_numberable when successful
func (m *CompaniesItemWithCompanyPatchRequestBody) GetPhoneNumber() CompaniesItemWithCompanyPatchRequestBody_phone_numberable {
	return m.phone_number
}

// GetWebsite gets the website property value. The website property
// returns a *string when successful
func (m *CompaniesItemWithCompanyPatchRequestBody) GetWebsite() *string {
	return m.website
}

// Serialize serializes information the current object
func (m *CompaniesItemWithCompanyPatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("address", m.GetAddress())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("company_name", m.GetCompanyName())
		if err != nil {
			return err
		}
	}
	if m.GetCustomFields() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomFields()))
		for i, v := range m.GetCustomFields() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("custom_fields", cast)
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
		err := writer.WriteObjectValue("fax_number", m.GetFaxNumber())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("notes", m.GetNotes())
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
		err := writer.WriteObjectValue("phone_number", m.GetPhoneNumber())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("website", m.GetWebsite())
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
func (m *CompaniesItemWithCompanyPatchRequestBody) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAddress sets the address property value. The address property
func (m *CompaniesItemWithCompanyPatchRequestBody) SetAddress(value CompaniesItemWithCompanyPatchRequestBody_addressable) {
	m.address = value
}

// SetCompanyName sets the company_name property value. The company_name property
func (m *CompaniesItemWithCompanyPatchRequestBody) SetCompanyName(value *string) {
	m.company_name = value
}

// SetCustomFields sets the custom_fields property value. The custom_fields property
func (m *CompaniesItemWithCompanyPatchRequestBody) SetCustomFields(value []CompaniesItemWithCompanyPatchRequestBody_custom_fieldsable) {
	m.custom_fields = value
}

// SetEmailAddress sets the email_address property value. The email_address property
func (m *CompaniesItemWithCompanyPatchRequestBody) SetEmailAddress(value *string) {
	m.email_address = value
}

// SetFaxNumber sets the fax_number property value. The fax_number property
func (m *CompaniesItemWithCompanyPatchRequestBody) SetFaxNumber(value CompaniesItemWithCompanyPatchRequestBody_fax_numberable) {
	m.fax_number = value
}

// SetNotes sets the notes property value. The notes property
func (m *CompaniesItemWithCompanyPatchRequestBody) SetNotes(value *string) {
	m.notes = value
}

// SetOptInReason sets the opt_in_reason property value. The opt_in_reason property
func (m *CompaniesItemWithCompanyPatchRequestBody) SetOptInReason(value *string) {
	m.opt_in_reason = value
}

// SetPhoneNumber sets the phone_number property value. The phone_number property
func (m *CompaniesItemWithCompanyPatchRequestBody) SetPhoneNumber(value CompaniesItemWithCompanyPatchRequestBody_phone_numberable) {
	m.phone_number = value
}

// SetWebsite sets the website property value. The website property
func (m *CompaniesItemWithCompanyPatchRequestBody) SetWebsite(value *string) {
	m.website = value
}

type CompaniesItemWithCompanyPatchRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAddress() CompaniesItemWithCompanyPatchRequestBody_addressable
	GetCompanyName() *string
	GetCustomFields() []CompaniesItemWithCompanyPatchRequestBody_custom_fieldsable
	GetEmailAddress() *string
	GetFaxNumber() CompaniesItemWithCompanyPatchRequestBody_fax_numberable
	GetNotes() *string
	GetOptInReason() *string
	GetPhoneNumber() CompaniesItemWithCompanyPatchRequestBody_phone_numberable
	GetWebsite() *string
	SetAddress(value CompaniesItemWithCompanyPatchRequestBody_addressable)
	SetCompanyName(value *string)
	SetCustomFields(value []CompaniesItemWithCompanyPatchRequestBody_custom_fieldsable)
	SetEmailAddress(value *string)
	SetFaxNumber(value CompaniesItemWithCompanyPatchRequestBody_fax_numberable)
	SetNotes(value *string)
	SetOptInReason(value *string)
	SetPhoneNumber(value CompaniesItemWithCompanyPatchRequestBody_phone_numberable)
	SetWebsite(value *string)
}
