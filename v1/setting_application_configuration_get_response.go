package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SettingApplicationConfigurationGetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The affiliate property
	affiliate SettingApplicationConfigurationGetResponse_affiliateable
	// The application property
	application SettingApplicationConfigurationGetResponse_applicationable
	// The appointment property
	appointment SettingApplicationConfigurationGetResponse_appointmentable
	// The contact property
	contact SettingApplicationConfigurationGetResponse_contactable
	// The ecommerce property
	ecommerce SettingApplicationConfigurationGetResponse_ecommerceable
	// The email property
	email SettingApplicationConfigurationGetResponse_emailable
	// The forms property
	forms SettingApplicationConfigurationGetResponse_formsable
	// The fulfillment property
	fulfillment SettingApplicationConfigurationGetResponse_fulfillmentable
	// The invoice property
	invoice SettingApplicationConfigurationGetResponse_invoiceable
	// The note property
	note SettingApplicationConfigurationGetResponse_noteable
	// The opportunity property
	opportunity SettingApplicationConfigurationGetResponse_opportunityable
	// The task property
	task SettingApplicationConfigurationGetResponse_taskable
	// The template property
	template SettingApplicationConfigurationGetResponse_templateable
}

// NewSettingApplicationConfigurationGetResponse instantiates a new SettingApplicationConfigurationGetResponse and sets the default values.
func NewSettingApplicationConfigurationGetResponse() *SettingApplicationConfigurationGetResponse {
	m := &SettingApplicationConfigurationGetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateSettingApplicationConfigurationGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSettingApplicationConfigurationGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSettingApplicationConfigurationGetResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SettingApplicationConfigurationGetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAffiliate gets the affiliate property value. The affiliate property
// returns a SettingApplicationConfigurationGetResponse_affiliateable when successful
func (m *SettingApplicationConfigurationGetResponse) GetAffiliate() SettingApplicationConfigurationGetResponse_affiliateable {
	return m.affiliate
}

// GetApplication gets the application property value. The application property
// returns a SettingApplicationConfigurationGetResponse_applicationable when successful
func (m *SettingApplicationConfigurationGetResponse) GetApplication() SettingApplicationConfigurationGetResponse_applicationable {
	return m.application
}

// GetAppointment gets the appointment property value. The appointment property
// returns a SettingApplicationConfigurationGetResponse_appointmentable when successful
func (m *SettingApplicationConfigurationGetResponse) GetAppointment() SettingApplicationConfigurationGetResponse_appointmentable {
	return m.appointment
}

// GetContact gets the contact property value. The contact property
// returns a SettingApplicationConfigurationGetResponse_contactable when successful
func (m *SettingApplicationConfigurationGetResponse) GetContact() SettingApplicationConfigurationGetResponse_contactable {
	return m.contact
}

// GetEcommerce gets the ecommerce property value. The ecommerce property
// returns a SettingApplicationConfigurationGetResponse_ecommerceable when successful
func (m *SettingApplicationConfigurationGetResponse) GetEcommerce() SettingApplicationConfigurationGetResponse_ecommerceable {
	return m.ecommerce
}

// GetEmail gets the email property value. The email property
// returns a SettingApplicationConfigurationGetResponse_emailable when successful
func (m *SettingApplicationConfigurationGetResponse) GetEmail() SettingApplicationConfigurationGetResponse_emailable {
	return m.email
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SettingApplicationConfigurationGetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["affiliate"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSettingApplicationConfigurationGetResponse_affiliateFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAffiliate(val.(SettingApplicationConfigurationGetResponse_affiliateable))
		}
		return nil
	}
	res["application"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSettingApplicationConfigurationGetResponse_applicationFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetApplication(val.(SettingApplicationConfigurationGetResponse_applicationable))
		}
		return nil
	}
	res["appointment"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSettingApplicationConfigurationGetResponse_appointmentFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAppointment(val.(SettingApplicationConfigurationGetResponse_appointmentable))
		}
		return nil
	}
	res["contact"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSettingApplicationConfigurationGetResponse_contactFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetContact(val.(SettingApplicationConfigurationGetResponse_contactable))
		}
		return nil
	}
	res["ecommerce"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSettingApplicationConfigurationGetResponse_ecommerceFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetEcommerce(val.(SettingApplicationConfigurationGetResponse_ecommerceable))
		}
		return nil
	}
	res["email"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSettingApplicationConfigurationGetResponse_emailFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetEmail(val.(SettingApplicationConfigurationGetResponse_emailable))
		}
		return nil
	}
	res["forms"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSettingApplicationConfigurationGetResponse_formsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetForms(val.(SettingApplicationConfigurationGetResponse_formsable))
		}
		return nil
	}
	res["fulfillment"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSettingApplicationConfigurationGetResponse_fulfillmentFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFulfillment(val.(SettingApplicationConfigurationGetResponse_fulfillmentable))
		}
		return nil
	}
	res["invoice"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSettingApplicationConfigurationGetResponse_invoiceFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetInvoice(val.(SettingApplicationConfigurationGetResponse_invoiceable))
		}
		return nil
	}
	res["note"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSettingApplicationConfigurationGetResponse_noteFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNote(val.(SettingApplicationConfigurationGetResponse_noteable))
		}
		return nil
	}
	res["opportunity"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSettingApplicationConfigurationGetResponse_opportunityFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetOpportunity(val.(SettingApplicationConfigurationGetResponse_opportunityable))
		}
		return nil
	}
	res["task"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSettingApplicationConfigurationGetResponse_taskFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTask(val.(SettingApplicationConfigurationGetResponse_taskable))
		}
		return nil
	}
	res["template"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSettingApplicationConfigurationGetResponse_templateFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTemplate(val.(SettingApplicationConfigurationGetResponse_templateable))
		}
		return nil
	}
	return res
}

