package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type HooksItemWithKeyGetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The eventKey property
	eventKey *string
	// The hookUrl property
	hookUrl *string
	// The key property
	key *string
}

// NewHooksItemWithKeyGetResponse instantiates a new HooksItemWithKeyGetResponse and sets the default values.
func NewHooksItemWithKeyGetResponse() *HooksItemWithKeyGetResponse {
	m := &HooksItemWithKeyGetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateHooksItemWithKeyGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateHooksItemWithKeyGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewHooksItemWithKeyGetResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *HooksItemWithKeyGetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetEventKey gets the eventKey property value. The eventKey property
// returns a *string when successful
func (m *HooksItemWithKeyGetResponse) GetEventKey() *string {
	return m.eventKey
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *HooksItemWithKeyGetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
	res["key"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetKey(val)
		}
		return nil
	}
	return res
}

// GetHookUrl gets the hookUrl property value. The hookUrl property
// returns a *string when successful
func (m *HooksItemWithKeyGetResponse) GetHookUrl() *string {
	return m.hookUrl
}

// GetKey gets the key property value. The key property
// returns a *string when successful
func (m *HooksItemWithKeyGetResponse) GetKey() *string {
	return m.key
}

// Serialize serializes information the current object
func (m *HooksItemWithKeyGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
		err := writer.WriteStringValue("key", m.GetKey())
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
func (m *HooksItemWithKeyGetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetEventKey sets the eventKey property value. The eventKey property
func (m *HooksItemWithKeyGetResponse) SetEventKey(value *string) {
	m.eventKey = value
}

// SetHookUrl sets the hookUrl property value. The hookUrl property
func (m *HooksItemWithKeyGetResponse) SetHookUrl(value *string) {
	m.hookUrl = value
}

// SetKey sets the key property value. The key property
func (m *HooksItemWithKeyGetResponse) SetKey(value *string) {
	m.key = value
}

type HooksItemWithKeyGetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetEventKey() *string
	GetHookUrl() *string
	GetKey() *string
	SetEventKey(value *string)
	SetHookUrl(value *string)
	SetKey(value *string)
}
