package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AutomationsGetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The automation_count property
	automation_count *int32
	// The automations property
	automations []AutomationsGetResponse_automationsable
	// The next_page_token property
	next_page_token *string
}

// NewAutomationsGetResponse instantiates a new AutomationsGetResponse and sets the default values.
func NewAutomationsGetResponse() *AutomationsGetResponse {
	m := &AutomationsGetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateAutomationsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAutomationsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAutomationsGetResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AutomationsGetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAutomationCount gets the automation_count property value. The automation_count property
// returns a *int32 when successful
func (m *AutomationsGetResponse) GetAutomationCount() *int32 {
	return m.automation_count
}

// GetAutomations gets the automations property value. The automations property
// returns a []AutomationsGetResponse_automationsable when successful
func (m *AutomationsGetResponse) GetAutomations() []AutomationsGetResponse_automationsable {
	return m.automations
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AutomationsGetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["automation_count"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAutomationCount(val)
		}
		return nil
	}
	res["automations"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateAutomationsGetResponse_automationsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]AutomationsGetResponse_automationsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(AutomationsGetResponse_automationsable)
				}
			}
			m.SetAutomations(res)
		}
		return nil
	}
	res["next_page_token"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNextPageToken(val)
		}
		return nil
	}
	return res
}

// GetNextPageToken gets the next_page_token property value. The next_page_token property
// returns a *string when successful
func (m *AutomationsGetResponse) GetNextPageToken() *string {
	return m.next_page_token
}

// Serialize serializes information the current object
func (m *AutomationsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetAutomations() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAutomations()))
		for i, v := range m.GetAutomations() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("automations", cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("automation_count", m.GetAutomationCount())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("next_page_token", m.GetNextPageToken())
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
func (m *AutomationsGetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAutomationCount sets the automation_count property value. The automation_count property
func (m *AutomationsGetResponse) SetAutomationCount(value *int32) {
	m.automation_count = value
}

// SetAutomations sets the automations property value. The automations property
func (m *AutomationsGetResponse) SetAutomations(value []AutomationsGetResponse_automationsable) {
	m.automations = value
}

// SetNextPageToken sets the next_page_token property value. The next_page_token property
func (m *AutomationsGetResponse) SetNextPageToken(value *string) {
	m.next_page_token = value
}

type AutomationsGetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAutomationCount() *int32
	GetAutomations() []AutomationsGetResponse_automationsable
	GetNextPageToken() *string
	SetAutomationCount(value *int32)
	SetAutomations(value []AutomationsGetResponse_automationsable)
	SetNextPageToken(value *string)
}
