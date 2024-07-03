package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SettingsApplicationsGetConfigurationGetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The affiliate property
	affiliate SettingsApplicationsGetConfigurationGetResponse_affiliateable
	// The application property
	application SettingsApplicationsGetConfigurationGetResponse_applicationable
	// The appointment property
	appointment SettingsApplicationsGetConfigurationGetResponse_appointmentable
	// The contact property
	contact SettingsApplicationsGetConfigurationGetResponse_contactable
	// The ecommerce property
	ecommerce SettingsApplicationsGetConfigurationGetResponse_ecommerceable
	// The email property
	email SettingsApplicationsGetConfigurationGetResponse_emailable
	// The forms property
	forms SettingsApplicationsGetConfigurationGetResponse_formsable
	// The fulfillment property
	fulfillment SettingsApplicationsGetConfigurationGetResponse_fulfillmentable
	// The invoice property
	invoice SettingsApplicationsGetConfigurationGetResponse_invoiceable
	// The note property
	note SettingsApplicationsGetConfigurationGetResponse_noteable
	// The opportunity property
	opportunity SettingsApplicationsGetConfigurationGetResponse_opportunityable
	// The task property
	task SettingsApplicationsGetConfigurationGetResponse_taskable
	// The template property
	template SettingsApplicationsGetConfigurationGetResponse_templateable
}

// NewSettingsApplicationsGetConfigurationGetResponse instantiates a new SettingsApplicationsGetConfigurationGetResponse and sets the default values.
func NewSettingsApplicationsGetConfigurationGetResponse() *SettingsApplicationsGetConfigurationGetResponse {
	m := &SettingsApplicationsGetConfigurationGetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateSettingsApplicationsGetConfigurationGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSettingsApplicationsGetConfigurationGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSettingsApplicationsGetConfigurationGetResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SettingsApplicationsGetConfigurationGetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAffiliate gets the affiliate property value. The affiliate property
// returns a SettingsApplicationsGetConfigurationGetResponse_affiliateable when successful
func (m *SettingsApplicationsGetConfigurationGetResponse) GetAffiliate() SettingsApplicationsGetConfigurationGetResponse_affiliateable {
	return m.affiliate
}

// GetApplication gets the application property value. The application property
// returns a SettingsApplicationsGetConfigurationGetResponse_applicationable when successful
func (m *SettingsApplicationsGetConfigurationGetResponse) GetApplication() SettingsApplicationsGetConfigurationGetResponse_applicationable {
	return m.application
}

// GetAppointment gets the appointment property value. The appointment property
// returns a SettingsApplicationsGetConfigurationGetResponse_appointmentable when successful
func (m *SettingsApplicationsGetConfigurationGetResponse) GetAppointment() SettingsApplicationsGetConfigurationGetResponse_appointmentable {
	return m.appointment
}

// GetContact gets the contact property value. The contact property
// returns a SettingsApplicationsGetConfigurationGetResponse_contactable when successful
func (m *SettingsApplicationsGetConfigurationGetResponse) GetContact() SettingsApplicationsGetConfigurationGetResponse_contactable {
	return m.contact
}

// GetEcommerce gets the ecommerce property value. The ecommerce property
// returns a SettingsApplicationsGetConfigurationGetResponse_ecommerceable when successful
func (m *SettingsApplicationsGetConfigurationGetResponse) GetEcommerce() SettingsApplicationsGetConfigurationGetResponse_ecommerceable {
	return m.ecommerce
}

// GetEmail gets the email property value. The email property
// returns a SettingsApplicationsGetConfigurationGetResponse_emailable when successful
func (m *SettingsApplicationsGetConfigurationGetResponse) GetEmail() SettingsApplicationsGetConfigurationGetResponse_emailable {
	return m.email
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SettingsApplicationsGetConfigurationGetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["affiliate"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSettingsApplicationsGetConfigurationGetResponse_affiliateFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAffiliate(val.(SettingsApplicationsGetConfigurationGetResponse_affiliateable))
		}
		return nil
	}
	res["application"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSettingsApplicationsGetConfigurationGetResponse_applicationFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetApplication(val.(SettingsApplicationsGetConfigurationGetResponse_applicationable))
		}
		return nil
	}
	res["appointment"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSettingsApplicationsGetConfigurationGetResponse_appointmentFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAppointment(val.(SettingsApplicationsGetConfigurationGetResponse_appointmentable))
		}
		return nil
	}
	res["contact"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSettingsApplicationsGetConfigurationGetResponse_contactFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetContact(val.(SettingsApplicationsGetConfigurationGetResponse_contactable))
		}
		return nil
	}
	res["ecommerce"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSettingsApplicationsGetConfigurationGetResponse_ecommerceFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetEcommerce(val.(SettingsApplicationsGetConfigurationGetResponse_ecommerceable))
		}
		return nil
	}
	res["email"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSettingsApplicationsGetConfigurationGetResponse_emailFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetEmail(val.(SettingsApplicationsGetConfigurationGetResponse_emailable))
		}
		return nil
	}
	res["forms"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSettingsApplicationsGetConfigurationGetResponse_formsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetForms(val.(SettingsApplicationsGetConfigurationGetResponse_formsable))
		}
		return nil
	}
	res["fulfillment"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSettingsApplicationsGetConfigurationGetResponse_fulfillmentFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFulfillment(val.(SettingsApplicationsGetConfigurationGetResponse_fulfillmentable))
		}
		return nil
	}
	res["invoice"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSettingsApplicationsGetConfigurationGetResponse_invoiceFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetInvoice(val.(SettingsApplicationsGetConfigurationGetResponse_invoiceable))
		}
		return nil
	}
	res["note"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSettingsApplicationsGetConfigurationGetResponse_noteFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNote(val.(SettingsApplicationsGetConfigurationGetResponse_noteable))
		}
		return nil
	}
	res["opportunity"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSettingsApplicationsGetConfigurationGetResponse_opportunityFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetOpportunity(val.(SettingsApplicationsGetConfigurationGetResponse_opportunityable))
		}
		return nil
	}
	res["task"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSettingsApplicationsGetConfigurationGetResponse_taskFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTask(val.(SettingsApplicationsGetConfigurationGetResponse_taskable))
		}
		return nil
	}
	res["template"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSettingsApplicationsGetConfigurationGetResponse_templateFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTemplate(val.(SettingsApplicationsGetConfigurationGetResponse_templateable))
		}
		return nil
	}
	return res
}

