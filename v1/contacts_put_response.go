package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
)

type ContactsPutResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The addresses property
	addresses []ContactsPutResponse_addressesable
	// The anniversary property
	anniversary *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	// The birthday property
	birthday *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	// The company property
	company ContactsPutResponse_companyable
	// The company_name property
	company_name *string
	// The contact_type property
	contact_type *string
	// The custom_fields property
	custom_fields []ContactsPutResponse_custom_fieldsable
	// The date_created property
	date_created *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	// The email_addresses property
	email_addresses []ContactsPutResponse_email_addressesable
	// The email_opted_in property
	email_opted_in *bool
	// The family_name property
	family_name *string
	// The fax_numbers property
	fax_numbers []ContactsPutResponse_fax_numbersable
	// The given_name property
	given_name *string
	// The id property
	id *int64
	// The job_title property
	job_title *string
	// The last_updated property
	last_updated *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	// The lead_source_id property
	lead_source_id *int64
	// The middle_name property
	middle_name *string
	// The opt_in_reason property
	opt_in_reason *string
	// The origin property
	origin ContactsPutResponse_originable
	// The owner_id property
	owner_id *int64
	// The phone_numbers property
	phone_numbers []ContactsPutResponse_phone_numbersable
	// The preferred_locale property
	preferred_locale *string
	// The preferred_name property
	preferred_name *string
	// The prefix property
	prefix *string
	// The relationships property
	relationships []ContactsPutResponse_relationshipsable
	// The ScoreValue property
	scoreValue *string
	// The social_accounts property
	social_accounts []ContactsPutResponse_social_accountsable
	// The spouse_name property
	spouse_name *string
	// The suffix property
	suffix *string
	// The tag_ids property
	tag_ids []int64
	// The time_zone property
	time_zone *string
	// The website property
	website *string
}

// NewContactsPutResponse instantiates a new ContactsPutResponse and sets the default values.
func NewContactsPutResponse() *ContactsPutResponse {
	m := &ContactsPutResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateContactsPutResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateContactsPutResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewContactsPutResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ContactsPutResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAddresses gets the addresses property value. The addresses property
// returns a []ContactsPutResponse_addressesable when successful
func (m *ContactsPutResponse) GetAddresses() []ContactsPutResponse_addressesable {
	return m.addresses
}

// GetAnniversary gets the anniversary property value. The anniversary property
// returns a *Time when successful
func (m *ContactsPutResponse) GetAnniversary() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.anniversary
}

// GetBirthday gets the birthday property value. The birthday property
// returns a *Time when successful
func (m *ContactsPutResponse) GetBirthday() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.birthday
}

// GetCompany gets the company property value. The company property
// returns a ContactsPutResponse_companyable when successful
func (m *ContactsPutResponse) GetCompany() ContactsPutResponse_companyable {
	return m.company
}

// GetCompanyName gets the company_name property value. The company_name property
// returns a *string when successful
func (m *ContactsPutResponse) GetCompanyName() *string {
	return m.company_name
}

// GetContactType gets the contact_type property value. The contact_type property
// returns a *string when successful
func (m *ContactsPutResponse) GetContactType() *string {
	return m.contact_type
}

// GetCustomFields gets the custom_fields property value. The custom_fields property
// returns a []ContactsPutResponse_custom_fieldsable when successful
func (m *ContactsPutResponse) GetCustomFields() []ContactsPutResponse_custom_fieldsable {
	return m.custom_fields
}

// GetDateCreated gets the date_created property value. The date_created property
// returns a *Time when successful
func (m *ContactsPutResponse) GetDateCreated() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.date_created
}

// GetEmailAddresses gets the email_addresses property value. The email_addresses property
// returns a []ContactsPutResponse_email_addressesable when successful
func (m *ContactsPutResponse) GetEmailAddresses() []ContactsPutResponse_email_addressesable {
	return m.email_addresses
}

// GetEmailOptedIn gets the email_opted_in property value. The email_opted_in property
// returns a *bool when successful
func (m *ContactsPutResponse) GetEmailOptedIn() *bool {
	return m.email_opted_in
}

