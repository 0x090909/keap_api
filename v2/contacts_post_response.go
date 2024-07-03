package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ContactsPostResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The addresses property
	addresses []ContactsPostResponse_addressesable
	// The anniversary_date property
	anniversary_date *string
	// The birth_date property
	birth_date *string
	// The company property
	company ContactsPostResponse_companyable
	// The contact_type property
	contact_type *string
	// The create_time property
	create_time *string
	// The custom_fields property
	custom_fields []ContactsPostResponse_custom_fieldsable
	// The email_addresses property
	email_addresses []ContactsPostResponse_email_addressesable
	// The family_name property
	family_name *string
	// The fax_numbers property
	fax_numbers []ContactsPostResponse_fax_numbersable
	// The given_name property
	given_name *string
	// The id property
	id *string
	// The job_title property
	job_title *string
	// The leadsource_id property
	leadsource_id *string
	// The links property
	links []ContactsPostResponse_linksable
	// The middle_name property
	middle_name *string
	// The origin property
	origin ContactsPostResponse_originable
	// The owner_id property
	owner_id *string
	// The phone_numbers property
	phone_numbers []ContactsPostResponse_phone_numbersable
	// The preferred_locale property
	preferred_locale *string
	// The preferred_name property
	preferred_name *string
	// The prefix property
	prefix *string
	// The referral_code property
	referral_code *string
	// The score_value property
	score_value *string
	// The social_accounts property
	social_accounts []ContactsPostResponse_social_accountsable
	// The spouse_name property
	spouse_name *string
	// The suffix property
	suffix *string
	// The tag_ids property
	tag_ids []string
	// The time_zone property
	time_zone *string
	// The update_time property
	update_time *string
	// The utm_parameters property
	utm_parameters []ContactsPostResponse_utm_parametersable
	// The website property
	website *string
}

