package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use TagsItemCompaniesGetResponseable instead.
type TagsItemCompaniesResponse struct {
	TagsItemCompaniesGetResponse
}

// NewTagsItemCompaniesResponse instantiates a new TagsItemCompaniesResponse and sets the default values.
func NewTagsItemCompaniesResponse() *TagsItemCompaniesResponse {
	m := &TagsItemCompaniesResponse{
		TagsItemCompaniesGetResponse: *NewTagsItemCompaniesGetResponse(),
	}
	return m
}

// CreateTagsItemCompaniesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTagsItemCompaniesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTagsItemCompaniesResponse(), nil
}

// Deprecated: This class is obsolete. Use TagsItemCompaniesGetResponseable instead.
type TagsItemCompaniesResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	TagsItemCompaniesGetResponseable
}
