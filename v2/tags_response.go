package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use TagsGetResponseable instead.
type TagsResponse struct {
	TagsGetResponse
}

// NewTagsResponse instantiates a new TagsResponse and sets the default values.
func NewTagsResponse() *TagsResponse {
	m := &TagsResponse{
		TagsGetResponse: *NewTagsGetResponse(),
	}
	return m
}

// CreateTagsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTagsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTagsResponse(), nil
}

// Deprecated: This class is obsolete. Use TagsGetResponseable instead.
type TagsResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	TagsGetResponseable
}
