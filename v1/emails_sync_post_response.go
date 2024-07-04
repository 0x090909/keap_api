package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EmailsSyncPostResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The emails property
	emails []EmailsSyncPostResponse_emailsable
	// The errors property
	errors []EmailsSyncPostResponse_errorsable
}

// NewEmailsSyncPostResponse instantiates a new EmailsSyncPostResponse and sets the default values.
func NewEmailsSyncPostResponse() *EmailsSyncPostResponse {
	m := &EmailsSyncPostResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateEmailsSyncPostResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEmailsSyncPostResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewEmailsSyncPostResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EmailsSyncPostResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetEmails gets the emails property value. The emails property
// returns a []EmailsSyncPostResponse_emailsable when successful
func (m *EmailsSyncPostResponse) GetEmails() []EmailsSyncPostResponse_emailsable {
	return m.emails
}

// GetErrors gets the errors property value. The errors property
// returns a []EmailsSyncPostResponse_errorsable when successful
func (m *EmailsSyncPostResponse) GetErrors() []EmailsSyncPostResponse_errorsable {
	return m.errors
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EmailsSyncPostResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["emails"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateEmailsSyncPostResponse_emailsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]EmailsSyncPostResponse_emailsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(EmailsSyncPostResponse_emailsable)
				}
			}
			m.SetEmails(res)
		}
		return nil
	}
	res["errors"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateEmailsSyncPostResponse_errorsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]EmailsSyncPostResponse_errorsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(EmailsSyncPostResponse_errorsable)
				}
			}
			m.SetErrors(res)
		}
		return nil
	}
	return res
}

// Serialize serializes information the current object
func (m *EmailsSyncPostResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetEmails() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEmails()))
		for i, v := range m.GetEmails() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("emails", cast)
		if err != nil {
			return err
		}
	}
	if m.GetErrors() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetErrors()))
		for i, v := range m.GetErrors() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("errors", cast)
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
func (m *EmailsSyncPostResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetEmails sets the emails property value. The emails property
func (m *EmailsSyncPostResponse) SetEmails(value []EmailsSyncPostResponse_emailsable) {
	m.emails = value
}

// SetErrors sets the errors property value. The errors property
func (m *EmailsSyncPostResponse) SetErrors(value []EmailsSyncPostResponse_errorsable) {
	m.errors = value
}

type EmailsSyncPostResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetEmails() []EmailsSyncPostResponse_emailsable
	GetErrors() []EmailsSyncPostResponse_errorsable
	SetEmails(value []EmailsSyncPostResponse_emailsable)
	SetErrors(value []EmailsSyncPostResponse_errorsable)
}
