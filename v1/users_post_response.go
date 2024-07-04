package v1

import (
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
	i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
)

type UsersPostResponse struct {
	// Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
	additionalData map[string]any
	// The address property
	address UsersPostResponse_addressable
	// The company_name property
	company_name *string
	// The created_by property
	created_by *int64
	// The date_created property
	date_created *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	// The email_address property
	email_address *string
	// The family_name property
	family_name *string
	// The fax_numbers property
	fax_numbers []UsersPostResponse_fax_numbersable
	// The given_name property
	given_name *string
	// The global_user_id property
	global_user_id *int64
	// The id property
	id *int64
	// The infusionsoft_id property
	infusionsoft_id *string
	// The job_title property
	job_title *string
	// The last_updated property
	last_updated *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	// The last_updated_by property
	last_updated_by *int64
	// The middle_name property
	middle_name *string
	// The partner property
	partner *bool
	// The phone_numbers property
	phone_numbers []UsersPostResponse_phone_numbersable
	// The preferred_name property
	preferred_name *string
	// The time_zone property
	time_zone *string
	// The website property
	website *string
}

// NewUsersPostResponse instantiates a new UsersPostResponse and sets the default values.
func NewUsersPostResponse() *UsersPostResponse {
	m := &UsersPostResponse{}
	m.SetAdditionalData(make(map[string]any))
	return m
}

// CreateUsersPostResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateUsersPostResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) (i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
	return NewUsersPostResponse(), nil
}

// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *UsersPostResponse) GetAdditionalData() map[string]any {
	return m.additionalData
}

// GetAddress gets the address property value. The address property
// returns a UsersPostResponse_addressable when successful
func (m *UsersPostResponse) GetAddress() UsersPostResponse_addressable {
	return m.address
}

// GetCompanyName gets the company_name property value. The company_name property
// returns a *string when successful
func (m *UsersPostResponse) GetCompanyName() *string {
	return m.company_name
}

// GetCreatedBy gets the created_by property value. The created_by property
// returns a *int64 when successful
func (m *UsersPostResponse) GetCreatedBy() *int64 {
	return m.created_by
}

// GetDateCreated gets the date_created property value. The date_created property
// returns a *Time when successful
func (m *UsersPostResponse) GetDateCreated() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.date_created
}

// GetEmailAddress gets the email_address property value. The email_address property
// returns a *string when successful
func (m *UsersPostResponse) GetEmailAddress() *string {
	return m.email_address
}

// GetFamilyName gets the family_name property value. The family_name property
// returns a *string when successful
func (m *UsersPostResponse) GetFamilyName() *string {
	return m.family_name
}

// GetFaxNumbers gets the fax_numbers property value. The fax_numbers property
// returns a []UsersPostResponse_fax_numbersable when successful
func (m *UsersPostResponse) GetFaxNumbers() []UsersPostResponse_fax_numbersable {
	return m.fax_numbers
}

// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *UsersPostResponse) GetFieldDeserializers() map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
	res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error)
	res["address"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetObjectValue(CreateUsersPostResponse_addressFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			m.SetAddress(val.(UsersPostResponse_addressable))
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
	res["created_by"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetCreatedBy(val)
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
	res["email_address"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetStringValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetEmailAddress(val)
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
		val, err := n.GetCollectionOfObjectValues(CreateUsersPostResponse_fax_numbersFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]UsersPostResponse_fax_numbersable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(UsersPostResponse_fax_numbersable)
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
	res["last_updated_by"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetInt64Value()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetLastUpdatedBy(val)
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
	res["partner"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetBoolValue()
		if err != nil {
			return err
		}
		if val != nil {
			m.SetPartner(val)
		}
		return nil
	}
	res["phone_numbers"] = func(n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
		val, err := n.GetCollectionOfObjectValues(CreateUsersPostResponse_phone_numbersFromDiscriminatorValue)
		if err != nil {
			return err
		}
		if val != nil {
			res := make([]UsersPostResponse_phone_numbersable, len(val))
			for i, v := range val {
				if v != nil {
					res[i] = v.(UsersPostResponse_phone_numbersable)
				}
			}
			m.SetPhoneNumbers(res)
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
func (m *UsersPostResponse) GetGivenName() *string {
	return m.given_name
}

// GetGlobalUserId gets the global_user_id property value. The global_user_id property
// returns a *int64 when successful
func (m *UsersPostResponse) GetGlobalUserId() *int64 {
	return m.global_user_id
}

// GetId gets the id property value. The id property
// returns a *int64 when successful
func (m *UsersPostResponse) GetId() *int64 {
	return m.id
}

// GetInfusionsoftId gets the infusionsoft_id property value. The infusionsoft_id property
// returns a *string when successful
func (m *UsersPostResponse) GetInfusionsoftId() *string {
	return m.infusionsoft_id
}

// GetJobTitle gets the job_title property value. The job_title property
// returns a *string when successful
func (m *UsersPostResponse) GetJobTitle() *string {
	return m.job_title
}

// GetLastUpdated gets the last_updated property value. The last_updated property
// returns a *Time when successful
func (m *UsersPostResponse) GetLastUpdated() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time {
	return m.last_updated
}

// GetLastUpdatedBy gets the last_updated_by property value. The last_updated_by property
// returns a *int64 when successful
func (m *UsersPostResponse) GetLastUpdatedBy() *int64 {
	return m.last_updated_by
}

// GetMiddleName gets the middle_name property value. The middle_name property
// returns a *string when successful
func (m *UsersPostResponse) GetMiddleName() *string {
	return m.middle_name
}

// GetPartner gets the partner property value. The partner property
// returns a *bool when successful
func (m *UsersPostResponse) GetPartner() *bool {
	return m.partner
}

// GetPhoneNumbers gets the phone_numbers property value. The phone_numbers property
// returns a []UsersPostResponse_phone_numbersable when successful
func (m *UsersPostResponse) GetPhoneNumbers() []UsersPostResponse_phone_numbersable {
	return m.phone_numbers
}

// GetPreferredName gets the preferred_name property value. The preferred_name property
// returns a *string when successful
func (m *UsersPostResponse) GetPreferredName() *string {
	return m.preferred_name
}

// GetTimeZone gets the time_zone property value. The time_zone property
// returns a *string when successful
func (m *UsersPostResponse) GetTimeZone() *string {
	return m.time_zone
}

// GetWebsite gets the website property value. The website property
// returns a *string when successful
func (m *UsersPostResponse) GetWebsite() *string {
	return m.website
}

// Serialize serializes information the current object
func (m *UsersPostResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter) error {
	{
		err := writer.WriteObjectValue("address", m.GetAddress())
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
		err := writer.WriteInt64Value("created_by", m.GetCreatedBy())
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
	{
		err := writer.WriteStringValue("email_address", m.GetEmailAddress())
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
		err := writer.WriteInt64Value("global_user_id", m.GetGlobalUserId())
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
		err := writer.WriteStringValue("infusionsoft_id", m.GetInfusionsoftId())
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
		err := writer.WriteInt64Value("last_updated_by", m.GetLastUpdatedBy())
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
		err := writer.WriteBoolValue("partner", m.GetPartner())
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
		err := writer.WriteStringValue("preferred_name", m.GetPreferredName())
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
func (m *UsersPostResponse) SetAdditionalData(value map[string]any) {
	m.additionalData = value
}

// SetAddress sets the address property value. The address property
func (m *UsersPostResponse) SetAddress(value UsersPostResponse_addressable) {
	m.address = value
}

// SetCompanyName sets the company_name property value. The company_name property
func (m *UsersPostResponse) SetCompanyName(value *string) {
	m.company_name = value
}

// SetCreatedBy sets the created_by property value. The created_by property
func (m *UsersPostResponse) SetCreatedBy(value *int64) {
	m.created_by = value
}

// SetDateCreated sets the date_created property value. The date_created property
func (m *UsersPostResponse) SetDateCreated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	m.date_created = value
}

// SetEmailAddress sets the email_address property value. The email_address property
func (m *UsersPostResponse) SetEmailAddress(value *string) {
	m.email_address = value
}

// SetFamilyName sets the family_name property value. The family_name property
func (m *UsersPostResponse) SetFamilyName(value *string) {
	m.family_name = value
}

// SetFaxNumbers sets the fax_numbers property value. The fax_numbers property
func (m *UsersPostResponse) SetFaxNumbers(value []UsersPostResponse_fax_numbersable) {
	m.fax_numbers = value
}

// SetGivenName sets the given_name property value. The given_name property
func (m *UsersPostResponse) SetGivenName(value *string) {
	m.given_name = value
}

// SetGlobalUserId sets the global_user_id property value. The global_user_id property
func (m *UsersPostResponse) SetGlobalUserId(value *int64) {
	m.global_user_id = value
}

// SetId sets the id property value. The id property
func (m *UsersPostResponse) SetId(value *int64) {
	m.id = value
}

// SetInfusionsoftId sets the infusionsoft_id property value. The infusionsoft_id property
func (m *UsersPostResponse) SetInfusionsoftId(value *string) {
	m.infusionsoft_id = value
}

// SetJobTitle sets the job_title property value. The job_title property
func (m *UsersPostResponse) SetJobTitle(value *string) {
	m.job_title = value
}

// SetLastUpdated sets the last_updated property value. The last_updated property
func (m *UsersPostResponse) SetLastUpdated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
	m.last_updated = value
}

// SetLastUpdatedBy sets the last_updated_by property value. The last_updated_by property
func (m *UsersPostResponse) SetLastUpdatedBy(value *int64) {
	m.last_updated_by = value
}

// SetMiddleName sets the middle_name property value. The middle_name property
func (m *UsersPostResponse) SetMiddleName(value *string) {
	m.middle_name = value
}

// SetPartner sets the partner property value. The partner property
func (m *UsersPostResponse) SetPartner(value *bool) {
	m.partner = value
}

// SetPhoneNumbers sets the phone_numbers property value. The phone_numbers property
func (m *UsersPostResponse) SetPhoneNumbers(value []UsersPostResponse_phone_numbersable) {
	m.phone_numbers = value
}

// SetPreferredName sets the preferred_name property value. The preferred_name property
func (m *UsersPostResponse) SetPreferredName(value *string) {
	m.preferred_name = value
}

// SetTimeZone sets the time_zone property value. The time_zone property
func (m *UsersPostResponse) SetTimeZone(value *string) {
	m.time_zone = value
}

// SetWebsite sets the website property value. The website property
func (m *UsersPostResponse) SetWebsite(value *string) {
	m.website = value
}

type UsersPostResponseable interface {
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
	i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
	GetAddress() UsersPostResponse_addressable
	GetCompanyName() *string
	GetCreatedBy() *int64
	GetDateCreated() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetEmailAddress() *string
	GetFamilyName() *string
	GetFaxNumbers() []UsersPostResponse_fax_numbersable
	GetGivenName() *string
	GetGlobalUserId() *int64
	GetId() *int64
	GetInfusionsoftId() *string
	GetJobTitle() *string
	GetLastUpdated() *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
	GetLastUpdatedBy() *int64
	GetMiddleName() *string
	GetPartner() *bool
	GetPhoneNumbers() []UsersPostResponse_phone_numbersable
	GetPreferredName() *string
	GetTimeZone() *string
	GetWebsite() *string
	SetAddress(value UsersPostResponse_addressable)
	SetCompanyName(value *string)
	SetCreatedBy(value *int64)
	SetDateCreated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	SetEmailAddress(value *string)
	SetFamilyName(value *string)
	SetFaxNumbers(value []UsersPostResponse_fax_numbersable)
	SetGivenName(value *string)
	SetGlobalUserId(value *int64)
	SetId(value *int64)
	SetInfusionsoftId(value *string)
	SetJobTitle(value *string)
	SetLastUpdated(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
	SetLastUpdatedBy(value *int64)
	SetMiddleName(value *string)
	SetPartner(value *bool)
	SetPhoneNumbers(value []UsersPostResponse_phone_numbersable)
	SetPreferredName(value *string)
	SetTimeZone(value *string)
	SetWebsite(value *string)
}
