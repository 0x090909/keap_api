package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CampaignsItemWithCampaign_GetResponse_sequences_historical_contact_count struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The ThreeZero_days property
	threeZero_days *int32
	// The TwoFour_hours property
	twoFour_hours *int32
}

// NewCampaignsItemWithCampaign_GetResponse_sequences_historical_contact_count instantiates a new CampaignsItemWithCampaign_GetResponse_sequences_historical_contact_count and sets the default values.
func NewCampaignsItemWithCampaign_GetResponse_sequences_historical_contact_count() *CampaignsItemWithCampaign_GetResponse_sequences_historical_contact_count {
	m := &CampaignsItemWithCampaign_GetResponse_sequences_historical_contact_count{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateCampaignsItemWithCampaign_GetResponse_sequences_historical_contact_countFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCampaignsItemWithCampaign_GetResponse_sequences_historical_contact_countFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCampaignsItemWithCampaign_GetResponse_sequences_historical_contact_count(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CampaignsItemWithCampaign_GetResponse_sequences_historical_contact_count) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CampaignsItemWithCampaign_GetResponse_sequences_historical_contact_count) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["30_days"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetThreeZeroDays(val)
		}
		return nil
	}
	res["24_hours"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTwoFourHours(val)
		}
		return nil
	}
	return res
}

// GetThreeZeroDays gets the 30_days property value. The ThreeZero_days property
// returns a *int32 when successful
func (m *CampaignsItemWithCampaign_GetResponse_sequences_historical_contact_count) GetThreeZeroDays() *int32 {
	return m.threeZero_days
}

// GetTwoFourHours gets the 24_hours property value. The TwoFour_hours property
// returns a *int32 when successful
func (m *CampaignsItemWithCampaign_GetResponse_sequences_historical_contact_count) GetTwoFourHours() *int32 {
	return m.twoFour_hours
}

// Serialize serializes information the current object
func (m *CampaignsItemWithCampaign_GetResponse_sequences_historical_contact_count) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt32Value("30_days", m.GetThreeZeroDays())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("24_hours", m.GetTwoFourHours())
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
func (m *CampaignsItemWithCampaign_GetResponse_sequences_historical_contact_count) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetThreeZeroDays sets the 30_days property value. The ThreeZero_days property
func (m *CampaignsItemWithCampaign_GetResponse_sequences_historical_contact_count) SetThreeZeroDays(value *int32) {
	m.threeZero_days = value
}

// SetTwoFourHours sets the 24_hours property value. The TwoFour_hours property
func (m *CampaignsItemWithCampaign_GetResponse_sequences_historical_contact_count) SetTwoFourHours(value *int32) {
	m.twoFour_hours = value
}

type CampaignsItemWithCampaign_GetResponse_sequences_historical_contact_countable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetThreeZeroDays() *int32
	GetTwoFourHours() *int32
	SetThreeZeroDays(value *int32)
	SetTwoFourHours(value *int32)
}
