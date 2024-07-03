package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TagsItemCompaniesGetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The next_page_token property
	next_page_token *string
	// The tagged_companies property
	tagged_companies []TagsItemCompaniesGetResponse_tagged_companiesable
}

// NewTagsItemCompaniesGetResponse instantiates a new TagsItemCompaniesGetResponse and sets the default values.
func NewTagsItemCompaniesGetResponse() *TagsItemCompaniesGetResponse {
	m := &TagsItemCompaniesGetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateTagsItemCompaniesGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTagsItemCompaniesGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTagsItemCompaniesGetResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TagsItemCompaniesGetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TagsItemCompaniesGetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["next_page_token"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNextPageToken(val)
		}
		return nil
	}
	res["tagged_companies"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateTagsItemCompaniesGetResponse_tagged_companiesFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]TagsItemCompaniesGetResponse_tagged_companiesable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(TagsItemCompaniesGetResponse_tagged_companiesable)
				}
			}
			m.SetTaggedCompanies(res)
		}
		return nil
	}
	return res
}

// GetNextPageToken gets the next_page_token property value. The next_page_token property
// returns a *string when successful
func (m *TagsItemCompaniesGetResponse) GetNextPageToken() *string {
	return m.next_page_token
}

// GetTaggedCompanies gets the tagged_companies property value. The tagged_companies property
// returns a []TagsItemCompaniesGetResponse_tagged_companiesable when successful
func (m *TagsItemCompaniesGetResponse) GetTaggedCompanies() []TagsItemCompaniesGetResponse_tagged_companiesable {
	return m.tagged_companies
}

// Serialize serializes information the current object
func (m *TagsItemCompaniesGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("next_page_token", m.GetNextPageToken())
		if err != nil {
			return err
		}
	}
	if m.GetTaggedCompanies() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTaggedCompanies()))
		for i, v := range m.GetTaggedCompanies() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("tagged_companies", cast)
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
func (m *TagsItemCompaniesGetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetNextPageToken sets the next_page_token property value. The next_page_token property
func (m *TagsItemCompaniesGetResponse) SetNextPageToken(value *string) {
	m.next_page_token = value
}

// SetTaggedCompanies sets the tagged_companies property value. The tagged_companies property
func (m *TagsItemCompaniesGetResponse) SetTaggedCompanies(value []TagsItemCompaniesGetResponse_tagged_companiesable) {
	m.tagged_companies = value
}

type TagsItemCompaniesGetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetNextPageToken() *string
	GetTaggedCompanies() []TagsItemCompaniesGetResponse_tagged_companiesable
	SetNextPageToken(value *string)
	SetTaggedCompanies(value []TagsItemCompaniesGetResponse_tagged_companiesable)
}
