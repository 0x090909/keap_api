package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ContactsModelGetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The custom_fields property
	custom_fields []ContactsModelGetResponse_custom_fieldsable
	// These fields are not transmitted by default on this model, but can be requested by specifying them in a comma-separated list in the optional_properties query parameter.
	optional_properties []string
}

// NewContactsModelGetResponse instantiates a new ContactsModelGetResponse and sets the default values.
func NewContactsModelGetResponse() *ContactsModelGetResponse {
	m := &ContactsModelGetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateContactsModelGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsModelGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsModelGetResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContactsModelGetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCustomFields gets the custom_fields property value. The custom_fields property
// returns a []ContactsModelGetResponse_custom_fieldsable when successful
func (m *ContactsModelGetResponse) GetCustomFields() []ContactsModelGetResponse_custom_fieldsable {
	return m.custom_fields
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContactsModelGetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["custom_fields"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateContactsModelGetResponse_custom_fieldsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ContactsModelGetResponse_custom_fieldsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ContactsModelGetResponse_custom_fieldsable)
				}
			}
			m.SetCustomFields(res)
		}
		return nil
	}
	res["optional_properties"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
			m.SetOptionalProperties(res)
		}
		return nil
	}
	return res
}

// GetOptionalProperties gets the optional_properties property value. These fields are not transmitted by default on this model, but can be requested by specifying them in a comma-separated list in the optional_properties query parameter.
// returns a []string when successful
func (m *ContactsModelGetResponse) GetOptionalProperties() []string {
	return m.optional_properties
}

// Serialize serializes information the current object
func (m *ContactsModelGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetCustomFields() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomFields()))
		for i, v := range m.GetCustomFields() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("custom_fields", cast)
		if err != nil {
			return err
		}
	}
	if m.GetOptionalProperties() != nil {
		err := writer.WriteCollectionOfStringValues("optional_properties", m.GetOptionalProperties())
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
func (m *ContactsModelGetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCustomFields sets the custom_fields property value. The custom_fields property
func (m *ContactsModelGetResponse) SetCustomFields(value []ContactsModelGetResponse_custom_fieldsable) {
	m.custom_fields = value
}

// SetOptionalProperties sets the optional_properties property value. These fields are not transmitted by default on this model, but can be requested by specifying them in a comma-separated list in the optional_properties query parameter.
func (m *ContactsModelGetResponse) SetOptionalProperties(value []string) {
	m.optional_properties = value
}

type ContactsModelGetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCustomFields() []ContactsModelGetResponse_custom_fieldsable
	GetOptionalProperties() []string
	SetCustomFields(value []ContactsModelGetResponse_custom_fieldsable)
	SetOptionalProperties(value []string)
}
