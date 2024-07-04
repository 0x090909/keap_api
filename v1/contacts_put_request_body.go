package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
)

type ContactsPutRequestBody struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The addresses property
	addresses []ContactsPutRequestBody_addressesable
	// The anniversary property
	anniversary *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	// The birthday property
	birthday *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	// The company property
	company ContactsPutRequestBody_companyable
	// The contact_type property
	contact_type *string
	// The custom_fields property
	custom_fields []ContactsPutRequestBody_custom_fieldsable
	// The email_addresses property
	email_addresses []ContactsPutRequestBody_email_addressesable
	// The family_name property
	family_name *string
	// The fax_numbers property
	fax_numbers []ContactsPutRequestBody_fax_numbersable
	// The given_name property
	given_name *string
	// The job_title property
	job_title *string
	// The lead_source_id property
	lead_source_id *int64
	// The middle_name property
	middle_name *string
	// The opt_in_reason property
	opt_in_reason *string
	// The origin property
	origin ContactsPutRequestBody_originable
	// The owner_id property
	owner_id *int64
	// The phone_numbers property
	phone_numbers []ContactsPutRequestBody_phone_numbersable
	// The preferred_locale property
	preferred_locale *string
	// The preferred_name property
	preferred_name *string
	// The prefix property
	prefix *string
	// The social_accounts property
	social_accounts []ContactsPutRequestBody_social_accountsable
	// The spouse_name property
	spouse_name *string
	// The suffix property
	suffix *string
	// The time_zone property
	time_zone *string
	// The website property
	website *string
}

// NewContactsPutRequestBody instantiates a new ContactsPutRequestBody and sets the default values.
func NewContactsPutRequestBody() *ContactsPutRequestBody {
	m := &ContactsPutRequestBody{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateContactsPutRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsPutRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsPutRequestBody(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContactsPutRequestBody) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAddresses gets the addresses property value. The addresses property
// returns a []ContactsPutRequestBody_addressesable when successful
func (m *ContactsPutRequestBody) GetAddresses() []ContactsPutRequestBody_addressesable {
	return m.addresses
}

// GetAnniversary gets the anniversary property value. The anniversary property
// returns a *Time when successful
func (m *ContactsPutRequestBody) GetAnniversary() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.anniversary
}

// GetBirthday gets the birthday property value. The birthday property
// returns a *Time when successful
func (m *ContactsPutRequestBody) GetBirthday() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.birthday
}

// GetCompany gets the company property value. The company property
// returns a ContactsPutRequestBody_companyable when successful
func (m *ContactsPutRequestBody) GetCompany() ContactsPutRequestBody_companyable {
	return m.company
}

// GetContactType gets the contact_type property value. The contact_type property
// returns a *string when successful
func (m *ContactsPutRequestBody) GetContactType() *string {
	return m.contact_type
}

// GetCustomFields gets the custom_fields property value. The custom_fields property
// returns a []ContactsPutRequestBody_custom_fieldsable when successful
func (m *ContactsPutRequestBody) GetCustomFields() []ContactsPutRequestBody_custom_fieldsable {
	return m.custom_fields
}

// GetEmailAddresses gets the email_addresses property value. The email_addresses property
// returns a []ContactsPutRequestBody_email_addressesable when successful
func (m *ContactsPutRequestBody) GetEmailAddresses() []ContactsPutRequestBody_email_addressesable {
	return m.email_addresses
}

// GetFamilyName gets the family_name property value. The family_name property
// returns a *string when successful
func (m *ContactsPutRequestBody) GetFamilyName() *string {
	return m.family_name
}

