package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ContactsLinksTypesPostRequestBody struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The max_links property
	max_links *int64
	// The name property
	name *string
}

// NewContactsLinksTypesPostRequestBody instantiates a new ContactsLinksTypesPostRequestBody and sets the default values.
func NewContactsLinksTypesPostRequestBody() *ContactsLinksTypesPostRequestBody {
	m := &ContactsLinksTypesPostRequestBody{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateContactsLinksTypesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsLinksTypesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsLinksTypesPostRequestBody(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContactsLinksTypesPostRequestBody) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContactsLinksTypesPostRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["max_links"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetMaxLinks(val)
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

// GetMaxLinks gets the max_links property value. The max_links property
// returns a *int64 when successful
func (m *ContactsLinksTypesPostRequestBody) GetMaxLinks() *int64 {
	return m.max_links
}

// GetName gets the name property value. The name property
// returns a *string when successful
func (m *ContactsLinksTypesPostRequestBody) GetName() *string {
	return m.name
}

// Serialize serializes information the current object
func (m *ContactsLinksTypesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt64Value("max_links", m.GetMaxLinks())
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
func (m *ContactsLinksTypesPostRequestBody) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetMaxLinks sets the max_links property value. The max_links property
func (m *ContactsLinksTypesPostRequestBody) SetMaxLinks(value *int64) {
	m.max_links = value
}

// SetName sets the name property value. The name property
func (m *ContactsLinksTypesPostRequestBody) SetName(value *string) {
	m.name = value
}

type ContactsLinksTypesPostRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetMaxLinks() *int64
	GetName() *string
	SetMaxLinks(value *int64)
	SetName(value *string)
}
