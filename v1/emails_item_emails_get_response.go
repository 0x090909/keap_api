package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
)

type EmailsItemEmailsGetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The clicked_date property
	clicked_date *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	// The contact_id property
	contact_id *int64
	// The headers property
	headers *string
	// Base64 encoded HTML
	html_content *string
	// The id property
	id *int64
	// The opened_date property
	opened_date *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	// The original_provider_id property
	original_provider_id *string
	// Base64 encoded plain text
	plain_content *string
	// The received_date property
	received_date *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	// The sent_date property
	sent_date *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	// The sent_from_address property
	sent_from_address *string
	// The sent_from_reply_address property
	sent_from_reply_address *string
	// The sent_to_address property
	sent_to_address *string
	// The sent_to_bcc_addresses property
	sent_to_bcc_addresses *string
	// The sent_to_cc_addresses property
	sent_to_cc_addresses *string
	// The subject property
	subject *string
}

// NewEmailsItemEmailsGetResponse instantiates a new EmailsItemEmailsGetResponse and sets the default values.
func NewEmailsItemEmailsGetResponse() *EmailsItemEmailsGetResponse {
	m := &EmailsItemEmailsGetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateEmailsItemEmailsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEmailsItemEmailsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewEmailsItemEmailsGetResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EmailsItemEmailsGetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetClickedDate gets the clicked_date property value. The clicked_date property
// returns a *Time when successful
func (m *EmailsItemEmailsGetResponse) GetClickedDate() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.clicked_date
}

// GetContactId gets the contact_id property value. The contact_id property
// returns a *int64 when successful
func (m *EmailsItemEmailsGetResponse) GetContactId() *int64 {
	return m.contact_id
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EmailsItemEmailsGetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["clicked_date"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetTimeValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetClickedDate(val)
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
	res["headers"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetHeaders(val)
		}
		return nil
	}
	res["html_content"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetHtmlContent(val)
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
	res["opened_date"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetTimeValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetOpenedDate(val)
		}
		return nil
	}
	res["original_provider_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetOriginalProviderId(val)
		}
		return nil
	}
	res["plain_content"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPlainContent(val)
		}
		return nil
	}
	res["received_date"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetTimeValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetReceivedDate(val)
		}
		return nil
	}
	res["sent_date"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetTimeValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSentDate(val)
		}
		return nil
	}
	res["sent_from_address"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSentFromAddress(val)
		}
		return nil
	}
	res["sent_from_reply_address"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSentFromReplyAddress(val)
		}
		return nil
	}
	res["sent_to_address"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSentToAddress(val)
		}
		return nil
	}
	res["sent_to_bcc_addresses"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSentToBccAddresses(val)
		}
		return nil
	}
	res["sent_to_cc_addresses"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSentToCcAddresses(val)
		}
		return nil
	}
	res["subject"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSubject(val)
		}
		return nil
	}
	return res
}

// GetHeaders gets the headers property value. The headers property
// returns a *string when successful
func (m *EmailsItemEmailsGetResponse) GetHeaders() *string {
	return m.headers
}

// GetHtmlContent gets the html_content property value. Base64 encoded HTML
// returns a *string when successful
func (m *EmailsItemEmailsGetResponse) GetHtmlContent() *string {
	return m.html_content
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *EmailsItemEmailsGetResponse) GetId() *int64 {
	return m.id
}

// GetOpenedDate gets the opened_date property value. The opened_date property
// returns a *Time when successful
func (m *EmailsItemEmailsGetResponse) GetOpenedDate() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.opened_date
}

// GetOriginalProviderId gets the original_provider_id property value. The original_provider_id property
// returns a *string when successful
func (m *EmailsItemEmailsGetResponse) GetOriginalProviderId() *string {
	return m.original_provider_id
}

// GetPlainContent gets the plain_content property value. Base64 encoded plain text
// returns a *string when successful
func (m *EmailsItemEmailsGetResponse) GetPlainContent() *string {
	return m.plain_content
}

// GetReceivedDate gets the received_date property value. The received_date property
// returns a *Time when successful
func (m *EmailsItemEmailsGetResponse) GetReceivedDate() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.received_date
}

// GetSentDate gets the sent_date property value. The sent_date property
// returns a *Time when successful
func (m *EmailsItemEmailsGetResponse) GetSentDate() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.sent_date
}

// GetSentFromAddress gets the sent_from_address property value. The sent_from_address property
// returns a *string when successful
func (m *EmailsItemEmailsGetResponse) GetSentFromAddress() *string {
	return m.sent_from_address
}

// GetSentFromReplyAddress gets the sent_from_reply_address property value. The sent_from_reply_address property
// returns a *string when successful
func (m *EmailsItemEmailsGetResponse) GetSentFromReplyAddress() *string {
	return m.sent_from_reply_address
}

// GetSentToAddress gets the sent_to_address property value. The sent_to_address property
// returns a *string when successful
func (m *EmailsItemEmailsGetResponse) GetSentToAddress() *string {
	return m.sent_to_address
}

// GetSentToBccAddresses gets the sent_to_bcc_addresses property value. The sent_to_bcc_addresses property
// returns a *string when successful
func (m *EmailsItemEmailsGetResponse) GetSentToBccAddresses() *string {
	return m.sent_to_bcc_addresses
}

