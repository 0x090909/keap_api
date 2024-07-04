package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type MerchantsGetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The default_merchant_account property
	default_merchant_account *int64
	// The merchant_accounts property
	merchant_accounts []MerchantsGetResponse_merchant_accountsable
}

// NewMerchantsGetResponse instantiates a new MerchantsGetResponse and sets the default values.
func NewMerchantsGetResponse() *MerchantsGetResponse {
	m := &MerchantsGetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateMerchantsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateMerchantsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewMerchantsGetResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *MerchantsGetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetDefaultMerchantAccount gets the default_merchant_account property value. The default_merchant_account property
// returns a *int64 when successful
func (m *MerchantsGetResponse) GetDefaultMerchantAccount() *int64 {
	return m.default_merchant_account
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *MerchantsGetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["default_merchant_account"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDefaultMerchantAccount(val)
		}
		return nil
	}
	res["merchant_accounts"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateMerchantsGetResponse_merchant_accountsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]MerchantsGetResponse_merchant_accountsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(MerchantsGetResponse_merchant_accountsable)
				}
			}
			m.SetMerchantAccounts(res)
		}
		return nil
	}
	return res
}

// GetMerchantAccounts gets the merchant_accounts property value. The merchant_accounts property
// returns a []MerchantsGetResponse_merchant_accountsable when successful
func (m *MerchantsGetResponse) GetMerchantAccounts() []MerchantsGetResponse_merchant_accountsable {
	return m.merchant_accounts
}

// Serialize serializes information the current object
func (m *MerchantsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt64Value("default_merchant_account", m.GetDefaultMerchantAccount())
		if err != nil {
			return err
		}
	}
	if m.GetMerchantAccounts() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMerchantAccounts()))
		for i, v := range m.GetMerchantAccounts() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("merchant_accounts", cast)
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
func (m *MerchantsGetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetDefaultMerchantAccount sets the default_merchant_account property value. The default_merchant_account property
func (m *MerchantsGetResponse) SetDefaultMerchantAccount(value *int64) {
	m.default_merchant_account = value
}

// SetMerchantAccounts sets the merchant_accounts property value. The merchant_accounts property
func (m *MerchantsGetResponse) SetMerchantAccounts(value []MerchantsGetResponse_merchant_accountsable) {
	m.merchant_accounts = value
}

type MerchantsGetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetDefaultMerchantAccount() *int64
	GetMerchantAccounts() []MerchantsGetResponse_merchant_accountsable
	SetDefaultMerchantAccount(value *int64)
	SetMerchantAccounts(value []MerchantsGetResponse_merchant_accountsable)
}
