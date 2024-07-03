package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CampaignsGetResponse_campaigns_goals_historical_contact_counts struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The ThreeZero_days property
	threeZero_days *int32
	// The TwoFour_hours property
	twoFour_hours *int32
}

// NewCampaignsGetResponse_campaigns_goals_historical_contact_counts instantiates a new CampaignsGetResponse_campaigns_goals_historical_contact_counts and sets the default values.
func NewCampaignsGetResponse_campaigns_goals_historical_contact_counts() *CampaignsGetResponse_campaigns_goals_historical_contact_counts {
	m := &CampaignsGetResponse_campaigns_goals_historical_contact_counts{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateCampaignsGetResponse_campaigns_goals_historical_contact_countsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCampaignsGetResponse_campaigns_goals_historical_contact_countsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCampaignsGetResponse_campaigns_goals_historical_contact_counts(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CampaignsGetResponse_campaigns_goals_historical_contact_counts) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CampaignsGetResponse_campaigns_goals_historical_contact_counts) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
func (m *CampaignsGetResponse_campaigns_goals_historical_contact_counts) GetThreeZeroDays() *int32 {
	return m.threeZero_days
}

// GetTwoFourHours gets the 24_hours property value. The TwoFour_hours property
// returns a *int32 when successful
func (m *CampaignsGetResponse_campaigns_goals_historical_contact_counts) GetTwoFourHours() *int32 {
	return m.twoFour_hours
}

// Serialize serializes information the current object
func (m *CampaignsGetResponse_campaigns_goals_historical_contact_counts) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *CampaignsGetResponse_campaigns_goals_historical_contact_counts) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetThreeZeroDays sets the 30_days property value. The ThreeZero_days property
func (m *CampaignsGetResponse_campaigns_goals_historical_contact_counts) SetThreeZeroDays(value *int32) {
	m.threeZero_days = value
}

// SetTwoFourHours sets the 24_hours property value. The TwoFour_hours property
func (m *CampaignsGetResponse_campaigns_goals_historical_contact_counts) SetTwoFourHours(value *int32) {
	m.twoFour_hours = value
}

type CampaignsGetResponse_campaigns_goals_historical_contact_countsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetThreeZeroDays() *int32
	GetTwoFourHours() *int32
	SetThreeZeroDays(value *int32)
	SetTwoFourHours(value *int32)
}
