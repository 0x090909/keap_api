package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AffiliatesItemClawbacksGetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The clawbacks property
	clawbacks []AffiliatesItemClawbacksGetResponse_clawbacksable
	// The count property
	count *int32
	// The next property
	next *string
	// The previous property
	previous *string
}

// NewAffiliatesItemClawbacksGetResponse instantiates a new AffiliatesItemClawbacksGetResponse and sets the default values.
func NewAffiliatesItemClawbacksGetResponse() *AffiliatesItemClawbacksGetResponse {
	m := &AffiliatesItemClawbacksGetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateAffiliatesItemClawbacksGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAffiliatesItemClawbacksGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAffiliatesItemClawbacksGetResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AffiliatesItemClawbacksGetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetClawbacks gets the clawbacks property value. The clawbacks property
// returns a []AffiliatesItemClawbacksGetResponse_clawbacksable when successful
func (m *AffiliatesItemClawbacksGetResponse) GetClawbacks() []AffiliatesItemClawbacksGetResponse_clawbacksable {
	return m.clawbacks
}

// GetCount gets the count property value. The count property
// returns a *int32 when successful
func (m *AffiliatesItemClawbacksGetResponse) GetCount() *int32 {
	return m.count
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AffiliatesItemClawbacksGetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["clawbacks"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateAffiliatesItemClawbacksGetResponse_clawbacksFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]AffiliatesItemClawbacksGetResponse_clawbacksable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(AffiliatesItemClawbacksGetResponse_clawbacksable)
				}
			}
			m.SetClawbacks(res)
		}
		return nil
	}
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
func (m *AffiliatesItemClawbacksGetResponse) GetNext() *string {
	return m.next
}

// GetPrevious gets the previous property value. The previous property
// returns a *string when successful
func (m *AffiliatesItemClawbacksGetResponse) GetPrevious() *string {
	return m.previous
}

// Serialize serializes information the current object
func (m *AffiliatesItemClawbacksGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetClawbacks() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetClawbacks()))
		for i, v := range m.GetClawbacks() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("clawbacks", cast)
		if err != nil {
			return err
		}
	}
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
func (m *AffiliatesItemClawbacksGetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetClawbacks sets the clawbacks property value. The clawbacks property
func (m *AffiliatesItemClawbacksGetResponse) SetClawbacks(value []AffiliatesItemClawbacksGetResponse_clawbacksable) {
	m.clawbacks = value
}

// SetCount sets the count property value. The count property
func (m *AffiliatesItemClawbacksGetResponse) SetCount(value *int32) {
	m.count = value
}

// SetNext sets the next property value. The next property
func (m *AffiliatesItemClawbacksGetResponse) SetNext(value *string) {
	m.next = value
}

// SetPrevious sets the previous property value. The previous property
func (m *AffiliatesItemClawbacksGetResponse) SetPrevious(value *string) {
	m.previous = value
}

type AffiliatesItemClawbacksGetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetClawbacks() []AffiliatesItemClawbacksGetResponse_clawbacksable
	GetCount() *int32
	GetNext() *string
	GetPrevious() *string
	SetClawbacks(value []AffiliatesItemClawbacksGetResponse_clawbacksable)
	SetCount(value *int32)
	SetNext(value *string)
	SetPrevious(value *string)
}