// GetSentToCcAddresses gets the sent_to_cc_addresses property value. The sent_to_cc_addresses property
// returns a *string when successful
func (m *EmailsItemEmailsGetResponse) GetSentToCcAddresses() *string {
	return m.sent_to_cc_addresses
}

// GetSubject gets the subject property value. The subject property
// returns a *string when successful
func (m *EmailsItemEmailsGetResponse) GetSubject() *string {
	return m.subject
}

// Serialize serializes information the current object
func (m *EmailsItemEmailsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteTimeValue("clicked_date", m.GetClickedDate())
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
		err := writer.WriteStringValue("headers", m.GetHeaders())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("html_content", m.GetHtmlContent())
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
		err := writer.WriteTimeValue("opened_date", m.GetOpenedDate())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("original_provider_id", m.GetOriginalProviderId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("plain_content", m.GetPlainContent())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteTimeValue("received_date", m.GetReceivedDate())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteTimeValue("sent_date", m.GetSentDate())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("sent_from_address", m.GetSentFromAddress())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("sent_from_reply_address", m.GetSentFromReplyAddress())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("sent_to_address", m.GetSentToAddress())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("sent_to_bcc_addresses", m.GetSentToBccAddresses())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("sent_to_cc_addresses", m.GetSentToCcAddresses())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("subject", m.GetSubject())
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
func (m *EmailsItemEmailsGetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetClickedDate sets the clicked_date property value. The clicked_date property
func (m *EmailsItemEmailsGetResponse) SetClickedDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	m.clicked_date = value
}

// SetContactId sets the contact_id property value. The contact_id property
func (m *EmailsItemEmailsGetResponse) SetContactId(value *int64) {
	m.contact_id = value
}

// SetHeaders sets the headers property value. The headers property
func (m *EmailsItemEmailsGetResponse) SetHeaders(value *string) {
	m.headers = value
}

// SetHtmlContent sets the html_content property value. Base64 encoded HTML
func (m *EmailsItemEmailsGetResponse) SetHtmlContent(value *string) {
	m.html_content = value
}

// SetId sets the id property value. The id property
func (m *EmailsItemEmailsGetResponse) SetId(value *int64) {
	m.id = value
}

// SetOpenedDate sets the opened_date property value. The opened_date property
func (m *EmailsItemEmailsGetResponse) SetOpenedDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	m.opened_date = value
}

// SetOriginalProviderId sets the original_provider_id property value. The original_provider_id property
func (m *EmailsItemEmailsGetResponse) SetOriginalProviderId(value *string) {
	m.original_provider_id = value
}

// SetPlainContent sets the plain_content property value. Base64 encoded plain text
func (m *EmailsItemEmailsGetResponse) SetPlainContent(value *string) {
	m.plain_content = value
}

// SetReceivedDate sets the received_date property value. The received_date property
func (m *EmailsItemEmailsGetResponse) SetReceivedDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	m.received_date = value
}

// SetSentDate sets the sent_date property value. The sent_date property
func (m *EmailsItemEmailsGetResponse) SetSentDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	m.sent_date = value
}

// SetSentFromAddress sets the sent_from_address property value. The sent_from_address property
func (m *EmailsItemEmailsGetResponse) SetSentFromAddress(value *string) {
	m.sent_from_address = value
}

// SetSentFromReplyAddress sets the sent_from_reply_address property value. The sent_from_reply_address property
func (m *EmailsItemEmailsGetResponse) SetSentFromReplyAddress(value *string) {
	m.sent_from_reply_address = value
}

// SetSentToAddress sets the sent_to_address property value. The sent_to_address property
func (m *EmailsItemEmailsGetResponse) SetSentToAddress(value *string) {
	m.sent_to_address = value
}

// SetSentToBccAddresses sets the sent_to_bcc_addresses property value. The sent_to_bcc_addresses property
func (m *EmailsItemEmailsGetResponse) SetSentToBccAddresses(value *string) {
	m.sent_to_bcc_addresses = value
}

// SetSentToCcAddresses sets the sent_to_cc_addresses property value. The sent_to_cc_addresses property
func (m *EmailsItemEmailsGetResponse) SetSentToCcAddresses(value *string) {
	m.sent_to_cc_addresses = value
}

// SetSubject sets the subject property value. The subject property
func (m *EmailsItemEmailsGetResponse) SetSubject(value *string) {
	m.subject = value
}

type EmailsItemEmailsGetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetClickedDate() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetContactId() *int64
	GetHeaders() *string
	GetHtmlContent() *string
	GetId() *int64
	GetOpenedDate() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetOriginalProviderId() *string
	GetPlainContent() *string
	GetReceivedDate() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetSentDate() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetSentFromAddress() *string
	GetSentFromReplyAddress() *string
	GetSentToAddress() *string
	GetSentToBccAddresses() *string
	GetSentToCcAddresses() *string
	GetSubject() *string
	SetClickedDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	SetContactId(value *int64)
	SetHeaders(value *string)
	SetHtmlContent(value *string)
	SetId(value *int64)
	SetOpenedDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	SetOriginalProviderId(value *string)
	SetPlainContent(value *string)
	SetReceivedDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	SetSentDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	SetSentFromAddress(value *string)
	SetSentFromReplyAddress(value *string)
	SetSentToAddress(value *string)
	SetSentToBccAddresses(value *string)
	SetSentToCcAddresses(value *string)
	SetSubject(value *string)
}