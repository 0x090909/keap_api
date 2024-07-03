package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TagsItemWithTag_PatchRequestBody struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The category property
	category TagsItemWithTag_PatchRequestBody_categoryable
	// The description property
	description *string
	// The name property
	name *string
}

// NewTagsItemWithTag_PatchRequestBody instantiates a new TagsItemWithTag_PatchRequestBody and sets the default values.
func NewTagsItemWithTag_PatchRequestBody() *TagsItemWithTag_PatchRequestBody {
	m := &TagsItemWithTag_PatchRequestBody{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateTagsItemWithTag_PatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTagsItemWithTag_PatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTagsItemWithTag_PatchRequestBody(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TagsItemWithTag_PatchRequestBody) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCategory gets the category property value. The category property
// returns a TagsItemWithTag_PatchRequestBody_categoryable when successful
func (m *TagsItemWithTag_PatchRequestBody) GetCategory() TagsItemWithTag_PatchRequestBody_categoryable {
	return m.category
}

// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *TagsItemWithTag_PatchRequestBody) GetDescription() *string {
	return m.description
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TagsItemWithTag_PatchRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["category"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateTagsItemWithTag_PatchRequestBody_categoryFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCategory(val.(TagsItemWithTag_PatchRequestBody_categoryable))
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

// GetName gets the name property value. The name property
// returns a *string when successful
func (m *TagsItemWithTag_PatchRequestBody) GetName() *string {
	return m.name
}

// Serialize serializes information the current object
func (m *TagsItemWithTag_PatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *TagsItemWithTag_PatchRequestBody) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCategory sets the category property value. The category property
func (m *TagsItemWithTag_PatchRequestBody) SetCategory(value TagsItemWithTag_PatchRequestBody_categoryable) {
	m.category = value
}

// SetDescription sets the description property value. The description property
func (m *TagsItemWithTag_PatchRequestBody) SetDescription(value *string) {
	m.description = value
}

// SetName sets the name property value. The name property
func (m *TagsItemWithTag_PatchRequestBody) SetName(value *string) {
	m.name = value
}

type TagsItemWithTag_PatchRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCategory() TagsItemWithTag_PatchRequestBody_categoryable
	GetDescription() *string
	GetName() *string
	SetCategory(value TagsItemWithTag_PatchRequestBody_categoryable)
	SetDescription(value *string)
	SetName(value *string)
}
