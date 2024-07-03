package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ContactsItemLinksGetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The links property
	links []ContactsItemLinksGetResponse_linksable
	// The next_page_token property
	next_page_token *string
}

// NewContactsItemLinksGetResponse instantiates a new ContactsItemLinksGetResponse and sets the default values.
func NewContactsItemLinksGetResponse() *ContactsItemLinksGetResponse {
	m := &ContactsItemLinksGetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateContactsItemLinksGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsItemLinksGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsItemLinksGetResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContactsItemLinksGetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContactsItemLinksGetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["links"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateContactsItemLinksGetResponse_linksFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ContactsItemLinksGetResponse_linksable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ContactsItemLinksGetResponse_linksable)
				}
			}
			m.SetLinks(res)
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

// GetLinks gets the links property value. The links property
// returns a []ContactsItemLinksGetResponse_linksable when successful
func (m *ContactsItemLinksGetResponse) GetLinks() []ContactsItemLinksGetResponse_linksable {
	return m.links
}

// GetNextPageToken gets the next_page_token property value. The next_page_token property
// returns a *string when successful
func (m *ContactsItemLinksGetResponse) GetNextPageToken() *string {
	return m.next_page_token
}

// Serialize serializes information the current object
func (m *ContactsItemLinksGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetLinks() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLinks()))
		for i, v := range m.GetLinks() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("links", cast)
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
func (m *ContactsItemLinksGetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetLinks sets the links property value. The links property
func (m *ContactsItemLinksGetResponse) SetLinks(value []ContactsItemLinksGetResponse_linksable) {
	m.links = value
}

// SetNextPageToken sets the next_page_token property value. The next_page_token property
func (m *ContactsItemLinksGetResponse) SetNextPageToken(value *string) {
	m.next_page_token = value
}

type ContactsItemLinksGetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetLinks() []ContactsItemLinksGetResponse_linksable
	GetNextPageToken() *string
	SetLinks(value []ContactsItemLinksGetResponse_linksable)
	SetNextPageToken(value *string)
}
