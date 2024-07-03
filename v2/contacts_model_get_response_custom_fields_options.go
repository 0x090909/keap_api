package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	ia90dd1409a2ba05682b4d7b08237e40a27cab9440e9a81b916d31aca6dbec751 "keapapi/models"
)

type ContactsModelGetResponse_custom_fields_options struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The id property
	id *string
	// The label property
	label *string
	// The options property
	options []ia90dd1409a2ba05682b4d7b08237e40a27cab9440e9a81b916d31aca6dbec751.CustomFieldOptionable
}

// NewContactsModelGetResponse_custom_fields_options instantiates a new ContactsModelGetResponse_custom_fields_options and sets the default values.
func NewContactsModelGetResponse_custom_fields_options() *ContactsModelGetResponse_custom_fields_options {
	m := &ContactsModelGetResponse_custom_fields_options{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateContactsModelGetResponse_custom_fields_optionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsModelGetResponse_custom_fields_optionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsModelGetResponse_custom_fields_options(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContactsModelGetResponse_custom_fields_options) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContactsModelGetResponse_custom_fields_options) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetId(val)
		}
		return nil
	}
	res["label"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLabel(val)
		}
		return nil
	}
	res["options"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(ia90dd1409a2ba05682b4d7b08237e40a27cab9440e9a81b916d31aca6dbec751.CreateCustomFieldOptionFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ia90dd1409a2ba05682b4d7b08237e40a27cab9440e9a81b916d31aca6dbec751.CustomFieldOptionable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ia90dd1409a2ba05682b4d7b08237e40a27cab9440e9a81b916d31aca6dbec751.CustomFieldOptionable)
				}
			}
			m.SetOptions(res)
		}
		return nil
	}
	return res
}

// GetId gets the id property value. The id property
// returns a *string when successful
func (m *ContactsModelGetResponse_custom_fields_options) GetId() *string {
	return m.id
}

// GetLabel gets the label property value. The label property
// returns a *string when successful
func (m *ContactsModelGetResponse_custom_fields_options) GetLabel() *string {
	return m.label
}

// GetOptions gets the options property value. The options property
// returns a []CustomFieldOptionable when successful
func (m *ContactsModelGetResponse_custom_fields_options) GetOptions() []ia90dd1409a2ba05682b4d7b08237e40a27cab9440e9a81b916d31aca6dbec751.CustomFieldOptionable {
	return m.options
}

// Serialize serializes information the current object
func (m *ContactsModelGetResponse_custom_fields_options) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("id", m.GetId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("label", m.GetLabel())
		if err != nil {
			return err
		}
	}
	if m.GetOptions() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOptions()))
		for i, v := range m.GetOptions() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("options", cast)
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
func (m *ContactsModelGetResponse_custom_fields_options) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetId sets the id property value. The id property
func (m *ContactsModelGetResponse_custom_fields_options) SetId(value *string) {
	m.id = value
}

// SetLabel sets the label property value. The label property
func (m *ContactsModelGetResponse_custom_fields_options) SetLabel(value *string) {
	m.label = value
}

// SetOptions sets the options property value. The options property
func (m *ContactsModelGetResponse_custom_fields_options) SetOptions(value []ia90dd1409a2ba05682b4d7b08237e40a27cab9440e9a81b916d31aca6dbec751.CustomFieldOptionable) {
	m.options = value
}

type ContactsModelGetResponse_custom_fields_optionsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetId() *string
	GetLabel() *string
	GetOptions() []ia90dd1409a2ba05682b4d7b08237e40a27cab9440e9a81b916d31aca6dbec751.CustomFieldOptionable
	SetId(value *string)
	SetLabel(value *string)
	SetOptions(value []ia90dd1409a2ba05682b4d7b08237e40a27cab9440e9a81b916d31aca6dbec751.CustomFieldOptionable)
}
