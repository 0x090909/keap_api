package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SettingApplicationConfigurationGetResponse_affiliate struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The choose_affiliate property
	choose_affiliate *bool
	// The commission property
	commission SettingApplicationConfigurationGetResponse_affiliate_commissionable
	// The custom_affiliate_url property
	custom_affiliate_url *string
	// The display_affiliate_ip_address property
	display_affiliate_ip_address *bool
	// The do_not_notify_affiliate property
	do_not_notify_affiliate *bool
	// The skip_pay_if_unused property
	skip_pay_if_unused *bool
	// The use_referral_history_if_no_tracking_cookie property
	use_referral_history_if_no_tracking_cookie *bool
}

// NewSettingApplicationConfigurationGetResponse_affiliate instantiates a new SettingApplicationConfigurationGetResponse_affiliate and sets the default values.
func NewSettingApplicationConfigurationGetResponse_affiliate() *SettingApplicationConfigurationGetResponse_affiliate {
	m := &SettingApplicationConfigurationGetResponse_affiliate{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateSettingApplicationConfigurationGetResponse_affiliateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSettingApplicationConfigurationGetResponse_affiliateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewSettingApplicationConfigurationGetResponse_affiliate(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *SettingApplicationConfigurationGetResponse_affiliate) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetChooseAffiliate gets the choose_affiliate property value. The choose_affiliate property
// returns a *bool when successful
func (m *SettingApplicationConfigurationGetResponse_affiliate) GetChooseAffiliate() *bool {
	return m.choose_affiliate
}

// GetCommission gets the commission property value. The commission property
// returns a SettingApplicationConfigurationGetResponse_affiliate_commissionable when successful
func (m *SettingApplicationConfigurationGetResponse_affiliate) GetCommission() SettingApplicationConfigurationGetResponse_affiliate_commissionable {
	return m.commission
}

// GetCustomAffiliateUrl gets the custom_affiliate_url property value. The custom_affiliate_url property
// returns a *string when successful
func (m *SettingApplicationConfigurationGetResponse_affiliate) GetCustomAffiliateUrl() *string {
	return m.custom_affiliate_url
}

// GetDisplayAffiliateIpAddress gets the display_affiliate_ip_address property value. The display_affiliate_ip_address property
// returns a *bool when successful
func (m *SettingApplicationConfigurationGetResponse_affiliate) GetDisplayAffiliateIpAddress() *bool {
	return m.display_affiliate_ip_address
}

// GetDoNotNotifyAffiliate gets the do_not_notify_affiliate property value. The do_not_notify_affiliate property
// returns a *bool when successful
func (m *SettingApplicationConfigurationGetResponse_affiliate) GetDoNotNotifyAffiliate() *bool {
	return m.do_not_notify_affiliate
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SettingApplicationConfigurationGetResponse_affiliate) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["choose_affiliate"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetChooseAffiliate(val)
		}
		return nil
	}
	res["commission"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateSettingApplicationConfigurationGetResponse_affiliate_commissionFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCommission(val.(SettingApplicationConfigurationGetResponse_affiliate_commissionable))
		}
		return nil
	}
	res["custom_affiliate_url"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCustomAffiliateUrl(val)
		}
		return nil
	}
	res["display_affiliate_ip_address"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDisplayAffiliateIpAddress(val)
		}
		return nil
	}
	res["do_not_notify_affiliate"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDoNotNotifyAffiliate(val)
		}
		return nil
	}
	res["skip_pay_if_unused"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSkipPayIfUnused(val)
		}
		return nil
	}
	res["use_referral_history_if_no_tracking_cookie"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetUseReferralHistoryIfNoTrackingCookie(val)
		}
		return nil
	}
	return res
}

// GetSkipPayIfUnused gets the skip_pay_if_unused property value. The skip_pay_if_unused property
// returns a *bool when successful
func (m *SettingApplicationConfigurationGetResponse_affiliate) GetSkipPayIfUnused() *bool {
	return m.skip_pay_if_unused
}

