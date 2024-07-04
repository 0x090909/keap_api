package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use AppointmentsItemWithAppointmentGetResponseable instead.
type AppointmentsItemWithAppointmentResponse struct {
	AppointmentsItemWithAppointmentGetResponse
}

// NewAppointmentsItemWithAppointmentResponse instantiates a new AppointmentsItemWithAppointmentResponse and sets the default values.
func NewAppointmentsItemWithAppointmentResponse() *AppointmentsItemWithAppointmentResponse {
	m := &AppointmentsItemWithAppointmentResponse{
		AppointmentsItemWithAppointmentGetResponse: *NewAppointmentsItemWithAppointmentGetResponse(),
	}
	return m
}

// CreateAppointmentsItemWithAppointmentResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAppointmentsItemWithAppointmentResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAppointmentsItemWithAppointmentResponse(), nil
}

// Deprecated: This class is obsolete. Use AppointmentsItemWithAppointmentGetResponseable instead.
type AppointmentsItemWithAppointmentResponseable interface {
	AppointmentsItemWithAppointmentGetResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