// GetFaxNumbers gets the fax_numbers property value. The fax_numbers property
// returns a []ContactsPutRequestBody_fax_numbersable when successful
func (m *ContactsPutRequestBody) GetFaxNumbers() []ContactsPutRequestBody_fax_numbersable {
	return m.fax_numbers
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContactsPutRequestBody) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["addresses"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateContactsPutRequestBody_addressesFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ContactsPutRequestBody_addressesable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ContactsPutRequestBody_addressesable)
				}
			}
			m.SetAddresses(res)
		}
		return nil
	}
	res["anniversary"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetTimeValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAnniversary(val)
		}
		return nil
	}
	res["birthday"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetTimeValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetBirthday(val)
		}
		return nil
	}
	res["company"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateContactsPutRequestBody_companyFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCompany(val.(ContactsPutRequestBody_companyable))
		}
		return nil
	}
	res["contact_type"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetContactType(val)
		}
		return nil
	}
	res["custom_fields"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateContactsPutRequestBody_custom_fieldsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ContactsPutRequestBody_custom_fieldsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ContactsPutRequestBody_custom_fieldsable)
				}
			}
			m.SetCustomFields(res)
		}
		return nil
	}
	res["email_addresses"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateContactsPutRequestBody_email_addressesFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ContactsPutRequestBody_email_addressesable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ContactsPutRequestBody_email_addressesable)
				}
			}
			m.SetEmailAddresses(res)
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
	res["fax_numbers"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateContactsPutRequestBody_fax_numbersFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ContactsPutRequestBody_fax_numbersable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ContactsPutRequestBody_fax_numbersable)
				}
			}
			m.SetFaxNumbers(res)
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
	res["job_title"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetJobTitle(val)
		}
		return nil
	}
	res["lead_source_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLeadSourceId(val)
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
	res["opt_in_reason"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetOptInReason(val)
		}
		return nil
	}
	res["origin"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateContactsPutRequestBody_originFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetOrigin(val.(ContactsPutRequestBody_originable))
		}
		return nil
	}
	res["owner_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetOwnerId(val)
		}
		return nil
	}
	res["phone_numbers"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateContactsPutRequestBody_phone_numbersFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ContactsPutRequestBody_phone_numbersable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ContactsPutRequestBody_phone_numbersable)
				}
			}
			m.SetPhoneNumbers(res)
		}
		return nil
	}
	res["preferred_locale"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPreferredLocale(val)
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
	res["prefix"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPrefix(val)
		}
		return nil
	}
	res["social_accounts"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateContactsPutRequestBody_social_accountsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ContactsPutRequestBody_social_accountsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ContactsPutRequestBody_social_accountsable)
				}
			}
			m.SetSocialAccounts(res)
		}
		return nil
	}
	res["spouse_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSpouseName(val)
		}
		return nil
	}
	res["suffix"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetSuffix(val)
		}
		return nil
	}
	res["time_zone"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetTimeZone(val)
		}
		return nil
	}
	res["website"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetWebsite(val)
		}
		return nil
	}
	return res
}

// GetGivenName gets the given_name property value. The given_name property
// returns a *string when successful
func (m *ContactsPutRequestBody) GetGivenName() *string {
	return m.given_name
}

// GetJobTitle gets the job_title property value. The job_title property
// returns a *string when successful
func (m *ContactsPutRequestBody) GetJobTitle() *string {
	return m.job_title
}

// GetLeadSourceId gets the lead_source_id property value. The lead_source_id property
// returns a *int64 when successful
func (m *ContactsPutRequestBody) GetLeadSourceId() *int64 {
	return m.lead_source_id
}

// GetMiddleName gets the middle_name property value. The middle_name property
// returns a *string when successful
func (m *ContactsPutRequestBody) GetMiddleName() *string {
	return m.middle_name
}

// GetOptInReason gets the opt_in_reason property value. The opt_in_reason property
// returns a *string when successful
func (m *ContactsPutRequestBody) GetOptInReason() *string {
	return m.opt_in_reason
}

// GetOrigin gets the origin property value. The origin property
// returns a ContactsPutRequestBody_originable when successful
func (m *ContactsPutRequestBody) GetOrigin() ContactsPutRequestBody_originable {
	return m.origin
}

// GetOwnerId gets the owner_id property value. The owner_id property
// returns a *int64 when successful
func (m *ContactsPutRequestBody) GetOwnerId() *int64 {
	return m.owner_id
}

// GetPhoneNumbers gets the phone_numbers property value. The phone_numbers property
// returns a []ContactsPutRequestBody_phone_numbersable when successful
func (m *ContactsPutRequestBody) GetPhoneNumbers() []ContactsPutRequestBody_phone_numbersable {
	return m.phone_numbers
}

// GetPreferredLocale gets the preferred_locale property value. The preferred_locale property
// returns a *string when successful
func (m *ContactsPutRequestBody) GetPreferredLocale() *string {
	return m.preferred_locale
}

// GetPreferredName gets the preferred_name property value. The preferred_name property
// returns a *string when successful
func (m *ContactsPutRequestBody) GetPreferredName() *string {
	return m.preferred_name
}

// GetPrefix gets the prefix property value. The prefix property
// returns a *string when successful
func (m *ContactsPutRequestBody) GetPrefix() *string {
	return m.prefix
}

// GetSocialAccounts gets the social_accounts property value. The social_accounts property
// returns a []ContactsPutRequestBody_social_accountsable when successful
func (m *ContactsPutRequestBody) GetSocialAccounts() []ContactsPutRequestBody_social_accountsable {
	return m.social_accounts
}

// GetSpouseName gets the spouse_name property value. The spouse_name property
// returns a *string when successful
func (m *ContactsPutRequestBody) GetSpouseName() *string {
	return m.spouse_name
}

