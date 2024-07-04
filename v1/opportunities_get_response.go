package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OpportunitiesGetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The count property
	count *int32
	// The next property
	next *string
	// The opportunities property
	opportunities []OpportunitiesGetResponse_opportunitiesable
	// The previous property
	previous *string
}

// NewOpportunitiesGetResponse instantiates a new OpportunitiesGetResponse and sets the default values.
func NewOpportunitiesGetResponse() *OpportunitiesGetResponse {
	m := &OpportunitiesGetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateOpportunitiesGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOpportunitiesGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewOpportunitiesGetResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OpportunitiesGetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCount gets the count property value. The count property
// returns a *int32 when successful
func (m *OpportunitiesGetResponse) GetCount() *int32 {
	return m.count
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OpportunitiesGetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
	res["opportunities"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateOpportunitiesGetResponse_opportunitiesFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]OpportunitiesGetResponse_opportunitiesable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(OpportunitiesGetResponse_opportunitiesable)
				}
			}
			m.SetOpportunities(res)
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
func (m *OpportunitiesGetResponse) GetNext() *string {
	return m.next
}

// GetOpportunities gets the opportunities property value. The opportunities property
// returns a []OpportunitiesGetResponse_opportunitiesable when successful
func (m *OpportunitiesGetResponse) GetOpportunities() []OpportunitiesGetResponse_opportunitiesable {
	return m.opportunities
}

// GetPrevious gets the previous property value. The previous property
// returns a *string when successful
func (m *OpportunitiesGetResponse) GetPrevious() *string {
	return m.previous
}

// Serialize serializes information the current object
func (m *OpportunitiesGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
	if m.GetOpportunities() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOpportunities()))
		for i, v := range m.GetOpportunities() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("opportunities", cast)
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
func (m *OpportunitiesGetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCount sets the count property value. The count property
func (m *OpportunitiesGetResponse) SetCount(value *int32) {
	m.count = value
}

// SetNext sets the next property value. The next property
func (m *OpportunitiesGetResponse) SetNext(value *string) {
	m.next = value
}

// SetOpportunities sets the opportunities property value. The opportunities property
func (m *OpportunitiesGetResponse) SetOpportunities(value []OpportunitiesGetResponse_opportunitiesable) {
	m.opportunities = value
}

// SetPrevious sets the previous property value. The previous property
func (m *OpportunitiesGetResponse) SetPrevious(value *string) {
	m.previous = value
}

type OpportunitiesGetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCount() *int32
	GetNext() *string
	GetOpportunities() []OpportunitiesGetResponse_opportunitiesable
	GetPrevious() *string
	SetCount(value *int32)
	SetNext(value *string)
	SetOpportunities(value []OpportunitiesGetResponse_opportunitiesable)
	SetPrevious(value *string)
}
