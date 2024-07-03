package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type FunnelIntegrationPostRequestBody struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The base URL of the integration, that will be used to call the integration related REST endpoints.
	base_url *string
	// The context of the integration, that will be used internally to identify the integration.
	context *string
	// The flag to enable or disable the integration.
	enabled *bool
	// The list of actions that will be installed with the integration.
	funnel_integration_actions []FunnelIntegrationPostRequestBody_funnel_integration_actionsable
	// The list of triggers that will be installed with the integration.
	funnel_integration_triggers []FunnelIntegrationPostRequestBody_funnel_integration_triggersable
	// The name of the integration, that will be looked for and deleted.
	integration_name *string
	// The name of the integration, that will be used internally to identify the integration.
	name *string
}

// NewFunnelIntegrationPostRequestBody instantiates a new FunnelIntegrationPostRequestBody and sets the default values.
func NewFunnelIntegrationPostRequestBody() *FunnelIntegrationPostRequestBody {
	m := &FunnelIntegrationPostRequestBody{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateFunnelIntegrationPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateFunnelIntegrationPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewFunnelIntegrationPostRequestBody(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *FunnelIntegrationPostRequestBody) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetBaseUrl gets the base_url property value. The base URL of the integration, that will be used to call the integration related REST endpoints.
// returns a *string when successful
func (m *FunnelIntegrationPostRequestBody) GetBaseUrl() *string {
	return m.base_url
}

// GetContext gets the context property value. The context of the integration, that will be used internally to identify the integration.
// returns a *string when successful
func (m *FunnelIntegrationPostRequestBody) GetContext() *string {
	return m.context
}

// GetEnabled gets the enabled property value. The flag to enable or disable the integration.
// returns a *bool when successful
func (m *FunnelIntegrationPostRequestBody) GetEnabled() *bool {
	return m.enabled
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *FunnelIntegrationPostRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
	res["funnel_integration_actions"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateFunnelIntegrationPostRequestBody_funnel_integration_actionsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]FunnelIntegrationPostRequestBody_funnel_integration_actionsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(FunnelIntegrationPostRequestBody_funnel_integration_actionsable)
				}
			}
			m.SetFunnelIntegrationActions(res)
		}
		return nil
	}
	res["funnel_integration_triggers"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateFunnelIntegrationPostRequestBody_funnel_integration_triggersFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]FunnelIntegrationPostRequestBody_funnel_integration_triggersable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(FunnelIntegrationPostRequestBody_funnel_integration_triggersable)
				}
			}
			m.SetFunnelIntegrationTriggers(res)
		}
		return nil
	}
	res["integration_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetIntegrationName(val)
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

// GetFunnelIntegrationActions gets the funnel_integration_actions property value. The list of actions that will be installed with the integration.
// returns a []FunnelIntegrationPostRequestBody_funnel_integration_actionsable when successful
func (m *FunnelIntegrationPostRequestBody) GetFunnelIntegrationActions() []FunnelIntegrationPostRequestBody_funnel_integration_actionsable {
	return m.funnel_integration_actions
}

// GetFunnelIntegrationTriggers gets the funnel_integration_triggers property value. The list of triggers that will be installed with the integration.
// returns a []FunnelIntegrationPostRequestBody_funnel_integration_triggersable when successful
func (m *FunnelIntegrationPostRequestBody) GetFunnelIntegrationTriggers() []FunnelIntegrationPostRequestBody_funnel_integration_triggersable {
	return m.funnel_integration_triggers
}

// GetIntegrationName gets the integration_name property value. The name of the integration, that will be looked for and deleted.
// returns a *string when successful
func (m *FunnelIntegrationPostRequestBody) GetIntegrationName() *string {
	return m.integration_name
}

// GetName gets the name property value. The name of the integration, that will be used internally to identify the integration.
// returns a *string when successful
func (m *FunnelIntegrationPostRequestBody) GetName() *string {
	return m.name
}

// Serialize serializes information the current object
func (m *FunnelIntegrationPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
	if m.GetFunnelIntegrationActions() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFunnelIntegrationActions()))
		for i, v := range m.GetFunnelIntegrationActions() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("funnel_integration_actions", cast)
		if err != nil {
			return err
		}
	}
	if m.GetFunnelIntegrationTriggers() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFunnelIntegrationTriggers()))
		for i, v := range m.GetFunnelIntegrationTriggers() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("funnel_integration_triggers", cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("integration_name", m.GetIntegrationName())
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
func (m *FunnelIntegrationPostRequestBody) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetBaseUrl sets the base_url property value. The base URL of the integration, that will be used to call the integration related REST endpoints.
func (m *FunnelIntegrationPostRequestBody) SetBaseUrl(value *string) {
	m.base_url = value
}

// SetContext sets the context property value. The context of the integration, that will be used internally to identify the integration.
func (m *FunnelIntegrationPostRequestBody) SetContext(value *string) {
	m.context = value
}

// SetEnabled sets the enabled property value. The flag to enable or disable the integration.
func (m *FunnelIntegrationPostRequestBody) SetEnabled(value *bool) {
	m.enabled = value
}

// SetFunnelIntegrationActions sets the funnel_integration_actions property value. The list of actions that will be installed with the integration.
func (m *FunnelIntegrationPostRequestBody) SetFunnelIntegrationActions(value []FunnelIntegrationPostRequestBody_funnel_integration_actionsable) {
	m.funnel_integration_actions = value
}

// SetFunnelIntegrationTriggers sets the funnel_integration_triggers property value. The list of triggers that will be installed with the integration.
func (m *FunnelIntegrationPostRequestBody) SetFunnelIntegrationTriggers(value []FunnelIntegrationPostRequestBody_funnel_integration_triggersable) {
	m.funnel_integration_triggers = value
}

// SetIntegrationName sets the integration_name property value. The name of the integration, that will be looked for and deleted.
func (m *FunnelIntegrationPostRequestBody) SetIntegrationName(value *string) {
	m.integration_name = value
}

// SetName sets the name property value. The name of the integration, that will be used internally to identify the integration.
func (m *FunnelIntegrationPostRequestBody) SetName(value *string) {
	m.name = value
}

type FunnelIntegrationPostRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetBaseUrl() *string
	GetContext() *string
	GetEnabled() *bool
	GetFunnelIntegrationActions() []FunnelIntegrationPostRequestBody_funnel_integration_actionsable
	GetFunnelIntegrationTriggers() []FunnelIntegrationPostRequestBody_funnel_integration_triggersable
	GetIntegrationName() *string
	GetName() *string
	SetBaseUrl(value *string)
	SetContext(value *string)
	SetEnabled(value *bool)
	SetFunnelIntegrationActions(value []FunnelIntegrationPostRequestBody_funnel_integration_actionsable)
	SetFunnelIntegrationTriggers(value []FunnelIntegrationPostRequestBody_funnel_integration_triggersable)
	SetIntegrationName(value *string)
	SetName(value *string)
}
