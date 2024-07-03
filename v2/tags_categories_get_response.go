package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TagsCategoriesGetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The next_page_token property
	next_page_token *string
	// The tag_categories property
	tag_categories []TagsCategoriesGetResponse_tag_categoriesable
}

// NewTagsCategoriesGetResponse instantiates a new TagsCategoriesGetResponse and sets the default values.
func NewTagsCategoriesGetResponse() *TagsCategoriesGetResponse {
	m := &TagsCategoriesGetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateTagsCategoriesGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTagsCategoriesGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTagsCategoriesGetResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TagsCategoriesGetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TagsCategoriesGetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
	res["tag_categories"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateTagsCategoriesGetResponse_tag_categoriesFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]TagsCategoriesGetResponse_tag_categoriesable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(TagsCategoriesGetResponse_tag_categoriesable)
				}
			}
			m.SetTagCategories(res)
		}
		return nil
	}
	return res
}

// GetNextPageToken gets the next_page_token property value. The next_page_token property
// returns a *string when successful
func (m *TagsCategoriesGetResponse) GetNextPageToken() *string {
	return m.next_page_token
}

// GetTagCategories gets the tag_categories property value. The tag_categories property
// returns a []TagsCategoriesGetResponse_tag_categoriesable when successful
func (m *TagsCategoriesGetResponse) GetTagCategories() []TagsCategoriesGetResponse_tag_categoriesable {
	return m.tag_categories
}

// Serialize serializes information the current object
func (m *TagsCategoriesGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("next_page_token", m.GetNextPageToken())
		if err != nil {
			return err
		}
	}
	if m.GetTagCategories() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTagCategories()))
		for i, v := range m.GetTagCategories() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("tag_categories", cast)
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
func (m *TagsCategoriesGetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetNextPageToken sets the next_page_token property value. The next_page_token property
func (m *TagsCategoriesGetResponse) SetNextPageToken(value *string) {
	m.next_page_token = value
}

// SetTagCategories sets the tag_categories property value. The tag_categories property
func (m *TagsCategoriesGetResponse) SetTagCategories(value []TagsCategoriesGetResponse_tag_categoriesable) {
	m.tag_categories = value
}

type TagsCategoriesGetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetNextPageToken() *string
	GetTagCategories() []TagsCategoriesGetResponse_tag_categoriesable
	SetNextPageToken(value *string)
	SetTagCategories(value []TagsCategoriesGetResponse_tag_categoriesable)
}