// NewContactsPostResponse instantiates a new ContactsPostResponse and sets the default values.
func NewContactsPostResponse() *ContactsPostResponse {
	m := &ContactsPostResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateContactsPostResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsPostResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsPostResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContactsPostResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAddresses gets the addresses property value. The addresses property
// returns a []ContactsPostResponse_addressesable when successful
func (m *ContactsPostResponse) GetAddresses() []ContactsPostResponse_addressesable {
	return m.addresses
}

// GetAnniversaryDate gets the anniversary_date property value. The anniversary_date property
// returns a *string when successful
func (m *ContactsPostResponse) GetAnniversaryDate() *string {
	return m.anniversary_date
}

// GetBirthDate gets the birth_date property value. The birth_date property
// returns a *string when successful
func (m *ContactsPostResponse) GetBirthDate() *string {
	return m.birth_date
}

// GetCompany gets the company property value. The company property
// returns a ContactsPostResponse_companyable when successful
func (m *ContactsPostResponse) GetCompany() ContactsPostResponse_companyable {
	return m.company
}

// GetContactType gets the contact_type property value. The contact_type property
// returns a *string when successful
func (m *ContactsPostResponse) GetContactType() *string {
	return m.contact_type
}

// GetCreateTime gets the create_time property value. The create_time property
// returns a *string when successful
func (m *ContactsPostResponse) GetCreateTime() *string {
	return m.create_time
}

// GetCustomFields gets the custom_fields property value. The custom_fields property
// returns a []ContactsPostResponse_custom_fieldsable when successful
func (m *ContactsPostResponse) GetCustomFields() []ContactsPostResponse_custom_fieldsable {
	return m.custom_fields
}

// GetEmailAddresses gets the email_addresses property value. The email_addresses property
// returns a []ContactsPostResponse_email_addressesable when successful
func (m *ContactsPostResponse) GetEmailAddresses() []ContactsPostResponse_email_addressesable {
	return m.email_addresses
}

// GetFamilyName gets the family_name property value. The family_name property
// returns a *string when successful
func (m *ContactsPostResponse) GetFamilyName() *string {
	return m.family_name
}

// GetFaxNumbers gets the fax_numbers property value. The fax_numbers property
// returns a []ContactsPostResponse_fax_numbersable when successful
func (m *ContactsPostResponse) GetFaxNumbers() []ContactsPostResponse_fax_numbersable {
	return m.fax_numbers
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContactsPostResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["addresses"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateContactsPostResponse_addressesFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ContactsPostResponse_addressesable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ContactsPostResponse_addressesable)
				}
			}
			m.SetAddresses(res)
		}
		return nil
	}
	res["anniversary_date"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAnniversaryDate(val)
		}
		return nil
	}
	res["birth_date"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetBirthDate(val)
		}
		return nil
	}
	res["company"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateContactsPostResponse_companyFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCompany(val.(ContactsPostResponse_companyable))
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
	res["create_time"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCreateTime(val)
		}
		return nil
	}
	res["custom_fields"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateContactsPostResponse_custom_fieldsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ContactsPostResponse_custom_fieldsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ContactsPostResponse_custom_fieldsable)
				}
			}
			m.SetCustomFields(res)
		}
		return nil
	}
	res["email_addresses"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateContactsPostResponse_email_addressesFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ContactsPostResponse_email_addressesable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ContactsPostResponse_email_addressesable)
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
		val, err := n.GetCollectionOfObjectValues(CreateContactsPostResponse_fax_numbersFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ContactsPostResponse_fax_numbersable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ContactsPostResponse_fax_numbersable)
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
	res["leadsource_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLeadsourceId(val)
		}
		return nil
	}
	res["links"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateContactsPostResponse_linksFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ContactsPostResponse_linksable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ContactsPostResponse_linksable)
				}
			}
			m.SetLinks(res)
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
	res["origin"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateContactsPostResponse_originFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetOrigin(val.(ContactsPostResponse_originable))
		}
		return nil
	}
	res["owner_id"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetOwnerId(val)
		}
		return nil
	}
	res["phone_numbers"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateContactsPostResponse_phone_numbersFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ContactsPostResponse_phone_numbersable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ContactsPostResponse_phone_numbersable)
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
	res["referral_code"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetReferralCode(val)
		}
		return nil
	}
	res["score_value"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetScoreValue(val)
		}
		return nil
	}
	res["social_accounts"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateContactsPostResponse_social_accountsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ContactsPostResponse_social_accountsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ContactsPostResponse_social_accountsable)
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
	res["tag_ids"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfPrimitiveValues("string")
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]string, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = *(v.(*string))
				}
			}
			m.SetTagIds(res)
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
	res["update_time"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetUpdateTime(val)
		}
		return nil
	}
	res["utm_parameters"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateContactsPostResponse_utm_parametersFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ContactsPostResponse_utm_parametersable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ContactsPostResponse_utm_parametersable)
				}
			}
			m.SetUtmParameters(res)
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
func (m *ContactsPostResponse) GetGivenName() *string {
	return m.given_name
}

// GetId gets the id property value. The id property
// returns a *string when successful
func (m *ContactsPostResponse) GetId() *string {
	return m.id
}

// GetJobTitle gets the job_title property value. The job_title property
// returns a *string when successful
func (m *ContactsPostResponse) GetJobTitle() *string {
	return m.job_title
}

// GetLeadsourceId gets the leadsource_id property value. The leadsource_id property
// returns a *string when successful
func (m *ContactsPostResponse) GetLeadsourceId() *string {
	return m.leadsource_id
}

// GetLinks gets the links property value. The links property
// returns a []ContactsPostResponse_linksable when successful
func (m *ContactsPostResponse) GetLinks() []ContactsPostResponse_linksable {
	return m.links
}

// GetMiddleName gets the middle_name property value. The middle_name property
// returns a *string when successful
func (m *ContactsPostResponse) GetMiddleName() *string {
	return m.middle_name
}

// GetOrigin gets the origin property value. The origin property
// returns a ContactsPostResponse_originable when successful
func (m *ContactsPostResponse) GetOrigin() ContactsPostResponse_originable {
	return m.origin
}

// GetOwnerId gets the owner_id property value. The owner_id property
// returns a *string when successful
func (m *ContactsPostResponse) GetOwnerId() *string {
	return m.owner_id
}