// GetFamilyName gets the family_name property value. The family_name property
// returns a *string when successful
func (m *ContactsPutResponse) GetFamilyName() *string {
	return m.family_name
}

// GetFaxNumbers gets the fax_numbers property value. The fax_numbers property
// returns a []ContactsPutResponse_fax_numbersable when successful
func (m *ContactsPutResponse) GetFaxNumbers() []ContactsPutResponse_fax_numbersable {
	return m.fax_numbers
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ContactsPutResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["addresses"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateContactsPutResponse_addressesFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ContactsPutResponse_addressesable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ContactsPutResponse_addressesable)
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
		val, err := n.GetObjectValue(CreateContactsPutResponse_companyFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCompany(val.(ContactsPutResponse_companyable))
		}
		return nil
	}
	res["company_name"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCompanyName(val)
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
		val, err := n.GetCollectionOfObjectValues(CreateContactsPutResponse_custom_fieldsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ContactsPutResponse_custom_fieldsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ContactsPutResponse_custom_fieldsable)
				}
			}
			m.SetCustomFields(res)
		}
		return nil
	}
	res["date_created"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetTimeValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetDateCreated(val)
		}
		return nil
	}
	res["email_addresses"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateContactsPutResponse_email_addressesFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ContactsPutResponse_email_addressesable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ContactsPutResponse_email_addressesable)
				}
			}
			m.SetEmailAddresses(res)
		}
		return nil
	}
	res["email_opted_in"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetEmailOptedIn(val)
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
		val, err := n.GetCollectionOfObjectValues(CreateContactsPutResponse_fax_numbersFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ContactsPutResponse_fax_numbersable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ContactsPutResponse_fax_numbersable)
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
		val, err := n.GetInt64Value()
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
	res["last_updated"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetTimeValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLastUpdated(val)
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
		val, err := n.GetObjectValue(CreateContactsPutResponse_originFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetOrigin(val.(ContactsPutResponse_originable))
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
		val, err := n.GetCollectionOfObjectValues(CreateContactsPutResponse_phone_numbersFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ContactsPutResponse_phone_numbersable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ContactsPutResponse_phone_numbersable)
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
	res["relationships"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateContactsPutResponse_relationshipsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ContactsPutResponse_relationshipsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ContactsPutResponse_relationshipsable)
				}
			}
			m.SetRelationships(res)
		}
		return nil
	}
	res["ScoreValue"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
		val, err := n.GetCollectionOfObjectValues(CreateContactsPutResponse_social_accountsFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]ContactsPutResponse_social_accountsable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(ContactsPutResponse_social_accountsable)
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
		val, err := n.GetCollectionOfPrimitiveValues("int64")
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]int64, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = *(v.(*int64))
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
func (m *ContactsPutResponse) GetGivenName() *string {
	return m.given_name
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *ContactsPutResponse) GetId() *int64 {
	return m.id
}

// GetJobTitle gets the job_title property value. The job_title property
// returns a *string when successful
func (m *ContactsPutResponse) GetJobTitle() *string {
	return m.job_title
}

// GetLastUpdated gets the last_updated property value. The last_updated property
// returns a *Time when successful
func (m *ContactsPutResponse) GetLastUpdated() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.last_updated
}

// GetLeadSourceId gets the lead_source_id property value. The lead_source_id property
// returns a *int64 when successful
func (m *ContactsPutResponse) GetLeadSourceId() *int64 {
	return m.lead_source_id
}

// GetMiddleName gets the middle_name property value. The middle_name property
// returns a *string when successful
func (m *ContactsPutResponse) GetMiddleName() *string {
	return m.middle_name
}

// GetOptInReason gets the opt_in_reason property value. The opt_in_reason property
// returns a *string when successful
func (m *ContactsPutResponse) GetOptInReason() *string {
	return m.opt_in_reason
}

// GetOrigin gets the origin property value. The origin property
// returns a ContactsPutResponse_originable when successful
func (m *ContactsPutResponse) GetOrigin() ContactsPutResponse_originable {
	return m.origin
}

