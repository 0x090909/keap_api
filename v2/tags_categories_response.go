package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use TagsCategoriesGetResponseable instead.
type TagsCategoriesResponse struct {
	TagsCategoriesGetResponse
}

// NewTagsCategoriesResponse instantiates a new TagsCategoriesResponse and sets the default values.
func NewTagsCategoriesResponse() *TagsCategoriesResponse {
	m := &TagsCategoriesResponse{
		TagsCategoriesGetResponse: *NewTagsCategoriesGetResponse(),
	}
	return m
}

// CreateTagsCategoriesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTagsCategoriesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTagsCategoriesResponse(), nil
}

// Deprecated: This class is obsolete. Use TagsCategoriesGetResponseable instead.
type TagsCategoriesResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	TagsCategoriesGetResponseable
}
