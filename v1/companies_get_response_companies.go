package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CompaniesGetResponse_companies struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The address property
	address CompaniesGetResponse_companies_addressable
	// The company_name property
	company_name *string
	// The custom_fields property
	custom_fields []CompaniesGetResponse_companies_custom_fieldsable
	// The email_address property
	email_address *string
	// The email_opted_in property
	email_opted_in *bool
	// The fax_number property
	fax_number CompaniesGetResponse_companies_fax_numberable
	// The id property
	id *int64
	// The notes property
	notes *string
	// The phone_number property
	phone_number CompaniesGetResponse_companies_phone_numberable
	// The website property
	website *string
}

// NewCompaniesGetResponse_companies instantiates a new CompaniesGetResponse_companies and sets the default values.
func NewCompaniesGetResponse_companies() *CompaniesGetResponse_companies {
	m := &CompaniesGetResponse_companies{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateCompaniesGetResponse_companiesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCompaniesGetResponse_companiesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCompaniesGetResponse_companies(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CompaniesGetResponse_companies) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAddress gets the address property value. The address property
// returns a CompaniesGetResponse_companies_addressable when successful
func (m *CompaniesGetResponse_companies) GetAddress() CompaniesGetResponse_companies_addressable {
	return m.address
}

// GetCompanyName gets the company_name property value. The company_name property
// returns a *string when successful
func (m *CompaniesGetResponse_companies) GetCompanyName() *string {
	return m.company_name
}

// GetCustomFields gets the custom_fields property value. The custom_fields property
// returns a []CompaniesGetResponse_companies_custom_fieldsable when successful
func (m *CompaniesGetResponse_companies) GetCustomFields() []CompaniesGetResponse_companies_custom_fieldsable {
	return m.custom_fields
}

// GetEmailAddress gets the email_address property value. The email_address property
// returns a *string when successful
func (m *CompaniesGetResponse_companies) GetEmailAddress() *string {
	return m.email_address
}

// GetEmailOptedIn gets the email_opted_in property value. The email_opted_in property
// returns a *bool when successful
func (m *CompaniesGetResponse_companies) GetEmailOptedIn() *bool {
	return m.email_opted_in
}

// GetFaxNumber gets the fax_number property value. The fax_number property
// returns a CompaniesGetResponse_companies_fax_numberable when successful
func (m *CompaniesGetResponse_companies) GetFaxNumber() CompaniesGetResponse_companies_fax_numberable {
	return m.fax_number
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CompaniesGetResponse_companies) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["address"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateCompaniesGetResponse_companies_addressFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAddress(val.(CompaniesGetResponse_companies_addressable))
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
		val, err := n.GetCollectionOfObjectValues(CreateCompaniesGetResponse_companies_custom_fieldsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]CompaniesGetResponse_companies_custom_fieldsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(CompaniesGetResponse_companies_custom_fieldsable)
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
	res["email_opted_in"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetEmailOptedIn(val)
		}
		return nil
	}
	res["fax_number"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateCompaniesGetResponse_companies_fax_numberFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFaxNumber(val.(CompaniesGetResponse_companies_fax_numberable))
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
		val, err := n.GetObjectValue(CreateCompaniesGetResponse_companies_phone_numberFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPhoneNumber(val.(CompaniesGetResponse_companies_phone_numberable))
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

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *CompaniesGetResponse_companies) GetId() *int64 {
	return m.id
}

// GetNotes gets the notes property value. The notes property
// returns a *string when successful
func (m *CompaniesGetResponse_companies) GetNotes() *string {
	return m.notes
}

// GetPhoneNumber gets the phone_number property value. The phone_number property
// returns a CompaniesGetResponse_companies_phone_numberable when successful
func (m *CompaniesGetResponse_companies) GetPhoneNumber() CompaniesGetResponse_companies_phone_numberable {
	return m.phone_number
}

// GetWebsite gets the website property value. The website property
// returns a *string when successful
func (m *CompaniesGetResponse_companies) GetWebsite() *string {
	return m.website
}

// Serialize serializes information the current object
func (m *CompaniesGetResponse_companies) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
		err := writer.WriteBoolValue("email_opted_in", m.GetEmailOptedIn())
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
		err := writer.WriteInt64Value("id", m.GetId())
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
func (m *CompaniesGetResponse_companies) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAddress sets the address property value. The address property
func (m *CompaniesGetResponse_companies) SetAddress(value CompaniesGetResponse_companies_addressable) {
	m.address = value
}

// SetCompanyName sets the company_name property value. The company_name property
func (m *CompaniesGetResponse_companies) SetCompanyName(value *string) {
	m.company_name = value
}

// SetCustomFields sets the custom_fields property value. The custom_fields property
func (m *CompaniesGetResponse_companies) SetCustomFields(value []CompaniesGetResponse_companies_custom_fieldsable) {
	m.custom_fields = value
}

// SetEmailAddress sets the email_address property value. The email_address property
func (m *CompaniesGetResponse_companies) SetEmailAddress(value *string) {
	m.email_address = value
}

// SetEmailOptedIn sets the email_opted_in property value. The email_opted_in property
func (m *CompaniesGetResponse_companies) SetEmailOptedIn(value *bool) {
	m.email_opted_in = value
}

// SetFaxNumber sets the fax_number property value. The fax_number property
func (m *CompaniesGetResponse_companies) SetFaxNumber(value CompaniesGetResponse_companies_fax_numberable) {
	m.fax_number = value
}

// SetId sets the id property value. The id property
func (m *CompaniesGetResponse_companies) SetId(value *int64) {
	m.id = value
}

// SetNotes sets the notes property value. The notes property
func (m *CompaniesGetResponse_companies) SetNotes(value *string) {
	m.notes = value
}

// SetPhoneNumber sets the phone_number property value. The phone_number property
func (m *CompaniesGetResponse_companies) SetPhoneNumber(value CompaniesGetResponse_companies_phone_numberable) {
	m.phone_number = value
}

// SetWebsite sets the website property value. The website property
func (m *CompaniesGetResponse_companies) SetWebsite(value *string) {
	m.website = value
}

type CompaniesGetResponse_companiesable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAddress() CompaniesGetResponse_companies_addressable
	GetCompanyName() *string
	GetCustomFields() []CompaniesGetResponse_companies_custom_fieldsable
	GetEmailAddress() *string
	GetEmailOptedIn() *bool
	GetFaxNumber() CompaniesGetResponse_companies_fax_numberable
	GetId() *int64
	GetNotes() *string
	GetPhoneNumber() CompaniesGetResponse_companies_phone_numberable
	GetWebsite() *string
	SetAddress(value CompaniesGetResponse_companies_addressable)
	SetCompanyName(value *string)
	SetCustomFields(value []CompaniesGetResponse_companies_custom_fieldsable)
	SetEmailAddress(value *string)
	SetEmailOptedIn(value *bool)
	SetFaxNumber(value CompaniesGetResponse_companies_fax_numberable)
	SetId(value *int64)
	SetNotes(value *string)
	SetPhoneNumber(value CompaniesGetResponse_companies_phone_numberable)
	SetWebsite(value *string)
}
