package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use NotesGetResponseable instead.
type NotesResponse struct {
	NotesGetResponse
}

// NewNotesResponse instantiates a new NotesResponse and sets the default values.
func NewNotesResponse() *NotesResponse {
	m := &NotesResponse{
		NotesGetResponse: *NewNotesGetResponse(),
	}
	return m
}

// CreateNotesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNotesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewNotesResponse(), nil
}

// Deprecated: This class is obsolete. Use NotesGetResponseable instead.
type NotesResponseable interface {
	NotesGetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
