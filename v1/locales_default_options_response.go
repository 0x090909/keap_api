package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use LocalesDefaultOptionsGetResponseable instead.
type LocalesDefaultOptionsResponse struct {
	LocalesDefaultOptionsGetResponse
}

// NewLocalesDefaultOptionsResponse instantiates a new LocalesDefaultOptionsResponse and sets the default values.
func NewLocalesDefaultOptionsResponse() *LocalesDefaultOptionsResponse {
	m := &LocalesDefaultOptionsResponse{
		LocalesDefaultOptionsGetResponse: *NewLocalesDefaultOptionsGetResponse(),
	}
	return m
}

// CreateLocalesDefaultOptionsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLocalesDefaultOptionsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewLocalesDefaultOptionsResponse(), nil
}

// Deprecated: This class is obsolete. Use LocalesDefaultOptionsGetResponseable instead.
type LocalesDefaultOptionsResponseable interface {
	LocalesDefaultOptionsGetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