// GetPhoneNumbers gets the phone_numbers property value. The phone_numbers property
// returns a []ContactsPostResponse_phone_numbersable when successful
func (m *ContactsPostResponse) GetPhoneNumbers() []ContactsPostResponse_phone_numbersable {
	return m.phone_numbers
}

// GetPreferredLocale gets the preferred_locale property value. The preferred_locale property
// returns a *string when successful
func (m *ContactsPostResponse) GetPreferredLocale() *string {
	return m.preferred_locale
}

// GetPreferredName gets the preferred_name property value. The preferred_name property
// returns a *string when successful
func (m *ContactsPostResponse) GetPreferredName() *string {
	return m.preferred_name
}

// GetPrefix gets the prefix property value. The prefix property
// returns a *string when successful
func (m *ContactsPostResponse) GetPrefix() *string {
	return m.prefix
}

// GetReferralCode gets the referral_code property value. The referral_code property
// returns a *string when successful
func (m *ContactsPostResponse) GetReferralCode() *string {
	return m.referral_code
}

// GetScoreValue gets the score_value property value. The score_value property
// returns a *string when successful
func (m *ContactsPostResponse) GetScoreValue() *string {
	return m.score_value
}

// GetSocialAccounts gets the social_accounts property value. The social_accounts property
// returns a []ContactsPostResponse_social_accountsable when successful
func (m *ContactsPostResponse) GetSocialAccounts() []ContactsPostResponse_social_accountsable {
	return m.social_accounts
}

// GetSpouseName gets the spouse_name property value. The spouse_name property
// returns a *string when successful
func (m *ContactsPostResponse) GetSpouseName() *string {
	return m.spouse_name
}

// GetSuffix gets the suffix property value. The suffix property
// returns a *string when successful
func (m *ContactsPostResponse) GetSuffix() *string {
	return m.suffix
}

// GetTagIds gets the tag_ids property value. The tag_ids property
// returns a []string when successful
func (m *ContactsPostResponse) GetTagIds() []string {
	return m.tag_ids
}

// GetTimeZone gets the time_zone property value. The time_zone property
// returns a *string when successful
func (m *ContactsPostResponse) GetTimeZone() *string {
	return m.time_zone
}

// GetUpdateTime gets the update_time property value. The update_time property
// returns a *string when successful
func (m *ContactsPostResponse) GetUpdateTime() *string {
	return m.update_time
}

// GetUtmParameters gets the utm_parameters property value. The utm_parameters property
// returns a []ContactsPostResponse_utm_parametersable when successful
func (m *ContactsPostResponse) GetUtmParameters() []ContactsPostResponse_utm_parametersable {
	return m.utm_parameters
}

// GetWebsite gets the website property value. The website property
// returns a *string when successful
func (m *ContactsPostResponse) GetWebsite() *string {
	return m.website
}