// GetUseReferralHistoryIfNoTrackingCookie gets the use_referral_history_if_no_tracking_cookie property value. The use_referral_history_if_no_tracking_cookie property
// returns a *bool when successful
func (m *SettingApplicationConfigurationGetResponse_affiliate) GetUseReferralHistoryIfNoTrackingCookie() *bool {
	return m.use_referral_history_if_no_tracking_cookie
}

// Serialize serializes information the current object
func (m *SettingApplicationConfigurationGetResponse_affiliate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteBoolValue("choose_affiliate", m.GetChooseAffiliate())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("commission", m.GetCommission())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("custom_affiliate_url", m.GetCustomAffiliateUrl())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("display_affiliate_ip_address", m.GetDisplayAffiliateIpAddress())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("do_not_notify_affiliate", m.GetDoNotNotifyAffiliate())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("skip_pay_if_unused", m.GetSkipPayIfUnused())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("use_referral_history_if_no_tracking_cookie", m.GetUseReferralHistoryIfNoTrackingCookie())
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
func (m *SettingApplicationConfigurationGetResponse_affiliate) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetChooseAffiliate sets the choose_affiliate property value. The choose_affiliate property
func (m *SettingApplicationConfigurationGetResponse_affiliate) SetChooseAffiliate(value *bool) {
	m.choose_affiliate = value
}

// SetCommission sets the commission property value. The commission property
func (m *SettingApplicationConfigurationGetResponse_affiliate) SetCommission(value SettingApplicationConfigurationGetResponse_affiliate_commissionable) {
	m.commission = value
}

// SetCustomAffiliateUrl sets the custom_affiliate_url property value. The custom_affiliate_url property
func (m *SettingApplicationConfigurationGetResponse_affiliate) SetCustomAffiliateUrl(value *string) {
	m.custom_affiliate_url = value
}

// SetDisplayAffiliateIpAddress sets the display_affiliate_ip_address property value. The display_affiliate_ip_address property
func (m *SettingApplicationConfigurationGetResponse_affiliate) SetDisplayAffiliateIpAddress(value *bool) {
	m.display_affiliate_ip_address = value
}

// SetDoNotNotifyAffiliate sets the do_not_notify_affiliate property value. The do_not_notify_affiliate property
func (m *SettingApplicationConfigurationGetResponse_affiliate) SetDoNotNotifyAffiliate(value *bool) {
	m.do_not_notify_affiliate = value
}

// SetSkipPayIfUnused sets the skip_pay_if_unused property value. The skip_pay_if_unused property
func (m *SettingApplicationConfigurationGetResponse_affiliate) SetSkipPayIfUnused(value *bool) {
	m.skip_pay_if_unused = value
}

// SetUseReferralHistoryIfNoTrackingCookie sets the use_referral_history_if_no_tracking_cookie property value. The use_referral_history_if_no_tracking_cookie property
func (m *SettingApplicationConfigurationGetResponse_affiliate) SetUseReferralHistoryIfNoTrackingCookie(value *bool) {
	m.use_referral_history_if_no_tracking_cookie = value
}

type SettingApplicationConfigurationGetResponse_affiliateable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetChooseAffiliate() *bool
	GetCommission() SettingApplicationConfigurationGetResponse_affiliate_commissionable
	GetCustomAffiliateUrl() *string
	GetDisplayAffiliateIpAddress() *bool
	GetDoNotNotifyAffiliate() *bool
	GetSkipPayIfUnused() *bool
	GetUseReferralHistoryIfNoTrackingCookie() *bool
	SetChooseAffiliate(value *bool)
	SetCommission(value SettingApplicationConfigurationGetResponse_affiliate_commissionable)
	SetCustomAffiliateUrl(value *string)
	SetDisplayAffiliateIpAddress(value *bool)
	SetDoNotNotifyAffiliate(value *bool)
	SetSkipPayIfUnused(value *bool)
	SetUseReferralHistoryIfNoTrackingCookie(value *bool)
}
