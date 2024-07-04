package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OpportunitiesPostRequestBody_stage struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The details property
	details OpportunitiesPostRequestBody_stage_detailsable
	// The id property
	id *int64
	// The name property
	name *string
	// The reasons property
	reasons []string
}

// NewOpportunitiesPostRequestBody_stage instantiates a new OpportunitiesPostRequestBody_stage and sets the default values.
func NewOpportunitiesPostRequestBody_stage() *OpportunitiesPostRequestBody_stage {
	m := &OpportunitiesPostRequestBody_stage{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateOpportunitiesPostRequestBody_stageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOpportunitiesPostRequestBody_stageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewOpportunitiesPostRequestBody_stage(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OpportunitiesPostRequestBody_stage) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetDetails gets the details property value. The details property
// returns a OpportunitiesPostRequestBody_stage_detailsable when successful
func (m *OpportunitiesPostRequestBody_stage) GetDetails() OpportunitiesPostRequestBody_stage_detailsable {
	return m.details
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OpportunitiesPostRequestBody_stage) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["details"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateOpportunitiesPostRequestBody_stage_detailsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDetails(val.(OpportunitiesPostRequestBody_stage_detailsable))
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
	res["name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetName(val)
		}
		return nil
	}
	res["reasons"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfPrimitiveValues("string")
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]string, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = *(v.(*string))
				}
			}
			m.SetReasons(res)
		}
		return nil
	}
	return res
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *OpportunitiesPostRequestBody_stage) GetId() *int64 {
	return m.id
}

// GetName gets the name property value. The name property
// returns a *string when successful
func (m *OpportunitiesPostRequestBody_stage) GetName() *string {
	return m.name
}

// GetReasons gets the reasons property value. The reasons property
// returns a []string when successful
func (m *OpportunitiesPostRequestBody_stage) GetReasons() []string {
	return m.reasons
}

// Serialize serializes information the current object
func (m *OpportunitiesPostRequestBody_stage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("details", m.GetDetails())
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
		err := writer.WriteStringValue("name", m.GetName())
		if err != nil {
			return err
		}
	}
	if m.GetReasons() != nil {
		err := writer.WriteCollectionOfStringValues("reasons", m.GetReasons())
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
func (m *OpportunitiesPostRequestBody_stage) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetDetails sets the details property value. The details property
func (m *OpportunitiesPostRequestBody_stage) SetDetails(value OpportunitiesPostRequestBody_stage_detailsable) {
	m.details = value
}

// SetId sets the id property value. The id property
func (m *OpportunitiesPostRequestBody_stage) SetId(value *int64) {
	m.id = value
}

// SetName sets the name property value. The name property
func (m *OpportunitiesPostRequestBody_stage) SetName(value *string) {
	m.name = value
}

// SetReasons sets the reasons property value. The reasons property
func (m *OpportunitiesPostRequestBody_stage) SetReasons(value []string) {
	m.reasons = value
}

type OpportunitiesPostRequestBody_stageable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetDetails() OpportunitiesPostRequestBody_stage_detailsable
	GetId() *int64
	GetName() *string
	GetReasons() []string
	SetDetails(value OpportunitiesPostRequestBody_stage_detailsable)
	SetId(value *int64)
	SetName(value *string)
	SetReasons(value []string)
}
