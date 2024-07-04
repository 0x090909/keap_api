package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AffiliatesGetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The affiliates property
	affiliates []AffiliatesGetResponse_affiliatesable
	// The count property
	count *int32
	// The next property
	next *string
	// The previous property
	previous *string
}

// NewAffiliatesGetResponse instantiates a new AffiliatesGetResponse and sets the default values.
func NewAffiliatesGetResponse() *AffiliatesGetResponse {
	m := &AffiliatesGetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateAffiliatesGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAffiliatesGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAffiliatesGetResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AffiliatesGetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAffiliates gets the affiliates property value. The affiliates property
// returns a []AffiliatesGetResponse_affiliatesable when successful
func (m *AffiliatesGetResponse) GetAffiliates() []AffiliatesGetResponse_affiliatesable {
	return m.affiliates
}

// GetCount gets the count property value. The count property
// returns a *int32 when successful
func (m *AffiliatesGetResponse) GetCount() *int32 {
	return m.count
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AffiliatesGetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["affiliates"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateAffiliatesGetResponse_affiliatesFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]AffiliatesGetResponse_affiliatesable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(AffiliatesGetResponse_affiliatesable)
				}
			}
			m.SetAffiliates(res)
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
func (m *AffiliatesGetResponse) GetNext() *string {
	return m.next
}

// GetPrevious gets the previous property value. The previous property
// returns a *string when successful
func (m *AffiliatesGetResponse) GetPrevious() *string {
	return m.previous
}

// Serialize serializes information the current object
func (m *AffiliatesGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetAffiliates() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAffiliates()))
		for i, v := range m.GetAffiliates() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("affiliates", cast)
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
func (m *AffiliatesGetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAffiliates sets the affiliates property value. The affiliates property
func (m *AffiliatesGetResponse) SetAffiliates(value []AffiliatesGetResponse_affiliatesable) {
	m.affiliates = value
}

// SetCount sets the count property value. The count property
func (m *AffiliatesGetResponse) SetCount(value *int32) {
	m.count = value
}

// SetNext sets the next property value. The next property
func (m *AffiliatesGetResponse) SetNext(value *string) {
	m.next = value
}

// SetPrevious sets the previous property value. The previous property
func (m *AffiliatesGetResponse) SetPrevious(value *string) {
	m.previous = value
}

type AffiliatesGetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAffiliates() []AffiliatesGetResponse_affiliatesable
	GetCount() *int32
	GetNext() *string
	GetPrevious() *string
	SetAffiliates(value []AffiliatesGetResponse_affiliatesable)
	SetCount(value *int32)
	SetNext(value *string)
	SetPrevious(value *string)
}
