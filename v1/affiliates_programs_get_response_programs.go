package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AffiliatesProgramsGetResponse_programs struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The affiliate_id property
	affiliate_id *int64
	// The id property
	id *int64
	// The name property
	name *string
	// The notes property
	notes *string
	// The priority property
	priority *int32
}

// NewAffiliatesProgramsGetResponse_programs instantiates a new AffiliatesProgramsGetResponse_programs and sets the default values.
func NewAffiliatesProgramsGetResponse_programs() *AffiliatesProgramsGetResponse_programs {
	m := &AffiliatesProgramsGetResponse_programs{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateAffiliatesProgramsGetResponse_programsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAffiliatesProgramsGetResponse_programsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAffiliatesProgramsGetResponse_programs(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AffiliatesProgramsGetResponse_programs) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAffiliateId gets the affiliate_id property value. The affiliate_id property
// returns a *int64 when successful
func (m *AffiliatesProgramsGetResponse_programs) GetAffiliateId() *int64 {
	return m.affiliate_id
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AffiliatesProgramsGetResponse_programs) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
	res["notes"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNotes(val)
		}
		return nil
	}
	res["priority"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPriority(val)
		}
		return nil
	}
	return res
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *AffiliatesProgramsGetResponse_programs) GetId() *int64 {
	return m.id
}

// GetName gets the name property value. The name property
// returns a *string when successful
func (m *AffiliatesProgramsGetResponse_programs) GetName() *string {
	return m.name
}

// GetNotes gets the notes property value. The notes property
// returns a *string when successful
func (m *AffiliatesProgramsGetResponse_programs) GetNotes() *string {
	return m.notes
}

// GetPriority gets the priority property value. The priority property
// returns a *int32 when successful
func (m *AffiliatesProgramsGetResponse_programs) GetPriority() *int32 {
	return m.priority
}

// Serialize serializes information the current object
func (m *AffiliatesProgramsGetResponse_programs) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt64Value("affiliate_id", m.GetAffiliateId())
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
	{
		err := writer.WriteStringValue("notes", m.GetNotes())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("priority", m.GetPriority())
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
func (m *AffiliatesProgramsGetResponse_programs) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAffiliateId sets the affiliate_id property value. The affiliate_id property
func (m *AffiliatesProgramsGetResponse_programs) SetAffiliateId(value *int64) {
	m.affiliate_id = value
}

// SetId sets the id property value. The id property
func (m *AffiliatesProgramsGetResponse_programs) SetId(value *int64) {
	m.id = value
}

// SetName sets the name property value. The name property
func (m *AffiliatesProgramsGetResponse_programs) SetName(value *string) {
	m.name = value
}

// SetNotes sets the notes property value. The notes property
func (m *AffiliatesProgramsGetResponse_programs) SetNotes(value *string) {
	m.notes = value
}

// SetPriority sets the priority property value. The priority property
func (m *AffiliatesProgramsGetResponse_programs) SetPriority(value *int32) {
	m.priority = value
}

type AffiliatesProgramsGetResponse_programsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAffiliateId() *int64
	GetId() *int64
	GetName() *string
	GetNotes() *string
	GetPriority() *int32
	SetAffiliateId(value *int64)
	SetId(value *int64)
	SetName(value *string)
	SetNotes(value *string)
	SetPriority(value *int32)
}
