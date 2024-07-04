package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type HooksPostRequestBody struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The eventKey property
	eventKey *string
	// The hookUrl property
	hookUrl *string
}

// NewHooksPostRequestBody instantiates a new HooksPostRequestBody and sets the default values.
func NewHooksPostRequestBody() *HooksPostRequestBody {
	m := &HooksPostRequestBody{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateHooksPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateHooksPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewHooksPostRequestBody(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *HooksPostRequestBody) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetEventKey gets the eventKey property value. The eventKey property
// returns a *string when successful
func (m *HooksPostRequestBody) GetEventKey() *string {
	return m.eventKey
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *HooksPostRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["eventKey"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetEventKey(val)
		}
		return nil
	}
	res["hookUrl"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetHookUrl(val)
		}
		return nil
	}
	return res
}

// GetHookUrl gets the hookUrl property value. The hookUrl property
// returns a *string when successful
func (m *HooksPostRequestBody) GetHookUrl() *string {
	return m.hookUrl
}

// Serialize serializes information the current object
func (m *HooksPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("eventKey", m.GetEventKey())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("hookUrl", m.GetHookUrl())
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
func (m *HooksPostRequestBody) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetEventKey sets the eventKey property value. The eventKey property
func (m *HooksPostRequestBody) SetEventKey(value *string) {
	m.eventKey = value
}

// SetHookUrl sets the hookUrl property value. The hookUrl property
func (m *HooksPostRequestBody) SetHookUrl(value *string) {
	m.hookUrl = value
}

type HooksPostRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetEventKey() *string
	GetHookUrl() *string
	SetEventKey(value *string)
	SetHookUrl(value *string)
}
