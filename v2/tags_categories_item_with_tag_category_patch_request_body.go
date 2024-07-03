package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TagsCategoriesItemWithTag_category_PatchRequestBody struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The description property
	description *string
	// The name property
	name *string
}

// NewTagsCategoriesItemWithTag_category_PatchRequestBody instantiates a new TagsCategoriesItemWithTag_category_PatchRequestBody and sets the default values.
func NewTagsCategoriesItemWithTag_category_PatchRequestBody() *TagsCategoriesItemWithTag_category_PatchRequestBody {
	m := &TagsCategoriesItemWithTag_category_PatchRequestBody{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateTagsCategoriesItemWithTag_category_PatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTagsCategoriesItemWithTag_category_PatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTagsCategoriesItemWithTag_category_PatchRequestBody(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TagsCategoriesItemWithTag_category_PatchRequestBody) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *TagsCategoriesItemWithTag_category_PatchRequestBody) GetDescription() *string {
	return m.description
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TagsCategoriesItemWithTag_category_PatchRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
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
func (m *TagsCategoriesItemWithTag_category_PatchRequestBody) GetName() *string {
	return m.name
}

// Serialize serializes information the current object
func (m *TagsCategoriesItemWithTag_category_PatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *TagsCategoriesItemWithTag_category_PatchRequestBody) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetDescription sets the description property value. The description property
func (m *TagsCategoriesItemWithTag_category_PatchRequestBody) SetDescription(value *string) {
	m.description = value
}

// SetName sets the name property value. The name property
func (m *TagsCategoriesItemWithTag_category_PatchRequestBody) SetName(value *string) {
	m.name = value
}

type TagsCategoriesItemWithTag_category_PatchRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetDescription() *string
	GetName() *string
	SetDescription(value *string)
	SetName(value *string)
}
