package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use AppointmentsModelCustomFieldsPostResponseable instead.
type AppointmentsModelCustomFieldsResponse struct {
	AppointmentsModelCustomFieldsPostResponse
}

// NewAppointmentsModelCustomFieldsResponse instantiates a new AppointmentsModelCustomFieldsResponse and sets the default values.
func NewAppointmentsModelCustomFieldsResponse() *AppointmentsModelCustomFieldsResponse {
	m := &AppointmentsModelCustomFieldsResponse{
		AppointmentsModelCustomFieldsPostResponse: *NewAppointmentsModelCustomFieldsPostResponse(),
	}
	return m
}

// CreateAppointmentsModelCustomFieldsResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAppointmentsModelCustomFieldsResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAppointmentsModelCustomFieldsResponse(), nil
}

// Deprecated: This class is obsolete. Use AppointmentsModelCustomFieldsPostResponseable instead.
type AppointmentsModelCustomFieldsResponseable interface {
	AppointmentsModelCustomFieldsPostResponseable
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
