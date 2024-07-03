package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use TagsCategoriesItemWithTag_category_GetResponseable instead.
type TagsCategoriesItemWithTag_category_Response struct {
	TagsCategoriesItemWithTag_category_GetResponse
}

// NewTagsCategoriesItemWithTag_category_Response instantiates a new TagsCategoriesItemWithTag_category_Response and sets the default values.
func NewTagsCategoriesItemWithTag_category_Response() *TagsCategoriesItemWithTag_category_Response {
	m := &TagsCategoriesItemWithTag_category_Response{
		TagsCategoriesItemWithTag_category_GetResponse: *NewTagsCategoriesItemWithTag_category_GetResponse(),
	}
	return m
}

// CreateTagsCategoriesItemWithTag_category_ResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTagsCategoriesItemWithTag_category_ResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTagsCategoriesItemWithTag_category_Response(), nil
}

// Deprecated: This class is obsolete. Use TagsCategoriesItemWithTag_category_GetResponseable instead.
type TagsCategoriesItemWithTag_category_Responseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	TagsCategoriesItemWithTag_category_GetResponseable
}
