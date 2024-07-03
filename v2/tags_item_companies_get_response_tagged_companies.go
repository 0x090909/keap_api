package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TagsItemCompaniesGetResponse_tagged_companies struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The applied_time property
	applied_time *string
	// The company property
	company TagsItemCompaniesGetResponse_tagged_companies_companyable
}

// NewTagsItemCompaniesGetResponse_tagged_companies instantiates a new TagsItemCompaniesGetResponse_tagged_companies and sets the default values.
func NewTagsItemCompaniesGetResponse_tagged_companies() *TagsItemCompaniesGetResponse_tagged_companies {
	m := &TagsItemCompaniesGetResponse_tagged_companies{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateTagsItemCompaniesGetResponse_tagged_companiesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTagsItemCompaniesGetResponse_tagged_companiesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTagsItemCompaniesGetResponse_tagged_companies(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TagsItemCompaniesGetResponse_tagged_companies) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAppliedTime gets the applied_time property value. The applied_time property
// returns a *string when successful
func (m *TagsItemCompaniesGetResponse_tagged_companies) GetAppliedTime() *string {
	return m.applied_time
}

// GetCompany gets the company property value. The company property
// returns a TagsItemCompaniesGetResponse_tagged_companies_companyable when successful
func (m *TagsItemCompaniesGetResponse_tagged_companies) GetCompany() TagsItemCompaniesGetResponse_tagged_companies_companyable {
	return m.company
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TagsItemCompaniesGetResponse_tagged_companies) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["applied_time"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAppliedTime(val)
		}
		return nil
	}
	res["company"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateTagsItemCompaniesGetResponse_tagged_companies_companyFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCompany(val.(TagsItemCompaniesGetResponse_tagged_companies_companyable))
		}
		return nil
	}
	return res
}

// Serialize serializes information the current object
func (m *TagsItemCompaniesGetResponse_tagged_companies) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("applied_time", m.GetAppliedTime())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("company", m.GetCompany())
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
func (m *TagsItemCompaniesGetResponse_tagged_companies) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAppliedTime sets the applied_time property value. The applied_time property
func (m *TagsItemCompaniesGetResponse_tagged_companies) SetAppliedTime(value *string) {
	m.applied_time = value
}

// SetCompany sets the company property value. The company property
func (m *TagsItemCompaniesGetResponse_tagged_companies) SetCompany(value TagsItemCompaniesGetResponse_tagged_companies_companyable) {
	m.company = value
}

type TagsItemCompaniesGetResponse_tagged_companiesable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAppliedTime() *string
	GetCompany() TagsItemCompaniesGetResponse_tagged_companies_companyable
	SetAppliedTime(value *string)
	SetCompany(value TagsItemCompaniesGetResponse_tagged_companies_companyable)
}
