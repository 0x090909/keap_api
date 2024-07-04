package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use NotesModelGetResponseable instead.
type NotesModelResponse struct {
	NotesModelGetResponse
}

// NewNotesModelResponse instantiates a new NotesModelResponse and sets the default values.
func NewNotesModelResponse() *NotesModelResponse {
	m := &NotesModelResponse{
		NotesModelGetResponse: *NewNotesModelGetResponse(),
	}
	return m
}

// CreateNotesModelResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNotesModelResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewNotesModelResponse(), nil
}

// Deprecated: This class is obsolete. Use NotesModelGetResponseable instead.
type NotesModelResponseable interface {
	NotesModelGetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
