package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
)

type TagsItemCompaniesGetResponse_companies struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The company property
	company TagsItemCompaniesGetResponse_companies_companyable
	// The date_applied property
	date_applied *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}

// NewTagsItemCompaniesGetResponse_companies instantiates a new TagsItemCompaniesGetResponse_companies and sets the default values.
func NewTagsItemCompaniesGetResponse_companies() *TagsItemCompaniesGetResponse_companies {
	m := &TagsItemCompaniesGetResponse_companies{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateTagsItemCompaniesGetResponse_companiesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTagsItemCompaniesGetResponse_companiesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTagsItemCompaniesGetResponse_companies(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TagsItemCompaniesGetResponse_companies) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCompany gets the company property value. The company property
// returns a TagsItemCompaniesGetResponse_companies_companyable when successful
func (m *TagsItemCompaniesGetResponse_companies) GetCompany() TagsItemCompaniesGetResponse_companies_companyable {
	return m.company
}

// GetDateApplied gets the date_applied property value. The date_applied property
// returns a *Time when successful
func (m *TagsItemCompaniesGetResponse_companies) GetDateApplied() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.date_applied
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TagsItemCompaniesGetResponse_companies) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["company"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateTagsItemCompaniesGetResponse_companies_companyFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCompany(val.(TagsItemCompaniesGetResponse_companies_companyable))
		}
		return nil
	}
	res["date_applied"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetTimeValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDateApplied(val)
		}
		return nil
	}
	return res
}

// Serialize serializes information the current object
func (m *TagsItemCompaniesGetResponse_companies) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("company", m.GetCompany())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteTimeValue("date_applied", m.GetDateApplied())
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
func (m *TagsItemCompaniesGetResponse_companies) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCompany sets the company property value. The company property
func (m *TagsItemCompaniesGetResponse_companies) SetCompany(value TagsItemCompaniesGetResponse_companies_companyable) {
	m.company = value
}

// SetDateApplied sets the date_applied property value. The date_applied property
func (m *TagsItemCompaniesGetResponse_companies) SetDateApplied(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	m.date_applied = value
}

type TagsItemCompaniesGetResponse_companiesable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCompany() TagsItemCompaniesGetResponse_companies_companyable
	GetDateApplied() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	SetCompany(value TagsItemCompaniesGetResponse_companies_companyable)
	SetDateApplied(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
}
