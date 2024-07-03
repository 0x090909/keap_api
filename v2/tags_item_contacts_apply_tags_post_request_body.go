package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TagsItemContactsApplyTagsPostRequestBody struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The contact_ids property
	contact_ids []string
}

// NewTagsItemContactsApplyTagsPostRequestBody instantiates a new TagsItemContactsApplyTagsPostRequestBody and sets the default values.
func NewTagsItemContactsApplyTagsPostRequestBody() *TagsItemContactsApplyTagsPostRequestBody {
	m := &TagsItemContactsApplyTagsPostRequestBody{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateTagsItemContactsApplyTagsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTagsItemContactsApplyTagsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTagsItemContactsApplyTagsPostRequestBody(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TagsItemContactsApplyTagsPostRequestBody) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetContactIds gets the contact_ids property value. The contact_ids property
// returns a []string when successful
func (m *TagsItemContactsApplyTagsPostRequestBody) GetContactIds() []string {
	return m.contact_ids
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TagsItemContactsApplyTagsPostRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["contact_ids"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfPrimitiveValues("string")
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]string, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = *(v.(*string))
				}
			}
			m.SetContactIds(res)
		}
		return nil
	}
	return res
}

// Serialize serializes information the current object
func (m *TagsItemContactsApplyTagsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetContactIds() != nil {
		err := writer.WriteCollectionOfStringValues("contact_ids", m.GetContactIds())
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
func (m *TagsItemContactsApplyTagsPostRequestBody) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetContactIds sets the contact_ids property value. The contact_ids property
func (m *TagsItemContactsApplyTagsPostRequestBody) SetContactIds(value []string) {
	m.contact_ids = value
}

type TagsItemContactsApplyTagsPostRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetContactIds() []string
	SetContactIds(value []string)
}
