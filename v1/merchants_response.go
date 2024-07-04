package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use MerchantsGetResponseable instead.
type MerchantsResponse struct {
	MerchantsGetResponse
}

// NewMerchantsResponse instantiates a new MerchantsResponse and sets the default values.
func NewMerchantsResponse() *MerchantsResponse {
	m := &MerchantsResponse{
		MerchantsGetResponse: *NewMerchantsGetResponse(),
	}
	return m
}

// CreateMerchantsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMerchantsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewMerchantsResponse(), nil
}

// Deprecated: This class is obsolete. Use MerchantsGetResponseable instead.
type MerchantsResponseable interface {
	MerchantsGetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
