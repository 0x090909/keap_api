package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AffiliatesSummariesGetResponse_summaries struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The affiliate_id property
	affiliate_id *int64
	// The amount_earned property
	amount_earned *float64
	// The balance property
	balance *float64
	// The clawbacks property
	clawbacks *float64
}

// NewAffiliatesSummariesGetResponse_summaries instantiates a new AffiliatesSummariesGetResponse_summaries and sets the default values.
func NewAffiliatesSummariesGetResponse_summaries() *AffiliatesSummariesGetResponse_summaries {
	m := &AffiliatesSummariesGetResponse_summaries{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateAffiliatesSummariesGetResponse_summariesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAffiliatesSummariesGetResponse_summariesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAffiliatesSummariesGetResponse_summaries(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AffiliatesSummariesGetResponse_summaries) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAffiliateId gets the affiliate_id property value. The affiliate_id property
// returns a *int64 when successful
func (m *AffiliatesSummariesGetResponse_summaries) GetAffiliateId() *int64 {
	return m.affiliate_id
}

// GetAmountEarned gets the amount_earned property value. The amount_earned property
// returns a *float64 when successful
func (m *AffiliatesSummariesGetResponse_summaries) GetAmountEarned() *float64 {
	return m.amount_earned
}

// GetBalance gets the balance property value. The balance property
// returns a *float64 when successful
func (m *AffiliatesSummariesGetResponse_summaries) GetBalance() *float64 {
	return m.balance
}

// GetClawbacks gets the clawbacks property value. The clawbacks property
// returns a *float64 when successful
func (m *AffiliatesSummariesGetResponse_summaries) GetClawbacks() *float64 {
	return m.clawbacks
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AffiliatesSummariesGetResponse_summaries) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["affiliate_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAffiliateId(val)
		}
		return nil
	}
	res["amount_earned"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetFloat64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAmountEarned(val)
		}
		return nil
	}
	res["balance"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetFloat64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetBalance(val)
		}
		return nil
	}
	res["clawbacks"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetFloat64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetClawbacks(val)
		}
		return nil
	}
	return res
}

// Serialize serializes information the current object
func (m *AffiliatesSummariesGetResponse_summaries) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt64Value("affiliate_id", m.GetAffiliateId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteFloat64Value("amount_earned", m.GetAmountEarned())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteFloat64Value("balance", m.GetBalance())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteFloat64Value("clawbacks", m.GetClawbacks())
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
func (m *AffiliatesSummariesGetResponse_summaries) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAffiliateId sets the affiliate_id property value. The affiliate_id property
func (m *AffiliatesSummariesGetResponse_summaries) SetAffiliateId(value *int64) {
	m.affiliate_id = value
}

// SetAmountEarned sets the amount_earned property value. The amount_earned property
func (m *AffiliatesSummariesGetResponse_summaries) SetAmountEarned(value *float64) {
	m.amount_earned = value
}

// SetBalance sets the balance property value. The balance property
func (m *AffiliatesSummariesGetResponse_summaries) SetBalance(value *float64) {
	m.balance = value
}

// SetClawbacks sets the clawbacks property value. The clawbacks property
func (m *AffiliatesSummariesGetResponse_summaries) SetClawbacks(value *float64) {
	m.clawbacks = value
}

type AffiliatesSummariesGetResponse_summariesable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAffiliateId() *int64
	GetAmountEarned() *float64
	GetBalance() *float64
	GetClawbacks() *float64
	SetAffiliateId(value *int64)
	SetAmountEarned(value *float64)
	SetBalance(value *float64)
	SetClawbacks(value *float64)
}