// GetOwnerId gets the owner_id property value. The owner_id property
// returns a *int64 when successful
func (m *ContactsPutResponse) GetOwnerId() *int64 {
	return m.owner_id
}

// GetPhoneNumbers gets the phone_numbers property value. The phone_numbers property
// returns a []ContactsPutResponse_phone_numbersable when successful
func (m *ContactsPutResponse) GetPhoneNumbers() []ContactsPutResponse_phone_numbersable {
	return m.phone_numbers
}

// GetPreferredLocale gets the preferred_locale property value. The preferred_locale property
// returns a *string when successful
func (m *ContactsPutResponse) GetPreferredLocale() *string {
	return m.preferred_locale
}

// GetPreferredName gets the preferred_name property value. The preferred_name property
// returns a *string when successful
func (m *ContactsPutResponse) GetPreferredName() *string {
	return m.preferred_name
}

// GetPrefix gets the prefix property value. The prefix property
// returns a *string when successful
func (m *ContactsPutResponse) GetPrefix() *string {
	return m.prefix
}

// GetRelationships gets the relationships property value. The relationships property
// returns a []ContactsPutResponse_relationshipsable when successful
func (m *ContactsPutResponse) GetRelationships() []ContactsPutResponse_relationshipsable {
	return m.relationships
}

// GetScoreValue gets the ScoreValue property value. The ScoreValue property
// returns a *string when successful
func (m *ContactsPutResponse) GetScoreValue() *string {
	return m.scoreValue
}

// GetSocialAccounts gets the social_accounts property value. The social_accounts property
// returns a []ContactsPutResponse_social_accountsable when successful
func (m *ContactsPutResponse) GetSocialAccounts() []ContactsPutResponse_social_accountsable {
	return m.social_accounts
}

// GetSpouseName gets the spouse_name property value. The spouse_name property
// returns a *string when successful
func (m *ContactsPutResponse) GetSpouseName() *string {
	return m.spouse_name
}

// GetSuffix gets the suffix property value. The suffix property
// returns a *string when successful
func (m *ContactsPutResponse) GetSuffix() *string {
	return m.suffix
}

// GetTagIds gets the tag_ids property value. The tag_ids property
// returns a []int64 when successful
func (m *ContactsPutResponse) GetTagIds() []int64 {
	return m.tag_ids
}

// GetTimeZone gets the time_zone property value. The time_zone property
// returns a *string when successful
func (m *ContactsPutResponse) GetTimeZone() *string {
	return m.time_zone
}

// GetWebsite gets the website property value. The website property
// returns a *string when successful
func (m *ContactsPutResponse) GetWebsite() *string {
	return m.website
}

