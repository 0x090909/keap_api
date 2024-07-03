package v2

import (
	i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
)

type ContactsItemPaymentMethodsGetResponse_records struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The cardInfo property
	cardInfo ContactsItemPaymentMethodsGetResponse_records_cardInfoable
	// The dateCreated property
	dateCreated *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	// The id property
	id *string
	// The paymentMethodType property
	paymentMethodType *string
	// The processorId property
	processorId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
	// The processorType property
	processorType *string
}

// NewContactsItemPaymentMethodsGetResponse_records instantiates a new ContactsItemPaymentMethodsGetResponse_records and sets the default values.
func NewContactsItemPaymentMethodsGetResponse_records() *ContactsItemPaymentMethodsGetResponse_records {
	m := &ContactsItemPaymentMethodsGetResponse_records{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateContactsItemPaymentMethodsGetResponse_recordsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsItemPaymentMethodsGetResponse_recordsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsItemPaymentMethodsGetResponse_records(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContactsItemPaymentMethodsGetResponse_records) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCardInfo gets the cardInfo property value. The cardInfo property
// returns a ContactsItemPaymentMethodsGetResponse_records_cardInfoable when successful
func (m *ContactsItemPaymentMethodsGetResponse_records) GetCardInfo() ContactsItemPaymentMethodsGetResponse_records_cardInfoable {
	return m.cardInfo
}

// GetDateCreated gets the dateCreated property value. The dateCreated property
// returns a *Time when successful
func (m *ContactsItemPaymentMethodsGetResponse_records) GetDateCreated() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.dateCreated
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContactsItemPaymentMethodsGetResponse_records) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["cardInfo"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateContactsItemPaymentMethodsGetResponse_records_cardInfoFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCardInfo(val.(ContactsItemPaymentMethodsGetResponse_records_cardInfoable))
		}
		return nil
	}
	res["dateCreated"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetTimeValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDateCreated(val)
		}
		return nil
	}
	res["id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetId(val)
		}
		return nil
	}
	res["paymentMethodType"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPaymentMethodType(val)
		}
		return nil
	}
	res["processorId"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetUUIDValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetProcessorId(val)
		}
		return nil
	}
	res["processorType"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetProcessorType(val)
		}
		return nil
	}
	return res
}

// GetId gets the id property value. The id property
// returns a *string when successful
func (m *ContactsItemPaymentMethodsGetResponse_records) GetId() *string {
	return m.id
}

// GetPaymentMethodType gets the paymentMethodType property value. The paymentMethodType property
// returns a *string when successful
func (m *ContactsItemPaymentMethodsGetResponse_records) GetPaymentMethodType() *string {
	return m.paymentMethodType
}

// GetProcessorId gets the processorId property value. The processorId property
// returns a *UUID when successful
func (m *ContactsItemPaymentMethodsGetResponse_records) GetProcessorId() *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID {
	return m.processorId
}

// GetProcessorType gets the processorType property value. The processorType property
// returns a *string when successful
func (m *ContactsItemPaymentMethodsGetResponse_records) GetProcessorType() *string {
	return m.processorType
}

// Serialize serializes information the current object
func (m *ContactsItemPaymentMethodsGetResponse_records) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("cardInfo", m.GetCardInfo())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteTimeValue("dateCreated", m.GetDateCreated())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("id", m.GetId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("paymentMethodType", m.GetPaymentMethodType())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteUUIDValue("processorId", m.GetProcessorId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("processorType", m.GetProcessorType())
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
func (m *ContactsItemPaymentMethodsGetResponse_records) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCardInfo sets the cardInfo property value. The cardInfo property
func (m *ContactsItemPaymentMethodsGetResponse_records) SetCardInfo(value ContactsItemPaymentMethodsGetResponse_records_cardInfoable) {
	m.cardInfo = value
}

// SetDateCreated sets the dateCreated property value. The dateCreated property
func (m *ContactsItemPaymentMethodsGetResponse_records) SetDateCreated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	m.dateCreated = value
}

// SetId sets the id property value. The id property
func (m *ContactsItemPaymentMethodsGetResponse_records) SetId(value *string) {
	m.id = value
}

// SetPaymentMethodType sets the paymentMethodType property value. The paymentMethodType property
func (m *ContactsItemPaymentMethodsGetResponse_records) SetPaymentMethodType(value *string) {
	m.paymentMethodType = value
}

// SetProcessorId sets the processorId property value. The processorId property
func (m *ContactsItemPaymentMethodsGetResponse_records) SetProcessorId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
	m.processorId = value
}

// SetProcessorType sets the processorType property value. The processorType property
func (m *ContactsItemPaymentMethodsGetResponse_records) SetProcessorType(value *string) {
	m.processorType = value
}

type ContactsItemPaymentMethodsGetResponse_recordsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCardInfo() ContactsItemPaymentMethodsGetResponse_records_cardInfoable
	GetDateCreated() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetId() *string
	GetPaymentMethodType() *string
	GetProcessorId() *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
	GetProcessorType() *string
	SetCardInfo(value ContactsItemPaymentMethodsGetResponse_records_cardInfoable)
	SetDateCreated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	SetId(value *string)
	SetPaymentMethodType(value *string)
	SetProcessorId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
	SetProcessorType(value *string)
}
