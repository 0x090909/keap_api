package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TagsItemCompaniesGetResponse_companies_company struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The company_name property
	company_name *string
	// The email property
	email *string
	// The id property
	id *int64
}

// NewTagsItemCompaniesGetResponse_companies_company instantiates a new TagsItemCompaniesGetResponse_companies_company and sets the default values.
func NewTagsItemCompaniesGetResponse_companies_company() *TagsItemCompaniesGetResponse_companies_company {
	m := &TagsItemCompaniesGetResponse_companies_company{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateTagsItemCompaniesGetResponse_companies_companyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTagsItemCompaniesGetResponse_companies_companyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTagsItemCompaniesGetResponse_companies_company(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TagsItemCompaniesGetResponse_companies_company) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCompanyName gets the company_name property value. The company_name property
// returns a *string when successful
func (m *TagsItemCompaniesGetResponse_companies_company) GetCompanyName() *string {
	return m.company_name
}

// GetEmail gets the email property value. The email property
// returns a *string when successful
func (m *TagsItemCompaniesGetResponse_companies_company) GetEmail() *string {
	return m.email
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TagsItemCompaniesGetResponse_companies_company) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
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
	res["email"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetEmail(val)
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
	return res
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *TagsItemCompaniesGetResponse_companies_company) GetId() *int64 {
	return m.id
}

// Serialize serializes information the current object
func (m *TagsItemCompaniesGetResponse_companies_company) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("company_name", m.GetCompanyName())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("email", m.GetEmail())
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
		err := writer.WriteAdditionalData(m.GetAdditionalData())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TagsItemCompaniesGetResponse_companies_company) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCompanyName sets the company_name property value. The company_name property
func (m *TagsItemCompaniesGetResponse_companies_company) SetCompanyName(value *string) {
	m.company_name = value
}

// SetEmail sets the email property value. The email property
func (m *TagsItemCompaniesGetResponse_companies_company) SetEmail(value *string) {
	m.email = value
}

// SetId sets the id property value. The id property
func (m *TagsItemCompaniesGetResponse_companies_company) SetId(value *int64) {
	m.id = value
}

type TagsItemCompaniesGetResponse_companies_companyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCompanyName() *string
	GetEmail() *string
	GetId() *int64
	SetCompanyName(value *string)
	SetEmail(value *string)
	SetId(value *int64)
}
