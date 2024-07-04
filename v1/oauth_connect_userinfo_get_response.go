package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OauthConnectUserinfoGetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The email property
	email *string
	// The family_name property
	family_name *string
	// The given_name property
	given_name *string
	// The global_user_id property
	global_user_id *int64
	// The infusionsoft_id property
	infusionsoft_id *string
	// The is_admin property
	is_admin *bool
	// The middle_name property
	middle_name *string
	// The preferred_name property
	preferred_name *string
	// The sub property
	sub *string
}

// NewOauthConnectUserinfoGetResponse instantiates a new OauthConnectUserinfoGetResponse and sets the default values.
func NewOauthConnectUserinfoGetResponse() *OauthConnectUserinfoGetResponse {
	m := &OauthConnectUserinfoGetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateOauthConnectUserinfoGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOauthConnectUserinfoGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewOauthConnectUserinfoGetResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *OauthConnectUserinfoGetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetEmail gets the email property value. The email property
// returns a *string when successful
func (m *OauthConnectUserinfoGetResponse) GetEmail() *string {
	return m.email
}

// GetFamilyName gets the family_name property value. The family_name property
// returns a *string when successful
func (m *OauthConnectUserinfoGetResponse) GetFamilyName() *string {
	return m.family_name
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OauthConnectUserinfoGetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["email"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetEmail(val)
		}
		return nil
	}
	res["family_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetFamilyName(val)
		}
		return nil
	}
	res["given_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetGivenName(val)
		}
		return nil
	}
	res["global_user_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetGlobalUserId(val)
		}
		return nil
	}
	res["infusionsoft_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetInfusionsoftId(val)
		}
		return nil
	}
	res["is_admin"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetIsAdmin(val)
		}
		return nil
	}
	res["middle_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetMiddleName(val)
		}
		return nil
	}
	res["preferred_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPreferredName(val)
		}
		return nil
	}
	res["sub"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSub(val)
		}
		return nil
	}
	return res
}

// GetGivenName gets the given_name property value. The given_name property
// returns a *string when successful
func (m *OauthConnectUserinfoGetResponse) GetGivenName() *string {
	return m.given_name
}

// GetGlobalUserId gets the global_user_id property value. The global_user_id property
// returns a *int64 when successful
func (m *OauthConnectUserinfoGetResponse) GetGlobalUserId() *int64 {
	return m.global_user_id
}

// GetInfusionsoftId gets the infusionsoft_id property value. The infusionsoft_id property
// returns a *string when successful
func (m *OauthConnectUserinfoGetResponse) GetInfusionsoftId() *string {
	return m.infusionsoft_id
}

// GetIsAdmin gets the is_admin property value. The is_admin property
// returns a *bool when successful
func (m *OauthConnectUserinfoGetResponse) GetIsAdmin() *bool {
	return m.is_admin
}

// GetMiddleName gets the middle_name property value. The middle_name property
// returns a *string when successful
func (m *OauthConnectUserinfoGetResponse) GetMiddleName() *string {
	return m.middle_name
}

// GetPreferredName gets the preferred_name property value. The preferred_name property
// returns a *string when successful
func (m *OauthConnectUserinfoGetResponse) GetPreferredName() *string {
	return m.preferred_name
}

// GetSub gets the sub property value. The sub property
// returns a *string when successful
func (m *OauthConnectUserinfoGetResponse) GetSub() *string {
	return m.sub
}

// Serialize serializes information the current object
func (m *OauthConnectUserinfoGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteStringValue("email", m.GetEmail())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("family_name", m.GetFamilyName())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("given_name", m.GetGivenName())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("global_user_id", m.GetGlobalUserId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("infusionsoft_id", m.GetInfusionsoftId())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteBoolValue("is_admin", m.GetIsAdmin())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("middle_name", m.GetMiddleName())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("preferred_name", m.GetPreferredName())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("sub", m.GetSub())
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
func (m *OauthConnectUserinfoGetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetEmail sets the email property value. The email property
func (m *OauthConnectUserinfoGetResponse) SetEmail(value *string) {
	m.email = value
}

// SetFamilyName sets the family_name property value. The family_name property
func (m *OauthConnectUserinfoGetResponse) SetFamilyName(value *string) {
	m.family_name = value
}

// SetGivenName sets the given_name property value. The given_name property
func (m *OauthConnectUserinfoGetResponse) SetGivenName(value *string) {
	m.given_name = value
}

// SetGlobalUserId sets the global_user_id property value. The global_user_id property
func (m *OauthConnectUserinfoGetResponse) SetGlobalUserId(value *int64) {
	m.global_user_id = value
}

// SetInfusionsoftId sets the infusionsoft_id property value. The infusionsoft_id property
func (m *OauthConnectUserinfoGetResponse) SetInfusionsoftId(value *string) {
	m.infusionsoft_id = value
}

// SetIsAdmin sets the is_admin property value. The is_admin property
func (m *OauthConnectUserinfoGetResponse) SetIsAdmin(value *bool) {
	m.is_admin = value
}

// SetMiddleName sets the middle_name property value. The middle_name property
func (m *OauthConnectUserinfoGetResponse) SetMiddleName(value *string) {
	m.middle_name = value
}

// SetPreferredName sets the preferred_name property value. The preferred_name property
func (m *OauthConnectUserinfoGetResponse) SetPreferredName(value *string) {
	m.preferred_name = value
}

// SetSub sets the sub property value. The sub property
func (m *OauthConnectUserinfoGetResponse) SetSub(value *string) {
	m.sub = value
}

type OauthConnectUserinfoGetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetEmail() *string
	GetFamilyName() *string
	GetGivenName() *string
	GetGlobalUserId() *int64
	GetInfusionsoftId() *string
	GetIsAdmin() *bool
	GetMiddleName() *string
	GetPreferredName() *string
	GetSub() *string
	SetEmail(value *string)
	SetFamilyName(value *string)
	SetGivenName(value *string)
	SetGlobalUserId(value *int64)
	SetInfusionsoftId(value *string)
	SetIsAdmin(value *bool)
	SetMiddleName(value *string)
	SetPreferredName(value *string)
	SetSub(value *string)
}
