package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BusinessProfileResponse profile information about the business that owns/uses this account
// Deprecated: This class is obsolete. Use BusinessProfileGetResponseable instead.
type BusinessProfileResponse struct {
	BusinessProfileGetResponse
}

// NewBusinessProfileResponse instantiates a new BusinessProfileResponse and sets the default values.
func NewBusinessProfileResponse() *BusinessProfileResponse {
	m := &BusinessProfileResponse{
		BusinessProfileGetResponse: *NewBusinessProfileGetResponse(),
	}
	return m
}

// CreateBusinessProfileResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateBusinessProfileResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewBusinessProfileResponse(), nil
}

// Deprecated: This class is obsolete. Use BusinessProfileGetResponseable instead.
type BusinessProfileResponseable interface {
	BusinessProfileGetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
