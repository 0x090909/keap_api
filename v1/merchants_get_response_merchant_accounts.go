package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MerchantsGetResponse_merchant_accounts struct {
	// The account_name property
	account_name *string
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The id property
	id *int64
	// The is_test property
	is_test *bool
	// The type property
	typeEscaped *string
}

// NewMerchantsGetResponse_merchant_accounts instantiates a new MerchantsGetResponse_merchant_accounts and sets the default values.
func NewMerchantsGetResponse_merchant_accounts() *MerchantsGetResponse_merchant_accounts {
	m := &MerchantsGetResponse_merchant_accounts{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateMerchantsGetResponse_merchant_accountsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMerchantsGetResponse_merchant_accountsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewMerchantsGetResponse_merchant_accounts(), nil
}

// GetAccountName gets the account_name property value. The account_name property
// returns a *string when successful
func (m *MerchantsGetResponse_merchant_accounts) GetAccountName() *string {
	return m.account_name
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *MerchantsGetResponse_merchant_accounts) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MerchantsGetResponse_merchant_accounts) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["account_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAccountName(val)
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
	res["is_test"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetIsTest(val)
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
func (m *MerchantsGetResponse_merchant_accounts) GetId() *int64 {
	return m.id
}

// GetIsTest gets the is_test property value. The is_test property
// returns a *bool when successful
func (m *MerchantsGetResponse_merchant_accounts) GetIsTest() *bool {
	return m.is_test
}

// GetTypeEscaped gets the type property value. The type property
// returns a *string when successful
func (m *MerchantsGetResponse_merchant_accounts) GetTypeEscaped() *string {
	return m.typeEscaped
}

// Serialize serializes information the current object
func (m *MerchantsGetResponse_merchant_accounts) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("account_name", m.GetAccountName())
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
		err := writer.WriteBoolValue("is_test", m.GetIsTest())
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

// SetAccountName sets the account_name property value. The account_name property
func (m *MerchantsGetResponse_merchant_accounts) SetAccountName(value *string) {
	m.account_name = value
}

// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MerchantsGetResponse_merchant_accounts) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetId sets the id property value. The id property
func (m *MerchantsGetResponse_merchant_accounts) SetId(value *int64) {
	m.id = value
}

// SetIsTest sets the is_test property value. The is_test property
func (m *MerchantsGetResponse_merchant_accounts) SetIsTest(value *bool) {
	m.is_test = value
}

// SetTypeEscaped sets the type property value. The type property
func (m *MerchantsGetResponse_merchant_accounts) SetTypeEscaped(value *string) {
	m.typeEscaped = value
}

type MerchantsGetResponse_merchant_accountsable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAccountName() *string
	GetId() *int64
	GetIsTest() *bool
	GetTypeEscaped() *string
	SetAccountName(value *string)
	SetId(value *int64)
	SetIsTest(value *bool)
	SetTypeEscaped(value *string)
}
