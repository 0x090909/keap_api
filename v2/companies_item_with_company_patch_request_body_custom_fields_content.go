package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type CompaniesItemWithCompany_PatchRequestBody_custom_fields_content struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
}

// NewCompaniesItemWithCompany_PatchRequestBody_custom_fields_content instantiates a new CompaniesItemWithCompany_PatchRequestBody_custom_fields_content and sets the default values.
func NewCompaniesItemWithCompany_PatchRequestBody_custom_fields_content() *CompaniesItemWithCompany_PatchRequestBody_custom_fields_content {
	m := &CompaniesItemWithCompany_PatchRequestBody_custom_fields_content{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateCompaniesItemWithCompany_PatchRequestBody_custom_fields_contentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateCompaniesItemWithCompany_PatchRequestBody_custom_fields_contentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewCompaniesItemWithCompany_PatchRequestBody_custom_fields_content(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *CompaniesItemWithCompany_PatchRequestBody_custom_fields_content) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *CompaniesItemWithCompany_PatchRequestBody_custom_fields_content) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	return res
}

// Serialize serializes information the current object
func (m *CompaniesItemWithCompany_PatchRequestBody_custom_fields_content) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteAdditionalData(m.GetAdditionalData())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CompaniesItemWithCompany_PatchRequestBody_custom_fields_content) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

type CompaniesItemWithCompany_PatchRequestBody_custom_fields_contentable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