// GetForms gets the forms property value. The forms property
// returns a SettingsApplicationsGetConfigurationGetResponse_formsable when successful
func (m *SettingsApplicationsGetConfigurationGetResponse) GetForms() SettingsApplicationsGetConfigurationGetResponse_formsable {
	return m.forms
}

// GetFulfillment gets the fulfillment property value. The fulfillment property
// returns a SettingsApplicationsGetConfigurationGetResponse_fulfillmentable when successful
func (m *SettingsApplicationsGetConfigurationGetResponse) GetFulfillment() SettingsApplicationsGetConfigurationGetResponse_fulfillmentable {
	return m.fulfillment
}

// GetInvoice gets the invoice property value. The invoice property
// returns a SettingsApplicationsGetConfigurationGetResponse_invoiceable when successful
func (m *SettingsApplicationsGetConfigurationGetResponse) GetInvoice() SettingsApplicationsGetConfigurationGetResponse_invoiceable {
	return m.invoice
}

// GetNote gets the note property value. The note property
// returns a SettingsApplicationsGetConfigurationGetResponse_noteable when successful
func (m *SettingsApplicationsGetConfigurationGetResponse) GetNote() SettingsApplicationsGetConfigurationGetResponse_noteable {
	return m.note
}

// GetOpportunity gets the opportunity property value. The opportunity property
// returns a SettingsApplicationsGetConfigurationGetResponse_opportunityable when successful
func (m *SettingsApplicationsGetConfigurationGetResponse) GetOpportunity() SettingsApplicationsGetConfigurationGetResponse_opportunityable {
	return m.opportunity
}

// GetTask gets the task property value. The task property
// returns a SettingsApplicationsGetConfigurationGetResponse_taskable when successful
func (m *SettingsApplicationsGetConfigurationGetResponse) GetTask() SettingsApplicationsGetConfigurationGetResponse_taskable {
	return m.task
}

// GetTemplate gets the template property value. The template property
// returns a SettingsApplicationsGetConfigurationGetResponse_templateable when successful
func (m *SettingsApplicationsGetConfigurationGetResponse) GetTemplate() SettingsApplicationsGetConfigurationGetResponse_templateable {
	return m.template
}

