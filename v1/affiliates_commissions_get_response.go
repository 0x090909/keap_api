package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AffiliatesCommissionsGetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The commissions property
	commissions []AffiliatesCommissionsGetResponse_commissionsable
	// The count property
	count *int32
	// The next property
	next *string
	// The previous property
	previous *string
}

// NewAffiliatesCommissionsGetResponse instantiates a new AffiliatesCommissionsGetResponse and sets the default values.
func NewAffiliatesCommissionsGetResponse() *AffiliatesCommissionsGetResponse {
	m := &AffiliatesCommissionsGetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateAffiliatesCommissionsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAffiliatesCommissionsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAffiliatesCommissionsGetResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AffiliatesCommissionsGetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCommissions gets the commissions property value. The commissions property
// returns a []AffiliatesCommissionsGetResponse_commissionsable when successful
func (m *AffiliatesCommissionsGetResponse) GetCommissions() []AffiliatesCommissionsGetResponse_commissionsable {
	return m.commissions
}

// GetCount gets the count property value. The count property
// returns a *int32 when successful
func (m *AffiliatesCommissionsGetResponse) GetCount() *int32 {
	return m.count
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AffiliatesCommissionsGetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["commissions"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateAffiliatesCommissionsGetResponse_commissionsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]AffiliatesCommissionsGetResponse_commissionsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(AffiliatesCommissionsGetResponse_commissionsable)
				}
			}
			m.SetCommissions(res)
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
func (m *AffiliatesCommissionsGetResponse) GetNext() *string {
	return m.next
}

// GetPrevious gets the previous property value. The previous property
// returns a *string when successful
func (m *AffiliatesCommissionsGetResponse) GetPrevious() *string {
	return m.previous
}

// Serialize serializes information the current object
func (m *AffiliatesCommissionsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetCommissions() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCommissions()))
		for i, v := range m.GetCommissions() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("commissions", cast)
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
func (m *AffiliatesCommissionsGetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCommissions sets the commissions property value. The commissions property
func (m *AffiliatesCommissionsGetResponse) SetCommissions(value []AffiliatesCommissionsGetResponse_commissionsable) {
	m.commissions = value
}

// SetCount sets the count property value. The count property
func (m *AffiliatesCommissionsGetResponse) SetCount(value *int32) {
	m.count = value
}

// SetNext sets the next property value. The next property
func (m *AffiliatesCommissionsGetResponse) SetNext(value *string) {
	m.next = value
}

// SetPrevious sets the previous property value. The previous property
func (m *AffiliatesCommissionsGetResponse) SetPrevious(value *string) {
	m.previous = value
}

type AffiliatesCommissionsGetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCommissions() []AffiliatesCommissionsGetResponse_commissionsable
	GetCount() *int32
	GetNext() *string
	GetPrevious() *string
	SetCommissions(value []AffiliatesCommissionsGetResponse_commissionsable)
	SetCount(value *int32)
	SetNext(value *string)
	SetPrevious(value *string)
}
