package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
)

type ContactsItemTagsGetResponse_tags struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The date_applied property
	date_applied *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	// The tag property
	tag ContactsItemTagsGetResponse_tags_tagable
}

// NewContactsItemTagsGetResponse_tags instantiates a new ContactsItemTagsGetResponse_tags and sets the default values.
func NewContactsItemTagsGetResponse_tags() *ContactsItemTagsGetResponse_tags {
	m := &ContactsItemTagsGetResponse_tags{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateContactsItemTagsGetResponse_tagsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsItemTagsGetResponse_tagsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsItemTagsGetResponse_tags(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContactsItemTagsGetResponse_tags) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetDateApplied gets the date_applied property value. The date_applied property
// returns a *Time when successful
func (m *ContactsItemTagsGetResponse_tags) GetDateApplied() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.date_applied
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContactsItemTagsGetResponse_tags) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
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
	res["tag"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateContactsItemTagsGetResponse_tags_tagFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTag(val.(ContactsItemTagsGetResponse_tags_tagable))
		}
		return nil
	}
	return res
}

// GetTag gets the tag property value. The tag property
// returns a ContactsItemTagsGetResponse_tags_tagable when successful
func (m *ContactsItemTagsGetResponse_tags) GetTag() ContactsItemTagsGetResponse_tags_tagable {
	return m.tag
}

// Serialize serializes information the current object
func (m *ContactsItemTagsGetResponse_tags) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteTimeValue("date_applied", m.GetDateApplied())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("tag", m.GetTag())
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
func (m *ContactsItemTagsGetResponse_tags) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetDateApplied sets the date_applied property value. The date_applied property
func (m *ContactsItemTagsGetResponse_tags) SetDateApplied(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	m.date_applied = value
}

// SetTag sets the tag property value. The tag property
func (m *ContactsItemTagsGetResponse_tags) SetTag(value ContactsItemTagsGetResponse_tags_tagable) {
	m.tag = value
}

type ContactsItemTagsGetResponse_tagsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetDateApplied() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetTag() ContactsItemTagsGetResponse_tags_tagable
	SetDateApplied(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	SetTag(value ContactsItemTagsGetResponse_tags_tagable)
}
