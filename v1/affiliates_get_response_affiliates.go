package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AffiliatesGetResponse_affiliates struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The code property
	code *string
	// The contact_id property
	contact_id *int64
	// The id property
	id *int64
	// The name property
	name *string
	// The notify_on_lead property
	notify_on_lead *bool
	// The notify_on_sale property
	notify_on_sale *bool
	// The parent_id property
	parent_id *int64
	// The status property
	status *string
	// The track_leads_for property
	track_leads_for *int32
}

// NewAffiliatesGetResponse_affiliates instantiates a new AffiliatesGetResponse_affiliates and sets the default values.
func NewAffiliatesGetResponse_affiliates() *AffiliatesGetResponse_affiliates {
	m := &AffiliatesGetResponse_affiliates{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateAffiliatesGetResponse_affiliatesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAffiliatesGetResponse_affiliatesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAffiliatesGetResponse_affiliates(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AffiliatesGetResponse_affiliates) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCode gets the code property value. The code property
// returns a *string when successful
func (m *AffiliatesGetResponse_affiliates) GetCode() *string {
	return m.code
}

// GetContactId gets the contact_id property value. The contact_id property
// returns a *int64 when successful
func (m *AffiliatesGetResponse_affiliates) GetContactId() *int64 {
	return m.contact_id
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AffiliatesGetResponse_affiliates) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
		val, err := n.GetInt64Value()
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
	res["notify_on_lead"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNotifyOnLead(val)
		}
		return nil
	}
	res["notify_on_sale"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNotifyOnSale(val)
		}
		return nil
	}
	res["parent_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetParentId(val)
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
	res["track_leads_for"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTrackLeadsFor(val)
		}
		return nil
	}
	return res
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *AffiliatesGetResponse_affiliates) GetId() *int64 {
	return m.id
}

// GetName gets the name property value. The name property
// returns a *string when successful
func (m *AffiliatesGetResponse_affiliates) GetName() *string {
	return m.name
}

// GetNotifyOnLead gets the notify_on_lead property value. The notify_on_lead property
// returns a *bool when successful
func (m *AffiliatesGetResponse_affiliates) GetNotifyOnLead() *bool {
	return m.notify_on_lead
}

// GetNotifyOnSale gets the notify_on_sale property value. The notify_on_sale property
// returns a *bool when successful
func (m *AffiliatesGetResponse_affiliates) GetNotifyOnSale() *bool {
	return m.notify_on_sale
}

// GetParentId gets the parent_id property value. The parent_id property
// returns a *int64 when successful
func (m *AffiliatesGetResponse_affiliates) GetParentId() *int64 {
	return m.parent_id
}

// GetStatus gets the status property value. The status property
// returns a *string when successful
func (m *AffiliatesGetResponse_affiliates) GetStatus() *string {
	return m.status
}

// GetTrackLeadsFor gets the track_leads_for property value. The track_leads_for property
// returns a *int32 when successful
func (m *AffiliatesGetResponse_affiliates) GetTrackLeadsFor() *int32 {
	return m.track_leads_for
}

// Serialize serializes information the current object
func (m *AffiliatesGetResponse_affiliates) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("code", m.GetCode())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("contact_id", m.GetContactId())
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
		err := writer.WriteBoolValue("notify_on_lead", m.GetNotifyOnLead())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("notify_on_sale", m.GetNotifyOnSale())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("parent_id", m.GetParentId())
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
		err := writer.WriteInt32Value("track_leads_for", m.GetTrackLeadsFor())
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
func (m *AffiliatesGetResponse_affiliates) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCode sets the code property value. The code property
func (m *AffiliatesGetResponse_affiliates) SetCode(value *string) {
	m.code = value
}

// SetContactId sets the contact_id property value. The contact_id property
func (m *AffiliatesGetResponse_affiliates) SetContactId(value *int64) {
	m.contact_id = value
}

// SetId sets the id property value. The id property
func (m *AffiliatesGetResponse_affiliates) SetId(value *int64) {
	m.id = value
}

// SetName sets the name property value. The name property
func (m *AffiliatesGetResponse_affiliates) SetName(value *string) {
	m.name = value
}

// SetNotifyOnLead sets the notify_on_lead property value. The notify_on_lead property
func (m *AffiliatesGetResponse_affiliates) SetNotifyOnLead(value *bool) {
	m.notify_on_lead = value
}

// SetNotifyOnSale sets the notify_on_sale property value. The notify_on_sale property
func (m *AffiliatesGetResponse_affiliates) SetNotifyOnSale(value *bool) {
	m.notify_on_sale = value
}

// SetParentId sets the parent_id property value. The parent_id property
func (m *AffiliatesGetResponse_affiliates) SetParentId(value *int64) {
	m.parent_id = value
}

// SetStatus sets the status property value. The status property
func (m *AffiliatesGetResponse_affiliates) SetStatus(value *string) {
	m.status = value
}

// SetTrackLeadsFor sets the track_leads_for property value. The track_leads_for property
func (m *AffiliatesGetResponse_affiliates) SetTrackLeadsFor(value *int32) {
	m.track_leads_for = value
}

type AffiliatesGetResponse_affiliatesable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCode() *string
	GetContactId() *int64
	GetId() *int64
	GetName() *string
	GetNotifyOnLead() *bool
	GetNotifyOnSale() *bool
	GetParentId() *int64
	GetStatus() *string
	GetTrackLeadsFor() *int32
	SetCode(value *string)
	SetContactId(value *int64)
	SetId(value *int64)
	SetName(value *string)
	SetNotifyOnLead(value *bool)
	SetNotifyOnSale(value *bool)
	SetParentId(value *int64)
	SetStatus(value *string)
	SetTrackLeadsFor(value *int32)
}
