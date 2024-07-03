package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EmailsPostResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The clicked_time property
	clicked_time *string
	// The contact_id property
	contact_id *string
	// The headers property
	headers *string
	// Base64 encoded HTML
	html_content *string
	// The id property
	id *string
	// The opened_time property
	opened_time *string
	// The original_provider_id property
	original_provider_id *string
	// Base64 encoded plain text
	plain_content *string
	// The provider_source_id property
	provider_source_id *string
	// The received_time property
	received_time *string
	// The sent_from_address property
	sent_from_address *string
	// The sent_from_reply_address property
	sent_from_reply_address *string
	// The sent_time property
	sent_time *string
	// The sent_to_address property
	sent_to_address *string
	// The sent_to_bcc_address_list property
	sent_to_bcc_address_list []string
	// The sent_to_cc_address_list property
	sent_to_cc_address_list []string
	// The subject property
	subject *string
}

// NewEmailsPostResponse instantiates a new EmailsPostResponse and sets the default values.
func NewEmailsPostResponse() *EmailsPostResponse {
	m := &EmailsPostResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateEmailsPostResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEmailsPostResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewEmailsPostResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EmailsPostResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetClickedTime gets the clicked_time property value. The clicked_time property
// returns a *string when successful
func (m *EmailsPostResponse) GetClickedTime() *string {
	return m.clicked_time
}

// GetContactId gets the contact_id property value. The contact_id property
// returns a *string when successful
func (m *EmailsPostResponse) GetContactId() *string {
	return m.contact_id
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EmailsPostResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["clicked_time"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetClickedTime(val)
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
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetId(val)
		}
		return nil
	}
	res["opened_time"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetOpenedTime(val)
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
	res["provider_source_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetProviderSourceId(val)
		}
		return nil
	}
	res["received_time"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetReceivedTime(val)
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
	res["sent_time"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSentTime(val)
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
	res["sent_to_bcc_address_list"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfPrimitiveValues("string")
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]string, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = *(v.(*string))
				}
			}
			m.SetSentToBccAddressList(res)
		}
		return nil
	}
	res["sent_to_cc_address_list"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfPrimitiveValues("string")
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]string, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = *(v.(*string))
				}
			}
			m.SetSentToCcAddressList(res)
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
func (m *EmailsPostResponse) GetHeaders() *string {
	return m.headers
}

// GetHtmlContent gets the html_content property value. Base64 encoded HTML
// returns a *string when successful
func (m *EmailsPostResponse) GetHtmlContent() *string {
	return m.html_content
}

// GetId gets the id property value. The id property
// returns a *string when successful
func (m *EmailsPostResponse) GetId() *string {
	return m.id
}

// GetOpenedTime gets the opened_time property value. The opened_time property
// returns a *string when successful
func (m *EmailsPostResponse) GetOpenedTime() *string {
	return m.opened_time
}

// GetOriginalProviderId gets the original_provider_id property value. The original_provider_id property
// returns a *string when successful
func (m *EmailsPostResponse) GetOriginalProviderId() *string {
	return m.original_provider_id
}

// GetPlainContent gets the plain_content property value. Base64 encoded plain text
// returns a *string when successful
func (m *EmailsPostResponse) GetPlainContent() *string {
	return m.plain_content
}

// GetProviderSourceId gets the provider_source_id property value. The provider_source_id property
// returns a *string when successful
func (m *EmailsPostResponse) GetProviderSourceId() *string {
	return m.provider_source_id
}

// GetReceivedTime gets the received_time property value. The received_time property
// returns a *string when successful
func (m *EmailsPostResponse) GetReceivedTime() *string {
	return m.received_time
}

// GetSentFromAddress gets the sent_from_address property value. The sent_from_address property
// returns a *string when successful
func (m *EmailsPostResponse) GetSentFromAddress() *string {
	return m.sent_from_address
}

// GetSentFromReplyAddress gets the sent_from_reply_address property value. The sent_from_reply_address property
// returns a *string when successful
func (m *EmailsPostResponse) GetSentFromReplyAddress() *string {
	return m.sent_from_reply_address
}

// GetSentTime gets the sent_time property value. The sent_time property
// returns a *string when successful
func (m *EmailsPostResponse) GetSentTime() *string {
	return m.sent_time
}

// GetSentToAddress gets the sent_to_address property value. The sent_to_address property
// returns a *string when successful
func (m *EmailsPostResponse) GetSentToAddress() *string {
	return m.sent_to_address
}

// GetSentToBccAddressList gets the sent_to_bcc_address_list property value. The sent_to_bcc_address_list property
// returns a []string when successful
func (m *EmailsPostResponse) GetSentToBccAddressList() []string {
	return m.sent_to_bcc_address_list
}

// GetSentToCcAddressList gets the sent_to_cc_address_list property value. The sent_to_cc_address_list property
// returns a []string when successful
func (m *EmailsPostResponse) GetSentToCcAddressList() []string {
	return m.sent_to_cc_address_list
}

// GetSubject gets the subject property value. The subject property
// returns a *string when successful
func (m *EmailsPostResponse) GetSubject() *string {
	return m.subject
}

