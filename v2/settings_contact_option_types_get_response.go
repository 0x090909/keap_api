package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SettingsContactOptionTypesGetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The option_types property
	option_types []string
}

// NewSettingsContactOptionTypesGetResponse instantiates a new SettingsContactOptionTypesGetResponse and sets the default values.
func NewSettingsContactOptionTypesGetResponse() *SettingsContactOptionTypesGetResponse {
	m := &SettingsContactOptionTypesGetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateSettingsContactOptionTypesGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSettingsContactOptionTypesGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSettingsContactOptionTypesGetResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SettingsContactOptionTypesGetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SettingsContactOptionTypesGetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["option_types"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
			m.SetOptionTypes(res)
		}
		return nil
	}
	return res
}

// GetOptionTypes gets the option_types property value. The option_types property
// returns a []string when successful
func (m *SettingsContactOptionTypesGetResponse) GetOptionTypes() []string {
	return m.option_types
}

// Serialize serializes information the current object
func (m *SettingsContactOptionTypesGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetOptionTypes() != nil {
		err := writer.WriteCollectionOfStringValues("option_types", m.GetOptionTypes())
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
func (m *SettingsContactOptionTypesGetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetOptionTypes sets the option_types property value. The option_types property
func (m *SettingsContactOptionTypesGetResponse) SetOptionTypes(value []string) {
	m.option_types = value
}

type SettingsContactOptionTypesGetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetOptionTypes() []string
	SetOptionTypes(value []string)
}