// GetSuffix gets the suffix property value. The suffix property
// returns a *string when successful
func (m *ContactsPutRequestBody) GetSuffix() *string {
	return m.suffix
}

// GetTimeZone gets the time_zone property value. The time_zone property
// returns a *string when successful
func (m *ContactsPutRequestBody) GetTimeZone() *string {
	return m.time_zone
}

// GetWebsite gets the website property value. The website property
// returns a *string when successful
func (m *ContactsPutRequestBody) GetWebsite() *string {
	return m.website
}

// Serialize serializes information the current object
func (m *ContactsPutRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	if m.GetAddresses() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAddresses()))
		for i, v := range m.GetAddresses() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("addresses", cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteTimeValue("anniversary", m.GetAnniversary())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteTimeValue("birthday", m.GetBirthday())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("company", m.GetCompany())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("contact_type", m.GetContactType())
		if err != nil {
			return err
		}
	}
	if m.GetCustomFields() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCustomFields()))
		for i, v := range m.GetCustomFields() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("custom_fields", cast)
		if err != nil {
			return err
		}
	}
	if m.GetEmailAddresses() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEmailAddresses()))
		for i, v := range m.GetEmailAddresses() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("email_addresses", cast)
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
	if m.GetFaxNumbers() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFaxNumbers()))
		for i, v := range m.GetFaxNumbers() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("fax_numbers", cast)
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
		err := writer.WriteStringValue("job_title", m.GetJobTitle())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("lead_source_id", m.GetLeadSourceId())
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
		err := writer.WriteStringValue("opt_in_reason", m.GetOptInReason())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteObjectValue("origin", m.GetOrigin())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteInt64Value("owner_id", m.GetOwnerId())
		if err != nil {
			return err
		}
	}
	if m.GetPhoneNumbers() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPhoneNumbers()))
		for i, v := range m.GetPhoneNumbers() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("phone_numbers", cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("preferred_locale", m.GetPreferredLocale())
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
		err := writer.WriteStringValue("prefix", m.GetPrefix())
		if err != nil {
			return err
		}
	}
	if m.GetSocialAccounts() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSocialAccounts()))
		for i, v := range m.GetSocialAccounts() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("social_accounts", cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("spouse_name", m.GetSpouseName())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("suffix", m.GetSuffix())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("time_zone", m.GetTimeZone())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("website", m.GetWebsite())
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
func (m *ContactsPutRequestBody) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAddresses sets the addresses property value. The addresses property
func (m *ContactsPutRequestBody) SetAddresses(value []ContactsPutRequestBody_addressesable) {
	m.addresses = value
}

// SetAnniversary sets the anniversary property value. The anniversary property
func (m *ContactsPutRequestBody) SetAnniversary(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	m.anniversary = value
}

// SetBirthday sets the birthday property value. The birthday property
func (m *ContactsPutRequestBody) SetBirthday(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	m.birthday = value
}

// SetCompany sets the company property value. The company property
func (m *ContactsPutRequestBody) SetCompany(value ContactsPutRequestBody_companyable) {
	m.company = value
}

// SetContactType sets the contact_type property value. The contact_type property
func (m *ContactsPutRequestBody) SetContactType(value *string) {
	m.contact_type = value
}

// SetCustomFields sets the custom_fields property value. The custom_fields property
func (m *ContactsPutRequestBody) SetCustomFields(value []ContactsPutRequestBody_custom_fieldsable) {
	m.custom_fields = value
}

// SetEmailAddresses sets the email_addresses property value. The email_addresses property
func (m *ContactsPutRequestBody) SetEmailAddresses(value []ContactsPutRequestBody_email_addressesable) {
	m.email_addresses = value
}

// SetFamilyName sets the family_name property value. The family_name property
func (m *ContactsPutRequestBody) SetFamilyName(value *string) {
	m.family_name = value
}

// SetFaxNumbers sets the fax_numbers property value. The fax_numbers property
func (m *ContactsPutRequestBody) SetFaxNumbers(value []ContactsPutRequestBody_fax_numbersable) {
	m.fax_numbers = value
}

// SetGivenName sets the given_name property value. The given_name property
func (m *ContactsPutRequestBody) SetGivenName(value *string) {
	m.given_name = value
}

// SetJobTitle sets the job_title property value. The job_title property
func (m *ContactsPutRequestBody) SetJobTitle(value *string) {
	m.job_title = value
}

// SetLeadSourceId sets the lead_source_id property value. The lead_source_id property
func (m *ContactsPutRequestBody) SetLeadSourceId(value *int64) {
	m.lead_source_id = value
}

// SetMiddleName sets the middle_name property value. The middle_name property
func (m *ContactsPutRequestBody) SetMiddleName(value *string) {
	m.middle_name = value
}

// SetOptInReason sets the opt_in_reason property value. The opt_in_reason property
func (m *ContactsPutRequestBody) SetOptInReason(value *string) {
	m.opt_in_reason = value
}

// SetOrigin sets the origin property value. The origin property
func (m *ContactsPutRequestBody) SetOrigin(value ContactsPutRequestBody_originable) {
	m.origin = value
}

// SetOwnerId sets the owner_id property value. The owner_id property
func (m *ContactsPutRequestBody) SetOwnerId(value *int64) {
	m.owner_id = value
}

// SetPhoneNumbers sets the phone_numbers property value. The phone_numbers property
func (m *ContactsPutRequestBody) SetPhoneNumbers(value []ContactsPutRequestBody_phone_numbersable) {
	m.phone_numbers = value
}

// SetPreferredLocale sets the preferred_locale property value. The preferred_locale property
func (m *ContactsPutRequestBody) SetPreferredLocale(value *string) {
	m.preferred_locale = value
}

// SetPreferredName sets the preferred_name property value. The preferred_name property
func (m *ContactsPutRequestBody) SetPreferredName(value *string) {
	m.preferred_name = value
}

// SetPrefix sets the prefix property value. The prefix property
func (m *ContactsPutRequestBody) SetPrefix(value *string) {
	m.prefix = value
}

// SetSocialAccounts sets the social_accounts property value. The social_accounts property
func (m *ContactsPutRequestBody) SetSocialAccounts(value []ContactsPutRequestBody_social_accountsable) {
	m.social_accounts = value
}

// SetSpouseName sets the spouse_name property value. The spouse_name property
func (m *ContactsPutRequestBody) SetSpouseName(value *string) {
	m.spouse_name = value
}

// SetSuffix sets the suffix property value. The suffix property
func (m *ContactsPutRequestBody) SetSuffix(value *string) {
	m.suffix = value
}

// SetTimeZone sets the time_zone property value. The time_zone property
func (m *ContactsPutRequestBody) SetTimeZone(value *string) {
	m.time_zone = value
}

// SetWebsite sets the website property value. The website property
func (m *ContactsPutRequestBody) SetWebsite(value *string) {
	m.website = value
}

type ContactsPutRequestBodyable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAddresses() []ContactsPutRequestBody_addressesable
	GetAnniversary() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetBirthday() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetCompany() ContactsPutRequestBody_companyable
	GetContactType() *string
	GetCustomFields() []ContactsPutRequestBody_custom_fieldsable
	GetEmailAddresses() []ContactsPutRequestBody_email_addressesable
	GetFamilyName() *string
	GetFaxNumbers() []ContactsPutRequestBody_fax_numbersable
	GetGivenName() *string
	GetJobTitle() *string
	GetLeadSourceId() *int64
	GetMiddleName() *string
	GetOptInReason() *string
	GetOrigin() ContactsPutRequestBody_originable
	GetOwnerId() *int64
	GetPhoneNumbers() []ContactsPutRequestBody_phone_numbersable
	GetPreferredLocale() *string
	GetPreferredName() *string
	GetPrefix() *string
	GetSocialAccounts() []ContactsPutRequestBody_social_accountsable
	GetSpouseName() *string
	GetSuffix() *string
	GetTimeZone() *string
	GetWebsite() *string
	SetAddresses(value []ContactsPutRequestBody_addressesable)
	SetAnniversary(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	SetBirthday(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	SetCompany(value ContactsPutRequestBody_companyable)
	SetContactType(value *string)
	SetCustomFields(value []ContactsPutRequestBody_custom_fieldsable)
	SetEmailAddresses(value []ContactsPutRequestBody_email_addressesable)
	SetFamilyName(value *string)
	SetFaxNumbers(value []ContactsPutRequestBody_fax_numbersable)
	SetGivenName(value *string)
	SetJobTitle(value *string)
	SetLeadSourceId(value *int64)
	SetMiddleName(value *string)
	SetOptInReason(value *string)
	SetOrigin(value ContactsPutRequestBody_originable)
	SetOwnerId(value *int64)
	SetPhoneNumbers(value []ContactsPutRequestBody_phone_numbersable)
	SetPreferredLocale(value *string)
	SetPreferredName(value *string)
	SetPrefix(value *string)
	SetSocialAccounts(value []ContactsPutRequestBody_social_accountsable)
	SetSpouseName(value *string)
	SetSuffix(value *string)
	SetTimeZone(value *string)
	SetWebsite(value *string)
}
