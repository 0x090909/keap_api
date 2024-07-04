package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FilesGetResponseable instead.
type FilesResponse struct {
	FilesGetResponse
}

// NewFilesResponse instantiates a new FilesResponse and sets the default values.
func NewFilesResponse() *FilesResponse {
	m := &FilesResponse{
		FilesGetResponse: *NewFilesGetResponse(),
	}
	return m
}

// CreateFilesResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFilesResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewFilesResponse(), nil
}

// Deprecated: This class is obsolete. Use FilesGetResponseable instead.
type FilesResponseable interface {
	FilesGetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
