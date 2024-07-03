package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CompaniesPostResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The address property
	address CompaniesPostResponse_addressable
	// The company_name property
	company_name *string
	// The create_time property
	create_time *string
	// The custom_fields property
	custom_fields []CompaniesPostResponse_custom_fieldsable
	// The email_address property
	email_address CompaniesPostResponse_email_addressable
	// The fax_number property
	fax_number CompaniesPostResponse_fax_numberable
	// The id property
	id *string
	// The notes property
	notes *string
	// The phone_number property
	phone_number CompaniesPostResponse_phone_numberable
	// The update_time property
	update_time *string
	// The website property
	website *string
}

// NewCompaniesPostResponse instantiates a new CompaniesPostResponse and sets the default values.
func NewCompaniesPostResponse() *CompaniesPostResponse {
	m := &CompaniesPostResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateCompaniesPostResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCompaniesPostResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCompaniesPostResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CompaniesPostResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAddress gets the address property value. The address property
// returns a CompaniesPostResponse_addressable when successful
func (m *CompaniesPostResponse) GetAddress() CompaniesPostResponse_addressable {
	return m.address
}

// GetCompanyName gets the company_name property value. The company_name property
// returns a *string when successful
func (m *CompaniesPostResponse) GetCompanyName() *string {
	return m.company_name
}

// GetCreateTime gets the create_time property value. The create_time property
// returns a *string when successful
func (m *CompaniesPostResponse) GetCreateTime() *string {
	return m.create_time
}

// GetCustomFields gets the custom_fields property value. The custom_fields property
// returns a []CompaniesPostResponse_custom_fieldsable when successful
func (m *CompaniesPostResponse) GetCustomFields() []CompaniesPostResponse_custom_fieldsable {
	return m.custom_fields
}

// GetEmailAddress gets the email_address property value. The email_address property
// returns a CompaniesPostResponse_email_addressable when successful
func (m *CompaniesPostResponse) GetEmailAddress() CompaniesPostResponse_email_addressable {
	return m.email_address
}

// GetFaxNumber gets the fax_number property value. The fax_number property
// returns a CompaniesPostResponse_fax_numberable when successful
func (m *CompaniesPostResponse) GetFaxNumber() CompaniesPostResponse_fax_numberable {
	return m.fax_number
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CompaniesPostResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["address"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateCompaniesPostResponse_addressFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAddress(val.(CompaniesPostResponse_addressable))
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
	res["create_time"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCreateTime(val)
		}
		return nil
	}
	res["custom_fields"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateCompaniesPostResponse_custom_fieldsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]CompaniesPostResponse_custom_fieldsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(CompaniesPostResponse_custom_fieldsable)
				}
			}
			m.SetCustomFields(res)
		}
		return nil
	}
	res["email_address"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateCompaniesPostResponse_email_addressFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetEmailAddress(val.(CompaniesPostResponse_email_addressable))
		}
		return nil
	}
	res["fax_number"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateCompaniesPostResponse_fax_numberFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFaxNumber(val.(CompaniesPostResponse_fax_numberable))
		}
		return nil
	}
	res["id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
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
		val, err := n.GetObjectValue(CreateCompaniesPostResponse_phone_numberFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPhoneNumber(val.(CompaniesPostResponse_phone_numberable))
		}
		return nil
	}
	res["update_time"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetUpdateTime(val)
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
// returns a *string when successful
func (m *CompaniesPostResponse) GetId() *string {
	return m.id
}

// GetNotes gets the notes property value. The notes property
// returns a *string when successful
func (m *CompaniesPostResponse) GetNotes() *string {
	return m.notes
}

// GetPhoneNumber gets the phone_number property value. The phone_number property
// returns a CompaniesPostResponse_phone_numberable when successful
func (m *CompaniesPostResponse) GetPhoneNumber() CompaniesPostResponse_phone_numberable {
	return m.phone_number
}

// GetUpdateTime gets the update_time property value. The update_time property
// returns a *string when successful
func (m *CompaniesPostResponse) GetUpdateTime() *string {
	return m.update_time
}

// GetWebsite gets the website property value. The website property
// returns a *string when successful
func (m *CompaniesPostResponse) GetWebsite() *string {
	return m.website
}

// Serialize serializes information the current object
func (m *CompaniesPostResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
	{
		err := writer.WriteStringValue("create_time", m.GetCreateTime())
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
		err := writer.WriteStringValue("id", m.GetId())
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
		err := writer.WriteStringValue("update_time", m.GetUpdateTime())
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
func (m *CompaniesPostResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAddress sets the address property value. The address property
func (m *CompaniesPostResponse) SetAddress(value CompaniesPostResponse_addressable) {
	m.address = value
}

// SetCompanyName sets the company_name property value. The company_name property
func (m *CompaniesPostResponse) SetCompanyName(value *string) {
	m.company_name = value
}

// SetCreateTime sets the create_time property value. The create_time property
func (m *CompaniesPostResponse) SetCreateTime(value *string) {
	m.create_time = value
}

// SetCustomFields sets the custom_fields property value. The custom_fields property
func (m *CompaniesPostResponse) SetCustomFields(value []CompaniesPostResponse_custom_fieldsable) {
	m.custom_fields = value
}

// SetEmailAddress sets the email_address property value. The email_address property
func (m *CompaniesPostResponse) SetEmailAddress(value CompaniesPostResponse_email_addressable) {
	m.email_address = value
}

// SetFaxNumber sets the fax_number property value. The fax_number property
func (m *CompaniesPostResponse) SetFaxNumber(value CompaniesPostResponse_fax_numberable) {
	m.fax_number = value
}

// SetId sets the id property value. The id property
func (m *CompaniesPostResponse) SetId(value *string) {
	m.id = value
}

// SetNotes sets the notes property value. The notes property
func (m *CompaniesPostResponse) SetNotes(value *string) {
	m.notes = value
}

// SetPhoneNumber sets the phone_number property value. The phone_number property
func (m *CompaniesPostResponse) SetPhoneNumber(value CompaniesPostResponse_phone_numberable) {
	m.phone_number = value
}

// SetUpdateTime sets the update_time property value. The update_time property
func (m *CompaniesPostResponse) SetUpdateTime(value *string) {
	m.update_time = value
}

// SetWebsite sets the website property value. The website property
func (m *CompaniesPostResponse) SetWebsite(value *string) {
	m.website = value
}

type CompaniesPostResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAddress() CompaniesPostResponse_addressable
	GetCompanyName() *string
	GetCreateTime() *string
	GetCustomFields() []CompaniesPostResponse_custom_fieldsable
	GetEmailAddress() CompaniesPostResponse_email_addressable
	GetFaxNumber() CompaniesPostResponse_fax_numberable
	GetId() *string
	GetNotes() *string
	GetPhoneNumber() CompaniesPostResponse_phone_numberable
	GetUpdateTime() *string
	GetWebsite() *string
	SetAddress(value CompaniesPostResponse_addressable)
	SetCompanyName(value *string)
	SetCreateTime(value *string)
	SetCustomFields(value []CompaniesPostResponse_custom_fieldsable)
	SetEmailAddress(value CompaniesPostResponse_email_addressable)
	SetFaxNumber(value CompaniesPostResponse_fax_numberable)
	SetId(value *string)
	SetNotes(value *string)
	SetPhoneNumber(value CompaniesPostResponse_phone_numberable)
	SetUpdateTime(value *string)
	SetWebsite(value *string)
}
