package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ContactsItemContactIdPatchResponse_company struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The company_name property
	company_name *string
	// The id property
	id *int64
}

// NewContactsItemContactIdPatchResponse_company instantiates a new ContactsItemContactIdPatchResponse_company and sets the default values.
func NewContactsItemContactIdPatchResponse_company() *ContactsItemContactIdPatchResponse_company {
	m := &ContactsItemContactIdPatchResponse_company{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateContactsItemContactIdPatchResponse_companyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsItemContactIdPatchResponse_companyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsItemContactIdPatchResponse_company(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContactsItemContactIdPatchResponse_company) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCompanyName gets the company_name property value. The company_name property
// returns a *string when successful
func (m *ContactsItemContactIdPatchResponse_company) GetCompanyName() *string {
	return m.company_name
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContactsItemContactIdPatchResponse_company) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["company_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCompanyName(val)
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
	return res
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *ContactsItemContactIdPatchResponse_company) GetId() *int64 {
	return m.id
}

// Serialize serializes information the current object
func (m *ContactsItemContactIdPatchResponse_company) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("company_name", m.GetCompanyName())
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
		err := writer.WriteAdditionalData(m.GetAdditionalData())
		if err != nil {
			return err
		}
	}
	return nil
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ContactsItemContactIdPatchResponse_company) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCompanyName sets the company_name property value. The company_name property
func (m *ContactsItemContactIdPatchResponse_company) SetCompanyName(value *string) {
	m.company_name = value
}

// SetId sets the id property value. The id property
func (m *ContactsItemContactIdPatchResponse_company) SetId(value *int64) {
	m.id = value
}

type ContactsItemContactIdPatchResponse_companyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCompanyName() *string
	GetId() *int64
	SetCompanyName(value *string)
	SetId(value *int64)
}
