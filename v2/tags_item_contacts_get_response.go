package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TagsItemContactsGetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The contacts property
	contacts []TagsItemContactsGetResponse_contactsable
	// The next_page_token property
	next_page_token *string
}

// NewTagsItemContactsGetResponse instantiates a new TagsItemContactsGetResponse and sets the default values.
func NewTagsItemContactsGetResponse() *TagsItemContactsGetResponse {
	m := &TagsItemContactsGetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateTagsItemContactsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTagsItemContactsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTagsItemContactsGetResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TagsItemContactsGetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetContacts gets the contacts property value. The contacts property
// returns a []TagsItemContactsGetResponse_contactsable when successful
func (m *TagsItemContactsGetResponse) GetContacts() []TagsItemContactsGetResponse_contactsable {
	return m.contacts
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TagsItemContactsGetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["contacts"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateTagsItemContactsGetResponse_contactsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]TagsItemContactsGetResponse_contactsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(TagsItemContactsGetResponse_contactsable)
				}
			}
			m.SetContacts(res)
		}
		return nil
	}
	res["next_page_token"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNextPageToken(val)
		}
		return nil
	}
	return res
}

// GetNextPageToken gets the next_page_token property value. The next_page_token property
// returns a *string when successful
func (m *TagsItemContactsGetResponse) GetNextPageToken() *string {
	return m.next_page_token
}

// Serialize serializes information the current object
func (m *TagsItemContactsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetContacts() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetContacts()))
		for i, v := range m.GetContacts() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("contacts", cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("next_page_token", m.GetNextPageToken())
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
func (m *TagsItemContactsGetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetContacts sets the contacts property value. The contacts property
func (m *TagsItemContactsGetResponse) SetContacts(value []TagsItemContactsGetResponse_contactsable) {
	m.contacts = value
}

// SetNextPageToken sets the next_page_token property value. The next_page_token property
func (m *TagsItemContactsGetResponse) SetNextPageToken(value *string) {
	m.next_page_token = value
}

type TagsItemContactsGetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetContacts() []TagsItemContactsGetResponse_contactsable
	GetNextPageToken() *string
	SetContacts(value []TagsItemContactsGetResponse_contactsable)
	SetNextPageToken(value *string)
}
