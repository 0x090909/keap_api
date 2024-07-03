package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EmailsBatchRemovePostRequestBody struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The email_ids property
	email_ids []string
}

// NewEmailsBatchRemovePostRequestBody instantiates a new EmailsBatchRemovePostRequestBody and sets the default values.
func NewEmailsBatchRemovePostRequestBody() *EmailsBatchRemovePostRequestBody {
	m := &EmailsBatchRemovePostRequestBody{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateEmailsBatchRemovePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEmailsBatchRemovePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewEmailsBatchRemovePostRequestBody(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EmailsBatchRemovePostRequestBody) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetEmailIds gets the email_ids property value. The email_ids property
// returns a []string when successful
func (m *EmailsBatchRemovePostRequestBody) GetEmailIds() []string {
	return m.email_ids
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EmailsBatchRemovePostRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["email_ids"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfPrimitiveValues("string")
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]string, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = *(v.(*string))
				}
			}
			m.SetEmailIds(res)
		}
		return nil
	}
	return res
}

// Serialize serializes information the current object
func (m *EmailsBatchRemovePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetEmailIds() != nil {
		err := writer.WriteCollectionOfStringValues("email_ids", m.GetEmailIds())
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
func (m *EmailsBatchRemovePostRequestBody) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetEmailIds sets the email_ids property value. The email_ids property
func (m *EmailsBatchRemovePostRequestBody) SetEmailIds(value []string) {
	m.email_ids = value
}

type EmailsBatchRemovePostRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetEmailIds() []string
	SetEmailIds(value []string)
}
