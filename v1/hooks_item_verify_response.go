package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use HooksItemVerifyPostResponseable instead.
type HooksItemVerifyResponse struct {
	HooksItemVerifyPostResponse
}

// NewHooksItemVerifyResponse instantiates a new HooksItemVerifyResponse and sets the default values.
func NewHooksItemVerifyResponse() *HooksItemVerifyResponse {
	m := &HooksItemVerifyResponse{
		HooksItemVerifyPostResponse: *NewHooksItemVerifyPostResponse(),
	}
	return m
}

// CreateHooksItemVerifyResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateHooksItemVerifyResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewHooksItemVerifyResponse(), nil
}

// Deprecated: This class is obsolete. Use HooksItemVerifyPostResponseable instead.
type HooksItemVerifyResponseable interface {
	HooksItemVerifyPostResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
