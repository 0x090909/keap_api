package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ContactsLinksTypesGetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The contact_link_types property
	contact_link_types []ContactsLinksTypesGetResponse_contact_link_typesable
	// The next_page_token property
	next_page_token *string
}

// NewContactsLinksTypesGetResponse instantiates a new ContactsLinksTypesGetResponse and sets the default values.
func NewContactsLinksTypesGetResponse() *ContactsLinksTypesGetResponse {
	m := &ContactsLinksTypesGetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateContactsLinksTypesGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsLinksTypesGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsLinksTypesGetResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContactsLinksTypesGetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetContactLinkTypes gets the contact_link_types property value. The contact_link_types property
// returns a []ContactsLinksTypesGetResponse_contact_link_typesable when successful
func (m *ContactsLinksTypesGetResponse) GetContactLinkTypes() []ContactsLinksTypesGetResponse_contact_link_typesable {
	return m.contact_link_types
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContactsLinksTypesGetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["contact_link_types"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateContactsLinksTypesGetResponse_contact_link_typesFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ContactsLinksTypesGetResponse_contact_link_typesable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ContactsLinksTypesGetResponse_contact_link_typesable)
				}
			}
			m.SetContactLinkTypes(res)
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
func (m *ContactsLinksTypesGetResponse) GetNextPageToken() *string {
	return m.next_page_token
}

// Serialize serializes information the current object
func (m *ContactsLinksTypesGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetContactLinkTypes() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetContactLinkTypes()))
		for i, v := range m.GetContactLinkTypes() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("contact_link_types", cast)
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
func (m *ContactsLinksTypesGetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetContactLinkTypes sets the contact_link_types property value. The contact_link_types property
func (m *ContactsLinksTypesGetResponse) SetContactLinkTypes(value []ContactsLinksTypesGetResponse_contact_link_typesable) {
	m.contact_link_types = value
}

// SetNextPageToken sets the next_page_token property value. The next_page_token property
func (m *ContactsLinksTypesGetResponse) SetNextPageToken(value *string) {
	m.next_page_token = value
}

type ContactsLinksTypesGetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetContactLinkTypes() []ContactsLinksTypesGetResponse_contact_link_typesable
	GetNextPageToken() *string
	SetContactLinkTypes(value []ContactsLinksTypesGetResponse_contact_link_typesable)
	SetNextPageToken(value *string)
}
