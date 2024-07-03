package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ContactsItemPaymentMethodsGetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The next_page_token property
	next_page_token *string
	// The records property
	records []ContactsItemPaymentMethodsGetResponse_recordsable
}

// NewContactsItemPaymentMethodsGetResponse instantiates a new ContactsItemPaymentMethodsGetResponse and sets the default values.
func NewContactsItemPaymentMethodsGetResponse() *ContactsItemPaymentMethodsGetResponse {
	m := &ContactsItemPaymentMethodsGetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateContactsItemPaymentMethodsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsItemPaymentMethodsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsItemPaymentMethodsGetResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContactsItemPaymentMethodsGetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContactsItemPaymentMethodsGetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
	res["records"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateContactsItemPaymentMethodsGetResponse_recordsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ContactsItemPaymentMethodsGetResponse_recordsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ContactsItemPaymentMethodsGetResponse_recordsable)
				}
			}
			m.SetRecords(res)
		}
		return nil
	}
	return res
}

// GetNextPageToken gets the next_page_token property value. The next_page_token property
// returns a *string when successful
func (m *ContactsItemPaymentMethodsGetResponse) GetNextPageToken() *string {
	return m.next_page_token
}

// GetRecords gets the records property value. The records property
// returns a []ContactsItemPaymentMethodsGetResponse_recordsable when successful
func (m *ContactsItemPaymentMethodsGetResponse) GetRecords() []ContactsItemPaymentMethodsGetResponse_recordsable {
	return m.records
}

// Serialize serializes information the current object
func (m *ContactsItemPaymentMethodsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("next_page_token", m.GetNextPageToken())
		if err != nil {
			return err
		}
	}
	if m.GetRecords() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRecords()))
		for i, v := range m.GetRecords() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("records", cast)
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
func (m *ContactsItemPaymentMethodsGetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetNextPageToken sets the next_page_token property value. The next_page_token property
func (m *ContactsItemPaymentMethodsGetResponse) SetNextPageToken(value *string) {
	m.next_page_token = value
}

// SetRecords sets the records property value. The records property
func (m *ContactsItemPaymentMethodsGetResponse) SetRecords(value []ContactsItemPaymentMethodsGetResponse_recordsable) {
	m.records = value
}

type ContactsItemPaymentMethodsGetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetNextPageToken() *string
	GetRecords() []ContactsItemPaymentMethodsGetResponse_recordsable
	SetNextPageToken(value *string)
	SetRecords(value []ContactsItemPaymentMethodsGetResponse_recordsable)
}