// Serialize serializes information the current object
func (m *ContactsPutResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
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
		err := writer.WriteStringValue("company_name", m.GetCompanyName())
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
	{
		err := writer.WriteTimeValue("date_created", m.GetDateCreated())
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
		err := writer.WriteBoolValue("email_opted_in", m.GetEmailOptedIn())
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
		err := writer.WriteInt64Value("id", m.GetId())
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
		err := writer.WriteTimeValue("last_updated", m.GetLastUpdated())
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
	if m.GetRelationships() != nil {
		cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRelationships()))
		for i, v := range m.GetRelationships() {
			if v != nil {
				cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
			}
		}
		err := writer.WriteCollectionOfObjectValues("relationships", cast)
		if err != nil {
			return err
		}
	}
	{
		err := writer.WriteStringValue("ScoreValue", m.GetScoreValue())
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
		err := writer.WriteCollectionOfInt64Values("tag_ids", m.GetTagIds())
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
func (m *ContactsPutResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAddresses sets the addresses property value. The addresses property
func (m *ContactsPutResponse) SetAddresses(value []ContactsPutResponse_addressesable) {
	m.addresses = value
}

// SetAnniversary sets the anniversary property value. The anniversary property
func (m *ContactsPutResponse) SetAnniversary(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	m.anniversary = value
}

// SetBirthday sets the birthday property value. The birthday property
func (m *ContactsPutResponse) SetBirthday(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	m.birthday = value
}

// SetCompany sets the company property value. The company property
func (m *ContactsPutResponse) SetCompany(value ContactsPutResponse_companyable) {
	m.company = value
}

// SetCompanyName sets the company_name property value. The company_name property
func (m *ContactsPutResponse) SetCompanyName(value *string) {
	m.company_name = value
}

// SetContactType sets the contact_type property value. The contact_type property
func (m *ContactsPutResponse) SetContactType(value *string) {
	m.contact_type = value
}

// SetCustomFields sets the custom_fields property value. The custom_fields property
func (m *ContactsPutResponse) SetCustomFields(value []ContactsPutResponse_custom_fieldsable) {
	m.custom_fields = value
}

// SetDateCreated sets the date_created property value. The date_created property
func (m *ContactsPutResponse) SetDateCreated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	m.date_created = value
}

// SetEmailAddresses sets the email_addresses property value. The email_addresses property
func (m *ContactsPutResponse) SetEmailAddresses(value []ContactsPutResponse_email_addressesable) {
	m.email_addresses = value
}

// SetEmailOptedIn sets the email_opted_in property value. The email_opted_in property
func (m *ContactsPutResponse) SetEmailOptedIn(value *bool) {
	m.email_opted_in = value
}

// SetFamilyName sets the family_name property value. The family_name property
func (m *ContactsPutResponse) SetFamilyName(value *string) {
	m.family_name = value
}

// SetFaxNumbers sets the fax_numbers property value. The fax_numbers property
func (m *ContactsPutResponse) SetFaxNumbers(value []ContactsPutResponse_fax_numbersable) {
	m.fax_numbers = value
}

// SetGivenName sets the given_name property value. The given_name property
func (m *ContactsPutResponse) SetGivenName(value *string) {
	m.given_name = value
}

// SetId sets the id property value. The id property
func (m *ContactsPutResponse) SetId(value *int64) {
	m.id = value
}

// SetJobTitle sets the job_title property value. The job_title property
func (m *ContactsPutResponse) SetJobTitle(value *string) {
	m.job_title = value
}

// SetLastUpdated sets the last_updated property value. The last_updated property
func (m *ContactsPutResponse) SetLastUpdated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	m.last_updated = value
}

// SetLeadSourceId sets the lead_source_id property value. The lead_source_id property
func (m *ContactsPutResponse) SetLeadSourceId(value *int64) {
	m.lead_source_id = value
}

// SetMiddleName sets the middle_name property value. The middle_name property
func (m *ContactsPutResponse) SetMiddleName(value *string) {
	m.middle_name = value
}

// SetOptInReason sets the opt_in_reason property value. The opt_in_reason property
func (m *ContactsPutResponse) SetOptInReason(value *string) {
	m.opt_in_reason = value
}

// SetOrigin sets the origin property value. The origin property
func (m *ContactsPutResponse) SetOrigin(value ContactsPutResponse_originable) {
	m.origin = value
}

// SetOwnerId sets the owner_id property value. The owner_id property
func (m *ContactsPutResponse) SetOwnerId(value *int64) {
	m.owner_id = value
}

// SetPhoneNumbers sets the phone_numbers property value. The phone_numbers property
func (m *ContactsPutResponse) SetPhoneNumbers(value []ContactsPutResponse_phone_numbersable) {
	m.phone_numbers = value
}

// SetPreferredLocale sets the preferred_locale property value. The preferred_locale property
func (m *ContactsPutResponse) SetPreferredLocale(value *string) {
	m.preferred_locale = value
}

// SetPreferredName sets the preferred_name property value. The preferred_name property
func (m *ContactsPutResponse) SetPreferredName(value *string) {
	m.preferred_name = value
}

// SetPrefix sets the prefix property value. The prefix property
func (m *ContactsPutResponse) SetPrefix(value *string) {
	m.prefix = value
}

// SetRelationships sets the relationships property value. The relationships property
func (m *ContactsPutResponse) SetRelationships(value []ContactsPutResponse_relationshipsable) {
	m.relationships = value
}