// Serialize serializes information the current object
func (m *EmailsPostResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("clicked_time", m.GetClickedTime())
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
		err := writer.WriteStringValue("id", m.GetId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("opened_time", m.GetOpenedTime())
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
		err := writer.WriteStringValue("provider_source_id", m.GetProviderSourceId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("received_time", m.GetReceivedTime())
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
		err := writer.WriteStringValue("sent_time", m.GetSentTime())
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
	if m.GetSentToBccAddressList() != nil {
		err := writer.WriteCollectionOfStringValues("sent_to_bcc_address_list", m.GetSentToBccAddressList())
		if err != nil {
			return err
		}
	}
	if m.GetSentToCcAddressList() != nil {
		err := writer.WriteCollectionOfStringValues("sent_to_cc_address_list", m.GetSentToCcAddressList())
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
func (m *EmailsPostResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetClickedTime sets the clicked_time property value. The clicked_time property
func (m *EmailsPostResponse) SetClickedTime(value *string) {
	m.clicked_time = value
}

// SetContactId sets the contact_id property value. The contact_id property
func (m *EmailsPostResponse) SetContactId(value *string) {
	m.contact_id = value
}

// SetHeaders sets the headers property value. The headers property
func (m *EmailsPostResponse) SetHeaders(value *string) {
	m.headers = value
}

// SetHtmlContent sets the html_content property value. Base64 encoded HTML
func (m *EmailsPostResponse) SetHtmlContent(value *string) {
	m.html_content = value
}

// SetId sets the id property value. The id property
func (m *EmailsPostResponse) SetId(value *string) {
	m.id = value
}

// SetOpenedTime sets the opened_time property value. The opened_time property
func (m *EmailsPostResponse) SetOpenedTime(value *string) {
	m.opened_time = value
}

// SetOriginalProviderId sets the original_provider_id property value. The original_provider_id property
func (m *EmailsPostResponse) SetOriginalProviderId(value *string) {
	m.original_provider_id = value
}

// SetPlainContent sets the plain_content property value. Base64 encoded plain text
func (m *EmailsPostResponse) SetPlainContent(value *string) {
	m.plain_content = value
}

// SetProviderSourceId sets the provider_source_id property value. The provider_source_id property
func (m *EmailsPostResponse) SetProviderSourceId(value *string) {
	m.provider_source_id = value
}

// SetReceivedTime sets the received_time property value. The received_time property
func (m *EmailsPostResponse) SetReceivedTime(value *string) {
	m.received_time = value
}

// SetSentFromAddress sets the sent_from_address property value. The sent_from_address property
func (m *EmailsPostResponse) SetSentFromAddress(value *string) {
	m.sent_from_address = value
}

// SetSentFromReplyAddress sets the sent_from_reply_address property value. The sent_from_reply_address property
func (m *EmailsPostResponse) SetSentFromReplyAddress(value *string) {
	m.sent_from_reply_address = value
}

// SetSentTime sets the sent_time property value. The sent_time property
func (m *EmailsPostResponse) SetSentTime(value *string) {
	m.sent_time = value
}

// SetSentToAddress sets the sent_to_address property value. The sent_to_address property
func (m *EmailsPostResponse) SetSentToAddress(value *string) {
	m.sent_to_address = value
}

// SetSentToBccAddressList sets the sent_to_bcc_address_list property value. The sent_to_bcc_address_list property
func (m *EmailsPostResponse) SetSentToBccAddressList(value []string) {
	m.sent_to_bcc_address_list = value
}

// SetSentToCcAddressList sets the sent_to_cc_address_list property value. The sent_to_cc_address_list property
func (m *EmailsPostResponse) SetSentToCcAddressList(value []string) {
	m.sent_to_cc_address_list = value
}

// SetSubject sets the subject property value. The subject property
func (m *EmailsPostResponse) SetSubject(value *string) {
	m.subject = value
}

type EmailsPostResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetClickedTime() *string
	GetContactId() *string
	GetHeaders() *string
	GetHtmlContent() *string
	GetId() *string
	GetOpenedTime() *string
	GetOriginalProviderId() *string
	GetPlainContent() *string
	GetProviderSourceId() *string
	GetReceivedTime() *string
	GetSentFromAddress() *string
	GetSentFromReplyAddress() *string
	GetSentTime() *string
	GetSentToAddress() *string
	GetSentToBccAddressList() []string
	GetSentToCcAddressList() []string
	GetSubject() *string
	SetClickedTime(value *string)
	SetContactId(value *string)
	SetHeaders(value *string)
	SetHtmlContent(value *string)
	SetId(value *string)
	SetOpenedTime(value *string)
	SetOriginalProviderId(value *string)
	SetPlainContent(value *string)
	SetProviderSourceId(value *string)
	SetReceivedTime(value *string)
	SetSentFromAddress(value *string)
	SetSentFromReplyAddress(value *string)
	SetSentTime(value *string)
	SetSentToAddress(value *string)
	SetSentToBccAddressList(value []string)
	SetSentToCcAddressList(value []string)
	SetSubject(value *string)
}
