package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use NotesItemWithNoteGetResponseable instead.
type NotesItemWithNoteResponse struct {
	NotesItemWithNoteGetResponse
}

// NewNotesItemWithNoteResponse instantiates a new NotesItemWithNoteResponse and sets the default values.
func NewNotesItemWithNoteResponse() *NotesItemWithNoteResponse {
	m := &NotesItemWithNoteResponse{
		NotesItemWithNoteGetResponse: *NewNotesItemWithNoteGetResponse(),
	}
	return m
}

// CreateNotesItemWithNoteResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNotesItemWithNoteResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewNotesItemWithNoteResponse(), nil
}

// Deprecated: This class is obsolete. Use NotesItemWithNoteGetResponseable instead.
type NotesItemWithNoteResponseable interface {
	NotesItemWithNoteGetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
