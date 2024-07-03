package v2

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ContactsItemWithContact_GetResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The addresses property
	addresses []ContactsItemWithContact_GetResponse_addressesable
	// The anniversary_date property
	anniversary_date *string
	// The birth_date property
	birth_date *string
	// The company property
	company ContactsItemWithContact_GetResponse_companyable
	// The contact_type property
	contact_type *string
	// The create_time property
	create_time *string
	// The custom_fields property
	custom_fields []ContactsItemWithContact_GetResponse_custom_fieldsable
	// The email_addresses property
	email_addresses []ContactsItemWithContact_GetResponse_email_addressesable
	// The family_name property
	family_name *string
	// The fax_numbers property
	fax_numbers []ContactsItemWithContact_GetResponse_fax_numbersable
	// The given_name property
	given_name *string
	// The id property
	id *string
	// The job_title property
	job_title *string
	// The leadsource_id property
	leadsource_id *string
	// The links property
	links []ContactsItemWithContact_GetResponse_linksable
	// The middle_name property
	middle_name *string
	// The origin property
	origin ContactsItemWithContact_GetResponse_originable
	// The owner_id property
	owner_id *string
	// The phone_numbers property
	phone_numbers []ContactsItemWithContact_GetResponse_phone_numbersable
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
	social_accounts []ContactsItemWithContact_GetResponse_social_accountsable
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
	utm_parameters []ContactsItemWithContact_GetResponse_utm_parametersable
	// The website property
	website *string
}

