package v1

import (
	ic811111da47f7c5d6f6d121f0e78585ed7060c3db2cbfa4a5d929d3a72d517ff "github.com/0x090909/keap_api/models"
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SubscriptionsModelGetResponse_custom_fields_options struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The id property
	id *string
	// The label property
	label *string
	// The options property
	options []ic811111da47f7c5d6f6d121f0e78585ed7060c3db2cbfa4a5d929d3a72d517ff.CustomFieldOptionable
}

// NewSubscriptionsModelGetResponse_custom_fields_options instantiates a new SubscriptionsModelGetResponse_custom_fields_options and sets the default values.
func NewSubscriptionsModelGetResponse_custom_fields_options() *SubscriptionsModelGetResponse_custom_fields_options {
	m := &SubscriptionsModelGetResponse_custom_fields_options{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateSubscriptionsModelGetResponse_custom_fields_optionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSubscriptionsModelGetResponse_custom_fields_optionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSubscriptionsModelGetResponse_custom_fields_options(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SubscriptionsModelGetResponse_custom_fields_options) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SubscriptionsModelGetResponse_custom_fields_options) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
		val, err := n.GetCollectionOfObjectValues(ic811111da47f7c5d6f6d121f0e78585ed7060c3db2cbfa4a5d929d3a72d517ff.CreateCustomFieldOptionFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ic811111da47f7c5d6f6d121f0e78585ed7060c3db2cbfa4a5d929d3a72d517ff.CustomFieldOptionable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ic811111da47f7c5d6f6d121f0e78585ed7060c3db2cbfa4a5d929d3a72d517ff.CustomFieldOptionable)
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
func (m *SubscriptionsModelGetResponse_custom_fields_options) GetId() *string {
	return m.id
}

// GetLabel gets the label property value. The label property
// returns a *string when successful
func (m *SubscriptionsModelGetResponse_custom_fields_options) GetLabel() *string {
	return m.label
}

// GetOptions gets the options property value. The options property
// returns a []CustomFieldOptionable when successful
func (m *SubscriptionsModelGetResponse_custom_fields_options) GetOptions() []ic811111da47f7c5d6f6d121f0e78585ed7060c3db2cbfa4a5d929d3a72d517ff.CustomFieldOptionable {
	return m.options
}

// Serialize serializes information the current object
func (m *SubscriptionsModelGetResponse_custom_fields_options) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *SubscriptionsModelGetResponse_custom_fields_options) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetId sets the id property value. The id property
func (m *SubscriptionsModelGetResponse_custom_fields_options) SetId(value *string) {
	m.id = value
}

// SetLabel sets the label property value. The label property
func (m *SubscriptionsModelGetResponse_custom_fields_options) SetLabel(value *string) {
	m.label = value
}

// SetOptions sets the options property value. The options property
func (m *SubscriptionsModelGetResponse_custom_fields_options) SetOptions(value []ic811111da47f7c5d6f6d121f0e78585ed7060c3db2cbfa4a5d929d3a72d517ff.CustomFieldOptionable) {
	m.options = value
}

type SubscriptionsModelGetResponse_custom_fields_optionsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetId() *string
	GetLabel() *string
	GetOptions() []ic811111da47f7c5d6f6d121f0e78585ed7060c3db2cbfa4a5d929d3a72d517ff.CustomFieldOptionable
	SetId(value *string)
	SetLabel(value *string)
	SetOptions(value []ic811111da47f7c5d6f6d121f0e78585ed7060c3db2cbfa4a5d929d3a72d517ff.CustomFieldOptionable)
}
