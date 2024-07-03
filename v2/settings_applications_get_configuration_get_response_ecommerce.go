package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SettingsApplicationsGetConfigurationGetResponse_ecommerce struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The credit_card_types property
	credit_card_types *string
	// The currency property
	currency *string
	// The default_charge_max_retry_attempts property
	default_charge_max_retry_attempts *int32
	// The default_country property
	default_country *string
	// The default_merchant property
	default_merchant *string
	// The default_number_of_days_between_charge_attempts property
	default_number_of_days_between_charge_attempts *int32
	// The default_tax property
	default_tax *string
	// The default_to_auto_charge property
	default_to_auto_charge *bool
	// The merchant_account_max_retry_attempts property
	merchant_account_max_retry_attempts *int32
	// The payment_types property
	payment_types *string
	// The promo_codes property
	promo_codes *string
	// The track_cost_per_unit property
	track_cost_per_unit *bool
	// The track_inventory property
	track_inventory *bool
}

// NewSettingsApplicationsGetConfigurationGetResponse_ecommerce instantiates a new SettingsApplicationsGetConfigurationGetResponse_ecommerce and sets the default values.
func NewSettingsApplicationsGetConfigurationGetResponse_ecommerce() *SettingsApplicationsGetConfigurationGetResponse_ecommerce {
	m := &SettingsApplicationsGetConfigurationGetResponse_ecommerce{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateSettingsApplicationsGetConfigurationGetResponse_ecommerceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSettingsApplicationsGetConfigurationGetResponse_ecommerceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSettingsApplicationsGetConfigurationGetResponse_ecommerce(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_ecommerce) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetCreditCardTypes gets the credit_card_types property value. The credit_card_types property
// returns a *string when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_ecommerce) GetCreditCardTypes() *string {
	return m.credit_card_types
}

// GetCurrency gets the currency property value. The currency property
// returns a *string when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_ecommerce) GetCurrency() *string {
	return m.currency
}

// GetDefaultChargeMaxRetryAttempts gets the default_charge_max_retry_attempts property value. The default_charge_max_retry_attempts property
// returns a *int32 when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_ecommerce) GetDefaultChargeMaxRetryAttempts() *int32 {
	return m.default_charge_max_retry_attempts
}

// GetDefaultCountry gets the default_country property value. The default_country property
// returns a *string when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_ecommerce) GetDefaultCountry() *string {
	return m.default_country
}

// GetDefaultMerchant gets the default_merchant property value. The default_merchant property
// returns a *string when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_ecommerce) GetDefaultMerchant() *string {
	return m.default_merchant
}

// GetDefaultNumberOfDaysBetweenChargeAttempts gets the default_number_of_days_between_charge_attempts property value. The default_number_of_days_between_charge_attempts property
// returns a *int32 when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_ecommerce) GetDefaultNumberOfDaysBetweenChargeAttempts() *int32 {
	return m.default_number_of_days_between_charge_attempts
}

// GetDefaultTax gets the default_tax property value. The default_tax property
// returns a *string when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_ecommerce) GetDefaultTax() *string {
	return m.default_tax
}

// GetDefaultToAutoCharge gets the default_to_auto_charge property value. The default_to_auto_charge property
// returns a *bool when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_ecommerce) GetDefaultToAutoCharge() *bool {
	return m.default_to_auto_charge
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_ecommerce) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["credit_card_types"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCreditCardTypes(val)
		}
		return nil
	}
	res["currency"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCurrency(val)
		}
		return nil
	}
	res["default_charge_max_retry_attempts"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDefaultChargeMaxRetryAttempts(val)
		}
		return nil
	}
	res["default_country"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDefaultCountry(val)
		}
		return nil
	}
	res["default_merchant"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDefaultMerchant(val)
		}
		return nil
	}
	res["default_number_of_days_between_charge_attempts"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDefaultNumberOfDaysBetweenChargeAttempts(val)
		}
		return nil
	}
	res["default_tax"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDefaultTax(val)
		}
		return nil
	}
	res["default_to_auto_charge"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDefaultToAutoCharge(val)
		}
		return nil
	}
	res["merchant_account_max_retry_attempts"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt32Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetMerchantAccountMaxRetryAttempts(val)
		}
		return nil
	}
	res["payment_types"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPaymentTypes(val)
		}
		return nil
	}
	res["promo_codes"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPromoCodes(val)
		}
		return nil
	}
	res["track_cost_per_unit"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTrackCostPerUnit(val)
		}
		return nil
	}
	res["track_inventory"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTrackInventory(val)
		}
		return nil
	}
	return res
}

// GetMerchantAccountMaxRetryAttempts gets the merchant_account_max_retry_attempts property value. The merchant_account_max_retry_attempts property
// returns a *int32 when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_ecommerce) GetMerchantAccountMaxRetryAttempts() *int32 {
	return m.merchant_account_max_retry_attempts
}

// GetPaymentTypes gets the payment_types property value. The payment_types property
// returns a *string when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_ecommerce) GetPaymentTypes() *string {
	return m.payment_types
}

// GetPromoCodes gets the promo_codes property value. The promo_codes property
// returns a *string when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_ecommerce) GetPromoCodes() *string {
	return m.promo_codes
}

// GetTrackCostPerUnit gets the track_cost_per_unit property value. The track_cost_per_unit property
// returns a *bool when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_ecommerce) GetTrackCostPerUnit() *bool {
	return m.track_cost_per_unit
}

// GetTrackInventory gets the track_inventory property value. The track_inventory property
// returns a *bool when successful
func (m *SettingsApplicationsGetConfigurationGetResponse_ecommerce) GetTrackInventory() *bool {
	return m.track_inventory
}

