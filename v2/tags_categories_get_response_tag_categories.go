package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TagsCategoriesGetResponse_tag_categories struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The category property
	category TagsCategoriesGetResponse_tag_categories_categoryable
	// The description property
	description *string
	// The id property
	id *string
	// The name property
	name *string
}

// NewTagsCategoriesGetResponse_tag_categories instantiates a new TagsCategoriesGetResponse_tag_categories and sets the default values.
func NewTagsCategoriesGetResponse_tag_categories() *TagsCategoriesGetResponse_tag_categories {
	m := &TagsCategoriesGetResponse_tag_categories{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateTagsCategoriesGetResponse_tag_categoriesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTagsCategoriesGetResponse_tag_categoriesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTagsCategoriesGetResponse_tag_categories(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TagsCategoriesGetResponse_tag_categories) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCategory gets the category property value. The category property
// returns a TagsCategoriesGetResponse_tag_categories_categoryable when successful
func (m *TagsCategoriesGetResponse_tag_categories) GetCategory() TagsCategoriesGetResponse_tag_categories_categoryable {
	return m.category
}

// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *TagsCategoriesGetResponse_tag_categories) GetDescription() *string {
	return m.description
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TagsCategoriesGetResponse_tag_categories) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["category"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateTagsCategoriesGetResponse_tag_categories_categoryFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCategory(val.(TagsCategoriesGetResponse_tag_categories_categoryable))
		}
		return nil
	}
	res["description"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDescription(val)
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
	res["name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetName(val)
		}
		return nil
	}
	return res
}

// GetId gets the id property value. The id property
// returns a *string when successful
func (m *TagsCategoriesGetResponse_tag_categories) GetId() *string {
	return m.id
}

// GetName gets the name property value. The name property
// returns a *string when successful
func (m *TagsCategoriesGetResponse_tag_categories) GetName() *string {
	return m.name
}

// Serialize serializes information the current object
func (m *TagsCategoriesGetResponse_tag_categories) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("category", m.GetCategory())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("description", m.GetDescription())
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
		err := writer.WriteStringValue("name", m.GetName())
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
func (m *TagsCategoriesGetResponse_tag_categories) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCategory sets the category property value. The category property
func (m *TagsCategoriesGetResponse_tag_categories) SetCategory(value TagsCategoriesGetResponse_tag_categories_categoryable) {
	m.category = value
}

// SetDescription sets the description property value. The description property
func (m *TagsCategoriesGetResponse_tag_categories) SetDescription(value *string) {
	m.description = value
}

// SetId sets the id property value. The id property
func (m *TagsCategoriesGetResponse_tag_categories) SetId(value *string) {
	m.id = value
}

// SetName sets the name property value. The name property
func (m *TagsCategoriesGetResponse_tag_categories) SetName(value *string) {
	m.name = value
}

type TagsCategoriesGetResponse_tag_categoriesable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCategory() TagsCategoriesGetResponse_tag_categories_categoryable
	GetDescription() *string
	GetId() *string
	GetName() *string
	SetCategory(value TagsCategoriesGetResponse_tag_categories_categoryable)
	SetDescription(value *string)
	SetId(value *string)
	SetName(value *string)
}
