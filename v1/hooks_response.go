package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use HooksPostResponseable instead.
type HooksResponse struct {
	HooksPostResponse
}

// NewHooksResponse instantiates a new HooksResponse and sets the default values.
func NewHooksResponse() *HooksResponse {
	m := &HooksResponse{
		HooksPostResponse: *NewHooksPostResponse(),
	}
	return m
}

// CreateHooksResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateHooksResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewHooksResponse(), nil
}

// Deprecated: This class is obsolete. Use HooksPostResponseable instead.
type HooksResponseable interface {
	HooksPostResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
