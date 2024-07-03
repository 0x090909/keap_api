package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ContactsModelGetResponse_custom_fields struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The default_value property
	default_value *string
	// The id property
	id *int64
	// The label property
	label *string
	// The options property
	options []ContactsModelGetResponse_custom_fields_optionsable
}

// NewContactsModelGetResponse_custom_fields instantiates a new ContactsModelGetResponse_custom_fields and sets the default values.
func NewContactsModelGetResponse_custom_fields() *ContactsModelGetResponse_custom_fields {
	m := &ContactsModelGetResponse_custom_fields{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateContactsModelGetResponse_custom_fieldsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsModelGetResponse_custom_fieldsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsModelGetResponse_custom_fields(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContactsModelGetResponse_custom_fields) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetDefaultValue gets the default_value property value. The default_value property
// returns a *string when successful
func (m *ContactsModelGetResponse_custom_fields) GetDefaultValue() *string {
	return m.default_value
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContactsModelGetResponse_custom_fields) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["default_value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDefaultValue(val)
		}
		return nil
	}
	res["id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
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
		val, err := n.GetCollectionOfObjectValues(CreateContactsModelGetResponse_custom_fields_optionsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ContactsModelGetResponse_custom_fields_optionsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ContactsModelGetResponse_custom_fields_optionsable)
				}
			}
			m.SetOptions(res)
		}
		return nil
	}
	return res
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *ContactsModelGetResponse_custom_fields) GetId() *int64 {
	return m.id
}

// GetLabel gets the label property value. The label property
// returns a *string when successful
func (m *ContactsModelGetResponse_custom_fields) GetLabel() *string {
	return m.label
}

// GetOptions gets the options property value. The options property
// returns a []ContactsModelGetResponse_custom_fields_optionsable when successful
func (m *ContactsModelGetResponse_custom_fields) GetOptions() []ContactsModelGetResponse_custom_fields_optionsable {
	return m.options
}

// Serialize serializes information the current object
func (m *ContactsModelGetResponse_custom_fields) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("default_value", m.GetDefaultValue())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("id", m.GetId())
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
func (m *ContactsModelGetResponse_custom_fields) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetDefaultValue sets the default_value property value. The default_value property
func (m *ContactsModelGetResponse_custom_fields) SetDefaultValue(value *string) {
	m.default_value = value
}

// SetId sets the id property value. The id property
func (m *ContactsModelGetResponse_custom_fields) SetId(value *int64) {
	m.id = value
}

// SetLabel sets the label property value. The label property
func (m *ContactsModelGetResponse_custom_fields) SetLabel(value *string) {
	m.label = value
}

// SetOptions sets the options property value. The options property
func (m *ContactsModelGetResponse_custom_fields) SetOptions(value []ContactsModelGetResponse_custom_fields_optionsable) {
	m.options = value
}

type ContactsModelGetResponse_custom_fieldsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetDefaultValue() *string
	GetId() *int64
	GetLabel() *string
	GetOptions() []ContactsModelGetResponse_custom_fields_optionsable
	SetDefaultValue(value *string)
	SetId(value *int64)
	SetLabel(value *string)
	SetOptions(value []ContactsModelGetResponse_custom_fields_optionsable)
}
