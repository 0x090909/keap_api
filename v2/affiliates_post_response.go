package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AffiliatesPostResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The Affiliate Code
	code *string
	// The ContactID
	contact_id *string
	// The AffiliateId
	id *int64
	// The Affiliate Name
	name *string
	// The Affiliate Status
	status *string
	// The Affiliate PortalSite Id
	unique_site_id *string
}

// NewAffiliatesPostResponse instantiates a new AffiliatesPostResponse and sets the default values.
func NewAffiliatesPostResponse() *AffiliatesPostResponse {
	m := &AffiliatesPostResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateAffiliatesPostResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAffiliatesPostResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAffiliatesPostResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AffiliatesPostResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCode gets the code property value. The Affiliate Code
// returns a *string when successful
func (m *AffiliatesPostResponse) GetCode() *string {
	return m.code
}

// GetContactId gets the contact_id property value. The ContactID
// returns a *string when successful
func (m *AffiliatesPostResponse) GetContactId() *string {
	return m.contact_id
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AffiliatesPostResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["code"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCode(val)
		}
		return nil
	}
	res["contact_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetContactId(val)
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
	res["status"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetStatus(val)
		}
		return nil
	}
	res["unique_site_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetUniqueSiteId(val)
		}
		return nil
	}
	return res
}

// GetId gets the id property value. The AffiliateId
// returns a *int64 when successful
func (m *AffiliatesPostResponse) GetId() *int64 {
	return m.id
}

// GetName gets the name property value. The Affiliate Name
// returns a *string when successful
func (m *AffiliatesPostResponse) GetName() *string {
	return m.name
}

// GetStatus gets the status property value. The Affiliate Status
// returns a *string when successful
func (m *AffiliatesPostResponse) GetStatus() *string {
	return m.status
}

// GetUniqueSiteId gets the unique_site_id property value. The Affiliate PortalSite Id
// returns a *string when successful
func (m *AffiliatesPostResponse) GetUniqueSiteId() *string {
	return m.unique_site_id
}

// Serialize serializes information the current object
func (m *AffiliatesPostResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("code", m.GetCode())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("contact_id", m.GetContactId())
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
		err := writer.WriteStringValue("status", m.GetStatus())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("unique_site_id", m.GetUniqueSiteId())
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
func (m *AffiliatesPostResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCode sets the code property value. The Affiliate Code
func (m *AffiliatesPostResponse) SetCode(value *string) {
	m.code = value
}

// SetContactId sets the contact_id property value. The ContactID
func (m *AffiliatesPostResponse) SetContactId(value *string) {
	m.contact_id = value
}

// SetId sets the id property value. The AffiliateId
func (m *AffiliatesPostResponse) SetId(value *int64) {
	m.id = value
}

// SetName sets the name property value. The Affiliate Name
func (m *AffiliatesPostResponse) SetName(value *string) {
	m.name = value
}

// SetStatus sets the status property value. The Affiliate Status
func (m *AffiliatesPostResponse) SetStatus(value *string) {
	m.status = value
}

// SetUniqueSiteId sets the unique_site_id property value. The Affiliate PortalSite Id
func (m *AffiliatesPostResponse) SetUniqueSiteId(value *string) {
	m.unique_site_id = value
}

type AffiliatesPostResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCode() *string
	GetContactId() *string
	GetId() *int64
	GetName() *string
	GetStatus() *string
	GetUniqueSiteId() *string
	SetCode(value *string)
	SetContactId(value *string)
	SetId(value *int64)
	SetName(value *string)
	SetStatus(value *string)
	SetUniqueSiteId(value *string)
}
