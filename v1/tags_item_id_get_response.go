package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TagsItemIdGetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The category property
	category TagsItemIdGetResponse_categoryable
	// The tag description
	description *string
	// The id for the tag
	id *int64
	// The tag name
	name *string
}

// NewTagsItemIdGetResponse instantiates a new TagsItemIdGetResponse and sets the default values.
func NewTagsItemIdGetResponse() *TagsItemIdGetResponse {
	m := &TagsItemIdGetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateTagsItemIdGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTagsItemIdGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTagsItemIdGetResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TagsItemIdGetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCategory gets the category property value. The category property
// returns a TagsItemIdGetResponse_categoryable when successful
func (m *TagsItemIdGetResponse) GetCategory() TagsItemIdGetResponse_categoryable {
	return m.category
}

// GetDescription gets the description property value. The tag description
// returns a *string when successful
func (m *TagsItemIdGetResponse) GetDescription() *string {
	return m.description
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TagsItemIdGetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["category"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateTagsItemIdGetResponse_categoryFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCategory(val.(TagsItemIdGetResponse_categoryable))
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
func (m *TagsItemIdGetResponse) GetId() *int64 {
	return m.id
}

// GetName gets the name property value. The tag name
// returns a *string when successful
func (m *TagsItemIdGetResponse) GetName() *string {
	return m.name
}

// Serialize serializes information the current object
func (m *TagsItemIdGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *TagsItemIdGetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCategory sets the category property value. The category property
func (m *TagsItemIdGetResponse) SetCategory(value TagsItemIdGetResponse_categoryable) {
	m.category = value
}

// SetDescription sets the description property value. The tag description
func (m *TagsItemIdGetResponse) SetDescription(value *string) {
	m.description = value
}

// SetId sets the id property value. The id for the tag
func (m *TagsItemIdGetResponse) SetId(value *int64) {
	m.id = value
}

// SetName sets the name property value. The tag name
func (m *TagsItemIdGetResponse) SetName(value *string) {
	m.name = value
}

type TagsItemIdGetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCategory() TagsItemIdGetResponse_categoryable
	GetDescription() *string
	GetId() *int64
	GetName() *string
	SetCategory(value TagsItemIdGetResponse_categoryable)
	SetDescription(value *string)
	SetId(value *int64)
	SetName(value *string)
}
