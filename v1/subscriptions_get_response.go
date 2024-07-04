package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SubscriptionsGetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The count property
	count *int32
	// The next property
	next *string
	// The previous property
	previous *string
	// The subscriptions property
	subscriptions []SubscriptionsGetResponse_subscriptionsable
}

// NewSubscriptionsGetResponse instantiates a new SubscriptionsGetResponse and sets the default values.
func NewSubscriptionsGetResponse() *SubscriptionsGetResponse {
	m := &SubscriptionsGetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateSubscriptionsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSubscriptionsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSubscriptionsGetResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SubscriptionsGetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCount gets the count property value. The count property
// returns a *int32 when successful
func (m *SubscriptionsGetResponse) GetCount() *int32 {
	return m.count
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SubscriptionsGetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
	res["subscriptions"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateSubscriptionsGetResponse_subscriptionsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]SubscriptionsGetResponse_subscriptionsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(SubscriptionsGetResponse_subscriptionsable)
				}
			}
			m.SetSubscriptions(res)
		}
		return nil
	}
	return res
}

// GetNext gets the next property value. The next property
// returns a *string when successful
func (m *SubscriptionsGetResponse) GetNext() *string {
	return m.next
}

// GetPrevious gets the previous property value. The previous property
// returns a *string when successful
func (m *SubscriptionsGetResponse) GetPrevious() *string {
	return m.previous
}

// GetSubscriptions gets the subscriptions property value. The subscriptions property
// returns a []SubscriptionsGetResponse_subscriptionsable when successful
func (m *SubscriptionsGetResponse) GetSubscriptions() []SubscriptionsGetResponse_subscriptionsable {
	return m.subscriptions
}

// Serialize serializes information the current object
func (m *SubscriptionsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
	if m.GetSubscriptions() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSubscriptions()))
		for i, v := range m.GetSubscriptions() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("subscriptions", cast)
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
func (m *SubscriptionsGetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCount sets the count property value. The count property
func (m *SubscriptionsGetResponse) SetCount(value *int32) {
	m.count = value
}

// SetNext sets the next property value. The next property
func (m *SubscriptionsGetResponse) SetNext(value *string) {
	m.next = value
}

// SetPrevious sets the previous property value. The previous property
func (m *SubscriptionsGetResponse) SetPrevious(value *string) {
	m.previous = value
}

// SetSubscriptions sets the subscriptions property value. The subscriptions property
func (m *SubscriptionsGetResponse) SetSubscriptions(value []SubscriptionsGetResponse_subscriptionsable) {
	m.subscriptions = value
}

type SubscriptionsGetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCount() *int32
	GetNext() *string
	GetPrevious() *string
	GetSubscriptions() []SubscriptionsGetResponse_subscriptionsable
	SetCount(value *int32)
	SetNext(value *string)
	SetPrevious(value *string)
	SetSubscriptions(value []SubscriptionsGetResponse_subscriptionsable)
}
