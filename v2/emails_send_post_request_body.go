package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EmailsSendPostRequestBody struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// Email field of each Contact record to address the email to, such as 'Email', 'EmailAddress2', 'EmailAddress3' or '_CustomFieldName', defaulting to the contact's primary email
	address_field *string
	// Attachments to be sent with each copy of the email, maximum of 10 with size of 1MB each
	attachments []EmailsSendPostRequestBody_attachmentsable
	// An array of Contact Ids to receive the email
	contacts []string
	// The HTML-formatted content of the email, encoded in Base64
	html_content *string
	// The plain-text content of the email, encoded in Base64
	plain_content *string
	// The subject line of the email
	subject *string
	// The infusionsoft user to send the email on behalf of
	user_id *string
}

// NewEmailsSendPostRequestBody instantiates a new EmailsSendPostRequestBody and sets the default values.
func NewEmailsSendPostRequestBody() *EmailsSendPostRequestBody {
	m := &EmailsSendPostRequestBody{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateEmailsSendPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEmailsSendPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewEmailsSendPostRequestBody(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *EmailsSendPostRequestBody) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAddressField gets the address_field property value. Email field of each Contact record to address the email to, such as 'Email', 'EmailAddress2', 'EmailAddress3' or '_CustomFieldName', defaulting to the contact's primary email
// returns a *string when successful
func (m *EmailsSendPostRequestBody) GetAddressField() *string {
	return m.address_field
}

// GetAttachments gets the attachments property value. Attachments to be sent with each copy of the email, maximum of 10 with size of 1MB each
// returns a []EmailsSendPostRequestBody_attachmentsable when successful
func (m *EmailsSendPostRequestBody) GetAttachments() []EmailsSendPostRequestBody_attachmentsable {
	return m.attachments
}

// GetContacts gets the contacts property value. An array of Contact Ids to receive the email
// returns a []string when successful
func (m *EmailsSendPostRequestBody) GetContacts() []string {
	return m.contacts
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EmailsSendPostRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["address_field"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAddressField(val)
		}
		return nil
	}
	res["attachments"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateEmailsSendPostRequestBody_attachmentsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]EmailsSendPostRequestBody_attachmentsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(EmailsSendPostRequestBody_attachmentsable)
				}
			}
			m.SetAttachments(res)
		}
		return nil
	}
	res["contacts"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
			m.SetContacts(res)
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
	res["user_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetUserId(val)
		}
		return nil
	}
	return res
}

// GetHtmlContent gets the html_content property value. The HTML-formatted content of the email, encoded in Base64
// returns a *string when successful
func (m *EmailsSendPostRequestBody) GetHtmlContent() *string {
	return m.html_content
}

// GetPlainContent gets the plain_content property value. The plain-text content of the email, encoded in Base64
// returns a *string when successful
func (m *EmailsSendPostRequestBody) GetPlainContent() *string {
	return m.plain_content
}

// GetSubject gets the subject property value. The subject line of the email
// returns a *string when successful
func (m *EmailsSendPostRequestBody) GetSubject() *string {
	return m.subject
}

// GetUserId gets the user_id property value. The infusionsoft user to send the email on behalf of
// returns a *string when successful
func (m *EmailsSendPostRequestBody) GetUserId() *string {
	return m.user_id
}

// Serialize serializes information the current object
func (m *EmailsSendPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("address_field", m.GetAddressField())
		if err != nil {
			return err
		}
	}
	if m.GetAttachments() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAttachments()))
		for i, v := range m.GetAttachments() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("attachments", cast)
		if err != nil {
			return err
		}
	}
	if m.GetContacts() != nil {
		err := writer.WriteCollectionOfStringValues("contacts", m.GetContacts())
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
		err := writer.WriteStringValue("plain_content", m.GetPlainContent())
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
		err := writer.WriteStringValue("user_id", m.GetUserId())
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
func (m *EmailsSendPostRequestBody) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAddressField sets the address_field property value. Email field of each Contact record to address the email to, such as 'Email', 'EmailAddress2', 'EmailAddress3' or '_CustomFieldName', defaulting to the contact's primary email
func (m *EmailsSendPostRequestBody) SetAddressField(value *string) {
	m.address_field = value
}

// SetAttachments sets the attachments property value. Attachments to be sent with each copy of the email, maximum of 10 with size of 1MB each
func (m *EmailsSendPostRequestBody) SetAttachments(value []EmailsSendPostRequestBody_attachmentsable) {
	m.attachments = value
}

// SetContacts sets the contacts property value. An array of Contact Ids to receive the email
func (m *EmailsSendPostRequestBody) SetContacts(value []string) {
	m.contacts = value
}

// SetHtmlContent sets the html_content property value. The HTML-formatted content of the email, encoded in Base64
func (m *EmailsSendPostRequestBody) SetHtmlContent(value *string) {
	m.html_content = value
}

// SetPlainContent sets the plain_content property value. The plain-text content of the email, encoded in Base64
func (m *EmailsSendPostRequestBody) SetPlainContent(value *string) {
	m.plain_content = value
}

// SetSubject sets the subject property value. The subject line of the email
func (m *EmailsSendPostRequestBody) SetSubject(value *string) {
	m.subject = value
}

// SetUserId sets the user_id property value. The infusionsoft user to send the email on behalf of
func (m *EmailsSendPostRequestBody) SetUserId(value *string) {
	m.user_id = value
}

type EmailsSendPostRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAddressField() *string
	GetAttachments() []EmailsSendPostRequestBody_attachmentsable
	GetContacts() []string
	GetHtmlContent() *string
	GetPlainContent() *string
	GetSubject() *string
	GetUserId() *string
	SetAddressField(value *string)
	SetAttachments(value []EmailsSendPostRequestBody_attachmentsable)
	SetContacts(value []string)
	SetHtmlContent(value *string)
	SetPlainContent(value *string)
	SetSubject(value *string)
	SetUserId(value *string)
}