// Serialize serializes information the current object
func (m *SettingsApplicationsGetConfigurationGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *SettingsApplicationsGetConfigurationGetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAffiliate sets the affiliate property value. The affiliate property
func (m *SettingsApplicationsGetConfigurationGetResponse) SetAffiliate(value SettingsApplicationsGetConfigurationGetResponse_affiliateable) {
	m.affiliate = value
}

// SetApplication sets the application property value. The application property
func (m *SettingsApplicationsGetConfigurationGetResponse) SetApplication(value SettingsApplicationsGetConfigurationGetResponse_applicationable) {
	m.application = value
}

// SetAppointment sets the appointment property value. The appointment property
func (m *SettingsApplicationsGetConfigurationGetResponse) SetAppointment(value SettingsApplicationsGetConfigurationGetResponse_appointmentable) {
	m.appointment = value
}

// SetContact sets the contact property value. The contact property
func (m *SettingsApplicationsGetConfigurationGetResponse) SetContact(value SettingsApplicationsGetConfigurationGetResponse_contactable) {
	m.contact = value
}

// SetEcommerce sets the ecommerce property value. The ecommerce property
func (m *SettingsApplicationsGetConfigurationGetResponse) SetEcommerce(value SettingsApplicationsGetConfigurationGetResponse_ecommerceable) {
	m.ecommerce = value
}

// SetEmail sets the email property value. The email property
func (m *SettingsApplicationsGetConfigurationGetResponse) SetEmail(value SettingsApplicationsGetConfigurationGetResponse_emailable) {
	m.email = value
}

// SetForms sets the forms property value. The forms property
func (m *SettingsApplicationsGetConfigurationGetResponse) SetForms(value SettingsApplicationsGetConfigurationGetResponse_formsable) {
	m.forms = value
}

// SetFulfillment sets the fulfillment property value. The fulfillment property
func (m *SettingsApplicationsGetConfigurationGetResponse) SetFulfillment(value SettingsApplicationsGetConfigurationGetResponse_fulfillmentable) {
	m.fulfillment = value
}

// SetInvoice sets the invoice property value. The invoice property
func (m *SettingsApplicationsGetConfigurationGetResponse) SetInvoice(value SettingsApplicationsGetConfigurationGetResponse_invoiceable) {
	m.invoice = value
}

// SetNote sets the note property value. The note property
func (m *SettingsApplicationsGetConfigurationGetResponse) SetNote(value SettingsApplicationsGetConfigurationGetResponse_noteable) {
	m.note = value
}

// SetOpportunity sets the opportunity property value. The opportunity property
func (m *SettingsApplicationsGetConfigurationGetResponse) SetOpportunity(value SettingsApplicationsGetConfigurationGetResponse_opportunityable) {
	m.opportunity = value
}

// SetTask sets the task property value. The task property
func (m *SettingsApplicationsGetConfigurationGetResponse) SetTask(value SettingsApplicationsGetConfigurationGetResponse_taskable) {
	m.task = value
}

// SetTemplate sets the template property value. The template property
func (m *SettingsApplicationsGetConfigurationGetResponse) SetTemplate(value SettingsApplicationsGetConfigurationGetResponse_templateable) {
	m.template = value
}

type SettingsApplicationsGetConfigurationGetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAffiliate() SettingsApplicationsGetConfigurationGetResponse_affiliateable
	GetApplication() SettingsApplicationsGetConfigurationGetResponse_applicationable
	GetAppointment() SettingsApplicationsGetConfigurationGetResponse_appointmentable
	GetContact() SettingsApplicationsGetConfigurationGetResponse_contactable
	GetEcommerce() SettingsApplicationsGetConfigurationGetResponse_ecommerceable
	GetEmail() SettingsApplicationsGetConfigurationGetResponse_emailable
	GetForms() SettingsApplicationsGetConfigurationGetResponse_formsable
	GetFulfillment() SettingsApplicationsGetConfigurationGetResponse_fulfillmentable
	GetInvoice() SettingsApplicationsGetConfigurationGetResponse_invoiceable
	GetNote() SettingsApplicationsGetConfigurationGetResponse_noteable
	GetOpportunity() SettingsApplicationsGetConfigurationGetResponse_opportunityable
	GetTask() SettingsApplicationsGetConfigurationGetResponse_taskable
	GetTemplate() SettingsApplicationsGetConfigurationGetResponse_templateable
	SetAffiliate(value SettingsApplicationsGetConfigurationGetResponse_affiliateable)
	SetApplication(value SettingsApplicationsGetConfigurationGetResponse_applicationable)
	SetAppointment(value SettingsApplicationsGetConfigurationGetResponse_appointmentable)
	SetContact(value SettingsApplicationsGetConfigurationGetResponse_contactable)
	SetEcommerce(value SettingsApplicationsGetConfigurationGetResponse_ecommerceable)
	SetEmail(value SettingsApplicationsGetConfigurationGetResponse_emailable)
	SetForms(value SettingsApplicationsGetConfigurationGetResponse_formsable)
	SetFulfillment(value SettingsApplicationsGetConfigurationGetResponse_fulfillmentable)
	SetInvoice(value SettingsApplicationsGetConfigurationGetResponse_invoiceable)
	SetNote(value SettingsApplicationsGetConfigurationGetResponse_noteable)
	SetOpportunity(value SettingsApplicationsGetConfigurationGetResponse_opportunityable)
	SetTask(value SettingsApplicationsGetConfigurationGetResponse_taskable)
	SetTemplate(value SettingsApplicationsGetConfigurationGetResponse_templateable)
}
