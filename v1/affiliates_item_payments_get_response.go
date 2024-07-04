package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AffiliatesItemPaymentsGetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The count property
	count *int32
	// The next property
	next *string
	// The payments property
	payments []AffiliatesItemPaymentsGetResponse_paymentsable
	// The previous property
	previous *string
}

// NewAffiliatesItemPaymentsGetResponse instantiates a new AffiliatesItemPaymentsGetResponse and sets the default values.
func NewAffiliatesItemPaymentsGetResponse() *AffiliatesItemPaymentsGetResponse {
	m := &AffiliatesItemPaymentsGetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateAffiliatesItemPaymentsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAffiliatesItemPaymentsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAffiliatesItemPaymentsGetResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AffiliatesItemPaymentsGetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCount gets the count property value. The count property
// returns a *int32 when successful
func (m *AffiliatesItemPaymentsGetResponse) GetCount() *int32 {
	return m.count
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AffiliatesItemPaymentsGetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["count"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCount(val)
		}
		return nil
	}
	res["next"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNext(val)
		}
		return nil
	}
	res["payments"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateAffiliatesItemPaymentsGetResponse_paymentsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]AffiliatesItemPaymentsGetResponse_paymentsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(AffiliatesItemPaymentsGetResponse_paymentsable)
				}
			}
			m.SetPayments(res)
		}
		return nil
	}
	res["previous"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPrevious(val)
		}
		return nil
	}
	return res
}

// GetNext gets the next property value. The next property
// returns a *string when successful
func (m *AffiliatesItemPaymentsGetResponse) GetNext() *string {
	return m.next
}

// GetPayments gets the payments property value. The payments property
// returns a []AffiliatesItemPaymentsGetResponse_paymentsable when successful
func (m *AffiliatesItemPaymentsGetResponse) GetPayments() []AffiliatesItemPaymentsGetResponse_paymentsable {
	return m.payments
}

// GetPrevious gets the previous property value. The previous property
// returns a *string when successful
func (m *AffiliatesItemPaymentsGetResponse) GetPrevious() *string {
	return m.previous
}

// Serialize serializes information the current object
func (m *AffiliatesItemPaymentsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt32Value("count", m.GetCount())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("next", m.GetNext())
		if err != nil {
			return err
		}
	}
	if m.GetPayments() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPayments()))
		for i, v := range m.GetPayments() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("payments", cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("previous", m.GetPrevious())
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
func (m *AffiliatesItemPaymentsGetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCount sets the count property value. The count property
func (m *AffiliatesItemPaymentsGetResponse) SetCount(value *int32) {
	m.count = value
}

// SetNext sets the next property value. The next property
func (m *AffiliatesItemPaymentsGetResponse) SetNext(value *string) {
	m.next = value
}

// SetPayments sets the payments property value. The payments property
func (m *AffiliatesItemPaymentsGetResponse) SetPayments(value []AffiliatesItemPaymentsGetResponse_paymentsable) {
	m.payments = value
}

// SetPrevious sets the previous property value. The previous property
func (m *AffiliatesItemPaymentsGetResponse) SetPrevious(value *string) {
	m.previous = value
}

type AffiliatesItemPaymentsGetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCount() *int32
	GetNext() *string
	GetPayments() []AffiliatesItemPaymentsGetResponse_paymentsable
	GetPrevious() *string
	SetCount(value *int32)
	SetNext(value *string)
	SetPayments(value []AffiliatesItemPaymentsGetResponse_paymentsable)
	SetPrevious(value *string)
}
