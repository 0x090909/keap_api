package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FunnelIntegrationPostRequestBody_funnel_integration_actions struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The base URL of the trigger, that will be used to call the integration related REST endpoints.
	base_url *string
	// The context of the trigger, that will be used internally to identify the integration.
	context *string
	// The flag to enable or disable the integration trigger.
	enabled *bool
	// The schema for this trigger that can be used in the builder context, and populated when a trigger is initiated.
	funnel_schema_json *string
	// The icon URL of the trigger, that will be used to display the icon of this specific integration trigger.
	icon_url *string
	// The name of the trigger, that will be used internally to identify and initiate the trigger.
	name *string
}

// NewFunnelIntegrationPostRequestBody_funnel_integration_actions instantiates a new FunnelIntegrationPostRequestBody_funnel_integration_actions and sets the default values.
func NewFunnelIntegrationPostRequestBody_funnel_integration_actions() *FunnelIntegrationPostRequestBody_funnel_integration_actions {
	m := &FunnelIntegrationPostRequestBody_funnel_integration_actions{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateFunnelIntegrationPostRequestBody_funnel_integration_actionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFunnelIntegrationPostRequestBody_funnel_integration_actionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewFunnelIntegrationPostRequestBody_funnel_integration_actions(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *FunnelIntegrationPostRequestBody_funnel_integration_actions) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetBaseUrl gets the base_url property value. The base URL of the trigger, that will be used to call the integration related REST endpoints.
// returns a *string when successful
func (m *FunnelIntegrationPostRequestBody_funnel_integration_actions) GetBaseUrl() *string {
	return m.base_url
}

// GetContext gets the context property value. The context of the trigger, that will be used internally to identify the integration.
// returns a *string when successful
func (m *FunnelIntegrationPostRequestBody_funnel_integration_actions) GetContext() *string {
	return m.context
}

// GetEnabled gets the enabled property value. The flag to enable or disable the integration trigger.
// returns a *bool when successful
func (m *FunnelIntegrationPostRequestBody_funnel_integration_actions) GetEnabled() *bool {
	return m.enabled
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FunnelIntegrationPostRequestBody_funnel_integration_actions) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["base_url"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetBaseUrl(val)
		}
		return nil
	}
	res["context"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetContext(val)
		}
		return nil
	}
	res["enabled"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetEnabled(val)
		}
		return nil
	}
	res["funnel_schema_json"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFunnelSchemaJson(val)
		}
		return nil
	}
	res["icon_url"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetIconUrl(val)
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

// GetFunnelSchemaJson gets the funnel_schema_json property value. The schema for this trigger that can be used in the builder context, and populated when a trigger is initiated.
// returns a *string when successful
func (m *FunnelIntegrationPostRequestBody_funnel_integration_actions) GetFunnelSchemaJson() *string {
	return m.funnel_schema_json
}

// GetIconUrl gets the icon_url property value. The icon URL of the trigger, that will be used to display the icon of this specific integration trigger.
// returns a *string when successful
func (m *FunnelIntegrationPostRequestBody_funnel_integration_actions) GetIconUrl() *string {
	return m.icon_url
}

// GetName gets the name property value. The name of the trigger, that will be used internally to identify and initiate the trigger.
// returns a *string when successful
func (m *FunnelIntegrationPostRequestBody_funnel_integration_actions) GetName() *string {
	return m.name
}

// Serialize serializes information the current object
func (m *FunnelIntegrationPostRequestBody_funnel_integration_actions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("base_url", m.GetBaseUrl())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("context", m.GetContext())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("enabled", m.GetEnabled())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("funnel_schema_json", m.GetFunnelSchemaJson())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("icon_url", m.GetIconUrl())
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
func (m *FunnelIntegrationPostRequestBody_funnel_integration_actions) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetBaseUrl sets the base_url property value. The base URL of the trigger, that will be used to call the integration related REST endpoints.
func (m *FunnelIntegrationPostRequestBody_funnel_integration_actions) SetBaseUrl(value *string) {
	m.base_url = value
}

// SetContext sets the context property value. The context of the trigger, that will be used internally to identify the integration.
func (m *FunnelIntegrationPostRequestBody_funnel_integration_actions) SetContext(value *string) {
	m.context = value
}

// SetEnabled sets the enabled property value. The flag to enable or disable the integration trigger.
func (m *FunnelIntegrationPostRequestBody_funnel_integration_actions) SetEnabled(value *bool) {
	m.enabled = value
}

// SetFunnelSchemaJson sets the funnel_schema_json property value. The schema for this trigger that can be used in the builder context, and populated when a trigger is initiated.
func (m *FunnelIntegrationPostRequestBody_funnel_integration_actions) SetFunnelSchemaJson(value *string) {
	m.funnel_schema_json = value
}

// SetIconUrl sets the icon_url property value. The icon URL of the trigger, that will be used to display the icon of this specific integration trigger.
func (m *FunnelIntegrationPostRequestBody_funnel_integration_actions) SetIconUrl(value *string) {
	m.icon_url = value
}

// SetName sets the name property value. The name of the trigger, that will be used internally to identify and initiate the trigger.
func (m *FunnelIntegrationPostRequestBody_funnel_integration_actions) SetName(value *string) {
	m.name = value
}

type FunnelIntegrationPostRequestBody_funnel_integration_actionsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetBaseUrl() *string
	GetContext() *string
	GetEnabled() *bool
	GetFunnelSchemaJson() *string
	GetIconUrl() *string
	GetName() *string
	SetBaseUrl(value *string)
	SetContext(value *string)
	SetEnabled(value *bool)
	SetFunnelSchemaJson(value *string)
	SetIconUrl(value *string)
	SetName(value *string)
}
