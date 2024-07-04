package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OpportunitiesPostResponse_user struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The first_name property
	first_name *string
	// The id property
	id *int64
	// The last_name property
	last_name *string
}

// NewOpportunitiesPostResponse_user instantiates a new OpportunitiesPostResponse_user and sets the default values.
func NewOpportunitiesPostResponse_user() *OpportunitiesPostResponse_user {
	m := &OpportunitiesPostResponse_user{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateOpportunitiesPostResponse_userFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOpportunitiesPostResponse_userFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewOpportunitiesPostResponse_user(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OpportunitiesPostResponse_user) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OpportunitiesPostResponse_user) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["first_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFirstName(val)
		}
		return nil
	}
	res["id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetId(val)
		}
		return nil
	}
	res["last_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLastName(val)
		}
		return nil
	}
	return res
}

// GetFirstName gets the first_name property value. The first_name property
// returns a *string when successful
func (m *OpportunitiesPostResponse_user) GetFirstName() *string {
	return m.first_name
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *OpportunitiesPostResponse_user) GetId() *int64 {
	return m.id
}

// GetLastName gets the last_name property value. The last_name property
// returns a *string when successful
func (m *OpportunitiesPostResponse_user) GetLastName() *string {
	return m.last_name
}

// Serialize serializes information the current object
func (m *OpportunitiesPostResponse_user) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("first_name", m.GetFirstName())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("id", m.GetId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("last_name", m.GetLastName())
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
func (m *OpportunitiesPostResponse_user) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetFirstName sets the first_name property value. The first_name property
func (m *OpportunitiesPostResponse_user) SetFirstName(value *string) {
	m.first_name = value
}

// SetId sets the id property value. The id property
func (m *OpportunitiesPostResponse_user) SetId(value *int64) {
	m.id = value
}

// SetLastName sets the last_name property value. The last_name property
func (m *OpportunitiesPostResponse_user) SetLastName(value *string) {
	m.last_name = value
}

type OpportunitiesPostResponse_userable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetFirstName() *string
	GetId() *int64
	GetLastName() *string
	SetFirstName(value *string)
	SetId(value *int64)
	SetLastName(value *string)
}
