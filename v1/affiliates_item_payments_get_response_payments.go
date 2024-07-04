package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AffiliatesItemPaymentsGetResponse_payments struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The amount property
	amount *float64
	// The date property
	date *string
	// The id property
	id *int64
	// The notes property
	notes *string
	// The type property
	typeEscaped *string
}

// NewAffiliatesItemPaymentsGetResponse_payments instantiates a new AffiliatesItemPaymentsGetResponse_payments and sets the default values.
func NewAffiliatesItemPaymentsGetResponse_payments() *AffiliatesItemPaymentsGetResponse_payments {
	m := &AffiliatesItemPaymentsGetResponse_payments{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateAffiliatesItemPaymentsGetResponse_paymentsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAffiliatesItemPaymentsGetResponse_paymentsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewAffiliatesItemPaymentsGetResponse_payments(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *AffiliatesItemPaymentsGetResponse_payments) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAmount gets the amount property value. The amount property
// returns a *float64 when successful
func (m *AffiliatesItemPaymentsGetResponse_payments) GetAmount() *float64 {
	return m.amount
}

// GetDate gets the date property value. The date property
// returns a *string when successful
func (m *AffiliatesItemPaymentsGetResponse_payments) GetDate() *string {
	return m.date
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AffiliatesItemPaymentsGetResponse_payments) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["amount"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetFloat64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAmount(val)
		}
		return nil
	}
	res["date"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDate(val)
		}
		return nil
	}
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
	res["notes"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetNotes(val)
		}
		return nil
	}
	res["type"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTypeEscaped(val)
		}
		return nil
	}
	return res
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *AffiliatesItemPaymentsGetResponse_payments) GetId() *int64 {
	return m.id
}

// GetNotes gets the notes property value. The notes property
// returns a *string when successful
func (m *AffiliatesItemPaymentsGetResponse_payments) GetNotes() *string {
	return m.notes
}

// GetTypeEscaped gets the type property value. The type property
// returns a *string when successful
func (m *AffiliatesItemPaymentsGetResponse_payments) GetTypeEscaped() *string {
	return m.typeEscaped
}

// Serialize serializes information the current object
func (m *AffiliatesItemPaymentsGetResponse_payments) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteFloat64Value("amount", m.GetAmount())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("date", m.GetDate())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("id", m.GetId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("notes", m.GetNotes())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("type", m.GetTypeEscaped())
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
func (m *AffiliatesItemPaymentsGetResponse_payments) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAmount sets the amount property value. The amount property
func (m *AffiliatesItemPaymentsGetResponse_payments) SetAmount(value *float64) {
	m.amount = value
}

// SetDate sets the date property value. The date property
func (m *AffiliatesItemPaymentsGetResponse_payments) SetDate(value *string) {
	m.date = value
}

// SetId sets the id property value. The id property
func (m *AffiliatesItemPaymentsGetResponse_payments) SetId(value *int64) {
	m.id = value
}

// SetNotes sets the notes property value. The notes property
func (m *AffiliatesItemPaymentsGetResponse_payments) SetNotes(value *string) {
	m.notes = value
}

// SetTypeEscaped sets the type property value. The type property
func (m *AffiliatesItemPaymentsGetResponse_payments) SetTypeEscaped(value *string) {
	m.typeEscaped = value
}

type AffiliatesItemPaymentsGetResponse_paymentsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAmount() *float64
	GetDate() *string
	GetId() *int64
	GetNotes() *string
	GetTypeEscaped() *string
	SetAmount(value *float64)
	SetDate(value *string)
	SetId(value *int64)
	SetNotes(value *string)
	SetTypeEscaped(value *string)
}
