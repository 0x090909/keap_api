package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AffiliatesItemAffiliate_idPatchRequestBody struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The Affiliate code which have some validations.1. The code should not have white spaces2. The code should starts with letters3. The code minimum 4 characters length
	code *string
	// The contactId identifier , Must be a positive number
	contact_id *string
	// The Affiliate name will be derived from the Contact,when not explicitly provided
	name *string
}

// NewAffiliatesItemAffiliate_idPatchRequestBody instantiates a new AffiliatesItemAffiliate_idPatchRequestBody and sets the default values.
func NewAffiliatesItemAffiliate_idPatchRequestBody() *AffiliatesItemAffiliate_idPatchRequestBody {
	m := &AffiliatesItemAffiliate_idPatchRequestBody{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateAffiliatesItemAffiliate_idPatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAffiliatesItemAffiliate_idPatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAffiliatesItemAffiliate_idPatchRequestBody(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AffiliatesItemAffiliate_idPatchRequestBody) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCode gets the code property value. The Affiliate code which have some validations.1. The code should not have white spaces2. The code should starts with letters3. The code minimum 4 characters length
// returns a *string when successful
func (m *AffiliatesItemAffiliate_idPatchRequestBody) GetCode() *string {
	return m.code
}

// GetContactId gets the contact_id property value. The contactId identifier , Must be a positive number
// returns a *string when successful
func (m *AffiliatesItemAffiliate_idPatchRequestBody) GetContactId() *string {
	return m.contact_id
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AffiliatesItemAffiliate_idPatchRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["code"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCode(val)
		}
		return nil
	}
	res["contact_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetContactId(val)
		}
		return nil
	}
	res["name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetName(val)
		}
		return nil
	}
	return res
}

// GetName gets the name property value. The Affiliate name will be derived from the Contact,when not explicitly provided
// returns a *string when successful
func (m *AffiliatesItemAffiliate_idPatchRequestBody) GetName() *string {
	return m.name
}

// Serialize serializes information the current object
func (m *AffiliatesItemAffiliate_idPatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("code", m.GetCode())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("contact_id", m.GetContactId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("name", m.GetName())
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
func (m *AffiliatesItemAffiliate_idPatchRequestBody) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCode sets the code property value. The Affiliate code which have some validations.1. The code should not have white spaces2. The code should starts with letters3. The code minimum 4 characters length
func (m *AffiliatesItemAffiliate_idPatchRequestBody) SetCode(value *string) {
	m.code = value
}

// SetContactId sets the contact_id property value. The contactId identifier , Must be a positive number
func (m *AffiliatesItemAffiliate_idPatchRequestBody) SetContactId(value *string) {
	m.contact_id = value
}

// SetName sets the name property value. The Affiliate name will be derived from the Contact,when not explicitly provided
func (m *AffiliatesItemAffiliate_idPatchRequestBody) SetName(value *string) {
	m.name = value
}

type AffiliatesItemAffiliate_idPatchRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCode() *string
	GetContactId() *string
	GetName() *string
	SetCode(value *string)
	SetContactId(value *string)
	SetName(value *string)
}
