package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ContactsItemNotesGetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The next_page_token property
	next_page_token *string
	// The notes property
	notes []ContactsItemNotesGetResponse_notesable
}

// NewContactsItemNotesGetResponse instantiates a new ContactsItemNotesGetResponse and sets the default values.
func NewContactsItemNotesGetResponse() *ContactsItemNotesGetResponse {
	m := &ContactsItemNotesGetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateContactsItemNotesGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsItemNotesGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsItemNotesGetResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContactsItemNotesGetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContactsItemNotesGetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
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
	res["notes"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateContactsItemNotesGetResponse_notesFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ContactsItemNotesGetResponse_notesable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ContactsItemNotesGetResponse_notesable)
				}
			}
			m.SetNotes(res)
		}
		return nil
	}
	return res
}

// GetNextPageToken gets the next_page_token property value. The next_page_token property
// returns a *string when successful
func (m *ContactsItemNotesGetResponse) GetNextPageToken() *string {
	return m.next_page_token
}

// GetNotes gets the notes property value. The notes property
// returns a []ContactsItemNotesGetResponse_notesable when successful
func (m *ContactsItemNotesGetResponse) GetNotes() []ContactsItemNotesGetResponse_notesable {
	return m.notes
}

// Serialize serializes information the current object
func (m *ContactsItemNotesGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("next_page_token", m.GetNextPageToken())
		if err != nil {
			return err
		}
	}
	if m.GetNotes() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNotes()))
		for i, v := range m.GetNotes() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("notes", cast)
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
func (m *ContactsItemNotesGetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetNextPageToken sets the next_page_token property value. The next_page_token property
func (m *ContactsItemNotesGetResponse) SetNextPageToken(value *string) {
	m.next_page_token = value
}

// SetNotes sets the notes property value. The notes property
func (m *ContactsItemNotesGetResponse) SetNotes(value []ContactsItemNotesGetResponse_notesable) {
	m.notes = value
}

type ContactsItemNotesGetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetNextPageToken() *string
	GetNotes() []ContactsItemNotesGetResponse_notesable
	SetNextPageToken(value *string)
	SetNotes(value []ContactsItemNotesGetResponse_notesable)
}
