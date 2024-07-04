package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TagsGetResponse_tags_category struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The category description
	description *string
	// The id for the tag category
	id *int64
	// The category name
	name *string
}

// NewTagsGetResponse_tags_category instantiates a new TagsGetResponse_tags_category and sets the default values.
func NewTagsGetResponse_tags_category() *TagsGetResponse_tags_category {
	m := &TagsGetResponse_tags_category{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateTagsGetResponse_tags_categoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTagsGetResponse_tags_categoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTagsGetResponse_tags_category(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TagsGetResponse_tags_category) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetDescription gets the description property value. The category description
// returns a *string when successful
func (m *TagsGetResponse_tags_category) GetDescription() *string {
	return m.description
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TagsGetResponse_tags_category) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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

// GetId gets the id property value. The id for the tag category
// returns a *int64 when successful
func (m *TagsGetResponse_tags_category) GetId() *int64 {
	return m.id
}

// GetName gets the name property value. The category name
// returns a *string when successful
func (m *TagsGetResponse_tags_category) GetName() *string {
	return m.name
}

// Serialize serializes information the current object
func (m *TagsGetResponse_tags_category) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *TagsGetResponse_tags_category) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetDescription sets the description property value. The category description
func (m *TagsGetResponse_tags_category) SetDescription(value *string) {
	m.description = value
}

// SetId sets the id property value. The id for the tag category
func (m *TagsGetResponse_tags_category) SetId(value *int64) {
	m.id = value
}

// SetName sets the name property value. The category name
func (m *TagsGetResponse_tags_category) SetName(value *string) {
	m.name = value
}

type TagsGetResponse_tags_categoryable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetDescription() *string
	GetId() *int64
	GetName() *string
	SetDescription(value *string)
	SetId(value *int64)
	SetName(value *string)
}
