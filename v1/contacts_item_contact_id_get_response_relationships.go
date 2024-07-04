package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ContactsItemContactIdGetResponse_relationships struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The id property
	id *int64
	// The linked_contact_id property
	linked_contact_id *int64
	// The relationship_type_id property
	relationship_type_id *int64
}

// NewContactsItemContactIdGetResponse_relationships instantiates a new ContactsItemContactIdGetResponse_relationships and sets the default values.
func NewContactsItemContactIdGetResponse_relationships() *ContactsItemContactIdGetResponse_relationships {
	m := &ContactsItemContactIdGetResponse_relationships{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateContactsItemContactIdGetResponse_relationshipsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsItemContactIdGetResponse_relationshipsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsItemContactIdGetResponse_relationships(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContactsItemContactIdGetResponse_relationships) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContactsItemContactIdGetResponse_relationships) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
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
	res["linked_contact_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLinkedContactId(val)
		}
		return nil
	}
	res["relationship_type_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetRelationshipTypeId(val)
		}
		return nil
	}
	return res
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *ContactsItemContactIdGetResponse_relationships) GetId() *int64 {
	return m.id
}

// GetLinkedContactId gets the linked_contact_id property value. The linked_contact_id property
// returns a *int64 when successful
func (m *ContactsItemContactIdGetResponse_relationships) GetLinkedContactId() *int64 {
	return m.linked_contact_id
}

// GetRelationshipTypeId gets the relationship_type_id property value. The relationship_type_id property
// returns a *int64 when successful
func (m *ContactsItemContactIdGetResponse_relationships) GetRelationshipTypeId() *int64 {
	return m.relationship_type_id
}

// Serialize serializes information the current object
func (m *ContactsItemContactIdGetResponse_relationships) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt64Value("id", m.GetId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("linked_contact_id", m.GetLinkedContactId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("relationship_type_id", m.GetRelationshipTypeId())
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
func (m *ContactsItemContactIdGetResponse_relationships) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetId sets the id property value. The id property
func (m *ContactsItemContactIdGetResponse_relationships) SetId(value *int64) {
	m.id = value
}

// SetLinkedContactId sets the linked_contact_id property value. The linked_contact_id property
func (m *ContactsItemContactIdGetResponse_relationships) SetLinkedContactId(value *int64) {
	m.linked_contact_id = value
}

// SetRelationshipTypeId sets the relationship_type_id property value. The relationship_type_id property
func (m *ContactsItemContactIdGetResponse_relationships) SetRelationshipTypeId(value *int64) {
	m.relationship_type_id = value
}

type ContactsItemContactIdGetResponse_relationshipsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetId() *int64
	GetLinkedContactId() *int64
	GetRelationshipTypeId() *int64
	SetId(value *int64)
	SetLinkedContactId(value *int64)
	SetRelationshipTypeId(value *int64)
}
