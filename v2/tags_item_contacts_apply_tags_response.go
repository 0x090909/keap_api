package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use TagsItemContactsApplyTagsPostResponseable instead.
type TagsItemContactsApplyTagsResponse struct {
	TagsItemContactsApplyTagsPostResponse
}

// NewTagsItemContactsApplyTagsResponse instantiates a new TagsItemContactsApplyTagsResponse and sets the default values.
func NewTagsItemContactsApplyTagsResponse() *TagsItemContactsApplyTagsResponse {
	m := &TagsItemContactsApplyTagsResponse{
		TagsItemContactsApplyTagsPostResponse: *NewTagsItemContactsApplyTagsPostResponse(),
	}
	return m
}

// CreateTagsItemContactsApplyTagsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTagsItemContactsApplyTagsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTagsItemContactsApplyTagsResponse(), nil
}

// Deprecated: This class is obsolete. Use TagsItemContactsApplyTagsPostResponseable instead.
type TagsItemContactsApplyTagsResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	TagsItemContactsApplyTagsPostResponseable
}