// GetForms gets the forms property value. The forms property
// returns a SettingApplicationConfigurationGetResponse_formsable when successful
func (m *SettingApplicationConfigurationGetResponse) GetForms() SettingApplicationConfigurationGetResponse_formsable {
	return m.forms
}

// GetFulfillment gets the fulfillment property value. The fulfillment property
// returns a SettingApplicationConfigurationGetResponse_fulfillmentable when successful
func (m *SettingApplicationConfigurationGetResponse) GetFulfillment() SettingApplicationConfigurationGetResponse_fulfillmentable {
	return m.fulfillment
}

// GetInvoice gets the invoice property value. The invoice property
// returns a SettingApplicationConfigurationGetResponse_invoiceable when successful
func (m *SettingApplicationConfigurationGetResponse) GetInvoice() SettingApplicationConfigurationGetResponse_invoiceable {
	return m.invoice
}

// GetNote gets the note property value. The note property
// returns a SettingApplicationConfigurationGetResponse_noteable when successful
func (m *SettingApplicationConfigurationGetResponse) GetNote() SettingApplicationConfigurationGetResponse_noteable {
	return m.note
}

// GetOpportunity gets the opportunity property value. The opportunity property
// returns a SettingApplicationConfigurationGetResponse_opportunityable when successful
func (m *SettingApplicationConfigurationGetResponse) GetOpportunity() SettingApplicationConfigurationGetResponse_opportunityable {
	return m.opportunity
}

// GetTask gets the task property value. The task property
// returns a SettingApplicationConfigurationGetResponse_taskable when successful
func (m *SettingApplicationConfigurationGetResponse) GetTask() SettingApplicationConfigurationGetResponse_taskable {
	return m.task
}

// GetTemplate gets the template property value. The template property
// returns a SettingApplicationConfigurationGetResponse_templateable when successful
func (m *SettingApplicationConfigurationGetResponse) GetTemplate() SettingApplicationConfigurationGetResponse_templateable {
	return m.template
}

