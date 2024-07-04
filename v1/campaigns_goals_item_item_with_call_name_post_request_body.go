package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CampaignsGoalsItemItemWithCallNamePostRequestBody struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The contact_id property
	contact_id *int64
}

// NewCampaignsGoalsItemItemWithCallNamePostRequestBody instantiates a new CampaignsGoalsItemItemWithCallNamePostRequestBody and sets the default values.
func NewCampaignsGoalsItemItemWithCallNamePostRequestBody() *CampaignsGoalsItemItemWithCallNamePostRequestBody {
	m := &CampaignsGoalsItemItemWithCallNamePostRequestBody{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateCampaignsGoalsItemItemWithCallNamePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCampaignsGoalsItemItemWithCallNamePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCampaignsGoalsItemItemWithCallNamePostRequestBody(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CampaignsGoalsItemItemWithCallNamePostRequestBody) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetContactId gets the contact_id property value. The contact_id property
// returns a *int64 when successful
func (m *CampaignsGoalsItemItemWithCallNamePostRequestBody) GetContactId() *int64 {
	return m.contact_id
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CampaignsGoalsItemItemWithCallNamePostRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
	return res
}

// Serialize serializes information the current object
func (m *CampaignsGoalsItemItemWithCallNamePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt64Value("contact_id", m.GetContactId())
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
func (m *CampaignsGoalsItemItemWithCallNamePostRequestBody) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetContactId sets the contact_id property value. The contact_id property
func (m *CampaignsGoalsItemItemWithCallNamePostRequestBody) SetContactId(value *int64) {
	m.contact_id = value
}

type CampaignsGoalsItemItemWithCallNamePostRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetContactId() *int64
	SetContactId(value *int64)
}
