package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EmailsBatchAddPostRequestBody struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The emails property
	emails []EmailsBatchAddPostRequestBody_emailsable
}

// NewEmailsBatchAddPostRequestBody instantiates a new EmailsBatchAddPostRequestBody and sets the default values.
func NewEmailsBatchAddPostRequestBody() *EmailsBatchAddPostRequestBody {
	m := &EmailsBatchAddPostRequestBody{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateEmailsBatchAddPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEmailsBatchAddPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewEmailsBatchAddPostRequestBody(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EmailsBatchAddPostRequestBody) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetEmails gets the emails property value. The emails property
// returns a []EmailsBatchAddPostRequestBody_emailsable when successful
func (m *EmailsBatchAddPostRequestBody) GetEmails() []EmailsBatchAddPostRequestBody_emailsable {
	return m.emails
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EmailsBatchAddPostRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["emails"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateEmailsBatchAddPostRequestBody_emailsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]EmailsBatchAddPostRequestBody_emailsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(EmailsBatchAddPostRequestBody_emailsable)
				}
			}
			m.SetEmails(res)
		}
		return nil
	}
	return res
}

// Serialize serializes information the current object
func (m *EmailsBatchAddPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
	{
		err := writer.WriteAdditionalData(m.GetAdditionalData())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EmailsBatchAddPostRequestBody) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetEmails sets the emails property value. The emails property
func (m *EmailsBatchAddPostRequestBody) SetEmails(value []EmailsBatchAddPostRequestBody_emailsable) {
	m.emails = value
}

type EmailsBatchAddPostRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetEmails() []EmailsBatchAddPostRequestBody_emailsable
	SetEmails(value []EmailsBatchAddPostRequestBody_emailsable)
}