// Serialize serializes information the current object
func (m *SettingApplicationConfigurationGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("affiliate", m.GetAffiliate())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("application", m.GetApplication())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("appointment", m.GetAppointment())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("contact", m.GetContact())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("ecommerce", m.GetEcommerce())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("email", m.GetEmail())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("forms", m.GetForms())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("fulfillment", m.GetFulfillment())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("invoice", m.GetInvoice())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("note", m.GetNote())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("opportunity", m.GetOpportunity())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("task", m.GetTask())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("template", m.GetTemplate())
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
func (m *SettingApplicationConfigurationGetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAffiliate sets the affiliate property value. The affiliate property
func (m *SettingApplicationConfigurationGetResponse) SetAffiliate(value SettingApplicationConfigurationGetResponse_affiliateable) {
	m.affiliate = value
}

// SetApplication sets the application property value. The application property
func (m *SettingApplicationConfigurationGetResponse) SetApplication(value SettingApplicationConfigurationGetResponse_applicationable) {
	m.application = value
}

// SetAppointment sets the appointment property value. The appointment property
func (m *SettingApplicationConfigurationGetResponse) SetAppointment(value SettingApplicationConfigurationGetResponse_appointmentable) {
	m.appointment = value
}

// SetContact sets the contact property value. The contact property
func (m *SettingApplicationConfigurationGetResponse) SetContact(value SettingApplicationConfigurationGetResponse_contactable) {
	m.contact = value
}

// SetEcommerce sets the ecommerce property value. The ecommerce property
func (m *SettingApplicationConfigurationGetResponse) SetEcommerce(value SettingApplicationConfigurationGetResponse_ecommerceable) {
	m.ecommerce = value
}

// SetEmail sets the email property value. The email property
func (m *SettingApplicationConfigurationGetResponse) SetEmail(value SettingApplicationConfigurationGetResponse_emailable) {
	m.email = value
}

// SetForms sets the forms property value. The forms property
func (m *SettingApplicationConfigurationGetResponse) SetForms(value SettingApplicationConfigurationGetResponse_formsable) {
	m.forms = value
}

// SetFulfillment sets the fulfillment property value. The fulfillment property
func (m *SettingApplicationConfigurationGetResponse) SetFulfillment(value SettingApplicationConfigurationGetResponse_fulfillmentable) {
	m.fulfillment = value
}

// SetInvoice sets the invoice property value. The invoice property
func (m *SettingApplicationConfigurationGetResponse) SetInvoice(value SettingApplicationConfigurationGetResponse_invoiceable) {
	m.invoice = value
}

// SetNote sets the note property value. The note property
func (m *SettingApplicationConfigurationGetResponse) SetNote(value SettingApplicationConfigurationGetResponse_noteable) {
	m.note = value
}

// SetOpportunity sets the opportunity property value. The opportunity property
func (m *SettingApplicationConfigurationGetResponse) SetOpportunity(value SettingApplicationConfigurationGetResponse_opportunityable) {
	m.opportunity = value
}

// SetTask sets the task property value. The task property
func (m *SettingApplicationConfigurationGetResponse) SetTask(value SettingApplicationConfigurationGetResponse_taskable) {
	m.task = value
}

// SetTemplate sets the template property value. The template property
func (m *SettingApplicationConfigurationGetResponse) SetTemplate(value SettingApplicationConfigurationGetResponse_templateable) {
	m.template = value
}

type SettingApplicationConfigurationGetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAffiliate() SettingApplicationConfigurationGetResponse_affiliateable
	GetApplication() SettingApplicationConfigurationGetResponse_applicationable
	GetAppointment() SettingApplicationConfigurationGetResponse_appointmentable
	GetContact() SettingApplicationConfigurationGetResponse_contactable
	GetEcommerce() SettingApplicationConfigurationGetResponse_ecommerceable
	GetEmail() SettingApplicationConfigurationGetResponse_emailable
	GetForms() SettingApplicationConfigurationGetResponse_formsable
	GetFulfillment() SettingApplicationConfigurationGetResponse_fulfillmentable
	GetInvoice() SettingApplicationConfigurationGetResponse_invoiceable
	GetNote() SettingApplicationConfigurationGetResponse_noteable
	GetOpportunity() SettingApplicationConfigurationGetResponse_opportunityable
	GetTask() SettingApplicationConfigurationGetResponse_taskable
	GetTemplate() SettingApplicationConfigurationGetResponse_templateable
	SetAffiliate(value SettingApplicationConfigurationGetResponse_affiliateable)
	SetApplication(value SettingApplicationConfigurationGetResponse_applicationable)
	SetAppointment(value SettingApplicationConfigurationGetResponse_appointmentable)
	SetContact(value SettingApplicationConfigurationGetResponse_contactable)
	SetEcommerce(value SettingApplicationConfigurationGetResponse_ecommerceable)
	SetEmail(value SettingApplicationConfigurationGetResponse_emailable)
	SetForms(value SettingApplicationConfigurationGetResponse_formsable)
	SetFulfillment(value SettingApplicationConfigurationGetResponse_fulfillmentable)
	SetInvoice(value SettingApplicationConfigurationGetResponse_invoiceable)
	SetNote(value SettingApplicationConfigurationGetResponse_noteable)
	SetOpportunity(value SettingApplicationConfigurationGetResponse_opportunityable)
	SetTask(value SettingApplicationConfigurationGetResponse_taskable)
	SetTemplate(value SettingApplicationConfigurationGetResponse_templateable)
}
