package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type TransactionsItemWithTransactionGetResponse_orders_payment_plan_payment_gateway struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The merchant_account_id property
	merchant_account_id *int64
	// The use_default property
	use_default *bool
}

// NewTransactionsItemWithTransactionGetResponse_orders_payment_plan_payment_gateway instantiates a new TransactionsItemWithTransactionGetResponse_orders_payment_plan_payment_gateway and sets the default values.
func NewTransactionsItemWithTransactionGetResponse_orders_payment_plan_payment_gateway() *TransactionsItemWithTransactionGetResponse_orders_payment_plan_payment_gateway {
	m := &TransactionsItemWithTransactionGetResponse_orders_payment_plan_payment_gateway{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateTransactionsItemWithTransactionGetResponse_orders_payment_plan_payment_gatewayFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateTransactionsItemWithTransactionGetResponse_orders_payment_plan_payment_gatewayFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewTransactionsItemWithTransactionGetResponse_orders_payment_plan_payment_gateway(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *TransactionsItemWithTransactionGetResponse_orders_payment_plan_payment_gateway) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *TransactionsItemWithTransactionGetResponse_orders_payment_plan_payment_gateway) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["merchant_account_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetMerchantAccountId(val)
		}
		return nil
	}
	res["use_default"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetUseDefault(val)
		}
		return nil
	}
	return res
}

// GetMerchantAccountId gets the merchant_account_id property value. The merchant_account_id property
// returns a *int64 when successful
func (m *TransactionsItemWithTransactionGetResponse_orders_payment_plan_payment_gateway) GetMerchantAccountId() *int64 {
	return m.merchant_account_id
}

// GetUseDefault gets the use_default property value. The use_default property
// returns a *bool when successful
func (m *TransactionsItemWithTransactionGetResponse_orders_payment_plan_payment_gateway) GetUseDefault() *bool {
	return m.use_default
}

// Serialize serializes information the current object
func (m *TransactionsItemWithTransactionGetResponse_orders_payment_plan_payment_gateway) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteInt64Value("merchant_account_id", m.GetMerchantAccountId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("use_default", m.GetUseDefault())
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
func (m *TransactionsItemWithTransactionGetResponse_orders_payment_plan_payment_gateway) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetMerchantAccountId sets the merchant_account_id property value. The merchant_account_id property
func (m *TransactionsItemWithTransactionGetResponse_orders_payment_plan_payment_gateway) SetMerchantAccountId(value *int64) {
	m.merchant_account_id = value
}

// SetUseDefault sets the use_default property value. The use_default property
func (m *TransactionsItemWithTransactionGetResponse_orders_payment_plan_payment_gateway) SetUseDefault(value *bool) {
	m.use_default = value
}

type TransactionsItemWithTransactionGetResponse_orders_payment_plan_payment_gatewayable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetMerchantAccountId() *int64
	GetUseDefault() *bool
	SetMerchantAccountId(value *int64)
	SetUseDefault(value *bool)
}