// Serialize serializes information the current object
func (m *ContactsPostResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
		err := writer.WriteStringValue("anniversary_date", m.GetAnniversaryDate())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("birth_date", m.GetBirthDate())
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
	{
		err := writer.WriteStringValue("create_time", m.GetCreateTime())
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
		err := writer.WriteStringValue("id", m.GetId())
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
		err := writer.WriteStringValue("leadsource_id", m.GetLeadsourceId())
		if err != nil {
			return err
		}
	}
	if m.GetLinks() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLinks()))
		for i, v := range m.GetLinks() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("links", cast)
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
		err := writer.WriteObjectValue("origin", m.GetOrigin())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("owner_id", m.GetOwnerId())
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
	{
		err := writer.WriteStringValue("referral_code", m.GetReferralCode())
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("score_value", m.GetScoreValue())
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
	if m.GetTagIds() != nil {
		err := writer.WriteCollectionOfStringValues("tag_ids", m.GetTagIds())
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
		err := writer.WriteStringValue("update_time", m.GetUpdateTime())
		if err != nil {
			return err
		}
	}
	if m.GetUtmParameters() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUtmParameters()))
		for i, v := range m.GetUtmParameters() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("utm_parameters", cast)
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
func (m *ContactsPostResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAddresses sets the addresses property value. The addresses property
func (m *ContactsPostResponse) SetAddresses(value []ContactsPostResponse_addressesable) {
	m.addresses = value
}

// SetAnniversaryDate sets the anniversary_date property value. The anniversary_date property
func (m *ContactsPostResponse) SetAnniversaryDate(value *string) {
	m.anniversary_date = value
}

// SetBirthDate sets the birth_date property value. The birth_date property
func (m *ContactsPostResponse) SetBirthDate(value *string) {
	m.birth_date = value
}

// SetCompany sets the company property value. The company property
func (m *ContactsPostResponse) SetCompany(value ContactsPostResponse_companyable) {
	m.company = value
}

// SetContactType sets the contact_type property value. The contact_type property
func (m *ContactsPostResponse) SetContactType(value *string) {
	m.contact_type = value
}

// SetCreateTime sets the create_time property value. The create_time property
func (m *ContactsPostResponse) SetCreateTime(value *string) {
	m.create_time = value
}

// SetCustomFields sets the custom_fields property value. The custom_fields property
func (m *ContactsPostResponse) SetCustomFields(value []ContactsPostResponse_custom_fieldsable) {
	m.custom_fields = value
}

// SetEmailAddresses sets the email_addresses property value. The email_addresses property
func (m *ContactsPostResponse) SetEmailAddresses(value []ContactsPostResponse_email_addressesable) {
	m.email_addresses = value
}

// SetFamilyName sets the family_name property value. The family_name property
func (m *ContactsPostResponse) SetFamilyName(value *string) {
	m.family_name = value
}

// SetFaxNumbers sets the fax_numbers property value. The fax_numbers property
func (m *ContactsPostResponse) SetFaxNumbers(value []ContactsPostResponse_fax_numbersable) {
	m.fax_numbers = value
}

// SetGivenName sets the given_name property value. The given_name property
func (m *ContactsPostResponse) SetGivenName(value *string) {
	m.given_name = value
}

// SetId sets the id property value. The id property
func (m *ContactsPostResponse) SetId(value *string) {
	m.id = value
}

// SetJobTitle sets the job_title property value. The job_title property
func (m *ContactsPostResponse) SetJobTitle(value *string) {
	m.job_title = value
}

// SetLeadsourceId sets the leadsource_id property value. The leadsource_id property
func (m *ContactsPostResponse) SetLeadsourceId(value *string) {
	m.leadsource_id = value
}

// SetLinks sets the links property value. The links property
func (m *ContactsPostResponse) SetLinks(value []ContactsPostResponse_linksable) {
	m.links = value
}

// SetMiddleName sets the middle_name property value. The middle_name property
func (m *ContactsPostResponse) SetMiddleName(value *string) {
	m.middle_name = value
}

// SetOrigin sets the origin property value. The origin property
func (m *ContactsPostResponse) SetOrigin(value ContactsPostResponse_originable) {
	m.origin = value
}

// SetOwnerId sets the owner_id property value. The owner_id property
func (m *ContactsPostResponse) SetOwnerId(value *string) {
	m.owner_id = value
}

// SetPhoneNumbers sets the phone_numbers property value. The phone_numbers property
func (m *ContactsPostResponse) SetPhoneNumbers(value []ContactsPostResponse_phone_numbersable) {
	m.phone_numbers = value
}

// SetPreferredLocale sets the preferred_locale property value. The preferred_locale property
func (m *ContactsPostResponse) SetPreferredLocale(value *string) {
	m.preferred_locale = value
}

// SetPreferredName sets the preferred_name property value. The preferred_name property
func (m *ContactsPostResponse) SetPreferredName(value *string) {
	m.preferred_name = value
}

// SetPrefix sets the prefix property value. The prefix property
func (m *ContactsPostResponse) SetPrefix(value *string) {
	m.prefix = value
}

// SetReferralCode sets the referral_code property value. The referral_code property
func (m *ContactsPostResponse) SetReferralCode(value *string) {
	m.referral_code = value
}

// SetScoreValue sets the score_value property value. The score_value property
func (m *ContactsPostResponse) SetScoreValue(value *string) {
	m.score_value = value
}

// SetSocialAccounts sets the social_accounts property value. The social_accounts property
func (m *ContactsPostResponse) SetSocialAccounts(value []ContactsPostResponse_social_accountsable) {
	m.social_accounts = value
}

// SetSpouseName sets the spouse_name property value. The spouse_name property
func (m *ContactsPostResponse) SetSpouseName(value *string) {
	m.spouse_name = value
}

// SetSuffix sets the suffix property value. The suffix property
func (m *ContactsPostResponse) SetSuffix(value *string) {
	m.suffix = value
}

// SetTagIds sets the tag_ids property value. The tag_ids property
func (m *ContactsPostResponse) SetTagIds(value []string) {
	m.tag_ids = value
}

// SetTimeZone sets the time_zone property value. The time_zone property
func (m *ContactsPostResponse) SetTimeZone(value *string) {
	m.time_zone = value
}

// SetUpdateTime sets the update_time property value. The update_time property
func (m *ContactsPostResponse) SetUpdateTime(value *string) {
	m.update_time = value
}

// SetUtmParameters sets the utm_parameters property value. The utm_parameters property
func (m *ContactsPostResponse) SetUtmParameters(value []ContactsPostResponse_utm_parametersable) {
	m.utm_parameters = value
}

// SetWebsite sets the website property value. The website property
func (m *ContactsPostResponse) SetWebsite(value *string) {
	m.website = value
}

type ContactsPostResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAddresses() []ContactsPostResponse_addressesable
	GetAnniversaryDate() *string
	GetBirthDate() *string
	GetCompany() ContactsPostResponse_companyable
	GetContactType() *string
	GetCreateTime() *string
	GetCustomFields() []ContactsPostResponse_custom_fieldsable
	GetEmailAddresses() []ContactsPostResponse_email_addressesable
	GetFamilyName() *string
	GetFaxNumbers() []ContactsPostResponse_fax_numbersable
	GetGivenName() *string
	GetId() *string
	GetJobTitle() *string
	GetLeadsourceId() *string
	GetLinks() []ContactsPostResponse_linksable
	GetMiddleName() *string
	GetOrigin() ContactsPostResponse_originable
	GetOwnerId() *string
	GetPhoneNumbers() []ContactsPostResponse_phone_numbersable
	GetPreferredLocale() *string
	GetPreferredName() *string
	GetPrefix() *string
	GetReferralCode() *string
	GetScoreValue() *string
	GetSocialAccounts() []ContactsPostResponse_social_accountsable
	GetSpouseName() *string
	GetSuffix() *string
	GetTagIds() []string
	GetTimeZone() *string
	GetUpdateTime() *string
	GetUtmParameters() []ContactsPostResponse_utm_parametersable
	GetWebsite() *string
	SetAddresses(value []ContactsPostResponse_addressesable)
	SetAnniversaryDate(value *string)
	SetBirthDate(value *string)
	SetCompany(value ContactsPostResponse_companyable)
	SetContactType(value *string)
	SetCreateTime(value *string)
	SetCustomFields(value []ContactsPostResponse_custom_fieldsable)
	SetEmailAddresses(value []ContactsPostResponse_email_addressesable)
	SetFamilyName(value *string)
	SetFaxNumbers(value []ContactsPostResponse_fax_numbersable)
	SetGivenName(value *string)
	SetId(value *string)
	SetJobTitle(value *string)
	SetLeadsourceId(value *string)
	SetLinks(value []ContactsPostResponse_linksable)
	SetMiddleName(value *string)
	SetOrigin(value ContactsPostResponse_originable)
	SetOwnerId(value *string)
	SetPhoneNumbers(value []ContactsPostResponse_phone_numbersable)
	SetPreferredLocale(value *string)
	SetPreferredName(value *string)
	SetPrefix(value *string)
	SetReferralCode(value *string)
	SetScoreValue(value *string)
	SetSocialAccounts(value []ContactsPostResponse_social_accountsable)
	SetSpouseName(value *string)
	SetSuffix(value *string)
	SetTagIds(value []string)
	SetTimeZone(value *string)
	SetUpdateTime(value *string)
	SetUtmParameters(value []ContactsPostResponse_utm_parametersable)
	SetWebsite(value *string)
}