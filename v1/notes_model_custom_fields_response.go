package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use NotesModelCustomFieldsPostResponseable instead.
type NotesModelCustomFieldsResponse struct {
	NotesModelCustomFieldsPostResponse
}

// NewNotesModelCustomFieldsResponse instantiates a new NotesModelCustomFieldsResponse and sets the default values.
func NewNotesModelCustomFieldsResponse() *NotesModelCustomFieldsResponse {
	m := &NotesModelCustomFieldsResponse{
		NotesModelCustomFieldsPostResponse: *NewNotesModelCustomFieldsPostResponse(),
	}
	return m
}

// CreateNotesModelCustomFieldsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNotesModelCustomFieldsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewNotesModelCustomFieldsResponse(), nil
}

// Deprecated: This class is obsolete. Use NotesModelCustomFieldsPostResponseable instead.
type NotesModelCustomFieldsResponseable interface {
	NotesModelCustomFieldsPostResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
