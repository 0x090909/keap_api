package models

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CreateRestCustomFieldOption struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The label property
	label *string
	// The options property
	options []CreateRestCustomFieldOptionable
}

// NewCreateRestCustomFieldOption instantiates a new CreateRestCustomFieldOption and sets the default values.
func NewCreateRestCustomFieldOption() *CreateRestCustomFieldOption {
	m := &CreateRestCustomFieldOption{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateCreateRestCustomFieldOptionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCreateRestCustomFieldOptionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCreateRestCustomFieldOption(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CreateRestCustomFieldOption) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CreateRestCustomFieldOption) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
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
		val, err := n.GetCollectionOfObjectValues(CreateCreateRestCustomFieldOptionFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]CreateRestCustomFieldOptionable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(CreateRestCustomFieldOptionable)
				}
			}
			m.SetOptions(res)
		}
		return nil
	}
	return res
}

// GetLabel gets the label property value. The label property
// returns a *string when successful
func (m *CreateRestCustomFieldOption) GetLabel() *string {
	return m.label
}

// GetOptions gets the options property value. The options property
// returns a []CreateRestCustomFieldOptionable when successful
func (m *CreateRestCustomFieldOption) GetOptions() []CreateRestCustomFieldOptionable {
	return m.options
}

// Serialize serializes information the current object
func (m *CreateRestCustomFieldOption) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *CreateRestCustomFieldOption) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetLabel sets the label property value. The label property
func (m *CreateRestCustomFieldOption) SetLabel(value *string) {
	m.label = value
}

// SetOptions sets the options property value. The options property
func (m *CreateRestCustomFieldOption) SetOptions(value []CreateRestCustomFieldOptionable) {
	m.options = value
}

type CreateRestCustomFieldOptionable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetLabel() *string
	GetOptions() []CreateRestCustomFieldOptionable
	SetLabel(value *string)
	SetOptions(value []CreateRestCustomFieldOptionable)
}
