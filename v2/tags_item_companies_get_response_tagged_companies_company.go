package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TagsItemCompaniesGetResponse_tagged_companies_company struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The address property
	address TagsItemCompaniesGetResponse_tagged_companies_company_addressable
	// The company_name property
	company_name *string
	// The create_time property
	create_time *string
	// The custom_fields property
	custom_fields []TagsItemCompaniesGetResponse_tagged_companies_company_custom_fieldsable
	// The email_address property
	email_address TagsItemCompaniesGetResponse_tagged_companies_company_email_addressable
	// The fax_number property
	fax_number TagsItemCompaniesGetResponse_tagged_companies_company_fax_numberable
	// The id property
	id *string
	// The notes property
	notes *string
	// The phone_number property
	phone_number TagsItemCompaniesGetResponse_tagged_companies_company_phone_numberable
	// The update_time property
	update_time *string
	// The website property
	website *string
}

// NewTagsItemCompaniesGetResponse_tagged_companies_company instantiates a new TagsItemCompaniesGetResponse_tagged_companies_company and sets the default values.
func NewTagsItemCompaniesGetResponse_tagged_companies_company() *TagsItemCompaniesGetResponse_tagged_companies_company {
	m := &TagsItemCompaniesGetResponse_tagged_companies_company{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateTagsItemCompaniesGetResponse_tagged_companies_companyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTagsItemCompaniesGetResponse_tagged_companies_companyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTagsItemCompaniesGetResponse_tagged_companies_company(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TagsItemCompaniesGetResponse_tagged_companies_company) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAddress gets the address property value. The address property
// returns a TagsItemCompaniesGetResponse_tagged_companies_company_addressable when successful
func (m *TagsItemCompaniesGetResponse_tagged_companies_company) GetAddress() TagsItemCompaniesGetResponse_tagged_companies_company_addressable {
	return m.address
}

// GetCompanyName gets the company_name property value. The company_name property
// returns a *string when successful
func (m *TagsItemCompaniesGetResponse_tagged_companies_company) GetCompanyName() *string {
	return m.company_name
}

// GetCreateTime gets the create_time property value. The create_time property
// returns a *string when successful
func (m *TagsItemCompaniesGetResponse_tagged_companies_company) GetCreateTime() *string {
	return m.create_time
}

// GetCustomFields gets the custom_fields property value. The custom_fields property
// returns a []TagsItemCompaniesGetResponse_tagged_companies_company_custom_fieldsable when successful
func (m *TagsItemCompaniesGetResponse_tagged_companies_company) GetCustomFields() []TagsItemCompaniesGetResponse_tagged_companies_company_custom_fieldsable {
	return m.custom_fields
}

// GetEmailAddress gets the email_address property value. The email_address property
// returns a TagsItemCompaniesGetResponse_tagged_companies_company_email_addressable when successful
func (m *TagsItemCompaniesGetResponse_tagged_companies_company) GetEmailAddress() TagsItemCompaniesGetResponse_tagged_companies_company_email_addressable {
	return m.email_address
}

// GetFaxNumber gets the fax_number property value. The fax_number property
// returns a TagsItemCompaniesGetResponse_tagged_companies_company_fax_numberable when successful
func (m *TagsItemCompaniesGetResponse_tagged_companies_company) GetFaxNumber() TagsItemCompaniesGetResponse_tagged_companies_company_fax_numberable {
	return m.fax_number
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TagsItemCompaniesGetResponse_tagged_companies_company) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["address"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateTagsItemCompaniesGetResponse_tagged_companies_company_addressFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAddress(val.(TagsItemCompaniesGetResponse_tagged_companies_company_addressable))
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
		val, err := n.GetCollectionOfObjectValues(CreateTagsItemCompaniesGetResponse_tagged_companies_company_custom_fieldsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]TagsItemCompaniesGetResponse_tagged_companies_company_custom_fieldsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(TagsItemCompaniesGetResponse_tagged_companies_company_custom_fieldsable)
				}
			}
			m.SetCustomFields(res)
		}
		return nil
	}
	res["email_address"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateTagsItemCompaniesGetResponse_tagged_companies_company_email_addressFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetEmailAddress(val.(TagsItemCompaniesGetResponse_tagged_companies_company_email_addressable))
		}
		return nil
	}
	res["fax_number"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateTagsItemCompaniesGetResponse_tagged_companies_company_fax_numberFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFaxNumber(val.(TagsItemCompaniesGetResponse_tagged_companies_company_fax_numberable))
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
		val, err := n.GetObjectValue(CreateTagsItemCompaniesGetResponse_tagged_companies_company_phone_numberFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPhoneNumber(val.(TagsItemCompaniesGetResponse_tagged_companies_company_phone_numberable))
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
func (m *TagsItemCompaniesGetResponse_tagged_companies_company) GetId() *string {
	return m.id
}

// GetNotes gets the notes property value. The notes property
// returns a *string when successful
func (m *TagsItemCompaniesGetResponse_tagged_companies_company) GetNotes() *string {
	return m.notes
}

// GetPhoneNumber gets the phone_number property value. The phone_number property
// returns a TagsItemCompaniesGetResponse_tagged_companies_company_phone_numberable when successful
func (m *TagsItemCompaniesGetResponse_tagged_companies_company) GetPhoneNumber() TagsItemCompaniesGetResponse_tagged_companies_company_phone_numberable {
	return m.phone_number
}

// GetUpdateTime gets the update_time property value. The update_time property
// returns a *string when successful
func (m *TagsItemCompaniesGetResponse_tagged_companies_company) GetUpdateTime() *string {
	return m.update_time
}

// GetWebsite gets the website property value. The website property
// returns a *string when successful
func (m *TagsItemCompaniesGetResponse_tagged_companies_company) GetWebsite() *string {
	return m.website
}

// Serialize serializes information the current object
func (m *TagsItemCompaniesGetResponse_tagged_companies_company) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *TagsItemCompaniesGetResponse_tagged_companies_company) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAddress sets the address property value. The address property
func (m *TagsItemCompaniesGetResponse_tagged_companies_company) SetAddress(value TagsItemCompaniesGetResponse_tagged_companies_company_addressable) {
	m.address = value
}

// SetCompanyName sets the company_name property value. The company_name property
func (m *TagsItemCompaniesGetResponse_tagged_companies_company) SetCompanyName(value *string) {
	m.company_name = value
}

// SetCreateTime sets the create_time property value. The create_time property
func (m *TagsItemCompaniesGetResponse_tagged_companies_company) SetCreateTime(value *string) {
	m.create_time = value
}

// SetCustomFields sets the custom_fields property value. The custom_fields property
func (m *TagsItemCompaniesGetResponse_tagged_companies_company) SetCustomFields(value []TagsItemCompaniesGetResponse_tagged_companies_company_custom_fieldsable) {
	m.custom_fields = value
}

// SetEmailAddress sets the email_address property value. The email_address property
func (m *TagsItemCompaniesGetResponse_tagged_companies_company) SetEmailAddress(value TagsItemCompaniesGetResponse_tagged_companies_company_email_addressable) {
	m.email_address = value
}

// SetFaxNumber sets the fax_number property value. The fax_number property
func (m *TagsItemCompaniesGetResponse_tagged_companies_company) SetFaxNumber(value TagsItemCompaniesGetResponse_tagged_companies_company_fax_numberable) {
	m.fax_number = value
}

// SetId sets the id property value. The id property
func (m *TagsItemCompaniesGetResponse_tagged_companies_company) SetId(value *string) {
	m.id = value
}

// SetNotes sets the notes property value. The notes property
func (m *TagsItemCompaniesGetResponse_tagged_companies_company) SetNotes(value *string) {
	m.notes = value
}

// SetPhoneNumber sets the phone_number property value. The phone_number property
func (m *TagsItemCompaniesGetResponse_tagged_companies_company) SetPhoneNumber(value TagsItemCompaniesGetResponse_tagged_companies_company_phone_numberable) {
	m.phone_number = value
}

// SetUpdateTime sets the update_time property value. The update_time property
func (m *TagsItemCompaniesGetResponse_tagged_companies_company) SetUpdateTime(value *string) {
	m.update_time = value
}

// SetWebsite sets the website property value. The website property
func (m *TagsItemCompaniesGetResponse_tagged_companies_company) SetWebsite(value *string) {
	m.website = value
}

type TagsItemCompaniesGetResponse_tagged_companies_companyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAddress() TagsItemCompaniesGetResponse_tagged_companies_company_addressable
	GetCompanyName() *string
	GetCreateTime() *string
	GetCustomFields() []TagsItemCompaniesGetResponse_tagged_companies_company_custom_fieldsable
	GetEmailAddress() TagsItemCompaniesGetResponse_tagged_companies_company_email_addressable
	GetFaxNumber() TagsItemCompaniesGetResponse_tagged_companies_company_fax_numberable
	GetId() *string
	GetNotes() *string
	GetPhoneNumber() TagsItemCompaniesGetResponse_tagged_companies_company_phone_numberable
	GetUpdateTime() *string
	GetWebsite() *string
	SetAddress(value TagsItemCompaniesGetResponse_tagged_companies_company_addressable)
	SetCompanyName(value *string)
	SetCreateTime(value *string)
	SetCustomFields(value []TagsItemCompaniesGetResponse_tagged_companies_company_custom_fieldsable)
	SetEmailAddress(value TagsItemCompaniesGetResponse_tagged_companies_company_email_addressable)
	SetFaxNumber(value TagsItemCompaniesGetResponse_tagged_companies_company_fax_numberable)
	SetId(value *string)
	SetNotes(value *string)
	SetPhoneNumber(value TagsItemCompaniesGetResponse_tagged_companies_company_phone_numberable)
	SetUpdateTime(value *string)
	SetWebsite(value *string)
}
