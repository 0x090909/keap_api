package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LocalesDefaultOptionsGetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The contact_types property
	contact_types []string
	// The fax_types property
	fax_types []string
	// The phone_types property
	phone_types []string
	// The suffix_types property
	suffix_types []string
	// The title_types property
	title_types []string
}

// NewLocalesDefaultOptionsGetResponse instantiates a new LocalesDefaultOptionsGetResponse and sets the default values.
func NewLocalesDefaultOptionsGetResponse() *LocalesDefaultOptionsGetResponse {
	m := &LocalesDefaultOptionsGetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateLocalesDefaultOptionsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLocalesDefaultOptionsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewLocalesDefaultOptionsGetResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *LocalesDefaultOptionsGetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetContactTypes gets the contact_types property value. The contact_types property
// returns a []string when successful
func (m *LocalesDefaultOptionsGetResponse) GetContactTypes() []string {
	return m.contact_types
}

// GetFaxTypes gets the fax_types property value. The fax_types property
// returns a []string when successful
func (m *LocalesDefaultOptionsGetResponse) GetFaxTypes() []string {
	return m.fax_types
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LocalesDefaultOptionsGetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["contact_types"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
			m.SetContactTypes(res)
		}
		return nil
	}
	res["fax_types"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
			m.SetFaxTypes(res)
		}
		return nil
	}
	res["phone_types"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
			m.SetPhoneTypes(res)
		}
		return nil
	}
	res["suffix_types"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
			m.SetSuffixTypes(res)
		}
		return nil
	}
	res["title_types"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
			m.SetTitleTypes(res)
		}
		return nil
	}
	return res
}

// GetPhoneTypes gets the phone_types property value. The phone_types property
// returns a []string when successful
func (m *LocalesDefaultOptionsGetResponse) GetPhoneTypes() []string {
	return m.phone_types
}

// GetSuffixTypes gets the suffix_types property value. The suffix_types property
// returns a []string when successful
func (m *LocalesDefaultOptionsGetResponse) GetSuffixTypes() []string {
	return m.suffix_types
}

// GetTitleTypes gets the title_types property value. The title_types property
// returns a []string when successful
func (m *LocalesDefaultOptionsGetResponse) GetTitleTypes() []string {
	return m.title_types
}

// Serialize serializes information the current object
func (m *LocalesDefaultOptionsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetContactTypes() != nil {
		err := writer.WriteCollectionOfStringValues("contact_types", m.GetContactTypes())
		if err != nil {
			return err
		}
	}
	if m.GetFaxTypes() != nil {
		err := writer.WriteCollectionOfStringValues("fax_types", m.GetFaxTypes())
		if err != nil {
			return err
		}
	}
	if m.GetPhoneTypes() != nil {
		err := writer.WriteCollectionOfStringValues("phone_types", m.GetPhoneTypes())
		if err != nil {
			return err
		}
	}
	if m.GetSuffixTypes() != nil {
		err := writer.WriteCollectionOfStringValues("suffix_types", m.GetSuffixTypes())
		if err != nil {
			return err
		}
	}
	if m.GetTitleTypes() != nil {
		err := writer.WriteCollectionOfStringValues("title_types", m.GetTitleTypes())
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
func (m *LocalesDefaultOptionsGetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetContactTypes sets the contact_types property value. The contact_types property
func (m *LocalesDefaultOptionsGetResponse) SetContactTypes(value []string) {
	m.contact_types = value
}

// SetFaxTypes sets the fax_types property value. The fax_types property
func (m *LocalesDefaultOptionsGetResponse) SetFaxTypes(value []string) {
	m.fax_types = value
}

// SetPhoneTypes sets the phone_types property value. The phone_types property
func (m *LocalesDefaultOptionsGetResponse) SetPhoneTypes(value []string) {
	m.phone_types = value
}

// SetSuffixTypes sets the suffix_types property value. The suffix_types property
func (m *LocalesDefaultOptionsGetResponse) SetSuffixTypes(value []string) {
	m.suffix_types = value
}

// SetTitleTypes sets the title_types property value. The title_types property
func (m *LocalesDefaultOptionsGetResponse) SetTitleTypes(value []string) {
	m.title_types = value
}

type LocalesDefaultOptionsGetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetContactTypes() []string
	GetFaxTypes() []string
	GetPhoneTypes() []string
	GetSuffixTypes() []string
	GetTitleTypes() []string
	SetContactTypes(value []string)
	SetFaxTypes(value []string)
	SetPhoneTypes(value []string)
	SetSuffixTypes(value []string)
	SetTitleTypes(value []string)
}