// NewContactsItemWithContact_GetResponse instantiates a new ContactsItemWithContact_GetResponse and sets the default values.
func NewContactsItemWithContact_GetResponse() *ContactsItemWithContact_GetResponse {
	m := &ContactsItemWithContact_GetResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateContactsItemWithContact_GetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsItemWithContact_GetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsItemWithContact_GetResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContactsItemWithContact_GetResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAddresses gets the addresses property value. The addresses property
// returns a []ContactsItemWithContact_GetResponse_addressesable when successful
func (m *ContactsItemWithContact_GetResponse) GetAddresses() []ContactsItemWithContact_GetResponse_addressesable {
	return m.addresses
}

// GetAnniversaryDate gets the anniversary_date property value. The anniversary_date property
// returns a *string when successful
func (m *ContactsItemWithContact_GetResponse) GetAnniversaryDate() *string {
	return m.anniversary_date
}

// GetBirthDate gets the birth_date property value. The birth_date property
// returns a *string when successful
func (m *ContactsItemWithContact_GetResponse) GetBirthDate() *string {
	return m.birth_date
}

// GetCompany gets the company property value. The company property
// returns a ContactsItemWithContact_GetResponse_companyable when successful
func (m *ContactsItemWithContact_GetResponse) GetCompany() ContactsItemWithContact_GetResponse_companyable {
	return m.company
}

// GetContactType gets the contact_type property value. The contact_type property
// returns a *string when successful
func (m *ContactsItemWithContact_GetResponse) GetContactType() *string {
	return m.contact_type
}

// GetCreateTime gets the create_time property value. The create_time property
// returns a *string when successful
func (m *ContactsItemWithContact_GetResponse) GetCreateTime() *string {
	return m.create_time
}

// GetCustomFields gets the custom_fields property value. The custom_fields property
// returns a []ContactsItemWithContact_GetResponse_custom_fieldsable when successful
func (m *ContactsItemWithContact_GetResponse) GetCustomFields() []ContactsItemWithContact_GetResponse_custom_fieldsable {
	return m.custom_fields
}

// GetEmailAddresses gets the email_addresses property value. The email_addresses property
// returns a []ContactsItemWithContact_GetResponse_email_addressesable when successful
func (m *ContactsItemWithContact_GetResponse) GetEmailAddresses() []ContactsItemWithContact_GetResponse_email_addressesable {
	return m.email_addresses
}

// GetFamilyName gets the family_name property value. The family_name property
// returns a *string when successful
func (m *ContactsItemWithContact_GetResponse) GetFamilyName() *string {
	return m.family_name
}

// GetFaxNumbers gets the fax_numbers property value. The fax_numbers property
// returns a []ContactsItemWithContact_GetResponse_fax_numbersable when successful
func (m *ContactsItemWithContact_GetResponse) GetFaxNumbers() []ContactsItemWithContact_GetResponse_fax_numbersable {
	return m.fax_numbers
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContactsItemWithContact_GetResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["addresses"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateContactsItemWithContact_GetResponse_addressesFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ContactsItemWithContact_GetResponse_addressesable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ContactsItemWithContact_GetResponse_addressesable)
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
		val, err := n.GetObjectValue(CreateContactsItemWithContact_GetResponse_companyFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCompany(val.(ContactsItemWithContact_GetResponse_companyable))
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
		val, err := n.GetCollectionOfObjectValues(CreateContactsItemWithContact_GetResponse_custom_fieldsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ContactsItemWithContact_GetResponse_custom_fieldsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ContactsItemWithContact_GetResponse_custom_fieldsable)
				}
			}
			m.SetCustomFields(res)
		}
		return nil
	}
	res["email_addresses"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateContactsItemWithContact_GetResponse_email_addressesFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ContactsItemWithContact_GetResponse_email_addressesable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ContactsItemWithContact_GetResponse_email_addressesable)
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
		val, err := n.GetCollectionOfObjectValues(CreateContactsItemWithContact_GetResponse_fax_numbersFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ContactsItemWithContact_GetResponse_fax_numbersable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ContactsItemWithContact_GetResponse_fax_numbersable)
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
		val, err := n.GetCollectionOfObjectValues(CreateContactsItemWithContact_GetResponse_linksFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ContactsItemWithContact_GetResponse_linksable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ContactsItemWithContact_GetResponse_linksable)
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
		val, err := n.GetObjectValue(CreateContactsItemWithContact_GetResponse_originFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetOrigin(val.(ContactsItemWithContact_GetResponse_originable))
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
		val, err := n.GetCollectionOfObjectValues(CreateContactsItemWithContact_GetResponse_phone_numbersFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ContactsItemWithContact_GetResponse_phone_numbersable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ContactsItemWithContact_GetResponse_phone_numbersable)
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
		val, err := n.GetCollectionOfObjectValues(CreateContactsItemWithContact_GetResponse_social_accountsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ContactsItemWithContact_GetResponse_social_accountsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ContactsItemWithContact_GetResponse_social_accountsable)
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
		val, err := n.GetCollectionOfObjectValues(CreateContactsItemWithContact_GetResponse_utm_parametersFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ContactsItemWithContact_GetResponse_utm_parametersable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ContactsItemWithContact_GetResponse_utm_parametersable)
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
func (m *ContactsItemWithContact_GetResponse) GetGivenName() *string {
	return m.given_name
}

// GetId gets the id property value. The id property
// returns a *string when successful
func (m *ContactsItemWithContact_GetResponse) GetId() *string {
	return m.id
}

// GetJobTitle gets the job_title property value. The job_title property
// returns a *string when successful
func (m *ContactsItemWithContact_GetResponse) GetJobTitle() *string {
	return m.job_title
}

// GetLeadsourceId gets the leadsource_id property value. The leadsource_id property
// returns a *string when successful
func (m *ContactsItemWithContact_GetResponse) GetLeadsourceId() *string {
	return m.leadsource_id
}

// GetLinks gets the links property value. The links property
// returns a []ContactsItemWithContact_GetResponse_linksable when successful
func (m *ContactsItemWithContact_GetResponse) GetLinks() []ContactsItemWithContact_GetResponse_linksable {
	return m.links
}

// GetMiddleName gets the middle_name property value. The middle_name property
// returns a *string when successful
func (m *ContactsItemWithContact_GetResponse) GetMiddleName() *string {
	return m.middle_name
}

// GetOrigin gets the origin property value. The origin property
// returns a ContactsItemWithContact_GetResponse_originable when successful
func (m *ContactsItemWithContact_GetResponse) GetOrigin() ContactsItemWithContact_GetResponse_originable {
	return m.origin
}

// GetOwnerId gets the owner_id property value. The owner_id property
// returns a *string when successful
func (m *ContactsItemWithContact_GetResponse) GetOwnerId() *string {
	return m.owner_id
}

// GetPhoneNumbers gets the phone_numbers property value. The phone_numbers property
// returns a []ContactsItemWithContact_GetResponse_phone_numbersable when successful
func (m *ContactsItemWithContact_GetResponse) GetPhoneNumbers() []ContactsItemWithContact_GetResponse_phone_numbersable {
	return m.phone_numbers
}

// GetPreferredLocale gets the preferred_locale property value. The preferred_locale property
// returns a *string when successful
func (m *ContactsItemWithContact_GetResponse) GetPreferredLocale() *string {
	return m.preferred_locale
}

// GetPreferredName gets the preferred_name property value. The preferred_name property
// returns a *string when successful
func (m *ContactsItemWithContact_GetResponse) GetPreferredName() *string {
	return m.preferred_name
}

// GetPrefix gets the prefix property value. The prefix property
// returns a *string when successful
func (m *ContactsItemWithContact_GetResponse) GetPrefix() *string {
	return m.prefix
}

// GetReferralCode gets the referral_code property value. The referral_code property
// returns a *string when successful
func (m *ContactsItemWithContact_GetResponse) GetReferralCode() *string {
	return m.referral_code
}

// GetScoreValue gets the score_value property value. The score_value property
// returns a *string when successful
func (m *ContactsItemWithContact_GetResponse) GetScoreValue() *string {
	return m.score_value
}

// GetSocialAccounts gets the social_accounts property value. The social_accounts property
// returns a []ContactsItemWithContact_GetResponse_social_accountsable when successful
func (m *ContactsItemWithContact_GetResponse) GetSocialAccounts() []ContactsItemWithContact_GetResponse_social_accountsable {
	return m.social_accounts
}

// GetSpouseName gets the spouse_name property value. The spouse_name property
// returns a *string when successful
func (m *ContactsItemWithContact_GetResponse) GetSpouseName() *string {
	return m.spouse_name
}

// GetSuffix gets the suffix property value. The suffix property
// returns a *string when successful
func (m *ContactsItemWithContact_GetResponse) GetSuffix() *string {
	return m.suffix
}

// GetTagIds gets the tag_ids property value. The tag_ids property
// returns a []string when successful
func (m *ContactsItemWithContact_GetResponse) GetTagIds() []string {
	return m.tag_ids
}

// GetTimeZone gets the time_zone property value. The time_zone property
// returns a *string when successful
func (m *ContactsItemWithContact_GetResponse) GetTimeZone() *string {
	return m.time_zone
}

// GetUpdateTime gets the update_time property value. The update_time property
// returns a *string when successful
func (m *ContactsItemWithContact_GetResponse) GetUpdateTime() *string {
	return m.update_time
}

// GetUtmParameters gets the utm_parameters property value. The utm_parameters property
// returns a []ContactsItemWithContact_GetResponse_utm_parametersable when successful
func (m *ContactsItemWithContact_GetResponse) GetUtmParameters() []ContactsItemWithContact_GetResponse_utm_parametersable {
	return m.utm_parameters
}

// GetWebsite gets the website property value. The website property
// returns a *string when successful
func (m *ContactsItemWithContact_GetResponse) GetWebsite() *string {
	return m.website
}

// Serialize serializes information the current object
func (m *ContactsItemWithContact_GetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
func (m *ContactsItemWithContact_GetResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAddresses sets the addresses property value. The addresses property
func (m *ContactsItemWithContact_GetResponse) SetAddresses(value []ContactsItemWithContact_GetResponse_addressesable) {
	m.addresses = value
}

// SetAnniversaryDate sets the anniversary_date property value. The anniversary_date property
func (m *ContactsItemWithContact_GetResponse) SetAnniversaryDate(value *string) {
	m.anniversary_date = value
}

// SetBirthDate sets the birth_date property value. The birth_date property
func (m *ContactsItemWithContact_GetResponse) SetBirthDate(value *string) {
	m.birth_date = value
}

// SetCompany sets the company property value. The company property
func (m *ContactsItemWithContact_GetResponse) SetCompany(value ContactsItemWithContact_GetResponse_companyable) {
	m.company = value
}

// SetContactType sets the contact_type property value. The contact_type property
func (m *ContactsItemWithContact_GetResponse) SetContactType(value *string) {
	m.contact_type = value
}

// SetCreateTime sets the create_time property value. The create_time property
func (m *ContactsItemWithContact_GetResponse) SetCreateTime(value *string) {
	m.create_time = value
}

// SetCustomFields sets the custom_fields property value. The custom_fields property
func (m *ContactsItemWithContact_GetResponse) SetCustomFields(value []ContactsItemWithContact_GetResponse_custom_fieldsable) {
	m.custom_fields = value
}

// SetEmailAddresses sets the email_addresses property value. The email_addresses property
func (m *ContactsItemWithContact_GetResponse) SetEmailAddresses(value []ContactsItemWithContact_GetResponse_email_addressesable) {
	m.email_addresses = value
}

// SetFamilyName sets the family_name property value. The family_name property
func (m *ContactsItemWithContact_GetResponse) SetFamilyName(value *string) {
	m.family_name = value
}

// SetFaxNumbers sets the fax_numbers property value. The fax_numbers property
func (m *ContactsItemWithContact_GetResponse) SetFaxNumbers(value []ContactsItemWithContact_GetResponse_fax_numbersable) {
	m.fax_numbers = value
}

// SetGivenName sets the given_name property value. The given_name property
func (m *ContactsItemWithContact_GetResponse) SetGivenName(value *string) {
	m.given_name = value
}

// SetId sets the id property value. The id property
func (m *ContactsItemWithContact_GetResponse) SetId(value *string) {
	m.id = value
}

// SetJobTitle sets the job_title property value. The job_title property
func (m *ContactsItemWithContact_GetResponse) SetJobTitle(value *string) {
	m.job_title = value
}

// SetLeadsourceId sets the leadsource_id property value. The leadsource_id property
func (m *ContactsItemWithContact_GetResponse) SetLeadsourceId(value *string) {
	m.leadsource_id = value
}

// SetLinks sets the links property value. The links property
func (m *ContactsItemWithContact_GetResponse) SetLinks(value []ContactsItemWithContact_GetResponse_linksable) {
	m.links = value
}

// SetMiddleName sets the middle_name property value. The middle_name property
func (m *ContactsItemWithContact_GetResponse) SetMiddleName(value *string) {
	m.middle_name = value
}

// SetOrigin sets the origin property value. The origin property
func (m *ContactsItemWithContact_GetResponse) SetOrigin(value ContactsItemWithContact_GetResponse_originable) {
	m.origin = value
}

// SetOwnerId sets the owner_id property value. The owner_id property
func (m *ContactsItemWithContact_GetResponse) SetOwnerId(value *string) {
	m.owner_id = value
}

// SetPhoneNumbers sets the phone_numbers property value. The phone_numbers property
func (m *ContactsItemWithContact_GetResponse) SetPhoneNumbers(value []ContactsItemWithContact_GetResponse_phone_numbersable) {
	m.phone_numbers = value
}

// SetPreferredLocale sets the preferred_locale property value. The preferred_locale property
func (m *ContactsItemWithContact_GetResponse) SetPreferredLocale(value *string) {
	m.preferred_locale = value
}

// SetPreferredName sets the preferred_name property value. The preferred_name property
func (m *ContactsItemWithContact_GetResponse) SetPreferredName(value *string) {
	m.preferred_name = value
}

// SetPrefix sets the prefix property value. The prefix property
func (m *ContactsItemWithContact_GetResponse) SetPrefix(value *string) {
	m.prefix = value
}

// SetReferralCode sets the referral_code property value. The referral_code property
func (m *ContactsItemWithContact_GetResponse) SetReferralCode(value *string) {
	m.referral_code = value
}

// SetScoreValue sets the score_value property value. The score_value property
func (m *ContactsItemWithContact_GetResponse) SetScoreValue(value *string) {
	m.score_value = value
}

// SetSocialAccounts sets the social_accounts property value. The social_accounts property
func (m *ContactsItemWithContact_GetResponse) SetSocialAccounts(value []ContactsItemWithContact_GetResponse_social_accountsable) {
	m.social_accounts = value
}

// SetSpouseName sets the spouse_name property value. The spouse_name property
func (m *ContactsItemWithContact_GetResponse) SetSpouseName(value *string) {
	m.spouse_name = value
}

// SetSuffix sets the suffix property value. The suffix property
func (m *ContactsItemWithContact_GetResponse) SetSuffix(value *string) {
	m.suffix = value
}

// SetTagIds sets the tag_ids property value. The tag_ids property
func (m *ContactsItemWithContact_GetResponse) SetTagIds(value []string) {
	m.tag_ids = value
}

// SetTimeZone sets the time_zone property value. The time_zone property
func (m *ContactsItemWithContact_GetResponse) SetTimeZone(value *string) {
	m.time_zone = value
}

// SetUpdateTime sets the update_time property value. The update_time property
func (m *ContactsItemWithContact_GetResponse) SetUpdateTime(value *string) {
	m.update_time = value
}

// SetUtmParameters sets the utm_parameters property value. The utm_parameters property
func (m *ContactsItemWithContact_GetResponse) SetUtmParameters(value []ContactsItemWithContact_GetResponse_utm_parametersable) {
	m.utm_parameters = value
}

// SetWebsite sets the website property value. The website property
func (m *ContactsItemWithContact_GetResponse) SetWebsite(value *string) {
	m.website = value
}

type ContactsItemWithContact_GetResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAddresses() []ContactsItemWithContact_GetResponse_addressesable
	GetAnniversaryDate() *string
	GetBirthDate() *string
	GetCompany() ContactsItemWithContact_GetResponse_companyable
	GetContactType() *string
	GetCreateTime() *string
	GetCustomFields() []ContactsItemWithContact_GetResponse_custom_fieldsable
	GetEmailAddresses() []ContactsItemWithContact_GetResponse_email_addressesable
	GetFamilyName() *string
	GetFaxNumbers() []ContactsItemWithContact_GetResponse_fax_numbersable
	GetGivenName() *string
	GetId() *string
	GetJobTitle() *string
	GetLeadsourceId() *string
	GetLinks() []ContactsItemWithContact_GetResponse_linksable
	GetMiddleName() *string
	GetOrigin() ContactsItemWithContact_GetResponse_originable
	GetOwnerId() *string
	GetPhoneNumbers() []ContactsItemWithContact_GetResponse_phone_numbersable
	GetPreferredLocale() *string
	GetPreferredName() *string
	GetPrefix() *string
	GetReferralCode() *string
	GetScoreValue() *string
	GetSocialAccounts() []ContactsItemWithContact_GetResponse_social_accountsable
	GetSpouseName() *string
	GetSuffix() *string
	GetTagIds() []string
	GetTimeZone() *string
	GetUpdateTime() *string
	GetUtmParameters() []ContactsItemWithContact_GetResponse_utm_parametersable
	GetWebsite() *string
	SetAddresses(value []ContactsItemWithContact_GetResponse_addressesable)
	SetAnniversaryDate(value *string)
	SetBirthDate(value *string)
	SetCompany(value ContactsItemWithContact_GetResponse_companyable)
	SetContactType(value *string)
	SetCreateTime(value *string)
	SetCustomFields(value []ContactsItemWithContact_GetResponse_custom_fieldsable)
	SetEmailAddresses(value []ContactsItemWithContact_GetResponse_email_addressesable)
	SetFamilyName(value *string)
	SetFaxNumbers(value []ContactsItemWithContact_GetResponse_fax_numbersable)
	SetGivenName(value *string)
	SetId(value *string)
	SetJobTitle(value *string)
	SetLeadsourceId(value *string)
	SetLinks(value []ContactsItemWithContact_GetResponse_linksable)
	SetMiddleName(value *string)
	SetOrigin(value ContactsItemWithContact_GetResponse_originable)
	SetOwnerId(value *string)
	SetPhoneNumbers(value []ContactsItemWithContact_GetResponse_phone_numbersable)
	SetPreferredLocale(value *string)
	SetPreferredName(value *string)
	SetPrefix(value *string)
	SetReferralCode(value *string)
	SetScoreValue(value *string)
	SetSocialAccounts(value []ContactsItemWithContact_GetResponse_social_accountsable)
	SetSpouseName(value *string)
	SetSuffix(value *string)
	SetTagIds(value []string)
	SetTimeZone(value *string)
	SetUpdateTime(value *string)
	SetUtmParameters(value []ContactsItemWithContact_GetResponse_utm_parametersable)
	SetWebsite(value *string)
}