// Serialize serializes information the current object
func (m *SettingsApplicationsGetConfigurationGetResponse_ecommerce) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("credit_card_types", m.GetCreditCardTypes())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("currency", m.GetCurrency())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("default_charge_max_retry_attempts", m.GetDefaultChargeMaxRetryAttempts())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("default_country", m.GetDefaultCountry())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("default_merchant", m.GetDefaultMerchant())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("default_number_of_days_between_charge_attempts", m.GetDefaultNumberOfDaysBetweenChargeAttempts())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("default_tax", m.GetDefaultTax())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("default_to_auto_charge", m.GetDefaultToAutoCharge())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt32Value("merchant_account_max_retry_attempts", m.GetMerchantAccountMaxRetryAttempts())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("payment_types", m.GetPaymentTypes())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("promo_codes", m.GetPromoCodes())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("track_cost_per_unit", m.GetTrackCostPerUnit())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("track_inventory", m.GetTrackInventory())
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
func (m *SettingsApplicationsGetConfigurationGetResponse_ecommerce) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetCreditCardTypes sets the credit_card_types property value. The credit_card_types property
func (m *SettingsApplicationsGetConfigurationGetResponse_ecommerce) SetCreditCardTypes(value *string) {
	m.credit_card_types = value
}

// SetCurrency sets the currency property value. The currency property
func (m *SettingsApplicationsGetConfigurationGetResponse_ecommerce) SetCurrency(value *string) {
	m.currency = value
}

// SetDefaultChargeMaxRetryAttempts sets the default_charge_max_retry_attempts property value. The default_charge_max_retry_attempts property
func (m *SettingsApplicationsGetConfigurationGetResponse_ecommerce) SetDefaultChargeMaxRetryAttempts(value *int32) {
	m.default_charge_max_retry_attempts = value
}

// SetDefaultCountry sets the default_country property value. The default_country property
func (m *SettingsApplicationsGetConfigurationGetResponse_ecommerce) SetDefaultCountry(value *string) {
	m.default_country = value
}

// SetDefaultMerchant sets the default_merchant property value. The default_merchant property
func (m *SettingsApplicationsGetConfigurationGetResponse_ecommerce) SetDefaultMerchant(value *string) {
	m.default_merchant = value
}

// SetDefaultNumberOfDaysBetweenChargeAttempts sets the default_number_of_days_between_charge_attempts property value. The default_number_of_days_between_charge_attempts property
func (m *SettingsApplicationsGetConfigurationGetResponse_ecommerce) SetDefaultNumberOfDaysBetweenChargeAttempts(value *int32) {
	m.default_number_of_days_between_charge_attempts = value
}

// SetDefaultTax sets the default_tax property value. The default_tax property
func (m *SettingsApplicationsGetConfigurationGetResponse_ecommerce) SetDefaultTax(value *string) {
	m.default_tax = value
}

// SetDefaultToAutoCharge sets the default_to_auto_charge property value. The default_to_auto_charge property
func (m *SettingsApplicationsGetConfigurationGetResponse_ecommerce) SetDefaultToAutoCharge(value *bool) {
	m.default_to_auto_charge = value
}

// SetMerchantAccountMaxRetryAttempts sets the merchant_account_max_retry_attempts property value. The merchant_account_max_retry_attempts property
func (m *SettingsApplicationsGetConfigurationGetResponse_ecommerce) SetMerchantAccountMaxRetryAttempts(value *int32) {
	m.merchant_account_max_retry_attempts = value
}

// SetPaymentTypes sets the payment_types property value. The payment_types property
func (m *SettingsApplicationsGetConfigurationGetResponse_ecommerce) SetPaymentTypes(value *string) {
	m.payment_types = value
}

// SetPromoCodes sets the promo_codes property value. The promo_codes property
func (m *SettingsApplicationsGetConfigurationGetResponse_ecommerce) SetPromoCodes(value *string) {
	m.promo_codes = value
}

// SetTrackCostPerUnit sets the track_cost_per_unit property value. The track_cost_per_unit property
func (m *SettingsApplicationsGetConfigurationGetResponse_ecommerce) SetTrackCostPerUnit(value *bool) {
	m.track_cost_per_unit = value
}

// SetTrackInventory sets the track_inventory property value. The track_inventory property
func (m *SettingsApplicationsGetConfigurationGetResponse_ecommerce) SetTrackInventory(value *bool) {
	m.track_inventory = value
}

type SettingsApplicationsGetConfigurationGetResponse_ecommerceable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetCreditCardTypes() *string
	GetCurrency() *string
	GetDefaultChargeMaxRetryAttempts() *int32
	GetDefaultCountry() *string
	GetDefaultMerchant() *string
	GetDefaultNumberOfDaysBetweenChargeAttempts() *int32
	GetDefaultTax() *string
	GetDefaultToAutoCharge() *bool
	GetMerchantAccountMaxRetryAttempts() *int32
	GetPaymentTypes() *string
	GetPromoCodes() *string
	GetTrackCostPerUnit() *bool
	GetTrackInventory() *bool
	SetCreditCardTypes(value *string)
	SetCurrency(value *string)
	SetDefaultChargeMaxRetryAttempts(value *int32)
	SetDefaultCountry(value *string)
	SetDefaultMerchant(value *string)
	SetDefaultNumberOfDaysBetweenChargeAttempts(value *int32)
	SetDefaultTax(value *string)
	SetDefaultToAutoCharge(value *bool)
	SetMerchantAccountMaxRetryAttempts(value *int32)
	SetPaymentTypes(value *string)
	SetPromoCodes(value *string)
	SetTrackCostPerUnit(value *bool)
	SetTrackInventory(value *bool)
}
