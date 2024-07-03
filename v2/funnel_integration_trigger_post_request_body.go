package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FunnelIntegrationTriggerPostRequestBody struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The contact_id property
	contact_id *int64
	// The funnel_integration_trigger_id property
	funnel_integration_trigger_id *string
	// The schema_data property
	schema_data *string
}

// NewFunnelIntegrationTriggerPostRequestBody instantiates a new FunnelIntegrationTriggerPostRequestBody and sets the default values.
func NewFunnelIntegrationTriggerPostRequestBody() *FunnelIntegrationTriggerPostRequestBody {
	m := &FunnelIntegrationTriggerPostRequestBody{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateFunnelIntegrationTriggerPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFunnelIntegrationTriggerPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewFunnelIntegrationTriggerPostRequestBody(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *FunnelIntegrationTriggerPostRequestBody) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetContactId gets the contact_id property value. The contact_id property
// returns a *int64 when successful
func (m *FunnelIntegrationTriggerPostRequestBody) GetContactId() *int64 {
	return m.contact_id
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FunnelIntegrationTriggerPostRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["contact_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetContactId(val)
		}
		return nil
	}
	res["funnel_integration_trigger_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFunnelIntegrationTriggerId(val)
		}
		return nil
	}
	res["schema_data"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSchemaData(val)
		}
		return nil
	}
	return res
}

// GetFunnelIntegrationTriggerId gets the funnel_integration_trigger_id property value. The funnel_integration_trigger_id property
// returns a *string when successful
func (m *FunnelIntegrationTriggerPostRequestBody) GetFunnelIntegrationTriggerId() *string {
	return m.funnel_integration_trigger_id
}

// GetSchemaData gets the schema_data property value. The schema_data property
// returns a *string when successful
func (m *FunnelIntegrationTriggerPostRequestBody) GetSchemaData() *string {
	return m.schema_data
}

// Serialize serializes information the current object
func (m *FunnelIntegrationTriggerPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt64Value("contact_id", m.GetContactId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("funnel_integration_trigger_id", m.GetFunnelIntegrationTriggerId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("schema_data", m.GetSchemaData())
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
func (m *FunnelIntegrationTriggerPostRequestBody) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetContactId sets the contact_id property value. The contact_id property
func (m *FunnelIntegrationTriggerPostRequestBody) SetContactId(value *int64) {
	m.contact_id = value
}

// SetFunnelIntegrationTriggerId sets the funnel_integration_trigger_id property value. The funnel_integration_trigger_id property
func (m *FunnelIntegrationTriggerPostRequestBody) SetFunnelIntegrationTriggerId(value *string) {
	m.funnel_integration_trigger_id = value
}

// SetSchemaData sets the schema_data property value. The schema_data property
func (m *FunnelIntegrationTriggerPostRequestBody) SetSchemaData(value *string) {
	m.schema_data = value
}

type FunnelIntegrationTriggerPostRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetContactId() *int64
	GetFunnelIntegrationTriggerId() *string
	GetSchemaData() *string
	SetContactId(value *int64)
	SetFunnelIntegrationTriggerId(value *string)
	SetSchemaData(value *string)
}
