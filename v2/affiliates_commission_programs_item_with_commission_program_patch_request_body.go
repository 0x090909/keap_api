package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AffiliatesCommissionProgramsItemWithCommission_program_PatchRequestBody struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The name of the Commission Program
	name *string
	// The notes of the Commission Program
	notes *string
	// The priority of the Commission Program
	priority *int32
}

// NewAffiliatesCommissionProgramsItemWithCommission_program_PatchRequestBody instantiates a new AffiliatesCommissionProgramsItemWithCommission_program_PatchRequestBody and sets the default values.
func NewAffiliatesCommissionProgramsItemWithCommission_program_PatchRequestBody() *AffiliatesCommissionProgramsItemWithCommission_program_PatchRequestBody {
	m := &AffiliatesCommissionProgramsItemWithCommission_program_PatchRequestBody{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateAffiliatesCommissionProgramsItemWithCommission_program_PatchRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAffiliatesCommissionProgramsItemWithCommission_program_PatchRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAffiliatesCommissionProgramsItemWithCommission_program_PatchRequestBody(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AffiliatesCommissionProgramsItemWithCommission_program_PatchRequestBody) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AffiliatesCommissionProgramsItemWithCommission_program_PatchRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
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

// GetName gets the name property value. The name of the Commission Program
// returns a *string when successful
func (m *AffiliatesCommissionProgramsItemWithCommission_program_PatchRequestBody) GetName() *string {
	return m.name
}

// GetNotes gets the notes property value. The notes of the Commission Program
// returns a *string when successful
func (m *AffiliatesCommissionProgramsItemWithCommission_program_PatchRequestBody) GetNotes() *string {
	return m.notes
}

// GetPriority gets the priority property value. The priority of the Commission Program
// returns a *int32 when successful
func (m *AffiliatesCommissionProgramsItemWithCommission_program_PatchRequestBody) GetPriority() *int32 {
	return m.priority
}

// Serialize serializes information the current object
func (m *AffiliatesCommissionProgramsItemWithCommission_program_PatchRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *AffiliatesCommissionProgramsItemWithCommission_program_PatchRequestBody) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetName sets the name property value. The name of the Commission Program
func (m *AffiliatesCommissionProgramsItemWithCommission_program_PatchRequestBody) SetName(value *string) {
	m.name = value
}

// SetNotes sets the notes property value. The notes of the Commission Program
func (m *AffiliatesCommissionProgramsItemWithCommission_program_PatchRequestBody) SetNotes(value *string) {
	m.notes = value
}

// SetPriority sets the priority property value. The priority of the Commission Program
func (m *AffiliatesCommissionProgramsItemWithCommission_program_PatchRequestBody) SetPriority(value *int32) {
	m.priority = value
}

type AffiliatesCommissionProgramsItemWithCommission_program_PatchRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetName() *string
	GetNotes() *string
	GetPriority() *int32
	SetName(value *string)
	SetNotes(value *string)
	SetPriority(value *int32)
}
