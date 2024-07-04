package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ContactsItemEmailsGetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The count property
	count *int32
	// The emails property
	emails []ContactsItemEmailsGetResponse_emailsable
	// The next property
	next *string
	// The previous property
	previous *string
}

// NewContactsItemEmailsGetResponse instantiates a new ContactsItemEmailsGetResponse and sets the default values.
func NewContactsItemEmailsGetResponse() *ContactsItemEmailsGetResponse {
	m := &ContactsItemEmailsGetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateContactsItemEmailsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsItemEmailsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsItemEmailsGetResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContactsItemEmailsGetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCount gets the count property value. The count property
// returns a *int32 when successful
func (m *ContactsItemEmailsGetResponse) GetCount() *int32 {
	return m.count
}

// GetEmails gets the emails property value. The emails property
// returns a []ContactsItemEmailsGetResponse_emailsable when successful
func (m *ContactsItemEmailsGetResponse) GetEmails() []ContactsItemEmailsGetResponse_emailsable {
	return m.emails
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContactsItemEmailsGetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["count"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCount(val)
		}
		return nil
	}
	res["emails"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateContactsItemEmailsGetResponse_emailsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ContactsItemEmailsGetResponse_emailsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ContactsItemEmailsGetResponse_emailsable)
				}
			}
			m.SetEmails(res)
		}
		return nil
	}
	res["next"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNext(val)
		}
		return nil
	}
	res["previous"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPrevious(val)
		}
		return nil
	}
	return res
}

// GetNext gets the next property value. The next property
// returns a *string when successful
func (m *ContactsItemEmailsGetResponse) GetNext() *string {
	return m.next
}

// GetPrevious gets the previous property value. The previous property
// returns a *string when successful
func (m *ContactsItemEmailsGetResponse) GetPrevious() *string {
	return m.previous
}

// Serialize serializes information the current object
func (m *ContactsItemEmailsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt32Value("count", m.GetCount())
		if err != nil {
			return err
		}
	}
	if m.GetEmails() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEmails()))
		for i, v := range m.GetEmails() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("emails", cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("next", m.GetNext())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("previous", m.GetPrevious())
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
func (m *ContactsItemEmailsGetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCount sets the count property value. The count property
func (m *ContactsItemEmailsGetResponse) SetCount(value *int32) {
	m.count = value
}

// SetEmails sets the emails property value. The emails property
func (m *ContactsItemEmailsGetResponse) SetEmails(value []ContactsItemEmailsGetResponse_emailsable) {
	m.emails = value
}

// SetNext sets the next property value. The next property
func (m *ContactsItemEmailsGetResponse) SetNext(value *string) {
	m.next = value
}

// SetPrevious sets the previous property value. The previous property
func (m *ContactsItemEmailsGetResponse) SetPrevious(value *string) {
	m.previous = value
}

type ContactsItemEmailsGetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCount() *int32
	GetEmails() []ContactsItemEmailsGetResponse_emailsable
	GetNext() *string
	GetPrevious() *string
	SetCount(value *int32)
	SetEmails(value []ContactsItemEmailsGetResponse_emailsable)
	SetNext(value *string)
	SetPrevious(value *string)
}
