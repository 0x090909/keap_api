package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TagsGetResponse_tags struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The category property
	category TagsGetResponse_tags_categoryable
	// The tag description
	description *string
	// The id for the tag
	id *int64
	// The tag name
	name *string
}

// NewTagsGetResponse_tags instantiates a new TagsGetResponse_tags and sets the default values.
func NewTagsGetResponse_tags() *TagsGetResponse_tags {
	m := &TagsGetResponse_tags{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateTagsGetResponse_tagsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTagsGetResponse_tagsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTagsGetResponse_tags(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TagsGetResponse_tags) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCategory gets the category property value. The category property
// returns a TagsGetResponse_tags_categoryable when successful
func (m *TagsGetResponse_tags) GetCategory() TagsGetResponse_tags_categoryable {
	return m.category
}

// GetDescription gets the description property value. The tag description
// returns a *string when successful
func (m *TagsGetResponse_tags) GetDescription() *string {
	return m.description
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TagsGetResponse_tags) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["category"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateTagsGetResponse_tags_categoryFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCategory(val.(TagsGetResponse_tags_categoryable))
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
		val, err := n.GetInt64Value()
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

// GetId gets the id property value. The id for the tag
// returns a *int64 when successful
func (m *TagsGetResponse_tags) GetId() *int64 {
	return m.id
}

// GetName gets the name property value. The tag name
// returns a *string when successful
func (m *TagsGetResponse_tags) GetName() *string {
	return m.name
}

// Serialize serializes information the current object
func (m *TagsGetResponse_tags) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
		err := writer.WriteInt64Value("id", m.GetId())
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
func (m *TagsGetResponse_tags) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCategory sets the category property value. The category property
func (m *TagsGetResponse_tags) SetCategory(value TagsGetResponse_tags_categoryable) {
	m.category = value
}

// SetDescription sets the description property value. The tag description
func (m *TagsGetResponse_tags) SetDescription(value *string) {
	m.description = value
}

// SetId sets the id property value. The id for the tag
func (m *TagsGetResponse_tags) SetId(value *int64) {
	m.id = value
}

// SetName sets the name property value. The tag name
func (m *TagsGetResponse_tags) SetName(value *string) {
	m.name = value
}

type TagsGetResponse_tagsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCategory() TagsGetResponse_tags_categoryable
	GetDescription() *string
	GetId() *int64
	GetName() *string
	SetCategory(value TagsGetResponse_tags_categoryable)
	SetDescription(value *string)
	SetId(value *int64)
	SetName(value *string)
}
