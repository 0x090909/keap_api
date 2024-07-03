package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TagsItemContactsApplyTagsPostResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The results property
	results TagsItemContactsApplyTagsPostResponse_resultsable
}

// NewTagsItemContactsApplyTagsPostResponse instantiates a new TagsItemContactsApplyTagsPostResponse and sets the default values.
func NewTagsItemContactsApplyTagsPostResponse() *TagsItemContactsApplyTagsPostResponse {
	m := &TagsItemContactsApplyTagsPostResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateTagsItemContactsApplyTagsPostResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTagsItemContactsApplyTagsPostResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTagsItemContactsApplyTagsPostResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TagsItemContactsApplyTagsPostResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TagsItemContactsApplyTagsPostResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["results"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateTagsItemContactsApplyTagsPostResponse_resultsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetResults(val.(TagsItemContactsApplyTagsPostResponse_resultsable))
		}
		return nil
	}
	return res
}

// GetResults gets the results property value. The results property
// returns a TagsItemContactsApplyTagsPostResponse_resultsable when successful
func (m *TagsItemContactsApplyTagsPostResponse) GetResults() TagsItemContactsApplyTagsPostResponse_resultsable {
	return m.results
}

// Serialize serializes information the current object
func (m *TagsItemContactsApplyTagsPostResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("results", m.GetResults())
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
func (m *TagsItemContactsApplyTagsPostResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetResults sets the results property value. The results property
func (m *TagsItemContactsApplyTagsPostResponse) SetResults(value TagsItemContactsApplyTagsPostResponse_resultsable) {
	m.results = value
}

type TagsItemContactsApplyTagsPostResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetResults() TagsItemContactsApplyTagsPostResponse_resultsable
	SetResults(value TagsItemContactsApplyTagsPostResponse_resultsable)
}
