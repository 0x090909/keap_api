package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CompaniesItemWithCompany_PatchRequestBody struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The address property
	address CompaniesItemWithCompany_PatchRequestBody_addressable
	// The company_name property
	company_name *string
	// The custom_fields property
	custom_fields []CompaniesItemWithCompany_PatchRequestBody_custom_fieldsable
	// The email_address property
	email_address CompaniesItemWithCompany_PatchRequestBody_email_addressable
	// The fax_number property
	fax_number CompaniesItemWithCompany_PatchRequestBody_fax_numberable
	// The notes property
	notes *string
	// The phone_number property
	phone_number CompaniesItemWithCompany_PatchRequestBody_phone_numberable
	// The website property
	website *string
}

// NewCompaniesItemWithCompany_PatchRequestBody instantiates a new CompaniesItemWithCompany_PatchRequestBody and sets the default values.
func NewCompaniesItemWithCompany_PatchRequestBody() *CompaniesItemWithCompany_PatchRequestBody {
	m := &CompaniesItemWithCompany_PatchRequestBody{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateCompaniesItemWithCompany_PatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCompaniesItemWithCompany_PatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCompaniesItemWithCompany_PatchRequestBody(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CompaniesItemWithCompany_PatchRequestBody) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAddress gets the address property value. The address property
// returns a CompaniesItemWithCompany_PatchRequestBody_addressable when successful
func (m *CompaniesItemWithCompany_PatchRequestBody) GetAddress() CompaniesItemWithCompany_PatchRequestBody_addressable {
	return m.address
}

// GetCompanyName gets the company_name property value. The company_name property
// returns a *string when successful
func (m *CompaniesItemWithCompany_PatchRequestBody) GetCompanyName() *string {
	return m.company_name
}

// GetCustomFields gets the custom_fields property value. The custom_fields property
// returns a []CompaniesItemWithCompany_PatchRequestBody_custom_fieldsable when successful
func (m *CompaniesItemWithCompany_PatchRequestBody) GetCustomFields() []CompaniesItemWithCompany_PatchRequestBody_custom_fieldsable {
	return m.custom_fields
}

// GetEmailAddress gets the email_address property value. The email_address property
// returns a CompaniesItemWithCompany_PatchRequestBody_email_addressable when successful
func (m *CompaniesItemWithCompany_PatchRequestBody) GetEmailAddress() CompaniesItemWithCompany_PatchRequestBody_email_addressable {
	return m.email_address
}

// GetFaxNumber gets the fax_number property value. The fax_number property
// returns a CompaniesItemWithCompany_PatchRequestBody_fax_numberable when successful
func (m *CompaniesItemWithCompany_PatchRequestBody) GetFaxNumber() CompaniesItemWithCompany_PatchRequestBody_fax_numberable {
	return m.fax_number
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CompaniesItemWithCompany_PatchRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["address"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateCompaniesItemWithCompany_PatchRequestBody_addressFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAddress(val.(CompaniesItemWithCompany_PatchRequestBody_addressable))
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
		val, err := n.GetCollectionOfObjectValues(CreateCompaniesItemWithCompany_PatchRequestBody_custom_fieldsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]CompaniesItemWithCompany_PatchRequestBody_custom_fieldsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(CompaniesItemWithCompany_PatchRequestBody_custom_fieldsable)
				}
			}
			m.SetCustomFields(res)
		}
		return nil
	}
	res["email_address"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateCompaniesItemWithCompany_PatchRequestBody_email_addressFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetEmailAddress(val.(CompaniesItemWithCompany_PatchRequestBody_email_addressable))
		}
		return nil
	}
	res["fax_number"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateCompaniesItemWithCompany_PatchRequestBody_fax_numberFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFaxNumber(val.(CompaniesItemWithCompany_PatchRequestBody_fax_numberable))
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
	res["phone_number"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateCompaniesItemWithCompany_PatchRequestBody_phone_numberFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPhoneNumber(val.(CompaniesItemWithCompany_PatchRequestBody_phone_numberable))
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
func (m *CompaniesItemWithCompany_PatchRequestBody) GetNotes() *string {
	return m.notes
}

// GetPhoneNumber gets the phone_number property value. The phone_number property
// returns a CompaniesItemWithCompany_PatchRequestBody_phone_numberable when successful
func (m *CompaniesItemWithCompany_PatchRequestBody) GetPhoneNumber() CompaniesItemWithCompany_PatchRequestBody_phone_numberable {
	return m.phone_number
}

// GetWebsite gets the website property value. The website property
// returns a *string when successful
func (m *CompaniesItemWithCompany_PatchRequestBody) GetWebsite() *string {
	return m.website
}

// Serialize serializes information the current object
func (m *CompaniesItemWithCompany_PatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
		err := writer.WriteObjectValue("email_address", m.GetEmailAddress())
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
func (m *CompaniesItemWithCompany_PatchRequestBody) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAddress sets the address property value. The address property
func (m *CompaniesItemWithCompany_PatchRequestBody) SetAddress(value CompaniesItemWithCompany_PatchRequestBody_addressable) {
	m.address = value
}

// SetCompanyName sets the company_name property value. The company_name property
func (m *CompaniesItemWithCompany_PatchRequestBody) SetCompanyName(value *string) {
	m.company_name = value
}

// SetCustomFields sets the custom_fields property value. The custom_fields property
func (m *CompaniesItemWithCompany_PatchRequestBody) SetCustomFields(value []CompaniesItemWithCompany_PatchRequestBody_custom_fieldsable) {
	m.custom_fields = value
}

// SetEmailAddress sets the email_address property value. The email_address property
func (m *CompaniesItemWithCompany_PatchRequestBody) SetEmailAddress(value CompaniesItemWithCompany_PatchRequestBody_email_addressable) {
	m.email_address = value
}

// SetFaxNumber sets the fax_number property value. The fax_number property
func (m *CompaniesItemWithCompany_PatchRequestBody) SetFaxNumber(value CompaniesItemWithCompany_PatchRequestBody_fax_numberable) {
	m.fax_number = value
}

// SetNotes sets the notes property value. The notes property
func (m *CompaniesItemWithCompany_PatchRequestBody) SetNotes(value *string) {
	m.notes = value
}

// SetPhoneNumber sets the phone_number property value. The phone_number property
func (m *CompaniesItemWithCompany_PatchRequestBody) SetPhoneNumber(value CompaniesItemWithCompany_PatchRequestBody_phone_numberable) {
	m.phone_number = value
}

// SetWebsite sets the website property value. The website property
func (m *CompaniesItemWithCompany_PatchRequestBody) SetWebsite(value *string) {
	m.website = value
}

type CompaniesItemWithCompany_PatchRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAddress() CompaniesItemWithCompany_PatchRequestBody_addressable
	GetCompanyName() *string
	GetCustomFields() []CompaniesItemWithCompany_PatchRequestBody_custom_fieldsable
	GetEmailAddress() CompaniesItemWithCompany_PatchRequestBody_email_addressable
	GetFaxNumber() CompaniesItemWithCompany_PatchRequestBody_fax_numberable
	GetNotes() *string
	GetPhoneNumber() CompaniesItemWithCompany_PatchRequestBody_phone_numberable
	GetWebsite() *string
	SetAddress(value CompaniesItemWithCompany_PatchRequestBody_addressable)
	SetCompanyName(value *string)
	SetCustomFields(value []CompaniesItemWithCompany_PatchRequestBody_custom_fieldsable)
	SetEmailAddress(value CompaniesItemWithCompany_PatchRequestBody_email_addressable)
	SetFaxNumber(value CompaniesItemWithCompany_PatchRequestBody_fax_numberable)
	SetNotes(value *string)
	SetPhoneNumber(value CompaniesItemWithCompany_PatchRequestBody_phone_numberable)
	SetWebsite(value *string)
}
