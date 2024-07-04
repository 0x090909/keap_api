package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use AppointmentsModelGetResponseable instead.
type AppointmentsModelResponse struct {
	AppointmentsModelGetResponse
}

// NewAppointmentsModelResponse instantiates a new AppointmentsModelResponse and sets the default values.
func NewAppointmentsModelResponse() *AppointmentsModelResponse {
	m := &AppointmentsModelResponse{
		AppointmentsModelGetResponse: *NewAppointmentsModelGetResponse(),
	}
	return m
}

// CreateAppointmentsModelResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAppointmentsModelResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAppointmentsModelResponse(), nil
}

// Deprecated: This class is obsolete. Use AppointmentsModelGetResponseable instead.
type AppointmentsModelResponseable interface {
	AppointmentsModelGetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
