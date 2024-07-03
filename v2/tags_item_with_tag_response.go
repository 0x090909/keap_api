package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use TagsItemWithTag_GetResponseable instead.
type TagsItemWithTag_Response struct {
	TagsItemWithTag_GetResponse
}

// NewTagsItemWithTag_Response instantiates a new TagsItemWithTag_Response and sets the default values.
func NewTagsItemWithTag_Response() *TagsItemWithTag_Response {
	m := &TagsItemWithTag_Response{
		TagsItemWithTag_GetResponse: *NewTagsItemWithTag_GetResponse(),
	}
	return m
}

// CreateTagsItemWithTag_ResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTagsItemWithTag_ResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTagsItemWithTag_Response(), nil
}

// Deprecated: This class is obsolete. Use TagsItemWithTag_GetResponseable instead.
type TagsItemWithTag_Responseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	TagsItemWithTag_GetResponseable
}
