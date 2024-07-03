package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AutomationCategoryPostResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The automation_count property
	automation_count *int64
	// The id property
	id *int64
	// The name property
	name *string
}

// NewAutomationCategoryPostResponse instantiates a new AutomationCategoryPostResponse and sets the default values.
func NewAutomationCategoryPostResponse() *AutomationCategoryPostResponse {
	m := &AutomationCategoryPostResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateAutomationCategoryPostResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAutomationCategoryPostResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAutomationCategoryPostResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AutomationCategoryPostResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAutomationCount gets the automation_count property value. The automation_count property
// returns a *int64 when successful
func (m *AutomationCategoryPostResponse) GetAutomationCount() *int64 {
	return m.automation_count
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AutomationCategoryPostResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["automation_count"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAutomationCount(val)
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

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *AutomationCategoryPostResponse) GetId() *int64 {
	return m.id
}

// GetName gets the name property value. The name property
// returns a *string when successful
func (m *AutomationCategoryPostResponse) GetName() *string {
	return m.name
}

// Serialize serializes information the current object
func (m *AutomationCategoryPostResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt64Value("automation_count", m.GetAutomationCount())
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
func (m *AutomationCategoryPostResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAutomationCount sets the automation_count property value. The automation_count property
func (m *AutomationCategoryPostResponse) SetAutomationCount(value *int64) {
	m.automation_count = value
}

// SetId sets the id property value. The id property
func (m *AutomationCategoryPostResponse) SetId(value *int64) {
	m.id = value
}

// SetName sets the name property value. The name property
func (m *AutomationCategoryPostResponse) SetName(value *string) {
	m.name = value
}

type AutomationCategoryPostResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAutomationCount() *int64
	GetId() *int64
	GetName() *string
	SetAutomationCount(value *int64)
	SetId(value *int64)
	SetName(value *string)
}
