package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AffiliatesCommissionProgramsItemWithCommission_program_PatchResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The id property
	id *string
	// The name property
	name *string
	// The notes property
	notes *string
	// The priority property
	priority *int32
}

// NewAffiliatesCommissionProgramsItemWithCommission_program_PatchResponse instantiates a new AffiliatesCommissionProgramsItemWithCommission_program_PatchResponse and sets the default values.
func NewAffiliatesCommissionProgramsItemWithCommission_program_PatchResponse() *AffiliatesCommissionProgramsItemWithCommission_program_PatchResponse {
	m := &AffiliatesCommissionProgramsItemWithCommission_program_PatchResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateAffiliatesCommissionProgramsItemWithCommission_program_PatchResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAffiliatesCommissionProgramsItemWithCommission_program_PatchResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAffiliatesCommissionProgramsItemWithCommission_program_PatchResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AffiliatesCommissionProgramsItemWithCommission_program_PatchResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AffiliatesCommissionProgramsItemWithCommission_program_PatchResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
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
// returns a *string when successful
func (m *AffiliatesCommissionProgramsItemWithCommission_program_PatchResponse) GetId() *string {
	return m.id
}

// GetName gets the name property value. The name property
// returns a *string when successful
func (m *AffiliatesCommissionProgramsItemWithCommission_program_PatchResponse) GetName() *string {
	return m.name
}

// GetNotes gets the notes property value. The notes property
// returns a *string when successful
func (m *AffiliatesCommissionProgramsItemWithCommission_program_PatchResponse) GetNotes() *string {
	return m.notes
}

// GetPriority gets the priority property value. The priority property
// returns a *int32 when successful
func (m *AffiliatesCommissionProgramsItemWithCommission_program_PatchResponse) GetPriority() *int32 {
	return m.priority
}

// Serialize serializes information the current object
func (m *AffiliatesCommissionProgramsItemWithCommission_program_PatchResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("id", m.GetId())
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
func (m *AffiliatesCommissionProgramsItemWithCommission_program_PatchResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetId sets the id property value. The id property
func (m *AffiliatesCommissionProgramsItemWithCommission_program_PatchResponse) SetId(value *string) {
	m.id = value
}

// SetName sets the name property value. The name property
func (m *AffiliatesCommissionProgramsItemWithCommission_program_PatchResponse) SetName(value *string) {
	m.name = value
}

// SetNotes sets the notes property value. The notes property
func (m *AffiliatesCommissionProgramsItemWithCommission_program_PatchResponse) SetNotes(value *string) {
	m.notes = value
}

// SetPriority sets the priority property value. The priority property
func (m *AffiliatesCommissionProgramsItemWithCommission_program_PatchResponse) SetPriority(value *int32) {
	m.priority = value
}

type AffiliatesCommissionProgramsItemWithCommission_program_PatchResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetId() *string
	GetName() *string
	GetNotes() *string
	GetPriority() *int32
	SetId(value *string)
	SetName(value *string)
	SetNotes(value *string)
	SetPriority(value *int32)
}
