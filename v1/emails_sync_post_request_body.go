package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EmailsSyncPostRequestBody struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The emails property
	emails []EmailsSyncPostRequestBody_emailsable
	// The errors property
	errors []EmailsSyncPostRequestBody_errorsable
}

// NewEmailsSyncPostRequestBody instantiates a new EmailsSyncPostRequestBody and sets the default values.
func NewEmailsSyncPostRequestBody() *EmailsSyncPostRequestBody {
	m := &EmailsSyncPostRequestBody{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateEmailsSyncPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEmailsSyncPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewEmailsSyncPostRequestBody(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EmailsSyncPostRequestBody) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetEmails gets the emails property value. The emails property
// returns a []EmailsSyncPostRequestBody_emailsable when successful
func (m *EmailsSyncPostRequestBody) GetEmails() []EmailsSyncPostRequestBody_emailsable {
	return m.emails
}

// GetErrors gets the errors property value. The errors property
// returns a []EmailsSyncPostRequestBody_errorsable when successful
func (m *EmailsSyncPostRequestBody) GetErrors() []EmailsSyncPostRequestBody_errorsable {
	return m.errors
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EmailsSyncPostRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["emails"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateEmailsSyncPostRequestBody_emailsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]EmailsSyncPostRequestBody_emailsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(EmailsSyncPostRequestBody_emailsable)
				}
			}
			m.SetEmails(res)
		}
		return nil
	}
	res["errors"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateEmailsSyncPostRequestBody_errorsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]EmailsSyncPostRequestBody_errorsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(EmailsSyncPostRequestBody_errorsable)
				}
			}
			m.SetErrors(res)
		}
		return nil
	}
	return res
}

// Serialize serializes information the current object
func (m *EmailsSyncPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *EmailsSyncPostRequestBody) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetEmails sets the emails property value. The emails property
func (m *EmailsSyncPostRequestBody) SetEmails(value []EmailsSyncPostRequestBody_emailsable) {
	m.emails = value
}

// SetErrors sets the errors property value. The errors property
func (m *EmailsSyncPostRequestBody) SetErrors(value []EmailsSyncPostRequestBody_errorsable) {
	m.errors = value
}

type EmailsSyncPostRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetEmails() []EmailsSyncPostRequestBody_emailsable
	GetErrors() []EmailsSyncPostRequestBody_errorsable
	SetEmails(value []EmailsSyncPostRequestBody_emailsable)
	SetErrors(value []EmailsSyncPostRequestBody_errorsable)
}