// SetScoreValue sets the ScoreValue property value. The ScoreValue property
func (m *ContactsPutResponse) SetScoreValue(value *string) {
	m.scoreValue = value
}

// SetSocialAccounts sets the social_accounts property value. The social_accounts property
func (m *ContactsPutResponse) SetSocialAccounts(value []ContactsPutResponse_social_accountsable) {
	m.social_accounts = value
}

// SetSpouseName sets the spouse_name property value. The spouse_name property
func (m *ContactsPutResponse) SetSpouseName(value *string) {
	m.spouse_name = value
}

// SetSuffix sets the suffix property value. The suffix property
func (m *ContactsPutResponse) SetSuffix(value *string) {
	m.suffix = value
}

// SetTagIds sets the tag_ids property value. The tag_ids property
func (m *ContactsPutResponse) SetTagIds(value []int64) {
	m.tag_ids = value
}

// SetTimeZone sets the time_zone property value. The time_zone property
func (m *ContactsPutResponse) SetTimeZone(value *string) {
	m.time_zone = value
}

// SetWebsite sets the website property value. The website property
func (m *ContactsPutResponse) SetWebsite(value *string) {
	m.website = value
}

type ContactsPutResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAddresses() []ContactsPutResponse_addressesable
	GetAnniversary() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetBirthday() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetCompany() ContactsPutResponse_companyable
	GetCompanyName() *string
	GetContactType() *string
	GetCustomFields() []ContactsPutResponse_custom_fieldsable
	GetDateCreated() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetEmailAddresses() []ContactsPutResponse_email_addressesable
	GetEmailOptedIn() *bool
	GetFamilyName() *string
	GetFaxNumbers() []ContactsPutResponse_fax_numbersable
	GetGivenName() *string
	GetId() *int64
	GetJobTitle() *string
	GetLastUpdated() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetLeadSourceId() *int64
	GetMiddleName() *string
	GetOptInReason() *string
	GetOrigin() ContactsPutResponse_originable
	GetOwnerId() *int64
	GetPhoneNumbers() []ContactsPutResponse_phone_numbersable
	GetPreferredLocale() *string
	GetPreferredName() *string
	GetPrefix() *string
	GetRelationships() []ContactsPutResponse_relationshipsable
	GetScoreValue() *string
	GetSocialAccounts() []ContactsPutResponse_social_accountsable
	GetSpouseName() *string
	GetSuffix() *string
	GetTagIds() []int64
	GetTimeZone() *string
	GetWebsite() *string
	SetAddresses(value []ContactsPutResponse_addressesable)
	SetAnniversary(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	SetBirthday(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	SetCompany(value ContactsPutResponse_companyable)
	SetCompanyName(value *string)
	SetContactType(value *string)
	SetCustomFields(value []ContactsPutResponse_custom_fieldsable)
	SetDateCreated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	SetEmailAddresses(value []ContactsPutResponse_email_addressesable)
	SetEmailOptedIn(value *bool)
	SetFamilyName(value *string)
	SetFaxNumbers(value []ContactsPutResponse_fax_numbersable)
	SetGivenName(value *string)
	SetId(value *int64)
	SetJobTitle(value *string)
	SetLastUpdated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	SetLeadSourceId(value *int64)
	SetMiddleName(value *string)
	SetOptInReason(value *string)
	SetOrigin(value ContactsPutResponse_originable)
	SetOwnerId(value *int64)
	SetPhoneNumbers(value []ContactsPutResponse_phone_numbersable)
	SetPreferredLocale(value *string)
	SetPreferredName(value *string)
	SetPrefix(value *string)
	SetRelationships(value []ContactsPutResponse_relationshipsable)
	SetScoreValue(value *string)
	SetSocialAccounts(value []ContactsPutResponse_social_accountsable)
	SetSpouseName(value *string)
	SetSuffix(value *string)
	SetTagIds(value []int64)
	SetTimeZone(value *string)
	SetWebsite(value *string)
}
