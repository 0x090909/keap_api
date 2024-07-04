package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AppointmentsGetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The appointments property
	appointments []AppointmentsGetResponse_appointmentsable
	// The count property
	count *int32
	// The next property
	next *string
	// The previous property
	previous *string
	// The sync_token property
	sync_token *string
}

// NewAppointmentsGetResponse instantiates a new AppointmentsGetResponse and sets the default values.
func NewAppointmentsGetResponse() *AppointmentsGetResponse {
	m := &AppointmentsGetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateAppointmentsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAppointmentsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAppointmentsGetResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AppointmentsGetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAppointments gets the appointments property value. The appointments property
// returns a []AppointmentsGetResponse_appointmentsable when successful
func (m *AppointmentsGetResponse) GetAppointments() []AppointmentsGetResponse_appointmentsable {
	return m.appointments
}

// GetCount gets the count property value. The count property
// returns a *int32 when successful
func (m *AppointmentsGetResponse) GetCount() *int32 {
	return m.count
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AppointmentsGetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["appointments"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateAppointmentsGetResponse_appointmentsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]AppointmentsGetResponse_appointmentsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(AppointmentsGetResponse_appointmentsable)
				}
			}
			m.SetAppointments(res)
		}
		return nil
	}
	res["count"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCount(val)
		}
		return nil
	}
	res["next"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNext(val)
		}
		return nil
	}
	res["previous"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPrevious(val)
		}
		return nil
	}
	res["sync_token"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSyncToken(val)
		}
		return nil
	}
	return res
}

// GetNext gets the next property value. The next property
// returns a *string when successful
func (m *AppointmentsGetResponse) GetNext() *string {
	return m.next
}

// GetPrevious gets the previous property value. The previous property
// returns a *string when successful
func (m *AppointmentsGetResponse) GetPrevious() *string {
	return m.previous
}

// GetSyncToken gets the sync_token property value. The sync_token property
// returns a *string when successful
func (m *AppointmentsGetResponse) GetSyncToken() *string {
	return m.sync_token
}

// Serialize serializes information the current object
func (m *AppointmentsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetAppointments() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAppointments()))
		for i, v := range m.GetAppointments() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("appointments", cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("count", m.GetCount())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("next", m.GetNext())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("previous", m.GetPrevious())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("sync_token", m.GetSyncToken())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteAdditionalData(m.GetAdditionalData())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AppointmentsGetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAppointments sets the appointments property value. The appointments property
func (m *AppointmentsGetResponse) SetAppointments(value []AppointmentsGetResponse_appointmentsable) {
	m.appointments = value
}

// SetCount sets the count property value. The count property
func (m *AppointmentsGetResponse) SetCount(value *int32) {
	m.count = value
}

// SetNext sets the next property value. The next property
func (m *AppointmentsGetResponse) SetNext(value *string) {
	m.next = value
}

// SetPrevious sets the previous property value. The previous property
func (m *AppointmentsGetResponse) SetPrevious(value *string) {
	m.previous = value
}

// SetSyncToken sets the sync_token property value. The sync_token property
func (m *AppointmentsGetResponse) SetSyncToken(value *string) {
	m.sync_token = value
}

type AppointmentsGetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAppointments() []AppointmentsGetResponse_appointmentsable
	GetCount() *int32
	GetNext() *string
	GetPrevious() *string
	GetSyncToken() *string
	SetAppointments(value []AppointmentsGetResponse_appointmentsable)
	SetCount(value *int32)
	SetNext(value *string)
	SetPrevious(value *string)
	SetSyncToken(value *string)
}
