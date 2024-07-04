package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TagsItemCompaniesGetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The companies property
	companies []TagsItemCompaniesGetResponse_companiesable
	// The count property
	count *int32
	// The next property
	next *string
	// The previous property
	previous *string
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

// GetCompanies gets the companies property value. The companies property
// returns a []TagsItemCompaniesGetResponse_companiesable when successful
func (m *TagsItemCompaniesGetResponse) GetCompanies() []TagsItemCompaniesGetResponse_companiesable {
	return m.companies
}

// GetCount gets the count property value. The count property
// returns a *int32 when successful
func (m *TagsItemCompaniesGetResponse) GetCount() *int32 {
	return m.count
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TagsItemCompaniesGetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["companies"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateTagsItemCompaniesGetResponse_companiesFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]TagsItemCompaniesGetResponse_companiesable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(TagsItemCompaniesGetResponse_companiesable)
				}
			}
			m.SetCompanies(res)
		}
		return nil
	}
	res["count"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCount(val)
		}
		return nil
	}
	res["next"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNext(val)
		}
		return nil
	}
	res["previous"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPrevious(val)
		}
		return nil
	}
	return res
}

// GetNext gets the next property value. The next property
// returns a *string when successful
func (m *TagsItemCompaniesGetResponse) GetNext() *string {
	return m.next
}

// GetPrevious gets the previous property value. The previous property
// returns a *string when successful
func (m *TagsItemCompaniesGetResponse) GetPrevious() *string {
	return m.previous
}

// Serialize serializes information the current object
func (m *TagsItemCompaniesGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetCompanies() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCompanies()))
		for i, v := range m.GetCompanies() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("companies", cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("count", m.GetCount())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("next", m.GetNext())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("previous", m.GetPrevious())
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

// SetCompanies sets the companies property value. The companies property
func (m *TagsItemCompaniesGetResponse) SetCompanies(value []TagsItemCompaniesGetResponse_companiesable) {
	m.companies = value
}

// SetCount sets the count property value. The count property
func (m *TagsItemCompaniesGetResponse) SetCount(value *int32) {
	m.count = value
}

// SetNext sets the next property value. The next property
func (m *TagsItemCompaniesGetResponse) SetNext(value *string) {
	m.next = value
}

// SetPrevious sets the previous property value. The previous property
func (m *TagsItemCompaniesGetResponse) SetPrevious(value *string) {
	m.previous = value
}

type TagsItemCompaniesGetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCompanies() []TagsItemCompaniesGetResponse_companiesable
	GetCount() *int32
	GetNext() *string
	GetPrevious() *string
	SetCompanies(value []TagsItemCompaniesGetResponse_companiesable)
	SetCount(value *int32)
	SetNext(value *string)
	SetPrevious(value *string)
}
