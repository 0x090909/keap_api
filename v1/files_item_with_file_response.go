package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use FilesItemWithFileGetResponseable instead.
type FilesItemWithFileResponse struct {
	FilesItemWithFileGetResponse
}

// NewFilesItemWithFileResponse instantiates a new FilesItemWithFileResponse and sets the default values.
func NewFilesItemWithFileResponse() *FilesItemWithFileResponse {
	m := &FilesItemWithFileResponse{
		FilesItemWithFileGetResponse: *NewFilesItemWithFileGetResponse(),
	}
	return m
}

// CreateFilesItemWithFileResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFilesItemWithFileResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewFilesItemWithFileResponse(), nil
}

// Deprecated: This class is obsolete. Use FilesItemWithFileGetResponseable instead.
type FilesItemWithFileResponseable interface {
	FilesItemWithFileGetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
