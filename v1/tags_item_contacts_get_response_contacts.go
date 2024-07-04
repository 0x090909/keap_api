package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
)

type TagsItemContactsGetResponse_contacts struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The contact property
	contact TagsItemContactsGetResponse_contacts_contactable
	// The date_applied property
	date_applied *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}

// NewTagsItemContactsGetResponse_contacts instantiates a new TagsItemContactsGetResponse_contacts and sets the default values.
func NewTagsItemContactsGetResponse_contacts() *TagsItemContactsGetResponse_contacts {
	m := &TagsItemContactsGetResponse_contacts{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateTagsItemContactsGetResponse_contactsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTagsItemContactsGetResponse_contactsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTagsItemContactsGetResponse_contacts(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TagsItemContactsGetResponse_contacts) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetContact gets the contact property value. The contact property
// returns a TagsItemContactsGetResponse_contacts_contactable when successful
func (m *TagsItemContactsGetResponse_contacts) GetContact() TagsItemContactsGetResponse_contacts_contactable {
	return m.contact
}

// GetDateApplied gets the date_applied property value. The date_applied property
// returns a *Time when successful
func (m *TagsItemContactsGetResponse_contacts) GetDateApplied() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.date_applied
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TagsItemContactsGetResponse_contacts) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["contact"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateTagsItemContactsGetResponse_contacts_contactFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetContact(val.(TagsItemContactsGetResponse_contacts_contactable))
		}
		return nil
	}
	res["date_applied"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetTimeValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDateApplied(val)
		}
		return nil
	}
	return res
}

// Serialize serializes information the current object
func (m *TagsItemContactsGetResponse_contacts) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("contact", m.GetContact())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteTimeValue("date_applied", m.GetDateApplied())
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
func (m *TagsItemContactsGetResponse_contacts) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetContact sets the contact property value. The contact property
func (m *TagsItemContactsGetResponse_contacts) SetContact(value TagsItemContactsGetResponse_contacts_contactable) {
	m.contact = value
}

// SetDateApplied sets the date_applied property value. The date_applied property
func (m *TagsItemContactsGetResponse_contacts) SetDateApplied(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	m.date_applied = value
}

type TagsItemContactsGetResponse_contactsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetContact() TagsItemContactsGetResponse_contacts_contactable
	GetDateApplied() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	SetContact(value TagsItemContactsGetResponse_contacts_contactable)
	SetDateApplied(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
}
