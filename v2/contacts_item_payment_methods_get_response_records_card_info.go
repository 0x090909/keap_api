package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ContactsItemPaymentMethodsGetResponse_records_cardInfo struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The brand property
	brand *string
	// The cardType property
	cardType *string
	// The expirationMonth property
	expirationMonth *string
	// The expirationYear property
	expirationYear *string
	// The lastFour property
	lastFour *string
}

// NewContactsItemPaymentMethodsGetResponse_records_cardInfo instantiates a new ContactsItemPaymentMethodsGetResponse_records_cardInfo and sets the default values.
func NewContactsItemPaymentMethodsGetResponse_records_cardInfo() *ContactsItemPaymentMethodsGetResponse_records_cardInfo {
	m := &ContactsItemPaymentMethodsGetResponse_records_cardInfo{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateContactsItemPaymentMethodsGetResponse_records_cardInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsItemPaymentMethodsGetResponse_records_cardInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsItemPaymentMethodsGetResponse_records_cardInfo(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContactsItemPaymentMethodsGetResponse_records_cardInfo) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetBrand gets the brand property value. The brand property
// returns a *string when successful
func (m *ContactsItemPaymentMethodsGetResponse_records_cardInfo) GetBrand() *string {
	return m.brand
}

// GetCardType gets the cardType property value. The cardType property
// returns a *string when successful
func (m *ContactsItemPaymentMethodsGetResponse_records_cardInfo) GetCardType() *string {
	return m.cardType
}

// GetExpirationMonth gets the expirationMonth property value. The expirationMonth property
// returns a *string when successful
func (m *ContactsItemPaymentMethodsGetResponse_records_cardInfo) GetExpirationMonth() *string {
	return m.expirationMonth
}

// GetExpirationYear gets the expirationYear property value. The expirationYear property
// returns a *string when successful
func (m *ContactsItemPaymentMethodsGetResponse_records_cardInfo) GetExpirationYear() *string {
	return m.expirationYear
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContactsItemPaymentMethodsGetResponse_records_cardInfo) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["brand"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetBrand(val)
		}
		return nil
	}
	res["cardType"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCardType(val)
		}
		return nil
	}
	res["expirationMonth"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetExpirationMonth(val)
		}
		return nil
	}
	res["expirationYear"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetExpirationYear(val)
		}
		return nil
	}
	res["lastFour"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLastFour(val)
		}
		return nil
	}
	return res
}

// GetLastFour gets the lastFour property value. The lastFour property
// returns a *string when successful
func (m *ContactsItemPaymentMethodsGetResponse_records_cardInfo) GetLastFour() *string {
	return m.lastFour
}

// Serialize serializes information the current object
func (m *ContactsItemPaymentMethodsGetResponse_records_cardInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("brand", m.GetBrand())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("cardType", m.GetCardType())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("expirationMonth", m.GetExpirationMonth())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("expirationYear", m.GetExpirationYear())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("lastFour", m.GetLastFour())
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
func (m *ContactsItemPaymentMethodsGetResponse_records_cardInfo) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetBrand sets the brand property value. The brand property
func (m *ContactsItemPaymentMethodsGetResponse_records_cardInfo) SetBrand(value *string) {
	m.brand = value
}

// SetCardType sets the cardType property value. The cardType property
func (m *ContactsItemPaymentMethodsGetResponse_records_cardInfo) SetCardType(value *string) {
	m.cardType = value
}

// SetExpirationMonth sets the expirationMonth property value. The expirationMonth property
func (m *ContactsItemPaymentMethodsGetResponse_records_cardInfo) SetExpirationMonth(value *string) {
	m.expirationMonth = value
}

// SetExpirationYear sets the expirationYear property value. The expirationYear property
func (m *ContactsItemPaymentMethodsGetResponse_records_cardInfo) SetExpirationYear(value *string) {
	m.expirationYear = value
}

// SetLastFour sets the lastFour property value. The lastFour property
func (m *ContactsItemPaymentMethodsGetResponse_records_cardInfo) SetLastFour(value *string) {
	m.lastFour = value
}

type ContactsItemPaymentMethodsGetResponse_records_cardInfoable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetBrand() *string
	GetCardType() *string
	GetExpirationMonth() *string
	GetExpirationYear() *string
	GetLastFour() *string
	SetBrand(value *string)
	SetCardType(value *string)
	SetExpirationMonth(value *string)
	SetExpirationYear(value *string)
	SetLastFour(value *string)
}
